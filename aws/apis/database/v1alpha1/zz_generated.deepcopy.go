// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstance) DeepCopyInto(out *RDSInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstance.
func (in *RDSInstance) DeepCopy() *RDSInstance {
	if in == nil {
		return nil
	}
	out := new(RDSInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceClass) DeepCopyInto(out *RDSInstanceClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceClass.
func (in *RDSInstanceClass) DeepCopy() *RDSInstanceClass {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSInstanceClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceClassList) DeepCopyInto(out *RDSInstanceClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RDSInstanceClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceClassList.
func (in *RDSInstanceClassList) DeepCopy() *RDSInstanceClassList {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSInstanceClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceClassSpecTemplate) DeepCopyInto(out *RDSInstanceClassSpecTemplate) {
	*out = *in
	in.ResourceClassSpecTemplate.DeepCopyInto(&out.ResourceClassSpecTemplate)
	in.RDSInstanceParameters.DeepCopyInto(&out.RDSInstanceParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceClassSpecTemplate.
func (in *RDSInstanceClassSpecTemplate) DeepCopy() *RDSInstanceClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceList) DeepCopyInto(out *RDSInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RDSInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceList.
func (in *RDSInstanceList) DeepCopy() *RDSInstanceList {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RDSInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceParameters) DeepCopyInto(out *RDSInstanceParameters) {
	*out = *in
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceParameters.
func (in *RDSInstanceParameters) DeepCopy() *RDSInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceSpec) DeepCopyInto(out *RDSInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.RDSInstanceParameters.DeepCopyInto(&out.RDSInstanceParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceSpec.
func (in *RDSInstanceSpec) DeepCopy() *RDSInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDSInstanceStatus) DeepCopyInto(out *RDSInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDSInstanceStatus.
func (in *RDSInstanceStatus) DeepCopy() *RDSInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(RDSInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
