// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-tc package operational data.
// 
// This module contains definitions
// for the following management objects:
//   traffic-collector: Global Traffic Collector configuration
//     commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_tc_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_tc_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-tc-oper traffic-collector}", reflect.TypeOf(TrafficCollector{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-tc-oper:traffic-collector", reflect.TypeOf(TrafficCollector{}))
}

// TcOperAfName represents Tc oper af name
type TcOperAfName string

const (
    // IPv4
    TcOperAfName_ipv4 TcOperAfName = "ipv4"

    // IPv6
    TcOperAfName_ipv6 TcOperAfName = "ipv6"
)

// TrafficCollector
// Global Traffic Collector configuration commands
type TrafficCollector struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // External Interface.
    ExternalInterfaces TrafficCollector_ExternalInterfaces

    // Traffic Collector summary.
    Summary TrafficCollector_Summary

    // VRF specific operational data.
    VrfTable TrafficCollector_VrfTable

    // Address Family specific operational data.
    Afs TrafficCollector_Afs
}

func (trafficCollector *TrafficCollector) GetEntityData() *types.CommonEntityData {
    trafficCollector.EntityData.YFilter = trafficCollector.YFilter
    trafficCollector.EntityData.YangName = "traffic-collector"
    trafficCollector.EntityData.BundleName = "cisco_ios_xr"
    trafficCollector.EntityData.ParentYangName = "Cisco-IOS-XR-infra-tc-oper"
    trafficCollector.EntityData.SegmentPath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector"
    trafficCollector.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficCollector.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficCollector.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficCollector.EntityData.Children = make(map[string]types.YChild)
    trafficCollector.EntityData.Children["external-interfaces"] = types.YChild{"ExternalInterfaces", &trafficCollector.ExternalInterfaces}
    trafficCollector.EntityData.Children["summary"] = types.YChild{"Summary", &trafficCollector.Summary}
    trafficCollector.EntityData.Children["vrf-table"] = types.YChild{"VrfTable", &trafficCollector.VrfTable}
    trafficCollector.EntityData.Children["afs"] = types.YChild{"Afs", &trafficCollector.Afs}
    trafficCollector.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trafficCollector.EntityData)
}

// TrafficCollector_ExternalInterfaces
// External Interface
type TrafficCollector_ExternalInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // External Interface . The type is slice of
    // TrafficCollector_ExternalInterfaces_ExternalInterface.
    ExternalInterface []TrafficCollector_ExternalInterfaces_ExternalInterface
}

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetEntityData() *types.CommonEntityData {
    externalInterfaces.EntityData.YFilter = externalInterfaces.YFilter
    externalInterfaces.EntityData.YangName = "external-interfaces"
    externalInterfaces.EntityData.BundleName = "cisco_ios_xr"
    externalInterfaces.EntityData.ParentYangName = "traffic-collector"
    externalInterfaces.EntityData.SegmentPath = "external-interfaces"
    externalInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    externalInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    externalInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    externalInterfaces.EntityData.Children = make(map[string]types.YChild)
    externalInterfaces.EntityData.Children["external-interface"] = types.YChild{"ExternalInterface", nil}
    for i := range externalInterfaces.ExternalInterface {
        externalInterfaces.EntityData.Children[types.GetSegmentPath(&externalInterfaces.ExternalInterface[i])] = types.YChild{"ExternalInterface", &externalInterfaces.ExternalInterface[i]}
    }
    externalInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(externalInterfaces.EntityData)
}

// TrafficCollector_ExternalInterfaces_ExternalInterface
// External Interface 
type TrafficCollector_ExternalInterfaces_ExternalInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Interface Name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name in Display format. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Interface VRF ID. The type is interface{} with range: 0..4294967295.
    Vrfid interface{}

    // Flag to indicate interface enabled or not. The type is bool.
    IsInterfaceEnabled interface{}
}

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetEntityData() *types.CommonEntityData {
    externalInterface.EntityData.YFilter = externalInterface.YFilter
    externalInterface.EntityData.YangName = "external-interface"
    externalInterface.EntityData.BundleName = "cisco_ios_xr"
    externalInterface.EntityData.ParentYangName = "external-interfaces"
    externalInterface.EntityData.SegmentPath = "external-interface" + "[interface-name='" + fmt.Sprintf("%v", externalInterface.InterfaceName) + "']"
    externalInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    externalInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    externalInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    externalInterface.EntityData.Children = make(map[string]types.YChild)
    externalInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    externalInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", externalInterface.InterfaceName}
    externalInterface.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", externalInterface.InterfaceNameXr}
    externalInterface.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", externalInterface.InterfaceHandle}
    externalInterface.EntityData.Leafs["vrfid"] = types.YLeaf{"Vrfid", externalInterface.Vrfid}
    externalInterface.EntityData.Leafs["is-interface-enabled"] = types.YLeaf{"IsInterfaceEnabled", externalInterface.IsInterfaceEnabled}
    return &(externalInterface.EntityData)
}

