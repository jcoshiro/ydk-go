// This module contains a collection of YANG definitions
// for Cisco IOS-XR manageability-object-tracking package operational data.
// 
// This module contains definitions
// for the following management objects:
//   object-tracking: Object Tracking operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package manageability_object_tracking_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package manageability_object_tracking_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-manageability-object-tracking-oper object-tracking}", reflect.TypeOf(ObjectTracking{}))
    ydk.RegisterEntity("Cisco-IOS-XR-manageability-object-tracking-oper:object-tracking", reflect.TypeOf(ObjectTracking{}))
}

// Track represents Track
type Track string

const (
    // Line protocol type
    Track_interface_type Track = "interface-type"

    // Route type
    Track_route_type Track = "route-type"

    // Boolean and type
    Track_bool_and_type Track = "bool-and-type"

    // Boolean or type
    Track_bool_or_type Track = "bool-or-type"

    // Ipsla track type
    Track_ipsla_type Track = "ipsla-type"

    // type undefined
    Track_undefined_type Track = "undefined-type"

    // type threshold weight
    Track_threshold_weight Track = "threshold-weight"

    // type threshold percentage
    Track_threshold_percentage Track = "threshold-percentage"

    // type bfd rtr
    Track_bfd_type Track = "bfd-type"
)

// ObjectTracking
// Object Tracking operational data
type ObjectTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object Tracking Type interface info.
    TrackTypeInterface ObjectTracking_TrackTypeInterface

    // Object Tracking Track table brief.
    TrackBriefs ObjectTracking_TrackBriefs

    // Object Tracking Type RTR Reachability info.
    TrackTypeRtrReachability ObjectTracking_TrackTypeRtrReachability

    // Object Tracking Type RTR Reachability brief info.
    TrackTypeRtrReachabilityBrief ObjectTracking_TrackTypeRtrReachabilityBrief

    // Object Tracking Track table.
    Tracks ObjectTracking_Tracks

    // Object Tracking Type Ipv4 Route brief info.
    TrackTypeIpv4RouteBrief ObjectTracking_TrackTypeIpv4RouteBrief

    // Object Tracking Type IPV4 route info.
    TrackTypeIpv4Route ObjectTracking_TrackTypeIpv4Route

    // Object Tracking Type Interface brief info.
    TrackTypeInterfaceBrief ObjectTracking_TrackTypeInterfaceBrief
}

func (objectTracking *ObjectTracking) GetEntityData() *types.CommonEntityData {
    objectTracking.EntityData.YFilter = objectTracking.YFilter
    objectTracking.EntityData.YangName = "object-tracking"
    objectTracking.EntityData.BundleName = "cisco_ios_xr"
    objectTracking.EntityData.ParentYangName = "Cisco-IOS-XR-manageability-object-tracking-oper"
    objectTracking.EntityData.SegmentPath = "Cisco-IOS-XR-manageability-object-tracking-oper:object-tracking"
    objectTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectTracking.EntityData.Children = make(map[string]types.YChild)
    objectTracking.EntityData.Children["track-type-interface"] = types.YChild{"TrackTypeInterface", &objectTracking.TrackTypeInterface}
    objectTracking.EntityData.Children["track-briefs"] = types.YChild{"TrackBriefs", &objectTracking.TrackBriefs}
    objectTracking.EntityData.Children["track-type-rtr-reachability"] = types.YChild{"TrackTypeRtrReachability", &objectTracking.TrackTypeRtrReachability}
    objectTracking.EntityData.Children["track-type-rtr-reachability-brief"] = types.YChild{"TrackTypeRtrReachabilityBrief", &objectTracking.TrackTypeRtrReachabilityBrief}
    objectTracking.EntityData.Children["tracks"] = types.YChild{"Tracks", &objectTracking.Tracks}
    objectTracking.EntityData.Children["track-type-ipv4-route-brief"] = types.YChild{"TrackTypeIpv4RouteBrief", &objectTracking.TrackTypeIpv4RouteBrief}
    objectTracking.EntityData.Children["track-type-ipv4-route"] = types.YChild{"TrackTypeIpv4Route", &objectTracking.TrackTypeIpv4Route}
    objectTracking.EntityData.Children["track-type-interface-brief"] = types.YChild{"TrackTypeInterfaceBrief", &objectTracking.TrackTypeInterfaceBrief}
    objectTracking.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(objectTracking.EntityData)
}

