// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-bfd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   bfd: BFD Configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_bfd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_bfd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-bfd-cfg bfd}", reflect.TypeOf(Bfd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-bfd-cfg:bfd", reflect.TypeOf(Bfd{}))
}

// BfdEchoStartupValidate represents Bfd echo startup validate
type BfdEchoStartupValidate string

const (
    // Disable echo startup validation
    BfdEchoStartupValidate_off BfdEchoStartupValidate = "off"

    // Enable echo startup validation
    BfdEchoStartupValidate_on BfdEchoStartupValidate = "on"

    // Force echo startup validation
    BfdEchoStartupValidate_force BfdEchoStartupValidate = "force"
)

// BfdBundleCoexistenceBobBlb represents Bfd bundle coexistence bob blb
type BfdBundleCoexistenceBobBlb string

const (
    // Inherited coexistence mode
    BfdBundleCoexistenceBobBlb_inherited BfdBundleCoexistenceBobBlb = "inherited"

    // Logical coexistence mode
    BfdBundleCoexistenceBobBlb_logical BfdBundleCoexistenceBobBlb = "logical"
)

// BfdIfEchoUsage represents Bfd if echo usage
type BfdIfEchoUsage string

const (
    // Enable echo
    BfdIfEchoUsage_enable BfdIfEchoUsage = "enable"

    // Disable echo
    BfdIfEchoUsage_disable BfdIfEchoUsage = "disable"
)

// BfdIfIpv6ChecksumUsage represents Bfd if ipv6 checksum usage
type BfdIfIpv6ChecksumUsage string

const (
    // Disable IPv6 checksum
    BfdIfIpv6ChecksumUsage_disable BfdIfIpv6ChecksumUsage = "disable"

    // Enable IPv6 checksum
    BfdIfIpv6ChecksumUsage_enable BfdIfIpv6ChecksumUsage = "enable"
)

// Bfd
// BFD Configuration
type Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Echo configuration. The type is interface{}.
    GlobalEchoUsage interface{}

    // To disable IPv6 checksum configuration. The type is interface{}.
    Ipv6ChecksumDisable interface{}

    // Configure echo min-interval for bundle interface. The type is interface{}
    // with range: 15..2000. Units are millisecond. The default value is 15.
    GlobalEchoMinInterval interface{}

    // Multihop TTL Drop Threshold. The type is interface{} with range: 0..254.
    TtlDropThreshold interface{}

    // Single hop trap configuration. The type is interface{}.
    SingleHopTrap interface{}

    // IPv4 echo source address configuration. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    GlobalIpv4EchoSource interface{}

    // Flapping class container.
    FlapDamp Bfd_FlapDamp

    // BFD echo latency feature class container.
    EchoLatency Bfd_EchoLatency

    // BFD echo startup feature class container.
    EchoStartup Bfd_EchoStartup

    // Interface configuration.
    Interfaces Bfd_Interfaces

    // Multipath Include configuration.
    MultiPathIncludes Bfd_MultiPathIncludes

    // Configuration related to BFD over Bundle.
    Bundle Bfd_Bundle
}

