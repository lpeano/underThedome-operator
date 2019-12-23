package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UnderthedomeSpec defines the desired state of Underthedome
// +k8s:openapi-gen=true
type UnderthedomeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	
	// Namespaces is list of namespaces to monitor
	Namespaces []string `json:"namespaces,omitempty"`

	// Repositories is list of valid repositories
	Repositories []string `json:"repositories,omitempty"`

	// Watchnamespace is namespace where crd can live
	Watchnamespace string `json:"watchnamespace,omitempty"`

}

// UnderthedomeStatus defines the observed state of Underthedome
// +k8s:openapi-gen=true
type UnderthedomeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Underthedome is the Schema for the underthedomes API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Underthedome struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UnderthedomeSpec   `json:"spec,omitempty"`
	Status UnderthedomeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UnderthedomeList contains a list of Underthedome
type UnderthedomeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Underthedome `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Underthedome{}, &UnderthedomeList{})
}
