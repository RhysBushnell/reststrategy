/*
Copyright DNITSCH

WTFPL
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/dnitsch/reststrategy/apis/reststrategy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRestStrategies implements RestStrategyInterface
type FakeRestStrategies struct {
	Fake *FakeReststrategyV1alpha1
	ns   string
}

var reststrategiesResource = schema.GroupVersionResource{Group: "reststrategy.dnitsch.net", Version: "v1alpha1", Resource: "reststrategies"}

var reststrategiesKind = schema.GroupVersionKind{Group: "reststrategy.dnitsch.net", Version: "v1alpha1", Kind: "RestStrategy"}

// Get takes name of the restStrategy, and returns the corresponding restStrategy object, and an error if there is any.
func (c *FakeRestStrategies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RestStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(reststrategiesResource, c.ns, name), &v1alpha1.RestStrategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RestStrategy), err
}

// List takes label and field selectors, and returns the list of RestStrategies that match those selectors.
func (c *FakeRestStrategies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RestStrategyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(reststrategiesResource, reststrategiesKind, c.ns, opts), &v1alpha1.RestStrategyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RestStrategyList{ListMeta: obj.(*v1alpha1.RestStrategyList).ListMeta}
	for _, item := range obj.(*v1alpha1.RestStrategyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested restStrategies.
func (c *FakeRestStrategies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(reststrategiesResource, c.ns, opts))

}

// Create takes the representation of a restStrategy and creates it.  Returns the server's representation of the restStrategy, and an error, if there is any.
func (c *FakeRestStrategies) Create(ctx context.Context, restStrategy *v1alpha1.RestStrategy, opts v1.CreateOptions) (result *v1alpha1.RestStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(reststrategiesResource, c.ns, restStrategy), &v1alpha1.RestStrategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RestStrategy), err
}

// Update takes the representation of a restStrategy and updates it. Returns the server's representation of the restStrategy, and an error, if there is any.
func (c *FakeRestStrategies) Update(ctx context.Context, restStrategy *v1alpha1.RestStrategy, opts v1.UpdateOptions) (result *v1alpha1.RestStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(reststrategiesResource, c.ns, restStrategy), &v1alpha1.RestStrategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RestStrategy), err
}

// Delete takes name of the restStrategy and deletes it. Returns an error if one occurs.
func (c *FakeRestStrategies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(reststrategiesResource, c.ns, name, opts), &v1alpha1.RestStrategy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRestStrategies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(reststrategiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RestStrategyList{})
	return err
}

// Patch applies the patch and returns the patched restStrategy.
func (c *FakeRestStrategies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RestStrategy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reststrategiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RestStrategy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RestStrategy), err
}
