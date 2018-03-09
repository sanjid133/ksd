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

package fake

import (
	sho_v1 "github.com/sanjid133/ksd/apis/sho/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKsds implements KsdInterface
type FakeKsds struct {
	Fake *FakeShoV1
	ns   string
}

var ksdsResource = schema.GroupVersionResource{Group: "sho.k8s.io", Version: "v1", Resource: "ksds"}

var ksdsKind = schema.GroupVersionKind{Group: "sho.k8s.io", Version: "v1", Kind: "Ksd"}

// Get takes name of the ksd, and returns the corresponding ksd object, and an error if there is any.
func (c *FakeKsds) Get(name string, options v1.GetOptions) (result *sho_v1.Ksd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ksdsResource, c.ns, name), &sho_v1.Ksd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*sho_v1.Ksd), err
}

// List takes label and field selectors, and returns the list of Ksds that match those selectors.
func (c *FakeKsds) List(opts v1.ListOptions) (result *sho_v1.KsdList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ksdsResource, ksdsKind, c.ns, opts), &sho_v1.KsdList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &sho_v1.KsdList{}
	for _, item := range obj.(*sho_v1.KsdList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ksds.
func (c *FakeKsds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ksdsResource, c.ns, opts))

}

// Create takes the representation of a ksd and creates it.  Returns the server's representation of the ksd, and an error, if there is any.
func (c *FakeKsds) Create(ksd *sho_v1.Ksd) (result *sho_v1.Ksd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ksdsResource, c.ns, ksd), &sho_v1.Ksd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*sho_v1.Ksd), err
}

// Update takes the representation of a ksd and updates it. Returns the server's representation of the ksd, and an error, if there is any.
func (c *FakeKsds) Update(ksd *sho_v1.Ksd) (result *sho_v1.Ksd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ksdsResource, c.ns, ksd), &sho_v1.Ksd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*sho_v1.Ksd), err
}

// Delete takes name of the ksd and deletes it. Returns an error if one occurs.
func (c *FakeKsds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ksdsResource, c.ns, name), &sho_v1.Ksd{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKsds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ksdsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &sho_v1.KsdList{})
	return err
}

// Patch applies the patch and returns the patched ksd.
func (c *FakeKsds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *sho_v1.Ksd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ksdsResource, c.ns, name, data, subresources...), &sho_v1.Ksd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*sho_v1.Ksd), err
}