func (bfd *Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "Cisco-IOS-XR-ip-bfd-cfg"
    bfd.EntityData.SegmentPath = "Cisco-IOS-XR-ip-bfd-cfg:bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = make(map[string]types.YChild)
    bfd.EntityData.Children["flap-damp"] = types.YChild{"FlapDamp", &bfd.FlapDamp}
    bfd.EntityData.Children["echo-latency"] = types.YChild{"EchoLatency", &bfd.EchoLatency}
    bfd.EntityData.Children["echo-startup"] = types.YChild{"EchoStartup", &bfd.EchoStartup}
    bfd.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &bfd.Interfaces}
    bfd.EntityData.Children["multi-path-includes"] = types.YChild{"MultiPathIncludes", &bfd.MultiPathIncludes}
    bfd.EntityData.Children["bundle"] = types.YChild{"Bundle", &bfd.Bundle}
    bfd.EntityData.Leafs = make(map[string]types.YLeaf)
    bfd.EntityData.Leafs["global-echo-usage"] = types.YLeaf{"GlobalEchoUsage", bfd.GlobalEchoUsage}
    bfd.EntityData.Leafs["ipv6-checksum-disable"] = types.YLeaf{"Ipv6ChecksumDisable", bfd.Ipv6ChecksumDisable}
    bfd.EntityData.Leafs["global-echo-min-interval"] = types.YLeaf{"GlobalEchoMinInterval", bfd.GlobalEchoMinInterval}
    bfd.EntityData.Leafs["ttl-drop-threshold"] = types.YLeaf{"TtlDropThreshold", bfd.TtlDropThreshold}
    bfd.EntityData.Leafs["single-hop-trap"] = types.YLeaf{"SingleHopTrap", bfd.SingleHopTrap}
    bfd.EntityData.Leafs["global-ipv4-echo-source"] = types.YLeaf{"GlobalIpv4EchoSource", bfd.GlobalIpv4EchoSource}
    return &(bfd.EntityData)
}

// Bfd_FlapDamp
// Flapping class container
type Bfd_FlapDamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stability threshold to enable dampening. The type is interface{} with
    // range: 60000..3600000. Units are millisecond. The default value is 120000.
    Threshold interface{}

    // Initial delay before bringing up session. The type is interface{} with
    // range: 1..3600000. Units are millisecond. The default value is 2000.
    InitialDelay interface{}

    // Maximum delay before bringing up session. The type is interface{} with
    // range: 1..3600000. Units are millisecond. The default value is 120000.
    MaximumDelay interface{}

    // Dampening Enable/Disable. The type is interface{}.
    DampenDisable interface{}

    // Secondary delay before bringing up session. The type is interface{} with
    // range: 1..3600000. Units are millisecond. The default value is 5000.
    SecondaryDelay interface{}

    // Bundle per member feature class container.
    BundleMember Bfd_FlapDamp_BundleMember

    // Extensions to the BFD dampening feature.
    Extensions Bfd_FlapDamp_Extensions
}

func (flapDamp *Bfd_FlapDamp) GetEntityData() *types.CommonEntityData {
    flapDamp.EntityData.YFilter = flapDamp.YFilter
    flapDamp.EntityData.YangName = "flap-damp"
    flapDamp.EntityData.BundleName = "cisco_ios_xr"
    flapDamp.EntityData.ParentYangName = "bfd"
    flapDamp.EntityData.SegmentPath = "flap-damp"
    flapDamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flapDamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flapDamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flapDamp.EntityData.Children = make(map[string]types.YChild)
    flapDamp.EntityData.Children["bundle-member"] = types.YChild{"BundleMember", &flapDamp.BundleMember}
    flapDamp.EntityData.Children["extensions"] = types.YChild{"Extensions", &flapDamp.Extensions}
    flapDamp.EntityData.Leafs = make(map[string]types.YLeaf)
    flapDamp.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", flapDamp.Threshold}
    flapDamp.EntityData.Leafs["initial-delay"] = types.YLeaf{"InitialDelay", flapDamp.InitialDelay}
    flapDamp.EntityData.Leafs["maximum-delay"] = types.YLeaf{"MaximumDelay", flapDamp.MaximumDelay}
    flapDamp.EntityData.Leafs["dampen-disable"] = types.YLeaf{"DampenDisable", flapDamp.DampenDisable}
    flapDamp.EntityData.Leafs["secondary-delay"] = types.YLeaf{"SecondaryDelay", flapDamp.SecondaryDelay}
    return &(flapDamp.EntityData)
}

