// This module contains a collection of YANG definitions
// for Cisco IOS-XR traffmon-netflow package configuration.
// 
// This module contains definitions
// for the following management objects:
//   net-flow: NetFlow Configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package traffmon_netflow_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package traffmon_netflow_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-traffmon-netflow-cfg net-flow}", reflect.TypeOf(NetFlow{}))
    ydk.RegisterEntity("Cisco-IOS-XR-traffmon-netflow-cfg:net-flow", reflect.TypeOf(NetFlow{}))
}

// NfCacheAgingMode represents Nf cache aging mode
type NfCacheAgingMode string

const (
    // Normal, caches age
    NfCacheAgingMode_normal NfCacheAgingMode = "normal"

    // Permanent, caches never age
    NfCacheAgingMode_permanent NfCacheAgingMode = "permanent"

    // Immediate, caches age immediately
    NfCacheAgingMode_immediate NfCacheAgingMode = "immediate"
)

// NfSamplingMode represents Nf sampling mode
type NfSamplingMode string

const (
    // Random sampling
    NfSamplingMode_random NfSamplingMode = "random"
)

// NetFlow
// NetFlow Configuration
type NetFlow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a flow exporter map.
    FlowExporterMaps NetFlow_FlowExporterMaps

    // Flow sampler map configuration.
    FlowSamplerMaps NetFlow_FlowSamplerMaps

    // Flow monitor map configuration.
    FlowMonitorMapTable NetFlow_FlowMonitorMapTable

    // Configure a performance traffic flow monitor map.
    FlowMonitorMapPerformanceTable NetFlow_FlowMonitorMapPerformanceTable
}

func (netFlow *NetFlow) GetEntityData() *types.CommonEntityData {
    netFlow.EntityData.YFilter = netFlow.YFilter
    netFlow.EntityData.YangName = "net-flow"
    netFlow.EntityData.BundleName = "cisco_ios_xr"
    netFlow.EntityData.ParentYangName = "Cisco-IOS-XR-traffmon-netflow-cfg"
    netFlow.EntityData.SegmentPath = "Cisco-IOS-XR-traffmon-netflow-cfg:net-flow"
    netFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netFlow.EntityData.Children = make(map[string]types.YChild)
    netFlow.EntityData.Children["flow-exporter-maps"] = types.YChild{"FlowExporterMaps", &netFlow.FlowExporterMaps}
    netFlow.EntityData.Children["flow-sampler-maps"] = types.YChild{"FlowSamplerMaps", &netFlow.FlowSamplerMaps}
    netFlow.EntityData.Children["flow-monitor-map-table"] = types.YChild{"FlowMonitorMapTable", &netFlow.FlowMonitorMapTable}
    netFlow.EntityData.Children["flow-monitor-map-performance-table"] = types.YChild{"FlowMonitorMapPerformanceTable", &netFlow.FlowMonitorMapPerformanceTable}
    netFlow.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(netFlow.EntityData)
}

// NetFlow_FlowExporterMaps
// Configure a flow exporter map
type NetFlow_FlowExporterMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exporter map name. The type is slice of
    // NetFlow_FlowExporterMaps_FlowExporterMap.
    FlowExporterMap []NetFlow_FlowExporterMaps_FlowExporterMap
}

func (flowExporterMaps *NetFlow_FlowExporterMaps) GetEntityData() *types.CommonEntityData {
    flowExporterMaps.EntityData.YFilter = flowExporterMaps.YFilter
    flowExporterMaps.EntityData.YangName = "flow-exporter-maps"
    flowExporterMaps.EntityData.BundleName = "cisco_ios_xr"
    flowExporterMaps.EntityData.ParentYangName = "net-flow"
    flowExporterMaps.EntityData.SegmentPath = "flow-exporter-maps"
    flowExporterMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowExporterMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowExporterMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowExporterMaps.EntityData.Children = make(map[string]types.YChild)
    flowExporterMaps.EntityData.Children["flow-exporter-map"] = types.YChild{"FlowExporterMap", nil}
    for i := range flowExporterMaps.FlowExporterMap {
        flowExporterMaps.EntityData.Children[types.GetSegmentPath(&flowExporterMaps.FlowExporterMap[i])] = types.YChild{"FlowExporterMap", &flowExporterMaps.FlowExporterMap[i]}
    }
    flowExporterMaps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flowExporterMaps.EntityData)
}

// NetFlow_FlowExporterMaps_FlowExporterMap
// Exporter map name
type NetFlow_FlowExporterMaps_FlowExporterMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Exporter map name. The type is string with length:
    // 1..32.
    ExporterMapName interface{}

    // Configure source interface for collector. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SourceInterface interface{}

    // Specify DSCP value for export packets. The type is interface{} with range:
    // 0..63.
    Dscp interface{}

    // Configure Maximum Value for Export Packet size. The type is interface{}
    // with range: 512..1468.
    PacketLength interface{}

    // Use UDP as transport protocol.
    Udp NetFlow_FlowExporterMaps_FlowExporterMap_Udp

    // Configure export destination (collector).
    Destination NetFlow_FlowExporterMaps_FlowExporterMap_Destination

    // Specify export version parameters.
    Version NetFlow_FlowExporterMaps_FlowExporterMap_Version
}

