package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Csye7374Spec defines the desired state of Csye7374
// +k8s:openapi-gen=true
type Csye7374Spec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Replicas           int    `json:"replicas"`
	ApplicationVersion string `json:"applicationVersion"`
}

// Csye7374Status defines the observed state of Csye7374
// +k8s:openapi-gen=true
type Csye7374Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Replicas           int    `json:"replicas"`
	ApplicationVersion string `json:"applicationVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Csye7374 is the Schema for the csye7374s API
// +k8s:openapi-gen=true
type Csye7374 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Csye7374Spec   `json:"spec,omitempty"`
	Status Csye7374Status `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Csye7374List contains a list of Csye7374
type Csye7374List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Csye7374 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Csye7374{}, &Csye7374List{})
}
