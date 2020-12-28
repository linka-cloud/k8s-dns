// +build !ignore_autogenerated

/*
Copyright 2020 The Linka Cloud Team.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ARecord) DeepCopyInto(out *ARecord) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ARecord.
func (in *ARecord) DeepCopy() *ARecord {
	if in == nil {
		return nil
	}
	out := new(ARecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecord) DeepCopyInto(out *DNSRecord) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecord.
func (in *DNSRecord) DeepCopy() *DNSRecord {
	if in == nil {
		return nil
	}
	out := new(DNSRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSRecord) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordList) DeepCopyInto(out *DNSRecordList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordList.
func (in *DNSRecordList) DeepCopy() *DNSRecordList {
	if in == nil {
		return nil
	}
	out := new(DNSRecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSRecordList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordSpec) DeepCopyInto(out *DNSRecordSpec) {
	*out = *in
	if in.ARecord != nil {
		in, out := &in.ARecord, &out.ARecord
		*out = new(ARecord)
		**out = **in
	}
	if in.TXTRecord != nil {
		in, out := &in.TXTRecord, &out.TXTRecord
		*out = new(TXTRecord)
		(*in).DeepCopyInto(*out)
	}
	if in.SRVRecord != nil {
		in, out := &in.SRVRecord, &out.SRVRecord
		*out = new(SRVRecord)
		**out = **in
	}
	if in.MXRecord != nil {
		in, out := &in.MXRecord, &out.MXRecord
		*out = new(MXRecord)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordSpec.
func (in *DNSRecordSpec) DeepCopy() *DNSRecordSpec {
	if in == nil {
		return nil
	}
	out := new(DNSRecordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordStatus) DeepCopyInto(out *DNSRecordStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordStatus.
func (in *DNSRecordStatus) DeepCopy() *DNSRecordStatus {
	if in == nil {
		return nil
	}
	out := new(DNSRecordStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MXRecord) DeepCopyInto(out *MXRecord) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MXRecord.
func (in *MXRecord) DeepCopy() *MXRecord {
	if in == nil {
		return nil
	}
	out := new(MXRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SRVRecord) DeepCopyInto(out *SRVRecord) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SRVRecord.
func (in *SRVRecord) DeepCopy() *SRVRecord {
	if in == nil {
		return nil
	}
	out := new(SRVRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TXTRecord) DeepCopyInto(out *TXTRecord) {
	*out = *in
	if in.Txt != nil {
		in, out := &in.Txt, &out.Txt
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TXTRecord.
func (in *TXTRecord) DeepCopy() *TXTRecord {
	if in == nil {
		return nil
	}
	out := new(TXTRecord)
	in.DeepCopyInto(out)
	return out
}
