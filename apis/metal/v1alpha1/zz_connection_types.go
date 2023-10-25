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

type ConnectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of connection ports - primary (ports[0]) and secondary (ports[1]). Schema of
	// port is described in documentation of the
	// equinix_metal_connection datasource.
	// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`)
	Ports []PortsObservation `json:"ports,omitempty" tf:"ports,omitempty"`

	// List of connection service tokens with attributes required to configure the connection in Equinix Fabric with the equinix_ecx_l2_connection resource or from the Equinix Fabric Portal. Scehma of service_token is described in documentation of the equinix_metal_connection datasource.
	// Only used with shared connection. List of service tokens required to continue the setup process with [equinix_ecx_l2_connection](https://registry.io/providers/equinix/equinix/latest/docs/resources/equinix_ecx_l2_connection) or from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
	ServiceTokens []ServiceTokensObservation `json:"serviceTokens,omitempty" tf:"service_tokens,omitempty"`

	// Status of the connection resource.
	// Status of the connection resource
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Deprecated) Fabric Token required to configure the connection in Equinix Fabric with the equinix_ecx_l2_connection resource or from the Equinix Fabric Portal. If your organization already has connection service tokens enabled, use service_tokens instead.
	// Only used with shared connection. Fabric Token required to continue the setup process with [equinix_ecx_l2_connection](https://registry.io/providers/equinix/equinix/latest/docs/resources/equinix_ecx_l2_connection) or from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
	Token *string `json:"token,omitempty" tf:"token,omitempty"`
}

type ConnectionParameters struct {

	// The preferred email used for communication and notifications about the Equinix Fabric interconnection. Required when using a Project API key. Optional and defaults to the primary user email address when using a User API key.
	// The preferred email used for communication and notifications about the Equinix Fabric interconnection. Required when using a Project API key. Optional and defaults to the primary user email address when using a User API key
	// +kubebuilder:validation:Optional
	ContactEmail *string `json:"contactEmail,omitempty" tf:"contact_email,omitempty"`

	// Description for the connection resource.
	// Description of the connection resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Deprecated) Facility where the connection will be created.   Use metro instead; read the facility to metro migration guide
	// Facility where the connection will be created
	// +kubebuilder:validation:Optional
	Facility *string `json:"facility,omitempty" tf:"facility,omitempty"`

	// Metro where the connection will be created.
	// Metro where the connection will be created
	// +kubebuilder:validation:Optional
	Metro *string `json:"metro,omitempty" tf:"metro,omitempty"`

	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Name of the connection resource
	// Name of the connection resource
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// ID of the organization where the connection is scoped to.
	// ID of the organization responsible for the connection. Applicable with type "dedicated"
	// +crossplane:generate:reference:type=Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Reference to a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// Selector for a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// ID of the project where the connection is scoped to, must be set for.
	// ID of the project where the connection is scoped to. Required with type "shared"
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Connection redundancy - redundant or primary.
	// Connection redundancy - redundant or primary
	// +kubebuilder:validation:Required
	Redundancy *string `json:"redundancy" tf:"redundancy,omitempty"`

	// Only used with shared connection. Type of service token to use for the connection, a_side or z_side. Type of service token to use for the connection, a_side or z_side
	// +kubebuilder:validation:Optional
	ServiceTokenType *string `json:"serviceTokenType,omitempty" tf:"service_token_type,omitempty"`

	// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
	// Port speed. Required for a_side connections. Allowed values are 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps
	// +kubebuilder:validation:Optional
	Speed *string `json:"speed,omitempty" tf:"speed,omitempty"`

	// String list of tags.
	// Tags attached to the connection
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Connection type - dedicated or shared.
	// Connection type - dedicated or shared
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.
	// Only used with shared connection. VLANs to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection
	// +kubebuilder:validation:Optional
	Vlans []*float64 `json:"vlans,omitempty" tf:"vlans,omitempty"`
}

type PortsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Status of the connection resource.
	LinkStatus *string `json:"linkStatus,omitempty" tf:"link_status,omitempty"`

	// Name of the connection resource
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
	Speed *float64 `json:"speed,omitempty" tf:"speed,omitempty"`

	// Status of the connection resource.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	VirtualCircuitIds []*string `json:"virtualCircuitIds,omitempty" tf:"virtual_circuit_ids,omitempty"`
}

type PortsParameters struct {
}

type ServiceTokensObservation struct {
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
	MaxAllowedSpeed *string `json:"maxAllowedSpeed,omitempty" tf:"max_allowed_speed,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Connection type - dedicated or shared.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServiceTokensParameters struct {
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
