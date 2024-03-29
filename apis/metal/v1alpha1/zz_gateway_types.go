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

type GatewayObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Status of the gateway resource.
	// Status of the gateway resource
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// UUID of the VRF associated with the IP Reservation
	// UUID of the VRF associated with the IP Reservation
	VrfID *string `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

type GatewayParameters struct {

	// UUID of Public or VRF IP Reservation to associate with the gateway, the
	// reservation must be in the same metro as the VLAN, conflicts with private_ipv4_subnet_size.
	// UUID of the Public or VRF IP Reservation to associate
	// +crossplane:generate:reference:type=ReservedIPBlock
	// +kubebuilder:validation:Optional
	IPReservationID *string `json:"ipReservationId,omitempty" tf:"ip_reservation_id,omitempty"`

	// Reference to a ReservedIPBlock to populate ipReservationId.
	// +kubebuilder:validation:Optional
	IPReservationIDRef *v1.Reference `json:"ipReservationIdRef,omitempty" tf:"-"`

	// Selector for a ReservedIPBlock to populate ipReservationId.
	// +kubebuilder:validation:Optional
	IPReservationIDSelector *v1.Selector `json:"ipReservationIdSelector,omitempty" tf:"-"`

	// Size of the private IPv4 subnet to create for this metal
	// gateway, must be one of 8, 16, 32, 64, 128. Conflicts with ip_reservation_id.
	// Size of the private IPv4 subnet to create for this gateway
	// +kubebuilder:validation:Optional
	PrivateIPv4SubnetSize *float64 `json:"privateIpv4SubnetSize,omitempty" tf:"private_ipv4_subnet_size,omitempty"`

	// UUID of the project where the gateway is scoped to.
	// UUID of the Project where the Gateway is scoped to
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// UUID of the VLAN where the gateway is scoped to.
	// UUID of the VLAN to associate
	// +crossplane:generate:reference:type=Vlan
	// +kubebuilder:validation:Optional
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// Reference to a Vlan to populate vlanId.
	// +kubebuilder:validation:Optional
	VlanIDRef *v1.Reference `json:"vlanIdRef,omitempty" tf:"-"`

	// Selector for a Vlan to populate vlanId.
	// +kubebuilder:validation:Optional
	VlanIDSelector *v1.Selector `json:"vlanIdSelector,omitempty" tf:"-"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Gateway is the Schema for the Gateways API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec   `json:"spec"`
	Status            GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	Gateway_Kind             = "Gateway"
	Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gateway_Kind}.String()
	Gateway_KindAPIVersion   = Gateway_Kind + "." + CRDGroupVersion.String()
	Gateway_GroupVersionKind = CRDGroupVersion.WithKind(Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