func (flowExporterMap *NetFlow_FlowExporterMaps_FlowExporterMap) GetEntityData() *types.CommonEntityData {
    flowExporterMap.EntityData.YFilter = flowExporterMap.YFilter
    flowExporterMap.EntityData.YangName = "flow-exporter-map"
    flowExporterMap.EntityData.BundleName = "cisco_ios_xr"
    flowExporterMap.EntityData.ParentYangName = "flow-exporter-maps"
    flowExporterMap.EntityData.SegmentPath = "flow-exporter-map" + "[exporter-map-name='" + fmt.Sprintf("%v", flowExporterMap.ExporterMapName) + "']"
    flowExporterMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowExporterMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowExporterMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowExporterMap.EntityData.Children = make(map[string]types.YChild)
    flowExporterMap.EntityData.Children["udp"] = types.YChild{"Udp", &flowExporterMap.Udp}
    flowExporterMap.EntityData.Children["destination"] = types.YChild{"Destination", &flowExporterMap.Destination}
    flowExporterMap.EntityData.Children["version"] = types.YChild{"Version", &flowExporterMap.Version}
    flowExporterMap.EntityData.Leafs = make(map[string]types.YLeaf)
    flowExporterMap.EntityData.Leafs["exporter-map-name"] = types.YLeaf{"ExporterMapName", flowExporterMap.ExporterMapName}
    flowExporterMap.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", flowExporterMap.SourceInterface}
    flowExporterMap.EntityData.Leafs["dscp"] = types.YLeaf{"Dscp", flowExporterMap.Dscp}
    flowExporterMap.EntityData.Leafs["packet-length"] = types.YLeaf{"PacketLength", flowExporterMap.PacketLength}
    return &(flowExporterMap.EntityData)
}

// NetFlow_FlowExporterMaps_FlowExporterMap_Udp
// Use UDP as transport protocol
type NetFlow_FlowExporterMaps_FlowExporterMap_Udp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Destination UDP port. The type is interface{} with range:
    // 1024..65535.
    DestinationPort interface{}
}

func (udp *NetFlow_FlowExporterMaps_FlowExporterMap_Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xr"
    udp.EntityData.ParentYangName = "flow-exporter-map"
    udp.EntityData.SegmentPath = "udp"
    udp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udp.EntityData.Children = make(map[string]types.YChild)
    udp.EntityData.Leafs = make(map[string]types.YLeaf)
    udp.EntityData.Leafs["destination-port"] = types.YLeaf{"DestinationPort", udp.DestinationPort}
    return &(udp.EntityData)
}

// NetFlow_FlowExporterMaps_FlowExporterMap_Destination
// Configure export destination (collector)
type NetFlow_FlowExporterMaps_FlowExporterMap_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // IPV6 address of the tunnel destination. The type is string.
    Ipv6Address interface{}

    // VRF name. The type is string.
    VrfName interface{}
}

func (destination *NetFlow_FlowExporterMaps_FlowExporterMap_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "cisco_ios_xr"
    destination.EntityData.ParentYangName = "flow-exporter-map"
    destination.EntityData.SegmentPath = "destination"
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = make(map[string]types.YChild)
    destination.EntityData.Leafs = make(map[string]types.YLeaf)
    destination.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", destination.IpAddress}
    destination.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", destination.Ipv6Address}
    destination.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", destination.VrfName}
    return &(destination.EntityData)
}

// NetFlow_FlowExporterMaps_FlowExporterMap_Version
// Specify export version parameters
type NetFlow_FlowExporterMaps_FlowExporterMap_Version struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Export version number. The type is interface{} with range: 9..10.
    VersionType interface{}

    // Option template configuration options. The type is interface{} with range:
    // 1..604800. Units are second. The default value is 1800.
    OptionsTemplateTimeout interface{}

    // Specify custom timeout for the template. The type is interface{} with
    // range: 1..604800. Units are second. The default value is 1800.
    CommonTemplateTimeout interface{}

    // Data template configuration options. The type is interface{} with range:
    // 1..604800. Units are second. The default value is 1800.
    DataTemplateTimeout interface{}

    // Specify options for exporting templates.
    Options NetFlow_FlowExporterMaps_FlowExporterMap_Version_Options
}

