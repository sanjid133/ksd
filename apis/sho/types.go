package sho

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Ksd struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec KsdSpec
}

// KsdSpec is the spec for a Ksd resource
type KsdSpec struct {
	SecretType string
	Info       []InfoSpec
}

// InfoSpec is the item for a KsdSpec resource
type InfoSpec struct {
	SecretName string
	Data       map[string]string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// KsdList is a list of Ksd resources
type KsdList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Ksd
}
