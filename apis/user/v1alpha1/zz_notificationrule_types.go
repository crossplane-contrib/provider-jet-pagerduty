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

type NotificationRuleInitParameters struct {

	// A contact method block, configured as a block described below.
	// +mapType=granular
	ContactMethod map[string]*string `json:"contactMethod,omitempty" tf:"contact_method,omitempty"`

	// The delay before firing the rule, in minutes.
	StartDelayInMinutes *float64 `json:"startDelayInMinutes,omitempty" tf:"start_delay_in_minutes,omitempty"`

	// Which incident urgency this rule is used for. Account must have the urgencies ability to have a low urgency notification rule. Can be high or low.
	Urgency *string `json:"urgency,omitempty" tf:"urgency,omitempty"`

	// The ID of the user.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/user/v1alpha1.User
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

type NotificationRuleObservation struct {

	// A contact method block, configured as a block described below.
	// +mapType=granular
	ContactMethod map[string]*string `json:"contactMethod,omitempty" tf:"contact_method,omitempty"`

	// The id of the referenced contact method.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The delay before firing the rule, in minutes.
	StartDelayInMinutes *float64 `json:"startDelayInMinutes,omitempty" tf:"start_delay_in_minutes,omitempty"`

	// Which incident urgency this rule is used for. Account must have the urgencies ability to have a low urgency notification rule. Can be high or low.
	Urgency *string `json:"urgency,omitempty" tf:"urgency,omitempty"`

	// The ID of the user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type NotificationRuleParameters struct {

	// A contact method block, configured as a block described below.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ContactMethod map[string]*string `json:"contactMethod,omitempty" tf:"contact_method,omitempty"`

	// The delay before firing the rule, in minutes.
	// +kubebuilder:validation:Optional
	StartDelayInMinutes *float64 `json:"startDelayInMinutes,omitempty" tf:"start_delay_in_minutes,omitempty"`

	// Which incident urgency this rule is used for. Account must have the urgencies ability to have a low urgency notification rule. Can be high or low.
	// +kubebuilder:validation:Optional
	Urgency *string `json:"urgency,omitempty" tf:"urgency,omitempty"`

	// The ID of the user.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/user/v1alpha1.User
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

// NotificationRuleSpec defines the desired state of NotificationRule
type NotificationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationRuleParameters `json:"forProvider"`
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
	InitProvider NotificationRuleInitParameters `json:"initProvider,omitempty"`
}

// NotificationRuleStatus defines the observed state of NotificationRule.
type NotificationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NotificationRule is the Schema for the NotificationRules API. Creates and manages notification rules for a user in PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type NotificationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.contactMethod) || (has(self.initProvider) && has(self.initProvider.contactMethod))",message="spec.forProvider.contactMethod is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.startDelayInMinutes) || (has(self.initProvider) && has(self.initProvider.startDelayInMinutes))",message="spec.forProvider.startDelayInMinutes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.urgency) || (has(self.initProvider) && has(self.initProvider.urgency))",message="spec.forProvider.urgency is a required parameter"
	Spec   NotificationRuleSpec   `json:"spec"`
	Status NotificationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationRuleList contains a list of NotificationRules
type NotificationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationRule `json:"items"`
}

// Repository type metadata.
var (
	NotificationRule_Kind             = "NotificationRule"
	NotificationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationRule_Kind}.String()
	NotificationRule_KindAPIVersion   = NotificationRule_Kind + "." + CRDGroupVersion.String()
	NotificationRule_GroupVersionKind = CRDGroupVersion.WithKind(NotificationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationRule{}, &NotificationRuleList{})
}