// Bfd_FlapDamp_BundleMember
// Bundle per member feature class container
type Bfd_FlapDamp_BundleMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Initial delay before bringing up session. The type is interface{} with
    // range: 1..518400000. Units are millisecond. The default value is 16000.
    InitialDelay interface{}

    // Maximum delay before bringing up session. The type is interface{} with
    // range: 1..518400000. Units are millisecond. The default value is 600000.
    MaximumDelay interface{}

    // Secondary delay before bringing up session. The type is interface{} with
    // range: 1..518400000. Units are millisecond. The default value is 20000.
    SecondaryDelay interface{}

    // Apply immediate dampening and only when failure is L3 specific. The type is
    // interface{}.
    L3OnlyMode interface{}
}

func (bundleMember *Bfd_FlapDamp_BundleMember) GetEntityData() *types.CommonEntityData {
    bundleMember.EntityData.YFilter = bundleMember.YFilter
    bundleMember.EntityData.YangName = "bundle-member"
    bundleMember.EntityData.BundleName = "cisco_ios_xr"
    bundleMember.EntityData.ParentYangName = "flap-damp"
    bundleMember.EntityData.SegmentPath = "bundle-member"
    bundleMember.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleMember.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleMember.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleMember.EntityData.Children = make(map[string]types.YChild)
    bundleMember.EntityData.Leafs = make(map[string]types.YLeaf)
    bundleMember.EntityData.Leafs["initial-delay"] = types.YLeaf{"InitialDelay", bundleMember.InitialDelay}
    bundleMember.EntityData.Leafs["maximum-delay"] = types.YLeaf{"MaximumDelay", bundleMember.MaximumDelay}
    bundleMember.EntityData.Leafs["secondary-delay"] = types.YLeaf{"SecondaryDelay", bundleMember.SecondaryDelay}
    bundleMember.EntityData.Leafs["l3-only-mode"] = types.YLeaf{"L3OnlyMode", bundleMember.L3OnlyMode}
    return &(bundleMember.EntityData)
}

// Bfd_FlapDamp_Extensions
// Extensions to the BFD dampening feature
type Bfd_FlapDamp_Extensions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set, DOWN state monitoring mode is enabled. The type is interface{}.
    DownMonitor interface{}
}

func (extensions *Bfd_FlapDamp_Extensions) GetEntityData() *types.CommonEntityData {
    extensions.EntityData.YFilter = extensions.YFilter
    extensions.EntityData.YangName = "extensions"
    extensions.EntityData.BundleName = "cisco_ios_xr"
    extensions.EntityData.ParentYangName = "flap-damp"
    extensions.EntityData.SegmentPath = "extensions"
    extensions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extensions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extensions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extensions.EntityData.Children = make(map[string]types.YChild)
    extensions.EntityData.Leafs = make(map[string]types.YLeaf)
    extensions.EntityData.Leafs["down-monitor"] = types.YLeaf{"DownMonitor", extensions.DownMonitor}
    return &(extensions.EntityData)
}

// Bfd_EchoLatency
// BFD echo latency feature class container
type Bfd_EchoLatency struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BFD echo latency detection.
    Detect Bfd_EchoLatency_Detect
}

func (echoLatency *Bfd_EchoLatency) GetEntityData() *types.CommonEntityData {
    echoLatency.EntityData.YFilter = echoLatency.YFilter
    echoLatency.EntityData.YangName = "echo-latency"
    echoLatency.EntityData.BundleName = "cisco_ios_xr"
    echoLatency.EntityData.ParentYangName = "bfd"
    echoLatency.EntityData.SegmentPath = "echo-latency"
    echoLatency.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoLatency.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoLatency.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoLatency.EntityData.Children = make(map[string]types.YChild)
    echoLatency.EntityData.Children["detect"] = types.YChild{"Detect", &echoLatency.Detect}
    echoLatency.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(echoLatency.EntityData)
}

// Bfd_EchoLatency_Detect
// BFD echo latency detection
type Bfd_EchoLatency_Detect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set, echo latency detect is enabled. The type is bool.
    LatencyDetectEnabled interface{}

    // Echo latency detect percentage. The type is interface{} with range:
    // 100..250. Units are percentage.
    LatencyDetectPercentage interface{}

    // Echo latency detect count. The type is interface{} with range: 1..10.
    LatencyDetectCount interface{}
}

