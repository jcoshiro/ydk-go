// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-rip package operational data.
// 
// This module contains definitions
// for the following management objects:
//   rip: RIP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_rip_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_rip_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-rip-oper rip}", reflect.TypeOf(Rip{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-rip-oper:rip", reflect.TypeOf(Rip{}))
}

// InterfaceState represents Interface state
type InterfaceState string

const (
    // Interface does not exist
    InterfaceState_interface_none InterfaceState = "interface-none"

    // Interface exists but IP is down
    InterfaceState_interface_down InterfaceState = "interface-down"

    // Interface exists and IP is up
    InterfaceState_interface_up InterfaceState = "interface-up"

    // Unknown state
    InterfaceState_interface_unknown InterfaceState = "interface-unknown"
)

// RipRouteOrigin represents Rip route origin
type RipRouteOrigin string

const (
    // configured 'network'
    RipRouteOrigin_rip_rt_org_runover RipRouteOrigin = "rip-rt-org-runover"

    // redistributed
    RipRouteOrigin_rip_rt_org_redist RipRouteOrigin = "rip-rt-org-redist"

    // auto summary
    RipRouteOrigin_rip_rt_org_auto_summary RipRouteOrigin = "rip-rt-org-auto-summary"

    // learned via RIP
    RipRouteOrigin_rip_rt_org_rip RipRouteOrigin = "rip-rt-org-rip"

    // interface summary-address
    RipRouteOrigin_rip_rt_org_intsummary RipRouteOrigin = "rip-rt-org-intsummary"

    // route stay in for triggered rip
    RipRouteOrigin_rip_rt_org_unused RipRouteOrigin = "rip-rt-org-unused"
)

// Rip
// RIP operational data
type Rip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF related operational data.
    Vrfs Rip_Vrfs

    // Protocol operational data.
    Protocol Rip_Protocol

    // RIP operational data for Default VRF.
    DefaultVrf Rip_DefaultVrf
}

func (rip *Rip) GetEntityData() *types.CommonEntityData {
    rip.EntityData.YFilter = rip.YFilter
    rip.EntityData.YangName = "rip"
    rip.EntityData.BundleName = "cisco_ios_xr"
    rip.EntityData.ParentYangName = "Cisco-IOS-XR-ip-rip-oper"
    rip.EntityData.SegmentPath = "Cisco-IOS-XR-ip-rip-oper:rip"
    rip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rip.EntityData.Children = make(map[string]types.YChild)
    rip.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &rip.Vrfs}
    rip.EntityData.Children["protocol"] = types.YChild{"Protocol", &rip.Protocol}
    rip.EntityData.Children["default-vrf"] = types.YChild{"DefaultVrf", &rip.DefaultVrf}
    rip.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rip.EntityData)
}

// Rip_Vrfs
// VRF related operational data
type Rip_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a particular VRF. The type is slice of Rip_Vrfs_Vrf.
    Vrf []Rip_Vrfs_Vrf
}

func (vrfs *Rip_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "rip"
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

// Rip_Vrfs_Vrf
// Operational data for a particular VRF
type Rip_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // RIP route database.
    Routes Rip_Vrfs_Vrf_Routes

    // RIP global configuration.
    Configuration Rip_Vrfs_Vrf_Configuration

    // RIP statistics information.
    Statistics Rip_Vrfs_Vrf_Statistics

    // RIP interfaces.
    Interfaces Rip_Vrfs_Vrf_Interfaces

    // Global Information .
    Global Rip_Vrfs_Vrf_Global
}

func (vrf *Rip_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["routes"] = types.YChild{"Routes", &vrf.Routes}
    vrf.EntityData.Children["configuration"] = types.YChild{"Configuration", &vrf.Configuration}
    vrf.EntityData.Children["statistics"] = types.YChild{"Statistics", &vrf.Statistics}
    vrf.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &vrf.Interfaces}
    vrf.EntityData.Children["global"] = types.YChild{"Global", &vrf.Global}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Rip_Vrfs_Vrf_Routes
// RIP route database
type Rip_Vrfs_Vrf_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A route in the RIP database. The type is slice of
    // Rip_Vrfs_Vrf_Routes_Route.
    Route []Rip_Vrfs_Vrf_Routes_Route
}

func (routes *Rip_Vrfs_Vrf_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "vrf"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = make(map[string]types.YChild)
    routes.EntityData.Children["route"] = types.YChild{"Route", nil}
    for i := range routes.Route {
        routes.EntityData.Children[types.GetSegmentPath(&routes.Route[i])] = types.YChild{"Route", &routes.Route[i]}
    }
    routes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routes.EntityData)
}

// Rip_Vrfs_Vrf_Routes_Route
// A route in the RIP database
type Rip_Vrfs_Vrf_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // Destination IP Address for this route. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLengthXr interface{}

    // Route administrative distance. The type is interface{} with range:
    // 0..65535.
    Distance interface{}

    // Hop count for this route. The type is interface{} with range: 0..65535.
    BgpCount interface{}

    // Type of this route. The type is interface{} with range: 0..65535.
    RouteType interface{}

    // Summary route placeholder indicator. The type is bool.
    RouteSummary interface{}

    // Generic route information. The type is interface{} with range: 0..65535.
    RouteTag interface{}

    // RIB supplied version number. The type is interface{} with range: 0..255.
    Version interface{}

    // RIB supplied route attributes. The type is interface{} with range:
    // 0..4294967295.
    Attributes interface{}

    // Active route indicator. The type is bool.
    Active interface{}

    // Where this route was learnt. The type is RipRouteOrigin.
    PathOrigin interface{}

    // Indicates whether route is in holddown. The type is bool.
    HoldDown interface{}

    // The paths for this route. The type is slice of
    // Rip_Vrfs_Vrf_Routes_Route_Paths.
    Paths []Rip_Vrfs_Vrf_Routes_Route_Paths
}

func (route *Rip_Vrfs_Vrf_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = make(map[string]types.YChild)
    route.EntityData.Children["paths"] = types.YChild{"Paths", nil}
    for i := range route.Paths {
        route.EntityData.Children[types.GetSegmentPath(&route.Paths[i])] = types.YChild{"Paths", &route.Paths[i]}
    }
    route.EntityData.Leafs = make(map[string]types.YLeaf)
    route.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", route.Prefix}
    route.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", route.PrefixLength}
    route.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", route.DestinationAddress}
    route.EntityData.Leafs["prefix-length-xr"] = types.YLeaf{"PrefixLengthXr", route.PrefixLengthXr}
    route.EntityData.Leafs["distance"] = types.YLeaf{"Distance", route.Distance}
    route.EntityData.Leafs["bgp-count"] = types.YLeaf{"BgpCount", route.BgpCount}
    route.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", route.RouteType}
    route.EntityData.Leafs["route-summary"] = types.YLeaf{"RouteSummary", route.RouteSummary}
    route.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", route.RouteTag}
    route.EntityData.Leafs["version"] = types.YLeaf{"Version", route.Version}
    route.EntityData.Leafs["attributes"] = types.YLeaf{"Attributes", route.Attributes}
    route.EntityData.Leafs["active"] = types.YLeaf{"Active", route.Active}
    route.EntityData.Leafs["path-origin"] = types.YLeaf{"PathOrigin", route.PathOrigin}
    route.EntityData.Leafs["hold-down"] = types.YLeaf{"HoldDown", route.HoldDown}
    return &(route.EntityData)
}

// Rip_Vrfs_Vrf_Routes_Route_Paths
// The paths for this route
type Rip_Vrfs_Vrf_Routes_Route_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Next hop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}

    // Tag. The type is interface{} with range: 0..65535.
    Tag interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Up time. The type is interface{} with range: 0..4294967295.
    Uptime interface{}

    // Permanent indicator. The type is bool.
    IsPermanent interface{}
}

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "route"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    paths.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", paths.SourceAddress}
    paths.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", paths.NextHopAddress}
    paths.EntityData.Leafs["metric"] = types.YLeaf{"Metric", paths.Metric}
    paths.EntityData.Leafs["tag"] = types.YLeaf{"Tag", paths.Tag}
    paths.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", paths.Interface_}
    paths.EntityData.Leafs["uptime"] = types.YLeaf{"Uptime", paths.Uptime}
    paths.EntityData.Leafs["is-permanent"] = types.YLeaf{"IsPermanent", paths.IsPermanent}
    return &(paths.EntityData)
}

// Rip_Vrfs_Vrf_Configuration
// RIP global configuration
type Rip_Vrfs_Vrf_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF active indicator. The type is bool.
    Active interface{}

    // VRF added to socket indicator. The type is bool.
    VrFisedSocket interface{}

    // Version of RIP configured. The type is interface{} with range:
    // -2147483648..2147483647.
    RipVersion interface{}

    // Default metric for RIP routes. The type is interface{} with range: 0..255.
    DefaultMetric interface{}

    // Maximum number of paths a route can have. The type is interface{} with
    // range: 0..255.
    MaximumPaths interface{}

    // Auto-summarization indicator. The type is bool.
    AutoSummarize interface{}

    // Use broadcast/multicast address indicator. The type is bool.
    MulticastAddress interface{}

    // Flash update threshold. The type is interface{} with range: 0..255.
    FlashThreshold interface{}

    // Length of the input queue. The type is interface{} with range: 0..65535.
    InputQLength interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Incoming packet source validation indicator. The type is bool.
    ValidationIndicator interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // NSF Enable status. The type is bool.
    NsfStatus interface{}

    // NSF life time. The type is interface{} with range: 0..4294967295.
    NsfLifeTime interface{}
}

func (configuration *Rip_Vrfs_Vrf_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "vrf"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = make(map[string]types.YChild)
    configuration.EntityData.Leafs = make(map[string]types.YLeaf)
    configuration.EntityData.Leafs["active"] = types.YLeaf{"Active", configuration.Active}
    configuration.EntityData.Leafs["vr-fised-socket"] = types.YLeaf{"VrFisedSocket", configuration.VrFisedSocket}
    configuration.EntityData.Leafs["rip-version"] = types.YLeaf{"RipVersion", configuration.RipVersion}
    configuration.EntityData.Leafs["default-metric"] = types.YLeaf{"DefaultMetric", configuration.DefaultMetric}
    configuration.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", configuration.MaximumPaths}
    configuration.EntityData.Leafs["auto-summarize"] = types.YLeaf{"AutoSummarize", configuration.AutoSummarize}
    configuration.EntityData.Leafs["multicast-address"] = types.YLeaf{"MulticastAddress", configuration.MulticastAddress}
    configuration.EntityData.Leafs["flash-threshold"] = types.YLeaf{"FlashThreshold", configuration.FlashThreshold}
    configuration.EntityData.Leafs["input-q-length"] = types.YLeaf{"InputQLength", configuration.InputQLength}
    configuration.EntityData.Leafs["triggered-rip"] = types.YLeaf{"TriggeredRip", configuration.TriggeredRip}
    configuration.EntityData.Leafs["validation-indicator"] = types.YLeaf{"ValidationIndicator", configuration.ValidationIndicator}
    configuration.EntityData.Leafs["update-timer"] = types.YLeaf{"UpdateTimer", configuration.UpdateTimer}
    configuration.EntityData.Leafs["next-update-time"] = types.YLeaf{"NextUpdateTime", configuration.NextUpdateTime}
    configuration.EntityData.Leafs["invalid-timer"] = types.YLeaf{"InvalidTimer", configuration.InvalidTimer}
    configuration.EntityData.Leafs["hold-down-timer"] = types.YLeaf{"HoldDownTimer", configuration.HoldDownTimer}
    configuration.EntityData.Leafs["flush-timer"] = types.YLeaf{"FlushTimer", configuration.FlushTimer}
    configuration.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", configuration.OomFlags}
    configuration.EntityData.Leafs["nsf-status"] = types.YLeaf{"NsfStatus", configuration.NsfStatus}
    configuration.EntityData.Leafs["nsf-life-time"] = types.YLeaf{"NsfLifeTime", configuration.NsfLifeTime}
    return &(configuration.EntityData)
}

