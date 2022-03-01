// Copyright 2020-2022 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "go.pinniped.dev/generated/1.21/apis/supervisor/idp/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OIDCIdentityProviderLister helps list OIDCIdentityProviders.
// All objects returned here must be treated as read-only.
type OIDCIdentityProviderLister interface {
	// List lists all OIDCIdentityProviders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OIDCIdentityProvider, err error)
	// OIDCIdentityProviders returns an object that can list and get OIDCIdentityProviders.
	OIDCIdentityProviders(namespace string) OIDCIdentityProviderNamespaceLister
	OIDCIdentityProviderListerExpansion
}

// oIDCIdentityProviderLister implements the OIDCIdentityProviderLister interface.
type oIDCIdentityProviderLister struct {
	indexer cache.Indexer
}

// NewOIDCIdentityProviderLister returns a new OIDCIdentityProviderLister.
func NewOIDCIdentityProviderLister(indexer cache.Indexer) OIDCIdentityProviderLister {
	return &oIDCIdentityProviderLister{indexer: indexer}
}

// List lists all OIDCIdentityProviders in the indexer.
func (s *oIDCIdentityProviderLister) List(selector labels.Selector) (ret []*v1alpha1.OIDCIdentityProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OIDCIdentityProvider))
	})
	return ret, err
}

// OIDCIdentityProviders returns an object that can list and get OIDCIdentityProviders.
func (s *oIDCIdentityProviderLister) OIDCIdentityProviders(namespace string) OIDCIdentityProviderNamespaceLister {
	return oIDCIdentityProviderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OIDCIdentityProviderNamespaceLister helps list and get OIDCIdentityProviders.
// All objects returned here must be treated as read-only.
type OIDCIdentityProviderNamespaceLister interface {
	// List lists all OIDCIdentityProviders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OIDCIdentityProvider, err error)
	// Get retrieves the OIDCIdentityProvider from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OIDCIdentityProvider, error)
	OIDCIdentityProviderNamespaceListerExpansion
}

// oIDCIdentityProviderNamespaceLister implements the OIDCIdentityProviderNamespaceLister
// interface.
type oIDCIdentityProviderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OIDCIdentityProviders in the indexer for a given namespace.
func (s oIDCIdentityProviderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OIDCIdentityProvider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OIDCIdentityProvider))
	})
	return ret, err
}

// Get retrieves the OIDCIdentityProvider from the indexer for a given namespace and name.
func (s oIDCIdentityProviderNamespaceLister) Get(name string) (*v1alpha1.OIDCIdentityProvider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("oidcidentityprovider"), name)
	}
	return obj.(*v1alpha1.OIDCIdentityProvider), nil
}
