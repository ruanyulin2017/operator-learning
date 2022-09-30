package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *TestDeployment) DeepCopyInto(out *TestDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *TestDeploymentSpec) DeepCopyInto(out *TestDeploymentSpec) {
	*out = *in
	return
}

func (in *TestDeploymentStatus) DeepCopyInto(out *TestDeploymentStatus) {
	*out = *in
	return
}

func (in *TestDeployment) DeepCopy() *TestDeployment {
	if in == nil {
		return nil
	}
	out := new(TestDeployment)
	in.DeepCopyInto(out)
	return out
}

func (in *TestDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *TestDeploymentList) DeepCopyInto(out *TestDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TestDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *TestDeploymentList) DeepCopy() *TestDeploymentList {
	if in == nil {
		return nil
	}
	out := new(TestDeploymentList)
	return out
}

func (in *TestDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
