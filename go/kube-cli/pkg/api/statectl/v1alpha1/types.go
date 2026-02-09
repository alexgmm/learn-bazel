package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceSpec   `json:"spec,omitempty"`
	Status InstanceStatus `json:"status,omitempty"`
}

type InstanceSpec struct {
	DesiredState string `json:"desiredState"`
}

type InstanceStatus struct {
	CurrentState string `json:"currentState"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Instance `json:"items"`
}

func (in *Instance) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}
	out := *in
	return &out
}

func (in *InstanceList) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}
	out := *in
	if in.Items != nil {
		out.Items = make([]Instance, len(in.Items))
		copy(out.Items, in.Items)
	}
	return &out
}