// Rip_Vrfs_Vrf_Statistics
// RIP statistics information
type Rip_Vrfs_Vrf_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total packets received. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Total discarded packets. The type is interface{} with range: 0..4294967295.
    DiscardedPackets interface{}

    // Total discarded routes. The type is interface{} with range: 0..4294967295.
    DiscardedRoutes interface{}

    // Packets rx in SRP. The type is interface{} with range: 0..4294967295.
    StandbyPacketsReceived interface{}

    // Number of messages sent. The type is interface{} with range: 0..4294967295.
    SentMessages interface{}

    // Number of message send failures. The type is interface{} with range:
    // 0..4294967295.
    SentMessageFailures interface{}

    // Number of RIP queries responded to. The type is interface{} with range:
    // 0..4294967295.
    QueryResponses interface{}

    // Number of periodic RIP updates. The type is interface{} with range:
    // 0..4294967295.
    PeriodicUpdates interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Number of failures to allocate memory for a route. The type is interface{}
    // with range: 0..4294967295.
    RouteMallocFailures interface{}

    // Number of failures to allocate memory for a path. The type is interface{}
    // with range: 0..4294967295.
    PathMallocFailures interface{}

    // Number of route updates to RIB made by RIP. The type is interface{} with
    // range: 0..4294967295.
    RibUpdates interface{}
}

func (statistics *Rip_Vrfs_Vrf_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "vrf"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets}
    statistics.EntityData.Leafs["discarded-packets"] = types.YLeaf{"DiscardedPackets", statistics.DiscardedPackets}
    statistics.EntityData.Leafs["discarded-routes"] = types.YLeaf{"DiscardedRoutes", statistics.DiscardedRoutes}
    statistics.EntityData.Leafs["standby-packets-received"] = types.YLeaf{"StandbyPacketsReceived", statistics.StandbyPacketsReceived}
    statistics.EntityData.Leafs["sent-messages"] = types.YLeaf{"SentMessages", statistics.SentMessages}
    statistics.EntityData.Leafs["sent-message-failures"] = types.YLeaf{"SentMessageFailures", statistics.SentMessageFailures}
    statistics.EntityData.Leafs["query-responses"] = types.YLeaf{"QueryResponses", statistics.QueryResponses}
    statistics.EntityData.Leafs["periodic-updates"] = types.YLeaf{"PeriodicUpdates", statistics.PeriodicUpdates}
    statistics.EntityData.Leafs["route-count"] = types.YLeaf{"RouteCount", statistics.RouteCount}
    statistics.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", statistics.PathCount}
    statistics.EntityData.Leafs["route-malloc-failures"] = types.YLeaf{"RouteMallocFailures", statistics.RouteMallocFailures}
    statistics.EntityData.Leafs["path-malloc-failures"] = types.YLeaf{"PathMallocFailures", statistics.PathMallocFailures}
    statistics.EntityData.Leafs["rib-updates"] = types.YLeaf{"RibUpdates", statistics.RibUpdates}
    return &(statistics.EntityData)
}

// Rip_Vrfs_Vrf_Interfaces
// RIP interfaces
type Rip_Vrfs_Vrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular RIP interface. The type is slice of
    // Rip_Vrfs_Vrf_Interfaces_Interface_.
    Interface_ []Rip_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "vrf"
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

// Rip_Vrfs_Vrf_Interfaces_Interface
// Information about a particular RIP interface
type Rip_Vrfs_Vrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    Interface_ interface{}

    // Interface handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    IfHandle interface{}

    // Whether RIP is enabled on this interface. The type is bool.
    RipEnabled interface{}

    // Passive interface indicator. The type is bool.
    IsPassiveInterface interface{}

    // Use broadcast address for v2 packets. The type is bool.
    MulticastAddress interface{}

    // Accept routes of metric 0 indicator. The type is bool.
    AcceptMetric interface{}

    // Versions that the interface is sending. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // Versions that the interface will recieve. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Current state of the interface. The type is InterfaceState.
    State interface{}

    // IP Address of this interface. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Prefix length of the IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Cost added to routes through this interface. The type is interface{} with
    // range: 0..4294967295.
    MetricCost interface{}

    // Split horizon enabled indicator. The type is bool.
    SplitHorizon interface{}

    // Poisoned reverse enabled indicator. The type is bool.
    PoisonHorizon interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Interface's triggered RIP neighbor. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborAddress interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // Multicast group join status. The type is bool.
    JoinStatus interface{}

    // LPTSState. The type is bool.
    LptsState interface{}

    // Authentication Mode. The type is interface{} with range: 0..4294967295.
    AuthMode interface{}

    // Authentication Keychain Name. The type is string.
    AuthKeychain interface{}

    // Authentication send key exists. The type is bool.
    SendAuthKeyExists interface{}

    // Authentication key programmed with MD5 algorithm. The type is bool.
    AuthKeyMd5 interface{}

    // Current active Send Authentication Key Id. The type is interface{} with
    // range: 0..18446744073709551615.
    AuthKeySendId interface{}

    // Total packets received. The type is interface{} with range: 0..4294967295.
    TotalPktRecvd interface{}

    // Packets dropped due to wrong keychain configured. The type is interface{}
    // with range: 0..4294967295.
    PktDropWrongKc interface{}

    // Packets dropped due to missing authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropNoAuth interface{}

    // Packets dropped due to invalid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropInvalidAuth interface{}

    // Packets accepted with valid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktAcceptedValidAuth interface{}

    // User defined summary addresses. The type is slice of
    // Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary.
    RipSummary []Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary

    // Neighbors on this interface. The type is slice of
    // Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer.
    RipPeer []Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["rip-summary"] = types.YChild{"RipSummary", nil}
    for i := range self.RipSummary {
        self.EntityData.Children[types.GetSegmentPath(&self.RipSummary[i])] = types.YChild{"RipSummary", &self.RipSummary[i]}
    }
    self.EntityData.Children["rip-peer"] = types.YChild{"RipPeer", nil}
    for i := range self.RipPeer {
        self.EntityData.Children[types.GetSegmentPath(&self.RipPeer[i])] = types.YChild{"RipPeer", &self.RipPeer[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", self.IfHandle}
    self.EntityData.Leafs["rip-enabled"] = types.YLeaf{"RipEnabled", self.RipEnabled}
    self.EntityData.Leafs["is-passive-interface"] = types.YLeaf{"IsPassiveInterface", self.IsPassiveInterface}
    self.EntityData.Leafs["multicast-address"] = types.YLeaf{"MulticastAddress", self.MulticastAddress}
    self.EntityData.Leafs["accept-metric"] = types.YLeaf{"AcceptMetric", self.AcceptMetric}
    self.EntityData.Leafs["send-version"] = types.YLeaf{"SendVersion", self.SendVersion}
    self.EntityData.Leafs["receive-version"] = types.YLeaf{"ReceiveVersion", self.ReceiveVersion}
    self.EntityData.Leafs["state"] = types.YLeaf{"State", self.State}
    self.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", self.DestinationAddress}
    self.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", self.PrefixLength}
    self.EntityData.Leafs["metric-cost"] = types.YLeaf{"MetricCost", self.MetricCost}
    self.EntityData.Leafs["split-horizon"] = types.YLeaf{"SplitHorizon", self.SplitHorizon}
    self.EntityData.Leafs["poison-horizon"] = types.YLeaf{"PoisonHorizon", self.PoisonHorizon}
    self.EntityData.Leafs["triggered-rip"] = types.YLeaf{"TriggeredRip", self.TriggeredRip}
    self.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", self.NeighborAddress}
    self.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", self.OomFlags}
    self.EntityData.Leafs["join-status"] = types.YLeaf{"JoinStatus", self.JoinStatus}
    self.EntityData.Leafs["lpts-state"] = types.YLeaf{"LptsState", self.LptsState}
    self.EntityData.Leafs["auth-mode"] = types.YLeaf{"AuthMode", self.AuthMode}
    self.EntityData.Leafs["auth-keychain"] = types.YLeaf{"AuthKeychain", self.AuthKeychain}
    self.EntityData.Leafs["send-auth-key-exists"] = types.YLeaf{"SendAuthKeyExists", self.SendAuthKeyExists}
    self.EntityData.Leafs["auth-key-md5"] = types.YLeaf{"AuthKeyMd5", self.AuthKeyMd5}
    self.EntityData.Leafs["auth-key-send-id"] = types.YLeaf{"AuthKeySendId", self.AuthKeySendId}
    self.EntityData.Leafs["total-pkt-recvd"] = types.YLeaf{"TotalPktRecvd", self.TotalPktRecvd}
    self.EntityData.Leafs["pkt-drop-wrong-kc"] = types.YLeaf{"PktDropWrongKc", self.PktDropWrongKc}
    self.EntityData.Leafs["pkt-drop-no-auth"] = types.YLeaf{"PktDropNoAuth", self.PktDropNoAuth}
    self.EntityData.Leafs["pkt-drop-invalid-auth"] = types.YLeaf{"PktDropInvalidAuth", self.PktDropInvalidAuth}
    self.EntityData.Leafs["pkt-accepted-valid-auth"] = types.YLeaf{"PktAcceptedValidAuth", self.PktAcceptedValidAuth}
    return &(self.EntityData)
}

// Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary
// User defined summary addresses
type Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary address prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Summary address prefix length. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Summary address next hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Summary metric. The type is interface{} with range:
    // -2147483648..2147483647.
    Metric interface{}
}

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetEntityData() *types.CommonEntityData {
    ripSummary.EntityData.YFilter = ripSummary.YFilter
    ripSummary.EntityData.YangName = "rip-summary"
    ripSummary.EntityData.BundleName = "cisco_ios_xr"
    ripSummary.EntityData.ParentYangName = "interface"
    ripSummary.EntityData.SegmentPath = "rip-summary"
    ripSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ripSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ripSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ripSummary.EntityData.Children = make(map[string]types.YChild)
    ripSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    ripSummary.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ripSummary.Prefix}
    ripSummary.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", ripSummary.PrefixLength}
    ripSummary.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", ripSummary.NextHopAddress}
    ripSummary.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ripSummary.Metric}
    return &(ripSummary.EntityData)
}

// Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer
// Neighbors on this interface
type Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Uptime of this peer. The type is interface{} with range: 0..4294967295.
    PeerUptime interface{}

    // IP Address of this peer. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // RIP version for this peer. The type is interface{} with range: 0..255.
    PeerVersion interface{}

    // Discarded packets from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerPackets interface{}

    // Discarded routes from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerRoutes interface{}
}

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetEntityData() *types.CommonEntityData {
    ripPeer.EntityData.YFilter = ripPeer.YFilter
    ripPeer.EntityData.YangName = "rip-peer"
    ripPeer.EntityData.BundleName = "cisco_ios_xr"
    ripPeer.EntityData.ParentYangName = "interface"
    ripPeer.EntityData.SegmentPath = "rip-peer"
    ripPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ripPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ripPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ripPeer.EntityData.Children = make(map[string]types.YChild)
    ripPeer.EntityData.Leafs = make(map[string]types.YLeaf)
    ripPeer.EntityData.Leafs["peer-uptime"] = types.YLeaf{"PeerUptime", ripPeer.PeerUptime}
    ripPeer.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", ripPeer.PeerAddress}
    ripPeer.EntityData.Leafs["peer-version"] = types.YLeaf{"PeerVersion", ripPeer.PeerVersion}
    ripPeer.EntityData.Leafs["discarded-peer-packets"] = types.YLeaf{"DiscardedPeerPackets", ripPeer.DiscardedPeerPackets}
    ripPeer.EntityData.Leafs["discarded-peer-routes"] = types.YLeaf{"DiscardedPeerRoutes", ripPeer.DiscardedPeerRoutes}
    return &(ripPeer.EntityData)
}

// Rip_Vrfs_Vrf_Global
// Global Information 
type Rip_Vrfs_Vrf_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF summary data.
    VrfSummary Rip_Vrfs_Vrf_Global_VrfSummary

    // List of Interfaces configured. The type is slice of
    // Rip_Vrfs_Vrf_Global_InterfaceSummary.
    InterfaceSummary []Rip_Vrfs_Vrf_Global_InterfaceSummary
}

func (global *Rip_Vrfs_Vrf_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "vrf"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["vrf-summary"] = types.YChild{"VrfSummary", &global.VrfSummary}
    global.EntityData.Children["interface-summary"] = types.YChild{"InterfaceSummary", nil}
    for i := range global.InterfaceSummary {
        global.EntityData.Children[types.GetSegmentPath(&global.InterfaceSummary[i])] = types.YChild{"InterfaceSummary", &global.InterfaceSummary[i]}
    }
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(global.EntityData)
}

// Rip_Vrfs_Vrf_Global_VrfSummary
// VRF summary data
type Rip_Vrfs_Vrf_Global_VrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Active indicator. The type is bool.
    Active interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Number of interfaces configured. The type is interface{} with range:
    // 0..4294967295.
    InterfaceConfiguredCount interface{}

    // Number of active interfaces. The type is interface{} with range:
    // 0..4294967295.
    ActiveInterfaceCount interface{}
}

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetEntityData() *types.CommonEntityData {
    vrfSummary.EntityData.YFilter = vrfSummary.YFilter
    vrfSummary.EntityData.YangName = "vrf-summary"
    vrfSummary.EntityData.BundleName = "cisco_ios_xr"
    vrfSummary.EntityData.ParentYangName = "global"
    vrfSummary.EntityData.SegmentPath = "vrf-summary"
    vrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummary.EntityData.Children = make(map[string]types.YChild)
    vrfSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfSummary.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrfSummary.VrfName}
    vrfSummary.EntityData.Leafs["active"] = types.YLeaf{"Active", vrfSummary.Active}
    vrfSummary.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", vrfSummary.OomFlags}
    vrfSummary.EntityData.Leafs["route-count"] = types.YLeaf{"RouteCount", vrfSummary.RouteCount}
    vrfSummary.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", vrfSummary.PathCount}
    vrfSummary.EntityData.Leafs["update-timer"] = types.YLeaf{"UpdateTimer", vrfSummary.UpdateTimer}
    vrfSummary.EntityData.Leafs["next-update-time"] = types.YLeaf{"NextUpdateTime", vrfSummary.NextUpdateTime}
    vrfSummary.EntityData.Leafs["invalid-timer"] = types.YLeaf{"InvalidTimer", vrfSummary.InvalidTimer}
    vrfSummary.EntityData.Leafs["hold-down-timer"] = types.YLeaf{"HoldDownTimer", vrfSummary.HoldDownTimer}
    vrfSummary.EntityData.Leafs["flush-timer"] = types.YLeaf{"FlushTimer", vrfSummary.FlushTimer}
    vrfSummary.EntityData.Leafs["interface-configured-count"] = types.YLeaf{"InterfaceConfiguredCount", vrfSummary.InterfaceConfiguredCount}
    vrfSummary.EntityData.Leafs["active-interface-count"] = types.YLeaf{"ActiveInterfaceCount", vrfSummary.ActiveInterfaceCount}
    return &(vrfSummary.EntityData)
}

// Rip_Vrfs_Vrf_Global_InterfaceSummary
// List of Interfaces configured
type Rip_Vrfs_Vrf_Global_InterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // RIP enabled indicator. The type is bool.
    Enabled interface{}

    // Interface state. The type is InterfaceState.
    State interface{}

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // RIP versions this interface sends out. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // RIP versions this interface will receive. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Number of neighbors seen. The type is interface{} with range:
    // 0..4294967295.
    NeighborCount interface{}
}

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetEntityData() *types.CommonEntityData {
    interfaceSummary.EntityData.YFilter = interfaceSummary.YFilter
    interfaceSummary.EntityData.YangName = "interface-summary"
    interfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummary.EntityData.ParentYangName = "global"
    interfaceSummary.EntityData.SegmentPath = "interface-summary"
    interfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummary.EntityData.Children = make(map[string]types.YChild)
    interfaceSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceSummary.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceSummary.InterfaceName}
    interfaceSummary.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", interfaceSummary.Enabled}
    interfaceSummary.EntityData.Leafs["state"] = types.YLeaf{"State", interfaceSummary.State}
    interfaceSummary.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", interfaceSummary.DestinationAddress}
    interfaceSummary.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", interfaceSummary.PrefixLength}
    interfaceSummary.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", interfaceSummary.OomFlags}
    interfaceSummary.EntityData.Leafs["send-version"] = types.YLeaf{"SendVersion", interfaceSummary.SendVersion}
    interfaceSummary.EntityData.Leafs["receive-version"] = types.YLeaf{"ReceiveVersion", interfaceSummary.ReceiveVersion}
    interfaceSummary.EntityData.Leafs["neighbor-count"] = types.YLeaf{"NeighborCount", interfaceSummary.NeighborCount}
    return &(interfaceSummary.EntityData)
}

// Rip_Protocol
// Protocol operational data
type Rip_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RIP global process .
    Process Rip_Protocol_Process

    // RIP operational data for Default VRF.
    DefaultVrf Rip_Protocol_DefaultVrf
}

func (protocol *Rip_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "rip"
    protocol.EntityData.SegmentPath = "protocol"
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = make(map[string]types.YChild)
    protocol.EntityData.Children["process"] = types.YChild{"Process", &protocol.Process}
    protocol.EntityData.Children["default-vrf"] = types.YChild{"DefaultVrf", &protocol.DefaultVrf}
    protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protocol.EntityData)
}

// Rip_Protocol_Process
// RIP global process 
type Rip_Protocol_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of VRFs configured. The type is interface{} with range:
    // 0..4294967295.
    VrfConfigCount interface{}

    // Number of active VRFs. The type is interface{} with range: 0..4294967295.
    VrfActiveCount interface{}

    // Socket descriptior. The type is interface{} with range:
    // -2147483648..2147483647.
    SocketDescriptor interface{}

    // Current OOM state. The type is interface{} with range:
    // -2147483648..2147483647.
    CurrentOomState interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // List of VRFs configured. The type is slice of
    // Rip_Protocol_Process_VrfSummary.
    VrfSummary []Rip_Protocol_Process_VrfSummary
}

func (process *Rip_Protocol_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "protocol"
    process.EntityData.SegmentPath = "process"
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = make(map[string]types.YChild)
    process.EntityData.Children["vrf-summary"] = types.YChild{"VrfSummary", nil}
    for i := range process.VrfSummary {
        process.EntityData.Children[types.GetSegmentPath(&process.VrfSummary[i])] = types.YChild{"VrfSummary", &process.VrfSummary[i]}
    }
    process.EntityData.Leafs = make(map[string]types.YLeaf)
    process.EntityData.Leafs["vrf-config-count"] = types.YLeaf{"VrfConfigCount", process.VrfConfigCount}
    process.EntityData.Leafs["vrf-active-count"] = types.YLeaf{"VrfActiveCount", process.VrfActiveCount}
    process.EntityData.Leafs["socket-descriptor"] = types.YLeaf{"SocketDescriptor", process.SocketDescriptor}
    process.EntityData.Leafs["current-oom-state"] = types.YLeaf{"CurrentOomState", process.CurrentOomState}
    process.EntityData.Leafs["route-count"] = types.YLeaf{"RouteCount", process.RouteCount}
    process.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", process.PathCount}
    return &(process.EntityData)
}

