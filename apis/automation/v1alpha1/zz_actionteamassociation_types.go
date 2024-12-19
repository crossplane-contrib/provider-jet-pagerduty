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

type ActionTeamAssociationInitParameters struct {

	// Id of the action.
	// +crossplane:generate:reference:type=Action
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// Reference to a Action to populate actionId.
	// +kubebuilder:validation:Optional
	ActionIDRef *v1.Reference `json:"actionIdRef,omitempty" tf:"-"`

	// Selector for a Action to populate actionId.
	// +kubebuilder:validation:Optional
	ActionIDSelector *v1.Selector `json:"actionIdSelector,omitempty" tf:"-"`

	// Id of the team associated to the action.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDRef *v1.Reference `json:"teamIdRef,omitempty" tf:"-"`

	// Selector for a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDSelector *v1.Selector `json:"teamIdSelector,omitempty" tf:"-"`
}

type ActionTeamAssociationObservation struct {

	// Id of the action.
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Id of the team associated to the action.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`
}

type ActionTeamAssociationParameters struct {

	// Id of the action.
	// +crossplane:generate:reference:type=Action
	// +kubebuilder:validation:Optional
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// Reference to a Action to populate actionId.
	// +kubebuilder:validation:Optional
	ActionIDRef *v1.Reference `json:"actionIdRef,omitempty" tf:"-"`

	// Selector for a Action to populate actionId.
	// +kubebuilder:validation:Optional
	ActionIDSelector *v1.Selector `json:"actionIdSelector,omitempty" tf:"-"`

	// Id of the team associated to the action.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDRef *v1.Reference `json:"teamIdRef,omitempty" tf:"-"`

	// Selector for a Team in team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDSelector *v1.Selector `json:"teamIdSelector,omitempty" tf:"-"`
}

// ActionTeamAssociationSpec defines the desired state of ActionTeamAssociation
type ActionTeamAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActionTeamAssociationParameters `json:"forProvider"`
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
	InitProvider ActionTeamAssociationInitParameters `json:"initProvider,omitempty"`
}

// ActionTeamAssociationStatus defines the observed state of ActionTeamAssociation.
type ActionTeamAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActionTeamAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ActionTeamAssociation is the Schema for the ActionTeamAssociations API. Creates and manages an Automation Actions action association with a Team in PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type ActionTeamAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActionTeamAssociationSpec   `json:"spec"`
	Status            ActionTeamAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActionTeamAssociationList contains a list of ActionTeamAssociations
type ActionTeamAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActionTeamAssociation `json:"items"`
}

// Repository type metadata.
var (
	ActionTeamAssociation_Kind             = "ActionTeamAssociation"
	ActionTeamAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActionTeamAssociation_Kind}.String()
	ActionTeamAssociation_KindAPIVersion   = ActionTeamAssociation_Kind + "." + CRDGroupVersion.String()
	ActionTeamAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ActionTeamAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ActionTeamAssociation{}, &ActionTeamAssociationList{})
}