// ObjectTracking_TrackTypeInterface
// Object Tracking Type interface info
type ObjectTracking_TrackTypeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // track info. The type is slice of
    // ObjectTracking_TrackTypeInterface_TrackInfo.
    TrackInfo []ObjectTracking_TrackTypeInterface_TrackInfo
}

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetEntityData() *types.CommonEntityData {
    trackTypeInterface.EntityData.YFilter = trackTypeInterface.YFilter
    trackTypeInterface.EntityData.YangName = "track-type-interface"
    trackTypeInterface.EntityData.BundleName = "cisco_ios_xr"
    trackTypeInterface.EntityData.ParentYangName = "object-tracking"
    trackTypeInterface.EntityData.SegmentPath = "track-type-interface"
    trackTypeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeInterface.EntityData.Children = make(map[string]types.YChild)
    trackTypeInterface.EntityData.Children["track-info"] = types.YChild{"TrackInfo", nil}
    for i := range trackTypeInterface.TrackInfo {
        trackTypeInterface.EntityData.Children[types.GetSegmentPath(&trackTypeInterface.TrackInfo[i])] = types.YChild{"TrackInfo", &trackTypeInterface.TrackInfo[i]}
    }
    trackTypeInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackTypeInterface.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo
// track info
type ObjectTracking_TrackTypeInterface_TrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type_ interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // State Change Counter. The type is interface{} with range: 0..4294967295.
    StateChangeCounter interface{}

    // Seconds Last Change. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecondsLastChange interface{}

    // User specified threshold upper limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdUp interface{}

    // User specified threshold lower limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdDown interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo

    // boolean objects.
    BoolTracks ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks

    // Threshold objects.
    ThresholdTracks ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks

    // Tracking Interfaces.
    TrackingInteraces ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces

    // Is the state change delay counter in progress.
    Delayed ObjectTracking_TrackTypeInterface_TrackInfo_Delayed
}

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetEntityData() *types.CommonEntityData {
    trackInfo.EntityData.YFilter = trackInfo.YFilter
    trackInfo.EntityData.YangName = "track-info"
    trackInfo.EntityData.BundleName = "cisco_ios_xr"
    trackInfo.EntityData.ParentYangName = "track-type-interface"
    trackInfo.EntityData.SegmentPath = "track-info"
    trackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackInfo.EntityData.Children = make(map[string]types.YChild)
    trackInfo.EntityData.Children["track-type-info"] = types.YChild{"TrackTypeInfo", &trackInfo.TrackTypeInfo}
    trackInfo.EntityData.Children["bool-tracks"] = types.YChild{"BoolTracks", &trackInfo.BoolTracks}
    trackInfo.EntityData.Children["threshold-tracks"] = types.YChild{"ThresholdTracks", &trackInfo.ThresholdTracks}
    trackInfo.EntityData.Children["tracking-interaces"] = types.YChild{"TrackingInteraces", &trackInfo.TrackingInteraces}
    trackInfo.EntityData.Children["delayed"] = types.YChild{"Delayed", &trackInfo.Delayed}
    trackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackInfo.EntityData.Leafs["tracke-name"] = types.YLeaf{"TrackeName", trackInfo.TrackeName}
    trackInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", trackInfo.Type_}
    trackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", trackInfo.TrackState}
    trackInfo.EntityData.Leafs["state-change-counter"] = types.YLeaf{"StateChangeCounter", trackInfo.StateChangeCounter}
    trackInfo.EntityData.Leafs["seconds-last-change"] = types.YLeaf{"SecondsLastChange", trackInfo.SecondsLastChange}
    trackInfo.EntityData.Leafs["threshold-up"] = types.YLeaf{"ThresholdUp", trackInfo.ThresholdUp}
    trackInfo.EntityData.Leafs["threshold-down"] = types.YLeaf{"ThresholdDown", trackInfo.ThresholdDown}
    return &(trackInfo.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetEntityData() *types.CommonEntityData {
    trackTypeInfo.EntityData.YFilter = trackTypeInfo.YFilter
    trackTypeInfo.EntityData.YangName = "track-type-info"
    trackTypeInfo.EntityData.BundleName = "cisco_ios_xr"
    trackTypeInfo.EntityData.ParentYangName = "track-info"
    trackTypeInfo.EntityData.SegmentPath = "track-type-info"
    trackTypeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeInfo.EntityData.Children = make(map[string]types.YChild)
    trackTypeInfo.EntityData.Children["interface-tracks"] = types.YChild{"InterfaceTracks", &trackTypeInfo.InterfaceTracks}
    trackTypeInfo.EntityData.Children["route-tracks"] = types.YChild{"RouteTracks", &trackTypeInfo.RouteTracks}
    trackTypeInfo.EntityData.Children["ipsla-tracks"] = types.YChild{"IpslaTracks", &trackTypeInfo.IpslaTracks}
    trackTypeInfo.EntityData.Children["bfd-tracks"] = types.YChild{"BfdTracks", &trackTypeInfo.BfdTracks}
    trackTypeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackTypeInfo.EntityData.Leafs["discriminant"] = types.YLeaf{"Discriminant", trackTypeInfo.Discriminant}
    return &(trackTypeInfo.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetEntityData() *types.CommonEntityData {
    interfaceTracks.EntityData.YFilter = interfaceTracks.YFilter
    interfaceTracks.EntityData.YangName = "interface-tracks"
    interfaceTracks.EntityData.BundleName = "cisco_ios_xr"
    interfaceTracks.EntityData.ParentYangName = "track-type-info"
    interfaceTracks.EntityData.SegmentPath = "interface-tracks"
    interfaceTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTracks.EntityData.Children = make(map[string]types.YChild)
    interfaceTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTracks.InterfaceName}
    return &(interfaceTracks.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetEntityData() *types.CommonEntityData {
    routeTracks.EntityData.YFilter = routeTracks.YFilter
    routeTracks.EntityData.YangName = "route-tracks"
    routeTracks.EntityData.BundleName = "cisco_ios_xr"
    routeTracks.EntityData.ParentYangName = "track-type-info"
    routeTracks.EntityData.SegmentPath = "route-tracks"
    routeTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTracks.EntityData.Children = make(map[string]types.YChild)
    routeTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTracks.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", routeTracks.Prefix}
    routeTracks.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", routeTracks.PrefixLength}
    routeTracks.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", routeTracks.Vrf}
    routeTracks.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", routeTracks.NextHop}
    return &(routeTracks.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetEntityData() *types.CommonEntityData {
    ipslaTracks.EntityData.YFilter = ipslaTracks.YFilter
    ipslaTracks.EntityData.YangName = "ipsla-tracks"
    ipslaTracks.EntityData.BundleName = "cisco_ios_xr"
    ipslaTracks.EntityData.ParentYangName = "track-type-info"
    ipslaTracks.EntityData.SegmentPath = "ipsla-tracks"
    ipslaTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipslaTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipslaTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipslaTracks.EntityData.Children = make(map[string]types.YChild)
    ipslaTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    ipslaTracks.EntityData.Leafs["ipsla-op-id"] = types.YLeaf{"IpslaOpId", ipslaTracks.IpslaOpId}
    ipslaTracks.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", ipslaTracks.Rtt}
    ipslaTracks.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", ipslaTracks.ReturnCode}
    return &(ipslaTracks.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetEntityData() *types.CommonEntityData {
    bfdTracks.EntityData.YFilter = bfdTracks.YFilter
    bfdTracks.EntityData.YangName = "bfd-tracks"
    bfdTracks.EntityData.BundleName = "cisco_ios_xr"
    bfdTracks.EntityData.ParentYangName = "track-type-info"
    bfdTracks.EntityData.SegmentPath = "bfd-tracks"
    bfdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdTracks.EntityData.Children = make(map[string]types.YChild)
    bfdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfdTracks.InterfaceName}
    bfdTracks.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", bfdTracks.DestinationAddress}
    bfdTracks.EntityData.Leafs["rate"] = types.YLeaf{"Rate", bfdTracks.Rate}
    bfdTracks.EntityData.Leafs["debounce-count"] = types.YLeaf{"DebounceCount", bfdTracks.DebounceCount}
    return &(bfdTracks.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks
// boolean objects
type ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bool track info. The type is slice of
    // ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo.
    BoolTrackInfo []ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo
}

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetEntityData() *types.CommonEntityData {
    boolTracks.EntityData.YFilter = boolTracks.YFilter
    boolTracks.EntityData.YangName = "bool-tracks"
    boolTracks.EntityData.BundleName = "cisco_ios_xr"
    boolTracks.EntityData.ParentYangName = "track-info"
    boolTracks.EntityData.SegmentPath = "bool-tracks"
    boolTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    boolTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    boolTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    boolTracks.EntityData.Children = make(map[string]types.YChild)
    boolTracks.EntityData.Children["bool-track-info"] = types.YChild{"BoolTrackInfo", nil}
    for i := range boolTracks.BoolTrackInfo {
        boolTracks.EntityData.Children[types.GetSegmentPath(&boolTracks.BoolTrackInfo[i])] = types.YChild{"BoolTrackInfo", &boolTracks.BoolTrackInfo[i]}
    }
    boolTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(boolTracks.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo
// bool track info
type ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object Name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track object with Not. The type is bool.
    WithNot interface{}
}

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetEntityData() *types.CommonEntityData {
    boolTrackInfo.EntityData.YFilter = boolTrackInfo.YFilter
    boolTrackInfo.EntityData.YangName = "bool-track-info"
    boolTrackInfo.EntityData.BundleName = "cisco_ios_xr"
    boolTrackInfo.EntityData.ParentYangName = "bool-tracks"
    boolTrackInfo.EntityData.SegmentPath = "bool-track-info"
    boolTrackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    boolTrackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    boolTrackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    boolTrackInfo.EntityData.Children = make(map[string]types.YChild)
    boolTrackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    boolTrackInfo.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", boolTrackInfo.ObjectName}
    boolTrackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", boolTrackInfo.TrackState}
    boolTrackInfo.EntityData.Leafs["with-not"] = types.YLeaf{"WithNot", boolTrackInfo.WithNot}
    return &(boolTrackInfo.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks
// Threshold objects
type ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // threshold track info. The type is slice of
    // ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo.
    ThresholdTrackInfo []ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo
}

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetEntityData() *types.CommonEntityData {
    thresholdTracks.EntityData.YFilter = thresholdTracks.YFilter
    thresholdTracks.EntityData.YangName = "threshold-tracks"
    thresholdTracks.EntityData.BundleName = "cisco_ios_xr"
    thresholdTracks.EntityData.ParentYangName = "track-info"
    thresholdTracks.EntityData.SegmentPath = "threshold-tracks"
    thresholdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdTracks.EntityData.Children = make(map[string]types.YChild)
    thresholdTracks.EntityData.Children["threshold-track-info"] = types.YChild{"ThresholdTrackInfo", nil}
    for i := range thresholdTracks.ThresholdTrackInfo {
        thresholdTracks.EntityData.Children[types.GetSegmentPath(&thresholdTracks.ThresholdTrackInfo[i])] = types.YChild{"ThresholdTrackInfo", &thresholdTracks.ThresholdTrackInfo[i]}
    }
    thresholdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdTracks.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo
// threshold track info
type ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. True means track is up; False means track is down. The type is
    // bool.
    TrackState interface{}

    // Weight is the number assigned to a track object . In case of a type
    // threshold weight( i.e. weighted sum list), weight is asigned by User at the
    // time of configuration. In case of a type threshold percentage (i.e.
    // percentage based list), weight is internally computed by (1/N)x100, where N
    // is the number of objects in the list. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Weight interface{}
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetEntityData() *types.CommonEntityData {
    thresholdTrackInfo.EntityData.YFilter = thresholdTrackInfo.YFilter
    thresholdTrackInfo.EntityData.YangName = "threshold-track-info"
    thresholdTrackInfo.EntityData.BundleName = "cisco_ios_xr"
    thresholdTrackInfo.EntityData.ParentYangName = "threshold-tracks"
    thresholdTrackInfo.EntityData.SegmentPath = "threshold-track-info"
    thresholdTrackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdTrackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdTrackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdTrackInfo.EntityData.Children = make(map[string]types.YChild)
    thresholdTrackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdTrackInfo.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", thresholdTrackInfo.ObjectName}
    thresholdTrackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", thresholdTrackInfo.TrackState}
    thresholdTrackInfo.EntityData.Leafs["weight"] = types.YLeaf{"Weight", thresholdTrackInfo.Weight}
    return &(thresholdTrackInfo.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces
// Tracking Interfaces
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface tracking info. The type is slice of
    // ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo.
    InterfaceTrackingInfo []ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
}

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetEntityData() *types.CommonEntityData {
    trackingInteraces.EntityData.YFilter = trackingInteraces.YFilter
    trackingInteraces.EntityData.YangName = "tracking-interaces"
    trackingInteraces.EntityData.BundleName = "cisco_ios_xr"
    trackingInteraces.EntityData.ParentYangName = "track-info"
    trackingInteraces.EntityData.SegmentPath = "tracking-interaces"
    trackingInteraces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackingInteraces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackingInteraces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackingInteraces.EntityData.Children = make(map[string]types.YChild)
    trackingInteraces.EntityData.Children["interface-tracking-info"] = types.YChild{"InterfaceTrackingInfo", nil}
    for i := range trackingInteraces.InterfaceTrackingInfo {
        trackingInteraces.EntityData.Children[types.GetSegmentPath(&trackingInteraces.InterfaceTrackingInfo[i])] = types.YChild{"InterfaceTrackingInfo", &trackingInteraces.InterfaceTrackingInfo[i]}
    }
    trackingInteraces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackingInteraces.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
// interface tracking info
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetEntityData() *types.CommonEntityData {
    interfaceTrackingInfo.EntityData.YFilter = interfaceTrackingInfo.YFilter
    interfaceTrackingInfo.EntityData.YangName = "interface-tracking-info"
    interfaceTrackingInfo.EntityData.BundleName = "cisco_ios_xr"
    interfaceTrackingInfo.EntityData.ParentYangName = "tracking-interaces"
    interfaceTrackingInfo.EntityData.SegmentPath = "interface-tracking-info"
    interfaceTrackingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTrackingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTrackingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTrackingInfo.EntityData.Children = make(map[string]types.YChild)
    interfaceTrackingInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTrackingInfo.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTrackingInfo.InterfaceName}
    return &(interfaceTrackingInfo.EntityData)
}

// ObjectTracking_TrackTypeInterface_TrackInfo_Delayed
// Is the state change delay counter in progress
type ObjectTracking_TrackTypeInterface_TrackInfo_Delayed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time remaining in seconds for the counter to trigger state change. The
    // type is interface{} with range: 0..4294967295. Units are second.
    TimeRemaining interface{}

    // State the track will transition to. Track state. True means track is up;
    // False means track is down. The type is bool.
    TrackState interface{}
}

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetEntityData() *types.CommonEntityData {
    delayed.EntityData.YFilter = delayed.YFilter
    delayed.EntityData.YangName = "delayed"
    delayed.EntityData.BundleName = "cisco_ios_xr"
    delayed.EntityData.ParentYangName = "track-info"
    delayed.EntityData.SegmentPath = "delayed"
    delayed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayed.EntityData.Children = make(map[string]types.YChild)
    delayed.EntityData.Leafs = make(map[string]types.YLeaf)
    delayed.EntityData.Leafs["time-remaining"] = types.YLeaf{"TimeRemaining", delayed.TimeRemaining}
    delayed.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", delayed.TrackState}
    return &(delayed.EntityData)
}

// ObjectTracking_TrackBriefs
// Object Tracking Track table brief
type ObjectTracking_TrackBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTracking_TrackBriefs_TrackBrief.
    TrackBrief []ObjectTracking_TrackBriefs_TrackBrief
}

func (trackBriefs *ObjectTracking_TrackBriefs) GetEntityData() *types.CommonEntityData {
    trackBriefs.EntityData.YFilter = trackBriefs.YFilter
    trackBriefs.EntityData.YangName = "track-briefs"
    trackBriefs.EntityData.BundleName = "cisco_ios_xr"
    trackBriefs.EntityData.ParentYangName = "object-tracking"
    trackBriefs.EntityData.SegmentPath = "track-briefs"
    trackBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackBriefs.EntityData.Children = make(map[string]types.YChild)
    trackBriefs.EntityData.Children["track-brief"] = types.YChild{"TrackBrief", nil}
    for i := range trackBriefs.TrackBrief {
        trackBriefs.EntityData.Children[types.GetSegmentPath(&trackBriefs.TrackBrief[i])] = types.YChild{"TrackBrief", &trackBriefs.TrackBrief[i]}
    }
    trackBriefs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackBriefs.EntityData)
}

// ObjectTracking_TrackBriefs_TrackBrief
// Track name - maximum 32 characters
type ObjectTracking_TrackBriefs_TrackBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Track name. The type is string with length: 1..32.
    TrackName interface{}

    // track info brief. The type is slice of
    // ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief.
    TrackInfoBrief []ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief
}

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetEntityData() *types.CommonEntityData {
    trackBrief.EntityData.YFilter = trackBrief.YFilter
    trackBrief.EntityData.YangName = "track-brief"
    trackBrief.EntityData.BundleName = "cisco_ios_xr"
    trackBrief.EntityData.ParentYangName = "track-briefs"
    trackBrief.EntityData.SegmentPath = "track-brief" + "[track-name='" + fmt.Sprintf("%v", trackBrief.TrackName) + "']"
    trackBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackBrief.EntityData.Children = make(map[string]types.YChild)
    trackBrief.EntityData.Children["track-info-brief"] = types.YChild{"TrackInfoBrief", nil}
    for i := range trackBrief.TrackInfoBrief {
        trackBrief.EntityData.Children[types.GetSegmentPath(&trackBrief.TrackInfoBrief[i])] = types.YChild{"TrackInfoBrief", &trackBrief.TrackInfoBrief[i]}
    }
    trackBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    trackBrief.EntityData.Leafs["track-name"] = types.YLeaf{"TrackName", trackBrief.TrackName}
    return &(trackBrief.EntityData)
}

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief
// track info brief
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type_ interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo
}

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetEntityData() *types.CommonEntityData {
    trackInfoBrief.EntityData.YFilter = trackInfoBrief.YFilter
    trackInfoBrief.EntityData.YangName = "track-info-brief"
    trackInfoBrief.EntityData.BundleName = "cisco_ios_xr"
    trackInfoBrief.EntityData.ParentYangName = "track-brief"
    trackInfoBrief.EntityData.SegmentPath = "track-info-brief"
    trackInfoBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackInfoBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackInfoBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackInfoBrief.EntityData.Children = make(map[string]types.YChild)
    trackInfoBrief.EntityData.Children["track-type-info"] = types.YChild{"TrackTypeInfo", &trackInfoBrief.TrackTypeInfo}
    trackInfoBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    trackInfoBrief.EntityData.Leafs["tracke-name"] = types.YLeaf{"TrackeName", trackInfoBrief.TrackeName}
    trackInfoBrief.EntityData.Leafs["type"] = types.YLeaf{"Type_", trackInfoBrief.Type_}
    trackInfoBrief.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", trackInfoBrief.TrackState}
    return &(trackInfoBrief.EntityData)
}

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo
// Track type information
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetEntityData() *types.CommonEntityData {
    trackTypeInfo.EntityData.YFilter = trackTypeInfo.YFilter
    trackTypeInfo.EntityData.YangName = "track-type-info"
    trackTypeInfo.EntityData.BundleName = "cisco_ios_xr"
    trackTypeInfo.EntityData.ParentYangName = "track-info-brief"
    trackTypeInfo.EntityData.SegmentPath = "track-type-info"
    trackTypeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeInfo.EntityData.Children = make(map[string]types.YChild)
    trackTypeInfo.EntityData.Children["interface-tracks"] = types.YChild{"InterfaceTracks", &trackTypeInfo.InterfaceTracks}
    trackTypeInfo.EntityData.Children["route-tracks"] = types.YChild{"RouteTracks", &trackTypeInfo.RouteTracks}
    trackTypeInfo.EntityData.Children["ipsla-tracks"] = types.YChild{"IpslaTracks", &trackTypeInfo.IpslaTracks}
    trackTypeInfo.EntityData.Children["bfd-tracks"] = types.YChild{"BfdTracks", &trackTypeInfo.BfdTracks}
    trackTypeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackTypeInfo.EntityData.Leafs["discriminant"] = types.YLeaf{"Discriminant", trackTypeInfo.Discriminant}
    return &(trackTypeInfo.EntityData)
}

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetEntityData() *types.CommonEntityData {
    interfaceTracks.EntityData.YFilter = interfaceTracks.YFilter
    interfaceTracks.EntityData.YangName = "interface-tracks"
    interfaceTracks.EntityData.BundleName = "cisco_ios_xr"
    interfaceTracks.EntityData.ParentYangName = "track-type-info"
    interfaceTracks.EntityData.SegmentPath = "interface-tracks"
    interfaceTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTracks.EntityData.Children = make(map[string]types.YChild)
    interfaceTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTracks.InterfaceName}
    return &(interfaceTracks.EntityData)
}

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetEntityData() *types.CommonEntityData {
    routeTracks.EntityData.YFilter = routeTracks.YFilter
    routeTracks.EntityData.YangName = "route-tracks"
    routeTracks.EntityData.BundleName = "cisco_ios_xr"
    routeTracks.EntityData.ParentYangName = "track-type-info"
    routeTracks.EntityData.SegmentPath = "route-tracks"
    routeTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTracks.EntityData.Children = make(map[string]types.YChild)
    routeTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTracks.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", routeTracks.Prefix}
    routeTracks.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", routeTracks.PrefixLength}
    routeTracks.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", routeTracks.Vrf}
    routeTracks.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", routeTracks.NextHop}
    return &(routeTracks.EntityData)
}

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetEntityData() *types.CommonEntityData {
    ipslaTracks.EntityData.YFilter = ipslaTracks.YFilter
    ipslaTracks.EntityData.YangName = "ipsla-tracks"
    ipslaTracks.EntityData.BundleName = "cisco_ios_xr"
    ipslaTracks.EntityData.ParentYangName = "track-type-info"
    ipslaTracks.EntityData.SegmentPath = "ipsla-tracks"
    ipslaTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipslaTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipslaTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipslaTracks.EntityData.Children = make(map[string]types.YChild)
    ipslaTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    ipslaTracks.EntityData.Leafs["ipsla-op-id"] = types.YLeaf{"IpslaOpId", ipslaTracks.IpslaOpId}
    ipslaTracks.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", ipslaTracks.Rtt}
    ipslaTracks.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", ipslaTracks.ReturnCode}
    return &(ipslaTracks.EntityData)
}

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetEntityData() *types.CommonEntityData {
    bfdTracks.EntityData.YFilter = bfdTracks.YFilter
    bfdTracks.EntityData.YangName = "bfd-tracks"
    bfdTracks.EntityData.BundleName = "cisco_ios_xr"
    bfdTracks.EntityData.ParentYangName = "track-type-info"
    bfdTracks.EntityData.SegmentPath = "bfd-tracks"
    bfdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdTracks.EntityData.Children = make(map[string]types.YChild)
    bfdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfdTracks.InterfaceName}
    bfdTracks.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", bfdTracks.DestinationAddress}
    bfdTracks.EntityData.Leafs["rate"] = types.YLeaf{"Rate", bfdTracks.Rate}
    bfdTracks.EntityData.Leafs["debounce-count"] = types.YLeaf{"DebounceCount", bfdTracks.DebounceCount}
    return &(bfdTracks.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability
// Object Tracking Type RTR Reachability info
type ObjectTracking_TrackTypeRtrReachability struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // track info. The type is slice of
    // ObjectTracking_TrackTypeRtrReachability_TrackInfo.
    TrackInfo []ObjectTracking_TrackTypeRtrReachability_TrackInfo
}

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetEntityData() *types.CommonEntityData {
    trackTypeRtrReachability.EntityData.YFilter = trackTypeRtrReachability.YFilter
    trackTypeRtrReachability.EntityData.YangName = "track-type-rtr-reachability"
    trackTypeRtrReachability.EntityData.BundleName = "cisco_ios_xr"
    trackTypeRtrReachability.EntityData.ParentYangName = "object-tracking"
    trackTypeRtrReachability.EntityData.SegmentPath = "track-type-rtr-reachability"
    trackTypeRtrReachability.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeRtrReachability.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeRtrReachability.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeRtrReachability.EntityData.Children = make(map[string]types.YChild)
    trackTypeRtrReachability.EntityData.Children["track-info"] = types.YChild{"TrackInfo", nil}
    for i := range trackTypeRtrReachability.TrackInfo {
        trackTypeRtrReachability.EntityData.Children[types.GetSegmentPath(&trackTypeRtrReachability.TrackInfo[i])] = types.YChild{"TrackInfo", &trackTypeRtrReachability.TrackInfo[i]}
    }
    trackTypeRtrReachability.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackTypeRtrReachability.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo
// track info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type_ interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // State Change Counter. The type is interface{} with range: 0..4294967295.
    StateChangeCounter interface{}

    // Seconds Last Change. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecondsLastChange interface{}

    // User specified threshold upper limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdUp interface{}

    // User specified threshold lower limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdDown interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo

    // boolean objects.
    BoolTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks

    // Threshold objects.
    ThresholdTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks

    // Tracking Interfaces.
    TrackingInteraces ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces

    // Is the state change delay counter in progress.
    Delayed ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed
}

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetEntityData() *types.CommonEntityData {
    trackInfo.EntityData.YFilter = trackInfo.YFilter
    trackInfo.EntityData.YangName = "track-info"
    trackInfo.EntityData.BundleName = "cisco_ios_xr"
    trackInfo.EntityData.ParentYangName = "track-type-rtr-reachability"
    trackInfo.EntityData.SegmentPath = "track-info"
    trackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackInfo.EntityData.Children = make(map[string]types.YChild)
    trackInfo.EntityData.Children["track-type-info"] = types.YChild{"TrackTypeInfo", &trackInfo.TrackTypeInfo}
    trackInfo.EntityData.Children["bool-tracks"] = types.YChild{"BoolTracks", &trackInfo.BoolTracks}
    trackInfo.EntityData.Children["threshold-tracks"] = types.YChild{"ThresholdTracks", &trackInfo.ThresholdTracks}
    trackInfo.EntityData.Children["tracking-interaces"] = types.YChild{"TrackingInteraces", &trackInfo.TrackingInteraces}
    trackInfo.EntityData.Children["delayed"] = types.YChild{"Delayed", &trackInfo.Delayed}
    trackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackInfo.EntityData.Leafs["tracke-name"] = types.YLeaf{"TrackeName", trackInfo.TrackeName}
    trackInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", trackInfo.Type_}
    trackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", trackInfo.TrackState}
    trackInfo.EntityData.Leafs["state-change-counter"] = types.YLeaf{"StateChangeCounter", trackInfo.StateChangeCounter}
    trackInfo.EntityData.Leafs["seconds-last-change"] = types.YLeaf{"SecondsLastChange", trackInfo.SecondsLastChange}
    trackInfo.EntityData.Leafs["threshold-up"] = types.YLeaf{"ThresholdUp", trackInfo.ThresholdUp}
    trackInfo.EntityData.Leafs["threshold-down"] = types.YLeaf{"ThresholdDown", trackInfo.ThresholdDown}
    return &(trackInfo.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetEntityData() *types.CommonEntityData {
    trackTypeInfo.EntityData.YFilter = trackTypeInfo.YFilter
    trackTypeInfo.EntityData.YangName = "track-type-info"
    trackTypeInfo.EntityData.BundleName = "cisco_ios_xr"
    trackTypeInfo.EntityData.ParentYangName = "track-info"
    trackTypeInfo.EntityData.SegmentPath = "track-type-info"
    trackTypeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeInfo.EntityData.Children = make(map[string]types.YChild)
    trackTypeInfo.EntityData.Children["interface-tracks"] = types.YChild{"InterfaceTracks", &trackTypeInfo.InterfaceTracks}
    trackTypeInfo.EntityData.Children["route-tracks"] = types.YChild{"RouteTracks", &trackTypeInfo.RouteTracks}
    trackTypeInfo.EntityData.Children["ipsla-tracks"] = types.YChild{"IpslaTracks", &trackTypeInfo.IpslaTracks}
    trackTypeInfo.EntityData.Children["bfd-tracks"] = types.YChild{"BfdTracks", &trackTypeInfo.BfdTracks}
    trackTypeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackTypeInfo.EntityData.Leafs["discriminant"] = types.YLeaf{"Discriminant", trackTypeInfo.Discriminant}
    return &(trackTypeInfo.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetEntityData() *types.CommonEntityData {
    interfaceTracks.EntityData.YFilter = interfaceTracks.YFilter
    interfaceTracks.EntityData.YangName = "interface-tracks"
    interfaceTracks.EntityData.BundleName = "cisco_ios_xr"
    interfaceTracks.EntityData.ParentYangName = "track-type-info"
    interfaceTracks.EntityData.SegmentPath = "interface-tracks"
    interfaceTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTracks.EntityData.Children = make(map[string]types.YChild)
    interfaceTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTracks.InterfaceName}
    return &(interfaceTracks.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetEntityData() *types.CommonEntityData {
    routeTracks.EntityData.YFilter = routeTracks.YFilter
    routeTracks.EntityData.YangName = "route-tracks"
    routeTracks.EntityData.BundleName = "cisco_ios_xr"
    routeTracks.EntityData.ParentYangName = "track-type-info"
    routeTracks.EntityData.SegmentPath = "route-tracks"
    routeTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTracks.EntityData.Children = make(map[string]types.YChild)
    routeTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTracks.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", routeTracks.Prefix}
    routeTracks.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", routeTracks.PrefixLength}
    routeTracks.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", routeTracks.Vrf}
    routeTracks.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", routeTracks.NextHop}
    return &(routeTracks.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetEntityData() *types.CommonEntityData {
    ipslaTracks.EntityData.YFilter = ipslaTracks.YFilter
    ipslaTracks.EntityData.YangName = "ipsla-tracks"
    ipslaTracks.EntityData.BundleName = "cisco_ios_xr"
    ipslaTracks.EntityData.ParentYangName = "track-type-info"
    ipslaTracks.EntityData.SegmentPath = "ipsla-tracks"
    ipslaTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipslaTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipslaTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipslaTracks.EntityData.Children = make(map[string]types.YChild)
    ipslaTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    ipslaTracks.EntityData.Leafs["ipsla-op-id"] = types.YLeaf{"IpslaOpId", ipslaTracks.IpslaOpId}
    ipslaTracks.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", ipslaTracks.Rtt}
    ipslaTracks.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", ipslaTracks.ReturnCode}
    return &(ipslaTracks.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetEntityData() *types.CommonEntityData {
    bfdTracks.EntityData.YFilter = bfdTracks.YFilter
    bfdTracks.EntityData.YangName = "bfd-tracks"
    bfdTracks.EntityData.BundleName = "cisco_ios_xr"
    bfdTracks.EntityData.ParentYangName = "track-type-info"
    bfdTracks.EntityData.SegmentPath = "bfd-tracks"
    bfdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdTracks.EntityData.Children = make(map[string]types.YChild)
    bfdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfdTracks.InterfaceName}
    bfdTracks.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", bfdTracks.DestinationAddress}
    bfdTracks.EntityData.Leafs["rate"] = types.YLeaf{"Rate", bfdTracks.Rate}
    bfdTracks.EntityData.Leafs["debounce-count"] = types.YLeaf{"DebounceCount", bfdTracks.DebounceCount}
    return &(bfdTracks.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks
// boolean objects
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bool track info. The type is slice of
    // ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo.
    BoolTrackInfo []ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo
}

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetEntityData() *types.CommonEntityData {
    boolTracks.EntityData.YFilter = boolTracks.YFilter
    boolTracks.EntityData.YangName = "bool-tracks"
    boolTracks.EntityData.BundleName = "cisco_ios_xr"
    boolTracks.EntityData.ParentYangName = "track-info"
    boolTracks.EntityData.SegmentPath = "bool-tracks"
    boolTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    boolTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    boolTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    boolTracks.EntityData.Children = make(map[string]types.YChild)
    boolTracks.EntityData.Children["bool-track-info"] = types.YChild{"BoolTrackInfo", nil}
    for i := range boolTracks.BoolTrackInfo {
        boolTracks.EntityData.Children[types.GetSegmentPath(&boolTracks.BoolTrackInfo[i])] = types.YChild{"BoolTrackInfo", &boolTracks.BoolTrackInfo[i]}
    }
    boolTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(boolTracks.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo
// bool track info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object Name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track object with Not. The type is bool.
    WithNot interface{}
}

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetEntityData() *types.CommonEntityData {
    boolTrackInfo.EntityData.YFilter = boolTrackInfo.YFilter
    boolTrackInfo.EntityData.YangName = "bool-track-info"
    boolTrackInfo.EntityData.BundleName = "cisco_ios_xr"
    boolTrackInfo.EntityData.ParentYangName = "bool-tracks"
    boolTrackInfo.EntityData.SegmentPath = "bool-track-info"
    boolTrackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    boolTrackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    boolTrackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    boolTrackInfo.EntityData.Children = make(map[string]types.YChild)
    boolTrackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    boolTrackInfo.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", boolTrackInfo.ObjectName}
    boolTrackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", boolTrackInfo.TrackState}
    boolTrackInfo.EntityData.Leafs["with-not"] = types.YLeaf{"WithNot", boolTrackInfo.WithNot}
    return &(boolTrackInfo.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks
// Threshold objects
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // threshold track info. The type is slice of
    // ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo.
    ThresholdTrackInfo []ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo
}

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetEntityData() *types.CommonEntityData {
    thresholdTracks.EntityData.YFilter = thresholdTracks.YFilter
    thresholdTracks.EntityData.YangName = "threshold-tracks"
    thresholdTracks.EntityData.BundleName = "cisco_ios_xr"
    thresholdTracks.EntityData.ParentYangName = "track-info"
    thresholdTracks.EntityData.SegmentPath = "threshold-tracks"
    thresholdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdTracks.EntityData.Children = make(map[string]types.YChild)
    thresholdTracks.EntityData.Children["threshold-track-info"] = types.YChild{"ThresholdTrackInfo", nil}
    for i := range thresholdTracks.ThresholdTrackInfo {
        thresholdTracks.EntityData.Children[types.GetSegmentPath(&thresholdTracks.ThresholdTrackInfo[i])] = types.YChild{"ThresholdTrackInfo", &thresholdTracks.ThresholdTrackInfo[i]}
    }
    thresholdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdTracks.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo
// threshold track info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. True means track is up; False means track is down. The type is
    // bool.
    TrackState interface{}

    // Weight is the number assigned to a track object . In case of a type
    // threshold weight( i.e. weighted sum list), weight is asigned by User at the
    // time of configuration. In case of a type threshold percentage (i.e.
    // percentage based list), weight is internally computed by (1/N)x100, where N
    // is the number of objects in the list. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Weight interface{}
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetEntityData() *types.CommonEntityData {
    thresholdTrackInfo.EntityData.YFilter = thresholdTrackInfo.YFilter
    thresholdTrackInfo.EntityData.YangName = "threshold-track-info"
    thresholdTrackInfo.EntityData.BundleName = "cisco_ios_xr"
    thresholdTrackInfo.EntityData.ParentYangName = "threshold-tracks"
    thresholdTrackInfo.EntityData.SegmentPath = "threshold-track-info"
    thresholdTrackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdTrackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdTrackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdTrackInfo.EntityData.Children = make(map[string]types.YChild)
    thresholdTrackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdTrackInfo.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", thresholdTrackInfo.ObjectName}
    thresholdTrackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", thresholdTrackInfo.TrackState}
    thresholdTrackInfo.EntityData.Leafs["weight"] = types.YLeaf{"Weight", thresholdTrackInfo.Weight}
    return &(thresholdTrackInfo.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces
// Tracking Interfaces
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface tracking info. The type is slice of
    // ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo.
    InterfaceTrackingInfo []ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
}

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetEntityData() *types.CommonEntityData {
    trackingInteraces.EntityData.YFilter = trackingInteraces.YFilter
    trackingInteraces.EntityData.YangName = "tracking-interaces"
    trackingInteraces.EntityData.BundleName = "cisco_ios_xr"
    trackingInteraces.EntityData.ParentYangName = "track-info"
    trackingInteraces.EntityData.SegmentPath = "tracking-interaces"
    trackingInteraces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackingInteraces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackingInteraces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackingInteraces.EntityData.Children = make(map[string]types.YChild)
    trackingInteraces.EntityData.Children["interface-tracking-info"] = types.YChild{"InterfaceTrackingInfo", nil}
    for i := range trackingInteraces.InterfaceTrackingInfo {
        trackingInteraces.EntityData.Children[types.GetSegmentPath(&trackingInteraces.InterfaceTrackingInfo[i])] = types.YChild{"InterfaceTrackingInfo", &trackingInteraces.InterfaceTrackingInfo[i]}
    }
    trackingInteraces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackingInteraces.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
// interface tracking info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetEntityData() *types.CommonEntityData {
    interfaceTrackingInfo.EntityData.YFilter = interfaceTrackingInfo.YFilter
    interfaceTrackingInfo.EntityData.YangName = "interface-tracking-info"
    interfaceTrackingInfo.EntityData.BundleName = "cisco_ios_xr"
    interfaceTrackingInfo.EntityData.ParentYangName = "tracking-interaces"
    interfaceTrackingInfo.EntityData.SegmentPath = "interface-tracking-info"
    interfaceTrackingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTrackingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTrackingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTrackingInfo.EntityData.Children = make(map[string]types.YChild)
    interfaceTrackingInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTrackingInfo.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTrackingInfo.InterfaceName}
    return &(interfaceTrackingInfo.EntityData)
}

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed
// Is the state change delay counter in progress
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time remaining in seconds for the counter to trigger state change. The
    // type is interface{} with range: 0..4294967295. Units are second.
    TimeRemaining interface{}

    // State the track will transition to. Track state. True means track is up;
    // False means track is down. The type is bool.
    TrackState interface{}
}

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetEntityData() *types.CommonEntityData {
    delayed.EntityData.YFilter = delayed.YFilter
    delayed.EntityData.YangName = "delayed"
    delayed.EntityData.BundleName = "cisco_ios_xr"
    delayed.EntityData.ParentYangName = "track-info"
    delayed.EntityData.SegmentPath = "delayed"
    delayed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayed.EntityData.Children = make(map[string]types.YChild)
    delayed.EntityData.Leafs = make(map[string]types.YLeaf)
    delayed.EntityData.Leafs["time-remaining"] = types.YLeaf{"TimeRemaining", delayed.TimeRemaining}
    delayed.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", delayed.TrackState}
    return &(delayed.EntityData)
}

// ObjectTracking_TrackTypeRtrReachabilityBrief
// Object Tracking Type RTR Reachability brief info
type ObjectTracking_TrackTypeRtrReachabilityBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // track info brief. The type is slice of
    // ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief.
    TrackInfoBrief []ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief
}

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetEntityData() *types.CommonEntityData {
    trackTypeRtrReachabilityBrief.EntityData.YFilter = trackTypeRtrReachabilityBrief.YFilter
    trackTypeRtrReachabilityBrief.EntityData.YangName = "track-type-rtr-reachability-brief"
    trackTypeRtrReachabilityBrief.EntityData.BundleName = "cisco_ios_xr"
    trackTypeRtrReachabilityBrief.EntityData.ParentYangName = "object-tracking"
    trackTypeRtrReachabilityBrief.EntityData.SegmentPath = "track-type-rtr-reachability-brief"
    trackTypeRtrReachabilityBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeRtrReachabilityBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeRtrReachabilityBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeRtrReachabilityBrief.EntityData.Children = make(map[string]types.YChild)
    trackTypeRtrReachabilityBrief.EntityData.Children["track-info-brief"] = types.YChild{"TrackInfoBrief", nil}
    for i := range trackTypeRtrReachabilityBrief.TrackInfoBrief {
        trackTypeRtrReachabilityBrief.EntityData.Children[types.GetSegmentPath(&trackTypeRtrReachabilityBrief.TrackInfoBrief[i])] = types.YChild{"TrackInfoBrief", &trackTypeRtrReachabilityBrief.TrackInfoBrief[i]}
    }
    trackTypeRtrReachabilityBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackTypeRtrReachabilityBrief.EntityData)
}

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief
// track info brief
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type_ interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo
}

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetEntityData() *types.CommonEntityData {
    trackInfoBrief.EntityData.YFilter = trackInfoBrief.YFilter
    trackInfoBrief.EntityData.YangName = "track-info-brief"
    trackInfoBrief.EntityData.BundleName = "cisco_ios_xr"
    trackInfoBrief.EntityData.ParentYangName = "track-type-rtr-reachability-brief"
    trackInfoBrief.EntityData.SegmentPath = "track-info-brief"
    trackInfoBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackInfoBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackInfoBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackInfoBrief.EntityData.Children = make(map[string]types.YChild)
    trackInfoBrief.EntityData.Children["track-type-info"] = types.YChild{"TrackTypeInfo", &trackInfoBrief.TrackTypeInfo}
    trackInfoBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    trackInfoBrief.EntityData.Leafs["tracke-name"] = types.YLeaf{"TrackeName", trackInfoBrief.TrackeName}
    trackInfoBrief.EntityData.Leafs["type"] = types.YLeaf{"Type_", trackInfoBrief.Type_}
    trackInfoBrief.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", trackInfoBrief.TrackState}
    return &(trackInfoBrief.EntityData)
}

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetEntityData() *types.CommonEntityData {
    trackTypeInfo.EntityData.YFilter = trackTypeInfo.YFilter
    trackTypeInfo.EntityData.YangName = "track-type-info"
    trackTypeInfo.EntityData.BundleName = "cisco_ios_xr"
    trackTypeInfo.EntityData.ParentYangName = "track-info-brief"
    trackTypeInfo.EntityData.SegmentPath = "track-type-info"
    trackTypeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeInfo.EntityData.Children = make(map[string]types.YChild)
    trackTypeInfo.EntityData.Children["interface-tracks"] = types.YChild{"InterfaceTracks", &trackTypeInfo.InterfaceTracks}
    trackTypeInfo.EntityData.Children["route-tracks"] = types.YChild{"RouteTracks", &trackTypeInfo.RouteTracks}
    trackTypeInfo.EntityData.Children["ipsla-tracks"] = types.YChild{"IpslaTracks", &trackTypeInfo.IpslaTracks}
    trackTypeInfo.EntityData.Children["bfd-tracks"] = types.YChild{"BfdTracks", &trackTypeInfo.BfdTracks}
    trackTypeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackTypeInfo.EntityData.Leafs["discriminant"] = types.YLeaf{"Discriminant", trackTypeInfo.Discriminant}
    return &(trackTypeInfo.EntityData)
}

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetEntityData() *types.CommonEntityData {
    interfaceTracks.EntityData.YFilter = interfaceTracks.YFilter
    interfaceTracks.EntityData.YangName = "interface-tracks"
    interfaceTracks.EntityData.BundleName = "cisco_ios_xr"
    interfaceTracks.EntityData.ParentYangName = "track-type-info"
    interfaceTracks.EntityData.SegmentPath = "interface-tracks"
    interfaceTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTracks.EntityData.Children = make(map[string]types.YChild)
    interfaceTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTracks.InterfaceName}
    return &(interfaceTracks.EntityData)
}

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetEntityData() *types.CommonEntityData {
    routeTracks.EntityData.YFilter = routeTracks.YFilter
    routeTracks.EntityData.YangName = "route-tracks"
    routeTracks.EntityData.BundleName = "cisco_ios_xr"
    routeTracks.EntityData.ParentYangName = "track-type-info"
    routeTracks.EntityData.SegmentPath = "route-tracks"
    routeTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTracks.EntityData.Children = make(map[string]types.YChild)
    routeTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTracks.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", routeTracks.Prefix}
    routeTracks.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", routeTracks.PrefixLength}
    routeTracks.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", routeTracks.Vrf}
    routeTracks.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", routeTracks.NextHop}
    return &(routeTracks.EntityData)
}

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetEntityData() *types.CommonEntityData {
    ipslaTracks.EntityData.YFilter = ipslaTracks.YFilter
    ipslaTracks.EntityData.YangName = "ipsla-tracks"
    ipslaTracks.EntityData.BundleName = "cisco_ios_xr"
    ipslaTracks.EntityData.ParentYangName = "track-type-info"
    ipslaTracks.EntityData.SegmentPath = "ipsla-tracks"
    ipslaTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipslaTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipslaTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipslaTracks.EntityData.Children = make(map[string]types.YChild)
    ipslaTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    ipslaTracks.EntityData.Leafs["ipsla-op-id"] = types.YLeaf{"IpslaOpId", ipslaTracks.IpslaOpId}
    ipslaTracks.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", ipslaTracks.Rtt}
    ipslaTracks.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", ipslaTracks.ReturnCode}
    return &(ipslaTracks.EntityData)
}

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetEntityData() *types.CommonEntityData {
    bfdTracks.EntityData.YFilter = bfdTracks.YFilter
    bfdTracks.EntityData.YangName = "bfd-tracks"
    bfdTracks.EntityData.BundleName = "cisco_ios_xr"
    bfdTracks.EntityData.ParentYangName = "track-type-info"
    bfdTracks.EntityData.SegmentPath = "bfd-tracks"
    bfdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdTracks.EntityData.Children = make(map[string]types.YChild)
    bfdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfdTracks.InterfaceName}
    bfdTracks.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", bfdTracks.DestinationAddress}
    bfdTracks.EntityData.Leafs["rate"] = types.YLeaf{"Rate", bfdTracks.Rate}
    bfdTracks.EntityData.Leafs["debounce-count"] = types.YLeaf{"DebounceCount", bfdTracks.DebounceCount}
    return &(bfdTracks.EntityData)
}

// ObjectTracking_Tracks
// Object Tracking Track table
type ObjectTracking_Tracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTracking_Tracks_Track.
    Track []ObjectTracking_Tracks_Track
}

func (tracks *ObjectTracking_Tracks) GetEntityData() *types.CommonEntityData {
    tracks.EntityData.YFilter = tracks.YFilter
    tracks.EntityData.YangName = "tracks"
    tracks.EntityData.BundleName = "cisco_ios_xr"
    tracks.EntityData.ParentYangName = "object-tracking"
    tracks.EntityData.SegmentPath = "tracks"
    tracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracks.EntityData.Children = make(map[string]types.YChild)
    tracks.EntityData.Children["track"] = types.YChild{"Track", nil}
    for i := range tracks.Track {
        tracks.EntityData.Children[types.GetSegmentPath(&tracks.Track[i])] = types.YChild{"Track", &tracks.Track[i]}
    }
    tracks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tracks.EntityData)
}

// ObjectTracking_Tracks_Track
// Track name - maximum 32 characters
type ObjectTracking_Tracks_Track struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Track name. The type is string with length: 1..32.
    TrackName interface{}

    // track info. The type is slice of ObjectTracking_Tracks_Track_TrackInfo.
    TrackInfo []ObjectTracking_Tracks_Track_TrackInfo
}

func (track *ObjectTracking_Tracks_Track) GetEntityData() *types.CommonEntityData {
    track.EntityData.YFilter = track.YFilter
    track.EntityData.YangName = "track"
    track.EntityData.BundleName = "cisco_ios_xr"
    track.EntityData.ParentYangName = "tracks"
    track.EntityData.SegmentPath = "track" + "[track-name='" + fmt.Sprintf("%v", track.TrackName) + "']"
    track.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    track.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    track.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    track.EntityData.Children = make(map[string]types.YChild)
    track.EntityData.Children["track-info"] = types.YChild{"TrackInfo", nil}
    for i := range track.TrackInfo {
        track.EntityData.Children[types.GetSegmentPath(&track.TrackInfo[i])] = types.YChild{"TrackInfo", &track.TrackInfo[i]}
    }
    track.EntityData.Leafs = make(map[string]types.YLeaf)
    track.EntityData.Leafs["track-name"] = types.YLeaf{"TrackName", track.TrackName}
    return &(track.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo
// track info
type ObjectTracking_Tracks_Track_TrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type_ interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // State Change Counter. The type is interface{} with range: 0..4294967295.
    StateChangeCounter interface{}

    // Seconds Last Change. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecondsLastChange interface{}

    // User specified threshold upper limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdUp interface{}

    // User specified threshold lower limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdDown interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo

    // boolean objects.
    BoolTracks ObjectTracking_Tracks_Track_TrackInfo_BoolTracks

    // Threshold objects.
    ThresholdTracks ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks

    // Tracking Interfaces.
    TrackingInteraces ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces

    // Is the state change delay counter in progress.
    Delayed ObjectTracking_Tracks_Track_TrackInfo_Delayed
}

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetEntityData() *types.CommonEntityData {
    trackInfo.EntityData.YFilter = trackInfo.YFilter
    trackInfo.EntityData.YangName = "track-info"
    trackInfo.EntityData.BundleName = "cisco_ios_xr"
    trackInfo.EntityData.ParentYangName = "track"
    trackInfo.EntityData.SegmentPath = "track-info"
    trackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackInfo.EntityData.Children = make(map[string]types.YChild)
    trackInfo.EntityData.Children["track-type-info"] = types.YChild{"TrackTypeInfo", &trackInfo.TrackTypeInfo}
    trackInfo.EntityData.Children["bool-tracks"] = types.YChild{"BoolTracks", &trackInfo.BoolTracks}
    trackInfo.EntityData.Children["threshold-tracks"] = types.YChild{"ThresholdTracks", &trackInfo.ThresholdTracks}
    trackInfo.EntityData.Children["tracking-interaces"] = types.YChild{"TrackingInteraces", &trackInfo.TrackingInteraces}
    trackInfo.EntityData.Children["delayed"] = types.YChild{"Delayed", &trackInfo.Delayed}
    trackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackInfo.EntityData.Leafs["tracke-name"] = types.YLeaf{"TrackeName", trackInfo.TrackeName}
    trackInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", trackInfo.Type_}
    trackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", trackInfo.TrackState}
    trackInfo.EntityData.Leafs["state-change-counter"] = types.YLeaf{"StateChangeCounter", trackInfo.StateChangeCounter}
    trackInfo.EntityData.Leafs["seconds-last-change"] = types.YLeaf{"SecondsLastChange", trackInfo.SecondsLastChange}
    trackInfo.EntityData.Leafs["threshold-up"] = types.YLeaf{"ThresholdUp", trackInfo.ThresholdUp}
    trackInfo.EntityData.Leafs["threshold-down"] = types.YLeaf{"ThresholdDown", trackInfo.ThresholdDown}
    return &(trackInfo.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo
// Track type information
type ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetEntityData() *types.CommonEntityData {
    trackTypeInfo.EntityData.YFilter = trackTypeInfo.YFilter
    trackTypeInfo.EntityData.YangName = "track-type-info"
    trackTypeInfo.EntityData.BundleName = "cisco_ios_xr"
    trackTypeInfo.EntityData.ParentYangName = "track-info"
    trackTypeInfo.EntityData.SegmentPath = "track-type-info"
    trackTypeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeInfo.EntityData.Children = make(map[string]types.YChild)
    trackTypeInfo.EntityData.Children["interface-tracks"] = types.YChild{"InterfaceTracks", &trackTypeInfo.InterfaceTracks}
    trackTypeInfo.EntityData.Children["route-tracks"] = types.YChild{"RouteTracks", &trackTypeInfo.RouteTracks}
    trackTypeInfo.EntityData.Children["ipsla-tracks"] = types.YChild{"IpslaTracks", &trackTypeInfo.IpslaTracks}
    trackTypeInfo.EntityData.Children["bfd-tracks"] = types.YChild{"BfdTracks", &trackTypeInfo.BfdTracks}
    trackTypeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackTypeInfo.EntityData.Leafs["discriminant"] = types.YLeaf{"Discriminant", trackTypeInfo.Discriminant}
    return &(trackTypeInfo.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetEntityData() *types.CommonEntityData {
    interfaceTracks.EntityData.YFilter = interfaceTracks.YFilter
    interfaceTracks.EntityData.YangName = "interface-tracks"
    interfaceTracks.EntityData.BundleName = "cisco_ios_xr"
    interfaceTracks.EntityData.ParentYangName = "track-type-info"
    interfaceTracks.EntityData.SegmentPath = "interface-tracks"
    interfaceTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTracks.EntityData.Children = make(map[string]types.YChild)
    interfaceTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTracks.InterfaceName}
    return &(interfaceTracks.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetEntityData() *types.CommonEntityData {
    routeTracks.EntityData.YFilter = routeTracks.YFilter
    routeTracks.EntityData.YangName = "route-tracks"
    routeTracks.EntityData.BundleName = "cisco_ios_xr"
    routeTracks.EntityData.ParentYangName = "track-type-info"
    routeTracks.EntityData.SegmentPath = "route-tracks"
    routeTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTracks.EntityData.Children = make(map[string]types.YChild)
    routeTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTracks.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", routeTracks.Prefix}
    routeTracks.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", routeTracks.PrefixLength}
    routeTracks.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", routeTracks.Vrf}
    routeTracks.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", routeTracks.NextHop}
    return &(routeTracks.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetEntityData() *types.CommonEntityData {
    ipslaTracks.EntityData.YFilter = ipslaTracks.YFilter
    ipslaTracks.EntityData.YangName = "ipsla-tracks"
    ipslaTracks.EntityData.BundleName = "cisco_ios_xr"
    ipslaTracks.EntityData.ParentYangName = "track-type-info"
    ipslaTracks.EntityData.SegmentPath = "ipsla-tracks"
    ipslaTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipslaTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipslaTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipslaTracks.EntityData.Children = make(map[string]types.YChild)
    ipslaTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    ipslaTracks.EntityData.Leafs["ipsla-op-id"] = types.YLeaf{"IpslaOpId", ipslaTracks.IpslaOpId}
    ipslaTracks.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", ipslaTracks.Rtt}
    ipslaTracks.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", ipslaTracks.ReturnCode}
    return &(ipslaTracks.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetEntityData() *types.CommonEntityData {
    bfdTracks.EntityData.YFilter = bfdTracks.YFilter
    bfdTracks.EntityData.YangName = "bfd-tracks"
    bfdTracks.EntityData.BundleName = "cisco_ios_xr"
    bfdTracks.EntityData.ParentYangName = "track-type-info"
    bfdTracks.EntityData.SegmentPath = "bfd-tracks"
    bfdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdTracks.EntityData.Children = make(map[string]types.YChild)
    bfdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfdTracks.InterfaceName}
    bfdTracks.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", bfdTracks.DestinationAddress}
    bfdTracks.EntityData.Leafs["rate"] = types.YLeaf{"Rate", bfdTracks.Rate}
    bfdTracks.EntityData.Leafs["debounce-count"] = types.YLeaf{"DebounceCount", bfdTracks.DebounceCount}
    return &(bfdTracks.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_BoolTracks
// boolean objects
type ObjectTracking_Tracks_Track_TrackInfo_BoolTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bool track info. The type is slice of
    // ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo.
    BoolTrackInfo []ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo
}

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetEntityData() *types.CommonEntityData {
    boolTracks.EntityData.YFilter = boolTracks.YFilter
    boolTracks.EntityData.YangName = "bool-tracks"
    boolTracks.EntityData.BundleName = "cisco_ios_xr"
    boolTracks.EntityData.ParentYangName = "track-info"
    boolTracks.EntityData.SegmentPath = "bool-tracks"
    boolTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    boolTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    boolTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    boolTracks.EntityData.Children = make(map[string]types.YChild)
    boolTracks.EntityData.Children["bool-track-info"] = types.YChild{"BoolTrackInfo", nil}
    for i := range boolTracks.BoolTrackInfo {
        boolTracks.EntityData.Children[types.GetSegmentPath(&boolTracks.BoolTrackInfo[i])] = types.YChild{"BoolTrackInfo", &boolTracks.BoolTrackInfo[i]}
    }
    boolTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(boolTracks.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo
// bool track info
type ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object Name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track object with Not. The type is bool.
    WithNot interface{}
}

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetEntityData() *types.CommonEntityData {
    boolTrackInfo.EntityData.YFilter = boolTrackInfo.YFilter
    boolTrackInfo.EntityData.YangName = "bool-track-info"
    boolTrackInfo.EntityData.BundleName = "cisco_ios_xr"
    boolTrackInfo.EntityData.ParentYangName = "bool-tracks"
    boolTrackInfo.EntityData.SegmentPath = "bool-track-info"
    boolTrackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    boolTrackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    boolTrackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    boolTrackInfo.EntityData.Children = make(map[string]types.YChild)
    boolTrackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    boolTrackInfo.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", boolTrackInfo.ObjectName}
    boolTrackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", boolTrackInfo.TrackState}
    boolTrackInfo.EntityData.Leafs["with-not"] = types.YLeaf{"WithNot", boolTrackInfo.WithNot}
    return &(boolTrackInfo.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks
// Threshold objects
type ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // threshold track info. The type is slice of
    // ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo.
    ThresholdTrackInfo []ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo
}

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetEntityData() *types.CommonEntityData {
    thresholdTracks.EntityData.YFilter = thresholdTracks.YFilter
    thresholdTracks.EntityData.YangName = "threshold-tracks"
    thresholdTracks.EntityData.BundleName = "cisco_ios_xr"
    thresholdTracks.EntityData.ParentYangName = "track-info"
    thresholdTracks.EntityData.SegmentPath = "threshold-tracks"
    thresholdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdTracks.EntityData.Children = make(map[string]types.YChild)
    thresholdTracks.EntityData.Children["threshold-track-info"] = types.YChild{"ThresholdTrackInfo", nil}
    for i := range thresholdTracks.ThresholdTrackInfo {
        thresholdTracks.EntityData.Children[types.GetSegmentPath(&thresholdTracks.ThresholdTrackInfo[i])] = types.YChild{"ThresholdTrackInfo", &thresholdTracks.ThresholdTrackInfo[i]}
    }
    thresholdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdTracks.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo
// threshold track info
type ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. True means track is up; False means track is down. The type is
    // bool.
    TrackState interface{}

    // Weight is the number assigned to a track object . In case of a type
    // threshold weight( i.e. weighted sum list), weight is asigned by User at the
    // time of configuration. In case of a type threshold percentage (i.e.
    // percentage based list), weight is internally computed by (1/N)x100, where N
    // is the number of objects in the list. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Weight interface{}
}

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetEntityData() *types.CommonEntityData {
    thresholdTrackInfo.EntityData.YFilter = thresholdTrackInfo.YFilter
    thresholdTrackInfo.EntityData.YangName = "threshold-track-info"
    thresholdTrackInfo.EntityData.BundleName = "cisco_ios_xr"
    thresholdTrackInfo.EntityData.ParentYangName = "threshold-tracks"
    thresholdTrackInfo.EntityData.SegmentPath = "threshold-track-info"
    thresholdTrackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdTrackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdTrackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdTrackInfo.EntityData.Children = make(map[string]types.YChild)
    thresholdTrackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdTrackInfo.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", thresholdTrackInfo.ObjectName}
    thresholdTrackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", thresholdTrackInfo.TrackState}
    thresholdTrackInfo.EntityData.Leafs["weight"] = types.YLeaf{"Weight", thresholdTrackInfo.Weight}
    return &(thresholdTrackInfo.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces
// Tracking Interfaces
type ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface tracking info. The type is slice of
    // ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo.
    InterfaceTrackingInfo []ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
}

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetEntityData() *types.CommonEntityData {
    trackingInteraces.EntityData.YFilter = trackingInteraces.YFilter
    trackingInteraces.EntityData.YangName = "tracking-interaces"
    trackingInteraces.EntityData.BundleName = "cisco_ios_xr"
    trackingInteraces.EntityData.ParentYangName = "track-info"
    trackingInteraces.EntityData.SegmentPath = "tracking-interaces"
    trackingInteraces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackingInteraces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackingInteraces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackingInteraces.EntityData.Children = make(map[string]types.YChild)
    trackingInteraces.EntityData.Children["interface-tracking-info"] = types.YChild{"InterfaceTrackingInfo", nil}
    for i := range trackingInteraces.InterfaceTrackingInfo {
        trackingInteraces.EntityData.Children[types.GetSegmentPath(&trackingInteraces.InterfaceTrackingInfo[i])] = types.YChild{"InterfaceTrackingInfo", &trackingInteraces.InterfaceTrackingInfo[i]}
    }
    trackingInteraces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackingInteraces.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
// interface tracking info
type ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetEntityData() *types.CommonEntityData {
    interfaceTrackingInfo.EntityData.YFilter = interfaceTrackingInfo.YFilter
    interfaceTrackingInfo.EntityData.YangName = "interface-tracking-info"
    interfaceTrackingInfo.EntityData.BundleName = "cisco_ios_xr"
    interfaceTrackingInfo.EntityData.ParentYangName = "tracking-interaces"
    interfaceTrackingInfo.EntityData.SegmentPath = "interface-tracking-info"
    interfaceTrackingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTrackingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTrackingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTrackingInfo.EntityData.Children = make(map[string]types.YChild)
    interfaceTrackingInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTrackingInfo.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTrackingInfo.InterfaceName}
    return &(interfaceTrackingInfo.EntityData)
}

// ObjectTracking_Tracks_Track_TrackInfo_Delayed
// Is the state change delay counter in progress
type ObjectTracking_Tracks_Track_TrackInfo_Delayed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time remaining in seconds for the counter to trigger state change. The
    // type is interface{} with range: 0..4294967295. Units are second.
    TimeRemaining interface{}

    // State the track will transition to. Track state. True means track is up;
    // False means track is down. The type is bool.
    TrackState interface{}
}

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetEntityData() *types.CommonEntityData {
    delayed.EntityData.YFilter = delayed.YFilter
    delayed.EntityData.YangName = "delayed"
    delayed.EntityData.BundleName = "cisco_ios_xr"
    delayed.EntityData.ParentYangName = "track-info"
    delayed.EntityData.SegmentPath = "delayed"
    delayed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayed.EntityData.Children = make(map[string]types.YChild)
    delayed.EntityData.Leafs = make(map[string]types.YLeaf)
    delayed.EntityData.Leafs["time-remaining"] = types.YLeaf{"TimeRemaining", delayed.TimeRemaining}
    delayed.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", delayed.TrackState}
    return &(delayed.EntityData)
}

// ObjectTracking_TrackTypeIpv4RouteBrief
// Object Tracking Type Ipv4 Route brief info
type ObjectTracking_TrackTypeIpv4RouteBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // track info brief. The type is slice of
    // ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief.
    TrackInfoBrief []ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief
}

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetEntityData() *types.CommonEntityData {
    trackTypeIpv4RouteBrief.EntityData.YFilter = trackTypeIpv4RouteBrief.YFilter
    trackTypeIpv4RouteBrief.EntityData.YangName = "track-type-ipv4-route-brief"
    trackTypeIpv4RouteBrief.EntityData.BundleName = "cisco_ios_xr"
    trackTypeIpv4RouteBrief.EntityData.ParentYangName = "object-tracking"
    trackTypeIpv4RouteBrief.EntityData.SegmentPath = "track-type-ipv4-route-brief"
    trackTypeIpv4RouteBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeIpv4RouteBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeIpv4RouteBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeIpv4RouteBrief.EntityData.Children = make(map[string]types.YChild)
    trackTypeIpv4RouteBrief.EntityData.Children["track-info-brief"] = types.YChild{"TrackInfoBrief", nil}
    for i := range trackTypeIpv4RouteBrief.TrackInfoBrief {
        trackTypeIpv4RouteBrief.EntityData.Children[types.GetSegmentPath(&trackTypeIpv4RouteBrief.TrackInfoBrief[i])] = types.YChild{"TrackInfoBrief", &trackTypeIpv4RouteBrief.TrackInfoBrief[i]}
    }
    trackTypeIpv4RouteBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackTypeIpv4RouteBrief.EntityData)
}

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief
// track info brief
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type_ interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo
}

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetEntityData() *types.CommonEntityData {
    trackInfoBrief.EntityData.YFilter = trackInfoBrief.YFilter
    trackInfoBrief.EntityData.YangName = "track-info-brief"
    trackInfoBrief.EntityData.BundleName = "cisco_ios_xr"
    trackInfoBrief.EntityData.ParentYangName = "track-type-ipv4-route-brief"
    trackInfoBrief.EntityData.SegmentPath = "track-info-brief"
    trackInfoBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackInfoBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackInfoBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackInfoBrief.EntityData.Children = make(map[string]types.YChild)
    trackInfoBrief.EntityData.Children["track-type-info"] = types.YChild{"TrackTypeInfo", &trackInfoBrief.TrackTypeInfo}
    trackInfoBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    trackInfoBrief.EntityData.Leafs["tracke-name"] = types.YLeaf{"TrackeName", trackInfoBrief.TrackeName}
    trackInfoBrief.EntityData.Leafs["type"] = types.YLeaf{"Type_", trackInfoBrief.Type_}
    trackInfoBrief.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", trackInfoBrief.TrackState}
    return &(trackInfoBrief.EntityData)
}

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetEntityData() *types.CommonEntityData {
    trackTypeInfo.EntityData.YFilter = trackTypeInfo.YFilter
    trackTypeInfo.EntityData.YangName = "track-type-info"
    trackTypeInfo.EntityData.BundleName = "cisco_ios_xr"
    trackTypeInfo.EntityData.ParentYangName = "track-info-brief"
    trackTypeInfo.EntityData.SegmentPath = "track-type-info"
    trackTypeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeInfo.EntityData.Children = make(map[string]types.YChild)
    trackTypeInfo.EntityData.Children["interface-tracks"] = types.YChild{"InterfaceTracks", &trackTypeInfo.InterfaceTracks}
    trackTypeInfo.EntityData.Children["route-tracks"] = types.YChild{"RouteTracks", &trackTypeInfo.RouteTracks}
    trackTypeInfo.EntityData.Children["ipsla-tracks"] = types.YChild{"IpslaTracks", &trackTypeInfo.IpslaTracks}
    trackTypeInfo.EntityData.Children["bfd-tracks"] = types.YChild{"BfdTracks", &trackTypeInfo.BfdTracks}
    trackTypeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackTypeInfo.EntityData.Leafs["discriminant"] = types.YLeaf{"Discriminant", trackTypeInfo.Discriminant}
    return &(trackTypeInfo.EntityData)
}

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetEntityData() *types.CommonEntityData {
    interfaceTracks.EntityData.YFilter = interfaceTracks.YFilter
    interfaceTracks.EntityData.YangName = "interface-tracks"
    interfaceTracks.EntityData.BundleName = "cisco_ios_xr"
    interfaceTracks.EntityData.ParentYangName = "track-type-info"
    interfaceTracks.EntityData.SegmentPath = "interface-tracks"
    interfaceTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTracks.EntityData.Children = make(map[string]types.YChild)
    interfaceTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTracks.InterfaceName}
    return &(interfaceTracks.EntityData)
}

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetEntityData() *types.CommonEntityData {
    routeTracks.EntityData.YFilter = routeTracks.YFilter
    routeTracks.EntityData.YangName = "route-tracks"
    routeTracks.EntityData.BundleName = "cisco_ios_xr"
    routeTracks.EntityData.ParentYangName = "track-type-info"
    routeTracks.EntityData.SegmentPath = "route-tracks"
    routeTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTracks.EntityData.Children = make(map[string]types.YChild)
    routeTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTracks.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", routeTracks.Prefix}
    routeTracks.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", routeTracks.PrefixLength}
    routeTracks.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", routeTracks.Vrf}
    routeTracks.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", routeTracks.NextHop}
    return &(routeTracks.EntityData)
}

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetEntityData() *types.CommonEntityData {
    ipslaTracks.EntityData.YFilter = ipslaTracks.YFilter
    ipslaTracks.EntityData.YangName = "ipsla-tracks"
    ipslaTracks.EntityData.BundleName = "cisco_ios_xr"
    ipslaTracks.EntityData.ParentYangName = "track-type-info"
    ipslaTracks.EntityData.SegmentPath = "ipsla-tracks"
    ipslaTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipslaTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipslaTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipslaTracks.EntityData.Children = make(map[string]types.YChild)
    ipslaTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    ipslaTracks.EntityData.Leafs["ipsla-op-id"] = types.YLeaf{"IpslaOpId", ipslaTracks.IpslaOpId}
    ipslaTracks.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", ipslaTracks.Rtt}
    ipslaTracks.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", ipslaTracks.ReturnCode}
    return &(ipslaTracks.EntityData)
}

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetEntityData() *types.CommonEntityData {
    bfdTracks.EntityData.YFilter = bfdTracks.YFilter
    bfdTracks.EntityData.YangName = "bfd-tracks"
    bfdTracks.EntityData.BundleName = "cisco_ios_xr"
    bfdTracks.EntityData.ParentYangName = "track-type-info"
    bfdTracks.EntityData.SegmentPath = "bfd-tracks"
    bfdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdTracks.EntityData.Children = make(map[string]types.YChild)
    bfdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfdTracks.InterfaceName}
    bfdTracks.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", bfdTracks.DestinationAddress}
    bfdTracks.EntityData.Leafs["rate"] = types.YLeaf{"Rate", bfdTracks.Rate}
    bfdTracks.EntityData.Leafs["debounce-count"] = types.YLeaf{"DebounceCount", bfdTracks.DebounceCount}
    return &(bfdTracks.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route
// Object Tracking Type IPV4 route info
type ObjectTracking_TrackTypeIpv4Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // track info. The type is slice of
    // ObjectTracking_TrackTypeIpv4Route_TrackInfo.
    TrackInfo []ObjectTracking_TrackTypeIpv4Route_TrackInfo
}

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetEntityData() *types.CommonEntityData {
    trackTypeIpv4Route.EntityData.YFilter = trackTypeIpv4Route.YFilter
    trackTypeIpv4Route.EntityData.YangName = "track-type-ipv4-route"
    trackTypeIpv4Route.EntityData.BundleName = "cisco_ios_xr"
    trackTypeIpv4Route.EntityData.ParentYangName = "object-tracking"
    trackTypeIpv4Route.EntityData.SegmentPath = "track-type-ipv4-route"
    trackTypeIpv4Route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeIpv4Route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeIpv4Route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeIpv4Route.EntityData.Children = make(map[string]types.YChild)
    trackTypeIpv4Route.EntityData.Children["track-info"] = types.YChild{"TrackInfo", nil}
    for i := range trackTypeIpv4Route.TrackInfo {
        trackTypeIpv4Route.EntityData.Children[types.GetSegmentPath(&trackTypeIpv4Route.TrackInfo[i])] = types.YChild{"TrackInfo", &trackTypeIpv4Route.TrackInfo[i]}
    }
    trackTypeIpv4Route.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackTypeIpv4Route.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo
// track info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type_ interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // State Change Counter. The type is interface{} with range: 0..4294967295.
    StateChangeCounter interface{}

    // Seconds Last Change. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecondsLastChange interface{}

    // User specified threshold upper limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdUp interface{}

    // User specified threshold lower limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdDown interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo

    // boolean objects.
    BoolTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks

    // Threshold objects.
    ThresholdTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks

    // Tracking Interfaces.
    TrackingInteraces ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces

    // Is the state change delay counter in progress.
    Delayed ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed
}

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetEntityData() *types.CommonEntityData {
    trackInfo.EntityData.YFilter = trackInfo.YFilter
    trackInfo.EntityData.YangName = "track-info"
    trackInfo.EntityData.BundleName = "cisco_ios_xr"
    trackInfo.EntityData.ParentYangName = "track-type-ipv4-route"
    trackInfo.EntityData.SegmentPath = "track-info"
    trackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackInfo.EntityData.Children = make(map[string]types.YChild)
    trackInfo.EntityData.Children["track-type-info"] = types.YChild{"TrackTypeInfo", &trackInfo.TrackTypeInfo}
    trackInfo.EntityData.Children["bool-tracks"] = types.YChild{"BoolTracks", &trackInfo.BoolTracks}
    trackInfo.EntityData.Children["threshold-tracks"] = types.YChild{"ThresholdTracks", &trackInfo.ThresholdTracks}
    trackInfo.EntityData.Children["tracking-interaces"] = types.YChild{"TrackingInteraces", &trackInfo.TrackingInteraces}
    trackInfo.EntityData.Children["delayed"] = types.YChild{"Delayed", &trackInfo.Delayed}
    trackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackInfo.EntityData.Leafs["tracke-name"] = types.YLeaf{"TrackeName", trackInfo.TrackeName}
    trackInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", trackInfo.Type_}
    trackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", trackInfo.TrackState}
    trackInfo.EntityData.Leafs["state-change-counter"] = types.YLeaf{"StateChangeCounter", trackInfo.StateChangeCounter}
    trackInfo.EntityData.Leafs["seconds-last-change"] = types.YLeaf{"SecondsLastChange", trackInfo.SecondsLastChange}
    trackInfo.EntityData.Leafs["threshold-up"] = types.YLeaf{"ThresholdUp", trackInfo.ThresholdUp}
    trackInfo.EntityData.Leafs["threshold-down"] = types.YLeaf{"ThresholdDown", trackInfo.ThresholdDown}
    return &(trackInfo.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetEntityData() *types.CommonEntityData {
    trackTypeInfo.EntityData.YFilter = trackTypeInfo.YFilter
    trackTypeInfo.EntityData.YangName = "track-type-info"
    trackTypeInfo.EntityData.BundleName = "cisco_ios_xr"
    trackTypeInfo.EntityData.ParentYangName = "track-info"
    trackTypeInfo.EntityData.SegmentPath = "track-type-info"
    trackTypeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeInfo.EntityData.Children = make(map[string]types.YChild)
    trackTypeInfo.EntityData.Children["interface-tracks"] = types.YChild{"InterfaceTracks", &trackTypeInfo.InterfaceTracks}
    trackTypeInfo.EntityData.Children["route-tracks"] = types.YChild{"RouteTracks", &trackTypeInfo.RouteTracks}
    trackTypeInfo.EntityData.Children["ipsla-tracks"] = types.YChild{"IpslaTracks", &trackTypeInfo.IpslaTracks}
    trackTypeInfo.EntityData.Children["bfd-tracks"] = types.YChild{"BfdTracks", &trackTypeInfo.BfdTracks}
    trackTypeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackTypeInfo.EntityData.Leafs["discriminant"] = types.YLeaf{"Discriminant", trackTypeInfo.Discriminant}
    return &(trackTypeInfo.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetEntityData() *types.CommonEntityData {
    interfaceTracks.EntityData.YFilter = interfaceTracks.YFilter
    interfaceTracks.EntityData.YangName = "interface-tracks"
    interfaceTracks.EntityData.BundleName = "cisco_ios_xr"
    interfaceTracks.EntityData.ParentYangName = "track-type-info"
    interfaceTracks.EntityData.SegmentPath = "interface-tracks"
    interfaceTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTracks.EntityData.Children = make(map[string]types.YChild)
    interfaceTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTracks.InterfaceName}
    return &(interfaceTracks.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetEntityData() *types.CommonEntityData {
    routeTracks.EntityData.YFilter = routeTracks.YFilter
    routeTracks.EntityData.YangName = "route-tracks"
    routeTracks.EntityData.BundleName = "cisco_ios_xr"
    routeTracks.EntityData.ParentYangName = "track-type-info"
    routeTracks.EntityData.SegmentPath = "route-tracks"
    routeTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTracks.EntityData.Children = make(map[string]types.YChild)
    routeTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTracks.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", routeTracks.Prefix}
    routeTracks.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", routeTracks.PrefixLength}
    routeTracks.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", routeTracks.Vrf}
    routeTracks.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", routeTracks.NextHop}
    return &(routeTracks.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetEntityData() *types.CommonEntityData {
    ipslaTracks.EntityData.YFilter = ipslaTracks.YFilter
    ipslaTracks.EntityData.YangName = "ipsla-tracks"
    ipslaTracks.EntityData.BundleName = "cisco_ios_xr"
    ipslaTracks.EntityData.ParentYangName = "track-type-info"
    ipslaTracks.EntityData.SegmentPath = "ipsla-tracks"
    ipslaTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipslaTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipslaTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipslaTracks.EntityData.Children = make(map[string]types.YChild)
    ipslaTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    ipslaTracks.EntityData.Leafs["ipsla-op-id"] = types.YLeaf{"IpslaOpId", ipslaTracks.IpslaOpId}
    ipslaTracks.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", ipslaTracks.Rtt}
    ipslaTracks.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", ipslaTracks.ReturnCode}
    return &(ipslaTracks.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetEntityData() *types.CommonEntityData {
    bfdTracks.EntityData.YFilter = bfdTracks.YFilter
    bfdTracks.EntityData.YangName = "bfd-tracks"
    bfdTracks.EntityData.BundleName = "cisco_ios_xr"
    bfdTracks.EntityData.ParentYangName = "track-type-info"
    bfdTracks.EntityData.SegmentPath = "bfd-tracks"
    bfdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdTracks.EntityData.Children = make(map[string]types.YChild)
    bfdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfdTracks.InterfaceName}
    bfdTracks.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", bfdTracks.DestinationAddress}
    bfdTracks.EntityData.Leafs["rate"] = types.YLeaf{"Rate", bfdTracks.Rate}
    bfdTracks.EntityData.Leafs["debounce-count"] = types.YLeaf{"DebounceCount", bfdTracks.DebounceCount}
    return &(bfdTracks.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks
// boolean objects
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bool track info. The type is slice of
    // ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo.
    BoolTrackInfo []ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo
}

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetEntityData() *types.CommonEntityData {
    boolTracks.EntityData.YFilter = boolTracks.YFilter
    boolTracks.EntityData.YangName = "bool-tracks"
    boolTracks.EntityData.BundleName = "cisco_ios_xr"
    boolTracks.EntityData.ParentYangName = "track-info"
    boolTracks.EntityData.SegmentPath = "bool-tracks"
    boolTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    boolTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    boolTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    boolTracks.EntityData.Children = make(map[string]types.YChild)
    boolTracks.EntityData.Children["bool-track-info"] = types.YChild{"BoolTrackInfo", nil}
    for i := range boolTracks.BoolTrackInfo {
        boolTracks.EntityData.Children[types.GetSegmentPath(&boolTracks.BoolTrackInfo[i])] = types.YChild{"BoolTrackInfo", &boolTracks.BoolTrackInfo[i]}
    }
    boolTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(boolTracks.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo
// bool track info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object Name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track object with Not. The type is bool.
    WithNot interface{}
}

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetEntityData() *types.CommonEntityData {
    boolTrackInfo.EntityData.YFilter = boolTrackInfo.YFilter
    boolTrackInfo.EntityData.YangName = "bool-track-info"
    boolTrackInfo.EntityData.BundleName = "cisco_ios_xr"
    boolTrackInfo.EntityData.ParentYangName = "bool-tracks"
    boolTrackInfo.EntityData.SegmentPath = "bool-track-info"
    boolTrackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    boolTrackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    boolTrackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    boolTrackInfo.EntityData.Children = make(map[string]types.YChild)
    boolTrackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    boolTrackInfo.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", boolTrackInfo.ObjectName}
    boolTrackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", boolTrackInfo.TrackState}
    boolTrackInfo.EntityData.Leafs["with-not"] = types.YLeaf{"WithNot", boolTrackInfo.WithNot}
    return &(boolTrackInfo.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks
// Threshold objects
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // threshold track info. The type is slice of
    // ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo.
    ThresholdTrackInfo []ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo
}

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetEntityData() *types.CommonEntityData {
    thresholdTracks.EntityData.YFilter = thresholdTracks.YFilter
    thresholdTracks.EntityData.YangName = "threshold-tracks"
    thresholdTracks.EntityData.BundleName = "cisco_ios_xr"
    thresholdTracks.EntityData.ParentYangName = "track-info"
    thresholdTracks.EntityData.SegmentPath = "threshold-tracks"
    thresholdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdTracks.EntityData.Children = make(map[string]types.YChild)
    thresholdTracks.EntityData.Children["threshold-track-info"] = types.YChild{"ThresholdTrackInfo", nil}
    for i := range thresholdTracks.ThresholdTrackInfo {
        thresholdTracks.EntityData.Children[types.GetSegmentPath(&thresholdTracks.ThresholdTrackInfo[i])] = types.YChild{"ThresholdTrackInfo", &thresholdTracks.ThresholdTrackInfo[i]}
    }
    thresholdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdTracks.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo
// threshold track info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. True means track is up; False means track is down. The type is
    // bool.
    TrackState interface{}

    // Weight is the number assigned to a track object . In case of a type
    // threshold weight( i.e. weighted sum list), weight is asigned by User at the
    // time of configuration. In case of a type threshold percentage (i.e.
    // percentage based list), weight is internally computed by (1/N)x100, where N
    // is the number of objects in the list. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Weight interface{}
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetEntityData() *types.CommonEntityData {
    thresholdTrackInfo.EntityData.YFilter = thresholdTrackInfo.YFilter
    thresholdTrackInfo.EntityData.YangName = "threshold-track-info"
    thresholdTrackInfo.EntityData.BundleName = "cisco_ios_xr"
    thresholdTrackInfo.EntityData.ParentYangName = "threshold-tracks"
    thresholdTrackInfo.EntityData.SegmentPath = "threshold-track-info"
    thresholdTrackInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdTrackInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdTrackInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdTrackInfo.EntityData.Children = make(map[string]types.YChild)
    thresholdTrackInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdTrackInfo.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", thresholdTrackInfo.ObjectName}
    thresholdTrackInfo.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", thresholdTrackInfo.TrackState}
    thresholdTrackInfo.EntityData.Leafs["weight"] = types.YLeaf{"Weight", thresholdTrackInfo.Weight}
    return &(thresholdTrackInfo.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces
// Tracking Interfaces
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface tracking info. The type is slice of
    // ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo.
    InterfaceTrackingInfo []ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
}

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetEntityData() *types.CommonEntityData {
    trackingInteraces.EntityData.YFilter = trackingInteraces.YFilter
    trackingInteraces.EntityData.YangName = "tracking-interaces"
    trackingInteraces.EntityData.BundleName = "cisco_ios_xr"
    trackingInteraces.EntityData.ParentYangName = "track-info"
    trackingInteraces.EntityData.SegmentPath = "tracking-interaces"
    trackingInteraces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackingInteraces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackingInteraces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackingInteraces.EntityData.Children = make(map[string]types.YChild)
    trackingInteraces.EntityData.Children["interface-tracking-info"] = types.YChild{"InterfaceTrackingInfo", nil}
    for i := range trackingInteraces.InterfaceTrackingInfo {
        trackingInteraces.EntityData.Children[types.GetSegmentPath(&trackingInteraces.InterfaceTrackingInfo[i])] = types.YChild{"InterfaceTrackingInfo", &trackingInteraces.InterfaceTrackingInfo[i]}
    }
    trackingInteraces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackingInteraces.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
// interface tracking info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetEntityData() *types.CommonEntityData {
    interfaceTrackingInfo.EntityData.YFilter = interfaceTrackingInfo.YFilter
    interfaceTrackingInfo.EntityData.YangName = "interface-tracking-info"
    interfaceTrackingInfo.EntityData.BundleName = "cisco_ios_xr"
    interfaceTrackingInfo.EntityData.ParentYangName = "tracking-interaces"
    interfaceTrackingInfo.EntityData.SegmentPath = "interface-tracking-info"
    interfaceTrackingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTrackingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTrackingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTrackingInfo.EntityData.Children = make(map[string]types.YChild)
    interfaceTrackingInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTrackingInfo.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTrackingInfo.InterfaceName}
    return &(interfaceTrackingInfo.EntityData)
}

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed
// Is the state change delay counter in progress
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time remaining in seconds for the counter to trigger state change. The
    // type is interface{} with range: 0..4294967295. Units are second.
    TimeRemaining interface{}

    // State the track will transition to. Track state. True means track is up;
    // False means track is down. The type is bool.
    TrackState interface{}
}

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetEntityData() *types.CommonEntityData {
    delayed.EntityData.YFilter = delayed.YFilter
    delayed.EntityData.YangName = "delayed"
    delayed.EntityData.BundleName = "cisco_ios_xr"
    delayed.EntityData.ParentYangName = "track-info"
    delayed.EntityData.SegmentPath = "delayed"
    delayed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayed.EntityData.Children = make(map[string]types.YChild)
    delayed.EntityData.Leafs = make(map[string]types.YLeaf)
    delayed.EntityData.Leafs["time-remaining"] = types.YLeaf{"TimeRemaining", delayed.TimeRemaining}
    delayed.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", delayed.TrackState}
    return &(delayed.EntityData)
}

// ObjectTracking_TrackTypeInterfaceBrief
// Object Tracking Type Interface brief info
type ObjectTracking_TrackTypeInterfaceBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // track info brief. The type is slice of
    // ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief.
    TrackInfoBrief []ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief
}

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetEntityData() *types.CommonEntityData {
    trackTypeInterfaceBrief.EntityData.YFilter = trackTypeInterfaceBrief.YFilter
    trackTypeInterfaceBrief.EntityData.YangName = "track-type-interface-brief"
    trackTypeInterfaceBrief.EntityData.BundleName = "cisco_ios_xr"
    trackTypeInterfaceBrief.EntityData.ParentYangName = "object-tracking"
    trackTypeInterfaceBrief.EntityData.SegmentPath = "track-type-interface-brief"
    trackTypeInterfaceBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeInterfaceBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeInterfaceBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeInterfaceBrief.EntityData.Children = make(map[string]types.YChild)
    trackTypeInterfaceBrief.EntityData.Children["track-info-brief"] = types.YChild{"TrackInfoBrief", nil}
    for i := range trackTypeInterfaceBrief.TrackInfoBrief {
        trackTypeInterfaceBrief.EntityData.Children[types.GetSegmentPath(&trackTypeInterfaceBrief.TrackInfoBrief[i])] = types.YChild{"TrackInfoBrief", &trackTypeInterfaceBrief.TrackInfoBrief[i]}
    }
    trackTypeInterfaceBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackTypeInterfaceBrief.EntityData)
}

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief
// track info brief
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type_ interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo
}

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetEntityData() *types.CommonEntityData {
    trackInfoBrief.EntityData.YFilter = trackInfoBrief.YFilter
    trackInfoBrief.EntityData.YangName = "track-info-brief"
    trackInfoBrief.EntityData.BundleName = "cisco_ios_xr"
    trackInfoBrief.EntityData.ParentYangName = "track-type-interface-brief"
    trackInfoBrief.EntityData.SegmentPath = "track-info-brief"
    trackInfoBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackInfoBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackInfoBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackInfoBrief.EntityData.Children = make(map[string]types.YChild)
    trackInfoBrief.EntityData.Children["track-type-info"] = types.YChild{"TrackTypeInfo", &trackInfoBrief.TrackTypeInfo}
    trackInfoBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    trackInfoBrief.EntityData.Leafs["tracke-name"] = types.YLeaf{"TrackeName", trackInfoBrief.TrackeName}
    trackInfoBrief.EntityData.Leafs["type"] = types.YLeaf{"Type_", trackInfoBrief.Type_}
    trackInfoBrief.EntityData.Leafs["track-state"] = types.YLeaf{"TrackState", trackInfoBrief.TrackState}
    return &(trackInfoBrief.EntityData)
}

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetEntityData() *types.CommonEntityData {
    trackTypeInfo.EntityData.YFilter = trackTypeInfo.YFilter
    trackTypeInfo.EntityData.YangName = "track-type-info"
    trackTypeInfo.EntityData.BundleName = "cisco_ios_xr"
    trackTypeInfo.EntityData.ParentYangName = "track-info-brief"
    trackTypeInfo.EntityData.SegmentPath = "track-type-info"
    trackTypeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackTypeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackTypeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackTypeInfo.EntityData.Children = make(map[string]types.YChild)
    trackTypeInfo.EntityData.Children["interface-tracks"] = types.YChild{"InterfaceTracks", &trackTypeInfo.InterfaceTracks}
    trackTypeInfo.EntityData.Children["route-tracks"] = types.YChild{"RouteTracks", &trackTypeInfo.RouteTracks}
    trackTypeInfo.EntityData.Children["ipsla-tracks"] = types.YChild{"IpslaTracks", &trackTypeInfo.IpslaTracks}
    trackTypeInfo.EntityData.Children["bfd-tracks"] = types.YChild{"BfdTracks", &trackTypeInfo.BfdTracks}
    trackTypeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trackTypeInfo.EntityData.Leafs["discriminant"] = types.YLeaf{"Discriminant", trackTypeInfo.Discriminant}
    return &(trackTypeInfo.EntityData)
}

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetEntityData() *types.CommonEntityData {
    interfaceTracks.EntityData.YFilter = interfaceTracks.YFilter
    interfaceTracks.EntityData.YangName = "interface-tracks"
    interfaceTracks.EntityData.BundleName = "cisco_ios_xr"
    interfaceTracks.EntityData.ParentYangName = "track-type-info"
    interfaceTracks.EntityData.SegmentPath = "interface-tracks"
    interfaceTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceTracks.EntityData.Children = make(map[string]types.YChild)
    interfaceTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceTracks.InterfaceName}
    return &(interfaceTracks.EntityData)
}

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetEntityData() *types.CommonEntityData {
    routeTracks.EntityData.YFilter = routeTracks.YFilter
    routeTracks.EntityData.YangName = "route-tracks"
    routeTracks.EntityData.BundleName = "cisco_ios_xr"
    routeTracks.EntityData.ParentYangName = "track-type-info"
    routeTracks.EntityData.SegmentPath = "route-tracks"
    routeTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTracks.EntityData.Children = make(map[string]types.YChild)
    routeTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTracks.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", routeTracks.Prefix}
    routeTracks.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", routeTracks.PrefixLength}
    routeTracks.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", routeTracks.Vrf}
    routeTracks.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", routeTracks.NextHop}
    return &(routeTracks.EntityData)
}

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetEntityData() *types.CommonEntityData {
    ipslaTracks.EntityData.YFilter = ipslaTracks.YFilter
    ipslaTracks.EntityData.YangName = "ipsla-tracks"
    ipslaTracks.EntityData.BundleName = "cisco_ios_xr"
    ipslaTracks.EntityData.ParentYangName = "track-type-info"
    ipslaTracks.EntityData.SegmentPath = "ipsla-tracks"
    ipslaTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipslaTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipslaTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipslaTracks.EntityData.Children = make(map[string]types.YChild)
    ipslaTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    ipslaTracks.EntityData.Leafs["ipsla-op-id"] = types.YLeaf{"IpslaOpId", ipslaTracks.IpslaOpId}
    ipslaTracks.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", ipslaTracks.Rtt}
    ipslaTracks.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", ipslaTracks.ReturnCode}
    return &(ipslaTracks.EntityData)
}

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetEntityData() *types.CommonEntityData {
    bfdTracks.EntityData.YFilter = bfdTracks.YFilter
    bfdTracks.EntityData.YangName = "bfd-tracks"
    bfdTracks.EntityData.BundleName = "cisco_ios_xr"
    bfdTracks.EntityData.ParentYangName = "track-type-info"
    bfdTracks.EntityData.SegmentPath = "bfd-tracks"
    bfdTracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdTracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdTracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdTracks.EntityData.Children = make(map[string]types.YChild)
    bfdTracks.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdTracks.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfdTracks.InterfaceName}
    bfdTracks.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", bfdTracks.DestinationAddress}
    bfdTracks.EntityData.Leafs["rate"] = types.YLeaf{"Rate", bfdTracks.Rate}
    bfdTracks.EntityData.Leafs["debounce-count"] = types.YLeaf{"DebounceCount", bfdTracks.DebounceCount}
    return &(bfdTracks.EntityData)
}

