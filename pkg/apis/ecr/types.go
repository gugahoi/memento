package ecr

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ECR ...
type ECR struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   ECRSpec
	Status ECRStatus
}

// ECRSpec ...
type ECRSpec struct {
	// Message string
}

// ECRStatus ...
type ECRStatus struct {
	ARN string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ECRList ...
type ECRList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []ECR
}
