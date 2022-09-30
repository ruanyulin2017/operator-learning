package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TestDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestDeploymentSpec   `json:"spec"`
	Status TestDeploymentStatus `json:"status"`
}

// spec for the resource
type TestDeploymentSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas       int32  `json:"replicas"`
}

// status for the resource
type TestDeploymentStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

type TestDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TestDeployment `json:"items"`
}
