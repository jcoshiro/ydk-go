// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-ma package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-network: IPv6 network operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_ma_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_ma_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-ma-oper ipv6-network}", reflect.TypeOf(Ipv6Network{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-ma-oper:ipv6-network", reflect.TypeOf(Ipv6Network{}))
}

// Ipv6MaOperState represents Interface oper states
type Ipv6MaOperState string

const (
    // Interface oper state is up
    Ipv6MaOperState_oper_up Ipv6MaOperState = "oper-up"

    // Interface oper state is down
    Ipv6MaOperState_oper_down Ipv6MaOperState = "oper-down"
)

// Ipv6MaIfAddrState represents Interface address states
type Ipv6MaIfAddrState string

const (
    // This is an active address that can appear as
    // the destination or source address of a packet
    Ipv6MaIfAddrState_active Ipv6MaIfAddrState = "active"

    // This is a valid but deprecated address that
    // should no longer be used as a source address in
    // new communications
    Ipv6MaIfAddrState_deprecated Ipv6MaIfAddrState = "deprecated"

    // This is a duplicate (invalid) address because
    // of conflict
    Ipv6MaIfAddrState_duplicate Ipv6MaIfAddrState = "duplicate"

    // This is not accessible because the interface to
    // which this address is assigned is not
    // operational
    Ipv6MaIfAddrState_inaccessible Ipv6MaIfAddrState = "inaccessible"

    // Status can not be determined for some reason
    Ipv6MaIfAddrState_tentative Ipv6MaIfAddrState = "tentative"
)

// Ipv6MaIfLineState represents Interface line states
type Ipv6MaIfLineState string

const (
    // Interface state is down
    Ipv6MaIfLineState_down Ipv6MaIfLineState = "down"

    // Interface state is up
    Ipv6MaIfLineState_up Ipv6MaIfLineState = "up"

    // Interface state is unknown
    Ipv6MaIfLineState_unknown Ipv6MaIfLineState = "unknown"

    // Interface state is incorrect
    Ipv6MaIfLineState_error Ipv6MaIfLineState = "error"
)

// Ipv6Network
// IPv6 network operational data
type Ipv6Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific IPv6 network operational data.
    Nodes Ipv6Network_Nodes
}

func (ipv6Network *Ipv6Network) GetEntityData() *types.CommonEntityData {
    ipv6Network.EntityData.YFilter = ipv6Network.YFilter
    ipv6Network.EntityData.YangName = "ipv6-network"
    ipv6Network.EntityData.BundleName = "cisco_ios_xr"
    ipv6Network.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-ma-oper"
    ipv6Network.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ma-oper:ipv6-network"
    ipv6Network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Network.EntityData.Children = make(map[string]types.YChild)
    ipv6Network.EntityData.Children["nodes"] = types.YChild{"Nodes", &ipv6Network.Nodes}
    ipv6Network.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6Network.EntityData)
}

// Ipv6Network_Nodes
// Node-specific IPv6 network operational data
type Ipv6Network_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 network operational data for a particular node. The type is slice of
    // Ipv6Network_Nodes_Node.
    Node []Ipv6Network_Nodes_Node
}

func (nodes *Ipv6Network_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ipv6-network"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Ipv6Network_Nodes_Node
// IPv6 network operational data for a particular
// node
type Ipv6Network_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // IPv6 network operational interface data.
    InterfaceData Ipv6Network_Nodes_Node_InterfaceData
}

func (node *Ipv6Network_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["interface-data"] = types.YChild{"InterfaceData", &node.InterfaceData}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData
// IPv6 network operational interface data
type Ipv6Network_Nodes_Node_InterfaceData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF specific IPv6 network operational interface data.
    Vrfs Ipv6Network_Nodes_Node_InterfaceData_Vrfs

    // Summary of IPv6 network operational interface data on a node.
    Summary Ipv6Network_Nodes_Node_InterfaceData_Summary
}

func (interfaceData *Ipv6Network_Nodes_Node_InterfaceData) GetEntityData() *types.CommonEntityData {
    interfaceData.EntityData.YFilter = interfaceData.YFilter
    interfaceData.EntityData.YangName = "interface-data"
    interfaceData.EntityData.BundleName = "cisco_ios_xr"
    interfaceData.EntityData.ParentYangName = "node"
    interfaceData.EntityData.SegmentPath = "interface-data"
    interfaceData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceData.EntityData.Children = make(map[string]types.YChild)
    interfaceData.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &interfaceData.Vrfs}
    interfaceData.EntityData.Children["summary"] = types.YChild{"Summary", &interfaceData.Summary}
    interfaceData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceData.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs
// VRF specific IPv6 network operational
// interface data
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF ID of an interface belong to. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf.
    Vrf []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf
}

func (vrfs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "interface-data"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf
// VRF ID of an interface belong to
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Brief interface IPv6 network operational data for a node.
    Briefs Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs

    // Detail interface IPv4 network operational data for global data.
    GlobalDetails Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails

    // Brief interface IPv6 network operational data from global data.
    GlobalBriefs Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs

    // Detail interface IPv4 network operational data for a node.
    Details Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details
}

func (vrf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["briefs"] = types.YChild{"Briefs", &vrf.Briefs}
    vrf.EntityData.Children["global-details"] = types.YChild{"GlobalDetails", &vrf.GlobalDetails}
    vrf.EntityData.Children["global-briefs"] = types.YChild{"GlobalBriefs", &vrf.GlobalBriefs}
    vrf.EntityData.Children["details"] = types.YChild{"Details", &vrf.Details}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs
// Brief interface IPv6 network operational
// data for a node
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief interface IPv6 network operational data for an interface. The type is
    // slice of Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief.
    Brief []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief
}