func (version *NetFlow_FlowExporterMaps_FlowExporterMap_Version) GetEntityData() *types.CommonEntityData {
    version.EntityData.YFilter = version.YFilter
    version.EntityData.YangName = "version"
    version.EntityData.BundleName = "cisco_ios_xr"
    version.EntityData.ParentYangName = "flow-exporter-map"
    version.EntityData.SegmentPath = "version"
    version.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    version.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    version.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    version.EntityData.Children = make(map[string]types.YChild)
    version.EntityData.Children["options"] = types.YChild{"Options", &version.Options}
    version.EntityData.Leafs = make(map[string]types.YLeaf)
    version.EntityData.Leafs["version-type"] = types.YLeaf{"VersionType", version.VersionType}
    version.EntityData.Leafs["options-template-timeout"] = types.YLeaf{"OptionsTemplateTimeout", version.OptionsTemplateTimeout}
    version.EntityData.Leafs["common-template-timeout"] = types.YLeaf{"CommonTemplateTimeout", version.CommonTemplateTimeout}
    version.EntityData.Leafs["data-template-timeout"] = types.YLeaf{"DataTemplateTimeout", version.DataTemplateTimeout}
    return &(version.EntityData)
}

// NetFlow_FlowExporterMaps_FlowExporterMap_Version_Options
// Specify options for exporting templates
type NetFlow_FlowExporterMaps_FlowExporterMap_Version_Options struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify timeout for exporting interface table. The type is interface{} with
    // range: 0..604800. Units are second.
    InterfaceTableExportTimeout interface{}

    // Specify timeout for exporting sampler table. The type is interface{} with
    // range: 0..604800. Units are second.
    SamplerTableExportTimeout interface{}

    // Specify timeout for exporting vrf table. The type is interface{} with
    // range: 0..604800. Units are second.
    VrfTableExportTimeout interface{}
}

func (options *NetFlow_FlowExporterMaps_FlowExporterMap_Version_Options) GetEntityData() *types.CommonEntityData {
    options.EntityData.YFilter = options.YFilter
    options.EntityData.YangName = "options"
    options.EntityData.BundleName = "cisco_ios_xr"
    options.EntityData.ParentYangName = "version"
    options.EntityData.SegmentPath = "options"
    options.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    options.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    options.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    options.EntityData.Children = make(map[string]types.YChild)
    options.EntityData.Leafs = make(map[string]types.YLeaf)
    options.EntityData.Leafs["interface-table-export-timeout"] = types.YLeaf{"InterfaceTableExportTimeout", options.InterfaceTableExportTimeout}
    options.EntityData.Leafs["sampler-table-export-timeout"] = types.YLeaf{"SamplerTableExportTimeout", options.SamplerTableExportTimeout}
    options.EntityData.Leafs["vrf-table-export-timeout"] = types.YLeaf{"VrfTableExportTimeout", options.VrfTableExportTimeout}
    return &(options.EntityData)
}

// NetFlow_FlowSamplerMaps
// Flow sampler map configuration
type NetFlow_FlowSamplerMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sampler map name. The type is slice of
    // NetFlow_FlowSamplerMaps_FlowSamplerMap.
    FlowSamplerMap []NetFlow_FlowSamplerMaps_FlowSamplerMap
}

func (flowSamplerMaps *NetFlow_FlowSamplerMaps) GetEntityData() *types.CommonEntityData {
    flowSamplerMaps.EntityData.YFilter = flowSamplerMaps.YFilter
    flowSamplerMaps.EntityData.YangName = "flow-sampler-maps"
    flowSamplerMaps.EntityData.BundleName = "cisco_ios_xr"
    flowSamplerMaps.EntityData.ParentYangName = "net-flow"
    flowSamplerMaps.EntityData.SegmentPath = "flow-sampler-maps"
    flowSamplerMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowSamplerMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowSamplerMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowSamplerMaps.EntityData.Children = make(map[string]types.YChild)
    flowSamplerMaps.EntityData.Children["flow-sampler-map"] = types.YChild{"FlowSamplerMap", nil}
    for i := range flowSamplerMaps.FlowSamplerMap {
        flowSamplerMaps.EntityData.Children[types.GetSegmentPath(&flowSamplerMaps.FlowSamplerMap[i])] = types.YChild{"FlowSamplerMap", &flowSamplerMaps.FlowSamplerMap[i]}
    }
    flowSamplerMaps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flowSamplerMaps.EntityData)
}

// NetFlow_FlowSamplerMaps_FlowSamplerMap
// Sampler map name
type NetFlow_FlowSamplerMaps_FlowSamplerMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Sampler map name. The type is string with length:
    // 1..32.
    SamplerMapName interface{}

    // Configure packet sampling mode.
    SamplingModes NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes
}