// TrafficCollector_Summary
// Traffic Collector summary
type TrafficCollector_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistic collection interval in minutes. The type is interface{} with
    // range: 0..255. Units are minute.
    CollectionInterval interface{}

    // TRUE if collection timer is running. The type is bool.
    CollectionTimerIsRunning interface{}

    // Statistic history timeout interval in hours. The type is interface{} with
    // range: 0..65535. Units are hour.
    TimeoutInterval interface{}

    // TRUE if history timeout timer is running. The type is bool.
    TimeoutTimerIsRunning interface{}

    // Statistics history size. The type is interface{} with range: 0..255.
    HistorySize interface{}

    // Database statistics for External Interface.
    DatabaseStatisticsExternalInterface TrafficCollector_Summary_DatabaseStatisticsExternalInterface

    // VRF table statistics. The type is slice of
    // TrafficCollector_Summary_VrfStatistic.
    VrfStatistic []TrafficCollector_Summary_VrfStatistic

    // Statistics per message type for STAT collector. The type is slice of
    // TrafficCollector_Summary_CollectionMessageStatistic.
    CollectionMessageStatistic []TrafficCollector_Summary_CollectionMessageStatistic

    // Statistics per message type for Chkpt. The type is slice of
    // TrafficCollector_Summary_CheckpointMessageStatistic.
    CheckpointMessageStatistic []TrafficCollector_Summary_CheckpointMessageStatistic
}

func (summary *TrafficCollector_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "traffic-collector"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["database-statistics-external-interface"] = types.YChild{"DatabaseStatisticsExternalInterface", &summary.DatabaseStatisticsExternalInterface}
    summary.EntityData.Children["vrf-statistic"] = types.YChild{"VrfStatistic", nil}
    for i := range summary.VrfStatistic {
        summary.EntityData.Children[types.GetSegmentPath(&summary.VrfStatistic[i])] = types.YChild{"VrfStatistic", &summary.VrfStatistic[i]}
    }
    summary.EntityData.Children["collection-message-statistic"] = types.YChild{"CollectionMessageStatistic", nil}
    for i := range summary.CollectionMessageStatistic {
        summary.EntityData.Children[types.GetSegmentPath(&summary.CollectionMessageStatistic[i])] = types.YChild{"CollectionMessageStatistic", &summary.CollectionMessageStatistic[i]}
    }
    summary.EntityData.Children["checkpoint-message-statistic"] = types.YChild{"CheckpointMessageStatistic", nil}
    for i := range summary.CheckpointMessageStatistic {
        summary.EntityData.Children[types.GetSegmentPath(&summary.CheckpointMessageStatistic[i])] = types.YChild{"CheckpointMessageStatistic", &summary.CheckpointMessageStatistic[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["collection-interval"] = types.YLeaf{"CollectionInterval", summary.CollectionInterval}
    summary.EntityData.Leafs["collection-timer-is-running"] = types.YLeaf{"CollectionTimerIsRunning", summary.CollectionTimerIsRunning}
    summary.EntityData.Leafs["timeout-interval"] = types.YLeaf{"TimeoutInterval", summary.TimeoutInterval}
    summary.EntityData.Leafs["timeout-timer-is-running"] = types.YLeaf{"TimeoutTimerIsRunning", summary.TimeoutTimerIsRunning}
    summary.EntityData.Leafs["history-size"] = types.YLeaf{"HistorySize", summary.HistorySize}
    return &(summary.EntityData)
}

// TrafficCollector_Summary_DatabaseStatisticsExternalInterface
// Database statistics for External Interface
type TrafficCollector_Summary_DatabaseStatisticsExternalInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of DB entries. The type is interface{} with range: 0..4294967295.
    NumberOfEntries interface{}

    // Number of stale  entries. The type is interface{} with range:
    // 0..4294967295.
    NumberOfStaleEntries interface{}

    // Number of add operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfAddOPerations interface{}

    // Number of delete operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfDeleteOPerations interface{}
}

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetEntityData() *types.CommonEntityData {
    databaseStatisticsExternalInterface.EntityData.YFilter = databaseStatisticsExternalInterface.YFilter
    databaseStatisticsExternalInterface.EntityData.YangName = "database-statistics-external-interface"
    databaseStatisticsExternalInterface.EntityData.BundleName = "cisco_ios_xr"
    databaseStatisticsExternalInterface.EntityData.ParentYangName = "summary"
    databaseStatisticsExternalInterface.EntityData.SegmentPath = "database-statistics-external-interface"
    databaseStatisticsExternalInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseStatisticsExternalInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseStatisticsExternalInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseStatisticsExternalInterface.EntityData.Children = make(map[string]types.YChild)
    databaseStatisticsExternalInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    databaseStatisticsExternalInterface.EntityData.Leafs["number-of-entries"] = types.YLeaf{"NumberOfEntries", databaseStatisticsExternalInterface.NumberOfEntries}
    databaseStatisticsExternalInterface.EntityData.Leafs["number-of-stale-entries"] = types.YLeaf{"NumberOfStaleEntries", databaseStatisticsExternalInterface.NumberOfStaleEntries}
    databaseStatisticsExternalInterface.EntityData.Leafs["number-of-add-o-perations"] = types.YLeaf{"NumberOfAddOPerations", databaseStatisticsExternalInterface.NumberOfAddOPerations}
    databaseStatisticsExternalInterface.EntityData.Leafs["number-of-delete-o-perations"] = types.YLeaf{"NumberOfDeleteOPerations", databaseStatisticsExternalInterface.NumberOfDeleteOPerations}
    return &(databaseStatisticsExternalInterface.EntityData)
}

// TrafficCollector_Summary_VrfStatistic
// VRF table statistics
type TrafficCollector_Summary_VrfStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string.
    VrfName interface{}

    // Database statistics for IPv4 table.
    DatabaseStatisticsIpv4 TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4

    // Database statistics for Tunnel table.
    DatabaseStatisticsTunnel TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel
}

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetEntityData() *types.CommonEntityData {
    vrfStatistic.EntityData.YFilter = vrfStatistic.YFilter
    vrfStatistic.EntityData.YangName = "vrf-statistic"
    vrfStatistic.EntityData.BundleName = "cisco_ios_xr"
    vrfStatistic.EntityData.ParentYangName = "summary"
    vrfStatistic.EntityData.SegmentPath = "vrf-statistic"
    vrfStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfStatistic.EntityData.Children = make(map[string]types.YChild)
    vrfStatistic.EntityData.Children["database-statistics-ipv4"] = types.YChild{"DatabaseStatisticsIpv4", &vrfStatistic.DatabaseStatisticsIpv4}
    vrfStatistic.EntityData.Children["database-statistics-tunnel"] = types.YChild{"DatabaseStatisticsTunnel", &vrfStatistic.DatabaseStatisticsTunnel}
    vrfStatistic.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfStatistic.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrfStatistic.VrfName}
    return &(vrfStatistic.EntityData)
}

// TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4
// Database statistics for IPv4 table
type TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of DB entries. The type is interface{} with range: 0..4294967295.
    NumberOfEntries interface{}

    // Number of stale  entries. The type is interface{} with range:
    // 0..4294967295.
    NumberOfStaleEntries interface{}

    // Number of add operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfAddOPerations interface{}

    // Number of delete operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfDeleteOPerations interface{}
}

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetEntityData() *types.CommonEntityData {
    databaseStatisticsIpv4.EntityData.YFilter = databaseStatisticsIpv4.YFilter
    databaseStatisticsIpv4.EntityData.YangName = "database-statistics-ipv4"
    databaseStatisticsIpv4.EntityData.BundleName = "cisco_ios_xr"
    databaseStatisticsIpv4.EntityData.ParentYangName = "vrf-statistic"
    databaseStatisticsIpv4.EntityData.SegmentPath = "database-statistics-ipv4"
    databaseStatisticsIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseStatisticsIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseStatisticsIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseStatisticsIpv4.EntityData.Children = make(map[string]types.YChild)
    databaseStatisticsIpv4.EntityData.Leafs = make(map[string]types.YLeaf)
    databaseStatisticsIpv4.EntityData.Leafs["number-of-entries"] = types.YLeaf{"NumberOfEntries", databaseStatisticsIpv4.NumberOfEntries}
    databaseStatisticsIpv4.EntityData.Leafs["number-of-stale-entries"] = types.YLeaf{"NumberOfStaleEntries", databaseStatisticsIpv4.NumberOfStaleEntries}
    databaseStatisticsIpv4.EntityData.Leafs["number-of-add-o-perations"] = types.YLeaf{"NumberOfAddOPerations", databaseStatisticsIpv4.NumberOfAddOPerations}
    databaseStatisticsIpv4.EntityData.Leafs["number-of-delete-o-perations"] = types.YLeaf{"NumberOfDeleteOPerations", databaseStatisticsIpv4.NumberOfDeleteOPerations}
    return &(databaseStatisticsIpv4.EntityData)
}

// TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel
// Database statistics for Tunnel table
type TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of DB entries. The type is interface{} with range: 0..4294967295.
    NumberOfEntries interface{}

    // Number of stale  entries. The type is interface{} with range:
    // 0..4294967295.
    NumberOfStaleEntries interface{}

    // Number of add operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfAddOPerations interface{}

    // Number of delete operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfDeleteOPerations interface{}
}

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetEntityData() *types.CommonEntityData {
    databaseStatisticsTunnel.EntityData.YFilter = databaseStatisticsTunnel.YFilter
    databaseStatisticsTunnel.EntityData.YangName = "database-statistics-tunnel"
    databaseStatisticsTunnel.EntityData.BundleName = "cisco_ios_xr"
    databaseStatisticsTunnel.EntityData.ParentYangName = "vrf-statistic"
    databaseStatisticsTunnel.EntityData.SegmentPath = "database-statistics-tunnel"
    databaseStatisticsTunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseStatisticsTunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseStatisticsTunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseStatisticsTunnel.EntityData.Children = make(map[string]types.YChild)
    databaseStatisticsTunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    databaseStatisticsTunnel.EntityData.Leafs["number-of-entries"] = types.YLeaf{"NumberOfEntries", databaseStatisticsTunnel.NumberOfEntries}
    databaseStatisticsTunnel.EntityData.Leafs["number-of-stale-entries"] = types.YLeaf{"NumberOfStaleEntries", databaseStatisticsTunnel.NumberOfStaleEntries}
    databaseStatisticsTunnel.EntityData.Leafs["number-of-add-o-perations"] = types.YLeaf{"NumberOfAddOPerations", databaseStatisticsTunnel.NumberOfAddOPerations}
    databaseStatisticsTunnel.EntityData.Leafs["number-of-delete-o-perations"] = types.YLeaf{"NumberOfDeleteOPerations", databaseStatisticsTunnel.NumberOfDeleteOPerations}
    return &(databaseStatisticsTunnel.EntityData)
}

