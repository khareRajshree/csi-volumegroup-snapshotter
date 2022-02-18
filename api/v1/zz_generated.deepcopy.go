//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DellCsiVolumeGroupSnapshot) DeepCopyInto(out *DellCsiVolumeGroupSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DellCsiVolumeGroupSnapshot.
func (in *DellCsiVolumeGroupSnapshot) DeepCopy() *DellCsiVolumeGroupSnapshot {
	if in == nil {
		return nil
	}
	out := new(DellCsiVolumeGroupSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DellCsiVolumeGroupSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DellCsiVolumeGroupSnapshotList) DeepCopyInto(out *DellCsiVolumeGroupSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DellCsiVolumeGroupSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DellCsiVolumeGroupSnapshotList.
func (in *DellCsiVolumeGroupSnapshotList) DeepCopy() *DellCsiVolumeGroupSnapshotList {
	if in == nil {
		return nil
	}
	out := new(DellCsiVolumeGroupSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DellCsiVolumeGroupSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DellCsiVolumeGroupSnapshotSpec) DeepCopyInto(out *DellCsiVolumeGroupSnapshotSpec) {
	*out = *in
	if in.PvcList != nil {
		in, out := &in.PvcList, &out.PvcList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DellCsiVolumeGroupSnapshotSpec.
func (in *DellCsiVolumeGroupSnapshotSpec) DeepCopy() *DellCsiVolumeGroupSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(DellCsiVolumeGroupSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DellCsiVolumeGroupSnapshotStatus) DeepCopyInto(out *DellCsiVolumeGroupSnapshotStatus) {
	*out = *in
	in.CreationTime.DeepCopyInto(&out.CreationTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DellCsiVolumeGroupSnapshotStatus.
func (in *DellCsiVolumeGroupSnapshotStatus) DeepCopy() *DellCsiVolumeGroupSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(DellCsiVolumeGroupSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}
