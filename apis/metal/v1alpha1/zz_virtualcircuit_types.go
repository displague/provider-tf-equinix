// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type VirtualCircuitInitParameters struct {

	// UUID of Connection where the VC is scoped to.
	// UUID of Connection where the VC is scoped to
	// +crossplane:generate:reference:type=Connection
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Reference to a Connection to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// Selector for a Connection to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	CustomerIP *string `json:"customerIp,omitempty" tf:"customer_ip,omitempty"`

	// Description for the Virtual Circuit resource.
	// Description of the Virtual Circuit resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	MetalIP *string `json:"metalIp,omitempty" tf:"metal_ip,omitempty"`

	// Name of the Virtual Circuit resource.
	// Name of the Virtual Circuit resource
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Equinix Metal network-to-network VLAN ID.
	// Equinix Metal network-to-network VLAN ID (optional when the connection has mode=tunnel)
	NniVlan *float64 `json:"nniVlan,omitempty" tf:"nni_vlan,omitempty"`

	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// UUID of the Connection Port where the VC is scoped to.
	// UUID of the Connection Port where the VC is scoped to
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// UUID of the Project where the VC is scoped to.
	// UUID of the Project where the VC is scoped to
	// +crossplane:generate:reference:type=Project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Speed of the Virtual Circuit resource.
	// Description of the Virtual Circuit speed. This is for information purposes and is computed when the connection type is shared.
	Speed *string `json:"speed,omitempty" tf:"speed,omitempty"`

	// A subnet from one of the IP
	// blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// Tags for the Virtual Circuit resource.
	// Tags attached to the virtual circuit
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// UUID of the VLAN to associate.
	// UUID of the VLAN to associate
	// +crossplane:generate:reference:type=Vlan
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// Reference to a Vlan to populate vlanId.
	// +kubebuilder:validation:Optional
	VlanIDRef *v1.Reference `json:"vlanIdRef,omitempty" tf:"-"`

	// Selector for a Vlan to populate vlanId.
	// +kubebuilder:validation:Optional
	VlanIDSelector *v1.Selector `json:"vlanIdSelector,omitempty" tf:"-"`

	// UUID of the VRF to associate.
	// UUID of the VRF to associate
	// +crossplane:generate:reference:type=Vrf
	VrfID *string `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`

	// Reference to a Vrf to populate vrfId.
	// +kubebuilder:validation:Optional
	VrfIDRef *v1.Reference `json:"vrfIdRef,omitempty" tf:"-"`

	// Selector for a Vrf to populate vrfId.
	// +kubebuilder:validation:Optional
	VrfIDSelector *v1.Selector `json:"vrfIdSelector,omitempty" tf:"-"`
}

