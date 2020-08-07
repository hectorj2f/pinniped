// +build !ignore_autogenerated

/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginDiscoveryConfig) DeepCopyInto(out *LoginDiscoveryConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginDiscoveryConfig.
func (in *LoginDiscoveryConfig) DeepCopy() *LoginDiscoveryConfig {
	if in == nil {
		return nil
	}
	out := new(LoginDiscoveryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoginDiscoveryConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginDiscoveryConfigList) DeepCopyInto(out *LoginDiscoveryConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoginDiscoveryConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginDiscoveryConfigList.
func (in *LoginDiscoveryConfigList) DeepCopy() *LoginDiscoveryConfigList {
	if in == nil {
		return nil
	}
	out := new(LoginDiscoveryConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoginDiscoveryConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginDiscoveryConfigSpec) DeepCopyInto(out *LoginDiscoveryConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginDiscoveryConfigSpec.
func (in *LoginDiscoveryConfigSpec) DeepCopy() *LoginDiscoveryConfigSpec {
	if in == nil {
		return nil
	}
	out := new(LoginDiscoveryConfigSpec)
	in.DeepCopyInto(out)
	return out
}
