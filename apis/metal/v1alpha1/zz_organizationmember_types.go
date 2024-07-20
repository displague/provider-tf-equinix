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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OrganizationMemberInitParameters struct {

	// The email address of the user to invite
	// The email address of the user to invite
	Invitee *string `json:"invitee,omitempty" tf:"invitee,omitempty"`

	// A message to include in the emailed invitation.
	// A message to the invitee (only used during the invitation stage)
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The organization to invite the user to
	// The organization to invite the user to
	// +crossplane:generate:reference:type=Organization
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Reference to a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// Selector for a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
	// Project IDs the member has access to within the organization. If the member is an 'owner', the projects list should be empty.
	// +listType=set
	ProjectsIds []*string `json:"projectsIds,omitempty" tf:"projects_ids,omitempty"`

	// Organization roles (admin, collaborator, limited_collaborator, billing)
	// Organization roles (owner, collaborator, limited_collaborator, billing)
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type OrganizationMemberObservation struct {

	// When the invitation was created (only known in the invitation stage)
	// When the invitation was created (only known in the invitation stage)
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// The unique ID of the membership.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The user_id of the user that sent the invitation (only known in the invitation stage)
	// The user id of the user that sent the invitation (only known in the invitation stage)
	InvitedBy *string `json:"invitedBy,omitempty" tf:"invited_by,omitempty"`

	// The email address of the user to invite
	// The email address of the user to invite
	Invitee *string `json:"invitee,omitempty" tf:"invitee,omitempty"`

	// A message to include in the emailed invitation.
	// A message to the invitee (only used during the invitation stage)
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The nonce for the invitation (only known in the invitation stage)
	// The nonce for the invitation (only known in the invitation stage)
	Nonce *string `json:"nonce,omitempty" tf:"nonce,omitempty"`

	// The organization to invite the user to
	// The organization to invite the user to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
	// Project IDs the member has access to within the organization. If the member is an 'owner', the projects list should be empty.
	// +listType=set
	ProjectsIds []*string `json:"projectsIds,omitempty" tf:"projects_ids,omitempty"`

	// Organization roles (admin, collaborator, limited_collaborator, billing)
	// Organization roles (owner, collaborator, limited_collaborator, billing)
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// The state of the membership ('invited' when an invitation is open, 'active' when the user is an organization member)
	// The state of the membership ('invited' when an invitation is open, 'active' when the user is an organization member)
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// When the invitation was updated (only known in the invitation stage)
	// When the invitation was updated (only known in the invitation stage)
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type OrganizationMemberParameters struct {

	// The email address of the user to invite
	// The email address of the user to invite
	// +kubebuilder:validation:Optional
	Invitee *string `json:"invitee,omitempty" tf:"invitee,omitempty"`

	// A message to include in the emailed invitation.
	// A message to the invitee (only used during the invitation stage)
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The organization to invite the user to
	// The organization to invite the user to
	// +crossplane:generate:reference:type=Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Reference to a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// Selector for a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.
	// Project IDs the member has access to within the organization. If the member is an 'owner', the projects list should be empty.
	// +kubebuilder:validation:Optional
	// +listType=set
	ProjectsIds []*string `json:"projectsIds,omitempty" tf:"projects_ids,omitempty"`

	// Organization roles (admin, collaborator, limited_collaborator, billing)
	// Organization roles (owner, collaborator, limited_collaborator, billing)
	// +kubebuilder:validation:Optional
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

// OrganizationMemberSpec defines the desired state of OrganizationMember
type OrganizationMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationMemberParameters `json:"forProvider"`
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
	InitProvider OrganizationMemberInitParameters `json:"initProvider,omitempty"`
}

// OrganizationMemberStatus defines the observed state of OrganizationMember.
type OrganizationMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OrganizationMember is the Schema for the OrganizationMembers API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type OrganizationMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.invitee) || (has(self.initProvider) && has(self.initProvider.invitee))",message="spec.forProvider.invitee is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectsIds) || (has(self.initProvider) && has(self.initProvider.projectsIds))",message="spec.forProvider.projectsIds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roles) || (has(self.initProvider) && has(self.initProvider.roles))",message="spec.forProvider.roles is a required parameter"
	Spec   OrganizationMemberSpec   `json:"spec"`
	Status OrganizationMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationMemberList contains a list of OrganizationMembers
type OrganizationMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationMember `json:"items"`
}

// Repository type metadata.
var (
	OrganizationMember_Kind             = "OrganizationMember"
	OrganizationMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationMember_Kind}.String()
	OrganizationMember_KindAPIVersion   = OrganizationMember_Kind + "." + CRDGroupVersion.String()
	OrganizationMember_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationMember_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationMember{}, &OrganizationMemberList{})
}
