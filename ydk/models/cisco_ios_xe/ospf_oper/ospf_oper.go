// This module contains a collection of YANG definitions for
// monitoring the operation of ospf protocol in a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package ospf_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ospf_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-ospf-oper ospf-oper-data}", reflect.TypeOf(OspfOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-ospf-oper:ospf-oper-data", reflect.TypeOf(OspfOperData{}))
}

// AddressFamily represents Address family type
type AddressFamily string

const (
    AddressFamily_address_family_ipv4 AddressFamily = "address-family-ipv4"

    AddressFamily_address_family_ipv6 AddressFamily = "address-family-ipv6"
)

// OspfOperationMode represents OSPF operational mode
type OspfOperationMode string

const (
    // Ships-in-the-night operation mode in which each OSPF instance carries only one address family
    OspfOperationMode_ospf_ships_in_the_night OspfOperationMode = "ospf-ships-in-the-night"
)

// OspfNetworkType represents OSPF network type
type OspfNetworkType string

const (
    // OSPF broadcast multi-access network
    OspfNetworkType_ospf_broadcast OspfNetworkType = "ospf-broadcast"

    // OSPF Non-Broadcast Multi-Access (NBMA) network
    OspfNetworkType_ospf_non_broadcast OspfNetworkType = "ospf-non-broadcast"

    // OSPF point-to-multipoint network
    OspfNetworkType_ospf_point_to_multipoint OspfNetworkType = "ospf-point-to-multipoint"

    // OSPF point-to-point network
    OspfNetworkType_ospf_point_to_point OspfNetworkType = "ospf-point-to-point"
)

// OspfAuthType represents OSPF Authentication type
type OspfAuthType string

const (
    OspfAuthType_ospf_auth_ipsec OspfAuthType = "ospf-auth-ipsec"

    OspfAuthType_ospf_auth_trailer_keychain OspfAuthType = "ospf-auth-trailer-keychain"

    OspfAuthType_ospf_auth_trailer_key OspfAuthType = "ospf-auth-trailer-key"

    OspfAuthType_ospf_auth_type_none OspfAuthType = "ospf-auth-type-none"
)

// NbrStateType represents OSPF neighbor state type
type NbrStateType string

const (
    // Neighbor state down
    NbrStateType_ospf_nbr_down NbrStateType = "ospf-nbr-down"

    // Neighbor attempt state
    NbrStateType_ospf_nbr_attempt NbrStateType = "ospf-nbr-attempt"

    // Neighbor init state
    NbrStateType_ospf_nbr_init NbrStateType = "ospf-nbr-init"

    // Neighbor 2-way state
    NbrStateType_ospf_nbr_two_way NbrStateType = "ospf-nbr-two-way"

    // Neighbor exchange start state
    NbrStateType_ospf_nbr_exchange_start NbrStateType = "ospf-nbr-exchange-start"

    // Neighbor exchange state
    NbrStateType_ospf_nbr_exchange NbrStateType = "ospf-nbr-exchange"

    // Neighbor loading state
    NbrStateType_ospf_nbr_loading NbrStateType = "ospf-nbr-loading"

    // Neighbor full state
    NbrStateType_ospf_nbr_full NbrStateType = "ospf-nbr-full"
)

// OspfOperData
// Operational state of ospf
type OspfOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF operational state.
    OspfState OspfOperData_OspfState
}

func (ospfOperData *OspfOperData) GetEntityData() *types.CommonEntityData {
    ospfOperData.EntityData.YFilter = ospfOperData.YFilter
    ospfOperData.EntityData.YangName = "ospf-oper-data"
    ospfOperData.EntityData.BundleName = "cisco_ios_xe"
    ospfOperData.EntityData.ParentYangName = "Cisco-IOS-XE-ospf-oper"
    ospfOperData.EntityData.SegmentPath = "Cisco-IOS-XE-ospf-oper:ospf-oper-data"
    ospfOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfOperData.EntityData.Children = make(map[string]types.YChild)
    ospfOperData.EntityData.Children["ospf-state"] = types.YChild{"OspfState", &ospfOperData.OspfState}
    ospfOperData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfOperData.EntityData)
}

// OspfOperData_OspfState
// OSPF operational state
// This type is a presence type.
type OspfOperData_OspfState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF operation mode. The type is OspfOperationMode.
    OpMode interface{}

    // OSPF routing protocol instance. The type is slice of
    // OspfOperData_OspfState_OspfInstance.
    OspfInstance []OspfOperData_OspfState_OspfInstance
}

func (ospfState *OspfOperData_OspfState) GetEntityData() *types.CommonEntityData {
    ospfState.EntityData.YFilter = ospfState.YFilter
    ospfState.EntityData.YangName = "ospf-state"
    ospfState.EntityData.BundleName = "cisco_ios_xe"
    ospfState.EntityData.ParentYangName = "ospf-oper-data"
    ospfState.EntityData.SegmentPath = "ospf-state"
    ospfState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfState.EntityData.Children = make(map[string]types.YChild)
    ospfState.EntityData.Children["ospf-instance"] = types.YChild{"OspfInstance", nil}
    for i := range ospfState.OspfInstance {
        ospfState.EntityData.Children[types.GetSegmentPath(&ospfState.OspfInstance[i])] = types.YChild{"OspfInstance", &ospfState.OspfInstance[i]}
    }
    ospfState.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfState.EntityData.Leafs["op-mode"] = types.YLeaf{"OpMode", ospfState.OpMode}
    return &(ospfState.EntityData)
}

// OspfOperData_OspfState_OspfInstance
// OSPF routing protocol instance
type OspfOperData_OspfState_OspfInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address-family of the instance. The type is
    // AddressFamily.
    Af interface{}

    // This attribute is a key. Defined in RFC 2328. A 32-bit number that uniquely
    // identifies the router. The type is interface{} with range: 0..4294967295.
    RouterId interface{}

    // List of ospf areas. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea.
    OspfArea []OspfOperData_OspfState_OspfInstance_OspfArea

    // List OSPF link scope LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas.
    LinkScopeLsas []OspfOperData_OspfState_OspfInstance_LinkScopeLsas

    // OSPF multi-topology interface augmentation. The type is slice of
    // OspfOperData_OspfState_OspfInstance_MultiTopology.
    MultiTopology []OspfOperData_OspfState_OspfInstance_MultiTopology
}

func (ospfInstance *OspfOperData_OspfState_OspfInstance) GetEntityData() *types.CommonEntityData {
    ospfInstance.EntityData.YFilter = ospfInstance.YFilter
    ospfInstance.EntityData.YangName = "ospf-instance"
    ospfInstance.EntityData.BundleName = "cisco_ios_xe"
    ospfInstance.EntityData.ParentYangName = "ospf-state"
    ospfInstance.EntityData.SegmentPath = "ospf-instance" + "[af='" + fmt.Sprintf("%v", ospfInstance.Af) + "']" + "[router-id='" + fmt.Sprintf("%v", ospfInstance.RouterId) + "']"
    ospfInstance.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfInstance.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfInstance.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfInstance.EntityData.Children = make(map[string]types.YChild)
    ospfInstance.EntityData.Children["ospf-area"] = types.YChild{"OspfArea", nil}
    for i := range ospfInstance.OspfArea {
        ospfInstance.EntityData.Children[types.GetSegmentPath(&ospfInstance.OspfArea[i])] = types.YChild{"OspfArea", &ospfInstance.OspfArea[i]}
    }
    ospfInstance.EntityData.Children["link-scope-lsas"] = types.YChild{"LinkScopeLsas", nil}
    for i := range ospfInstance.LinkScopeLsas {
        ospfInstance.EntityData.Children[types.GetSegmentPath(&ospfInstance.LinkScopeLsas[i])] = types.YChild{"LinkScopeLsas", &ospfInstance.LinkScopeLsas[i]}
    }
    ospfInstance.EntityData.Children["multi-topology"] = types.YChild{"MultiTopology", nil}
    for i := range ospfInstance.MultiTopology {
        ospfInstance.EntityData.Children[types.GetSegmentPath(&ospfInstance.MultiTopology[i])] = types.YChild{"MultiTopology", &ospfInstance.MultiTopology[i]}
    }
    ospfInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfInstance.EntityData.Leafs["af"] = types.YLeaf{"Af", ospfInstance.Af}
    ospfInstance.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospfInstance.RouterId}
    return &(ospfInstance.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea
// List of ospf areas
type OspfOperData_OspfState_OspfInstance_OspfArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF area ID. The type is interface{} with range:
    // 0..4294967295.
    AreaId interface{}

    // List of OSPF interfaces. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface.
    OspfInterface []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface

    // List of OSPF area scope LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa.
    AreaScopeLsa []OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa
}

