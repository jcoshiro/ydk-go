// This module contains a collection of YANG definitions
// for Cisco IOS-XR prm-server package operational data.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: PRM data
//   prm: prm
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package prm_server_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package prm_server_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-prm-server-oper hardware-module}", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-prm-server-oper:hardware-module", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-prm-server-oper prm}", reflect.TypeOf(Prm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-prm-server-oper:prm", reflect.TypeOf(Prm{}))
}

// HardwareModule
// PRM data
type HardwareModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of PRM Nodes.
    Nodes HardwareModule_Nodes
}

func (hardwareModule *HardwareModule) GetEntityData() *types.CommonEntityData {
    hardwareModule.EntityData.YFilter = hardwareModule.YFilter
    hardwareModule.EntityData.YangName = "hardware-module"
    hardwareModule.EntityData.BundleName = "cisco_ios_xr"
    hardwareModule.EntityData.ParentYangName = "Cisco-IOS-XR-prm-server-oper"
    hardwareModule.EntityData.SegmentPath = "Cisco-IOS-XR-prm-server-oper:hardware-module"
    hardwareModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModule.EntityData.Children = make(map[string]types.YChild)
    hardwareModule.EntityData.Children["nodes"] = types.YChild{"Nodes", &hardwareModule.Nodes}
    hardwareModule.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareModule.EntityData)
}

// HardwareModule_Nodes
// List of PRM Nodes
type HardwareModule_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Information. The type is slice of HardwareModule_Nodes_Node.
    Node []HardwareModule_Nodes_Node
}

func (nodes *HardwareModule_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "hardware-module"
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

// HardwareModule_Nodes_Node
// Node Information
type HardwareModule_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Server specific.
    Np HardwareModule_Nodes_Node_Np
}

func (node *HardwareModule_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["np"] = types.YChild{"Np", &node.Np}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// HardwareModule_Nodes_Node_Np
// Server specific
type HardwareModule_Nodes_Node_Np struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource specific.
    Cpu HardwareModule_Nodes_Node_Np_Cpu

    // Platform drops.
    PlatformDrop HardwareModule_Nodes_Node_Np_PlatformDrop
}

func (np *HardwareModule_Nodes_Node_Np) GetEntityData() *types.CommonEntityData {
    np.EntityData.YFilter = np.YFilter
    np.EntityData.YangName = "np"
    np.EntityData.BundleName = "cisco_ios_xr"
    np.EntityData.ParentYangName = "node"
    np.EntityData.SegmentPath = "np"
    np.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    np.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    np.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    np.EntityData.Children = make(map[string]types.YChild)
    np.EntityData.Children["cpu"] = types.YChild{"Cpu", &np.Cpu}
    np.EntityData.Children["platform-drop"] = types.YChild{"PlatformDrop", &np.PlatformDrop}
    np.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(np.EntityData)
}

// HardwareModule_Nodes_Node_Np_Cpu
// Resource specific
type HardwareModule_Nodes_Node_Np_Cpu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data for software resource.
    Indexes HardwareModule_Nodes_Node_Np_Cpu_Indexes
}

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetEntityData() *types.CommonEntityData {
    cpu.EntityData.YFilter = cpu.YFilter
    cpu.EntityData.YangName = "cpu"
    cpu.EntityData.BundleName = "cisco_ios_xr"
    cpu.EntityData.ParentYangName = "np"
    cpu.EntityData.SegmentPath = "cpu"
    cpu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpu.EntityData.Children = make(map[string]types.YChild)
    cpu.EntityData.Children["indexes"] = types.YChild{"Indexes", &cpu.Indexes}
    cpu.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpu.EntityData)
}

