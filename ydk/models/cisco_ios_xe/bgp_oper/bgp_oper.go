// This module contains a collection of YANG definitions
// for bgp operational data.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package bgp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bgp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-bgp-oper bgp-state-data}", reflect.TypeOf(BgpStateData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-bgp-oper:bgp-state-data", reflect.TypeOf(BgpStateData{}))
}

// BgpLink represents Operational state relevent to bgp global neighbor
type BgpLink string

const (
    // iBGP neighbors
    BgpLink_internal BgpLink = "internal"

    // eBGP neighbors
    BgpLink_external BgpLink = "external"
)

// BgpFsmState represents BGP FSM State
type BgpFsmState string

const (
    // neighbor is in Idle state
    BgpFsmState_fsm_idle BgpFsmState = "fsm-idle"

    // neighbor is in Connect state
    BgpFsmState_fsm_connect BgpFsmState = "fsm-connect"

    // neighbor is in Active state
    BgpFsmState_fsm_active BgpFsmState = "fsm-active"

    // neighbor is in OpenSent state
    BgpFsmState_fsm_opensent BgpFsmState = "fsm-opensent"

    // neighbor is in OpenConfirm state
    BgpFsmState_fsm_openconfirm BgpFsmState = "fsm-openconfirm"

    // neighbor is in Established state
    BgpFsmState_fsm_established BgpFsmState = "fsm-established"

    // neighbor is Non Negotiated
    BgpFsmState_fsm_nonnegotiated BgpFsmState = "fsm-nonnegotiated"
)

// BgpMode represents BGP mode
type BgpMode string

const (
    // active connection
    BgpMode_mode_active BgpMode = "mode-active"

    // passive connection
    BgpMode_mode_passive BgpMode = "mode-passive"
)

// BgpStateData
// BGP operational state data
type BgpStateData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP neighbor information.
    Neighbors BgpStateData_Neighbors

    // BGP address family.
    AddressFamilies BgpStateData_AddressFamilies

    // BGP VRFs.
    BgpRouteVrfs BgpStateData_BgpRouteVrfs

    // BGP RDs.
    BgpRouteRds BgpStateData_BgpRouteRds
}

func (bgpStateData *BgpStateData) GetEntityData() *types.CommonEntityData {
    bgpStateData.EntityData.YFilter = bgpStateData.YFilter
    bgpStateData.EntityData.YangName = "bgp-state-data"
    bgpStateData.EntityData.BundleName = "cisco_ios_xe"
    bgpStateData.EntityData.ParentYangName = "Cisco-IOS-XE-bgp-oper"
    bgpStateData.EntityData.SegmentPath = "Cisco-IOS-XE-bgp-oper:bgp-state-data"
    bgpStateData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpStateData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpStateData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpStateData.EntityData.Children = make(map[string]types.YChild)
    bgpStateData.EntityData.Children["neighbors"] = types.YChild{"Neighbors", &bgpStateData.Neighbors}
    bgpStateData.EntityData.Children["address-families"] = types.YChild{"AddressFamilies", &bgpStateData.AddressFamilies}
    bgpStateData.EntityData.Children["bgp-route-vrfs"] = types.YChild{"BgpRouteVrfs", &bgpStateData.BgpRouteVrfs}
    bgpStateData.EntityData.Children["bgp-route-rds"] = types.YChild{"BgpRouteRds", &bgpStateData.BgpRouteRds}
    bgpStateData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpStateData.EntityData)
}

// BgpStateData_Neighbors
// BGP neighbor information
type BgpStateData_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP neighbors. The type is slice of
    // BgpStateData_Neighbors_Neighbor.
    Neighbor []BgpStateData_Neighbors_Neighbor
}

func (neighbors *BgpStateData_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xe"
    neighbors.EntityData.ParentYangName = "bgp-state-data"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    neighbors.EntityData.Children = make(map[string]types.YChild)
    neighbors.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children[types.GetSegmentPath(&neighbors.Neighbor[i])] = types.YChild{"Neighbor", &neighbors.Neighbor[i]}
    }
    neighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighbors.EntityData)
}

// BgpStateData_Neighbors_Neighbor
// List of BGP neighbors
type BgpStateData_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Afi-safi key. The type is AfiSafi.
    AfiSafi interface{}

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // This attribute is a key. Neighbor identifier. The type is string.
    NeighborId interface{}

    // Neighbor description string. The type is string.
    Description interface{}

    // BGP version being used to communicate with the remote router. The type is
    // interface{} with range: 0..65535.
    BgpVersion interface{}

    // Neighbor link type. The type is BgpLink.
    Link interface{}

    // Amout of time the bgp session has been up since being established. The type
    // is string.
    UpTime interface{}

    // Time since BGP last sent a message to the neighbor. The type is string.
    LastWrite interface{}

    // Time since BGP last received a message from the neighbor. The type is
    // string.
    LastRead interface{}

    // The number of installed prefixes. The type is interface{} with range:
    // 0..4294967295.
    InstalledPrefixes interface{}

    // BGP neighbor session state. The type is BgpFsmState.
    SessionState interface{}

    // Negotiated capabilities for neighbor session. The type is slice of string.
    NegotiatedCap []interface{}

    // BGP neighbor AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // Negotiated keepalive timers status of BGP neighbor.
    NegotiatedKeepaliveTimers BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers

    // BGP neighbor session counters.
    BgpNeighborCounters BgpStateData_Neighbors_Neighbor_BgpNeighborCounters

    // BGP neighbor connection.
    Connection BgpStateData_Neighbors_Neighbor_Connection

    // BGP neighbor transport.
    Transport BgpStateData_Neighbors_Neighbor_Transport

    // BGP neighbor activity.
    PrefixActivity BgpStateData_Neighbors_Neighbor_PrefixActivity
}

func (neighbor *BgpStateData_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xe"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + "[afi-safi='" + fmt.Sprintf("%v", neighbor.AfiSafi) + "']" + "[vrf-name='" + fmt.Sprintf("%v", neighbor.VrfName) + "']" + "[neighbor-id='" + fmt.Sprintf("%v", neighbor.NeighborId) + "']"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Children["negotiated-keepalive-timers"] = types.YChild{"NegotiatedKeepaliveTimers", &neighbor.NegotiatedKeepaliveTimers}
    neighbor.EntityData.Children["bgp-neighbor-counters"] = types.YChild{"BgpNeighborCounters", &neighbor.BgpNeighborCounters}
    neighbor.EntityData.Children["connection"] = types.YChild{"Connection", &neighbor.Connection}
    neighbor.EntityData.Children["transport"] = types.YChild{"Transport", &neighbor.Transport}
    neighbor.EntityData.Children["prefix-activity"] = types.YChild{"PrefixActivity", &neighbor.PrefixActivity}
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["afi-safi"] = types.YLeaf{"AfiSafi", neighbor.AfiSafi}
    neighbor.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", neighbor.VrfName}
    neighbor.EntityData.Leafs["neighbor-id"] = types.YLeaf{"NeighborId", neighbor.NeighborId}
    neighbor.EntityData.Leafs["description"] = types.YLeaf{"Description", neighbor.Description}
    neighbor.EntityData.Leafs["bgp-version"] = types.YLeaf{"BgpVersion", neighbor.BgpVersion}
    neighbor.EntityData.Leafs["link"] = types.YLeaf{"Link", neighbor.Link}
    neighbor.EntityData.Leafs["up-time"] = types.YLeaf{"UpTime", neighbor.UpTime}
    neighbor.EntityData.Leafs["last-write"] = types.YLeaf{"LastWrite", neighbor.LastWrite}
    neighbor.EntityData.Leafs["last-read"] = types.YLeaf{"LastRead", neighbor.LastRead}
    neighbor.EntityData.Leafs["installed-prefixes"] = types.YLeaf{"InstalledPrefixes", neighbor.InstalledPrefixes}
    neighbor.EntityData.Leafs["session-state"] = types.YLeaf{"SessionState", neighbor.SessionState}
    neighbor.EntityData.Leafs["negotiated-cap"] = types.YLeaf{"NegotiatedCap", neighbor.NegotiatedCap}
    neighbor.EntityData.Leafs["as"] = types.YLeaf{"As", neighbor.As}
    return &(neighbor.EntityData)
}

// BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers
// Negotiated keepalive timers status of BGP neighbor
type BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Keepalive interval. The type is interface{} with range: 0..65535.
    KeepaliveInterval interface{}
}

func (negotiatedKeepaliveTimers *BgpStateData_Neighbors_Neighbor_NegotiatedKeepaliveTimers) GetEntityData() *types.CommonEntityData {
    negotiatedKeepaliveTimers.EntityData.YFilter = negotiatedKeepaliveTimers.YFilter
    negotiatedKeepaliveTimers.EntityData.YangName = "negotiated-keepalive-timers"
    negotiatedKeepaliveTimers.EntityData.BundleName = "cisco_ios_xe"
    negotiatedKeepaliveTimers.EntityData.ParentYangName = "neighbor"
    negotiatedKeepaliveTimers.EntityData.SegmentPath = "negotiated-keepalive-timers"
    negotiatedKeepaliveTimers.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    negotiatedKeepaliveTimers.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    negotiatedKeepaliveTimers.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    negotiatedKeepaliveTimers.EntityData.Children = make(map[string]types.YChild)
    negotiatedKeepaliveTimers.EntityData.Leafs = make(map[string]types.YLeaf)
    negotiatedKeepaliveTimers.EntityData.Leafs["hold-time"] = types.YLeaf{"HoldTime", negotiatedKeepaliveTimers.HoldTime}
    negotiatedKeepaliveTimers.EntityData.Leafs["keepalive-interval"] = types.YLeaf{"KeepaliveInterval", negotiatedKeepaliveTimers.KeepaliveInterval}
    return &(negotiatedKeepaliveTimers.EntityData)
}

// BgpStateData_Neighbors_Neighbor_BgpNeighborCounters
// BGP neighbor session counters
type BgpStateData_Neighbors_Neighbor_BgpNeighborCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input Q depth. The type is interface{} with range: 0..4294967295.
    InqDepth interface{}

    // Output Q depth. The type is interface{} with range: 0..4294967295.
    OutqDepth interface{}

    // Number of mesaged sent.
    Sent BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent

    // Number of mesaged received.
    Received BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received
}

func (bgpNeighborCounters *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters) GetEntityData() *types.CommonEntityData {
    bgpNeighborCounters.EntityData.YFilter = bgpNeighborCounters.YFilter
    bgpNeighborCounters.EntityData.YangName = "bgp-neighbor-counters"
    bgpNeighborCounters.EntityData.BundleName = "cisco_ios_xe"
    bgpNeighborCounters.EntityData.ParentYangName = "neighbor"
    bgpNeighborCounters.EntityData.SegmentPath = "bgp-neighbor-counters"
    bgpNeighborCounters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpNeighborCounters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpNeighborCounters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpNeighborCounters.EntityData.Children = make(map[string]types.YChild)
    bgpNeighborCounters.EntityData.Children["sent"] = types.YChild{"Sent", &bgpNeighborCounters.Sent}
    bgpNeighborCounters.EntityData.Children["received"] = types.YChild{"Received", &bgpNeighborCounters.Received}
    bgpNeighborCounters.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpNeighborCounters.EntityData.Leafs["inq-depth"] = types.YLeaf{"InqDepth", bgpNeighborCounters.InqDepth}
    bgpNeighborCounters.EntityData.Leafs["outq-depth"] = types.YLeaf{"OutqDepth", bgpNeighborCounters.OutqDepth}
    return &(bgpNeighborCounters.EntityData)
}

// BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent
// Number of mesaged sent
type BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OPEN message count. The type is interface{} with range: 0..4294967295.
    Opens interface{}

    // UPDATE message count. The type is interface{} with range: 0..4294967295.
    Updates interface{}

    // NOTIFICATION message count. The type is interface{} with range:
    // 0..4294967295.
    Notifications interface{}

    // KEEPALIVE message count. The type is interface{} with range: 0..4294967295.
    Keepalives interface{}

    // Route refresh message count. The type is interface{} with range:
    // 0..4294967295.
    RouteRefreshes interface{}
}

func (sent *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Sent) GetEntityData() *types.CommonEntityData {
    sent.EntityData.YFilter = sent.YFilter
    sent.EntityData.YangName = "sent"
    sent.EntityData.BundleName = "cisco_ios_xe"
    sent.EntityData.ParentYangName = "bgp-neighbor-counters"
    sent.EntityData.SegmentPath = "sent"
    sent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sent.EntityData.Children = make(map[string]types.YChild)
    sent.EntityData.Leafs = make(map[string]types.YLeaf)
    sent.EntityData.Leafs["opens"] = types.YLeaf{"Opens", sent.Opens}
    sent.EntityData.Leafs["updates"] = types.YLeaf{"Updates", sent.Updates}
    sent.EntityData.Leafs["notifications"] = types.YLeaf{"Notifications", sent.Notifications}
    sent.EntityData.Leafs["keepalives"] = types.YLeaf{"Keepalives", sent.Keepalives}
    sent.EntityData.Leafs["route-refreshes"] = types.YLeaf{"RouteRefreshes", sent.RouteRefreshes}
    return &(sent.EntityData)
}

// BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received
// Number of mesaged received
type BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OPEN message count. The type is interface{} with range: 0..4294967295.
    Opens interface{}

    // UPDATE message count. The type is interface{} with range: 0..4294967295.
    Updates interface{}

    // NOTIFICATION message count. The type is interface{} with range:
    // 0..4294967295.
    Notifications interface{}

    // KEEPALIVE message count. The type is interface{} with range: 0..4294967295.
    Keepalives interface{}

    // Route refresh message count. The type is interface{} with range:
    // 0..4294967295.
    RouteRefreshes interface{}
}

func (received *BgpStateData_Neighbors_Neighbor_BgpNeighborCounters_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xe"
    received.EntityData.ParentYangName = "bgp-neighbor-counters"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    received.EntityData.Children = make(map[string]types.YChild)
    received.EntityData.Leafs = make(map[string]types.YLeaf)
    received.EntityData.Leafs["opens"] = types.YLeaf{"Opens", received.Opens}
    received.EntityData.Leafs["updates"] = types.YLeaf{"Updates", received.Updates}
    received.EntityData.Leafs["notifications"] = types.YLeaf{"Notifications", received.Notifications}
    received.EntityData.Leafs["keepalives"] = types.YLeaf{"Keepalives", received.Keepalives}
    received.EntityData.Leafs["route-refreshes"] = types.YLeaf{"RouteRefreshes", received.RouteRefreshes}
    return &(received.EntityData)
}

// BgpStateData_Neighbors_Neighbor_Connection
// BGP neighbor connection
type BgpStateData_Neighbors_Neighbor_Connection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TCP FSM state. The type is TcpFsmState.
    State interface{}

    // BGP transport connection mode. The type is BgpMode.
    Mode interface{}

    // The number of times a TCP and BGP  connection has been successfully
    // established. The type is interface{} with range: 0..4294967295.
    TotalEstablished interface{}

    // The number of times that a valid session has failed or been taken down. The
    // type is interface{} with range: 0..4294967295.
    TotalDropped interface{}

    // Time since peering session was last reset. The type is string.
    LastReset interface{}

    // The reason for the last reset. The type is string.
    ResetReason interface{}
}

func (connection *BgpStateData_Neighbors_Neighbor_Connection) GetEntityData() *types.CommonEntityData {
    connection.EntityData.YFilter = connection.YFilter
    connection.EntityData.YangName = "connection"
    connection.EntityData.BundleName = "cisco_ios_xe"
    connection.EntityData.ParentYangName = "neighbor"
    connection.EntityData.SegmentPath = "connection"
    connection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    connection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    connection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    connection.EntityData.Children = make(map[string]types.YChild)
    connection.EntityData.Leafs = make(map[string]types.YLeaf)
    connection.EntityData.Leafs["state"] = types.YLeaf{"State", connection.State}
    connection.EntityData.Leafs["mode"] = types.YLeaf{"Mode", connection.Mode}
    connection.EntityData.Leafs["total-established"] = types.YLeaf{"TotalEstablished", connection.TotalEstablished}
    connection.EntityData.Leafs["total-dropped"] = types.YLeaf{"TotalDropped", connection.TotalDropped}
    connection.EntityData.Leafs["last-reset"] = types.YLeaf{"LastReset", connection.LastReset}
    connection.EntityData.Leafs["reset-reason"] = types.YLeaf{"ResetReason", connection.ResetReason}
    return &(connection.EntityData)
}

// BgpStateData_Neighbors_Neighbor_Transport
// BGP neighbor transport
type BgpStateData_Neighbors_Neighbor_Transport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indication whether path MTU discovrey is enabled. The type is bool.
    PathMtuDiscovery interface{}

    // Local TCP port used for TCP session. The type is interface{} with range:
    // 0..4294967295.
    LocalPort interface{}

    // Local address used for the TCP session. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalHost interface{}

    // Remote port used by the peer for the TCP session. The type is interface{}
    // with range: 0..4294967295.
    ForeignPort interface{}

    // Remote address of the BGP session. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForeignHost interface{}

    // Maximum Data segment size. The type is interface{} with range:
    // 0..4294967295.
    Mss interface{}
}

