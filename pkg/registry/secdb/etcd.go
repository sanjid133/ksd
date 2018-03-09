package secdb

import (
	"github.com/sanjid133/ksd/apis/sho"
	"github.com/sanjid133/ksd/pkg/registry"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
)

func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &sho.Ksd{} },
		NewListFunc:              func() runtime.Object { return &sho.KsdList{} },
		PredicateFunc:            MatchSecDb,
		DefaultQualifiedResource: sho.Resource("ksd"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,
	}

	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}

	return &registry.REST{store}, nil

}
