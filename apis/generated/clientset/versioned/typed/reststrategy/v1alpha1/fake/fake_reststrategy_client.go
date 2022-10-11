/*
Copyright DNITSCH

WTFPL
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/dnitsch/reststrategy/apis/generated/clientset/versioned/typed/reststrategy/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeReststrategyV1alpha1 struct {
	*testing.Fake
}

func (c *FakeReststrategyV1alpha1) RestStrategies(namespace string) v1alpha1.RestStrategyInterface {
	return &FakeRestStrategies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeReststrategyV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
