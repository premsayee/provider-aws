// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
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
func (in *NodeGroup) DeepCopyInto(out *NodeGroup) {
	*out = *in
	if in.NodeGroupMembers != nil {
		in, out := &in.NodeGroupMembers, &out.NodeGroupMembers
		*out = make([]NodeGroupMember, len(*in))
		copy(*out, *in)
	}
	out.PrimaryEndpoint = in.PrimaryEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroup.
func (in *NodeGroup) DeepCopy() *NodeGroup {
	if in == nil {
		return nil
	}
	out := new(NodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupConfigurationSpec) DeepCopyInto(out *NodeGroupConfigurationSpec) {
	*out = *in
	if in.PrimaryAvailabilityZone != nil {
		in, out := &in.PrimaryAvailabilityZone, &out.PrimaryAvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.ReplicaAvailabilityZones != nil {
		in, out := &in.ReplicaAvailabilityZones, &out.ReplicaAvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(int)
		**out = **in
	}
	if in.Slots != nil {
		in, out := &in.Slots, &out.Slots
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupConfigurationSpec.
func (in *NodeGroupConfigurationSpec) DeepCopy() *NodeGroupConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(NodeGroupConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupMember) DeepCopyInto(out *NodeGroupMember) {
	*out = *in
	out.ReadEndpoint = in.ReadEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupMember.
func (in *NodeGroupMember) DeepCopy() *NodeGroupMember {
	if in == nil {
		return nil
	}
	out := new(NodeGroupMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroup) DeepCopyInto(out *ReplicationGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroup.
func (in *ReplicationGroup) DeepCopy() *ReplicationGroup {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupClass) DeepCopyInto(out *ReplicationGroupClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupClass.
func (in *ReplicationGroupClass) DeepCopy() *ReplicationGroupClass {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationGroupClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupClassList) DeepCopyInto(out *ReplicationGroupClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationGroupClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupClassList.
func (in *ReplicationGroupClassList) DeepCopy() *ReplicationGroupClassList {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationGroupClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupClassSpecTemplate) DeepCopyInto(out *ReplicationGroupClassSpecTemplate) {
	*out = *in
	in.ClassSpecTemplate.DeepCopyInto(&out.ClassSpecTemplate)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupClassSpecTemplate.
func (in *ReplicationGroupClassSpecTemplate) DeepCopy() *ReplicationGroupClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupList) DeepCopyInto(out *ReplicationGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupList.
func (in *ReplicationGroupList) DeepCopy() *ReplicationGroupList {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupObservation) DeepCopyInto(out *ReplicationGroupObservation) {
	*out = *in
	out.ConfigurationEndpoint = in.ConfigurationEndpoint
	if in.MemberClusters != nil {
		in, out := &in.MemberClusters, &out.MemberClusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeGroups != nil {
		in, out := &in.NodeGroups, &out.NodeGroups
		*out = make([]NodeGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.PendingModifiedValues = in.PendingModifiedValues
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupObservation.
func (in *ReplicationGroupObservation) DeepCopy() *ReplicationGroupObservation {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupParameters) DeepCopyInto(out *ReplicationGroupParameters) {
	*out = *in
	if in.AtRestEncryptionEnabled != nil {
		in, out := &in.AtRestEncryptionEnabled, &out.AtRestEncryptionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthEnabled != nil {
		in, out := &in.AuthEnabled, &out.AuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AutomaticFailoverEnabled != nil {
		in, out := &in.AutomaticFailoverEnabled, &out.AutomaticFailoverEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CacheParameterGroupName != nil {
		in, out := &in.CacheParameterGroupName, &out.CacheParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.CacheSecurityGroupNames != nil {
		in, out := &in.CacheSecurityGroupNames, &out.CacheSecurityGroupNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CacheSecurityGroupNameRefs != nil {
		in, out := &in.CacheSecurityGroupNameRefs, &out.CacheSecurityGroupNameRefs
		*out = make([]*SecurityGroupNameReferencerForReplicationGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SecurityGroupNameReferencerForReplicationGroup)
				**out = **in
			}
		}
	}
	if in.CacheSubnetGroupName != nil {
		in, out := &in.CacheSubnetGroupName, &out.CacheSubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.NodeGroupConfiguration != nil {
		in, out := &in.NodeGroupConfiguration, &out.NodeGroupConfiguration
		*out = make([]NodeGroupConfigurationSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NotificationTopicARN != nil {
		in, out := &in.NotificationTopicARN, &out.NotificationTopicARN
		*out = new(string)
		**out = **in
	}
	if in.NotificationTopicStatus != nil {
		in, out := &in.NotificationTopicStatus, &out.NotificationTopicStatus
		*out = new(string)
		**out = **in
	}
	if in.NumCacheClusters != nil {
		in, out := &in.NumCacheClusters, &out.NumCacheClusters
		*out = new(int)
		**out = **in
	}
	if in.NumNodeGroups != nil {
		in, out := &in.NumNodeGroups, &out.NumNodeGroups
		*out = new(int)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.PreferredCacheClusterAZs != nil {
		in, out := &in.PreferredCacheClusterAZs, &out.PreferredCacheClusterAZs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.PrimaryClusterID != nil {
		in, out := &in.PrimaryClusterID, &out.PrimaryClusterID
		*out = new(string)
		**out = **in
	}
	if in.ReplicasPerNodeGroup != nil {
		in, out := &in.ReplicasPerNodeGroup, &out.ReplicasPerNodeGroup
		*out = new(int)
		**out = **in
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDRefs != nil {
		in, out := &in.SecurityGroupIDRefs, &out.SecurityGroupIDRefs
		*out = make([]*SecurityGroupIDReferencerForReplicationGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SecurityGroupIDReferencerForReplicationGroup)
				**out = **in
			}
		}
	}
	if in.SnapshotARNs != nil {
		in, out := &in.SnapshotARNs, &out.SnapshotARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotName != nil {
		in, out := &in.SnapshotName, &out.SnapshotName
		*out = new(string)
		**out = **in
	}
	if in.SnapshotRetentionLimit != nil {
		in, out := &in.SnapshotRetentionLimit, &out.SnapshotRetentionLimit
		*out = new(int)
		**out = **in
	}
	if in.SnapshotWindow != nil {
		in, out := &in.SnapshotWindow, &out.SnapshotWindow
		*out = new(string)
		**out = **in
	}
	if in.SnapshottingClusterID != nil {
		in, out := &in.SnapshottingClusterID, &out.SnapshottingClusterID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
	if in.TransitEncryptionEnabled != nil {
		in, out := &in.TransitEncryptionEnabled, &out.TransitEncryptionEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupParameters.
func (in *ReplicationGroupParameters) DeepCopy() *ReplicationGroupParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupPendingModifiedValues) DeepCopyInto(out *ReplicationGroupPendingModifiedValues) {
	*out = *in
	out.Resharding = in.Resharding
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupPendingModifiedValues.
func (in *ReplicationGroupPendingModifiedValues) DeepCopy() *ReplicationGroupPendingModifiedValues {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupPendingModifiedValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupSpec) DeepCopyInto(out *ReplicationGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupSpec.
func (in *ReplicationGroupSpec) DeepCopy() *ReplicationGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupStatus) DeepCopyInto(out *ReplicationGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupStatus.
func (in *ReplicationGroupStatus) DeepCopy() *ReplicationGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReshardingStatus) DeepCopyInto(out *ReshardingStatus) {
	*out = *in
	out.SlotMigration = in.SlotMigration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReshardingStatus.
func (in *ReshardingStatus) DeepCopy() *ReshardingStatus {
	if in == nil {
		return nil
	}
	out := new(ReshardingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupIDReferencerForReplicationGroup) DeepCopyInto(out *SecurityGroupIDReferencerForReplicationGroup) {
	*out = *in
	out.SecurityGroupIDReferencer = in.SecurityGroupIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupIDReferencerForReplicationGroup.
func (in *SecurityGroupIDReferencerForReplicationGroup) DeepCopy() *SecurityGroupIDReferencerForReplicationGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupIDReferencerForReplicationGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupNameReferencerForReplicationGroup) DeepCopyInto(out *SecurityGroupNameReferencerForReplicationGroup) {
	*out = *in
	out.SecurityGroupNameReferencer = in.SecurityGroupNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupNameReferencerForReplicationGroup.
func (in *SecurityGroupNameReferencerForReplicationGroup) DeepCopy() *SecurityGroupNameReferencerForReplicationGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupNameReferencerForReplicationGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlotMigration) DeepCopyInto(out *SlotMigration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlotMigration.
func (in *SlotMigration) DeepCopy() *SlotMigration {
	if in == nil {
		return nil
	}
	out := new(SlotMigration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
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