func (flowSamplerMap *NetFlow_FlowSamplerMaps_FlowSamplerMap) GetEntityData() *types.CommonEntityData {
    flowSamplerMap.EntityData.YFilter = flowSamplerMap.YFilter
    flowSamplerMap.EntityData.YangName = "flow-sampler-map"
    flowSamplerMap.EntityData.BundleName = "cisco_ios_xr"
    flowSamplerMap.EntityData.ParentYangName = "flow-sampler-maps"
    flowSamplerMap.EntityData.SegmentPath = "flow-sampler-map" + "[sampler-map-name='" + fmt.Sprintf("%v", flowSamplerMap.SamplerMapName) + "']"
    flowSamplerMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowSamplerMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowSamplerMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowSamplerMap.EntityData.Children = make(map[string]types.YChild)
    flowSamplerMap.EntityData.Children["sampling-modes"] = types.YChild{"SamplingModes", &flowSamplerMap.SamplingModes}
    flowSamplerMap.EntityData.Leafs = make(map[string]types.YLeaf)
    flowSamplerMap.EntityData.Leafs["sampler-map-name"] = types.YLeaf{"SamplerMapName", flowSamplerMap.SamplerMapName}
    return &(flowSamplerMap.EntityData)
}

// NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes
// Configure packet sampling mode
type NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure sampling mode. The type is slice of
    // NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode.
    SamplingMode []NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode
}

func (samplingModes *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes) GetEntityData() *types.CommonEntityData {
    samplingModes.EntityData.YFilter = samplingModes.YFilter
    samplingModes.EntityData.YangName = "sampling-modes"
    samplingModes.EntityData.BundleName = "cisco_ios_xr"
    samplingModes.EntityData.ParentYangName = "flow-sampler-map"
    samplingModes.EntityData.SegmentPath = "sampling-modes"
    samplingModes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samplingModes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samplingModes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samplingModes.EntityData.Children = make(map[string]types.YChild)
    samplingModes.EntityData.Children["sampling-mode"] = types.YChild{"SamplingMode", nil}
    for i := range samplingModes.SamplingMode {
        samplingModes.EntityData.Children[types.GetSegmentPath(&samplingModes.SamplingMode[i])] = types.YChild{"SamplingMode", &samplingModes.SamplingMode[i]}
    }
    samplingModes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(samplingModes.EntityData)
}

// NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode
// Configure sampling mode
type NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Sampling mode. The type is NfSamplingMode.
    Mode interface{}

    // Number of packets to be sampled in the sampling interval. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    SampleNumber interface{}

    // Sampling interval in units of packets. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    Interval interface{}
}

func (samplingMode *NetFlow_FlowSamplerMaps_FlowSamplerMap_SamplingModes_SamplingMode) GetEntityData() *types.CommonEntityData {
    samplingMode.EntityData.YFilter = samplingMode.YFilter
    samplingMode.EntityData.YangName = "sampling-mode"
    samplingMode.EntityData.BundleName = "cisco_ios_xr"
    samplingMode.EntityData.ParentYangName = "sampling-modes"
    samplingMode.EntityData.SegmentPath = "sampling-mode" + "[mode='" + fmt.Sprintf("%v", samplingMode.Mode) + "']"
    samplingMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samplingMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samplingMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samplingMode.EntityData.Children = make(map[string]types.YChild)
    samplingMode.EntityData.Leafs = make(map[string]types.YLeaf)
    samplingMode.EntityData.Leafs["mode"] = types.YLeaf{"Mode", samplingMode.Mode}
    samplingMode.EntityData.Leafs["sample-number"] = types.YLeaf{"SampleNumber", samplingMode.SampleNumber}
    samplingMode.EntityData.Leafs["interval"] = types.YLeaf{"Interval", samplingMode.Interval}
    return &(samplingMode.EntityData)
}

// NetFlow_FlowMonitorMapTable
// Flow monitor map configuration
type NetFlow_FlowMonitorMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor map name. The type is slice of
    // NetFlow_FlowMonitorMapTable_FlowMonitorMap.
    FlowMonitorMap []NetFlow_FlowMonitorMapTable_FlowMonitorMap
}

func (flowMonitorMapTable *NetFlow_FlowMonitorMapTable) GetEntityData() *types.CommonEntityData {
    flowMonitorMapTable.EntityData.YFilter = flowMonitorMapTable.YFilter
    flowMonitorMapTable.EntityData.YangName = "flow-monitor-map-table"
    flowMonitorMapTable.EntityData.BundleName = "cisco_ios_xr"
    flowMonitorMapTable.EntityData.ParentYangName = "net-flow"
    flowMonitorMapTable.EntityData.SegmentPath = "flow-monitor-map-table"
    flowMonitorMapTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowMonitorMapTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowMonitorMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowMonitorMapTable.EntityData.Children = make(map[string]types.YChild)
    flowMonitorMapTable.EntityData.Children["flow-monitor-map"] = types.YChild{"FlowMonitorMap", nil}
    for i := range flowMonitorMapTable.FlowMonitorMap {
        flowMonitorMapTable.EntityData.Children[types.GetSegmentPath(&flowMonitorMapTable.FlowMonitorMap[i])] = types.YChild{"FlowMonitorMap", &flowMonitorMapTable.FlowMonitorMap[i]}
    }
    flowMonitorMapTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flowMonitorMapTable.EntityData)
}

