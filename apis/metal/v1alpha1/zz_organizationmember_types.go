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

type OrganizationMemberObservation struct {
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InvitedBy *string `json:"invitedBy,omitempty" tf:"invited_by,omitempty"`

	Nonce *string `json:"nonce,omitempty" tf:"nonce,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type OrganizationMemberParameters struct {

	// The email address of the user to invite
	// +kubebuilder:validation:Required
	Invitee *string `json:"invitee" tf:"invitee,omitempty"`

	// A message to the invitee (only used during the invitation stage)
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// The organization to invite the user to
	// +kubebuilder:validation:Required
	OrganizationID *string `json:"organizationId" tf:"organization_id,omitempty"`

	// Project IDs the member has access to within the organization. If the member is an 'owner', the projects list should be empty.
	// +kubebuilder:validation:Required
	ProjectsIds []*string `json:"projectsIds" tf:"projects_ids,omitempty"`

	// Organization roles (owner, collaborator, limited_collaborator, billing)
	// +kubebuilder:validation:Required
	Roles []*string `json:"roles" tf:"roles,omitempty"`
}

// OrganizationMemberSpec defines the desired state of OrganizationMember
type OrganizationMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationMemberParameters `json:"forProvider"`
}

// OrganizationMemberStatus defines the observed state of OrganizationMember.
type OrganizationMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationMember is the Schema for the OrganizationMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinixjet}
type OrganizationMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationMemberSpec   `json:"spec"`
	Status            OrganizationMemberStatus `json:"status,omitempty"`
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