// HardwareModule_Nodes_Node_Np_Cpu_Indexes
// Data for software resource
type HardwareModule_Nodes_Node_Np_Cpu_Indexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue Stats. The type is slice of
    // HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index.
    Index []HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index
}

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetEntityData() *types.CommonEntityData {
    indexes.EntityData.YFilter = indexes.YFilter
    indexes.EntityData.YangName = "indexes"
    indexes.EntityData.BundleName = "cisco_ios_xr"
    indexes.EntityData.ParentYangName = "cpu"
    indexes.EntityData.SegmentPath = "indexes"
    indexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indexes.EntityData.Children = make(map[string]types.YChild)
    indexes.EntityData.Children["index"] = types.YChild{"Index", nil}
    for i := range indexes.Index {
        indexes.EntityData.Children[types.GetSegmentPath(&indexes.Index[i])] = types.YChild{"Index", &indexes.Index[i]}
    }
    indexes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(indexes.EntityData)
}

// HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index
// Queue Stats
type HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index value. The type is interface{} with range:
    // -2147483648..2147483647.
    Index interface{}

    // CosQ Name. The type is string with length: 0..1024.
    CosQName interface{}

    // CosQ No. The type is interface{} with range: 0..255.
    CosQ interface{}

    // Rx DMA Channel. The type is interface{} with range: 0..4294967295.
    RxChannel interface{}

    // Flow Rate. The type is interface{} with range: 0..4294967295.
    FlowRate interface{}

    // Burst. The type is interface{} with range: 0..4294967295.
    Burst interface{}

    // Accepted. The type is interface{} with range: 0..18446744073709551615.
    Accepted interface{}

    // Dropped. The type is interface{} with range: 0..18446744073709551615.
    Dropped interface{}
}

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetEntityData() *types.CommonEntityData {
    index.EntityData.YFilter = index.YFilter
    index.EntityData.YangName = "index"
    index.EntityData.BundleName = "cisco_ios_xr"
    index.EntityData.ParentYangName = "indexes"
    index.EntityData.SegmentPath = "index" + "[index='" + fmt.Sprintf("%v", index.Index) + "']"
    index.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    index.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    index.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    index.EntityData.Children = make(map[string]types.YChild)
    index.EntityData.Leafs = make(map[string]types.YLeaf)
    index.EntityData.Leafs["index"] = types.YLeaf{"Index", index.Index}
    index.EntityData.Leafs["cos-q-name"] = types.YLeaf{"CosQName", index.CosQName}
    index.EntityData.Leafs["cos-q"] = types.YLeaf{"CosQ", index.CosQ}
    index.EntityData.Leafs["rx-channel"] = types.YLeaf{"RxChannel", index.RxChannel}
    index.EntityData.Leafs["flow-rate"] = types.YLeaf{"FlowRate", index.FlowRate}
    index.EntityData.Leafs["burst"] = types.YLeaf{"Burst", index.Burst}
    index.EntityData.Leafs["accepted"] = types.YLeaf{"Accepted", index.Accepted}
    index.EntityData.Leafs["dropped"] = types.YLeaf{"Dropped", index.Dropped}
    return &(index.EntityData)
}

// HardwareModule_Nodes_Node_Np_PlatformDrop
// Platform drops
type HardwareModule_Nodes_Node_Np_PlatformDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Captured Packets.
    Indxes HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes

    // Stats for Drop packets.
    Idxes HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes
}

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetEntityData() *types.CommonEntityData {
    platformDrop.EntityData.YFilter = platformDrop.YFilter
    platformDrop.EntityData.YangName = "platform-drop"
    platformDrop.EntityData.BundleName = "cisco_ios_xr"
    platformDrop.EntityData.ParentYangName = "np"
    platformDrop.EntityData.SegmentPath = "platform-drop"
    platformDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformDrop.EntityData.Children = make(map[string]types.YChild)
    platformDrop.EntityData.Children["indxes"] = types.YChild{"Indxes", &platformDrop.Indxes}
    platformDrop.EntityData.Children["idxes"] = types.YChild{"Idxes", &platformDrop.Idxes}
    platformDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platformDrop.EntityData)
}

// HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes
// Captured Packets
type HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Captured packets. The type is slice of
    // HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx.
    Indx []HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx
}

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetEntityData() *types.CommonEntityData {
    indxes.EntityData.YFilter = indxes.YFilter
    indxes.EntityData.YangName = "indxes"
    indxes.EntityData.BundleName = "cisco_ios_xr"
    indxes.EntityData.ParentYangName = "platform-drop"
    indxes.EntityData.SegmentPath = "indxes"
    indxes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indxes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indxes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indxes.EntityData.Children = make(map[string]types.YChild)
    indxes.EntityData.Children["indx"] = types.YChild{"Indx", nil}
    for i := range indxes.Indx {
        indxes.EntityData.Children[types.GetSegmentPath(&indxes.Indx[i])] = types.YChild{"Indx", &indxes.Indx[i]}
    }
    indxes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(indxes.EntityData)
}

// HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx
// Captured packets
type HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index value. The type is interface{} with range:
    // -2147483648..2147483647.
    Index interface{}

    // Total packets Captured. The type is interface{} with range: 0..4294967295.
    TotalCaptured interface{}

    // Captured Packet. The type is string with length: 0..1024.
    CapturedPak interface{}

    // Packet Index. The type is interface{} with range: 0..255.
    PktIndex interface{}

    // If Handle. The type is interface{} with range: 0..4294967295.
    Ifhandle interface{}

    // Buffer Length. The type is interface{} with range: 0..4294967295.
    BufferLen interface{}

    // Reason Hi. The type is interface{} with range: 0..4294967295.
    ReasonHi interface{}

    // Reason. The type is interface{} with range: 0..4294967295.
    Reason interface{}

    // Year. The type is interface{} with range: 0..18446744073709551615.
    Years interface{}

    // Hours. The type is interface{} with range: 0..18446744073709551615. Units
    // are hour.
    Hours interface{}

    // Days. The type is interface{} with range: 0..18446744073709551615. Units
    // are day.
    Days interface{}

    // Minutes. The type is interface{} with range: 0..18446744073709551615. Units
    // are minute.
    Mins interface{}

    // Seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Secs interface{}
}

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetEntityData() *types.CommonEntityData {
    indx.EntityData.YFilter = indx.YFilter
    indx.EntityData.YangName = "indx"
    indx.EntityData.BundleName = "cisco_ios_xr"
    indx.EntityData.ParentYangName = "indxes"
    indx.EntityData.SegmentPath = "indx" + "[index='" + fmt.Sprintf("%v", indx.Index) + "']"
    indx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indx.EntityData.Children = make(map[string]types.YChild)
    indx.EntityData.Leafs = make(map[string]types.YLeaf)
    indx.EntityData.Leafs["index"] = types.YLeaf{"Index", indx.Index}
    indx.EntityData.Leafs["total-captured"] = types.YLeaf{"TotalCaptured", indx.TotalCaptured}
    indx.EntityData.Leafs["captured-pak"] = types.YLeaf{"CapturedPak", indx.CapturedPak}
    indx.EntityData.Leafs["pkt-index"] = types.YLeaf{"PktIndex", indx.PktIndex}
    indx.EntityData.Leafs["ifhandle"] = types.YLeaf{"Ifhandle", indx.Ifhandle}
    indx.EntityData.Leafs["buffer-len"] = types.YLeaf{"BufferLen", indx.BufferLen}
    indx.EntityData.Leafs["reason-hi"] = types.YLeaf{"ReasonHi", indx.ReasonHi}
    indx.EntityData.Leafs["reason"] = types.YLeaf{"Reason", indx.Reason}
    indx.EntityData.Leafs["years"] = types.YLeaf{"Years", indx.Years}
    indx.EntityData.Leafs["hours"] = types.YLeaf{"Hours", indx.Hours}
    indx.EntityData.Leafs["days"] = types.YLeaf{"Days", indx.Days}
    indx.EntityData.Leafs["mins"] = types.YLeaf{"Mins", indx.Mins}
    indx.EntityData.Leafs["secs"] = types.YLeaf{"Secs", indx.Secs}
    return &(indx.EntityData)
}

// HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes
// Stats for Drop packets
type HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Drop Stats. The type is slice of
    // HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx.
    Idx []HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx
}

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetEntityData() *types.CommonEntityData {
    idxes.EntityData.YFilter = idxes.YFilter
    idxes.EntityData.YangName = "idxes"
    idxes.EntityData.BundleName = "cisco_ios_xr"
    idxes.EntityData.ParentYangName = "platform-drop"
    idxes.EntityData.SegmentPath = "idxes"
    idxes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idxes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idxes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idxes.EntityData.Children = make(map[string]types.YChild)
    idxes.EntityData.Children["idx"] = types.YChild{"Idx", nil}
    for i := range idxes.Idx {
        idxes.EntityData.Children[types.GetSegmentPath(&idxes.Idx[i])] = types.YChild{"Idx", &idxes.Idx[i]}
    }
    idxes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(idxes.EntityData)
}

// HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx
// Drop Stats
type HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index value. The type is interface{} with range:
    // -2147483648..2147483647.
    Index interface{}

    // Drop Reason. The type is string with length: 0..1024.
    DropReason interface{}

    // Counter. The type is interface{} with range: 0..4294967295.
    Counters interface{}
}

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetEntityData() *types.CommonEntityData {
    idx.EntityData.YFilter = idx.YFilter
    idx.EntityData.YangName = "idx"
    idx.EntityData.BundleName = "cisco_ios_xr"
    idx.EntityData.ParentYangName = "idxes"
    idx.EntityData.SegmentPath = "idx" + "[index='" + fmt.Sprintf("%v", idx.Index) + "']"
    idx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idx.EntityData.Children = make(map[string]types.YChild)
    idx.EntityData.Leafs = make(map[string]types.YLeaf)
    idx.EntityData.Leafs["index"] = types.YLeaf{"Index", idx.Index}
    idx.EntityData.Leafs["drop-reason"] = types.YLeaf{"DropReason", idx.DropReason}
    idx.EntityData.Leafs["counters"] = types.YLeaf{"Counters", idx.Counters}
    return &(idx.EntityData)
}

// Prm
// prm
type Prm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of PRM Nodes.
    Nodes Prm_Nodes
}

func (prm *Prm) GetEntityData() *types.CommonEntityData {
    prm.EntityData.YFilter = prm.YFilter
    prm.EntityData.YangName = "prm"
    prm.EntityData.BundleName = "cisco_ios_xr"
    prm.EntityData.ParentYangName = "Cisco-IOS-XR-prm-server-oper"
    prm.EntityData.SegmentPath = "Cisco-IOS-XR-prm-server-oper:prm"
    prm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prm.EntityData.Children = make(map[string]types.YChild)
    prm.EntityData.Children["nodes"] = types.YChild{"Nodes", &prm.Nodes}
    prm.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prm.EntityData)
}

// Prm_Nodes
// List of PRM Nodes
type Prm_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Information. The type is slice of Prm_Nodes_Node.
    Node []Prm_Nodes_Node
}

func (nodes *Prm_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "prm"
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

// Prm_Nodes_Node
// Node Information
type Prm_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Server specific.
    Server Prm_Nodes_Node_Server
}

func (node *Prm_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["server"] = types.YChild{"Server", &node.Server}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Prm_Nodes_Node_Server
// Server specific
type Prm_Nodes_Node_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource specific.
    Resource Prm_Nodes_Node_Server_Resource
}

func (server *Prm_Nodes_Node_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "node"
    server.EntityData.SegmentPath = "server"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Children["resource"] = types.YChild{"Resource", &server.Resource}
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(server.EntityData)
}

// Prm_Nodes_Node_Server_Resource
// Resource specific
type Prm_Nodes_Node_Server_Resource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data for software resource.
    Indexes Prm_Nodes_Node_Server_Resource_Indexes
}

func (resource *Prm_Nodes_Node_Server_Resource) GetEntityData() *types.CommonEntityData {
    resource.EntityData.YFilter = resource.YFilter
    resource.EntityData.YangName = "resource"
    resource.EntityData.BundleName = "cisco_ios_xr"
    resource.EntityData.ParentYangName = "server"
    resource.EntityData.SegmentPath = "resource"
    resource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resource.EntityData.Children = make(map[string]types.YChild)
    resource.EntityData.Children["indexes"] = types.YChild{"Indexes", &resource.Indexes}
    resource.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(resource.EntityData)
}