// NetFlow_FlowMonitorMapTable_FlowMonitorMap
// Monitor map name
type NetFlow_FlowMonitorMapTable_FlowMonitorMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Monitor map name. The type is string with length:
    // 1..32.
    MonitorMapName interface{}

    // Specify the update flow cache aging timeout. The type is interface{} with
    // range: 1..604800. Units are second. The default value is 1800.
    CacheUpdateAgingTimeout interface{}

    // Specify the number of entries in the flow cache. The type is interface{}
    // with range: 4096..1000000. The default value is 65535.
    CacheEntries interface{}

    // Specify the inactive flow cache aging timeout. The type is interface{} with
    // range: 0..604800. Units are second. The default value is 15.
    CacheInactiveAgingTimeout interface{}

    // Specify the active flow cache aging timeout. The type is interface{} with
    // range: 1..604800. Units are second. The default value is 1800.
    CacheActiveAgingTimeout interface{}

    // Specify the maximum number of entries to age each second. The type is
    // interface{} with range: 1..1000000. The default value is 2000.
    CacheTimeoutRateLimit interface{}

    // Specify the flow cache aging mode. The type is NfCacheAgingMode. The
    // default value is normal.
    CacheAgingMode interface{}

    // Specify an option for the flow cache.
    Option NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option

    // Configure exporters to be used by the monitor-map.
    Exporters NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters

    // Specify a flow record format.
    Record NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record
}

func (flowMonitorMap *NetFlow_FlowMonitorMapTable_FlowMonitorMap) GetEntityData() *types.CommonEntityData {
    flowMonitorMap.EntityData.YFilter = flowMonitorMap.YFilter
    flowMonitorMap.EntityData.YangName = "flow-monitor-map"
    flowMonitorMap.EntityData.BundleName = "cisco_ios_xr"
    flowMonitorMap.EntityData.ParentYangName = "flow-monitor-map-table"
    flowMonitorMap.EntityData.SegmentPath = "flow-monitor-map" + "[monitor-map-name='" + fmt.Sprintf("%v", flowMonitorMap.MonitorMapName) + "']"
    flowMonitorMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowMonitorMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowMonitorMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowMonitorMap.EntityData.Children = make(map[string]types.YChild)
    flowMonitorMap.EntityData.Children["option"] = types.YChild{"Option", &flowMonitorMap.Option}
    flowMonitorMap.EntityData.Children["exporters"] = types.YChild{"Exporters", &flowMonitorMap.Exporters}
    flowMonitorMap.EntityData.Children["record"] = types.YChild{"Record", &flowMonitorMap.Record}
    flowMonitorMap.EntityData.Leafs = make(map[string]types.YLeaf)
    flowMonitorMap.EntityData.Leafs["monitor-map-name"] = types.YLeaf{"MonitorMapName", flowMonitorMap.MonitorMapName}
    flowMonitorMap.EntityData.Leafs["cache-update-aging-timeout"] = types.YLeaf{"CacheUpdateAgingTimeout", flowMonitorMap.CacheUpdateAgingTimeout}
    flowMonitorMap.EntityData.Leafs["cache-entries"] = types.YLeaf{"CacheEntries", flowMonitorMap.CacheEntries}
    flowMonitorMap.EntityData.Leafs["cache-inactive-aging-timeout"] = types.YLeaf{"CacheInactiveAgingTimeout", flowMonitorMap.CacheInactiveAgingTimeout}
    flowMonitorMap.EntityData.Leafs["cache-active-aging-timeout"] = types.YLeaf{"CacheActiveAgingTimeout", flowMonitorMap.CacheActiveAgingTimeout}
    flowMonitorMap.EntityData.Leafs["cache-timeout-rate-limit"] = types.YLeaf{"CacheTimeoutRateLimit", flowMonitorMap.CacheTimeoutRateLimit}
    flowMonitorMap.EntityData.Leafs["cache-aging-mode"] = types.YLeaf{"CacheAgingMode", flowMonitorMap.CacheAgingMode}
    return &(flowMonitorMap.EntityData)
}

// NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option
// Specify an option for the flow cache
type NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether data should be filtered. The type is interface{}.
    Filtered interface{}

    // Specify whether to export physical ifh for bundle interface. The type is
    // interface{}.
    OutBundleMember interface{}

    // Specify whether it exports the physical output interface. The type is
    // interface{}.
    OutPhysInt interface{}

    // Specify if BGP Attributes AS_PATH STD_COMM should be exported. The type is
    // interface{}.
    BgpAttr interface{}
}