// TrafficCollector_Summary_CollectionMessageStatistic
// Statistics per message type for STAT collector
type TrafficCollector_Summary_CollectionMessageStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketSent interface{}

    // Number of bytes sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ByteSent interface{}

    // Number of packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketReceived interface{}

    // Number of bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ByteReceived interface{}

    // Maximum roundtrip latency in msec. The type is interface{} with range:
    // 0..4294967295.
    MaximumRoundtripLatency interface{}

    // Timestamp of maximum latency. The type is interface{} with range:
    // 0..18446744073709551615.
    MaimumLatencyTimestamp interface{}
}

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetEntityData() *types.CommonEntityData {
    collectionMessageStatistic.EntityData.YFilter = collectionMessageStatistic.YFilter
    collectionMessageStatistic.EntityData.YangName = "collection-message-statistic"
    collectionMessageStatistic.EntityData.BundleName = "cisco_ios_xr"
    collectionMessageStatistic.EntityData.ParentYangName = "summary"
    collectionMessageStatistic.EntityData.SegmentPath = "collection-message-statistic"
    collectionMessageStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collectionMessageStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collectionMessageStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collectionMessageStatistic.EntityData.Children = make(map[string]types.YChild)
    collectionMessageStatistic.EntityData.Leafs = make(map[string]types.YLeaf)
    collectionMessageStatistic.EntityData.Leafs["packet-sent"] = types.YLeaf{"PacketSent", collectionMessageStatistic.PacketSent}
    collectionMessageStatistic.EntityData.Leafs["byte-sent"] = types.YLeaf{"ByteSent", collectionMessageStatistic.ByteSent}
    collectionMessageStatistic.EntityData.Leafs["packet-received"] = types.YLeaf{"PacketReceived", collectionMessageStatistic.PacketReceived}
    collectionMessageStatistic.EntityData.Leafs["byte-received"] = types.YLeaf{"ByteReceived", collectionMessageStatistic.ByteReceived}
    collectionMessageStatistic.EntityData.Leafs["maximum-roundtrip-latency"] = types.YLeaf{"MaximumRoundtripLatency", collectionMessageStatistic.MaximumRoundtripLatency}
    collectionMessageStatistic.EntityData.Leafs["maimum-latency-timestamp"] = types.YLeaf{"MaimumLatencyTimestamp", collectionMessageStatistic.MaimumLatencyTimestamp}
    return &(collectionMessageStatistic.EntityData)
}

// TrafficCollector_Summary_CheckpointMessageStatistic
// Statistics per message type for Chkpt
type TrafficCollector_Summary_CheckpointMessageStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketSent interface{}

    // Number of bytes sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ByteSent interface{}

    // Number of packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketReceived interface{}

    // Number of bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ByteReceived interface{}

    // Maximum roundtrip latency in msec. The type is interface{} with range:
    // 0..4294967295.
    MaximumRoundtripLatency interface{}

    // Timestamp of maximum latency. The type is interface{} with range:
    // 0..18446744073709551615.
    MaimumLatencyTimestamp interface{}
}

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetEntityData() *types.CommonEntityData {
    checkpointMessageStatistic.EntityData.YFilter = checkpointMessageStatistic.YFilter
    checkpointMessageStatistic.EntityData.YangName = "checkpoint-message-statistic"
    checkpointMessageStatistic.EntityData.BundleName = "cisco_ios_xr"
    checkpointMessageStatistic.EntityData.ParentYangName = "summary"
    checkpointMessageStatistic.EntityData.SegmentPath = "checkpoint-message-statistic"
    checkpointMessageStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    checkpointMessageStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    checkpointMessageStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    checkpointMessageStatistic.EntityData.Children = make(map[string]types.YChild)
    checkpointMessageStatistic.EntityData.Leafs = make(map[string]types.YLeaf)
    checkpointMessageStatistic.EntityData.Leafs["packet-sent"] = types.YLeaf{"PacketSent", checkpointMessageStatistic.PacketSent}
    checkpointMessageStatistic.EntityData.Leafs["byte-sent"] = types.YLeaf{"ByteSent", checkpointMessageStatistic.ByteSent}
    checkpointMessageStatistic.EntityData.Leafs["packet-received"] = types.YLeaf{"PacketReceived", checkpointMessageStatistic.PacketReceived}
    checkpointMessageStatistic.EntityData.Leafs["byte-received"] = types.YLeaf{"ByteReceived", checkpointMessageStatistic.ByteReceived}
    checkpointMessageStatistic.EntityData.Leafs["maximum-roundtrip-latency"] = types.YLeaf{"MaximumRoundtripLatency", checkpointMessageStatistic.MaximumRoundtripLatency}
    checkpointMessageStatistic.EntityData.Leafs["maimum-latency-timestamp"] = types.YLeaf{"MaimumLatencyTimestamp", checkpointMessageStatistic.MaimumLatencyTimestamp}
    return &(checkpointMessageStatistic.EntityData)
}

// TrafficCollector_VrfTable
// VRF specific operational data
type TrafficCollector_VrfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DefaultVRF specific operational data.
    DefaultVrf TrafficCollector_VrfTable_DefaultVrf
}

