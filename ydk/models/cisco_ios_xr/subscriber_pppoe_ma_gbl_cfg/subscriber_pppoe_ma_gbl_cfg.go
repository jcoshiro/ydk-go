// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-pppoe-ma-gbl package configuration.
// 
// This module contains definitions
// for the following management objects:
//   pppoe-cfg: PPPOE configuration data
// 
// This YANG module augments the
//   Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_pppoe_ma_gbl_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_pppoe_ma_gbl_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg pppoe-cfg}", reflect.TypeOf(PppoeCfg{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg:pppoe-cfg", reflect.TypeOf(PppoeCfg{}))
}

// PppoeInvalidSessionIdBehavior represents Pppoe invalid session id behavior
type PppoeInvalidSessionIdBehavior string

const (
    // Drop packets with an invalid session ID
    PppoeInvalidSessionIdBehavior_drop PppoeInvalidSessionIdBehavior = "drop"

    // Log packets with an invalid session ID
    PppoeInvalidSessionIdBehavior_log PppoeInvalidSessionIdBehavior = "log"
)

// PppoeCfg
// PPPOE configuration data
type PppoeCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable per-parent session ID spaces. The type is interface{}.
    SessionIdSpaceFlat interface{}

    // Set the PPPoE in-flight window size. The type is interface{} with range:
    // 1..20000.
    InFlightWindow interface{}

    // PPPoE BBA-Group configuration data.
    PppoeBbaGroups PppoeCfg_PppoeBbaGroups
}

func (pppoeCfg *PppoeCfg) GetEntityData() *types.CommonEntityData {
    pppoeCfg.EntityData.YFilter = pppoeCfg.YFilter
    pppoeCfg.EntityData.YangName = "pppoe-cfg"
    pppoeCfg.EntityData.BundleName = "cisco_ios_xr"
    pppoeCfg.EntityData.ParentYangName = "Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg"
    pppoeCfg.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg:pppoe-cfg"
    pppoeCfg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoeCfg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoeCfg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoeCfg.EntityData.Children = make(map[string]types.YChild)
    pppoeCfg.EntityData.Children["pppoe-bba-groups"] = types.YChild{"PppoeBbaGroups", &pppoeCfg.PppoeBbaGroups}
    pppoeCfg.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoeCfg.EntityData.Leafs["session-id-space-flat"] = types.YLeaf{"SessionIdSpaceFlat", pppoeCfg.SessionIdSpaceFlat}
    pppoeCfg.EntityData.Leafs["in-flight-window"] = types.YLeaf{"InFlightWindow", pppoeCfg.InFlightWindow}
    return &(pppoeCfg.EntityData)
}