func (ospfArea *OspfOperData_OspfState_OspfInstance_OspfArea) GetEntityData() *types.CommonEntityData {
    ospfArea.EntityData.YFilter = ospfArea.YFilter
    ospfArea.EntityData.YangName = "ospf-area"
    ospfArea.EntityData.BundleName = "cisco_ios_xe"
    ospfArea.EntityData.ParentYangName = "ospf-instance"
    ospfArea.EntityData.SegmentPath = "ospf-area" + "[area-id='" + fmt.Sprintf("%v", ospfArea.AreaId) + "']"
    ospfArea.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfArea.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfArea.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfArea.EntityData.Children = make(map[string]types.YChild)
    ospfArea.EntityData.Children["ospf-interface"] = types.YChild{"OspfInterface", nil}
    for i := range ospfArea.OspfInterface {
        ospfArea.EntityData.Children[types.GetSegmentPath(&ospfArea.OspfInterface[i])] = types.YChild{"OspfInterface", &ospfArea.OspfInterface[i]}
    }
    ospfArea.EntityData.Children["area-scope-lsa"] = types.YChild{"AreaScopeLsa", nil}
    for i := range ospfArea.AreaScopeLsa {
        ospfArea.EntityData.Children[types.GetSegmentPath(&ospfArea.AreaScopeLsa[i])] = types.YChild{"AreaScopeLsa", &ospfArea.AreaScopeLsa[i]}
    }
    ospfArea.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfArea.EntityData.Leafs["area-id"] = types.YLeaf{"AreaId", ospfArea.AreaId}
    return &(ospfArea.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface
// List of OSPF interfaces
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string.
    Name interface{}

    // Network type. The type is OspfNetworkType.
    NetworkType interface{}

    // Enable/Disable passive. The type is bool.
    Passive interface{}

    // Enable/Disable demand circuit. The type is bool.
    DemandCircuit interface{}

    // Set prefix as a node representative prefix. The type is bool.
    NodeFlag interface{}

    // Interface cost. The type is interface{} with range: 0..65535.
    Cost interface{}

    // Time between hello packets. The type is interface{} with range: 0..65535.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead. The type is interface{}
    // with range: 0..65535.
    DeadInterval interface{}

    // Time between retransmitting unacknowledged Link State Advertisements
    // (LSAs). The type is interface{} with range: 0..65535.
    RetransmitInterval interface{}

    // Estimated time needed to send link-state update. The type is interface{}
    // with range: 0..65535.
    TransmitDelay interface{}

    // Enable/Disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/Disable link-local signaling (LLS) support. The type is bool.
    Lls interface{}

    // Suppress advertisement of the prefixes. The type is bool.
    PrefixSuppression interface{}

    // Enable/disable bfd. The type is bool.
    Bfd interface{}

    // Enable/disable protocol on the interface. The type is bool.
    Enable interface{}

    // Interface state. The type is string.
    State interface{}

    // Hello timer. The type is interface{} with range: 0..4294967295.
    HelloTimer interface{}

    // Wait timer. The type is interface{} with range: 0..4294967295.
    WaitTimer interface{}

    // Designated Router. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Dr interface{}

    // Backup Designated Router. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Bdr interface{}

    // Configure OSPF router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Multi Area.
    MultiArea OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_MultiArea

    // Staticly configured neighbors. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_StaticNeighbor.
    StaticNeighbor []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_StaticNeighbor

    // Fast reroute config.
    FastReroute OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_FastReroute

    // TTL security.
    TtlSecurity OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_TtlSecurity

    // Authentication configuration.
    Authentication OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication

    // List of OSPF neighbors. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor.
    OspfNeighbor []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor

    // List OSPF link scope LSAs. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas.
    IntfLinkScopeLsas []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas

    // OSPF interface topology. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfMultiTopology.
    IntfMultiTopology []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfMultiTopology
}

func (ospfInterface *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface) GetEntityData() *types.CommonEntityData {
    ospfInterface.EntityData.YFilter = ospfInterface.YFilter
    ospfInterface.EntityData.YangName = "ospf-interface"
    ospfInterface.EntityData.BundleName = "cisco_ios_xe"
    ospfInterface.EntityData.ParentYangName = "ospf-area"
    ospfInterface.EntityData.SegmentPath = "ospf-interface" + "[name='" + fmt.Sprintf("%v", ospfInterface.Name) + "']"
    ospfInterface.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfInterface.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfInterface.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfInterface.EntityData.Children = make(map[string]types.YChild)
    ospfInterface.EntityData.Children["multi-area"] = types.YChild{"MultiArea", &ospfInterface.MultiArea}
    ospfInterface.EntityData.Children["static-neighbor"] = types.YChild{"StaticNeighbor", nil}
    for i := range ospfInterface.StaticNeighbor {
        ospfInterface.EntityData.Children[types.GetSegmentPath(&ospfInterface.StaticNeighbor[i])] = types.YChild{"StaticNeighbor", &ospfInterface.StaticNeighbor[i]}
    }
    ospfInterface.EntityData.Children["fast-reroute"] = types.YChild{"FastReroute", &ospfInterface.FastReroute}
    ospfInterface.EntityData.Children["ttl-security"] = types.YChild{"TtlSecurity", &ospfInterface.TtlSecurity}
    ospfInterface.EntityData.Children["authentication"] = types.YChild{"Authentication", &ospfInterface.Authentication}
    ospfInterface.EntityData.Children["ospf-neighbor"] = types.YChild{"OspfNeighbor", nil}
    for i := range ospfInterface.OspfNeighbor {
        ospfInterface.EntityData.Children[types.GetSegmentPath(&ospfInterface.OspfNeighbor[i])] = types.YChild{"OspfNeighbor", &ospfInterface.OspfNeighbor[i]}
    }
    ospfInterface.EntityData.Children["intf-link-scope-lsas"] = types.YChild{"IntfLinkScopeLsas", nil}
    for i := range ospfInterface.IntfLinkScopeLsas {
        ospfInterface.EntityData.Children[types.GetSegmentPath(&ospfInterface.IntfLinkScopeLsas[i])] = types.YChild{"IntfLinkScopeLsas", &ospfInterface.IntfLinkScopeLsas[i]}
    }
    ospfInterface.EntityData.Children["intf-multi-topology"] = types.YChild{"IntfMultiTopology", nil}
    for i := range ospfInterface.IntfMultiTopology {
        ospfInterface.EntityData.Children[types.GetSegmentPath(&ospfInterface.IntfMultiTopology[i])] = types.YChild{"IntfMultiTopology", &ospfInterface.IntfMultiTopology[i]}
    }
    ospfInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfInterface.EntityData.Leafs["name"] = types.YLeaf{"Name", ospfInterface.Name}
    ospfInterface.EntityData.Leafs["network-type"] = types.YLeaf{"NetworkType", ospfInterface.NetworkType}
    ospfInterface.EntityData.Leafs["passive"] = types.YLeaf{"Passive", ospfInterface.Passive}
    ospfInterface.EntityData.Leafs["demand-circuit"] = types.YLeaf{"DemandCircuit", ospfInterface.DemandCircuit}
    ospfInterface.EntityData.Leafs["node-flag"] = types.YLeaf{"NodeFlag", ospfInterface.NodeFlag}
    ospfInterface.EntityData.Leafs["cost"] = types.YLeaf{"Cost", ospfInterface.Cost}
    ospfInterface.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", ospfInterface.HelloInterval}
    ospfInterface.EntityData.Leafs["dead-interval"] = types.YLeaf{"DeadInterval", ospfInterface.DeadInterval}
    ospfInterface.EntityData.Leafs["retransmit-interval"] = types.YLeaf{"RetransmitInterval", ospfInterface.RetransmitInterval}
    ospfInterface.EntityData.Leafs["transmit-delay"] = types.YLeaf{"TransmitDelay", ospfInterface.TransmitDelay}
    ospfInterface.EntityData.Leafs["mtu-ignore"] = types.YLeaf{"MtuIgnore", ospfInterface.MtuIgnore}
    ospfInterface.EntityData.Leafs["lls"] = types.YLeaf{"Lls", ospfInterface.Lls}
    ospfInterface.EntityData.Leafs["prefix-suppression"] = types.YLeaf{"PrefixSuppression", ospfInterface.PrefixSuppression}
    ospfInterface.EntityData.Leafs["bfd"] = types.YLeaf{"Bfd", ospfInterface.Bfd}
    ospfInterface.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ospfInterface.Enable}
    ospfInterface.EntityData.Leafs["state"] = types.YLeaf{"State", ospfInterface.State}
    ospfInterface.EntityData.Leafs["hello-timer"] = types.YLeaf{"HelloTimer", ospfInterface.HelloTimer}
    ospfInterface.EntityData.Leafs["wait-timer"] = types.YLeaf{"WaitTimer", ospfInterface.WaitTimer}
    ospfInterface.EntityData.Leafs["dr"] = types.YLeaf{"Dr", ospfInterface.Dr}
    ospfInterface.EntityData.Leafs["bdr"] = types.YLeaf{"Bdr", ospfInterface.Bdr}
    ospfInterface.EntityData.Leafs["priority"] = types.YLeaf{"Priority", ospfInterface.Priority}
    return &(ospfInterface.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_MultiArea
// Multi Area
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_MultiArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multi-area ID. The type is interface{} with range: 0..4294967295.
    MultiAreaId interface{}

    // Interface cost for multi-area. The type is interface{} with range:
    // 0..65535.
    Cost interface{}
}

func (multiArea *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_MultiArea) GetEntityData() *types.CommonEntityData {
    multiArea.EntityData.YFilter = multiArea.YFilter
    multiArea.EntityData.YangName = "multi-area"
    multiArea.EntityData.BundleName = "cisco_ios_xe"
    multiArea.EntityData.ParentYangName = "ospf-interface"
    multiArea.EntityData.SegmentPath = "multi-area"
    multiArea.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    multiArea.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    multiArea.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    multiArea.EntityData.Children = make(map[string]types.YChild)
    multiArea.EntityData.Leafs = make(map[string]types.YLeaf)
    multiArea.EntityData.Leafs["multi-area-id"] = types.YLeaf{"MultiAreaId", multiArea.MultiAreaId}
    multiArea.EntityData.Leafs["cost"] = types.YLeaf{"Cost", multiArea.Cost}
    return &(multiArea.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_StaticNeighbor
// Staticly configured neighbors
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_StaticNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Neighbor cost. The type is interface{} with range: 0..65535.
    Cost interface{}

    // Neighbor polling intervali in seconds. The type is interface{} with range:
    // 0..65535. Units are seconds.
    PollInterval interface{}
}

func (staticNeighbor *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_StaticNeighbor) GetEntityData() *types.CommonEntityData {
    staticNeighbor.EntityData.YFilter = staticNeighbor.YFilter
    staticNeighbor.EntityData.YangName = "static-neighbor"
    staticNeighbor.EntityData.BundleName = "cisco_ios_xe"
    staticNeighbor.EntityData.ParentYangName = "ospf-interface"
    staticNeighbor.EntityData.SegmentPath = "static-neighbor" + "[address='" + fmt.Sprintf("%v", staticNeighbor.Address) + "']"
    staticNeighbor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    staticNeighbor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    staticNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    staticNeighbor.EntityData.Children = make(map[string]types.YChild)
    staticNeighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    staticNeighbor.EntityData.Leafs["address"] = types.YLeaf{"Address", staticNeighbor.Address}
    staticNeighbor.EntityData.Leafs["cost"] = types.YLeaf{"Cost", staticNeighbor.Cost}
    staticNeighbor.EntityData.Leafs["poll-interval"] = types.YLeaf{"PollInterval", staticNeighbor.PollInterval}
    return &(staticNeighbor.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_FastReroute
// Fast reroute config
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prevent the interface to be used as backup. The type is bool.
    CandidateDisabled interface{}

    // Activates LFA. This model assumes activation of per-prefix LFA. The type is
    // bool.
    Enabled interface{}

    // Activates remote LFA. The type is bool.
    RemoteLfaEnabled interface{}
}

func (fastReroute *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xe"
    fastReroute.EntityData.ParentYangName = "ospf-interface"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fastReroute.EntityData.Children = make(map[string]types.YChild)
    fastReroute.EntityData.Leafs = make(map[string]types.YLeaf)
    fastReroute.EntityData.Leafs["candidate-disabled"] = types.YLeaf{"CandidateDisabled", fastReroute.CandidateDisabled}
    fastReroute.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", fastReroute.Enabled}
    fastReroute.EntityData.Leafs["remote-lfa-enabled"] = types.YLeaf{"RemoteLfaEnabled", fastReroute.RemoteLfaEnabled}
    return &(fastReroute.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_TtlSecurity
// TTL security
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_TtlSecurity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable TTL security check. The type is bool.
    Enabled interface{}

    // Maximum number of hops that a OSPF packet may have traveled. The type is
    // interface{} with range: 0..255.
    Hops interface{}
}

func (ttlSecurity *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_TtlSecurity) GetEntityData() *types.CommonEntityData {
    ttlSecurity.EntityData.YFilter = ttlSecurity.YFilter
    ttlSecurity.EntityData.YangName = "ttl-security"
    ttlSecurity.EntityData.BundleName = "cisco_ios_xe"
    ttlSecurity.EntityData.ParentYangName = "ospf-interface"
    ttlSecurity.EntityData.SegmentPath = "ttl-security"
    ttlSecurity.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ttlSecurity.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ttlSecurity.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ttlSecurity.EntityData.Children = make(map[string]types.YChild)
    ttlSecurity.EntityData.Leafs = make(map[string]types.YLeaf)
    ttlSecurity.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", ttlSecurity.Enabled}
    ttlSecurity.EntityData.Leafs["hops"] = types.YLeaf{"Hops", ttlSecurity.Hops}
    return &(ttlSecurity.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication
// Authentication configuration
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA name. The type is string.
    Sa interface{}

    // key-chain name. The type is string.
    KeyChain interface{}

    // Key string in ASCII format. The type is string.
    KeyString interface{}

    // No authentication enabled. The type is interface{} with range:
    // 0..4294967295.
    NoAuth interface{}

    // Crypto algorithm.
    CryptoAlgorithmVal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication_CryptoAlgorithmVal
}

func (authentication *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xe"
    authentication.EntityData.ParentYangName = "ospf-interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["crypto-algorithm-val"] = types.YChild{"CryptoAlgorithmVal", &authentication.CryptoAlgorithmVal}
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["sa"] = types.YLeaf{"Sa", authentication.Sa}
    authentication.EntityData.Leafs["key-chain"] = types.YLeaf{"KeyChain", authentication.KeyChain}
    authentication.EntityData.Leafs["key-string"] = types.YLeaf{"KeyString", authentication.KeyString}
    authentication.EntityData.Leafs["no-auth"] = types.YLeaf{"NoAuth", authentication.NoAuth}
    return &(authentication.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication_CryptoAlgorithmVal
// Crypto algorithm
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication_CryptoAlgorithmVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // SHA-1 algorithm. The type is interface{}.
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

func (cryptoAlgorithmVal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_Authentication_CryptoAlgorithmVal) GetEntityData() *types.CommonEntityData {
    cryptoAlgorithmVal.EntityData.YFilter = cryptoAlgorithmVal.YFilter
    cryptoAlgorithmVal.EntityData.YangName = "crypto-algorithm-val"
    cryptoAlgorithmVal.EntityData.BundleName = "cisco_ios_xe"
    cryptoAlgorithmVal.EntityData.ParentYangName = "authentication"
    cryptoAlgorithmVal.EntityData.SegmentPath = "crypto-algorithm-val"
    cryptoAlgorithmVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cryptoAlgorithmVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cryptoAlgorithmVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cryptoAlgorithmVal.EntityData.Children = make(map[string]types.YChild)
    cryptoAlgorithmVal.EntityData.Leafs = make(map[string]types.YLeaf)
    cryptoAlgorithmVal.EntityData.Leafs["hmac-sha1-12"] = types.YLeaf{"HmacSha112", cryptoAlgorithmVal.HmacSha112}
    cryptoAlgorithmVal.EntityData.Leafs["hmac-sha1-20"] = types.YLeaf{"HmacSha120", cryptoAlgorithmVal.HmacSha120}
    cryptoAlgorithmVal.EntityData.Leafs["md5"] = types.YLeaf{"Md5", cryptoAlgorithmVal.Md5}
    cryptoAlgorithmVal.EntityData.Leafs["sha-1"] = types.YLeaf{"Sha1", cryptoAlgorithmVal.Sha1}
    cryptoAlgorithmVal.EntityData.Leafs["hmac-sha-1"] = types.YLeaf{"HmacSha1", cryptoAlgorithmVal.HmacSha1}
    cryptoAlgorithmVal.EntityData.Leafs["hmac-sha-256"] = types.YLeaf{"HmacSha256", cryptoAlgorithmVal.HmacSha256}
    cryptoAlgorithmVal.EntityData.Leafs["hmac-sha-384"] = types.YLeaf{"HmacSha384", cryptoAlgorithmVal.HmacSha384}
    cryptoAlgorithmVal.EntityData.Leafs["hmac-sha-512"] = types.YLeaf{"HmacSha512", cryptoAlgorithmVal.HmacSha512}
    return &(cryptoAlgorithmVal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor
// List of OSPF neighbors
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF neighbor ID. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NeighborId interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Designated Router. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Dr interface{}

    // Backup Designated Router. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Bdr interface{}

    // OSPF neighbor state. The type is NbrStateType.
    State interface{}

    // Per-neighbor statistics.
    Stats OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor_Stats
}

func (ospfNeighbor *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor) GetEntityData() *types.CommonEntityData {
    ospfNeighbor.EntityData.YFilter = ospfNeighbor.YFilter
    ospfNeighbor.EntityData.YangName = "ospf-neighbor"
    ospfNeighbor.EntityData.BundleName = "cisco_ios_xe"
    ospfNeighbor.EntityData.ParentYangName = "ospf-interface"
    ospfNeighbor.EntityData.SegmentPath = "ospf-neighbor" + "[neighbor-id='" + fmt.Sprintf("%v", ospfNeighbor.NeighborId) + "']"
    ospfNeighbor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfNeighbor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfNeighbor.EntityData.Children = make(map[string]types.YChild)
    ospfNeighbor.EntityData.Children["stats"] = types.YChild{"Stats", &ospfNeighbor.Stats}
    ospfNeighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfNeighbor.EntityData.Leafs["neighbor-id"] = types.YLeaf{"NeighborId", ospfNeighbor.NeighborId}
    ospfNeighbor.EntityData.Leafs["address"] = types.YLeaf{"Address", ospfNeighbor.Address}
    ospfNeighbor.EntityData.Leafs["dr"] = types.YLeaf{"Dr", ospfNeighbor.Dr}
    ospfNeighbor.EntityData.Leafs["bdr"] = types.YLeaf{"Bdr", ospfNeighbor.Bdr}
    ospfNeighbor.EntityData.Leafs["state"] = types.YLeaf{"State", ospfNeighbor.State}
    return &(ospfNeighbor.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor_Stats
// Per-neighbor statistics
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of time this neighbor has changed state or an error has
    // occurred. The type is interface{} with range: 0..4294967295.
    NbrEventCount interface{}

    // The current length of the retransmission queue. The type is interface{}
    // with range: 0..4294967295.
    NbrRetransQlen interface{}
}

func (stats *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_OspfNeighbor_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xe"
    stats.EntityData.ParentYangName = "ospf-neighbor"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["nbr-event-count"] = types.YLeaf{"NbrEventCount", stats.NbrEventCount}
    stats.EntityData.Leafs["nbr-retrans-qlen"] = types.YLeaf{"NbrRetransQlen", stats.NbrRetransQlen}
    return &(stats.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas
// List OSPF link scope LSAs
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF link scope LSA type. The type is interface{}
    // with range: 0..4294967295.
    LsaType interface{}

    // List of OSPF link scope LSAs. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa.
    LinkScopeLsa []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa

    // List OSPF area scope LSA databases. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa.
    AreaScopeLsa []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa
}

func (intfLinkScopeLsas *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas) GetEntityData() *types.CommonEntityData {
    intfLinkScopeLsas.EntityData.YFilter = intfLinkScopeLsas.YFilter
    intfLinkScopeLsas.EntityData.YangName = "intf-link-scope-lsas"
    intfLinkScopeLsas.EntityData.BundleName = "cisco_ios_xe"
    intfLinkScopeLsas.EntityData.ParentYangName = "ospf-interface"
    intfLinkScopeLsas.EntityData.SegmentPath = "intf-link-scope-lsas" + "[lsa-type='" + fmt.Sprintf("%v", intfLinkScopeLsas.LsaType) + "']"
    intfLinkScopeLsas.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intfLinkScopeLsas.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intfLinkScopeLsas.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intfLinkScopeLsas.EntityData.Children = make(map[string]types.YChild)
    intfLinkScopeLsas.EntityData.Children["link-scope-lsa"] = types.YChild{"LinkScopeLsa", nil}
    for i := range intfLinkScopeLsas.LinkScopeLsa {
        intfLinkScopeLsas.EntityData.Children[types.GetSegmentPath(&intfLinkScopeLsas.LinkScopeLsa[i])] = types.YChild{"LinkScopeLsa", &intfLinkScopeLsas.LinkScopeLsa[i]}
    }
    intfLinkScopeLsas.EntityData.Children["area-scope-lsa"] = types.YChild{"AreaScopeLsa", nil}
    for i := range intfLinkScopeLsas.AreaScopeLsa {
        intfLinkScopeLsas.EntityData.Children[types.GetSegmentPath(&intfLinkScopeLsas.AreaScopeLsa[i])] = types.YChild{"AreaScopeLsa", &intfLinkScopeLsas.AreaScopeLsa[i]}
    }
    intfLinkScopeLsas.EntityData.Leafs = make(map[string]types.YLeaf)
    intfLinkScopeLsas.EntityData.Leafs["lsa-type"] = types.YLeaf{"LsaType", intfLinkScopeLsas.LsaType}
    return &(intfLinkScopeLsas.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa
// List of OSPF link scope LSAs
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSA ID. The type is interface{} with range:
    // 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is slice of interface{} with range: 0..255.
    RawData []interface{}

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Router address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RouterAddress interface{}

    // OSPFv2 LSA.
    Ospfv2Lsa OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa

    // OSPFv2 LSA link. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link.
    Ospfv2Link []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link

    // Summary LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Topology.
    Ospfv2Topology []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Topology

    // External LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2External.
    Ospfv2External []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2External

    // OSPFv2 Unknown TLV. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv.
    Ospfv2UnknownTlv []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv

    // OSPFv3 LSA.
    Ospfv3LsaVal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal

    // OSPFv3 links. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3Link.
    Ospfv3Link []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3Link

    // OSPFv3 prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList.
    Ospfv3PrefixList []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList

    // OSPFv3 intra-area prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix.
    Ospfv3IaPrefix []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix

    // OSPF multi-topology interface augmentation. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_MultiTopology.
    MultiTopology []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_MultiTopology

    // Link TLV.
    Tlv OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Tlv

    // OSPFv2 Unknown sub TLV. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_UnknownSubTlv.
    UnknownSubTlv []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_UnknownSubTlv
}

func (linkScopeLsa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa) GetEntityData() *types.CommonEntityData {
    linkScopeLsa.EntityData.YFilter = linkScopeLsa.YFilter
    linkScopeLsa.EntityData.YangName = "link-scope-lsa"
    linkScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    linkScopeLsa.EntityData.ParentYangName = "intf-link-scope-lsas"
    linkScopeLsa.EntityData.SegmentPath = "link-scope-lsa" + "[lsa-id='" + fmt.Sprintf("%v", linkScopeLsa.LsaId) + "']" + "[adv-router='" + fmt.Sprintf("%v", linkScopeLsa.AdvRouter) + "']"
    linkScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkScopeLsa.EntityData.Children = make(map[string]types.YChild)
    linkScopeLsa.EntityData.Children["ospfv2-lsa"] = types.YChild{"Ospfv2Lsa", &linkScopeLsa.Ospfv2Lsa}
    linkScopeLsa.EntityData.Children["ospfv2-link"] = types.YChild{"Ospfv2Link", nil}
    for i := range linkScopeLsa.Ospfv2Link {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv2Link[i])] = types.YChild{"Ospfv2Link", &linkScopeLsa.Ospfv2Link[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv2-topology"] = types.YChild{"Ospfv2Topology", nil}
    for i := range linkScopeLsa.Ospfv2Topology {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv2Topology[i])] = types.YChild{"Ospfv2Topology", &linkScopeLsa.Ospfv2Topology[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv2-external"] = types.YChild{"Ospfv2External", nil}
    for i := range linkScopeLsa.Ospfv2External {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv2External[i])] = types.YChild{"Ospfv2External", &linkScopeLsa.Ospfv2External[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv2-unknown-tlv"] = types.YChild{"Ospfv2UnknownTlv", nil}
    for i := range linkScopeLsa.Ospfv2UnknownTlv {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv2UnknownTlv[i])] = types.YChild{"Ospfv2UnknownTlv", &linkScopeLsa.Ospfv2UnknownTlv[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv3-lsa-val"] = types.YChild{"Ospfv3LsaVal", &linkScopeLsa.Ospfv3LsaVal}
    linkScopeLsa.EntityData.Children["ospfv3-link"] = types.YChild{"Ospfv3Link", nil}
    for i := range linkScopeLsa.Ospfv3Link {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv3Link[i])] = types.YChild{"Ospfv3Link", &linkScopeLsa.Ospfv3Link[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv3-prefix-list"] = types.YChild{"Ospfv3PrefixList", nil}
    for i := range linkScopeLsa.Ospfv3PrefixList {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv3PrefixList[i])] = types.YChild{"Ospfv3PrefixList", &linkScopeLsa.Ospfv3PrefixList[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv3-ia-prefix"] = types.YChild{"Ospfv3IaPrefix", nil}
    for i := range linkScopeLsa.Ospfv3IaPrefix {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv3IaPrefix[i])] = types.YChild{"Ospfv3IaPrefix", &linkScopeLsa.Ospfv3IaPrefix[i]}
    }
    linkScopeLsa.EntityData.Children["multi-topology"] = types.YChild{"MultiTopology", nil}
    for i := range linkScopeLsa.MultiTopology {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.MultiTopology[i])] = types.YChild{"MultiTopology", &linkScopeLsa.MultiTopology[i]}
    }
    linkScopeLsa.EntityData.Children["tlv"] = types.YChild{"Tlv", &linkScopeLsa.Tlv}
    linkScopeLsa.EntityData.Children["unknown-sub-tlv"] = types.YChild{"UnknownSubTlv", nil}
    for i := range linkScopeLsa.UnknownSubTlv {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.UnknownSubTlv[i])] = types.YChild{"UnknownSubTlv", &linkScopeLsa.UnknownSubTlv[i]}
    }
    linkScopeLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    linkScopeLsa.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", linkScopeLsa.LsaId}
    linkScopeLsa.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", linkScopeLsa.AdvRouter}
    linkScopeLsa.EntityData.Leafs["decoded-completed"] = types.YLeaf{"DecodedCompleted", linkScopeLsa.DecodedCompleted}
    linkScopeLsa.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", linkScopeLsa.RawData}
    linkScopeLsa.EntityData.Leafs["version"] = types.YLeaf{"Version", linkScopeLsa.Version}
    linkScopeLsa.EntityData.Leafs["router-address"] = types.YLeaf{"RouterAddress", linkScopeLsa.RouterAddress}
    return &(linkScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa
// OSPFv2 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header

    // Decoded OSPFv2 LSA body data.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody
}

func (ospfv2Lsa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa) GetEntityData() *types.CommonEntityData {
    ospfv2Lsa.EntityData.YFilter = ospfv2Lsa.YFilter
    ospfv2Lsa.EntityData.YangName = "ospfv2-lsa"
    ospfv2Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Lsa.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Lsa.EntityData.SegmentPath = "ospfv2-lsa"
    ospfv2Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Lsa.EntityData.Children = make(map[string]types.YChild)
    ospfv2Lsa.EntityData.Children["header"] = types.YChild{"Header", &ospfv2Lsa.Header}
    ospfv2Lsa.EntityData.Children["lsa-body"] = types.YChild{"LsaBody", &ospfv2Lsa.LsaBody}
    ospfv2Lsa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header
// Decoded OSPFv2 LSA header data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255.
    OpaqueType interface{}

    // Opaque ID. The type is interface{} with range: 0..4294967295.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type_ interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}

    // LSA options. The type is map[string]bool.
    FlagOptions interface{}
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv2-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["opaque-type"] = types.YLeaf{"OpaqueType", header.OpaqueType}
    header.EntityData.Leafs["opaque-id"] = types.YLeaf{"OpaqueId", header.OpaqueId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    header.EntityData.Leafs["flag-options"] = types.YLeaf{"FlagOptions", header.FlagOptions}
    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody
// Decoded OSPFv2 LSA body data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Summary mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryMask interface{}

    // External mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ExternalMask interface{}

    // LSA body flags. The type is map[string]bool.
    BodyFlagOptions interface{}

    // Network details.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv2-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = make(map[string]types.YChild)
    lsaBody.EntityData.Children["network"] = types.YChild{"Network", &lsaBody.Network}
    lsaBody.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaBody.EntityData.Leafs["num-of-links"] = types.YLeaf{"NumOfLinks", lsaBody.NumOfLinks}
    lsaBody.EntityData.Leafs["summary-mask"] = types.YLeaf{"SummaryMask", lsaBody.SummaryMask}
    lsaBody.EntityData.Leafs["external-mask"] = types.YLeaf{"ExternalMask", lsaBody.ExternalMask}
    lsaBody.EntityData.Leafs["body-flag-options"] = types.YLeaf{"BodyFlagOptions", lsaBody.BodyFlagOptions}
    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network
// Network details
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP network mask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", network.NetworkMask}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link
// OSPFv2 LSA link
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is interface{} with range:
    // 0..4294967295.
    LinkId interface{}

    // This attribute is a key. Link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Topology specific information. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology.
    Ospfv2Topology []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology
}

func (ospfv2Link *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link) GetEntityData() *types.CommonEntityData {
    ospfv2Link.EntityData.YFilter = ospfv2Link.YFilter
    ospfv2Link.EntityData.YangName = "ospfv2-link"
    ospfv2Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Link.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Link.EntityData.SegmentPath = "ospfv2-link" + "[link-id='" + fmt.Sprintf("%v", ospfv2Link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", ospfv2Link.LinkData) + "']"
    ospfv2Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Link.EntityData.Children = make(map[string]types.YChild)
    ospfv2Link.EntityData.Children["ospfv2-topology"] = types.YChild{"Ospfv2Topology", nil}
    for i := range ospfv2Link.Ospfv2Topology {
        ospfv2Link.EntityData.Children[types.GetSegmentPath(&ospfv2Link.Ospfv2Topology[i])] = types.YChild{"Ospfv2Topology", &ospfv2Link.Ospfv2Topology[i]}
    }
    ospfv2Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Link.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", ospfv2Link.LinkId}
    ospfv2Link.EntityData.Leafs["link-data"] = types.YLeaf{"LinkData", ospfv2Link.LinkData}
    ospfv2Link.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv2Link.Type_}
    return &(ospfv2Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology
// Topology specific information
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "ospfv2-link"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + "[mt-id='" + fmt.Sprintf("%v", ospfv2Topology.MtId) + "']"
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = make(map[string]types.YChild)
    ospfv2Topology.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2Topology.MtId}
    ospfv2Topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2Topology.Metric}
    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Topology
// Summary LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + "[mt-id='" + fmt.Sprintf("%v", ospfv2Topology.MtId) + "']"
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = make(map[string]types.YChild)
    ospfv2Topology.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2Topology.MtId}
    ospfv2Topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2Topology.Metric}
    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2External
// External LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled on the link. The type
    // is interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (ospfv2External *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2External) GetEntityData() *types.CommonEntityData {
    ospfv2External.EntityData.YFilter = ospfv2External.YFilter
    ospfv2External.EntityData.YangName = "ospfv2-external"
    ospfv2External.EntityData.BundleName = "cisco_ios_xe"
    ospfv2External.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2External.EntityData.SegmentPath = "ospfv2-external" + "[mt-id='" + fmt.Sprintf("%v", ospfv2External.MtId) + "']"
    ospfv2External.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2External.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2External.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2External.EntityData.Children = make(map[string]types.YChild)
    ospfv2External.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2External.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2External.MtId}
    ospfv2External.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2External.Metric}
    ospfv2External.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", ospfv2External.ForwardingAddress}
    ospfv2External.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", ospfv2External.ExternalRouteTag}
    return &(ospfv2External.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv
// OSPFv2 Unknown TLV
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type_ interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (ospfv2UnknownTlv *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv) GetEntityData() *types.CommonEntityData {
    ospfv2UnknownTlv.EntityData.YFilter = ospfv2UnknownTlv.YFilter
    ospfv2UnknownTlv.EntityData.YangName = "ospfv2-unknown-tlv"
    ospfv2UnknownTlv.EntityData.BundleName = "cisco_ios_xe"
    ospfv2UnknownTlv.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2UnknownTlv.EntityData.SegmentPath = "ospfv2-unknown-tlv" + "[type='" + fmt.Sprintf("%v", ospfv2UnknownTlv.Type_) + "']"
    ospfv2UnknownTlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2UnknownTlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2UnknownTlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2UnknownTlv.EntityData.Children = make(map[string]types.YChild)
    ospfv2UnknownTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2UnknownTlv.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv2UnknownTlv.Type_}
    ospfv2UnknownTlv.EntityData.Leafs["length"] = types.YLeaf{"Length", ospfv2UnknownTlv.Length}
    ospfv2UnknownTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", ospfv2UnknownTlv.Value}
    return &(ospfv2UnknownTlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal
// OSPFv3 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header

    // Decoded OSPFv3 LSA body.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody
}

func (ospfv3LsaVal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal) GetEntityData() *types.CommonEntityData {
    ospfv3LsaVal.EntityData.YFilter = ospfv3LsaVal.YFilter
    ospfv3LsaVal.EntityData.YangName = "ospfv3-lsa-val"
    ospfv3LsaVal.EntityData.BundleName = "cisco_ios_xe"
    ospfv3LsaVal.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3LsaVal.EntityData.SegmentPath = "ospfv3-lsa-val"
    ospfv3LsaVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3LsaVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3LsaVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3LsaVal.EntityData.Children = make(map[string]types.YChild)
    ospfv3LsaVal.EntityData.Children["header"] = types.YChild{"Header", &ospfv3LsaVal.Header}
    ospfv3LsaVal.EntityData.Children["lsa-body"] = types.YChild{"LsaBody", &ospfv3LsaVal.LsaBody}
    ospfv3LsaVal.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3LsaVal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header
// Decoded OSPFv3 LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // OSPFv3 LSA options. The type is map[string]bool.
    LsaHdrOptions interface{}

    // LSA header.
    LsaHeader OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv3-lsa-val"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Children["lsa-header"] = types.YChild{"LsaHeader", &header.LsaHeader}
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["lsa-hdr-options"] = types.YLeaf{"LsaHdrOptions", header.LsaHdrOptions}
    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader
// LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type_ interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}
}

func (lsaHeader *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader) GetEntityData() *types.CommonEntityData {
    lsaHeader.EntityData.YFilter = lsaHeader.YFilter
    lsaHeader.EntityData.YangName = "lsa-header"
    lsaHeader.EntityData.BundleName = "cisco_ios_xe"
    lsaHeader.EntityData.ParentYangName = "header"
    lsaHeader.EntityData.SegmentPath = "lsa-header"
    lsaHeader.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaHeader.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaHeader.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaHeader.EntityData.Children = make(map[string]types.YChild)
    lsaHeader.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaHeader.EntityData.Leafs["age"] = types.YLeaf{"Age", lsaHeader.Age}
    lsaHeader.EntityData.Leafs["type"] = types.YLeaf{"Type_", lsaHeader.Type_}
    lsaHeader.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", lsaHeader.AdvRouter}
    lsaHeader.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", lsaHeader.SeqNum}
    lsaHeader.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", lsaHeader.Checksum}
    lsaHeader.EntityData.Leafs["length"] = types.YLeaf{"Length", lsaHeader.Length}
    return &(lsaHeader.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody
// Decoded OSPFv3 LSA body
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA options. The type is map[string]bool.
    LsaFlagOptions interface{}

    // LSA Body Flags. The type is map[string]bool.
    LsaBodyFlags interface{}

    // OSPFv3 network.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network

    // OSPFv3 inter area prefix.
    Prefix OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix

    // OSPFv3 inter area router.
    IaRouter OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter

    // OSPFv3 LSA external.
    LsaExternal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal

    // OSPFv3 NSSA.
    Nssa OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa

    // OSPFv3 Link data.
    LinkData OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData

    // OSPFv3 Intra area prefixes.
    IaPrefix OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv3-lsa-val"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = make(map[string]types.YChild)
    lsaBody.EntityData.Children["network"] = types.YChild{"Network", &lsaBody.Network}
    lsaBody.EntityData.Children["prefix"] = types.YChild{"Prefix", &lsaBody.Prefix}
    lsaBody.EntityData.Children["ia-router"] = types.YChild{"IaRouter", &lsaBody.IaRouter}
    lsaBody.EntityData.Children["lsa-external"] = types.YChild{"LsaExternal", &lsaBody.LsaExternal}
    lsaBody.EntityData.Children["nssa"] = types.YChild{"Nssa", &lsaBody.Nssa}
    lsaBody.EntityData.Children["link-data"] = types.YChild{"LinkData", &lsaBody.LinkData}
    lsaBody.EntityData.Children["ia-prefix"] = types.YChild{"IaPrefix", &lsaBody.IaPrefix}
    lsaBody.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaBody.EntityData.Leafs["lsa-flag-options"] = types.YLeaf{"LsaFlagOptions", lsaBody.LsaFlagOptions}
    lsaBody.EntityData.Leafs["lsa-body-flags"] = types.YLeaf{"LsaBodyFlags", lsaBody.LsaBodyFlags}
    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network
// OSPFv3 network
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}

    // Network LSA options. The type is map[string]bool.
    LsaNetOptions interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    network.EntityData.Leafs["lsa-net-options"] = types.YLeaf{"LsaNetOptions", network.LsaNetOptions}
    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix
// OSPFv3 inter area prefix
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Inter area Prefix. The type is string.
    IaPrefix interface{}

    // Inter area prefix options. The type is string.
    IaPrefixOptions interface{}
}

func (prefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xe"
    prefix.EntityData.ParentYangName = "lsa-body"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["metric"] = types.YLeaf{"Metric", prefix.Metric}
    prefix.EntityData.Leafs["ia-prefix"] = types.YLeaf{"IaPrefix", prefix.IaPrefix}
    prefix.EntityData.Leafs["ia-prefix-options"] = types.YLeaf{"IaPrefixOptions", prefix.IaPrefixOptions}
    return &(prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter
// OSPFv3 inter area router
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Router ID of the router being described by the LSA. The type is interface{}
    // with range: 0..4294967295.
    DestinationRouterId interface{}

    // Inter area LSA options. The type is map[string]bool.
    LsaIaOptions interface{}
}

func (iaRouter *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter) GetEntityData() *types.CommonEntityData {
    iaRouter.EntityData.YFilter = iaRouter.YFilter
    iaRouter.EntityData.YangName = "ia-router"
    iaRouter.EntityData.BundleName = "cisco_ios_xe"
    iaRouter.EntityData.ParentYangName = "lsa-body"
    iaRouter.EntityData.SegmentPath = "ia-router"
    iaRouter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaRouter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaRouter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaRouter.EntityData.Children = make(map[string]types.YChild)
    iaRouter.EntityData.Leafs = make(map[string]types.YLeaf)
    iaRouter.EntityData.Leafs["metric"] = types.YLeaf{"Metric", iaRouter.Metric}
    iaRouter.EntityData.Leafs["destination-router-id"] = types.YLeaf{"DestinationRouterId", iaRouter.DestinationRouterId}
    iaRouter.EntityData.Leafs["lsa-ia-options"] = types.YLeaf{"LsaIaOptions", iaRouter.LsaIaOptions}
    return &(iaRouter.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal
// OSPFv3 LSA external
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags
}

func (lsaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal) GetEntityData() *types.CommonEntityData {
    lsaExternal.EntityData.YFilter = lsaExternal.YFilter
    lsaExternal.EntityData.YangName = "lsa-external"
    lsaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaExternal.EntityData.ParentYangName = "lsa-body"
    lsaExternal.EntityData.SegmentPath = "lsa-external"
    lsaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaExternal.EntityData.Children = make(map[string]types.YChild)
    lsaExternal.EntityData.Children["flags"] = types.YChild{"Flags", &lsaExternal.Flags}
    lsaExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", lsaExternal.Metric}
    lsaExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", lsaExternal.ReferencedLsType}
    lsaExternal.EntityData.Leafs["external-prefix"] = types.YLeaf{"ExternalPrefix", lsaExternal.ExternalPrefix}
    lsaExternal.EntityData.Leafs["external-prefix-options"] = types.YLeaf{"ExternalPrefixOptions", lsaExternal.ExternalPrefixOptions}
    lsaExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", lsaExternal.ForwardingAddress}
    lsaExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", lsaExternal.ExternalRouteTag}
    lsaExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", lsaExternal.ReferencedLinkStateId}
    return &(lsaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    flags.EntityData.Leafs["e-flag"] = types.YLeaf{"EFlag", flags.EFlag}
    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa
// OSPFv3 NSSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA LSA.
    LsaNssaExternal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal
}

func (nssa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xe"
    nssa.EntityData.ParentYangName = "lsa-body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssa.EntityData.Children = make(map[string]types.YChild)
    nssa.EntityData.Children["lsa-nssa-external"] = types.YChild{"LsaNssaExternal", &nssa.LsaNssaExternal}
    nssa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nssa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal
// NSSA LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags
}

func (lsaNssaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal) GetEntityData() *types.CommonEntityData {
    lsaNssaExternal.EntityData.YFilter = lsaNssaExternal.YFilter
    lsaNssaExternal.EntityData.YangName = "lsa-nssa-external"
    lsaNssaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaNssaExternal.EntityData.ParentYangName = "nssa"
    lsaNssaExternal.EntityData.SegmentPath = "lsa-nssa-external"
    lsaNssaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaNssaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaNssaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaNssaExternal.EntityData.Children = make(map[string]types.YChild)
    lsaNssaExternal.EntityData.Children["flags"] = types.YChild{"Flags", &lsaNssaExternal.Flags}
    lsaNssaExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaNssaExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", lsaNssaExternal.Metric}
    lsaNssaExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", lsaNssaExternal.ReferencedLsType}
    lsaNssaExternal.EntityData.Leafs["external-prefix"] = types.YLeaf{"ExternalPrefix", lsaNssaExternal.ExternalPrefix}
    lsaNssaExternal.EntityData.Leafs["external-prefix-options"] = types.YLeaf{"ExternalPrefixOptions", lsaNssaExternal.ExternalPrefixOptions}
    lsaNssaExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", lsaNssaExternal.ForwardingAddress}
    lsaNssaExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", lsaNssaExternal.ExternalRouteTag}
    lsaNssaExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", lsaNssaExternal.ReferencedLinkStateId}
    return &(lsaNssaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-nssa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    flags.EntityData.Leafs["e-flag"] = types.YLeaf{"EFlag", flags.EFlag}
    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData
// OSPFv3 Link data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router priority of the interce. The type is interface{} with range: 0..255.
    RtrPriority interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // Link data LSA options. The type is map[string]bool.
    LsaIdOptions interface{}
}

func (linkData *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xe"
    linkData.EntityData.ParentYangName = "lsa-body"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkData.EntityData.Children = make(map[string]types.YChild)
    linkData.EntityData.Leafs = make(map[string]types.YLeaf)
    linkData.EntityData.Leafs["rtr-priority"] = types.YLeaf{"RtrPriority", linkData.RtrPriority}
    linkData.EntityData.Leafs["link-local-interface-address"] = types.YLeaf{"LinkLocalInterfaceAddress", linkData.LinkLocalInterfaceAddress}
    linkData.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", linkData.NumOfPrefixes}
    linkData.EntityData.Leafs["lsa-id-options"] = types.YLeaf{"LsaIdOptions", linkData.LsaIdOptions}
    return &(linkData.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix
// OSPFv3 Intra area prefixes
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}
}

func (iaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix) GetEntityData() *types.CommonEntityData {
    iaPrefix.EntityData.YFilter = iaPrefix.YFilter
    iaPrefix.EntityData.YangName = "ia-prefix"
    iaPrefix.EntityData.BundleName = "cisco_ios_xe"
    iaPrefix.EntityData.ParentYangName = "lsa-body"
    iaPrefix.EntityData.SegmentPath = "ia-prefix"
    iaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaPrefix.EntityData.Children = make(map[string]types.YChild)
    iaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    iaPrefix.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", iaPrefix.ReferencedLsType}
    iaPrefix.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", iaPrefix.ReferencedLinkStateId}
    iaPrefix.EntityData.Leafs["referenced-adv-router"] = types.YLeaf{"ReferencedAdvRouter", iaPrefix.ReferencedAdvRouter}
    iaPrefix.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", iaPrefix.NumOfPrefixes}
    return &(iaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3Link
// OSPFv3 links
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor router ID. The type is interface{} with
    // range: 0..4294967295.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv3Link *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3Link) GetEntityData() *types.CommonEntityData {
    ospfv3Link.EntityData.YFilter = ospfv3Link.YFilter
    ospfv3Link.EntityData.YangName = "ospfv3-link"
    ospfv3Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Link.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3Link.EntityData.SegmentPath = "ospfv3-link" + "[interface-id='" + fmt.Sprintf("%v", ospfv3Link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", ospfv3Link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", ospfv3Link.NeighborRouterId) + "']"
    ospfv3Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Link.EntityData.Children = make(map[string]types.YChild)
    ospfv3Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3Link.EntityData.Leafs["interface-id"] = types.YLeaf{"InterfaceId", ospfv3Link.InterfaceId}
    ospfv3Link.EntityData.Leafs["neighbor-interface-id"] = types.YLeaf{"NeighborInterfaceId", ospfv3Link.NeighborInterfaceId}
    ospfv3Link.EntityData.Leafs["neighbor-router-id"] = types.YLeaf{"NeighborRouterId", ospfv3Link.NeighborRouterId}
    ospfv3Link.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv3Link.Type_}
    ospfv3Link.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv3Link.Metric}
    return &(ospfv3Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList
// OSPFv3 prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3PrefixList *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList) GetEntityData() *types.CommonEntityData {
    ospfv3PrefixList.EntityData.YFilter = ospfv3PrefixList.YFilter
    ospfv3PrefixList.EntityData.YangName = "ospfv3-prefix-list"
    ospfv3PrefixList.EntityData.BundleName = "cisco_ios_xe"
    ospfv3PrefixList.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3PrefixList.EntityData.SegmentPath = "ospfv3-prefix-list" + "[prefix='" + fmt.Sprintf("%v", ospfv3PrefixList.Prefix) + "']"
    ospfv3PrefixList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3PrefixList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3PrefixList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3PrefixList.EntityData.Children = make(map[string]types.YChild)
    ospfv3PrefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3PrefixList.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ospfv3PrefixList.Prefix}
    ospfv3PrefixList.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", ospfv3PrefixList.PrefixOptions}
    return &(ospfv3PrefixList.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix
// OSPFv3 intra-area prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3IaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix) GetEntityData() *types.CommonEntityData {
    ospfv3IaPrefix.EntityData.YFilter = ospfv3IaPrefix.YFilter
    ospfv3IaPrefix.EntityData.YangName = "ospfv3-ia-prefix"
    ospfv3IaPrefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3IaPrefix.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3IaPrefix.EntityData.SegmentPath = "ospfv3-ia-prefix" + "[prefix='" + fmt.Sprintf("%v", ospfv3IaPrefix.Prefix) + "']"
    ospfv3IaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3IaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3IaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3IaPrefix.EntityData.Children = make(map[string]types.YChild)
    ospfv3IaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3IaPrefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ospfv3IaPrefix.Prefix}
    ospfv3IaPrefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", ospfv3IaPrefix.PrefixOptions}
    return &(ospfv3IaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_MultiTopology
// OSPF multi-topology interface augmentation
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_MultiTopology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string.
    Name interface{}
}

func (multiTopology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_MultiTopology) GetEntityData() *types.CommonEntityData {
    multiTopology.EntityData.YFilter = multiTopology.YFilter
    multiTopology.EntityData.YangName = "multi-topology"
    multiTopology.EntityData.BundleName = "cisco_ios_xe"
    multiTopology.EntityData.ParentYangName = "link-scope-lsa"
    multiTopology.EntityData.SegmentPath = "multi-topology" + "[name='" + fmt.Sprintf("%v", multiTopology.Name) + "']"
    multiTopology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    multiTopology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    multiTopology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    multiTopology.EntityData.Children = make(map[string]types.YChild)
    multiTopology.EntityData.Leafs = make(map[string]types.YLeaf)
    multiTopology.EntityData.Leafs["name"] = types.YLeaf{"Name", multiTopology.Name}
    return &(multiTopology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Tlv
// Link TLV
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Tlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link type. The type is interface{} with range: 0..255.
    LinkType interface{}

    // Link ID. The type is interface{} with range: 0..4294967295.
    LinkId interface{}

    // List of local interface IPv4 addresses. The type is one of the following
    // types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalIfIpv4Addr []interface{}

    // List of remote interface IPv4 addresses. The type is one of the following
    // types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalRemoteIpv4Addr []interface{}

    // TE metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Maximum bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxBandwidth interface{}

    // Maximum reservable bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxReservableBandwidth interface{}

    // Unrseerved bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    UnreservedBandwidth interface{}

    // Administrative group/Resource class/Color. The type is interface{} with
    // range: 0..4294967295.
    AdminGroup interface{}
}

func (tlv *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_Tlv) GetEntityData() *types.CommonEntityData {
    tlv.EntityData.YFilter = tlv.YFilter
    tlv.EntityData.YangName = "tlv"
    tlv.EntityData.BundleName = "cisco_ios_xe"
    tlv.EntityData.ParentYangName = "link-scope-lsa"
    tlv.EntityData.SegmentPath = "tlv"
    tlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tlv.EntityData.Children = make(map[string]types.YChild)
    tlv.EntityData.Leafs = make(map[string]types.YLeaf)
    tlv.EntityData.Leafs["link-type"] = types.YLeaf{"LinkType", tlv.LinkType}
    tlv.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", tlv.LinkId}
    tlv.EntityData.Leafs["local-if-ipv4-addr"] = types.YLeaf{"LocalIfIpv4Addr", tlv.LocalIfIpv4Addr}
    tlv.EntityData.Leafs["local-remote-ipv4-addr"] = types.YLeaf{"LocalRemoteIpv4Addr", tlv.LocalRemoteIpv4Addr}
    tlv.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", tlv.TeMetric}
    tlv.EntityData.Leafs["max-bandwidth"] = types.YLeaf{"MaxBandwidth", tlv.MaxBandwidth}
    tlv.EntityData.Leafs["max-reservable-bandwidth"] = types.YLeaf{"MaxReservableBandwidth", tlv.MaxReservableBandwidth}
    tlv.EntityData.Leafs["unreserved-bandwidth"] = types.YLeaf{"UnreservedBandwidth", tlv.UnreservedBandwidth}
    tlv.EntityData.Leafs["admin-group"] = types.YLeaf{"AdminGroup", tlv.AdminGroup}
    return &(tlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_UnknownSubTlv
// OSPFv2 Unknown sub TLV
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_UnknownSubTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type_ interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (unknownSubTlv *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_LinkScopeLsa_UnknownSubTlv) GetEntityData() *types.CommonEntityData {
    unknownSubTlv.EntityData.YFilter = unknownSubTlv.YFilter
    unknownSubTlv.EntityData.YangName = "unknown-sub-tlv"
    unknownSubTlv.EntityData.BundleName = "cisco_ios_xe"
    unknownSubTlv.EntityData.ParentYangName = "link-scope-lsa"
    unknownSubTlv.EntityData.SegmentPath = "unknown-sub-tlv" + "[type='" + fmt.Sprintf("%v", unknownSubTlv.Type_) + "']"
    unknownSubTlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    unknownSubTlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    unknownSubTlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    unknownSubTlv.EntityData.Children = make(map[string]types.YChild)
    unknownSubTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownSubTlv.EntityData.Leafs["type"] = types.YLeaf{"Type_", unknownSubTlv.Type_}
    unknownSubTlv.EntityData.Leafs["length"] = types.YLeaf{"Length", unknownSubTlv.Length}
    unknownSubTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownSubTlv.Value}
    return &(unknownSubTlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa
// List OSPF area scope LSA databases
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSA Type. The type is interface{} with range:
    // 0..4294967295.
    LsaType interface{}

    // This attribute is a key. Advertising router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is slice of interface{} with range: 0..255.
    RawData []interface{}

    // OSPFv2 LSA.
    Ospfv2Lsa OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa

    // Router LSA link. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link.
    Ospfv2Link []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link

    // Summary LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Topology.
    Ospfv2Topology []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Topology

    // External LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2External.
    Ospfv2External []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2External

    // OSPFv3 LSA.
    Ospfv3Lsa OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa

    // OSPFv3 links. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Link.
    Ospfv3Link []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Link

    // OSPFv3 prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Prefix.
    Ospfv3Prefix []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Prefix

    // OSPFv3 intra-area prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix.
    Ospfv3IaPrefix []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix
}

func (areaScopeLsa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa) GetEntityData() *types.CommonEntityData {
    areaScopeLsa.EntityData.YFilter = areaScopeLsa.YFilter
    areaScopeLsa.EntityData.YangName = "area-scope-lsa"
    areaScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    areaScopeLsa.EntityData.ParentYangName = "intf-link-scope-lsas"
    areaScopeLsa.EntityData.SegmentPath = "area-scope-lsa" + "[lsa-type='" + fmt.Sprintf("%v", areaScopeLsa.LsaType) + "']" + "[adv-router='" + fmt.Sprintf("%v", areaScopeLsa.AdvRouter) + "']"
    areaScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    areaScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    areaScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    areaScopeLsa.EntityData.Children = make(map[string]types.YChild)
    areaScopeLsa.EntityData.Children["ospfv2-lsa"] = types.YChild{"Ospfv2Lsa", &areaScopeLsa.Ospfv2Lsa}
    areaScopeLsa.EntityData.Children["ospfv2-link"] = types.YChild{"Ospfv2Link", nil}
    for i := range areaScopeLsa.Ospfv2Link {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv2Link[i])] = types.YChild{"Ospfv2Link", &areaScopeLsa.Ospfv2Link[i]}
    }
    areaScopeLsa.EntityData.Children["ospfv2-topology"] = types.YChild{"Ospfv2Topology", nil}
    for i := range areaScopeLsa.Ospfv2Topology {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv2Topology[i])] = types.YChild{"Ospfv2Topology", &areaScopeLsa.Ospfv2Topology[i]}
    }
    areaScopeLsa.EntityData.Children["ospfv2-external"] = types.YChild{"Ospfv2External", nil}
    for i := range areaScopeLsa.Ospfv2External {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv2External[i])] = types.YChild{"Ospfv2External", &areaScopeLsa.Ospfv2External[i]}
    }
    areaScopeLsa.EntityData.Children["ospfv3-lsa"] = types.YChild{"Ospfv3Lsa", &areaScopeLsa.Ospfv3Lsa}
    areaScopeLsa.EntityData.Children["ospfv3-link"] = types.YChild{"Ospfv3Link", nil}
    for i := range areaScopeLsa.Ospfv3Link {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv3Link[i])] = types.YChild{"Ospfv3Link", &areaScopeLsa.Ospfv3Link[i]}
    }
    areaScopeLsa.EntityData.Children["ospfv3-prefix"] = types.YChild{"Ospfv3Prefix", nil}
    for i := range areaScopeLsa.Ospfv3Prefix {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv3Prefix[i])] = types.YChild{"Ospfv3Prefix", &areaScopeLsa.Ospfv3Prefix[i]}
    }
    areaScopeLsa.EntityData.Children["ospfv3-ia-prefix"] = types.YChild{"Ospfv3IaPrefix", nil}
    for i := range areaScopeLsa.Ospfv3IaPrefix {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv3IaPrefix[i])] = types.YChild{"Ospfv3IaPrefix", &areaScopeLsa.Ospfv3IaPrefix[i]}
    }
    areaScopeLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    areaScopeLsa.EntityData.Leafs["lsa-type"] = types.YLeaf{"LsaType", areaScopeLsa.LsaType}
    areaScopeLsa.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", areaScopeLsa.AdvRouter}
    areaScopeLsa.EntityData.Leafs["decoded-completed"] = types.YLeaf{"DecodedCompleted", areaScopeLsa.DecodedCompleted}
    areaScopeLsa.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", areaScopeLsa.RawData}
    return &(areaScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa
// OSPFv2 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header

    // Decoded OSPFv2 LSA body data.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody
}

func (ospfv2Lsa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa) GetEntityData() *types.CommonEntityData {
    ospfv2Lsa.EntityData.YFilter = ospfv2Lsa.YFilter
    ospfv2Lsa.EntityData.YangName = "ospfv2-lsa"
    ospfv2Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Lsa.EntityData.SegmentPath = "ospfv2-lsa"
    ospfv2Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Lsa.EntityData.Children = make(map[string]types.YChild)
    ospfv2Lsa.EntityData.Children["header"] = types.YChild{"Header", &ospfv2Lsa.Header}
    ospfv2Lsa.EntityData.Children["lsa-body"] = types.YChild{"LsaBody", &ospfv2Lsa.LsaBody}
    ospfv2Lsa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header
// Decoded OSPFv2 LSA header data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255.
    OpaqueType interface{}

    // Opaque ID. The type is interface{} with range: 0..4294967295.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type_ interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}

    // LSA options. The type is map[string]bool.
    FlagOptions interface{}
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv2-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["opaque-type"] = types.YLeaf{"OpaqueType", header.OpaqueType}
    header.EntityData.Leafs["opaque-id"] = types.YLeaf{"OpaqueId", header.OpaqueId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    header.EntityData.Leafs["flag-options"] = types.YLeaf{"FlagOptions", header.FlagOptions}
    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody
// Decoded OSPFv2 LSA body data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Summary mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryMask interface{}

    // External mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ExternalMask interface{}

    // LSA body flags. The type is map[string]bool.
    BodyFlagOptions interface{}

    // Network details.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv2-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = make(map[string]types.YChild)
    lsaBody.EntityData.Children["network"] = types.YChild{"Network", &lsaBody.Network}
    lsaBody.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaBody.EntityData.Leafs["num-of-links"] = types.YLeaf{"NumOfLinks", lsaBody.NumOfLinks}
    lsaBody.EntityData.Leafs["summary-mask"] = types.YLeaf{"SummaryMask", lsaBody.SummaryMask}
    lsaBody.EntityData.Leafs["external-mask"] = types.YLeaf{"ExternalMask", lsaBody.ExternalMask}
    lsaBody.EntityData.Leafs["body-flag-options"] = types.YLeaf{"BodyFlagOptions", lsaBody.BodyFlagOptions}
    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network
// Network details
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP network mask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", network.NetworkMask}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link
// Router LSA link
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is interface{} with range:
    // 0..4294967295.
    LinkId interface{}

    // This attribute is a key. Link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Topology specific information. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology.
    Ospfv2Topology []OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology
}

func (ospfv2Link *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link) GetEntityData() *types.CommonEntityData {
    ospfv2Link.EntityData.YFilter = ospfv2Link.YFilter
    ospfv2Link.EntityData.YangName = "ospfv2-link"
    ospfv2Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Link.EntityData.SegmentPath = "ospfv2-link" + "[link-id='" + fmt.Sprintf("%v", ospfv2Link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", ospfv2Link.LinkData) + "']"
    ospfv2Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Link.EntityData.Children = make(map[string]types.YChild)
    ospfv2Link.EntityData.Children["ospfv2-topology"] = types.YChild{"Ospfv2Topology", nil}
    for i := range ospfv2Link.Ospfv2Topology {
        ospfv2Link.EntityData.Children[types.GetSegmentPath(&ospfv2Link.Ospfv2Topology[i])] = types.YChild{"Ospfv2Topology", &ospfv2Link.Ospfv2Topology[i]}
    }
    ospfv2Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Link.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", ospfv2Link.LinkId}
    ospfv2Link.EntityData.Leafs["link-data"] = types.YLeaf{"LinkData", ospfv2Link.LinkData}
    ospfv2Link.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv2Link.Type_}
    return &(ospfv2Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology
// Topology specific information
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "ospfv2-link"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + "[mt-id='" + fmt.Sprintf("%v", ospfv2Topology.MtId) + "']"
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = make(map[string]types.YChild)
    ospfv2Topology.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2Topology.MtId}
    ospfv2Topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2Topology.Metric}
    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Topology
// Summary LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + "[mt-id='" + fmt.Sprintf("%v", ospfv2Topology.MtId) + "']"
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = make(map[string]types.YChild)
    ospfv2Topology.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2Topology.MtId}
    ospfv2Topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2Topology.Metric}
    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2External
// External LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled on the link. The type
    // is interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (ospfv2External *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv2External) GetEntityData() *types.CommonEntityData {
    ospfv2External.EntityData.YFilter = ospfv2External.YFilter
    ospfv2External.EntityData.YangName = "ospfv2-external"
    ospfv2External.EntityData.BundleName = "cisco_ios_xe"
    ospfv2External.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2External.EntityData.SegmentPath = "ospfv2-external" + "[mt-id='" + fmt.Sprintf("%v", ospfv2External.MtId) + "']"
    ospfv2External.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2External.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2External.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2External.EntityData.Children = make(map[string]types.YChild)
    ospfv2External.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2External.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2External.MtId}
    ospfv2External.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2External.Metric}
    ospfv2External.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", ospfv2External.ForwardingAddress}
    ospfv2External.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", ospfv2External.ExternalRouteTag}
    return &(ospfv2External.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa
// OSPFv3 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header

    // Decoded OSPFv3 LSA body.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody
}

func (ospfv3Lsa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa) GetEntityData() *types.CommonEntityData {
    ospfv3Lsa.EntityData.YFilter = ospfv3Lsa.YFilter
    ospfv3Lsa.EntityData.YangName = "ospfv3-lsa"
    ospfv3Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Lsa.EntityData.SegmentPath = "ospfv3-lsa"
    ospfv3Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Lsa.EntityData.Children = make(map[string]types.YChild)
    ospfv3Lsa.EntityData.Children["header"] = types.YChild{"Header", &ospfv3Lsa.Header}
    ospfv3Lsa.EntityData.Children["lsa-body"] = types.YChild{"LsaBody", &ospfv3Lsa.LsaBody}
    ospfv3Lsa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header
// Decoded OSPFv3 LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // OSPFv3 LSA options. The type is map[string]bool.
    LsaHdrOptions interface{}

    // LSA header.
    LsaHeader OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv3-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Children["lsa-header"] = types.YChild{"LsaHeader", &header.LsaHeader}
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["lsa-hdr-options"] = types.YLeaf{"LsaHdrOptions", header.LsaHdrOptions}
    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader
// LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type_ interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}
}

func (lsaHeader *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader) GetEntityData() *types.CommonEntityData {
    lsaHeader.EntityData.YFilter = lsaHeader.YFilter
    lsaHeader.EntityData.YangName = "lsa-header"
    lsaHeader.EntityData.BundleName = "cisco_ios_xe"
    lsaHeader.EntityData.ParentYangName = "header"
    lsaHeader.EntityData.SegmentPath = "lsa-header"
    lsaHeader.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaHeader.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaHeader.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaHeader.EntityData.Children = make(map[string]types.YChild)
    lsaHeader.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaHeader.EntityData.Leafs["age"] = types.YLeaf{"Age", lsaHeader.Age}
    lsaHeader.EntityData.Leafs["type"] = types.YLeaf{"Type_", lsaHeader.Type_}
    lsaHeader.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", lsaHeader.AdvRouter}
    lsaHeader.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", lsaHeader.SeqNum}
    lsaHeader.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", lsaHeader.Checksum}
    lsaHeader.EntityData.Leafs["length"] = types.YLeaf{"Length", lsaHeader.Length}
    return &(lsaHeader.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody
// Decoded OSPFv3 LSA body
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA options. The type is map[string]bool.
    LsaFlagOptions interface{}

    // LSA Body Flags. The type is map[string]bool.
    LsaBodyFlags interface{}

    // OSPFv3 network.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network

    // OSPFv3 inter area prefix.
    Prefix OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix

    // OSPFv3 inter area router.
    IaRouter OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter

    // OSPFv3 LSA external.
    LsaExternal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal

    // OSPFv3 NSSA.
    Nssa OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa

    // OSPFv3 Link data.
    LinkData OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData

    // OSPFv3 Intra area prefixes.
    IaPrefix OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv3-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = make(map[string]types.YChild)
    lsaBody.EntityData.Children["network"] = types.YChild{"Network", &lsaBody.Network}
    lsaBody.EntityData.Children["prefix"] = types.YChild{"Prefix", &lsaBody.Prefix}
    lsaBody.EntityData.Children["ia-router"] = types.YChild{"IaRouter", &lsaBody.IaRouter}
    lsaBody.EntityData.Children["lsa-external"] = types.YChild{"LsaExternal", &lsaBody.LsaExternal}
    lsaBody.EntityData.Children["nssa"] = types.YChild{"Nssa", &lsaBody.Nssa}
    lsaBody.EntityData.Children["link-data"] = types.YChild{"LinkData", &lsaBody.LinkData}
    lsaBody.EntityData.Children["ia-prefix"] = types.YChild{"IaPrefix", &lsaBody.IaPrefix}
    lsaBody.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaBody.EntityData.Leafs["lsa-flag-options"] = types.YLeaf{"LsaFlagOptions", lsaBody.LsaFlagOptions}
    lsaBody.EntityData.Leafs["lsa-body-flags"] = types.YLeaf{"LsaBodyFlags", lsaBody.LsaBodyFlags}
    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network
// OSPFv3 network
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}

    // Network LSA options. The type is map[string]bool.
    LsaNetOptions interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    network.EntityData.Leafs["lsa-net-options"] = types.YLeaf{"LsaNetOptions", network.LsaNetOptions}
    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix
// OSPFv3 inter area prefix
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Inter area Prefix. The type is string.
    IaPrefix interface{}

    // Inter area prefix options. The type is string.
    IaPrefixOptions interface{}
}

func (prefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xe"
    prefix.EntityData.ParentYangName = "lsa-body"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["metric"] = types.YLeaf{"Metric", prefix.Metric}
    prefix.EntityData.Leafs["ia-prefix"] = types.YLeaf{"IaPrefix", prefix.IaPrefix}
    prefix.EntityData.Leafs["ia-prefix-options"] = types.YLeaf{"IaPrefixOptions", prefix.IaPrefixOptions}
    return &(prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter
// OSPFv3 inter area router
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Router ID of the router being described by the LSA. The type is interface{}
    // with range: 0..4294967295.
    DestinationRouterId interface{}

    // Inter area LSA options. The type is map[string]bool.
    LsaIaOptions interface{}
}

func (iaRouter *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter) GetEntityData() *types.CommonEntityData {
    iaRouter.EntityData.YFilter = iaRouter.YFilter
    iaRouter.EntityData.YangName = "ia-router"
    iaRouter.EntityData.BundleName = "cisco_ios_xe"
    iaRouter.EntityData.ParentYangName = "lsa-body"
    iaRouter.EntityData.SegmentPath = "ia-router"
    iaRouter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaRouter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaRouter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaRouter.EntityData.Children = make(map[string]types.YChild)
    iaRouter.EntityData.Leafs = make(map[string]types.YLeaf)
    iaRouter.EntityData.Leafs["metric"] = types.YLeaf{"Metric", iaRouter.Metric}
    iaRouter.EntityData.Leafs["destination-router-id"] = types.YLeaf{"DestinationRouterId", iaRouter.DestinationRouterId}
    iaRouter.EntityData.Leafs["lsa-ia-options"] = types.YLeaf{"LsaIaOptions", iaRouter.LsaIaOptions}
    return &(iaRouter.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal
// OSPFv3 LSA external
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags
}

func (lsaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal) GetEntityData() *types.CommonEntityData {
    lsaExternal.EntityData.YFilter = lsaExternal.YFilter
    lsaExternal.EntityData.YangName = "lsa-external"
    lsaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaExternal.EntityData.ParentYangName = "lsa-body"
    lsaExternal.EntityData.SegmentPath = "lsa-external"
    lsaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaExternal.EntityData.Children = make(map[string]types.YChild)
    lsaExternal.EntityData.Children["flags"] = types.YChild{"Flags", &lsaExternal.Flags}
    lsaExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", lsaExternal.Metric}
    lsaExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", lsaExternal.ReferencedLsType}
    lsaExternal.EntityData.Leafs["external-prefix"] = types.YLeaf{"ExternalPrefix", lsaExternal.ExternalPrefix}
    lsaExternal.EntityData.Leafs["external-prefix-options"] = types.YLeaf{"ExternalPrefixOptions", lsaExternal.ExternalPrefixOptions}
    lsaExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", lsaExternal.ForwardingAddress}
    lsaExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", lsaExternal.ExternalRouteTag}
    lsaExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", lsaExternal.ReferencedLinkStateId}
    return &(lsaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    flags.EntityData.Leafs["e-flag"] = types.YLeaf{"EFlag", flags.EFlag}
    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa
// OSPFv3 NSSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA LSA.
    LsaNssaExternal OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
}

func (nssa *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xe"
    nssa.EntityData.ParentYangName = "lsa-body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssa.EntityData.Children = make(map[string]types.YChild)
    nssa.EntityData.Children["lsa-nssa-external"] = types.YChild{"LsaNssaExternal", &nssa.LsaNssaExternal}
    nssa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nssa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
// NSSA LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
}

func (lsaNssaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal) GetEntityData() *types.CommonEntityData {
    lsaNssaExternal.EntityData.YFilter = lsaNssaExternal.YFilter
    lsaNssaExternal.EntityData.YangName = "lsa-nssa-external"
    lsaNssaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaNssaExternal.EntityData.ParentYangName = "nssa"
    lsaNssaExternal.EntityData.SegmentPath = "lsa-nssa-external"
    lsaNssaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaNssaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaNssaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaNssaExternal.EntityData.Children = make(map[string]types.YChild)
    lsaNssaExternal.EntityData.Children["flags"] = types.YChild{"Flags", &lsaNssaExternal.Flags}
    lsaNssaExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaNssaExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", lsaNssaExternal.Metric}
    lsaNssaExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", lsaNssaExternal.ReferencedLsType}
    lsaNssaExternal.EntityData.Leafs["external-prefix"] = types.YLeaf{"ExternalPrefix", lsaNssaExternal.ExternalPrefix}
    lsaNssaExternal.EntityData.Leafs["external-prefix-options"] = types.YLeaf{"ExternalPrefixOptions", lsaNssaExternal.ExternalPrefixOptions}
    lsaNssaExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", lsaNssaExternal.ForwardingAddress}
    lsaNssaExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", lsaNssaExternal.ExternalRouteTag}
    lsaNssaExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", lsaNssaExternal.ReferencedLinkStateId}
    return &(lsaNssaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-nssa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    flags.EntityData.Leafs["e-flag"] = types.YLeaf{"EFlag", flags.EFlag}
    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData
// OSPFv3 Link data
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router priority of the interce. The type is interface{} with range: 0..255.
    RtrPriority interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // Link data LSA options. The type is map[string]bool.
    LsaIdOptions interface{}
}

func (linkData *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xe"
    linkData.EntityData.ParentYangName = "lsa-body"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkData.EntityData.Children = make(map[string]types.YChild)
    linkData.EntityData.Leafs = make(map[string]types.YLeaf)
    linkData.EntityData.Leafs["rtr-priority"] = types.YLeaf{"RtrPriority", linkData.RtrPriority}
    linkData.EntityData.Leafs["link-local-interface-address"] = types.YLeaf{"LinkLocalInterfaceAddress", linkData.LinkLocalInterfaceAddress}
    linkData.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", linkData.NumOfPrefixes}
    linkData.EntityData.Leafs["lsa-id-options"] = types.YLeaf{"LsaIdOptions", linkData.LsaIdOptions}
    return &(linkData.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix
// OSPFv3 Intra area prefixes
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}
}

func (iaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix) GetEntityData() *types.CommonEntityData {
    iaPrefix.EntityData.YFilter = iaPrefix.YFilter
    iaPrefix.EntityData.YangName = "ia-prefix"
    iaPrefix.EntityData.BundleName = "cisco_ios_xe"
    iaPrefix.EntityData.ParentYangName = "lsa-body"
    iaPrefix.EntityData.SegmentPath = "ia-prefix"
    iaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaPrefix.EntityData.Children = make(map[string]types.YChild)
    iaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    iaPrefix.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", iaPrefix.ReferencedLsType}
    iaPrefix.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", iaPrefix.ReferencedLinkStateId}
    iaPrefix.EntityData.Leafs["referenced-adv-router"] = types.YLeaf{"ReferencedAdvRouter", iaPrefix.ReferencedAdvRouter}
    iaPrefix.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", iaPrefix.NumOfPrefixes}
    return &(iaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Link
// OSPFv3 links
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor router ID. The type is interface{} with
    // range: 0..4294967295.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv3Link *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Link) GetEntityData() *types.CommonEntityData {
    ospfv3Link.EntityData.YFilter = ospfv3Link.YFilter
    ospfv3Link.EntityData.YangName = "ospfv3-link"
    ospfv3Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Link.EntityData.SegmentPath = "ospfv3-link" + "[interface-id='" + fmt.Sprintf("%v", ospfv3Link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", ospfv3Link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", ospfv3Link.NeighborRouterId) + "']"
    ospfv3Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Link.EntityData.Children = make(map[string]types.YChild)
    ospfv3Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3Link.EntityData.Leafs["interface-id"] = types.YLeaf{"InterfaceId", ospfv3Link.InterfaceId}
    ospfv3Link.EntityData.Leafs["neighbor-interface-id"] = types.YLeaf{"NeighborInterfaceId", ospfv3Link.NeighborInterfaceId}
    ospfv3Link.EntityData.Leafs["neighbor-router-id"] = types.YLeaf{"NeighborRouterId", ospfv3Link.NeighborRouterId}
    ospfv3Link.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv3Link.Type_}
    ospfv3Link.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv3Link.Metric}
    return &(ospfv3Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Prefix
// OSPFv3 prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3Prefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3Prefix) GetEntityData() *types.CommonEntityData {
    ospfv3Prefix.EntityData.YFilter = ospfv3Prefix.YFilter
    ospfv3Prefix.EntityData.YangName = "ospfv3-prefix"
    ospfv3Prefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Prefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Prefix.EntityData.SegmentPath = "ospfv3-prefix" + "[prefix='" + fmt.Sprintf("%v", ospfv3Prefix.Prefix) + "']"
    ospfv3Prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Prefix.EntityData.Children = make(map[string]types.YChild)
    ospfv3Prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3Prefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ospfv3Prefix.Prefix}
    ospfv3Prefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", ospfv3Prefix.PrefixOptions}
    return &(ospfv3Prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix
// OSPFv3 intra-area prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3IaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfLinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix) GetEntityData() *types.CommonEntityData {
    ospfv3IaPrefix.EntityData.YFilter = ospfv3IaPrefix.YFilter
    ospfv3IaPrefix.EntityData.YangName = "ospfv3-ia-prefix"
    ospfv3IaPrefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3IaPrefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3IaPrefix.EntityData.SegmentPath = "ospfv3-ia-prefix" + "[prefix='" + fmt.Sprintf("%v", ospfv3IaPrefix.Prefix) + "']"
    ospfv3IaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3IaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3IaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3IaPrefix.EntityData.Children = make(map[string]types.YChild)
    ospfv3IaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3IaPrefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ospfv3IaPrefix.Prefix}
    ospfv3IaPrefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", ospfv3IaPrefix.PrefixOptions}
    return &(ospfv3IaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfMultiTopology
// OSPF interface topology
type OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfMultiTopology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string.
    Name interface{}
}

func (intfMultiTopology *OspfOperData_OspfState_OspfInstance_OspfArea_OspfInterface_IntfMultiTopology) GetEntityData() *types.CommonEntityData {
    intfMultiTopology.EntityData.YFilter = intfMultiTopology.YFilter
    intfMultiTopology.EntityData.YangName = "intf-multi-topology"
    intfMultiTopology.EntityData.BundleName = "cisco_ios_xe"
    intfMultiTopology.EntityData.ParentYangName = "ospf-interface"
    intfMultiTopology.EntityData.SegmentPath = "intf-multi-topology" + "[name='" + fmt.Sprintf("%v", intfMultiTopology.Name) + "']"
    intfMultiTopology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intfMultiTopology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intfMultiTopology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intfMultiTopology.EntityData.Children = make(map[string]types.YChild)
    intfMultiTopology.EntityData.Leafs = make(map[string]types.YLeaf)
    intfMultiTopology.EntityData.Leafs["name"] = types.YLeaf{"Name", intfMultiTopology.Name}
    return &(intfMultiTopology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa
// List of OSPF area scope LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF link scope LSA type. The type is interface{}
    // with range: 0..4294967295.
    LsaType interface{}

    // List of OSPF link scope LSAs. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa.
    AreaScopeLsa []OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_
}

func (areaScopeLsa *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa) GetEntityData() *types.CommonEntityData {
    areaScopeLsa.EntityData.YFilter = areaScopeLsa.YFilter
    areaScopeLsa.EntityData.YangName = "area-scope-lsa"
    areaScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    areaScopeLsa.EntityData.ParentYangName = "ospf-area"
    areaScopeLsa.EntityData.SegmentPath = "area-scope-lsa" + "[lsa-type='" + fmt.Sprintf("%v", areaScopeLsa.LsaType) + "']"
    areaScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    areaScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    areaScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    areaScopeLsa.EntityData.Children = make(map[string]types.YChild)
    areaScopeLsa.EntityData.Children["area-scope-lsa"] = types.YChild{"AreaScopeLsa", nil}
    for i := range areaScopeLsa.AreaScopeLsa {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.AreaScopeLsa[i])] = types.YChild{"AreaScopeLsa", &areaScopeLsa.AreaScopeLsa[i]}
    }
    areaScopeLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    areaScopeLsa.EntityData.Leafs["lsa-type"] = types.YLeaf{"LsaType", areaScopeLsa.LsaType}
    return &(areaScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_
// List of OSPF link scope LSAs
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSA Type. The type is interface{} with range:
    // 0..4294967295.
    LsaType interface{}

    // This attribute is a key. Advertising router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is slice of interface{} with range: 0..255.
    RawData []interface{}

    // OSPFv2 LSA.
    Ospfv2Lsa OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa

    // Router LSA link. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Link.
    Ospfv2Link []OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Link

    // Summary LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Topology.
    Ospfv2Topology []OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Topology

    // External LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2External.
    Ospfv2External []OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2External

    // OSPFv3 LSA.
    Ospfv3Lsa OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa

    // OSPFv3 links. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Link.
    Ospfv3Link []OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Link

    // OSPFv3 prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Prefix.
    Ospfv3Prefix []OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Prefix

    // OSPFv3 intra-area prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3IaPrefix.
    Ospfv3IaPrefix []OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3IaPrefix
}

func (areaScopeLsa_ *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa_) GetEntityData() *types.CommonEntityData {
    areaScopeLsa_.EntityData.YFilter = areaScopeLsa_.YFilter
    areaScopeLsa_.EntityData.YangName = "area-scope-lsa"
    areaScopeLsa_.EntityData.BundleName = "cisco_ios_xe"
    areaScopeLsa_.EntityData.ParentYangName = "area-scope-lsa"
    areaScopeLsa_.EntityData.SegmentPath = "area-scope-lsa" + "[lsa-type='" + fmt.Sprintf("%v", areaScopeLsa_.LsaType) + "']" + "[adv-router='" + fmt.Sprintf("%v", areaScopeLsa_.AdvRouter) + "']"
    areaScopeLsa_.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    areaScopeLsa_.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    areaScopeLsa_.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    areaScopeLsa_.EntityData.Children = make(map[string]types.YChild)
    areaScopeLsa_.EntityData.Children["ospfv2-lsa"] = types.YChild{"Ospfv2Lsa", &areaScopeLsa_.Ospfv2Lsa}
    areaScopeLsa_.EntityData.Children["ospfv2-link"] = types.YChild{"Ospfv2Link", nil}
    for i := range areaScopeLsa_.Ospfv2Link {
        areaScopeLsa_.EntityData.Children[types.GetSegmentPath(&areaScopeLsa_.Ospfv2Link[i])] = types.YChild{"Ospfv2Link", &areaScopeLsa_.Ospfv2Link[i]}
    }
    areaScopeLsa_.EntityData.Children["ospfv2-topology"] = types.YChild{"Ospfv2Topology", nil}
    for i := range areaScopeLsa_.Ospfv2Topology {
        areaScopeLsa_.EntityData.Children[types.GetSegmentPath(&areaScopeLsa_.Ospfv2Topology[i])] = types.YChild{"Ospfv2Topology", &areaScopeLsa_.Ospfv2Topology[i]}
    }
    areaScopeLsa_.EntityData.Children["ospfv2-external"] = types.YChild{"Ospfv2External", nil}
    for i := range areaScopeLsa_.Ospfv2External {
        areaScopeLsa_.EntityData.Children[types.GetSegmentPath(&areaScopeLsa_.Ospfv2External[i])] = types.YChild{"Ospfv2External", &areaScopeLsa_.Ospfv2External[i]}
    }
    areaScopeLsa_.EntityData.Children["ospfv3-lsa"] = types.YChild{"Ospfv3Lsa", &areaScopeLsa_.Ospfv3Lsa}
    areaScopeLsa_.EntityData.Children["ospfv3-link"] = types.YChild{"Ospfv3Link", nil}
    for i := range areaScopeLsa_.Ospfv3Link {
        areaScopeLsa_.EntityData.Children[types.GetSegmentPath(&areaScopeLsa_.Ospfv3Link[i])] = types.YChild{"Ospfv3Link", &areaScopeLsa_.Ospfv3Link[i]}
    }
    areaScopeLsa_.EntityData.Children["ospfv3-prefix"] = types.YChild{"Ospfv3Prefix", nil}
    for i := range areaScopeLsa_.Ospfv3Prefix {
        areaScopeLsa_.EntityData.Children[types.GetSegmentPath(&areaScopeLsa_.Ospfv3Prefix[i])] = types.YChild{"Ospfv3Prefix", &areaScopeLsa_.Ospfv3Prefix[i]}
    }
    areaScopeLsa_.EntityData.Children["ospfv3-ia-prefix"] = types.YChild{"Ospfv3IaPrefix", nil}
    for i := range areaScopeLsa_.Ospfv3IaPrefix {
        areaScopeLsa_.EntityData.Children[types.GetSegmentPath(&areaScopeLsa_.Ospfv3IaPrefix[i])] = types.YChild{"Ospfv3IaPrefix", &areaScopeLsa_.Ospfv3IaPrefix[i]}
    }
    areaScopeLsa_.EntityData.Leafs = make(map[string]types.YLeaf)
    areaScopeLsa_.EntityData.Leafs["lsa-type"] = types.YLeaf{"LsaType", areaScopeLsa_.LsaType}
    areaScopeLsa_.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", areaScopeLsa_.AdvRouter}
    areaScopeLsa_.EntityData.Leafs["decoded-completed"] = types.YLeaf{"DecodedCompleted", areaScopeLsa_.DecodedCompleted}
    areaScopeLsa_.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", areaScopeLsa_.RawData}
    return &(areaScopeLsa_.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa
// OSPFv2 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_Header

    // Decoded OSPFv2 LSA body data.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_LsaBody
}

func (ospfv2Lsa *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa) GetEntityData() *types.CommonEntityData {
    ospfv2Lsa.EntityData.YFilter = ospfv2Lsa.YFilter
    ospfv2Lsa.EntityData.YangName = "ospfv2-lsa"
    ospfv2Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Lsa.EntityData.SegmentPath = "ospfv2-lsa"
    ospfv2Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Lsa.EntityData.Children = make(map[string]types.YChild)
    ospfv2Lsa.EntityData.Children["header"] = types.YChild{"Header", &ospfv2Lsa.Header}
    ospfv2Lsa.EntityData.Children["lsa-body"] = types.YChild{"LsaBody", &ospfv2Lsa.LsaBody}
    ospfv2Lsa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_Header
// Decoded OSPFv2 LSA header data
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255.
    OpaqueType interface{}

    // Opaque ID. The type is interface{} with range: 0..4294967295.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type_ interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}

    // LSA options. The type is map[string]bool.
    FlagOptions interface{}
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv2-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["opaque-type"] = types.YLeaf{"OpaqueType", header.OpaqueType}
    header.EntityData.Leafs["opaque-id"] = types.YLeaf{"OpaqueId", header.OpaqueId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    header.EntityData.Leafs["flag-options"] = types.YLeaf{"FlagOptions", header.FlagOptions}
    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_LsaBody
// Decoded OSPFv2 LSA body data
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Summary mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryMask interface{}

    // External mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ExternalMask interface{}

    // LSA body flags. The type is map[string]bool.
    BodyFlagOptions interface{}

    // Network details.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_LsaBody_Network
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv2-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = make(map[string]types.YChild)
    lsaBody.EntityData.Children["network"] = types.YChild{"Network", &lsaBody.Network}
    lsaBody.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaBody.EntityData.Leafs["num-of-links"] = types.YLeaf{"NumOfLinks", lsaBody.NumOfLinks}
    lsaBody.EntityData.Leafs["summary-mask"] = types.YLeaf{"SummaryMask", lsaBody.SummaryMask}
    lsaBody.EntityData.Leafs["external-mask"] = types.YLeaf{"ExternalMask", lsaBody.ExternalMask}
    lsaBody.EntityData.Leafs["body-flag-options"] = types.YLeaf{"BodyFlagOptions", lsaBody.BodyFlagOptions}
    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_LsaBody_Network
// Network details
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP network mask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", network.NetworkMask}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Link
// Router LSA link
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is interface{} with range:
    // 0..4294967295.
    LinkId interface{}

    // This attribute is a key. Link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Topology specific information. The type is slice of
    // OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Link_Ospfv2Topology.
    Ospfv2Topology []OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Link_Ospfv2Topology
}

func (ospfv2Link *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Link) GetEntityData() *types.CommonEntityData {
    ospfv2Link.EntityData.YFilter = ospfv2Link.YFilter
    ospfv2Link.EntityData.YangName = "ospfv2-link"
    ospfv2Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Link.EntityData.SegmentPath = "ospfv2-link" + "[link-id='" + fmt.Sprintf("%v", ospfv2Link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", ospfv2Link.LinkData) + "']"
    ospfv2Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Link.EntityData.Children = make(map[string]types.YChild)
    ospfv2Link.EntityData.Children["ospfv2-topology"] = types.YChild{"Ospfv2Topology", nil}
    for i := range ospfv2Link.Ospfv2Topology {
        ospfv2Link.EntityData.Children[types.GetSegmentPath(&ospfv2Link.Ospfv2Topology[i])] = types.YChild{"Ospfv2Topology", &ospfv2Link.Ospfv2Topology[i]}
    }
    ospfv2Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Link.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", ospfv2Link.LinkId}
    ospfv2Link.EntityData.Leafs["link-data"] = types.YLeaf{"LinkData", ospfv2Link.LinkData}
    ospfv2Link.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv2Link.Type_}
    return &(ospfv2Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Link_Ospfv2Topology
// Topology specific information
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Link_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Link_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "ospfv2-link"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + "[mt-id='" + fmt.Sprintf("%v", ospfv2Topology.MtId) + "']"
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = make(map[string]types.YChild)
    ospfv2Topology.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2Topology.MtId}
    ospfv2Topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2Topology.Metric}
    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Topology
// Summary LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + "[mt-id='" + fmt.Sprintf("%v", ospfv2Topology.MtId) + "']"
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = make(map[string]types.YChild)
    ospfv2Topology.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2Topology.MtId}
    ospfv2Topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2Topology.Metric}
    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2External
// External LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled on the link. The type
    // is interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (ospfv2External *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv2External) GetEntityData() *types.CommonEntityData {
    ospfv2External.EntityData.YFilter = ospfv2External.YFilter
    ospfv2External.EntityData.YangName = "ospfv2-external"
    ospfv2External.EntityData.BundleName = "cisco_ios_xe"
    ospfv2External.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2External.EntityData.SegmentPath = "ospfv2-external" + "[mt-id='" + fmt.Sprintf("%v", ospfv2External.MtId) + "']"
    ospfv2External.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2External.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2External.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2External.EntityData.Children = make(map[string]types.YChild)
    ospfv2External.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2External.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2External.MtId}
    ospfv2External.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2External.Metric}
    ospfv2External.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", ospfv2External.ForwardingAddress}
    ospfv2External.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", ospfv2External.ExternalRouteTag}
    return &(ospfv2External.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa
// OSPFv3 LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header.
    Header OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_Header

    // Decoded OSPFv3 LSA body.
    LsaBody OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody
}

func (ospfv3Lsa *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa) GetEntityData() *types.CommonEntityData {
    ospfv3Lsa.EntityData.YFilter = ospfv3Lsa.YFilter
    ospfv3Lsa.EntityData.YangName = "ospfv3-lsa"
    ospfv3Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Lsa.EntityData.SegmentPath = "ospfv3-lsa"
    ospfv3Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Lsa.EntityData.Children = make(map[string]types.YChild)
    ospfv3Lsa.EntityData.Children["header"] = types.YChild{"Header", &ospfv3Lsa.Header}
    ospfv3Lsa.EntityData.Children["lsa-body"] = types.YChild{"LsaBody", &ospfv3Lsa.LsaBody}
    ospfv3Lsa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_Header
// Decoded OSPFv3 LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // OSPFv3 LSA options. The type is map[string]bool.
    LsaHdrOptions interface{}

    // LSA header.
    LsaHeader OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_Header_LsaHeader
}

func (header *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv3-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Children["lsa-header"] = types.YChild{"LsaHeader", &header.LsaHeader}
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["lsa-hdr-options"] = types.YLeaf{"LsaHdrOptions", header.LsaHdrOptions}
    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_Header_LsaHeader
// LSA header
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_Header_LsaHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type_ interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}
}

func (lsaHeader *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_Header_LsaHeader) GetEntityData() *types.CommonEntityData {
    lsaHeader.EntityData.YFilter = lsaHeader.YFilter
    lsaHeader.EntityData.YangName = "lsa-header"
    lsaHeader.EntityData.BundleName = "cisco_ios_xe"
    lsaHeader.EntityData.ParentYangName = "header"
    lsaHeader.EntityData.SegmentPath = "lsa-header"
    lsaHeader.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaHeader.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaHeader.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaHeader.EntityData.Children = make(map[string]types.YChild)
    lsaHeader.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaHeader.EntityData.Leafs["age"] = types.YLeaf{"Age", lsaHeader.Age}
    lsaHeader.EntityData.Leafs["type"] = types.YLeaf{"Type_", lsaHeader.Type_}
    lsaHeader.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", lsaHeader.AdvRouter}
    lsaHeader.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", lsaHeader.SeqNum}
    lsaHeader.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", lsaHeader.Checksum}
    lsaHeader.EntityData.Leafs["length"] = types.YLeaf{"Length", lsaHeader.Length}
    return &(lsaHeader.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody
// Decoded OSPFv3 LSA body
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA options. The type is map[string]bool.
    LsaFlagOptions interface{}

    // LSA Body Flags. The type is map[string]bool.
    LsaBodyFlags interface{}

    // OSPFv3 network.
    Network OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Network

    // OSPFv3 inter area prefix.
    Prefix OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Prefix

    // OSPFv3 inter area router.
    IaRouter OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_IaRouter

    // OSPFv3 LSA external.
    LsaExternal OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LsaExternal

    // OSPFv3 NSSA.
    Nssa OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa

    // OSPFv3 Link data.
    LinkData OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LinkData

    // OSPFv3 Intra area prefixes.
    IaPrefix OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_IaPrefix
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv3-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = make(map[string]types.YChild)
    lsaBody.EntityData.Children["network"] = types.YChild{"Network", &lsaBody.Network}
    lsaBody.EntityData.Children["prefix"] = types.YChild{"Prefix", &lsaBody.Prefix}
    lsaBody.EntityData.Children["ia-router"] = types.YChild{"IaRouter", &lsaBody.IaRouter}
    lsaBody.EntityData.Children["lsa-external"] = types.YChild{"LsaExternal", &lsaBody.LsaExternal}
    lsaBody.EntityData.Children["nssa"] = types.YChild{"Nssa", &lsaBody.Nssa}
    lsaBody.EntityData.Children["link-data"] = types.YChild{"LinkData", &lsaBody.LinkData}
    lsaBody.EntityData.Children["ia-prefix"] = types.YChild{"IaPrefix", &lsaBody.IaPrefix}
    lsaBody.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaBody.EntityData.Leafs["lsa-flag-options"] = types.YLeaf{"LsaFlagOptions", lsaBody.LsaFlagOptions}
    lsaBody.EntityData.Leafs["lsa-body-flags"] = types.YLeaf{"LsaBodyFlags", lsaBody.LsaBodyFlags}
    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Network
// OSPFv3 network
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}

    // Network LSA options. The type is map[string]bool.
    LsaNetOptions interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    network.EntityData.Leafs["lsa-net-options"] = types.YLeaf{"LsaNetOptions", network.LsaNetOptions}
    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Prefix
// OSPFv3 inter area prefix
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Inter area Prefix. The type is string.
    IaPrefix interface{}

    // Inter area prefix options. The type is string.
    IaPrefixOptions interface{}
}

func (prefix *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xe"
    prefix.EntityData.ParentYangName = "lsa-body"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["metric"] = types.YLeaf{"Metric", prefix.Metric}
    prefix.EntityData.Leafs["ia-prefix"] = types.YLeaf{"IaPrefix", prefix.IaPrefix}
    prefix.EntityData.Leafs["ia-prefix-options"] = types.YLeaf{"IaPrefixOptions", prefix.IaPrefixOptions}
    return &(prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_IaRouter
// OSPFv3 inter area router
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_IaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Router ID of the router being described by the LSA. The type is interface{}
    // with range: 0..4294967295.
    DestinationRouterId interface{}

    // Inter area LSA options. The type is map[string]bool.
    LsaIaOptions interface{}
}

func (iaRouter *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_IaRouter) GetEntityData() *types.CommonEntityData {
    iaRouter.EntityData.YFilter = iaRouter.YFilter
    iaRouter.EntityData.YangName = "ia-router"
    iaRouter.EntityData.BundleName = "cisco_ios_xe"
    iaRouter.EntityData.ParentYangName = "lsa-body"
    iaRouter.EntityData.SegmentPath = "ia-router"
    iaRouter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaRouter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaRouter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaRouter.EntityData.Children = make(map[string]types.YChild)
    iaRouter.EntityData.Leafs = make(map[string]types.YLeaf)
    iaRouter.EntityData.Leafs["metric"] = types.YLeaf{"Metric", iaRouter.Metric}
    iaRouter.EntityData.Leafs["destination-router-id"] = types.YLeaf{"DestinationRouterId", iaRouter.DestinationRouterId}
    iaRouter.EntityData.Leafs["lsa-ia-options"] = types.YLeaf{"LsaIaOptions", iaRouter.LsaIaOptions}
    return &(iaRouter.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LsaExternal
// OSPFv3 LSA external
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LsaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LsaExternal_Flags
}

func (lsaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LsaExternal) GetEntityData() *types.CommonEntityData {
    lsaExternal.EntityData.YFilter = lsaExternal.YFilter
    lsaExternal.EntityData.YangName = "lsa-external"
    lsaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaExternal.EntityData.ParentYangName = "lsa-body"
    lsaExternal.EntityData.SegmentPath = "lsa-external"
    lsaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaExternal.EntityData.Children = make(map[string]types.YChild)
    lsaExternal.EntityData.Children["flags"] = types.YChild{"Flags", &lsaExternal.Flags}
    lsaExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", lsaExternal.Metric}
    lsaExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", lsaExternal.ReferencedLsType}
    lsaExternal.EntityData.Leafs["external-prefix"] = types.YLeaf{"ExternalPrefix", lsaExternal.ExternalPrefix}
    lsaExternal.EntityData.Leafs["external-prefix-options"] = types.YLeaf{"ExternalPrefixOptions", lsaExternal.ExternalPrefixOptions}
    lsaExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", lsaExternal.ForwardingAddress}
    lsaExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", lsaExternal.ExternalRouteTag}
    lsaExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", lsaExternal.ReferencedLinkStateId}
    return &(lsaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LsaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LsaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LsaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    flags.EntityData.Leafs["e-flag"] = types.YLeaf{"EFlag", flags.EFlag}
    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa
// OSPFv3 NSSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA LSA.
    LsaNssaExternal OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
}

func (nssa *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xe"
    nssa.EntityData.ParentYangName = "lsa-body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssa.EntityData.Children = make(map[string]types.YChild)
    nssa.EntityData.Children["lsa-nssa-external"] = types.YChild{"LsaNssaExternal", &nssa.LsaNssaExternal}
    nssa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nssa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
// NSSA LSA
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
}

func (lsaNssaExternal *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal) GetEntityData() *types.CommonEntityData {
    lsaNssaExternal.EntityData.YFilter = lsaNssaExternal.YFilter
    lsaNssaExternal.EntityData.YangName = "lsa-nssa-external"
    lsaNssaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaNssaExternal.EntityData.ParentYangName = "nssa"
    lsaNssaExternal.EntityData.SegmentPath = "lsa-nssa-external"
    lsaNssaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaNssaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaNssaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaNssaExternal.EntityData.Children = make(map[string]types.YChild)
    lsaNssaExternal.EntityData.Children["flags"] = types.YChild{"Flags", &lsaNssaExternal.Flags}
    lsaNssaExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaNssaExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", lsaNssaExternal.Metric}
    lsaNssaExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", lsaNssaExternal.ReferencedLsType}
    lsaNssaExternal.EntityData.Leafs["external-prefix"] = types.YLeaf{"ExternalPrefix", lsaNssaExternal.ExternalPrefix}
    lsaNssaExternal.EntityData.Leafs["external-prefix-options"] = types.YLeaf{"ExternalPrefixOptions", lsaNssaExternal.ExternalPrefixOptions}
    lsaNssaExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", lsaNssaExternal.ForwardingAddress}
    lsaNssaExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", lsaNssaExternal.ExternalRouteTag}
    lsaNssaExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", lsaNssaExternal.ReferencedLinkStateId}
    return &(lsaNssaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-nssa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    flags.EntityData.Leafs["e-flag"] = types.YLeaf{"EFlag", flags.EFlag}
    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LinkData
// OSPFv3 Link data
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router priority of the interce. The type is interface{} with range: 0..255.
    RtrPriority interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // Link data LSA options. The type is map[string]bool.
    LsaIdOptions interface{}
}

func (linkData *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xe"
    linkData.EntityData.ParentYangName = "lsa-body"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkData.EntityData.Children = make(map[string]types.YChild)
    linkData.EntityData.Leafs = make(map[string]types.YLeaf)
    linkData.EntityData.Leafs["rtr-priority"] = types.YLeaf{"RtrPriority", linkData.RtrPriority}
    linkData.EntityData.Leafs["link-local-interface-address"] = types.YLeaf{"LinkLocalInterfaceAddress", linkData.LinkLocalInterfaceAddress}
    linkData.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", linkData.NumOfPrefixes}
    linkData.EntityData.Leafs["lsa-id-options"] = types.YLeaf{"LsaIdOptions", linkData.LsaIdOptions}
    return &(linkData.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_IaPrefix
// OSPFv3 Intra area prefixes
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}
}

func (iaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Lsa_LsaBody_IaPrefix) GetEntityData() *types.CommonEntityData {
    iaPrefix.EntityData.YFilter = iaPrefix.YFilter
    iaPrefix.EntityData.YangName = "ia-prefix"
    iaPrefix.EntityData.BundleName = "cisco_ios_xe"
    iaPrefix.EntityData.ParentYangName = "lsa-body"
    iaPrefix.EntityData.SegmentPath = "ia-prefix"
    iaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaPrefix.EntityData.Children = make(map[string]types.YChild)
    iaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    iaPrefix.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", iaPrefix.ReferencedLsType}
    iaPrefix.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", iaPrefix.ReferencedLinkStateId}
    iaPrefix.EntityData.Leafs["referenced-adv-router"] = types.YLeaf{"ReferencedAdvRouter", iaPrefix.ReferencedAdvRouter}
    iaPrefix.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", iaPrefix.NumOfPrefixes}
    return &(iaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Link
// OSPFv3 links
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor router ID. The type is interface{} with
    // range: 0..4294967295.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv3Link *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Link) GetEntityData() *types.CommonEntityData {
    ospfv3Link.EntityData.YFilter = ospfv3Link.YFilter
    ospfv3Link.EntityData.YangName = "ospfv3-link"
    ospfv3Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Link.EntityData.SegmentPath = "ospfv3-link" + "[interface-id='" + fmt.Sprintf("%v", ospfv3Link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", ospfv3Link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", ospfv3Link.NeighborRouterId) + "']"
    ospfv3Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Link.EntityData.Children = make(map[string]types.YChild)
    ospfv3Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3Link.EntityData.Leafs["interface-id"] = types.YLeaf{"InterfaceId", ospfv3Link.InterfaceId}
    ospfv3Link.EntityData.Leafs["neighbor-interface-id"] = types.YLeaf{"NeighborInterfaceId", ospfv3Link.NeighborInterfaceId}
    ospfv3Link.EntityData.Leafs["neighbor-router-id"] = types.YLeaf{"NeighborRouterId", ospfv3Link.NeighborRouterId}
    ospfv3Link.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv3Link.Type_}
    ospfv3Link.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv3Link.Metric}
    return &(ospfv3Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Prefix
// OSPFv3 prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3Prefix *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3Prefix) GetEntityData() *types.CommonEntityData {
    ospfv3Prefix.EntityData.YFilter = ospfv3Prefix.YFilter
    ospfv3Prefix.EntityData.YangName = "ospfv3-prefix"
    ospfv3Prefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Prefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Prefix.EntityData.SegmentPath = "ospfv3-prefix" + "[prefix='" + fmt.Sprintf("%v", ospfv3Prefix.Prefix) + "']"
    ospfv3Prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Prefix.EntityData.Children = make(map[string]types.YChild)
    ospfv3Prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3Prefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ospfv3Prefix.Prefix}
    ospfv3Prefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", ospfv3Prefix.PrefixOptions}
    return &(ospfv3Prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3IaPrefix
// OSPFv3 intra-area prefix-list
type OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3IaPrefix *OspfOperData_OspfState_OspfInstance_OspfArea_AreaScopeLsa_AreaScopeLsa__Ospfv3IaPrefix) GetEntityData() *types.CommonEntityData {
    ospfv3IaPrefix.EntityData.YFilter = ospfv3IaPrefix.YFilter
    ospfv3IaPrefix.EntityData.YangName = "ospfv3-ia-prefix"
    ospfv3IaPrefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3IaPrefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3IaPrefix.EntityData.SegmentPath = "ospfv3-ia-prefix" + "[prefix='" + fmt.Sprintf("%v", ospfv3IaPrefix.Prefix) + "']"
    ospfv3IaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3IaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3IaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3IaPrefix.EntityData.Children = make(map[string]types.YChild)
    ospfv3IaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3IaPrefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ospfv3IaPrefix.Prefix}
    ospfv3IaPrefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", ospfv3IaPrefix.PrefixOptions}
    return &(ospfv3IaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas
// List OSPF link scope LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF link scope LSA type. The type is interface{}
    // with range: 0..4294967295.
    LsaType interface{}

    // List of OSPF link scope LSAs. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa.
    LinkScopeLsa []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa

    // List OSPF area scope LSA databases. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa.
    AreaScopeLsa []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa
}

func (linkScopeLsas *OspfOperData_OspfState_OspfInstance_LinkScopeLsas) GetEntityData() *types.CommonEntityData {
    linkScopeLsas.EntityData.YFilter = linkScopeLsas.YFilter
    linkScopeLsas.EntityData.YangName = "link-scope-lsas"
    linkScopeLsas.EntityData.BundleName = "cisco_ios_xe"
    linkScopeLsas.EntityData.ParentYangName = "ospf-instance"
    linkScopeLsas.EntityData.SegmentPath = "link-scope-lsas" + "[lsa-type='" + fmt.Sprintf("%v", linkScopeLsas.LsaType) + "']"
    linkScopeLsas.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkScopeLsas.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkScopeLsas.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkScopeLsas.EntityData.Children = make(map[string]types.YChild)
    linkScopeLsas.EntityData.Children["link-scope-lsa"] = types.YChild{"LinkScopeLsa", nil}
    for i := range linkScopeLsas.LinkScopeLsa {
        linkScopeLsas.EntityData.Children[types.GetSegmentPath(&linkScopeLsas.LinkScopeLsa[i])] = types.YChild{"LinkScopeLsa", &linkScopeLsas.LinkScopeLsa[i]}
    }
    linkScopeLsas.EntityData.Children["area-scope-lsa"] = types.YChild{"AreaScopeLsa", nil}
    for i := range linkScopeLsas.AreaScopeLsa {
        linkScopeLsas.EntityData.Children[types.GetSegmentPath(&linkScopeLsas.AreaScopeLsa[i])] = types.YChild{"AreaScopeLsa", &linkScopeLsas.AreaScopeLsa[i]}
    }
    linkScopeLsas.EntityData.Leafs = make(map[string]types.YLeaf)
    linkScopeLsas.EntityData.Leafs["lsa-type"] = types.YLeaf{"LsaType", linkScopeLsas.LsaType}
    return &(linkScopeLsas.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa
// List of OSPF link scope LSAs
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSA ID. The type is interface{} with range:
    // 0..4294967295.
    LsaId interface{}

    // This attribute is a key. Advertising router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is slice of interface{} with range: 0..255.
    RawData []interface{}

    // Version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Router address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RouterAddress interface{}

    // OSPFv2 LSA.
    Ospfv2Lsa OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa

    // OSPFv2 LSA link. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link.
    Ospfv2Link []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link

    // Summary LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Topology.
    Ospfv2Topology []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Topology

    // External LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2External.
    Ospfv2External []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2External

    // OSPFv2 Unknown TLV. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv.
    Ospfv2UnknownTlv []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv

    // OSPFv3 LSA.
    Ospfv3LsaVal OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal

    // OSPFv3 links. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3Link.
    Ospfv3Link []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3Link

    // OSPFv3 prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList.
    Ospfv3PrefixList []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList

    // OSPFv3 intra-area prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix.
    Ospfv3IaPrefix []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix

    // OSPF multi-topology interface augmentation. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_MultiTopology.
    MultiTopology []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_MultiTopology

    // Link TLV.
    Tlv OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Tlv

    // OSPFv2 Unknown sub TLV. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_UnknownSubTlv.
    UnknownSubTlv []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_UnknownSubTlv
}

func (linkScopeLsa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa) GetEntityData() *types.CommonEntityData {
    linkScopeLsa.EntityData.YFilter = linkScopeLsa.YFilter
    linkScopeLsa.EntityData.YangName = "link-scope-lsa"
    linkScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    linkScopeLsa.EntityData.ParentYangName = "link-scope-lsas"
    linkScopeLsa.EntityData.SegmentPath = "link-scope-lsa" + "[lsa-id='" + fmt.Sprintf("%v", linkScopeLsa.LsaId) + "']" + "[adv-router='" + fmt.Sprintf("%v", linkScopeLsa.AdvRouter) + "']"
    linkScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkScopeLsa.EntityData.Children = make(map[string]types.YChild)
    linkScopeLsa.EntityData.Children["ospfv2-lsa"] = types.YChild{"Ospfv2Lsa", &linkScopeLsa.Ospfv2Lsa}
    linkScopeLsa.EntityData.Children["ospfv2-link"] = types.YChild{"Ospfv2Link", nil}
    for i := range linkScopeLsa.Ospfv2Link {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv2Link[i])] = types.YChild{"Ospfv2Link", &linkScopeLsa.Ospfv2Link[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv2-topology"] = types.YChild{"Ospfv2Topology", nil}
    for i := range linkScopeLsa.Ospfv2Topology {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv2Topology[i])] = types.YChild{"Ospfv2Topology", &linkScopeLsa.Ospfv2Topology[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv2-external"] = types.YChild{"Ospfv2External", nil}
    for i := range linkScopeLsa.Ospfv2External {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv2External[i])] = types.YChild{"Ospfv2External", &linkScopeLsa.Ospfv2External[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv2-unknown-tlv"] = types.YChild{"Ospfv2UnknownTlv", nil}
    for i := range linkScopeLsa.Ospfv2UnknownTlv {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv2UnknownTlv[i])] = types.YChild{"Ospfv2UnknownTlv", &linkScopeLsa.Ospfv2UnknownTlv[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv3-lsa-val"] = types.YChild{"Ospfv3LsaVal", &linkScopeLsa.Ospfv3LsaVal}
    linkScopeLsa.EntityData.Children["ospfv3-link"] = types.YChild{"Ospfv3Link", nil}
    for i := range linkScopeLsa.Ospfv3Link {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv3Link[i])] = types.YChild{"Ospfv3Link", &linkScopeLsa.Ospfv3Link[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv3-prefix-list"] = types.YChild{"Ospfv3PrefixList", nil}
    for i := range linkScopeLsa.Ospfv3PrefixList {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv3PrefixList[i])] = types.YChild{"Ospfv3PrefixList", &linkScopeLsa.Ospfv3PrefixList[i]}
    }
    linkScopeLsa.EntityData.Children["ospfv3-ia-prefix"] = types.YChild{"Ospfv3IaPrefix", nil}
    for i := range linkScopeLsa.Ospfv3IaPrefix {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.Ospfv3IaPrefix[i])] = types.YChild{"Ospfv3IaPrefix", &linkScopeLsa.Ospfv3IaPrefix[i]}
    }
    linkScopeLsa.EntityData.Children["multi-topology"] = types.YChild{"MultiTopology", nil}
    for i := range linkScopeLsa.MultiTopology {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.MultiTopology[i])] = types.YChild{"MultiTopology", &linkScopeLsa.MultiTopology[i]}
    }
    linkScopeLsa.EntityData.Children["tlv"] = types.YChild{"Tlv", &linkScopeLsa.Tlv}
    linkScopeLsa.EntityData.Children["unknown-sub-tlv"] = types.YChild{"UnknownSubTlv", nil}
    for i := range linkScopeLsa.UnknownSubTlv {
        linkScopeLsa.EntityData.Children[types.GetSegmentPath(&linkScopeLsa.UnknownSubTlv[i])] = types.YChild{"UnknownSubTlv", &linkScopeLsa.UnknownSubTlv[i]}
    }
    linkScopeLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    linkScopeLsa.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", linkScopeLsa.LsaId}
    linkScopeLsa.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", linkScopeLsa.AdvRouter}
    linkScopeLsa.EntityData.Leafs["decoded-completed"] = types.YLeaf{"DecodedCompleted", linkScopeLsa.DecodedCompleted}
    linkScopeLsa.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", linkScopeLsa.RawData}
    linkScopeLsa.EntityData.Leafs["version"] = types.YLeaf{"Version", linkScopeLsa.Version}
    linkScopeLsa.EntityData.Leafs["router-address"] = types.YLeaf{"RouterAddress", linkScopeLsa.RouterAddress}
    return &(linkScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa
// OSPFv2 LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header

    // Decoded OSPFv2 LSA body data.
    LsaBody OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody
}

func (ospfv2Lsa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa) GetEntityData() *types.CommonEntityData {
    ospfv2Lsa.EntityData.YFilter = ospfv2Lsa.YFilter
    ospfv2Lsa.EntityData.YangName = "ospfv2-lsa"
    ospfv2Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Lsa.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Lsa.EntityData.SegmentPath = "ospfv2-lsa"
    ospfv2Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Lsa.EntityData.Children = make(map[string]types.YChild)
    ospfv2Lsa.EntityData.Children["header"] = types.YChild{"Header", &ospfv2Lsa.Header}
    ospfv2Lsa.EntityData.Children["lsa-body"] = types.YChild{"LsaBody", &ospfv2Lsa.LsaBody}
    ospfv2Lsa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header
// Decoded OSPFv2 LSA header data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255.
    OpaqueType interface{}

    // Opaque ID. The type is interface{} with range: 0..4294967295.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type_ interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}

    // LSA options. The type is map[string]bool.
    FlagOptions interface{}
}

func (header *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv2-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["opaque-type"] = types.YLeaf{"OpaqueType", header.OpaqueType}
    header.EntityData.Leafs["opaque-id"] = types.YLeaf{"OpaqueId", header.OpaqueId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    header.EntityData.Leafs["flag-options"] = types.YLeaf{"FlagOptions", header.FlagOptions}
    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody
// Decoded OSPFv2 LSA body data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Summary mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryMask interface{}

    // External mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ExternalMask interface{}

    // LSA body flags. The type is map[string]bool.
    BodyFlagOptions interface{}

    // Network details.
    Network OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv2-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = make(map[string]types.YChild)
    lsaBody.EntityData.Children["network"] = types.YChild{"Network", &lsaBody.Network}
    lsaBody.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaBody.EntityData.Leafs["num-of-links"] = types.YLeaf{"NumOfLinks", lsaBody.NumOfLinks}
    lsaBody.EntityData.Leafs["summary-mask"] = types.YLeaf{"SummaryMask", lsaBody.SummaryMask}
    lsaBody.EntityData.Leafs["external-mask"] = types.YLeaf{"ExternalMask", lsaBody.ExternalMask}
    lsaBody.EntityData.Leafs["body-flag-options"] = types.YLeaf{"BodyFlagOptions", lsaBody.BodyFlagOptions}
    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network
// Network details
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP network mask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", network.NetworkMask}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link
// OSPFv2 LSA link
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is interface{} with range:
    // 0..4294967295.
    LinkId interface{}

    // This attribute is a key. Link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Topology specific information. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology.
    Ospfv2Topology []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology
}

func (ospfv2Link *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link) GetEntityData() *types.CommonEntityData {
    ospfv2Link.EntityData.YFilter = ospfv2Link.YFilter
    ospfv2Link.EntityData.YangName = "ospfv2-link"
    ospfv2Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Link.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Link.EntityData.SegmentPath = "ospfv2-link" + "[link-id='" + fmt.Sprintf("%v", ospfv2Link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", ospfv2Link.LinkData) + "']"
    ospfv2Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Link.EntityData.Children = make(map[string]types.YChild)
    ospfv2Link.EntityData.Children["ospfv2-topology"] = types.YChild{"Ospfv2Topology", nil}
    for i := range ospfv2Link.Ospfv2Topology {
        ospfv2Link.EntityData.Children[types.GetSegmentPath(&ospfv2Link.Ospfv2Topology[i])] = types.YChild{"Ospfv2Topology", &ospfv2Link.Ospfv2Topology[i]}
    }
    ospfv2Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Link.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", ospfv2Link.LinkId}
    ospfv2Link.EntityData.Leafs["link-data"] = types.YLeaf{"LinkData", ospfv2Link.LinkData}
    ospfv2Link.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv2Link.Type_}
    return &(ospfv2Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology
// Topology specific information
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Link_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "ospfv2-link"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + "[mt-id='" + fmt.Sprintf("%v", ospfv2Topology.MtId) + "']"
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = make(map[string]types.YChild)
    ospfv2Topology.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2Topology.MtId}
    ospfv2Topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2Topology.Metric}
    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Topology
// Summary LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + "[mt-id='" + fmt.Sprintf("%v", ospfv2Topology.MtId) + "']"
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = make(map[string]types.YChild)
    ospfv2Topology.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2Topology.MtId}
    ospfv2Topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2Topology.Metric}
    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2External
// External LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled on the link. The type
    // is interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (ospfv2External *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2External) GetEntityData() *types.CommonEntityData {
    ospfv2External.EntityData.YFilter = ospfv2External.YFilter
    ospfv2External.EntityData.YangName = "ospfv2-external"
    ospfv2External.EntityData.BundleName = "cisco_ios_xe"
    ospfv2External.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2External.EntityData.SegmentPath = "ospfv2-external" + "[mt-id='" + fmt.Sprintf("%v", ospfv2External.MtId) + "']"
    ospfv2External.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2External.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2External.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2External.EntityData.Children = make(map[string]types.YChild)
    ospfv2External.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2External.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2External.MtId}
    ospfv2External.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2External.Metric}
    ospfv2External.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", ospfv2External.ForwardingAddress}
    ospfv2External.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", ospfv2External.ExternalRouteTag}
    return &(ospfv2External.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv
// OSPFv2 Unknown TLV
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type_ interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (ospfv2UnknownTlv *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv2UnknownTlv) GetEntityData() *types.CommonEntityData {
    ospfv2UnknownTlv.EntityData.YFilter = ospfv2UnknownTlv.YFilter
    ospfv2UnknownTlv.EntityData.YangName = "ospfv2-unknown-tlv"
    ospfv2UnknownTlv.EntityData.BundleName = "cisco_ios_xe"
    ospfv2UnknownTlv.EntityData.ParentYangName = "link-scope-lsa"
    ospfv2UnknownTlv.EntityData.SegmentPath = "ospfv2-unknown-tlv" + "[type='" + fmt.Sprintf("%v", ospfv2UnknownTlv.Type_) + "']"
    ospfv2UnknownTlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2UnknownTlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2UnknownTlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2UnknownTlv.EntityData.Children = make(map[string]types.YChild)
    ospfv2UnknownTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2UnknownTlv.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv2UnknownTlv.Type_}
    ospfv2UnknownTlv.EntityData.Leafs["length"] = types.YLeaf{"Length", ospfv2UnknownTlv.Length}
    ospfv2UnknownTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", ospfv2UnknownTlv.Value}
    return &(ospfv2UnknownTlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal
// OSPFv3 LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header.
    Header OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header

    // Decoded OSPFv3 LSA body.
    LsaBody OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody
}

func (ospfv3LsaVal *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal) GetEntityData() *types.CommonEntityData {
    ospfv3LsaVal.EntityData.YFilter = ospfv3LsaVal.YFilter
    ospfv3LsaVal.EntityData.YangName = "ospfv3-lsa-val"
    ospfv3LsaVal.EntityData.BundleName = "cisco_ios_xe"
    ospfv3LsaVal.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3LsaVal.EntityData.SegmentPath = "ospfv3-lsa-val"
    ospfv3LsaVal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3LsaVal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3LsaVal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3LsaVal.EntityData.Children = make(map[string]types.YChild)
    ospfv3LsaVal.EntityData.Children["header"] = types.YChild{"Header", &ospfv3LsaVal.Header}
    ospfv3LsaVal.EntityData.Children["lsa-body"] = types.YChild{"LsaBody", &ospfv3LsaVal.LsaBody}
    ospfv3LsaVal.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3LsaVal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header
// Decoded OSPFv3 LSA header
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // OSPFv3 LSA options. The type is map[string]bool.
    LsaHdrOptions interface{}

    // LSA header.
    LsaHeader OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader
}

func (header *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv3-lsa-val"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Children["lsa-header"] = types.YChild{"LsaHeader", &header.LsaHeader}
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["lsa-hdr-options"] = types.YLeaf{"LsaHdrOptions", header.LsaHdrOptions}
    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader
// LSA header
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type_ interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}
}

func (lsaHeader *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_Header_LsaHeader) GetEntityData() *types.CommonEntityData {
    lsaHeader.EntityData.YFilter = lsaHeader.YFilter
    lsaHeader.EntityData.YangName = "lsa-header"
    lsaHeader.EntityData.BundleName = "cisco_ios_xe"
    lsaHeader.EntityData.ParentYangName = "header"
    lsaHeader.EntityData.SegmentPath = "lsa-header"
    lsaHeader.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaHeader.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaHeader.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaHeader.EntityData.Children = make(map[string]types.YChild)
    lsaHeader.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaHeader.EntityData.Leafs["age"] = types.YLeaf{"Age", lsaHeader.Age}
    lsaHeader.EntityData.Leafs["type"] = types.YLeaf{"Type_", lsaHeader.Type_}
    lsaHeader.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", lsaHeader.AdvRouter}
    lsaHeader.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", lsaHeader.SeqNum}
    lsaHeader.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", lsaHeader.Checksum}
    lsaHeader.EntityData.Leafs["length"] = types.YLeaf{"Length", lsaHeader.Length}
    return &(lsaHeader.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody
// Decoded OSPFv3 LSA body
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA options. The type is map[string]bool.
    LsaFlagOptions interface{}

    // LSA Body Flags. The type is map[string]bool.
    LsaBodyFlags interface{}

    // OSPFv3 network.
    Network OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network

    // OSPFv3 inter area prefix.
    Prefix OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix

    // OSPFv3 inter area router.
    IaRouter OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter

    // OSPFv3 LSA external.
    LsaExternal OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal

    // OSPFv3 NSSA.
    Nssa OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa

    // OSPFv3 Link data.
    LinkData OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData

    // OSPFv3 Intra area prefixes.
    IaPrefix OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv3-lsa-val"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = make(map[string]types.YChild)
    lsaBody.EntityData.Children["network"] = types.YChild{"Network", &lsaBody.Network}
    lsaBody.EntityData.Children["prefix"] = types.YChild{"Prefix", &lsaBody.Prefix}
    lsaBody.EntityData.Children["ia-router"] = types.YChild{"IaRouter", &lsaBody.IaRouter}
    lsaBody.EntityData.Children["lsa-external"] = types.YChild{"LsaExternal", &lsaBody.LsaExternal}
    lsaBody.EntityData.Children["nssa"] = types.YChild{"Nssa", &lsaBody.Nssa}
    lsaBody.EntityData.Children["link-data"] = types.YChild{"LinkData", &lsaBody.LinkData}
    lsaBody.EntityData.Children["ia-prefix"] = types.YChild{"IaPrefix", &lsaBody.IaPrefix}
    lsaBody.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaBody.EntityData.Leafs["lsa-flag-options"] = types.YLeaf{"LsaFlagOptions", lsaBody.LsaFlagOptions}
    lsaBody.EntityData.Leafs["lsa-body-flags"] = types.YLeaf{"LsaBodyFlags", lsaBody.LsaBodyFlags}
    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network
// OSPFv3 network
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}

    // Network LSA options. The type is map[string]bool.
    LsaNetOptions interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    network.EntityData.Leafs["lsa-net-options"] = types.YLeaf{"LsaNetOptions", network.LsaNetOptions}
    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix
// OSPFv3 inter area prefix
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Inter area Prefix. The type is string.
    IaPrefix interface{}

    // Inter area prefix options. The type is string.
    IaPrefixOptions interface{}
}

func (prefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xe"
    prefix.EntityData.ParentYangName = "lsa-body"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["metric"] = types.YLeaf{"Metric", prefix.Metric}
    prefix.EntityData.Leafs["ia-prefix"] = types.YLeaf{"IaPrefix", prefix.IaPrefix}
    prefix.EntityData.Leafs["ia-prefix-options"] = types.YLeaf{"IaPrefixOptions", prefix.IaPrefixOptions}
    return &(prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter
// OSPFv3 inter area router
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Router ID of the router being described by the LSA. The type is interface{}
    // with range: 0..4294967295.
    DestinationRouterId interface{}

    // Inter area LSA options. The type is map[string]bool.
    LsaIaOptions interface{}
}

func (iaRouter *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaRouter) GetEntityData() *types.CommonEntityData {
    iaRouter.EntityData.YFilter = iaRouter.YFilter
    iaRouter.EntityData.YangName = "ia-router"
    iaRouter.EntityData.BundleName = "cisco_ios_xe"
    iaRouter.EntityData.ParentYangName = "lsa-body"
    iaRouter.EntityData.SegmentPath = "ia-router"
    iaRouter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaRouter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaRouter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaRouter.EntityData.Children = make(map[string]types.YChild)
    iaRouter.EntityData.Leafs = make(map[string]types.YLeaf)
    iaRouter.EntityData.Leafs["metric"] = types.YLeaf{"Metric", iaRouter.Metric}
    iaRouter.EntityData.Leafs["destination-router-id"] = types.YLeaf{"DestinationRouterId", iaRouter.DestinationRouterId}
    iaRouter.EntityData.Leafs["lsa-ia-options"] = types.YLeaf{"LsaIaOptions", iaRouter.LsaIaOptions}
    return &(iaRouter.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal
// OSPFv3 LSA external
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags
}

func (lsaExternal *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal) GetEntityData() *types.CommonEntityData {
    lsaExternal.EntityData.YFilter = lsaExternal.YFilter
    lsaExternal.EntityData.YangName = "lsa-external"
    lsaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaExternal.EntityData.ParentYangName = "lsa-body"
    lsaExternal.EntityData.SegmentPath = "lsa-external"
    lsaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaExternal.EntityData.Children = make(map[string]types.YChild)
    lsaExternal.EntityData.Children["flags"] = types.YChild{"Flags", &lsaExternal.Flags}
    lsaExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", lsaExternal.Metric}
    lsaExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", lsaExternal.ReferencedLsType}
    lsaExternal.EntityData.Leafs["external-prefix"] = types.YLeaf{"ExternalPrefix", lsaExternal.ExternalPrefix}
    lsaExternal.EntityData.Leafs["external-prefix-options"] = types.YLeaf{"ExternalPrefixOptions", lsaExternal.ExternalPrefixOptions}
    lsaExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", lsaExternal.ForwardingAddress}
    lsaExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", lsaExternal.ExternalRouteTag}
    lsaExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", lsaExternal.ReferencedLinkStateId}
    return &(lsaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LsaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    flags.EntityData.Leafs["e-flag"] = types.YLeaf{"EFlag", flags.EFlag}
    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa
// OSPFv3 NSSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA LSA.
    LsaNssaExternal OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal
}

func (nssa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xe"
    nssa.EntityData.ParentYangName = "lsa-body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssa.EntityData.Children = make(map[string]types.YChild)
    nssa.EntityData.Children["lsa-nssa-external"] = types.YChild{"LsaNssaExternal", &nssa.LsaNssaExternal}
    nssa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nssa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal
// NSSA LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags
}

func (lsaNssaExternal *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal) GetEntityData() *types.CommonEntityData {
    lsaNssaExternal.EntityData.YFilter = lsaNssaExternal.YFilter
    lsaNssaExternal.EntityData.YangName = "lsa-nssa-external"
    lsaNssaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaNssaExternal.EntityData.ParentYangName = "nssa"
    lsaNssaExternal.EntityData.SegmentPath = "lsa-nssa-external"
    lsaNssaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaNssaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaNssaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaNssaExternal.EntityData.Children = make(map[string]types.YChild)
    lsaNssaExternal.EntityData.Children["flags"] = types.YChild{"Flags", &lsaNssaExternal.Flags}
    lsaNssaExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaNssaExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", lsaNssaExternal.Metric}
    lsaNssaExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", lsaNssaExternal.ReferencedLsType}
    lsaNssaExternal.EntityData.Leafs["external-prefix"] = types.YLeaf{"ExternalPrefix", lsaNssaExternal.ExternalPrefix}
    lsaNssaExternal.EntityData.Leafs["external-prefix-options"] = types.YLeaf{"ExternalPrefixOptions", lsaNssaExternal.ExternalPrefixOptions}
    lsaNssaExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", lsaNssaExternal.ForwardingAddress}
    lsaNssaExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", lsaNssaExternal.ExternalRouteTag}
    lsaNssaExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", lsaNssaExternal.ReferencedLinkStateId}
    return &(lsaNssaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_Nssa_LsaNssaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-nssa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    flags.EntityData.Leafs["e-flag"] = types.YLeaf{"EFlag", flags.EFlag}
    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData
// OSPFv3 Link data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router priority of the interce. The type is interface{} with range: 0..255.
    RtrPriority interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // Link data LSA options. The type is map[string]bool.
    LsaIdOptions interface{}
}

func (linkData *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xe"
    linkData.EntityData.ParentYangName = "lsa-body"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkData.EntityData.Children = make(map[string]types.YChild)
    linkData.EntityData.Leafs = make(map[string]types.YLeaf)
    linkData.EntityData.Leafs["rtr-priority"] = types.YLeaf{"RtrPriority", linkData.RtrPriority}
    linkData.EntityData.Leafs["link-local-interface-address"] = types.YLeaf{"LinkLocalInterfaceAddress", linkData.LinkLocalInterfaceAddress}
    linkData.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", linkData.NumOfPrefixes}
    linkData.EntityData.Leafs["lsa-id-options"] = types.YLeaf{"LsaIdOptions", linkData.LsaIdOptions}
    return &(linkData.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix
// OSPFv3 Intra area prefixes
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}
}

func (iaPrefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3LsaVal_LsaBody_IaPrefix) GetEntityData() *types.CommonEntityData {
    iaPrefix.EntityData.YFilter = iaPrefix.YFilter
    iaPrefix.EntityData.YangName = "ia-prefix"
    iaPrefix.EntityData.BundleName = "cisco_ios_xe"
    iaPrefix.EntityData.ParentYangName = "lsa-body"
    iaPrefix.EntityData.SegmentPath = "ia-prefix"
    iaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaPrefix.EntityData.Children = make(map[string]types.YChild)
    iaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    iaPrefix.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", iaPrefix.ReferencedLsType}
    iaPrefix.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", iaPrefix.ReferencedLinkStateId}
    iaPrefix.EntityData.Leafs["referenced-adv-router"] = types.YLeaf{"ReferencedAdvRouter", iaPrefix.ReferencedAdvRouter}
    iaPrefix.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", iaPrefix.NumOfPrefixes}
    return &(iaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3Link
// OSPFv3 links
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor router ID. The type is interface{} with
    // range: 0..4294967295.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv3Link *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3Link) GetEntityData() *types.CommonEntityData {
    ospfv3Link.EntityData.YFilter = ospfv3Link.YFilter
    ospfv3Link.EntityData.YangName = "ospfv3-link"
    ospfv3Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Link.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3Link.EntityData.SegmentPath = "ospfv3-link" + "[interface-id='" + fmt.Sprintf("%v", ospfv3Link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", ospfv3Link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", ospfv3Link.NeighborRouterId) + "']"
    ospfv3Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Link.EntityData.Children = make(map[string]types.YChild)
    ospfv3Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3Link.EntityData.Leafs["interface-id"] = types.YLeaf{"InterfaceId", ospfv3Link.InterfaceId}
    ospfv3Link.EntityData.Leafs["neighbor-interface-id"] = types.YLeaf{"NeighborInterfaceId", ospfv3Link.NeighborInterfaceId}
    ospfv3Link.EntityData.Leafs["neighbor-router-id"] = types.YLeaf{"NeighborRouterId", ospfv3Link.NeighborRouterId}
    ospfv3Link.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv3Link.Type_}
    ospfv3Link.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv3Link.Metric}
    return &(ospfv3Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList
// OSPFv3 prefix-list
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3PrefixList *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3PrefixList) GetEntityData() *types.CommonEntityData {
    ospfv3PrefixList.EntityData.YFilter = ospfv3PrefixList.YFilter
    ospfv3PrefixList.EntityData.YangName = "ospfv3-prefix-list"
    ospfv3PrefixList.EntityData.BundleName = "cisco_ios_xe"
    ospfv3PrefixList.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3PrefixList.EntityData.SegmentPath = "ospfv3-prefix-list" + "[prefix='" + fmt.Sprintf("%v", ospfv3PrefixList.Prefix) + "']"
    ospfv3PrefixList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3PrefixList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3PrefixList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3PrefixList.EntityData.Children = make(map[string]types.YChild)
    ospfv3PrefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3PrefixList.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ospfv3PrefixList.Prefix}
    ospfv3PrefixList.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", ospfv3PrefixList.PrefixOptions}
    return &(ospfv3PrefixList.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix
// OSPFv3 intra-area prefix-list
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3IaPrefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Ospfv3IaPrefix) GetEntityData() *types.CommonEntityData {
    ospfv3IaPrefix.EntityData.YFilter = ospfv3IaPrefix.YFilter
    ospfv3IaPrefix.EntityData.YangName = "ospfv3-ia-prefix"
    ospfv3IaPrefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3IaPrefix.EntityData.ParentYangName = "link-scope-lsa"
    ospfv3IaPrefix.EntityData.SegmentPath = "ospfv3-ia-prefix" + "[prefix='" + fmt.Sprintf("%v", ospfv3IaPrefix.Prefix) + "']"
    ospfv3IaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3IaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3IaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3IaPrefix.EntityData.Children = make(map[string]types.YChild)
    ospfv3IaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3IaPrefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ospfv3IaPrefix.Prefix}
    ospfv3IaPrefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", ospfv3IaPrefix.PrefixOptions}
    return &(ospfv3IaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_MultiTopology
// OSPF multi-topology interface augmentation
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_MultiTopology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string.
    Name interface{}
}

func (multiTopology *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_MultiTopology) GetEntityData() *types.CommonEntityData {
    multiTopology.EntityData.YFilter = multiTopology.YFilter
    multiTopology.EntityData.YangName = "multi-topology"
    multiTopology.EntityData.BundleName = "cisco_ios_xe"
    multiTopology.EntityData.ParentYangName = "link-scope-lsa"
    multiTopology.EntityData.SegmentPath = "multi-topology" + "[name='" + fmt.Sprintf("%v", multiTopology.Name) + "']"
    multiTopology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    multiTopology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    multiTopology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    multiTopology.EntityData.Children = make(map[string]types.YChild)
    multiTopology.EntityData.Leafs = make(map[string]types.YLeaf)
    multiTopology.EntityData.Leafs["name"] = types.YLeaf{"Name", multiTopology.Name}
    return &(multiTopology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Tlv
// Link TLV
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Tlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link type. The type is interface{} with range: 0..255.
    LinkType interface{}

    // Link ID. The type is interface{} with range: 0..4294967295.
    LinkId interface{}

    // List of local interface IPv4 addresses. The type is one of the following
    // types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalIfIpv4Addr []interface{}

    // List of remote interface IPv4 addresses. The type is one of the following
    // types: slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalRemoteIpv4Addr []interface{}

    // TE metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Maximum bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxBandwidth interface{}

    // Maximum reservable bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    MaxReservableBandwidth interface{}

    // Unrseerved bandwidth. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    UnreservedBandwidth interface{}

    // Administrative group/Resource class/Color. The type is interface{} with
    // range: 0..4294967295.
    AdminGroup interface{}
}

func (tlv *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_Tlv) GetEntityData() *types.CommonEntityData {
    tlv.EntityData.YFilter = tlv.YFilter
    tlv.EntityData.YangName = "tlv"
    tlv.EntityData.BundleName = "cisco_ios_xe"
    tlv.EntityData.ParentYangName = "link-scope-lsa"
    tlv.EntityData.SegmentPath = "tlv"
    tlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tlv.EntityData.Children = make(map[string]types.YChild)
    tlv.EntityData.Leafs = make(map[string]types.YLeaf)
    tlv.EntityData.Leafs["link-type"] = types.YLeaf{"LinkType", tlv.LinkType}
    tlv.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", tlv.LinkId}
    tlv.EntityData.Leafs["local-if-ipv4-addr"] = types.YLeaf{"LocalIfIpv4Addr", tlv.LocalIfIpv4Addr}
    tlv.EntityData.Leafs["local-remote-ipv4-addr"] = types.YLeaf{"LocalRemoteIpv4Addr", tlv.LocalRemoteIpv4Addr}
    tlv.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", tlv.TeMetric}
    tlv.EntityData.Leafs["max-bandwidth"] = types.YLeaf{"MaxBandwidth", tlv.MaxBandwidth}
    tlv.EntityData.Leafs["max-reservable-bandwidth"] = types.YLeaf{"MaxReservableBandwidth", tlv.MaxReservableBandwidth}
    tlv.EntityData.Leafs["unreserved-bandwidth"] = types.YLeaf{"UnreservedBandwidth", tlv.UnreservedBandwidth}
    tlv.EntityData.Leafs["admin-group"] = types.YLeaf{"AdminGroup", tlv.AdminGroup}
    return &(tlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_UnknownSubTlv
// OSPFv2 Unknown sub TLV
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_UnknownSubTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TLV type. The type is interface{} with range:
    // 0..65535.
    Type_ interface{}

    // TLV length. The type is interface{} with range: 0..65535.
    Length interface{}

    // TLV value. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (unknownSubTlv *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_LinkScopeLsa_UnknownSubTlv) GetEntityData() *types.CommonEntityData {
    unknownSubTlv.EntityData.YFilter = unknownSubTlv.YFilter
    unknownSubTlv.EntityData.YangName = "unknown-sub-tlv"
    unknownSubTlv.EntityData.BundleName = "cisco_ios_xe"
    unknownSubTlv.EntityData.ParentYangName = "link-scope-lsa"
    unknownSubTlv.EntityData.SegmentPath = "unknown-sub-tlv" + "[type='" + fmt.Sprintf("%v", unknownSubTlv.Type_) + "']"
    unknownSubTlv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    unknownSubTlv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    unknownSubTlv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    unknownSubTlv.EntityData.Children = make(map[string]types.YChild)
    unknownSubTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownSubTlv.EntityData.Leafs["type"] = types.YLeaf{"Type_", unknownSubTlv.Type_}
    unknownSubTlv.EntityData.Leafs["length"] = types.YLeaf{"Length", unknownSubTlv.Length}
    unknownSubTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownSubTlv.Value}
    return &(unknownSubTlv.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa
// List OSPF area scope LSA databases
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSA Type. The type is interface{} with range:
    // 0..4294967295.
    LsaType interface{}

    // This attribute is a key. Advertising router. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    AdvRouter interface{}

    // The OSPF LSA body is fully decoded. The type is bool.
    DecodedCompleted interface{}

    // The complete LSA in network byte order as received/sent over the wire. The
    // type is slice of interface{} with range: 0..255.
    RawData []interface{}

    // OSPFv2 LSA.
    Ospfv2Lsa OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa

    // Router LSA link. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link.
    Ospfv2Link []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link

    // Summary LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Topology.
    Ospfv2Topology []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Topology

    // External LSA. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2External.
    Ospfv2External []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2External

    // OSPFv3 LSA.
    Ospfv3Lsa OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa

    // OSPFv3 links. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Link.
    Ospfv3Link []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Link

    // OSPFv3 prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Prefix.
    Ospfv3Prefix []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Prefix

    // OSPFv3 intra-area prefix-list. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix.
    Ospfv3IaPrefix []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix
}

func (areaScopeLsa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa) GetEntityData() *types.CommonEntityData {
    areaScopeLsa.EntityData.YFilter = areaScopeLsa.YFilter
    areaScopeLsa.EntityData.YangName = "area-scope-lsa"
    areaScopeLsa.EntityData.BundleName = "cisco_ios_xe"
    areaScopeLsa.EntityData.ParentYangName = "link-scope-lsas"
    areaScopeLsa.EntityData.SegmentPath = "area-scope-lsa" + "[lsa-type='" + fmt.Sprintf("%v", areaScopeLsa.LsaType) + "']" + "[adv-router='" + fmt.Sprintf("%v", areaScopeLsa.AdvRouter) + "']"
    areaScopeLsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    areaScopeLsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    areaScopeLsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    areaScopeLsa.EntityData.Children = make(map[string]types.YChild)
    areaScopeLsa.EntityData.Children["ospfv2-lsa"] = types.YChild{"Ospfv2Lsa", &areaScopeLsa.Ospfv2Lsa}
    areaScopeLsa.EntityData.Children["ospfv2-link"] = types.YChild{"Ospfv2Link", nil}
    for i := range areaScopeLsa.Ospfv2Link {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv2Link[i])] = types.YChild{"Ospfv2Link", &areaScopeLsa.Ospfv2Link[i]}
    }
    areaScopeLsa.EntityData.Children["ospfv2-topology"] = types.YChild{"Ospfv2Topology", nil}
    for i := range areaScopeLsa.Ospfv2Topology {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv2Topology[i])] = types.YChild{"Ospfv2Topology", &areaScopeLsa.Ospfv2Topology[i]}
    }
    areaScopeLsa.EntityData.Children["ospfv2-external"] = types.YChild{"Ospfv2External", nil}
    for i := range areaScopeLsa.Ospfv2External {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv2External[i])] = types.YChild{"Ospfv2External", &areaScopeLsa.Ospfv2External[i]}
    }
    areaScopeLsa.EntityData.Children["ospfv3-lsa"] = types.YChild{"Ospfv3Lsa", &areaScopeLsa.Ospfv3Lsa}
    areaScopeLsa.EntityData.Children["ospfv3-link"] = types.YChild{"Ospfv3Link", nil}
    for i := range areaScopeLsa.Ospfv3Link {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv3Link[i])] = types.YChild{"Ospfv3Link", &areaScopeLsa.Ospfv3Link[i]}
    }
    areaScopeLsa.EntityData.Children["ospfv3-prefix"] = types.YChild{"Ospfv3Prefix", nil}
    for i := range areaScopeLsa.Ospfv3Prefix {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv3Prefix[i])] = types.YChild{"Ospfv3Prefix", &areaScopeLsa.Ospfv3Prefix[i]}
    }
    areaScopeLsa.EntityData.Children["ospfv3-ia-prefix"] = types.YChild{"Ospfv3IaPrefix", nil}
    for i := range areaScopeLsa.Ospfv3IaPrefix {
        areaScopeLsa.EntityData.Children[types.GetSegmentPath(&areaScopeLsa.Ospfv3IaPrefix[i])] = types.YChild{"Ospfv3IaPrefix", &areaScopeLsa.Ospfv3IaPrefix[i]}
    }
    areaScopeLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    areaScopeLsa.EntityData.Leafs["lsa-type"] = types.YLeaf{"LsaType", areaScopeLsa.LsaType}
    areaScopeLsa.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", areaScopeLsa.AdvRouter}
    areaScopeLsa.EntityData.Leafs["decoded-completed"] = types.YLeaf{"DecodedCompleted", areaScopeLsa.DecodedCompleted}
    areaScopeLsa.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", areaScopeLsa.RawData}
    return &(areaScopeLsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa
// OSPFv2 LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv2 LSA header data.
    Header OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header

    // Decoded OSPFv2 LSA body data.
    LsaBody OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody
}

func (ospfv2Lsa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa) GetEntityData() *types.CommonEntityData {
    ospfv2Lsa.EntityData.YFilter = ospfv2Lsa.YFilter
    ospfv2Lsa.EntityData.YangName = "ospfv2-lsa"
    ospfv2Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Lsa.EntityData.SegmentPath = "ospfv2-lsa"
    ospfv2Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Lsa.EntityData.Children = make(map[string]types.YChild)
    ospfv2Lsa.EntityData.Children["header"] = types.YChild{"Header", &ospfv2Lsa.Header}
    ospfv2Lsa.EntityData.Children["lsa-body"] = types.YChild{"LsaBody", &ospfv2Lsa.LsaBody}
    ospfv2Lsa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header
// Decoded OSPFv2 LSA header data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // Opaque type. The type is interface{} with range: 0..255.
    OpaqueType interface{}

    // Opaque ID. The type is interface{} with range: 0..4294967295.
    OpaqueId interface{}

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type_ interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}

    // LSA options. The type is map[string]bool.
    FlagOptions interface{}
}

func (header *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv2-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["opaque-type"] = types.YLeaf{"OpaqueType", header.OpaqueType}
    header.EntityData.Leafs["opaque-id"] = types.YLeaf{"OpaqueId", header.OpaqueId}
    header.EntityData.Leafs["age"] = types.YLeaf{"Age", header.Age}
    header.EntityData.Leafs["type"] = types.YLeaf{"Type_", header.Type_}
    header.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", header.AdvRouter}
    header.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", header.SeqNum}
    header.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", header.Checksum}
    header.EntityData.Leafs["length"] = types.YLeaf{"Length", header.Length}
    header.EntityData.Leafs["flag-options"] = types.YLeaf{"FlagOptions", header.FlagOptions}
    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody
// Decoded OSPFv2 LSA body data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of links. The type is interface{} with range: 0..65535.
    NumOfLinks interface{}

    // Summary mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SummaryMask interface{}

    // External mask. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ExternalMask interface{}

    // LSA body flags. The type is map[string]bool.
    BodyFlagOptions interface{}

    // Network details.
    Network OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv2-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = make(map[string]types.YChild)
    lsaBody.EntityData.Children["network"] = types.YChild{"Network", &lsaBody.Network}
    lsaBody.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaBody.EntityData.Leafs["num-of-links"] = types.YLeaf{"NumOfLinks", lsaBody.NumOfLinks}
    lsaBody.EntityData.Leafs["summary-mask"] = types.YLeaf{"SummaryMask", lsaBody.SummaryMask}
    lsaBody.EntityData.Leafs["external-mask"] = types.YLeaf{"ExternalMask", lsaBody.ExternalMask}
    lsaBody.EntityData.Leafs["body-flag-options"] = types.YLeaf{"BodyFlagOptions", lsaBody.BodyFlagOptions}
    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network
// Network details
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP network mask. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NetworkMask interface{}

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["network-mask"] = types.YLeaf{"NetworkMask", network.NetworkMask}
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link
// Router LSA link
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link ID. The type is interface{} with range:
    // 0..4294967295.
    LinkId interface{}

    // This attribute is a key. Link data. The type is interface{} with range:
    // 0..4294967295.
    LinkData interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Topology specific information. The type is slice of
    // OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology.
    Ospfv2Topology []OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology
}

func (ospfv2Link *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link) GetEntityData() *types.CommonEntityData {
    ospfv2Link.EntityData.YFilter = ospfv2Link.YFilter
    ospfv2Link.EntityData.YangName = "ospfv2-link"
    ospfv2Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Link.EntityData.SegmentPath = "ospfv2-link" + "[link-id='" + fmt.Sprintf("%v", ospfv2Link.LinkId) + "']" + "[link-data='" + fmt.Sprintf("%v", ospfv2Link.LinkData) + "']"
    ospfv2Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Link.EntityData.Children = make(map[string]types.YChild)
    ospfv2Link.EntityData.Children["ospfv2-topology"] = types.YChild{"Ospfv2Topology", nil}
    for i := range ospfv2Link.Ospfv2Topology {
        ospfv2Link.EntityData.Children[types.GetSegmentPath(&ospfv2Link.Ospfv2Topology[i])] = types.YChild{"Ospfv2Topology", &ospfv2Link.Ospfv2Topology[i]}
    }
    ospfv2Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Link.EntityData.Leafs["link-id"] = types.YLeaf{"LinkId", ospfv2Link.LinkId}
    ospfv2Link.EntityData.Leafs["link-data"] = types.YLeaf{"LinkData", ospfv2Link.LinkData}
    ospfv2Link.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv2Link.Type_}
    return &(ospfv2Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology
// Topology specific information
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Link_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "ospfv2-link"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + "[mt-id='" + fmt.Sprintf("%v", ospfv2Topology.MtId) + "']"
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = make(map[string]types.YChild)
    ospfv2Topology.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2Topology.MtId}
    ospfv2Topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2Topology.Metric}
    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Topology
// Summary LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Topology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled link. The type is
    // interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv2Topology *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2Topology) GetEntityData() *types.CommonEntityData {
    ospfv2Topology.EntityData.YFilter = ospfv2Topology.YFilter
    ospfv2Topology.EntityData.YangName = "ospfv2-topology"
    ospfv2Topology.EntityData.BundleName = "cisco_ios_xe"
    ospfv2Topology.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2Topology.EntityData.SegmentPath = "ospfv2-topology" + "[mt-id='" + fmt.Sprintf("%v", ospfv2Topology.MtId) + "']"
    ospfv2Topology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2Topology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2Topology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2Topology.EntityData.Children = make(map[string]types.YChild)
    ospfv2Topology.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Topology.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2Topology.MtId}
    ospfv2Topology.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2Topology.Metric}
    return &(ospfv2Topology.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2External
// External LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2External struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MT-ID for topology enabled on the link. The type
    // is interface{} with range: 0..4294967295.
    MtId interface{}

    // Metric for the topology. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}
}

