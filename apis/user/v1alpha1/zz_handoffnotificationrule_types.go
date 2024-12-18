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

type HandoffNotificationRuleContactMethodInitParameters struct {

	// The ID of the contact method.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of the contact method. May be (email_contact_method, email_contact_method_reference, phone_contact_method, phone_contact_method_reference, push_notification_contact_method, push_notification_contact_method_reference, sms_contact_method, sms_contact_method_reference).
	// The type of contact method to notify for. Possible values are 'email_contact_method', 'email_contact_method_reference', 'phone_contact_method', 'phone_contact_method_reference', 'push_notification_contact_method', 'push_notification_contact_method_reference', 'sms_contact_method', 'sms_contact_method_reference'.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HandoffNotificationRuleContactMethodObservation struct {

	// The ID of the contact method.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of the contact method. May be (email_contact_method, email_contact_method_reference, phone_contact_method, phone_contact_method_reference, push_notification_contact_method, push_notification_contact_method_reference, sms_contact_method, sms_contact_method_reference).
	// The type of contact method to notify for. Possible values are 'email_contact_method', 'email_contact_method_reference', 'phone_contact_method', 'phone_contact_method_reference', 'push_notification_contact_method', 'push_notification_contact_method_reference', 'sms_contact_method', 'sms_contact_method_reference'.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HandoffNotificationRuleContactMethodParameters struct {

	// The ID of the contact method.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// The type of the contact method. May be (email_contact_method, email_contact_method_reference, phone_contact_method, phone_contact_method_reference, push_notification_contact_method, push_notification_contact_method_reference, sms_contact_method, sms_contact_method_reference).
	// The type of contact method to notify for. Possible values are 'email_contact_method', 'email_contact_method_reference', 'phone_contact_method', 'phone_contact_method_reference', 'push_notification_contact_method', 'push_notification_contact_method_reference', 'sms_contact_method', 'sms_contact_method_reference'.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type HandoffNotificationRuleInitParameters struct {

	// The contact method to notify the user. Contact method documented below.
	// The contact method to notify for the user handoff notification rule.
	ContactMethod []HandoffNotificationRuleContactMethodInitParameters `json:"contactMethod,omitempty" tf:"contact_method,omitempty"`

	// The type of handoff to notify the user about. Possible values are oncall, offcall, both.
	// The type of handoff to notify for. Possible values are 'both', 'oncall', 'offcall'.
	HandoffType *string `json:"handoffType,omitempty" tf:"handoff_type,omitempty"`

	// The number of minutes before the handoff that the user should be notified. Must be a positive integer greater than or equal to 0.
	// The number of minutes before the handoff to notify the user. Must be greater than or equal to 0.
	NotifyAdvanceInMinutes *float64 `json:"notifyAdvanceInMinutes,omitempty" tf:"notify_advance_in_minutes,omitempty"`

	// The ID of the user.
	// +crossplane:generate:reference:type=User
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

type HandoffNotificationRuleObservation struct {

	// The contact method to notify the user. Contact method documented below.
	// The contact method to notify for the user handoff notification rule.
	ContactMethod []HandoffNotificationRuleContactMethodObservation `json:"contactMethod,omitempty" tf:"contact_method,omitempty"`

	// The type of handoff to notify the user about. Possible values are oncall, offcall, both.
	// The type of handoff to notify for. Possible values are 'both', 'oncall', 'offcall'.
	HandoffType *string `json:"handoffType,omitempty" tf:"handoff_type,omitempty"`

	// The ID of the contact method.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The number of minutes before the handoff that the user should be notified. Must be a positive integer greater than or equal to 0.
	// The number of minutes before the handoff to notify the user. Must be greater than or equal to 0.
	NotifyAdvanceInMinutes *float64 `json:"notifyAdvanceInMinutes,omitempty" tf:"notify_advance_in_minutes,omitempty"`

	// The ID of the user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type HandoffNotificationRuleParameters struct {

	// The contact method to notify the user. Contact method documented below.
	// The contact method to notify for the user handoff notification rule.
	// +kubebuilder:validation:Optional
	ContactMethod []HandoffNotificationRuleContactMethodParameters `json:"contactMethod,omitempty" tf:"contact_method,omitempty"`

	// The type of handoff to notify the user about. Possible values are oncall, offcall, both.
	// The type of handoff to notify for. Possible values are 'both', 'oncall', 'offcall'.
	// +kubebuilder:validation:Optional
	HandoffType *string `json:"handoffType,omitempty" tf:"handoff_type,omitempty"`

	// The number of minutes before the handoff that the user should be notified. Must be a positive integer greater than or equal to 0.
	// The number of minutes before the handoff to notify the user. Must be greater than or equal to 0.
	// +kubebuilder:validation:Optional
	NotifyAdvanceInMinutes *float64 `json:"notifyAdvanceInMinutes,omitempty" tf:"notify_advance_in_minutes,omitempty"`

	// The ID of the user.
	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

// HandoffNotificationRuleSpec defines the desired state of HandoffNotificationRule
type HandoffNotificationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HandoffNotificationRuleParameters `json:"forProvider"`
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
	InitProvider HandoffNotificationRuleInitParameters `json:"initProvider,omitempty"`
}

// HandoffNotificationRuleStatus defines the observed state of HandoffNotificationRule.
type HandoffNotificationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HandoffNotificationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HandoffNotificationRule is the Schema for the HandoffNotificationRules API. Creates and manages an user handoff notification rule in PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type HandoffNotificationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notifyAdvanceInMinutes) || (has(self.initProvider) && has(self.initProvider.notifyAdvanceInMinutes))",message="spec.forProvider.notifyAdvanceInMinutes is a required parameter"
	Spec   HandoffNotificationRuleSpec   `json:"spec"`
	Status HandoffNotificationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HandoffNotificationRuleList contains a list of HandoffNotificationRules
type HandoffNotificationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HandoffNotificationRule `json:"items"`
}

// Repository type metadata.
var (
	HandoffNotificationRule_Kind             = "HandoffNotificationRule"
	HandoffNotificationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HandoffNotificationRule_Kind}.String()
	HandoffNotificationRule_KindAPIVersion   = HandoffNotificationRule_Kind + "." + CRDGroupVersion.String()
	HandoffNotificationRule_GroupVersionKind = CRDGroupVersion.WithKind(HandoffNotificationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&HandoffNotificationRule{}, &HandoffNotificationRuleList{})
}
