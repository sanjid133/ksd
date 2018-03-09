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

package internalversion

import (
	sho "github.com/sanjid133/ksd/apis/sho"
	scheme "github.com/sanjid133/ksd/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KsdsGetter has a method to return a KsdInterface.
// A group's client should implement this interface.
type KsdsGetter interface {
	Ksds(namespace string) KsdInterface
}

// KsdInterface has methods to work with Ksd resources.
type KsdInterface interface {
	Create(*sho.Ksd) (*sho.Ksd, error)
	Update(*sho.Ksd) (*sho.Ksd, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*sho.Ksd, error)
	List(opts v1.ListOptions) (*sho.KsdList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *sho.Ksd, err error)
	KsdExpansion
}

// ksds implements KsdInterface
type ksds struct {
	client rest.Interface
	ns     string
}

// newKsds returns a Ksds
func newKsds(c *ShoClient, namespace string) *ksds {
	return &ksds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ksd, and returns the corresponding ksd object, and an error if there is any.
func (c *ksds) Get(name string, options v1.GetOptions) (result *sho.Ksd, err error) {
	result = &sho.Ksd{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ksds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Ksds that match those selectors.
func (c *ksds) List(opts v1.ListOptions) (result *sho.KsdList, err error) {
	result = &sho.KsdList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ksds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ksds.
func (c *ksds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ksds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a ksd and creates it.  Returns the server's representation of the ksd, and an error, if there is any.
func (c *ksds) Create(ksd *sho.Ksd) (result *sho.Ksd, err error) {
	result = &sho.Ksd{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ksds").
		Body(ksd).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ksd and updates it. Returns the server's representation of the ksd, and an error, if there is any.
func (c *ksds) Update(ksd *sho.Ksd) (result *sho.Ksd, err error) {
	result = &sho.Ksd{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ksds").
		Name(ksd.Name).
		Body(ksd).
		Do().
		Into(result)
	return
}

// Delete takes name of the ksd and deletes it. Returns an error if one occurs.
func (c *ksds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ksds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ksds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ksds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ksd.
func (c *ksds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *sho.Ksd, err error) {
	result = &sho.Ksd{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ksds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
