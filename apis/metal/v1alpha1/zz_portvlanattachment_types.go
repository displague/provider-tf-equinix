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

type PortVlanAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type PortVlanAttachmentParameters struct {

	// ID of device to be assigned to the VLAN
	// +kubebuilder:validation:Required
	DeviceID *string `json:"deviceId" tf:"device_id,omitempty"`

	// Add port back to the bond when this resource is removed. Default is false
	// +kubebuilder:validation:Optional
	ForceBond *bool `json:"forceBond,omitempty" tf:"force_bond,omitempty"`

	// Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use depends_on pointing to another equinix_metal_port_vlan_attachment, just like in the layer2-individual example above
	// +kubebuilder:validation:Optional
	Native *bool `json:"native,omitempty" tf:"native,omitempty"`

	// Name of network port to be assigned to the VLAN
	// +kubebuilder:validation:Required
	PortName *string `json:"portName" tf:"port_name,omitempty"`

	// VXLAN Network Identifier, integer
	// +kubebuilder:validation:Required
	VlanVnid *float64 `json:"vlanVnid" tf:"vlan_vnid,omitempty"`
}

// PortVlanAttachmentSpec defines the desired state of PortVlanAttachment
type PortVlanAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PortVlanAttachmentParameters `json:"forProvider"`
}

// PortVlanAttachmentStatus defines the observed state of PortVlanAttachment.
type PortVlanAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PortVlanAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PortVlanAttachment is the Schema for the PortVlanAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinixjet}
type PortVlanAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PortVlanAttachmentSpec   `json:"spec"`
	Status            PortVlanAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PortVlanAttachmentList contains a list of PortVlanAttachments
type PortVlanAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PortVlanAttachment `json:"items"`
}

// Repository type metadata.
var (
	PortVlanAttachment_Kind             = "PortVlanAttachment"
	PortVlanAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PortVlanAttachment_Kind}.String()
	PortVlanAttachment_KindAPIVersion   = PortVlanAttachment_Kind + "." + CRDGroupVersion.String()
	PortVlanAttachment_GroupVersionKind = CRDGroupVersion.WithKind(PortVlanAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&PortVlanAttachment{}, &PortVlanAttachmentList{})
}