// Prm_Nodes_Node_Server_Resource_Indexes
// Data for software resource
type Prm_Nodes_Node_Server_Resource_Indexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data for software resource. The type is slice of
    // Prm_Nodes_Node_Server_Resource_Indexes_Index.
    Index []Prm_Nodes_Node_Server_Resource_Indexes_Index
}

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetEntityData() *types.CommonEntityData {
    indexes.EntityData.YFilter = indexes.YFilter
    indexes.EntityData.YangName = "indexes"
    indexes.EntityData.BundleName = "cisco_ios_xr"
    indexes.EntityData.ParentYangName = "resource"
    indexes.EntityData.SegmentPath = "indexes"
    indexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indexes.EntityData.Children = make(map[string]types.YChild)
    indexes.EntityData.Children["index"] = types.YChild{"Index", nil}
    for i := range indexes.Index {
        indexes.EntityData.Children[types.GetSegmentPath(&indexes.Index[i])] = types.YChild{"Index", &indexes.Index[i]}
    }
    indexes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(indexes.EntityData)
}

// Prm_Nodes_Node_Server_Resource_Indexes_Index
// Data for software resource
type Prm_Nodes_Node_Server_Resource_Indexes_Index struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index value. The type is interface{} with range:
    // -2147483648..2147483647.
    Index interface{}

    // Resource Name. The type is string with length: 0..1024.
    ResourceName interface{}

    // Resource Type. The type is interface{} with range: 0..4294967295.
    ResourceType interface{}

    // Total Resource Count. The type is interface{} with range: 0..4294967295.
    TotalNum interface{}

    // Free Resource Count. The type is interface{} with range: 0..4294967295.
    FreeNum interface{}

    // Next Free Index. The type is interface{} with range: 0..4294967295.
    FirstAvailableIndex interface{}

    // Start Index. The type is interface{} with range: 0..4294967295.
    StartIndex interface{}

    // Availability Status. The type is bool.
    AvailabilityStatus interface{}

    // Resource Flags. The type is interface{} with range: 0..255.
    Flags interface{}

    // Inconsistice Flags. The type is bool.
    Inconsistent interface{}
}

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetEntityData() *types.CommonEntityData {
    index.EntityData.YFilter = index.YFilter
    index.EntityData.YangName = "index"
    index.EntityData.BundleName = "cisco_ios_xr"
    index.EntityData.ParentYangName = "indexes"
    index.EntityData.SegmentPath = "index" + "[index='" + fmt.Sprintf("%v", index.Index) + "']"
    index.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    index.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    index.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    index.EntityData.Children = make(map[string]types.YChild)
    index.EntityData.Leafs = make(map[string]types.YLeaf)
    index.EntityData.Leafs["index"] = types.YLeaf{"Index", index.Index}
    index.EntityData.Leafs["resource-name"] = types.YLeaf{"ResourceName", index.ResourceName}
    index.EntityData.Leafs["resource-type"] = types.YLeaf{"ResourceType", index.ResourceType}
    index.EntityData.Leafs["total-num"] = types.YLeaf{"TotalNum", index.TotalNum}
    index.EntityData.Leafs["free-num"] = types.YLeaf{"FreeNum", index.FreeNum}
    index.EntityData.Leafs["first-available-index"] = types.YLeaf{"FirstAvailableIndex", index.FirstAvailableIndex}
    index.EntityData.Leafs["start-index"] = types.YLeaf{"StartIndex", index.StartIndex}
    index.EntityData.Leafs["availability-status"] = types.YLeaf{"AvailabilityStatus", index.AvailabilityStatus}
    index.EntityData.Leafs["flags"] = types.YLeaf{"Flags", index.Flags}
    index.EntityData.Leafs["inconsistent"] = types.YLeaf{"Inconsistent", index.Inconsistent}
    return &(index.EntityData)
}

