//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.ActiveNodes != nil {
		in, out := &in.ActiveNodes, &out.ActiveNodes
		*out = new(int64)
		**out = **in
	}
	if in.ClusterARN != nil {
		in, out := &in.ClusterARN, &out.ClusterARN
		*out = new(string)
		**out = **in
	}
	if in.ClusterDiscoveryEndpoint != nil {
		in, out := &in.ClusterDiscoveryEndpoint, &out.ClusterDiscoveryEndpoint
		*out = new(Endpoint)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleARN != nil {
		in, out := &in.IAMRoleARN, &out.IAMRoleARN
		*out = new(string)
		**out = **in
	}
	if in.NodeIDsToRemove != nil {
		in, out := &in.NodeIDsToRemove, &out.NodeIDsToRemove
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]*Node, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Node)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.NotificationConfiguration != nil {
		in, out := &in.NotificationConfiguration, &out.NotificationConfiguration
		*out = new(NotificationConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ParameterGroup != nil {
		in, out := &in.ParameterGroup, &out.ParameterGroup
		*out = new(ParameterGroupStatus_SDK)
		(*in).DeepCopyInto(*out)
	}
	if in.SSEDescription != nil {
		in, out := &in.SSEDescription, &out.SSEDescription
		*out = new(SSEDescription)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*SecurityGroupMembership, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SecurityGroupMembership)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SubnetGroup != nil {
		in, out := &in.SubnetGroup, &out.SubnetGroup
		*out = new(string)
		**out = **in
	}
	if in.TotalNodes != nil {
		in, out := &in.TotalNodes, &out.TotalNodes
		*out = new(int64)
		**out = **in
	}
	out.CustomClusterObservation = in.CustomClusterObservation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ClusterEndpointEncryptionType != nil {
		in, out := &in.ClusterEndpointEncryptionType, &out.ClusterEndpointEncryptionType
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.ReplicationFactor != nil {
		in, out := &in.ReplicationFactor, &out.ReplicationFactor
		*out = new(int64)
		**out = **in
	}
	if in.SSESpecification != nil {
		in, out := &in.SSESpecification, &out.SSESpecification
		*out = new(SSESpecification)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.CustomClusterParameters.DeepCopyInto(&out.CustomClusterParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster_SDK) DeepCopyInto(out *Cluster_SDK) {
	*out = *in
	if in.ActiveNodes != nil {
		in, out := &in.ActiveNodes, &out.ActiveNodes
		*out = new(int64)
		**out = **in
	}
	if in.ClusterARN != nil {
		in, out := &in.ClusterARN, &out.ClusterARN
		*out = new(string)
		**out = **in
	}
	if in.ClusterDiscoveryEndpoint != nil {
		in, out := &in.ClusterDiscoveryEndpoint, &out.ClusterDiscoveryEndpoint
		*out = new(Endpoint)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterEndpointEncryptionType != nil {
		in, out := &in.ClusterEndpointEncryptionType, &out.ClusterEndpointEncryptionType
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleARN != nil {
		in, out := &in.IAMRoleARN, &out.IAMRoleARN
		*out = new(string)
		**out = **in
	}
	if in.NodeIDsToRemove != nil {
		in, out := &in.NodeIDsToRemove, &out.NodeIDsToRemove
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]*Node, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Node)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.NotificationConfiguration != nil {
		in, out := &in.NotificationConfiguration, &out.NotificationConfiguration
		*out = new(NotificationConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ParameterGroup != nil {
		in, out := &in.ParameterGroup, &out.ParameterGroup
		*out = new(ParameterGroupStatus_SDK)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.SSEDescription != nil {
		in, out := &in.SSEDescription, &out.SSEDescription
		*out = new(SSEDescription)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*SecurityGroupMembership, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SecurityGroupMembership)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SubnetGroup != nil {
		in, out := &in.SubnetGroup, &out.SubnetGroup
		*out = new(string)
		**out = **in
	}
	if in.TotalNodes != nil {
		in, out := &in.TotalNodes, &out.TotalNodes
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster_SDK.
func (in *Cluster_SDK) DeepCopy() *Cluster_SDK {
	if in == nil {
		return nil
	}
	out := new(Cluster_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomClusterObservation) DeepCopyInto(out *CustomClusterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomClusterObservation.
func (in *CustomClusterObservation) DeepCopy() *CustomClusterObservation {
	if in == nil {
		return nil
	}
	out := new(CustomClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomClusterParameters) DeepCopyInto(out *CustomClusterParameters) {
	*out = *in
	if in.IAMRoleARN != nil {
		in, out := &in.IAMRoleARN, &out.IAMRoleARN
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleARNRef != nil {
		in, out := &in.IAMRoleARNRef, &out.IAMRoleARNRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IAMRoleARNSelector != nil {
		in, out := &in.IAMRoleARNSelector, &out.IAMRoleARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ParameterGroupName != nil {
		in, out := &in.ParameterGroupName, &out.ParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.ParameterGroupNameRef != nil {
		in, out := &in.ParameterGroupNameRef, &out.ParameterGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ParameterGroupNameSelector != nil {
		in, out := &in.ParameterGroupNameSelector, &out.ParameterGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetGroupName != nil {
		in, out := &in.SubnetGroupName, &out.SubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.SubnetGroupNameRef != nil {
		in, out := &in.SubnetGroupNameRef, &out.SubnetGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetGroupNameSelector != nil {
		in, out := &in.SubnetGroupNameSelector, &out.SubnetGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupIDRefs != nil {
		in, out := &in.SecurityGroupIDRefs, &out.SecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroupIDSelector != nil {
		in, out := &in.SecurityGroupIDSelector, &out.SecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.NotificationTopicARN != nil {
		in, out := &in.NotificationTopicARN, &out.NotificationTopicARN
		*out = new(string)
		**out = **in
	}
	if in.NotificationTopicARNRef != nil {
		in, out := &in.NotificationTopicARNRef, &out.NotificationTopicARNRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NotificationTopicARNSelector != nil {
		in, out := &in.NotificationTopicARNSelector, &out.NotificationTopicARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomClusterParameters.
func (in *CustomClusterParameters) DeepCopy() *CustomClusterParameters {
	if in == nil {
		return nil
	}
	out := new(CustomClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomParameterGroupObservation) DeepCopyInto(out *CustomParameterGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomParameterGroupObservation.
func (in *CustomParameterGroupObservation) DeepCopy() *CustomParameterGroupObservation {
	if in == nil {
		return nil
	}
	out := new(CustomParameterGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomParameterGroupParameters) DeepCopyInto(out *CustomParameterGroupParameters) {
	*out = *in
	if in.ParameterNameValues != nil {
		in, out := &in.ParameterNameValues, &out.ParameterNameValues
		*out = make([]*ParameterNameValue, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ParameterNameValue)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomParameterGroupParameters.
func (in *CustomParameterGroupParameters) DeepCopy() *CustomParameterGroupParameters {
	if in == nil {
		return nil
	}
	out := new(CustomParameterGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSubnetGroupObservation) DeepCopyInto(out *CustomSubnetGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSubnetGroupObservation.
func (in *CustomSubnetGroupObservation) DeepCopy() *CustomSubnetGroupObservation {
	if in == nil {
		return nil
	}
	out := new(CustomSubnetGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSubnetGroupParameters) DeepCopyInto(out *CustomSubnetGroupParameters) {
	*out = *in
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSubnetGroupParameters.
func (in *CustomSubnetGroupParameters) DeepCopy() *CustomSubnetGroupParameters {
	if in == nil {
		return nil
	}
	out := new(CustomSubnetGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Event) DeepCopyInto(out *Event) {
	*out = *in
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = (*in).DeepCopy()
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.SourceName != nil {
		in, out := &in.SourceName, &out.SourceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Event.
func (in *Event) DeepCopy() *Event {
	if in == nil {
		return nil
	}
	out := new(Event)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(Endpoint)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeCreateTime != nil {
		in, out := &in.NodeCreateTime, &out.NodeCreateTime
		*out = (*in).DeepCopy()
	}
	if in.NodeID != nil {
		in, out := &in.NodeID, &out.NodeID
		*out = new(string)
		**out = **in
	}
	if in.NodeStatus != nil {
		in, out := &in.NodeStatus, &out.NodeStatus
		*out = new(string)
		**out = **in
	}
	if in.ParameterGroupStatus != nil {
		in, out := &in.ParameterGroupStatus, &out.ParameterGroupStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTypeSpecificValue) DeepCopyInto(out *NodeTypeSpecificValue) {
	*out = *in
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTypeSpecificValue.
func (in *NodeTypeSpecificValue) DeepCopy() *NodeTypeSpecificValue {
	if in == nil {
		return nil
	}
	out := new(NodeTypeSpecificValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationConfiguration) DeepCopyInto(out *NotificationConfiguration) {
	*out = *in
	if in.TopicARN != nil {
		in, out := &in.TopicARN, &out.TopicARN
		*out = new(string)
		**out = **in
	}
	if in.TopicStatus != nil {
		in, out := &in.TopicStatus, &out.TopicStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationConfiguration.
func (in *NotificationConfiguration) DeepCopy() *NotificationConfiguration {
	if in == nil {
		return nil
	}
	out := new(NotificationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
	if in.AllowedValues != nil {
		in, out := &in.AllowedValues, &out.AllowedValues
		*out = new(string)
		**out = **in
	}
	if in.DataType != nil {
		in, out := &in.DataType, &out.DataType
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ParameterName != nil {
		in, out := &in.ParameterName, &out.ParameterName
		*out = new(string)
		**out = **in
	}
	if in.ParameterValue != nil {
		in, out := &in.ParameterValue, &out.ParameterValue
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroup) DeepCopyInto(out *ParameterGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroup.
func (in *ParameterGroup) DeepCopy() *ParameterGroup {
	if in == nil {
		return nil
	}
	out := new(ParameterGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParameterGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupList) DeepCopyInto(out *ParameterGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ParameterGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupList.
func (in *ParameterGroupList) DeepCopy() *ParameterGroupList {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParameterGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupObservation) DeepCopyInto(out *ParameterGroupObservation) {
	*out = *in
	if in.ParameterGroupName != nil {
		in, out := &in.ParameterGroupName, &out.ParameterGroupName
		*out = new(string)
		**out = **in
	}
	out.CustomParameterGroupObservation = in.CustomParameterGroupObservation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupObservation.
func (in *ParameterGroupObservation) DeepCopy() *ParameterGroupObservation {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupParameters) DeepCopyInto(out *ParameterGroupParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	in.CustomParameterGroupParameters.DeepCopyInto(&out.CustomParameterGroupParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupParameters.
func (in *ParameterGroupParameters) DeepCopy() *ParameterGroupParameters {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupSpec) DeepCopyInto(out *ParameterGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupSpec.
func (in *ParameterGroupSpec) DeepCopy() *ParameterGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupStatus) DeepCopyInto(out *ParameterGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupStatus.
func (in *ParameterGroupStatus) DeepCopy() *ParameterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupStatus_SDK) DeepCopyInto(out *ParameterGroupStatus_SDK) {
	*out = *in
	if in.NodeIDsToReboot != nil {
		in, out := &in.NodeIDsToReboot, &out.NodeIDsToReboot
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ParameterApplyStatus != nil {
		in, out := &in.ParameterApplyStatus, &out.ParameterApplyStatus
		*out = new(string)
		**out = **in
	}
	if in.ParameterGroupName != nil {
		in, out := &in.ParameterGroupName, &out.ParameterGroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupStatus_SDK.
func (in *ParameterGroupStatus_SDK) DeepCopy() *ParameterGroupStatus_SDK {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupStatus_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroup_SDK) DeepCopyInto(out *ParameterGroup_SDK) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ParameterGroupName != nil {
		in, out := &in.ParameterGroupName, &out.ParameterGroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroup_SDK.
func (in *ParameterGroup_SDK) DeepCopy() *ParameterGroup_SDK {
	if in == nil {
		return nil
	}
	out := new(ParameterGroup_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterNameValue) DeepCopyInto(out *ParameterNameValue) {
	*out = *in
	if in.ParameterName != nil {
		in, out := &in.ParameterName, &out.ParameterName
		*out = new(string)
		**out = **in
	}
	if in.ParameterValue != nil {
		in, out := &in.ParameterValue, &out.ParameterValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterNameValue.
func (in *ParameterNameValue) DeepCopy() *ParameterNameValue {
	if in == nil {
		return nil
	}
	out := new(ParameterNameValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSEDescription) DeepCopyInto(out *SSEDescription) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSEDescription.
func (in *SSEDescription) DeepCopy() *SSEDescription {
	if in == nil {
		return nil
	}
	out := new(SSEDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSESpecification) DeepCopyInto(out *SSESpecification) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSESpecification.
func (in *SSESpecification) DeepCopy() *SSESpecification {
	if in == nil {
		return nil
	}
	out := new(SSESpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupMembership) DeepCopyInto(out *SecurityGroupMembership) {
	*out = *in
	if in.SecurityGroupIdentifier != nil {
		in, out := &in.SecurityGroupIdentifier, &out.SecurityGroupIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupMembership.
func (in *SecurityGroupMembership) DeepCopy() *SecurityGroupMembership {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	if in.SubnetAvailabilityZone != nil {
		in, out := &in.SubnetAvailabilityZone, &out.SubnetAvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.SubnetIdentifier != nil {
		in, out := &in.SubnetIdentifier, &out.SubnetIdentifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroup) DeepCopyInto(out *SubnetGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroup.
func (in *SubnetGroup) DeepCopy() *SubnetGroup {
	if in == nil {
		return nil
	}
	out := new(SubnetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroupList) DeepCopyInto(out *SubnetGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SubnetGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroupList.
func (in *SubnetGroupList) DeepCopy() *SubnetGroupList {
	if in == nil {
		return nil
	}
	out := new(SubnetGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroupObservation) DeepCopyInto(out *SubnetGroupObservation) {
	*out = *in
	if in.SubnetGroupName != nil {
		in, out := &in.SubnetGroupName, &out.SubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]*Subnet, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Subnet)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	out.CustomSubnetGroupObservation = in.CustomSubnetGroupObservation
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroupObservation.
func (in *SubnetGroupObservation) DeepCopy() *SubnetGroupObservation {
	if in == nil {
		return nil
	}
	out := new(SubnetGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroupParameters) DeepCopyInto(out *SubnetGroupParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	in.CustomSubnetGroupParameters.DeepCopyInto(&out.CustomSubnetGroupParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroupParameters.
func (in *SubnetGroupParameters) DeepCopy() *SubnetGroupParameters {
	if in == nil {
		return nil
	}
	out := new(SubnetGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroupSpec) DeepCopyInto(out *SubnetGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroupSpec.
func (in *SubnetGroupSpec) DeepCopy() *SubnetGroupSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroupStatus) DeepCopyInto(out *SubnetGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroupStatus.
func (in *SubnetGroupStatus) DeepCopy() *SubnetGroupStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroup_SDK) DeepCopyInto(out *SubnetGroup_SDK) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.SubnetGroupName != nil {
		in, out := &in.SubnetGroupName, &out.SubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]*Subnet, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Subnet)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroup_SDK.
func (in *SubnetGroup_SDK) DeepCopy() *SubnetGroup_SDK {
	if in == nil {
		return nil
	}
	out := new(SubnetGroup_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}
