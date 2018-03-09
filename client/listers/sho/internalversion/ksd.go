/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package internalversion

import (
	sho "github.com/sanjid133/ksd/apis/sho"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KsdLister helps list Ksds.
type KsdLister interface {
	// List lists all Ksds in the indexer.
	List(selector labels.Selector) (ret []*sho.Ksd, err error)
	// Ksds returns an object that can list and get Ksds.
	Ksds(namespace string) KsdNamespaceLister
	KsdListerExpansion
}

// ksdLister implements the KsdLister interface.
type ksdLister struct {
	indexer cache.Indexer
}

// NewKsdLister returns a new KsdLister.
func NewKsdLister(indexer cache.Indexer) KsdLister {
	return &ksdLister{indexer: indexer}
}

// List lists all Ksds in the indexer.
func (s *ksdLister) List(selector labels.Selector) (ret []*sho.Ksd, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*sho.Ksd))
	})
	return ret, err
}

// Ksds returns an object that can list and get Ksds.
func (s *ksdLister) Ksds(namespace string) KsdNamespaceLister {
	return ksdNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KsdNamespaceLister helps list and get Ksds.
type KsdNamespaceLister interface {
	// List lists all Ksds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*sho.Ksd, err error)
	// Get retrieves the Ksd from the indexer for a given namespace and name.
	Get(name string) (*sho.Ksd, error)
	KsdNamespaceListerExpansion
}

// ksdNamespaceLister implements the KsdNamespaceLister
// interface.
type ksdNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ksds in the indexer for a given namespace.
func (s ksdNamespaceLister) List(selector labels.Selector) (ret []*sho.Ksd, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*sho.Ksd))
	})
	return ret, err
}

// Get retrieves the Ksd from the indexer for a given namespace and name.
func (s ksdNamespaceLister) Get(name string) (*sho.Ksd, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(sho.Resource("ksd"), name)
	}
	return obj.(*sho.Ksd), nil
}