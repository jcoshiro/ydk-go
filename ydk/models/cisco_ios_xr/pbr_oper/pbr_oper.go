// This module contains a collection of YANG definitions
// for Cisco IOS-XR pbr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pbr: PBR operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package pbr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pbr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-pbr-oper pbr}", reflect.TypeOf(Pbr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-pbr-oper:pbr", reflect.TypeOf(Pbr{}))
}

// PolicyState represents Different Interface states
type PolicyState string

const (
    // active
    PolicyState_active PolicyState = "active"

    // suspended
    PolicyState_suspended PolicyState = "suspended"
)

// Pbr
// PBR operational data
type Pbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific PBR operational data.
    Nodes Pbr_Nodes
}

func (pbr *Pbr) GetEntityData() *types.CommonEntityData {
    pbr.EntityData.YFilter = pbr.YFilter
    pbr.EntityData.YangName = "pbr"
    pbr.EntityData.BundleName = "cisco_ios_xr"
    pbr.EntityData.ParentYangName = "Cisco-IOS-XR-pbr-oper"
    pbr.EntityData.SegmentPath = "Cisco-IOS-XR-pbr-oper:pbr"
    pbr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbr.EntityData.Children = make(map[string]types.YChild)
    pbr.EntityData.Children["nodes"] = types.YChild{"Nodes", &pbr.Nodes}
    pbr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pbr.EntityData)
}

// Pbr_Nodes
// Node-specific PBR operational data
type Pbr_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBR operational data for a particular node. The type is slice of
    // Pbr_Nodes_Node.
    Node []Pbr_Nodes_Node
}

func (nodes *Pbr_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "pbr"
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

// Pbr_Nodes_Node
// PBR operational data for a particular node
type Pbr_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Operational data for policymaps.
    PolicyMap Pbr_Nodes_Node_PolicyMap
}

func (node *Pbr_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["policy-map"] = types.YChild{"PolicyMap", &node.PolicyMap}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Pbr_Nodes_Node_PolicyMap
// Operational data for policymaps
type Pbr_Nodes_Node_PolicyMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for all interfaces.
    Interfaces Pbr_Nodes_Node_PolicyMap_Interfaces
}

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetEntityData() *types.CommonEntityData {
    policyMap.EntityData.YFilter = policyMap.YFilter
    policyMap.EntityData.YangName = "policy-map"
    policyMap.EntityData.BundleName = "cisco_ios_xr"
    policyMap.EntityData.ParentYangName = "node"
    policyMap.EntityData.SegmentPath = "policy-map"
    policyMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyMap.EntityData.Children = make(map[string]types.YChild)
    policyMap.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &policyMap.Interfaces}
    policyMap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policyMap.EntityData)
}

