// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerChannel) DeepCopyInto(out *BrokerChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerChannel.
func (in *BrokerChannel) DeepCopy() *BrokerChannel {
	if in == nil {
		return nil
	}
	out := new(BrokerChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BrokerChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerChannelList) DeepCopyInto(out *BrokerChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BrokerChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerChannelList.
func (in *BrokerChannelList) DeepCopy() *BrokerChannelList {
	if in == nil {
		return nil
	}
	out := new(BrokerChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BrokerChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerChannelSpec) DeepCopyInto(out *BrokerChannelSpec) {
	*out = *in
	in.SourceSpec.DeepCopyInto(&out.SourceSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerChannelSpec.
func (in *BrokerChannelSpec) DeepCopy() *BrokerChannelSpec {
	if in == nil {
		return nil
	}
	out := new(BrokerChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerChannelStatus) DeepCopyInto(out *BrokerChannelStatus) {
	*out = *in
	in.SourceStatus.DeepCopyInto(&out.SourceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerChannelStatus.
func (in *BrokerChannelStatus) DeepCopy() *BrokerChannelStatus {
	if in == nil {
		return nil
	}
	out := new(BrokerChannelStatus)
	in.DeepCopyInto(out)
	return out
}
