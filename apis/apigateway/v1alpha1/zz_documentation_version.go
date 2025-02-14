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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// DocumentationVersionParameters defines the desired state of DocumentationVersion
type DocumentationVersionParameters struct {
	// Region is which region the DocumentationVersion will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// A description about the new documentation snapshot.
	Description *string `json:"description,omitempty"`
	// The version identifier of the new snapshot.
	// +kubebuilder:validation:Required
	DocumentationVersion *string `json:"documentationVersion"`
	// The stage name to be associated with the new documentation snapshot.
	StageName                            *string `json:"stageName,omitempty"`
	CustomDocumentationVersionParameters `json:",inline"`
}

// DocumentationVersionSpec defines the desired state of DocumentationVersion
type DocumentationVersionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DocumentationVersionParameters `json:"forProvider"`
}

// DocumentationVersionObservation defines the observed state of DocumentationVersion
type DocumentationVersionObservation struct {
	// The date when the API documentation snapshot is created.
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`
	// The version identifier of the API documentation snapshot.
	Version *string `json:"version,omitempty"`

	CustomDocumentationVersionObservation `json:",inline"`
}

// DocumentationVersionStatus defines the observed state of DocumentationVersion.
type DocumentationVersionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DocumentationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationVersion is the Schema for the DocumentationVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DocumentationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocumentationVersionSpec   `json:"spec"`
	Status            DocumentationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationVersionList contains a list of DocumentationVersions
type DocumentationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocumentationVersion `json:"items"`
}

// Repository type metadata.
var (
	DocumentationVersionKind             = "DocumentationVersion"
	DocumentationVersionGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DocumentationVersionKind}.String()
	DocumentationVersionKindAPIVersion   = DocumentationVersionKind + "." + GroupVersion.String()
	DocumentationVersionGroupVersionKind = GroupVersion.WithKind(DocumentationVersionKind)
)

func init() {
	SchemeBuilder.Register(&DocumentationVersion{}, &DocumentationVersionList{})
}
