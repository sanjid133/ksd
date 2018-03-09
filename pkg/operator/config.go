package operator

import (
	clientset "github.com/sanjid133/ksd/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	informers "github.com/sanjid133/ksd/client/informers/externalversions"
	"k8s.io/client-go/tools/clientcmd"
	kubeinformers "k8s.io/client-go/informers"
	"time"
)

type Config struct {
	Kubeconfig string
	MasterUrl string
}

type OperatorConfig struct {
	Config
	ClientConfig *rest.Config
	KubeClient   kubernetes.Interface
	KsdClient    clientset.Interface
}

func NewOperatorConfig(kc, mu string) OperatorConfig  {
	return OperatorConfig{
		Config: Config{
			Kubeconfig: kc,
			MasterUrl: mu,
		},
	}
}
func (c *OperatorConfig) New() (*Operator, error) {
	cfg, err := clientcmd.BuildConfigFromFlags(c.MasterUrl, c.Kubeconfig)
	if err != nil {
		return nil, err
	}

	kubeclient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	demoClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	kif := kubeinformers.NewSharedInformerFactory(kubeclient, time.Second*30)
	dif := informers.NewSharedInformerFactory(demoClient, time.Second*30)
	ctrl := NewController(kubeclient, demoClient, kif, dif)

	stop := make(chan struct{})
	defer close(stop)
	go kif.Start(stop)
	go dif.Start(stop)
	return ctrl, nil
	//ctrl.Run(2, stop)

}
