/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/suzerain-io/placeholder-name/kubernetes/1.19/client-go/clientset/versioned/typed/placeholder/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePlaceholderV1alpha1 struct {
	*testing.Fake
}

func (c *FakePlaceholderV1alpha1) LoginRequests() v1alpha1.LoginRequestInterface {
	return &FakeLoginRequests{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePlaceholderV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
