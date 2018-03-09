package server

import (
	"github.com/sanjid133/ksd/apis/sho"
	"github.com/sanjid133/ksd/apis/sho/install"
	"github.com/sanjid133/ksd/apis/sho/v1"
	"github.com/sanjid133/ksd/pkg/operator"
	ksdregistry "github.com/sanjid133/ksd/pkg/registry"
	ksdstorage "github.com/sanjid133/ksd/pkg/registry/secdb"
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"

	"fmt"
)

var (
	groupFactoryRegistry = make(announced.APIGroupFactoryRegistry)
	registry             = registered.NewOrDie("")

	Scheme = runtime.NewScheme()
	Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
	install.Install(groupFactoryRegistry, registry, Scheme)

	// we need to add the options to empty v1
	// TODO fix the server code to avoid this
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})

	// TODO: keep the generic API server from wanting this
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

type KsdConfig struct {
	GenericConfig    *genericapiserver.RecommendedConfig
	OperatorConfig operator.OperatorConfig
}

type KsdServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
	Operator *operator.Operator
}

type completedConfig struct {
	GenericConfig    genericapiserver.CompletedConfig
	OperatorConfig *operator.OperatorConfig
}

type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

func (c *KsdConfig) Complete() CompletedConfig {
	completedCfg := completedConfig{
		c.GenericConfig.Complete(),
		&c.OperatorConfig,
	}
	c.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}
	return CompletedConfig{&completedCfg}
}

func (c CompletedConfig) New() (*KsdServer, error) {
	fmt.Println("1.....................")
	genericServer, err := c.GenericConfig.New("ksd-apiserver", genericapiserver.EmptyDelegate)
	if err != nil {
		return nil, err
	}


	s := &KsdServer{
		GenericAPIServer: genericServer,
	}

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(sho.GroupName, registry, Scheme, metav1.ParameterCodec, Codecs)
	apiGroupInfo.GroupMeta.GroupVersion = v1.SchemeGroupVersion
	v1Storage := map[string]rest.Storage{}
	v1Storage["ksds"] = ksdregistry.RESTInPeace(ksdstorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter))
	apiGroupInfo.VersionedResourcesStorageMap["v1"] = v1Storage
	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}
	return s, nil
}