// Rip_Protocol_Process_VrfSummary
// List of VRFs configured
type Rip_Protocol_Process_VrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Active indicator. The type is bool.
    Active interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Number of interfaces configured. The type is interface{} with range:
    // 0..4294967295.
    InterfaceConfiguredCount interface{}

    // Number of active interfaces. The type is interface{} with range:
    // 0..4294967295.
    ActiveInterfaceCount interface{}
}

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetEntityData() *types.CommonEntityData {
    vrfSummary.EntityData.YFilter = vrfSummary.YFilter
    vrfSummary.EntityData.YangName = "vrf-summary"
    vrfSummary.EntityData.BundleName = "cisco_ios_xr"
    vrfSummary.EntityData.ParentYangName = "process"
    vrfSummary.EntityData.SegmentPath = "vrf-summary"
    vrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummary.EntityData.Children = make(map[string]types.YChild)
    vrfSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfSummary.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrfSummary.VrfName}
    vrfSummary.EntityData.Leafs["active"] = types.YLeaf{"Active", vrfSummary.Active}
    vrfSummary.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", vrfSummary.OomFlags}
    vrfSummary.EntityData.Leafs["route-count"] = types.YLeaf{"RouteCount", vrfSummary.RouteCount}
    vrfSummary.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", vrfSummary.PathCount}
    vrfSummary.EntityData.Leafs["update-timer"] = types.YLeaf{"UpdateTimer", vrfSummary.UpdateTimer}
    vrfSummary.EntityData.Leafs["next-update-time"] = types.YLeaf{"NextUpdateTime", vrfSummary.NextUpdateTime}
    vrfSummary.EntityData.Leafs["invalid-timer"] = types.YLeaf{"InvalidTimer", vrfSummary.InvalidTimer}
    vrfSummary.EntityData.Leafs["hold-down-timer"] = types.YLeaf{"HoldDownTimer", vrfSummary.HoldDownTimer}
    vrfSummary.EntityData.Leafs["flush-timer"] = types.YLeaf{"FlushTimer", vrfSummary.FlushTimer}
    vrfSummary.EntityData.Leafs["interface-configured-count"] = types.YLeaf{"InterfaceConfiguredCount", vrfSummary.InterfaceConfiguredCount}
    vrfSummary.EntityData.Leafs["active-interface-count"] = types.YLeaf{"ActiveInterfaceCount", vrfSummary.ActiveInterfaceCount}
    return &(vrfSummary.EntityData)
}

// Rip_Protocol_DefaultVrf
// RIP operational data for Default VRF
type Rip_Protocol_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RIP route database.
    Routes Rip_Protocol_DefaultVrf_Routes

    // RIP global configuration.
    Configuration Rip_Protocol_DefaultVrf_Configuration

    // RIP statistics information.
    Statistics Rip_Protocol_DefaultVrf_Statistics

    // RIP interfaces.
    Interfaces Rip_Protocol_DefaultVrf_Interfaces

    // Global Information .
    Global Rip_Protocol_DefaultVrf_Global
}

func (defaultVrf *Rip_Protocol_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "protocol"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = make(map[string]types.YChild)
    defaultVrf.EntityData.Children["routes"] = types.YChild{"Routes", &defaultVrf.Routes}
    defaultVrf.EntityData.Children["configuration"] = types.YChild{"Configuration", &defaultVrf.Configuration}
    defaultVrf.EntityData.Children["statistics"] = types.YChild{"Statistics", &defaultVrf.Statistics}
    defaultVrf.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &defaultVrf.Interfaces}
    defaultVrf.EntityData.Children["global"] = types.YChild{"Global", &defaultVrf.Global}
    defaultVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(defaultVrf.EntityData)
}

// Rip_Protocol_DefaultVrf_Routes
// RIP route database
type Rip_Protocol_DefaultVrf_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A route in the RIP database. The type is slice of
    // Rip_Protocol_DefaultVrf_Routes_Route.
    Route []Rip_Protocol_DefaultVrf_Routes_Route
}

func (routes *Rip_Protocol_DefaultVrf_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "default-vrf"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = make(map[string]types.YChild)
    routes.EntityData.Children["route"] = types.YChild{"Route", nil}
    for i := range routes.Route {
        routes.EntityData.Children[types.GetSegmentPath(&routes.Route[i])] = types.YChild{"Route", &routes.Route[i]}
    }
    routes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routes.EntityData)
}

// Rip_Protocol_DefaultVrf_Routes_Route
// A route in the RIP database
type Rip_Protocol_DefaultVrf_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // Destination IP Address for this route. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLengthXr interface{}

    // Route administrative distance. The type is interface{} with range:
    // 0..65535.
    Distance interface{}

    // Hop count for this route. The type is interface{} with range: 0..65535.
    BgpCount interface{}

    // Type of this route. The type is interface{} with range: 0..65535.
    RouteType interface{}

    // Summary route placeholder indicator. The type is bool.
    RouteSummary interface{}

    // Generic route information. The type is interface{} with range: 0..65535.
    RouteTag interface{}

    // RIB supplied version number. The type is interface{} with range: 0..255.
    Version interface{}

    // RIB supplied route attributes. The type is interface{} with range:
    // 0..4294967295.
    Attributes interface{}

    // Active route indicator. The type is bool.
    Active interface{}

    // Where this route was learnt. The type is RipRouteOrigin.
    PathOrigin interface{}

    // Indicates whether route is in holddown. The type is bool.
    HoldDown interface{}

    // The paths for this route. The type is slice of
    // Rip_Protocol_DefaultVrf_Routes_Route_Paths.
    Paths []Rip_Protocol_DefaultVrf_Routes_Route_Paths
}

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = make(map[string]types.YChild)
    route.EntityData.Children["paths"] = types.YChild{"Paths", nil}
    for i := range route.Paths {
        route.EntityData.Children[types.GetSegmentPath(&route.Paths[i])] = types.YChild{"Paths", &route.Paths[i]}
    }
    route.EntityData.Leafs = make(map[string]types.YLeaf)
    route.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", route.Prefix}
    route.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", route.PrefixLength}
    route.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", route.DestinationAddress}
    route.EntityData.Leafs["prefix-length-xr"] = types.YLeaf{"PrefixLengthXr", route.PrefixLengthXr}
    route.EntityData.Leafs["distance"] = types.YLeaf{"Distance", route.Distance}
    route.EntityData.Leafs["bgp-count"] = types.YLeaf{"BgpCount", route.BgpCount}
    route.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", route.RouteType}
    route.EntityData.Leafs["route-summary"] = types.YLeaf{"RouteSummary", route.RouteSummary}
    route.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", route.RouteTag}
    route.EntityData.Leafs["version"] = types.YLeaf{"Version", route.Version}
    route.EntityData.Leafs["attributes"] = types.YLeaf{"Attributes", route.Attributes}
    route.EntityData.Leafs["active"] = types.YLeaf{"Active", route.Active}
    route.EntityData.Leafs["path-origin"] = types.YLeaf{"PathOrigin", route.PathOrigin}
    route.EntityData.Leafs["hold-down"] = types.YLeaf{"HoldDown", route.HoldDown}
    return &(route.EntityData)
}

// Rip_Protocol_DefaultVrf_Routes_Route_Paths
// The paths for this route
type Rip_Protocol_DefaultVrf_Routes_Route_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Next hop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}

    // Tag. The type is interface{} with range: 0..65535.
    Tag interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Up time. The type is interface{} with range: 0..4294967295.
    Uptime interface{}

    // Permanent indicator. The type is bool.
    IsPermanent interface{}
}

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "route"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    paths.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", paths.SourceAddress}
    paths.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", paths.NextHopAddress}
    paths.EntityData.Leafs["metric"] = types.YLeaf{"Metric", paths.Metric}
    paths.EntityData.Leafs["tag"] = types.YLeaf{"Tag", paths.Tag}
    paths.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", paths.Interface_}
    paths.EntityData.Leafs["uptime"] = types.YLeaf{"Uptime", paths.Uptime}
    paths.EntityData.Leafs["is-permanent"] = types.YLeaf{"IsPermanent", paths.IsPermanent}
    return &(paths.EntityData)
}

// Rip_Protocol_DefaultVrf_Configuration
// RIP global configuration
type Rip_Protocol_DefaultVrf_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF active indicator. The type is bool.
    Active interface{}

    // VRF added to socket indicator. The type is bool.
    VrFisedSocket interface{}

    // Version of RIP configured. The type is interface{} with range:
    // -2147483648..2147483647.
    RipVersion interface{}

    // Default metric for RIP routes. The type is interface{} with range: 0..255.
    DefaultMetric interface{}

    // Maximum number of paths a route can have. The type is interface{} with
    // range: 0..255.
    MaximumPaths interface{}

    // Auto-summarization indicator. The type is bool.
    AutoSummarize interface{}

    // Use broadcast/multicast address indicator. The type is bool.
    MulticastAddress interface{}

    // Flash update threshold. The type is interface{} with range: 0..255.
    FlashThreshold interface{}

    // Length of the input queue. The type is interface{} with range: 0..65535.
    InputQLength interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Incoming packet source validation indicator. The type is bool.
    ValidationIndicator interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // NSF Enable status. The type is bool.
    NsfStatus interface{}

    // NSF life time. The type is interface{} with range: 0..4294967295.
    NsfLifeTime interface{}
}

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "default-vrf"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = make(map[string]types.YChild)
    configuration.EntityData.Leafs = make(map[string]types.YLeaf)
    configuration.EntityData.Leafs["active"] = types.YLeaf{"Active", configuration.Active}
    configuration.EntityData.Leafs["vr-fised-socket"] = types.YLeaf{"VrFisedSocket", configuration.VrFisedSocket}
    configuration.EntityData.Leafs["rip-version"] = types.YLeaf{"RipVersion", configuration.RipVersion}
    configuration.EntityData.Leafs["default-metric"] = types.YLeaf{"DefaultMetric", configuration.DefaultMetric}
    configuration.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", configuration.MaximumPaths}
    configuration.EntityData.Leafs["auto-summarize"] = types.YLeaf{"AutoSummarize", configuration.AutoSummarize}
    configuration.EntityData.Leafs["multicast-address"] = types.YLeaf{"MulticastAddress", configuration.MulticastAddress}
    configuration.EntityData.Leafs["flash-threshold"] = types.YLeaf{"FlashThreshold", configuration.FlashThreshold}
    configuration.EntityData.Leafs["input-q-length"] = types.YLeaf{"InputQLength", configuration.InputQLength}
    configuration.EntityData.Leafs["triggered-rip"] = types.YLeaf{"TriggeredRip", configuration.TriggeredRip}
    configuration.EntityData.Leafs["validation-indicator"] = types.YLeaf{"ValidationIndicator", configuration.ValidationIndicator}
    configuration.EntityData.Leafs["update-timer"] = types.YLeaf{"UpdateTimer", configuration.UpdateTimer}
    configuration.EntityData.Leafs["next-update-time"] = types.YLeaf{"NextUpdateTime", configuration.NextUpdateTime}
    configuration.EntityData.Leafs["invalid-timer"] = types.YLeaf{"InvalidTimer", configuration.InvalidTimer}
    configuration.EntityData.Leafs["hold-down-timer"] = types.YLeaf{"HoldDownTimer", configuration.HoldDownTimer}
    configuration.EntityData.Leafs["flush-timer"] = types.YLeaf{"FlushTimer", configuration.FlushTimer}
    configuration.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", configuration.OomFlags}
    configuration.EntityData.Leafs["nsf-status"] = types.YLeaf{"NsfStatus", configuration.NsfStatus}
    configuration.EntityData.Leafs["nsf-life-time"] = types.YLeaf{"NsfLifeTime", configuration.NsfLifeTime}
    return &(configuration.EntityData)
}