func (briefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs) GetEntityData() *types.CommonEntityData {
    briefs.EntityData.YFilter = briefs.YFilter
    briefs.EntityData.YangName = "briefs"
    briefs.EntityData.BundleName = "cisco_ios_xr"
    briefs.EntityData.ParentYangName = "vrf"
    briefs.EntityData.SegmentPath = "briefs"
    briefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefs.EntityData.Children = make(map[string]types.YChild)
    briefs.EntityData.Children["brief"] = types.YChild{"Brief", nil}
    for i := range briefs.Brief {
        briefs.EntityData.Children[types.GetSegmentPath(&briefs.Brief[i])] = types.YChild{"Brief", &briefs.Brief[i]}
    }
    briefs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(briefs.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief
// Brief interface IPv6 network operational
// data for an interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // State of Interface Line. The type is Ipv6MaIfLineState.
    LineState interface{}

    // VRF Name. The type is string with length: 0..32.
    VrfName interface{}

    // Link Local Address.
    LinkLocalAddress Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress

    // Address List. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address.
    Address []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address
}

func (brief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "briefs"
    brief.EntityData.SegmentPath = "brief" + "[interface-name='" + fmt.Sprintf("%v", brief.InterfaceName) + "']"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Children["link-local-address"] = types.YChild{"LinkLocalAddress", &brief.LinkLocalAddress}
    brief.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range brief.Address {
        brief.EntityData.Children[types.GetSegmentPath(&brief.Address[i])] = types.YChild{"Address", &brief.Address[i]}
    }
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    brief.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", brief.InterfaceName}
    brief.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", brief.LineState}
    brief.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", brief.VrfName}
    return &(brief.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress
// Link Local Address
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_LinkLocalAddress) GetEntityData() *types.CommonEntityData {
    linkLocalAddress.EntityData.YFilter = linkLocalAddress.YFilter
    linkLocalAddress.EntityData.YangName = "link-local-address"
    linkLocalAddress.EntityData.BundleName = "cisco_ios_xr"
    linkLocalAddress.EntityData.ParentYangName = "brief"
    linkLocalAddress.EntityData.SegmentPath = "link-local-address"
    linkLocalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkLocalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkLocalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkLocalAddress.EntityData.Children = make(map[string]types.YChild)
    linkLocalAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    linkLocalAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", linkLocalAddress.Address}
    linkLocalAddress.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", linkLocalAddress.PrefixLength}
    linkLocalAddress.EntityData.Leafs["address-state"] = types.YLeaf{"AddressState", linkLocalAddress.AddressState}
    linkLocalAddress.EntityData.Leafs["is-anycast"] = types.YLeaf{"IsAnycast", linkLocalAddress.IsAnycast}
    linkLocalAddress.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", linkLocalAddress.RouteTag}
    return &(linkLocalAddress.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address
// Address List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Briefs_Brief_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "brief"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address"] = types.YLeaf{"Address", address.Address}
    address.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", address.PrefixLength}
    address.EntityData.Leafs["address-state"] = types.YLeaf{"AddressState", address.AddressState}
    address.EntityData.Leafs["is-anycast"] = types.YLeaf{"IsAnycast", address.IsAnycast}
    address.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", address.RouteTag}
    return &(address.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails
// Detail interface IPv4 network operational
// data for global data
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detail interface IPv6 network operational data for an interface. The type
    // is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail.
    GlobalDetail []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail
}

func (globalDetails *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails) GetEntityData() *types.CommonEntityData {
    globalDetails.EntityData.YFilter = globalDetails.YFilter
    globalDetails.EntityData.YangName = "global-details"
    globalDetails.EntityData.BundleName = "cisco_ios_xr"
    globalDetails.EntityData.ParentYangName = "vrf"
    globalDetails.EntityData.SegmentPath = "global-details"
    globalDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalDetails.EntityData.Children = make(map[string]types.YChild)
    globalDetails.EntityData.Children["global-detail"] = types.YChild{"GlobalDetail", nil}
    for i := range globalDetails.GlobalDetail {
        globalDetails.EntityData.Children[types.GetSegmentPath(&globalDetails.GlobalDetail[i])] = types.YChild{"GlobalDetail", &globalDetails.GlobalDetail[i]}
    }
    globalDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(globalDetails.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail
// Detail interface IPv6 network operational
// data for an interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // State of Interface Line. The type is Ipv6MaIfLineState.
    LineState interface{}

    // IPv6 MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // IPv6 Operation State. The type is Ipv6MaOperState.
    OperationState interface{}

    // VRF Name. The type is string with length: 0..32.
    VrfName interface{}

    // ICMP unreach Enable. The type is bool.
    IsIcmpUnreachEnabled interface{}

    // Does ICCP RG ID exist on the interface?. The type is bool.
    RgIdExists interface{}

    // Is mLACP state Active (valid if RG ID exists). The type is bool.
    MlacpActive interface{}

    // Is BGP Flow Tag Source is enable. The type is bool.
    FlowTagSrc interface{}

    // Is BGP Flow Tag Destination is enable. The type is bool.
    FlowTagDst interface{}

    // Link Local Address.
    LinkLocalAddress Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress

    // IPv6 Access Control List.
    AccessControlList Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList

    // Multi IPv6 Access Control List.
    MultiAccessControlList Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList

    // RPF config on the interface.
    Rpf Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf

    // BGP PA config on the interface.
    BgpPa Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa

    // Address Publish Time.
    Utime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime

    // IDB Create Time.
    IdbUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime

    // CAPS Add Time.
    CapsUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime

    // FWD ENABLE Time.
    FwdEnUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime

    // FWD DISABLE Time.
    FwdDisUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime

    // IPv6 Multicast Group. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup.
    MulticastGroup []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup

    // Address List. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address.
    Address []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address

    // IPv6 Client Multicast Group. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup.
    ClientMulticastGroup []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup
}

func (globalDetail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail) GetEntityData() *types.CommonEntityData {
    globalDetail.EntityData.YFilter = globalDetail.YFilter
    globalDetail.EntityData.YangName = "global-detail"
    globalDetail.EntityData.BundleName = "cisco_ios_xr"
    globalDetail.EntityData.ParentYangName = "global-details"
    globalDetail.EntityData.SegmentPath = "global-detail" + "[interface-name='" + fmt.Sprintf("%v", globalDetail.InterfaceName) + "']"
    globalDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalDetail.EntityData.Children = make(map[string]types.YChild)
    globalDetail.EntityData.Children["link-local-address"] = types.YChild{"LinkLocalAddress", &globalDetail.LinkLocalAddress}
    globalDetail.EntityData.Children["access-control-list"] = types.YChild{"AccessControlList", &globalDetail.AccessControlList}
    globalDetail.EntityData.Children["multi-access-control-list"] = types.YChild{"MultiAccessControlList", &globalDetail.MultiAccessControlList}
    globalDetail.EntityData.Children["rpf"] = types.YChild{"Rpf", &globalDetail.Rpf}
    globalDetail.EntityData.Children["bgp-pa"] = types.YChild{"BgpPa", &globalDetail.BgpPa}
    globalDetail.EntityData.Children["utime"] = types.YChild{"Utime", &globalDetail.Utime}
    globalDetail.EntityData.Children["idb-utime"] = types.YChild{"IdbUtime", &globalDetail.IdbUtime}
    globalDetail.EntityData.Children["caps-utime"] = types.YChild{"CapsUtime", &globalDetail.CapsUtime}
    globalDetail.EntityData.Children["fwd-en-utime"] = types.YChild{"FwdEnUtime", &globalDetail.FwdEnUtime}
    globalDetail.EntityData.Children["fwd-dis-utime"] = types.YChild{"FwdDisUtime", &globalDetail.FwdDisUtime}
    globalDetail.EntityData.Children["multicast-group"] = types.YChild{"MulticastGroup", nil}
    for i := range globalDetail.MulticastGroup {
        globalDetail.EntityData.Children[types.GetSegmentPath(&globalDetail.MulticastGroup[i])] = types.YChild{"MulticastGroup", &globalDetail.MulticastGroup[i]}
    }
    globalDetail.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range globalDetail.Address {
        globalDetail.EntityData.Children[types.GetSegmentPath(&globalDetail.Address[i])] = types.YChild{"Address", &globalDetail.Address[i]}
    }
    globalDetail.EntityData.Children["client-multicast-group"] = types.YChild{"ClientMulticastGroup", nil}
    for i := range globalDetail.ClientMulticastGroup {
        globalDetail.EntityData.Children[types.GetSegmentPath(&globalDetail.ClientMulticastGroup[i])] = types.YChild{"ClientMulticastGroup", &globalDetail.ClientMulticastGroup[i]}
    }
    globalDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    globalDetail.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", globalDetail.InterfaceName}
    globalDetail.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", globalDetail.LineState}
    globalDetail.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", globalDetail.Mtu}
    globalDetail.EntityData.Leafs["operation-state"] = types.YLeaf{"OperationState", globalDetail.OperationState}
    globalDetail.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", globalDetail.VrfName}
    globalDetail.EntityData.Leafs["is-icmp-unreach-enabled"] = types.YLeaf{"IsIcmpUnreachEnabled", globalDetail.IsIcmpUnreachEnabled}
    globalDetail.EntityData.Leafs["rg-id-exists"] = types.YLeaf{"RgIdExists", globalDetail.RgIdExists}
    globalDetail.EntityData.Leafs["mlacp-active"] = types.YLeaf{"MlacpActive", globalDetail.MlacpActive}
    globalDetail.EntityData.Leafs["flow-tag-src"] = types.YLeaf{"FlowTagSrc", globalDetail.FlowTagSrc}
    globalDetail.EntityData.Leafs["flow-tag-dst"] = types.YLeaf{"FlowTagDst", globalDetail.FlowTagDst}
    return &(globalDetail.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress
// Link Local Address
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_LinkLocalAddress) GetEntityData() *types.CommonEntityData {
    linkLocalAddress.EntityData.YFilter = linkLocalAddress.YFilter
    linkLocalAddress.EntityData.YangName = "link-local-address"
    linkLocalAddress.EntityData.BundleName = "cisco_ios_xr"
    linkLocalAddress.EntityData.ParentYangName = "global-detail"
    linkLocalAddress.EntityData.SegmentPath = "link-local-address"
    linkLocalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkLocalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkLocalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkLocalAddress.EntityData.Children = make(map[string]types.YChild)
    linkLocalAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    linkLocalAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", linkLocalAddress.Address}
    linkLocalAddress.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", linkLocalAddress.PrefixLength}
    linkLocalAddress.EntityData.Leafs["address-state"] = types.YLeaf{"AddressState", linkLocalAddress.AddressState}
    linkLocalAddress.EntityData.Leafs["is-anycast"] = types.YLeaf{"IsAnycast", linkLocalAddress.IsAnycast}
    linkLocalAddress.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", linkLocalAddress.RouteTag}
    return &(linkLocalAddress.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList
// IPv6 Access Control List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL applied to incoming packets. The type is string.
    InBound interface{}

    // ACL applied to outgoing packets. The type is string.
    OutBound interface{}

    // Common ACL applied to incoming packets. The type is string.
    CommonInBound interface{}

    // Common ACL applied to outgoing packets. The type is string.
    CommonOutBound interface{}
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_AccessControlList) GetEntityData() *types.CommonEntityData {
    accessControlList.EntityData.YFilter = accessControlList.YFilter
    accessControlList.EntityData.YangName = "access-control-list"
    accessControlList.EntityData.BundleName = "cisco_ios_xr"
    accessControlList.EntityData.ParentYangName = "global-detail"
    accessControlList.EntityData.SegmentPath = "access-control-list"
    accessControlList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessControlList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessControlList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessControlList.EntityData.Children = make(map[string]types.YChild)
    accessControlList.EntityData.Leafs = make(map[string]types.YLeaf)
    accessControlList.EntityData.Leafs["in-bound"] = types.YLeaf{"InBound", accessControlList.InBound}
    accessControlList.EntityData.Leafs["out-bound"] = types.YLeaf{"OutBound", accessControlList.OutBound}
    accessControlList.EntityData.Leafs["common-in-bound"] = types.YLeaf{"CommonInBound", accessControlList.CommonInBound}
    accessControlList.EntityData.Leafs["common-out-bound"] = types.YLeaf{"CommonOutBound", accessControlList.CommonOutBound}
    return &(accessControlList.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList
// Multi IPv6 Access Control List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inbound ACLs. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Inbound.
    Inbound []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Inbound

    // Outbound ACLs. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Outbound.
    Outbound []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Outbound

    // Common ACLs. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Common.
    Common []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Common
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList) GetEntityData() *types.CommonEntityData {
    multiAccessControlList.EntityData.YFilter = multiAccessControlList.YFilter
    multiAccessControlList.EntityData.YangName = "multi-access-control-list"
    multiAccessControlList.EntityData.BundleName = "cisco_ios_xr"
    multiAccessControlList.EntityData.ParentYangName = "global-detail"
    multiAccessControlList.EntityData.SegmentPath = "multi-access-control-list"
    multiAccessControlList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multiAccessControlList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multiAccessControlList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multiAccessControlList.EntityData.Children = make(map[string]types.YChild)
    multiAccessControlList.EntityData.Children["inbound"] = types.YChild{"Inbound", nil}
    for i := range multiAccessControlList.Inbound {
        multiAccessControlList.EntityData.Children[types.GetSegmentPath(&multiAccessControlList.Inbound[i])] = types.YChild{"Inbound", &multiAccessControlList.Inbound[i]}
    }
    multiAccessControlList.EntityData.Children["outbound"] = types.YChild{"Outbound", nil}
    for i := range multiAccessControlList.Outbound {
        multiAccessControlList.EntityData.Children[types.GetSegmentPath(&multiAccessControlList.Outbound[i])] = types.YChild{"Outbound", &multiAccessControlList.Outbound[i]}
    }
    multiAccessControlList.EntityData.Children["common"] = types.YChild{"Common", nil}
    for i := range multiAccessControlList.Common {
        multiAccessControlList.EntityData.Children[types.GetSegmentPath(&multiAccessControlList.Common[i])] = types.YChild{"Common", &multiAccessControlList.Common[i]}
    }
    multiAccessControlList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(multiAccessControlList.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Inbound
// Inbound ACLs
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Inbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (inbound *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Inbound) GetEntityData() *types.CommonEntityData {
    inbound.EntityData.YFilter = inbound.YFilter
    inbound.EntityData.YangName = "inbound"
    inbound.EntityData.BundleName = "cisco_ios_xr"
    inbound.EntityData.ParentYangName = "multi-access-control-list"
    inbound.EntityData.SegmentPath = "inbound"
    inbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inbound.EntityData.Children = make(map[string]types.YChild)
    inbound.EntityData.Leafs = make(map[string]types.YLeaf)
    inbound.EntityData.Leafs["entry"] = types.YLeaf{"Entry", inbound.Entry}
    return &(inbound.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Outbound
// Outbound ACLs
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Outbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (outbound *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Outbound) GetEntityData() *types.CommonEntityData {
    outbound.EntityData.YFilter = outbound.YFilter
    outbound.EntityData.YangName = "outbound"
    outbound.EntityData.BundleName = "cisco_ios_xr"
    outbound.EntityData.ParentYangName = "multi-access-control-list"
    outbound.EntityData.SegmentPath = "outbound"
    outbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outbound.EntityData.Children = make(map[string]types.YChild)
    outbound.EntityData.Leafs = make(map[string]types.YLeaf)
    outbound.EntityData.Leafs["entry"] = types.YLeaf{"Entry", outbound.Entry}
    return &(outbound.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Common
// Common ACLs
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Common struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (common *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MultiAccessControlList_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xr"
    common.EntityData.ParentYangName = "multi-access-control-list"
    common.EntityData.SegmentPath = "common"
    common.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    common.EntityData.Children = make(map[string]types.YChild)
    common.EntityData.Leafs = make(map[string]types.YLeaf)
    common.EntityData.Leafs["entry"] = types.YLeaf{"Entry", common.Entry}
    return &(common.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf
// RPF config on the interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable RPF config. The type is bool.
    Enable interface{}

    // Allow Default Route. The type is bool.
    AllowDefaultRoute interface{}

    // Allow Self Ping. The type is bool.
    AllowSelfPing interface{}

    // RPF Mode (loose/strict). The type is interface{} with range: 0..4294967295.
    Mode interface{}
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Rpf) GetEntityData() *types.CommonEntityData {
    rpf.EntityData.YFilter = rpf.YFilter
    rpf.EntityData.YangName = "rpf"
    rpf.EntityData.BundleName = "cisco_ios_xr"
    rpf.EntityData.ParentYangName = "global-detail"
    rpf.EntityData.SegmentPath = "rpf"
    rpf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpf.EntityData.Children = make(map[string]types.YChild)
    rpf.EntityData.Leafs = make(map[string]types.YLeaf)
    rpf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", rpf.Enable}
    rpf.EntityData.Leafs["allow-default-route"] = types.YLeaf{"AllowDefaultRoute", rpf.AllowDefaultRoute}
    rpf.EntityData.Leafs["allow-self-ping"] = types.YLeaf{"AllowSelfPing", rpf.AllowSelfPing}
    rpf.EntityData.Leafs["mode"] = types.YLeaf{"Mode", rpf.Mode}
    return &(rpf.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa
// BGP PA config on the interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP PA input config.
    Input Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input

    // BGP PA output config.
    Output Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa) GetEntityData() *types.CommonEntityData {
    bgpPa.EntityData.YFilter = bgpPa.YFilter
    bgpPa.EntityData.YangName = "bgp-pa"
    bgpPa.EntityData.BundleName = "cisco_ios_xr"
    bgpPa.EntityData.ParentYangName = "global-detail"
    bgpPa.EntityData.SegmentPath = "bgp-pa"
    bgpPa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpPa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpPa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpPa.EntityData.Children = make(map[string]types.YChild)
    bgpPa.EntityData.Children["input"] = types.YChild{"Input", &bgpPa.Input}
    bgpPa.EntityData.Children["output"] = types.YChild{"Output", &bgpPa.Output}
    bgpPa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpPa.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input
// BGP PA input config
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is interface{} with range:
    // 0..4294967295.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "bgp-pa"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["enable"] = types.YLeaf{"Enable", input.Enable}
    input.EntityData.Leafs["source"] = types.YLeaf{"Source", input.Source}
    input.EntityData.Leafs["destination"] = types.YLeaf{"Destination", input.Destination}
    return &(input.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output
// BGP PA output config
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is interface{} with range:
    // 0..4294967295.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_BgpPa_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "bgp-pa"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["enable"] = types.YLeaf{"Enable", output.Enable}
    output.EntityData.Leafs["source"] = types.YLeaf{"Source", output.Source}
    output.EntityData.Leafs["destination"] = types.YLeaf{"Destination", output.Destination}
    return &(output.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime
// Address Publish Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Utime) GetEntityData() *types.CommonEntityData {
    utime.EntityData.YFilter = utime.YFilter
    utime.EntityData.YangName = "utime"
    utime.EntityData.BundleName = "cisco_ios_xr"
    utime.EntityData.ParentYangName = "global-detail"
    utime.EntityData.SegmentPath = "utime"
    utime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utime.EntityData.Children = make(map[string]types.YChild)
    utime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(utime.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime
// IDB Create Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_IdbUtime) GetEntityData() *types.CommonEntityData {
    idbUtime.EntityData.YFilter = idbUtime.YFilter
    idbUtime.EntityData.YangName = "idb-utime"
    idbUtime.EntityData.BundleName = "cisco_ios_xr"
    idbUtime.EntityData.ParentYangName = "global-detail"
    idbUtime.EntityData.SegmentPath = "idb-utime"
    idbUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idbUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idbUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idbUtime.EntityData.Children = make(map[string]types.YChild)
    idbUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(idbUtime.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime
// CAPS Add Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_CapsUtime) GetEntityData() *types.CommonEntityData {
    capsUtime.EntityData.YFilter = capsUtime.YFilter
    capsUtime.EntityData.YangName = "caps-utime"
    capsUtime.EntityData.BundleName = "cisco_ios_xr"
    capsUtime.EntityData.ParentYangName = "global-detail"
    capsUtime.EntityData.SegmentPath = "caps-utime"
    capsUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capsUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capsUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capsUtime.EntityData.Children = make(map[string]types.YChild)
    capsUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(capsUtime.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime
// FWD ENABLE Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdEnUtime) GetEntityData() *types.CommonEntityData {
    fwdEnUtime.EntityData.YFilter = fwdEnUtime.YFilter
    fwdEnUtime.EntityData.YangName = "fwd-en-utime"
    fwdEnUtime.EntityData.BundleName = "cisco_ios_xr"
    fwdEnUtime.EntityData.ParentYangName = "global-detail"
    fwdEnUtime.EntityData.SegmentPath = "fwd-en-utime"
    fwdEnUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwdEnUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwdEnUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwdEnUtime.EntityData.Children = make(map[string]types.YChild)
    fwdEnUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fwdEnUtime.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime
// FWD DISABLE Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_FwdDisUtime) GetEntityData() *types.CommonEntityData {
    fwdDisUtime.EntityData.YFilter = fwdDisUtime.YFilter
    fwdDisUtime.EntityData.YangName = "fwd-dis-utime"
    fwdDisUtime.EntityData.BundleName = "cisco_ios_xr"
    fwdDisUtime.EntityData.ParentYangName = "global-detail"
    fwdDisUtime.EntityData.SegmentPath = "fwd-dis-utime"
    fwdDisUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwdDisUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwdDisUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwdDisUtime.EntityData.Children = make(map[string]types.YChild)
    fwdDisUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fwdDisUtime.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup
// IPv6 Multicast Group
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address of Multicast Group. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_MulticastGroup) GetEntityData() *types.CommonEntityData {
    multicastGroup.EntityData.YFilter = multicastGroup.YFilter
    multicastGroup.EntityData.YangName = "multicast-group"
    multicastGroup.EntityData.BundleName = "cisco_ios_xr"
    multicastGroup.EntityData.ParentYangName = "global-detail"
    multicastGroup.EntityData.SegmentPath = "multicast-group"
    multicastGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastGroup.EntityData.Children = make(map[string]types.YChild)
    multicastGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    multicastGroup.EntityData.Leafs["address"] = types.YLeaf{"Address", multicastGroup.Address}
    return &(multicastGroup.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address
// Address List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "global-detail"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address"] = types.YLeaf{"Address", address.Address}
    address.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", address.PrefixLength}
    address.EntityData.Leafs["address-state"] = types.YLeaf{"AddressState", address.AddressState}
    address.EntityData.Leafs["is-anycast"] = types.YLeaf{"IsAnycast", address.IsAnycast}
    address.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", address.RouteTag}
    return &(address.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup
// IPv6 Client Multicast Group
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address of Multicast Group. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalDetails_GlobalDetail_ClientMulticastGroup) GetEntityData() *types.CommonEntityData {
    clientMulticastGroup.EntityData.YFilter = clientMulticastGroup.YFilter
    clientMulticastGroup.EntityData.YangName = "client-multicast-group"
    clientMulticastGroup.EntityData.BundleName = "cisco_ios_xr"
    clientMulticastGroup.EntityData.ParentYangName = "global-detail"
    clientMulticastGroup.EntityData.SegmentPath = "client-multicast-group"
    clientMulticastGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientMulticastGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientMulticastGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientMulticastGroup.EntityData.Children = make(map[string]types.YChild)
    clientMulticastGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    clientMulticastGroup.EntityData.Leafs["address"] = types.YLeaf{"Address", clientMulticastGroup.Address}
    return &(clientMulticastGroup.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs
// Brief interface IPv6 network operational
// data from global data
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief interface IPv6 network operational data for an interface. The type is
    // slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief.
    GlobalBrief []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief
}

func (globalBriefs *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs) GetEntityData() *types.CommonEntityData {
    globalBriefs.EntityData.YFilter = globalBriefs.YFilter
    globalBriefs.EntityData.YangName = "global-briefs"
    globalBriefs.EntityData.BundleName = "cisco_ios_xr"
    globalBriefs.EntityData.ParentYangName = "vrf"
    globalBriefs.EntityData.SegmentPath = "global-briefs"
    globalBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalBriefs.EntityData.Children = make(map[string]types.YChild)
    globalBriefs.EntityData.Children["global-brief"] = types.YChild{"GlobalBrief", nil}
    for i := range globalBriefs.GlobalBrief {
        globalBriefs.EntityData.Children[types.GetSegmentPath(&globalBriefs.GlobalBrief[i])] = types.YChild{"GlobalBrief", &globalBriefs.GlobalBrief[i]}
    }
    globalBriefs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(globalBriefs.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief
// Brief interface IPv6 network operational
// data for an interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // State of Interface Line. The type is Ipv6MaIfLineState.
    LineState interface{}

    // VRF Name. The type is string with length: 0..32.
    VrfName interface{}

    // Link Local Address.
    LinkLocalAddress Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress

    // Address List. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address.
    Address []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address
}

func (globalBrief *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief) GetEntityData() *types.CommonEntityData {
    globalBrief.EntityData.YFilter = globalBrief.YFilter
    globalBrief.EntityData.YangName = "global-brief"
    globalBrief.EntityData.BundleName = "cisco_ios_xr"
    globalBrief.EntityData.ParentYangName = "global-briefs"
    globalBrief.EntityData.SegmentPath = "global-brief" + "[interface-name='" + fmt.Sprintf("%v", globalBrief.InterfaceName) + "']"
    globalBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalBrief.EntityData.Children = make(map[string]types.YChild)
    globalBrief.EntityData.Children["link-local-address"] = types.YChild{"LinkLocalAddress", &globalBrief.LinkLocalAddress}
    globalBrief.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range globalBrief.Address {
        globalBrief.EntityData.Children[types.GetSegmentPath(&globalBrief.Address[i])] = types.YChild{"Address", &globalBrief.Address[i]}
    }
    globalBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    globalBrief.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", globalBrief.InterfaceName}
    globalBrief.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", globalBrief.LineState}
    globalBrief.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", globalBrief.VrfName}
    return &(globalBrief.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress
// Link Local Address
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_LinkLocalAddress) GetEntityData() *types.CommonEntityData {
    linkLocalAddress.EntityData.YFilter = linkLocalAddress.YFilter
    linkLocalAddress.EntityData.YangName = "link-local-address"
    linkLocalAddress.EntityData.BundleName = "cisco_ios_xr"
    linkLocalAddress.EntityData.ParentYangName = "global-brief"
    linkLocalAddress.EntityData.SegmentPath = "link-local-address"
    linkLocalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkLocalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkLocalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkLocalAddress.EntityData.Children = make(map[string]types.YChild)
    linkLocalAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    linkLocalAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", linkLocalAddress.Address}
    linkLocalAddress.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", linkLocalAddress.PrefixLength}
    linkLocalAddress.EntityData.Leafs["address-state"] = types.YLeaf{"AddressState", linkLocalAddress.AddressState}
    linkLocalAddress.EntityData.Leafs["is-anycast"] = types.YLeaf{"IsAnycast", linkLocalAddress.IsAnycast}
    linkLocalAddress.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", linkLocalAddress.RouteTag}
    return &(linkLocalAddress.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address
// Address List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_GlobalBriefs_GlobalBrief_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "global-brief"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address"] = types.YLeaf{"Address", address.Address}
    address.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", address.PrefixLength}
    address.EntityData.Leafs["address-state"] = types.YLeaf{"AddressState", address.AddressState}
    address.EntityData.Leafs["is-anycast"] = types.YLeaf{"IsAnycast", address.IsAnycast}
    address.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", address.RouteTag}
    return &(address.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details
// Detail interface IPv4 network operational
// data for a node
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detail interface IPv6 network operational data for an interface. The type
    // is slice of Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail.
    Detail []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail
}

func (details *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "vrf"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["detail"] = types.YChild{"Detail", nil}
    for i := range details.Detail {
        details.EntityData.Children[types.GetSegmentPath(&details.Detail[i])] = types.YChild{"Detail", &details.Detail[i]}
    }
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail
// Detail interface IPv6 network operational
// data for an interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // State of Interface Line. The type is Ipv6MaIfLineState.
    LineState interface{}

    // IPv6 MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // IPv6 Operation State. The type is Ipv6MaOperState.
    OperationState interface{}

    // VRF Name. The type is string with length: 0..32.
    VrfName interface{}

    // ICMP unreach Enable. The type is bool.
    IsIcmpUnreachEnabled interface{}

    // Does ICCP RG ID exist on the interface?. The type is bool.
    RgIdExists interface{}

    // Is mLACP state Active (valid if RG ID exists). The type is bool.
    MlacpActive interface{}

    // Is BGP Flow Tag Source is enable. The type is bool.
    FlowTagSrc interface{}

    // Is BGP Flow Tag Destination is enable. The type is bool.
    FlowTagDst interface{}

    // Link Local Address.
    LinkLocalAddress Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress

    // IPv6 Access Control List.
    AccessControlList Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList

    // Multi IPv6 Access Control List.
    MultiAccessControlList Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList

    // RPF config on the interface.
    Rpf Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf

    // BGP PA config on the interface.
    BgpPa Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa

    // Address Publish Time.
    Utime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime

    // IDB Create Time.
    IdbUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime

    // CAPS Add Time.
    CapsUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime

    // FWD ENABLE Time.
    FwdEnUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime

    // FWD DISABLE Time.
    FwdDisUtime Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime

    // IPv6 Multicast Group. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup.
    MulticastGroup []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup

    // Address List. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address.
    Address []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address

    // IPv6 Client Multicast Group. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup.
    ClientMulticastGroup []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup
}

func (detail *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "details"
    detail.EntityData.SegmentPath = "detail" + "[interface-name='" + fmt.Sprintf("%v", detail.InterfaceName) + "']"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["link-local-address"] = types.YChild{"LinkLocalAddress", &detail.LinkLocalAddress}
    detail.EntityData.Children["access-control-list"] = types.YChild{"AccessControlList", &detail.AccessControlList}
    detail.EntityData.Children["multi-access-control-list"] = types.YChild{"MultiAccessControlList", &detail.MultiAccessControlList}
    detail.EntityData.Children["rpf"] = types.YChild{"Rpf", &detail.Rpf}
    detail.EntityData.Children["bgp-pa"] = types.YChild{"BgpPa", &detail.BgpPa}
    detail.EntityData.Children["utime"] = types.YChild{"Utime", &detail.Utime}
    detail.EntityData.Children["idb-utime"] = types.YChild{"IdbUtime", &detail.IdbUtime}
    detail.EntityData.Children["caps-utime"] = types.YChild{"CapsUtime", &detail.CapsUtime}
    detail.EntityData.Children["fwd-en-utime"] = types.YChild{"FwdEnUtime", &detail.FwdEnUtime}
    detail.EntityData.Children["fwd-dis-utime"] = types.YChild{"FwdDisUtime", &detail.FwdDisUtime}
    detail.EntityData.Children["multicast-group"] = types.YChild{"MulticastGroup", nil}
    for i := range detail.MulticastGroup {
        detail.EntityData.Children[types.GetSegmentPath(&detail.MulticastGroup[i])] = types.YChild{"MulticastGroup", &detail.MulticastGroup[i]}
    }
    detail.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range detail.Address {
        detail.EntityData.Children[types.GetSegmentPath(&detail.Address[i])] = types.YChild{"Address", &detail.Address[i]}
    }
    detail.EntityData.Children["client-multicast-group"] = types.YChild{"ClientMulticastGroup", nil}
    for i := range detail.ClientMulticastGroup {
        detail.EntityData.Children[types.GetSegmentPath(&detail.ClientMulticastGroup[i])] = types.YChild{"ClientMulticastGroup", &detail.ClientMulticastGroup[i]}
    }
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", detail.InterfaceName}
    detail.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", detail.LineState}
    detail.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", detail.Mtu}
    detail.EntityData.Leafs["operation-state"] = types.YLeaf{"OperationState", detail.OperationState}
    detail.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", detail.VrfName}
    detail.EntityData.Leafs["is-icmp-unreach-enabled"] = types.YLeaf{"IsIcmpUnreachEnabled", detail.IsIcmpUnreachEnabled}
    detail.EntityData.Leafs["rg-id-exists"] = types.YLeaf{"RgIdExists", detail.RgIdExists}
    detail.EntityData.Leafs["mlacp-active"] = types.YLeaf{"MlacpActive", detail.MlacpActive}
    detail.EntityData.Leafs["flow-tag-src"] = types.YLeaf{"FlowTagSrc", detail.FlowTagSrc}
    detail.EntityData.Leafs["flow-tag-dst"] = types.YLeaf{"FlowTagDst", detail.FlowTagDst}
    return &(detail.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress
// Link Local Address
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (linkLocalAddress *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_LinkLocalAddress) GetEntityData() *types.CommonEntityData {
    linkLocalAddress.EntityData.YFilter = linkLocalAddress.YFilter
    linkLocalAddress.EntityData.YangName = "link-local-address"
    linkLocalAddress.EntityData.BundleName = "cisco_ios_xr"
    linkLocalAddress.EntityData.ParentYangName = "detail"
    linkLocalAddress.EntityData.SegmentPath = "link-local-address"
    linkLocalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkLocalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkLocalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkLocalAddress.EntityData.Children = make(map[string]types.YChild)
    linkLocalAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    linkLocalAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", linkLocalAddress.Address}
    linkLocalAddress.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", linkLocalAddress.PrefixLength}
    linkLocalAddress.EntityData.Leafs["address-state"] = types.YLeaf{"AddressState", linkLocalAddress.AddressState}
    linkLocalAddress.EntityData.Leafs["is-anycast"] = types.YLeaf{"IsAnycast", linkLocalAddress.IsAnycast}
    linkLocalAddress.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", linkLocalAddress.RouteTag}
    return &(linkLocalAddress.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList
// IPv6 Access Control List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL applied to incoming packets. The type is string.
    InBound interface{}

    // ACL applied to outgoing packets. The type is string.
    OutBound interface{}

    // Common ACL applied to incoming packets. The type is string.
    CommonInBound interface{}

    // Common ACL applied to outgoing packets. The type is string.
    CommonOutBound interface{}
}

func (accessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_AccessControlList) GetEntityData() *types.CommonEntityData {
    accessControlList.EntityData.YFilter = accessControlList.YFilter
    accessControlList.EntityData.YangName = "access-control-list"
    accessControlList.EntityData.BundleName = "cisco_ios_xr"
    accessControlList.EntityData.ParentYangName = "detail"
    accessControlList.EntityData.SegmentPath = "access-control-list"
    accessControlList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessControlList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessControlList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessControlList.EntityData.Children = make(map[string]types.YChild)
    accessControlList.EntityData.Leafs = make(map[string]types.YLeaf)
    accessControlList.EntityData.Leafs["in-bound"] = types.YLeaf{"InBound", accessControlList.InBound}
    accessControlList.EntityData.Leafs["out-bound"] = types.YLeaf{"OutBound", accessControlList.OutBound}
    accessControlList.EntityData.Leafs["common-in-bound"] = types.YLeaf{"CommonInBound", accessControlList.CommonInBound}
    accessControlList.EntityData.Leafs["common-out-bound"] = types.YLeaf{"CommonOutBound", accessControlList.CommonOutBound}
    return &(accessControlList.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList
// Multi IPv6 Access Control List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inbound ACLs. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Inbound.
    Inbound []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Inbound

    // Outbound ACLs. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Outbound.
    Outbound []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Outbound

    // Common ACLs. The type is slice of
    // Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Common.
    Common []Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Common
}

func (multiAccessControlList *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList) GetEntityData() *types.CommonEntityData {
    multiAccessControlList.EntityData.YFilter = multiAccessControlList.YFilter
    multiAccessControlList.EntityData.YangName = "multi-access-control-list"
    multiAccessControlList.EntityData.BundleName = "cisco_ios_xr"
    multiAccessControlList.EntityData.ParentYangName = "detail"
    multiAccessControlList.EntityData.SegmentPath = "multi-access-control-list"
    multiAccessControlList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multiAccessControlList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multiAccessControlList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multiAccessControlList.EntityData.Children = make(map[string]types.YChild)
    multiAccessControlList.EntityData.Children["inbound"] = types.YChild{"Inbound", nil}
    for i := range multiAccessControlList.Inbound {
        multiAccessControlList.EntityData.Children[types.GetSegmentPath(&multiAccessControlList.Inbound[i])] = types.YChild{"Inbound", &multiAccessControlList.Inbound[i]}
    }
    multiAccessControlList.EntityData.Children["outbound"] = types.YChild{"Outbound", nil}
    for i := range multiAccessControlList.Outbound {
        multiAccessControlList.EntityData.Children[types.GetSegmentPath(&multiAccessControlList.Outbound[i])] = types.YChild{"Outbound", &multiAccessControlList.Outbound[i]}
    }
    multiAccessControlList.EntityData.Children["common"] = types.YChild{"Common", nil}
    for i := range multiAccessControlList.Common {
        multiAccessControlList.EntityData.Children[types.GetSegmentPath(&multiAccessControlList.Common[i])] = types.YChild{"Common", &multiAccessControlList.Common[i]}
    }
    multiAccessControlList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(multiAccessControlList.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Inbound
// Inbound ACLs
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Inbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (inbound *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Inbound) GetEntityData() *types.CommonEntityData {
    inbound.EntityData.YFilter = inbound.YFilter
    inbound.EntityData.YangName = "inbound"
    inbound.EntityData.BundleName = "cisco_ios_xr"
    inbound.EntityData.ParentYangName = "multi-access-control-list"
    inbound.EntityData.SegmentPath = "inbound"
    inbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inbound.EntityData.Children = make(map[string]types.YChild)
    inbound.EntityData.Leafs = make(map[string]types.YLeaf)
    inbound.EntityData.Leafs["entry"] = types.YLeaf{"Entry", inbound.Entry}
    return &(inbound.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Outbound
// Outbound ACLs
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Outbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (outbound *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Outbound) GetEntityData() *types.CommonEntityData {
    outbound.EntityData.YFilter = outbound.YFilter
    outbound.EntityData.YangName = "outbound"
    outbound.EntityData.BundleName = "cisco_ios_xr"
    outbound.EntityData.ParentYangName = "multi-access-control-list"
    outbound.EntityData.SegmentPath = "outbound"
    outbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outbound.EntityData.Children = make(map[string]types.YChild)
    outbound.EntityData.Leafs = make(map[string]types.YLeaf)
    outbound.EntityData.Leafs["entry"] = types.YLeaf{"Entry", outbound.Entry}
    return &(outbound.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Common
// Common ACLs
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Common struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Entry interface{}
}

func (common *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MultiAccessControlList_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xr"
    common.EntityData.ParentYangName = "multi-access-control-list"
    common.EntityData.SegmentPath = "common"
    common.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    common.EntityData.Children = make(map[string]types.YChild)
    common.EntityData.Leafs = make(map[string]types.YLeaf)
    common.EntityData.Leafs["entry"] = types.YLeaf{"Entry", common.Entry}
    return &(common.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf
// RPF config on the interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable RPF config. The type is bool.
    Enable interface{}

    // Allow Default Route. The type is bool.
    AllowDefaultRoute interface{}

    // Allow Self Ping. The type is bool.
    AllowSelfPing interface{}

    // RPF Mode (loose/strict). The type is interface{} with range: 0..4294967295.
    Mode interface{}
}

func (rpf *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Rpf) GetEntityData() *types.CommonEntityData {
    rpf.EntityData.YFilter = rpf.YFilter
    rpf.EntityData.YangName = "rpf"
    rpf.EntityData.BundleName = "cisco_ios_xr"
    rpf.EntityData.ParentYangName = "detail"
    rpf.EntityData.SegmentPath = "rpf"
    rpf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpf.EntityData.Children = make(map[string]types.YChild)
    rpf.EntityData.Leafs = make(map[string]types.YLeaf)
    rpf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", rpf.Enable}
    rpf.EntityData.Leafs["allow-default-route"] = types.YLeaf{"AllowDefaultRoute", rpf.AllowDefaultRoute}
    rpf.EntityData.Leafs["allow-self-ping"] = types.YLeaf{"AllowSelfPing", rpf.AllowSelfPing}
    rpf.EntityData.Leafs["mode"] = types.YLeaf{"Mode", rpf.Mode}
    return &(rpf.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa
// BGP PA config on the interface
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP PA input config.
    Input Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input

    // BGP PA output config.
    Output Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output
}

func (bgpPa *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa) GetEntityData() *types.CommonEntityData {
    bgpPa.EntityData.YFilter = bgpPa.YFilter
    bgpPa.EntityData.YangName = "bgp-pa"
    bgpPa.EntityData.BundleName = "cisco_ios_xr"
    bgpPa.EntityData.ParentYangName = "detail"
    bgpPa.EntityData.SegmentPath = "bgp-pa"
    bgpPa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpPa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpPa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpPa.EntityData.Children = make(map[string]types.YChild)
    bgpPa.EntityData.Children["input"] = types.YChild{"Input", &bgpPa.Input}
    bgpPa.EntityData.Children["output"] = types.YChild{"Output", &bgpPa.Output}
    bgpPa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpPa.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input
// BGP PA input config
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is interface{} with range:
    // 0..4294967295.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (input *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "bgp-pa"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["enable"] = types.YLeaf{"Enable", input.Enable}
    input.EntityData.Leafs["source"] = types.YLeaf{"Source", input.Source}
    input.EntityData.Leafs["destination"] = types.YLeaf{"Destination", input.Destination}
    return &(input.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output
// BGP PA output config
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP PA for ingress/egress. The type is interface{} with range:
    // 0..4294967295.
    Enable interface{}

    // Enable source accouting. The type is bool.
    Source interface{}

    // Enable destination accouting. The type is bool.
    Destination interface{}
}

func (output *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_BgpPa_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "bgp-pa"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["enable"] = types.YLeaf{"Enable", output.Enable}
    output.EntityData.Leafs["source"] = types.YLeaf{"Source", output.Source}
    output.EntityData.Leafs["destination"] = types.YLeaf{"Destination", output.Destination}
    return &(output.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime
// Address Publish Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (utime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Utime) GetEntityData() *types.CommonEntityData {
    utime.EntityData.YFilter = utime.YFilter
    utime.EntityData.YangName = "utime"
    utime.EntityData.BundleName = "cisco_ios_xr"
    utime.EntityData.ParentYangName = "detail"
    utime.EntityData.SegmentPath = "utime"
    utime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    utime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    utime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    utime.EntityData.Children = make(map[string]types.YChild)
    utime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(utime.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime
// IDB Create Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (idbUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_IdbUtime) GetEntityData() *types.CommonEntityData {
    idbUtime.EntityData.YFilter = idbUtime.YFilter
    idbUtime.EntityData.YangName = "idb-utime"
    idbUtime.EntityData.BundleName = "cisco_ios_xr"
    idbUtime.EntityData.ParentYangName = "detail"
    idbUtime.EntityData.SegmentPath = "idb-utime"
    idbUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idbUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idbUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idbUtime.EntityData.Children = make(map[string]types.YChild)
    idbUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(idbUtime.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime
// CAPS Add Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (capsUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_CapsUtime) GetEntityData() *types.CommonEntityData {
    capsUtime.EntityData.YFilter = capsUtime.YFilter
    capsUtime.EntityData.YangName = "caps-utime"
    capsUtime.EntityData.BundleName = "cisco_ios_xr"
    capsUtime.EntityData.ParentYangName = "detail"
    capsUtime.EntityData.SegmentPath = "caps-utime"
    capsUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capsUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capsUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capsUtime.EntityData.Children = make(map[string]types.YChild)
    capsUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(capsUtime.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime
// FWD ENABLE Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (fwdEnUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdEnUtime) GetEntityData() *types.CommonEntityData {
    fwdEnUtime.EntityData.YFilter = fwdEnUtime.YFilter
    fwdEnUtime.EntityData.YangName = "fwd-en-utime"
    fwdEnUtime.EntityData.BundleName = "cisco_ios_xr"
    fwdEnUtime.EntityData.ParentYangName = "detail"
    fwdEnUtime.EntityData.SegmentPath = "fwd-en-utime"
    fwdEnUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwdEnUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwdEnUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwdEnUtime.EntityData.Children = make(map[string]types.YChild)
    fwdEnUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fwdEnUtime.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime
// FWD DISABLE Time
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (fwdDisUtime *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_FwdDisUtime) GetEntityData() *types.CommonEntityData {
    fwdDisUtime.EntityData.YFilter = fwdDisUtime.YFilter
    fwdDisUtime.EntityData.YangName = "fwd-dis-utime"
    fwdDisUtime.EntityData.BundleName = "cisco_ios_xr"
    fwdDisUtime.EntityData.ParentYangName = "detail"
    fwdDisUtime.EntityData.SegmentPath = "fwd-dis-utime"
    fwdDisUtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwdDisUtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwdDisUtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwdDisUtime.EntityData.Children = make(map[string]types.YChild)
    fwdDisUtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fwdDisUtime.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup
// IPv6 Multicast Group
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address of Multicast Group. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (multicastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_MulticastGroup) GetEntityData() *types.CommonEntityData {
    multicastGroup.EntityData.YFilter = multicastGroup.YFilter
    multicastGroup.EntityData.YangName = "multicast-group"
    multicastGroup.EntityData.BundleName = "cisco_ios_xr"
    multicastGroup.EntityData.ParentYangName = "detail"
    multicastGroup.EntityData.SegmentPath = "multicast-group"
    multicastGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multicastGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multicastGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multicastGroup.EntityData.Children = make(map[string]types.YChild)
    multicastGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    multicastGroup.EntityData.Leafs["address"] = types.YLeaf{"Address", multicastGroup.Address}
    return &(multicastGroup.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address
// Address List
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Prefix Length of IPv6 Address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // State of Address. The type is Ipv6MaIfAddrState.
    AddressState interface{}

    // Anycast address. The type is bool.
    IsAnycast interface{}

    // Route-tag of the Address. The type is interface{} with range:
    // 0..4294967295.
    RouteTag interface{}
}

func (address *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "detail"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    address.EntityData.Leafs["address"] = types.YLeaf{"Address", address.Address}
    address.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", address.PrefixLength}
    address.EntityData.Leafs["address-state"] = types.YLeaf{"AddressState", address.AddressState}
    address.EntityData.Leafs["is-anycast"] = types.YLeaf{"IsAnycast", address.IsAnycast}
    address.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", address.RouteTag}
    return &(address.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup
// IPv6 Client Multicast Group
type Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Address of Multicast Group. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (clientMulticastGroup *Ipv6Network_Nodes_Node_InterfaceData_Vrfs_Vrf_Details_Detail_ClientMulticastGroup) GetEntityData() *types.CommonEntityData {
    clientMulticastGroup.EntityData.YFilter = clientMulticastGroup.YFilter
    clientMulticastGroup.EntityData.YangName = "client-multicast-group"
    clientMulticastGroup.EntityData.BundleName = "cisco_ios_xr"
    clientMulticastGroup.EntityData.ParentYangName = "detail"
    clientMulticastGroup.EntityData.SegmentPath = "client-multicast-group"
    clientMulticastGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientMulticastGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientMulticastGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientMulticastGroup.EntityData.Children = make(map[string]types.YChild)
    clientMulticastGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    clientMulticastGroup.EntityData.Leafs["address"] = types.YLeaf{"Address", clientMulticastGroup.Address}
    return &(clientMulticastGroup.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Summary
// Summary of IPv6 network operational interface
// data on a node
type Ipv6Network_Nodes_Node_InterfaceData_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces (up,down) with basecaps up. The type is interface{}
    // with range: 0..4294967295.
    IfUpDownBasecapsUp interface{}

    // Number of interfaces (up,up).
    IfUpUp Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp

    // Number of interfaces (up,down).
    IfUpDown Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown

    // Number of interfaces (down,down).
    IfDownDown Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown

    // Number of interfaces (shutdown,down).
    IfShutdownDown Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown
}

func (summary *Ipv6Network_Nodes_Node_InterfaceData_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "interface-data"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["if-up-up"] = types.YChild{"IfUpUp", &summary.IfUpUp}
    summary.EntityData.Children["if-up-down"] = types.YChild{"IfUpDown", &summary.IfUpDown}
    summary.EntityData.Children["if-down-down"] = types.YChild{"IfDownDown", &summary.IfDownDown}
    summary.EntityData.Children["if-shutdown-down"] = types.YChild{"IfShutdownDown", &summary.IfShutdownDown}
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["if-up-down-basecaps-up"] = types.YLeaf{"IfUpDownBasecapsUp", summary.IfUpDownBasecapsUp}
    return &(summary.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp
// Number of interfaces (up,up)
type Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces without explicit address. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifUpUp *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpUp) GetEntityData() *types.CommonEntityData {
    ifUpUp.EntityData.YFilter = ifUpUp.YFilter
    ifUpUp.EntityData.YangName = "if-up-up"
    ifUpUp.EntityData.BundleName = "cisco_ios_xr"
    ifUpUp.EntityData.ParentYangName = "summary"
    ifUpUp.EntityData.SegmentPath = "if-up-up"
    ifUpUp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifUpUp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifUpUp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifUpUp.EntityData.Children = make(map[string]types.YChild)
    ifUpUp.EntityData.Leafs = make(map[string]types.YLeaf)
    ifUpUp.EntityData.Leafs["ip-assigned"] = types.YLeaf{"IpAssigned", ifUpUp.IpAssigned}
    ifUpUp.EntityData.Leafs["ip-unnumbered"] = types.YLeaf{"IpUnnumbered", ifUpUp.IpUnnumbered}
    ifUpUp.EntityData.Leafs["ip-unassigned"] = types.YLeaf{"IpUnassigned", ifUpUp.IpUnassigned}
    return &(ifUpUp.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown
// Number of interfaces (up,down)
type Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces without explicit address. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifUpDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfUpDown) GetEntityData() *types.CommonEntityData {
    ifUpDown.EntityData.YFilter = ifUpDown.YFilter
    ifUpDown.EntityData.YangName = "if-up-down"
    ifUpDown.EntityData.BundleName = "cisco_ios_xr"
    ifUpDown.EntityData.ParentYangName = "summary"
    ifUpDown.EntityData.SegmentPath = "if-up-down"
    ifUpDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifUpDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifUpDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifUpDown.EntityData.Children = make(map[string]types.YChild)
    ifUpDown.EntityData.Leafs = make(map[string]types.YLeaf)
    ifUpDown.EntityData.Leafs["ip-assigned"] = types.YLeaf{"IpAssigned", ifUpDown.IpAssigned}
    ifUpDown.EntityData.Leafs["ip-unnumbered"] = types.YLeaf{"IpUnnumbered", ifUpDown.IpUnnumbered}
    ifUpDown.EntityData.Leafs["ip-unassigned"] = types.YLeaf{"IpUnassigned", ifUpDown.IpUnassigned}
    return &(ifUpDown.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown
// Number of interfaces (down,down)
type Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces without explicit address. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifDownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfDownDown) GetEntityData() *types.CommonEntityData {
    ifDownDown.EntityData.YFilter = ifDownDown.YFilter
    ifDownDown.EntityData.YangName = "if-down-down"
    ifDownDown.EntityData.BundleName = "cisco_ios_xr"
    ifDownDown.EntityData.ParentYangName = "summary"
    ifDownDown.EntityData.SegmentPath = "if-down-down"
    ifDownDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifDownDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifDownDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifDownDown.EntityData.Children = make(map[string]types.YChild)
    ifDownDown.EntityData.Leafs = make(map[string]types.YLeaf)
    ifDownDown.EntityData.Leafs["ip-assigned"] = types.YLeaf{"IpAssigned", ifDownDown.IpAssigned}
    ifDownDown.EntityData.Leafs["ip-unnumbered"] = types.YLeaf{"IpUnnumbered", ifDownDown.IpUnnumbered}
    ifDownDown.EntityData.Leafs["ip-unassigned"] = types.YLeaf{"IpUnassigned", ifDownDown.IpUnassigned}
    return &(ifDownDown.EntityData)
}

// Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown
// Number of interfaces (shutdown,down)
type Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces with explicit addresses. The type is interface{} with
    // range: 0..4294967295.
    IpAssigned interface{}

    // Number of unnumbered interfaces with explicit addresses. The type is
    // interface{} with range: 0..4294967295.
    IpUnnumbered interface{}

    // Number of unassigned interfaces without explicit address. The type is
    // interface{} with range: 0..4294967295.
    IpUnassigned interface{}
}

func (ifShutdownDown *Ipv6Network_Nodes_Node_InterfaceData_Summary_IfShutdownDown) GetEntityData() *types.CommonEntityData {
    ifShutdownDown.EntityData.YFilter = ifShutdownDown.YFilter
    ifShutdownDown.EntityData.YangName = "if-shutdown-down"
    ifShutdownDown.EntityData.BundleName = "cisco_ios_xr"
    ifShutdownDown.EntityData.ParentYangName = "summary"
    ifShutdownDown.EntityData.SegmentPath = "if-shutdown-down"
    ifShutdownDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifShutdownDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifShutdownDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifShutdownDown.EntityData.Children = make(map[string]types.YChild)
    ifShutdownDown.EntityData.Leafs = make(map[string]types.YLeaf)
    ifShutdownDown.EntityData.Leafs["ip-assigned"] = types.YLeaf{"IpAssigned", ifShutdownDown.IpAssigned}
    ifShutdownDown.EntityData.Leafs["ip-unnumbered"] = types.YLeaf{"IpUnnumbered", ifShutdownDown.IpUnnumbered}
    ifShutdownDown.EntityData.Leafs["ip-unassigned"] = types.YLeaf{"IpUnassigned", ifShutdownDown.IpUnassigned}
    return &(ifShutdownDown.EntityData)
}