func (detect *Bfd_EchoLatency_Detect) GetEntityData() *types.CommonEntityData {
    detect.EntityData.YFilter = detect.YFilter
    detect.EntityData.YangName = "detect"
    detect.EntityData.BundleName = "cisco_ios_xr"
    detect.EntityData.ParentYangName = "echo-latency"
    detect.EntityData.SegmentPath = "detect"
    detect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detect.EntityData.Children = make(map[string]types.YChild)
    detect.EntityData.Leafs = make(map[string]types.YLeaf)
    detect.EntityData.Leafs["latency-detect-enabled"] = types.YLeaf{"LatencyDetectEnabled", detect.LatencyDetectEnabled}
    detect.EntityData.Leafs["latency-detect-percentage"] = types.YLeaf{"LatencyDetectPercentage", detect.LatencyDetectPercentage}
    detect.EntityData.Leafs["latency-detect-count"] = types.YLeaf{"LatencyDetectCount", detect.LatencyDetectCount}
    return &(detect.EntityData)
}

// Bfd_EchoStartup
// BFD echo startup feature class container
type Bfd_EchoStartup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BFD echo validation prior to session bringup. The type is
    // BfdEchoStartupValidate. The default value is off.
    Validate interface{}
}

func (echoStartup *Bfd_EchoStartup) GetEntityData() *types.CommonEntityData {
    echoStartup.EntityData.YFilter = echoStartup.YFilter
    echoStartup.EntityData.YangName = "echo-startup"
    echoStartup.EntityData.BundleName = "cisco_ios_xr"
    echoStartup.EntityData.ParentYangName = "bfd"
    echoStartup.EntityData.SegmentPath = "echo-startup"
    echoStartup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoStartup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoStartup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoStartup.EntityData.Children = make(map[string]types.YChild)
    echoStartup.EntityData.Leafs = make(map[string]types.YLeaf)
    echoStartup.EntityData.Leafs["validate"] = types.YLeaf{"Validate", echoStartup.Validate}
    return &(echoStartup.EntityData)
}

// Bfd_Interfaces
// Interface configuration
type Bfd_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single interface configuration. The type is slice of
    // Bfd_Interfaces_Interface_.
    Interface_ []Bfd_Interfaces_Interface
}

func (interfaces *Bfd_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "bfd"
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

// Bfd_Interfaces_Interface
// Single interface configuration
type Bfd_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Echo usage for this interface. The type is BfdIfEchoUsage. The default
    // value is enable.
    InterfaceEchoUsage interface{}

    // IPv6 checksum usage for this interface - Interface config will always take
    // precedence over global config. The type is BfdIfIpv6ChecksumUsage. The
    // default value is enable.
    Ipv6Checksum interface{}

    // Interface IPv4 echo source address configuration. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    InterfaceIpv4EchoSource interface{}
}

func (self *Bfd_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface-echo-usage"] = types.YLeaf{"InterfaceEchoUsage", self.InterfaceEchoUsage}
    self.EntityData.Leafs["ipv6-checksum"] = types.YLeaf{"Ipv6Checksum", self.Ipv6Checksum}
    self.EntityData.Leafs["interface-ipv4-echo-source"] = types.YLeaf{"InterfaceIpv4EchoSource", self.InterfaceIpv4EchoSource}
    return &(self.EntityData)
}

// Bfd_MultiPathIncludes
// Multipath Include configuration
type Bfd_MultiPathIncludes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location configuration. The type is slice of
    // Bfd_MultiPathIncludes_MultiPathInclude.
    MultiPathInclude []Bfd_MultiPathIncludes_MultiPathInclude
}

