// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-fsi package operational data.
// 
// This module contains definitions
// for the following management objects:
//   fabric-stats: Fabric stats operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_fsi_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_fsi_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-fsi-oper fabric-stats}", reflect.TypeOf(FabricStats{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-fsi-oper:fabric-stats", reflect.TypeOf(FabricStats{}))
}

// FabricStats
// Fabric stats operational data
type FabricStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Nodes.
    Nodes FabricStats_Nodes
}

func (fabricStats *FabricStats) GetEntityData() *types.CommonEntityData {
    fabricStats.EntityData.YFilter = fabricStats.YFilter
    fabricStats.EntityData.YangName = "fabric-stats"
    fabricStats.EntityData.BundleName = "cisco_ios_xr"
    fabricStats.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-fsi-oper"
    fabricStats.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-fsi-oper:fabric-stats"
    fabricStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabricStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabricStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabricStats.EntityData.Children = make(map[string]types.YChild)
    fabricStats.EntityData.Children["nodes"] = types.YChild{"Nodes", &fabricStats.Nodes}
    fabricStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fabricStats.EntityData)
}

// FabricStats_Nodes
// Table of Nodes
type FabricStats_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular node. The type is slice of
    // FabricStats_Nodes_Node.
    Node []FabricStats_Nodes_Node
}

func (nodes *FabricStats_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "fabric-stats"
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

// FabricStats_Nodes_Node
// Information about a particular node
type FabricStats_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Table of stats information.
    Statses FabricStats_Nodes_Node_Statses
}

func (node *FabricStats_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["statses"] = types.YChild{"Statses", &node.Statses}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// FabricStats_Nodes_Node_Statses
// Table of stats information
type FabricStats_Nodes_Node_Statses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats information for a particular type. The type is slice of
    // FabricStats_Nodes_Node_Statses_Stats.
    Stats []FabricStats_Nodes_Node_Statses_Stats
}

func (statses *FabricStats_Nodes_Node_Statses) GetEntityData() *types.CommonEntityData {
    statses.EntityData.YFilter = statses.YFilter
    statses.EntityData.YangName = "statses"
    statses.EntityData.BundleName = "cisco_ios_xr"
    statses.EntityData.ParentYangName = "node"
    statses.EntityData.SegmentPath = "statses"
    statses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statses.EntityData.Children = make(map[string]types.YChild)
    statses.EntityData.Children["stats"] = types.YChild{"Stats", nil}
    for i := range statses.Stats {
        statses.EntityData.Children[types.GetSegmentPath(&statses.Stats[i])] = types.YChild{"Stats", &statses.Stats[i]}
    }
    statses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statses.EntityData)
}

// FabricStats_Nodes_Node_Statses_Stats
// Stats information for a particular type
type FabricStats_Nodes_Node_Statses_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fabric asic type. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Type_ interface{}

    // Last Clear Time. The type is interface{} with range:
    // 0..18446744073709551615.
    LastClearTime interface{}

    // Stat Table Name. The type is string.
    StatTableName interface{}

    // Array of counters . The type is slice of
    // FabricStats_Nodes_Node_Statses_Stats_StatsTable.
    StatsTable []FabricStats_Nodes_Node_Statses_Stats_StatsTable
}

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "statses"
    stats.EntityData.SegmentPath = "stats" + "[type='" + fmt.Sprintf("%v", stats.Type_) + "']"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Children["stats-table"] = types.YChild{"StatsTable", nil}
    for i := range stats.StatsTable {
        stats.EntityData.Children[types.GetSegmentPath(&stats.StatsTable[i])] = types.YChild{"StatsTable", &stats.StatsTable[i]}
    }
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["type"] = types.YLeaf{"Type_", stats.Type_}
    stats.EntityData.Leafs["last-clear-time"] = types.YLeaf{"LastClearTime", stats.LastClearTime}
    stats.EntityData.Leafs["stat-table-name"] = types.YLeaf{"StatTableName", stats.StatTableName}
    return &(stats.EntityData)
}

// FabricStats_Nodes_Node_Statses_Stats_StatsTable
// Array of counters 
type FabricStats_Nodes_Node_Statses_Stats_StatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats Table. The type is slice of
    // FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat.
    FsiStat []FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat
}

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetEntityData() *types.CommonEntityData {
    statsTable.EntityData.YFilter = statsTable.YFilter
    statsTable.EntityData.YangName = "stats-table"
    statsTable.EntityData.BundleName = "cisco_ios_xr"
    statsTable.EntityData.ParentYangName = "stats"
    statsTable.EntityData.SegmentPath = "stats-table"
    statsTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsTable.EntityData.Children = make(map[string]types.YChild)
    statsTable.EntityData.Children["fsi-stat"] = types.YChild{"FsiStat", nil}
    for i := range statsTable.FsiStat {
        statsTable.EntityData.Children[types.GetSegmentPath(&statsTable.FsiStat[i])] = types.YChild{"FsiStat", &statsTable.FsiStat[i]}
    }
    statsTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statsTable.EntityData)
}

// FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat
// Stats Table
type FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Counter. The type is interface{} with range: 0..18446744073709551615.
    Count interface{}

    // Counter Name. The type is string.
    CounterName interface{}
}

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetEntityData() *types.CommonEntityData {
    fsiStat.EntityData.YFilter = fsiStat.YFilter
    fsiStat.EntityData.YangName = "fsi-stat"
    fsiStat.EntityData.BundleName = "cisco_ios_xr"
    fsiStat.EntityData.ParentYangName = "stats-table"
    fsiStat.EntityData.SegmentPath = "fsi-stat"
    fsiStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fsiStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fsiStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fsiStat.EntityData.Children = make(map[string]types.YChild)
    fsiStat.EntityData.Leafs = make(map[string]types.YLeaf)
    fsiStat.EntityData.Leafs["count"] = types.YLeaf{"Count", fsiStat.Count}
    fsiStat.EntityData.Leafs["counter-name"] = types.YLeaf{"CounterName", fsiStat.CounterName}
    return &(fsiStat.EntityData)
}