type VirtualCircuitObservation struct {

	// UUID of Connection where the VC is scoped to.
	// UUID of Connection where the VC is scoped to
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	CustomerIP *string `json:"customerIp,omitempty" tf:"customer_ip,omitempty"`

	// Description for the Virtual Circuit resource.
	// Description of the Virtual Circuit resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	MetalIP *string `json:"metalIp,omitempty" tf:"metal_ip,omitempty"`

	// Name of the Virtual Circuit resource.
	// Name of the Virtual Circuit resource
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Equinix Metal network-to-network VLAN ID.
	// Equinix Metal network-to-network VLAN ID (optional when the connection has mode=tunnel)
	NniVlan *float64 `json:"nniVlan,omitempty" tf:"nni_vlan,omitempty"`

	// NNI VLAN parameters, see the documentation for Equinix Fabric.
	// Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
	NniVnid *float64 `json:"nniVnid,omitempty" tf:"nni_vnid,omitempty"`

	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// UUID of the Connection Port where the VC is scoped to.
	// UUID of the Connection Port where the VC is scoped to
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// UUID of the Project where the VC is scoped to.
	// UUID of the Project where the VC is scoped to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Speed of the Virtual Circuit resource.
	// Description of the Virtual Circuit speed. This is for information purposes and is computed when the connection type is shared.
	Speed *string `json:"speed,omitempty" tf:"speed,omitempty"`

	// Status of the virtal circuit.
	// Status of the virtual circuit resource
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A subnet from one of the IP
	// blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// Tags for the Virtual Circuit resource.
	// Tags attached to the virtual circuit
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// UUID of the VLAN to associate.
	// UUID of the VLAN to associate
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// VNID VLAN parameter, see the documentation for Equinix Fabric.
	// VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
	Vnid *float64 `json:"vnid,omitempty" tf:"vnid,omitempty"`

	// UUID of the VRF to associate.
	// UUID of the VRF to associate
	VrfID *string `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

type VirtualCircuitParameters struct {

	// UUID of Connection where the VC is scoped to.
	// UUID of Connection where the VC is scoped to
	// +crossplane:generate:reference:type=Connection
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Reference to a Connection to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// Selector for a Connection to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	// +kubebuilder:validation:Optional
	CustomerIP *string `json:"customerIp,omitempty" tf:"customer_ip,omitempty"`

	// Description for the Virtual Circuit resource.
	// Description of the Virtual Circuit resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The password that can be set for the VRF BGP peer
	// The password that can be set for the VRF BGP peer
	// +kubebuilder:validation:Optional
	Md5SecretRef *v1.SecretKeySelector `json:"md5SecretRef,omitempty" tf:"-"`

	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	// +kubebuilder:validation:Optional
	MetalIP *string `json:"metalIp,omitempty" tf:"metal_ip,omitempty"`

	// Name of the Virtual Circuit resource.
	// Name of the Virtual Circuit resource
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Equinix Metal network-to-network VLAN ID.
	// Equinix Metal network-to-network VLAN ID (optional when the connection has mode=tunnel)
	// +kubebuilder:validation:Optional
	NniVlan *float64 `json:"nniVlan,omitempty" tf:"nni_vlan,omitempty"`

	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	// +kubebuilder:validation:Optional
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// UUID of the Connection Port where the VC is scoped to.
	// UUID of the Connection Port where the VC is scoped to
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// UUID of the Project where the VC is scoped to.
	// UUID of the Project where the VC is scoped to
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Speed of the Virtual Circuit resource.
	// Description of the Virtual Circuit speed. This is for information purposes and is computed when the connection type is shared.
	// +kubebuilder:validation:Optional
	Speed *string `json:"speed,omitempty" tf:"speed,omitempty"`

	// A subnet from one of the IP
	// blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	// +kubebuilder:validation:Optional
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// Tags for the Virtual Circuit resource.
	// Tags attached to the virtual circuit
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// UUID of the VLAN to associate.
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

	// UUID of the VRF to associate.
	// UUID of the VRF to associate
	// +crossplane:generate:reference:type=Vrf
	// +kubebuilder:validation:Optional
	VrfID *string `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`

	// Reference to a Vrf to populate vrfId.
	// +kubebuilder:validation:Optional
	VrfIDRef *v1.Reference `json:"vrfIdRef,omitempty" tf:"-"`

	// Selector for a Vrf to populate vrfId.
	// +kubebuilder:validation:Optional
	VrfIDSelector *v1.Selector `json:"vrfIdSelector,omitempty" tf:"-"`
}

// VirtualCircuitSpec defines the desired state of VirtualCircuit
type VirtualCircuitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualCircuitParameters `json:"forProvider"`
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
	InitProvider VirtualCircuitInitParameters `json:"initProvider,omitempty"`
}

// VirtualCircuitStatus defines the observed state of VirtualCircuit.
type VirtualCircuitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualCircuitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VirtualCircuit is the Schema for the VirtualCircuits API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type VirtualCircuit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.portId) || (has(self.initProvider) && has(self.initProvider.portId))",message="spec.forProvider.portId is a required parameter"
	Spec   VirtualCircuitSpec   `json:"spec"`
	Status VirtualCircuitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualCircuitList contains a list of VirtualCircuits
type VirtualCircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualCircuit `json:"items"`
}

// Repository type metadata.
var (
	VirtualCircuit_Kind             = "VirtualCircuit"
	VirtualCircuit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualCircuit_Kind}.String()
	VirtualCircuit_KindAPIVersion   = VirtualCircuit_Kind + "." + CRDGroupVersion.String()
	VirtualCircuit_GroupVersionKind = CRDGroupVersion.WithKind(VirtualCircuit_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualCircuit{}, &VirtualCircuitList{})
}
