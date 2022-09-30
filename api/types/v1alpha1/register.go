package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GroupName = "test.example.com"
var GroupVersion = "v1alpha1"
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: GroupVersion}

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&TestDeployment{},
		&TestDeploymentList{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)