// Rip_Protocol_DefaultVrf_Statistics
// RIP statistics information
type Rip_Protocol_DefaultVrf_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total packets received. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Total discarded packets. The type is interface{} with range: 0..4294967295.
    DiscardedPackets interface{}

    // Total discarded routes. The type is interface{} with range: 0..4294967295.
    DiscardedRoutes interface{}

    // Packets rx in SRP. The type is interface{} with range: 0..4294967295.
    StandbyPacketsReceived interface{}

    // Number of messages sent. The type is interface{} with range: 0..4294967295.
    SentMessages interface{}

    // Number of message send failures. The type is interface{} with range:
    // 0..4294967295.
    SentMessageFailures interface{}

    // Number of RIP queries responded to. The type is interface{} with range:
    // 0..4294967295.
    QueryResponses interface{}

    // Number of periodic RIP updates. The type is interface{} with range:
    // 0..4294967295.
    PeriodicUpdates interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Number of failures to allocate memory for a route. The type is interface{}
    // with range: 0..4294967295.
    RouteMallocFailures interface{}

    // Number of failures to allocate memory for a path. The type is interface{}
    // with range: 0..4294967295.
    PathMallocFailures interface{}

    // Number of route updates to RIB made by RIP. The type is interface{} with
    // range: 0..4294967295.
    RibUpdates interface{}
}

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "default-vrf"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets}
    statistics.EntityData.Leafs["discarded-packets"] = types.YLeaf{"DiscardedPackets", statistics.DiscardedPackets}
    statistics.EntityData.Leafs["discarded-routes"] = types.YLeaf{"DiscardedRoutes", statistics.DiscardedRoutes}
    statistics.EntityData.Leafs["standby-packets-received"] = types.YLeaf{"StandbyPacketsReceived", statistics.StandbyPacketsReceived}
    statistics.EntityData.Leafs["sent-messages"] = types.YLeaf{"SentMessages", statistics.SentMessages}
    statistics.EntityData.Leafs["sent-message-failures"] = types.YLeaf{"SentMessageFailures", statistics.SentMessageFailures}
    statistics.EntityData.Leafs["query-responses"] = types.YLeaf{"QueryResponses", statistics.QueryResponses}
    statistics.EntityData.Leafs["periodic-updates"] = types.YLeaf{"PeriodicUpdates", statistics.PeriodicUpdates}
    statistics.EntityData.Leafs["route-count"] = types.YLeaf{"RouteCount", statistics.RouteCount}
    statistics.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", statistics.PathCount}
    statistics.EntityData.Leafs["route-malloc-failures"] = types.YLeaf{"RouteMallocFailures", statistics.RouteMallocFailures}
    statistics.EntityData.Leafs["path-malloc-failures"] = types.YLeaf{"PathMallocFailures", statistics.PathMallocFailures}
    statistics.EntityData.Leafs["rib-updates"] = types.YLeaf{"RibUpdates", statistics.RibUpdates}
    return &(statistics.EntityData)
}

// Rip_Protocol_DefaultVrf_Interfaces
// RIP interfaces
type Rip_Protocol_DefaultVrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular RIP interface. The type is slice of
    // Rip_Protocol_DefaultVrf_Interfaces_Interface_.
    Interface_ []Rip_Protocol_DefaultVrf_Interfaces_Interface
}

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "default-vrf"
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

// Rip_Protocol_DefaultVrf_Interfaces_Interface
// Information about a particular RIP interface
type Rip_Protocol_DefaultVrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    Interface_ interface{}

    // Interface handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    IfHandle interface{}

    // Whether RIP is enabled on this interface. The type is bool.
    RipEnabled interface{}

    // Passive interface indicator. The type is bool.
    IsPassiveInterface interface{}

    // Use broadcast address for v2 packets. The type is bool.
    MulticastAddress interface{}

    // Accept routes of metric 0 indicator. The type is bool.
    AcceptMetric interface{}

    // Versions that the interface is sending. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // Versions that the interface will recieve. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Current state of the interface. The type is InterfaceState.
    State interface{}

    // IP Address of this interface. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Prefix length of the IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Cost added to routes through this interface. The type is interface{} with
    // range: 0..4294967295.
    MetricCost interface{}

    // Split horizon enabled indicator. The type is bool.
    SplitHorizon interface{}

    // Poisoned reverse enabled indicator. The type is bool.
    PoisonHorizon interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Interface's triggered RIP neighbor. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborAddress interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // Multicast group join status. The type is bool.
    JoinStatus interface{}

    // LPTSState. The type is bool.
    LptsState interface{}

    // Authentication Mode. The type is interface{} with range: 0..4294967295.
    AuthMode interface{}

    // Authentication Keychain Name. The type is string.
    AuthKeychain interface{}

    // Authentication send key exists. The type is bool.
    SendAuthKeyExists interface{}

    // Authentication key programmed with MD5 algorithm. The type is bool.
    AuthKeyMd5 interface{}

    // Current active Send Authentication Key Id. The type is interface{} with
    // range: 0..18446744073709551615.
    AuthKeySendId interface{}

    // Total packets received. The type is interface{} with range: 0..4294967295.
    TotalPktRecvd interface{}

    // Packets dropped due to wrong keychain configured. The type is interface{}
    // with range: 0..4294967295.
    PktDropWrongKc interface{}

    // Packets dropped due to missing authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropNoAuth interface{}

    // Packets dropped due to invalid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropInvalidAuth interface{}

    // Packets accepted with valid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktAcceptedValidAuth interface{}

    // User defined summary addresses. The type is slice of
    // Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary.
    RipSummary []Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary

    // Neighbors on this interface. The type is slice of
    // Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer.
    RipPeer []Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer
}

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["rip-summary"] = types.YChild{"RipSummary", nil}
    for i := range self.RipSummary {
        self.EntityData.Children[types.GetSegmentPath(&self.RipSummary[i])] = types.YChild{"RipSummary", &self.RipSummary[i]}
    }
    self.EntityData.Children["rip-peer"] = types.YChild{"RipPeer", nil}
    for i := range self.RipPeer {
        self.EntityData.Children[types.GetSegmentPath(&self.RipPeer[i])] = types.YChild{"RipPeer", &self.RipPeer[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", self.IfHandle}
    self.EntityData.Leafs["rip-enabled"] = types.YLeaf{"RipEnabled", self.RipEnabled}
    self.EntityData.Leafs["is-passive-interface"] = types.YLeaf{"IsPassiveInterface", self.IsPassiveInterface}
    self.EntityData.Leafs["multicast-address"] = types.YLeaf{"MulticastAddress", self.MulticastAddress}
    self.EntityData.Leafs["accept-metric"] = types.YLeaf{"AcceptMetric", self.AcceptMetric}
    self.EntityData.Leafs["send-version"] = types.YLeaf{"SendVersion", self.SendVersion}
    self.EntityData.Leafs["receive-version"] = types.YLeaf{"ReceiveVersion", self.ReceiveVersion}
    self.EntityData.Leafs["state"] = types.YLeaf{"State", self.State}
    self.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", self.DestinationAddress}
    self.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", self.PrefixLength}
    self.EntityData.Leafs["metric-cost"] = types.YLeaf{"MetricCost", self.MetricCost}
    self.EntityData.Leafs["split-horizon"] = types.YLeaf{"SplitHorizon", self.SplitHorizon}
    self.EntityData.Leafs["poison-horizon"] = types.YLeaf{"PoisonHorizon", self.PoisonHorizon}
    self.EntityData.Leafs["triggered-rip"] = types.YLeaf{"TriggeredRip", self.TriggeredRip}
    self.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", self.NeighborAddress}
    self.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", self.OomFlags}
    self.EntityData.Leafs["join-status"] = types.YLeaf{"JoinStatus", self.JoinStatus}
    self.EntityData.Leafs["lpts-state"] = types.YLeaf{"LptsState", self.LptsState}
    self.EntityData.Leafs["auth-mode"] = types.YLeaf{"AuthMode", self.AuthMode}
    self.EntityData.Leafs["auth-keychain"] = types.YLeaf{"AuthKeychain", self.AuthKeychain}
    self.EntityData.Leafs["send-auth-key-exists"] = types.YLeaf{"SendAuthKeyExists", self.SendAuthKeyExists}
    self.EntityData.Leafs["auth-key-md5"] = types.YLeaf{"AuthKeyMd5", self.AuthKeyMd5}
    self.EntityData.Leafs["auth-key-send-id"] = types.YLeaf{"AuthKeySendId", self.AuthKeySendId}
    self.EntityData.Leafs["total-pkt-recvd"] = types.YLeaf{"TotalPktRecvd", self.TotalPktRecvd}
    self.EntityData.Leafs["pkt-drop-wrong-kc"] = types.YLeaf{"PktDropWrongKc", self.PktDropWrongKc}
    self.EntityData.Leafs["pkt-drop-no-auth"] = types.YLeaf{"PktDropNoAuth", self.PktDropNoAuth}
    self.EntityData.Leafs["pkt-drop-invalid-auth"] = types.YLeaf{"PktDropInvalidAuth", self.PktDropInvalidAuth}
    self.EntityData.Leafs["pkt-accepted-valid-auth"] = types.YLeaf{"PktAcceptedValidAuth", self.PktAcceptedValidAuth}
    return &(self.EntityData)
}

// Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary
// User defined summary addresses
type Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary address prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Summary address prefix length. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Summary address next hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Summary metric. The type is interface{} with range:
    // -2147483648..2147483647.
    Metric interface{}
}

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetEntityData() *types.CommonEntityData {
    ripSummary.EntityData.YFilter = ripSummary.YFilter
    ripSummary.EntityData.YangName = "rip-summary"
    ripSummary.EntityData.BundleName = "cisco_ios_xr"
    ripSummary.EntityData.ParentYangName = "interface"
    ripSummary.EntityData.SegmentPath = "rip-summary"
    ripSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ripSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ripSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ripSummary.EntityData.Children = make(map[string]types.YChild)
    ripSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    ripSummary.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ripSummary.Prefix}
    ripSummary.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", ripSummary.PrefixLength}
    ripSummary.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", ripSummary.NextHopAddress}
    ripSummary.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ripSummary.Metric}
    return &(ripSummary.EntityData)
}

// Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer
// Neighbors on this interface
type Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Uptime of this peer. The type is interface{} with range: 0..4294967295.
    PeerUptime interface{}

    // IP Address of this peer. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // RIP version for this peer. The type is interface{} with range: 0..255.
    PeerVersion interface{}

    // Discarded packets from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerPackets interface{}

    // Discarded routes from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerRoutes interface{}
}

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetEntityData() *types.CommonEntityData {
    ripPeer.EntityData.YFilter = ripPeer.YFilter
    ripPeer.EntityData.YangName = "rip-peer"
    ripPeer.EntityData.BundleName = "cisco_ios_xr"
    ripPeer.EntityData.ParentYangName = "interface"
    ripPeer.EntityData.SegmentPath = "rip-peer"
    ripPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ripPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ripPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ripPeer.EntityData.Children = make(map[string]types.YChild)
    ripPeer.EntityData.Leafs = make(map[string]types.YLeaf)
    ripPeer.EntityData.Leafs["peer-uptime"] = types.YLeaf{"PeerUptime", ripPeer.PeerUptime}
    ripPeer.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", ripPeer.PeerAddress}
    ripPeer.EntityData.Leafs["peer-version"] = types.YLeaf{"PeerVersion", ripPeer.PeerVersion}
    ripPeer.EntityData.Leafs["discarded-peer-packets"] = types.YLeaf{"DiscardedPeerPackets", ripPeer.DiscardedPeerPackets}
    ripPeer.EntityData.Leafs["discarded-peer-routes"] = types.YLeaf{"DiscardedPeerRoutes", ripPeer.DiscardedPeerRoutes}
    return &(ripPeer.EntityData)
}

