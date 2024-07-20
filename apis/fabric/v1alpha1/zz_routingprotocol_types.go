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

type BGPIPv4InitParameters struct {

	// Customer side peering ip
	CustomerPeerIP *string `json:"customerPeerIp,omitempty" tf:"customer_peer_ip,omitempty"`

	// Admin status for the BGP session
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BGPIPv4Observation struct {

	// Customer side peering ip
	CustomerPeerIP *string `json:"customerPeerIp,omitempty" tf:"customer_peer_ip,omitempty"`

	// Admin status for the BGP session
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Equinix side peering ip
	EquinixPeerIP *string `json:"equinixPeerIp,omitempty" tf:"equinix_peer_ip,omitempty"`
}

type BGPIPv4Parameters struct {

	// Customer side peering ip
	// +kubebuilder:validation:Optional
	CustomerPeerIP *string `json:"customerPeerIp" tf:"customer_peer_ip,omitempty"`

	// Admin status for the BGP session
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BGPIPv6InitParameters struct {

	// Customer side peering ip
	CustomerPeerIP *string `json:"customerPeerIp,omitempty" tf:"customer_peer_ip,omitempty"`

	// Admin status for the BGP session
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BGPIPv6Observation struct {

	// Customer side peering ip
	CustomerPeerIP *string `json:"customerPeerIp,omitempty" tf:"customer_peer_ip,omitempty"`

	// Admin status for the BGP session
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Equinix side peering ip
	EquinixPeerIP *string `json:"equinixPeerIp,omitempty" tf:"equinix_peer_ip,omitempty"`
}

type BGPIPv6Parameters struct {

	// Customer side peering ip
	// +kubebuilder:validation:Optional
	CustomerPeerIP *string `json:"customerPeerIp" tf:"customer_peer_ip,omitempty"`

	// Admin status for the BGP session
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BfdInitParameters struct {

	// Bidirectional Forwarding Detection enablement
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval range between the received BFD control packets
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`
}

type BfdObservation struct {

	// Bidirectional Forwarding Detection enablement
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval range between the received BFD control packets
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`
}

type BfdParameters struct {

	// Bidirectional Forwarding Detection enablement
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Interval range between the received BFD control packets
	// +kubebuilder:validation:Optional
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`
}

type DirectIPv4InitParameters struct {

	// Equinix side Interface IP address
	EquinixIfaceIP *string `json:"equinixIfaceIp,omitempty" tf:"equinix_iface_ip,omitempty"`
}

type DirectIPv4Observation struct {

	// Equinix side Interface IP address
	EquinixIfaceIP *string `json:"equinixIfaceIp,omitempty" tf:"equinix_iface_ip,omitempty"`
}

type DirectIPv4Parameters struct {

	// Equinix side Interface IP address
	// +kubebuilder:validation:Optional
	EquinixIfaceIP *string `json:"equinixIfaceIp" tf:"equinix_iface_ip,omitempty"`
}

type DirectIPv6InitParameters struct {

	// Equinix side Interface IP address
	EquinixIfaceIP *string `json:"equinixIfaceIp,omitempty" tf:"equinix_iface_ip,omitempty"`
}

type DirectIPv6Observation struct {

	// Equinix side Interface IP address
	EquinixIfaceIP *string `json:"equinixIfaceIp,omitempty" tf:"equinix_iface_ip,omitempty"`
}

type DirectIPv6Parameters struct {

	// Equinix side Interface IP address
	// +kubebuilder:validation:Optional
	EquinixIfaceIP *string `json:"equinixIfaceIp,omitempty" tf:"equinix_iface_ip,omitempty"`
}

type OperationErrorsAdditionalInfoInitParameters struct {
}

type OperationErrorsAdditionalInfoObservation struct {
	Property *string `json:"property,omitempty" tf:"property,omitempty"`

	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`
}

type OperationErrorsAdditionalInfoParameters struct {
}

type OperationErrorsInitParameters struct {
}

type OperationErrorsObservation struct {
	AdditionalInfo []OperationErrorsAdditionalInfoObservation `json:"additionalInfo,omitempty" tf:"additional_info,omitempty"`

	CorrelationID *string `json:"correlationId,omitempty" tf:"correlation_id,omitempty"`

	Details *string `json:"details,omitempty" tf:"details,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty" tf:"error_code,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty" tf:"error_message,omitempty"`

	Help *string `json:"help,omitempty" tf:"help,omitempty"`
}

type OperationErrorsParameters struct {
}

type RoutingProtocolChangeInitParameters struct {
}

type RoutingProtocolChangeLogInitParameters struct {
}

type RoutingProtocolChangeLogObservation struct {
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	CreatedByEmail *string `json:"createdByEmail,omitempty" tf:"created_by_email,omitempty"`

	CreatedByFullName *string `json:"createdByFullName,omitempty" tf:"created_by_full_name,omitempty"`

	CreatedDateTime *string `json:"createdDateTime,omitempty" tf:"created_date_time,omitempty"`

	DeletedBy *string `json:"deletedBy,omitempty" tf:"deleted_by,omitempty"`

	DeletedByEmail *string `json:"deletedByEmail,omitempty" tf:"deleted_by_email,omitempty"`

	DeletedByFullName *string `json:"deletedByFullName,omitempty" tf:"deleted_by_full_name,omitempty"`

	DeletedDateTime *string `json:"deletedDateTime,omitempty" tf:"deleted_date_time,omitempty"`

	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by,omitempty"`

	UpdatedByEmail *string `json:"updatedByEmail,omitempty" tf:"updated_by_email,omitempty"`

	UpdatedByFullName *string `json:"updatedByFullName,omitempty" tf:"updated_by_full_name,omitempty"`

	UpdatedDateTime *string `json:"updatedDateTime,omitempty" tf:"updated_date_time,omitempty"`
}

type RoutingProtocolChangeLogParameters struct {
}

type RoutingProtocolChangeObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type RoutingProtocolChangeParameters struct {
}

type RoutingProtocolInitParameters struct {

	// BGP authorization key
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// Routing Protocol BGP IPv4
	BGPIPv4 []BGPIPv4InitParameters `json:"bgpIpv4,omitempty" tf:"bgp_ipv4,omitempty"`

	// Routing Protocol BGP IPv6
	BGPIPv6 []BGPIPv6InitParameters `json:"bgpIpv6,omitempty" tf:"bgp_ipv6,omitempty"`

	// Bidirectional Forwarding Detection
	Bfd []BfdInitParameters `json:"bfd,omitempty" tf:"bfd,omitempty"`

	// Connection URI associated with Routing Protocol
	ConnectionUUID *string `json:"connectionUuid,omitempty" tf:"connection_uuid,omitempty"`

	// Customer-provided ASN
	CustomerAsn *float64 `json:"customerAsn,omitempty" tf:"customer_asn,omitempty"`

	// Customer-provided Fabric Routing Protocol description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Routing Protocol Direct IPv4
	DirectIPv4 []DirectIPv4InitParameters `json:"directIpv4,omitempty" tf:"direct_ipv4,omitempty"`

	// Routing Protocol Direct IPv6
	DirectIPv6 []DirectIPv6InitParameters `json:"directIpv6,omitempty" tf:"direct_ipv6,omitempty"`

	// Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Defines the routing protocol type like BGP or DIRECT
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Equinix-assigned routing protocol identifier
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type RoutingProtocolObservation struct {

	// BGP authorization key
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// Routing Protocol BGP IPv4
	BGPIPv4 []BGPIPv4Observation `json:"bgpIpv4,omitempty" tf:"bgp_ipv4,omitempty"`

	// Routing Protocol BGP IPv6
	BGPIPv6 []BGPIPv6Observation `json:"bgpIpv6,omitempty" tf:"bgp_ipv6,omitempty"`

	// Bidirectional Forwarding Detection
	Bfd []BfdObservation `json:"bfd,omitempty" tf:"bfd,omitempty"`

	// Routing Protocol configuration Changes
	Change []RoutingProtocolChangeObservation `json:"change,omitempty" tf:"change,omitempty"`

	// Captures Routing Protocol lifecycle change information
	ChangeLog []RoutingProtocolChangeLogObservation `json:"changeLog,omitempty" tf:"change_log,omitempty"`

	// Connection URI associated with Routing Protocol
	ConnectionUUID *string `json:"connectionUuid,omitempty" tf:"connection_uuid,omitempty"`

	// Customer-provided ASN
	CustomerAsn *float64 `json:"customerAsn,omitempty" tf:"customer_asn,omitempty"`

	// Customer-provided Fabric Routing Protocol description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Routing Protocol Direct IPv4
	DirectIPv4 []DirectIPv4Observation `json:"directIpv4,omitempty" tf:"direct_ipv4,omitempty"`

	// Routing Protocol Direct IPv6
	DirectIPv6 []DirectIPv6Observation `json:"directIpv6,omitempty" tf:"direct_ipv6,omitempty"`

	// Equinix ASN
	EquinixAsn *float64 `json:"equinixAsn,omitempty" tf:"equinix_asn,omitempty"`

	// Routing Protocol URI information
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Routing Protocol type-specific operational data
	Operation []RoutingProtocolOperationObservation `json:"operation,omitempty" tf:"operation,omitempty"`

	// Routing Protocol overall state
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Defines the routing protocol type like BGP or DIRECT
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Equinix-assigned routing protocol identifier
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type RoutingProtocolOperationInitParameters struct {
}

type RoutingProtocolOperationObservation struct {
	Errors []OperationErrorsObservation `json:"errors,omitempty" tf:"errors,omitempty"`
}

type RoutingProtocolOperationParameters struct {
}

type RoutingProtocolParameters struct {

	// BGP authorization key
	// +kubebuilder:validation:Optional
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// Routing Protocol BGP IPv4
	// +kubebuilder:validation:Optional
	BGPIPv4 []BGPIPv4Parameters `json:"bgpIpv4,omitempty" tf:"bgp_ipv4,omitempty"`

	// Routing Protocol BGP IPv6
	// +kubebuilder:validation:Optional
	BGPIPv6 []BGPIPv6Parameters `json:"bgpIpv6,omitempty" tf:"bgp_ipv6,omitempty"`

	// Bidirectional Forwarding Detection
	// +kubebuilder:validation:Optional
	Bfd []BfdParameters `json:"bfd,omitempty" tf:"bfd,omitempty"`

	// Connection URI associated with Routing Protocol
	// +kubebuilder:validation:Optional
	ConnectionUUID *string `json:"connectionUuid,omitempty" tf:"connection_uuid,omitempty"`

	// Customer-provided ASN
	// +kubebuilder:validation:Optional
	CustomerAsn *float64 `json:"customerAsn,omitempty" tf:"customer_asn,omitempty"`

	// Customer-provided Fabric Routing Protocol description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Routing Protocol Direct IPv4
	// +kubebuilder:validation:Optional
	DirectIPv4 []DirectIPv4Parameters `json:"directIpv4,omitempty" tf:"direct_ipv4,omitempty"`

	// Routing Protocol Direct IPv6
	// +kubebuilder:validation:Optional
	DirectIPv6 []DirectIPv6Parameters `json:"directIpv6,omitempty" tf:"direct_ipv6,omitempty"`

	// Routing Protocol name. An alpha-numeric 24 characters string which can include only hyphens and underscores
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Defines the routing protocol type like BGP or DIRECT
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Equinix-assigned routing protocol identifier
	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

// RoutingProtocolSpec defines the desired state of RoutingProtocol
type RoutingProtocolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoutingProtocolParameters `json:"forProvider"`
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
	InitProvider RoutingProtocolInitParameters `json:"initProvider,omitempty"`
}

// RoutingProtocolStatus defines the observed state of RoutingProtocol.
type RoutingProtocolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoutingProtocolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RoutingProtocol is the Schema for the RoutingProtocols API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type RoutingProtocol struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionUuid) || (has(self.initProvider) && has(self.initProvider.connectionUuid))",message="spec.forProvider.connectionUuid is a required parameter"
	Spec   RoutingProtocolSpec   `json:"spec"`
	Status RoutingProtocolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoutingProtocolList contains a list of RoutingProtocols
type RoutingProtocolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoutingProtocol `json:"items"`
}

// Repository type metadata.
var (
	RoutingProtocol_Kind             = "RoutingProtocol"
	RoutingProtocol_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoutingProtocol_Kind}.String()
	RoutingProtocol_KindAPIVersion   = RoutingProtocol_Kind + "." + CRDGroupVersion.String()
	RoutingProtocol_GroupVersionKind = CRDGroupVersion.WithKind(RoutingProtocol_Kind)
)

func init() {
	SchemeBuilder.Register(&RoutingProtocol{}, &RoutingProtocolList{})
}
