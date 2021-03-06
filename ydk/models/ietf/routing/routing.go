// This YANG module defines essential components for the management
// of a routing subsystem.
// 
// Copyright (c) 2014 IETF Trust and the persons identified as
// authors of the code. All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject to
// the license terms contained in, the Simplified BSD License set
// forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// The key words 'MUST', 'MUST NOT', 'REQUIRED', 'SHALL', 'SHALL
// NOT', 'SHOULD', 'SHOULD NOT', 'RECOMMENDED', 'MAY', and
// 'OPTIONAL' in the module text are to be interpreted as described
// in RFC 2119 (http://tools.ietf.org/html/rfc2119).
// 
// This version of this YANG module is part of RFC XXXX
// (http://tools.ietf.org/html/rfcXXXX); see the RFC itself for
// full legal notices.
package routing

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package routing"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-routing routing-state}", reflect.TypeOf(RoutingState{}))
    ydk.RegisterEntity("ietf-routing:routing-state", reflect.TypeOf(RoutingState{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-routing routing}", reflect.TypeOf(Routing{}))
    ydk.RegisterEntity("ietf-routing:routing", reflect.TypeOf(Routing{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-routing fib-route}", reflect.TypeOf(FibRoute{}))
    ydk.RegisterEntity("ietf-routing:fib-route", reflect.TypeOf(FibRoute{}))
}

type AddressFamily struct {
}

func (id AddressFamily) String() string {
	return "ietf-routing:address-family"
}

type Ipv4 struct {
}

func (id Ipv4) String() string {
	return "ietf-routing:ipv4"
}

type Ipv6 struct {
}

func (id Ipv6) String() string {
	return "ietf-routing:ipv6"
}

type RoutingInstance struct {
}

func (id RoutingInstance) String() string {
	return "ietf-routing:routing-instance"
}

type DefaultRoutingInstance struct {
}

func (id DefaultRoutingInstance) String() string {
	return "ietf-routing:default-routing-instance"
}

type VrfRoutingInstance struct {
}

func (id VrfRoutingInstance) String() string {
	return "ietf-routing:vrf-routing-instance"
}

type RoutingProtocol struct {
}

func (id RoutingProtocol) String() string {
	return "ietf-routing:routing-protocol"
}

type Direct struct {
}

func (id Direct) String() string {
	return "ietf-routing:direct"
}

type Static struct {
}

func (id Static) String() string {
	return "ietf-routing:static"
}

// RoutingState
// State data of the routing subsystem.
type RoutingState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each list entry is a container for state data of a routing instance.  An
    // implementation MUST support routing instance(s) of the type
    // 'rt:default-routing-instance', and MAY support other types. An
    // implementation MAY restrict the number of routing instances of each
    // supported type.  An implementation SHOULD create at least one
    // system-controlled instance, and MAY allow the clients to create
    // user-controlled routing instances in configuration. The type is slice of
    // RoutingState_RoutingInstance.
    RoutingInstance []RoutingState_RoutingInstance
}

func (routingState *RoutingState) GetEntityData() *types.CommonEntityData {
    routingState.EntityData.YFilter = routingState.YFilter
    routingState.EntityData.YangName = "routing-state"
    routingState.EntityData.BundleName = "ietf"
    routingState.EntityData.ParentYangName = "ietf-routing"
    routingState.EntityData.SegmentPath = "ietf-routing:routing-state"
    routingState.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routingState.EntityData.NamespaceTable = ietf.GetNamespaces()
    routingState.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routingState.EntityData.Children = make(map[string]types.YChild)
    routingState.EntityData.Children["routing-instance"] = types.YChild{"RoutingInstance", nil}
    for i := range routingState.RoutingInstance {
        routingState.EntityData.Children[types.GetSegmentPath(&routingState.RoutingInstance[i])] = types.YChild{"RoutingInstance", &routingState.RoutingInstance[i]}
    }
    routingState.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routingState.EntityData)
}

// RoutingState_RoutingInstance
// Each list entry is a container for state data of a routing
// instance.
// 
// An implementation MUST support routing instance(s) of the
// type 'rt:default-routing-instance', and MAY support other
// types. An implementation MAY restrict the number of routing
// instances of each supported type.
// 
// An implementation SHOULD create at least one
// system-controlled instance, and MAY allow the clients to
// create user-controlled routing instances in
// configuration.
type RoutingState_RoutingInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the routing instance.  For
    // system-controlled instances the name is persistent, i.e., it SHOULD NOT
    // change across reboots. The type is string.
    Name interface{}

    // The routing instance type. The type is one of the following:
    // DefaultRoutingInstanceVrfRoutingInstance.
    Type_ interface{}

    // A 32-bit number in the form of a dotted quad that is used by some routing
    // protocols identifying a router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    RouterId interface{}

    // Network layer interfaces belonging to the routing instance.
    Interfaces RoutingState_RoutingInstance_Interfaces

    // Container for the list of routing protocol instances.
    RoutingProtocols RoutingState_RoutingInstance_RoutingProtocols

    // Container for RIBs.
    Ribs RoutingState_RoutingInstance_Ribs
}

func (routingInstance *RoutingState_RoutingInstance) GetEntityData() *types.CommonEntityData {
    routingInstance.EntityData.YFilter = routingInstance.YFilter
    routingInstance.EntityData.YangName = "routing-instance"
    routingInstance.EntityData.BundleName = "ietf"
    routingInstance.EntityData.ParentYangName = "routing-state"
    routingInstance.EntityData.SegmentPath = "routing-instance" + "[name='" + fmt.Sprintf("%v", routingInstance.Name) + "']"
    routingInstance.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routingInstance.EntityData.NamespaceTable = ietf.GetNamespaces()
    routingInstance.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routingInstance.EntityData.Children = make(map[string]types.YChild)
    routingInstance.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &routingInstance.Interfaces}
    routingInstance.EntityData.Children["routing-protocols"] = types.YChild{"RoutingProtocols", &routingInstance.RoutingProtocols}
    routingInstance.EntityData.Children["ribs"] = types.YChild{"Ribs", &routingInstance.Ribs}
    routingInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    routingInstance.EntityData.Leafs["name"] = types.YLeaf{"Name", routingInstance.Name}
    routingInstance.EntityData.Leafs["type"] = types.YLeaf{"Type_", routingInstance.Type_}
    routingInstance.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", routingInstance.RouterId}
    return &(routingInstance.EntityData)
}

// RoutingState_RoutingInstance_Interfaces
// Network layer interfaces belonging to the routing
// instance.
type RoutingState_RoutingInstance_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry is a reference to the name of a configured network layer
    // interface. The type is slice of string. Refers to
    // interfaces.InterfacesState_Interface_Name
    Interface_ []interface{}
}

