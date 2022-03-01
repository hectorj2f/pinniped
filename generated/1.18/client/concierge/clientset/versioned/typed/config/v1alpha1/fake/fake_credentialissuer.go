// Copyright 2020-2022 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "go.pinniped.dev/generated/1.18/apis/concierge/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCredentialIssuers implements CredentialIssuerInterface
type FakeCredentialIssuers struct {
	Fake *FakeConfigV1alpha1
}

var credentialissuersResource = schema.GroupVersionResource{Group: "config.concierge.pinniped.dev", Version: "v1alpha1", Resource: "credentialissuers"}

var credentialissuersKind = schema.GroupVersionKind{Group: "config.concierge.pinniped.dev", Version: "v1alpha1", Kind: "CredentialIssuer"}

// Get takes name of the credentialIssuer, and returns the corresponding credentialIssuer object, and an error if there is any.
func (c *FakeCredentialIssuers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CredentialIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(credentialissuersResource, name), &v1alpha1.CredentialIssuer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuer), err
}

// List takes label and field selectors, and returns the list of CredentialIssuers that match those selectors.
func (c *FakeCredentialIssuers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CredentialIssuerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(credentialissuersResource, credentialissuersKind, opts), &v1alpha1.CredentialIssuerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CredentialIssuerList{ListMeta: obj.(*v1alpha1.CredentialIssuerList).ListMeta}
	for _, item := range obj.(*v1alpha1.CredentialIssuerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested credentialIssuers.
func (c *FakeCredentialIssuers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(credentialissuersResource, opts))
}

// Create takes the representation of a credentialIssuer and creates it.  Returns the server's representation of the credentialIssuer, and an error, if there is any.
func (c *FakeCredentialIssuers) Create(ctx context.Context, credentialIssuer *v1alpha1.CredentialIssuer, opts v1.CreateOptions) (result *v1alpha1.CredentialIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(credentialissuersResource, credentialIssuer), &v1alpha1.CredentialIssuer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuer), err
}

// Update takes the representation of a credentialIssuer and updates it. Returns the server's representation of the credentialIssuer, and an error, if there is any.
func (c *FakeCredentialIssuers) Update(ctx context.Context, credentialIssuer *v1alpha1.CredentialIssuer, opts v1.UpdateOptions) (result *v1alpha1.CredentialIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(credentialissuersResource, credentialIssuer), &v1alpha1.CredentialIssuer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCredentialIssuers) UpdateStatus(ctx context.Context, credentialIssuer *v1alpha1.CredentialIssuer, opts v1.UpdateOptions) (*v1alpha1.CredentialIssuer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(credentialissuersResource, "status", credentialIssuer), &v1alpha1.CredentialIssuer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuer), err
}

// Delete takes name of the credentialIssuer and deletes it. Returns an error if one occurs.
func (c *FakeCredentialIssuers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(credentialissuersResource, name), &v1alpha1.CredentialIssuer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCredentialIssuers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(credentialissuersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CredentialIssuerList{})
	return err
}

// Patch applies the patch and returns the patched credentialIssuer.
func (c *FakeCredentialIssuers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CredentialIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(credentialissuersResource, name, pt, data, subresources...), &v1alpha1.CredentialIssuer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CredentialIssuer), err
}