func (option *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Option) GetEntityData() *types.CommonEntityData {
    option.EntityData.YFilter = option.YFilter
    option.EntityData.YangName = "option"
    option.EntityData.BundleName = "cisco_ios_xr"
    option.EntityData.ParentYangName = "flow-monitor-map"
    option.EntityData.SegmentPath = "option"
    option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    option.EntityData.Children = make(map[string]types.YChild)
    option.EntityData.Leafs = make(map[string]types.YLeaf)
    option.EntityData.Leafs["filtered"] = types.YLeaf{"Filtered", option.Filtered}
    option.EntityData.Leafs["out-bundle-member"] = types.YLeaf{"OutBundleMember", option.OutBundleMember}
    option.EntityData.Leafs["out-phys-int"] = types.YLeaf{"OutPhysInt", option.OutPhysInt}
    option.EntityData.Leafs["bgp-attr"] = types.YLeaf{"BgpAttr", option.BgpAttr}
    return &(option.EntityData)
}

// NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters
// Configure exporters to be used by the
// monitor-map
type NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure exporter to be used by the monitor-map. The type is slice of
    // NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter.
    Exporter []NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter
}

func (exporters *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters) GetEntityData() *types.CommonEntityData {
    exporters.EntityData.YFilter = exporters.YFilter
    exporters.EntityData.YangName = "exporters"
    exporters.EntityData.BundleName = "cisco_ios_xr"
    exporters.EntityData.ParentYangName = "flow-monitor-map"
    exporters.EntityData.SegmentPath = "exporters"
    exporters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exporters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exporters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exporters.EntityData.Children = make(map[string]types.YChild)
    exporters.EntityData.Children["exporter"] = types.YChild{"Exporter", nil}
    for i := range exporters.Exporter {
        exporters.EntityData.Children[types.GetSegmentPath(&exporters.Exporter[i])] = types.YChild{"Exporter", &exporters.Exporter[i]}
    }
    exporters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(exporters.EntityData)
}

// NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter
// Configure exporter to be used by the
// monitor-map
type NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Exporter name. The type is string with length:
    // 1..32.
    ExporterName interface{}
}

func (exporter *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Exporters_Exporter) GetEntityData() *types.CommonEntityData {
    exporter.EntityData.YFilter = exporter.YFilter
    exporter.EntityData.YangName = "exporter"
    exporter.EntityData.BundleName = "cisco_ios_xr"
    exporter.EntityData.ParentYangName = "exporters"
    exporter.EntityData.SegmentPath = "exporter" + "[exporter-name='" + fmt.Sprintf("%v", exporter.ExporterName) + "']"
    exporter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exporter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exporter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exporter.EntityData.Children = make(map[string]types.YChild)
    exporter.EntityData.Leafs = make(map[string]types.YLeaf)
    exporter.EntityData.Leafs["exporter-name"] = types.YLeaf{"ExporterName", exporter.ExporterName}
    return &(exporter.EntityData)
}

// NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record
// Specify a flow record format
// This type is a presence type.
type NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow record format (Either 'ipv4-raw' ,'ipv4-peer-as', 'ipv6', 'mpls',
    // 'mpls-ipv4', 'mpls-ipv6', 'mpls-ipv4-ipv6', 'ipv6-peer-as'). The type is
    // string with length: 1..32. This attribute is mandatory.
    RecordName interface{}

    // Enter label value for MPLS record type. The type is interface{} with range:
    // 1..6.
    Label interface{}
}

func (record *NetFlow_FlowMonitorMapTable_FlowMonitorMap_Record) GetEntityData() *types.CommonEntityData {
    record.EntityData.YFilter = record.YFilter
    record.EntityData.YangName = "record"
    record.EntityData.BundleName = "cisco_ios_xr"
    record.EntityData.ParentYangName = "flow-monitor-map"
    record.EntityData.SegmentPath = "record"
    record.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    record.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    record.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    record.EntityData.Children = make(map[string]types.YChild)
    record.EntityData.Leafs = make(map[string]types.YLeaf)
    record.EntityData.Leafs["record-name"] = types.YLeaf{"RecordName", record.RecordName}
    record.EntityData.Leafs["label"] = types.YLeaf{"Label", record.Label}
    return &(record.EntityData)
}

// NetFlow_FlowMonitorMapPerformanceTable
// Configure a performance traffic flow monitor map
type NetFlow_FlowMonitorMapPerformanceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor map name. The type is slice of
    // NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap.
    FlowMonitorMap []NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap
}