func (ospfv2External *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv2External) GetEntityData() *types.CommonEntityData {
    ospfv2External.EntityData.YFilter = ospfv2External.YFilter
    ospfv2External.EntityData.YangName = "ospfv2-external"
    ospfv2External.EntityData.BundleName = "cisco_ios_xe"
    ospfv2External.EntityData.ParentYangName = "area-scope-lsa"
    ospfv2External.EntityData.SegmentPath = "ospfv2-external" + "[mt-id='" + fmt.Sprintf("%v", ospfv2External.MtId) + "']"
    ospfv2External.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv2External.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv2External.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv2External.EntityData.Children = make(map[string]types.YChild)
    ospfv2External.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2External.EntityData.Leafs["mt-id"] = types.YLeaf{"MtId", ospfv2External.MtId}
    ospfv2External.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv2External.Metric}
    ospfv2External.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", ospfv2External.ForwardingAddress}
    ospfv2External.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", ospfv2External.ExternalRouteTag}
    return &(ospfv2External.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa
// OSPFv3 LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Decoded OSPFv3 LSA header.
    Header OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header

    // Decoded OSPFv3 LSA body.
    LsaBody OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody
}

func (ospfv3Lsa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa) GetEntityData() *types.CommonEntityData {
    ospfv3Lsa.EntityData.YFilter = ospfv3Lsa.YFilter
    ospfv3Lsa.EntityData.YangName = "ospfv3-lsa"
    ospfv3Lsa.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Lsa.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Lsa.EntityData.SegmentPath = "ospfv3-lsa"
    ospfv3Lsa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Lsa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Lsa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Lsa.EntityData.Children = make(map[string]types.YChild)
    ospfv3Lsa.EntityData.Children["header"] = types.YChild{"Header", &ospfv3Lsa.Header}
    ospfv3Lsa.EntityData.Children["lsa-body"] = types.YChild{"LsaBody", &ospfv3Lsa.LsaBody}
    ospfv3Lsa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3Lsa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header
// Decoded OSPFv3 LSA header
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA ID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LsaId interface{}

    // OSPFv3 LSA options. The type is map[string]bool.
    LsaHdrOptions interface{}

    // LSA header.
    LsaHeader OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader
}

