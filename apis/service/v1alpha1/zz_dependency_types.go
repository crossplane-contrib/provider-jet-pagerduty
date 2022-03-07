/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DependencyDependencyObservation struct {
}

type DependencyDependencyParameters struct {

	// +kubebuilder:validation:Required
	DependentService []DependentServiceParameters `json:"dependentService" tf:"dependent_service,omitempty"`

	// +kubebuilder:validation:Required
	SupportingService []SupportingServiceParameters `json:"supportingService" tf:"supporting_service,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DependencyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DependencyParameters struct {

	// +kubebuilder:validation:Required
	Dependency []DependencyDependencyParameters `json:"dependency" tf:"dependency,omitempty"`
}

type DependentServiceObservation struct {
}

type DependentServiceParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SupportingServiceObservation struct {
}

type SupportingServiceParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// DependencySpec defines the desired state of Dependency
type DependencySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DependencyParameters `json:"forProvider"`
}

// DependencyStatus defines the observed state of Dependency.
type DependencyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DependencyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Dependency is the Schema for the Dependencys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerdutyjet}
type Dependency struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DependencySpec   `json:"spec"`
	Status            DependencyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DependencyList contains a list of Dependencys
type DependencyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dependency `json:"items"`
}

// Repository type metadata.
var (
	Dependency_Kind             = "Dependency"
	Dependency_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dependency_Kind}.String()
	Dependency_KindAPIVersion   = Dependency_Kind + "." + CRDGroupVersion.String()
	Dependency_GroupVersionKind = CRDGroupVersion.WithKind(Dependency_Kind)
)

func init() {
	SchemeBuilder.Register(&Dependency{}, &DependencyList{})
}
