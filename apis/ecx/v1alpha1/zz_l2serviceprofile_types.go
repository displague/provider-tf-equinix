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

type FeaturesInitParameters struct {

	// Indicates whether or not connections to this profile
	// can be created from remote metro locations.
	// Indicates whether or not connections to this profile can be created from remote metro locations
	AllowRemoteConnections *bool `json:"allowRemoteConnections,omitempty" tf:"allow_remote_connections,omitempty"`

	// (Deprecated) Indicates whether or not this profile can be used for test
	// connections.
	// Indicates whether or not this profile can be used for test connections
	TestProfile *bool `json:"testProfile,omitempty" tf:"test_profile,omitempty"`
}

type FeaturesObservation struct {

	// Indicates whether or not connections to this profile
	// can be created from remote metro locations.
	// Indicates whether or not connections to this profile can be created from remote metro locations
	AllowRemoteConnections *bool `json:"allowRemoteConnections,omitempty" tf:"allow_remote_connections,omitempty"`

	// (Deprecated) Indicates whether or not this profile can be used for test
	// connections.
	// Indicates whether or not this profile can be used for test connections
	TestProfile *bool `json:"testProfile,omitempty" tf:"test_profile,omitempty"`
}

type FeaturesParameters struct {

	// Indicates whether or not connections to this profile
	// can be created from remote metro locations.
	// Indicates whether or not connections to this profile can be created from remote metro locations
	// +kubebuilder:validation:Optional
	AllowRemoteConnections *bool `json:"allowRemoteConnections" tf:"allow_remote_connections,omitempty"`

	// (Deprecated) Indicates whether or not this profile can be used for test
	// connections.
	// Indicates whether or not this profile can be used for test connections
	// +kubebuilder:validation:Optional
	TestProfile *bool `json:"testProfile,omitempty" tf:"test_profile,omitempty"`
}