// Rip_Protocol_DefaultVrf_Global
// Global Information 
type Rip_Protocol_DefaultVrf_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF summary data.
    VrfSummary Rip_Protocol_DefaultVrf_Global_VrfSummary

    // List of Interfaces configured. The type is slice of
    // Rip_Protocol_DefaultVrf_Global_InterfaceSummary.
    InterfaceSummary []Rip_Protocol_DefaultVrf_Global_InterfaceSummary
}

func (global *Rip_Protocol_DefaultVrf_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "default-vrf"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["vrf-summary"] = types.YChild{"VrfSummary", &global.VrfSummary}
    global.EntityData.Children["interface-summary"] = types.YChild{"InterfaceSummary", nil}
    for i := range global.InterfaceSummary {
        global.EntityData.Children[types.GetSegmentPath(&global.InterfaceSummary[i])] = types.YChild{"InterfaceSummary", &global.InterfaceSummary[i]}
    }
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(global.EntityData)
}

// Rip_Protocol_DefaultVrf_Global_VrfSummary
// VRF summary data
type Rip_Protocol_DefaultVrf_Global_VrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Active indicator. The type is bool.
    Active interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Number of interfaces configured. The type is interface{} with range:
    // 0..4294967295.
    InterfaceConfiguredCount interface{}

    // Number of active interfaces. The type is interface{} with range:
    // 0..4294967295.
    ActiveInterfaceCount interface{}
}

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetEntityData() *types.CommonEntityData {
    vrfSummary.EntityData.YFilter = vrfSummary.YFilter
    vrfSummary.EntityData.YangName = "vrf-summary"
    vrfSummary.EntityData.BundleName = "cisco_ios_xr"
    vrfSummary.EntityData.ParentYangName = "global"
    vrfSummary.EntityData.SegmentPath = "vrf-summary"
    vrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummary.EntityData.Children = make(map[string]types.YChild)
    vrfSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfSummary.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrfSummary.VrfName}
    vrfSummary.EntityData.Leafs["active"] = types.YLeaf{"Active", vrfSummary.Active}
    vrfSummary.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", vrfSummary.OomFlags}
    vrfSummary.EntityData.Leafs["route-count"] = types.YLeaf{"RouteCount", vrfSummary.RouteCount}
    vrfSummary.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", vrfSummary.PathCount}
    vrfSummary.EntityData.Leafs["update-timer"] = types.YLeaf{"UpdateTimer", vrfSummary.UpdateTimer}
    vrfSummary.EntityData.Leafs["next-update-time"] = types.YLeaf{"NextUpdateTime", vrfSummary.NextUpdateTime}
    vrfSummary.EntityData.Leafs["invalid-timer"] = types.YLeaf{"InvalidTimer", vrfSummary.InvalidTimer}
    vrfSummary.EntityData.Leafs["hold-down-timer"] = types.YLeaf{"HoldDownTimer", vrfSummary.HoldDownTimer}
    vrfSummary.EntityData.Leafs["flush-timer"] = types.YLeaf{"FlushTimer", vrfSummary.FlushTimer}
    vrfSummary.EntityData.Leafs["interface-configured-count"] = types.YLeaf{"InterfaceConfiguredCount", vrfSummary.InterfaceConfiguredCount}
    vrfSummary.EntityData.Leafs["active-interface-count"] = types.YLeaf{"ActiveInterfaceCount", vrfSummary.ActiveInterfaceCount}
    return &(vrfSummary.EntityData)
}

// Rip_Protocol_DefaultVrf_Global_InterfaceSummary
// List of Interfaces configured
type Rip_Protocol_DefaultVrf_Global_InterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // RIP enabled indicator. The type is bool.
    Enabled interface{}

    // Interface state. The type is InterfaceState.
    State interface{}

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // RIP versions this interface sends out. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // RIP versions this interface will receive. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Number of neighbors seen. The type is interface{} with range:
    // 0..4294967295.
    NeighborCount interface{}
}

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetEntityData() *types.CommonEntityData {
    interfaceSummary.EntityData.YFilter = interfaceSummary.YFilter
    interfaceSummary.EntityData.YangName = "interface-summary"
    interfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummary.EntityData.ParentYangName = "global"
    interfaceSummary.EntityData.SegmentPath = "interface-summary"
    interfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummary.EntityData.Children = make(map[string]types.YChild)
    interfaceSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceSummary.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceSummary.InterfaceName}
    interfaceSummary.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", interfaceSummary.Enabled}
    interfaceSummary.EntityData.Leafs["state"] = types.YLeaf{"State", interfaceSummary.State}
    interfaceSummary.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", interfaceSummary.DestinationAddress}
    interfaceSummary.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", interfaceSummary.PrefixLength}
    interfaceSummary.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", interfaceSummary.OomFlags}
    interfaceSummary.EntityData.Leafs["send-version"] = types.YLeaf{"SendVersion", interfaceSummary.SendVersion}
    interfaceSummary.EntityData.Leafs["receive-version"] = types.YLeaf{"ReceiveVersion", interfaceSummary.ReceiveVersion}
    interfaceSummary.EntityData.Leafs["neighbor-count"] = types.YLeaf{"NeighborCount", interfaceSummary.NeighborCount}
    return &(interfaceSummary.EntityData)
}

// Rip_DefaultVrf
// RIP operational data for Default VRF
type Rip_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RIP route database.
    Routes Rip_DefaultVrf_Routes

    // RIP global configuration.
    Configuration Rip_DefaultVrf_Configuration

    // RIP statistics information.
    Statistics Rip_DefaultVrf_Statistics

    // RIP interfaces.
    Interfaces Rip_DefaultVrf_Interfaces

    // Global Information .
    Global Rip_DefaultVrf_Global
}

func (defaultVrf *Rip_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "rip"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = make(map[string]types.YChild)
    defaultVrf.EntityData.Children["routes"] = types.YChild{"Routes", &defaultVrf.Routes}
    defaultVrf.EntityData.Children["configuration"] = types.YChild{"Configuration", &defaultVrf.Configuration}
    defaultVrf.EntityData.Children["statistics"] = types.YChild{"Statistics", &defaultVrf.Statistics}
    defaultVrf.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &defaultVrf.Interfaces}
    defaultVrf.EntityData.Children["global"] = types.YChild{"Global", &defaultVrf.Global}
    defaultVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(defaultVrf.EntityData)
}

// Rip_DefaultVrf_Routes
// RIP route database
type Rip_DefaultVrf_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A route in the RIP database. The type is slice of
    // Rip_DefaultVrf_Routes_Route.
    Route []Rip_DefaultVrf_Routes_Route
}

func (routes *Rip_DefaultVrf_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "default-vrf"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = make(map[string]types.YChild)
    routes.EntityData.Children["route"] = types.YChild{"Route", nil}
    for i := range routes.Route {
        routes.EntityData.Children[types.GetSegmentPath(&routes.Route[i])] = types.YChild{"Route", &routes.Route[i]}
    }
    routes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routes.EntityData)
}

// Rip_DefaultVrf_Routes_Route
// A route in the RIP database
type Rip_DefaultVrf_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // Destination IP Address for this route. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLengthXr interface{}

    // Route administrative distance. The type is interface{} with range:
    // 0..65535.
    Distance interface{}

    // Hop count for this route. The type is interface{} with range: 0..65535.
    BgpCount interface{}

    // Type of this route. The type is interface{} with range: 0..65535.
    RouteType interface{}

    // Summary route placeholder indicator. The type is bool.
    RouteSummary interface{}

    // Generic route information. The type is interface{} with range: 0..65535.
    RouteTag interface{}

    // RIB supplied version number. The type is interface{} with range: 0..255.
    Version interface{}

    // RIB supplied route attributes. The type is interface{} with range:
    // 0..4294967295.
    Attributes interface{}

    // Active route indicator. The type is bool.
    Active interface{}

    // Where this route was learnt. The type is RipRouteOrigin.
    PathOrigin interface{}

    // Indicates whether route is in holddown. The type is bool.
    HoldDown interface{}

    // The paths for this route. The type is slice of
    // Rip_DefaultVrf_Routes_Route_Paths.
    Paths []Rip_DefaultVrf_Routes_Route_Paths
}

func (route *Rip_DefaultVrf_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = make(map[string]types.YChild)
    route.EntityData.Children["paths"] = types.YChild{"Paths", nil}
    for i := range route.Paths {
        route.EntityData.Children[types.GetSegmentPath(&route.Paths[i])] = types.YChild{"Paths", &route.Paths[i]}
    }
    route.EntityData.Leafs = make(map[string]types.YLeaf)
    route.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", route.Prefix}
    route.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", route.PrefixLength}
    route.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", route.DestinationAddress}
    route.EntityData.Leafs["prefix-length-xr"] = types.YLeaf{"PrefixLengthXr", route.PrefixLengthXr}
    route.EntityData.Leafs["distance"] = types.YLeaf{"Distance", route.Distance}
    route.EntityData.Leafs["bgp-count"] = types.YLeaf{"BgpCount", route.BgpCount}
    route.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", route.RouteType}
    route.EntityData.Leafs["route-summary"] = types.YLeaf{"RouteSummary", route.RouteSummary}
    route.EntityData.Leafs["route-tag"] = types.YLeaf{"RouteTag", route.RouteTag}
    route.EntityData.Leafs["version"] = types.YLeaf{"Version", route.Version}
    route.EntityData.Leafs["attributes"] = types.YLeaf{"Attributes", route.Attributes}
    route.EntityData.Leafs["active"] = types.YLeaf{"Active", route.Active}
    route.EntityData.Leafs["path-origin"] = types.YLeaf{"PathOrigin", route.PathOrigin}
    route.EntityData.Leafs["hold-down"] = types.YLeaf{"HoldDown", route.HoldDown}
    return &(route.EntityData)
}

