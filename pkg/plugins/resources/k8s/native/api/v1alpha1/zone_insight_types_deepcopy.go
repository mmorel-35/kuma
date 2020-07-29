package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (cb *ZoneInsight) DeepCopyInto(out *ZoneInsight) {
	*out = *cb
	out.TypeMeta = cb.TypeMeta
	cb.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = runtime.DeepCopyJSON(cb.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPermission.
func (cb *ZoneInsight) DeepCopy() *ZoneInsight {
	if cb == nil {
		return nil
	}
	out := new(ZoneInsight)
	cb.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (cb *ZoneInsight) DeepCopyObject() runtime.Object {
	if c := cb.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (l *ZoneInsightList) DeepCopyInto(out *ZoneInsightList) {
	*out = *l
	out.TypeMeta = l.TypeMeta
	out.ListMeta = l.ListMeta
	if l.Items != nil {
		in, out := &l.Items, &out.Items
		*out = make([]ZoneInsight, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPermissionList.
func (l *ZoneInsightList) DeepCopy() *ZoneInsightList {
	if l == nil {
		return nil
	}
	out := new(ZoneInsightList)
	l.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (l *ZoneInsightList) DeepCopyObject() runtime.Object {
	if c := l.DeepCopy(); c != nil {
		return c
	}
	return nil
}