type L2ServiceprofileInitParameters struct {

	// Boolean value that determines if API integration is enabled. It
	// allows you to complete connection provisioning in less than five minutes. Without API Integration,
	// additional manual steps will be required and the provisioning will likely take longer.
	// Specifies the API integration ID that was provided to the customer during onboarding
	APIIntegration *bool `json:"apiIntegration,omitempty" tf:"api_integration,omitempty"`

	// Name of the authentication key label to be used by the
	// Authentication Key service. It allows Service Providers with QinQ ports to accept groups of
	// connections or VLANs from Dot1q customers. This is similar to S-Tag/C-Tag capabilities.
	// Name of the authentication key label to be used by the Authentication Key service
	AuthkeyLabel *string `json:"authkeyLabel,omitempty" tf:"authkey_label,omitempty"`

	// Specifies the port bandwidth threshold percentage. If
	// the bandwidth limit is met or exceeded, an alert is sent to the seller.
	// Specifies the port bandwidth threshold percentage. If the bandwidth limit is met or exceeded, an alert is sent to the seller
	BandwidthAlertThreshold *float64 `json:"bandwidthAlertThreshold,omitempty" tf:"bandwidth_alert_threshold,omitempty"`

	// A list of email addresses that will receive
	// notifications about bandwidth thresholds.
	// A list of email addresses that will receive notifications about bandwidth thresholds
	// +listType=set
	BandwidthThresholdNotifications []*string `json:"bandwidthThresholdNotifications,omitempty" tf:"bandwidth_threshold_notifications,omitempty"`

	// Custom name used for calling a connections
	// e.g. circuit. Defaults to Connection.
	// Custom name used for calling a connections i.e. circuit. Defaults to Connection
	ConnectionNameLabel *string `json:"connectionNameLabel,omitempty" tf:"connection_name_label,omitempty"`

	// C-Tag/Inner-Tag label name for the connections.
	// C-Tag/Inner-Tag label name for the connections
	CtagLabel *string `json:"ctagLabel,omitempty" tf:"ctag_label,omitempty"`

	// Description of the service profile.
	// Description of the service profile
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Applicable when api_integration is set to true. It
	// indicates whether the port and VLAN details are managed by Equinix.
	// Boolean value that indicates whether the port and VLAN details are managed by Equinix
	EquinixManagedPortVlan *bool `json:"equinixManagedPortVlan,omitempty" tf:"equinix_managed_port_vlan,omitempty"`

	// Block of profile features configuration. See Features below
	// for more details.
	// Block of profile features configuration
	Features []FeaturesInitParameters `json:"features,omitempty" tf:"features,omitempty"`

	// Specifies the API integration ID that was provided to the customer
	// during onboarding. You can validate your API integration ID using the validateIntegrationId API.
	// Specifies the API integration ID that was provided to the customer during onboarding
	IntegrationID *string `json:"integrationId,omitempty" tf:"integration_id,omitempty"`

	// Name of the service profile. An alpha-numeric 50 characters string which can
	// include only hyphens and underscores.
	// Name of the service profile. An alpha-numeric 50 characters string which can include only hyphens and underscores
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// You can set an alert for when a percentage of your profile has
	// been sold. Service providers like to use this functionality to alert them when they need to add
	// more ports or when they need to create a new service profile. Required with
	// oversubscription_allowed, defaults to 1x.
	// Oversubscription limit that will cause alerting. Default is 1x
	Oversubscription *string `json:"oversubscription,omitempty" tf:"oversubscription,omitempty"`

	// Boolean value that determines if, regardless of the
	// utilization, Equinix Fabric will continue to add connections to your links until we reach the
	// oversubscription limit. By selecting this service, you acknowledge that you will manage decisions
	// on when to increase capacity on these link.
	// Boolean value that determines if, regardless of the utilization, Equinix Fabric will continue to add connections to your links until we reach the oversubscription limit
	OversubscriptionAllowed *bool `json:"oversubscriptionAllowed,omitempty" tf:"oversubscription_allowed,omitempty"`

	// One or more definitions of ports residing in locations, from which your
	// customers will be able to access services using this service profile. See Port below for
	// more details.
	// One or more definitions of ports associated with the profile
	Port []PortInitParameters `json:"port,omitempty" tf:"port,omitempty"`

	// Boolean value that indicates whether or not this is a private profile,
	// i.e. not public like AWS/Azure/Oracle/Google, etc. If private, it can only be available for
	// creating connections if correct permissions are granted.
	// Boolean value that indicates whether or not this is a private profile.
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// An array of users email ids who have permission to access this
	// service profile. Argument is required when profile is set as private.
	// A list of email addresses associated to users that will be allowed to access this service profile. Applicable for private profiles
	// +listType=set
	PrivateUserEmails []*string `json:"privateUserEmails,omitempty" tf:"private_user_emails,omitempty"`

	// A list of email addresses that will receive
	// notifications about profile status changes.
	// A list of email addresses that will receive notifications about profile status changes
	// +listType=set
	ProfileStatuschangeNotifications []*string `json:"profileStatuschangeNotifications,omitempty" tf:"profile_statuschange_notifications,omitempty"`

	// Boolean value that determines if your connections will require
	// redundancy. if yes, then users need to create a secondary redundant connection.
	// Boolean value that determines if yourconnections will require redundancy
	RedundancyRequired *bool `json:"redundancyRequired,omitempty" tf:"redundancy_required,omitempty"`

	// Indicates whether the VLAN ID of. the secondary
	// connection is the same as the primary connection.
	// Indicates whether the VLAN ID of the secondary connection is the same as the primary connection
	SecondaryVlanFromPrimary *bool `json:"secondaryVlanFromPrimary,omitempty" tf:"secondary_vlan_from_primary,omitempty"`

	// Boolean value that indicates whether multiple connections
	// can be created with the same authorization key to connect to this service profile after the first
	// connection has been approved by the seller.
	// Boolean value that indicates whether multiple connections can be created with the same authorization key
	ServicekeyAutogenerated *bool `json:"servicekeyAutogenerated,omitempty" tf:"servicekey_autogenerated,omitempty"`

	// One or more definitions of supported speed/bandwidth. Argument is
	// required when speed_from_api is set to false. See Speed Band below for more
	// details.
	// One or more definitions of supported speed/bandwidth configurations
	SpeedBand []SpeedBandInitParameters `json:"speedBand,omitempty" tf:"speed_band,omitempty"`

	// Boolean value that determines if customer is allowed
	// to enter a custom connection speed.
	// Boolean value that determines if customer is allowed to enter a custom connection speed
	SpeedCustomizationAllowed *bool `json:"speedCustomizationAllowed,omitempty" tf:"speed_customization_allowed,omitempty"`

	// Boolean valuta that determines if connection speed will be derived
	// from an API call. Argument has to be specified when api_integration is enabled.
	// Boolean valuta that determines if connection speed will be derived from an API call
	SpeedFromAPI *bool `json:"speedFromApi,omitempty" tf:"speed_from_api,omitempty"`

	// Specifies additional tagging information required by the seller profile
	// for Dot1Q to QinQ translation. See Enhance Dot1q to QinQ translation support
	// for additional information. Valid values are:
	// Specifies additional tagging information required by the seller profile for Dot1Q to QinQ translation
	TagType *string `json:"tagType,omitempty" tf:"tag_type,omitempty"`

	// A list of email addresses that will receive
	// notifications about connections approvals and rejections.
	// A list of email addresses that will receive notifications about connections approvals and rejections
	// +listType=set
	VcStatuschangeNotifications []*string `json:"vcStatuschangeNotifications,omitempty" tf:"vc_statuschange_notifications,omitempty"`
}

