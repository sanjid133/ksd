package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Ksd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KsdSpec `json:"spec"`
}

// KsdSpec is the spec for a Ksd resource
type KsdSpec struct {
	SecretType string     `json:"secretType"`
	Info       []InfoSpec `json:"info"`
}

// InfoSpec is the item for a KsdSpec resource
type InfoSpec struct {
	SecretName string            `json:"secretName"`
	Data       map[string]string `json:"data"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KsdList is a list of Ksd resources
type KsdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is a list of ksds
	Items []Ksd `json:"items"`
}
