/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/suzerain-io/placeholder-name/kubernetes/1.19/client-go/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// LoginRequests returns a LoginRequestInformer.
	LoginRequests() LoginRequestInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// LoginRequests returns a LoginRequestInformer.
func (v *version) LoginRequests() LoginRequestInformer {
	return &loginRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
