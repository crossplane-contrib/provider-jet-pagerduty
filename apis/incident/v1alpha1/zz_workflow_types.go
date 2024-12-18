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

type InlineStepsInputInitParameters struct {

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The steps in the workflow.
	Step []InlineStepsInputStepInitParameters `json:"step,omitempty" tf:"step,omitempty"`
}

type InlineStepsInputObservation struct {

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The steps in the workflow.
	Step []InlineStepsInputStepObservation `json:"step,omitempty" tf:"step,omitempty"`
}

type InlineStepsInputParameters struct {

	// The name of the workflow.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The steps in the workflow.
	// +kubebuilder:validation:Optional
	Step []InlineStepsInputStepParameters `json:"step,omitempty" tf:"step,omitempty"`
}

type InlineStepsInputStepInitParameters struct {

	// The action id for the workflow step, including the version. A list of actions available can be retrieved using the PagerDuty API.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The list of standard inputs for the workflow action.
	Input []InputInitParameters `json:"input,omitempty" tf:"input,omitempty"`

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type InlineStepsInputStepObservation struct {

	// The action id for the workflow step, including the version. A list of actions available can be retrieved using the PagerDuty API.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The list of standard inputs for the workflow action.
	Input []InputObservation `json:"input,omitempty" tf:"input,omitempty"`

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type InlineStepsInputStepParameters struct {

	// The action id for the workflow step, including the version. A list of actions available can be retrieved using the PagerDuty API.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// The list of standard inputs for the workflow action.
	// +kubebuilder:validation:Optional
	Input []InputParameters `json:"input,omitempty" tf:"input,omitempty"`

	// The name of the workflow.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type InputInitParameters struct {

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the input.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type InputObservation struct {
	Generated *bool `json:"generated,omitempty" tf:"generated,omitempty"`

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the input.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type InputParameters struct {

	// The name of the workflow.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of the input.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type StepInitParameters struct {

	// The action id for the workflow step, including the version. A list of actions available can be retrieved using the PagerDuty API.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The list of inputs that contain a series of inline steps for the workflow action.
	InlineStepsInput []InlineStepsInputInitParameters `json:"inlineStepsInput,omitempty" tf:"inline_steps_input,omitempty"`

	// The list of standard inputs for the workflow action.
	Input []StepInputInitParameters `json:"input,omitempty" tf:"input,omitempty"`

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StepInputInitParameters struct {

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the input.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StepInputObservation struct {
	Generated *bool `json:"generated,omitempty" tf:"generated,omitempty"`

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the input.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StepInputParameters struct {

	// The name of the workflow.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of the input.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type StepObservation struct {

	// The action id for the workflow step, including the version. A list of actions available can be retrieved using the PagerDuty API.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ID of the incident workflow.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of inputs that contain a series of inline steps for the workflow action.
	InlineStepsInput []InlineStepsInputObservation `json:"inlineStepsInput,omitempty" tf:"inline_steps_input,omitempty"`

	// The list of standard inputs for the workflow action.
	Input []StepInputObservation `json:"input,omitempty" tf:"input,omitempty"`

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StepParameters struct {

	// The action id for the workflow step, including the version. A list of actions available can be retrieved using the PagerDuty API.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// The list of inputs that contain a series of inline steps for the workflow action.
	// +kubebuilder:validation:Optional
	InlineStepsInput []InlineStepsInputParameters `json:"inlineStepsInput,omitempty" tf:"inline_steps_input,omitempty"`

	// The list of standard inputs for the workflow action.
	// +kubebuilder:validation:Optional
	Input []StepInputParameters `json:"input,omitempty" tf:"input,omitempty"`

	// The name of the workflow.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type WorkflowInitParameters struct {

	// The description of the workflow.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The steps in the workflow.
	Step []StepInitParameters `json:"step,omitempty" tf:"step,omitempty"`

	// A team ID. If specified then workflow edit permissions will be scoped to members of this team.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team
	// +crossplane:generate:reference:refFieldName=TeamRefs
	// +crossplane:generate:reference:selectorFieldName=TeamSelector
	Team *string `json:"team,omitempty" tf:"team,omitempty"`

	// Reference to a Team in team to populate team.
	// +kubebuilder:validation:Optional
	TeamRefs *v1.Reference `json:"teamRefs,omitempty" tf:"-"`

	// Selector for a Team in team to populate team.
	// +kubebuilder:validation:Optional
	TeamSelector *v1.Selector `json:"teamSelector,omitempty" tf:"-"`
}

type WorkflowObservation struct {

	// The description of the workflow.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the incident workflow.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the workflow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The steps in the workflow.
	Step []StepObservation `json:"step,omitempty" tf:"step,omitempty"`

	// A team ID. If specified then workflow edit permissions will be scoped to members of this team.
	Team *string `json:"team,omitempty" tf:"team,omitempty"`
}

type WorkflowParameters struct {

	// The description of the workflow.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the workflow.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The steps in the workflow.
	// +kubebuilder:validation:Optional
	Step []StepParameters `json:"step,omitempty" tf:"step,omitempty"`

	// A team ID. If specified then workflow edit permissions will be scoped to members of this team.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team
	// +crossplane:generate:reference:refFieldName=TeamRefs
	// +crossplane:generate:reference:selectorFieldName=TeamSelector
	// +kubebuilder:validation:Optional
	Team *string `json:"team,omitempty" tf:"team,omitempty"`

	// Reference to a Team in team to populate team.
	// +kubebuilder:validation:Optional
	TeamRefs *v1.Reference `json:"teamRefs,omitempty" tf:"-"`

	// Selector for a Team in team to populate team.
	// +kubebuilder:validation:Optional
	TeamSelector *v1.Selector `json:"teamSelector,omitempty" tf:"-"`
}

// WorkflowSpec defines the desired state of Workflow
type WorkflowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkflowParameters `json:"forProvider"`
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
	InitProvider WorkflowInitParameters `json:"initProvider,omitempty"`
}

// WorkflowStatus defines the observed state of Workflow.
type WorkflowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkflowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Workflow is the Schema for the Workflows API. Creates and manages an incident workflow in PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   WorkflowSpec   `json:"spec"`
	Status WorkflowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkflowList contains a list of Workflows
type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workflow `json:"items"`
}

// Repository type metadata.
var (
	Workflow_Kind             = "Workflow"
	Workflow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workflow_Kind}.String()
	Workflow_KindAPIVersion   = Workflow_Kind + "." + CRDGroupVersion.String()
	Workflow_GroupVersionKind = CRDGroupVersion.WithKind(Workflow_Kind)
)

func init() {
	SchemeBuilder.Register(&Workflow{}, &WorkflowList{})
}