// PppoeCfg_PppoeBbaGroups
// PPPoE BBA-Group configuration data
type PppoeCfg_PppoeBbaGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE BBA-Group configuration data. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup.
    PppoeBbaGroup []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup
}

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetEntityData() *types.CommonEntityData {
    pppoeBbaGroups.EntityData.YFilter = pppoeBbaGroups.YFilter
    pppoeBbaGroups.EntityData.YangName = "pppoe-bba-groups"
    pppoeBbaGroups.EntityData.BundleName = "cisco_ios_xr"
    pppoeBbaGroups.EntityData.ParentYangName = "pppoe-cfg"
    pppoeBbaGroups.EntityData.SegmentPath = "pppoe-bba-groups"
    pppoeBbaGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoeBbaGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoeBbaGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoeBbaGroups.EntityData.Children = make(map[string]types.YChild)
    pppoeBbaGroups.EntityData.Children["pppoe-bba-group"] = types.YChild{"PppoeBbaGroup", nil}
    for i := range pppoeBbaGroups.PppoeBbaGroup {
        pppoeBbaGroups.EntityData.Children[types.GetSegmentPath(&pppoeBbaGroups.PppoeBbaGroup[i])] = types.YChild{"PppoeBbaGroup", &pppoeBbaGroups.PppoeBbaGroup[i]}
    }
    pppoeBbaGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pppoeBbaGroups.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup
// PPPoE BBA-Group configuration data
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. BBA-Group name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    BbaGroup interface{}

    // PPPoE session completion timeout. The type is interface{} with range:
    // 30..600.
    CompletionTimeout interface{}

    // Invalid session-ID behavior. The type is PppoeInvalidSessionIdBehavior.
    InvalidSessionId interface{}

    // Enable padt after shutdown. The type is interface{}.
    EnablePadtAfterShutDown interface{}

    // PPPoE default MTU. The type is interface{} with range: 500..2000.
    Mtu interface{}

    // PPPoE tag configuration data.
    Tag PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag

    // PPPoE session configuration data.
    Sessions PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions

    // PPPoE control-packet configuration data.
    ControlPackets PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets

    // PPPoE PADO delay configuration data.
    PaDoDelay PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay
}

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetEntityData() *types.CommonEntityData {
    pppoeBbaGroup.EntityData.YFilter = pppoeBbaGroup.YFilter
    pppoeBbaGroup.EntityData.YangName = "pppoe-bba-group"
    pppoeBbaGroup.EntityData.BundleName = "cisco_ios_xr"
    pppoeBbaGroup.EntityData.ParentYangName = "pppoe-bba-groups"
    pppoeBbaGroup.EntityData.SegmentPath = "pppoe-bba-group" + "[bba-group='" + fmt.Sprintf("%v", pppoeBbaGroup.BbaGroup) + "']"
    pppoeBbaGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoeBbaGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoeBbaGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoeBbaGroup.EntityData.Children = make(map[string]types.YChild)
    pppoeBbaGroup.EntityData.Children["tag"] = types.YChild{"Tag", &pppoeBbaGroup.Tag}
    pppoeBbaGroup.EntityData.Children["sessions"] = types.YChild{"Sessions", &pppoeBbaGroup.Sessions}
    pppoeBbaGroup.EntityData.Children["control-packets"] = types.YChild{"ControlPackets", &pppoeBbaGroup.ControlPackets}
    pppoeBbaGroup.EntityData.Children["pa-do-delay"] = types.YChild{"PaDoDelay", &pppoeBbaGroup.PaDoDelay}
    pppoeBbaGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoeBbaGroup.EntityData.Leafs["bba-group"] = types.YLeaf{"BbaGroup", pppoeBbaGroup.BbaGroup}
    pppoeBbaGroup.EntityData.Leafs["completion-timeout"] = types.YLeaf{"CompletionTimeout", pppoeBbaGroup.CompletionTimeout}
    pppoeBbaGroup.EntityData.Leafs["invalid-session-id"] = types.YLeaf{"InvalidSessionId", pppoeBbaGroup.InvalidSessionId}
    pppoeBbaGroup.EntityData.Leafs["enable-padt-after-shut-down"] = types.YLeaf{"EnablePadtAfterShutDown", pppoeBbaGroup.EnablePadtAfterShutDown}
    pppoeBbaGroup.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", pppoeBbaGroup.Mtu}
    return &(pppoeBbaGroup.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag
// PPPoE tag configuration data
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ignore the ppp-max-payload tag. The type is interface{}.
    PppMaxPayloadDeny interface{}

    // Disable advertising of unrequested service names. The type is interface{}.
    ServiceSelectionDisable interface{}

    // The name to include in the AC tag. The type is string.
    AcName interface{}

    // PADR packets.
    Padr PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr

    // Service name.
    ServiceNameConfigureds PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds

    // Minimum and maximum payloads.
    PppMaxPayload PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload
}

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetEntityData() *types.CommonEntityData {
    tag.EntityData.YFilter = tag.YFilter
    tag.EntityData.YangName = "tag"
    tag.EntityData.BundleName = "cisco_ios_xr"
    tag.EntityData.ParentYangName = "pppoe-bba-group"
    tag.EntityData.SegmentPath = "tag"
    tag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tag.EntityData.Children = make(map[string]types.YChild)
    tag.EntityData.Children["padr"] = types.YChild{"Padr", &tag.Padr}
    tag.EntityData.Children["service-name-configureds"] = types.YChild{"ServiceNameConfigureds", &tag.ServiceNameConfigureds}
    tag.EntityData.Children["ppp-max-payload"] = types.YChild{"PppMaxPayload", &tag.PppMaxPayload}
    tag.EntityData.Leafs = make(map[string]types.YLeaf)
    tag.EntityData.Leafs["ppp-max-payload-deny"] = types.YLeaf{"PppMaxPayloadDeny", tag.PppMaxPayloadDeny}
    tag.EntityData.Leafs["service-selection-disable"] = types.YLeaf{"ServiceSelectionDisable", tag.ServiceSelectionDisable}
    tag.EntityData.Leafs["ac-name"] = types.YLeaf{"AcName", tag.AcName}
    return &(tag.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr
// PADR packets
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow sessions to come up with unique relay-session-id in padr. The type is
    // interface{}.
    SessionUniqueRelaySessionId interface{}

    // Allow sessions to come up with invalid-payload. The type is interface{}.
    InvalidPayloadAllow interface{}
}

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetEntityData() *types.CommonEntityData {
    padr.EntityData.YFilter = padr.YFilter
    padr.EntityData.YangName = "padr"
    padr.EntityData.BundleName = "cisco_ios_xr"
    padr.EntityData.ParentYangName = "tag"
    padr.EntityData.SegmentPath = "padr"
    padr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    padr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    padr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    padr.EntityData.Children = make(map[string]types.YChild)
    padr.EntityData.Leafs = make(map[string]types.YLeaf)
    padr.EntityData.Leafs["session-unique-relay-session-id"] = types.YLeaf{"SessionUniqueRelaySessionId", padr.SessionUniqueRelaySessionId}
    padr.EntityData.Leafs["invalid-payload-allow"] = types.YLeaf{"InvalidPayloadAllow", padr.InvalidPayloadAllow}
    return &(padr.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds
// Service name
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service name supported on this group. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured.
    ServiceNameConfigured []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured
}

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetEntityData() *types.CommonEntityData {
    serviceNameConfigureds.EntityData.YFilter = serviceNameConfigureds.YFilter
    serviceNameConfigureds.EntityData.YangName = "service-name-configureds"
    serviceNameConfigureds.EntityData.BundleName = "cisco_ios_xr"
    serviceNameConfigureds.EntityData.ParentYangName = "tag"
    serviceNameConfigureds.EntityData.SegmentPath = "service-name-configureds"
    serviceNameConfigureds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceNameConfigureds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceNameConfigureds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceNameConfigureds.EntityData.Children = make(map[string]types.YChild)
    serviceNameConfigureds.EntityData.Children["service-name-configured"] = types.YChild{"ServiceNameConfigured", nil}
    for i := range serviceNameConfigureds.ServiceNameConfigured {
        serviceNameConfigureds.EntityData.Children[types.GetSegmentPath(&serviceNameConfigureds.ServiceNameConfigured[i])] = types.YChild{"ServiceNameConfigured", &serviceNameConfigureds.ServiceNameConfigured[i]}
    }
    serviceNameConfigureds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serviceNameConfigureds.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured
// Service name supported on this group
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Service name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}
}

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetEntityData() *types.CommonEntityData {
    serviceNameConfigured.EntityData.YFilter = serviceNameConfigured.YFilter
    serviceNameConfigured.EntityData.YangName = "service-name-configured"
    serviceNameConfigured.EntityData.BundleName = "cisco_ios_xr"
    serviceNameConfigured.EntityData.ParentYangName = "service-name-configureds"
    serviceNameConfigured.EntityData.SegmentPath = "service-name-configured" + "[name='" + fmt.Sprintf("%v", serviceNameConfigured.Name) + "']"
    serviceNameConfigured.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceNameConfigured.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceNameConfigured.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceNameConfigured.EntityData.Children = make(map[string]types.YChild)
    serviceNameConfigured.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceNameConfigured.EntityData.Leafs["name"] = types.YLeaf{"Name", serviceNameConfigured.Name}
    return &(serviceNameConfigured.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload
// Minimum and maximum payloads
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum payload. The type is interface{} with range: 500..2000. This
    // attribute is mandatory.
    Min interface{}

    // Maximum payload. The type is interface{} with range: 500..2000. This
    // attribute is mandatory.
    Max interface{}
}

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetEntityData() *types.CommonEntityData {
    pppMaxPayload.EntityData.YFilter = pppMaxPayload.YFilter
    pppMaxPayload.EntityData.YangName = "ppp-max-payload"
    pppMaxPayload.EntityData.BundleName = "cisco_ios_xr"
    pppMaxPayload.EntityData.ParentYangName = "tag"
    pppMaxPayload.EntityData.SegmentPath = "ppp-max-payload"
    pppMaxPayload.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppMaxPayload.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppMaxPayload.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppMaxPayload.EntityData.Children = make(map[string]types.YChild)
    pppMaxPayload.EntityData.Leafs = make(map[string]types.YLeaf)
    pppMaxPayload.EntityData.Leafs["min"] = types.YLeaf{"Min", pppMaxPayload.Min}
    pppMaxPayload.EntityData.Leafs["max"] = types.YLeaf{"Max", pppMaxPayload.Max}
    return &(pppMaxPayload.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions
// PPPoE session configuration data
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set VLAN (inner + outer tags) session throttle.
    VlanThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle

    // Set Inner VLAN session throttle.
    InnerVlanThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle

    // Set Remote ID session limit and threshold.
    RemoteIdLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit

    // Set per-MAC/Access interface throttle for IWF sessions.
    MacIwfAccessInterfaceThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle

    // Set per-access interface limit.
    AccessInterfaceLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit

    // Set per-MAC/Access Interface throttle.
    MacAccessInterfaceThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle

    // Set Outer VLAN session limit and threshold.
    OuterVlanLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit

    // Set Circuit ID session throttle.
    CircuitIdThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle

    // Set per-MAC address session limit and threshold.
    MacLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit

    // Set Circuit ID session limit and threshold.
    CircuitIdLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit

    // Set per-MAC session limit and threshold for IWF sessions.
    MacIwfLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit

    // Set per-MAC access interface session limit and threshold for IWF sessions.
    MacIwfAccessInterfaceLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit

    // Set Inner VLAN session limit and threshold.
    InnerVlanLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit

    // Set Outer VLAN session throttle.
    OuterVlanThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle

    // Set per-MAC throttle.
    MacThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle

    // Set Circuit ID and Remote ID session limit/threshold.
    CircuitIdAndRemoteIdLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit

    // Set VLAN (inner + outer tags) session limit and threshold.
    VlanLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit

    // Set per-MAC access interface session limit and threshold.
    MacAccessInterfaceLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit

    // Set Remote ID session throttle.
    RemoteIdThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle

    // Set per-card session limit and threshold.
    MaxLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit

    // Set Circuit ID and Remote ID session throttle.
    CircuitIdAndRemoteIdThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle
}

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "pppoe-bba-group"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["vlan-throttle"] = types.YChild{"VlanThrottle", &sessions.VlanThrottle}
    sessions.EntityData.Children["inner-vlan-throttle"] = types.YChild{"InnerVlanThrottle", &sessions.InnerVlanThrottle}
    sessions.EntityData.Children["remote-id-limit"] = types.YChild{"RemoteIdLimit", &sessions.RemoteIdLimit}
    sessions.EntityData.Children["mac-iwf-access-interface-throttle"] = types.YChild{"MacIwfAccessInterfaceThrottle", &sessions.MacIwfAccessInterfaceThrottle}
    sessions.EntityData.Children["access-interface-limit"] = types.YChild{"AccessInterfaceLimit", &sessions.AccessInterfaceLimit}
    sessions.EntityData.Children["mac-access-interface-throttle"] = types.YChild{"MacAccessInterfaceThrottle", &sessions.MacAccessInterfaceThrottle}
    sessions.EntityData.Children["outer-vlan-limit"] = types.YChild{"OuterVlanLimit", &sessions.OuterVlanLimit}
    sessions.EntityData.Children["circuit-id-throttle"] = types.YChild{"CircuitIdThrottle", &sessions.CircuitIdThrottle}
    sessions.EntityData.Children["mac-limit"] = types.YChild{"MacLimit", &sessions.MacLimit}
    sessions.EntityData.Children["circuit-id-limit"] = types.YChild{"CircuitIdLimit", &sessions.CircuitIdLimit}
    sessions.EntityData.Children["mac-iwf-limit"] = types.YChild{"MacIwfLimit", &sessions.MacIwfLimit}
    sessions.EntityData.Children["mac-iwf-access-interface-limit"] = types.YChild{"MacIwfAccessInterfaceLimit", &sessions.MacIwfAccessInterfaceLimit}
    sessions.EntityData.Children["inner-vlan-limit"] = types.YChild{"InnerVlanLimit", &sessions.InnerVlanLimit}
    sessions.EntityData.Children["outer-vlan-throttle"] = types.YChild{"OuterVlanThrottle", &sessions.OuterVlanThrottle}
    sessions.EntityData.Children["mac-throttle"] = types.YChild{"MacThrottle", &sessions.MacThrottle}
    sessions.EntityData.Children["circuit-id-and-remote-id-limit"] = types.YChild{"CircuitIdAndRemoteIdLimit", &sessions.CircuitIdAndRemoteIdLimit}
    sessions.EntityData.Children["vlan-limit"] = types.YChild{"VlanLimit", &sessions.VlanLimit}
    sessions.EntityData.Children["mac-access-interface-limit"] = types.YChild{"MacAccessInterfaceLimit", &sessions.MacAccessInterfaceLimit}
    sessions.EntityData.Children["remote-id-throttle"] = types.YChild{"RemoteIdThrottle", &sessions.RemoteIdThrottle}
    sessions.EntityData.Children["max-limit"] = types.YChild{"MaxLimit", &sessions.MaxLimit}
    sessions.EntityData.Children["circuit-id-and-remote-id-throttle"] = types.YChild{"CircuitIdAndRemoteIdThrottle", &sessions.CircuitIdAndRemoteIdThrottle}
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle
// Set VLAN (inner + outer tags) session
// throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetEntityData() *types.CommonEntityData {
    vlanThrottle.EntityData.YFilter = vlanThrottle.YFilter
    vlanThrottle.EntityData.YangName = "vlan-throttle"
    vlanThrottle.EntityData.BundleName = "cisco_ios_xr"
    vlanThrottle.EntityData.ParentYangName = "sessions"
    vlanThrottle.EntityData.SegmentPath = "vlan-throttle"
    vlanThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanThrottle.EntityData.Children = make(map[string]types.YChild)
    vlanThrottle.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanThrottle.EntityData.Leafs["throttle"] = types.YLeaf{"Throttle", vlanThrottle.Throttle}
    vlanThrottle.EntityData.Leafs["request-period"] = types.YLeaf{"RequestPeriod", vlanThrottle.RequestPeriod}
    vlanThrottle.EntityData.Leafs["blocking-period"] = types.YLeaf{"BlockingPeriod", vlanThrottle.BlockingPeriod}
    return &(vlanThrottle.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle
// Set Inner VLAN session throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetEntityData() *types.CommonEntityData {
    innerVlanThrottle.EntityData.YFilter = innerVlanThrottle.YFilter
    innerVlanThrottle.EntityData.YangName = "inner-vlan-throttle"
    innerVlanThrottle.EntityData.BundleName = "cisco_ios_xr"
    innerVlanThrottle.EntityData.ParentYangName = "sessions"
    innerVlanThrottle.EntityData.SegmentPath = "inner-vlan-throttle"
    innerVlanThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    innerVlanThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    innerVlanThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    innerVlanThrottle.EntityData.Children = make(map[string]types.YChild)
    innerVlanThrottle.EntityData.Leafs = make(map[string]types.YLeaf)
    innerVlanThrottle.EntityData.Leafs["throttle"] = types.YLeaf{"Throttle", innerVlanThrottle.Throttle}
    innerVlanThrottle.EntityData.Leafs["request-period"] = types.YLeaf{"RequestPeriod", innerVlanThrottle.RequestPeriod}
    innerVlanThrottle.EntityData.Leafs["blocking-period"] = types.YLeaf{"BlockingPeriod", innerVlanThrottle.BlockingPeriod}
    return &(innerVlanThrottle.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit
// Set Remote ID session limit and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-Remote ID limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-Remote ID threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetEntityData() *types.CommonEntityData {
    remoteIdLimit.EntityData.YFilter = remoteIdLimit.YFilter
    remoteIdLimit.EntityData.YangName = "remote-id-limit"
    remoteIdLimit.EntityData.BundleName = "cisco_ios_xr"
    remoteIdLimit.EntityData.ParentYangName = "sessions"
    remoteIdLimit.EntityData.SegmentPath = "remote-id-limit"
    remoteIdLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteIdLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteIdLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteIdLimit.EntityData.Children = make(map[string]types.YChild)
    remoteIdLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteIdLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", remoteIdLimit.Limit}
    remoteIdLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", remoteIdLimit.Threshold}
    return &(remoteIdLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle
// Set per-MAC/Access interface throttle for IWF
// sessions
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetEntityData() *types.CommonEntityData {
    macIwfAccessInterfaceThrottle.EntityData.YFilter = macIwfAccessInterfaceThrottle.YFilter
    macIwfAccessInterfaceThrottle.EntityData.YangName = "mac-iwf-access-interface-throttle"
    macIwfAccessInterfaceThrottle.EntityData.BundleName = "cisco_ios_xr"
    macIwfAccessInterfaceThrottle.EntityData.ParentYangName = "sessions"
    macIwfAccessInterfaceThrottle.EntityData.SegmentPath = "mac-iwf-access-interface-throttle"
    macIwfAccessInterfaceThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macIwfAccessInterfaceThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macIwfAccessInterfaceThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macIwfAccessInterfaceThrottle.EntityData.Children = make(map[string]types.YChild)
    macIwfAccessInterfaceThrottle.EntityData.Leafs = make(map[string]types.YLeaf)
    macIwfAccessInterfaceThrottle.EntityData.Leafs["throttle"] = types.YLeaf{"Throttle", macIwfAccessInterfaceThrottle.Throttle}
    macIwfAccessInterfaceThrottle.EntityData.Leafs["request-period"] = types.YLeaf{"RequestPeriod", macIwfAccessInterfaceThrottle.RequestPeriod}
    macIwfAccessInterfaceThrottle.EntityData.Leafs["blocking-period"] = types.YLeaf{"BlockingPeriod", macIwfAccessInterfaceThrottle.BlockingPeriod}
    return &(macIwfAccessInterfaceThrottle.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit
// Set per-access interface limit
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-access interface session limit. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    Limit interface{}

    // Per-access interface session threshold. The type is interface{} with range:
    // 1..65535.
    Threshold interface{}
}

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetEntityData() *types.CommonEntityData {
    accessInterfaceLimit.EntityData.YFilter = accessInterfaceLimit.YFilter
    accessInterfaceLimit.EntityData.YangName = "access-interface-limit"
    accessInterfaceLimit.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceLimit.EntityData.ParentYangName = "sessions"
    accessInterfaceLimit.EntityData.SegmentPath = "access-interface-limit"
    accessInterfaceLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceLimit.EntityData.Children = make(map[string]types.YChild)
    accessInterfaceLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    accessInterfaceLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", accessInterfaceLimit.Limit}
    accessInterfaceLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", accessInterfaceLimit.Threshold}
    return &(accessInterfaceLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle
// Set per-MAC/Access Interface throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetEntityData() *types.CommonEntityData {
    macAccessInterfaceThrottle.EntityData.YFilter = macAccessInterfaceThrottle.YFilter
    macAccessInterfaceThrottle.EntityData.YangName = "mac-access-interface-throttle"
    macAccessInterfaceThrottle.EntityData.BundleName = "cisco_ios_xr"
    macAccessInterfaceThrottle.EntityData.ParentYangName = "sessions"
    macAccessInterfaceThrottle.EntityData.SegmentPath = "mac-access-interface-throttle"
    macAccessInterfaceThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAccessInterfaceThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAccessInterfaceThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAccessInterfaceThrottle.EntityData.Children = make(map[string]types.YChild)
    macAccessInterfaceThrottle.EntityData.Leafs = make(map[string]types.YLeaf)
    macAccessInterfaceThrottle.EntityData.Leafs["throttle"] = types.YLeaf{"Throttle", macAccessInterfaceThrottle.Throttle}
    macAccessInterfaceThrottle.EntityData.Leafs["request-period"] = types.YLeaf{"RequestPeriod", macAccessInterfaceThrottle.RequestPeriod}
    macAccessInterfaceThrottle.EntityData.Leafs["blocking-period"] = types.YLeaf{"BlockingPeriod", macAccessInterfaceThrottle.BlockingPeriod}
    return &(macAccessInterfaceThrottle.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit
// Set Outer VLAN session limit and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-Outer VLAN limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-Outer VLAN threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetEntityData() *types.CommonEntityData {
    outerVlanLimit.EntityData.YFilter = outerVlanLimit.YFilter
    outerVlanLimit.EntityData.YangName = "outer-vlan-limit"
    outerVlanLimit.EntityData.BundleName = "cisco_ios_xr"
    outerVlanLimit.EntityData.ParentYangName = "sessions"
    outerVlanLimit.EntityData.SegmentPath = "outer-vlan-limit"
    outerVlanLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outerVlanLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outerVlanLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outerVlanLimit.EntityData.Children = make(map[string]types.YChild)
    outerVlanLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    outerVlanLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", outerVlanLimit.Limit}
    outerVlanLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", outerVlanLimit.Threshold}
    return &(outerVlanLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle
// Set Circuit ID session throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetEntityData() *types.CommonEntityData {
    circuitIdThrottle.EntityData.YFilter = circuitIdThrottle.YFilter
    circuitIdThrottle.EntityData.YangName = "circuit-id-throttle"
    circuitIdThrottle.EntityData.BundleName = "cisco_ios_xr"
    circuitIdThrottle.EntityData.ParentYangName = "sessions"
    circuitIdThrottle.EntityData.SegmentPath = "circuit-id-throttle"
    circuitIdThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitIdThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitIdThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitIdThrottle.EntityData.Children = make(map[string]types.YChild)
    circuitIdThrottle.EntityData.Leafs = make(map[string]types.YLeaf)
    circuitIdThrottle.EntityData.Leafs["throttle"] = types.YLeaf{"Throttle", circuitIdThrottle.Throttle}
    circuitIdThrottle.EntityData.Leafs["request-period"] = types.YLeaf{"RequestPeriod", circuitIdThrottle.RequestPeriod}
    circuitIdThrottle.EntityData.Leafs["blocking-period"] = types.YLeaf{"BlockingPeriod", circuitIdThrottle.BlockingPeriod}
    return &(circuitIdThrottle.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit
// Set per-MAC address session limit and
// threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-MAC session limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-MAC session threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetEntityData() *types.CommonEntityData {
    macLimit.EntityData.YFilter = macLimit.YFilter
    macLimit.EntityData.YangName = "mac-limit"
    macLimit.EntityData.BundleName = "cisco_ios_xr"
    macLimit.EntityData.ParentYangName = "sessions"
    macLimit.EntityData.SegmentPath = "mac-limit"
    macLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macLimit.EntityData.Children = make(map[string]types.YChild)
    macLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    macLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", macLimit.Limit}
    macLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", macLimit.Threshold}
    return &(macLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit
// Set Circuit ID session limit and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-Circuit ID limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-Circuit ID threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetEntityData() *types.CommonEntityData {
    circuitIdLimit.EntityData.YFilter = circuitIdLimit.YFilter
    circuitIdLimit.EntityData.YangName = "circuit-id-limit"
    circuitIdLimit.EntityData.BundleName = "cisco_ios_xr"
    circuitIdLimit.EntityData.ParentYangName = "sessions"
    circuitIdLimit.EntityData.SegmentPath = "circuit-id-limit"
    circuitIdLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitIdLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitIdLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitIdLimit.EntityData.Children = make(map[string]types.YChild)
    circuitIdLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    circuitIdLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", circuitIdLimit.Limit}
    circuitIdLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", circuitIdLimit.Threshold}
    return &(circuitIdLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit
// Set per-MAC session limit and threshold for
// IWF sessions
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-MAC session limit for IWF sessions. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    Limit interface{}

    // Per-MAC session threshold for IWF sessions. The type is interface{} with
    // range: 1..65535.
    Threshold interface{}
}

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetEntityData() *types.CommonEntityData {
    macIwfLimit.EntityData.YFilter = macIwfLimit.YFilter
    macIwfLimit.EntityData.YangName = "mac-iwf-limit"
    macIwfLimit.EntityData.BundleName = "cisco_ios_xr"
    macIwfLimit.EntityData.ParentYangName = "sessions"
    macIwfLimit.EntityData.SegmentPath = "mac-iwf-limit"
    macIwfLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macIwfLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macIwfLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macIwfLimit.EntityData.Children = make(map[string]types.YChild)
    macIwfLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    macIwfLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", macIwfLimit.Limit}
    macIwfLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", macIwfLimit.Threshold}
    return &(macIwfLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit
// Set per-MAC access interface session limit
// and threshold for IWF sessions
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-MAC access interface session limit for IWF sessions. The type is
    // interface{} with range: 1..65535. This attribute is mandatory.
    Limit interface{}

    // Per-MAC access interface session threshold for IWF sessions. The type is
    // interface{} with range: 1..65535.
    Threshold interface{}
}

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetEntityData() *types.CommonEntityData {
    macIwfAccessInterfaceLimit.EntityData.YFilter = macIwfAccessInterfaceLimit.YFilter
    macIwfAccessInterfaceLimit.EntityData.YangName = "mac-iwf-access-interface-limit"
    macIwfAccessInterfaceLimit.EntityData.BundleName = "cisco_ios_xr"
    macIwfAccessInterfaceLimit.EntityData.ParentYangName = "sessions"
    macIwfAccessInterfaceLimit.EntityData.SegmentPath = "mac-iwf-access-interface-limit"
    macIwfAccessInterfaceLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macIwfAccessInterfaceLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macIwfAccessInterfaceLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macIwfAccessInterfaceLimit.EntityData.Children = make(map[string]types.YChild)
    macIwfAccessInterfaceLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    macIwfAccessInterfaceLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", macIwfAccessInterfaceLimit.Limit}
    macIwfAccessInterfaceLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", macIwfAccessInterfaceLimit.Threshold}
    return &(macIwfAccessInterfaceLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit
// Set Inner VLAN session limit and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-Inner VLAN limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-Inner VLAN threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetEntityData() *types.CommonEntityData {
    innerVlanLimit.EntityData.YFilter = innerVlanLimit.YFilter
    innerVlanLimit.EntityData.YangName = "inner-vlan-limit"
    innerVlanLimit.EntityData.BundleName = "cisco_ios_xr"
    innerVlanLimit.EntityData.ParentYangName = "sessions"
    innerVlanLimit.EntityData.SegmentPath = "inner-vlan-limit"
    innerVlanLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    innerVlanLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    innerVlanLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    innerVlanLimit.EntityData.Children = make(map[string]types.YChild)
    innerVlanLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    innerVlanLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", innerVlanLimit.Limit}
    innerVlanLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", innerVlanLimit.Threshold}
    return &(innerVlanLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle
// Set Outer VLAN session throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetEntityData() *types.CommonEntityData {
    outerVlanThrottle.EntityData.YFilter = outerVlanThrottle.YFilter
    outerVlanThrottle.EntityData.YangName = "outer-vlan-throttle"
    outerVlanThrottle.EntityData.BundleName = "cisco_ios_xr"
    outerVlanThrottle.EntityData.ParentYangName = "sessions"
    outerVlanThrottle.EntityData.SegmentPath = "outer-vlan-throttle"
    outerVlanThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outerVlanThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outerVlanThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outerVlanThrottle.EntityData.Children = make(map[string]types.YChild)
    outerVlanThrottle.EntityData.Leafs = make(map[string]types.YLeaf)
    outerVlanThrottle.EntityData.Leafs["throttle"] = types.YLeaf{"Throttle", outerVlanThrottle.Throttle}
    outerVlanThrottle.EntityData.Leafs["request-period"] = types.YLeaf{"RequestPeriod", outerVlanThrottle.RequestPeriod}
    outerVlanThrottle.EntityData.Leafs["blocking-period"] = types.YLeaf{"BlockingPeriod", outerVlanThrottle.BlockingPeriod}
    return &(outerVlanThrottle.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle
// Set per-MAC throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetEntityData() *types.CommonEntityData {
    macThrottle.EntityData.YFilter = macThrottle.YFilter
    macThrottle.EntityData.YangName = "mac-throttle"
    macThrottle.EntityData.BundleName = "cisco_ios_xr"
    macThrottle.EntityData.ParentYangName = "sessions"
    macThrottle.EntityData.SegmentPath = "mac-throttle"
    macThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macThrottle.EntityData.Children = make(map[string]types.YChild)
    macThrottle.EntityData.Leafs = make(map[string]types.YLeaf)
    macThrottle.EntityData.Leafs["throttle"] = types.YLeaf{"Throttle", macThrottle.Throttle}
    macThrottle.EntityData.Leafs["request-period"] = types.YLeaf{"RequestPeriod", macThrottle.RequestPeriod}
    macThrottle.EntityData.Leafs["blocking-period"] = types.YLeaf{"BlockingPeriod", macThrottle.BlockingPeriod}
    return &(macThrottle.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit
// Set Circuit ID and Remote ID session
// limit/threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-Circuit ID limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-Circuit ID threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetEntityData() *types.CommonEntityData {
    circuitIdAndRemoteIdLimit.EntityData.YFilter = circuitIdAndRemoteIdLimit.YFilter
    circuitIdAndRemoteIdLimit.EntityData.YangName = "circuit-id-and-remote-id-limit"
    circuitIdAndRemoteIdLimit.EntityData.BundleName = "cisco_ios_xr"
    circuitIdAndRemoteIdLimit.EntityData.ParentYangName = "sessions"
    circuitIdAndRemoteIdLimit.EntityData.SegmentPath = "circuit-id-and-remote-id-limit"
    circuitIdAndRemoteIdLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitIdAndRemoteIdLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitIdAndRemoteIdLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitIdAndRemoteIdLimit.EntityData.Children = make(map[string]types.YChild)
    circuitIdAndRemoteIdLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    circuitIdAndRemoteIdLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", circuitIdAndRemoteIdLimit.Limit}
    circuitIdAndRemoteIdLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", circuitIdAndRemoteIdLimit.Threshold}
    return &(circuitIdAndRemoteIdLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit
// Set VLAN (inner + outer tags) session limit
// and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-VLAN (inner + outer tags) limit. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    Limit interface{}

    // Per-VLAN (inner + outer tags) threshold. The type is interface{} with
    // range: 1..65535.
    Threshold interface{}
}

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetEntityData() *types.CommonEntityData {
    vlanLimit.EntityData.YFilter = vlanLimit.YFilter
    vlanLimit.EntityData.YangName = "vlan-limit"
    vlanLimit.EntityData.BundleName = "cisco_ios_xr"
    vlanLimit.EntityData.ParentYangName = "sessions"
    vlanLimit.EntityData.SegmentPath = "vlan-limit"
    vlanLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanLimit.EntityData.Children = make(map[string]types.YChild)
    vlanLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", vlanLimit.Limit}
    vlanLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", vlanLimit.Threshold}
    return &(vlanLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit
// Set per-MAC access interface session limit
// and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-MAC access interface session limit. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    Limit interface{}

    // Per-MAC access interface session threshold. The type is interface{} with
    // range: 1..65535.
    Threshold interface{}
}

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetEntityData() *types.CommonEntityData {
    macAccessInterfaceLimit.EntityData.YFilter = macAccessInterfaceLimit.YFilter
    macAccessInterfaceLimit.EntityData.YangName = "mac-access-interface-limit"
    macAccessInterfaceLimit.EntityData.BundleName = "cisco_ios_xr"
    macAccessInterfaceLimit.EntityData.ParentYangName = "sessions"
    macAccessInterfaceLimit.EntityData.SegmentPath = "mac-access-interface-limit"
    macAccessInterfaceLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macAccessInterfaceLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macAccessInterfaceLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macAccessInterfaceLimit.EntityData.Children = make(map[string]types.YChild)
    macAccessInterfaceLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    macAccessInterfaceLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", macAccessInterfaceLimit.Limit}
    macAccessInterfaceLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", macAccessInterfaceLimit.Threshold}
    return &(macAccessInterfaceLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle
// Set Remote ID session throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetEntityData() *types.CommonEntityData {
    remoteIdThrottle.EntityData.YFilter = remoteIdThrottle.YFilter
    remoteIdThrottle.EntityData.YangName = "remote-id-throttle"
    remoteIdThrottle.EntityData.BundleName = "cisco_ios_xr"
    remoteIdThrottle.EntityData.ParentYangName = "sessions"
    remoteIdThrottle.EntityData.SegmentPath = "remote-id-throttle"
    remoteIdThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteIdThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteIdThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteIdThrottle.EntityData.Children = make(map[string]types.YChild)
    remoteIdThrottle.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteIdThrottle.EntityData.Leafs["throttle"] = types.YLeaf{"Throttle", remoteIdThrottle.Throttle}
    remoteIdThrottle.EntityData.Leafs["request-period"] = types.YLeaf{"RequestPeriod", remoteIdThrottle.RequestPeriod}
    remoteIdThrottle.EntityData.Leafs["blocking-period"] = types.YLeaf{"BlockingPeriod", remoteIdThrottle.BlockingPeriod}
    return &(remoteIdThrottle.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit
// Set per-card session limit and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-card session limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-card session threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetEntityData() *types.CommonEntityData {
    maxLimit.EntityData.YFilter = maxLimit.YFilter
    maxLimit.EntityData.YangName = "max-limit"
    maxLimit.EntityData.BundleName = "cisco_ios_xr"
    maxLimit.EntityData.ParentYangName = "sessions"
    maxLimit.EntityData.SegmentPath = "max-limit"
    maxLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxLimit.EntityData.Children = make(map[string]types.YChild)
    maxLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    maxLimit.EntityData.Leafs["limit"] = types.YLeaf{"Limit", maxLimit.Limit}
    maxLimit.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", maxLimit.Threshold}
    return &(maxLimit.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle
// Set Circuit ID and Remote ID session throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetEntityData() *types.CommonEntityData {
    circuitIdAndRemoteIdThrottle.EntityData.YFilter = circuitIdAndRemoteIdThrottle.YFilter
    circuitIdAndRemoteIdThrottle.EntityData.YangName = "circuit-id-and-remote-id-throttle"
    circuitIdAndRemoteIdThrottle.EntityData.BundleName = "cisco_ios_xr"
    circuitIdAndRemoteIdThrottle.EntityData.ParentYangName = "sessions"
    circuitIdAndRemoteIdThrottle.EntityData.SegmentPath = "circuit-id-and-remote-id-throttle"
    circuitIdAndRemoteIdThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitIdAndRemoteIdThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitIdAndRemoteIdThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitIdAndRemoteIdThrottle.EntityData.Children = make(map[string]types.YChild)
    circuitIdAndRemoteIdThrottle.EntityData.Leafs = make(map[string]types.YLeaf)
    circuitIdAndRemoteIdThrottle.EntityData.Leafs["throttle"] = types.YLeaf{"Throttle", circuitIdAndRemoteIdThrottle.Throttle}
    circuitIdAndRemoteIdThrottle.EntityData.Leafs["request-period"] = types.YLeaf{"RequestPeriod", circuitIdAndRemoteIdThrottle.RequestPeriod}
    circuitIdAndRemoteIdThrottle.EntityData.Leafs["blocking-period"] = types.YLeaf{"BlockingPeriod", circuitIdAndRemoteIdThrottle.BlockingPeriod}
    return &(circuitIdAndRemoteIdThrottle.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets
// PPPoE control-packet configuration data
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the Priority to use for PPP and PPPoE control packets. The type is
    // interface{} with range: 0..7.
    Priority interface{}
}

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetEntityData() *types.CommonEntityData {
    controlPackets.EntityData.YFilter = controlPackets.YFilter
    controlPackets.EntityData.YangName = "control-packets"
    controlPackets.EntityData.BundleName = "cisco_ios_xr"
    controlPackets.EntityData.ParentYangName = "pppoe-bba-group"
    controlPackets.EntityData.SegmentPath = "control-packets"
    controlPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controlPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controlPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controlPackets.EntityData.Children = make(map[string]types.YChild)
    controlPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    controlPackets.EntityData.Leafs["priority"] = types.YLeaf{"Priority", controlPackets.Priority}
    return &(controlPackets.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay
// PPPoE PADO delay configuration data
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // Units are millisecond.
    Default_ interface{}

    // Configure PADO delay dependent on the received Circuit ID. The type is
    // interface{} with range: 0..10000. Units are millisecond.
    CircuitId interface{}

    // Configure PADO delay dependent on the received Remote ID. The type is
    // interface{} with range: 0..10000. Units are millisecond.
    RemoteId interface{}

    // Delay the PADO response when the received Remote ID contains the given
    // string.
    RemoteIdSubstrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings

    // Delay the PADO response when there is an exact match on the received Remote
    // ID.
    RemoteIdStrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings

    // Delay the PADO response when there is an exact match on the received
    // Service Name.
    ServiceNameStrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings

    // Delay the PADO response when the received Circuit ID contains the given
    // string.
    CircuitIdSubstrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings

    // Delay the PADO response when the received Service Name contains the given
    // string.
    ServiceNameSubstrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings

    // Delay the PADO response when there is an exact match on the received
    // Circuit ID.
    CircuitIdStrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings
}

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetEntityData() *types.CommonEntityData {
    paDoDelay.EntityData.YFilter = paDoDelay.YFilter
    paDoDelay.EntityData.YangName = "pa-do-delay"
    paDoDelay.EntityData.BundleName = "cisco_ios_xr"
    paDoDelay.EntityData.ParentYangName = "pppoe-bba-group"
    paDoDelay.EntityData.SegmentPath = "pa-do-delay"
    paDoDelay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paDoDelay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paDoDelay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paDoDelay.EntityData.Children = make(map[string]types.YChild)
    paDoDelay.EntityData.Children["remote-id-substrings"] = types.YChild{"RemoteIdSubstrings", &paDoDelay.RemoteIdSubstrings}
    paDoDelay.EntityData.Children["remote-id-strings"] = types.YChild{"RemoteIdStrings", &paDoDelay.RemoteIdStrings}
    paDoDelay.EntityData.Children["service-name-strings"] = types.YChild{"ServiceNameStrings", &paDoDelay.ServiceNameStrings}
    paDoDelay.EntityData.Children["circuit-id-substrings"] = types.YChild{"CircuitIdSubstrings", &paDoDelay.CircuitIdSubstrings}
    paDoDelay.EntityData.Children["service-name-substrings"] = types.YChild{"ServiceNameSubstrings", &paDoDelay.ServiceNameSubstrings}
    paDoDelay.EntityData.Children["circuit-id-strings"] = types.YChild{"CircuitIdStrings", &paDoDelay.CircuitIdStrings}
    paDoDelay.EntityData.Leafs = make(map[string]types.YLeaf)
    paDoDelay.EntityData.Leafs["default"] = types.YLeaf{"Default_", paDoDelay.Default_}
    paDoDelay.EntityData.Leafs["circuit-id"] = types.YLeaf{"CircuitId", paDoDelay.CircuitId}
    paDoDelay.EntityData.Leafs["remote-id"] = types.YLeaf{"RemoteId", paDoDelay.RemoteId}
    return &(paDoDelay.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings
// Delay the PADO response when the received
// Remote ID contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay the PADO response when the received Remote ID contains the given
    // string. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring.
    RemoteIdSubstring []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring
}

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetEntityData() *types.CommonEntityData {
    remoteIdSubstrings.EntityData.YFilter = remoteIdSubstrings.YFilter
    remoteIdSubstrings.EntityData.YangName = "remote-id-substrings"
    remoteIdSubstrings.EntityData.BundleName = "cisco_ios_xr"
    remoteIdSubstrings.EntityData.ParentYangName = "pa-do-delay"
    remoteIdSubstrings.EntityData.SegmentPath = "remote-id-substrings"
    remoteIdSubstrings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteIdSubstrings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteIdSubstrings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteIdSubstrings.EntityData.Children = make(map[string]types.YChild)
    remoteIdSubstrings.EntityData.Children["remote-id-substring"] = types.YChild{"RemoteIdSubstring", nil}
    for i := range remoteIdSubstrings.RemoteIdSubstring {
        remoteIdSubstrings.EntityData.Children[types.GetSegmentPath(&remoteIdSubstrings.RemoteIdSubstring[i])] = types.YChild{"RemoteIdSubstring", &remoteIdSubstrings.RemoteIdSubstring[i]}
    }
    remoteIdSubstrings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(remoteIdSubstrings.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring
// Delay the PADO response when the received
// Remote ID contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The string to be contained within the received
    // Remote ID. The type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetEntityData() *types.CommonEntityData {
    remoteIdSubstring.EntityData.YFilter = remoteIdSubstring.YFilter
    remoteIdSubstring.EntityData.YangName = "remote-id-substring"
    remoteIdSubstring.EntityData.BundleName = "cisco_ios_xr"
    remoteIdSubstring.EntityData.ParentYangName = "remote-id-substrings"
    remoteIdSubstring.EntityData.SegmentPath = "remote-id-substring" + "[name='" + fmt.Sprintf("%v", remoteIdSubstring.Name) + "']"
    remoteIdSubstring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteIdSubstring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteIdSubstring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteIdSubstring.EntityData.Children = make(map[string]types.YChild)
    remoteIdSubstring.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteIdSubstring.EntityData.Leafs["name"] = types.YLeaf{"Name", remoteIdSubstring.Name}
    remoteIdSubstring.EntityData.Leafs["delay"] = types.YLeaf{"Delay", remoteIdSubstring.Delay}
    return &(remoteIdSubstring.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings
// Delay the PADO response when there is an
// exact match on the received Remote ID
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay the PADO response when there is an exact match on the received Remote
    // ID. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString.
    RemoteIdString []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString
}

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetEntityData() *types.CommonEntityData {
    remoteIdStrings.EntityData.YFilter = remoteIdStrings.YFilter
    remoteIdStrings.EntityData.YangName = "remote-id-strings"
    remoteIdStrings.EntityData.BundleName = "cisco_ios_xr"
    remoteIdStrings.EntityData.ParentYangName = "pa-do-delay"
    remoteIdStrings.EntityData.SegmentPath = "remote-id-strings"
    remoteIdStrings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteIdStrings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteIdStrings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteIdStrings.EntityData.Children = make(map[string]types.YChild)
    remoteIdStrings.EntityData.Children["remote-id-string"] = types.YChild{"RemoteIdString", nil}
    for i := range remoteIdStrings.RemoteIdString {
        remoteIdStrings.EntityData.Children[types.GetSegmentPath(&remoteIdStrings.RemoteIdString[i])] = types.YChild{"RemoteIdString", &remoteIdStrings.RemoteIdString[i]}
    }
    remoteIdStrings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(remoteIdStrings.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString
// Delay the PADO response when there is an
// exact match on the received Remote ID
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The string to exactly match the received Remote
    // ID. The type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetEntityData() *types.CommonEntityData {
    remoteIdString.EntityData.YFilter = remoteIdString.YFilter
    remoteIdString.EntityData.YangName = "remote-id-string"
    remoteIdString.EntityData.BundleName = "cisco_ios_xr"
    remoteIdString.EntityData.ParentYangName = "remote-id-strings"
    remoteIdString.EntityData.SegmentPath = "remote-id-string" + "[name='" + fmt.Sprintf("%v", remoteIdString.Name) + "']"
    remoteIdString.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteIdString.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteIdString.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteIdString.EntityData.Children = make(map[string]types.YChild)
    remoteIdString.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteIdString.EntityData.Leafs["name"] = types.YLeaf{"Name", remoteIdString.Name}
    remoteIdString.EntityData.Leafs["delay"] = types.YLeaf{"Delay", remoteIdString.Delay}
    return &(remoteIdString.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings
// Delay the PADO response when there is an
// exact match on the received Service Name
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay the PADO response when there is an exact match on the received
    // Service Name. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString.
    ServiceNameString []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString
}

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetEntityData() *types.CommonEntityData {
    serviceNameStrings.EntityData.YFilter = serviceNameStrings.YFilter
    serviceNameStrings.EntityData.YangName = "service-name-strings"
    serviceNameStrings.EntityData.BundleName = "cisco_ios_xr"
    serviceNameStrings.EntityData.ParentYangName = "pa-do-delay"
    serviceNameStrings.EntityData.SegmentPath = "service-name-strings"
    serviceNameStrings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceNameStrings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceNameStrings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceNameStrings.EntityData.Children = make(map[string]types.YChild)
    serviceNameStrings.EntityData.Children["service-name-string"] = types.YChild{"ServiceNameString", nil}
    for i := range serviceNameStrings.ServiceNameString {
        serviceNameStrings.EntityData.Children[types.GetSegmentPath(&serviceNameStrings.ServiceNameString[i])] = types.YChild{"ServiceNameString", &serviceNameStrings.ServiceNameString[i]}
    }
    serviceNameStrings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serviceNameStrings.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString
// Delay the PADO response when there is an
// exact match on the received Service Name
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The string to exactly match the received Service
    // Name. The type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetEntityData() *types.CommonEntityData {
    serviceNameString.EntityData.YFilter = serviceNameString.YFilter
    serviceNameString.EntityData.YangName = "service-name-string"
    serviceNameString.EntityData.BundleName = "cisco_ios_xr"
    serviceNameString.EntityData.ParentYangName = "service-name-strings"
    serviceNameString.EntityData.SegmentPath = "service-name-string" + "[name='" + fmt.Sprintf("%v", serviceNameString.Name) + "']"
    serviceNameString.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceNameString.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceNameString.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceNameString.EntityData.Children = make(map[string]types.YChild)
    serviceNameString.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceNameString.EntityData.Leafs["name"] = types.YLeaf{"Name", serviceNameString.Name}
    serviceNameString.EntityData.Leafs["delay"] = types.YLeaf{"Delay", serviceNameString.Delay}
    return &(serviceNameString.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings
// Delay the PADO response when the received
// Circuit ID contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay the PADO response when the received Circuit ID contains the given
    // string. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring.
    CircuitIdSubstring []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring
}

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetEntityData() *types.CommonEntityData {
    circuitIdSubstrings.EntityData.YFilter = circuitIdSubstrings.YFilter
    circuitIdSubstrings.EntityData.YangName = "circuit-id-substrings"
    circuitIdSubstrings.EntityData.BundleName = "cisco_ios_xr"
    circuitIdSubstrings.EntityData.ParentYangName = "pa-do-delay"
    circuitIdSubstrings.EntityData.SegmentPath = "circuit-id-substrings"
    circuitIdSubstrings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitIdSubstrings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitIdSubstrings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitIdSubstrings.EntityData.Children = make(map[string]types.YChild)
    circuitIdSubstrings.EntityData.Children["circuit-id-substring"] = types.YChild{"CircuitIdSubstring", nil}
    for i := range circuitIdSubstrings.CircuitIdSubstring {
        circuitIdSubstrings.EntityData.Children[types.GetSegmentPath(&circuitIdSubstrings.CircuitIdSubstring[i])] = types.YChild{"CircuitIdSubstring", &circuitIdSubstrings.CircuitIdSubstring[i]}
    }
    circuitIdSubstrings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(circuitIdSubstrings.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring
// Delay the PADO response when the received
// Circuit ID contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The string to be contained within the received
    // Circuit ID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetEntityData() *types.CommonEntityData {
    circuitIdSubstring.EntityData.YFilter = circuitIdSubstring.YFilter
    circuitIdSubstring.EntityData.YangName = "circuit-id-substring"
    circuitIdSubstring.EntityData.BundleName = "cisco_ios_xr"
    circuitIdSubstring.EntityData.ParentYangName = "circuit-id-substrings"
    circuitIdSubstring.EntityData.SegmentPath = "circuit-id-substring" + "[name='" + fmt.Sprintf("%v", circuitIdSubstring.Name) + "']"
    circuitIdSubstring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitIdSubstring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitIdSubstring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitIdSubstring.EntityData.Children = make(map[string]types.YChild)
    circuitIdSubstring.EntityData.Leafs = make(map[string]types.YLeaf)
    circuitIdSubstring.EntityData.Leafs["name"] = types.YLeaf{"Name", circuitIdSubstring.Name}
    circuitIdSubstring.EntityData.Leafs["delay"] = types.YLeaf{"Delay", circuitIdSubstring.Delay}
    return &(circuitIdSubstring.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings
// Delay the PADO response when the received
// Service Name contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay the PADO response when the received Service Name contains the given
    // string. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring.
    ServiceNameSubstring []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring
}

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetEntityData() *types.CommonEntityData {
    serviceNameSubstrings.EntityData.YFilter = serviceNameSubstrings.YFilter
    serviceNameSubstrings.EntityData.YangName = "service-name-substrings"
    serviceNameSubstrings.EntityData.BundleName = "cisco_ios_xr"
    serviceNameSubstrings.EntityData.ParentYangName = "pa-do-delay"
    serviceNameSubstrings.EntityData.SegmentPath = "service-name-substrings"
    serviceNameSubstrings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceNameSubstrings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceNameSubstrings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceNameSubstrings.EntityData.Children = make(map[string]types.YChild)
    serviceNameSubstrings.EntityData.Children["service-name-substring"] = types.YChild{"ServiceNameSubstring", nil}
    for i := range serviceNameSubstrings.ServiceNameSubstring {
        serviceNameSubstrings.EntityData.Children[types.GetSegmentPath(&serviceNameSubstrings.ServiceNameSubstring[i])] = types.YChild{"ServiceNameSubstring", &serviceNameSubstrings.ServiceNameSubstring[i]}
    }
    serviceNameSubstrings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serviceNameSubstrings.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring
// Delay the PADO response when the received
// Service Name contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The string to be contained within the received
    // Service Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetEntityData() *types.CommonEntityData {
    serviceNameSubstring.EntityData.YFilter = serviceNameSubstring.YFilter
    serviceNameSubstring.EntityData.YangName = "service-name-substring"
    serviceNameSubstring.EntityData.BundleName = "cisco_ios_xr"
    serviceNameSubstring.EntityData.ParentYangName = "service-name-substrings"
    serviceNameSubstring.EntityData.SegmentPath = "service-name-substring" + "[name='" + fmt.Sprintf("%v", serviceNameSubstring.Name) + "']"
    serviceNameSubstring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceNameSubstring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceNameSubstring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceNameSubstring.EntityData.Children = make(map[string]types.YChild)
    serviceNameSubstring.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceNameSubstring.EntityData.Leafs["name"] = types.YLeaf{"Name", serviceNameSubstring.Name}
    serviceNameSubstring.EntityData.Leafs["delay"] = types.YLeaf{"Delay", serviceNameSubstring.Delay}
    return &(serviceNameSubstring.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings
// Delay the PADO response when there is an
// exact match on the received Circuit ID
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay the PADO response when there is an exact match on the received
    // Circuit ID. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString.
    CircuitIdString []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString
}

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetEntityData() *types.CommonEntityData {
    circuitIdStrings.EntityData.YFilter = circuitIdStrings.YFilter
    circuitIdStrings.EntityData.YangName = "circuit-id-strings"
    circuitIdStrings.EntityData.BundleName = "cisco_ios_xr"
    circuitIdStrings.EntityData.ParentYangName = "pa-do-delay"
    circuitIdStrings.EntityData.SegmentPath = "circuit-id-strings"
    circuitIdStrings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitIdStrings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitIdStrings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitIdStrings.EntityData.Children = make(map[string]types.YChild)
    circuitIdStrings.EntityData.Children["circuit-id-string"] = types.YChild{"CircuitIdString", nil}
    for i := range circuitIdStrings.CircuitIdString {
        circuitIdStrings.EntityData.Children[types.GetSegmentPath(&circuitIdStrings.CircuitIdString[i])] = types.YChild{"CircuitIdString", &circuitIdStrings.CircuitIdString[i]}
    }
    circuitIdStrings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(circuitIdStrings.EntityData)
}

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString
// Delay the PADO response when there is an
// exact match on the received Circuit ID
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The string to exactly match the received Circuit
    // ID. The type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetEntityData() *types.CommonEntityData {
    circuitIdString.EntityData.YFilter = circuitIdString.YFilter
    circuitIdString.EntityData.YangName = "circuit-id-string"
    circuitIdString.EntityData.BundleName = "cisco_ios_xr"
    circuitIdString.EntityData.ParentYangName = "circuit-id-strings"
    circuitIdString.EntityData.SegmentPath = "circuit-id-string" + "[name='" + fmt.Sprintf("%v", circuitIdString.Name) + "']"
    circuitIdString.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    circuitIdString.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    circuitIdString.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    circuitIdString.EntityData.Children = make(map[string]types.YChild)
    circuitIdString.EntityData.Leafs = make(map[string]types.YLeaf)
    circuitIdString.EntityData.Leafs["name"] = types.YLeaf{"Name", circuitIdString.Name}
    circuitIdString.EntityData.Leafs["delay"] = types.YLeaf{"Delay", circuitIdString.Delay}
    return &(circuitIdString.EntityData)
}