func (vrfTable *TrafficCollector_VrfTable) GetEntityData() *types.CommonEntityData {
    vrfTable.EntityData.YFilter = vrfTable.YFilter
    vrfTable.EntityData.YangName = "vrf-table"
    vrfTable.EntityData.BundleName = "cisco_ios_xr"
    vrfTable.EntityData.ParentYangName = "traffic-collector"
    vrfTable.EntityData.SegmentPath = "vrf-table"
    vrfTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfTable.EntityData.Children = make(map[string]types.YChild)
    vrfTable.EntityData.Children["default-vrf"] = types.YChild{"DefaultVrf", &vrfTable.DefaultVrf}
    vrfTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfTable.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf
// DefaultVRF specific operational data
type TrafficCollector_VrfTable_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family specific operational data.
    Afs TrafficCollector_VrfTable_DefaultVrf_Afs
}

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "vrf-table"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = make(map[string]types.YChild)
    defaultVrf.EntityData.Children["afs"] = types.YChild{"Afs", &defaultVrf.Afs}
    defaultVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(defaultVrf.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs
// Address Family specific operational data
type TrafficCollector_VrfTable_DefaultVrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given Address Family. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af.
    Af []TrafficCollector_VrfTable_DefaultVrf_Afs_Af
}

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "default-vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = make(map[string]types.YChild)
    afs.EntityData.Children["af"] = types.YChild{"Af", nil}
    for i := range afs.Af {
        afs.EntityData.Children[types.GetSegmentPath(&afs.Af[i])] = types.YChild{"Af", &afs.Af[i]}
    }
    afs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(afs.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af
// Operational data for given Address Family
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is TcOperAfName.
    AfName interface{}

    // Show Counters.
    Counters TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters
}

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = make(map[string]types.YChild)
    af.EntityData.Children["counters"] = types.YChild{"Counters", &af.Counters}
    af.EntityData.Leafs = make(map[string]types.YLeaf)
    af.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", af.AfName}
    return &(af.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters
// Show Counters
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix Database.
    Prefixes TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes

    // Tunnels.
    Tunnels TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels
}

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "af"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Children["prefixes"] = types.YChild{"Prefixes", &counters.Prefixes}
    counters.EntityData.Children["tunnels"] = types.YChild{"Tunnels", &counters.Tunnels}
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(counters.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes
// Prefix Database
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show Prefix Counter. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix.
    Prefix []TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix
}

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetEntityData() *types.CommonEntityData {
    prefixes.EntityData.YFilter = prefixes.YFilter
    prefixes.EntityData.YangName = "prefixes"
    prefixes.EntityData.BundleName = "cisco_ios_xr"
    prefixes.EntityData.ParentYangName = "counters"
    prefixes.EntityData.SegmentPath = "prefixes"
    prefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixes.EntityData.Children = make(map[string]types.YChild)
    prefixes.EntityData.Children["prefix"] = types.YChild{"Prefix", nil}
    for i := range prefixes.Prefix {
        prefixes.EntityData.Children[types.GetSegmentPath(&prefixes.Prefix[i])] = types.YChild{"Prefix", &prefixes.Prefix[i]}
    }
    prefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixes.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix
// Show Prefix Counter
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Ipaddr interface{}

    // Prefix Mask. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Mask interface{}

    // Local Label. The type is interface{} with range: 16..1048575.
    Label interface{}

    // Prefix Address (V4 or V6 Format). The type is string.
    Prefix interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    LabelXr interface{}

    // Prefix is Active and collecting new Statistics. The type is bool.
    IsActive interface{}

    // Base counter statistics.
    BaseCounterStatistics TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics

    // Traffic Matrix (TM) counter statistics.
    TrafficMatrixCounterStatistics TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
}

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefixes"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Children["base-counter-statistics"] = types.YChild{"BaseCounterStatistics", &prefix.BaseCounterStatistics}
    prefix.EntityData.Children["traffic-matrix-counter-statistics"] = types.YChild{"TrafficMatrixCounterStatistics", &prefix.TrafficMatrixCounterStatistics}
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["ipaddr"] = types.YLeaf{"Ipaddr", prefix.Ipaddr}
    prefix.EntityData.Leafs["mask"] = types.YLeaf{"Mask", prefix.Mask}
    prefix.EntityData.Leafs["label"] = types.YLeaf{"Label", prefix.Label}
    prefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", prefix.Prefix}
    prefix.EntityData.Leafs["label-xr"] = types.YLeaf{"LabelXr", prefix.LabelXr}
    prefix.EntityData.Leafs["is-active"] = types.YLeaf{"IsActive", prefix.IsActive}
    return &(prefix.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetEntityData() *types.CommonEntityData {
    baseCounterStatistics.EntityData.YFilter = baseCounterStatistics.YFilter
    baseCounterStatistics.EntityData.YangName = "base-counter-statistics"
    baseCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    baseCounterStatistics.EntityData.ParentYangName = "prefix"
    baseCounterStatistics.EntityData.SegmentPath = "base-counter-statistics"
    baseCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseCounterStatistics.EntityData.Children = make(map[string]types.YChild)
    baseCounterStatistics.EntityData.Children["count-history"] = types.YChild{"CountHistory", nil}
    for i := range baseCounterStatistics.CountHistory {
        baseCounterStatistics.EntityData.Children[types.GetSegmentPath(&baseCounterStatistics.CountHistory[i])] = types.YChild{"CountHistory", &baseCounterStatistics.CountHistory[i]}
    }
    baseCounterStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    baseCounterStatistics.EntityData.Leafs["transmit-packets-per-second-switched"] = types.YLeaf{"TransmitPacketsPerSecondSwitched", baseCounterStatistics.TransmitPacketsPerSecondSwitched}
    baseCounterStatistics.EntityData.Leafs["transmit-bytes-per-second-switched"] = types.YLeaf{"TransmitBytesPerSecondSwitched", baseCounterStatistics.TransmitBytesPerSecondSwitched}
    return &(baseCounterStatistics.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "base-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history"
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = make(map[string]types.YChild)
    countHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    countHistory.EntityData.Leafs["event-start-timestamp"] = types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp}
    countHistory.EntityData.Leafs["event-end-timestamp"] = types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp}
    countHistory.EntityData.Leafs["transmit-number-of-packets-switched"] = types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched}
    countHistory.EntityData.Leafs["transmit-number-of-bytes-switched"] = types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched}
    countHistory.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", countHistory.IsValid}
    return &(countHistory.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
// Traffic Matrix (TM) counter statistics
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
}

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetEntityData() *types.CommonEntityData {
    trafficMatrixCounterStatistics.EntityData.YFilter = trafficMatrixCounterStatistics.YFilter
    trafficMatrixCounterStatistics.EntityData.YangName = "traffic-matrix-counter-statistics"
    trafficMatrixCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    trafficMatrixCounterStatistics.EntityData.ParentYangName = "prefix"
    trafficMatrixCounterStatistics.EntityData.SegmentPath = "traffic-matrix-counter-statistics"
    trafficMatrixCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficMatrixCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficMatrixCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficMatrixCounterStatistics.EntityData.Children = make(map[string]types.YChild)
    trafficMatrixCounterStatistics.EntityData.Children["count-history"] = types.YChild{"CountHistory", nil}
    for i := range trafficMatrixCounterStatistics.CountHistory {
        trafficMatrixCounterStatistics.EntityData.Children[types.GetSegmentPath(&trafficMatrixCounterStatistics.CountHistory[i])] = types.YChild{"CountHistory", &trafficMatrixCounterStatistics.CountHistory[i]}
    }
    trafficMatrixCounterStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficMatrixCounterStatistics.EntityData.Leafs["transmit-packets-per-second-switched"] = types.YLeaf{"TransmitPacketsPerSecondSwitched", trafficMatrixCounterStatistics.TransmitPacketsPerSecondSwitched}
    trafficMatrixCounterStatistics.EntityData.Leafs["transmit-bytes-per-second-switched"] = types.YLeaf{"TransmitBytesPerSecondSwitched", trafficMatrixCounterStatistics.TransmitBytesPerSecondSwitched}
    return &(trafficMatrixCounterStatistics.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
// Counter History
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "traffic-matrix-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history"
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = make(map[string]types.YChild)
    countHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    countHistory.EntityData.Leafs["event-start-timestamp"] = types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp}
    countHistory.EntityData.Leafs["event-end-timestamp"] = types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp}
    countHistory.EntityData.Leafs["transmit-number-of-packets-switched"] = types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched}
    countHistory.EntityData.Leafs["transmit-number-of-bytes-switched"] = types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched}
    countHistory.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", countHistory.IsValid}
    return &(countHistory.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels
// Tunnels
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel.
    Tunnel []TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel
}

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "counters"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = make(map[string]types.YChild)
    tunnels.EntityData.Children["tunnel"] = types.YChild{"Tunnel", nil}
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children[types.GetSegmentPath(&tunnels.Tunnel[i])] = types.YChild{"Tunnel", &tunnels.Tunnel[i]}
    }
    tunnels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnels.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel
// Tunnel information
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Interface Name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name in Display format. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Interface VRF ID. The type is interface{} with range: 0..4294967295.
    Vrfid interface{}

    // Interface is Active and collecting new Statistics. The type is bool.
    IsActive interface{}

    // Base counter statistics.
    BaseCounterStatistics TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics
}

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + "[interface-name='" + fmt.Sprintf("%v", tunnel.InterfaceName) + "']"
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Children["base-counter-statistics"] = types.YChild{"BaseCounterStatistics", &tunnel.BaseCounterStatistics}
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnel.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", tunnel.InterfaceName}
    tunnel.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", tunnel.InterfaceNameXr}
    tunnel.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", tunnel.InterfaceHandle}
    tunnel.EntityData.Leafs["vrfid"] = types.YLeaf{"Vrfid", tunnel.Vrfid}
    tunnel.EntityData.Leafs["is-active"] = types.YLeaf{"IsActive", tunnel.IsActive}
    return &(tunnel.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetEntityData() *types.CommonEntityData {
    baseCounterStatistics.EntityData.YFilter = baseCounterStatistics.YFilter
    baseCounterStatistics.EntityData.YangName = "base-counter-statistics"
    baseCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    baseCounterStatistics.EntityData.ParentYangName = "tunnel"
    baseCounterStatistics.EntityData.SegmentPath = "base-counter-statistics"
    baseCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseCounterStatistics.EntityData.Children = make(map[string]types.YChild)
    baseCounterStatistics.EntityData.Children["count-history"] = types.YChild{"CountHistory", nil}
    for i := range baseCounterStatistics.CountHistory {
        baseCounterStatistics.EntityData.Children[types.GetSegmentPath(&baseCounterStatistics.CountHistory[i])] = types.YChild{"CountHistory", &baseCounterStatistics.CountHistory[i]}
    }
    baseCounterStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    baseCounterStatistics.EntityData.Leafs["transmit-packets-per-second-switched"] = types.YLeaf{"TransmitPacketsPerSecondSwitched", baseCounterStatistics.TransmitPacketsPerSecondSwitched}
    baseCounterStatistics.EntityData.Leafs["transmit-bytes-per-second-switched"] = types.YLeaf{"TransmitBytesPerSecondSwitched", baseCounterStatistics.TransmitBytesPerSecondSwitched}
    return &(baseCounterStatistics.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "base-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history"
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = make(map[string]types.YChild)
    countHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    countHistory.EntityData.Leafs["event-start-timestamp"] = types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp}
    countHistory.EntityData.Leafs["event-end-timestamp"] = types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp}
    countHistory.EntityData.Leafs["transmit-number-of-packets-switched"] = types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched}
    countHistory.EntityData.Leafs["transmit-number-of-bytes-switched"] = types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched}
    countHistory.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", countHistory.IsValid}
    return &(countHistory.EntityData)
}

// TrafficCollector_Afs
// Address Family specific operational data
type TrafficCollector_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given Address Family. The type is slice of
    // TrafficCollector_Afs_Af.
    Af []TrafficCollector_Afs_Af
}

func (afs *TrafficCollector_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "traffic-collector"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = make(map[string]types.YChild)
    afs.EntityData.Children["af"] = types.YChild{"Af", nil}
    for i := range afs.Af {
        afs.EntityData.Children[types.GetSegmentPath(&afs.Af[i])] = types.YChild{"Af", &afs.Af[i]}
    }
    afs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(afs.EntityData)
}