func (transport *BgpStateData_Neighbors_Neighbor_Transport) GetEntityData() *types.CommonEntityData {
    transport.EntityData.YFilter = transport.YFilter
    transport.EntityData.YangName = "transport"
    transport.EntityData.BundleName = "cisco_ios_xe"
    transport.EntityData.ParentYangName = "neighbor"
    transport.EntityData.SegmentPath = "transport"
    transport.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    transport.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    transport.EntityData.Children = make(map[string]types.YChild)
    transport.EntityData.Leafs = make(map[string]types.YLeaf)
    transport.EntityData.Leafs["path-mtu-discovery"] = types.YLeaf{"PathMtuDiscovery", transport.PathMtuDiscovery}
    transport.EntityData.Leafs["local-port"] = types.YLeaf{"LocalPort", transport.LocalPort}
    transport.EntityData.Leafs["local-host"] = types.YLeaf{"LocalHost", transport.LocalHost}
    transport.EntityData.Leafs["foreign-port"] = types.YLeaf{"ForeignPort", transport.ForeignPort}
    transport.EntityData.Leafs["foreign-host"] = types.YLeaf{"ForeignHost", transport.ForeignHost}
    transport.EntityData.Leafs["mss"] = types.YLeaf{"Mss", transport.Mss}
    return &(transport.EntityData)
}

// BgpStateData_Neighbors_Neighbor_PrefixActivity
// BGP neighbor activity
type BgpStateData_Neighbors_Neighbor_PrefixActivity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of prefixes sent.
    Sent BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent

    // Number of prefixes received.
    Received BgpStateData_Neighbors_Neighbor_PrefixActivity_Received
}

func (prefixActivity *BgpStateData_Neighbors_Neighbor_PrefixActivity) GetEntityData() *types.CommonEntityData {
    prefixActivity.EntityData.YFilter = prefixActivity.YFilter
    prefixActivity.EntityData.YangName = "prefix-activity"
    prefixActivity.EntityData.BundleName = "cisco_ios_xe"
    prefixActivity.EntityData.ParentYangName = "neighbor"
    prefixActivity.EntityData.SegmentPath = "prefix-activity"
    prefixActivity.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefixActivity.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefixActivity.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefixActivity.EntityData.Children = make(map[string]types.YChild)
    prefixActivity.EntityData.Children["sent"] = types.YChild{"Sent", &prefixActivity.Sent}
    prefixActivity.EntityData.Children["received"] = types.YChild{"Received", &prefixActivity.Received}
    prefixActivity.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixActivity.EntityData)
}

// BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent
// Number of prefixes sent
type BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The current number of accepted prefixes. The type is interface{} with
    // range: 0..18446744073709551615.
    CurrentPrefixes interface{}

    // The total number of accepted prefixes. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalPrefixes interface{}

    // The number of times a prefix has been withdrawn and readvertised. The type
    // is interface{} with range: 0..18446744073709551615.
    ImplicitWithdraw interface{}

    // The number of times a prefix has been withdrawn because it is no longer
    // feasible. The type is interface{} with range: 0..18446744073709551615.
    ExplicitWithdraw interface{}

    // The number of received prefixes installed as best paths. The type is
    // interface{} with range: 0..18446744073709551615.
    Bestpaths interface{}

    // The number of received prefixes installed as multipaths. The type is
    // interface{} with range: 0..18446744073709551615.
    Multipaths interface{}
}

func (sent *BgpStateData_Neighbors_Neighbor_PrefixActivity_Sent) GetEntityData() *types.CommonEntityData {
    sent.EntityData.YFilter = sent.YFilter
    sent.EntityData.YangName = "sent"
    sent.EntityData.BundleName = "cisco_ios_xe"
    sent.EntityData.ParentYangName = "prefix-activity"
    sent.EntityData.SegmentPath = "sent"
    sent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sent.EntityData.Children = make(map[string]types.YChild)
    sent.EntityData.Leafs = make(map[string]types.YLeaf)
    sent.EntityData.Leafs["current-prefixes"] = types.YLeaf{"CurrentPrefixes", sent.CurrentPrefixes}
    sent.EntityData.Leafs["total-prefixes"] = types.YLeaf{"TotalPrefixes", sent.TotalPrefixes}
    sent.EntityData.Leafs["implicit-withdraw"] = types.YLeaf{"ImplicitWithdraw", sent.ImplicitWithdraw}
    sent.EntityData.Leafs["explicit-withdraw"] = types.YLeaf{"ExplicitWithdraw", sent.ExplicitWithdraw}
    sent.EntityData.Leafs["bestpaths"] = types.YLeaf{"Bestpaths", sent.Bestpaths}
    sent.EntityData.Leafs["multipaths"] = types.YLeaf{"Multipaths", sent.Multipaths}
    return &(sent.EntityData)
}

// BgpStateData_Neighbors_Neighbor_PrefixActivity_Received
// Number of prefixes received
type BgpStateData_Neighbors_Neighbor_PrefixActivity_Received struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The current number of accepted prefixes. The type is interface{} with
    // range: 0..18446744073709551615.
    CurrentPrefixes interface{}

    // The total number of accepted prefixes. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalPrefixes interface{}

    // The number of times a prefix has been withdrawn and readvertised. The type
    // is interface{} with range: 0..18446744073709551615.
    ImplicitWithdraw interface{}

    // The number of times a prefix has been withdrawn because it is no longer
    // feasible. The type is interface{} with range: 0..18446744073709551615.
    ExplicitWithdraw interface{}

    // The number of received prefixes installed as best paths. The type is
    // interface{} with range: 0..18446744073709551615.
    Bestpaths interface{}

    // The number of received prefixes installed as multipaths. The type is
    // interface{} with range: 0..18446744073709551615.
    Multipaths interface{}
}

func (received *BgpStateData_Neighbors_Neighbor_PrefixActivity_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "cisco_ios_xe"
    received.EntityData.ParentYangName = "prefix-activity"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    received.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    received.EntityData.Children = make(map[string]types.YChild)
    received.EntityData.Leafs = make(map[string]types.YLeaf)
    received.EntityData.Leafs["current-prefixes"] = types.YLeaf{"CurrentPrefixes", received.CurrentPrefixes}
    received.EntityData.Leafs["total-prefixes"] = types.YLeaf{"TotalPrefixes", received.TotalPrefixes}
    received.EntityData.Leafs["implicit-withdraw"] = types.YLeaf{"ImplicitWithdraw", received.ImplicitWithdraw}
    received.EntityData.Leafs["explicit-withdraw"] = types.YLeaf{"ExplicitWithdraw", received.ExplicitWithdraw}
    received.EntityData.Leafs["bestpaths"] = types.YLeaf{"Bestpaths", received.Bestpaths}
    received.EntityData.Leafs["multipaths"] = types.YLeaf{"Multipaths", received.Multipaths}
    return &(received.EntityData)
}

// BgpStateData_AddressFamilies
// BGP address family
type BgpStateData_AddressFamilies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP address families. The type is slice of
    // BgpStateData_AddressFamilies_AddressFamily.
    AddressFamily []BgpStateData_AddressFamilies_AddressFamily
}

func (addressFamilies *BgpStateData_AddressFamilies) GetEntityData() *types.CommonEntityData {
    addressFamilies.EntityData.YFilter = addressFamilies.YFilter
    addressFamilies.EntityData.YangName = "address-families"
    addressFamilies.EntityData.BundleName = "cisco_ios_xe"
    addressFamilies.EntityData.ParentYangName = "bgp-state-data"
    addressFamilies.EntityData.SegmentPath = "address-families"
    addressFamilies.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressFamilies.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressFamilies.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressFamilies.EntityData.Children = make(map[string]types.YChild)
    addressFamilies.EntityData.Children["address-family"] = types.YChild{"AddressFamily", nil}
    for i := range addressFamilies.AddressFamily {
        addressFamilies.EntityData.Children[types.GetSegmentPath(&addressFamilies.AddressFamily[i])] = types.YChild{"AddressFamily", &addressFamilies.AddressFamily[i]}
    }
    addressFamilies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilies.EntityData)
}

// BgpStateData_AddressFamilies_AddressFamily
// List of BGP address families
type BgpStateData_AddressFamilies_AddressFamily struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Afi-safi value. The type is AfiSafi.
    AfiSafi interface{}

    // This attribute is a key. VRF name. The type is string.
    VrfName interface{}

    // Router ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // BGP table version number. The type is interface{} with range:
    // 0..18446744073709551615.
    BgpTableVersion interface{}

    // Routing table version number. The type is interface{} with range:
    // 0..18446744073709551615.
    RoutingTableVersion interface{}

    // Total memory in use. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalMemory interface{}

    // Local AS number. The type is interface{} with range: 0..4294967295.
    LocalAs interface{}

    // Prefix entry statistics.
    Prefixes BgpStateData_AddressFamilies_AddressFamily_Prefixes

    // Path entry statistics.
    Path BgpStateData_AddressFamilies_AddressFamily_Path

    // AS path entry statistics.
    AsPath BgpStateData_AddressFamilies_AddressFamily_AsPath

    // Route map entry statistics.
    RouteMap BgpStateData_AddressFamilies_AddressFamily_RouteMap

    // Filter list entry statistics.
    FilterList BgpStateData_AddressFamilies_AddressFamily_FilterList

    // BGP activity information.
    Activities BgpStateData_AddressFamilies_AddressFamily_Activities

    // Neighbor summary.
    BgpNeighborSummaries BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries
}

func (addressFamily *BgpStateData_AddressFamilies_AddressFamily) GetEntityData() *types.CommonEntityData {
    addressFamily.EntityData.YFilter = addressFamily.YFilter
    addressFamily.EntityData.YangName = "address-family"
    addressFamily.EntityData.BundleName = "cisco_ios_xe"
    addressFamily.EntityData.ParentYangName = "address-families"
    addressFamily.EntityData.SegmentPath = "address-family" + "[afi-safi='" + fmt.Sprintf("%v", addressFamily.AfiSafi) + "']" + "[vrf-name='" + fmt.Sprintf("%v", addressFamily.VrfName) + "']"
    addressFamily.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressFamily.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressFamily.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressFamily.EntityData.Children = make(map[string]types.YChild)
    addressFamily.EntityData.Children["prefixes"] = types.YChild{"Prefixes", &addressFamily.Prefixes}
    addressFamily.EntityData.Children["path"] = types.YChild{"Path", &addressFamily.Path}
    addressFamily.EntityData.Children["as-path"] = types.YChild{"AsPath", &addressFamily.AsPath}
    addressFamily.EntityData.Children["route-map"] = types.YChild{"RouteMap", &addressFamily.RouteMap}
    addressFamily.EntityData.Children["filter-list"] = types.YChild{"FilterList", &addressFamily.FilterList}
    addressFamily.EntityData.Children["activities"] = types.YChild{"Activities", &addressFamily.Activities}
    addressFamily.EntityData.Children["bgp-neighbor-summaries"] = types.YChild{"BgpNeighborSummaries", &addressFamily.BgpNeighborSummaries}
    addressFamily.EntityData.Leafs = make(map[string]types.YLeaf)
    addressFamily.EntityData.Leafs["afi-safi"] = types.YLeaf{"AfiSafi", addressFamily.AfiSafi}
    addressFamily.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", addressFamily.VrfName}
    addressFamily.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", addressFamily.RouterId}
    addressFamily.EntityData.Leafs["bgp-table-version"] = types.YLeaf{"BgpTableVersion", addressFamily.BgpTableVersion}
    addressFamily.EntityData.Leafs["routing-table-version"] = types.YLeaf{"RoutingTableVersion", addressFamily.RoutingTableVersion}
    addressFamily.EntityData.Leafs["total-memory"] = types.YLeaf{"TotalMemory", addressFamily.TotalMemory}
    addressFamily.EntityData.Leafs["local-as"] = types.YLeaf{"LocalAs", addressFamily.LocalAs}
    return &(addressFamily.EntityData)
}

// BgpStateData_AddressFamilies_AddressFamily_Prefixes
// Prefix entry statistics
type BgpStateData_AddressFamilies_AddressFamily_Prefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of prefix entries. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEntries interface{}

    // Total memory usage in byte. The type is interface{} with range:
    // 0..18446744073709551615.
    MemoryUsage interface{}
}

func (prefixes *BgpStateData_AddressFamilies_AddressFamily_Prefixes) GetEntityData() *types.CommonEntityData {
    prefixes.EntityData.YFilter = prefixes.YFilter
    prefixes.EntityData.YangName = "prefixes"
    prefixes.EntityData.BundleName = "cisco_ios_xe"
    prefixes.EntityData.ParentYangName = "address-family"
    prefixes.EntityData.SegmentPath = "prefixes"
    prefixes.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefixes.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefixes.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefixes.EntityData.Children = make(map[string]types.YChild)
    prefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixes.EntityData.Leafs["total-entries"] = types.YLeaf{"TotalEntries", prefixes.TotalEntries}
    prefixes.EntityData.Leafs["memory-usage"] = types.YLeaf{"MemoryUsage", prefixes.MemoryUsage}
    return &(prefixes.EntityData)
}

// BgpStateData_AddressFamilies_AddressFamily_Path
// Path entry statistics
type BgpStateData_AddressFamilies_AddressFamily_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of prefix entries. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEntries interface{}

    // Total memory usage in byte. The type is interface{} with range:
    // 0..18446744073709551615.
    MemoryUsage interface{}
}

func (path *BgpStateData_AddressFamilies_AddressFamily_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xe"
    path.EntityData.ParentYangName = "address-family"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["total-entries"] = types.YLeaf{"TotalEntries", path.TotalEntries}
    path.EntityData.Leafs["memory-usage"] = types.YLeaf{"MemoryUsage", path.MemoryUsage}
    return &(path.EntityData)
}

// BgpStateData_AddressFamilies_AddressFamily_AsPath
// AS path entry statistics
type BgpStateData_AddressFamilies_AddressFamily_AsPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of prefix entries. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEntries interface{}

    // Total memory usage in byte. The type is interface{} with range:
    // 0..18446744073709551615.
    MemoryUsage interface{}
}

func (asPath *BgpStateData_AddressFamilies_AddressFamily_AsPath) GetEntityData() *types.CommonEntityData {
    asPath.EntityData.YFilter = asPath.YFilter
    asPath.EntityData.YangName = "as-path"
    asPath.EntityData.BundleName = "cisco_ios_xe"
    asPath.EntityData.ParentYangName = "address-family"
    asPath.EntityData.SegmentPath = "as-path"
    asPath.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    asPath.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    asPath.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    asPath.EntityData.Children = make(map[string]types.YChild)
    asPath.EntityData.Leafs = make(map[string]types.YLeaf)
    asPath.EntityData.Leafs["total-entries"] = types.YLeaf{"TotalEntries", asPath.TotalEntries}
    asPath.EntityData.Leafs["memory-usage"] = types.YLeaf{"MemoryUsage", asPath.MemoryUsage}
    return &(asPath.EntityData)
}

// BgpStateData_AddressFamilies_AddressFamily_RouteMap
// Route map entry statistics
type BgpStateData_AddressFamilies_AddressFamily_RouteMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of prefix entries. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEntries interface{}

    // Total memory usage in byte. The type is interface{} with range:
    // 0..18446744073709551615.
    MemoryUsage interface{}
}

func (routeMap *BgpStateData_AddressFamilies_AddressFamily_RouteMap) GetEntityData() *types.CommonEntityData {
    routeMap.EntityData.YFilter = routeMap.YFilter
    routeMap.EntityData.YangName = "route-map"
    routeMap.EntityData.BundleName = "cisco_ios_xe"
    routeMap.EntityData.ParentYangName = "address-family"
    routeMap.EntityData.SegmentPath = "route-map"
    routeMap.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    routeMap.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    routeMap.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    routeMap.EntityData.Children = make(map[string]types.YChild)
    routeMap.EntityData.Leafs = make(map[string]types.YLeaf)
    routeMap.EntityData.Leafs["total-entries"] = types.YLeaf{"TotalEntries", routeMap.TotalEntries}
    routeMap.EntityData.Leafs["memory-usage"] = types.YLeaf{"MemoryUsage", routeMap.MemoryUsage}
    return &(routeMap.EntityData)
}

// BgpStateData_AddressFamilies_AddressFamily_FilterList
// Filter list entry statistics
type BgpStateData_AddressFamilies_AddressFamily_FilterList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of prefix entries. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalEntries interface{}

    // Total memory usage in byte. The type is interface{} with range:
    // 0..18446744073709551615.
    MemoryUsage interface{}
}

func (filterList *BgpStateData_AddressFamilies_AddressFamily_FilterList) GetEntityData() *types.CommonEntityData {
    filterList.EntityData.YFilter = filterList.YFilter
    filterList.EntityData.YangName = "filter-list"
    filterList.EntityData.BundleName = "cisco_ios_xe"
    filterList.EntityData.ParentYangName = "address-family"
    filterList.EntityData.SegmentPath = "filter-list"
    filterList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    filterList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    filterList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    filterList.EntityData.Children = make(map[string]types.YChild)
    filterList.EntityData.Leafs = make(map[string]types.YLeaf)
    filterList.EntityData.Leafs["total-entries"] = types.YLeaf{"TotalEntries", filterList.TotalEntries}
    filterList.EntityData.Leafs["memory-usage"] = types.YLeaf{"MemoryUsage", filterList.MemoryUsage}
    return &(filterList.EntityData)
}

// BgpStateData_AddressFamilies_AddressFamily_Activities
// BGP activity information
type BgpStateData_AddressFamilies_AddressFamily_Activities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of prefixes. The type is interface{} with range:
    // 0..18446744073709551615.
    Prefixes interface{}

    // Total number of paths. The type is interface{} with range:
    // 0..18446744073709551615.
    Paths interface{}

    // Scan interval in seconds. The type is string.
    ScanInterval interface{}
}