// Rip_DefaultVrf_Routes_Route_Paths
// The paths for this route
type Rip_DefaultVrf_Routes_Route_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Next hop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}

    // Tag. The type is interface{} with range: 0..65535.
    Tag interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Up time. The type is interface{} with range: 0..4294967295.
    Uptime interface{}

    // Permanent indicator. The type is bool.
    IsPermanent interface{}
}

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "route"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    paths.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", paths.SourceAddress}
    paths.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", paths.NextHopAddress}
    paths.EntityData.Leafs["metric"] = types.YLeaf{"Metric", paths.Metric}
    paths.EntityData.Leafs["tag"] = types.YLeaf{"Tag", paths.Tag}
    paths.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", paths.Interface_}
    paths.EntityData.Leafs["uptime"] = types.YLeaf{"Uptime", paths.Uptime}
    paths.EntityData.Leafs["is-permanent"] = types.YLeaf{"IsPermanent", paths.IsPermanent}
    return &(paths.EntityData)
}

// Rip_DefaultVrf_Configuration
// RIP global configuration
type Rip_DefaultVrf_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF active indicator. The type is bool.
    Active interface{}

    // VRF added to socket indicator. The type is bool.
    VrFisedSocket interface{}

    // Version of RIP configured. The type is interface{} with range:
    // -2147483648..2147483647.
    RipVersion interface{}

    // Default metric for RIP routes. The type is interface{} with range: 0..255.
    DefaultMetric interface{}

    // Maximum number of paths a route can have. The type is interface{} with
    // range: 0..255.
    MaximumPaths interface{}

    // Auto-summarization indicator. The type is bool.
    AutoSummarize interface{}

    // Use broadcast/multicast address indicator. The type is bool.
    MulticastAddress interface{}

    // Flash update threshold. The type is interface{} with range: 0..255.
    FlashThreshold interface{}

    // Length of the input queue. The type is interface{} with range: 0..65535.
    InputQLength interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Incoming packet source validation indicator. The type is bool.
    ValidationIndicator interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // NSF Enable status. The type is bool.
    NsfStatus interface{}

    // NSF life time. The type is interface{} with range: 0..4294967295.
    NsfLifeTime interface{}
}

func (configuration *Rip_DefaultVrf_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "default-vrf"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = make(map[string]types.YChild)
    configuration.EntityData.Leafs = make(map[string]types.YLeaf)
    configuration.EntityData.Leafs["active"] = types.YLeaf{"Active", configuration.Active}
    configuration.EntityData.Leafs["vr-fised-socket"] = types.YLeaf{"VrFisedSocket", configuration.VrFisedSocket}
    configuration.EntityData.Leafs["rip-version"] = types.YLeaf{"RipVersion", configuration.RipVersion}
    configuration.EntityData.Leafs["default-metric"] = types.YLeaf{"DefaultMetric", configuration.DefaultMetric}
    configuration.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", configuration.MaximumPaths}
    configuration.EntityData.Leafs["auto-summarize"] = types.YLeaf{"AutoSummarize", configuration.AutoSummarize}
    configuration.EntityData.Leafs["multicast-address"] = types.YLeaf{"MulticastAddress", configuration.MulticastAddress}
    configuration.EntityData.Leafs["flash-threshold"] = types.YLeaf{"FlashThreshold", configuration.FlashThreshold}
    configuration.EntityData.Leafs["input-q-length"] = types.YLeaf{"InputQLength", configuration.InputQLength}
    configuration.EntityData.Leafs["triggered-rip"] = types.YLeaf{"TriggeredRip", configuration.TriggeredRip}
    configuration.EntityData.Leafs["validation-indicator"] = types.YLeaf{"ValidationIndicator", configuration.ValidationIndicator}
    configuration.EntityData.Leafs["update-timer"] = types.YLeaf{"UpdateTimer", configuration.UpdateTimer}
    configuration.EntityData.Leafs["next-update-time"] = types.YLeaf{"NextUpdateTime", configuration.NextUpdateTime}
    configuration.EntityData.Leafs["invalid-timer"] = types.YLeaf{"InvalidTimer", configuration.InvalidTimer}
    configuration.EntityData.Leafs["hold-down-timer"] = types.YLeaf{"HoldDownTimer", configuration.HoldDownTimer}
    configuration.EntityData.Leafs["flush-timer"] = types.YLeaf{"FlushTimer", configuration.FlushTimer}
    configuration.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", configuration.OomFlags}
    configuration.EntityData.Leafs["nsf-status"] = types.YLeaf{"NsfStatus", configuration.NsfStatus}
    configuration.EntityData.Leafs["nsf-life-time"] = types.YLeaf{"NsfLifeTime", configuration.NsfLifeTime}
    return &(configuration.EntityData)
}

// Rip_DefaultVrf_Statistics
// RIP statistics information
type Rip_DefaultVrf_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total packets received. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Total discarded packets. The type is interface{} with range: 0..4294967295.
    DiscardedPackets interface{}

    // Total discarded routes. The type is interface{} with range: 0..4294967295.
    DiscardedRoutes interface{}

    // Packets rx in SRP. The type is interface{} with range: 0..4294967295.
    StandbyPacketsReceived interface{}

    // Number of messages sent. The type is interface{} with range: 0..4294967295.
    SentMessages interface{}

    // Number of message send failures. The type is interface{} with range:
    // 0..4294967295.
    SentMessageFailures interface{}

    // Number of RIP queries responded to. The type is interface{} with range:
    // 0..4294967295.
    QueryResponses interface{}

    // Number of periodic RIP updates. The type is interface{} with range:
    // 0..4294967295.
    PeriodicUpdates interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Number of failures to allocate memory for a route. The type is interface{}
    // with range: 0..4294967295.
    RouteMallocFailures interface{}

    // Number of failures to allocate memory for a path. The type is interface{}
    // with range: 0..4294967295.
    PathMallocFailures interface{}

    // Number of route updates to RIB made by RIP. The type is interface{} with
    // range: 0..4294967295.
    RibUpdates interface{}
}

func (statistics *Rip_DefaultVrf_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "default-vrf"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets}
    statistics.EntityData.Leafs["discarded-packets"] = types.YLeaf{"DiscardedPackets", statistics.DiscardedPackets}
    statistics.EntityData.Leafs["discarded-routes"] = types.YLeaf{"DiscardedRoutes", statistics.DiscardedRoutes}
    statistics.EntityData.Leafs["standby-packets-received"] = types.YLeaf{"StandbyPacketsReceived", statistics.StandbyPacketsReceived}
    statistics.EntityData.Leafs["sent-messages"] = types.YLeaf{"SentMessages", statistics.SentMessages}
    statistics.EntityData.Leafs["sent-message-failures"] = types.YLeaf{"SentMessageFailures", statistics.SentMessageFailures}
    statistics.EntityData.Leafs["query-responses"] = types.YLeaf{"QueryResponses", statistics.QueryResponses}
    statistics.EntityData.Leafs["periodic-updates"] = types.YLeaf{"PeriodicUpdates", statistics.PeriodicUpdates}
    statistics.EntityData.Leafs["route-count"] = types.YLeaf{"RouteCount", statistics.RouteCount}
    statistics.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", statistics.PathCount}
    statistics.EntityData.Leafs["route-malloc-failures"] = types.YLeaf{"RouteMallocFailures", statistics.RouteMallocFailures}
    statistics.EntityData.Leafs["path-malloc-failures"] = types.YLeaf{"PathMallocFailures", statistics.PathMallocFailures}
    statistics.EntityData.Leafs["rib-updates"] = types.YLeaf{"RibUpdates", statistics.RibUpdates}
    return &(statistics.EntityData)
}

// Rip_DefaultVrf_Interfaces
// RIP interfaces
type Rip_DefaultVrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular RIP interface. The type is slice of
    // Rip_DefaultVrf_Interfaces_Interface_.
    Interface_ []Rip_DefaultVrf_Interfaces_Interface
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "default-vrf"
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

