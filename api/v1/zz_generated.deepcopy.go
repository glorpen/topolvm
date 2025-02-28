//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalVolume) DeepCopyInto(out *LogicalVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalVolume.
func (in *LogicalVolume) DeepCopy() *LogicalVolume {
	if in == nil {
		return nil
	}
	out := new(LogicalVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogicalVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalVolumeList) DeepCopyInto(out *LogicalVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LogicalVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalVolumeList.
func (in *LogicalVolumeList) DeepCopy() *LogicalVolumeList {
	if in == nil {
		return nil
	}
	out := new(LogicalVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogicalVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalVolumeSpec) DeepCopyInto(out *LogicalVolumeSpec) {
	*out = *in
	out.Size = in.Size.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalVolumeSpec.
func (in *LogicalVolumeSpec) DeepCopy() *LogicalVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(LogicalVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalVolumeStatus) DeepCopyInto(out *LogicalVolumeStatus) {
	*out = *in
	if in.CurrentSize != nil {
		in, out := &in.CurrentSize, &out.CurrentSize
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalVolumeStatus.
func (in *LogicalVolumeStatus) DeepCopy() *LogicalVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(LogicalVolumeStatus)
	in.DeepCopyInto(out)
	return out
}
