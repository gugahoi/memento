package registry

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Registry ...
type Registry struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   RegistrySpec
	Status RegistryStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RegistryList ...
type RegistryList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []Registry
}

// RegistrySpec ...
type RegistrySpec struct {
	// Message string
}

// RegistryStatus ...
type RegistryStatus struct {
	Sent bool
}