// Rip_DefaultVrf_Interfaces_Interface
// Information about a particular RIP interface
type Rip_DefaultVrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    Interface_ interface{}

    // Interface handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    IfHandle interface{}

    // Whether RIP is enabled on this interface. The type is bool.
    RipEnabled interface{}

    // Passive interface indicator. The type is bool.
    IsPassiveInterface interface{}

    // Use broadcast address for v2 packets. The type is bool.
    MulticastAddress interface{}

    // Accept routes of metric 0 indicator. The type is bool.
    AcceptMetric interface{}

    // Versions that the interface is sending. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // Versions that the interface will recieve. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Current state of the interface. The type is InterfaceState.
    State interface{}

    // IP Address of this interface. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Prefix length of the IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Cost added to routes through this interface. The type is interface{} with
    // range: 0..4294967295.
    MetricCost interface{}

    // Split horizon enabled indicator. The type is bool.
    SplitHorizon interface{}

    // Poisoned reverse enabled indicator. The type is bool.
    PoisonHorizon interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Interface's triggered RIP neighbor. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborAddress interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // Multicast group join status. The type is bool.
    JoinStatus interface{}

    // LPTSState. The type is bool.
    LptsState interface{}

    // Authentication Mode. The type is interface{} with range: 0..4294967295.
    AuthMode interface{}

    // Authentication Keychain Name. The type is string.
    AuthKeychain interface{}

    // Authentication send key exists. The type is bool.
    SendAuthKeyExists interface{}

    // Authentication key programmed with MD5 algorithm. The type is bool.
    AuthKeyMd5 interface{}

    // Current active Send Authentication Key Id. The type is interface{} with
    // range: 0..18446744073709551615.
    AuthKeySendId interface{}

    // Total packets received. The type is interface{} with range: 0..4294967295.
    TotalPktRecvd interface{}

    // Packets dropped due to wrong keychain configured. The type is interface{}
    // with range: 0..4294967295.
    PktDropWrongKc interface{}

    // Packets dropped due to missing authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropNoAuth interface{}

    // Packets dropped due to invalid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropInvalidAuth interface{}

    // Packets accepted with valid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktAcceptedValidAuth interface{}

    // User defined summary addresses. The type is slice of
    // Rip_DefaultVrf_Interfaces_Interface_RipSummary.
    RipSummary []Rip_DefaultVrf_Interfaces_Interface_RipSummary

    // Neighbors on this interface. The type is slice of
    // Rip_DefaultVrf_Interfaces_Interface_RipPeer.
    RipPeer []Rip_DefaultVrf_Interfaces_Interface_RipPeer
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["rip-summary"] = types.YChild{"RipSummary", nil}
    for i := range self.RipSummary {
        self.EntityData.Children[types.GetSegmentPath(&self.RipSummary[i])] = types.YChild{"RipSummary", &self.RipSummary[i]}
    }
    self.EntityData.Children["rip-peer"] = types.YChild{"RipPeer", nil}
    for i := range self.RipPeer {
        self.EntityData.Children[types.GetSegmentPath(&self.RipPeer[i])] = types.YChild{"RipPeer", &self.RipPeer[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", self.IfHandle}
    self.EntityData.Leafs["rip-enabled"] = types.YLeaf{"RipEnabled", self.RipEnabled}
    self.EntityData.Leafs["is-passive-interface"] = types.YLeaf{"IsPassiveInterface", self.IsPassiveInterface}
    self.EntityData.Leafs["multicast-address"] = types.YLeaf{"MulticastAddress", self.MulticastAddress}
    self.EntityData.Leafs["accept-metric"] = types.YLeaf{"AcceptMetric", self.AcceptMetric}
    self.EntityData.Leafs["send-version"] = types.YLeaf{"SendVersion", self.SendVersion}
    self.EntityData.Leafs["receive-version"] = types.YLeaf{"ReceiveVersion", self.ReceiveVersion}
    self.EntityData.Leafs["state"] = types.YLeaf{"State", self.State}
    self.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", self.DestinationAddress}
    self.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", self.PrefixLength}
    self.EntityData.Leafs["metric-cost"] = types.YLeaf{"MetricCost", self.MetricCost}
    self.EntityData.Leafs["split-horizon"] = types.YLeaf{"SplitHorizon", self.SplitHorizon}
    self.EntityData.Leafs["poison-horizon"] = types.YLeaf{"PoisonHorizon", self.PoisonHorizon}
    self.EntityData.Leafs["triggered-rip"] = types.YLeaf{"TriggeredRip", self.TriggeredRip}
    self.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", self.NeighborAddress}
    self.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", self.OomFlags}
    self.EntityData.Leafs["join-status"] = types.YLeaf{"JoinStatus", self.JoinStatus}
    self.EntityData.Leafs["lpts-state"] = types.YLeaf{"LptsState", self.LptsState}
    self.EntityData.Leafs["auth-mode"] = types.YLeaf{"AuthMode", self.AuthMode}
    self.EntityData.Leafs["auth-keychain"] = types.YLeaf{"AuthKeychain", self.AuthKeychain}
    self.EntityData.Leafs["send-auth-key-exists"] = types.YLeaf{"SendAuthKeyExists", self.SendAuthKeyExists}
    self.EntityData.Leafs["auth-key-md5"] = types.YLeaf{"AuthKeyMd5", self.AuthKeyMd5}
    self.EntityData.Leafs["auth-key-send-id"] = types.YLeaf{"AuthKeySendId", self.AuthKeySendId}
    self.EntityData.Leafs["total-pkt-recvd"] = types.YLeaf{"TotalPktRecvd", self.TotalPktRecvd}
    self.EntityData.Leafs["pkt-drop-wrong-kc"] = types.YLeaf{"PktDropWrongKc", self.PktDropWrongKc}
    self.EntityData.Leafs["pkt-drop-no-auth"] = types.YLeaf{"PktDropNoAuth", self.PktDropNoAuth}
    self.EntityData.Leafs["pkt-drop-invalid-auth"] = types.YLeaf{"PktDropInvalidAuth", self.PktDropInvalidAuth}
    self.EntityData.Leafs["pkt-accepted-valid-auth"] = types.YLeaf{"PktAcceptedValidAuth", self.PktAcceptedValidAuth}
    return &(self.EntityData)
}

// Rip_DefaultVrf_Interfaces_Interface_RipSummary
// User defined summary addresses
type Rip_DefaultVrf_Interfaces_Interface_RipSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary address prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Summary address prefix length. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Summary address next hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Summary metric. The type is interface{} with range:
    // -2147483648..2147483647.
    Metric interface{}
}

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetEntityData() *types.CommonEntityData {
    ripSummary.EntityData.YFilter = ripSummary.YFilter
    ripSummary.EntityData.YangName = "rip-summary"
    ripSummary.EntityData.BundleName = "cisco_ios_xr"
    ripSummary.EntityData.ParentYangName = "interface"
    ripSummary.EntityData.SegmentPath = "rip-summary"
    ripSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ripSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ripSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ripSummary.EntityData.Children = make(map[string]types.YChild)
    ripSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    ripSummary.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ripSummary.Prefix}
    ripSummary.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", ripSummary.PrefixLength}
    ripSummary.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", ripSummary.NextHopAddress}
    ripSummary.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ripSummary.Metric}
    return &(ripSummary.EntityData)
}

// Rip_DefaultVrf_Interfaces_Interface_RipPeer
// Neighbors on this interface
type Rip_DefaultVrf_Interfaces_Interface_RipPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Uptime of this peer. The type is interface{} with range: 0..4294967295.
    PeerUptime interface{}

    // IP Address of this peer. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // RIP version for this peer. The type is interface{} with range: 0..255.
    PeerVersion interface{}

    // Discarded packets from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerPackets interface{}

    // Discarded routes from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerRoutes interface{}
}

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetEntityData() *types.CommonEntityData {
    ripPeer.EntityData.YFilter = ripPeer.YFilter
    ripPeer.EntityData.YangName = "rip-peer"
    ripPeer.EntityData.BundleName = "cisco_ios_xr"
    ripPeer.EntityData.ParentYangName = "interface"
    ripPeer.EntityData.SegmentPath = "rip-peer"
    ripPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ripPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ripPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ripPeer.EntityData.Children = make(map[string]types.YChild)
    ripPeer.EntityData.Leafs = make(map[string]types.YLeaf)
    ripPeer.EntityData.Leafs["peer-uptime"] = types.YLeaf{"PeerUptime", ripPeer.PeerUptime}
    ripPeer.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", ripPeer.PeerAddress}
    ripPeer.EntityData.Leafs["peer-version"] = types.YLeaf{"PeerVersion", ripPeer.PeerVersion}
    ripPeer.EntityData.Leafs["discarded-peer-packets"] = types.YLeaf{"DiscardedPeerPackets", ripPeer.DiscardedPeerPackets}
    ripPeer.EntityData.Leafs["discarded-peer-routes"] = types.YLeaf{"DiscardedPeerRoutes", ripPeer.DiscardedPeerRoutes}
    return &(ripPeer.EntityData)
}

// Rip_DefaultVrf_Global
// Global Information 
type Rip_DefaultVrf_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF summary data.
    VrfSummary Rip_DefaultVrf_Global_VrfSummary

    // List of Interfaces configured. The type is slice of
    // Rip_DefaultVrf_Global_InterfaceSummary.
    InterfaceSummary []Rip_DefaultVrf_Global_InterfaceSummary
}

func (global *Rip_DefaultVrf_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "default-vrf"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["vrf-summary"] = types.YChild{"VrfSummary", &global.VrfSummary}
    global.EntityData.Children["interface-summary"] = types.YChild{"InterfaceSummary", nil}
    for i := range global.InterfaceSummary {
        global.EntityData.Children[types.GetSegmentPath(&global.InterfaceSummary[i])] = types.YChild{"InterfaceSummary", &global.InterfaceSummary[i]}
    }
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(global.EntityData)
}

// Rip_DefaultVrf_Global_VrfSummary
// VRF summary data
type Rip_DefaultVrf_Global_VrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Active indicator. The type is bool.
    Active interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Number of interfaces configured. The type is interface{} with range:
    // 0..4294967295.
    InterfaceConfiguredCount interface{}

    // Number of active interfaces. The type is interface{} with range:
    // 0..4294967295.
    ActiveInterfaceCount interface{}
}

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetEntityData() *types.CommonEntityData {
    vrfSummary.EntityData.YFilter = vrfSummary.YFilter
    vrfSummary.EntityData.YangName = "vrf-summary"
    vrfSummary.EntityData.BundleName = "cisco_ios_xr"
    vrfSummary.EntityData.ParentYangName = "global"
    vrfSummary.EntityData.SegmentPath = "vrf-summary"
    vrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummary.EntityData.Children = make(map[string]types.YChild)
    vrfSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfSummary.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrfSummary.VrfName}
    vrfSummary.EntityData.Leafs["active"] = types.YLeaf{"Active", vrfSummary.Active}
    vrfSummary.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", vrfSummary.OomFlags}
    vrfSummary.EntityData.Leafs["route-count"] = types.YLeaf{"RouteCount", vrfSummary.RouteCount}
    vrfSummary.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", vrfSummary.PathCount}
    vrfSummary.EntityData.Leafs["update-timer"] = types.YLeaf{"UpdateTimer", vrfSummary.UpdateTimer}
    vrfSummary.EntityData.Leafs["next-update-time"] = types.YLeaf{"NextUpdateTime", vrfSummary.NextUpdateTime}
    vrfSummary.EntityData.Leafs["invalid-timer"] = types.YLeaf{"InvalidTimer", vrfSummary.InvalidTimer}
    vrfSummary.EntityData.Leafs["hold-down-timer"] = types.YLeaf{"HoldDownTimer", vrfSummary.HoldDownTimer}
    vrfSummary.EntityData.Leafs["flush-timer"] = types.YLeaf{"FlushTimer", vrfSummary.FlushTimer}
    vrfSummary.EntityData.Leafs["interface-configured-count"] = types.YLeaf{"InterfaceConfiguredCount", vrfSummary.InterfaceConfiguredCount}
    vrfSummary.EntityData.Leafs["active-interface-count"] = types.YLeaf{"ActiveInterfaceCount", vrfSummary.ActiveInterfaceCount}
    return &(vrfSummary.EntityData)
}

// Rip_DefaultVrf_Global_InterfaceSummary
// List of Interfaces configured
type Rip_DefaultVrf_Global_InterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // RIP enabled indicator. The type is bool.
    Enabled interface{}

    // Interface state. The type is InterfaceState.
    State interface{}

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // RIP versions this interface sends out. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // RIP versions this interface will receive. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Number of neighbors seen. The type is interface{} with range:
    // 0..4294967295.
    NeighborCount interface{}
}

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetEntityData() *types.CommonEntityData {
    interfaceSummary.EntityData.YFilter = interfaceSummary.YFilter
    interfaceSummary.EntityData.YangName = "interface-summary"
    interfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummary.EntityData.ParentYangName = "global"
    interfaceSummary.EntityData.SegmentPath = "interface-summary"
    interfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummary.EntityData.Children = make(map[string]types.YChild)
    interfaceSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceSummary.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceSummary.InterfaceName}
    interfaceSummary.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", interfaceSummary.Enabled}
    interfaceSummary.EntityData.Leafs["state"] = types.YLeaf{"State", interfaceSummary.State}
    interfaceSummary.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", interfaceSummary.DestinationAddress}
    interfaceSummary.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", interfaceSummary.PrefixLength}
    interfaceSummary.EntityData.Leafs["oom-flags"] = types.YLeaf{"OomFlags", interfaceSummary.OomFlags}
    interfaceSummary.EntityData.Leafs["send-version"] = types.YLeaf{"SendVersion", interfaceSummary.SendVersion}
    interfaceSummary.EntityData.Leafs["receive-version"] = types.YLeaf{"ReceiveVersion", interfaceSummary.ReceiveVersion}
    interfaceSummary.EntityData.Leafs["neighbor-count"] = types.YLeaf{"NeighborCount", interfaceSummary.NeighborCount}
    return &(interfaceSummary.EntityData)
}