type L2ServiceprofileObservation struct {

	// Boolean value that determines if API integration is enabled. It
	// allows you to complete connection provisioning in less than five minutes. Without API Integration,
	// additional manual steps will be required and the provisioning will likely take longer.
	// Specifies the API integration ID that was provided to the customer during onboarding
	APIIntegration *bool `json:"apiIntegration,omitempty" tf:"api_integration,omitempty"`

	// Name of the authentication key label to be used by the
	// Authentication Key service. It allows Service Providers with QinQ ports to accept groups of
	// connections or VLANs from Dot1q customers. This is similar to S-Tag/C-Tag capabilities.
	// Name of the authentication key label to be used by the Authentication Key service
	AuthkeyLabel *string `json:"authkeyLabel,omitempty" tf:"authkey_label,omitempty"`

	// Specifies the port bandwidth threshold percentage. If
	// the bandwidth limit is met or exceeded, an alert is sent to the seller.
	// Specifies the port bandwidth threshold percentage. If the bandwidth limit is met or exceeded, an alert is sent to the seller
	BandwidthAlertThreshold *float64 `json:"bandwidthAlertThreshold,omitempty" tf:"bandwidth_alert_threshold,omitempty"`

	// A list of email addresses that will receive
	// notifications about bandwidth thresholds.
	// A list of email addresses that will receive notifications about bandwidth thresholds
	// +listType=set
	BandwidthThresholdNotifications []*string `json:"bandwidthThresholdNotifications,omitempty" tf:"bandwidth_threshold_notifications,omitempty"`

	// Custom name used for calling a connections
	// e.g. circuit. Defaults to Connection.
	// Custom name used for calling a connections i.e. circuit. Defaults to Connection
	ConnectionNameLabel *string `json:"connectionNameLabel,omitempty" tf:"connection_name_label,omitempty"`

	// C-Tag/Inner-Tag label name for the connections.
	// C-Tag/Inner-Tag label name for the connections
	CtagLabel *string `json:"ctagLabel,omitempty" tf:"ctag_label,omitempty"`

	// Description of the service profile.
	// Description of the service profile
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Applicable when api_integration is set to true. It
	// indicates whether the port and VLAN details are managed by Equinix.
	// Boolean value that indicates whether the port and VLAN details are managed by Equinix
	EquinixManagedPortVlan *bool `json:"equinixManagedPortVlan,omitempty" tf:"equinix_managed_port_vlan,omitempty"`

	// Block of profile features configuration. See Features below
	// for more details.
	// Block of profile features configuration
	Features []FeaturesObservation `json:"features,omitempty" tf:"features,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the API integration ID that was provided to the customer
	// during onboarding. You can validate your API integration ID using the validateIntegrationId API.
	// Specifies the API integration ID that was provided to the customer during onboarding
	IntegrationID *string `json:"integrationId,omitempty" tf:"integration_id,omitempty"`

	// Name of the service profile. An alpha-numeric 50 characters string which can
	// include only hyphens and underscores.
	// Name of the service profile. An alpha-numeric 50 characters string which can include only hyphens and underscores
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// You can set an alert for when a percentage of your profile has
	// been sold. Service providers like to use this functionality to alert them when they need to add
	// more ports or when they need to create a new service profile. Required with
	// oversubscription_allowed, defaults to 1x.
	// Oversubscription limit that will cause alerting. Default is 1x
	Oversubscription *string `json:"oversubscription,omitempty" tf:"oversubscription,omitempty"`

	// Boolean value that determines if, regardless of the
	// utilization, Equinix Fabric will continue to add connections to your links until we reach the
	// oversubscription limit. By selecting this service, you acknowledge that you will manage decisions
	// on when to increase capacity on these link.
	// Boolean value that determines if, regardless of the utilization, Equinix Fabric will continue to add connections to your links until we reach the oversubscription limit
	OversubscriptionAllowed *bool `json:"oversubscriptionAllowed,omitempty" tf:"oversubscription_allowed,omitempty"`

	// One or more definitions of ports residing in locations, from which your
	// customers will be able to access services using this service profile. See Port below for
	// more details.
	// One or more definitions of ports associated with the profile
	Port []PortObservation `json:"port,omitempty" tf:"port,omitempty"`

	// Boolean value that indicates whether or not this is a private profile,
	// i.e. not public like AWS/Azure/Oracle/Google, etc. If private, it can only be available for
	// creating connections if correct permissions are granted.
	// Boolean value that indicates whether or not this is a private profile.
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// An array of users email ids who have permission to access this
	// service profile. Argument is required when profile is set as private.
	// A list of email addresses associated to users that will be allowed to access this service profile. Applicable for private profiles
	// +listType=set
	PrivateUserEmails []*string `json:"privateUserEmails,omitempty" tf:"private_user_emails,omitempty"`

	// A list of email addresses that will receive
	// notifications about profile status changes.
	// A list of email addresses that will receive notifications about profile status changes
	// +listType=set
	ProfileStatuschangeNotifications []*string `json:"profileStatuschangeNotifications,omitempty" tf:"profile_statuschange_notifications,omitempty"`

	// Boolean value that determines if your connections will require
	// redundancy. if yes, then users need to create a secondary redundant connection.
	// Boolean value that determines if yourconnections will require redundancy
	RedundancyRequired *bool `json:"redundancyRequired,omitempty" tf:"redundancy_required,omitempty"`

	// Indicates whether the VLAN ID of. the secondary
	// connection is the same as the primary connection.
	// Indicates whether the VLAN ID of the secondary connection is the same as the primary connection
	SecondaryVlanFromPrimary *bool `json:"secondaryVlanFromPrimary,omitempty" tf:"secondary_vlan_from_primary,omitempty"`

	// Boolean value that indicates whether multiple connections
	// can be created with the same authorization key to connect to this service profile after the first
	// connection has been approved by the seller.
	// Boolean value that indicates whether multiple connections can be created with the same authorization key
	ServicekeyAutogenerated *bool `json:"servicekeyAutogenerated,omitempty" tf:"servicekey_autogenerated,omitempty"`

	// One or more definitions of supported speed/bandwidth. Argument is
	// required when speed_from_api is set to false. See Speed Band below for more
	// details.
	// One or more definitions of supported speed/bandwidth configurations
	SpeedBand []SpeedBandObservation `json:"speedBand,omitempty" tf:"speed_band,omitempty"`

	// Boolean value that determines if customer is allowed
	// to enter a custom connection speed.
	// Boolean value that determines if customer is allowed to enter a custom connection speed
	SpeedCustomizationAllowed *bool `json:"speedCustomizationAllowed,omitempty" tf:"speed_customization_allowed,omitempty"`

	// Boolean valuta that determines if connection speed will be derived
	// from an API call. Argument has to be specified when api_integration is enabled.
	// Boolean valuta that determines if connection speed will be derived from an API call
	SpeedFromAPI *bool `json:"speedFromApi,omitempty" tf:"speed_from_api,omitempty"`

	// Service profile provisioning status.
	// Service profile provisioning status
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Specifies additional tagging information required by the seller profile
	// for Dot1Q to QinQ translation. See Enhance Dot1q to QinQ translation support
	// for additional information. Valid values are:
	// Specifies additional tagging information required by the seller profile for Dot1Q to QinQ translation
	TagType *string `json:"tagType,omitempty" tf:"tag_type,omitempty"`

	// Unique identifier of the service profile.
	// Unique identifier of the service profile
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// A list of email addresses that will receive
	// notifications about connections approvals and rejections.
	// A list of email addresses that will receive notifications about connections approvals and rejections
	// +listType=set
	VcStatuschangeNotifications []*string `json:"vcStatuschangeNotifications,omitempty" tf:"vc_statuschange_notifications,omitempty"`
}

type L2ServiceprofileParameters struct {

	// Boolean value that determines if API integration is enabled. It
	// allows you to complete connection provisioning in less than five minutes. Without API Integration,
	// additional manual steps will be required and the provisioning will likely take longer.
	// Specifies the API integration ID that was provided to the customer during onboarding
	// +kubebuilder:validation:Optional
	APIIntegration *bool `json:"apiIntegration,omitempty" tf:"api_integration,omitempty"`

	// Name of the authentication key label to be used by the
	// Authentication Key service. It allows Service Providers with QinQ ports to accept groups of
	// connections or VLANs from Dot1q customers. This is similar to S-Tag/C-Tag capabilities.
	// Name of the authentication key label to be used by the Authentication Key service
	// +kubebuilder:validation:Optional
	AuthkeyLabel *string `json:"authkeyLabel,omitempty" tf:"authkey_label,omitempty"`

	// Specifies the port bandwidth threshold percentage. If
	// the bandwidth limit is met or exceeded, an alert is sent to the seller.
	// Specifies the port bandwidth threshold percentage. If the bandwidth limit is met or exceeded, an alert is sent to the seller
	// +kubebuilder:validation:Optional
	BandwidthAlertThreshold *float64 `json:"bandwidthAlertThreshold,omitempty" tf:"bandwidth_alert_threshold,omitempty"`

	// A list of email addresses that will receive
	// notifications about bandwidth thresholds.
	// A list of email addresses that will receive notifications about bandwidth thresholds
	// +kubebuilder:validation:Optional
	// +listType=set
	BandwidthThresholdNotifications []*string `json:"bandwidthThresholdNotifications,omitempty" tf:"bandwidth_threshold_notifications,omitempty"`

	// Custom name used for calling a connections
	// e.g. circuit. Defaults to Connection.
	// Custom name used for calling a connections i.e. circuit. Defaults to Connection
	// +kubebuilder:validation:Optional
	ConnectionNameLabel *string `json:"connectionNameLabel,omitempty" tf:"connection_name_label,omitempty"`

	// C-Tag/Inner-Tag label name for the connections.
	// C-Tag/Inner-Tag label name for the connections
	// +kubebuilder:validation:Optional
	CtagLabel *string `json:"ctagLabel,omitempty" tf:"ctag_label,omitempty"`

	// Description of the service profile.
	// Description of the service profile
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Applicable when api_integration is set to true. It
	// indicates whether the port and VLAN details are managed by Equinix.
	// Boolean value that indicates whether the port and VLAN details are managed by Equinix
	// +kubebuilder:validation:Optional
	EquinixManagedPortVlan *bool `json:"equinixManagedPortVlan,omitempty" tf:"equinix_managed_port_vlan,omitempty"`

	// Block of profile features configuration. See Features below
	// for more details.
	// Block of profile features configuration
	// +kubebuilder:validation:Optional
	Features []FeaturesParameters `json:"features,omitempty" tf:"features,omitempty"`

	// Specifies the API integration ID that was provided to the customer
	// during onboarding. You can validate your API integration ID using the validateIntegrationId API.
	// Specifies the API integration ID that was provided to the customer during onboarding
	// +kubebuilder:validation:Optional
	IntegrationID *string `json:"integrationId,omitempty" tf:"integration_id,omitempty"`

	// Name of the service profile. An alpha-numeric 50 characters string which can
	// include only hyphens and underscores.
	// Name of the service profile. An alpha-numeric 50 characters string which can include only hyphens and underscores
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// You can set an alert for when a percentage of your profile has
	// been sold. Service providers like to use this functionality to alert them when they need to add
	// more ports or when they need to create a new service profile. Required with
	// oversubscription_allowed, defaults to 1x.
	// Oversubscription limit that will cause alerting. Default is 1x
	// +kubebuilder:validation:Optional
	Oversubscription *string `json:"oversubscription,omitempty" tf:"oversubscription,omitempty"`

	// Boolean value that determines if, regardless of the
	// utilization, Equinix Fabric will continue to add connections to your links until we reach the
	// oversubscription limit. By selecting this service, you acknowledge that you will manage decisions
	// on when to increase capacity on these link.
	// Boolean value that determines if, regardless of the utilization, Equinix Fabric will continue to add connections to your links until we reach the oversubscription limit
	// +kubebuilder:validation:Optional
	OversubscriptionAllowed *bool `json:"oversubscriptionAllowed,omitempty" tf:"oversubscription_allowed,omitempty"`

	// One or more definitions of ports residing in locations, from which your
	// customers will be able to access services using this service profile. See Port below for
	// more details.
	// One or more definitions of ports associated with the profile
	// +kubebuilder:validation:Optional
	Port []PortParameters `json:"port,omitempty" tf:"port,omitempty"`

	// Boolean value that indicates whether or not this is a private profile,
	// i.e. not public like AWS/Azure/Oracle/Google, etc. If private, it can only be available for
	// creating connections if correct permissions are granted.
	// Boolean value that indicates whether or not this is a private profile.
	// +kubebuilder:validation:Optional
	Private *bool `json:"private,omitempty" tf:"private,omitempty"`

	// An array of users email ids who have permission to access this
	// service profile. Argument is required when profile is set as private.
	// A list of email addresses associated to users that will be allowed to access this service profile. Applicable for private profiles
	// +kubebuilder:validation:Optional
	// +listType=set
	PrivateUserEmails []*string `json:"privateUserEmails,omitempty" tf:"private_user_emails,omitempty"`

	// A list of email addresses that will receive
	// notifications about profile status changes.
	// A list of email addresses that will receive notifications about profile status changes
	// +kubebuilder:validation:Optional
	// +listType=set
	ProfileStatuschangeNotifications []*string `json:"profileStatuschangeNotifications,omitempty" tf:"profile_statuschange_notifications,omitempty"`

	// Boolean value that determines if your connections will require
	// redundancy. if yes, then users need to create a secondary redundant connection.
	// Boolean value that determines if yourconnections will require redundancy
	// +kubebuilder:validation:Optional
	RedundancyRequired *bool `json:"redundancyRequired,omitempty" tf:"redundancy_required,omitempty"`

	// Indicates whether the VLAN ID of. the secondary
	// connection is the same as the primary connection.
	// Indicates whether the VLAN ID of the secondary connection is the same as the primary connection
	// +kubebuilder:validation:Optional
	SecondaryVlanFromPrimary *bool `json:"secondaryVlanFromPrimary,omitempty" tf:"secondary_vlan_from_primary,omitempty"`

	// Boolean value that indicates whether multiple connections
	// can be created with the same authorization key to connect to this service profile after the first
	// connection has been approved by the seller.
	// Boolean value that indicates whether multiple connections can be created with the same authorization key
	// +kubebuilder:validation:Optional
	ServicekeyAutogenerated *bool `json:"servicekeyAutogenerated,omitempty" tf:"servicekey_autogenerated,omitempty"`

	// One or more definitions of supported speed/bandwidth. Argument is
	// required when speed_from_api is set to false. See Speed Band below for more
	// details.
	// One or more definitions of supported speed/bandwidth configurations
	// +kubebuilder:validation:Optional
	SpeedBand []SpeedBandParameters `json:"speedBand,omitempty" tf:"speed_band,omitempty"`

	// Boolean value that determines if customer is allowed
	// to enter a custom connection speed.
	// Boolean value that determines if customer is allowed to enter a custom connection speed
	// +kubebuilder:validation:Optional
	SpeedCustomizationAllowed *bool `json:"speedCustomizationAllowed,omitempty" tf:"speed_customization_allowed,omitempty"`

	// Boolean valuta that determines if connection speed will be derived
	// from an API call. Argument has to be specified when api_integration is enabled.
	// Boolean valuta that determines if connection speed will be derived from an API call
	// +kubebuilder:validation:Optional
	SpeedFromAPI *bool `json:"speedFromApi,omitempty" tf:"speed_from_api,omitempty"`

	// Specifies additional tagging information required by the seller profile
	// for Dot1Q to QinQ translation. See Enhance Dot1q to QinQ translation support
	// for additional information. Valid values are:
	// Specifies additional tagging information required by the seller profile for Dot1Q to QinQ translation
	// +kubebuilder:validation:Optional
	TagType *string `json:"tagType,omitempty" tf:"tag_type,omitempty"`

	// A list of email addresses that will receive
	// notifications about connections approvals and rejections.
	// A list of email addresses that will receive notifications about connections approvals and rejections
	// +kubebuilder:validation:Optional
	// +listType=set
	VcStatuschangeNotifications []*string `json:"vcStatuschangeNotifications,omitempty" tf:"vc_statuschange_notifications,omitempty"`
}

type PortInitParameters struct {

	// The metro code of location where the port resides.
	// Port location metro code
	MetroCode *string `json:"metroCode,omitempty" tf:"metro_code,omitempty"`

	// Unique identifier of the port.
	// Unique identifier of the port
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type PortObservation struct {

	// The metro code of location where the port resides.
	// Port location metro code
	MetroCode *string `json:"metroCode,omitempty" tf:"metro_code,omitempty"`

	// Unique identifier of the port.
	// Unique identifier of the port
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type PortParameters struct {

	// The metro code of location where the port resides.
	// Port location metro code
	// +kubebuilder:validation:Optional
	MetroCode *string `json:"metroCode" tf:"metro_code,omitempty"`

	// Unique identifier of the port.
	// Unique identifier of the port
	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid" tf:"uuid,omitempty"`
}