func (interfaces *RoutingState_RoutingInstance_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "ietf"
    interfaces.EntityData.ParentYangName = "routing-instance"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interfaces.EntityData.NamespaceTable = ietf.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaces.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", interfaces.Interface_}
    return &(interfaces.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols
// Container for the list of routing protocol instances.
type RoutingState_RoutingInstance_RoutingProtocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State data of a routing protocol instance.  An implementation MUST provide
    // exactly one system-controlled instance of the type 'direct'. Other
    // instances MAY be created by configuration. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol.
    RoutingProtocol []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol
}

func (routingProtocols *RoutingState_RoutingInstance_RoutingProtocols) GetEntityData() *types.CommonEntityData {
    routingProtocols.EntityData.YFilter = routingProtocols.YFilter
    routingProtocols.EntityData.YangName = "routing-protocols"
    routingProtocols.EntityData.BundleName = "ietf"
    routingProtocols.EntityData.ParentYangName = "routing-instance"
    routingProtocols.EntityData.SegmentPath = "routing-protocols"
    routingProtocols.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routingProtocols.EntityData.NamespaceTable = ietf.GetNamespaces()
    routingProtocols.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routingProtocols.EntityData.Children = make(map[string]types.YChild)
    routingProtocols.EntityData.Children["routing-protocol"] = types.YChild{"RoutingProtocol", nil}
    for i := range routingProtocols.RoutingProtocol {
        routingProtocols.EntityData.Children[types.GetSegmentPath(&routingProtocols.RoutingProtocol[i])] = types.YChild{"RoutingProtocol", &routingProtocols.RoutingProtocol[i]}
    }
    routingProtocols.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routingProtocols.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol
// State data of a routing protocol instance.
// 
// An implementation MUST provide exactly one
// system-controlled instance of the type 'direct'. Other
// instances MAY be created by configuration.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Type of the routing protocol. The type is one of
    // the following: OspfOspfv2Ospfv3DirectStatic.
    Type_ interface{}

    // This attribute is a key. The name of the routing protocol instance.  For
    // system-controlled instances this name is persistent, i.e., it SHOULD NOT
    // change across reboots. The type is string.
    Name interface{}

    // OSPF.
    Ospf RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf
}

func (routingProtocol *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol) GetEntityData() *types.CommonEntityData {
    routingProtocol.EntityData.YFilter = routingProtocol.YFilter
    routingProtocol.EntityData.YangName = "routing-protocol"
    routingProtocol.EntityData.BundleName = "ietf"
    routingProtocol.EntityData.ParentYangName = "routing-protocols"
    routingProtocol.EntityData.SegmentPath = "routing-protocol" + "[type='" + fmt.Sprintf("%v", routingProtocol.Type_) + "']" + "[name='" + fmt.Sprintf("%v", routingProtocol.Name) + "']"
    routingProtocol.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routingProtocol.EntityData.NamespaceTable = ietf.GetNamespaces()
    routingProtocol.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routingProtocol.EntityData.Children = make(map[string]types.YChild)
    routingProtocol.EntityData.Children["ietf-ospf:ospf"] = types.YChild{"Ospf", &routingProtocol.Ospf}
    routingProtocol.EntityData.Leafs = make(map[string]types.YLeaf)
    routingProtocol.EntityData.Leafs["type"] = types.YLeaf{"Type_", routingProtocol.Type_}
    routingProtocol.EntityData.Leafs["name"] = types.YLeaf{"Name", routingProtocol.Name}
    return &(routingProtocol.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf
// OSPF
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF operation mode. The type is one of the following: ShipsInTheNight.
    OperationMode interface{}

    // An OSPF routing protocol instance. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance.
    Instance []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance
}

func (ospf *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "ietf"
    ospf.EntityData.ParentYangName = "routing-protocol"
    ospf.EntityData.SegmentPath = "ietf-ospf:ospf"
    ospf.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ospf.EntityData.NamespaceTable = ietf.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Children["instance"] = types.YChild{"Instance", nil}
    for i := range ospf.Instance {
        ospf.EntityData.Children[types.GetSegmentPath(&ospf.Instance[i])] = types.YChild{"Instance", &ospf.Instance[i]}
    }
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["operation-mode"] = types.YLeaf{"OperationMode", ospf.OperationMode}
    return &(ospf.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance
// An OSPF routing protocol instance.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address-family of the instance. The type is one of
    // the following: Ipv4Ipv4UnicastIpv6Ipv6Unicast.
    Af interface{}

    // Defined in RFC 2328. A 32-bit number that uniquely identifies the router.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    RouterId interface{}

    // List of OSPF areas. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area.
    Area []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area

    // List OSPF AS scope LSA databases. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas.
    AsScopeLsas []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas

    // OSPF topology. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology
}

func (instance *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "ietf"
    instance.EntityData.ParentYangName = "ospf"
    instance.EntityData.SegmentPath = "instance" + "[af='" + fmt.Sprintf("%v", instance.Af) + "']"
    instance.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    instance.EntityData.NamespaceTable = ietf.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Children["area"] = types.YChild{"Area", nil}
    for i := range instance.Area {
        instance.EntityData.Children[types.GetSegmentPath(&instance.Area[i])] = types.YChild{"Area", &instance.Area[i]}
    }
    instance.EntityData.Children["as-scope-lsas"] = types.YChild{"AsScopeLsas", nil}
    for i := range instance.AsScopeLsas {
        instance.EntityData.Children[types.GetSegmentPath(&instance.AsScopeLsas[i])] = types.YChild{"AsScopeLsas", &instance.AsScopeLsas[i]}
    }
    instance.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range instance.Topology {
        instance.EntityData.Children[types.GetSegmentPath(&instance.Topology[i])] = types.YChild{"Topology", &instance.Topology[i]}
    }
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["af"] = types.YLeaf{"Af", instance.Af}
    instance.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", instance.RouterId}
    return &(instance.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area
// List of OSPF areas
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Area ID. The type is one of the following types:
    // int with range: 0..4294967295, or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    AreaId interface{}

    // List of OSPF interfaces. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces.
    Interfaces []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces

    // List OSPF area scope LSA databases. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas.
    AreaScopeLsas []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetEntityData() *types.CommonEntityData {
    area.EntityData.YFilter = area.YFilter
    area.EntityData.YangName = "area"
    area.EntityData.BundleName = "ietf"
    area.EntityData.ParentYangName = "instance"
    area.EntityData.SegmentPath = "area" + "[area-id='" + fmt.Sprintf("%v", area.AreaId) + "']"
    area.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    area.EntityData.NamespaceTable = ietf.GetNamespaces()
    area.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    area.EntityData.Children = make(map[string]types.YChild)
    area.EntityData.Children["interfaces"] = types.YChild{"Interfaces", nil}
    for i := range area.Interfaces {
        area.EntityData.Children[types.GetSegmentPath(&area.Interfaces[i])] = types.YChild{"Interfaces", &area.Interfaces[i]}
    }
    area.EntityData.Children["area-scope-lsas"] = types.YChild{"AreaScopeLsas", nil}
    for i := range area.AreaScopeLsas {
        area.EntityData.Children[types.GetSegmentPath(&area.AreaScopeLsas[i])] = types.YChild{"AreaScopeLsas", &area.AreaScopeLsas[i]}
    }
    area.EntityData.Leafs = make(map[string]types.YLeaf)
    area.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", area.AreaId}
    return &(area.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces
// List of OSPF interfaces.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string.
    Interface_ interface{}

    // Network type. The type is NetworkType.
    NetworkType interface{}

    // Enable/Disable passive. The type is bool.
    Passive interface{}

    // Enable/Disable demand circuit. The type is bool.
    DemandCircuit interface{}

    // Set prefix as a node representative prefix. The type is bool. The default
    // value is false.
    NodeFlag interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Time between hello packets. The type is interface{} with range: 1..65535.
    // Units are seconds.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead. The type is interface{}
    // with range: 1..65535. Units are seconds.
    DeadInterval interface{}

    // Time between retransmitting unacknowledged Link State Advertisements
    // (LSAs). The type is interface{} with range: 1..65535. Units are seconds.
    RetransmitInterval interface{}

    // Estimated time needed to send link-state update. The type is interface{}
    // with range: 1..65535. Units are seconds.
    TransmitDelay interface{}

    // Enable/Disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/Disable link-local signaling (LLS) support. The type is bool.
    Lls interface{}

    // Suppress advertisement of the prefixes. The type is bool.
    PrefixSuppression interface{}

    // Enable/disable bfd. The type is bool.
    Bfd interface{}

    // Enable/disable protocol on the interface. The type is bool. The default
    // value is true.
    Enable interface{}

    // Interface state. The type is string.
    State interface{}

    // Hello timer. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    HelloTimer interface{}

    // Wait timer. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    WaitTimer interface{}

    // DR. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Dr interface{}

    // BDR. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bdr interface{}

    // Configure ospf multi-area.
    MultiArea RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea

    // Static configured neighbors.
    StaticNeighbors RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors

    // Fast-reroute configuration.
    FastReroute RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute

    // TTL security check.
    TtlSecurity RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity

    // Authentication configuration.
    Authentication RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication

    // List of OSPF neighbors. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor.
    Neighbor []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor

    // List OSPF link scope LSA databases. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas.
    LinkScopeLsas []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas

    // OSPF interface topology. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology
}

func (interfaces *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "ietf"
    interfaces.EntityData.ParentYangName = "area"
    interfaces.EntityData.SegmentPath = "interfaces" + "[interface='" + fmt.Sprintf("%v", interfaces.Interface_) + "']"
    interfaces.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interfaces.EntityData.NamespaceTable = ietf.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["multi-area"] = types.YChild{"MultiArea", &interfaces.MultiArea}
    interfaces.EntityData.Children["static-neighbors"] = types.YChild{"StaticNeighbors", &interfaces.StaticNeighbors}
    interfaces.EntityData.Children["fast-reroute"] = types.YChild{"FastReroute", &interfaces.FastReroute}
    interfaces.EntityData.Children["ttl-security"] = types.YChild{"TtlSecurity", &interfaces.TtlSecurity}
    interfaces.EntityData.Children["authentication"] = types.YChild{"Authentication", &interfaces.Authentication}
    interfaces.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range interfaces.Neighbor {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Neighbor[i])] = types.YChild{"Neighbor", &interfaces.Neighbor[i]}
    }
    interfaces.EntityData.Children["link-scope-lsas"] = types.YChild{"LinkScopeLsas", nil}
    for i := range interfaces.LinkScopeLsas {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.LinkScopeLsas[i])] = types.YChild{"LinkScopeLsas", &interfaces.LinkScopeLsas[i]}
    }
    interfaces.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range interfaces.Topology {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Topology[i])] = types.YChild{"Topology", &interfaces.Topology[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaces.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", interfaces.Interface_}
    interfaces.EntityData.Leafs["network-type"] = types.YLeaf{"NetworkType", interfaces.NetworkType}
    interfaces.EntityData.Leafs["passive"] = types.YLeaf{"Passive", interfaces.Passive}
    interfaces.EntityData.Leafs["demand-circuit"] = types.YLeaf{"DemandCircuit", interfaces.DemandCircuit}
    interfaces.EntityData.Leafs["node-flag"] = types.YLeaf{"NodeFlag", interfaces.NodeFlag}
    interfaces.EntityData.Leafs["cost"] = types.YLeaf{"Cost", interfaces.Cost}
    interfaces.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", interfaces.HelloInterval}
    interfaces.EntityData.Leafs["dead-interval"] = types.YLeaf{"DeadInterval", interfaces.DeadInterval}
    interfaces.EntityData.Leafs["retransmit-interval"] = types.YLeaf{"RetransmitInterval", interfaces.RetransmitInterval}
    interfaces.EntityData.Leafs["transmit-delay"] = types.YLeaf{"TransmitDelay", interfaces.TransmitDelay}
    interfaces.EntityData.Leafs["mtu-ignore"] = types.YLeaf{"MtuIgnore", interfaces.MtuIgnore}
    interfaces.EntityData.Leafs["lls"] = types.YLeaf{"Lls", interfaces.Lls}
    interfaces.EntityData.Leafs["prefix-suppression"] = types.YLeaf{"PrefixSuppression", interfaces.PrefixSuppression}
    interfaces.EntityData.Leafs["bfd"] = types.YLeaf{"Bfd", interfaces.Bfd}
    interfaces.EntityData.Leafs["enable"] = types.YLeaf{"Enable", interfaces.Enable}
    interfaces.EntityData.Leafs["state"] = types.YLeaf{"State", interfaces.State}
    interfaces.EntityData.Leafs["hello-timer"] = types.YLeaf{"HelloTimer", interfaces.HelloTimer}
    interfaces.EntityData.Leafs["wait-timer"] = types.YLeaf{"WaitTimer", interfaces.WaitTimer}
    interfaces.EntityData.Leafs["dr"] = types.YLeaf{"Dr", interfaces.Dr}
    interfaces.EntityData.Leafs["bdr"] = types.YLeaf{"Bdr", interfaces.Bdr}
    return &(interfaces.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea
// Configure ospf multi-area.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multi-area ID. The type is one of the following types: int with range:
    // 0..4294967295, or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    MultiAreaId interface{}

    // Interface cost for multi-area. The type is interface{} with range:
    // 0..65535.
    Cost interface{}
}

func (multiArea *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_MultiArea) GetEntityData() *types.CommonEntityData {
    multiArea.EntityData.YFilter = multiArea.YFilter
    multiArea.EntityData.YangName = "multi-area"
    multiArea.EntityData.BundleName = "ietf"
    multiArea.EntityData.ParentYangName = "interfaces"
    multiArea.EntityData.SegmentPath = "multi-area"
    multiArea.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    multiArea.EntityData.NamespaceTable = ietf.GetNamespaces()
    multiArea.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    multiArea.EntityData.Children = make(map[string]types.YChild)
    multiArea.EntityData.Leafs = make(map[string]types.YLeaf)
    multiArea.EntityData.Leafs["multi-area-id"] = types.YLeaf{"MultiAreaId", multiArea.MultiAreaId}
    multiArea.EntityData.Leafs["cost"] = types.YLeaf{"Cost", multiArea.Cost}
    return &(multiArea.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors
// Static configured neighbors.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify a neighbor router. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor.
    Neighbor []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor
}

func (staticNeighbors *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors) GetEntityData() *types.CommonEntityData {
    staticNeighbors.EntityData.YFilter = staticNeighbors.YFilter
    staticNeighbors.EntityData.YangName = "static-neighbors"
    staticNeighbors.EntityData.BundleName = "ietf"
    staticNeighbors.EntityData.ParentYangName = "interfaces"
    staticNeighbors.EntityData.SegmentPath = "static-neighbors"
    staticNeighbors.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    staticNeighbors.EntityData.NamespaceTable = ietf.GetNamespaces()
    staticNeighbors.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    staticNeighbors.EntityData.Children = make(map[string]types.YChild)
    staticNeighbors.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range staticNeighbors.Neighbor {
        staticNeighbors.EntityData.Children[types.GetSegmentPath(&staticNeighbors.Neighbor[i])] = types.YChild{"Neighbor", &staticNeighbors.Neighbor[i]}
    }
    staticNeighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticNeighbors.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor
// Specify a neighbor router.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Neighbor cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Neighbor poll interval. The type is interface{} with range: 1..65535. Units
    // are seconds.
    PollInterval interface{}

    // Neighbor priority for DR election. The type is interface{} with range:
    // 1..255.
    Priority interface{}
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_StaticNeighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "ietf"
    neighbor.EntityData.ParentYangName = "static-neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + "[address='" + fmt.Sprintf("%v", neighbor.Address) + "']"
    neighbor.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    neighbor.EntityData.NamespaceTable = ietf.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["address"] = types.YLeaf{"Address", neighbor.Address}
    neighbor.EntityData.Leafs["cost"] = types.YLeaf{"Cost", neighbor.Cost}
    neighbor.EntityData.Leafs["poll-interval"] = types.YLeaf{"PollInterval", neighbor.PollInterval}
    neighbor.EntityData.Leafs["priority"] = types.YLeaf{"Priority", neighbor.Priority}
    return &(neighbor.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute
// Fast-reroute configuration.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LFA configuration.
    Lfa RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa
}

func (fastReroute *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "ietf"
    fastReroute.EntityData.ParentYangName = "interfaces"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = ietf.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    fastReroute.EntityData.Children = make(map[string]types.YChild)
    fastReroute.EntityData.Children["lfa"] = types.YChild{"Lfa", &fastReroute.Lfa}
    fastReroute.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fastReroute.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa
// LFA configuration.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prevent the interface to be used as backup. The type is bool.
    CandidateDisabled interface{}

    // Activates LFA. This model assumes activation of per-prefix LFA. The type is
    // bool.
    Enabled interface{}

    // Remote LFA configuration.
    RemoteLfa RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa
}

func (lfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa) GetEntityData() *types.CommonEntityData {
    lfa.EntityData.YFilter = lfa.YFilter
    lfa.EntityData.YangName = "lfa"
    lfa.EntityData.BundleName = "ietf"
    lfa.EntityData.ParentYangName = "fast-reroute"
    lfa.EntityData.SegmentPath = "lfa"
    lfa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    lfa.EntityData.NamespaceTable = ietf.GetNamespaces()
    lfa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    lfa.EntityData.Children = make(map[string]types.YChild)
    lfa.EntityData.Children["remote-lfa"] = types.YChild{"RemoteLfa", &lfa.RemoteLfa}
    lfa.EntityData.Leafs = make(map[string]types.YLeaf)
    lfa.EntityData.Leafs["candidate-disabled"] = types.YLeaf{"CandidateDisabled", lfa.CandidateDisabled}
    lfa.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", lfa.Enabled}
    return &(lfa.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa
// Remote LFA configuration.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Activates remote LFA. The type is bool.
    Enabled interface{}
}

func (remoteLfa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_FastReroute_Lfa_RemoteLfa) GetEntityData() *types.CommonEntityData {
    remoteLfa.EntityData.YFilter = remoteLfa.YFilter
    remoteLfa.EntityData.YangName = "remote-lfa"
    remoteLfa.EntityData.BundleName = "ietf"
    remoteLfa.EntityData.ParentYangName = "lfa"
    remoteLfa.EntityData.SegmentPath = "remote-lfa"
    remoteLfa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    remoteLfa.EntityData.NamespaceTable = ietf.GetNamespaces()
    remoteLfa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    remoteLfa.EntityData.Children = make(map[string]types.YChild)
    remoteLfa.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteLfa.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", remoteLfa.Enabled}
    return &(remoteLfa.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity
// TTL security check.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable TTL security check. The type is bool.
    Enable interface{}

    // Maximum number of hops that a OSPF packet may have traveled. The type is
    // interface{} with range: 1..254.
    Hops interface{}
}

func (ttlSecurity *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_TtlSecurity) GetEntityData() *types.CommonEntityData {
    ttlSecurity.EntityData.YFilter = ttlSecurity.YFilter
    ttlSecurity.EntityData.YangName = "ttl-security"
    ttlSecurity.EntityData.BundleName = "ietf"
    ttlSecurity.EntityData.ParentYangName = "interfaces"
    ttlSecurity.EntityData.SegmentPath = "ttl-security"
    ttlSecurity.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ttlSecurity.EntityData.NamespaceTable = ietf.GetNamespaces()
    ttlSecurity.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ttlSecurity.EntityData.Children = make(map[string]types.YChild)
    ttlSecurity.EntityData.Leafs = make(map[string]types.YLeaf)
    ttlSecurity.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ttlSecurity.Enable}
    ttlSecurity.EntityData.Leafs["hops"] = types.YLeaf{"Hops", ttlSecurity.Hops}
    return &(ttlSecurity.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication
// Authentication configuration.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA name. The type is string.
    Sa interface{}

    // key-chain name. The type is string. Refers to key_chain.KeyChains_Name
    KeyChain interface{}

    // Key string in ASCII format. The type is string.
    Key interface{}

    // Cryptographic algorithm associated with key.
    CryptoAlgorithm RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm
}

func (authentication *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "ietf"
    authentication.EntityData.ParentYangName = "interfaces"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    authentication.EntityData.NamespaceTable = ietf.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["crypto-algorithm"] = types.YChild{"CryptoAlgorithm", &authentication.CryptoAlgorithm}
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["sa"] = types.YLeaf{"Sa", authentication.Sa}
    authentication.EntityData.Leafs["key-chain"] = types.YLeaf{"KeyChain", authentication.KeyChain}
    authentication.EntityData.Leafs["key"] = types.YLeaf{"Key", authentication.Key}
    return &(authentication.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm
// Cryptographic algorithm associated with key.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // The HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // The MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // The SHA-1 algorithm. The type is interface{}.
    Sha1 interface{}

    // HMAC-SHA-1 authentication algorithm. The type is interface{}.
    HmacSha1 interface{}

    // HMAC-SHA-256 authentication algorithm. The type is interface{}.
    HmacSha256 interface{}

    // HMAC-SHA-384 authentication algorithm. The type is interface{}.
    HmacSha384 interface{}

    // HMAC-SHA-512 authentication algorithm. The type is interface{}.
    HmacSha512 interface{}
}

func (cryptoAlgorithm *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Authentication_CryptoAlgorithm) GetEntityData() *types.CommonEntityData {
    cryptoAlgorithm.EntityData.YFilter = cryptoAlgorithm.YFilter
    cryptoAlgorithm.EntityData.YangName = "crypto-algorithm"
    cryptoAlgorithm.EntityData.BundleName = "ietf"
    cryptoAlgorithm.EntityData.ParentYangName = "authentication"
    cryptoAlgorithm.EntityData.SegmentPath = "crypto-algorithm"
    cryptoAlgorithm.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    cryptoAlgorithm.EntityData.NamespaceTable = ietf.GetNamespaces()
    cryptoAlgorithm.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    cryptoAlgorithm.EntityData.Children = make(map[string]types.YChild)
    cryptoAlgorithm.EntityData.Leafs = make(map[string]types.YLeaf)
    cryptoAlgorithm.EntityData.Leafs["hmac-sha1-12"] = types.YLeaf{"HmacSha112", cryptoAlgorithm.HmacSha112}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha1-20"] = types.YLeaf{"HmacSha120", cryptoAlgorithm.HmacSha120}
    cryptoAlgorithm.EntityData.Leafs["md5"] = types.YLeaf{"Md5", cryptoAlgorithm.Md5}
    cryptoAlgorithm.EntityData.Leafs["sha-1"] = types.YLeaf{"Sha1", cryptoAlgorithm.Sha1}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-1"] = types.YLeaf{"HmacSha1", cryptoAlgorithm.HmacSha1}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-256"] = types.YLeaf{"HmacSha256", cryptoAlgorithm.HmacSha256}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-384"] = types.YLeaf{"HmacSha384", cryptoAlgorithm.HmacSha384}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-512"] = types.YLeaf{"HmacSha512", cryptoAlgorithm.HmacSha512}
    return &(cryptoAlgorithm.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor
// List of OSPF neighbors.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborId interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Designated Router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Dr interface{}

    // Backup Designated Router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Bdr interface{}

    // OSPF neighbor state. The type is NbrStateType.
    State interface{}
}

func (neighbor *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "ietf"
    neighbor.EntityData.ParentYangName = "interfaces"
    neighbor.EntityData.SegmentPath = "neighbor" + "[neighbor-id='" + fmt.Sprintf("%v", neighbor.NeighborId) + "']"
    neighbor.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    neighbor.EntityData.NamespaceTable = ietf.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["neighbor-id"] = types.YLeaf{"NeighborId", neighbor.NeighborId}
    neighbor.EntityData.Leafs["address"] = types.YLeaf{"Address", neighbor.Address}
    neighbor.EntityData.Leafs["dr"] = types.YLeaf{"Dr", neighbor.Dr}
    neighbor.EntityData.Leafs["bdr"] = types.YLeaf{"Bdr", neighbor.Bdr}
    neighbor.EntityData.Leafs["state"] = types.YLeaf{"State", neighbor.State}
    return &(neighbor.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas
// List OSPF link scope LSA databases
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF link scope LSA type. The type is interface{}
    // with range: 0..255.
    LsaType interface{}

    // List of OSPF link scope LSAs. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa.
    LinkScopeLsa []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa
}

func (linkScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas) GetEntityData() *types.CommonEntityData {
    linkScopeLsas.EntityData.YFilter = linkScopeLsas.YFilter
    linkScopeLsas.EntityData.YangName = "link-scope-lsas"
    linkScopeLsas.EntityData.BundleName = "ietf"
    linkScopeLsas.EntityData.ParentYangName = "interfaces"
    linkScopeLsas.EntityData.SegmentPath = "link-scope-lsas" + "[lsa-type='" + fmt.Sprintf("%v", linkScopeLsas.LsaType) + "']"
    linkScopeLsas.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    linkScopeLsas.EntityData.NamespaceTable = ietf.GetNamespaces()
    linkScopeLsas.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    linkScopeLsas.EntityData.Children = make(map[string]types.YChild)
    linkScopeLsas.EntityData.Children["link-scope-lsa"] = types.YChild{"LinkScopeLsa", nil}
    for i := range linkScopeLsas.LinkScopeLsa {
        linkScopeLsas.EntityData.Children[types.GetSegmentPath(&linkScopeLsas.LinkScopeLsa[i])] = types.YChild{"LinkScopeLsa", &linkScopeLsas.LinkScopeLsa[i]}
    }
    linkScopeLsas.EntityData.Leafs = make(map[string]types.YLeaf)
    linkScopeLsas.EntityData.Leafs["lsa-type"] = types.YLeaf{"LsaType", linkScopeLsas.LsaType}
    return &(linkScopeLsas.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa
// List of OSPF link scope LSAs
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSA ID. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or int with range: 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is string with pattern: b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    RawData interface{}

    // OSPFv2 LSA.
    Ospfv2 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2

    // OSPFv3 LSA.
    Ospfv3 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3
}

func (linkScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa) GetEntityData() *types.CommonEntityData {
    linkScopeLsa.EntityData.YFilter = linkScopeLsa.YFilter
    linkScopeLsa.EntityData.YangName = "link-scope-lsa"
    linkScopeLsa.EntityData.BundleName = "ietf"
    linkScopeLsa.EntityData.ParentYangName = "link-scope-lsas"
    linkScopeLsa.EntityData.SegmentPath = "link-scope-lsa" + "[lsa-id='" + fmt.Sprintf("%v", linkScopeLsa.LsaId) + "']" + "[adv-router='" + fmt.Sprintf("%v", linkScopeLsa.AdvRouter) + "']"
    linkScopeLsa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    linkScopeLsa.EntityData.NamespaceTable = ietf.GetNamespaces()
    linkScopeLsa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    linkScopeLsa.EntityData.Children = make(map[string]types.YChild)
    linkScopeLsa.EntityData.Children["ospfv2"] = types.YChild{"Ospfv2", &linkScopeLsa.Ospfv2}
    linkScopeLsa.EntityData.Children["ospfv3"] = types.YChild{"Ospfv3", &linkScopeLsa.Ospfv3}
    linkScopeLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    linkScopeLsa.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", linkScopeLsa.LsaId}
    linkScopeLsa.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", linkScopeLsa.AdvRouter}
    linkScopeLsa.EntityData.Leafs["decoded-completed"] = types.YLeaf{"DecodedCompleted", linkScopeLsa.DecodedCompleted}
    linkScopeLsa.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", linkScopeLsa.RawData}
    return &(linkScopeLsa.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2
// OSPFv2 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header

    // Decoded OSPFv2 LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2) GetEntityData() *types.CommonEntityData {
    ospfv2.EntityData.YFilter = ospfv2.YFilter
    ospfv2.EntityData.YangName = "ospfv2"
    ospfv2.EntityData.BundleName = "ietf"
    ospfv2.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2.EntityData.SegmentPath = "ospfv2"
    ospfv2.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ospfv2.EntityData.NamespaceTable = ietf.GetNamespaces()
    ospfv2.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ospfv2.EntityData.Children = make(map[string]types.YChild)
    ospfv2.EntityData.Children["header"] = types.YChild{"Header", &ospfv2.Header}
    ospfv2.EntityData.Children["body"] = types.YChild{"Body", &ospfv2.Body}
    ospfv2.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header
// Decoded OSPFv2 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Options interface{}

    // LSA ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    OpaqueType interface{}

    // Opaque id. The type is interface{} with range: 0..16777215. This attribute
    // is mandatory.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type_ interface{}

    // LSA advertising router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "ietf"
    header.EntityData.ParentYangName = "ospfv2"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    header.EntityData.NamespaceTable = ietf.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["options"] = types.YLeaf{"Options", header.Options}
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["opaque-type"] = types.YLeaf{"OpaqueType", header.OpaqueType}
    header.EntityData.Leafs["opaque-id"] = types.YLeaf{"OpaqueId", header.OpaqueId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    return &(header.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body
// Decoded OSPFv2 LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network

    // Summary LSA.
    Summary RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary

    // External LSA.
    External RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External

    // Opaque LSA.
    Opaque RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body) GetEntityData() *types.CommonEntityData {
    body.EntityData.YFilter = body.YFilter
    body.EntityData.YangName = "body"
    body.EntityData.BundleName = "ietf"
    body.EntityData.ParentYangName = "ospfv2"
    body.EntityData.SegmentPath = "body"
    body.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    body.EntityData.NamespaceTable = ietf.GetNamespaces()
    body.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    body.EntityData.Children = make(map[string]types.YChild)
    body.EntityData.Children["router"] = types.YChild{"Router", &body.Router}
    body.EntityData.Children["network"] = types.YChild{"Network", &body.Network}
    body.EntityData.Children["summary"] = types.YChild{"Summary", &body.Summary}
    body.EntityData.Children["external"] = types.YChild{"External", &body.External}
    body.EntityData.Children["opaque"] = types.YChild{"Opaque", &body.Opaque}
    body.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(body.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router) GetEntityData() *types.CommonEntityData {
    router.EntityData.YFilter = router.YFilter
    router.EntityData.YangName = "router"
    router.EntityData.BundleName = "ietf"
    router.EntityData.ParentYangName = "body"
    router.EntityData.SegmentPath = "router"
    router.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    router.EntityData.NamespaceTable = ietf.GetNamespaces()
    router.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    router.EntityData.Children = make(map[string]types.YChild)
    router.EntityData.Children["link"] = types.YChild{"Link", nil}
    for i := range router.Link {
        router.EntityData.Children[types.GetSegmentPath(&router.Link[i])] = types.YChild{"Link", &router.Link[i]}
    }
    router.EntityData.Leafs = make(map[string]types.YLeaf)
    router.EntityData.Leafs["flags"] = types.YLeaf{"Flags", router.Flags}
    router.EntityData.Leafs["num-of-links"] = types.YLeaf{"NumOfLinks", router.NumOfLinks}
    return &(router.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    LinkId interface{}

    // This attribute is a key. Link data. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or int with range: 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "ietf"
    link.EntityData.ParentYangName = "router"
    link.EntityData.SegmentPath = "link" + "[link-id='" + fmt.Sprintf("%v", link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", link.LinkData) + "']"
    link.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    link.EntityData.NamespaceTable = ietf.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    link.EntityData.Children = make(map[string]types.YChild)
    link.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range link.Topology {
        link.EntityData.Children[types.GetSegmentPath(&link.Topology[i])] = types.YChild{"Topology", &link.Topology[i]}
    }
    link.EntityData.Leafs = make(map[string]types.YLeaf)
    link.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", link.LinkId}
    link.EntityData.Leafs["link-data"] = types.YLeaf{"LinkData", link.LinkData}
    link.EntityData.Leafs["type"] = types.YLeaf{"Type_", link.Type_}
    return &(link.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Router_Link_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "link"
    topology.EntityData.SegmentPath = "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", topology.MtId}
    topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", topology.Metric}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "ietf"
    network.EntityData.ParentYangName = "body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    network.EntityData.NamespaceTable = ietf.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", network.NetworkMask}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary
// Summary LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "ietf"
    summary.EntityData.ParentYangName = "body"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    summary.EntityData.NamespaceTable = ietf.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range summary.Topology {
        summary.EntityData.Children[types.GetSegmentPath(&summary.Topology[i])] = types.YChild{"Topology", &summary.Topology[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", summary.NetworkMask}
    return &(summary.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Summary_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "summary"
    topology.EntityData.SegmentPath = "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", topology.MtId}
    topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", topology.Metric}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External
// External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External) GetEntityData() *types.CommonEntityData {
    external.EntityData.YFilter = external.YFilter
    external.EntityData.YangName = "external"
    external.EntityData.BundleName = "ietf"
    external.EntityData.ParentYangName = "body"
    external.EntityData.SegmentPath = "external"
    external.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    external.EntityData.NamespaceTable = ietf.GetNamespaces()
    external.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    external.EntityData.Children = make(map[string]types.YChild)
    external.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range external.Topology {
        external.EntityData.Children[types.GetSegmentPath(&external.Topology[i])] = types.YChild{"Topology", &external.Topology[i]}
    }
    external.EntityData.Leafs = make(map[string]types.YLeaf)
    external.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", external.NetworkMask}
    return &(external.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Forwarding address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_External_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "external"
    topology.EntityData.SegmentPath = "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", topology.MtId}
    topology.EntityData.Leafs["flags"] = types.YLeaf{"Flags", topology.Flags}
    topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", topology.Metric}
    topology.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", topology.ForwardingAddress}
    topology.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", topology.ExternalRouteTag}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque
// Opaque LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv.
    UnknownTlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv

    // Router address TLV.
    RouterAddressTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv

    // Link TLV.
    LinkTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque) GetEntityData() *types.CommonEntityData {
    opaque.EntityData.YFilter = opaque.YFilter
    opaque.EntityData.YangName = "opaque"
    opaque.EntityData.BundleName = "ietf"
    opaque.EntityData.ParentYangName = "body"
    opaque.EntityData.SegmentPath = "opaque"
    opaque.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    opaque.EntityData.NamespaceTable = ietf.GetNamespaces()
    opaque.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    opaque.EntityData.Children = make(map[string]types.YChild)
    opaque.EntityData.Children["unknown-tlv"] = types.YChild{"UnknownTlv", nil}
    for i := range opaque.UnknownTlv {
        opaque.EntityData.Children[types.GetSegmentPath(&opaque.UnknownTlv[i])] = types.YChild{"UnknownTlv", &opaque.UnknownTlv[i]}
    }
    opaque.EntityData.Children["router-address-tlv"] = types.YChild{"RouterAddressTlv", &opaque.RouterAddressTlv}
    opaque.EntityData.Children["link-tlv"] = types.YChild{"LinkTlv", &opaque.LinkTlv}
    opaque.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(opaque.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv
// Unknown TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type_ interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetEntityData() *types.CommonEntityData {
    unknownTlv.EntityData.YFilter = unknownTlv.YFilter
    unknownTlv.EntityData.YangName = "unknown-tlv"
    unknownTlv.EntityData.BundleName = "ietf"
    unknownTlv.EntityData.ParentYangName = "opaque"
    unknownTlv.EntityData.SegmentPath = "unknown-tlv" + "[type='" + fmt.Sprintf("%v", unknownTlv.Type_) + "']"
    unknownTlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    unknownTlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    unknownTlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    unknownTlv.EntityData.Children = make(map[string]types.YChild)
    unknownTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownTlv.EntityData.Leafs["type"] = types.YLeaf{"Type_", unknownTlv.Type_}
    unknownTlv.EntityData.Leafs["length"] = types.YLeaf{"Length", unknownTlv.Length}
    unknownTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownTlv.Value}
    return &(unknownTlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv
// Router address TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterAddress interface{}
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetEntityData() *types.CommonEntityData {
    routerAddressTlv.EntityData.YFilter = routerAddressTlv.YFilter
    routerAddressTlv.EntityData.YangName = "router-address-tlv"
    routerAddressTlv.EntityData.BundleName = "ietf"
    routerAddressTlv.EntityData.ParentYangName = "opaque"
    routerAddressTlv.EntityData.SegmentPath = "router-address-tlv"
    routerAddressTlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routerAddressTlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    routerAddressTlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routerAddressTlv.EntityData.Children = make(map[string]types.YChild)
    routerAddressTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    routerAddressTlv.EntityData.Leafs["router-address"] = types.YLeaf{"RouterAddress", routerAddressTlv.RouterAddress}
    return &(routerAddressTlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv
// Link TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    LinkType interface{}

    // Link ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'
    // This attribute is mandatory..
    LinkId interface{}

    // List of local interface IPv4 addresses. The type is slice of string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalIfIpv4Addr []interface{}

    // List of remote interface IPv4 addresses. The type is slice of string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalRemoteIpv4Addr []interface{}

    // TE metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Maximum bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxBandwidth interface{}

    // Maximum reservable bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxReservableBandwidth interface{}

    // Unreserved bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    UnreservedBandwidth interface{}

    // Administrative group/Resource class/Color. The type is interface{} with
    // range: 0..4294967295.
    AdminGroup interface{}

    // Unknown sub-TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv.
    UnknownSubtlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetEntityData() *types.CommonEntityData {
    linkTlv.EntityData.YFilter = linkTlv.YFilter
    linkTlv.EntityData.YangName = "link-tlv"
    linkTlv.EntityData.BundleName = "ietf"
    linkTlv.EntityData.ParentYangName = "opaque"
    linkTlv.EntityData.SegmentPath = "link-tlv"
    linkTlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    linkTlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    linkTlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    linkTlv.EntityData.Children = make(map[string]types.YChild)
    linkTlv.EntityData.Children["unknown-subtlv"] = types.YChild{"UnknownSubtlv", nil}
    for i := range linkTlv.UnknownSubtlv {
        linkTlv.EntityData.Children[types.GetSegmentPath(&linkTlv.UnknownSubtlv[i])] = types.YChild{"UnknownSubtlv", &linkTlv.UnknownSubtlv[i]}
    }
    linkTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    linkTlv.EntityData.Leafs["link-type"] = types.YLeaf{"LinkType", linkTlv.LinkType}
    linkTlv.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", linkTlv.LinkId}
    linkTlv.EntityData.Leafs["local-if-ipv4-addr"] = types.YLeaf{"LocalIfIpv4Addr", linkTlv.LocalIfIpv4Addr}
    linkTlv.EntityData.Leafs["local-remote-ipv4-addr"] = types.YLeaf{"LocalRemoteIpv4Addr", linkTlv.LocalRemoteIpv4Addr}
    linkTlv.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", linkTlv.TeMetric}
    linkTlv.EntityData.Leafs["max-bandwidth"] = types.YLeaf{"MaxBandwidth", linkTlv.MaxBandwidth}
    linkTlv.EntityData.Leafs["max-reservable-bandwidth"] = types.YLeaf{"MaxReservableBandwidth", linkTlv.MaxReservableBandwidth}
    linkTlv.EntityData.Leafs["unreserved-bandwidth"] = types.YLeaf{"UnreservedBandwidth", linkTlv.UnreservedBandwidth}
    linkTlv.EntityData.Leafs["admin-group"] = types.YLeaf{"AdminGroup", linkTlv.AdminGroup}
    return &(linkTlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
// Unknown sub-TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type_ interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetEntityData() *types.CommonEntityData {
    unknownSubtlv.EntityData.YFilter = unknownSubtlv.YFilter
    unknownSubtlv.EntityData.YangName = "unknown-subtlv"
    unknownSubtlv.EntityData.BundleName = "ietf"
    unknownSubtlv.EntityData.ParentYangName = "link-tlv"
    unknownSubtlv.EntityData.SegmentPath = "unknown-subtlv" + "[type='" + fmt.Sprintf("%v", unknownSubtlv.Type_) + "']"
    unknownSubtlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    unknownSubtlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    unknownSubtlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    unknownSubtlv.EntityData.Children = make(map[string]types.YChild)
    unknownSubtlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownSubtlv.EntityData.Leafs["type"] = types.YLeaf{"Type_", unknownSubtlv.Type_}
    unknownSubtlv.EntityData.Leafs["length"] = types.YLeaf{"Length", unknownSubtlv.Length}
    unknownSubtlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownSubtlv.Value}
    return &(unknownSubtlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3
// OSPFv3 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header

    // Decoded OSPF LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3) GetEntityData() *types.CommonEntityData {
    ospfv3.EntityData.YFilter = ospfv3.YFilter
    ospfv3.EntityData.YangName = "ospfv3"
    ospfv3.EntityData.BundleName = "ietf"
    ospfv3.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3.EntityData.SegmentPath = "ospfv3"
    ospfv3.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ospfv3.EntityData.NamespaceTable = ietf.GetNamespaces()
    ospfv3.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ospfv3.EntityData.Children = make(map[string]types.YChild)
    ospfv3.EntityData.Children["header"] = types.YChild{"Header", &ospfv3.Header}
    ospfv3.EntityData.Children["body"] = types.YChild{"Body", &ospfv3.Body}
    ospfv3.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header
// Decoded OSPFv3 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is interface{} with range: 0..4294967295. This attribute
    // is mandatory.
    LsaId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type_ interface{}

    // LSA advertising router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "ietf"
    header.EntityData.ParentYangName = "ospfv3"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    header.EntityData.NamespaceTable = ietf.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    header.EntityData.Leafs["options"] = types.YLeaf{"Options", header.Options}
    return &(header.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body
// Decoded OSPF LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network

    // Inter-Area-Prefix LSA.
    InterAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix

    // Inter-Area-Router LSA.
    InterAreaRouter RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter

    // AS-External LSA.
    AsExternal RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal

    // NSSA LSA.
    Nssa RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa

    // Link LSA.
    Link RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link

    // Intra-Area-Prefix LSA.
    IntraAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body) GetEntityData() *types.CommonEntityData {
    body.EntityData.YFilter = body.YFilter
    body.EntityData.YangName = "body"
    body.EntityData.BundleName = "ietf"
    body.EntityData.ParentYangName = "ospfv3"
    body.EntityData.SegmentPath = "body"
    body.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    body.EntityData.NamespaceTable = ietf.GetNamespaces()
    body.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    body.EntityData.Children = make(map[string]types.YChild)
    body.EntityData.Children["router"] = types.YChild{"Router", &body.Router}
    body.EntityData.Children["network"] = types.YChild{"Network", &body.Network}
    body.EntityData.Children["inter-area-prefix"] = types.YChild{"InterAreaPrefix", &body.InterAreaPrefix}
    body.EntityData.Children["inter-area-router"] = types.YChild{"InterAreaRouter", &body.InterAreaRouter}
    body.EntityData.Children["as-external"] = types.YChild{"AsExternal", &body.AsExternal}
    body.EntityData.Children["nssa"] = types.YChild{"Nssa", &body.Nssa}
    body.EntityData.Children["link"] = types.YChild{"Link", &body.Link}
    body.EntityData.Children["intra-area-prefix"] = types.YChild{"IntraAreaPrefix", &body.IntraAreaPrefix}
    body.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(body.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Flags interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router) GetEntityData() *types.CommonEntityData {
    router.EntityData.YFilter = router.YFilter
    router.EntityData.YangName = "router"
    router.EntityData.BundleName = "ietf"
    router.EntityData.ParentYangName = "body"
    router.EntityData.SegmentPath = "router"
    router.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    router.EntityData.NamespaceTable = ietf.GetNamespaces()
    router.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    router.EntityData.Children = make(map[string]types.YChild)
    router.EntityData.Children["link"] = types.YChild{"Link", nil}
    for i := range router.Link {
        router.EntityData.Children[types.GetSegmentPath(&router.Link[i])] = types.YChild{"Link", &router.Link[i]}
    }
    router.EntityData.Leafs = make(map[string]types.YLeaf)
    router.EntityData.Leafs["flags"] = types.YLeaf{"Flags", router.Flags}
    router.EntityData.Leafs["options"] = types.YLeaf{"Options", router.Options}
    return &(router.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor Interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor Router ID. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Router_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "ietf"
    link.EntityData.ParentYangName = "router"
    link.EntityData.SegmentPath = "link" + "[interface-id='" + fmt.Sprintf("%v", link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", link.NeighborRouterId) + "']"
    link.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    link.EntityData.NamespaceTable = ietf.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    link.EntityData.Children = make(map[string]types.YChild)
    link.EntityData.Leafs = make(map[string]types.YLeaf)
    link.EntityData.Leafs["interface-id"] = types.YLeaf{"InterfaceId", link.InterfaceId}
    link.EntityData.Leafs["neighbor-interface-id"] = types.YLeaf{"NeighborInterfaceId", link.NeighborInterfaceId}
    link.EntityData.Leafs["neighbor-router-id"] = types.YLeaf{"NeighborRouterId", link.NeighborRouterId}
    link.EntityData.Leafs["type"] = types.YLeaf{"Type_", link.Type_}
    link.EntityData.Leafs["metric"] = types.YLeaf{"Metric", link.Metric}
    return &(link.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "ietf"
    network.EntityData.ParentYangName = "body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    network.EntityData.NamespaceTable = ietf.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["options"] = types.YLeaf{"Options", network.Options}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix
// Inter-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaPrefix) GetEntityData() *types.CommonEntityData {
    interAreaPrefix.EntityData.YFilter = interAreaPrefix.YFilter
    interAreaPrefix.EntityData.YangName = "inter-area-prefix"
    interAreaPrefix.EntityData.BundleName = "ietf"
    interAreaPrefix.EntityData.ParentYangName = "body"
    interAreaPrefix.EntityData.SegmentPath = "inter-area-prefix"
    interAreaPrefix.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interAreaPrefix.EntityData.NamespaceTable = ietf.GetNamespaces()
    interAreaPrefix.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interAreaPrefix.EntityData.Children = make(map[string]types.YChild)
    interAreaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    interAreaPrefix.EntityData.Leafs["metric"] = types.YLeaf{"Metric", interAreaPrefix.Metric}
    interAreaPrefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", interAreaPrefix.Prefix}
    interAreaPrefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", interAreaPrefix.PrefixOptions}
    return &(interAreaPrefix.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter
// Inter-Area-Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // The Router ID of the router being described by the LSA. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    DestinationRouterId interface{}
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_InterAreaRouter) GetEntityData() *types.CommonEntityData {
    interAreaRouter.EntityData.YFilter = interAreaRouter.YFilter
    interAreaRouter.EntityData.YangName = "inter-area-router"
    interAreaRouter.EntityData.BundleName = "ietf"
    interAreaRouter.EntityData.ParentYangName = "body"
    interAreaRouter.EntityData.SegmentPath = "inter-area-router"
    interAreaRouter.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interAreaRouter.EntityData.NamespaceTable = ietf.GetNamespaces()
    interAreaRouter.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interAreaRouter.EntityData.Children = make(map[string]types.YChild)
    interAreaRouter.EntityData.Leafs = make(map[string]types.YLeaf)
    interAreaRouter.EntityData.Leafs["options"] = types.YLeaf{"Options", interAreaRouter.Options}
    interAreaRouter.EntityData.Leafs["metric"] = types.YLeaf{"Metric", interAreaRouter.Metric}
    interAreaRouter.EntityData.Leafs["destination-router-id"] = types.YLeaf{"DestinationRouterId", interAreaRouter.DestinationRouterId}
    return &(interAreaRouter.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal
// AS-External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_AsExternal) GetEntityData() *types.CommonEntityData {
    asExternal.EntityData.YFilter = asExternal.YFilter
    asExternal.EntityData.YangName = "as-external"
    asExternal.EntityData.BundleName = "ietf"
    asExternal.EntityData.ParentYangName = "body"
    asExternal.EntityData.SegmentPath = "as-external"
    asExternal.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    asExternal.EntityData.NamespaceTable = ietf.GetNamespaces()
    asExternal.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    asExternal.EntityData.Children = make(map[string]types.YChild)
    asExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    asExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", asExternal.Metric}
    asExternal.EntityData.Leafs["flags"] = types.YLeaf{"Flags", asExternal.Flags}
    asExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", asExternal.ReferencedLsType}
    asExternal.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", asExternal.Prefix}
    asExternal.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", asExternal.PrefixOptions}
    asExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", asExternal.ForwardingAddress}
    asExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", asExternal.ExternalRouteTag}
    asExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", asExternal.ReferencedLinkStateId}
    return &(asExternal.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa
// NSSA LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "ietf"
    nssa.EntityData.ParentYangName = "body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    nssa.EntityData.NamespaceTable = ietf.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    nssa.EntityData.Children = make(map[string]types.YChild)
    nssa.EntityData.Leafs = make(map[string]types.YLeaf)
    nssa.EntityData.Leafs["metric"] = types.YLeaf{"Metric", nssa.Metric}
    nssa.EntityData.Leafs["flags"] = types.YLeaf{"Flags", nssa.Flags}
    nssa.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", nssa.ReferencedLsType}
    nssa.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", nssa.Prefix}
    nssa.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", nssa.PrefixOptions}
    nssa.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", nssa.ForwardingAddress}
    nssa.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", nssa.ExternalRouteTag}
    nssa.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", nssa.ReferencedLinkStateId}
    return &(nssa.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link
// Link LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Priority of the interface. The type is interface{} with range:
    // 0..255.
    RtrPriority interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "ietf"
    link.EntityData.ParentYangName = "body"
    link.EntityData.SegmentPath = "link"
    link.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    link.EntityData.NamespaceTable = ietf.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    link.EntityData.Children = make(map[string]types.YChild)
    link.EntityData.Children["prefix-list"] = types.YChild{"PrefixList", nil}
    for i := range link.PrefixList {
        link.EntityData.Children[types.GetSegmentPath(&link.PrefixList[i])] = types.YChild{"PrefixList", &link.PrefixList[i]}
    }
    link.EntityData.Leafs = make(map[string]types.YLeaf)
    link.EntityData.Leafs["rtr-priority"] = types.YLeaf{"RtrPriority", link.RtrPriority}
    link.EntityData.Leafs["options"] = types.YLeaf{"Options", link.Options}
    link.EntityData.Leafs["link-local-interface-address"] = types.YLeaf{"LinkLocalInterfaceAddress", link.LinkLocalInterfaceAddress}
    link.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", link.NumOfPrefixes}
    return &(link.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_Link_PrefixList) GetEntityData() *types.CommonEntityData {
    prefixList.EntityData.YFilter = prefixList.YFilter
    prefixList.EntityData.YangName = "prefix-list"
    prefixList.EntityData.BundleName = "ietf"
    prefixList.EntityData.ParentYangName = "link"
    prefixList.EntityData.SegmentPath = "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
    prefixList.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    prefixList.EntityData.NamespaceTable = ietf.GetNamespaces()
    prefixList.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    prefixList.EntityData.Children = make(map[string]types.YChild)
    prefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixList.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", prefixList.Prefix}
    prefixList.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", prefixList.PrefixOptions}
    return &(prefixList.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix
// Intra-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetEntityData() *types.CommonEntityData {
    intraAreaPrefix.EntityData.YFilter = intraAreaPrefix.YFilter
    intraAreaPrefix.EntityData.YangName = "intra-area-prefix"
    intraAreaPrefix.EntityData.BundleName = "ietf"
    intraAreaPrefix.EntityData.ParentYangName = "body"
    intraAreaPrefix.EntityData.SegmentPath = "intra-area-prefix"
    intraAreaPrefix.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    intraAreaPrefix.EntityData.NamespaceTable = ietf.GetNamespaces()
    intraAreaPrefix.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    intraAreaPrefix.EntityData.Children = make(map[string]types.YChild)
    intraAreaPrefix.EntityData.Children["prefix-list"] = types.YChild{"PrefixList", nil}
    for i := range intraAreaPrefix.PrefixList {
        intraAreaPrefix.EntityData.Children[types.GetSegmentPath(&intraAreaPrefix.PrefixList[i])] = types.YChild{"PrefixList", &intraAreaPrefix.PrefixList[i]}
    }
    intraAreaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    intraAreaPrefix.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", intraAreaPrefix.ReferencedLsType}
    intraAreaPrefix.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", intraAreaPrefix.ReferencedLinkStateId}
    intraAreaPrefix.EntityData.Leafs["referenced-adv-router"] = types.YLeaf{"ReferencedAdvRouter", intraAreaPrefix.ReferencedAdvRouter}
    intraAreaPrefix.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", intraAreaPrefix.NumOfPrefixes}
    return &(intraAreaPrefix.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_LinkScopeLsas_LinkScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetEntityData() *types.CommonEntityData {
    prefixList.EntityData.YFilter = prefixList.YFilter
    prefixList.EntityData.YangName = "prefix-list"
    prefixList.EntityData.BundleName = "ietf"
    prefixList.EntityData.ParentYangName = "intra-area-prefix"
    prefixList.EntityData.SegmentPath = "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
    prefixList.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    prefixList.EntityData.NamespaceTable = ietf.GetNamespaces()
    prefixList.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    prefixList.EntityData.Children = make(map[string]types.YChild)
    prefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixList.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", prefixList.Prefix}
    prefixList.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", prefixList.PrefixOptions}
    prefixList.EntityData.Leafs["metric"] = types.YLeaf{"Metric", prefixList.Metric}
    return &(prefixList.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology
// OSPF interface topology.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string. Refers to routing.Routing_RoutingInstance_Ribs_Rib_Name
    Name interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "interfaces"
    topology.EntityData.SegmentPath = "topology" + "[name='" + fmt.Sprintf("%v", topology.Name) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["name"] = types.YLeaf{"Name", topology.Name}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType represents Network type.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType string

const (
    // Specify OSPF broadcast multi-access network.
    RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType_broadcast RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType = "broadcast"

    // Specify OSPF Non-Broadcast Multi-Access
    // (NBMA) network.
    RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType_non_broadcast RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType = "non-broadcast"

    // Specify OSPF point-to-multipoint network.
    RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType_point_to_multipoint RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType = "point-to-multipoint"

    // Specify OSPF point-to-point network.
    RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType_point_to_point RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interfaces_NetworkType = "point-to-point"
)

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas
// List OSPF area scope LSA databases
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF area scope LSA type. The type is interface{}
    // with range: 0..255.
    LsaType interface{}

    // List of OSPF area scope LSAs. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa.
    AreaScopeLsa []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa
}

func (areaScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas) GetEntityData() *types.CommonEntityData {
    areaScopeLsas.EntityData.YFilter = areaScopeLsas.YFilter
    areaScopeLsas.EntityData.YangName = "area-scope-lsas"
    areaScopeLsas.EntityData.BundleName = "ietf"
    areaScopeLsas.EntityData.ParentYangName = "area"
    areaScopeLsas.EntityData.SegmentPath = "area-scope-lsas" + "[lsa-type='" + fmt.Sprintf("%v", areaScopeLsas.LsaType) + "']"
    areaScopeLsas.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    areaScopeLsas.EntityData.NamespaceTable = ietf.GetNamespaces()
    areaScopeLsas.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    areaScopeLsas.EntityData.Children = make(map[string]types.YChild)
    areaScopeLsas.EntityData.Children["area-scope-lsa"] = types.YChild{"AreaScopeLsa", nil}
    for i := range areaScopeLsas.AreaScopeLsa {
        areaScopeLsas.EntityData.Children[types.GetSegmentPath(&areaScopeLsas.AreaScopeLsa[i])] = types.YChild{"AreaScopeLsa", &areaScopeLsas.AreaScopeLsa[i]}
    }
    areaScopeLsas.EntityData.Leafs = make(map[string]types.YLeaf)
    areaScopeLsas.EntityData.Leafs["lsa-type"] = types.YLeaf{"LsaType", areaScopeLsas.LsaType}
    return &(areaScopeLsas.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa
// List of OSPF area scope LSAs
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSA ID. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or int with range: 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is string with pattern: b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    RawData interface{}

    // OSPFv2 LSA.
    Ospfv2 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2

    // OSPFv3 LSA.
    Ospfv3 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3
}

func (areaScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa) GetEntityData() *types.CommonEntityData {
    areaScopeLsa.EntityData.YFilter = areaScopeLsa.YFilter
    areaScopeLsa.EntityData.YangName = "area-scope-lsa"
    areaScopeLsa.EntityData.BundleName = "ietf"
    areaScopeLsa.EntityData.ParentYangName = "area-scope-lsas"
    areaScopeLsa.EntityData.SegmentPath = "area-scope-lsa" + "[lsa-id='" + fmt.Sprintf("%v", areaScopeLsa.LsaId) + "']" + "[adv-router='" + fmt.Sprintf("%v", areaScopeLsa.AdvRouter) + "']"
    areaScopeLsa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    areaScopeLsa.EntityData.NamespaceTable = ietf.GetNamespaces()
    areaScopeLsa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    areaScopeLsa.EntityData.Children = make(map[string]types.YChild)
    areaScopeLsa.EntityData.Children["ospfv2"] = types.YChild{"Ospfv2", &areaScopeLsa.Ospfv2}
    areaScopeLsa.EntityData.Children["ospfv3"] = types.YChild{"Ospfv3", &areaScopeLsa.Ospfv3}
    areaScopeLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    areaScopeLsa.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", areaScopeLsa.LsaId}
    areaScopeLsa.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", areaScopeLsa.AdvRouter}
    areaScopeLsa.EntityData.Leafs["decoded-completed"] = types.YLeaf{"DecodedCompleted", areaScopeLsa.DecodedCompleted}
    areaScopeLsa.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", areaScopeLsa.RawData}
    return &(areaScopeLsa.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2
// OSPFv2 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header

    // Decoded OSPFv2 LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2) GetEntityData() *types.CommonEntityData {
    ospfv2.EntityData.YFilter = ospfv2.YFilter
    ospfv2.EntityData.YangName = "ospfv2"
    ospfv2.EntityData.BundleName = "ietf"
    ospfv2.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2.EntityData.SegmentPath = "ospfv2"
    ospfv2.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ospfv2.EntityData.NamespaceTable = ietf.GetNamespaces()
    ospfv2.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ospfv2.EntityData.Children = make(map[string]types.YChild)
    ospfv2.EntityData.Children["header"] = types.YChild{"Header", &ospfv2.Header}
    ospfv2.EntityData.Children["body"] = types.YChild{"Body", &ospfv2.Body}
    ospfv2.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header
// Decoded OSPFv2 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Options interface{}

    // LSA ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    OpaqueType interface{}

    // Opaque id. The type is interface{} with range: 0..16777215. This attribute
    // is mandatory.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type_ interface{}

    // LSA advertising router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "ietf"
    header.EntityData.ParentYangName = "ospfv2"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    header.EntityData.NamespaceTable = ietf.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["options"] = types.YLeaf{"Options", header.Options}
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["opaque-type"] = types.YLeaf{"OpaqueType", header.OpaqueType}
    header.EntityData.Leafs["opaque-id"] = types.YLeaf{"OpaqueId", header.OpaqueId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    return &(header.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body
// Decoded OSPFv2 LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network

    // Summary LSA.
    Summary RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary

    // External LSA.
    External RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External

    // Opaque LSA.
    Opaque RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body) GetEntityData() *types.CommonEntityData {
    body.EntityData.YFilter = body.YFilter
    body.EntityData.YangName = "body"
    body.EntityData.BundleName = "ietf"
    body.EntityData.ParentYangName = "ospfv2"
    body.EntityData.SegmentPath = "body"
    body.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    body.EntityData.NamespaceTable = ietf.GetNamespaces()
    body.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    body.EntityData.Children = make(map[string]types.YChild)
    body.EntityData.Children["router"] = types.YChild{"Router", &body.Router}
    body.EntityData.Children["network"] = types.YChild{"Network", &body.Network}
    body.EntityData.Children["summary"] = types.YChild{"Summary", &body.Summary}
    body.EntityData.Children["external"] = types.YChild{"External", &body.External}
    body.EntityData.Children["opaque"] = types.YChild{"Opaque", &body.Opaque}
    body.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(body.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router) GetEntityData() *types.CommonEntityData {
    router.EntityData.YFilter = router.YFilter
    router.EntityData.YangName = "router"
    router.EntityData.BundleName = "ietf"
    router.EntityData.ParentYangName = "body"
    router.EntityData.SegmentPath = "router"
    router.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    router.EntityData.NamespaceTable = ietf.GetNamespaces()
    router.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    router.EntityData.Children = make(map[string]types.YChild)
    router.EntityData.Children["link"] = types.YChild{"Link", nil}
    for i := range router.Link {
        router.EntityData.Children[types.GetSegmentPath(&router.Link[i])] = types.YChild{"Link", &router.Link[i]}
    }
    router.EntityData.Leafs = make(map[string]types.YLeaf)
    router.EntityData.Leafs["flags"] = types.YLeaf{"Flags", router.Flags}
    router.EntityData.Leafs["num-of-links"] = types.YLeaf{"NumOfLinks", router.NumOfLinks}
    return &(router.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    LinkId interface{}

    // This attribute is a key. Link data. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or int with range: 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "ietf"
    link.EntityData.ParentYangName = "router"
    link.EntityData.SegmentPath = "link" + "[link-id='" + fmt.Sprintf("%v", link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", link.LinkData) + "']"
    link.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    link.EntityData.NamespaceTable = ietf.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    link.EntityData.Children = make(map[string]types.YChild)
    link.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range link.Topology {
        link.EntityData.Children[types.GetSegmentPath(&link.Topology[i])] = types.YChild{"Topology", &link.Topology[i]}
    }
    link.EntityData.Leafs = make(map[string]types.YLeaf)
    link.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", link.LinkId}
    link.EntityData.Leafs["link-data"] = types.YLeaf{"LinkData", link.LinkData}
    link.EntityData.Leafs["type"] = types.YLeaf{"Type_", link.Type_}
    return &(link.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Router_Link_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "link"
    topology.EntityData.SegmentPath = "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", topology.MtId}
    topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", topology.Metric}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "ietf"
    network.EntityData.ParentYangName = "body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    network.EntityData.NamespaceTable = ietf.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", network.NetworkMask}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary
// Summary LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "ietf"
    summary.EntityData.ParentYangName = "body"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    summary.EntityData.NamespaceTable = ietf.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range summary.Topology {
        summary.EntityData.Children[types.GetSegmentPath(&summary.Topology[i])] = types.YChild{"Topology", &summary.Topology[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", summary.NetworkMask}
    return &(summary.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Summary_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "summary"
    topology.EntityData.SegmentPath = "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", topology.MtId}
    topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", topology.Metric}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External
// External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External) GetEntityData() *types.CommonEntityData {
    external.EntityData.YFilter = external.YFilter
    external.EntityData.YangName = "external"
    external.EntityData.BundleName = "ietf"
    external.EntityData.ParentYangName = "body"
    external.EntityData.SegmentPath = "external"
    external.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    external.EntityData.NamespaceTable = ietf.GetNamespaces()
    external.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    external.EntityData.Children = make(map[string]types.YChild)
    external.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range external.Topology {
        external.EntityData.Children[types.GetSegmentPath(&external.Topology[i])] = types.YChild{"Topology", &external.Topology[i]}
    }
    external.EntityData.Leafs = make(map[string]types.YLeaf)
    external.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", external.NetworkMask}
    return &(external.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Forwarding address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_External_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "external"
    topology.EntityData.SegmentPath = "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", topology.MtId}
    topology.EntityData.Leafs["flags"] = types.YLeaf{"Flags", topology.Flags}
    topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", topology.Metric}
    topology.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", topology.ForwardingAddress}
    topology.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", topology.ExternalRouteTag}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque
// Opaque LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv.
    UnknownTlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv

    // Router address TLV.
    RouterAddressTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv

    // Link TLV.
    LinkTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque) GetEntityData() *types.CommonEntityData {
    opaque.EntityData.YFilter = opaque.YFilter
    opaque.EntityData.YangName = "opaque"
    opaque.EntityData.BundleName = "ietf"
    opaque.EntityData.ParentYangName = "body"
    opaque.EntityData.SegmentPath = "opaque"
    opaque.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    opaque.EntityData.NamespaceTable = ietf.GetNamespaces()
    opaque.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    opaque.EntityData.Children = make(map[string]types.YChild)
    opaque.EntityData.Children["unknown-tlv"] = types.YChild{"UnknownTlv", nil}
    for i := range opaque.UnknownTlv {
        opaque.EntityData.Children[types.GetSegmentPath(&opaque.UnknownTlv[i])] = types.YChild{"UnknownTlv", &opaque.UnknownTlv[i]}
    }
    opaque.EntityData.Children["router-address-tlv"] = types.YChild{"RouterAddressTlv", &opaque.RouterAddressTlv}
    opaque.EntityData.Children["link-tlv"] = types.YChild{"LinkTlv", &opaque.LinkTlv}
    opaque.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(opaque.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv
// Unknown TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type_ interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetEntityData() *types.CommonEntityData {
    unknownTlv.EntityData.YFilter = unknownTlv.YFilter
    unknownTlv.EntityData.YangName = "unknown-tlv"
    unknownTlv.EntityData.BundleName = "ietf"
    unknownTlv.EntityData.ParentYangName = "opaque"
    unknownTlv.EntityData.SegmentPath = "unknown-tlv" + "[type='" + fmt.Sprintf("%v", unknownTlv.Type_) + "']"
    unknownTlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    unknownTlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    unknownTlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    unknownTlv.EntityData.Children = make(map[string]types.YChild)
    unknownTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownTlv.EntityData.Leafs["type"] = types.YLeaf{"Type_", unknownTlv.Type_}
    unknownTlv.EntityData.Leafs["length"] = types.YLeaf{"Length", unknownTlv.Length}
    unknownTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownTlv.Value}
    return &(unknownTlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv
// Router address TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterAddress interface{}
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetEntityData() *types.CommonEntityData {
    routerAddressTlv.EntityData.YFilter = routerAddressTlv.YFilter
    routerAddressTlv.EntityData.YangName = "router-address-tlv"
    routerAddressTlv.EntityData.BundleName = "ietf"
    routerAddressTlv.EntityData.ParentYangName = "opaque"
    routerAddressTlv.EntityData.SegmentPath = "router-address-tlv"
    routerAddressTlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routerAddressTlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    routerAddressTlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routerAddressTlv.EntityData.Children = make(map[string]types.YChild)
    routerAddressTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    routerAddressTlv.EntityData.Leafs["router-address"] = types.YLeaf{"RouterAddress", routerAddressTlv.RouterAddress}
    return &(routerAddressTlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv
// Link TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    LinkType interface{}

    // Link ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'
    // This attribute is mandatory..
    LinkId interface{}

    // List of local interface IPv4 addresses. The type is slice of string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalIfIpv4Addr []interface{}

    // List of remote interface IPv4 addresses. The type is slice of string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalRemoteIpv4Addr []interface{}

    // TE metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Maximum bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxBandwidth interface{}

    // Maximum reservable bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxReservableBandwidth interface{}

    // Unreserved bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    UnreservedBandwidth interface{}

    // Administrative group/Resource class/Color. The type is interface{} with
    // range: 0..4294967295.
    AdminGroup interface{}

    // Unknown sub-TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv.
    UnknownSubtlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetEntityData() *types.CommonEntityData {
    linkTlv.EntityData.YFilter = linkTlv.YFilter
    linkTlv.EntityData.YangName = "link-tlv"
    linkTlv.EntityData.BundleName = "ietf"
    linkTlv.EntityData.ParentYangName = "opaque"
    linkTlv.EntityData.SegmentPath = "link-tlv"
    linkTlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    linkTlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    linkTlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    linkTlv.EntityData.Children = make(map[string]types.YChild)
    linkTlv.EntityData.Children["unknown-subtlv"] = types.YChild{"UnknownSubtlv", nil}
    for i := range linkTlv.UnknownSubtlv {
        linkTlv.EntityData.Children[types.GetSegmentPath(&linkTlv.UnknownSubtlv[i])] = types.YChild{"UnknownSubtlv", &linkTlv.UnknownSubtlv[i]}
    }
    linkTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    linkTlv.EntityData.Leafs["link-type"] = types.YLeaf{"LinkType", linkTlv.LinkType}
    linkTlv.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", linkTlv.LinkId}
    linkTlv.EntityData.Leafs["local-if-ipv4-addr"] = types.YLeaf{"LocalIfIpv4Addr", linkTlv.LocalIfIpv4Addr}
    linkTlv.EntityData.Leafs["local-remote-ipv4-addr"] = types.YLeaf{"LocalRemoteIpv4Addr", linkTlv.LocalRemoteIpv4Addr}
    linkTlv.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", linkTlv.TeMetric}
    linkTlv.EntityData.Leafs["max-bandwidth"] = types.YLeaf{"MaxBandwidth", linkTlv.MaxBandwidth}
    linkTlv.EntityData.Leafs["max-reservable-bandwidth"] = types.YLeaf{"MaxReservableBandwidth", linkTlv.MaxReservableBandwidth}
    linkTlv.EntityData.Leafs["unreserved-bandwidth"] = types.YLeaf{"UnreservedBandwidth", linkTlv.UnreservedBandwidth}
    linkTlv.EntityData.Leafs["admin-group"] = types.YLeaf{"AdminGroup", linkTlv.AdminGroup}
    return &(linkTlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
// Unknown sub-TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type_ interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetEntityData() *types.CommonEntityData {
    unknownSubtlv.EntityData.YFilter = unknownSubtlv.YFilter
    unknownSubtlv.EntityData.YangName = "unknown-subtlv"
    unknownSubtlv.EntityData.BundleName = "ietf"
    unknownSubtlv.EntityData.ParentYangName = "link-tlv"
    unknownSubtlv.EntityData.SegmentPath = "unknown-subtlv" + "[type='" + fmt.Sprintf("%v", unknownSubtlv.Type_) + "']"
    unknownSubtlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    unknownSubtlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    unknownSubtlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    unknownSubtlv.EntityData.Children = make(map[string]types.YChild)
    unknownSubtlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownSubtlv.EntityData.Leafs["type"] = types.YLeaf{"Type_", unknownSubtlv.Type_}
    unknownSubtlv.EntityData.Leafs["length"] = types.YLeaf{"Length", unknownSubtlv.Length}
    unknownSubtlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownSubtlv.Value}
    return &(unknownSubtlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3
// OSPFv3 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header

    // Decoded OSPF LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3) GetEntityData() *types.CommonEntityData {
    ospfv3.EntityData.YFilter = ospfv3.YFilter
    ospfv3.EntityData.YangName = "ospfv3"
    ospfv3.EntityData.BundleName = "ietf"
    ospfv3.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3.EntityData.SegmentPath = "ospfv3"
    ospfv3.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ospfv3.EntityData.NamespaceTable = ietf.GetNamespaces()
    ospfv3.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ospfv3.EntityData.Children = make(map[string]types.YChild)
    ospfv3.EntityData.Children["header"] = types.YChild{"Header", &ospfv3.Header}
    ospfv3.EntityData.Children["body"] = types.YChild{"Body", &ospfv3.Body}
    ospfv3.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header
// Decoded OSPFv3 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is interface{} with range: 0..4294967295. This attribute
    // is mandatory.
    LsaId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type_ interface{}

    // LSA advertising router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "ietf"
    header.EntityData.ParentYangName = "ospfv3"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    header.EntityData.NamespaceTable = ietf.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    header.EntityData.Leafs["options"] = types.YLeaf{"Options", header.Options}
    return &(header.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body
// Decoded OSPF LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network

    // Inter-Area-Prefix LSA.
    InterAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix

    // Inter-Area-Router LSA.
    InterAreaRouter RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter

    // AS-External LSA.
    AsExternal RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal

    // NSSA LSA.
    Nssa RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa

    // Link LSA.
    Link RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link

    // Intra-Area-Prefix LSA.
    IntraAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body) GetEntityData() *types.CommonEntityData {
    body.EntityData.YFilter = body.YFilter
    body.EntityData.YangName = "body"
    body.EntityData.BundleName = "ietf"
    body.EntityData.ParentYangName = "ospfv3"
    body.EntityData.SegmentPath = "body"
    body.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    body.EntityData.NamespaceTable = ietf.GetNamespaces()
    body.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    body.EntityData.Children = make(map[string]types.YChild)
    body.EntityData.Children["router"] = types.YChild{"Router", &body.Router}
    body.EntityData.Children["network"] = types.YChild{"Network", &body.Network}
    body.EntityData.Children["inter-area-prefix"] = types.YChild{"InterAreaPrefix", &body.InterAreaPrefix}
    body.EntityData.Children["inter-area-router"] = types.YChild{"InterAreaRouter", &body.InterAreaRouter}
    body.EntityData.Children["as-external"] = types.YChild{"AsExternal", &body.AsExternal}
    body.EntityData.Children["nssa"] = types.YChild{"Nssa", &body.Nssa}
    body.EntityData.Children["link"] = types.YChild{"Link", &body.Link}
    body.EntityData.Children["intra-area-prefix"] = types.YChild{"IntraAreaPrefix", &body.IntraAreaPrefix}
    body.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(body.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Flags interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router) GetEntityData() *types.CommonEntityData {
    router.EntityData.YFilter = router.YFilter
    router.EntityData.YangName = "router"
    router.EntityData.BundleName = "ietf"
    router.EntityData.ParentYangName = "body"
    router.EntityData.SegmentPath = "router"
    router.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    router.EntityData.NamespaceTable = ietf.GetNamespaces()
    router.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    router.EntityData.Children = make(map[string]types.YChild)
    router.EntityData.Children["link"] = types.YChild{"Link", nil}
    for i := range router.Link {
        router.EntityData.Children[types.GetSegmentPath(&router.Link[i])] = types.YChild{"Link", &router.Link[i]}
    }
    router.EntityData.Leafs = make(map[string]types.YLeaf)
    router.EntityData.Leafs["flags"] = types.YLeaf{"Flags", router.Flags}
    router.EntityData.Leafs["options"] = types.YLeaf{"Options", router.Options}
    return &(router.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor Interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor Router ID. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Router_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "ietf"
    link.EntityData.ParentYangName = "router"
    link.EntityData.SegmentPath = "link" + "[interface-id='" + fmt.Sprintf("%v", link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", link.NeighborRouterId) + "']"
    link.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    link.EntityData.NamespaceTable = ietf.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    link.EntityData.Children = make(map[string]types.YChild)
    link.EntityData.Leafs = make(map[string]types.YLeaf)
    link.EntityData.Leafs["interface-id"] = types.YLeaf{"InterfaceId", link.InterfaceId}
    link.EntityData.Leafs["neighbor-interface-id"] = types.YLeaf{"NeighborInterfaceId", link.NeighborInterfaceId}
    link.EntityData.Leafs["neighbor-router-id"] = types.YLeaf{"NeighborRouterId", link.NeighborRouterId}
    link.EntityData.Leafs["type"] = types.YLeaf{"Type_", link.Type_}
    link.EntityData.Leafs["metric"] = types.YLeaf{"Metric", link.Metric}
    return &(link.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "ietf"
    network.EntityData.ParentYangName = "body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    network.EntityData.NamespaceTable = ietf.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["options"] = types.YLeaf{"Options", network.Options}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix
// Inter-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaPrefix) GetEntityData() *types.CommonEntityData {
    interAreaPrefix.EntityData.YFilter = interAreaPrefix.YFilter
    interAreaPrefix.EntityData.YangName = "inter-area-prefix"
    interAreaPrefix.EntityData.BundleName = "ietf"
    interAreaPrefix.EntityData.ParentYangName = "body"
    interAreaPrefix.EntityData.SegmentPath = "inter-area-prefix"
    interAreaPrefix.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interAreaPrefix.EntityData.NamespaceTable = ietf.GetNamespaces()
    interAreaPrefix.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interAreaPrefix.EntityData.Children = make(map[string]types.YChild)
    interAreaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    interAreaPrefix.EntityData.Leafs["metric"] = types.YLeaf{"Metric", interAreaPrefix.Metric}
    interAreaPrefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", interAreaPrefix.Prefix}
    interAreaPrefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", interAreaPrefix.PrefixOptions}
    return &(interAreaPrefix.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter
// Inter-Area-Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // The Router ID of the router being described by the LSA. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    DestinationRouterId interface{}
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_InterAreaRouter) GetEntityData() *types.CommonEntityData {
    interAreaRouter.EntityData.YFilter = interAreaRouter.YFilter
    interAreaRouter.EntityData.YangName = "inter-area-router"
    interAreaRouter.EntityData.BundleName = "ietf"
    interAreaRouter.EntityData.ParentYangName = "body"
    interAreaRouter.EntityData.SegmentPath = "inter-area-router"
    interAreaRouter.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interAreaRouter.EntityData.NamespaceTable = ietf.GetNamespaces()
    interAreaRouter.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interAreaRouter.EntityData.Children = make(map[string]types.YChild)
    interAreaRouter.EntityData.Leafs = make(map[string]types.YLeaf)
    interAreaRouter.EntityData.Leafs["options"] = types.YLeaf{"Options", interAreaRouter.Options}
    interAreaRouter.EntityData.Leafs["metric"] = types.YLeaf{"Metric", interAreaRouter.Metric}
    interAreaRouter.EntityData.Leafs["destination-router-id"] = types.YLeaf{"DestinationRouterId", interAreaRouter.DestinationRouterId}
    return &(interAreaRouter.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal
// AS-External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_AsExternal) GetEntityData() *types.CommonEntityData {
    asExternal.EntityData.YFilter = asExternal.YFilter
    asExternal.EntityData.YangName = "as-external"
    asExternal.EntityData.BundleName = "ietf"
    asExternal.EntityData.ParentYangName = "body"
    asExternal.EntityData.SegmentPath = "as-external"
    asExternal.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    asExternal.EntityData.NamespaceTable = ietf.GetNamespaces()
    asExternal.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    asExternal.EntityData.Children = make(map[string]types.YChild)
    asExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    asExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", asExternal.Metric}
    asExternal.EntityData.Leafs["flags"] = types.YLeaf{"Flags", asExternal.Flags}
    asExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", asExternal.ReferencedLsType}
    asExternal.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", asExternal.Prefix}
    asExternal.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", asExternal.PrefixOptions}
    asExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", asExternal.ForwardingAddress}
    asExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", asExternal.ExternalRouteTag}
    asExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", asExternal.ReferencedLinkStateId}
    return &(asExternal.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa
// NSSA LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "ietf"
    nssa.EntityData.ParentYangName = "body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    nssa.EntityData.NamespaceTable = ietf.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    nssa.EntityData.Children = make(map[string]types.YChild)
    nssa.EntityData.Leafs = make(map[string]types.YLeaf)
    nssa.EntityData.Leafs["metric"] = types.YLeaf{"Metric", nssa.Metric}
    nssa.EntityData.Leafs["flags"] = types.YLeaf{"Flags", nssa.Flags}
    nssa.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", nssa.ReferencedLsType}
    nssa.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", nssa.Prefix}
    nssa.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", nssa.PrefixOptions}
    nssa.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", nssa.ForwardingAddress}
    nssa.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", nssa.ExternalRouteTag}
    nssa.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", nssa.ReferencedLinkStateId}
    return &(nssa.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link
// Link LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Priority of the interface. The type is interface{} with range:
    // 0..255.
    RtrPriority interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "ietf"
    link.EntityData.ParentYangName = "body"
    link.EntityData.SegmentPath = "link"
    link.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    link.EntityData.NamespaceTable = ietf.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    link.EntityData.Children = make(map[string]types.YChild)
    link.EntityData.Children["prefix-list"] = types.YChild{"PrefixList", nil}
    for i := range link.PrefixList {
        link.EntityData.Children[types.GetSegmentPath(&link.PrefixList[i])] = types.YChild{"PrefixList", &link.PrefixList[i]}
    }
    link.EntityData.Leafs = make(map[string]types.YLeaf)
    link.EntityData.Leafs["rtr-priority"] = types.YLeaf{"RtrPriority", link.RtrPriority}
    link.EntityData.Leafs["options"] = types.YLeaf{"Options", link.Options}
    link.EntityData.Leafs["link-local-interface-address"] = types.YLeaf{"LinkLocalInterfaceAddress", link.LinkLocalInterfaceAddress}
    link.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", link.NumOfPrefixes}
    return &(link.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_Link_PrefixList) GetEntityData() *types.CommonEntityData {
    prefixList.EntityData.YFilter = prefixList.YFilter
    prefixList.EntityData.YangName = "prefix-list"
    prefixList.EntityData.BundleName = "ietf"
    prefixList.EntityData.ParentYangName = "link"
    prefixList.EntityData.SegmentPath = "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
    prefixList.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    prefixList.EntityData.NamespaceTable = ietf.GetNamespaces()
    prefixList.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    prefixList.EntityData.Children = make(map[string]types.YChild)
    prefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixList.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", prefixList.Prefix}
    prefixList.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", prefixList.PrefixOptions}
    return &(prefixList.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix
// Intra-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetEntityData() *types.CommonEntityData {
    intraAreaPrefix.EntityData.YFilter = intraAreaPrefix.YFilter
    intraAreaPrefix.EntityData.YangName = "intra-area-prefix"
    intraAreaPrefix.EntityData.BundleName = "ietf"
    intraAreaPrefix.EntityData.ParentYangName = "body"
    intraAreaPrefix.EntityData.SegmentPath = "intra-area-prefix"
    intraAreaPrefix.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    intraAreaPrefix.EntityData.NamespaceTable = ietf.GetNamespaces()
    intraAreaPrefix.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    intraAreaPrefix.EntityData.Children = make(map[string]types.YChild)
    intraAreaPrefix.EntityData.Children["prefix-list"] = types.YChild{"PrefixList", nil}
    for i := range intraAreaPrefix.PrefixList {
        intraAreaPrefix.EntityData.Children[types.GetSegmentPath(&intraAreaPrefix.PrefixList[i])] = types.YChild{"PrefixList", &intraAreaPrefix.PrefixList[i]}
    }
    intraAreaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    intraAreaPrefix.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", intraAreaPrefix.ReferencedLsType}
    intraAreaPrefix.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", intraAreaPrefix.ReferencedLinkStateId}
    intraAreaPrefix.EntityData.Leafs["referenced-adv-router"] = types.YLeaf{"ReferencedAdvRouter", intraAreaPrefix.ReferencedAdvRouter}
    intraAreaPrefix.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", intraAreaPrefix.NumOfPrefixes}
    return &(intraAreaPrefix.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AreaScopeLsas_AreaScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetEntityData() *types.CommonEntityData {
    prefixList.EntityData.YFilter = prefixList.YFilter
    prefixList.EntityData.YangName = "prefix-list"
    prefixList.EntityData.BundleName = "ietf"
    prefixList.EntityData.ParentYangName = "intra-area-prefix"
    prefixList.EntityData.SegmentPath = "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
    prefixList.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    prefixList.EntityData.NamespaceTable = ietf.GetNamespaces()
    prefixList.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    prefixList.EntityData.Children = make(map[string]types.YChild)
    prefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixList.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", prefixList.Prefix}
    prefixList.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", prefixList.PrefixOptions}
    prefixList.EntityData.Leafs["metric"] = types.YLeaf{"Metric", prefixList.Metric}
    return &(prefixList.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas
// List OSPF AS scope LSA databases
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF AS scope LSA type. The type is interface{}
    // with range: 0..255.
    LsaType interface{}

    // List of OSPF AS scope LSAs. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa.
    AsScopeLsa []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa
}

func (asScopeLsas *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas) GetEntityData() *types.CommonEntityData {
    asScopeLsas.EntityData.YFilter = asScopeLsas.YFilter
    asScopeLsas.EntityData.YangName = "as-scope-lsas"
    asScopeLsas.EntityData.BundleName = "ietf"
    asScopeLsas.EntityData.ParentYangName = "instance"
    asScopeLsas.EntityData.SegmentPath = "as-scope-lsas" + "[lsa-type='" + fmt.Sprintf("%v", asScopeLsas.LsaType) + "']"
    asScopeLsas.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    asScopeLsas.EntityData.NamespaceTable = ietf.GetNamespaces()
    asScopeLsas.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    asScopeLsas.EntityData.Children = make(map[string]types.YChild)
    asScopeLsas.EntityData.Children["as-scope-lsa"] = types.YChild{"AsScopeLsa", nil}
    for i := range asScopeLsas.AsScopeLsa {
        asScopeLsas.EntityData.Children[types.GetSegmentPath(&asScopeLsas.AsScopeLsa[i])] = types.YChild{"AsScopeLsa", &asScopeLsas.AsScopeLsa[i]}
    }
    asScopeLsas.EntityData.Leafs = make(map[string]types.YLeaf)
    asScopeLsas.EntityData.Leafs["lsa-type"] = types.YLeaf{"LsaType", asScopeLsas.LsaType}
    return &(asScopeLsas.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa
// List of OSPF AS scope LSAs
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSA ID. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or int with range: 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is string with pattern: b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    RawData interface{}

    // OSPFv2 LSA.
    Ospfv2 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2

    // OSPFv3 LSA.
    Ospfv3 RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3
}

func (asScopeLsa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa) GetEntityData() *types.CommonEntityData {
    asScopeLsa.EntityData.YFilter = asScopeLsa.YFilter
    asScopeLsa.EntityData.YangName = "as-scope-lsa"
    asScopeLsa.EntityData.BundleName = "ietf"
    asScopeLsa.EntityData.ParentYangName = "as-scope-lsas"
    asScopeLsa.EntityData.SegmentPath = "as-scope-lsa" + "[lsa-id='" + fmt.Sprintf("%v", asScopeLsa.LsaId) + "']" + "[adv-router='" + fmt.Sprintf("%v", asScopeLsa.AdvRouter) + "']"
    asScopeLsa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    asScopeLsa.EntityData.NamespaceTable = ietf.GetNamespaces()
    asScopeLsa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    asScopeLsa.EntityData.Children = make(map[string]types.YChild)
    asScopeLsa.EntityData.Children["ospfv2"] = types.YChild{"Ospfv2", &asScopeLsa.Ospfv2}
    asScopeLsa.EntityData.Children["ospfv3"] = types.YChild{"Ospfv3", &asScopeLsa.Ospfv3}
    asScopeLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    asScopeLsa.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", asScopeLsa.LsaId}
    asScopeLsa.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", asScopeLsa.AdvRouter}
    asScopeLsa.EntityData.Leafs["decoded-completed"] = types.YLeaf{"DecodedCompleted", asScopeLsa.DecodedCompleted}
    asScopeLsa.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", asScopeLsa.RawData}
    return &(asScopeLsa.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2
// OSPFv2 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header

    // Decoded OSPFv2 LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body
}

func (ospfv2 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2) GetEntityData() *types.CommonEntityData {
    ospfv2.EntityData.YFilter = ospfv2.YFilter
    ospfv2.EntityData.YangName = "ospfv2"
    ospfv2.EntityData.BundleName = "ietf"
    ospfv2.EntityData.ParentYangName = "as-scope-lsa"
    ospfv2.EntityData.SegmentPath = "ospfv2"
    ospfv2.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ospfv2.EntityData.NamespaceTable = ietf.GetNamespaces()
    ospfv2.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ospfv2.EntityData.Children = make(map[string]types.YChild)
    ospfv2.EntityData.Children["header"] = types.YChild{"Header", &ospfv2.Header}
    ospfv2.EntityData.Children["body"] = types.YChild{"Body", &ospfv2.Body}
    ospfv2.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header
// Decoded OSPFv2 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Options interface{}

    // LSA ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    OpaqueType interface{}

    // Opaque id. The type is interface{} with range: 0..16777215. This attribute
    // is mandatory.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type_ interface{}

    // LSA advertising router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "ietf"
    header.EntityData.ParentYangName = "ospfv2"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    header.EntityData.NamespaceTable = ietf.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["options"] = types.YLeaf{"Options", header.Options}
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["opaque-type"] = types.YLeaf{"OpaqueType", header.OpaqueType}
    header.EntityData.Leafs["opaque-id"] = types.YLeaf{"OpaqueId", header.OpaqueId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    return &(header.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body
// Decoded OSPFv2 LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network

    // Summary LSA.
    Summary RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary

    // External LSA.
    External RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External

    // Opaque LSA.
    Opaque RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body) GetEntityData() *types.CommonEntityData {
    body.EntityData.YFilter = body.YFilter
    body.EntityData.YangName = "body"
    body.EntityData.BundleName = "ietf"
    body.EntityData.ParentYangName = "ospfv2"
    body.EntityData.SegmentPath = "body"
    body.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    body.EntityData.NamespaceTable = ietf.GetNamespaces()
    body.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    body.EntityData.Children = make(map[string]types.YChild)
    body.EntityData.Children["router"] = types.YChild{"Router", &body.Router}
    body.EntityData.Children["network"] = types.YChild{"Network", &body.Network}
    body.EntityData.Children["summary"] = types.YChild{"Summary", &body.Summary}
    body.EntityData.Children["external"] = types.YChild{"External", &body.External}
    body.EntityData.Children["opaque"] = types.YChild{"Opaque", &body.Opaque}
    body.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(body.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router) GetEntityData() *types.CommonEntityData {
    router.EntityData.YFilter = router.YFilter
    router.EntityData.YangName = "router"
    router.EntityData.BundleName = "ietf"
    router.EntityData.ParentYangName = "body"
    router.EntityData.SegmentPath = "router"
    router.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    router.EntityData.NamespaceTable = ietf.GetNamespaces()
    router.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    router.EntityData.Children = make(map[string]types.YChild)
    router.EntityData.Children["link"] = types.YChild{"Link", nil}
    for i := range router.Link {
        router.EntityData.Children[types.GetSegmentPath(&router.Link[i])] = types.YChild{"Link", &router.Link[i]}
    }
    router.EntityData.Leafs = make(map[string]types.YLeaf)
    router.EntityData.Leafs["flags"] = types.YLeaf{"Flags", router.Flags}
    router.EntityData.Leafs["num-of-links"] = types.YLeaf{"NumOfLinks", router.NumOfLinks}
    return &(router.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    LinkId interface{}

    // This attribute is a key. Link data. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or int with range: 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "ietf"
    link.EntityData.ParentYangName = "router"
    link.EntityData.SegmentPath = "link" + "[link-id='" + fmt.Sprintf("%v", link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", link.LinkData) + "']"
    link.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    link.EntityData.NamespaceTable = ietf.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    link.EntityData.Children = make(map[string]types.YChild)
    link.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range link.Topology {
        link.EntityData.Children[types.GetSegmentPath(&link.Topology[i])] = types.YChild{"Topology", &link.Topology[i]}
    }
    link.EntityData.Leafs = make(map[string]types.YLeaf)
    link.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", link.LinkId}
    link.EntityData.Leafs["link-data"] = types.YLeaf{"LinkData", link.LinkData}
    link.EntityData.Leafs["type"] = types.YLeaf{"Type_", link.Type_}
    return &(link.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Router_Link_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "link"
    topology.EntityData.SegmentPath = "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", topology.MtId}
    topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", topology.Metric}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "ietf"
    network.EntityData.ParentYangName = "body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    network.EntityData.NamespaceTable = ietf.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", network.NetworkMask}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary
// Summary LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology
}

func (summary *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "ietf"
    summary.EntityData.ParentYangName = "body"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    summary.EntityData.NamespaceTable = ietf.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range summary.Topology {
        summary.EntityData.Children[types.GetSegmentPath(&summary.Topology[i])] = types.YChild{"Topology", &summary.Topology[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", summary.NetworkMask}
    return &(summary.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Summary_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "summary"
    topology.EntityData.SegmentPath = "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", topology.MtId}
    topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", topology.Metric}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External
// External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP address mask for the network. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // Topology specific information. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology.
    Topology []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology
}

func (external *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External) GetEntityData() *types.CommonEntityData {
    external.EntityData.YFilter = external.YFilter
    external.EntityData.YangName = "external"
    external.EntityData.BundleName = "ietf"
    external.EntityData.ParentYangName = "body"
    external.EntityData.SegmentPath = "external"
    external.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    external.EntityData.NamespaceTable = ietf.GetNamespaces()
    external.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    external.EntityData.Children = make(map[string]types.YChild)
    external.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range external.Topology {
        external.EntityData.Children[types.GetSegmentPath(&external.Topology[i])] = types.YChild{"Topology", &external.Topology[i]}
    }
    external.EntityData.Leafs = make(map[string]types.YLeaf)
    external.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", external.NetworkMask}
    return &(external.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology
// Topology specific information.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MT-ID for topology enabled on the link. The
    // type is interface{} with range: 0..255.
    MtId interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Metric for the topology. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Forwarding address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_External_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "external"
    topology.EntityData.SegmentPath = "topology" + "[mt-id='" + fmt.Sprintf("%v", topology.MtId) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", topology.MtId}
    topology.EntityData.Leafs["flags"] = types.YLeaf{"Flags", topology.Flags}
    topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", topology.Metric}
    topology.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", topology.ForwardingAddress}
    topology.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", topology.ExternalRouteTag}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque
// Opaque LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv.
    UnknownTlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv

    // Router address TLV.
    RouterAddressTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv

    // Link TLV.
    LinkTlv RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv
}

func (opaque *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque) GetEntityData() *types.CommonEntityData {
    opaque.EntityData.YFilter = opaque.YFilter
    opaque.EntityData.YangName = "opaque"
    opaque.EntityData.BundleName = "ietf"
    opaque.EntityData.ParentYangName = "body"
    opaque.EntityData.SegmentPath = "opaque"
    opaque.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    opaque.EntityData.NamespaceTable = ietf.GetNamespaces()
    opaque.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    opaque.EntityData.Children = make(map[string]types.YChild)
    opaque.EntityData.Children["unknown-tlv"] = types.YChild{"UnknownTlv", nil}
    for i := range opaque.UnknownTlv {
        opaque.EntityData.Children[types.GetSegmentPath(&opaque.UnknownTlv[i])] = types.YChild{"UnknownTlv", &opaque.UnknownTlv[i]}
    }
    opaque.EntityData.Children["router-address-tlv"] = types.YChild{"RouterAddressTlv", &opaque.RouterAddressTlv}
    opaque.EntityData.Children["link-tlv"] = types.YChild{"LinkTlv", &opaque.LinkTlv}
    opaque.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(opaque.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv
// Unknown TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type_ interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (unknownTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_UnknownTlv) GetEntityData() *types.CommonEntityData {
    unknownTlv.EntityData.YFilter = unknownTlv.YFilter
    unknownTlv.EntityData.YangName = "unknown-tlv"
    unknownTlv.EntityData.BundleName = "ietf"
    unknownTlv.EntityData.ParentYangName = "opaque"
    unknownTlv.EntityData.SegmentPath = "unknown-tlv" + "[type='" + fmt.Sprintf("%v", unknownTlv.Type_) + "']"
    unknownTlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    unknownTlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    unknownTlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    unknownTlv.EntityData.Children = make(map[string]types.YChild)
    unknownTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownTlv.EntityData.Leafs["type"] = types.YLeaf{"Type_", unknownTlv.Type_}
    unknownTlv.EntityData.Leafs["length"] = types.YLeaf{"Length", unknownTlv.Length}
    unknownTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownTlv.Value}
    return &(unknownTlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv
// Router address TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterAddress interface{}
}

func (routerAddressTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_RouterAddressTlv) GetEntityData() *types.CommonEntityData {
    routerAddressTlv.EntityData.YFilter = routerAddressTlv.YFilter
    routerAddressTlv.EntityData.YangName = "router-address-tlv"
    routerAddressTlv.EntityData.BundleName = "ietf"
    routerAddressTlv.EntityData.ParentYangName = "opaque"
    routerAddressTlv.EntityData.SegmentPath = "router-address-tlv"
    routerAddressTlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routerAddressTlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    routerAddressTlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routerAddressTlv.EntityData.Children = make(map[string]types.YChild)
    routerAddressTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    routerAddressTlv.EntityData.Leafs["router-address"] = types.YLeaf{"RouterAddress", routerAddressTlv.RouterAddress}
    return &(routerAddressTlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv
// Link TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link type. The type is interface{} with range: 0..255. This attribute is
    // mandatory.
    LinkType interface{}

    // Link ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'
    // This attribute is mandatory..
    LinkId interface{}

    // List of local interface IPv4 addresses. The type is slice of string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalIfIpv4Addr []interface{}

    // List of remote interface IPv4 addresses. The type is slice of string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalRemoteIpv4Addr []interface{}

    // TE metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Maximum bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxBandwidth interface{}

    // Maximum reservable bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxReservableBandwidth interface{}

    // Unreserved bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    UnreservedBandwidth interface{}

    // Administrative group/Resource class/Color. The type is interface{} with
    // range: 0..4294967295.
    AdminGroup interface{}

    // Unknown sub-TLV. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv.
    UnknownSubtlv []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
}

func (linkTlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv) GetEntityData() *types.CommonEntityData {
    linkTlv.EntityData.YFilter = linkTlv.YFilter
    linkTlv.EntityData.YangName = "link-tlv"
    linkTlv.EntityData.BundleName = "ietf"
    linkTlv.EntityData.ParentYangName = "opaque"
    linkTlv.EntityData.SegmentPath = "link-tlv"
    linkTlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    linkTlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    linkTlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    linkTlv.EntityData.Children = make(map[string]types.YChild)
    linkTlv.EntityData.Children["unknown-subtlv"] = types.YChild{"UnknownSubtlv", nil}
    for i := range linkTlv.UnknownSubtlv {
        linkTlv.EntityData.Children[types.GetSegmentPath(&linkTlv.UnknownSubtlv[i])] = types.YChild{"UnknownSubtlv", &linkTlv.UnknownSubtlv[i]}
    }
    linkTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    linkTlv.EntityData.Leafs["link-type"] = types.YLeaf{"LinkType", linkTlv.LinkType}
    linkTlv.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", linkTlv.LinkId}
    linkTlv.EntityData.Leafs["local-if-ipv4-addr"] = types.YLeaf{"LocalIfIpv4Addr", linkTlv.LocalIfIpv4Addr}
    linkTlv.EntityData.Leafs["local-remote-ipv4-addr"] = types.YLeaf{"LocalRemoteIpv4Addr", linkTlv.LocalRemoteIpv4Addr}
    linkTlv.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", linkTlv.TeMetric}
    linkTlv.EntityData.Leafs["max-bandwidth"] = types.YLeaf{"MaxBandwidth", linkTlv.MaxBandwidth}
    linkTlv.EntityData.Leafs["max-reservable-bandwidth"] = types.YLeaf{"MaxReservableBandwidth", linkTlv.MaxReservableBandwidth}
    linkTlv.EntityData.Leafs["unreserved-bandwidth"] = types.YLeaf{"UnreservedBandwidth", linkTlv.UnreservedBandwidth}
    linkTlv.EntityData.Leafs["admin-group"] = types.YLeaf{"AdminGroup", linkTlv.AdminGroup}
    return &(linkTlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv
// Unknown sub-TLV.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type_ interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (unknownSubtlv *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv2_Body_Opaque_LinkTlv_UnknownSubtlv) GetEntityData() *types.CommonEntityData {
    unknownSubtlv.EntityData.YFilter = unknownSubtlv.YFilter
    unknownSubtlv.EntityData.YangName = "unknown-subtlv"
    unknownSubtlv.EntityData.BundleName = "ietf"
    unknownSubtlv.EntityData.ParentYangName = "link-tlv"
    unknownSubtlv.EntityData.SegmentPath = "unknown-subtlv" + "[type='" + fmt.Sprintf("%v", unknownSubtlv.Type_) + "']"
    unknownSubtlv.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    unknownSubtlv.EntityData.NamespaceTable = ietf.GetNamespaces()
    unknownSubtlv.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    unknownSubtlv.EntityData.Children = make(map[string]types.YChild)
    unknownSubtlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownSubtlv.EntityData.Leafs["type"] = types.YLeaf{"Type_", unknownSubtlv.Type_}
    unknownSubtlv.EntityData.Leafs["length"] = types.YLeaf{"Length", unknownSubtlv.Length}
    unknownSubtlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownSubtlv.Value}
    return &(unknownSubtlv.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3
// OSPFv3 LSA
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header data.
    Header RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header

    // Decoded OSPF LSA body data.
    Body RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body
}

func (ospfv3 *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3) GetEntityData() *types.CommonEntityData {
    ospfv3.EntityData.YFilter = ospfv3.YFilter
    ospfv3.EntityData.YangName = "ospfv3"
    ospfv3.EntityData.BundleName = "ietf"
    ospfv3.EntityData.ParentYangName = "as-scope-lsa"
    ospfv3.EntityData.SegmentPath = "ospfv3"
    ospfv3.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ospfv3.EntityData.NamespaceTable = ietf.GetNamespaces()
    ospfv3.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ospfv3.EntityData.Children = make(map[string]types.YChild)
    ospfv3.EntityData.Children["header"] = types.YChild{"Header", &ospfv3.Header}
    ospfv3.EntityData.Children["body"] = types.YChild{"Body", &ospfv3.Body}
    ospfv3.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header
// Decoded OSPFv3 LSA header data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is interface{} with range: 0..4294967295. This attribute
    // is mandatory.
    LsaId interface{}

    // LSA age. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Type_ interface{}

    // LSA advertising router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    // This attribute is mandatory.
    AdvRouter interface{}

    // LSA sequence number. The type is string. This attribute is mandatory.
    SeqNum interface{}

    // LSA checksum. The type is string. This attribute is mandatory.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535. This attribute is
    // mandatory.
    Length interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}
}

func (header *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "ietf"
    header.EntityData.ParentYangName = "ospfv3"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    header.EntityData.NamespaceTable = ietf.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    header.EntityData.Leafs["options"] = types.YLeaf{"Options", header.Options}
    return &(header.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body
// Decoded OSPF LSA body data.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router LSA.
    Router RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router

    // Network LSA.
    Network RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network

    // Inter-Area-Prefix LSA.
    InterAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix

    // Inter-Area-Router LSA.
    InterAreaRouter RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter

    // AS-External LSA.
    AsExternal RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal

    // NSSA LSA.
    Nssa RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa

    // Link LSA.
    Link RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link

    // Intra-Area-Prefix LSA.
    IntraAreaPrefix RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix
}

func (body *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body) GetEntityData() *types.CommonEntityData {
    body.EntityData.YFilter = body.YFilter
    body.EntityData.YangName = "body"
    body.EntityData.BundleName = "ietf"
    body.EntityData.ParentYangName = "ospfv3"
    body.EntityData.SegmentPath = "body"
    body.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    body.EntityData.NamespaceTable = ietf.GetNamespaces()
    body.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    body.EntityData.Children = make(map[string]types.YChild)
    body.EntityData.Children["router"] = types.YChild{"Router", &body.Router}
    body.EntityData.Children["network"] = types.YChild{"Network", &body.Network}
    body.EntityData.Children["inter-area-prefix"] = types.YChild{"InterAreaPrefix", &body.InterAreaPrefix}
    body.EntityData.Children["inter-area-router"] = types.YChild{"InterAreaRouter", &body.InterAreaRouter}
    body.EntityData.Children["as-external"] = types.YChild{"AsExternal", &body.AsExternal}
    body.EntityData.Children["nssa"] = types.YChild{"Nssa", &body.Nssa}
    body.EntityData.Children["link"] = types.YChild{"Link", &body.Link}
    body.EntityData.Children["intra-area-prefix"] = types.YChild{"IntraAreaPrefix", &body.IntraAreaPrefix}
    body.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(body.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router
// Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA option. The type is map[string]bool. This attribute is mandatory.
    Flags interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Router LSA link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link.
    Link []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link
}

func (router *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router) GetEntityData() *types.CommonEntityData {
    router.EntityData.YFilter = router.YFilter
    router.EntityData.YangName = "router"
    router.EntityData.BundleName = "ietf"
    router.EntityData.ParentYangName = "body"
    router.EntityData.SegmentPath = "router"
    router.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    router.EntityData.NamespaceTable = ietf.GetNamespaces()
    router.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    router.EntityData.Children = make(map[string]types.YChild)
    router.EntityData.Children["link"] = types.YChild{"Link", nil}
    for i := range router.Link {
        router.EntityData.Children[types.GetSegmentPath(&router.Link[i])] = types.YChild{"Link", &router.Link[i]}
    }
    router.EntityData.Leafs = make(map[string]types.YLeaf)
    router.EntityData.Leafs["flags"] = types.YLeaf{"Flags", router.Flags}
    router.EntityData.Leafs["options"] = types.YLeaf{"Options", router.Options}
    return &(router.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link
// Router LSA link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor Interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor Router ID. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Router_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "ietf"
    link.EntityData.ParentYangName = "router"
    link.EntityData.SegmentPath = "link" + "[interface-id='" + fmt.Sprintf("%v", link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", link.NeighborRouterId) + "']"
    link.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    link.EntityData.NamespaceTable = ietf.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    link.EntityData.Children = make(map[string]types.YChild)
    link.EntityData.Leafs = make(map[string]types.YLeaf)
    link.EntityData.Leafs["interface-id"] = types.YLeaf{"InterfaceId", link.InterfaceId}
    link.EntityData.Leafs["neighbor-interface-id"] = types.YLeaf{"NeighborInterfaceId", link.NeighborInterfaceId}
    link.EntityData.Leafs["neighbor-router-id"] = types.YLeaf{"NeighborRouterId", link.NeighborRouterId}
    link.EntityData.Leafs["type"] = types.YLeaf{"Type_", link.Type_}
    link.EntityData.Leafs["metric"] = types.YLeaf{"Metric", link.Metric}
    return &(link.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network
// Network LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // List of the routers attached to the network. The type is slice of string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    AttachedRouter []interface{}
}

func (network *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "ietf"
    network.EntityData.ParentYangName = "body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    network.EntityData.NamespaceTable = ietf.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["options"] = types.YLeaf{"Options", network.Options}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix
// Inter-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (interAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaPrefix) GetEntityData() *types.CommonEntityData {
    interAreaPrefix.EntityData.YFilter = interAreaPrefix.YFilter
    interAreaPrefix.EntityData.YangName = "inter-area-prefix"
    interAreaPrefix.EntityData.BundleName = "ietf"
    interAreaPrefix.EntityData.ParentYangName = "body"
    interAreaPrefix.EntityData.SegmentPath = "inter-area-prefix"
    interAreaPrefix.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interAreaPrefix.EntityData.NamespaceTable = ietf.GetNamespaces()
    interAreaPrefix.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interAreaPrefix.EntityData.Children = make(map[string]types.YChild)
    interAreaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    interAreaPrefix.EntityData.Leafs["metric"] = types.YLeaf{"Metric", interAreaPrefix.Metric}
    interAreaPrefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", interAreaPrefix.Prefix}
    interAreaPrefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", interAreaPrefix.PrefixOptions}
    return &(interAreaPrefix.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter
// Inter-Area-Router LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // The Router ID of the router being described by the LSA. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    DestinationRouterId interface{}
}

func (interAreaRouter *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_InterAreaRouter) GetEntityData() *types.CommonEntityData {
    interAreaRouter.EntityData.YFilter = interAreaRouter.YFilter
    interAreaRouter.EntityData.YangName = "inter-area-router"
    interAreaRouter.EntityData.BundleName = "ietf"
    interAreaRouter.EntityData.ParentYangName = "body"
    interAreaRouter.EntityData.SegmentPath = "inter-area-router"
    interAreaRouter.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interAreaRouter.EntityData.NamespaceTable = ietf.GetNamespaces()
    interAreaRouter.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interAreaRouter.EntityData.Children = make(map[string]types.YChild)
    interAreaRouter.EntityData.Leafs = make(map[string]types.YLeaf)
    interAreaRouter.EntityData.Leafs["options"] = types.YLeaf{"Options", interAreaRouter.Options}
    interAreaRouter.EntityData.Leafs["metric"] = types.YLeaf{"Metric", interAreaRouter.Metric}
    interAreaRouter.EntityData.Leafs["destination-router-id"] = types.YLeaf{"DestinationRouterId", interAreaRouter.DestinationRouterId}
    return &(interAreaRouter.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal
// AS-External LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (asExternal *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_AsExternal) GetEntityData() *types.CommonEntityData {
    asExternal.EntityData.YFilter = asExternal.YFilter
    asExternal.EntityData.YangName = "as-external"
    asExternal.EntityData.BundleName = "ietf"
    asExternal.EntityData.ParentYangName = "body"
    asExternal.EntityData.SegmentPath = "as-external"
    asExternal.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    asExternal.EntityData.NamespaceTable = ietf.GetNamespaces()
    asExternal.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    asExternal.EntityData.Children = make(map[string]types.YChild)
    asExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    asExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", asExternal.Metric}
    asExternal.EntityData.Leafs["flags"] = types.YLeaf{"Flags", asExternal.Flags}
    asExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", asExternal.ReferencedLsType}
    asExternal.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", asExternal.Prefix}
    asExternal.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", asExternal.PrefixOptions}
    asExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", asExternal.ForwardingAddress}
    asExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", asExternal.ExternalRouteTag}
    asExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", asExternal.ReferencedLinkStateId}
    return &(asExternal.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa
// NSSA LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Flags. The type is map[string]bool.
    Flags interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Forwarding address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}
}

func (nssa *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "ietf"
    nssa.EntityData.ParentYangName = "body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    nssa.EntityData.NamespaceTable = ietf.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    nssa.EntityData.Children = make(map[string]types.YChild)
    nssa.EntityData.Leafs = make(map[string]types.YLeaf)
    nssa.EntityData.Leafs["metric"] = types.YLeaf{"Metric", nssa.Metric}
    nssa.EntityData.Leafs["flags"] = types.YLeaf{"Flags", nssa.Flags}
    nssa.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", nssa.ReferencedLsType}
    nssa.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", nssa.Prefix}
    nssa.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", nssa.PrefixOptions}
    nssa.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", nssa.ForwardingAddress}
    nssa.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", nssa.ExternalRouteTag}
    nssa.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", nssa.ReferencedLinkStateId}
    return &(nssa.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link
// Link LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router Priority of the interface. The type is interface{} with range:
    // 0..255.
    RtrPriority interface{}

    // OSPFv3 LSA options. The type is map[string]bool. This attribute is
    // mandatory.
    Options interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList
}

func (link *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "ietf"
    link.EntityData.ParentYangName = "body"
    link.EntityData.SegmentPath = "link"
    link.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    link.EntityData.NamespaceTable = ietf.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    link.EntityData.Children = make(map[string]types.YChild)
    link.EntityData.Children["prefix-list"] = types.YChild{"PrefixList", nil}
    for i := range link.PrefixList {
        link.EntityData.Children[types.GetSegmentPath(&link.PrefixList[i])] = types.YChild{"PrefixList", &link.PrefixList[i]}
    }
    link.EntityData.Leafs = make(map[string]types.YLeaf)
    link.EntityData.Leafs["rtr-priority"] = types.YLeaf{"RtrPriority", link.RtrPriority}
    link.EntityData.Leafs["options"] = types.YLeaf{"Options", link.Options}
    link.EntityData.Leafs["link-local-interface-address"] = types.YLeaf{"LinkLocalInterfaceAddress", link.LinkLocalInterfaceAddress}
    link.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", link.NumOfPrefixes}
    return &(link.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_Link_PrefixList) GetEntityData() *types.CommonEntityData {
    prefixList.EntityData.YFilter = prefixList.YFilter
    prefixList.EntityData.YangName = "prefix-list"
    prefixList.EntityData.BundleName = "ietf"
    prefixList.EntityData.ParentYangName = "link"
    prefixList.EntityData.SegmentPath = "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
    prefixList.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    prefixList.EntityData.NamespaceTable = ietf.GetNamespaces()
    prefixList.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    prefixList.EntityData.Children = make(map[string]types.YChild)
    prefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixList.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", prefixList.Prefix}
    prefixList.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", prefixList.PrefixOptions}
    return &(prefixList.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix
// Intra-Area-Prefix LSA.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}

    // List of prefixes associated with the link. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList.
    PrefixList []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
}

func (intraAreaPrefix *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix) GetEntityData() *types.CommonEntityData {
    intraAreaPrefix.EntityData.YFilter = intraAreaPrefix.YFilter
    intraAreaPrefix.EntityData.YangName = "intra-area-prefix"
    intraAreaPrefix.EntityData.BundleName = "ietf"
    intraAreaPrefix.EntityData.ParentYangName = "body"
    intraAreaPrefix.EntityData.SegmentPath = "intra-area-prefix"
    intraAreaPrefix.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    intraAreaPrefix.EntityData.NamespaceTable = ietf.GetNamespaces()
    intraAreaPrefix.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    intraAreaPrefix.EntityData.Children = make(map[string]types.YChild)
    intraAreaPrefix.EntityData.Children["prefix-list"] = types.YChild{"PrefixList", nil}
    for i := range intraAreaPrefix.PrefixList {
        intraAreaPrefix.EntityData.Children[types.GetSegmentPath(&intraAreaPrefix.PrefixList[i])] = types.YChild{"PrefixList", &intraAreaPrefix.PrefixList[i]}
    }
    intraAreaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    intraAreaPrefix.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", intraAreaPrefix.ReferencedLsType}
    intraAreaPrefix.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", intraAreaPrefix.ReferencedLinkStateId}
    intraAreaPrefix.EntityData.Leafs["referenced-adv-router"] = types.YLeaf{"ReferencedAdvRouter", intraAreaPrefix.ReferencedAdvRouter}
    intraAreaPrefix.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", intraAreaPrefix.NumOfPrefixes}
    return &(intraAreaPrefix.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList
// List of prefixes associated with the link.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string. This attribute is mandatory.
    PrefixOptions interface{}

    // Metric. The type is interface{} with range: 0..16777215.
    Metric interface{}
}

func (prefixList *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AsScopeLsas_AsScopeLsa_Ospfv3_Body_IntraAreaPrefix_PrefixList) GetEntityData() *types.CommonEntityData {
    prefixList.EntityData.YFilter = prefixList.YFilter
    prefixList.EntityData.YangName = "prefix-list"
    prefixList.EntityData.BundleName = "ietf"
    prefixList.EntityData.ParentYangName = "intra-area-prefix"
    prefixList.EntityData.SegmentPath = "prefix-list" + "[prefix='" + fmt.Sprintf("%v", prefixList.Prefix) + "']"
    prefixList.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    prefixList.EntityData.NamespaceTable = ietf.GetNamespaces()
    prefixList.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    prefixList.EntityData.Children = make(map[string]types.YChild)
    prefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixList.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", prefixList.Prefix}
    prefixList.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", prefixList.PrefixOptions}
    prefixList.EntityData.Leafs["metric"] = types.YLeaf{"Metric", prefixList.Metric}
    return &(prefixList.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology
// OSPF topology.
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. RIB. The type is string. Refers to
    // routing.Routing_RoutingInstance_Ribs_Rib_Name
    Name interface{}

    // List of ospf areas. The type is slice of
    // RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area.
    Area []RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area
}

func (topology *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "instance"
    topology.EntityData.SegmentPath = "topology" + "[name='" + fmt.Sprintf("%v", topology.Name) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Children["area"] = types.YChild{"Area", nil}
    for i := range topology.Area {
        topology.EntityData.Children[types.GetSegmentPath(&topology.Area[i])] = types.YChild{"Area", &topology.Area[i]}
    }
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["name"] = types.YLeaf{"Name", topology.Name}
    return &(topology.EntityData)
}

// RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area
// List of ospf areas
type RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Area ID. The type is one of the following types:
    // int with range: 0..4294967295, or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    AreaId interface{}
}

func (area *RoutingState_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetEntityData() *types.CommonEntityData {
    area.EntityData.YFilter = area.YFilter
    area.EntityData.YangName = "area"
    area.EntityData.BundleName = "ietf"
    area.EntityData.ParentYangName = "topology"
    area.EntityData.SegmentPath = "area" + "[area-id='" + fmt.Sprintf("%v", area.AreaId) + "']"
    area.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    area.EntityData.NamespaceTable = ietf.GetNamespaces()
    area.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    area.EntityData.Children = make(map[string]types.YChild)
    area.EntityData.Leafs = make(map[string]types.YLeaf)
    area.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", area.AreaId}
    return &(area.EntityData)
}

// RoutingState_RoutingInstance_Ribs
// Container for RIBs.
type RoutingState_RoutingInstance_Ribs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a RIB identified by the 'name' key. All routes in a
    // RIB MUST belong to the same address family.  For each routing instance, an
    // implementation SHOULD provide one system-controlled default RIB for each
    // supported address family. The type is slice of
    // RoutingState_RoutingInstance_Ribs_Rib.
    Rib []RoutingState_RoutingInstance_Ribs_Rib
}

func (ribs *RoutingState_RoutingInstance_Ribs) GetEntityData() *types.CommonEntityData {
    ribs.EntityData.YFilter = ribs.YFilter
    ribs.EntityData.YangName = "ribs"
    ribs.EntityData.BundleName = "ietf"
    ribs.EntityData.ParentYangName = "routing-instance"
    ribs.EntityData.SegmentPath = "ribs"
    ribs.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ribs.EntityData.NamespaceTable = ietf.GetNamespaces()
    ribs.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ribs.EntityData.Children = make(map[string]types.YChild)
    ribs.EntityData.Children["rib"] = types.YChild{"Rib", nil}
    for i := range ribs.Rib {
        ribs.EntityData.Children[types.GetSegmentPath(&ribs.Rib[i])] = types.YChild{"Rib", &ribs.Rib[i]}
    }
    ribs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ribs.EntityData)
}

// RoutingState_RoutingInstance_Ribs_Rib
// Each entry represents a RIB identified by the 'name'
// key. All routes in a RIB MUST belong to the same address
// family.
// 
// For each routing instance, an implementation SHOULD
// provide one system-controlled default RIB for each
// supported address family.
type RoutingState_RoutingInstance_Ribs_Rib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the RIB. The type is string.
    Name interface{}

    // Address family. The type is one of the following:
    // Ipv4Ipv4UnicastIpv6Ipv6Unicast. This attribute is mandatory.
    AddressFamily interface{}

    // This flag has the value of 'true' if and only if the RIB is the default RIB
    // for the given address family.  A default RIB always receives direct routes.
    // By default it also receives routes from all routing protocols. The type is
    // bool. The default value is true.
    DefaultRib interface{}

    // Current content of the RIB.
    Routes RoutingState_RoutingInstance_Ribs_Rib_Routes
}

func (rib *RoutingState_RoutingInstance_Ribs_Rib) GetEntityData() *types.CommonEntityData {
    rib.EntityData.YFilter = rib.YFilter
    rib.EntityData.YangName = "rib"
    rib.EntityData.BundleName = "ietf"
    rib.EntityData.ParentYangName = "ribs"
    rib.EntityData.SegmentPath = "rib" + "[name='" + fmt.Sprintf("%v", rib.Name) + "']"
    rib.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    rib.EntityData.NamespaceTable = ietf.GetNamespaces()
    rib.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    rib.EntityData.Children = make(map[string]types.YChild)
    rib.EntityData.Children["routes"] = types.YChild{"Routes", &rib.Routes}
    rib.EntityData.Leafs = make(map[string]types.YLeaf)
    rib.EntityData.Leafs["name"] = types.YLeaf{"Name", rib.Name}
    rib.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", rib.AddressFamily}
    rib.EntityData.Leafs["default-rib"] = types.YLeaf{"DefaultRib", rib.DefaultRib}
    return &(rib.EntityData)
}

// RoutingState_RoutingInstance_Ribs_Rib_Routes
// Current content of the RIB.
type RoutingState_RoutingInstance_Ribs_Rib_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A RIB route entry. This data node MUST be augmented with information
    // specific for routes of each address family. The type is slice of
    // RoutingState_RoutingInstance_Ribs_Rib_Routes_Route.
    Route []RoutingState_RoutingInstance_Ribs_Rib_Routes_Route
}

func (routes *RoutingState_RoutingInstance_Ribs_Rib_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "ietf"
    routes.EntityData.ParentYangName = "rib"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routes.EntityData.NamespaceTable = ietf.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routes.EntityData.Children = make(map[string]types.YChild)
    routes.EntityData.Children["route"] = types.YChild{"Route", nil}
    for i := range routes.Route {
        routes.EntityData.Children[types.GetSegmentPath(&routes.Route[i])] = types.YChild{"Route", &routes.Route[i]}
    }
    routes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routes.EntityData)
}

// RoutingState_RoutingInstance_Ribs_Rib_Routes_Route
// A RIB route entry. This data node MUST be augmented
// with information specific for routes of each address
// family.
type RoutingState_RoutingInstance_Ribs_Rib_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Destination IP address with prefix. The type is
    // string.
    DestinationPrefix interface{}

    // This route attribute, also known as administrative distance, allows for
    // selecting the preferred route among routes with the same destination
    // prefix. A smaller value means a more preferred route. The type is
    // interface{} with range: 0..4294967295.
    RoutePreference interface{}

    // Route metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Type of the routing protocol from which the route originated. The type is
    // one of the following: OspfOspfv2Ospfv3DirectStatic. This attribute is
    // mandatory.
    SourceProtocol interface{}

    // Presence of this leaf indicates that the route is preferred among all
    // routes in the same RIB that have the same destination prefix. The type is
    // interface{}.
    Active interface{}

    // Time stamp of the last modification of the route. If the route was never
    // modified, it is the time when the route was inserted into the RIB. The type
    // is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    LastUpdated interface{}

    // Update source for the route. The type is string.
    UpdateSource interface{}

    // OSPF route tag. The type is interface{} with range: 0..4294967295. The
    // default value is 0.
    Tag interface{}

    // OSPF route type. The type is RouteType.
    RouteType interface{}

    // Route's next-hop attribute.
    NextHop RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop
}

func (route *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "ietf"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + "[destination-prefix='" + fmt.Sprintf("%v", route.DestinationPrefix) + "']"
    route.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    route.EntityData.NamespaceTable = ietf.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    route.EntityData.Children = make(map[string]types.YChild)
    route.EntityData.Children["next-hop"] = types.YChild{"NextHop", &route.NextHop}
    route.EntityData.Leafs = make(map[string]types.YLeaf)
    route.EntityData.Leafs["destination-prefix"] = types.YLeaf{"DestinationPrefix", route.DestinationPrefix}
    route.EntityData.Leafs["route-preference"] = types.YLeaf{"RoutePreference", route.RoutePreference}
    route.EntityData.Leafs["metric"] = types.YLeaf{"Metric", route.Metric}
    route.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", route.SourceProtocol}
    route.EntityData.Leafs["active"] = types.YLeaf{"Active", route.Active}
    route.EntityData.Leafs["last-updated"] = types.YLeaf{"LastUpdated", route.LastUpdated}
    route.EntityData.Leafs["update-source"] = types.YLeaf{"UpdateSource", route.UpdateSource}
    route.EntityData.Leafs["tag"] = types.YLeaf{"Tag", route.Tag}
    route.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", route.RouteType}
    return &(route.EntityData)
}

// RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop
// Route's next-hop attribute.
type RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the outgoing interface. The type is string.
    OutgoingInterface interface{}

    // IP address. The type is string.
    NextHopAddress interface{}

    // Special next-hop options. The type is SpecialNextHop.
    SpecialNextHop interface{}
}

func (nextHop *RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "ietf"
    nextHop.EntityData.ParentYangName = "route"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    nextHop.EntityData.NamespaceTable = ietf.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["outgoing-interface"] = types.YLeaf{"OutgoingInterface", nextHop.OutgoingInterface}
    nextHop.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", nextHop.NextHopAddress}
    nextHop.EntityData.Leafs["special-next-hop"] = types.YLeaf{"SpecialNextHop", nextHop.SpecialNextHop}
    return &(nextHop.EntityData)
}

// RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop represents Special next-hop options.
type RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop string

const (
    // Silently discard the packet.
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop_blackhole RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop = "blackhole"

    // Discard the packet and notify the sender with an error
    // message indicating that the destination host is
    // unreachable.
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop_unreachable RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop = "unreachable"

    // Discard the packet and notify the sender with an error
    // message indicating that the communication is
    // administratively prohibited.
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop_prohibit RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop = "prohibit"

    // The packet will be received by the local system.
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop_receive RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_NextHop_SpecialNextHop = "receive"
)

// RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType represents OSPF route type
type RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType string

const (
    // OSPF intra-area route
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_intra_area RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "intra-area"

    // OSPF inter-area route
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_inter_area RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "inter-area"

    // OSPF external route type 1
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_external_1 RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "external-1"

    // OSPF External route type 2
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_external_2 RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "external-2"

    // OSPF NSSA external route type 1
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_nssa_1 RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "nssa-1"

    // OSPF NSSA external route type 2
    RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType_nssa_2 RoutingState_RoutingInstance_Ribs_Rib_Routes_Route_RouteType = "nssa-2"
)

// Routing
// Configuration parameters for the routing subsystem.
type Routing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of a routing instance. The type is slice of
    // Routing_RoutingInstance.
    RoutingInstance []Routing_RoutingInstance
}

func (routing *Routing) GetEntityData() *types.CommonEntityData {
    routing.EntityData.YFilter = routing.YFilter
    routing.EntityData.YangName = "routing"
    routing.EntityData.BundleName = "ietf"
    routing.EntityData.ParentYangName = "ietf-routing"
    routing.EntityData.SegmentPath = "ietf-routing:routing"
    routing.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routing.EntityData.NamespaceTable = ietf.GetNamespaces()
    routing.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routing.EntityData.Children = make(map[string]types.YChild)
    routing.EntityData.Children["routing-instance"] = types.YChild{"RoutingInstance", nil}
    for i := range routing.RoutingInstance {
        routing.EntityData.Children[types.GetSegmentPath(&routing.RoutingInstance[i])] = types.YChild{"RoutingInstance", &routing.RoutingInstance[i]}
    }
    routing.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routing.EntityData)
}

// Routing_RoutingInstance
// Configuration of a routing instance.
type Routing_RoutingInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the routing instance.  For
    // system-controlled entries, the value of this leaf must be the same as the
    // name of the corresponding entry in state data.  For user-controlled
    // entries, an arbitrary name can be used. The type is string.
    Name interface{}

    // The type of the routing instance. The type is one of the following:
    // DefaultRoutingInstanceVrfRoutingInstance. The default value is
    // rt:default-routing-instance.
    Type_ interface{}

    // Enable/disable the routing instance.  If this parameter is false, the
    // parent routing instance is disabled and does not appear in state data,
    // despite any other configuration that might be present. The type is bool.
    // The default value is true.
    Enabled interface{}

    // A 32-bit number in the form of a dotted quad that is used by some routing
    // protocols identifying a router. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    RouterId interface{}

    // Textual description of the routing instance. The type is string.
    Description interface{}

    // Assignment of the routing instance's interfaces.
    Interfaces Routing_RoutingInstance_Interfaces

    // Configuration of routing protocol instances.
    RoutingProtocols Routing_RoutingInstance_RoutingProtocols

    // Configuration of RIBs.
    Ribs Routing_RoutingInstance_Ribs
}

func (routingInstance *Routing_RoutingInstance) GetEntityData() *types.CommonEntityData {
    routingInstance.EntityData.YFilter = routingInstance.YFilter
    routingInstance.EntityData.YangName = "routing-instance"
    routingInstance.EntityData.BundleName = "ietf"
    routingInstance.EntityData.ParentYangName = "routing"
    routingInstance.EntityData.SegmentPath = "routing-instance" + "[name='" + fmt.Sprintf("%v", routingInstance.Name) + "']"
    routingInstance.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routingInstance.EntityData.NamespaceTable = ietf.GetNamespaces()
    routingInstance.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routingInstance.EntityData.Children = make(map[string]types.YChild)
    routingInstance.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &routingInstance.Interfaces}
    routingInstance.EntityData.Children["routing-protocols"] = types.YChild{"RoutingProtocols", &routingInstance.RoutingProtocols}
    routingInstance.EntityData.Children["ribs"] = types.YChild{"Ribs", &routingInstance.Ribs}
    routingInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    routingInstance.EntityData.Leafs["name"] = types.YLeaf{"Name", routingInstance.Name}
    routingInstance.EntityData.Leafs["type"] = types.YLeaf{"Type_", routingInstance.Type_}
    routingInstance.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", routingInstance.Enabled}
    routingInstance.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", routingInstance.RouterId}
    routingInstance.EntityData.Leafs["description"] = types.YLeaf{"Description", routingInstance.Description}
    return &(routingInstance.EntityData)
}

// Routing_RoutingInstance_Interfaces
// Assignment of the routing instance's interfaces.
type Routing_RoutingInstance_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of a configured network layer interface to be assigned to the
    // routing-instance. The type is slice of string. Refers to
    // interfaces.Interfaces_Interface_Name
    Interface_ []interface{}
}

func (interfaces *Routing_RoutingInstance_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "ietf"
    interfaces.EntityData.ParentYangName = "routing-instance"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    interfaces.EntityData.NamespaceTable = ietf.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaces.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", interfaces.Interface_}
    return &(interfaces.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols
// Configuration of routing protocol instances.
type Routing_RoutingInstance_RoutingProtocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains configuration of a routing protocol instance. The type
    // is slice of Routing_RoutingInstance_RoutingProtocols_RoutingProtocol.
    RoutingProtocol []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol
}

func (routingProtocols *Routing_RoutingInstance_RoutingProtocols) GetEntityData() *types.CommonEntityData {
    routingProtocols.EntityData.YFilter = routingProtocols.YFilter
    routingProtocols.EntityData.YangName = "routing-protocols"
    routingProtocols.EntityData.BundleName = "ietf"
    routingProtocols.EntityData.ParentYangName = "routing-instance"
    routingProtocols.EntityData.SegmentPath = "routing-protocols"
    routingProtocols.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routingProtocols.EntityData.NamespaceTable = ietf.GetNamespaces()
    routingProtocols.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routingProtocols.EntityData.Children = make(map[string]types.YChild)
    routingProtocols.EntityData.Children["routing-protocol"] = types.YChild{"RoutingProtocol", nil}
    for i := range routingProtocols.RoutingProtocol {
        routingProtocols.EntityData.Children[types.GetSegmentPath(&routingProtocols.RoutingProtocol[i])] = types.YChild{"RoutingProtocol", &routingProtocols.RoutingProtocol[i]}
    }
    routingProtocols.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routingProtocols.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol
// Each entry contains configuration of a routing protocol
// instance.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Type of the routing protocol - an identity derived
    // from the 'routing-protocol' base identity. The type is one of the
    // following: OspfOspfv2Ospfv3DirectStatic.
    Type_ interface{}

    // This attribute is a key. An arbitrary name of the routing protocol
    // instance. The type is string.
    Name interface{}

    // Textual description of the routing protocol instance. The type is string.
    Description interface{}

    // Configuration of the 'static' pseudo-protocol.  Address-family-specific
    // modules augment this node with their lists of routes.
    StaticRoutes Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes

    // OSPF.
    Ospf Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf
}

func (routingProtocol *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol) GetEntityData() *types.CommonEntityData {
    routingProtocol.EntityData.YFilter = routingProtocol.YFilter
    routingProtocol.EntityData.YangName = "routing-protocol"
    routingProtocol.EntityData.BundleName = "ietf"
    routingProtocol.EntityData.ParentYangName = "routing-protocols"
    routingProtocol.EntityData.SegmentPath = "routing-protocol" + "[type='" + fmt.Sprintf("%v", routingProtocol.Type_) + "']" + "[name='" + fmt.Sprintf("%v", routingProtocol.Name) + "']"
    routingProtocol.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    routingProtocol.EntityData.NamespaceTable = ietf.GetNamespaces()
    routingProtocol.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    routingProtocol.EntityData.Children = make(map[string]types.YChild)
    routingProtocol.EntityData.Children["static-routes"] = types.YChild{"StaticRoutes", &routingProtocol.StaticRoutes}
    routingProtocol.EntityData.Children["ietf-ospf:ospf"] = types.YChild{"Ospf", &routingProtocol.Ospf}
    routingProtocol.EntityData.Leafs = make(map[string]types.YLeaf)
    routingProtocol.EntityData.Leafs["type"] = types.YLeaf{"Type_", routingProtocol.Type_}
    routingProtocol.EntityData.Leafs["name"] = types.YLeaf{"Name", routingProtocol.Name}
    routingProtocol.EntityData.Leafs["description"] = types.YLeaf{"Description", routingProtocol.Description}
    return &(routingProtocol.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes
// Configuration of the 'static' pseudo-protocol.
// 
// Address-family-specific modules augment this node with
// their lists of routes.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of a 'static' pseudo-protocol instance consists of a list of
    // routes.
    Ipv4 Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4

    // Configuration of a 'static' pseudo-protocol instance consists of a list of
    // routes.
    Ipv6 Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6
}

func (staticRoutes *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes) GetEntityData() *types.CommonEntityData {
    staticRoutes.EntityData.YFilter = staticRoutes.YFilter
    staticRoutes.EntityData.YangName = "static-routes"
    staticRoutes.EntityData.BundleName = "ietf"
    staticRoutes.EntityData.ParentYangName = "routing-protocol"
    staticRoutes.EntityData.SegmentPath = "static-routes"
    staticRoutes.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    staticRoutes.EntityData.NamespaceTable = ietf.GetNamespaces()
    staticRoutes.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    staticRoutes.EntityData.Children = make(map[string]types.YChild)
    staticRoutes.EntityData.Children["ietf-ipv4-unicast-routing:ipv4"] = types.YChild{"Ipv4", &staticRoutes.Ipv4}
    staticRoutes.EntityData.Children["ietf-ipv6-unicast-routing:ipv6"] = types.YChild{"Ipv6", &staticRoutes.Ipv6}
    staticRoutes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticRoutes.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4
// Configuration of a 'static' pseudo-protocol instance
// consists of a list of routes.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A user-ordered list of static routes. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route.
    Route []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route
}

func (ipv4 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "ietf"
    ipv4.EntityData.ParentYangName = "static-routes"
    ipv4.EntityData.SegmentPath = "ietf-ipv4-unicast-routing:ipv4"
    ipv4.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ipv4.EntityData.NamespaceTable = ietf.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["route"] = types.YChild{"Route", nil}
    for i := range ipv4.Route {
        ipv4.EntityData.Children[types.GetSegmentPath(&ipv4.Route[i])] = types.YChild{"Route", &ipv4.Route[i]}
    }
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route
// A user-ordered list of static routes.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 destination prefix. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))'.
    // This attribute is mandatory.
    DestinationPrefix interface{}

    // Textual description of the route. The type is string.
    Description interface{}

    // Configuration of next-hop.
    NextHop Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "ietf"
    route.EntityData.ParentYangName = "ipv4"
    route.EntityData.SegmentPath = "route" + "[destination-prefix='" + fmt.Sprintf("%v", route.DestinationPrefix) + "']"
    route.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    route.EntityData.NamespaceTable = ietf.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    route.EntityData.Children = make(map[string]types.YChild)
    route.EntityData.Children["next-hop"] = types.YChild{"NextHop", &route.NextHop}
    route.EntityData.Leafs = make(map[string]types.YLeaf)
    route.EntityData.Leafs["destination-prefix"] = types.YLeaf{"DestinationPrefix", route.DestinationPrefix}
    route.EntityData.Leafs["description"] = types.YLeaf{"Description", route.Description}
    return &(route.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop
// Configuration of next-hop.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the outgoing interface. The type is string.
    OutgoingInterface interface{}

    // Special next-hop options. The type is SpecialNextHop.
    SpecialNextHop interface{}

    // IPv4 address of the next-hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "ietf"
    nextHop.EntityData.ParentYangName = "route"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    nextHop.EntityData.NamespaceTable = ietf.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["outgoing-interface"] = types.YLeaf{"OutgoingInterface", nextHop.OutgoingInterface}
    nextHop.EntityData.Leafs["special-next-hop"] = types.YLeaf{"SpecialNextHop", nextHop.SpecialNextHop}
    nextHop.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", nextHop.NextHopAddress}
    return &(nextHop.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop represents Special next-hop options.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop string

const (
    // Silently discard the packet.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop_blackhole Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop = "blackhole"

    // Discard the packet and notify the sender with an error
    // message indicating that the destination host is
    // unreachable.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop_unreachable Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop = "unreachable"

    // Discard the packet and notify the sender with an error
    // message indicating that the communication is
    // administratively prohibited.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop_prohibit Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop = "prohibit"

    // The packet will be received by the local system.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop_receive Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv4_Route_NextHop_SpecialNextHop = "receive"
)

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6
// Configuration of a 'static' pseudo-protocol instance
// consists of a list of routes.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A user-ordered list of static routes. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route.
    Route []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route
}

func (ipv6 *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "ietf"
    ipv6.EntityData.ParentYangName = "static-routes"
    ipv6.EntityData.SegmentPath = "ietf-ipv6-unicast-routing:ipv6"
    ipv6.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ipv6.EntityData.NamespaceTable = ietf.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["route"] = types.YChild{"Route", nil}
    for i := range ipv6.Route {
        ipv6.EntityData.Children[types.GetSegmentPath(&ipv6.Route[i])] = types.YChild{"Route", &ipv6.Route[i]}
    }
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route
// A user-ordered list of static routes.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 destination prefix. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    // This attribute is mandatory.
    DestinationPrefix interface{}

    // Textual description of the route. The type is string.
    Description interface{}

    // Configuration of next-hop.
    NextHop Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop
}

func (route *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "ietf"
    route.EntityData.ParentYangName = "ipv6"
    route.EntityData.SegmentPath = "route" + "[destination-prefix='" + fmt.Sprintf("%v", route.DestinationPrefix) + "']"
    route.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    route.EntityData.NamespaceTable = ietf.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    route.EntityData.Children = make(map[string]types.YChild)
    route.EntityData.Children["next-hop"] = types.YChild{"NextHop", &route.NextHop}
    route.EntityData.Leafs = make(map[string]types.YLeaf)
    route.EntityData.Leafs["destination-prefix"] = types.YLeaf{"DestinationPrefix", route.DestinationPrefix}
    route.EntityData.Leafs["description"] = types.YLeaf{"Description", route.Description}
    return &(route.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop
// Configuration of next-hop.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the outgoing interface. The type is string.
    OutgoingInterface interface{}

    // Special next-hop options. The type is SpecialNextHop.
    SpecialNextHop interface{}

    // IPv6 address of the next-hop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}
}

func (nextHop *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "ietf"
    nextHop.EntityData.ParentYangName = "route"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    nextHop.EntityData.NamespaceTable = ietf.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["outgoing-interface"] = types.YLeaf{"OutgoingInterface", nextHop.OutgoingInterface}
    nextHop.EntityData.Leafs["special-next-hop"] = types.YLeaf{"SpecialNextHop", nextHop.SpecialNextHop}
    nextHop.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", nextHop.NextHopAddress}
    return &(nextHop.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop represents Special next-hop options.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop string

const (
    // Silently discard the packet.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop_blackhole Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop = "blackhole"

    // Discard the packet and notify the sender with an error
    // message indicating that the destination host is
    // unreachable.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop_unreachable Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop = "unreachable"

    // Discard the packet and notify the sender with an error
    // message indicating that the communication is
    // administratively prohibited.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop_prohibit Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop = "prohibit"

    // The packet will be received by the local system.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop_receive Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_StaticRoutes_Ipv6_Route_NextHop_SpecialNextHop = "receive"
)

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf
// OSPF.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF operation mode. The type is one of the following: ShipsInTheNight. The
    // default value is ospf:ships-in-the-night.
    OperationMode interface{}

    // Inheritance support to all instances.
    AllInstancesInherit Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit

    // An OSPF routing protocol instance. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance.
    Instance []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance
}

func (ospf *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "ietf"
    ospf.EntityData.ParentYangName = "routing-protocol"
    ospf.EntityData.SegmentPath = "ietf-ospf:ospf"
    ospf.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ospf.EntityData.NamespaceTable = ietf.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Children["all-instances-inherit"] = types.YChild{"AllInstancesInherit", &ospf.AllInstancesInherit}
    ospf.EntityData.Children["instance"] = types.YChild{"Instance", nil}
    for i := range ospf.Instance {
        ospf.EntityData.Children[types.GetSegmentPath(&ospf.Instance[i])] = types.YChild{"Instance", &ospf.Instance[i]}
    }
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["operation-mode"] = types.YLeaf{"OperationMode", ospf.OperationMode}
    return &(ospf.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit
// Inheritance support to all instances.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Area config to be inherited by all areas in all instances.
    Area Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area

    // Interface config to be inherited by all interfaces in all instances.
    Interface_ Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface
}

func (allInstancesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit) GetEntityData() *types.CommonEntityData {
    allInstancesInherit.EntityData.YFilter = allInstancesInherit.YFilter
    allInstancesInherit.EntityData.YangName = "all-instances-inherit"
    allInstancesInherit.EntityData.BundleName = "ietf"
    allInstancesInherit.EntityData.ParentYangName = "ospf"
    allInstancesInherit.EntityData.SegmentPath = "all-instances-inherit"
    allInstancesInherit.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    allInstancesInherit.EntityData.NamespaceTable = ietf.GetNamespaces()
    allInstancesInherit.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    allInstancesInherit.EntityData.Children = make(map[string]types.YChild)
    allInstancesInherit.EntityData.Children["area"] = types.YChild{"Area", &allInstancesInherit.Area}
    allInstancesInherit.EntityData.Children["interface"] = types.YChild{"Interface_", &allInstancesInherit.Interface_}
    allInstancesInherit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(allInstancesInherit.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area
// Area config to be inherited by all areas in
// all instances.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Area) GetEntityData() *types.CommonEntityData {
    area.EntityData.YFilter = area.YFilter
    area.EntityData.YangName = "area"
    area.EntityData.BundleName = "ietf"
    area.EntityData.ParentYangName = "all-instances-inherit"
    area.EntityData.SegmentPath = "area"
    area.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    area.EntityData.NamespaceTable = ietf.GetNamespaces()
    area.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    area.EntityData.Children = make(map[string]types.YChild)
    area.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(area.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface
// Interface config to be inherited by all interfaces
// in all instances.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_AllInstancesInherit_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "ietf"
    self.EntityData.ParentYangName = "all-instances-inherit"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    self.EntityData.NamespaceTable = ietf.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance
// An OSPF routing protocol instance.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address-family of the instance. The type is one of
    // the following: Ipv4Ipv4UnicastIpv6Ipv6Unicast.
    Af interface{}

    // Defined in RFC 2328. A 32-bit number that uniquely identifies the router.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    RouterId interface{}

    // Enable/Disable the protocol. The type is bool. The default value is true.
    Enable interface{}

    // Admin distance config state.
    AdminDistance Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance

    // NSR config state.
    Nsr Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr

    // Graceful restart config state.
    GracefulRestart Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart

    // Auto cost config state.
    AutoCost Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost

    // SPF calculation control.
    SpfControl Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl

    // Database maintenance control.
    DatabaseControl Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl

    // Protocol reload control.
    ReloadControl Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl

    // OSPF MPLS config state.
    Mpls Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls

    // This container may be augmented with global parameters for IPFRR.
    FastReroute Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute

    // Inheritance for all areas.
    AllAreasInherit Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit

    // List of ospf areas. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area.
    Area []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area

    // OSPF topology. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology.
    Topology []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology
}

func (instance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "ietf"
    instance.EntityData.ParentYangName = "ospf"
    instance.EntityData.SegmentPath = "instance" + "[af='" + fmt.Sprintf("%v", instance.Af) + "']"
    instance.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    instance.EntityData.NamespaceTable = ietf.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Children["admin-distance"] = types.YChild{"AdminDistance", &instance.AdminDistance}
    instance.EntityData.Children["nsr"] = types.YChild{"Nsr", &instance.Nsr}
    instance.EntityData.Children["graceful-restart"] = types.YChild{"GracefulRestart", &instance.GracefulRestart}
    instance.EntityData.Children["auto-cost"] = types.YChild{"AutoCost", &instance.AutoCost}
    instance.EntityData.Children["spf-control"] = types.YChild{"SpfControl", &instance.SpfControl}
    instance.EntityData.Children["database-control"] = types.YChild{"DatabaseControl", &instance.DatabaseControl}
    instance.EntityData.Children["reload-control"] = types.YChild{"ReloadControl", &instance.ReloadControl}
    instance.EntityData.Children["mpls"] = types.YChild{"Mpls", &instance.Mpls}
    instance.EntityData.Children["fast-reroute"] = types.YChild{"FastReroute", &instance.FastReroute}
    instance.EntityData.Children["all-areas-inherit"] = types.YChild{"AllAreasInherit", &instance.AllAreasInherit}
    instance.EntityData.Children["area"] = types.YChild{"Area", nil}
    for i := range instance.Area {
        instance.EntityData.Children[types.GetSegmentPath(&instance.Area[i])] = types.YChild{"Area", &instance.Area[i]}
    }
    instance.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range instance.Topology {
        instance.EntityData.Children[types.GetSegmentPath(&instance.Topology[i])] = types.YChild{"Topology", &instance.Topology[i]}
    }
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["af"] = types.YLeaf{"Af", instance.Af}
    instance.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", instance.RouterId}
    instance.EntityData.Leafs["enable"] = types.YLeaf{"Enable", instance.Enable}
    return &(instance.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance
// Admin distance config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Admin distance for intra-area route. The type is interface{} with range:
    // 0..255.
    IntraArea interface{}

    // Admin distance for inter-area route. The type is interface{} with range:
    // 0..255.
    InterArea interface{}

    // Admin distance for both intra-area and inter-area route. The type is
    // interface{} with range: 0..255.
    Internal interface{}

    // Admin distance for both external route. The type is interface{} with range:
    // 0..255.
    External interface{}
}

func (adminDistance *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AdminDistance) GetEntityData() *types.CommonEntityData {
    adminDistance.EntityData.YFilter = adminDistance.YFilter
    adminDistance.EntityData.YangName = "admin-distance"
    adminDistance.EntityData.BundleName = "ietf"
    adminDistance.EntityData.ParentYangName = "instance"
    adminDistance.EntityData.SegmentPath = "admin-distance"
    adminDistance.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    adminDistance.EntityData.NamespaceTable = ietf.GetNamespaces()
    adminDistance.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    adminDistance.EntityData.Children = make(map[string]types.YChild)
    adminDistance.EntityData.Leafs = make(map[string]types.YLeaf)
    adminDistance.EntityData.Leafs["intra-area"] = types.YLeaf{"IntraArea", adminDistance.IntraArea}
    adminDistance.EntityData.Leafs["inter-area"] = types.YLeaf{"InterArea", adminDistance.InterArea}
    adminDistance.EntityData.Leafs["internal"] = types.YLeaf{"Internal", adminDistance.Internal}
    adminDistance.EntityData.Leafs["external"] = types.YLeaf{"External", adminDistance.External}
    return &(adminDistance.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr
// NSR config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable NSR. The type is bool.
    Enable interface{}
}

func (nsr *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Nsr) GetEntityData() *types.CommonEntityData {
    nsr.EntityData.YFilter = nsr.YFilter
    nsr.EntityData.YangName = "nsr"
    nsr.EntityData.BundleName = "ietf"
    nsr.EntityData.ParentYangName = "instance"
    nsr.EntityData.SegmentPath = "nsr"
    nsr.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    nsr.EntityData.NamespaceTable = ietf.GetNamespaces()
    nsr.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    nsr.EntityData.Children = make(map[string]types.YChild)
    nsr.EntityData.Leafs = make(map[string]types.YLeaf)
    nsr.EntityData.Leafs["enable"] = types.YLeaf{"Enable", nsr.Enable}
    return &(nsr.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart
// Graceful restart config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable graceful restart as defined in RFC 3623. The type is bool.
    Enable interface{}

    // Enable RestartHelperSupport in RFC 3623 Section B.2. The type is bool.
    HelperEnable interface{}

    // RestartInterval option in RFC 3623 Section B.1. The type is interface{}
    // with range: 1..1800. Units are seconds. The default value is 120.
    RestartInterval interface{}

    // RestartHelperStrictLSAChecking option in RFC 3623 Section B.2. The type is
    // bool.
    HelperStrictLsaChecking interface{}
}

func (gracefulRestart *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "ietf"
    gracefulRestart.EntityData.ParentYangName = "instance"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = ietf.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    gracefulRestart.EntityData.Children = make(map[string]types.YChild)
    gracefulRestart.EntityData.Leafs = make(map[string]types.YLeaf)
    gracefulRestart.EntityData.Leafs["enable"] = types.YLeaf{"Enable", gracefulRestart.Enable}
    gracefulRestart.EntityData.Leafs["helper-enable"] = types.YLeaf{"HelperEnable", gracefulRestart.HelperEnable}
    gracefulRestart.EntityData.Leafs["restart-interval"] = types.YLeaf{"RestartInterval", gracefulRestart.RestartInterval}
    gracefulRestart.EntityData.Leafs["helper-strict-lsa-checking"] = types.YLeaf{"HelperStrictLsaChecking", gracefulRestart.HelperStrictLsaChecking}
    return &(gracefulRestart.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost
// Auto cost config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable auto cost. The type is bool.
    Enable interface{}

    // Configure reference bandwidth in term of Mbits. The type is interface{}
    // with range: 1..4294967. Units are Mbits.
    ReferenceBandwidth interface{}
}

func (autoCost *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AutoCost) GetEntityData() *types.CommonEntityData {
    autoCost.EntityData.YFilter = autoCost.YFilter
    autoCost.EntityData.YangName = "auto-cost"
    autoCost.EntityData.BundleName = "ietf"
    autoCost.EntityData.ParentYangName = "instance"
    autoCost.EntityData.SegmentPath = "auto-cost"
    autoCost.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    autoCost.EntityData.NamespaceTable = ietf.GetNamespaces()
    autoCost.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    autoCost.EntityData.Children = make(map[string]types.YChild)
    autoCost.EntityData.Leafs = make(map[string]types.YLeaf)
    autoCost.EntityData.Leafs["enable"] = types.YLeaf{"Enable", autoCost.Enable}
    autoCost.EntityData.Leafs["reference-bandwidth"] = types.YLeaf{"ReferenceBandwidth", autoCost.ReferenceBandwidth}
    return &(autoCost.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl
// SPF calculation control.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of ECMP paths. The type is interface{} with range: 1..32.
    Paths interface{}
}

func (spfControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_SpfControl) GetEntityData() *types.CommonEntityData {
    spfControl.EntityData.YFilter = spfControl.YFilter
    spfControl.EntityData.YangName = "spf-control"
    spfControl.EntityData.BundleName = "ietf"
    spfControl.EntityData.ParentYangName = "instance"
    spfControl.EntityData.SegmentPath = "spf-control"
    spfControl.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    spfControl.EntityData.NamespaceTable = ietf.GetNamespaces()
    spfControl.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    spfControl.EntityData.Children = make(map[string]types.YChild)
    spfControl.EntityData.Leafs = make(map[string]types.YLeaf)
    spfControl.EntityData.Leafs["paths"] = types.YLeaf{"Paths", spfControl.Paths}
    return &(spfControl.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl
// Database maintenance control.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of LSAs OSPF will receive. The type is interface{} with
    // range: 1..4294967294.
    MaxLsa interface{}
}

func (databaseControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_DatabaseControl) GetEntityData() *types.CommonEntityData {
    databaseControl.EntityData.YFilter = databaseControl.YFilter
    databaseControl.EntityData.YangName = "database-control"
    databaseControl.EntityData.BundleName = "ietf"
    databaseControl.EntityData.ParentYangName = "instance"
    databaseControl.EntityData.SegmentPath = "database-control"
    databaseControl.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    databaseControl.EntityData.NamespaceTable = ietf.GetNamespaces()
    databaseControl.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    databaseControl.EntityData.Children = make(map[string]types.YChild)
    databaseControl.EntityData.Leafs = make(map[string]types.YLeaf)
    databaseControl.EntityData.Leafs["max-lsa"] = types.YLeaf{"MaxLsa", databaseControl.MaxLsa}
    return &(databaseControl.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl
// Protocol reload control.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (reloadControl *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_ReloadControl) GetEntityData() *types.CommonEntityData {
    reloadControl.EntityData.YFilter = reloadControl.YFilter
    reloadControl.EntityData.YangName = "reload-control"
    reloadControl.EntityData.BundleName = "ietf"
    reloadControl.EntityData.ParentYangName = "instance"
    reloadControl.EntityData.SegmentPath = "reload-control"
    reloadControl.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    reloadControl.EntityData.NamespaceTable = ietf.GetNamespaces()
    reloadControl.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    reloadControl.EntityData.Children = make(map[string]types.YChild)
    reloadControl.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reloadControl.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls
// OSPF MPLS config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traffic Engineering stable IP address for system.
    TeRid Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid

    // OSPF MPLS LDP config state.
    Ldp Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp
}

func (mpls *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls) GetEntityData() *types.CommonEntityData {
    mpls.EntityData.YFilter = mpls.YFilter
    mpls.EntityData.YangName = "mpls"
    mpls.EntityData.BundleName = "ietf"
    mpls.EntityData.ParentYangName = "instance"
    mpls.EntityData.SegmentPath = "mpls"
    mpls.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    mpls.EntityData.NamespaceTable = ietf.GetNamespaces()
    mpls.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    mpls.EntityData.Children = make(map[string]types.YChild)
    mpls.EntityData.Children["te-rid"] = types.YChild{"TeRid", &mpls.TeRid}
    mpls.EntityData.Children["ldp"] = types.YChild{"Ldp", &mpls.Ldp}
    mpls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mpls.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid
// Traffic Engineering stable IP address for system.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Take the interface's IPv4 address as TE router ID. The type is string.
    // Refers to interfaces.Interfaces_Interface_Name
    Interface_ interface{}

    // Explicitly configure the TE router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}
}

func (teRid *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_TeRid) GetEntityData() *types.CommonEntityData {
    teRid.EntityData.YFilter = teRid.YFilter
    teRid.EntityData.YangName = "te-rid"
    teRid.EntityData.BundleName = "ietf"
    teRid.EntityData.ParentYangName = "mpls"
    teRid.EntityData.SegmentPath = "te-rid"
    teRid.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    teRid.EntityData.NamespaceTable = ietf.GetNamespaces()
    teRid.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    teRid.EntityData.Children = make(map[string]types.YChild)
    teRid.EntityData.Leafs = make(map[string]types.YLeaf)
    teRid.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", teRid.Interface_}
    teRid.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", teRid.RouterId}
    return &(teRid.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp
// OSPF MPLS LDP config state.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable LDP IGP synchronization. The type is bool.
    IgpSync interface{}

    // Enable LDP IGP interface auto-configuration. The type is bool.
    Autoconfig interface{}
}

func (ldp *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Mpls_Ldp) GetEntityData() *types.CommonEntityData {
    ldp.EntityData.YFilter = ldp.YFilter
    ldp.EntityData.YangName = "ldp"
    ldp.EntityData.BundleName = "ietf"
    ldp.EntityData.ParentYangName = "mpls"
    ldp.EntityData.SegmentPath = "ldp"
    ldp.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ldp.EntityData.NamespaceTable = ietf.GetNamespaces()
    ldp.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ldp.EntityData.Children = make(map[string]types.YChild)
    ldp.EntityData.Leafs = make(map[string]types.YLeaf)
    ldp.EntityData.Leafs["igp-sync"] = types.YLeaf{"IgpSync", ldp.IgpSync}
    ldp.EntityData.Leafs["autoconfig"] = types.YLeaf{"Autoconfig", ldp.Autoconfig}
    return &(ldp.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute
// This container may be augmented with global
// parameters for IPFRR.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This container may be augmented with global parameters for LFA. Creating
    // the container has no effect on LFA activation.
    Lfa Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "ietf"
    fastReroute.EntityData.ParentYangName = "instance"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = ietf.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    fastReroute.EntityData.Children = make(map[string]types.YChild)
    fastReroute.EntityData.Children["lfa"] = types.YChild{"Lfa", &fastReroute.Lfa}
    fastReroute.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fastReroute.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa
// This container may be augmented with
// global parameters for LFA.
// Creating the container has no effect on
// LFA activation.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_FastReroute_Lfa) GetEntityData() *types.CommonEntityData {
    lfa.EntityData.YFilter = lfa.YFilter
    lfa.EntityData.YangName = "lfa"
    lfa.EntityData.BundleName = "ietf"
    lfa.EntityData.ParentYangName = "fast-reroute"
    lfa.EntityData.SegmentPath = "lfa"
    lfa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    lfa.EntityData.NamespaceTable = ietf.GetNamespaces()
    lfa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    lfa.EntityData.Children = make(map[string]types.YChild)
    lfa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lfa.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit
// Inheritance for all areas.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Area config to be inherited by all areas.
    Area Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area

    // Interface config to be inherited by all interfaces in all areas.
    Interface_ Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface
}

func (allAreasInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit) GetEntityData() *types.CommonEntityData {
    allAreasInherit.EntityData.YFilter = allAreasInherit.YFilter
    allAreasInherit.EntityData.YangName = "all-areas-inherit"
    allAreasInherit.EntityData.BundleName = "ietf"
    allAreasInherit.EntityData.ParentYangName = "instance"
    allAreasInherit.EntityData.SegmentPath = "all-areas-inherit"
    allAreasInherit.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    allAreasInherit.EntityData.NamespaceTable = ietf.GetNamespaces()
    allAreasInherit.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    allAreasInherit.EntityData.Children = make(map[string]types.YChild)
    allAreasInherit.EntityData.Children["area"] = types.YChild{"Area", &allAreasInherit.Area}
    allAreasInherit.EntityData.Children["interface"] = types.YChild{"Interface_", &allAreasInherit.Interface_}
    allAreasInherit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(allAreasInherit.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area
// Area config to be inherited by all areas.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Area) GetEntityData() *types.CommonEntityData {
    area.EntityData.YFilter = area.YFilter
    area.EntityData.YangName = "area"
    area.EntityData.BundleName = "ietf"
    area.EntityData.ParentYangName = "all-areas-inherit"
    area.EntityData.SegmentPath = "area"
    area.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    area.EntityData.NamespaceTable = ietf.GetNamespaces()
    area.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    area.EntityData.Children = make(map[string]types.YChild)
    area.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(area.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface
// Interface config to be inherited by all interfaces
// in all areas.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_AllAreasInherit_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "ietf"
    self.EntityData.ParentYangName = "all-areas-inherit"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    self.EntityData.NamespaceTable = ietf.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area
// List of ospf areas
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Area ID. The type is one of the following types:
    // int with range: 0..4294967295, or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    AreaId interface{}

    // Area type. The type is one of the following: NormalStubNssa. The default
    // value is normal.
    AreaType interface{}

    // Enable/Disable summary generation to the stub or NSSA area. The type is
    // bool.
    Summary interface{}

    // Set the summary default-cost for a stub or NSSA area. The type is
    // interface{} with range: 1..16777215.
    DefaultCost interface{}

    // Summarize routes matching address/mask (border routers only). The type is
    // slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range_.
    Range_ []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range

    // Inheritance for all interfaces.
    AllInterfacesInherit Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit

    // OSPF virtual link. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink.
    VirtualLink []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink

    // OSPF sham link. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink.
    ShamLink []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink

    // List of OSPF interfaces. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_.
    Interface_ []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area) GetEntityData() *types.CommonEntityData {
    area.EntityData.YFilter = area.YFilter
    area.EntityData.YangName = "area"
    area.EntityData.BundleName = "ietf"
    area.EntityData.ParentYangName = "instance"
    area.EntityData.SegmentPath = "area" + "[area-id='" + fmt.Sprintf("%v", area.AreaId) + "']"
    area.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    area.EntityData.NamespaceTable = ietf.GetNamespaces()
    area.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    area.EntityData.Children = make(map[string]types.YChild)
    area.EntityData.Children["range"] = types.YChild{"Range_", nil}
    for i := range area.Range_ {
        area.EntityData.Children[types.GetSegmentPath(&area.Range_[i])] = types.YChild{"Range_", &area.Range_[i]}
    }
    area.EntityData.Children["all-interfaces-inherit"] = types.YChild{"AllInterfacesInherit", &area.AllInterfacesInherit}
    area.EntityData.Children["virtual-link"] = types.YChild{"VirtualLink", nil}
    for i := range area.VirtualLink {
        area.EntityData.Children[types.GetSegmentPath(&area.VirtualLink[i])] = types.YChild{"VirtualLink", &area.VirtualLink[i]}
    }
    area.EntityData.Children["sham-link"] = types.YChild{"ShamLink", nil}
    for i := range area.ShamLink {
        area.EntityData.Children[types.GetSegmentPath(&area.ShamLink[i])] = types.YChild{"ShamLink", &area.ShamLink[i]}
    }
    area.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range area.Interface_ {
        area.EntityData.Children[types.GetSegmentPath(&area.Interface_[i])] = types.YChild{"Interface_", &area.Interface_[i]}
    }
    area.EntityData.Leafs = make(map[string]types.YLeaf)
    area.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", area.AreaId}
    area.EntityData.Leafs["area-type"] = types.YLeaf{"AreaType", area.AreaType}
    area.EntityData.Leafs["summary"] = types.YLeaf{"Summary", area.Summary}
    area.EntityData.Leafs["default-cost"] = types.YLeaf{"DefaultCost", area.DefaultCost}
    return &(area.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range
// Summarize routes matching address/mask (border
// routers only)
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 or IPv6 prefix. The type is string.
    Prefix interface{}

    // Advertise or hide. The type is bool.
    Advertise interface{}

    // Cost of summary route. The type is interface{} with range: 0..16777214.
    Cost interface{}
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "ietf"
    self.EntityData.ParentYangName = "area"
    self.EntityData.SegmentPath = "range" + "[prefix='" + fmt.Sprintf("%v", self.Prefix) + "']"
    self.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    self.EntityData.NamespaceTable = ietf.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", self.Prefix}
    self.EntityData.Leafs["advertise"] = types.YLeaf{"Advertise", self.Advertise}
    self.EntityData.Leafs["cost"] = types.YLeaf{"Cost", self.Cost}
    return &(self.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit
// Inheritance for all interfaces
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface config to be inherited by all interfaces.
    Interface_ Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface
}

func (allInterfacesInherit *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit) GetEntityData() *types.CommonEntityData {
    allInterfacesInherit.EntityData.YFilter = allInterfacesInherit.YFilter
    allInterfacesInherit.EntityData.YangName = "all-interfaces-inherit"
    allInterfacesInherit.EntityData.BundleName = "ietf"
    allInterfacesInherit.EntityData.ParentYangName = "area"
    allInterfacesInherit.EntityData.SegmentPath = "all-interfaces-inherit"
    allInterfacesInherit.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    allInterfacesInherit.EntityData.NamespaceTable = ietf.GetNamespaces()
    allInterfacesInherit.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    allInterfacesInherit.EntityData.Children = make(map[string]types.YChild)
    allInterfacesInherit.EntityData.Children["interface"] = types.YChild{"Interface_", &allInterfacesInherit.Interface_}
    allInterfacesInherit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(allInterfacesInherit.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface
// Interface config to be inherited by all
// interfaces.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_AllInterfacesInherit_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "ietf"
    self.EntityData.ParentYangName = "all-interfaces-inherit"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    self.EntityData.NamespaceTable = ietf.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink
// OSPF virtual link
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Virtual link router ID. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    RouterId interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Time between hello packets. The type is interface{} with range: 1..65535.
    // Units are seconds.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead. The type is interface{}
    // with range: 1..65535. Units are seconds.
    DeadInterval interface{}

    // Time between retransmitting unacknowledged Link State Advertisements
    // (LSAs). The type is interface{} with range: 1..65535. Units are seconds.
    RetransmitInterval interface{}

    // Estimated time needed to send link-state update. The type is interface{}
    // with range: 1..65535. Units are seconds.
    TransmitDelay interface{}

    // Enable/Disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/Disable link-local signaling (LLS) support. The type is bool.
    Lls interface{}

    // Suppress advertisement of the prefixes. The type is bool.
    PrefixSuppression interface{}

    // Enable/disable bfd. The type is bool.
    Bfd interface{}

    // Enable/disable protocol on the interface. The type is bool. The default
    // value is true.
    Enable interface{}

    // TTL security check.
    TtlSecurity Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity

    // Authentication configuration.
    Authentication Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication
}

func (virtualLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink) GetEntityData() *types.CommonEntityData {
    virtualLink.EntityData.YFilter = virtualLink.YFilter
    virtualLink.EntityData.YangName = "virtual-link"
    virtualLink.EntityData.BundleName = "ietf"
    virtualLink.EntityData.ParentYangName = "area"
    virtualLink.EntityData.SegmentPath = "virtual-link" + "[router-id='" + fmt.Sprintf("%v", virtualLink.RouterId) + "']"
    virtualLink.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    virtualLink.EntityData.NamespaceTable = ietf.GetNamespaces()
    virtualLink.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    virtualLink.EntityData.Children = make(map[string]types.YChild)
    virtualLink.EntityData.Children["ttl-security"] = types.YChild{"TtlSecurity", &virtualLink.TtlSecurity}
    virtualLink.EntityData.Children["authentication"] = types.YChild{"Authentication", &virtualLink.Authentication}
    virtualLink.EntityData.Leafs = make(map[string]types.YLeaf)
    virtualLink.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", virtualLink.RouterId}
    virtualLink.EntityData.Leafs["cost"] = types.YLeaf{"Cost", virtualLink.Cost}
    virtualLink.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", virtualLink.HelloInterval}
    virtualLink.EntityData.Leafs["dead-interval"] = types.YLeaf{"DeadInterval", virtualLink.DeadInterval}
    virtualLink.EntityData.Leafs["retransmit-interval"] = types.YLeaf{"RetransmitInterval", virtualLink.RetransmitInterval}
    virtualLink.EntityData.Leafs["transmit-delay"] = types.YLeaf{"TransmitDelay", virtualLink.TransmitDelay}
    virtualLink.EntityData.Leafs["mtu-ignore"] = types.YLeaf{"MtuIgnore", virtualLink.MtuIgnore}
    virtualLink.EntityData.Leafs["lls"] = types.YLeaf{"Lls", virtualLink.Lls}
    virtualLink.EntityData.Leafs["prefix-suppression"] = types.YLeaf{"PrefixSuppression", virtualLink.PrefixSuppression}
    virtualLink.EntityData.Leafs["bfd"] = types.YLeaf{"Bfd", virtualLink.Bfd}
    virtualLink.EntityData.Leafs["enable"] = types.YLeaf{"Enable", virtualLink.Enable}
    return &(virtualLink.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity
// TTL security check.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable TTL security check. The type is bool.
    Enable interface{}

    // Maximum number of hops that a OSPF packet may have traveled. The type is
    // interface{} with range: 1..254.
    Hops interface{}
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_TtlSecurity) GetEntityData() *types.CommonEntityData {
    ttlSecurity.EntityData.YFilter = ttlSecurity.YFilter
    ttlSecurity.EntityData.YangName = "ttl-security"
    ttlSecurity.EntityData.BundleName = "ietf"
    ttlSecurity.EntityData.ParentYangName = "virtual-link"
    ttlSecurity.EntityData.SegmentPath = "ttl-security"
    ttlSecurity.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ttlSecurity.EntityData.NamespaceTable = ietf.GetNamespaces()
    ttlSecurity.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ttlSecurity.EntityData.Children = make(map[string]types.YChild)
    ttlSecurity.EntityData.Leafs = make(map[string]types.YLeaf)
    ttlSecurity.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ttlSecurity.Enable}
    ttlSecurity.EntityData.Leafs["hops"] = types.YLeaf{"Hops", ttlSecurity.Hops}
    return &(ttlSecurity.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication
// Authentication configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA name. The type is string.
    Sa interface{}

    // key-chain name. The type is string. Refers to key_chain.KeyChains_Name
    KeyChain interface{}

    // Key string in ASCII format. The type is string.
    Key interface{}

    // Cryptographic algorithm associated with key.
    CryptoAlgorithm Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "ietf"
    authentication.EntityData.ParentYangName = "virtual-link"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    authentication.EntityData.NamespaceTable = ietf.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["crypto-algorithm"] = types.YChild{"CryptoAlgorithm", &authentication.CryptoAlgorithm}
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["sa"] = types.YLeaf{"Sa", authentication.Sa}
    authentication.EntityData.Leafs["key-chain"] = types.YLeaf{"KeyChain", authentication.KeyChain}
    authentication.EntityData.Leafs["key"] = types.YLeaf{"Key", authentication.Key}
    return &(authentication.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm
// Cryptographic algorithm associated with key.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // The HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // The MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // The SHA-1 algorithm. The type is interface{}.
    Sha1 interface{}

    // HMAC-SHA-1 authentication algorithm. The type is interface{}.
    HmacSha1 interface{}

    // HMAC-SHA-256 authentication algorithm. The type is interface{}.
    HmacSha256 interface{}

    // HMAC-SHA-384 authentication algorithm. The type is interface{}.
    HmacSha384 interface{}

    // HMAC-SHA-512 authentication algorithm. The type is interface{}.
    HmacSha512 interface{}
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_VirtualLink_Authentication_CryptoAlgorithm) GetEntityData() *types.CommonEntityData {
    cryptoAlgorithm.EntityData.YFilter = cryptoAlgorithm.YFilter
    cryptoAlgorithm.EntityData.YangName = "crypto-algorithm"
    cryptoAlgorithm.EntityData.BundleName = "ietf"
    cryptoAlgorithm.EntityData.ParentYangName = "authentication"
    cryptoAlgorithm.EntityData.SegmentPath = "crypto-algorithm"
    cryptoAlgorithm.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    cryptoAlgorithm.EntityData.NamespaceTable = ietf.GetNamespaces()
    cryptoAlgorithm.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    cryptoAlgorithm.EntityData.Children = make(map[string]types.YChild)
    cryptoAlgorithm.EntityData.Leafs = make(map[string]types.YLeaf)
    cryptoAlgorithm.EntityData.Leafs["hmac-sha1-12"] = types.YLeaf{"HmacSha112", cryptoAlgorithm.HmacSha112}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha1-20"] = types.YLeaf{"HmacSha120", cryptoAlgorithm.HmacSha120}
    cryptoAlgorithm.EntityData.Leafs["md5"] = types.YLeaf{"Md5", cryptoAlgorithm.Md5}
    cryptoAlgorithm.EntityData.Leafs["sha-1"] = types.YLeaf{"Sha1", cryptoAlgorithm.Sha1}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-1"] = types.YLeaf{"HmacSha1", cryptoAlgorithm.HmacSha1}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-256"] = types.YLeaf{"HmacSha256", cryptoAlgorithm.HmacSha256}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-384"] = types.YLeaf{"HmacSha384", cryptoAlgorithm.HmacSha384}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-512"] = types.YLeaf{"HmacSha512", cryptoAlgorithm.HmacSha512}
    return &(cryptoAlgorithm.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink
// OSPF sham link
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address of the local end-point. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalId interface{}

    // This attribute is a key. Address of the remote end-point. The type is one
    // of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RemoteId interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Time between hello packets. The type is interface{} with range: 1..65535.
    // Units are seconds.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead. The type is interface{}
    // with range: 1..65535. Units are seconds.
    DeadInterval interface{}

    // Time between retransmitting unacknowledged Link State Advertisements
    // (LSAs). The type is interface{} with range: 1..65535. Units are seconds.
    RetransmitInterval interface{}

    // Estimated time needed to send link-state update. The type is interface{}
    // with range: 1..65535. Units are seconds.
    TransmitDelay interface{}

    // Enable/Disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/Disable link-local signaling (LLS) support. The type is bool.
    Lls interface{}

    // Suppress advertisement of the prefixes. The type is bool.
    PrefixSuppression interface{}

    // Enable/disable bfd. The type is bool.
    Bfd interface{}

    // Enable/disable protocol on the interface. The type is bool. The default
    // value is true.
    Enable interface{}

    // TTL security check.
    TtlSecurity Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity

    // Authentication configuration.
    Authentication Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication
}

func (shamLink *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink) GetEntityData() *types.CommonEntityData {
    shamLink.EntityData.YFilter = shamLink.YFilter
    shamLink.EntityData.YangName = "sham-link"
    shamLink.EntityData.BundleName = "ietf"
    shamLink.EntityData.ParentYangName = "area"
    shamLink.EntityData.SegmentPath = "sham-link" + "[local-id='" + fmt.Sprintf("%v", shamLink.LocalId) + "']" + "[remote-id='" + fmt.Sprintf("%v", shamLink.RemoteId) + "']"
    shamLink.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    shamLink.EntityData.NamespaceTable = ietf.GetNamespaces()
    shamLink.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    shamLink.EntityData.Children = make(map[string]types.YChild)
    shamLink.EntityData.Children["ttl-security"] = types.YChild{"TtlSecurity", &shamLink.TtlSecurity}
    shamLink.EntityData.Children["authentication"] = types.YChild{"Authentication", &shamLink.Authentication}
    shamLink.EntityData.Leafs = make(map[string]types.YLeaf)
    shamLink.EntityData.Leafs["local-id"] = types.YLeaf{"LocalId", shamLink.LocalId}
    shamLink.EntityData.Leafs["remote-id"] = types.YLeaf{"RemoteId", shamLink.RemoteId}
    shamLink.EntityData.Leafs["cost"] = types.YLeaf{"Cost", shamLink.Cost}
    shamLink.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", shamLink.HelloInterval}
    shamLink.EntityData.Leafs["dead-interval"] = types.YLeaf{"DeadInterval", shamLink.DeadInterval}
    shamLink.EntityData.Leafs["retransmit-interval"] = types.YLeaf{"RetransmitInterval", shamLink.RetransmitInterval}
    shamLink.EntityData.Leafs["transmit-delay"] = types.YLeaf{"TransmitDelay", shamLink.TransmitDelay}
    shamLink.EntityData.Leafs["mtu-ignore"] = types.YLeaf{"MtuIgnore", shamLink.MtuIgnore}
    shamLink.EntityData.Leafs["lls"] = types.YLeaf{"Lls", shamLink.Lls}
    shamLink.EntityData.Leafs["prefix-suppression"] = types.YLeaf{"PrefixSuppression", shamLink.PrefixSuppression}
    shamLink.EntityData.Leafs["bfd"] = types.YLeaf{"Bfd", shamLink.Bfd}
    shamLink.EntityData.Leafs["enable"] = types.YLeaf{"Enable", shamLink.Enable}
    return &(shamLink.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity
// TTL security check.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable TTL security check. The type is bool.
    Enable interface{}

    // Maximum number of hops that a OSPF packet may have traveled. The type is
    // interface{} with range: 1..254.
    Hops interface{}
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_TtlSecurity) GetEntityData() *types.CommonEntityData {
    ttlSecurity.EntityData.YFilter = ttlSecurity.YFilter
    ttlSecurity.EntityData.YangName = "ttl-security"
    ttlSecurity.EntityData.BundleName = "ietf"
    ttlSecurity.EntityData.ParentYangName = "sham-link"
    ttlSecurity.EntityData.SegmentPath = "ttl-security"
    ttlSecurity.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ttlSecurity.EntityData.NamespaceTable = ietf.GetNamespaces()
    ttlSecurity.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ttlSecurity.EntityData.Children = make(map[string]types.YChild)
    ttlSecurity.EntityData.Leafs = make(map[string]types.YLeaf)
    ttlSecurity.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ttlSecurity.Enable}
    ttlSecurity.EntityData.Leafs["hops"] = types.YLeaf{"Hops", ttlSecurity.Hops}
    return &(ttlSecurity.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication
// Authentication configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA name. The type is string.
    Sa interface{}

    // key-chain name. The type is string. Refers to key_chain.KeyChains_Name
    KeyChain interface{}

    // Key string in ASCII format. The type is string.
    Key interface{}

    // Cryptographic algorithm associated with key.
    CryptoAlgorithm Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "ietf"
    authentication.EntityData.ParentYangName = "sham-link"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    authentication.EntityData.NamespaceTable = ietf.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["crypto-algorithm"] = types.YChild{"CryptoAlgorithm", &authentication.CryptoAlgorithm}
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["sa"] = types.YLeaf{"Sa", authentication.Sa}
    authentication.EntityData.Leafs["key-chain"] = types.YLeaf{"KeyChain", authentication.KeyChain}
    authentication.EntityData.Leafs["key"] = types.YLeaf{"Key", authentication.Key}
    return &(authentication.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm
// Cryptographic algorithm associated with key.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // The HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // The MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // The SHA-1 algorithm. The type is interface{}.
    Sha1 interface{}

    // HMAC-SHA-1 authentication algorithm. The type is interface{}.
    HmacSha1 interface{}

    // HMAC-SHA-256 authentication algorithm. The type is interface{}.
    HmacSha256 interface{}

    // HMAC-SHA-384 authentication algorithm. The type is interface{}.
    HmacSha384 interface{}

    // HMAC-SHA-512 authentication algorithm. The type is interface{}.
    HmacSha512 interface{}
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_ShamLink_Authentication_CryptoAlgorithm) GetEntityData() *types.CommonEntityData {
    cryptoAlgorithm.EntityData.YFilter = cryptoAlgorithm.YFilter
    cryptoAlgorithm.EntityData.YangName = "crypto-algorithm"
    cryptoAlgorithm.EntityData.BundleName = "ietf"
    cryptoAlgorithm.EntityData.ParentYangName = "authentication"
    cryptoAlgorithm.EntityData.SegmentPath = "crypto-algorithm"
    cryptoAlgorithm.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    cryptoAlgorithm.EntityData.NamespaceTable = ietf.GetNamespaces()
    cryptoAlgorithm.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    cryptoAlgorithm.EntityData.Children = make(map[string]types.YChild)
    cryptoAlgorithm.EntityData.Leafs = make(map[string]types.YLeaf)
    cryptoAlgorithm.EntityData.Leafs["hmac-sha1-12"] = types.YLeaf{"HmacSha112", cryptoAlgorithm.HmacSha112}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha1-20"] = types.YLeaf{"HmacSha120", cryptoAlgorithm.HmacSha120}
    cryptoAlgorithm.EntityData.Leafs["md5"] = types.YLeaf{"Md5", cryptoAlgorithm.Md5}
    cryptoAlgorithm.EntityData.Leafs["sha-1"] = types.YLeaf{"Sha1", cryptoAlgorithm.Sha1}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-1"] = types.YLeaf{"HmacSha1", cryptoAlgorithm.HmacSha1}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-256"] = types.YLeaf{"HmacSha256", cryptoAlgorithm.HmacSha256}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-384"] = types.YLeaf{"HmacSha384", cryptoAlgorithm.HmacSha384}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-512"] = types.YLeaf{"HmacSha512", cryptoAlgorithm.HmacSha512}
    return &(cryptoAlgorithm.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface
// List of OSPF interfaces.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Interface_ interface{}

    // Network type. The type is NetworkType.
    NetworkType interface{}

    // Enable/Disable passive. The type is bool.
    Passive interface{}

    // Enable/Disable demand circuit. The type is bool.
    DemandCircuit interface{}

    // Set prefix as a node representative prefix. The type is bool. The default
    // value is false.
    NodeFlag interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Time between hello packets. The type is interface{} with range: 1..65535.
    // Units are seconds.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead. The type is interface{}
    // with range: 1..65535. Units are seconds.
    DeadInterval interface{}

    // Time between retransmitting unacknowledged Link State Advertisements
    // (LSAs). The type is interface{} with range: 1..65535. Units are seconds.
    RetransmitInterval interface{}

    // Estimated time needed to send link-state update. The type is interface{}
    // with range: 1..65535. Units are seconds.
    TransmitDelay interface{}

    // Enable/Disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/Disable link-local signaling (LLS) support. The type is bool.
    Lls interface{}

    // Suppress advertisement of the prefixes. The type is bool.
    PrefixSuppression interface{}

    // Enable/disable bfd. The type is bool.
    Bfd interface{}

    // Enable/disable protocol on the interface. The type is bool. The default
    // value is true.
    Enable interface{}

    // Configure ospf multi-area.
    MultiArea Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea

    // Static configured neighbors.
    StaticNeighbors Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors

    // Fast-reroute configuration.
    FastReroute Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute

    // TTL security check.
    TtlSecurity Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity

    // Authentication configuration.
    Authentication Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication

    // OSPF interface topology. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology.
    Topology []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "ietf"
    self.EntityData.ParentYangName = "area"
    self.EntityData.SegmentPath = "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface_) + "']"
    self.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    self.EntityData.NamespaceTable = ietf.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["multi-area"] = types.YChild{"MultiArea", &self.MultiArea}
    self.EntityData.Children["static-neighbors"] = types.YChild{"StaticNeighbors", &self.StaticNeighbors}
    self.EntityData.Children["fast-reroute"] = types.YChild{"FastReroute", &self.FastReroute}
    self.EntityData.Children["ttl-security"] = types.YChild{"TtlSecurity", &self.TtlSecurity}
    self.EntityData.Children["authentication"] = types.YChild{"Authentication", &self.Authentication}
    self.EntityData.Children["topology"] = types.YChild{"Topology", nil}
    for i := range self.Topology {
        self.EntityData.Children[types.GetSegmentPath(&self.Topology[i])] = types.YChild{"Topology", &self.Topology[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["network-type"] = types.YLeaf{"NetworkType", self.NetworkType}
    self.EntityData.Leafs["passive"] = types.YLeaf{"Passive", self.Passive}
    self.EntityData.Leafs["demand-circuit"] = types.YLeaf{"DemandCircuit", self.DemandCircuit}
    self.EntityData.Leafs["node-flag"] = types.YLeaf{"NodeFlag", self.NodeFlag}
    self.EntityData.Leafs["cost"] = types.YLeaf{"Cost", self.Cost}
    self.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", self.HelloInterval}
    self.EntityData.Leafs["dead-interval"] = types.YLeaf{"DeadInterval", self.DeadInterval}
    self.EntityData.Leafs["retransmit-interval"] = types.YLeaf{"RetransmitInterval", self.RetransmitInterval}
    self.EntityData.Leafs["transmit-delay"] = types.YLeaf{"TransmitDelay", self.TransmitDelay}
    self.EntityData.Leafs["mtu-ignore"] = types.YLeaf{"MtuIgnore", self.MtuIgnore}
    self.EntityData.Leafs["lls"] = types.YLeaf{"Lls", self.Lls}
    self.EntityData.Leafs["prefix-suppression"] = types.YLeaf{"PrefixSuppression", self.PrefixSuppression}
    self.EntityData.Leafs["bfd"] = types.YLeaf{"Bfd", self.Bfd}
    self.EntityData.Leafs["enable"] = types.YLeaf{"Enable", self.Enable}
    return &(self.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea
// Configure ospf multi-area.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multi-area ID. The type is one of the following types: int with range:
    // 0..4294967295, or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    MultiAreaId interface{}

    // Interface cost for multi-area. The type is interface{} with range:
    // 0..65535.
    Cost interface{}
}

func (multiArea *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_MultiArea) GetEntityData() *types.CommonEntityData {
    multiArea.EntityData.YFilter = multiArea.YFilter
    multiArea.EntityData.YangName = "multi-area"
    multiArea.EntityData.BundleName = "ietf"
    multiArea.EntityData.ParentYangName = "interface"
    multiArea.EntityData.SegmentPath = "multi-area"
    multiArea.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    multiArea.EntityData.NamespaceTable = ietf.GetNamespaces()
    multiArea.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    multiArea.EntityData.Children = make(map[string]types.YChild)
    multiArea.EntityData.Leafs = make(map[string]types.YLeaf)
    multiArea.EntityData.Leafs["multi-area-id"] = types.YLeaf{"MultiAreaId", multiArea.MultiAreaId}
    multiArea.EntityData.Leafs["cost"] = types.YLeaf{"Cost", multiArea.Cost}
    return &(multiArea.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors
// Static configured neighbors.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify a neighbor router. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor.
    Neighbor []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor
}

func (staticNeighbors *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors) GetEntityData() *types.CommonEntityData {
    staticNeighbors.EntityData.YFilter = staticNeighbors.YFilter
    staticNeighbors.EntityData.YangName = "static-neighbors"
    staticNeighbors.EntityData.BundleName = "ietf"
    staticNeighbors.EntityData.ParentYangName = "interface"
    staticNeighbors.EntityData.SegmentPath = "static-neighbors"
    staticNeighbors.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    staticNeighbors.EntityData.NamespaceTable = ietf.GetNamespaces()
    staticNeighbors.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    staticNeighbors.EntityData.Children = make(map[string]types.YChild)
    staticNeighbors.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range staticNeighbors.Neighbor {
        staticNeighbors.EntityData.Children[types.GetSegmentPath(&staticNeighbors.Neighbor[i])] = types.YChild{"Neighbor", &staticNeighbors.Neighbor[i]}
    }
    staticNeighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticNeighbors.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor
// Specify a neighbor router.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Neighbor cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Neighbor poll interval. The type is interface{} with range: 1..65535. Units
    // are seconds.
    PollInterval interface{}

    // Neighbor priority for DR election. The type is interface{} with range:
    // 1..255.
    Priority interface{}
}

func (neighbor *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_StaticNeighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "ietf"
    neighbor.EntityData.ParentYangName = "static-neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + "[address='" + fmt.Sprintf("%v", neighbor.Address) + "']"
    neighbor.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    neighbor.EntityData.NamespaceTable = ietf.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["address"] = types.YLeaf{"Address", neighbor.Address}
    neighbor.EntityData.Leafs["cost"] = types.YLeaf{"Cost", neighbor.Cost}
    neighbor.EntityData.Leafs["poll-interval"] = types.YLeaf{"PollInterval", neighbor.PollInterval}
    neighbor.EntityData.Leafs["priority"] = types.YLeaf{"Priority", neighbor.Priority}
    return &(neighbor.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute
// Fast-reroute configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LFA configuration.
    Lfa Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa
}

func (fastReroute *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "ietf"
    fastReroute.EntityData.ParentYangName = "interface"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = ietf.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    fastReroute.EntityData.Children = make(map[string]types.YChild)
    fastReroute.EntityData.Children["lfa"] = types.YChild{"Lfa", &fastReroute.Lfa}
    fastReroute.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fastReroute.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa
// LFA configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prevent the interface to be used as backup. The type is bool.
    CandidateDisabled interface{}

    // Activates LFA. This model assumes activation of per-prefix LFA. The type is
    // bool.
    Enabled interface{}

    // Remote LFA configuration.
    RemoteLfa Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa
}

func (lfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa) GetEntityData() *types.CommonEntityData {
    lfa.EntityData.YFilter = lfa.YFilter
    lfa.EntityData.YangName = "lfa"
    lfa.EntityData.BundleName = "ietf"
    lfa.EntityData.ParentYangName = "fast-reroute"
    lfa.EntityData.SegmentPath = "lfa"
    lfa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    lfa.EntityData.NamespaceTable = ietf.GetNamespaces()
    lfa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    lfa.EntityData.Children = make(map[string]types.YChild)
    lfa.EntityData.Children["remote-lfa"] = types.YChild{"RemoteLfa", &lfa.RemoteLfa}
    lfa.EntityData.Leafs = make(map[string]types.YLeaf)
    lfa.EntityData.Leafs["candidate-disabled"] = types.YLeaf{"CandidateDisabled", lfa.CandidateDisabled}
    lfa.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", lfa.Enabled}
    return &(lfa.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa
// Remote LFA configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Activates remote LFA. The type is bool.
    Enabled interface{}
}

func (remoteLfa *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_FastReroute_Lfa_RemoteLfa) GetEntityData() *types.CommonEntityData {
    remoteLfa.EntityData.YFilter = remoteLfa.YFilter
    remoteLfa.EntityData.YangName = "remote-lfa"
    remoteLfa.EntityData.BundleName = "ietf"
    remoteLfa.EntityData.ParentYangName = "lfa"
    remoteLfa.EntityData.SegmentPath = "remote-lfa"
    remoteLfa.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    remoteLfa.EntityData.NamespaceTable = ietf.GetNamespaces()
    remoteLfa.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    remoteLfa.EntityData.Children = make(map[string]types.YChild)
    remoteLfa.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteLfa.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", remoteLfa.Enabled}
    return &(remoteLfa.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity
// TTL security check.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable TTL security check. The type is bool.
    Enable interface{}

    // Maximum number of hops that a OSPF packet may have traveled. The type is
    // interface{} with range: 1..254.
    Hops interface{}
}

func (ttlSecurity *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_TtlSecurity) GetEntityData() *types.CommonEntityData {
    ttlSecurity.EntityData.YFilter = ttlSecurity.YFilter
    ttlSecurity.EntityData.YangName = "ttl-security"
    ttlSecurity.EntityData.BundleName = "ietf"
    ttlSecurity.EntityData.ParentYangName = "interface"
    ttlSecurity.EntityData.SegmentPath = "ttl-security"
    ttlSecurity.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ttlSecurity.EntityData.NamespaceTable = ietf.GetNamespaces()
    ttlSecurity.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ttlSecurity.EntityData.Children = make(map[string]types.YChild)
    ttlSecurity.EntityData.Leafs = make(map[string]types.YLeaf)
    ttlSecurity.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ttlSecurity.Enable}
    ttlSecurity.EntityData.Leafs["hops"] = types.YLeaf{"Hops", ttlSecurity.Hops}
    return &(ttlSecurity.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication
// Authentication configuration.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA name. The type is string.
    Sa interface{}

    // key-chain name. The type is string. Refers to key_chain.KeyChains_Name
    KeyChain interface{}

    // Key string in ASCII format. The type is string.
    Key interface{}

    // Cryptographic algorithm associated with key.
    CryptoAlgorithm Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm
}

func (authentication *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "ietf"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    authentication.EntityData.NamespaceTable = ietf.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["crypto-algorithm"] = types.YChild{"CryptoAlgorithm", &authentication.CryptoAlgorithm}
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["sa"] = types.YLeaf{"Sa", authentication.Sa}
    authentication.EntityData.Leafs["key-chain"] = types.YLeaf{"KeyChain", authentication.KeyChain}
    authentication.EntityData.Leafs["key"] = types.YLeaf{"Key", authentication.Key}
    return &(authentication.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm
// Cryptographic algorithm associated with key.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // The HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // The MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // The SHA-1 algorithm. The type is interface{}.
    Sha1 interface{}

    // HMAC-SHA-1 authentication algorithm. The type is interface{}.
    HmacSha1 interface{}

    // HMAC-SHA-256 authentication algorithm. The type is interface{}.
    HmacSha256 interface{}

    // HMAC-SHA-384 authentication algorithm. The type is interface{}.
    HmacSha384 interface{}

    // HMAC-SHA-512 authentication algorithm. The type is interface{}.
    HmacSha512 interface{}
}

func (cryptoAlgorithm *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Authentication_CryptoAlgorithm) GetEntityData() *types.CommonEntityData {
    cryptoAlgorithm.EntityData.YFilter = cryptoAlgorithm.YFilter
    cryptoAlgorithm.EntityData.YangName = "crypto-algorithm"
    cryptoAlgorithm.EntityData.BundleName = "ietf"
    cryptoAlgorithm.EntityData.ParentYangName = "authentication"
    cryptoAlgorithm.EntityData.SegmentPath = "crypto-algorithm"
    cryptoAlgorithm.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    cryptoAlgorithm.EntityData.NamespaceTable = ietf.GetNamespaces()
    cryptoAlgorithm.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    cryptoAlgorithm.EntityData.Children = make(map[string]types.YChild)
    cryptoAlgorithm.EntityData.Leafs = make(map[string]types.YLeaf)
    cryptoAlgorithm.EntityData.Leafs["hmac-sha1-12"] = types.YLeaf{"HmacSha112", cryptoAlgorithm.HmacSha112}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha1-20"] = types.YLeaf{"HmacSha120", cryptoAlgorithm.HmacSha120}
    cryptoAlgorithm.EntityData.Leafs["md5"] = types.YLeaf{"Md5", cryptoAlgorithm.Md5}
    cryptoAlgorithm.EntityData.Leafs["sha-1"] = types.YLeaf{"Sha1", cryptoAlgorithm.Sha1}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-1"] = types.YLeaf{"HmacSha1", cryptoAlgorithm.HmacSha1}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-256"] = types.YLeaf{"HmacSha256", cryptoAlgorithm.HmacSha256}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-384"] = types.YLeaf{"HmacSha384", cryptoAlgorithm.HmacSha384}
    cryptoAlgorithm.EntityData.Leafs["hmac-sha-512"] = types.YLeaf{"HmacSha512", cryptoAlgorithm.HmacSha512}
    return &(cryptoAlgorithm.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology
// OSPF interface topology.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string. Refers to routing.Routing_RoutingInstance_Ribs_Rib_Name
    Name interface{}

    // Interface cost for this topology. The type is interface{} with range:
    // 0..4294967295.
    Cost interface{}
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "interface"
    topology.EntityData.SegmentPath = "topology" + "[name='" + fmt.Sprintf("%v", topology.Name) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["name"] = types.YLeaf{"Name", topology.Name}
    topology.EntityData.Leafs["cost"] = types.YLeaf{"Cost", topology.Cost}
    return &(topology.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType represents Network type.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType string

const (
    // Specify OSPF broadcast multi-access network.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType_broadcast Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType = "broadcast"

    // Specify OSPF Non-Broadcast Multi-Access
    // (NBMA) network.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType_non_broadcast Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType = "non-broadcast"

    // Specify OSPF point-to-multipoint network.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType_point_to_multipoint Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType = "point-to-multipoint"

    // Specify OSPF point-to-point network.
    Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType_point_to_point Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Area_Interface_NetworkType = "point-to-point"
)

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology
// OSPF topology.
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. RIB. The type is string. Refers to
    // routing.Routing_RoutingInstance_Ribs_Rib_Name
    Name interface{}

    // List of ospf areas. The type is slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area.
    Area []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area
}

func (topology *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology) GetEntityData() *types.CommonEntityData {
    topology.EntityData.YFilter = topology.YFilter
    topology.EntityData.YangName = "topology"
    topology.EntityData.BundleName = "ietf"
    topology.EntityData.ParentYangName = "instance"
    topology.EntityData.SegmentPath = "topology" + "[name='" + fmt.Sprintf("%v", topology.Name) + "']"
    topology.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    topology.EntityData.NamespaceTable = ietf.GetNamespaces()
    topology.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    topology.EntityData.Children = make(map[string]types.YChild)
    topology.EntityData.Children["area"] = types.YChild{"Area", nil}
    for i := range topology.Area {
        topology.EntityData.Children[types.GetSegmentPath(&topology.Area[i])] = types.YChild{"Area", &topology.Area[i]}
    }
    topology.EntityData.Leafs = make(map[string]types.YLeaf)
    topology.EntityData.Leafs["name"] = types.YLeaf{"Name", topology.Name}
    return &(topology.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area
// List of ospf areas
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Area ID. The type is one of the following types:
    // int with range: 0..4294967295, or string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    AreaId interface{}

    // Area type. The type is one of the following: NormalStubNssa. The default
    // value is normal.
    AreaType interface{}

    // Enable/Disable summary generation to the stub or NSSA area. The type is
    // bool.
    Summary interface{}

    // Set the summary default-cost for a stub or NSSA area. The type is
    // interface{} with range: 1..16777215.
    DefaultCost interface{}

    // Summarize routes matching address/mask (border routers only). The type is
    // slice of
    // Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range_.
    Range_ []Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range
}

func (area *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area) GetEntityData() *types.CommonEntityData {
    area.EntityData.YFilter = area.YFilter
    area.EntityData.YangName = "area"
    area.EntityData.BundleName = "ietf"
    area.EntityData.ParentYangName = "topology"
    area.EntityData.SegmentPath = "area" + "[area-id='" + fmt.Sprintf("%v", area.AreaId) + "']"
    area.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    area.EntityData.NamespaceTable = ietf.GetNamespaces()
    area.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    area.EntityData.Children = make(map[string]types.YChild)
    area.EntityData.Children["range"] = types.YChild{"Range_", nil}
    for i := range area.Range_ {
        area.EntityData.Children[types.GetSegmentPath(&area.Range_[i])] = types.YChild{"Range_", &area.Range_[i]}
    }
    area.EntityData.Leafs = make(map[string]types.YLeaf)
    area.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", area.AreaId}
    area.EntityData.Leafs["area-type"] = types.YLeaf{"AreaType", area.AreaType}
    area.EntityData.Leafs["summary"] = types.YLeaf{"Summary", area.Summary}
    area.EntityData.Leafs["default-cost"] = types.YLeaf{"DefaultCost", area.DefaultCost}
    return &(area.EntityData)
}

// Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range
// Summarize routes matching address/mask (border
// routers only)
type Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 or IPv6 prefix. The type is string.
    Prefix interface{}

    // Advertise or hide. The type is bool.
    Advertise interface{}

    // Cost of summary route. The type is interface{} with range: 0..16777214.
    Cost interface{}
}

func (self *Routing_RoutingInstance_RoutingProtocols_RoutingProtocol_Ospf_Instance_Topology_Area_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "ietf"
    self.EntityData.ParentYangName = "area"
    self.EntityData.SegmentPath = "range" + "[prefix='" + fmt.Sprintf("%v", self.Prefix) + "']"
    self.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    self.EntityData.NamespaceTable = ietf.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", self.Prefix}
    self.EntityData.Leafs["advertise"] = types.YLeaf{"Advertise", self.Advertise}
    self.EntityData.Leafs["cost"] = types.YLeaf{"Cost", self.Cost}
    return &(self.EntityData)
}

// Routing_RoutingInstance_Ribs
// Configuration of RIBs.
type Routing_RoutingInstance_Ribs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains configuration for a RIB identified by the 'name' key. 
    // Entries having the same key as a system-controlled entry of the list
    // /routing-state/routing-instance/ribs/rib are used for configuring
    // parameters of that entry. Other entries define additional user-controlled
    // RIBs. The type is slice of Routing_RoutingInstance_Ribs_Rib.
    Rib []Routing_RoutingInstance_Ribs_Rib
}

func (ribs *Routing_RoutingInstance_Ribs) GetEntityData() *types.CommonEntityData {
    ribs.EntityData.YFilter = ribs.YFilter
    ribs.EntityData.YangName = "ribs"
    ribs.EntityData.BundleName = "ietf"
    ribs.EntityData.ParentYangName = "routing-instance"
    ribs.EntityData.SegmentPath = "ribs"
    ribs.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ribs.EntityData.NamespaceTable = ietf.GetNamespaces()
    ribs.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ribs.EntityData.Children = make(map[string]types.YChild)
    ribs.EntityData.Children["rib"] = types.YChild{"Rib", nil}
    for i := range ribs.Rib {
        ribs.EntityData.Children[types.GetSegmentPath(&ribs.Rib[i])] = types.YChild{"Rib", &ribs.Rib[i]}
    }
    ribs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ribs.EntityData)
}

// Routing_RoutingInstance_Ribs_Rib
// Each entry contains configuration for a RIB identified
// by the 'name' key.
// 
// Entries having the same key as a system-controlled entry
// of the list /routing-state/routing-instance/ribs/rib are
// used for configuring parameters of that entry. Other
// entries define additional user-controlled RIBs.
type Routing_RoutingInstance_Ribs_Rib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the RIB.  For system-controlled
    // entries, the value of this leaf must be the same as the name of the
    // corresponding entry in state data.  For user-controlled entries, an
    // arbitrary name can be used. The type is string.
    Name interface{}

    // Address family. The type is one of the following:
    // Ipv4Ipv4UnicastIpv6Ipv6Unicast.
    AddressFamily interface{}

    // Textual description of the RIB. The type is string.
    Description interface{}
}

func (rib *Routing_RoutingInstance_Ribs_Rib) GetEntityData() *types.CommonEntityData {
    rib.EntityData.YFilter = rib.YFilter
    rib.EntityData.YangName = "rib"
    rib.EntityData.BundleName = "ietf"
    rib.EntityData.ParentYangName = "ribs"
    rib.EntityData.SegmentPath = "rib" + "[name='" + fmt.Sprintf("%v", rib.Name) + "']"
    rib.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    rib.EntityData.NamespaceTable = ietf.GetNamespaces()
    rib.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    rib.EntityData.Children = make(map[string]types.YChild)
    rib.EntityData.Leafs = make(map[string]types.YLeaf)
    rib.EntityData.Leafs["name"] = types.YLeaf{"Name", rib.Name}
    rib.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", rib.AddressFamily}
    rib.EntityData.Leafs["description"] = types.YLeaf{"Description", rib.Description}
    return &(rib.EntityData)
}

// FibRoute
// Return the active FIB route that a routing-instance uses for
// sending packets to a destination address.
type FibRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input FibRoute_Input

    
    Output FibRoute_Output
}

func (fibRoute *FibRoute) GetEntityData() *types.CommonEntityData {
    fibRoute.EntityData.YFilter = fibRoute.YFilter
    fibRoute.EntityData.YangName = "fib-route"
    fibRoute.EntityData.BundleName = "ietf"
    fibRoute.EntityData.ParentYangName = "ietf-routing"
    fibRoute.EntityData.SegmentPath = "ietf-routing:fib-route"
    fibRoute.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    fibRoute.EntityData.NamespaceTable = ietf.GetNamespaces()
    fibRoute.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    fibRoute.EntityData.Children = make(map[string]types.YChild)
    fibRoute.EntityData.Children["input"] = types.YChild{"Input", &fibRoute.Input}
    fibRoute.EntityData.Children["output"] = types.YChild{"Output", &fibRoute.Output}
    fibRoute.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fibRoute.EntityData)
}

// FibRoute_Input
type FibRoute_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the routing instance whose forwarding information base is being
    // queried.  If the routing instance with name equal to the value of this
    // parameter doesn't exist, then this operation SHALL fail with error-tag
    // 'data-missing' and error-app-tag 'routing-instance-not-found'. The type is
    // string. This attribute is mandatory.
    RoutingInstanceName interface{}

    // Network layer destination address.  Address family specific modules MUST
    // augment this container with a leaf named 'address'.
    DestinationAddress FibRoute_Input_DestinationAddress
}

func (input *FibRoute_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "fib-route"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["destination-address"] = types.YChild{"DestinationAddress", &input.DestinationAddress}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["routing-instance-name"] = types.YLeaf{"RoutingInstanceName", input.RoutingInstanceName}
    return &(input.EntityData)
}

// FibRoute_Input_DestinationAddress
// Network layer destination address.
// 
// Address family specific modules MUST augment this
// container with a leaf named 'address'.
type FibRoute_Input_DestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address family. The type is one of the following:
    // Ipv4Ipv4UnicastIpv6Ipv6Unicast. This attribute is mandatory.
    AddressFamily interface{}

    // IPv4 destination address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IetfIpv4UnicastRoutingAddress interface{}

    // IPv6 destination address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IetfIpv6UnicastRoutingAddress interface{}
}

func (destinationAddress *FibRoute_Input_DestinationAddress) GetEntityData() *types.CommonEntityData {
    destinationAddress.EntityData.YFilter = destinationAddress.YFilter
    destinationAddress.EntityData.YangName = "destination-address"
    destinationAddress.EntityData.BundleName = "ietf"
    destinationAddress.EntityData.ParentYangName = "input"
    destinationAddress.EntityData.SegmentPath = "destination-address"
    destinationAddress.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    destinationAddress.EntityData.NamespaceTable = ietf.GetNamespaces()
    destinationAddress.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    destinationAddress.EntityData.Children = make(map[string]types.YChild)
    destinationAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationAddress.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", destinationAddress.AddressFamily}
    destinationAddress.EntityData.Leafs["address"] = types.YLeaf{"IetfIpv4UnicastRoutingAddress", destinationAddress.IetfIpv4UnicastRoutingAddress}
    destinationAddress.EntityData.Leafs["address"] = types.YLeaf{"IetfIpv6UnicastRoutingAddress", destinationAddress.IetfIpv6UnicastRoutingAddress}
    return &(destinationAddress.EntityData)
}

// FibRoute_Output
type FibRoute_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The active FIB route for the specified destination.  If the routing
    // instance has no active FIB route for the destination address, no output is
    // returned - the server SHALL send an <rpc-reply> containing a single element
    // <ok>.  Address family specific modules MUST augment this list with
    // appropriate route contents.
    Route FibRoute_Output_Route
}

func (output *FibRoute_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "ietf"
    output.EntityData.ParentYangName = "fib-route"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    output.EntityData.NamespaceTable = ietf.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Children["route"] = types.YChild{"Route", &output.Route}
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(output.EntityData)
}

// FibRoute_Output_Route
// The active FIB route for the specified destination.
// 
// If the routing instance has no active FIB route for the
// destination address, no output is returned - the server
// SHALL send an <rpc-reply> containing a single element
// <ok>.
// 
// Address family specific modules MUST augment this list
// with appropriate route contents.
type FibRoute_Output_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address family. The type is one of the following:
    // Ipv4Ipv4UnicastIpv6Ipv6Unicast. This attribute is mandatory.
    AddressFamily interface{}

    // Type of the routing protocol from which the route originated. The type is
    // one of the following: OspfOspfv2Ospfv3DirectStatic. This attribute is
    // mandatory.
    SourceProtocol interface{}

    // Presence of this leaf indicates that the route is preferred among all
    // routes in the same RIB that have the same destination prefix. The type is
    // interface{}.
    Active interface{}

    // Time stamp of the last modification of the route. If the route was never
    // modified, it is the time when the route was inserted into the RIB. The type
    // is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    LastUpdated interface{}

    // IPv4 destination prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))'.
    IetfIpv4UnicastRoutingDestinationPrefix interface{}

    // IPv6 destination prefix. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    IetfIpv6UnicastRoutingDestinationPrefix interface{}

    // Route's next-hop attribute.
    NextHop FibRoute_Output_Route_NextHop
}

func (route *FibRoute_Output_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "ietf"
    route.EntityData.ParentYangName = "output"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    route.EntityData.NamespaceTable = ietf.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    route.EntityData.Children = make(map[string]types.YChild)
    route.EntityData.Children["next-hop"] = types.YChild{"NextHop", &route.NextHop}
    route.EntityData.Leafs = make(map[string]types.YLeaf)
    route.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", route.AddressFamily}
    route.EntityData.Leafs["source-protocol"] = types.YLeaf{"SourceProtocol", route.SourceProtocol}
    route.EntityData.Leafs["active"] = types.YLeaf{"Active", route.Active}
    route.EntityData.Leafs["last-updated"] = types.YLeaf{"LastUpdated", route.LastUpdated}
    route.EntityData.Leafs["destination-prefix"] = types.YLeaf{"IetfIpv4UnicastRoutingDestinationPrefix", route.IetfIpv4UnicastRoutingDestinationPrefix}
    route.EntityData.Leafs["destination-prefix"] = types.YLeaf{"IetfIpv6UnicastRoutingDestinationPrefix", route.IetfIpv6UnicastRoutingDestinationPrefix}
    return &(route.EntityData)
}

// FibRoute_Output_Route_NextHop
// Route's next-hop attribute.
type FibRoute_Output_Route_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the outgoing interface. The type is string.
    OutgoingInterface interface{}

    // IP address. The type is string.
    IetfRoutingNextHopAddress interface{}

    // IPv4 address of the next-hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IetfIpv4UnicastRoutingNextHopAddress interface{}

    // IPv6 address of the next-hop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IetfIpv6UnicastRoutingNextHopAddress interface{}

    // Special next-hop options. The type is SpecialNextHop.
    SpecialNextHop interface{}
}

func (nextHop *FibRoute_Output_Route_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "ietf"
    nextHop.EntityData.ParentYangName = "route"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    nextHop.EntityData.NamespaceTable = ietf.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["outgoing-interface"] = types.YLeaf{"OutgoingInterface", nextHop.OutgoingInterface}
    nextHop.EntityData.Leafs["next-hop-address"] = types.YLeaf{"IetfRoutingNextHopAddress", nextHop.IetfRoutingNextHopAddress}
    nextHop.EntityData.Leafs["next-hop-address"] = types.YLeaf{"IetfIpv4UnicastRoutingNextHopAddress", nextHop.IetfIpv4UnicastRoutingNextHopAddress}
    nextHop.EntityData.Leafs["next-hop-address"] = types.YLeaf{"IetfIpv6UnicastRoutingNextHopAddress", nextHop.IetfIpv6UnicastRoutingNextHopAddress}
    nextHop.EntityData.Leafs["special-next-hop"] = types.YLeaf{"SpecialNextHop", nextHop.SpecialNextHop}
    return &(nextHop.EntityData)
}

// FibRoute_Output_Route_NextHop_SpecialNextHop represents Special next-hop options.
type FibRoute_Output_Route_NextHop_SpecialNextHop string

const (
    // Silently discard the packet.
    FibRoute_Output_Route_NextHop_SpecialNextHop_blackhole FibRoute_Output_Route_NextHop_SpecialNextHop = "blackhole"

    // Discard the packet and notify the sender with an error
    // message indicating that the destination host is
    // unreachable.
    FibRoute_Output_Route_NextHop_SpecialNextHop_unreachable FibRoute_Output_Route_NextHop_SpecialNextHop = "unreachable"

    // Discard the packet and notify the sender with an error
    // message indicating that the communication is
    // administratively prohibited.
    FibRoute_Output_Route_NextHop_SpecialNextHop_prohibit FibRoute_Output_Route_NextHop_SpecialNextHop = "prohibit"

    // The packet will be received by the local system.
    FibRoute_Output_Route_NextHop_SpecialNextHop_receive FibRoute_Output_Route_NextHop_SpecialNextHop = "receive"
)

