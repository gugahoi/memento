package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ECR ...
type ECR struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ECRSpec   `json:"spec,omitempty"`
	Status ECRStatus `json:"status,omitempty"`
}

// ECRSpec ...
type ECRSpec struct {
	Policy string `json:"type"`
}

// ECRStatus ...
type ECRStatus struct {
	ARN string `json:"arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ECRList ...
type ECRList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []ECR `json:"items"`
}
