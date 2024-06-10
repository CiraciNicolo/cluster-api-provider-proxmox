//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalNetworkDevice) DeepCopyInto(out *AdditionalNetworkDevice) {
	*out = *in
	in.NetworkDevice.DeepCopyInto(&out.NetworkDevice)
	in.InterfaceConfig.DeepCopyInto(&out.InterfaceConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalNetworkDevice.
func (in *AdditionalNetworkDevice) DeepCopy() *AdditionalNetworkDevice {
	if in == nil {
		return nil
	}
	out := new(AdditionalNetworkDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerFlags) DeepCopyInto(out *ControllerFlags) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerFlags.
func (in *ControllerFlags) DeepCopy() *ControllerFlags {
	if in == nil {
		return nil
	}
	out := new(ControllerFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSize) DeepCopyInto(out *DiskSize) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSize.
func (in *DiskSize) DeepCopy() *DiskSize {
	if in == nil {
		return nil
	}
	out := new(DiskSize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddress) DeepCopyInto(out *IPAddress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddress.
func (in *IPAddress) DeepCopy() *IPAddress {
	if in == nil {
		return nil
	}
	out := new(IPAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPConfigSpec) DeepCopyInto(out *IPConfigSpec) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPConfigSpec.
func (in *IPConfigSpec) DeepCopy() *IPConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IPConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceConfig) DeepCopyInto(out *InterfaceConfig) {
	*out = *in
	if in.IPv4PoolRef != nil {
		in, out := &in.IPv4PoolRef, &out.IPv4PoolRef
		*out = new(v1.TypedLocalObjectReference)
		(*in).DeepCopyInto(*out)
	}
	if in.IPv6PoolRef != nil {
		in, out := &in.IPv6PoolRef, &out.IPv6PoolRef
		*out = new(v1.TypedLocalObjectReference)
		(*in).DeepCopyInto(*out)
	}
	if in.DNSServers != nil {
		in, out := &in.DNSServers, &out.DNSServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Routing.DeepCopyInto(&out.Routing)
	if in.LinkMTU != nil {
		in, out := &in.LinkMTU, &out.LinkMTU
		*out = new(uint16)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceConfig.
func (in *InterfaceConfig) DeepCopy() *InterfaceConfig {
	if in == nil {
		return nil
	}
	out := new(InterfaceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkDevice) DeepCopyInto(out *NetworkDevice) {
	*out = *in
	if in.Model != nil {
		in, out := &in.Model, &out.Model
		*out = new(string)
		**out = **in
	}
	if in.MTU != nil {
		in, out := &in.MTU, &out.MTU
		*out = new(uint16)
		**out = **in
	}
	if in.VLAN != nil {
		in, out := &in.VLAN, &out.VLAN
		*out = new(uint16)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkDevice.
func (in *NetworkDevice) DeepCopy() *NetworkDevice {
	if in == nil {
		return nil
	}
	out := new(NetworkDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(NetworkDevice)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalDevices != nil {
		in, out := &in.AdditionalDevices, &out.AdditionalDevices
		*out = make([]AdditionalNetworkDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.VirtualNetworkDevices.DeepCopyInto(&out.VirtualNetworkDevices)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	if in.IPAddrs != nil {
		in, out := &in.IPAddrs, &out.IPAddrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeLocation) DeepCopyInto(out *NodeLocation) {
	*out = *in
	out.Machine = in.Machine
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeLocation.
func (in *NodeLocation) DeepCopy() *NodeLocation {
	if in == nil {
		return nil
	}
	out := new(NodeLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeLocations) DeepCopyInto(out *NodeLocations) {
	*out = *in
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = make([]NodeLocation, len(*in))
		copy(*out, *in)
	}
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]NodeLocation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeLocations.
func (in *NodeLocations) DeepCopy() *NodeLocations {
	if in == nil {
		return nil
	}
	out := new(NodeLocations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxCluster) DeepCopyInto(out *ProxmoxCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxCluster.
func (in *ProxmoxCluster) DeepCopy() *ProxmoxCluster {
	if in == nil {
		return nil
	}
	out := new(ProxmoxCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterCloneSpec) DeepCopyInto(out *ProxmoxClusterCloneSpec) {
	*out = *in
	if in.ProxmoxMachineSpec != nil {
		in, out := &in.ProxmoxMachineSpec, &out.ProxmoxMachineSpec
		*out = make(map[string]ProxmoxMachineSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.SSHAuthorizedKeys != nil {
		in, out := &in.SSHAuthorizedKeys, &out.SSHAuthorizedKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterCloneSpec.
func (in *ProxmoxClusterCloneSpec) DeepCopy() *ProxmoxClusterCloneSpec {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterCloneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterList) DeepCopyInto(out *ProxmoxClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProxmoxCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterList.
func (in *ProxmoxClusterList) DeepCopy() *ProxmoxClusterList {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterSpec) DeepCopyInto(out *ProxmoxClusterSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	if in.AllowedNodes != nil {
		in, out := &in.AllowedNodes, &out.AllowedNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SchedulerHints != nil {
		in, out := &in.SchedulerHints, &out.SchedulerHints
		*out = new(SchedulerHints)
		(*in).DeepCopyInto(*out)
	}
	if in.IPv4Config != nil {
		in, out := &in.IPv4Config, &out.IPv4Config
		*out = new(IPConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.IPv6Config != nil {
		in, out := &in.IPv6Config, &out.IPv6Config
		*out = new(IPConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DNSServers != nil {
		in, out := &in.DNSServers, &out.DNSServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CloneSpec != nil {
		in, out := &in.CloneSpec, &out.CloneSpec
		*out = new(ProxmoxClusterCloneSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterSpec.
func (in *ProxmoxClusterSpec) DeepCopy() *ProxmoxClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterStatus) DeepCopyInto(out *ProxmoxClusterStatus) {
	*out = *in
	if in.InClusterIPPoolRef != nil {
		in, out := &in.InClusterIPPoolRef, &out.InClusterIPPoolRef
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.NodeLocations != nil {
		in, out := &in.NodeLocations, &out.NodeLocations
		*out = new(NodeLocations)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterStatus.
func (in *ProxmoxClusterStatus) DeepCopy() *ProxmoxClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterTemplate) DeepCopyInto(out *ProxmoxClusterTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterTemplate.
func (in *ProxmoxClusterTemplate) DeepCopy() *ProxmoxClusterTemplate {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxClusterTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterTemplateList) DeepCopyInto(out *ProxmoxClusterTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProxmoxClusterTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterTemplateList.
func (in *ProxmoxClusterTemplateList) DeepCopy() *ProxmoxClusterTemplateList {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxClusterTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterTemplateResource) DeepCopyInto(out *ProxmoxClusterTemplateResource) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterTemplateResource.
func (in *ProxmoxClusterTemplateResource) DeepCopy() *ProxmoxClusterTemplateResource {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterTemplateSpec) DeepCopyInto(out *ProxmoxClusterTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterTemplateSpec.
func (in *ProxmoxClusterTemplateSpec) DeepCopy() *ProxmoxClusterTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachine) DeepCopyInto(out *ProxmoxMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachine.
func (in *ProxmoxMachine) DeepCopy() *ProxmoxMachine {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachineList) DeepCopyInto(out *ProxmoxMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProxmoxMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachineList.
func (in *ProxmoxMachineList) DeepCopy() *ProxmoxMachineList {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachineSpec) DeepCopyInto(out *ProxmoxMachineSpec) {
	*out = *in
	in.VirtualMachineCloneSpec.DeepCopyInto(&out.VirtualMachineCloneSpec)
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineID != nil {
		in, out := &in.VirtualMachineID, &out.VirtualMachineID
		*out = new(int64)
		**out = **in
	}
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = new(Storage)
		(*in).DeepCopyInto(*out)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(NetworkSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachineSpec.
func (in *ProxmoxMachineSpec) DeepCopy() *ProxmoxMachineSpec {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachineStatus) DeepCopyInto(out *ProxmoxMachineStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1beta1.MachineAddress, len(*in))
		copy(*out, *in)
	}
	if in.BootstrapDataProvided != nil {
		in, out := &in.BootstrapDataProvided, &out.BootstrapDataProvided
		*out = new(bool)
		**out = **in
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make(map[string]IPAddress, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProxmoxNode != nil {
		in, out := &in.ProxmoxNode, &out.ProxmoxNode
		*out = new(string)
		**out = **in
	}
	if in.TaskRef != nil {
		in, out := &in.TaskRef, &out.TaskRef
		*out = new(string)
		**out = **in
	}
	in.RetryAfter.DeepCopyInto(&out.RetryAfter)
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachineStatus.
func (in *ProxmoxMachineStatus) DeepCopy() *ProxmoxMachineStatus {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachineTemplate) DeepCopyInto(out *ProxmoxMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachineTemplate.
func (in *ProxmoxMachineTemplate) DeepCopy() *ProxmoxMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachineTemplateList) DeepCopyInto(out *ProxmoxMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProxmoxMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachineTemplateList.
func (in *ProxmoxMachineTemplateList) DeepCopy() *ProxmoxMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachineTemplateResource) DeepCopyInto(out *ProxmoxMachineTemplateResource) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachineTemplateResource.
func (in *ProxmoxMachineTemplateResource) DeepCopy() *ProxmoxMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachineTemplateSpec) DeepCopyInto(out *ProxmoxMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachineTemplateSpec.
func (in *ProxmoxMachineTemplateSpec) DeepCopy() *ProxmoxMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Routing) DeepCopyInto(out *Routing) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RouteSpec, len(*in))
		copy(*out, *in)
	}
	if in.RoutingPolicy != nil {
		in, out := &in.RoutingPolicy, &out.RoutingPolicy
		*out = make([]RoutingPolicySpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Routing.
func (in *Routing) DeepCopy() *Routing {
	if in == nil {
		return nil
	}
	out := new(Routing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingPolicySpec) DeepCopyInto(out *RoutingPolicySpec) {
	*out = *in
	if in.Table != nil {
		in, out := &in.Table, &out.Table
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingPolicySpec.
func (in *RoutingPolicySpec) DeepCopy() *RoutingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RoutingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerHints) DeepCopyInto(out *SchedulerHints) {
	*out = *in
	if in.MemoryAdjustment != nil {
		in, out := &in.MemoryAdjustment, &out.MemoryAdjustment
		*out = new(uint64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerHints.
func (in *SchedulerHints) DeepCopy() *SchedulerHints {
	if in == nil {
		return nil
	}
	out := new(SchedulerHints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.BootVolume != nil {
		in, out := &in.BootVolume, &out.BootVolume
		*out = new(DiskSize)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VRFDevice) DeepCopyInto(out *VRFDevice) {
	*out = *in
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Routing.DeepCopyInto(&out.Routing)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VRFDevice.
func (in *VRFDevice) DeepCopy() *VRFDevice {
	if in == nil {
		return nil
	}
	out := new(VRFDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachine) DeepCopyInto(out *VirtualMachine) {
	*out = *in
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachine.
func (in *VirtualMachine) DeepCopy() *VirtualMachine {
	if in == nil {
		return nil
	}
	out := new(VirtualMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineCloneSpec) DeepCopyInto(out *VirtualMachineCloneSpec) {
	*out = *in
	if in.TemplateID != nil {
		in, out := &in.TemplateID, &out.TemplateID
		*out = new(int32)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(TargetFileStorageFormat)
		**out = **in
	}
	if in.Full != nil {
		in, out := &in.Full, &out.Full
		*out = new(bool)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(string)
		**out = **in
	}
	if in.SnapName != nil {
		in, out := &in.SnapName, &out.SnapName
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineCloneSpec.
func (in *VirtualMachineCloneSpec) DeepCopy() *VirtualMachineCloneSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineCloneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkDevices) DeepCopyInto(out *VirtualNetworkDevices) {
	*out = *in
	if in.VRFs != nil {
		in, out := &in.VRFs, &out.VRFs
		*out = make([]VRFDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkDevices.
func (in *VirtualNetworkDevices) DeepCopy() *VirtualNetworkDevices {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkDevices)
	in.DeepCopyInto(out)
	return out
}