func (flowMonitorMapPerformanceTable *NetFlow_FlowMonitorMapPerformanceTable) GetEntityData() *types.CommonEntityData {
    flowMonitorMapPerformanceTable.EntityData.YFilter = flowMonitorMapPerformanceTable.YFilter
    flowMonitorMapPerformanceTable.EntityData.YangName = "flow-monitor-map-performance-table"
    flowMonitorMapPerformanceTable.EntityData.BundleName = "cisco_ios_xr"
    flowMonitorMapPerformanceTable.EntityData.ParentYangName = "net-flow"
    flowMonitorMapPerformanceTable.EntityData.SegmentPath = "flow-monitor-map-performance-table"
    flowMonitorMapPerformanceTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowMonitorMapPerformanceTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowMonitorMapPerformanceTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowMonitorMapPerformanceTable.EntityData.Children = make(map[string]types.YChild)
    flowMonitorMapPerformanceTable.EntityData.Children["flow-monitor-map"] = types.YChild{"FlowMonitorMap", nil}
    for i := range flowMonitorMapPerformanceTable.FlowMonitorMap {
        flowMonitorMapPerformanceTable.EntityData.Children[types.GetSegmentPath(&flowMonitorMapPerformanceTable.FlowMonitorMap[i])] = types.YChild{"FlowMonitorMap", &flowMonitorMapPerformanceTable.FlowMonitorMap[i]}
    }
    flowMonitorMapPerformanceTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flowMonitorMapPerformanceTable.EntityData)
}

// NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap
// Monitor map name
type NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Monitor map name. The type is string with length:
    // 1..32.
    MonitorMapName interface{}

    // Specify the update flow cache aging timeout. The type is interface{} with
    // range: 1..604800. Units are second. The default value is 1800.
    CacheUpdateAgingTimeout interface{}

    // Specify the number of entries in the flow cache. The type is interface{}
    // with range: 4096..1000000. The default value is 65535.
    CacheEntries interface{}

    // Specify the inactive flow cache aging timeout. The type is interface{} with
    // range: 0..604800. Units are second. The default value is 15.
    CacheInactiveAgingTimeout interface{}

    // Specify the active flow cache aging timeout. The type is interface{} with
    // range: 1..604800. Units are second. The default value is 1800.
    CacheActiveAgingTimeout interface{}

    // Specify the maximum number of entries to age each second. The type is
    // interface{} with range: 1..1000000. The default value is 2000.
    CacheTimeoutRateLimit interface{}

    // Specify the flow cache aging mode. The type is NfCacheAgingMode. The
    // default value is normal.
    CacheAgingMode interface{}

    // Specify an option for the flow cache.
    Option NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option

    // Configure exporters to be used by the monitor-map.
    Exporters NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters

    // Specify a flow record format.
    Record NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record
}

func (flowMonitorMap *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap) GetEntityData() *types.CommonEntityData {
    flowMonitorMap.EntityData.YFilter = flowMonitorMap.YFilter
    flowMonitorMap.EntityData.YangName = "flow-monitor-map"
    flowMonitorMap.EntityData.BundleName = "cisco_ios_xr"
    flowMonitorMap.EntityData.ParentYangName = "flow-monitor-map-performance-table"
    flowMonitorMap.EntityData.SegmentPath = "flow-monitor-map" + "[monitor-map-name='" + fmt.Sprintf("%v", flowMonitorMap.MonitorMapName) + "']"
    flowMonitorMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowMonitorMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowMonitorMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowMonitorMap.EntityData.Children = make(map[string]types.YChild)
    flowMonitorMap.EntityData.Children["option"] = types.YChild{"Option", &flowMonitorMap.Option}
    flowMonitorMap.EntityData.Children["exporters"] = types.YChild{"Exporters", &flowMonitorMap.Exporters}
    flowMonitorMap.EntityData.Children["record"] = types.YChild{"Record", &flowMonitorMap.Record}
    flowMonitorMap.EntityData.Leafs = make(map[string]types.YLeaf)
    flowMonitorMap.EntityData.Leafs["monitor-map-name"] = types.YLeaf{"MonitorMapName", flowMonitorMap.MonitorMapName}
    flowMonitorMap.EntityData.Leafs["cache-update-aging-timeout"] = types.YLeaf{"CacheUpdateAgingTimeout", flowMonitorMap.CacheUpdateAgingTimeout}
    flowMonitorMap.EntityData.Leafs["cache-entries"] = types.YLeaf{"CacheEntries", flowMonitorMap.CacheEntries}
    flowMonitorMap.EntityData.Leafs["cache-inactive-aging-timeout"] = types.YLeaf{"CacheInactiveAgingTimeout", flowMonitorMap.CacheInactiveAgingTimeout}
    flowMonitorMap.EntityData.Leafs["cache-active-aging-timeout"] = types.YLeaf{"CacheActiveAgingTimeout", flowMonitorMap.CacheActiveAgingTimeout}
    flowMonitorMap.EntityData.Leafs["cache-timeout-rate-limit"] = types.YLeaf{"CacheTimeoutRateLimit", flowMonitorMap.CacheTimeoutRateLimit}
    flowMonitorMap.EntityData.Leafs["cache-aging-mode"] = types.YLeaf{"CacheAgingMode", flowMonitorMap.CacheAgingMode}
    return &(flowMonitorMap.EntityData)
}

// NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option
// Specify an option for the flow cache
type NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether data should be filtered. The type is interface{}.
    Filtered interface{}

    // Specify whether to export physical ifh for bundle interface. The type is
    // interface{}.
    OutBundleMember interface{}

    // Specify whether it exports the physical output interface. The type is
    // interface{}.
    OutPhysInt interface{}

    // Specify if BGP Attributes AS_PATH STD_COMM should be exported. The type is
    // interface{}.
    BgpAttr interface{}
}

func (option *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Option) GetEntityData() *types.CommonEntityData {
    option.EntityData.YFilter = option.YFilter
    option.EntityData.YangName = "option"
    option.EntityData.BundleName = "cisco_ios_xr"
    option.EntityData.ParentYangName = "flow-monitor-map"
    option.EntityData.SegmentPath = "option"
    option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    option.EntityData.Children = make(map[string]types.YChild)
    option.EntityData.Leafs = make(map[string]types.YLeaf)
    option.EntityData.Leafs["filtered"] = types.YLeaf{"Filtered", option.Filtered}
    option.EntityData.Leafs["out-bundle-member"] = types.YLeaf{"OutBundleMember", option.OutBundleMember}
    option.EntityData.Leafs["out-phys-int"] = types.YLeaf{"OutPhysInt", option.OutPhysInt}
    option.EntityData.Leafs["bgp-attr"] = types.YLeaf{"BgpAttr", option.BgpAttr}
    return &(option.EntityData)
}

// NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters
// Configure exporters to be used by the
// monitor-map
type NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure exporter to be used by the monitor-map. The type is slice of
    // NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter.
    Exporter []NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter
}

func (exporters *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters) GetEntityData() *types.CommonEntityData {
    exporters.EntityData.YFilter = exporters.YFilter
    exporters.EntityData.YangName = "exporters"
    exporters.EntityData.BundleName = "cisco_ios_xr"
    exporters.EntityData.ParentYangName = "flow-monitor-map"
    exporters.EntityData.SegmentPath = "exporters"
    exporters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exporters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exporters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exporters.EntityData.Children = make(map[string]types.YChild)
    exporters.EntityData.Children["exporter"] = types.YChild{"Exporter", nil}
    for i := range exporters.Exporter {
        exporters.EntityData.Children[types.GetSegmentPath(&exporters.Exporter[i])] = types.YChild{"Exporter", &exporters.Exporter[i]}
    }
    exporters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(exporters.EntityData)
}

// NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter
// Configure exporter to be used by the
// monitor-map
type NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Exporter name. The type is string with length:
    // 1..32.
    ExporterName interface{}
}

func (exporter *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Exporters_Exporter) GetEntityData() *types.CommonEntityData {
    exporter.EntityData.YFilter = exporter.YFilter
    exporter.EntityData.YangName = "exporter"
    exporter.EntityData.BundleName = "cisco_ios_xr"
    exporter.EntityData.ParentYangName = "exporters"
    exporter.EntityData.SegmentPath = "exporter" + "[exporter-name='" + fmt.Sprintf("%v", exporter.ExporterName) + "']"
    exporter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exporter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exporter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exporter.EntityData.Children = make(map[string]types.YChild)
    exporter.EntityData.Leafs = make(map[string]types.YLeaf)
    exporter.EntityData.Leafs["exporter-name"] = types.YLeaf{"ExporterName", exporter.ExporterName}
    return &(exporter.EntityData)
}

// NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record
// Specify a flow record format
// This type is a presence type.
type NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow record format (Either 'ipv4-raw' ,'ipv4-peer-as', 'ipv6', 'mpls',
    // 'mpls-ipv4', 'mpls-ipv6', 'mpls-ipv4-ipv6', 'ipv6-peer-as'). The type is
    // string with length: 1..32. This attribute is mandatory.
    RecordName interface{}

    // Enter label value for MPLS record type. The type is interface{} with range:
    // 1..6.
    Label interface{}
}

func (record *NetFlow_FlowMonitorMapPerformanceTable_FlowMonitorMap_Record) GetEntityData() *types.CommonEntityData {
    record.EntityData.YFilter = record.YFilter
    record.EntityData.YangName = "record"
    record.EntityData.BundleName = "cisco_ios_xr"
    record.EntityData.ParentYangName = "flow-monitor-map"
    record.EntityData.SegmentPath = "record"
    record.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    record.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    record.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    record.EntityData.Children = make(map[string]types.YChild)
    record.EntityData.Leafs = make(map[string]types.YLeaf)
    record.EntityData.Leafs["record-name"] = types.YLeaf{"RecordName", record.RecordName}
    record.EntityData.Leafs["label"] = types.YLeaf{"Label", record.Label}
    return &(record.EntityData)
}