func (activities *BgpStateData_AddressFamilies_AddressFamily_Activities) GetEntityData() *types.CommonEntityData {
    activities.EntityData.YFilter = activities.YFilter
    activities.EntityData.YangName = "activities"
    activities.EntityData.BundleName = "cisco_ios_xe"
    activities.EntityData.ParentYangName = "address-family"
    activities.EntityData.SegmentPath = "activities"
    activities.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    activities.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    activities.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    activities.EntityData.Children = make(map[string]types.YChild)
    activities.EntityData.Leafs = make(map[string]types.YLeaf)
    activities.EntityData.Leafs["prefixes"] = types.YLeaf{"Prefixes", activities.Prefixes}
    activities.EntityData.Leafs["paths"] = types.YLeaf{"Paths", activities.Paths}
    activities.EntityData.Leafs["scan-interval"] = types.YLeaf{"ScanInterval", activities.ScanInterval}
    return &(activities.EntityData)
}

// BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries
// Neighbor summary
type BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of neighbor summaries. The type is slice of
    // BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary.
    BgpNeighborSummary []BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary
}

func (bgpNeighborSummaries *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries) GetEntityData() *types.CommonEntityData {
    bgpNeighborSummaries.EntityData.YFilter = bgpNeighborSummaries.YFilter
    bgpNeighborSummaries.EntityData.YangName = "bgp-neighbor-summaries"
    bgpNeighborSummaries.EntityData.BundleName = "cisco_ios_xe"
    bgpNeighborSummaries.EntityData.ParentYangName = "address-family"
    bgpNeighborSummaries.EntityData.SegmentPath = "bgp-neighbor-summaries"
    bgpNeighborSummaries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpNeighborSummaries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpNeighborSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpNeighborSummaries.EntityData.Children = make(map[string]types.YChild)
    bgpNeighborSummaries.EntityData.Children["bgp-neighbor-summary"] = types.YChild{"BgpNeighborSummary", nil}
    for i := range bgpNeighborSummaries.BgpNeighborSummary {
        bgpNeighborSummaries.EntityData.Children[types.GetSegmentPath(&bgpNeighborSummaries.BgpNeighborSummary[i])] = types.YChild{"BgpNeighborSummary", &bgpNeighborSummaries.BgpNeighborSummary[i]}
    }
    bgpNeighborSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpNeighborSummaries.EntityData)
}

// BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary
// List of neighbor summaries
type BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor address. The type is string.
    Id interface{}

    // BGP protocol version. The type is interface{} with range: 0..4294967295.
    BgpVersion interface{}

    // Number of messages received from this neighbor. The type is interface{}
    // with range: 0..18446744073709551615.
    MessagesReceived interface{}

    // Number of messages sent to this neighbor. The type is interface{} with
    // range: 0..18446744073709551615.
    MessagesSent interface{}

    // BGP table version. The type is interface{} with range:
    // 0..18446744073709551615.
    TableVersion interface{}

    // Number of messages in input queue. The type is interface{} with range:
    // 0..18446744073709551615.
    InputQueue interface{}

    // Number of messages in output queue. The type is interface{} with range:
    // 0..18446744073709551615.
    OutputQueue interface{}

    // Neighbor session uptime. The type is string.
    UpTime interface{}

    // BGP session state. The type is BgpFsmState.
    State interface{}

    // Number of prefixes received from the neighbor. The type is interface{} with
    // range: 0..18446744073709551615.
    PrefixesReceived interface{}

    // Indication of whether the neighbor was dyanmically configured. The type is
    // bool.
    DynamicallyConfigured interface{}

    // BGP neighbor AS number. The type is interface{} with range: 0..4294967295.
    As interface{}
}

func (bgpNeighborSummary *BgpStateData_AddressFamilies_AddressFamily_BgpNeighborSummaries_BgpNeighborSummary) GetEntityData() *types.CommonEntityData {
    bgpNeighborSummary.EntityData.YFilter = bgpNeighborSummary.YFilter
    bgpNeighborSummary.EntityData.YangName = "bgp-neighbor-summary"
    bgpNeighborSummary.EntityData.BundleName = "cisco_ios_xe"
    bgpNeighborSummary.EntityData.ParentYangName = "bgp-neighbor-summaries"
    bgpNeighborSummary.EntityData.SegmentPath = "bgp-neighbor-summary" + "[id='" + fmt.Sprintf("%v", bgpNeighborSummary.Id) + "']"
    bgpNeighborSummary.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpNeighborSummary.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpNeighborSummary.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpNeighborSummary.EntityData.Children = make(map[string]types.YChild)
    bgpNeighborSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpNeighborSummary.EntityData.Leafs["id"] = types.YLeaf{"Id", bgpNeighborSummary.Id}
    bgpNeighborSummary.EntityData.Leafs["bgp-version"] = types.YLeaf{"BgpVersion", bgpNeighborSummary.BgpVersion}
    bgpNeighborSummary.EntityData.Leafs["messages-received"] = types.YLeaf{"MessagesReceived", bgpNeighborSummary.MessagesReceived}
    bgpNeighborSummary.EntityData.Leafs["messages-sent"] = types.YLeaf{"MessagesSent", bgpNeighborSummary.MessagesSent}
    bgpNeighborSummary.EntityData.Leafs["table-version"] = types.YLeaf{"TableVersion", bgpNeighborSummary.TableVersion}
    bgpNeighborSummary.EntityData.Leafs["input-queue"] = types.YLeaf{"InputQueue", bgpNeighborSummary.InputQueue}
    bgpNeighborSummary.EntityData.Leafs["output-queue"] = types.YLeaf{"OutputQueue", bgpNeighborSummary.OutputQueue}
    bgpNeighborSummary.EntityData.Leafs["up-time"] = types.YLeaf{"UpTime", bgpNeighborSummary.UpTime}
    bgpNeighborSummary.EntityData.Leafs["state"] = types.YLeaf{"State", bgpNeighborSummary.State}
    bgpNeighborSummary.EntityData.Leafs["prefixes-received"] = types.YLeaf{"PrefixesReceived", bgpNeighborSummary.PrefixesReceived}
    bgpNeighborSummary.EntityData.Leafs["dynamically-configured"] = types.YLeaf{"DynamicallyConfigured", bgpNeighborSummary.DynamicallyConfigured}
    bgpNeighborSummary.EntityData.Leafs["as"] = types.YLeaf{"As", bgpNeighborSummary.As}
    return &(bgpNeighborSummary.EntityData)
}

// BgpStateData_BgpRouteVrfs
// BGP VRFs
type BgpStateData_BgpRouteVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP VRFs. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf.
    BgpRouteVrf []BgpStateData_BgpRouteVrfs_BgpRouteVrf
}

func (bgpRouteVrfs *BgpStateData_BgpRouteVrfs) GetEntityData() *types.CommonEntityData {
    bgpRouteVrfs.EntityData.YFilter = bgpRouteVrfs.YFilter
    bgpRouteVrfs.EntityData.YangName = "bgp-route-vrfs"
    bgpRouteVrfs.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteVrfs.EntityData.ParentYangName = "bgp-state-data"
    bgpRouteVrfs.EntityData.SegmentPath = "bgp-route-vrfs"
    bgpRouteVrfs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteVrfs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteVrfs.EntityData.Children = make(map[string]types.YChild)
    bgpRouteVrfs.EntityData.Children["bgp-route-vrf"] = types.YChild{"BgpRouteVrf", nil}
    for i := range bgpRouteVrfs.BgpRouteVrf {
        bgpRouteVrfs.EntityData.Children[types.GetSegmentPath(&bgpRouteVrfs.BgpRouteVrf[i])] = types.YChild{"BgpRouteVrf", &bgpRouteVrfs.BgpRouteVrf[i]}
    }
    bgpRouteVrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRouteVrfs.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf
// List of BGP VRFs
type BgpStateData_BgpRouteVrfs_BgpRouteVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP vrf. The type is string.
    Vrf interface{}

    // BGP address families.
    BgpRouteAfs BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs
}

