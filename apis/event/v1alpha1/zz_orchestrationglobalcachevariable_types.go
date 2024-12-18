/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConfigurationInitParameters struct {

	// A RE2 regular expression that will be matched against the field specified via the source argument. This field is only used when type is recent_value
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// The path to the event field where the regex will be applied to extract a value. You can use any valid PCL path. This field is only used when type is recent_value
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The number of seconds indicating how long to count incoming trigger events for. This field is only used when type is trigger_event_count
	TTLSeconds *float64 `json:"ttlSeconds,omitempty" tf:"ttl_seconds,omitempty"`

	// The type of value to store into the Cache Variable. Can be one of: recent_value or trigger_event_count.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConfigurationObservation struct {

	// A RE2 regular expression that will be matched against the field specified via the source argument. This field is only used when type is recent_value
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// The path to the event field where the regex will be applied to extract a value. You can use any valid PCL path. This field is only used when type is recent_value
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The number of seconds indicating how long to count incoming trigger events for. This field is only used when type is trigger_event_count
	TTLSeconds *float64 `json:"ttlSeconds,omitempty" tf:"ttl_seconds,omitempty"`

	// The type of value to store into the Cache Variable. Can be one of: recent_value or trigger_event_count.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConfigurationParameters struct {

	// A RE2 regular expression that will be matched against the field specified via the source argument. This field is only used when type is recent_value
	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// The path to the event field where the regex will be applied to extract a value. You can use any valid PCL path. This field is only used when type is recent_value
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The number of seconds indicating how long to count incoming trigger events for. This field is only used when type is trigger_event_count
	// +kubebuilder:validation:Optional
	TTLSeconds *float64 `json:"ttlSeconds,omitempty" tf:"ttl_seconds,omitempty"`

	// The type of value to store into the Cache Variable. Can be one of: recent_value or trigger_event_count.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type OrchestrationGlobalCacheVariableConditionInitParameters struct {

	// A PCL condition string.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`
}

type OrchestrationGlobalCacheVariableConditionObservation struct {

	// A PCL condition string.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`
}

type OrchestrationGlobalCacheVariableConditionParameters struct {

	// A PCL condition string.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`
}

type OrchestrationGlobalCacheVariableInitParameters struct {

	// Conditions to be evaluated in order to determine whether or not to update the Cache Variable's stored value.
	Condition []OrchestrationGlobalCacheVariableConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// A configuration object to define what and how values will be stored in the Cache Variable.
	Configuration []ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Indicates whether the Cache Variable is disabled and would therefore not be evaluated.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// ID of the Global Event Orchestration to which this Cache Variable belongs.
	// +crossplane:generate:reference:type=Orchestration
	EventOrchestration *string `json:"eventOrchestration,omitempty" tf:"event_orchestration,omitempty"`

	// Reference to a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationRef *v1.Reference `json:"eventOrchestrationRef,omitempty" tf:"-"`

	// Selector for a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationSelector *v1.Selector `json:"eventOrchestrationSelector,omitempty" tf:"-"`

	// Name of the Cache Variable associated with the Global Event Orchestration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type OrchestrationGlobalCacheVariableObservation struct {

	// Conditions to be evaluated in order to determine whether or not to update the Cache Variable's stored value.
	Condition []OrchestrationGlobalCacheVariableConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// A configuration object to define what and how values will be stored in the Cache Variable.
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Indicates whether the Cache Variable is disabled and would therefore not be evaluated.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// ID of the Global Event Orchestration to which this Cache Variable belongs.
	EventOrchestration *string `json:"eventOrchestration,omitempty" tf:"event_orchestration,omitempty"`

	// ID of this Cache Variable.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Cache Variable associated with the Global Event Orchestration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type OrchestrationGlobalCacheVariableParameters struct {

	// Conditions to be evaluated in order to determine whether or not to update the Cache Variable's stored value.
	// +kubebuilder:validation:Optional
	Condition []OrchestrationGlobalCacheVariableConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// A configuration object to define what and how values will be stored in the Cache Variable.
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Indicates whether the Cache Variable is disabled and would therefore not be evaluated.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// ID of the Global Event Orchestration to which this Cache Variable belongs.
	// +crossplane:generate:reference:type=Orchestration
	// +kubebuilder:validation:Optional
	EventOrchestration *string `json:"eventOrchestration,omitempty" tf:"event_orchestration,omitempty"`

	// Reference to a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationRef *v1.Reference `json:"eventOrchestrationRef,omitempty" tf:"-"`

	// Selector for a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationSelector *v1.Selector `json:"eventOrchestrationSelector,omitempty" tf:"-"`

	// Name of the Cache Variable associated with the Global Event Orchestration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// OrchestrationGlobalCacheVariableSpec defines the desired state of OrchestrationGlobalCacheVariable
type OrchestrationGlobalCacheVariableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrchestrationGlobalCacheVariableParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider OrchestrationGlobalCacheVariableInitParameters `json:"initProvider,omitempty"`
}

// OrchestrationGlobalCacheVariableStatus defines the observed state of OrchestrationGlobalCacheVariable.
type OrchestrationGlobalCacheVariableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrchestrationGlobalCacheVariableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OrchestrationGlobalCacheVariable is the Schema for the OrchestrationGlobalCacheVariables API. Creates and manages a Cache Variable for a Global Event Orchestration.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type OrchestrationGlobalCacheVariable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configuration) || (has(self.initProvider) && has(self.initProvider.configuration))",message="spec.forProvider.configuration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   OrchestrationGlobalCacheVariableSpec   `json:"spec"`
	Status OrchestrationGlobalCacheVariableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrchestrationGlobalCacheVariableList contains a list of OrchestrationGlobalCacheVariables
type OrchestrationGlobalCacheVariableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrchestrationGlobalCacheVariable `json:"items"`
}

// Repository type metadata.
var (
	OrchestrationGlobalCacheVariable_Kind             = "OrchestrationGlobalCacheVariable"
	OrchestrationGlobalCacheVariable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrchestrationGlobalCacheVariable_Kind}.String()
	OrchestrationGlobalCacheVariable_KindAPIVersion   = OrchestrationGlobalCacheVariable_Kind + "." + CRDGroupVersion.String()
	OrchestrationGlobalCacheVariable_GroupVersionKind = CRDGroupVersion.WithKind(OrchestrationGlobalCacheVariable_Kind)
)

func init() {
	SchemeBuilder.Register(&OrchestrationGlobalCacheVariable{}, &OrchestrationGlobalCacheVariableList{})
}