func (header *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xe"
    header.EntityData.ParentYangName = "ospfv3-lsa"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Children["lsa-header"] = types.YChild{"LsaHeader", &header.LsaHeader}
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["lsa-id"] = types.YLeaf{"LsaId", header.LsaId}
    header.EntityData.Leafs["lsa-hdr-options"] = types.YLeaf{"LsaHdrOptions", header.LsaHdrOptions}
    return &(header.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader
// LSA header
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA age. The type is interface{} with range: 0..65535.
    Age interface{}

    // LSA type. The type is interface{} with range: 0..65535.
    Type_ interface{}

    // LSA advertising router. The type is interface{} with range: 0..4294967295.
    AdvRouter interface{}

    // LSA sequence number. The type is string.
    SeqNum interface{}

    // LSA checksum. The type is string.
    Checksum interface{}

    // LSA length. The type is interface{} with range: 0..65535.
    Length interface{}
}

func (lsaHeader *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_Header_LsaHeader) GetEntityData() *types.CommonEntityData {
    lsaHeader.EntityData.YFilter = lsaHeader.YFilter
    lsaHeader.EntityData.YangName = "lsa-header"
    lsaHeader.EntityData.BundleName = "cisco_ios_xe"
    lsaHeader.EntityData.ParentYangName = "header"
    lsaHeader.EntityData.SegmentPath = "lsa-header"
    lsaHeader.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaHeader.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaHeader.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaHeader.EntityData.Children = make(map[string]types.YChild)
    lsaHeader.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaHeader.EntityData.Leafs["age"] = types.YLeaf{"Age", lsaHeader.Age}
    lsaHeader.EntityData.Leafs["type"] = types.YLeaf{"Type_", lsaHeader.Type_}
    lsaHeader.EntityData.Leafs["adv-router"] = types.YLeaf{"AdvRouter", lsaHeader.AdvRouter}
    lsaHeader.EntityData.Leafs["seq-num"] = types.YLeaf{"SeqNum", lsaHeader.SeqNum}
    lsaHeader.EntityData.Leafs["checksum"] = types.YLeaf{"Checksum", lsaHeader.Checksum}
    lsaHeader.EntityData.Leafs["length"] = types.YLeaf{"Length", lsaHeader.Length}
    return &(lsaHeader.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody
// Decoded OSPFv3 LSA body
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA options. The type is map[string]bool.
    LsaFlagOptions interface{}

    // LSA Body Flags. The type is map[string]bool.
    LsaBodyFlags interface{}

    // OSPFv3 network.
    Network OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network

    // OSPFv3 inter area prefix.
    Prefix OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix

    // OSPFv3 inter area router.
    IaRouter OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter

    // OSPFv3 LSA external.
    LsaExternal OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal

    // OSPFv3 NSSA.
    Nssa OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa

    // OSPFv3 Link data.
    LinkData OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData

    // OSPFv3 Intra area prefixes.
    IaPrefix OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix
}

func (lsaBody *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody) GetEntityData() *types.CommonEntityData {
    lsaBody.EntityData.YFilter = lsaBody.YFilter
    lsaBody.EntityData.YangName = "lsa-body"
    lsaBody.EntityData.BundleName = "cisco_ios_xe"
    lsaBody.EntityData.ParentYangName = "ospfv3-lsa"
    lsaBody.EntityData.SegmentPath = "lsa-body"
    lsaBody.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaBody.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaBody.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaBody.EntityData.Children = make(map[string]types.YChild)
    lsaBody.EntityData.Children["network"] = types.YChild{"Network", &lsaBody.Network}
    lsaBody.EntityData.Children["prefix"] = types.YChild{"Prefix", &lsaBody.Prefix}
    lsaBody.EntityData.Children["ia-router"] = types.YChild{"IaRouter", &lsaBody.IaRouter}
    lsaBody.EntityData.Children["lsa-external"] = types.YChild{"LsaExternal", &lsaBody.LsaExternal}
    lsaBody.EntityData.Children["nssa"] = types.YChild{"Nssa", &lsaBody.Nssa}
    lsaBody.EntityData.Children["link-data"] = types.YChild{"LinkData", &lsaBody.LinkData}
    lsaBody.EntityData.Children["ia-prefix"] = types.YChild{"IaPrefix", &lsaBody.IaPrefix}
    lsaBody.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaBody.EntityData.Leafs["lsa-flag-options"] = types.YLeaf{"LsaFlagOptions", lsaBody.LsaFlagOptions}
    lsaBody.EntityData.Leafs["lsa-body-flags"] = types.YLeaf{"LsaBodyFlags", lsaBody.LsaBodyFlags}
    return &(lsaBody.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network
// OSPFv3 network
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of the routers attached to the network. The type is slice of
    // interface{} with range: 0..4294967295.
    AttachedRouter []interface{}

    // Network LSA options. The type is map[string]bool.
    LsaNetOptions interface{}
}

func (network *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xe"
    network.EntityData.ParentYangName = "lsa-body"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    network.EntityData.Children = make(map[string]types.YChild)
    network.EntityData.Leafs = make(map[string]types.YLeaf)
    network.EntityData.Leafs["attached-router"] = types.YLeaf{"AttachedRouter", network.AttachedRouter}
    network.EntityData.Leafs["lsa-net-options"] = types.YLeaf{"LsaNetOptions", network.LsaNetOptions}
    return &(network.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix
// OSPFv3 inter area prefix
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Inter area Prefix. The type is string.
    IaPrefix interface{}

    // Inter area prefix options. The type is string.
    IaPrefixOptions interface{}
}

func (prefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xe"
    prefix.EntityData.ParentYangName = "lsa-body"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["metric"] = types.YLeaf{"Metric", prefix.Metric}
    prefix.EntityData.Leafs["ia-prefix"] = types.YLeaf{"IaPrefix", prefix.IaPrefix}
    prefix.EntityData.Leafs["ia-prefix-options"] = types.YLeaf{"IaPrefixOptions", prefix.IaPrefixOptions}
    return &(prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter
// OSPFv3 inter area router
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Router ID of the router being described by the LSA. The type is interface{}
    // with range: 0..4294967295.
    DestinationRouterId interface{}

    // Inter area LSA options. The type is map[string]bool.
    LsaIaOptions interface{}
}

func (iaRouter *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaRouter) GetEntityData() *types.CommonEntityData {
    iaRouter.EntityData.YFilter = iaRouter.YFilter
    iaRouter.EntityData.YangName = "ia-router"
    iaRouter.EntityData.BundleName = "cisco_ios_xe"
    iaRouter.EntityData.ParentYangName = "lsa-body"
    iaRouter.EntityData.SegmentPath = "ia-router"
    iaRouter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaRouter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaRouter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaRouter.EntityData.Children = make(map[string]types.YChild)
    iaRouter.EntityData.Leafs = make(map[string]types.YLeaf)
    iaRouter.EntityData.Leafs["metric"] = types.YLeaf{"Metric", iaRouter.Metric}
    iaRouter.EntityData.Leafs["destination-router-id"] = types.YLeaf{"DestinationRouterId", iaRouter.DestinationRouterId}
    iaRouter.EntityData.Leafs["lsa-ia-options"] = types.YLeaf{"LsaIaOptions", iaRouter.LsaIaOptions}
    return &(iaRouter.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal
// OSPFv3 LSA external
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags
}

func (lsaExternal *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal) GetEntityData() *types.CommonEntityData {
    lsaExternal.EntityData.YFilter = lsaExternal.YFilter
    lsaExternal.EntityData.YangName = "lsa-external"
    lsaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaExternal.EntityData.ParentYangName = "lsa-body"
    lsaExternal.EntityData.SegmentPath = "lsa-external"
    lsaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaExternal.EntityData.Children = make(map[string]types.YChild)
    lsaExternal.EntityData.Children["flags"] = types.YChild{"Flags", &lsaExternal.Flags}
    lsaExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", lsaExternal.Metric}
    lsaExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", lsaExternal.ReferencedLsType}
    lsaExternal.EntityData.Leafs["external-prefix"] = types.YLeaf{"ExternalPrefix", lsaExternal.ExternalPrefix}
    lsaExternal.EntityData.Leafs["external-prefix-options"] = types.YLeaf{"ExternalPrefixOptions", lsaExternal.ExternalPrefixOptions}
    lsaExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", lsaExternal.ForwardingAddress}
    lsaExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", lsaExternal.ExternalRouteTag}
    lsaExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", lsaExternal.ReferencedLinkStateId}
    return &(lsaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LsaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    flags.EntityData.Leafs["e-flag"] = types.YLeaf{"EFlag", flags.EFlag}
    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa
// OSPFv3 NSSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA LSA.
    LsaNssaExternal OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
}

func (nssa *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xe"
    nssa.EntityData.ParentYangName = "lsa-body"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nssa.EntityData.Children = make(map[string]types.YChild)
    nssa.EntityData.Children["lsa-nssa-external"] = types.YChild{"LsaNssaExternal", &nssa.LsaNssaExternal}
    nssa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nssa.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal
// NSSA LSA
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Prefix. The type is string.
    ExternalPrefix interface{}

    // Prefix options. The type is string.
    ExternalPrefixOptions interface{}

    // Forwarding address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ForwardingAddress interface{}

    // Route tag. The type is interface{} with range: 0..4294967295.
    ExternalRouteTag interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // LSA Flags.
    Flags OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
}

func (lsaNssaExternal *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal) GetEntityData() *types.CommonEntityData {
    lsaNssaExternal.EntityData.YFilter = lsaNssaExternal.YFilter
    lsaNssaExternal.EntityData.YangName = "lsa-nssa-external"
    lsaNssaExternal.EntityData.BundleName = "cisco_ios_xe"
    lsaNssaExternal.EntityData.ParentYangName = "nssa"
    lsaNssaExternal.EntityData.SegmentPath = "lsa-nssa-external"
    lsaNssaExternal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lsaNssaExternal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lsaNssaExternal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lsaNssaExternal.EntityData.Children = make(map[string]types.YChild)
    lsaNssaExternal.EntityData.Children["flags"] = types.YChild{"Flags", &lsaNssaExternal.Flags}
    lsaNssaExternal.EntityData.Leafs = make(map[string]types.YLeaf)
    lsaNssaExternal.EntityData.Leafs["metric"] = types.YLeaf{"Metric", lsaNssaExternal.Metric}
    lsaNssaExternal.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", lsaNssaExternal.ReferencedLsType}
    lsaNssaExternal.EntityData.Leafs["external-prefix"] = types.YLeaf{"ExternalPrefix", lsaNssaExternal.ExternalPrefix}
    lsaNssaExternal.EntityData.Leafs["external-prefix-options"] = types.YLeaf{"ExternalPrefixOptions", lsaNssaExternal.ExternalPrefixOptions}
    lsaNssaExternal.EntityData.Leafs["forwarding-address"] = types.YLeaf{"ForwardingAddress", lsaNssaExternal.ForwardingAddress}
    lsaNssaExternal.EntityData.Leafs["external-route-tag"] = types.YLeaf{"ExternalRouteTag", lsaNssaExternal.ExternalRouteTag}
    lsaNssaExternal.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", lsaNssaExternal.ReferencedLinkStateId}
    return &(lsaNssaExternal.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags
// LSA Flags
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When set, the metric specified is a Type 2 external metric. The type is
    // bool.
    EFlag interface{}
}

func (flags *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_Nssa_LsaNssaExternal_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xe"
    flags.EntityData.ParentYangName = "lsa-nssa-external"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    flags.EntityData.Leafs["e-flag"] = types.YLeaf{"EFlag", flags.EFlag}
    return &(flags.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData
// OSPFv3 Link data
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router priority of the interce. The type is interface{} with range: 0..255.
    RtrPriority interface{}

    // The originating router's link-local interface address on the link. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LinkLocalInterfaceAddress interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumOfPrefixes interface{}

    // Link data LSA options. The type is map[string]bool.
    LsaIdOptions interface{}
}

func (linkData *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_LinkData) GetEntityData() *types.CommonEntityData {
    linkData.EntityData.YFilter = linkData.YFilter
    linkData.EntityData.YangName = "link-data"
    linkData.EntityData.BundleName = "cisco_ios_xe"
    linkData.EntityData.ParentYangName = "lsa-body"
    linkData.EntityData.SegmentPath = "link-data"
    linkData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkData.EntityData.Children = make(map[string]types.YChild)
    linkData.EntityData.Leafs = make(map[string]types.YLeaf)
    linkData.EntityData.Leafs["rtr-priority"] = types.YLeaf{"RtrPriority", linkData.RtrPriority}
    linkData.EntityData.Leafs["link-local-interface-address"] = types.YLeaf{"LinkLocalInterfaceAddress", linkData.LinkLocalInterfaceAddress}
    linkData.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", linkData.NumOfPrefixes}
    linkData.EntityData.Leafs["lsa-id-options"] = types.YLeaf{"LsaIdOptions", linkData.LsaIdOptions}
    return &(linkData.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix
// OSPFv3 Intra area prefixes
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Referenced Link State type. The type is interface{} with range: 0..65535.
    ReferencedLsType interface{}

    // Referenced Link State ID. The type is interface{} with range:
    // 0..4294967295.
    ReferencedLinkStateId interface{}

    // Referenced Advertising Router. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ReferencedAdvRouter interface{}

    // Number of prefixes. The type is interface{} with range: 0..65535.
    NumOfPrefixes interface{}
}

func (iaPrefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Lsa_LsaBody_IaPrefix) GetEntityData() *types.CommonEntityData {
    iaPrefix.EntityData.YFilter = iaPrefix.YFilter
    iaPrefix.EntityData.YangName = "ia-prefix"
    iaPrefix.EntityData.BundleName = "cisco_ios_xe"
    iaPrefix.EntityData.ParentYangName = "lsa-body"
    iaPrefix.EntityData.SegmentPath = "ia-prefix"
    iaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iaPrefix.EntityData.Children = make(map[string]types.YChild)
    iaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    iaPrefix.EntityData.Leafs["referenced-ls-type"] = types.YLeaf{"ReferencedLsType", iaPrefix.ReferencedLsType}
    iaPrefix.EntityData.Leafs["referenced-link-state-id"] = types.YLeaf{"ReferencedLinkStateId", iaPrefix.ReferencedLinkStateId}
    iaPrefix.EntityData.Leafs["referenced-adv-router"] = types.YLeaf{"ReferencedAdvRouter", iaPrefix.ReferencedAdvRouter}
    iaPrefix.EntityData.Leafs["num-of-prefixes"] = types.YLeaf{"NumOfPrefixes", iaPrefix.NumOfPrefixes}
    return &(iaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Link
// OSPFv3 links
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceId interface{}

    // This attribute is a key. Neighbor interface ID. The type is interface{}
    // with range: 0..4294967295.
    NeighborInterfaceId interface{}

    // This attribute is a key. Neighbor router ID. The type is interface{} with
    // range: 0..4294967295.
    NeighborRouterId interface{}

    // Link type. The type is interface{} with range: 0..255.
    Type_ interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}
}

func (ospfv3Link *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Link) GetEntityData() *types.CommonEntityData {
    ospfv3Link.EntityData.YFilter = ospfv3Link.YFilter
    ospfv3Link.EntityData.YangName = "ospfv3-link"
    ospfv3Link.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Link.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Link.EntityData.SegmentPath = "ospfv3-link" + "[interface-id='" + fmt.Sprintf("%v", ospfv3Link.InterfaceId) + "']" + "[neighbor-interface-id='" + fmt.Sprintf("%v", ospfv3Link.NeighborInterfaceId) + "']" + "[neighbor-router-id='" + fmt.Sprintf("%v", ospfv3Link.NeighborRouterId) + "']"
    ospfv3Link.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Link.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Link.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Link.EntityData.Children = make(map[string]types.YChild)
    ospfv3Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3Link.EntityData.Leafs["interface-id"] = types.YLeaf{"InterfaceId", ospfv3Link.InterfaceId}
    ospfv3Link.EntityData.Leafs["neighbor-interface-id"] = types.YLeaf{"NeighborInterfaceId", ospfv3Link.NeighborInterfaceId}
    ospfv3Link.EntityData.Leafs["neighbor-router-id"] = types.YLeaf{"NeighborRouterId", ospfv3Link.NeighborRouterId}
    ospfv3Link.EntityData.Leafs["type"] = types.YLeaf{"Type_", ospfv3Link.Type_}
    ospfv3Link.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfv3Link.Metric}
    return &(ospfv3Link.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Prefix
// OSPFv3 prefix-list
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3Prefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3Prefix) GetEntityData() *types.CommonEntityData {
    ospfv3Prefix.EntityData.YFilter = ospfv3Prefix.YFilter
    ospfv3Prefix.EntityData.YangName = "ospfv3-prefix"
    ospfv3Prefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3Prefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3Prefix.EntityData.SegmentPath = "ospfv3-prefix" + "[prefix='" + fmt.Sprintf("%v", ospfv3Prefix.Prefix) + "']"
    ospfv3Prefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3Prefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3Prefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3Prefix.EntityData.Children = make(map[string]types.YChild)
    ospfv3Prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3Prefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ospfv3Prefix.Prefix}
    ospfv3Prefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", ospfv3Prefix.PrefixOptions}
    return &(ospfv3Prefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix
// OSPFv3 intra-area prefix-list
type OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix. The type is string.
    Prefix interface{}

    // Prefix options. The type is string.
    PrefixOptions interface{}
}

func (ospfv3IaPrefix *OspfOperData_OspfState_OspfInstance_LinkScopeLsas_AreaScopeLsa_Ospfv3IaPrefix) GetEntityData() *types.CommonEntityData {
    ospfv3IaPrefix.EntityData.YFilter = ospfv3IaPrefix.YFilter
    ospfv3IaPrefix.EntityData.YangName = "ospfv3-ia-prefix"
    ospfv3IaPrefix.EntityData.BundleName = "cisco_ios_xe"
    ospfv3IaPrefix.EntityData.ParentYangName = "area-scope-lsa"
    ospfv3IaPrefix.EntityData.SegmentPath = "ospfv3-ia-prefix" + "[prefix='" + fmt.Sprintf("%v", ospfv3IaPrefix.Prefix) + "']"
    ospfv3IaPrefix.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfv3IaPrefix.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfv3IaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfv3IaPrefix.EntityData.Children = make(map[string]types.YChild)
    ospfv3IaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3IaPrefix.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", ospfv3IaPrefix.Prefix}
    ospfv3IaPrefix.EntityData.Leafs["prefix-options"] = types.YLeaf{"PrefixOptions", ospfv3IaPrefix.PrefixOptions}
    return &(ospfv3IaPrefix.EntityData)
}

// OspfOperData_OspfState_OspfInstance_MultiTopology
// OSPF multi-topology interface augmentation
type OspfOperData_OspfState_OspfInstance_MultiTopology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. One of the topology enabled on this interface. The
    // type is string.
    Name interface{}
}

func (multiTopology *OspfOperData_OspfState_OspfInstance_MultiTopology) GetEntityData() *types.CommonEntityData {
    multiTopology.EntityData.YFilter = multiTopology.YFilter
    multiTopology.EntityData.YangName = "multi-topology"
    multiTopology.EntityData.BundleName = "cisco_ios_xe"
    multiTopology.EntityData.ParentYangName = "ospf-instance"
    multiTopology.EntityData.SegmentPath = "multi-topology" + "[name='" + fmt.Sprintf("%v", multiTopology.Name) + "']"
    multiTopology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    multiTopology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    multiTopology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    multiTopology.EntityData.Children = make(map[string]types.YChild)
    multiTopology.EntityData.Leafs = make(map[string]types.YLeaf)
    multiTopology.EntityData.Leafs["name"] = types.YLeaf{"Name", multiTopology.Name}
    return &(multiTopology.EntityData)
}