// Pbr_Nodes_Node_PolicyMap_Interfaces
// Operational data for all interfaces
type Pbr_Nodes_Node_PolicyMap_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBR action data for a particular interface. The type is slice of
    // Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_.
    Interface_ []Pbr_Nodes_Node_PolicyMap_Interfaces_Interface
}

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "policy-map"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface
// PBR action data for a particular interface
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // PBR direction.
    Direction Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction
}

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["direction"] = types.YChild{"Direction", &self.Direction}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction
// PBR direction
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBR policy statistics.
    Input Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input
}

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetEntityData() *types.CommonEntityData {
    direction.EntityData.YFilter = direction.YFilter
    direction.EntityData.YangName = "direction"
    direction.EntityData.BundleName = "cisco_ios_xr"
    direction.EntityData.ParentYangName = "interface"
    direction.EntityData.SegmentPath = "direction"
    direction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    direction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    direction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    direction.EntityData.Children = make(map[string]types.YChild)
    direction.EntityData.Children["input"] = types.YChild{"Input", &direction.Input}
    direction.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(direction.EntityData)
}

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input
// PBR policy statistics
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NodeName. The type is string with length: 0..42.
    NodeName interface{}

    // PolicyName. The type is string with length: 0..65.
    PolicyName interface{}

    // State. The type is PolicyState.
    State interface{}

    // StateDescription. The type is string with length: 0..128.
    StateDescription interface{}

    // Array of classes contained in policy. The type is slice of
    // Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat.
    ClassStat []Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat
}

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "direction"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["class-stat"] = types.YChild{"ClassStat", nil}
    for i := range input.ClassStat {
        input.EntityData.Children[types.GetSegmentPath(&input.ClassStat[i])] = types.YChild{"ClassStat", &input.ClassStat[i]}
    }
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", input.NodeName}
    input.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", input.PolicyName}
    input.EntityData.Leafs["state"] = types.YLeaf{"State", input.State}
    input.EntityData.Leafs["state-description"] = types.YLeaf{"StateDescription", input.StateDescription}
    return &(input.EntityData)
}

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat
// Array of classes contained in policy
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bitmask to indicate which counter or counters are undetermined. Counters
    // will be marked undetermined when one or more classes share queues with
    // class-default because in such cases the value of counters for each class is
    // invalid. Based on the flag(s) set, the following counters will be marked
    // undetermined. For example, if value of this object returned is 0x00000101,
    // counters TransmitPackets/TransmitBytes/TotalTransmitRate and
    // DropPackets/DropBytes are undetermined .0x00000001 - Transmit
    // (TransmitPackets/TransmitBytes/TotalTransmitRate ), 0x00000002 - Drop
    // (TotalDropPackets/TotalDropBytes/TotalDropRate), 0x00000004 - Httpr
    // (HttprTransmitPackets/HttprTransmitBytes), . The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    CounterValidityBitmask interface{}

    // ClassName. The type is string with length: 0..65.
    ClassName interface{}

    // ClassId. The type is interface{} with range: 0..4294967295.
    ClassId interface{}

    // general stats.
    GeneralStats Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats

    // HTTPR stats.
    HttprStats Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats
}

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetEntityData() *types.CommonEntityData {
    classStat.EntityData.YFilter = classStat.YFilter
    classStat.EntityData.YangName = "class-stat"
    classStat.EntityData.BundleName = "cisco_ios_xr"
    classStat.EntityData.ParentYangName = "input"
    classStat.EntityData.SegmentPath = "class-stat"
    classStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classStat.EntityData.Children = make(map[string]types.YChild)
    classStat.EntityData.Children["general-stats"] = types.YChild{"GeneralStats", &classStat.GeneralStats}
    classStat.EntityData.Children["httpr-stats"] = types.YChild{"HttprStats", &classStat.HttprStats}
    classStat.EntityData.Leafs = make(map[string]types.YLeaf)
    classStat.EntityData.Leafs["counter-validity-bitmask"] = types.YLeaf{"CounterValidityBitmask", classStat.CounterValidityBitmask}
    classStat.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", classStat.ClassName}
    classStat.EntityData.Leafs["class-id"] = types.YLeaf{"ClassId", classStat.ClassId}
    return &(classStat.EntityData)
}

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats
// general stats
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmitted packets (packets/bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TransmitPackets interface{}

    // Transmitted bytes (packets/bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TransmitBytes interface{}

    // Dropped packets (packets/bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TotalDropPackets interface{}

    // Dropped bytes (packets/bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TotalDropBytes interface{}

    // Total drop rate (packets/bytes). The type is interface{} with range:
    // 0..4294967295. Units are byte.
    TotalDropRate interface{}

    // Incoming matched data rate in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    MatchDataRate interface{}

    // Total transmit rate in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    TotalTransmitRate interface{}

    // Matched pkts before applying policy. The type is interface{} with range:
    // 0..18446744073709551615.
    PrePolicyMatchedPackets interface{}

    // Matched bytes before applying policy. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    PrePolicyMatchedBytes interface{}
}

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetEntityData() *types.CommonEntityData {
    generalStats.EntityData.YFilter = generalStats.YFilter
    generalStats.EntityData.YangName = "general-stats"
    generalStats.EntityData.BundleName = "cisco_ios_xr"
    generalStats.EntityData.ParentYangName = "class-stat"
    generalStats.EntityData.SegmentPath = "general-stats"
    generalStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    generalStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    generalStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    generalStats.EntityData.Children = make(map[string]types.YChild)
    generalStats.EntityData.Leafs = make(map[string]types.YLeaf)
    generalStats.EntityData.Leafs["transmit-packets"] = types.YLeaf{"TransmitPackets", generalStats.TransmitPackets}
    generalStats.EntityData.Leafs["transmit-bytes"] = types.YLeaf{"TransmitBytes", generalStats.TransmitBytes}
    generalStats.EntityData.Leafs["total-drop-packets"] = types.YLeaf{"TotalDropPackets", generalStats.TotalDropPackets}
    generalStats.EntityData.Leafs["total-drop-bytes"] = types.YLeaf{"TotalDropBytes", generalStats.TotalDropBytes}
    generalStats.EntityData.Leafs["total-drop-rate"] = types.YLeaf{"TotalDropRate", generalStats.TotalDropRate}
    generalStats.EntityData.Leafs["match-data-rate"] = types.YLeaf{"MatchDataRate", generalStats.MatchDataRate}
    generalStats.EntityData.Leafs["total-transmit-rate"] = types.YLeaf{"TotalTransmitRate", generalStats.TotalTransmitRate}
    generalStats.EntityData.Leafs["pre-policy-matched-packets"] = types.YLeaf{"PrePolicyMatchedPackets", generalStats.PrePolicyMatchedPackets}
    generalStats.EntityData.Leafs["pre-policy-matched-bytes"] = types.YLeaf{"PrePolicyMatchedBytes", generalStats.PrePolicyMatchedBytes}
    return &(generalStats.EntityData)
}

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats
// HTTPR stats
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TotalNum of pkts HTTP request received. The type is interface{} with range:
    // 0..18446744073709551615.
    RqstRcvdPackets interface{}

    // TotalNum of Bytes HTTP request received. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    RqstRcvdBytes interface{}

    // Dropped  packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DropPackets interface{}

    // Dropped bytes. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    DropBytes interface{}

    // TotalNum of pkts HTTPR response sent. The type is interface{} with range:
    // 0..18446744073709551615.
    RespSentPackets interface{}

    // TotalNum of Bytes HTTPR response sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    RespSentBytes interface{}
}

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetEntityData() *types.CommonEntityData {
    httprStats.EntityData.YFilter = httprStats.YFilter
    httprStats.EntityData.YangName = "httpr-stats"
    httprStats.EntityData.BundleName = "cisco_ios_xr"
    httprStats.EntityData.ParentYangName = "class-stat"
    httprStats.EntityData.SegmentPath = "httpr-stats"
    httprStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    httprStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    httprStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    httprStats.EntityData.Children = make(map[string]types.YChild)
    httprStats.EntityData.Leafs = make(map[string]types.YLeaf)
    httprStats.EntityData.Leafs["rqst-rcvd-packets"] = types.YLeaf{"RqstRcvdPackets", httprStats.RqstRcvdPackets}
    httprStats.EntityData.Leafs["rqst-rcvd-bytes"] = types.YLeaf{"RqstRcvdBytes", httprStats.RqstRcvdBytes}
    httprStats.EntityData.Leafs["drop-packets"] = types.YLeaf{"DropPackets", httprStats.DropPackets}
    httprStats.EntityData.Leafs["drop-bytes"] = types.YLeaf{"DropBytes", httprStats.DropBytes}
    httprStats.EntityData.Leafs["resp-sent-packets"] = types.YLeaf{"RespSentPackets", httprStats.RespSentPackets}
    httprStats.EntityData.Leafs["resp-sent-bytes"] = types.YLeaf{"RespSentBytes", httprStats.RespSentBytes}
    return &(httprStats.EntityData)
}

