// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-pfilter package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pfilter-ma: Root class of PfilterMa Oper schema
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_pfilter_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_pfilter_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-pfilter-oper pfilter-ma}", reflect.TypeOf(PfilterMa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-pfilter-oper:pfilter-ma", reflect.TypeOf(PfilterMa{}))
}

// PfilterMa
// Root class of PfilterMa Oper schema
type PfilterMa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific operational data.
    Nodes PfilterMa_Nodes
}

func (pfilterMa *PfilterMa) GetEntityData() *types.CommonEntityData {
    pfilterMa.EntityData.YFilter = pfilterMa.YFilter
    pfilterMa.EntityData.YangName = "pfilter-ma"
    pfilterMa.EntityData.BundleName = "cisco_ios_xr"
    pfilterMa.EntityData.ParentYangName = "Cisco-IOS-XR-ip-pfilter-oper"
    pfilterMa.EntityData.SegmentPath = "Cisco-IOS-XR-ip-pfilter-oper:pfilter-ma"
    pfilterMa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pfilterMa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pfilterMa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pfilterMa.EntityData.Children = make(map[string]types.YChild)
    pfilterMa.EntityData.Children["nodes"] = types.YChild{"Nodes", &pfilterMa.Nodes}
    pfilterMa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pfilterMa.EntityData)
}

// PfilterMa_Nodes
// Node-specific operational data
type PfilterMa_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PfilterMa operational data for a particular node. The type is slice of
    // PfilterMa_Nodes_Node.
    Node []PfilterMa_Nodes_Node
}

func (nodes *PfilterMa_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "pfilter-ma"
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

// PfilterMa_Nodes_Node
// PfilterMa operational data for a particular
// node
type PfilterMa_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Operational data for pfilter.
    Process PfilterMa_Nodes_Node_Process
}

func (node *PfilterMa_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["process"] = types.YChild{"Process", &node.Process}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// PfilterMa_Nodes_Node_Process
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for pfilter.
    Ipv6 PfilterMa_Nodes_Node_Process_Ipv6

    // Operational data for pfilter.
    Ipv4 PfilterMa_Nodes_Node_Process_Ipv4
}

func (process *PfilterMa_Nodes_Node_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "node"
    process.EntityData.SegmentPath = "process"
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = make(map[string]types.YChild)
    process.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &process.Ipv6}
    process.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &process.Ipv4}
    process.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(process.EntityData)
}

// PfilterMa_Nodes_Node_Process_Ipv6
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for pfilter.
    AclInfoTable PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable
}

func (ipv6 *PfilterMa_Nodes_Node_Process_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "process"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["acl-info-table"] = types.YChild{"AclInfoTable", &ipv6.AclInfoTable}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6.EntityData)
}

// PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for pfilter.
    InterfaceInfos PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable) GetEntityData() *types.CommonEntityData {
    aclInfoTable.EntityData.YFilter = aclInfoTable.YFilter
    aclInfoTable.EntityData.YangName = "acl-info-table"
    aclInfoTable.EntityData.BundleName = "cisco_ios_xr"
    aclInfoTable.EntityData.ParentYangName = "ipv6"
    aclInfoTable.EntityData.SegmentPath = "acl-info-table"
    aclInfoTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aclInfoTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aclInfoTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aclInfoTable.EntityData.Children = make(map[string]types.YChild)
    aclInfoTable.EntityData.Children["interface-infos"] = types.YChild{"InterfaceInfos", &aclInfoTable.InterfaceInfos}
    aclInfoTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aclInfoTable.EntityData)
}

// PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for pfilter in bag. The type is slice of
    // PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo.
    InterfaceInfo []PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos) GetEntityData() *types.CommonEntityData {
    interfaceInfos.EntityData.YFilter = interfaceInfos.YFilter
    interfaceInfos.EntityData.YangName = "interface-infos"
    interfaceInfos.EntityData.BundleName = "cisco_ios_xr"
    interfaceInfos.EntityData.ParentYangName = "acl-info-table"
    interfaceInfos.EntityData.SegmentPath = "interface-infos"
    interfaceInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceInfos.EntityData.Children = make(map[string]types.YChild)
    interfaceInfos.EntityData.Children["interface-info"] = types.YChild{"InterfaceInfo", nil}
    for i := range interfaceInfos.InterfaceInfo {
        interfaceInfos.EntityData.Children[types.GetSegmentPath(&interfaceInfos.InterfaceInfo[i])] = types.YChild{"InterfaceInfo", &interfaceInfos.InterfaceInfo[i]}
    }
    interfaceInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceInfos.EntityData)
}

// PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo
// Operational data for pfilter in bag
type PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // acl information. The type is string.
    AclInfo interface{}
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv6_AclInfoTable_InterfaceInfos_InterfaceInfo) GetEntityData() *types.CommonEntityData {
    interfaceInfo.EntityData.YFilter = interfaceInfo.YFilter
    interfaceInfo.EntityData.YangName = "interface-info"
    interfaceInfo.EntityData.BundleName = "cisco_ios_xr"
    interfaceInfo.EntityData.ParentYangName = "interface-infos"
    interfaceInfo.EntityData.SegmentPath = "interface-info" + "[interface-name='" + fmt.Sprintf("%v", interfaceInfo.InterfaceName) + "']"
    interfaceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceInfo.EntityData.Children = make(map[string]types.YChild)
    interfaceInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceInfo.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceInfo.InterfaceName}
    interfaceInfo.EntityData.Leafs["acl-info"] = types.YLeaf{"AclInfo", interfaceInfo.AclInfo}
    return &(interfaceInfo.EntityData)
}

// PfilterMa_Nodes_Node_Process_Ipv4
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for pfilter.
    AclInfoTable PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable
}

func (ipv4 *PfilterMa_Nodes_Node_Process_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "process"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["acl-info-table"] = types.YChild{"AclInfoTable", &ipv4.AclInfoTable}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4.EntityData)
}

// PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for pfilter.
    InterfaceInfos PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos
}

func (aclInfoTable *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable) GetEntityData() *types.CommonEntityData {
    aclInfoTable.EntityData.YFilter = aclInfoTable.YFilter
    aclInfoTable.EntityData.YangName = "acl-info-table"
    aclInfoTable.EntityData.BundleName = "cisco_ios_xr"
    aclInfoTable.EntityData.ParentYangName = "ipv4"
    aclInfoTable.EntityData.SegmentPath = "acl-info-table"
    aclInfoTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aclInfoTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aclInfoTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aclInfoTable.EntityData.Children = make(map[string]types.YChild)
    aclInfoTable.EntityData.Children["interface-infos"] = types.YChild{"InterfaceInfos", &aclInfoTable.InterfaceInfos}
    aclInfoTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aclInfoTable.EntityData)
}

// PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos
// Operational data for pfilter
type PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for pfilter in bag. The type is slice of
    // PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo.
    InterfaceInfo []PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo
}

func (interfaceInfos *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos) GetEntityData() *types.CommonEntityData {
    interfaceInfos.EntityData.YFilter = interfaceInfos.YFilter
    interfaceInfos.EntityData.YangName = "interface-infos"
    interfaceInfos.EntityData.BundleName = "cisco_ios_xr"
    interfaceInfos.EntityData.ParentYangName = "acl-info-table"
    interfaceInfos.EntityData.SegmentPath = "interface-infos"
    interfaceInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceInfos.EntityData.Children = make(map[string]types.YChild)
    interfaceInfos.EntityData.Children["interface-info"] = types.YChild{"InterfaceInfo", nil}
    for i := range interfaceInfos.InterfaceInfo {
        interfaceInfos.EntityData.Children[types.GetSegmentPath(&interfaceInfos.InterfaceInfo[i])] = types.YChild{"InterfaceInfo", &interfaceInfos.InterfaceInfo[i]}
    }
    interfaceInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceInfos.EntityData)
}

// PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo
// Operational data for pfilter in bag
type PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // acl information. The type is string.
    AclInfo interface{}
}

func (interfaceInfo *PfilterMa_Nodes_Node_Process_Ipv4_AclInfoTable_InterfaceInfos_InterfaceInfo) GetEntityData() *types.CommonEntityData {
    interfaceInfo.EntityData.YFilter = interfaceInfo.YFilter
    interfaceInfo.EntityData.YangName = "interface-info"
    interfaceInfo.EntityData.BundleName = "cisco_ios_xr"
    interfaceInfo.EntityData.ParentYangName = "interface-infos"
    interfaceInfo.EntityData.SegmentPath = "interface-info" + "[interface-name='" + fmt.Sprintf("%v", interfaceInfo.InterfaceName) + "']"
    interfaceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceInfo.EntityData.Children = make(map[string]types.YChild)
    interfaceInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceInfo.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceInfo.InterfaceName}
    interfaceInfo.EntityData.Leafs["acl-info"] = types.YLeaf{"AclInfo", interfaceInfo.AclInfo}
    return &(interfaceInfo.EntityData)
}