// TrafficCollector_Afs_Af
// Operational data for given Address Family
type TrafficCollector_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is TcOperAfName.
    AfName interface{}

    // Show Counters.
    Counters TrafficCollector_Afs_Af_Counters
}

func (af *TrafficCollector_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = make(map[string]types.YChild)
    af.EntityData.Children["counters"] = types.YChild{"Counters", &af.Counters}
    af.EntityData.Leafs = make(map[string]types.YLeaf)
    af.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", af.AfName}
    return &(af.EntityData)
}

// TrafficCollector_Afs_Af_Counters
// Show Counters
type TrafficCollector_Afs_Af_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix Database.
    Prefixes TrafficCollector_Afs_Af_Counters_Prefixes

    // Tunnels.
    Tunnels TrafficCollector_Afs_Af_Counters_Tunnels
}

func (counters *TrafficCollector_Afs_Af_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "af"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Children["prefixes"] = types.YChild{"Prefixes", &counters.Prefixes}
    counters.EntityData.Children["tunnels"] = types.YChild{"Tunnels", &counters.Tunnels}
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(counters.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes
// Prefix Database
type TrafficCollector_Afs_Af_Counters_Prefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show Prefix Counter. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Prefixes_Prefix.
    Prefix []TrafficCollector_Afs_Af_Counters_Prefixes_Prefix
}

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetEntityData() *types.CommonEntityData {
    prefixes.EntityData.YFilter = prefixes.YFilter
    prefixes.EntityData.YangName = "prefixes"
    prefixes.EntityData.BundleName = "cisco_ios_xr"
    prefixes.EntityData.ParentYangName = "counters"
    prefixes.EntityData.SegmentPath = "prefixes"
    prefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixes.EntityData.Children = make(map[string]types.YChild)
    prefixes.EntityData.Children["prefix"] = types.YChild{"Prefix", nil}
    for i := range prefixes.Prefix {
        prefixes.EntityData.Children[types.GetSegmentPath(&prefixes.Prefix[i])] = types.YChild{"Prefix", &prefixes.Prefix[i]}
    }
    prefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixes.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix
// Show Prefix Counter
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Ipaddr interface{}

    // Prefix Mask. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Mask interface{}

    // Local Label. The type is interface{} with range: 16..1048575.
    Label interface{}

    // Prefix Address (V4 or V6 Format). The type is string.
    Prefix interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    LabelXr interface{}

    // Prefix is Active and collecting new Statistics. The type is bool.
    IsActive interface{}

    // Base counter statistics.
    BaseCounterStatistics TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics

    // Traffic Matrix (TM) counter statistics.
    TrafficMatrixCounterStatistics TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
}

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefixes"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Children["base-counter-statistics"] = types.YChild{"BaseCounterStatistics", &prefix.BaseCounterStatistics}
    prefix.EntityData.Children["traffic-matrix-counter-statistics"] = types.YChild{"TrafficMatrixCounterStatistics", &prefix.TrafficMatrixCounterStatistics}
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["ipaddr"] = types.YLeaf{"Ipaddr", prefix.Ipaddr}
    prefix.EntityData.Leafs["mask"] = types.YLeaf{"Mask", prefix.Mask}
    prefix.EntityData.Leafs["label"] = types.YLeaf{"Label", prefix.Label}
    prefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", prefix.Prefix}
    prefix.EntityData.Leafs["label-xr"] = types.YLeaf{"LabelXr", prefix.LabelXr}
    prefix.EntityData.Leafs["is-active"] = types.YLeaf{"IsActive", prefix.IsActive}
    return &(prefix.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetEntityData() *types.CommonEntityData {
    baseCounterStatistics.EntityData.YFilter = baseCounterStatistics.YFilter
    baseCounterStatistics.EntityData.YangName = "base-counter-statistics"
    baseCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    baseCounterStatistics.EntityData.ParentYangName = "prefix"
    baseCounterStatistics.EntityData.SegmentPath = "base-counter-statistics"
    baseCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseCounterStatistics.EntityData.Children = make(map[string]types.YChild)
    baseCounterStatistics.EntityData.Children["count-history"] = types.YChild{"CountHistory", nil}
    for i := range baseCounterStatistics.CountHistory {
        baseCounterStatistics.EntityData.Children[types.GetSegmentPath(&baseCounterStatistics.CountHistory[i])] = types.YChild{"CountHistory", &baseCounterStatistics.CountHistory[i]}
    }
    baseCounterStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    baseCounterStatistics.EntityData.Leafs["transmit-packets-per-second-switched"] = types.YLeaf{"TransmitPacketsPerSecondSwitched", baseCounterStatistics.TransmitPacketsPerSecondSwitched}
    baseCounterStatistics.EntityData.Leafs["transmit-bytes-per-second-switched"] = types.YLeaf{"TransmitBytesPerSecondSwitched", baseCounterStatistics.TransmitBytesPerSecondSwitched}
    return &(baseCounterStatistics.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "base-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history"
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = make(map[string]types.YChild)
    countHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    countHistory.EntityData.Leafs["event-start-timestamp"] = types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp}
    countHistory.EntityData.Leafs["event-end-timestamp"] = types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp}
    countHistory.EntityData.Leafs["transmit-number-of-packets-switched"] = types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched}
    countHistory.EntityData.Leafs["transmit-number-of-bytes-switched"] = types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched}
    countHistory.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", countHistory.IsValid}
    return &(countHistory.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
// Traffic Matrix (TM) counter statistics
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
}

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetEntityData() *types.CommonEntityData {
    trafficMatrixCounterStatistics.EntityData.YFilter = trafficMatrixCounterStatistics.YFilter
    trafficMatrixCounterStatistics.EntityData.YangName = "traffic-matrix-counter-statistics"
    trafficMatrixCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    trafficMatrixCounterStatistics.EntityData.ParentYangName = "prefix"
    trafficMatrixCounterStatistics.EntityData.SegmentPath = "traffic-matrix-counter-statistics"
    trafficMatrixCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficMatrixCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficMatrixCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficMatrixCounterStatistics.EntityData.Children = make(map[string]types.YChild)
    trafficMatrixCounterStatistics.EntityData.Children["count-history"] = types.YChild{"CountHistory", nil}
    for i := range trafficMatrixCounterStatistics.CountHistory {
        trafficMatrixCounterStatistics.EntityData.Children[types.GetSegmentPath(&trafficMatrixCounterStatistics.CountHistory[i])] = types.YChild{"CountHistory", &trafficMatrixCounterStatistics.CountHistory[i]}
    }
    trafficMatrixCounterStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficMatrixCounterStatistics.EntityData.Leafs["transmit-packets-per-second-switched"] = types.YLeaf{"TransmitPacketsPerSecondSwitched", trafficMatrixCounterStatistics.TransmitPacketsPerSecondSwitched}
    trafficMatrixCounterStatistics.EntityData.Leafs["transmit-bytes-per-second-switched"] = types.YLeaf{"TransmitBytesPerSecondSwitched", trafficMatrixCounterStatistics.TransmitBytesPerSecondSwitched}
    return &(trafficMatrixCounterStatistics.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
// Counter History
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "traffic-matrix-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history"
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = make(map[string]types.YChild)
    countHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    countHistory.EntityData.Leafs["event-start-timestamp"] = types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp}
    countHistory.EntityData.Leafs["event-end-timestamp"] = types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp}
    countHistory.EntityData.Leafs["transmit-number-of-packets-switched"] = types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched}
    countHistory.EntityData.Leafs["transmit-number-of-bytes-switched"] = types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched}
    countHistory.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", countHistory.IsValid}
    return &(countHistory.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Tunnels
// Tunnels
type TrafficCollector_Afs_Af_Counters_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel.
    Tunnel []TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel
}

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "counters"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = make(map[string]types.YChild)
    tunnels.EntityData.Children["tunnel"] = types.YChild{"Tunnel", nil}
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children[types.GetSegmentPath(&tunnels.Tunnel[i])] = types.YChild{"Tunnel", &tunnels.Tunnel[i]}
    }
    tunnels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnels.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel
// Tunnel information
type TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Interface Name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name in Display format. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Interface VRF ID. The type is interface{} with range: 0..4294967295.
    Vrfid interface{}

    // Interface is Active and collecting new Statistics. The type is bool.
    IsActive interface{}

    // Base counter statistics.
    BaseCounterStatistics TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics
}

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + "[interface-name='" + fmt.Sprintf("%v", tunnel.InterfaceName) + "']"
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Children["base-counter-statistics"] = types.YChild{"BaseCounterStatistics", &tunnel.BaseCounterStatistics}
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnel.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", tunnel.InterfaceName}
    tunnel.EntityData.Leafs["interface-name-xr"] = types.YLeaf{"InterfaceNameXr", tunnel.InterfaceNameXr}
    tunnel.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", tunnel.InterfaceHandle}
    tunnel.EntityData.Leafs["vrfid"] = types.YLeaf{"Vrfid", tunnel.Vrfid}
    tunnel.EntityData.Leafs["is-active"] = types.YLeaf{"IsActive", tunnel.IsActive}
    return &(tunnel.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetEntityData() *types.CommonEntityData {
    baseCounterStatistics.EntityData.YFilter = baseCounterStatistics.YFilter
    baseCounterStatistics.EntityData.YangName = "base-counter-statistics"
    baseCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    baseCounterStatistics.EntityData.ParentYangName = "tunnel"
    baseCounterStatistics.EntityData.SegmentPath = "base-counter-statistics"
    baseCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseCounterStatistics.EntityData.Children = make(map[string]types.YChild)
    baseCounterStatistics.EntityData.Children["count-history"] = types.YChild{"CountHistory", nil}
    for i := range baseCounterStatistics.CountHistory {
        baseCounterStatistics.EntityData.Children[types.GetSegmentPath(&baseCounterStatistics.CountHistory[i])] = types.YChild{"CountHistory", &baseCounterStatistics.CountHistory[i]}
    }
    baseCounterStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    baseCounterStatistics.EntityData.Leafs["transmit-packets-per-second-switched"] = types.YLeaf{"TransmitPacketsPerSecondSwitched", baseCounterStatistics.TransmitPacketsPerSecondSwitched}
    baseCounterStatistics.EntityData.Leafs["transmit-bytes-per-second-switched"] = types.YLeaf{"TransmitBytesPerSecondSwitched", baseCounterStatistics.TransmitBytesPerSecondSwitched}
    return &(baseCounterStatistics.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "base-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history"
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = make(map[string]types.YChild)
    countHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    countHistory.EntityData.Leafs["event-start-timestamp"] = types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp}
    countHistory.EntityData.Leafs["event-end-timestamp"] = types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp}
    countHistory.EntityData.Leafs["transmit-number-of-packets-switched"] = types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched}
    countHistory.EntityData.Leafs["transmit-number-of-bytes-switched"] = types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched}
    countHistory.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", countHistory.IsValid}
    return &(countHistory.EntityData)
}

