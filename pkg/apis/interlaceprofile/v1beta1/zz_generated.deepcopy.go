//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationConfig) DeepCopyInto(out *ApplicationConfig) {
	*out = *in
	out.Selector = in.Selector
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]ApplicationAction, len(*in))
		copy(*out, *in)
	}
	out.VerifyConfig = in.VerifyConfig
	in.SignConfig.DeepCopyInto(&out.SignConfig)
	out.ProvenanceConfig = in.ProvenanceConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationConfig.
func (in *ApplicationConfig) DeepCopy() *ApplicationConfig {
	if in == nil {
		return nil
	}
	out := new(ApplicationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplictionSelector) DeepCopyInto(out *ApplictionSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplictionSelector.
func (in *ApplictionSelector) DeepCopy() *ApplictionSelector {
	if in == nil {
		return nil
	}
	out := new(ApplictionSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterlaceProfile) DeepCopyInto(out *InterlaceProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterlaceProfile.
func (in *InterlaceProfile) DeepCopy() *InterlaceProfile {
	if in == nil {
		return nil
	}
	out := new(InterlaceProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InterlaceProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterlaceProfileList) DeepCopyInto(out *InterlaceProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InterlaceProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterlaceProfileList.
func (in *InterlaceProfileList) DeepCopy() *InterlaceProfileList {
	if in == nil {
		return nil
	}
	out := new(InterlaceProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InterlaceProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterlaceProfileSpec) DeepCopyInto(out *InterlaceProfileSpec) {
	*out = *in
	if in.Applications != nil {
		in, out := &in.Applications, &out.Applications
		*out = make([]ApplicationConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Protection.DeepCopyInto(&out.Protection)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterlaceProfileSpec.
func (in *InterlaceProfileSpec) DeepCopy() *InterlaceProfileSpec {
	if in == nil {
		return nil
	}
	out := new(InterlaceProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterlaceProfileStatus) DeepCopyInto(out *InterlaceProfileStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterlaceProfileStatus.
func (in *InterlaceProfileStatus) DeepCopy() *InterlaceProfileStatus {
	if in == nil {
		return nil
	}
	out := new(InterlaceProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyConfig) DeepCopyInto(out *KeyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyConfig.
func (in *KeyConfig) DeepCopy() *KeyConfig {
	if in == nil {
		return nil
	}
	out := new(KeyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionConfig) DeepCopyInto(out *ProtectionConfig) {
	*out = *in
	if in.Generators != nil {
		in, out := &in.Generators, &out.Generators
		*out = make([]Generator, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				(*out)[i] = (*in)[i].DeepCopyGenerator()
			}
		}
	}
	out.Destination = in.Destination
	out.PolicySource = in.PolicySource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionConfig.
func (in *ProtectionConfig) DeepCopy() *ProtectionConfig {
	if in == nil {
		return nil
	}
	out := new(ProtectionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectionPolicySourceConfig) DeepCopyInto(out *ProtectionPolicySourceConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectionPolicySourceConfig.
func (in *ProtectionPolicySourceConfig) DeepCopy() *ProtectionPolicySourceConfig {
	if in == nil {
		return nil
	}
	out := new(ProtectionPolicySourceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvenanceConfig) DeepCopyInto(out *ProvenanceConfig) {
	*out = *in
	out.KeyConfig = in.KeyConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvenanceConfig.
func (in *ProvenanceConfig) DeepCopy() *ProvenanceConfig {
	if in == nil {
		return nil
	}
	out := new(ProvenanceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMatchPattern) DeepCopyInto(out *ResourceMatchPattern) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMatchPattern.
func (in *ResourceMatchPattern) DeepCopy() *ResourceMatchPattern {
	if in == nil {
		return nil
	}
	out := new(ResourceMatchPattern)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignConfig) DeepCopyInto(out *SignConfig) {
	*out = *in
	out.KeyConfig = in.KeyConfig
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]ResourceMatchPattern, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignConfig.
func (in *SignConfig) DeepCopy() *SignConfig {
	if in == nil {
		return nil
	}
	out := new(SignConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerifyConfig) DeepCopyInto(out *VerifyConfig) {
	*out = *in
	out.KeyConfig = in.KeyConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerifyConfig.
func (in *VerifyConfig) DeepCopy() *VerifyConfig {
	if in == nil {
		return nil
	}
	out := new(VerifyConfig)
	in.DeepCopyInto(out)
	return out
}