func (multiPathIncludes *Bfd_MultiPathIncludes) GetEntityData() *types.CommonEntityData {
    multiPathIncludes.EntityData.YFilter = multiPathIncludes.YFilter
    multiPathIncludes.EntityData.YangName = "multi-path-includes"
    multiPathIncludes.EntityData.BundleName = "cisco_ios_xr"
    multiPathIncludes.EntityData.ParentYangName = "bfd"
    multiPathIncludes.EntityData.SegmentPath = "multi-path-includes"
    multiPathIncludes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multiPathIncludes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multiPathIncludes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multiPathIncludes.EntityData.Children = make(map[string]types.YChild)
    multiPathIncludes.EntityData.Children["multi-path-include"] = types.YChild{"MultiPathInclude", nil}
    for i := range multiPathIncludes.MultiPathInclude {
        multiPathIncludes.EntityData.Children[types.GetSegmentPath(&multiPathIncludes.MultiPathInclude[i])] = types.YChild{"MultiPathInclude", &multiPathIncludes.MultiPathInclude[i]}
    }
    multiPathIncludes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(multiPathIncludes.EntityData)
}

// Bfd_MultiPathIncludes_MultiPathInclude
// Location configuration
type Bfd_MultiPathIncludes_MultiPathInclude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Location interface{}
}

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetEntityData() *types.CommonEntityData {
    multiPathInclude.EntityData.YFilter = multiPathInclude.YFilter
    multiPathInclude.EntityData.YangName = "multi-path-include"
    multiPathInclude.EntityData.BundleName = "cisco_ios_xr"
    multiPathInclude.EntityData.ParentYangName = "multi-path-includes"
    multiPathInclude.EntityData.SegmentPath = "multi-path-include" + "[location='" + fmt.Sprintf("%v", multiPathInclude.Location) + "']"
    multiPathInclude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    multiPathInclude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    multiPathInclude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    multiPathInclude.EntityData.Children = make(map[string]types.YChild)
    multiPathInclude.EntityData.Leafs = make(map[string]types.YLeaf)
    multiPathInclude.EntityData.Leafs["location"] = types.YLeaf{"Location", multiPathInclude.Location}
    return &(multiPathInclude.EntityData)
}

// Bfd_Bundle
// Configuration related to BFD over Bundle
type Bfd_Bundle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Coexistence mode for BFD on bundle feature.
    Coexistence Bfd_Bundle_Coexistence
}

func (bundle *Bfd_Bundle) GetEntityData() *types.CommonEntityData {
    bundle.EntityData.YFilter = bundle.YFilter
    bundle.EntityData.YangName = "bundle"
    bundle.EntityData.BundleName = "cisco_ios_xr"
    bundle.EntityData.ParentYangName = "bfd"
    bundle.EntityData.SegmentPath = "bundle"
    bundle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundle.EntityData.Children = make(map[string]types.YChild)
    bundle.EntityData.Children["coexistence"] = types.YChild{"Coexistence", &bundle.Coexistence}
    bundle.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundle.EntityData)
}

// Bfd_Bundle_Coexistence
// Coexistence mode for BFD on bundle feature
type Bfd_Bundle_Coexistence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Coexistence mode for BoB and BLB feature. The type is
    // BfdBundleCoexistenceBobBlb.
    BobBlb interface{}
}

func (coexistence *Bfd_Bundle_Coexistence) GetEntityData() *types.CommonEntityData {
    coexistence.EntityData.YFilter = coexistence.YFilter
    coexistence.EntityData.YangName = "coexistence"
    coexistence.EntityData.BundleName = "cisco_ios_xr"
    coexistence.EntityData.ParentYangName = "bundle"
    coexistence.EntityData.SegmentPath = "coexistence"
    coexistence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coexistence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coexistence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coexistence.EntityData.Children = make(map[string]types.YChild)
    coexistence.EntityData.Leafs = make(map[string]types.YLeaf)
    coexistence.EntityData.Leafs["bob-blb"] = types.YLeaf{"BobBlb", coexistence.BobBlb}
    return &(coexistence.EntityData)
}

