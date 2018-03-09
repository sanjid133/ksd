package registry

import (
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

type REST struct {
	*genericregistry.Store
}

func RESTInPeace(storage rest.StandardStorage, err error) rest.StandardStorage {
	if err != nil {
		panic(err)
	}
	return storage
}