type SpeedBandInitParameters struct {

	// Speed/bandwidth supported by this service profile.
	// Speed/bandwidth supported by given service profile
	Speed *float64 `json:"speed,omitempty" tf:"speed,omitempty"`

	// Unit of the speed/bandwidth supported by this service profile. One of
	// MB, GB.
	// Unit of the speed/bandwidth supported by given service profile
	SpeedUnit *string `json:"speedUnit,omitempty" tf:"speed_unit,omitempty"`
}

type SpeedBandObservation struct {

	// Speed/bandwidth supported by this service profile.
	// Speed/bandwidth supported by given service profile
	Speed *float64 `json:"speed,omitempty" tf:"speed,omitempty"`

	// Unit of the speed/bandwidth supported by this service profile. One of
	// MB, GB.
	// Unit of the speed/bandwidth supported by given service profile
	SpeedUnit *string `json:"speedUnit,omitempty" tf:"speed_unit,omitempty"`
}

type SpeedBandParameters struct {

	// Speed/bandwidth supported by this service profile.
	// Speed/bandwidth supported by given service profile
	// +kubebuilder:validation:Optional
	Speed *float64 `json:"speed" tf:"speed,omitempty"`

	// Unit of the speed/bandwidth supported by this service profile. One of
	// MB, GB.
	// Unit of the speed/bandwidth supported by given service profile
	// +kubebuilder:validation:Optional
	SpeedUnit *string `json:"speedUnit" tf:"speed_unit,omitempty"`
}