func (bgpRouteVrf *BgpStateData_BgpRouteVrfs_BgpRouteVrf) GetEntityData() *types.CommonEntityData {
    bgpRouteVrf.EntityData.YFilter = bgpRouteVrf.YFilter
    bgpRouteVrf.EntityData.YangName = "bgp-route-vrf"
    bgpRouteVrf.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteVrf.EntityData.ParentYangName = "bgp-route-vrfs"
    bgpRouteVrf.EntityData.SegmentPath = "bgp-route-vrf" + "[vrf='" + fmt.Sprintf("%v", bgpRouteVrf.Vrf) + "']"
    bgpRouteVrf.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteVrf.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteVrf.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteVrf.EntityData.Children = make(map[string]types.YChild)
    bgpRouteVrf.EntityData.Children["bgp-route-afs"] = types.YChild{"BgpRouteAfs", &bgpRouteVrf.BgpRouteAfs}
    bgpRouteVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRouteVrf.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", bgpRouteVrf.Vrf}
    return &(bgpRouteVrf.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs
// BGP address families
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP address families. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf.
    BgpRouteAf []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf
}

func (bgpRouteAfs *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs) GetEntityData() *types.CommonEntityData {
    bgpRouteAfs.EntityData.YFilter = bgpRouteAfs.YFilter
    bgpRouteAfs.EntityData.YangName = "bgp-route-afs"
    bgpRouteAfs.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteAfs.EntityData.ParentYangName = "bgp-route-vrf"
    bgpRouteAfs.EntityData.SegmentPath = "bgp-route-afs"
    bgpRouteAfs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteAfs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteAfs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteAfs.EntityData.Children = make(map[string]types.YChild)
    bgpRouteAfs.EntityData.Children["bgp-route-af"] = types.YChild{"BgpRouteAf", nil}
    for i := range bgpRouteAfs.BgpRouteAf {
        bgpRouteAfs.EntityData.Children[types.GetSegmentPath(&bgpRouteAfs.BgpRouteAf[i])] = types.YChild{"BgpRouteAf", &bgpRouteAfs.BgpRouteAf[i]}
    }
    bgpRouteAfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRouteAfs.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf
// List of BGP address families
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP address family. The type is AfiSafi.
    AfiSafi interface{}

    // BGP route filters.
    BgpRouteFilters BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters

    // BGP route neighbors.
    BgpRouteNeighbors BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors

    // BGP peer groups.
    BgpPeerGroups BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpPeerGroups
}

func (bgpRouteAf *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf) GetEntityData() *types.CommonEntityData {
    bgpRouteAf.EntityData.YFilter = bgpRouteAf.YFilter
    bgpRouteAf.EntityData.YangName = "bgp-route-af"
    bgpRouteAf.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteAf.EntityData.ParentYangName = "bgp-route-afs"
    bgpRouteAf.EntityData.SegmentPath = "bgp-route-af" + "[afi-safi='" + fmt.Sprintf("%v", bgpRouteAf.AfiSafi) + "']"
    bgpRouteAf.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteAf.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteAf.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteAf.EntityData.Children = make(map[string]types.YChild)
    bgpRouteAf.EntityData.Children["bgp-route-filters"] = types.YChild{"BgpRouteFilters", &bgpRouteAf.BgpRouteFilters}
    bgpRouteAf.EntityData.Children["bgp-route-neighbors"] = types.YChild{"BgpRouteNeighbors", &bgpRouteAf.BgpRouteNeighbors}
    bgpRouteAf.EntityData.Children["bgp-peer-groups"] = types.YChild{"BgpPeerGroups", &bgpRouteAf.BgpPeerGroups}
    bgpRouteAf.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRouteAf.EntityData.Leafs["afi-safi"] = types.YLeaf{"AfiSafi", bgpRouteAf.AfiSafi}
    return &(bgpRouteAf.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters
// BGP route filters
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP route filters. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter.
    BgpRouteFilter []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter
}

func (bgpRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters) GetEntityData() *types.CommonEntityData {
    bgpRouteFilters.EntityData.YFilter = bgpRouteFilters.YFilter
    bgpRouteFilters.EntityData.YangName = "bgp-route-filters"
    bgpRouteFilters.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteFilters.EntityData.ParentYangName = "bgp-route-af"
    bgpRouteFilters.EntityData.SegmentPath = "bgp-route-filters"
    bgpRouteFilters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteFilters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteFilters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteFilters.EntityData.Children = make(map[string]types.YChild)
    bgpRouteFilters.EntityData.Children["bgp-route-filter"] = types.YChild{"BgpRouteFilter", nil}
    for i := range bgpRouteFilters.BgpRouteFilter {
        bgpRouteFilters.EntityData.Children[types.GetSegmentPath(&bgpRouteFilters.BgpRouteFilter[i])] = types.YChild{"BgpRouteFilter", &bgpRouteFilters.BgpRouteFilter[i]}
    }
    bgpRouteFilters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRouteFilters.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter
// List of BGP route filters
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP route filter. The type is BgpRouteFilters.
    RouteFilter interface{}

    // BGP route entries.
    BgpRouteEntries BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries
}

func (bgpRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter) GetEntityData() *types.CommonEntityData {
    bgpRouteFilter.EntityData.YFilter = bgpRouteFilter.YFilter
    bgpRouteFilter.EntityData.YangName = "bgp-route-filter"
    bgpRouteFilter.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteFilter.EntityData.ParentYangName = "bgp-route-filters"
    bgpRouteFilter.EntityData.SegmentPath = "bgp-route-filter" + "[route-filter='" + fmt.Sprintf("%v", bgpRouteFilter.RouteFilter) + "']"
    bgpRouteFilter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteFilter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteFilter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteFilter.EntityData.Children = make(map[string]types.YChild)
    bgpRouteFilter.EntityData.Children["bgp-route-entries"] = types.YChild{"BgpRouteEntries", &bgpRouteFilter.BgpRouteEntries}
    bgpRouteFilter.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRouteFilter.EntityData.Leafs["route-filter"] = types.YLeaf{"RouteFilter", bgpRouteFilter.RouteFilter}
    return &(bgpRouteFilter.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries
// BGP route entries
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP route entries. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry.
    BgpRouteEntry []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry
}

func (bgpRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries) GetEntityData() *types.CommonEntityData {
    bgpRouteEntries.EntityData.YFilter = bgpRouteEntries.YFilter
    bgpRouteEntries.EntityData.YangName = "bgp-route-entries"
    bgpRouteEntries.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteEntries.EntityData.ParentYangName = "bgp-route-filter"
    bgpRouteEntries.EntityData.SegmentPath = "bgp-route-entries"
    bgpRouteEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteEntries.EntityData.Children = make(map[string]types.YChild)
    bgpRouteEntries.EntityData.Children["bgp-route-entry"] = types.YChild{"BgpRouteEntry", nil}
    for i := range bgpRouteEntries.BgpRouteEntry {
        bgpRouteEntries.EntityData.Children[types.GetSegmentPath(&bgpRouteEntries.BgpRouteEntry[i])] = types.YChild{"BgpRouteEntry", &bgpRouteEntries.BgpRouteEntry[i]}
    }
    bgpRouteEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRouteEntries.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry
// List of BGP route entries
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Routing table entry prefix. The type is string.
    Prefix interface{}

    // Routing table version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Number of paths available for BGP prefix. The type is interface{} with
    // range: 0..4294967295.
    AvailablePaths interface{}

    // Whom is thie prefix advertized to. The type is string.
    AdvertisedTo interface{}

    // Prefix next hop details.
    BgpPathEntries BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries
}

func (bgpRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry) GetEntityData() *types.CommonEntityData {
    bgpRouteEntry.EntityData.YFilter = bgpRouteEntry.YFilter
    bgpRouteEntry.EntityData.YangName = "bgp-route-entry"
    bgpRouteEntry.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteEntry.EntityData.ParentYangName = "bgp-route-entries"
    bgpRouteEntry.EntityData.SegmentPath = "bgp-route-entry" + "[prefix='" + fmt.Sprintf("%v", bgpRouteEntry.Prefix) + "']"
    bgpRouteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteEntry.EntityData.Children = make(map[string]types.YChild)
    bgpRouteEntry.EntityData.Children["bgp-path-entries"] = types.YChild{"BgpPathEntries", &bgpRouteEntry.BgpPathEntries}
    bgpRouteEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRouteEntry.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", bgpRouteEntry.Prefix}
    bgpRouteEntry.EntityData.Leafs["version"] = types.YLeaf{"Version", bgpRouteEntry.Version}
    bgpRouteEntry.EntityData.Leafs["available-paths"] = types.YLeaf{"AvailablePaths", bgpRouteEntry.AvailablePaths}
    bgpRouteEntry.EntityData.Leafs["advertised-to"] = types.YLeaf{"AdvertisedTo", bgpRouteEntry.AdvertisedTo}
    return &(bgpRouteEntry.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries
// Prefix next hop details
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of prefix next hop details. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry.
    BgpPathEntry []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry
}

func (bgpPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries) GetEntityData() *types.CommonEntityData {
    bgpPathEntries.EntityData.YFilter = bgpPathEntries.YFilter
    bgpPathEntries.EntityData.YangName = "bgp-path-entries"
    bgpPathEntries.EntityData.BundleName = "cisco_ios_xe"
    bgpPathEntries.EntityData.ParentYangName = "bgp-route-entry"
    bgpPathEntries.EntityData.SegmentPath = "bgp-path-entries"
    bgpPathEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpPathEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpPathEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpPathEntries.EntityData.Children = make(map[string]types.YChild)
    bgpPathEntries.EntityData.Children["bgp-path-entry"] = types.YChild{"BgpPathEntry", nil}
    for i := range bgpPathEntries.BgpPathEntry {
        bgpPathEntries.EntityData.Children[types.GetSegmentPath(&bgpPathEntries.BgpPathEntry[i])] = types.YChild{"BgpPathEntry", &bgpPathEntries.BgpPathEntry[i]}
    }
    bgpPathEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpPathEntries.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry
// List of prefix next hop details
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Next hop for this path. The type is string.
    Nexthop interface{}

    // Metric associated with the path. The type is interface{} with range:
    // 0..4294967295.
    Metric interface{}

    // Local preference for this path. The type is interface{} with range:
    // 0..4294967295.
    LocalPref interface{}

    // Path weight. The type is interface{} with range: 0..4294967295.
    Weight interface{}

    // AS path. The type is string.
    AsPath interface{}

    // Path origin. The type is BgpOriginCode.
    Origin interface{}

    // RPKI path status. The type is BgpRpkiStatus.
    RpkiStatus interface{}

    // Community label for the path. The type is string.
    Community interface{}

    // MPLS label in for the path. The type is string.
    MplsIn interface{}

    // MPLS label out for the path. The type is string.
    MplsOut interface{}

    // SR profile name for the path. The type is string.
    SrProfileName interface{}

    // SR binding sid for the path. The type is interface{} with range:
    // 0..4294967295.
    SrBindingSid interface{}

    // SR label index for the path. The type is interface{} with range:
    // 0..4294967295.
    SrLabelIndx interface{}

    // path using 4-octet AS numbers. The type is string.
    As4Path interface{}

    // attribute indicating whether or not the prefix is an atomic aggregate. The
    // type is bool.
    AtomicAggregate interface{}

    // AS number of autonomous system them performed the aggregation. The type is
    // interface{} with range: 0..4294967295.
    AggrAsNumber interface{}

    // AS4 number of autonomous system them performed the aggregation. The type is
    // interface{} with range: 0..4294967295.
    AggrAs4Number interface{}

    // IP address of the router that performed the aggregation. The type is
    // string.
    AggrAddress interface{}

    // the router ID of the originator of the route in the local AS. The type is
    // string.
    OriginatorId interface{}

    // the reflection path the route has passed. The type is string.
    ClusterList interface{}

    // BGP extended community attribute. The type is string.
    ExtendedCommunity interface{}

    // the accumulated IGP metric for the path. The type is interface{} with
    // range: 0..18446744073709551615.
    ExtAigpMetric interface{}

    // path identifier used to uniquely identify a route. The type is interface{}
    // with range: 0..4294967295.
    PathId interface{}

    // Path status.
    PathStatus BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus
}

func (bgpPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry) GetEntityData() *types.CommonEntityData {
    bgpPathEntry.EntityData.YFilter = bgpPathEntry.YFilter
    bgpPathEntry.EntityData.YangName = "bgp-path-entry"
    bgpPathEntry.EntityData.BundleName = "cisco_ios_xe"
    bgpPathEntry.EntityData.ParentYangName = "bgp-path-entries"
    bgpPathEntry.EntityData.SegmentPath = "bgp-path-entry" + "[nexthop='" + fmt.Sprintf("%v", bgpPathEntry.Nexthop) + "']"
    bgpPathEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpPathEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpPathEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpPathEntry.EntityData.Children = make(map[string]types.YChild)
    bgpPathEntry.EntityData.Children["path-status"] = types.YChild{"PathStatus", &bgpPathEntry.PathStatus}
    bgpPathEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpPathEntry.EntityData.Leafs["nexthop"] = types.YLeaf{"Nexthop", bgpPathEntry.Nexthop}
    bgpPathEntry.EntityData.Leafs["metric"] = types.YLeaf{"Metric", bgpPathEntry.Metric}
    bgpPathEntry.EntityData.Leafs["local-pref"] = types.YLeaf{"LocalPref", bgpPathEntry.LocalPref}
    bgpPathEntry.EntityData.Leafs["weight"] = types.YLeaf{"Weight", bgpPathEntry.Weight}
    bgpPathEntry.EntityData.Leafs["as-path"] = types.YLeaf{"AsPath", bgpPathEntry.AsPath}
    bgpPathEntry.EntityData.Leafs["origin"] = types.YLeaf{"Origin", bgpPathEntry.Origin}
    bgpPathEntry.EntityData.Leafs["rpki-status"] = types.YLeaf{"RpkiStatus", bgpPathEntry.RpkiStatus}
    bgpPathEntry.EntityData.Leafs["community"] = types.YLeaf{"Community", bgpPathEntry.Community}
    bgpPathEntry.EntityData.Leafs["mpls-in"] = types.YLeaf{"MplsIn", bgpPathEntry.MplsIn}
    bgpPathEntry.EntityData.Leafs["mpls-out"] = types.YLeaf{"MplsOut", bgpPathEntry.MplsOut}
    bgpPathEntry.EntityData.Leafs["sr-profile-name"] = types.YLeaf{"SrProfileName", bgpPathEntry.SrProfileName}
    bgpPathEntry.EntityData.Leafs["sr-binding-sid"] = types.YLeaf{"SrBindingSid", bgpPathEntry.SrBindingSid}
    bgpPathEntry.EntityData.Leafs["sr-label-indx"] = types.YLeaf{"SrLabelIndx", bgpPathEntry.SrLabelIndx}
    bgpPathEntry.EntityData.Leafs["as4-path"] = types.YLeaf{"As4Path", bgpPathEntry.As4Path}
    bgpPathEntry.EntityData.Leafs["atomic-aggregate"] = types.YLeaf{"AtomicAggregate", bgpPathEntry.AtomicAggregate}
    bgpPathEntry.EntityData.Leafs["aggr-as-number"] = types.YLeaf{"AggrAsNumber", bgpPathEntry.AggrAsNumber}
    bgpPathEntry.EntityData.Leafs["aggr-as4-number"] = types.YLeaf{"AggrAs4Number", bgpPathEntry.AggrAs4Number}
    bgpPathEntry.EntityData.Leafs["aggr-address"] = types.YLeaf{"AggrAddress", bgpPathEntry.AggrAddress}
    bgpPathEntry.EntityData.Leafs["originator-id"] = types.YLeaf{"OriginatorId", bgpPathEntry.OriginatorId}
    bgpPathEntry.EntityData.Leafs["cluster-list"] = types.YLeaf{"ClusterList", bgpPathEntry.ClusterList}
    bgpPathEntry.EntityData.Leafs["extended-community"] = types.YLeaf{"ExtendedCommunity", bgpPathEntry.ExtendedCommunity}
    bgpPathEntry.EntityData.Leafs["ext-aigp-metric"] = types.YLeaf{"ExtAigpMetric", bgpPathEntry.ExtAigpMetric}
    bgpPathEntry.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", bgpPathEntry.PathId}
    return &(bgpPathEntry.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus
// Path status
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed path. The type is interface{}.
    Suppressed interface{}

    // Damped path. The type is interface{}.
    Damped interface{}

    // History path. The type is interface{}.
    History interface{}

    // Valid path. The type is interface{}.
    Valid interface{}

    // Sourced path. The type is interface{}.
    Sourced interface{}

    // Best path. The type is interface{}.
    Bestpath interface{}

    // Internal path. The type is interface{}.
    Internal interface{}

    // RIB-fail path. The type is interface{}.
    RibFail interface{}

    // Stale path. The type is interface{}.
    Stale interface{}

    // Multipath path. The type is interface{}.
    Multipath interface{}

    // Backup path. The type is interface{}.
    BackupPath interface{}

    // RT filter path. The type is interface{}.
    RtFilter interface{}

    // Best externa path. The type is interface{}.
    BestExternal interface{}

    // Additional path. The type is interface{}.
    AdditionalPath interface{}

    // RIB compressed path. The type is interface{}.
    RibCompressed interface{}
}

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteFilters_BgpRouteFilter_BgpRouteEntries_BgpRouteEntry_BgpPathEntries_BgpPathEntry_PathStatus) GetEntityData() *types.CommonEntityData {
    pathStatus.EntityData.YFilter = pathStatus.YFilter
    pathStatus.EntityData.YangName = "path-status"
    pathStatus.EntityData.BundleName = "cisco_ios_xe"
    pathStatus.EntityData.ParentYangName = "bgp-path-entry"
    pathStatus.EntityData.SegmentPath = "path-status"
    pathStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pathStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pathStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pathStatus.EntityData.Children = make(map[string]types.YChild)
    pathStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    pathStatus.EntityData.Leafs["suppressed"] = types.YLeaf{"Suppressed", pathStatus.Suppressed}
    pathStatus.EntityData.Leafs["damped"] = types.YLeaf{"Damped", pathStatus.Damped}
    pathStatus.EntityData.Leafs["history"] = types.YLeaf{"History", pathStatus.History}
    pathStatus.EntityData.Leafs["valid"] = types.YLeaf{"Valid", pathStatus.Valid}
    pathStatus.EntityData.Leafs["sourced"] = types.YLeaf{"Sourced", pathStatus.Sourced}
    pathStatus.EntityData.Leafs["bestpath"] = types.YLeaf{"Bestpath", pathStatus.Bestpath}
    pathStatus.EntityData.Leafs["internal"] = types.YLeaf{"Internal", pathStatus.Internal}
    pathStatus.EntityData.Leafs["rib-fail"] = types.YLeaf{"RibFail", pathStatus.RibFail}
    pathStatus.EntityData.Leafs["stale"] = types.YLeaf{"Stale", pathStatus.Stale}
    pathStatus.EntityData.Leafs["multipath"] = types.YLeaf{"Multipath", pathStatus.Multipath}
    pathStatus.EntityData.Leafs["backup-path"] = types.YLeaf{"BackupPath", pathStatus.BackupPath}
    pathStatus.EntityData.Leafs["rt-filter"] = types.YLeaf{"RtFilter", pathStatus.RtFilter}
    pathStatus.EntityData.Leafs["best-external"] = types.YLeaf{"BestExternal", pathStatus.BestExternal}
    pathStatus.EntityData.Leafs["additional-path"] = types.YLeaf{"AdditionalPath", pathStatus.AdditionalPath}
    pathStatus.EntityData.Leafs["rib-compressed"] = types.YLeaf{"RibCompressed", pathStatus.RibCompressed}
    return &(pathStatus.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors
// BGP route neighbors
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP route neighbors. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor.
    BgpRouteNeighbor []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor
}

func (bgpRouteNeighbors *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors) GetEntityData() *types.CommonEntityData {
    bgpRouteNeighbors.EntityData.YFilter = bgpRouteNeighbors.YFilter
    bgpRouteNeighbors.EntityData.YangName = "bgp-route-neighbors"
    bgpRouteNeighbors.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteNeighbors.EntityData.ParentYangName = "bgp-route-af"
    bgpRouteNeighbors.EntityData.SegmentPath = "bgp-route-neighbors"
    bgpRouteNeighbors.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteNeighbors.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteNeighbors.EntityData.Children = make(map[string]types.YChild)
    bgpRouteNeighbors.EntityData.Children["bgp-route-neighbor"] = types.YChild{"BgpRouteNeighbor", nil}
    for i := range bgpRouteNeighbors.BgpRouteNeighbor {
        bgpRouteNeighbors.EntityData.Children[types.GetSegmentPath(&bgpRouteNeighbors.BgpRouteNeighbor[i])] = types.YChild{"BgpRouteNeighbor", &bgpRouteNeighbors.BgpRouteNeighbor[i]}
    }
    bgpRouteNeighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRouteNeighbors.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor
// List of BGP route neighbors
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP neighbor ID. The type is string.
    NbrId interface{}

    // BGP neighbor route filters.
    BgpNeighborRouteFilters BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters
}

func (bgpRouteNeighbor *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor) GetEntityData() *types.CommonEntityData {
    bgpRouteNeighbor.EntityData.YFilter = bgpRouteNeighbor.YFilter
    bgpRouteNeighbor.EntityData.YangName = "bgp-route-neighbor"
    bgpRouteNeighbor.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteNeighbor.EntityData.ParentYangName = "bgp-route-neighbors"
    bgpRouteNeighbor.EntityData.SegmentPath = "bgp-route-neighbor" + "[nbr-id='" + fmt.Sprintf("%v", bgpRouteNeighbor.NbrId) + "']"
    bgpRouteNeighbor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteNeighbor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteNeighbor.EntityData.Children = make(map[string]types.YChild)
    bgpRouteNeighbor.EntityData.Children["bgp-neighbor-route-filters"] = types.YChild{"BgpNeighborRouteFilters", &bgpRouteNeighbor.BgpNeighborRouteFilters}
    bgpRouteNeighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRouteNeighbor.EntityData.Leafs["nbr-id"] = types.YLeaf{"NbrId", bgpRouteNeighbor.NbrId}
    return &(bgpRouteNeighbor.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters
// BGP neighbor route filters
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP neighbor route filters. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter.
    BgpNeighborRouteFilter []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter
}

func (bgpNeighborRouteFilters *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters) GetEntityData() *types.CommonEntityData {
    bgpNeighborRouteFilters.EntityData.YFilter = bgpNeighborRouteFilters.YFilter
    bgpNeighborRouteFilters.EntityData.YangName = "bgp-neighbor-route-filters"
    bgpNeighborRouteFilters.EntityData.BundleName = "cisco_ios_xe"
    bgpNeighborRouteFilters.EntityData.ParentYangName = "bgp-route-neighbor"
    bgpNeighborRouteFilters.EntityData.SegmentPath = "bgp-neighbor-route-filters"
    bgpNeighborRouteFilters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpNeighborRouteFilters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpNeighborRouteFilters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpNeighborRouteFilters.EntityData.Children = make(map[string]types.YChild)
    bgpNeighborRouteFilters.EntityData.Children["bgp-neighbor-route-filter"] = types.YChild{"BgpNeighborRouteFilter", nil}
    for i := range bgpNeighborRouteFilters.BgpNeighborRouteFilter {
        bgpNeighborRouteFilters.EntityData.Children[types.GetSegmentPath(&bgpNeighborRouteFilters.BgpNeighborRouteFilter[i])] = types.YChild{"BgpNeighborRouteFilter", &bgpNeighborRouteFilters.BgpNeighborRouteFilter[i]}
    }
    bgpNeighborRouteFilters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpNeighborRouteFilters.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter
// List of BGP neighbor route filters
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP neighbor route filter. The type is
    // BgpNeighborRouteFilters.
    NbrFltr interface{}

    // BGP neighbor route entries.
    BgpNeighborRouteEntries BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries
}

func (bgpNeighborRouteFilter *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter) GetEntityData() *types.CommonEntityData {
    bgpNeighborRouteFilter.EntityData.YFilter = bgpNeighborRouteFilter.YFilter
    bgpNeighborRouteFilter.EntityData.YangName = "bgp-neighbor-route-filter"
    bgpNeighborRouteFilter.EntityData.BundleName = "cisco_ios_xe"
    bgpNeighborRouteFilter.EntityData.ParentYangName = "bgp-neighbor-route-filters"
    bgpNeighborRouteFilter.EntityData.SegmentPath = "bgp-neighbor-route-filter" + "[nbr-fltr='" + fmt.Sprintf("%v", bgpNeighborRouteFilter.NbrFltr) + "']"
    bgpNeighborRouteFilter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpNeighborRouteFilter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpNeighborRouteFilter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpNeighborRouteFilter.EntityData.Children = make(map[string]types.YChild)
    bgpNeighborRouteFilter.EntityData.Children["bgp-neighbor-route-entries"] = types.YChild{"BgpNeighborRouteEntries", &bgpNeighborRouteFilter.BgpNeighborRouteEntries}
    bgpNeighborRouteFilter.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpNeighborRouteFilter.EntityData.Leafs["nbr-fltr"] = types.YLeaf{"NbrFltr", bgpNeighborRouteFilter.NbrFltr}
    return &(bgpNeighborRouteFilter.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries
// BGP neighbor route entries
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP neighbor route entries. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry.
    BgpNeighborRouteEntry []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry
}

func (bgpNeighborRouteEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries) GetEntityData() *types.CommonEntityData {
    bgpNeighborRouteEntries.EntityData.YFilter = bgpNeighborRouteEntries.YFilter
    bgpNeighborRouteEntries.EntityData.YangName = "bgp-neighbor-route-entries"
    bgpNeighborRouteEntries.EntityData.BundleName = "cisco_ios_xe"
    bgpNeighborRouteEntries.EntityData.ParentYangName = "bgp-neighbor-route-filter"
    bgpNeighborRouteEntries.EntityData.SegmentPath = "bgp-neighbor-route-entries"
    bgpNeighborRouteEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpNeighborRouteEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpNeighborRouteEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpNeighborRouteEntries.EntityData.Children = make(map[string]types.YChild)
    bgpNeighborRouteEntries.EntityData.Children["bgp-neighbor-route-entry"] = types.YChild{"BgpNeighborRouteEntry", nil}
    for i := range bgpNeighborRouteEntries.BgpNeighborRouteEntry {
        bgpNeighborRouteEntries.EntityData.Children[types.GetSegmentPath(&bgpNeighborRouteEntries.BgpNeighborRouteEntry[i])] = types.YChild{"BgpNeighborRouteEntry", &bgpNeighborRouteEntries.BgpNeighborRouteEntry[i]}
    }
    bgpNeighborRouteEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpNeighborRouteEntries.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry
// List of BGP neighbor route entries
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor routing table entry prefix. The type is
    // string.
    Prefix interface{}

    // Neighbor routing table version. The type is interface{} with range:
    // 0..4294967295.
    Version interface{}

    // Number of paths available for BGP prefix. The type is interface{} with
    // range: 0..4294967295.
    AvailablePaths interface{}

    // Who this prefix was advertized to. The type is string.
    AdvertisedTo interface{}

    // Prefix next hop details.
    BgpNeighborPathEntries BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries
}

func (bgpNeighborRouteEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry) GetEntityData() *types.CommonEntityData {
    bgpNeighborRouteEntry.EntityData.YFilter = bgpNeighborRouteEntry.YFilter
    bgpNeighborRouteEntry.EntityData.YangName = "bgp-neighbor-route-entry"
    bgpNeighborRouteEntry.EntityData.BundleName = "cisco_ios_xe"
    bgpNeighborRouteEntry.EntityData.ParentYangName = "bgp-neighbor-route-entries"
    bgpNeighborRouteEntry.EntityData.SegmentPath = "bgp-neighbor-route-entry" + "[prefix='" + fmt.Sprintf("%v", bgpNeighborRouteEntry.Prefix) + "']"
    bgpNeighborRouteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpNeighborRouteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpNeighborRouteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpNeighborRouteEntry.EntityData.Children = make(map[string]types.YChild)
    bgpNeighborRouteEntry.EntityData.Children["bgp-neighbor-path-entries"] = types.YChild{"BgpNeighborPathEntries", &bgpNeighborRouteEntry.BgpNeighborPathEntries}
    bgpNeighborRouteEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpNeighborRouteEntry.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", bgpNeighborRouteEntry.Prefix}
    bgpNeighborRouteEntry.EntityData.Leafs["version"] = types.YLeaf{"Version", bgpNeighborRouteEntry.Version}
    bgpNeighborRouteEntry.EntityData.Leafs["available-paths"] = types.YLeaf{"AvailablePaths", bgpNeighborRouteEntry.AvailablePaths}
    bgpNeighborRouteEntry.EntityData.Leafs["advertised-to"] = types.YLeaf{"AdvertisedTo", bgpNeighborRouteEntry.AdvertisedTo}
    return &(bgpNeighborRouteEntry.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries
// Prefix next hop details
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of prefix next hop details. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries_BgpNeighborPathEntry.
    BgpNeighborPathEntry []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries_BgpNeighborPathEntry
}

func (bgpNeighborPathEntries *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries) GetEntityData() *types.CommonEntityData {
    bgpNeighborPathEntries.EntityData.YFilter = bgpNeighborPathEntries.YFilter
    bgpNeighborPathEntries.EntityData.YangName = "bgp-neighbor-path-entries"
    bgpNeighborPathEntries.EntityData.BundleName = "cisco_ios_xe"
    bgpNeighborPathEntries.EntityData.ParentYangName = "bgp-neighbor-route-entry"
    bgpNeighborPathEntries.EntityData.SegmentPath = "bgp-neighbor-path-entries"
    bgpNeighborPathEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpNeighborPathEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpNeighborPathEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpNeighborPathEntries.EntityData.Children = make(map[string]types.YChild)
    bgpNeighborPathEntries.EntityData.Children["bgp-neighbor-path-entry"] = types.YChild{"BgpNeighborPathEntry", nil}
    for i := range bgpNeighborPathEntries.BgpNeighborPathEntry {
        bgpNeighborPathEntries.EntityData.Children[types.GetSegmentPath(&bgpNeighborPathEntries.BgpNeighborPathEntry[i])] = types.YChild{"BgpNeighborPathEntry", &bgpNeighborPathEntries.BgpNeighborPathEntry[i]}
    }
    bgpNeighborPathEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpNeighborPathEntries.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries_BgpNeighborPathEntry
// List of prefix next hop details
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries_BgpNeighborPathEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Next hop for this path. The type is string.
    Nexthop interface{}

    // Metric associated with the path. The type is interface{} with range:
    // 0..4294967295.
    Metric interface{}

    // Local preference for this path. The type is interface{} with range:
    // 0..4294967295.
    LocalPref interface{}

    // Path weight. The type is interface{} with range: 0..4294967295.
    Weight interface{}

    // AS path. The type is string.
    AsPath interface{}

    // Path origin. The type is BgpOriginCode.
    Origin interface{}

    // RPKI path status. The type is BgpRpkiStatus.
    RpkiStatus interface{}

    // Community label for the path. The type is string.
    Community interface{}

    // MPLS label in for the path. The type is string.
    MplsIn interface{}

    // MPLS label out for the path. The type is string.
    MplsOut interface{}

    // SR profile name for the path. The type is string.
    SrProfileName interface{}

    // SR binding sid for the path. The type is interface{} with range:
    // 0..4294967295.
    SrBindingSid interface{}

    // SR label index for the path. The type is interface{} with range:
    // 0..4294967295.
    SrLabelIndx interface{}

    // path using 4-octet AS numbers. The type is string.
    As4Path interface{}

    // attribute indicating whether or not the prefix is an atomic aggregate. The
    // type is bool.
    AtomicAggregate interface{}

    // AS number of autonomous system them performed the aggregation. The type is
    // interface{} with range: 0..4294967295.
    AggrAsNumber interface{}

    // AS4 number of autonomous system them performed the aggregation. The type is
    // interface{} with range: 0..4294967295.
    AggrAs4Number interface{}

    // IP address of the router that performed the aggregation. The type is
    // string.
    AggrAddress interface{}

    // the router ID of the originator of the route in the local AS. The type is
    // string.
    OriginatorId interface{}

    // the reflection path the route has passed. The type is string.
    ClusterList interface{}

    // BGP extended community attribute. The type is string.
    ExtendedCommunity interface{}

    // the accumulated IGP metric for the path. The type is interface{} with
    // range: 0..18446744073709551615.
    ExtAigpMetric interface{}

    // path identifier used to uniquely identify a route. The type is interface{}
    // with range: 0..4294967295.
    PathId interface{}

    // Path status.
    PathStatus BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries_BgpNeighborPathEntry_PathStatus
}

func (bgpNeighborPathEntry *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries_BgpNeighborPathEntry) GetEntityData() *types.CommonEntityData {
    bgpNeighborPathEntry.EntityData.YFilter = bgpNeighborPathEntry.YFilter
    bgpNeighborPathEntry.EntityData.YangName = "bgp-neighbor-path-entry"
    bgpNeighborPathEntry.EntityData.BundleName = "cisco_ios_xe"
    bgpNeighborPathEntry.EntityData.ParentYangName = "bgp-neighbor-path-entries"
    bgpNeighborPathEntry.EntityData.SegmentPath = "bgp-neighbor-path-entry" + "[nexthop='" + fmt.Sprintf("%v", bgpNeighborPathEntry.Nexthop) + "']"
    bgpNeighborPathEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpNeighborPathEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpNeighborPathEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpNeighborPathEntry.EntityData.Children = make(map[string]types.YChild)
    bgpNeighborPathEntry.EntityData.Children["path-status"] = types.YChild{"PathStatus", &bgpNeighborPathEntry.PathStatus}
    bgpNeighborPathEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpNeighborPathEntry.EntityData.Leafs["nexthop"] = types.YLeaf{"Nexthop", bgpNeighborPathEntry.Nexthop}
    bgpNeighborPathEntry.EntityData.Leafs["metric"] = types.YLeaf{"Metric", bgpNeighborPathEntry.Metric}
    bgpNeighborPathEntry.EntityData.Leafs["local-pref"] = types.YLeaf{"LocalPref", bgpNeighborPathEntry.LocalPref}
    bgpNeighborPathEntry.EntityData.Leafs["weight"] = types.YLeaf{"Weight", bgpNeighborPathEntry.Weight}
    bgpNeighborPathEntry.EntityData.Leafs["as-path"] = types.YLeaf{"AsPath", bgpNeighborPathEntry.AsPath}
    bgpNeighborPathEntry.EntityData.Leafs["origin"] = types.YLeaf{"Origin", bgpNeighborPathEntry.Origin}
    bgpNeighborPathEntry.EntityData.Leafs["rpki-status"] = types.YLeaf{"RpkiStatus", bgpNeighborPathEntry.RpkiStatus}
    bgpNeighborPathEntry.EntityData.Leafs["community"] = types.YLeaf{"Community", bgpNeighborPathEntry.Community}
    bgpNeighborPathEntry.EntityData.Leafs["mpls-in"] = types.YLeaf{"MplsIn", bgpNeighborPathEntry.MplsIn}
    bgpNeighborPathEntry.EntityData.Leafs["mpls-out"] = types.YLeaf{"MplsOut", bgpNeighborPathEntry.MplsOut}
    bgpNeighborPathEntry.EntityData.Leafs["sr-profile-name"] = types.YLeaf{"SrProfileName", bgpNeighborPathEntry.SrProfileName}
    bgpNeighborPathEntry.EntityData.Leafs["sr-binding-sid"] = types.YLeaf{"SrBindingSid", bgpNeighborPathEntry.SrBindingSid}
    bgpNeighborPathEntry.EntityData.Leafs["sr-label-indx"] = types.YLeaf{"SrLabelIndx", bgpNeighborPathEntry.SrLabelIndx}
    bgpNeighborPathEntry.EntityData.Leafs["as4-path"] = types.YLeaf{"As4Path", bgpNeighborPathEntry.As4Path}
    bgpNeighborPathEntry.EntityData.Leafs["atomic-aggregate"] = types.YLeaf{"AtomicAggregate", bgpNeighborPathEntry.AtomicAggregate}
    bgpNeighborPathEntry.EntityData.Leafs["aggr-as-number"] = types.YLeaf{"AggrAsNumber", bgpNeighborPathEntry.AggrAsNumber}
    bgpNeighborPathEntry.EntityData.Leafs["aggr-as4-number"] = types.YLeaf{"AggrAs4Number", bgpNeighborPathEntry.AggrAs4Number}
    bgpNeighborPathEntry.EntityData.Leafs["aggr-address"] = types.YLeaf{"AggrAddress", bgpNeighborPathEntry.AggrAddress}
    bgpNeighborPathEntry.EntityData.Leafs["originator-id"] = types.YLeaf{"OriginatorId", bgpNeighborPathEntry.OriginatorId}
    bgpNeighborPathEntry.EntityData.Leafs["cluster-list"] = types.YLeaf{"ClusterList", bgpNeighborPathEntry.ClusterList}
    bgpNeighborPathEntry.EntityData.Leafs["extended-community"] = types.YLeaf{"ExtendedCommunity", bgpNeighborPathEntry.ExtendedCommunity}
    bgpNeighborPathEntry.EntityData.Leafs["ext-aigp-metric"] = types.YLeaf{"ExtAigpMetric", bgpNeighborPathEntry.ExtAigpMetric}
    bgpNeighborPathEntry.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", bgpNeighborPathEntry.PathId}
    return &(bgpNeighborPathEntry.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries_BgpNeighborPathEntry_PathStatus
// Path status
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries_BgpNeighborPathEntry_PathStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed path. The type is interface{}.
    Suppressed interface{}

    // Damped path. The type is interface{}.
    Damped interface{}

    // History path. The type is interface{}.
    History interface{}

    // Valid path. The type is interface{}.
    Valid interface{}

    // Sourced path. The type is interface{}.
    Sourced interface{}

    // Best path. The type is interface{}.
    Bestpath interface{}

    // Internal path. The type is interface{}.
    Internal interface{}

    // RIB-fail path. The type is interface{}.
    RibFail interface{}

    // Stale path. The type is interface{}.
    Stale interface{}

    // Multipath path. The type is interface{}.
    Multipath interface{}

    // Backup path. The type is interface{}.
    BackupPath interface{}

    // RT filter path. The type is interface{}.
    RtFilter interface{}

    // Best externa path. The type is interface{}.
    BestExternal interface{}

    // Additional path. The type is interface{}.
    AdditionalPath interface{}

    // RIB compressed path. The type is interface{}.
    RibCompressed interface{}
}

func (pathStatus *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpRouteNeighbors_BgpRouteNeighbor_BgpNeighborRouteFilters_BgpNeighborRouteFilter_BgpNeighborRouteEntries_BgpNeighborRouteEntry_BgpNeighborPathEntries_BgpNeighborPathEntry_PathStatus) GetEntityData() *types.CommonEntityData {
    pathStatus.EntityData.YFilter = pathStatus.YFilter
    pathStatus.EntityData.YangName = "path-status"
    pathStatus.EntityData.BundleName = "cisco_ios_xe"
    pathStatus.EntityData.ParentYangName = "bgp-neighbor-path-entry"
    pathStatus.EntityData.SegmentPath = "path-status"
    pathStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pathStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pathStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pathStatus.EntityData.Children = make(map[string]types.YChild)
    pathStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    pathStatus.EntityData.Leafs["suppressed"] = types.YLeaf{"Suppressed", pathStatus.Suppressed}
    pathStatus.EntityData.Leafs["damped"] = types.YLeaf{"Damped", pathStatus.Damped}
    pathStatus.EntityData.Leafs["history"] = types.YLeaf{"History", pathStatus.History}
    pathStatus.EntityData.Leafs["valid"] = types.YLeaf{"Valid", pathStatus.Valid}
    pathStatus.EntityData.Leafs["sourced"] = types.YLeaf{"Sourced", pathStatus.Sourced}
    pathStatus.EntityData.Leafs["bestpath"] = types.YLeaf{"Bestpath", pathStatus.Bestpath}
    pathStatus.EntityData.Leafs["internal"] = types.YLeaf{"Internal", pathStatus.Internal}
    pathStatus.EntityData.Leafs["rib-fail"] = types.YLeaf{"RibFail", pathStatus.RibFail}
    pathStatus.EntityData.Leafs["stale"] = types.YLeaf{"Stale", pathStatus.Stale}
    pathStatus.EntityData.Leafs["multipath"] = types.YLeaf{"Multipath", pathStatus.Multipath}
    pathStatus.EntityData.Leafs["backup-path"] = types.YLeaf{"BackupPath", pathStatus.BackupPath}
    pathStatus.EntityData.Leafs["rt-filter"] = types.YLeaf{"RtFilter", pathStatus.RtFilter}
    pathStatus.EntityData.Leafs["best-external"] = types.YLeaf{"BestExternal", pathStatus.BestExternal}
    pathStatus.EntityData.Leafs["additional-path"] = types.YLeaf{"AdditionalPath", pathStatus.AdditionalPath}
    pathStatus.EntityData.Leafs["rib-compressed"] = types.YLeaf{"RibCompressed", pathStatus.RibCompressed}
    return &(pathStatus.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpPeerGroups
// BGP peer groups
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpPeerGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP peer groups. The type is slice of
    // BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpPeerGroups_BgpPeerGroup.
    BgpPeerGroup []BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpPeerGroups_BgpPeerGroup
}

func (bgpPeerGroups *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpPeerGroups) GetEntityData() *types.CommonEntityData {
    bgpPeerGroups.EntityData.YFilter = bgpPeerGroups.YFilter
    bgpPeerGroups.EntityData.YangName = "bgp-peer-groups"
    bgpPeerGroups.EntityData.BundleName = "cisco_ios_xe"
    bgpPeerGroups.EntityData.ParentYangName = "bgp-route-af"
    bgpPeerGroups.EntityData.SegmentPath = "bgp-peer-groups"
    bgpPeerGroups.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpPeerGroups.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpPeerGroups.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpPeerGroups.EntityData.Children = make(map[string]types.YChild)
    bgpPeerGroups.EntityData.Children["bgp-peer-group"] = types.YChild{"BgpPeerGroup", nil}
    for i := range bgpPeerGroups.BgpPeerGroup {
        bgpPeerGroups.EntityData.Children[types.GetSegmentPath(&bgpPeerGroups.BgpPeerGroup[i])] = types.YChild{"BgpPeerGroup", &bgpPeerGroups.BgpPeerGroup[i]}
    }
    bgpPeerGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpPeerGroups.EntityData)
}

// BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpPeerGroups_BgpPeerGroup
// List of BGP peer groups
type BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpPeerGroups_BgpPeerGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP peer group name. The type is string.
    Name interface{}

    // Peer Group description string. The type is string.
    Description interface{}

    // Peer Group Remote as value. The type is interface{} with range:
    // 0..4294967295.
    RemoteAs interface{}

    // BGP version being used to communicate with the remote router. The type is
    // interface{} with range: 0..65535.
    BgpVersion interface{}

    // Peer Group minimum time. The type is interface{} with range: 0..65535.
    MinTime interface{}

    // Number of Sessions for peer group. The type is interface{} with range:
    // 0..4294967295.
    NumOfSessions interface{}

    // Number of established sessions  for peer group. The type is interface{}
    // with range: 0..4294967295.
    NumEstabSessions interface{}

    // Number of SSO sesssions for peer group. The type is interface{} with range:
    // 0..4294967295.
    NumSsoSessions interface{}

    // BGP peer group members. The type is slice of string.
    PeerMembers []interface{}

    // BGP peer group Formatted Group Index value. The type is interface{} with
    // range: 0..65535.
    FmtGrpIx interface{}

    // BGP peer group Advertised Index value. The type is interface{} with range:
    // 0..65535.
    AdvIx interface{}

    // BGP peer group aspath  filter in value. The type is interface{} with range:
    // 0..4294967295.
    AspathIn interface{}

    // BGP peer group aspath  filter out value. The type is interface{} with
    // range: 0..4294967295.
    AspathOut interface{}

    // BGP peer group Route Map in value. The type is string.
    RoutemapIn interface{}

    // BGP peer group Route Map out value. The type is string.
    RoutemapOut interface{}

    // BGP peer group Updated Messages value. The type is interface{} with range:
    // 0..18446744073709551615.
    UpdatedMessages interface{}

    // BGP peer group Replicated Count value. The type is interface{} with range:
    // 0..4294967295.
    RepCount interface{}

    // BGP peer group slow peer Threshold value. The type is interface{} with
    // range: 0..65535.
    SlowpeerDetectionValue interface{}

    // BGP weight value. The type is interface{} with range: 0..65535.
    Weight interface{}

    // BGP Send Community status. The type is bool.
    SendCommunity interface{}

    // BGP Extended Community status. The type is bool.
    ExtendedCommunity interface{}

    // BGP Remove PrivateAs status. The type is bool.
    RemovePrivateAs interface{}
}

func (bgpPeerGroup *BgpStateData_BgpRouteVrfs_BgpRouteVrf_BgpRouteAfs_BgpRouteAf_BgpPeerGroups_BgpPeerGroup) GetEntityData() *types.CommonEntityData {
    bgpPeerGroup.EntityData.YFilter = bgpPeerGroup.YFilter
    bgpPeerGroup.EntityData.YangName = "bgp-peer-group"
    bgpPeerGroup.EntityData.BundleName = "cisco_ios_xe"
    bgpPeerGroup.EntityData.ParentYangName = "bgp-peer-groups"
    bgpPeerGroup.EntityData.SegmentPath = "bgp-peer-group" + "[name='" + fmt.Sprintf("%v", bgpPeerGroup.Name) + "']"
    bgpPeerGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpPeerGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpPeerGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpPeerGroup.EntityData.Children = make(map[string]types.YChild)
    bgpPeerGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpPeerGroup.EntityData.Leafs["name"] = types.YLeaf{"Name", bgpPeerGroup.Name}
    bgpPeerGroup.EntityData.Leafs["description"] = types.YLeaf{"Description", bgpPeerGroup.Description}
    bgpPeerGroup.EntityData.Leafs["remote-as"] = types.YLeaf{"RemoteAs", bgpPeerGroup.RemoteAs}
    bgpPeerGroup.EntityData.Leafs["bgp-version"] = types.YLeaf{"BgpVersion", bgpPeerGroup.BgpVersion}
    bgpPeerGroup.EntityData.Leafs["min-time"] = types.YLeaf{"MinTime", bgpPeerGroup.MinTime}
    bgpPeerGroup.EntityData.Leafs["num-of-sessions"] = types.YLeaf{"NumOfSessions", bgpPeerGroup.NumOfSessions}
    bgpPeerGroup.EntityData.Leafs["num-estab-sessions"] = types.YLeaf{"NumEstabSessions", bgpPeerGroup.NumEstabSessions}
    bgpPeerGroup.EntityData.Leafs["num-sso-sessions"] = types.YLeaf{"NumSsoSessions", bgpPeerGroup.NumSsoSessions}
    bgpPeerGroup.EntityData.Leafs["peer-members"] = types.YLeaf{"PeerMembers", bgpPeerGroup.PeerMembers}
    bgpPeerGroup.EntityData.Leafs["fmt-grp-ix"] = types.YLeaf{"FmtGrpIx", bgpPeerGroup.FmtGrpIx}
    bgpPeerGroup.EntityData.Leafs["adv-ix"] = types.YLeaf{"AdvIx", bgpPeerGroup.AdvIx}
    bgpPeerGroup.EntityData.Leafs["aspath-in"] = types.YLeaf{"AspathIn", bgpPeerGroup.AspathIn}
    bgpPeerGroup.EntityData.Leafs["aspath-out"] = types.YLeaf{"AspathOut", bgpPeerGroup.AspathOut}
    bgpPeerGroup.EntityData.Leafs["routemap-in"] = types.YLeaf{"RoutemapIn", bgpPeerGroup.RoutemapIn}
    bgpPeerGroup.EntityData.Leafs["routemap-out"] = types.YLeaf{"RoutemapOut", bgpPeerGroup.RoutemapOut}
    bgpPeerGroup.EntityData.Leafs["updated-messages"] = types.YLeaf{"UpdatedMessages", bgpPeerGroup.UpdatedMessages}
    bgpPeerGroup.EntityData.Leafs["rep-count"] = types.YLeaf{"RepCount", bgpPeerGroup.RepCount}
    bgpPeerGroup.EntityData.Leafs["slowpeer-detection-value"] = types.YLeaf{"SlowpeerDetectionValue", bgpPeerGroup.SlowpeerDetectionValue}
    bgpPeerGroup.EntityData.Leafs["weight"] = types.YLeaf{"Weight", bgpPeerGroup.Weight}
    bgpPeerGroup.EntityData.Leafs["send-community"] = types.YLeaf{"SendCommunity", bgpPeerGroup.SendCommunity}
    bgpPeerGroup.EntityData.Leafs["extended-community"] = types.YLeaf{"ExtendedCommunity", bgpPeerGroup.ExtendedCommunity}
    bgpPeerGroup.EntityData.Leafs["remove-private-as"] = types.YLeaf{"RemovePrivateAs", bgpPeerGroup.RemovePrivateAs}
    return &(bgpPeerGroup.EntityData)
}

// BgpStateData_BgpRouteRds
// BGP RDs
type BgpStateData_BgpRouteRds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP RDs. The type is slice of BgpStateData_BgpRouteRds_BgpRouteRd.
    BgpRouteRd []BgpStateData_BgpRouteRds_BgpRouteRd
}

func (bgpRouteRds *BgpStateData_BgpRouteRds) GetEntityData() *types.CommonEntityData {
    bgpRouteRds.EntityData.YFilter = bgpRouteRds.YFilter
    bgpRouteRds.EntityData.YangName = "bgp-route-rds"
    bgpRouteRds.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteRds.EntityData.ParentYangName = "bgp-state-data"
    bgpRouteRds.EntityData.SegmentPath = "bgp-route-rds"
    bgpRouteRds.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteRds.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteRds.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteRds.EntityData.Children = make(map[string]types.YChild)
    bgpRouteRds.EntityData.Children["bgp-route-rd"] = types.YChild{"BgpRouteRd", nil}
    for i := range bgpRouteRds.BgpRouteRd {
        bgpRouteRds.EntityData.Children[types.GetSegmentPath(&bgpRouteRds.BgpRouteRd[i])] = types.YChild{"BgpRouteRd", &bgpRouteRds.BgpRouteRd[i]}
    }
    bgpRouteRds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRouteRds.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd
// List of BGP RDs
type BgpStateData_BgpRouteRds_BgpRouteRd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP rd value. The type is string.
    RdValue interface{}

    // BGP rd address families.
    BgpRdRouteAfs BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs
}

func (bgpRouteRd *BgpStateData_BgpRouteRds_BgpRouteRd) GetEntityData() *types.CommonEntityData {
    bgpRouteRd.EntityData.YFilter = bgpRouteRd.YFilter
    bgpRouteRd.EntityData.YangName = "bgp-route-rd"
    bgpRouteRd.EntityData.BundleName = "cisco_ios_xe"
    bgpRouteRd.EntityData.ParentYangName = "bgp-route-rds"
    bgpRouteRd.EntityData.SegmentPath = "bgp-route-rd" + "[rd-value='" + fmt.Sprintf("%v", bgpRouteRd.RdValue) + "']"
    bgpRouteRd.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRouteRd.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRouteRd.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRouteRd.EntityData.Children = make(map[string]types.YChild)
    bgpRouteRd.EntityData.Children["bgp-rd-route-afs"] = types.YChild{"BgpRdRouteAfs", &bgpRouteRd.BgpRdRouteAfs}
    bgpRouteRd.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRouteRd.EntityData.Leafs["rd-value"] = types.YLeaf{"RdValue", bgpRouteRd.RdValue}
    return &(bgpRouteRd.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs
// BGP rd address families
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP RD address families. The type is slice of
    // BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf.
    BgpRdRouteAf []BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf
}

func (bgpRdRouteAfs *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs) GetEntityData() *types.CommonEntityData {
    bgpRdRouteAfs.EntityData.YFilter = bgpRdRouteAfs.YFilter
    bgpRdRouteAfs.EntityData.YangName = "bgp-rd-route-afs"
    bgpRdRouteAfs.EntityData.BundleName = "cisco_ios_xe"
    bgpRdRouteAfs.EntityData.ParentYangName = "bgp-route-rd"
    bgpRdRouteAfs.EntityData.SegmentPath = "bgp-rd-route-afs"
    bgpRdRouteAfs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdRouteAfs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdRouteAfs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdRouteAfs.EntityData.Children = make(map[string]types.YChild)
    bgpRdRouteAfs.EntityData.Children["bgp-rd-route-af"] = types.YChild{"BgpRdRouteAf", nil}
    for i := range bgpRdRouteAfs.BgpRdRouteAf {
        bgpRdRouteAfs.EntityData.Children[types.GetSegmentPath(&bgpRdRouteAfs.BgpRdRouteAf[i])] = types.YChild{"BgpRdRouteAf", &bgpRdRouteAfs.BgpRdRouteAf[i]}
    }
    bgpRdRouteAfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRdRouteAfs.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf
// List of BGP RD address families
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP address family. The type is AfiSafi.
    AfiSafi interface{}

    // BGP RD route filters.
    BgpRdRouteFilters BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters

    // BGP RD route neighbors.
    BgpRdRouteNeighbors BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors
}

func (bgpRdRouteAf *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf) GetEntityData() *types.CommonEntityData {
    bgpRdRouteAf.EntityData.YFilter = bgpRdRouteAf.YFilter
    bgpRdRouteAf.EntityData.YangName = "bgp-rd-route-af"
    bgpRdRouteAf.EntityData.BundleName = "cisco_ios_xe"
    bgpRdRouteAf.EntityData.ParentYangName = "bgp-rd-route-afs"
    bgpRdRouteAf.EntityData.SegmentPath = "bgp-rd-route-af" + "[afi-safi='" + fmt.Sprintf("%v", bgpRdRouteAf.AfiSafi) + "']"
    bgpRdRouteAf.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdRouteAf.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdRouteAf.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdRouteAf.EntityData.Children = make(map[string]types.YChild)
    bgpRdRouteAf.EntityData.Children["bgp-rd-route-filters"] = types.YChild{"BgpRdRouteFilters", &bgpRdRouteAf.BgpRdRouteFilters}
    bgpRdRouteAf.EntityData.Children["bgp-rd-route-neighbors"] = types.YChild{"BgpRdRouteNeighbors", &bgpRdRouteAf.BgpRdRouteNeighbors}
    bgpRdRouteAf.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRdRouteAf.EntityData.Leafs["afi-safi"] = types.YLeaf{"AfiSafi", bgpRdRouteAf.AfiSafi}
    return &(bgpRdRouteAf.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters
// BGP RD route filters
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP RD route filters. The type is slice of
    // BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter.
    BgpRdRouteFilter []BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter
}

func (bgpRdRouteFilters *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters) GetEntityData() *types.CommonEntityData {
    bgpRdRouteFilters.EntityData.YFilter = bgpRdRouteFilters.YFilter
    bgpRdRouteFilters.EntityData.YangName = "bgp-rd-route-filters"
    bgpRdRouteFilters.EntityData.BundleName = "cisco_ios_xe"
    bgpRdRouteFilters.EntityData.ParentYangName = "bgp-rd-route-af"
    bgpRdRouteFilters.EntityData.SegmentPath = "bgp-rd-route-filters"
    bgpRdRouteFilters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdRouteFilters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdRouteFilters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdRouteFilters.EntityData.Children = make(map[string]types.YChild)
    bgpRdRouteFilters.EntityData.Children["bgp-rd-route-filter"] = types.YChild{"BgpRdRouteFilter", nil}
    for i := range bgpRdRouteFilters.BgpRdRouteFilter {
        bgpRdRouteFilters.EntityData.Children[types.GetSegmentPath(&bgpRdRouteFilters.BgpRdRouteFilter[i])] = types.YChild{"BgpRdRouteFilter", &bgpRdRouteFilters.BgpRdRouteFilter[i]}
    }
    bgpRdRouteFilters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRdRouteFilters.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter
// List of BGP RD route filters
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP RD route filter. The type is BgpRouteFilters.
    RouteFilter interface{}

    // BGP RD route entries.
    BgpRdRouteEntries BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries
}

func (bgpRdRouteFilter *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter) GetEntityData() *types.CommonEntityData {
    bgpRdRouteFilter.EntityData.YFilter = bgpRdRouteFilter.YFilter
    bgpRdRouteFilter.EntityData.YangName = "bgp-rd-route-filter"
    bgpRdRouteFilter.EntityData.BundleName = "cisco_ios_xe"
    bgpRdRouteFilter.EntityData.ParentYangName = "bgp-rd-route-filters"
    bgpRdRouteFilter.EntityData.SegmentPath = "bgp-rd-route-filter" + "[route-filter='" + fmt.Sprintf("%v", bgpRdRouteFilter.RouteFilter) + "']"
    bgpRdRouteFilter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdRouteFilter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdRouteFilter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdRouteFilter.EntityData.Children = make(map[string]types.YChild)
    bgpRdRouteFilter.EntityData.Children["bgp-rd-route-entries"] = types.YChild{"BgpRdRouteEntries", &bgpRdRouteFilter.BgpRdRouteEntries}
    bgpRdRouteFilter.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRdRouteFilter.EntityData.Leafs["route-filter"] = types.YLeaf{"RouteFilter", bgpRdRouteFilter.RouteFilter}
    return &(bgpRdRouteFilter.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries
// BGP RD route entries
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP RD route entries. The type is slice of
    // BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry.
    BgpRdRouteEntry []BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry
}

func (bgpRdRouteEntries *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries) GetEntityData() *types.CommonEntityData {
    bgpRdRouteEntries.EntityData.YFilter = bgpRdRouteEntries.YFilter
    bgpRdRouteEntries.EntityData.YangName = "bgp-rd-route-entries"
    bgpRdRouteEntries.EntityData.BundleName = "cisco_ios_xe"
    bgpRdRouteEntries.EntityData.ParentYangName = "bgp-rd-route-filter"
    bgpRdRouteEntries.EntityData.SegmentPath = "bgp-rd-route-entries"
    bgpRdRouteEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdRouteEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdRouteEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdRouteEntries.EntityData.Children = make(map[string]types.YChild)
    bgpRdRouteEntries.EntityData.Children["bgp-rd-route-entry"] = types.YChild{"BgpRdRouteEntry", nil}
    for i := range bgpRdRouteEntries.BgpRdRouteEntry {
        bgpRdRouteEntries.EntityData.Children[types.GetSegmentPath(&bgpRdRouteEntries.BgpRdRouteEntry[i])] = types.YChild{"BgpRdRouteEntry", &bgpRdRouteEntries.BgpRdRouteEntry[i]}
    }
    bgpRdRouteEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRdRouteEntries.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry
// List of BGP RD route entries
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. RD Routing table entry prefix. The type is string.
    Prefix interface{}

    // RD Routing table version. The type is interface{} with range:
    // 0..4294967295.
    Version interface{}

    // Number of  paths available for BGP prefix. The type is interface{} with
    // range: 0..4294967295.
    AvailablePaths interface{}

    // Whom is thie prefix advertized to. The type is string.
    AdvertisedTo interface{}

    // Prefix next hop details.
    BgpRdPathEntries BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries
}

func (bgpRdRouteEntry *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry) GetEntityData() *types.CommonEntityData {
    bgpRdRouteEntry.EntityData.YFilter = bgpRdRouteEntry.YFilter
    bgpRdRouteEntry.EntityData.YangName = "bgp-rd-route-entry"
    bgpRdRouteEntry.EntityData.BundleName = "cisco_ios_xe"
    bgpRdRouteEntry.EntityData.ParentYangName = "bgp-rd-route-entries"
    bgpRdRouteEntry.EntityData.SegmentPath = "bgp-rd-route-entry" + "[prefix='" + fmt.Sprintf("%v", bgpRdRouteEntry.Prefix) + "']"
    bgpRdRouteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdRouteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdRouteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdRouteEntry.EntityData.Children = make(map[string]types.YChild)
    bgpRdRouteEntry.EntityData.Children["bgp-rd-path-entries"] = types.YChild{"BgpRdPathEntries", &bgpRdRouteEntry.BgpRdPathEntries}
    bgpRdRouteEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRdRouteEntry.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", bgpRdRouteEntry.Prefix}
    bgpRdRouteEntry.EntityData.Leafs["version"] = types.YLeaf{"Version", bgpRdRouteEntry.Version}
    bgpRdRouteEntry.EntityData.Leafs["available-paths"] = types.YLeaf{"AvailablePaths", bgpRdRouteEntry.AvailablePaths}
    bgpRdRouteEntry.EntityData.Leafs["advertised-to"] = types.YLeaf{"AdvertisedTo", bgpRdRouteEntry.AdvertisedTo}
    return &(bgpRdRouteEntry.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries
// Prefix next hop details
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of prefix next hop details. The type is slice of
    // BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries_BgpRdPathEntry.
    BgpRdPathEntry []BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries_BgpRdPathEntry
}

func (bgpRdPathEntries *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries) GetEntityData() *types.CommonEntityData {
    bgpRdPathEntries.EntityData.YFilter = bgpRdPathEntries.YFilter
    bgpRdPathEntries.EntityData.YangName = "bgp-rd-path-entries"
    bgpRdPathEntries.EntityData.BundleName = "cisco_ios_xe"
    bgpRdPathEntries.EntityData.ParentYangName = "bgp-rd-route-entry"
    bgpRdPathEntries.EntityData.SegmentPath = "bgp-rd-path-entries"
    bgpRdPathEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdPathEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdPathEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdPathEntries.EntityData.Children = make(map[string]types.YChild)
    bgpRdPathEntries.EntityData.Children["bgp-rd-path-entry"] = types.YChild{"BgpRdPathEntry", nil}
    for i := range bgpRdPathEntries.BgpRdPathEntry {
        bgpRdPathEntries.EntityData.Children[types.GetSegmentPath(&bgpRdPathEntries.BgpRdPathEntry[i])] = types.YChild{"BgpRdPathEntry", &bgpRdPathEntries.BgpRdPathEntry[i]}
    }
    bgpRdPathEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRdPathEntries.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries_BgpRdPathEntry
// List of prefix next hop details
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries_BgpRdPathEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Next hop for this path. The type is string.
    Nexthop interface{}

    // Metric associated with the path. The type is interface{} with range:
    // 0..4294967295.
    Metric interface{}

    // Local preference for this path. The type is interface{} with range:
    // 0..4294967295.
    LocalPref interface{}

    // Path weight. The type is interface{} with range: 0..4294967295.
    Weight interface{}

    // AS path. The type is string.
    AsPath interface{}

    // Path origin. The type is BgpOriginCode.
    Origin interface{}

    // RPKI path status. The type is BgpRpkiStatus.
    RpkiStatus interface{}

    // Community label for the path. The type is string.
    Community interface{}

    // MPLS label in for the path. The type is string.
    MplsIn interface{}

    // MPLS label out for the path. The type is string.
    MplsOut interface{}

    // SR profile name for the path. The type is string.
    SrProfileName interface{}

    // SR binding sid for the path. The type is interface{} with range:
    // 0..4294967295.
    SrBindingSid interface{}

    // SR label index for the path. The type is interface{} with range:
    // 0..4294967295.
    SrLabelIndx interface{}

    // path using 4-octet AS numbers. The type is string.
    As4Path interface{}

    // attribute indicating whether or not the prefix is an atomic aggregate. The
    // type is bool.
    AtomicAggregate interface{}

    // AS number of autonomous system them performed the aggregation. The type is
    // interface{} with range: 0..4294967295.
    AggrAsNumber interface{}

    // AS4 number of autonomous system them performed the aggregation. The type is
    // interface{} with range: 0..4294967295.
    AggrAs4Number interface{}

    // IP address of the router that performed the aggregation. The type is
    // string.
    AggrAddress interface{}

    // the router ID of the originator of the route in the local AS. The type is
    // string.
    OriginatorId interface{}

    // the reflection path the route has passed. The type is string.
    ClusterList interface{}

    // BGP extended community attribute. The type is string.
    ExtendedCommunity interface{}

    // the accumulated IGP metric for the path. The type is interface{} with
    // range: 0..18446744073709551615.
    ExtAigpMetric interface{}

    // path identifier used to uniquely identify a route. The type is interface{}
    // with range: 0..4294967295.
    PathId interface{}

    // Path status.
    PathStatus BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries_BgpRdPathEntry_PathStatus
}

func (bgpRdPathEntry *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries_BgpRdPathEntry) GetEntityData() *types.CommonEntityData {
    bgpRdPathEntry.EntityData.YFilter = bgpRdPathEntry.YFilter
    bgpRdPathEntry.EntityData.YangName = "bgp-rd-path-entry"
    bgpRdPathEntry.EntityData.BundleName = "cisco_ios_xe"
    bgpRdPathEntry.EntityData.ParentYangName = "bgp-rd-path-entries"
    bgpRdPathEntry.EntityData.SegmentPath = "bgp-rd-path-entry" + "[nexthop='" + fmt.Sprintf("%v", bgpRdPathEntry.Nexthop) + "']"
    bgpRdPathEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdPathEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdPathEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdPathEntry.EntityData.Children = make(map[string]types.YChild)
    bgpRdPathEntry.EntityData.Children["path-status"] = types.YChild{"PathStatus", &bgpRdPathEntry.PathStatus}
    bgpRdPathEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRdPathEntry.EntityData.Leafs["nexthop"] = types.YLeaf{"Nexthop", bgpRdPathEntry.Nexthop}
    bgpRdPathEntry.EntityData.Leafs["metric"] = types.YLeaf{"Metric", bgpRdPathEntry.Metric}
    bgpRdPathEntry.EntityData.Leafs["local-pref"] = types.YLeaf{"LocalPref", bgpRdPathEntry.LocalPref}
    bgpRdPathEntry.EntityData.Leafs["weight"] = types.YLeaf{"Weight", bgpRdPathEntry.Weight}
    bgpRdPathEntry.EntityData.Leafs["as-path"] = types.YLeaf{"AsPath", bgpRdPathEntry.AsPath}
    bgpRdPathEntry.EntityData.Leafs["origin"] = types.YLeaf{"Origin", bgpRdPathEntry.Origin}
    bgpRdPathEntry.EntityData.Leafs["rpki-status"] = types.YLeaf{"RpkiStatus", bgpRdPathEntry.RpkiStatus}
    bgpRdPathEntry.EntityData.Leafs["community"] = types.YLeaf{"Community", bgpRdPathEntry.Community}
    bgpRdPathEntry.EntityData.Leafs["mpls-in"] = types.YLeaf{"MplsIn", bgpRdPathEntry.MplsIn}
    bgpRdPathEntry.EntityData.Leafs["mpls-out"] = types.YLeaf{"MplsOut", bgpRdPathEntry.MplsOut}
    bgpRdPathEntry.EntityData.Leafs["sr-profile-name"] = types.YLeaf{"SrProfileName", bgpRdPathEntry.SrProfileName}
    bgpRdPathEntry.EntityData.Leafs["sr-binding-sid"] = types.YLeaf{"SrBindingSid", bgpRdPathEntry.SrBindingSid}
    bgpRdPathEntry.EntityData.Leafs["sr-label-indx"] = types.YLeaf{"SrLabelIndx", bgpRdPathEntry.SrLabelIndx}
    bgpRdPathEntry.EntityData.Leafs["as4-path"] = types.YLeaf{"As4Path", bgpRdPathEntry.As4Path}
    bgpRdPathEntry.EntityData.Leafs["atomic-aggregate"] = types.YLeaf{"AtomicAggregate", bgpRdPathEntry.AtomicAggregate}
    bgpRdPathEntry.EntityData.Leafs["aggr-as-number"] = types.YLeaf{"AggrAsNumber", bgpRdPathEntry.AggrAsNumber}
    bgpRdPathEntry.EntityData.Leafs["aggr-as4-number"] = types.YLeaf{"AggrAs4Number", bgpRdPathEntry.AggrAs4Number}
    bgpRdPathEntry.EntityData.Leafs["aggr-address"] = types.YLeaf{"AggrAddress", bgpRdPathEntry.AggrAddress}
    bgpRdPathEntry.EntityData.Leafs["originator-id"] = types.YLeaf{"OriginatorId", bgpRdPathEntry.OriginatorId}
    bgpRdPathEntry.EntityData.Leafs["cluster-list"] = types.YLeaf{"ClusterList", bgpRdPathEntry.ClusterList}
    bgpRdPathEntry.EntityData.Leafs["extended-community"] = types.YLeaf{"ExtendedCommunity", bgpRdPathEntry.ExtendedCommunity}
    bgpRdPathEntry.EntityData.Leafs["ext-aigp-metric"] = types.YLeaf{"ExtAigpMetric", bgpRdPathEntry.ExtAigpMetric}
    bgpRdPathEntry.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", bgpRdPathEntry.PathId}
    return &(bgpRdPathEntry.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries_BgpRdPathEntry_PathStatus
// Path status
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries_BgpRdPathEntry_PathStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed path. The type is interface{}.
    Suppressed interface{}

    // Damped path. The type is interface{}.
    Damped interface{}

    // History path. The type is interface{}.
    History interface{}

    // Valid path. The type is interface{}.
    Valid interface{}

    // Sourced path. The type is interface{}.
    Sourced interface{}

    // Best path. The type is interface{}.
    Bestpath interface{}

    // Internal path. The type is interface{}.
    Internal interface{}

    // RIB-fail path. The type is interface{}.
    RibFail interface{}

    // Stale path. The type is interface{}.
    Stale interface{}

    // Multipath path. The type is interface{}.
    Multipath interface{}

    // Backup path. The type is interface{}.
    BackupPath interface{}

    // RT filter path. The type is interface{}.
    RtFilter interface{}

    // Best externa path. The type is interface{}.
    BestExternal interface{}

    // Additional path. The type is interface{}.
    AdditionalPath interface{}

    // RIB compressed path. The type is interface{}.
    RibCompressed interface{}
}

func (pathStatus *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteFilters_BgpRdRouteFilter_BgpRdRouteEntries_BgpRdRouteEntry_BgpRdPathEntries_BgpRdPathEntry_PathStatus) GetEntityData() *types.CommonEntityData {
    pathStatus.EntityData.YFilter = pathStatus.YFilter
    pathStatus.EntityData.YangName = "path-status"
    pathStatus.EntityData.BundleName = "cisco_ios_xe"
    pathStatus.EntityData.ParentYangName = "bgp-rd-path-entry"
    pathStatus.EntityData.SegmentPath = "path-status"
    pathStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pathStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pathStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pathStatus.EntityData.Children = make(map[string]types.YChild)
    pathStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    pathStatus.EntityData.Leafs["suppressed"] = types.YLeaf{"Suppressed", pathStatus.Suppressed}
    pathStatus.EntityData.Leafs["damped"] = types.YLeaf{"Damped", pathStatus.Damped}
    pathStatus.EntityData.Leafs["history"] = types.YLeaf{"History", pathStatus.History}
    pathStatus.EntityData.Leafs["valid"] = types.YLeaf{"Valid", pathStatus.Valid}
    pathStatus.EntityData.Leafs["sourced"] = types.YLeaf{"Sourced", pathStatus.Sourced}
    pathStatus.EntityData.Leafs["bestpath"] = types.YLeaf{"Bestpath", pathStatus.Bestpath}
    pathStatus.EntityData.Leafs["internal"] = types.YLeaf{"Internal", pathStatus.Internal}
    pathStatus.EntityData.Leafs["rib-fail"] = types.YLeaf{"RibFail", pathStatus.RibFail}
    pathStatus.EntityData.Leafs["stale"] = types.YLeaf{"Stale", pathStatus.Stale}
    pathStatus.EntityData.Leafs["multipath"] = types.YLeaf{"Multipath", pathStatus.Multipath}
    pathStatus.EntityData.Leafs["backup-path"] = types.YLeaf{"BackupPath", pathStatus.BackupPath}
    pathStatus.EntityData.Leafs["rt-filter"] = types.YLeaf{"RtFilter", pathStatus.RtFilter}
    pathStatus.EntityData.Leafs["best-external"] = types.YLeaf{"BestExternal", pathStatus.BestExternal}
    pathStatus.EntityData.Leafs["additional-path"] = types.YLeaf{"AdditionalPath", pathStatus.AdditionalPath}
    pathStatus.EntityData.Leafs["rib-compressed"] = types.YLeaf{"RibCompressed", pathStatus.RibCompressed}
    return &(pathStatus.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors
// BGP RD route neighbors
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP RD route neighbors. The type is slice of
    // BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor.
    BgpRdRouteNeighbor []BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor
}

func (bgpRdRouteNeighbors *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors) GetEntityData() *types.CommonEntityData {
    bgpRdRouteNeighbors.EntityData.YFilter = bgpRdRouteNeighbors.YFilter
    bgpRdRouteNeighbors.EntityData.YangName = "bgp-rd-route-neighbors"
    bgpRdRouteNeighbors.EntityData.BundleName = "cisco_ios_xe"
    bgpRdRouteNeighbors.EntityData.ParentYangName = "bgp-rd-route-af"
    bgpRdRouteNeighbors.EntityData.SegmentPath = "bgp-rd-route-neighbors"
    bgpRdRouteNeighbors.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdRouteNeighbors.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdRouteNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdRouteNeighbors.EntityData.Children = make(map[string]types.YChild)
    bgpRdRouteNeighbors.EntityData.Children["bgp-rd-route-neighbor"] = types.YChild{"BgpRdRouteNeighbor", nil}
    for i := range bgpRdRouteNeighbors.BgpRdRouteNeighbor {
        bgpRdRouteNeighbors.EntityData.Children[types.GetSegmentPath(&bgpRdRouteNeighbors.BgpRdRouteNeighbor[i])] = types.YChild{"BgpRdRouteNeighbor", &bgpRdRouteNeighbors.BgpRdRouteNeighbor[i]}
    }
    bgpRdRouteNeighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRdRouteNeighbors.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor
// List of BGP RD route neighbors
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP RD neighbor ID. The type is string.
    NeighborId interface{}

    // BGP RD neighbor route filters.
    BgpRdNeighborRouteFilters BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters
}

func (bgpRdRouteNeighbor *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor) GetEntityData() *types.CommonEntityData {
    bgpRdRouteNeighbor.EntityData.YFilter = bgpRdRouteNeighbor.YFilter
    bgpRdRouteNeighbor.EntityData.YangName = "bgp-rd-route-neighbor"
    bgpRdRouteNeighbor.EntityData.BundleName = "cisco_ios_xe"
    bgpRdRouteNeighbor.EntityData.ParentYangName = "bgp-rd-route-neighbors"
    bgpRdRouteNeighbor.EntityData.SegmentPath = "bgp-rd-route-neighbor" + "[neighbor-id='" + fmt.Sprintf("%v", bgpRdRouteNeighbor.NeighborId) + "']"
    bgpRdRouteNeighbor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdRouteNeighbor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdRouteNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdRouteNeighbor.EntityData.Children = make(map[string]types.YChild)
    bgpRdRouteNeighbor.EntityData.Children["bgp-rd-neighbor-route-filters"] = types.YChild{"BgpRdNeighborRouteFilters", &bgpRdRouteNeighbor.BgpRdNeighborRouteFilters}
    bgpRdRouteNeighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRdRouteNeighbor.EntityData.Leafs["neighbor-id"] = types.YLeaf{"NeighborId", bgpRdRouteNeighbor.NeighborId}
    return &(bgpRdRouteNeighbor.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters
// BGP RD neighbor route filters
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP RD neighbor route filters. The type is slice of
    // BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter.
    BgpRdNeighborRouteFilter []BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter
}

func (bgpRdNeighborRouteFilters *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters) GetEntityData() *types.CommonEntityData {
    bgpRdNeighborRouteFilters.EntityData.YFilter = bgpRdNeighborRouteFilters.YFilter
    bgpRdNeighborRouteFilters.EntityData.YangName = "bgp-rd-neighbor-route-filters"
    bgpRdNeighborRouteFilters.EntityData.BundleName = "cisco_ios_xe"
    bgpRdNeighborRouteFilters.EntityData.ParentYangName = "bgp-rd-route-neighbor"
    bgpRdNeighborRouteFilters.EntityData.SegmentPath = "bgp-rd-neighbor-route-filters"
    bgpRdNeighborRouteFilters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdNeighborRouteFilters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdNeighborRouteFilters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdNeighborRouteFilters.EntityData.Children = make(map[string]types.YChild)
    bgpRdNeighborRouteFilters.EntityData.Children["bgp-rd-neighbor-route-filter"] = types.YChild{"BgpRdNeighborRouteFilter", nil}
    for i := range bgpRdNeighborRouteFilters.BgpRdNeighborRouteFilter {
        bgpRdNeighborRouteFilters.EntityData.Children[types.GetSegmentPath(&bgpRdNeighborRouteFilters.BgpRdNeighborRouteFilter[i])] = types.YChild{"BgpRdNeighborRouteFilter", &bgpRdNeighborRouteFilters.BgpRdNeighborRouteFilter[i]}
    }
    bgpRdNeighborRouteFilters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRdNeighborRouteFilters.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter
// List of BGP RD neighbor route filters
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BGP RD neighbor route filter. The type is
    // BgpNeighborRouteFilters.
    NeighborFilter interface{}

    // BGP RD neighbor route entries.
    BgpRdNeighborRouteEntries BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries
}

func (bgpRdNeighborRouteFilter *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter) GetEntityData() *types.CommonEntityData {
    bgpRdNeighborRouteFilter.EntityData.YFilter = bgpRdNeighborRouteFilter.YFilter
    bgpRdNeighborRouteFilter.EntityData.YangName = "bgp-rd-neighbor-route-filter"
    bgpRdNeighborRouteFilter.EntityData.BundleName = "cisco_ios_xe"
    bgpRdNeighborRouteFilter.EntityData.ParentYangName = "bgp-rd-neighbor-route-filters"
    bgpRdNeighborRouteFilter.EntityData.SegmentPath = "bgp-rd-neighbor-route-filter" + "[neighbor-filter='" + fmt.Sprintf("%v", bgpRdNeighborRouteFilter.NeighborFilter) + "']"
    bgpRdNeighborRouteFilter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdNeighborRouteFilter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdNeighborRouteFilter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdNeighborRouteFilter.EntityData.Children = make(map[string]types.YChild)
    bgpRdNeighborRouteFilter.EntityData.Children["bgp-rd-neighbor-route-entries"] = types.YChild{"BgpRdNeighborRouteEntries", &bgpRdNeighborRouteFilter.BgpRdNeighborRouteEntries}
    bgpRdNeighborRouteFilter.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRdNeighborRouteFilter.EntityData.Leafs["neighbor-filter"] = types.YLeaf{"NeighborFilter", bgpRdNeighborRouteFilter.NeighborFilter}
    return &(bgpRdNeighborRouteFilter.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries
// BGP RD neighbor route entries
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP RD neighbor route entries. The type is slice of
    // BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry.
    BgpRdNeighborRouteEntry []BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry
}

func (bgpRdNeighborRouteEntries *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries) GetEntityData() *types.CommonEntityData {
    bgpRdNeighborRouteEntries.EntityData.YFilter = bgpRdNeighborRouteEntries.YFilter
    bgpRdNeighborRouteEntries.EntityData.YangName = "bgp-rd-neighbor-route-entries"
    bgpRdNeighborRouteEntries.EntityData.BundleName = "cisco_ios_xe"
    bgpRdNeighborRouteEntries.EntityData.ParentYangName = "bgp-rd-neighbor-route-filter"
    bgpRdNeighborRouteEntries.EntityData.SegmentPath = "bgp-rd-neighbor-route-entries"
    bgpRdNeighborRouteEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdNeighborRouteEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdNeighborRouteEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdNeighborRouteEntries.EntityData.Children = make(map[string]types.YChild)
    bgpRdNeighborRouteEntries.EntityData.Children["bgp-rd-neighbor-route-entry"] = types.YChild{"BgpRdNeighborRouteEntry", nil}
    for i := range bgpRdNeighborRouteEntries.BgpRdNeighborRouteEntry {
        bgpRdNeighborRouteEntries.EntityData.Children[types.GetSegmentPath(&bgpRdNeighborRouteEntries.BgpRdNeighborRouteEntry[i])] = types.YChild{"BgpRdNeighborRouteEntry", &bgpRdNeighborRouteEntries.BgpRdNeighborRouteEntry[i]}
    }
    bgpRdNeighborRouteEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRdNeighborRouteEntries.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry
// List of BGP RD neighbor route entries
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. RD neighbor routing table entry prefix. The type
    // is string.
    Prefix interface{}

    // RD neighbor routing table version. The type is interface{} with range:
    // 0..4294967295.
    Version interface{}

    // Number of  paths available for BGP prefix. The type is interface{} with
    // range: 0..4294967295.
    AvailablePaths interface{}

    // Who this prefix was advertized to. The type is string.
    AdvertisedTo interface{}

    // Prefix next hop details.
    BgpRdNeighborPathEntries BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries
}

func (bgpRdNeighborRouteEntry *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry) GetEntityData() *types.CommonEntityData {
    bgpRdNeighborRouteEntry.EntityData.YFilter = bgpRdNeighborRouteEntry.YFilter
    bgpRdNeighborRouteEntry.EntityData.YangName = "bgp-rd-neighbor-route-entry"
    bgpRdNeighborRouteEntry.EntityData.BundleName = "cisco_ios_xe"
    bgpRdNeighborRouteEntry.EntityData.ParentYangName = "bgp-rd-neighbor-route-entries"
    bgpRdNeighborRouteEntry.EntityData.SegmentPath = "bgp-rd-neighbor-route-entry" + "[prefix='" + fmt.Sprintf("%v", bgpRdNeighborRouteEntry.Prefix) + "']"
    bgpRdNeighborRouteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdNeighborRouteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdNeighborRouteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdNeighborRouteEntry.EntityData.Children = make(map[string]types.YChild)
    bgpRdNeighborRouteEntry.EntityData.Children["bgp-rd-neighbor-path-entries"] = types.YChild{"BgpRdNeighborPathEntries", &bgpRdNeighborRouteEntry.BgpRdNeighborPathEntries}
    bgpRdNeighborRouteEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRdNeighborRouteEntry.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", bgpRdNeighborRouteEntry.Prefix}
    bgpRdNeighborRouteEntry.EntityData.Leafs["version"] = types.YLeaf{"Version", bgpRdNeighborRouteEntry.Version}
    bgpRdNeighborRouteEntry.EntityData.Leafs["available-paths"] = types.YLeaf{"AvailablePaths", bgpRdNeighborRouteEntry.AvailablePaths}
    bgpRdNeighborRouteEntry.EntityData.Leafs["advertised-to"] = types.YLeaf{"AdvertisedTo", bgpRdNeighborRouteEntry.AdvertisedTo}
    return &(bgpRdNeighborRouteEntry.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries
// Prefix next hop details
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of prefix next hop details. The type is slice of
    // BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries_BgpRdNeighborPathEntry.
    BgpRdNeighborPathEntry []BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries_BgpRdNeighborPathEntry
}

func (bgpRdNeighborPathEntries *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries) GetEntityData() *types.CommonEntityData {
    bgpRdNeighborPathEntries.EntityData.YFilter = bgpRdNeighborPathEntries.YFilter
    bgpRdNeighborPathEntries.EntityData.YangName = "bgp-rd-neighbor-path-entries"
    bgpRdNeighborPathEntries.EntityData.BundleName = "cisco_ios_xe"
    bgpRdNeighborPathEntries.EntityData.ParentYangName = "bgp-rd-neighbor-route-entry"
    bgpRdNeighborPathEntries.EntityData.SegmentPath = "bgp-rd-neighbor-path-entries"
    bgpRdNeighborPathEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdNeighborPathEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdNeighborPathEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdNeighborPathEntries.EntityData.Children = make(map[string]types.YChild)
    bgpRdNeighborPathEntries.EntityData.Children["bgp-rd-neighbor-path-entry"] = types.YChild{"BgpRdNeighborPathEntry", nil}
    for i := range bgpRdNeighborPathEntries.BgpRdNeighborPathEntry {
        bgpRdNeighborPathEntries.EntityData.Children[types.GetSegmentPath(&bgpRdNeighborPathEntries.BgpRdNeighborPathEntry[i])] = types.YChild{"BgpRdNeighborPathEntry", &bgpRdNeighborPathEntries.BgpRdNeighborPathEntry[i]}
    }
    bgpRdNeighborPathEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpRdNeighborPathEntries.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries_BgpRdNeighborPathEntry
// List of prefix next hop details
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries_BgpRdNeighborPathEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Next hop for this path. The type is string.
    Nexthop interface{}

    // Metric associated with the path. The type is interface{} with range:
    // 0..4294967295.
    Metric interface{}

    // Local preference for this path. The type is interface{} with range:
    // 0..4294967295.
    LocalPref interface{}

    // Path weight. The type is interface{} with range: 0..4294967295.
    Weight interface{}

    // AS path. The type is string.
    AsPath interface{}

    // Path origin. The type is BgpOriginCode.
    Origin interface{}

    // RPKI path status. The type is BgpRpkiStatus.
    RpkiStatus interface{}

    // Community label for the path. The type is string.
    Community interface{}

    // MPLS label in for the path. The type is string.
    MplsIn interface{}

    // MPLS label out for the path. The type is string.
    MplsOut interface{}

    // SR profile name for the path. The type is string.
    SrProfileName interface{}

    // SR binding sid for the path. The type is interface{} with range:
    // 0..4294967295.
    SrBindingSid interface{}

    // SR label index for the path. The type is interface{} with range:
    // 0..4294967295.
    SrLabelIndx interface{}

    // path using 4-octet AS numbers. The type is string.
    As4Path interface{}

    // attribute indicating whether or not the prefix is an atomic aggregate. The
    // type is bool.
    AtomicAggregate interface{}

    // AS number of autonomous system them performed the aggregation. The type is
    // interface{} with range: 0..4294967295.
    AggrAsNumber interface{}

    // AS4 number of autonomous system them performed the aggregation. The type is
    // interface{} with range: 0..4294967295.
    AggrAs4Number interface{}

    // IP address of the router that performed the aggregation. The type is
    // string.
    AggrAddress interface{}

    // the router ID of the originator of the route in the local AS. The type is
    // string.
    OriginatorId interface{}

    // the reflection path the route has passed. The type is string.
    ClusterList interface{}

    // BGP extended community attribute. The type is string.
    ExtendedCommunity interface{}

    // the accumulated IGP metric for the path. The type is interface{} with
    // range: 0..18446744073709551615.
    ExtAigpMetric interface{}

    // path identifier used to uniquely identify a route. The type is interface{}
    // with range: 0..4294967295.
    PathId interface{}

    // Path status.
    PathStatus BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries_BgpRdNeighborPathEntry_PathStatus
}

func (bgpRdNeighborPathEntry *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries_BgpRdNeighborPathEntry) GetEntityData() *types.CommonEntityData {
    bgpRdNeighborPathEntry.EntityData.YFilter = bgpRdNeighborPathEntry.YFilter
    bgpRdNeighborPathEntry.EntityData.YangName = "bgp-rd-neighbor-path-entry"
    bgpRdNeighborPathEntry.EntityData.BundleName = "cisco_ios_xe"
    bgpRdNeighborPathEntry.EntityData.ParentYangName = "bgp-rd-neighbor-path-entries"
    bgpRdNeighborPathEntry.EntityData.SegmentPath = "bgp-rd-neighbor-path-entry" + "[nexthop='" + fmt.Sprintf("%v", bgpRdNeighborPathEntry.Nexthop) + "']"
    bgpRdNeighborPathEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRdNeighborPathEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRdNeighborPathEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRdNeighborPathEntry.EntityData.Children = make(map[string]types.YChild)
    bgpRdNeighborPathEntry.EntityData.Children["path-status"] = types.YChild{"PathStatus", &bgpRdNeighborPathEntry.PathStatus}
    bgpRdNeighborPathEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpRdNeighborPathEntry.EntityData.Leafs["nexthop"] = types.YLeaf{"Nexthop", bgpRdNeighborPathEntry.Nexthop}
    bgpRdNeighborPathEntry.EntityData.Leafs["metric"] = types.YLeaf{"Metric", bgpRdNeighborPathEntry.Metric}
    bgpRdNeighborPathEntry.EntityData.Leafs["local-pref"] = types.YLeaf{"LocalPref", bgpRdNeighborPathEntry.LocalPref}
    bgpRdNeighborPathEntry.EntityData.Leafs["weight"] = types.YLeaf{"Weight", bgpRdNeighborPathEntry.Weight}
    bgpRdNeighborPathEntry.EntityData.Leafs["as-path"] = types.YLeaf{"AsPath", bgpRdNeighborPathEntry.AsPath}
    bgpRdNeighborPathEntry.EntityData.Leafs["origin"] = types.YLeaf{"Origin", bgpRdNeighborPathEntry.Origin}
    bgpRdNeighborPathEntry.EntityData.Leafs["rpki-status"] = types.YLeaf{"RpkiStatus", bgpRdNeighborPathEntry.RpkiStatus}
    bgpRdNeighborPathEntry.EntityData.Leafs["community"] = types.YLeaf{"Community", bgpRdNeighborPathEntry.Community}
    bgpRdNeighborPathEntry.EntityData.Leafs["mpls-in"] = types.YLeaf{"MplsIn", bgpRdNeighborPathEntry.MplsIn}
    bgpRdNeighborPathEntry.EntityData.Leafs["mpls-out"] = types.YLeaf{"MplsOut", bgpRdNeighborPathEntry.MplsOut}
    bgpRdNeighborPathEntry.EntityData.Leafs["sr-profile-name"] = types.YLeaf{"SrProfileName", bgpRdNeighborPathEntry.SrProfileName}
    bgpRdNeighborPathEntry.EntityData.Leafs["sr-binding-sid"] = types.YLeaf{"SrBindingSid", bgpRdNeighborPathEntry.SrBindingSid}
    bgpRdNeighborPathEntry.EntityData.Leafs["sr-label-indx"] = types.YLeaf{"SrLabelIndx", bgpRdNeighborPathEntry.SrLabelIndx}
    bgpRdNeighborPathEntry.EntityData.Leafs["as4-path"] = types.YLeaf{"As4Path", bgpRdNeighborPathEntry.As4Path}
    bgpRdNeighborPathEntry.EntityData.Leafs["atomic-aggregate"] = types.YLeaf{"AtomicAggregate", bgpRdNeighborPathEntry.AtomicAggregate}
    bgpRdNeighborPathEntry.EntityData.Leafs["aggr-as-number"] = types.YLeaf{"AggrAsNumber", bgpRdNeighborPathEntry.AggrAsNumber}
    bgpRdNeighborPathEntry.EntityData.Leafs["aggr-as4-number"] = types.YLeaf{"AggrAs4Number", bgpRdNeighborPathEntry.AggrAs4Number}
    bgpRdNeighborPathEntry.EntityData.Leafs["aggr-address"] = types.YLeaf{"AggrAddress", bgpRdNeighborPathEntry.AggrAddress}
    bgpRdNeighborPathEntry.EntityData.Leafs["originator-id"] = types.YLeaf{"OriginatorId", bgpRdNeighborPathEntry.OriginatorId}
    bgpRdNeighborPathEntry.EntityData.Leafs["cluster-list"] = types.YLeaf{"ClusterList", bgpRdNeighborPathEntry.ClusterList}
    bgpRdNeighborPathEntry.EntityData.Leafs["extended-community"] = types.YLeaf{"ExtendedCommunity", bgpRdNeighborPathEntry.ExtendedCommunity}
    bgpRdNeighborPathEntry.EntityData.Leafs["ext-aigp-metric"] = types.YLeaf{"ExtAigpMetric", bgpRdNeighborPathEntry.ExtAigpMetric}
    bgpRdNeighborPathEntry.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", bgpRdNeighborPathEntry.PathId}
    return &(bgpRdNeighborPathEntry.EntityData)
}

// BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries_BgpRdNeighborPathEntry_PathStatus
// Path status
type BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries_BgpRdNeighborPathEntry_PathStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppressed path. The type is interface{}.
    Suppressed interface{}

    // Damped path. The type is interface{}.
    Damped interface{}

    // History path. The type is interface{}.
    History interface{}

    // Valid path. The type is interface{}.
    Valid interface{}

    // Sourced path. The type is interface{}.
    Sourced interface{}

    // Best path. The type is interface{}.
    Bestpath interface{}

    // Internal path. The type is interface{}.
    Internal interface{}

    // RIB-fail path. The type is interface{}.
    RibFail interface{}

    // Stale path. The type is interface{}.
    Stale interface{}

    // Multipath path. The type is interface{}.
    Multipath interface{}

    // Backup path. The type is interface{}.
    BackupPath interface{}

    // RT filter path. The type is interface{}.
    RtFilter interface{}

    // Best externa path. The type is interface{}.
    BestExternal interface{}

    // Additional path. The type is interface{}.
    AdditionalPath interface{}

    // RIB compressed path. The type is interface{}.
    RibCompressed interface{}
}

func (pathStatus *BgpStateData_BgpRouteRds_BgpRouteRd_BgpRdRouteAfs_BgpRdRouteAf_BgpRdRouteNeighbors_BgpRdRouteNeighbor_BgpRdNeighborRouteFilters_BgpRdNeighborRouteFilter_BgpRdNeighborRouteEntries_BgpRdNeighborRouteEntry_BgpRdNeighborPathEntries_BgpRdNeighborPathEntry_PathStatus) GetEntityData() *types.CommonEntityData {
    pathStatus.EntityData.YFilter = pathStatus.YFilter
    pathStatus.EntityData.YangName = "path-status"
    pathStatus.EntityData.BundleName = "cisco_ios_xe"
    pathStatus.EntityData.ParentYangName = "bgp-rd-neighbor-path-entry"
    pathStatus.EntityData.SegmentPath = "path-status"
    pathStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pathStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pathStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pathStatus.EntityData.Children = make(map[string]types.YChild)
    pathStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    pathStatus.EntityData.Leafs["suppressed"] = types.YLeaf{"Suppressed", pathStatus.Suppressed}
    pathStatus.EntityData.Leafs["damped"] = types.YLeaf{"Damped", pathStatus.Damped}
    pathStatus.EntityData.Leafs["history"] = types.YLeaf{"History", pathStatus.History}
    pathStatus.EntityData.Leafs["valid"] = types.YLeaf{"Valid", pathStatus.Valid}
    pathStatus.EntityData.Leafs["sourced"] = types.YLeaf{"Sourced", pathStatus.Sourced}
    pathStatus.EntityData.Leafs["bestpath"] = types.YLeaf{"Bestpath", pathStatus.Bestpath}
    pathStatus.EntityData.Leafs["internal"] = types.YLeaf{"Internal", pathStatus.Internal}
    pathStatus.EntityData.Leafs["rib-fail"] = types.YLeaf{"RibFail", pathStatus.RibFail}
    pathStatus.EntityData.Leafs["stale"] = types.YLeaf{"Stale", pathStatus.Stale}
    pathStatus.EntityData.Leafs["multipath"] = types.YLeaf{"Multipath", pathStatus.Multipath}
    pathStatus.EntityData.Leafs["backup-path"] = types.YLeaf{"BackupPath", pathStatus.BackupPath}
    pathStatus.EntityData.Leafs["rt-filter"] = types.YLeaf{"RtFilter", pathStatus.RtFilter}
    pathStatus.EntityData.Leafs["best-external"] = types.YLeaf{"BestExternal", pathStatus.BestExternal}
    pathStatus.EntityData.Leafs["additional-path"] = types.YLeaf{"AdditionalPath", pathStatus.AdditionalPath}
    pathStatus.EntityData.Leafs["rib-compressed"] = types.YLeaf{"RibCompressed", pathStatus.RibCompressed}
    return &(pathStatus.EntityData)
}

