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

type AccountObservation struct {
}

type AccountParameters struct {

	// Account Number
	// +kubebuilder:validation:Optional
	AccountNumber *float64 `json:"accountNumber,omitempty" tf:"account_number,omitempty"`
}

type ChangeLogObservation struct {
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

type ChangeLogParameters struct {
}

type CloudRouterObservation struct {

	// Number of IPv4 BGP routes in use (including non-distinct prefixes)
	BGPIPv4RoutesCount *float64 `json:"bgpIpv4RoutesCount,omitempty" tf:"bgp_ipv4_routes_count,omitempty"`

	// Number of IPv6 BGP routes in use (including non-distinct prefixes)
	BGPIPv6RoutesCount *float64 `json:"bgpIpv6RoutesCount,omitempty" tf:"bgp_ipv6_routes_count,omitempty"`

	// Captures Fabric Cloud Router lifecycle change information
	ChangeLog []ChangeLogObservation `json:"changeLog,omitempty" tf:"change_log,omitempty"`

	// Number of connections associated with this Fabric Cloud Router instance
	ConnectionsCount *float64 `json:"connectionsCount,omitempty" tf:"connections_count,omitempty"`

	// Number of distinct IPv4 routes
	DistinctIPv4PrefixesCount *float64 `json:"distinctIpv4PrefixesCount,omitempty" tf:"distinct_ipv4_prefixes_count,omitempty"`

	// Number of distinct IPv6 routes
	DistinctIPv6PrefixesCount *float64 `json:"distinctIpv6PrefixesCount,omitempty" tf:"distinct_ipv6_prefixes_count,omitempty"`

	// Equinix ASN
	EquinixAsn *float64 `json:"equinixAsn,omitempty" tf:"equinix_asn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Fabric Cloud Router overall state
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type CloudRouterParameters struct {

	// Customer account information that is associated with this Fabric Cloud Router
	// +kubebuilder:validation:Required
	Account []AccountParameters `json:"account" tf:"account,omitempty"`

	// Customer-provided Fabric Cloud Router description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Fabric Cloud Router URI information
	// +kubebuilder:validation:Optional
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// Fabric Cloud Router location
	// +kubebuilder:validation:Required
	Location []LocationParameters `json:"location" tf:"location,omitempty"`

	// Fabric Cloud Router name. An alpha-numeric 24 characters string which can include only hyphens and underscores
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Preferences for notifications on Fabric Cloud Router configuration or status changes
	// +kubebuilder:validation:Required
	Notifications []NotificationsParameters `json:"notifications" tf:"notifications,omitempty"`

	// Order information related to this Fabric Cloud Router
	// +kubebuilder:validation:Required
	Order []OrderParameters `json:"order" tf:"order,omitempty"`

	// Fabric Cloud Router Package Type
	// +kubebuilder:validation:Required
	Package []PackageParameters `json:"package" tf:"package,omitempty"`

	// Customer resource hierarchy project information.Applicable to customers onboarded to Equinix Identity and Access Management. For more information see Identity and Access Management: Projects
	// +kubebuilder:validation:Required
	Project []ProjectParameters `json:"project" tf:"project,omitempty"`

	// Defines the FCR type like; XF_ROUTER
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Equinix-assigned Fabric Cloud Router identifier
	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type LocationObservation struct {
}

type LocationParameters struct {

	// IBX Code
	// +kubebuilder:validation:Optional
	Ibx *string `json:"ibx,omitempty" tf:"ibx,omitempty"`

	// Access point metro code
	// +kubebuilder:validation:Optional
	MetroCode *string `json:"metroCode,omitempty" tf:"metro_code,omitempty"`

	// Access point metro name
	// +kubebuilder:validation:Optional
	MetroName *string `json:"metroName,omitempty" tf:"metro_name,omitempty"`

	// Access point region
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type NotificationsObservation struct {
}

type NotificationsParameters struct {

	// Array of contact emails
	// +kubebuilder:validation:Required
	Emails []*string `json:"emails" tf:"emails,omitempty"`

	// Send interval
	// +kubebuilder:validation:Optional
	SendInterval *string `json:"sendInterval,omitempty" tf:"send_interval,omitempty"`

	// Notification Type - ALL,CONNECTION_APPROVAL,SALES_REP_NOTIFICATIONS, NOTIFICATIONS
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type OrderObservation struct {
}

type OrderParameters struct {

	// Billing tier for connection bandwidth
	// +kubebuilder:validation:Optional
	BillingTier *string `json:"billingTier,omitempty" tf:"billing_tier,omitempty"`

	// Order Identification
	// +kubebuilder:validation:Optional
	OrderID *string `json:"orderId,omitempty" tf:"order_id,omitempty"`

	// Order Reference Number
	// +kubebuilder:validation:Optional
	OrderNumber *string `json:"orderNumber,omitempty" tf:"order_number,omitempty"`

	// Purchase order number
	// +kubebuilder:validation:Optional
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty" tf:"purchase_order_number,omitempty"`
}

type PackageObservation struct {
}

type PackageParameters struct {

	// Fabric Cloud Router package code
	// +kubebuilder:validation:Required
	Code *string `json:"code" tf:"code,omitempty"`
}

type ProjectObservation struct {
}

type ProjectParameters struct {

	// Unique Resource URL
	// +kubebuilder:validation:Optional
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// Project Id
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

// CloudRouterSpec defines the desired state of CloudRouter
type CloudRouterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudRouterParameters `json:"forProvider"`
}

// CloudRouterStatus defines the observed state of CloudRouter.
type CloudRouterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudRouterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudRouter is the Schema for the CloudRouters API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type CloudRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudRouterSpec   `json:"spec"`
	Status            CloudRouterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudRouterList contains a list of CloudRouters
type CloudRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudRouter `json:"items"`
}

// Repository type metadata.
var (
	CloudRouter_Kind             = "CloudRouter"
	CloudRouter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudRouter_Kind}.String()
	CloudRouter_KindAPIVersion   = CloudRouter_Kind + "." + CRDGroupVersion.String()
	CloudRouter_GroupVersionKind = CRDGroupVersion.WithKind(CloudRouter_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudRouter{}, &CloudRouterList{})
}
