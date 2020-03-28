package v1alpha1

import (
//	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// -------------------------------------------------------------------
// TASK
// -------------------------------------------------------------------

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=sirocco.cloud

// Task is a top-level type
type Pubkey struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Status PubkeyStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec PubkeySpec `json:"spec,omitempty"`
}

type PubkeyStatus struct {
	State         string `json:"state"`
}

type PubkeySpec struct {
	// +optional
	Address string `json:"address,omitempty"`
	// +optional
	Keytype string `json:"keytype,omitempty"`
	// +optional
	Active bool `json:"active,omitempty"`
	// +optional
	Key string `json:"key,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// no client needed for list as it's been created in above
type PubkeyList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []Pubkey `json:"items"`
}