// L2ServiceprofileSpec defines the desired state of L2Serviceprofile
type L2ServiceprofileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     L2ServiceprofileParameters `json:"forProvider"`
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
	InitProvider L2ServiceprofileInitParameters `json:"initProvider,omitempty"`
}

// L2ServiceprofileStatus defines the observed state of L2Serviceprofile.
type L2ServiceprofileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        L2ServiceprofileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// L2Serviceprofile is the Schema for the L2Serviceprofiles API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type L2Serviceprofile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bandwidthThresholdNotifications) || (has(self.initProvider) && has(self.initProvider.bandwidthThresholdNotifications))",message="spec.forProvider.bandwidthThresholdNotifications is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.features) || (has(self.initProvider) && has(self.initProvider.features))",message="spec.forProvider.features is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.port) || (has(self.initProvider) && has(self.initProvider.port))",message="spec.forProvider.port is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.profileStatuschangeNotifications) || (has(self.initProvider) && has(self.initProvider.profileStatuschangeNotifications))",message="spec.forProvider.profileStatuschangeNotifications is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vcStatuschangeNotifications) || (has(self.initProvider) && has(self.initProvider.vcStatuschangeNotifications))",message="spec.forProvider.vcStatuschangeNotifications is a required parameter"
	Spec   L2ServiceprofileSpec   `json:"spec"`
	Status L2ServiceprofileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// L2ServiceprofileList contains a list of L2Serviceprofiles
type L2ServiceprofileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []L2Serviceprofile `json:"items"`
}

// Repository type metadata.
var (
	L2Serviceprofile_Kind             = "L2Serviceprofile"
	L2Serviceprofile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: L2Serviceprofile_Kind}.String()
	L2Serviceprofile_KindAPIVersion   = L2Serviceprofile_Kind + "." + CRDGroupVersion.String()
	L2Serviceprofile_GroupVersionKind = CRDGroupVersion.WithKind(L2Serviceprofile_Kind)
)

func init() {
	SchemeBuilder.Register(&L2Serviceprofile{}, &L2ServiceprofileList{})
}
