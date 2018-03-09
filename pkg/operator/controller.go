package operator

import (
	clientset "github.com/sanjid133/ksd/client/clientset/versioned"
	informers "github.com/sanjid133/ksd/client/informers/externalversions"
	listers "github.com/sanjid133/ksd/client/listers/sho/v1"
	kubeinformers "k8s.io/client-go/informers"
	ksdscheme "github.com/sanjid133/ksd/client/clientset/versioned/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/client-go/kubernetes/scheme"
	"github.com/golang/glog"
	core "k8s.io/api/core/v1"
	"fmt"
	sho "github.com/sanjid133/ksd/apis/sho/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

type Operator struct {
	kubeClient kubernetes.Interface
	ksdClient  clientset.Interface

	secretsLister corelisters.SecretLister
	secretSynced  cache.InformerSynced

	secdbLister listers.KsdLister
	secdbSynced cache.InformerSynced

	queue workqueue.RateLimitingInterface

	recorder record.EventRecorder
}

const controllerAgentName = "ksd"

func NewController(
	kubeClient kubernetes.Interface,
	ksdClient clientset.Interface,
	kubeInformerFactory kubeinformers.SharedInformerFactory,
	ksdInformerFactory informers.SharedInformerFactory) *Operator {

	secretInformer := kubeInformerFactory.Core().V1().Secrets()
	secdbInformer := ksdInformerFactory.Sho().V1().Ksds()


	/*lw := &cache.ListWatch{
		ListFunc: func(opts metav1.ListOptions) (rt.Object, error) {
			return demoClient.DemocontrollerV1().SecDbs(core.NamespaceAll).List(metav1.ListOptions{})
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return demoClient.DemocontrollerV1().SecDbs(core.NamespaceAll).Watch(metav1.ListOptions{})
		},
	}*/

	ksdscheme.AddToScheme(scheme.Scheme)
	glog.V(4).Info("Creating event broadcaster")
	fmt.Println("creating................")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, core.EventSource{Component: controllerAgentName})

	controller := &Operator{
		kubeClient:    kubeClient,
		ksdClient:     ksdClient,
		secretsLister: secretInformer.Lister(),
		secretSynced:  secretInformer.Informer().HasSynced,
		secdbLister:   secdbInformer.Lister(),
		secdbSynced:   secdbInformer.Informer().HasSynced,
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Ksd"),
		recorder:      recorder,
	}
	fmt.Println("setting up event handler")

	secdbInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("adddddddddddddddddddd")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				controller.queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("updateeeeeeeeeeeeee")
			newFiDb := newObj.(*sho.Ksd)
			oldFiDb := oldObj.(*sho.Ksd)
			if newFiDb.ResourceVersion == oldFiDb.ResourceVersion {
				return
			}
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err == nil {
				controller.queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("deleteeeeeeeeeeeeeeee")
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				controller.queue.Add(key)
			}
		},
	})

	secretInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handleObject,
		UpdateFunc: func(old, new interface{}) {
			newSec := new.(*core.Secret)
			oldSec := old.(*core.Secret)
			if newSec.ResourceVersion == oldSec.ResourceVersion {
				// Periodic resync will send update events for all known Deployments.
				// Two different versions of the same Deployment will always have different RVs.
				return
			}
			controller.handleObject(new)
		},
		DeleteFunc: controller.handleObject,
	})
	return controller
}

func (c *Operator) Run(thrediness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	fmt.Println("Starting fidb controller")

	fmt.Println("waiting for informer caches to sync")
	fmt.Println("waiting 2")

	if !cache.WaitForCacheSync(stopCh, c.secretSynced, c.secdbSynced) {
		runtime.HandleError(fmt.Errorf("waiting timeout for caches sync"))
		return
	}
	

	fmt.Println("starting workers")
	for i := 0; i < thrediness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	fmt.Println("started workers")
	<-stopCh
	fmt.Println("shutting down workers")
	fmt.Println("Stopping fidb controller")
}

func (c *Operator) runWorker() {
	for c.processNextItem() {
	}
}

func (c *Operator) processNextItem() bool {
	key, shutdown := c.queue.Get()
	if shutdown {
		return false
	}

	defer c.queue.Done(key)

	var k string
	var ok bool
	if k, ok = key.(string); !ok {
		c.queue.Forget(key)
		runtime.HandleError(fmt.Errorf("expected string but found %v", key))
		return true
	}

	err := c.demoSyncHandler(k)
	if err == nil {
		//fmt.Println("no error, object is proccess so removing from queue")
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < 5 {
		fmt.Println("error on processing. will retry with error ", err)
		c.queue.AddRateLimited(key)
	} else {
		fmt.Println("error processing. give up with error ", err)
		c.queue.Forget(key)
		runtime.HandleError(err)
	}

	return true

}

func (c *Operator) handleObject(obj interface{}) {
	var object metav1.Object
	var ok bool

	if object, ok = obj.(metav1.Object); !ok {
		dobj, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			runtime.HandleError(fmt.Errorf("error decoding object, invalaid type"))
			return
		}
		object, ok = dobj.Obj.(metav1.Object)
		if !ok {
			runtime.HandleError(fmt.Errorf("error decoding cache object, invalid type"))
			return
		}
		fmt.Println("recovered delete object ", object.GetName())
	}
	fmt.Println("processing object..........")
	if ownerRef := metav1.GetControllerOf(object); ownerRef != nil {
		fmt.Println("controller kind = ", ownerRef.Kind)
		if ownerRef.Kind != sho.KindKsd || object.GetDeletionTimestamp() != nil{
			return
		}

		fidb, err := c.secdbLister.Ksds(object.GetNamespace()).Get(ownerRef.Name)
		if err != nil {
			fmt.Println(fmt.Sprintf("orphand object %v, err %v", object.GetSelfLink(), err))
			return
		}
		c.enqueueFoo(fidb)
		return

	}
	fmt.Println("nothing found here........")
}

func (c *Operator) enqueueFoo(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}
	c.queue.AddRateLimited(key)
}
