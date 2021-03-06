// This module contains a collection of YANG definitions
// for Cisco IOS-XR dwdm-ui package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dwdm: DWDM operational data
//   vtxp: vtxp
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package dwdm_ui_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dwdm_ui_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dwdm-ui-oper dwdm}", reflect.TypeOf(Dwdm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dwdm-ui-oper:dwdm", reflect.TypeOf(Dwdm{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dwdm-ui-oper vtxp}", reflect.TypeOf(Vtxp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dwdm-ui-oper:vtxp", reflect.TypeOf(Vtxp{}))
}

// G709apsByte represents G709aps byte
type G709apsByte string

const (
    // No Protection
    G709apsByte_pp_no_protect G709apsByte = "pp-no-protect"

    // No Request
    G709apsByte_pp_no_request G709apsByte = "pp-no-request"

    // Regenerator Degrade
    G709apsByte_pp_regen_degrade G709apsByte = "pp-regen-degrade"

    // Signal Degrade
    G709apsByte_pp_sig_degrade G709apsByte = "pp-sig-degrade"

    // Maintenance Request
    G709apsByte_pp_remote_main G709apsByte = "pp-remote-main"

    // Unknown
    G709apsByte_pp_aps_unknown G709apsByte = "pp-aps-unknown"
)

// G709ppintfState represents G709ppintf state
type G709ppintfState string

const (
    // Interface is Up
    G709ppintfState_pp_intf_up G709ppintfState = "pp-intf-up"

    // Interface is Going Down
    G709ppintfState_pp_intf_failing G709ppintfState = "pp-intf-failing"

    // Interface Down
    G709ppintfState_pp_intf_down G709ppintfState = "pp-intf-down"
)

// G709ppfsmState represents G709ppfsm state
type G709ppfsmState string

const (
    // In Active
    G709ppfsmState_in_active G709ppfsmState = "in-active"

    // Disabled
    G709ppfsmState_disabled G709ppfsmState = "disabled"

    // Normal
    G709ppfsmState_normal_state G709ppfsmState = "normal-state"

    // Local Failing
    G709ppfsmState_local_failing G709ppfsmState = "local-failing"

    // Remote Failing
    G709ppfsmState_remote_failing G709ppfsmState = "remote-failing"

    // Maintance Failing
    G709ppfsmState_main_t_failing G709ppfsmState = "main-t-failing"

    // Regenerator Failing
    G709ppfsmState_regen_failing G709ppfsmState = "regen-failing"

    // Local Failed
    G709ppfsmState_local_failed G709ppfsmState = "local-failed"

    // Remote Failed
    G709ppfsmState_remote_failed G709ppfsmState = "remote-failed"

    // Maintance Failed
    G709ppfsmState_main_t_failed G709ppfsmState = "main-t-failed"

    // Regenerator Failed
    G709ppfsmState_regen_failed G709ppfsmState = "regen-failed"
)

// G709ppfsmMode represents G709ppfsm mode
type G709ppfsmMode string

const (
    // OFF
    G709ppfsmMode_pp_disable G709ppfsmMode = "pp-disable"

    // ON (Default Mode)
    G709ppfsmMode_pp_default_mode G709ppfsmMode = "pp-default-mode"

    // ON (Graceful Mode)
    G709ppfsmMode_pp_graceful_mode G709ppfsmMode = "pp-graceful-mode"
)

// DwdmWaveChannelOwner represents Dwdm wave channel owner
type DwdmWaveChannelOwner string

const (
    // Hardware Default
    DwdmWaveChannelOwner_default_ DwdmWaveChannelOwner = "default"

    // Configuration
    DwdmWaveChannelOwner_configuration DwdmWaveChannelOwner = "configuration"

    // GMPLS Signaled
    DwdmWaveChannelOwner_gmpls DwdmWaveChannelOwner = "gmpls"
)

// G709efecMode represents G709efec mode
type G709efecMode string

const (
    // 
    G709efecMode_g975_none G709efecMode = "g975-none"

    // G975.1 I.4
    G709efecMode_g975_1_i4 G709efecMode = "g975-1-i4"

    // G975.1 I.7
    G709efecMode_g975_1_i7 G709efecMode = "g975-1-i7"
)

// DwdmtasState represents Dwdmtas state
type DwdmtasState string

const (
    // Out of Service
    DwdmtasState_tas_oos DwdmtasState = "tas-oos"

    // In Service
    DwdmtasState_tas_is DwdmtasState = "tas-is"

    // Out of Service Maintenance
    DwdmtasState_tas_oos_mt DwdmtasState = "tas-oos-mt"

    // In Service Config allowed
    DwdmtasState_tas_is_cfg DwdmtasState = "tas-is-cfg"
)

// DwdmControllerState represents Dwdm controller state
type DwdmControllerState string

const (
    // Up
    DwdmControllerState_dwdm_ui_state_up DwdmControllerState = "dwdm-ui-state-up"

    // Down
    DwdmControllerState_dwdm_ui_state_down DwdmControllerState = "dwdm-ui-state-down"

    // Administratively Down
    DwdmControllerState_dwdm_ui_state_admin_down DwdmControllerState = "dwdm-ui-state-admin-down"
)

// G709prbsMode represents G709prbs mode
type G709prbsMode string

const (
    // mode source
    G709prbsMode_mode_source G709prbsMode = "mode-source"

    // mode sink
    G709prbsMode_mode_sink G709prbsMode = "mode-sink"

    // mode source sink
    G709prbsMode_mode_source_sink G709prbsMode = "mode-source-sink"

    // mode invalid
    G709prbsMode_mode_invalid G709prbsMode = "mode-invalid"
)

// G709prbsPattern represents G709prbs pattern
type G709prbsPattern string

const (
    // pattern none
    G709prbsPattern_pattern_none G709prbsPattern = "pattern-none"

    // pattern null
    G709prbsPattern_pattern_null G709prbsPattern = "pattern-null"

    // pattern pn11
    G709prbsPattern_pattern_pn11 G709prbsPattern = "pattern-pn11"

    // pattern pn23
    G709prbsPattern_pattern_pn23 G709prbsPattern = "pattern-pn23"

    // pattern pn31
    G709prbsPattern_pattern_pn31 G709prbsPattern = "pattern-pn31"
)

// G709prbsInterval represents PRBS test interval information
type G709prbsInterval string

const (
    // Current interval
    G709prbsInterval_current_interval G709prbsInterval = "current-interval"

    // Previous interval
    G709prbsInterval_previous_interval G709prbsInterval = "previous-interval"

    // Previous interval 2
    G709prbsInterval_previous_interval2 G709prbsInterval = "previous-interval2"

    // Previous interval 3
    G709prbsInterval_previous_interval3 G709prbsInterval = "previous-interval3"

    // Previous interval 4
    G709prbsInterval_previous_interval4 G709prbsInterval = "previous-interval4"

    // Previous interval 5
    G709prbsInterval_previous_interval5 G709prbsInterval = "previous-interval5"

    // Previous interval 6
    G709prbsInterval_previous_interval6 G709prbsInterval = "previous-interval6"

    // Previous interval 7
    G709prbsInterval_previous_interval7 G709prbsInterval = "previous-interval7"

    // Previous interval 8
    G709prbsInterval_previous_interval8 G709prbsInterval = "previous-interval8"

    // Previous interval 9
    G709prbsInterval_previous_interval9 G709prbsInterval = "previous-interval9"

    // Previous interval 10
    G709prbsInterval_previous_interval10 G709prbsInterval = "previous-interval10"

    // Previous interval 11
    G709prbsInterval_previous_interval11 G709prbsInterval = "previous-interval11"

    // Previous interval 12
    G709prbsInterval_previous_interval12 G709prbsInterval = "previous-interval12"

    // Previous interval 13
    G709prbsInterval_previous_interval13 G709prbsInterval = "previous-interval13"

    // Previous interval 14
    G709prbsInterval_previous_interval14 G709prbsInterval = "previous-interval14"

    // Previous interval 15
    G709prbsInterval_previous_interval15 G709prbsInterval = "previous-interval15"

    // Previous interval 16
    G709prbsInterval_previous_interval16 G709prbsInterval = "previous-interval16"

    // Previous interval 17
    G709prbsInterval_previous_interval17 G709prbsInterval = "previous-interval17"

    // Previous interval 18
    G709prbsInterval_previous_interval18 G709prbsInterval = "previous-interval18"

    // Previous interval 19
    G709prbsInterval_previous_interval19 G709prbsInterval = "previous-interval19"

    // Previous interval 20
    G709prbsInterval_previous_interval20 G709prbsInterval = "previous-interval20"

    // Previous interval 21
    G709prbsInterval_previous_interval21 G709prbsInterval = "previous-interval21"

    // Previous interval 22
    G709prbsInterval_previous_interval22 G709prbsInterval = "previous-interval22"

    // Previous interval 23
    G709prbsInterval_previous_interval23 G709prbsInterval = "previous-interval23"

    // Previous interval 24
    G709prbsInterval_previous_interval24 G709prbsInterval = "previous-interval24"

    // Previous interval 25
    G709prbsInterval_previous_interval25 G709prbsInterval = "previous-interval25"

    // Previous interval 26
    G709prbsInterval_previous_interval26 G709prbsInterval = "previous-interval26"

    // Previous interval 27
    G709prbsInterval_previous_interval27 G709prbsInterval = "previous-interval27"

    // Previous interval 28
    G709prbsInterval_previous_interval28 G709prbsInterval = "previous-interval28"

    // Previous interval 29
    G709prbsInterval_previous_interval29 G709prbsInterval = "previous-interval29"

    // Previous interval 30
    G709prbsInterval_previous_interval30 G709prbsInterval = "previous-interval30"

    // Previous interval 31
    G709prbsInterval_previous_interval31 G709prbsInterval = "previous-interval31"

    // Previous interval 32
    G709prbsInterval_previous_interval32 G709prbsInterval = "previous-interval32"
)

// Dwdm
// DWDM operational data
type Dwdm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All DWDM Port operational data.
    Ports Dwdm_Ports
}

func (dwdm *Dwdm) GetEntityData() *types.CommonEntityData {
    dwdm.EntityData.YFilter = dwdm.YFilter
    dwdm.EntityData.YangName = "dwdm"
    dwdm.EntityData.BundleName = "cisco_ios_xr"
    dwdm.EntityData.ParentYangName = "Cisco-IOS-XR-dwdm-ui-oper"
    dwdm.EntityData.SegmentPath = "Cisco-IOS-XR-dwdm-ui-oper:dwdm"
    dwdm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dwdm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dwdm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dwdm.EntityData.Children = make(map[string]types.YChild)
    dwdm.EntityData.Children["ports"] = types.YChild{"Ports", &dwdm.Ports}
    dwdm.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dwdm.EntityData)
}

// Dwdm_Ports
// All DWDM Port operational data
type Dwdm_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DWDM Port operational data. The type is slice of Dwdm_Ports_Port.
    Port []Dwdm_Ports_Port
}

func (ports *Dwdm_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "cisco_ios_xr"
    ports.EntityData.ParentYangName = "dwdm"
    ports.EntityData.SegmentPath = "ports"
    ports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ports.EntityData.Children = make(map[string]types.YChild)
    ports.EntityData.Children["port"] = types.YChild{"Port", nil}
    for i := range ports.Port {
        ports.EntityData.Children[types.GetSegmentPath(&ports.Port[i])] = types.YChild{"Port", &ports.Port[i]}
    }
    ports.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ports.EntityData)
}

// Dwdm_Ports_Port
// DWDM Port operational data
type Dwdm_Ports_Port struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // DWDM Port PRBS related data.
    Prbs Dwdm_Ports_Port_Prbs

    // DWDM Port optics operational data.
    Optics Dwdm_Ports_Port_Optics

    // DWDM port operational data.
    Info Dwdm_Ports_Port_Info
}

func (port *Dwdm_Ports_Port) GetEntityData() *types.CommonEntityData {
    port.EntityData.YFilter = port.YFilter
    port.EntityData.YangName = "port"
    port.EntityData.BundleName = "cisco_ios_xr"
    port.EntityData.ParentYangName = "ports"
    port.EntityData.SegmentPath = "port" + "[name='" + fmt.Sprintf("%v", port.Name) + "']"
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = make(map[string]types.YChild)
    port.EntityData.Children["prbs"] = types.YChild{"Prbs", &port.Prbs}
    port.EntityData.Children["optics"] = types.YChild{"Optics", &port.Optics}
    port.EntityData.Children["info"] = types.YChild{"Info", &port.Info}
    port.EntityData.Leafs = make(map[string]types.YLeaf)
    port.EntityData.Leafs["name"] = types.YLeaf{"Name", port.Name}
    return &(port.EntityData)
}

// Dwdm_Ports_Port_Prbs
// DWDM Port PRBS related data
type Dwdm_Ports_Port_Prbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port 24-hour PRBS statistics table.
    TwentyFourHoursBucket Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket

    // Port 15-minute PRBS statistics table.
    FifteenMinutesBucket Dwdm_Ports_Port_Prbs_FifteenMinutesBucket
}

func (prbs *Dwdm_Ports_Port_Prbs) GetEntityData() *types.CommonEntityData {
    prbs.EntityData.YFilter = prbs.YFilter
    prbs.EntityData.YangName = "prbs"
    prbs.EntityData.BundleName = "cisco_ios_xr"
    prbs.EntityData.ParentYangName = "port"
    prbs.EntityData.SegmentPath = "prbs"
    prbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prbs.EntityData.Children = make(map[string]types.YChild)
    prbs.EntityData.Children["twenty-four-hours-bucket"] = types.YChild{"TwentyFourHoursBucket", &prbs.TwentyFourHoursBucket}
    prbs.EntityData.Children["fifteen-minutes-bucket"] = types.YChild{"FifteenMinutesBucket", &prbs.FifteenMinutesBucket}
    prbs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prbs.EntityData)
}

// Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket
// Port 24-hour PRBS statistics table
type Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port 24-hour PRBS statistics data.
    TwentyFourHoursStatistics Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics
}

func (twentyFourHoursBucket *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket) GetEntityData() *types.CommonEntityData {
    twentyFourHoursBucket.EntityData.YFilter = twentyFourHoursBucket.YFilter
    twentyFourHoursBucket.EntityData.YangName = "twenty-four-hours-bucket"
    twentyFourHoursBucket.EntityData.BundleName = "cisco_ios_xr"
    twentyFourHoursBucket.EntityData.ParentYangName = "prbs"
    twentyFourHoursBucket.EntityData.SegmentPath = "twenty-four-hours-bucket"
    twentyFourHoursBucket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twentyFourHoursBucket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twentyFourHoursBucket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twentyFourHoursBucket.EntityData.Children = make(map[string]types.YChild)
    twentyFourHoursBucket.EntityData.Children["twenty-four-hours-statistics"] = types.YChild{"TwentyFourHoursStatistics", &twentyFourHoursBucket.TwentyFourHoursStatistics}
    twentyFourHoursBucket.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(twentyFourHoursBucket.EntityData)
}

// Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics
// Port 24-hour PRBS statistics data
type Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 'True' if PRBS is enabled 'False' otherwise. The type is bool.
    IsPrbsEnabled interface{}

    // Configured mode of PRBS test. The type is G709prbsMode.
    PrbsConfigMode interface{}

    // History consists of 15-minute/24-hour intervals. The type is slice of
    // Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry.
    PrbsEntry []Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry
}

func (twentyFourHoursStatistics *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics) GetEntityData() *types.CommonEntityData {
    twentyFourHoursStatistics.EntityData.YFilter = twentyFourHoursStatistics.YFilter
    twentyFourHoursStatistics.EntityData.YangName = "twenty-four-hours-statistics"
    twentyFourHoursStatistics.EntityData.BundleName = "cisco_ios_xr"
    twentyFourHoursStatistics.EntityData.ParentYangName = "twenty-four-hours-bucket"
    twentyFourHoursStatistics.EntityData.SegmentPath = "twenty-four-hours-statistics"
    twentyFourHoursStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twentyFourHoursStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twentyFourHoursStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twentyFourHoursStatistics.EntityData.Children = make(map[string]types.YChild)
    twentyFourHoursStatistics.EntityData.Children["prbs-entry"] = types.YChild{"PrbsEntry", nil}
    for i := range twentyFourHoursStatistics.PrbsEntry {
        twentyFourHoursStatistics.EntityData.Children[types.GetSegmentPath(&twentyFourHoursStatistics.PrbsEntry[i])] = types.YChild{"PrbsEntry", &twentyFourHoursStatistics.PrbsEntry[i]}
    }
    twentyFourHoursStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    twentyFourHoursStatistics.EntityData.Leafs["is-prbs-enabled"] = types.YLeaf{"IsPrbsEnabled", twentyFourHoursStatistics.IsPrbsEnabled}
    twentyFourHoursStatistics.EntityData.Leafs["prbs-config-mode"] = types.YLeaf{"PrbsConfigMode", twentyFourHoursStatistics.PrbsConfigMode}
    return &(twentyFourHoursStatistics.EntityData)
}

// Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry
// History consists of 15-minute/24-hour intervals
type Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index of bucket, current and previous. The type is G709prbsInterval.
    IntervalIndex interface{}

    // Configured pattern of PRBS test. The type is G709prbsPattern.
    ConfiguredPattern interface{}

    // Interval start timestamp. The type is string with length: 0..64.
    StartAt interface{}

    // Interval stop timestamp. The type is string with length: 0..64.
    StopAt interface{}

    // Received Pattern of PRBS Test. The type is G709prbsPattern.
    ReceivedPattern interface{}

    // Bit Error Count. The type is interface{} with range:
    // 0..18446744073709551615.
    BitErrorCount interface{}

    // Count of pattern found in interval. The type is interface{} with range:
    // 0..18446744073709551615.
    FoundCount interface{}

    // Count of pattern lost in interval. The type is interface{} with range:
    // 0..18446744073709551615.
    LostCount interface{}

    // Pattern first found at timestamp. The type is string with length: 0..64.
    FoundAt interface{}

    // Pattern first lost at timestamp. The type is string with length: 0..64.
    LostAt interface{}
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_TwentyFourHoursBucket_TwentyFourHoursStatistics_PrbsEntry) GetEntityData() *types.CommonEntityData {
    prbsEntry.EntityData.YFilter = prbsEntry.YFilter
    prbsEntry.EntityData.YangName = "prbs-entry"
    prbsEntry.EntityData.BundleName = "cisco_ios_xr"
    prbsEntry.EntityData.ParentYangName = "twenty-four-hours-statistics"
    prbsEntry.EntityData.SegmentPath = "prbs-entry"
    prbsEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prbsEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prbsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prbsEntry.EntityData.Children = make(map[string]types.YChild)
    prbsEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    prbsEntry.EntityData.Leafs["interval-index"] = types.YLeaf{"IntervalIndex", prbsEntry.IntervalIndex}
    prbsEntry.EntityData.Leafs["configured-pattern"] = types.YLeaf{"ConfiguredPattern", prbsEntry.ConfiguredPattern}
    prbsEntry.EntityData.Leafs["start-at"] = types.YLeaf{"StartAt", prbsEntry.StartAt}
    prbsEntry.EntityData.Leafs["stop-at"] = types.YLeaf{"StopAt", prbsEntry.StopAt}
    prbsEntry.EntityData.Leafs["received-pattern"] = types.YLeaf{"ReceivedPattern", prbsEntry.ReceivedPattern}
    prbsEntry.EntityData.Leafs["bit-error-count"] = types.YLeaf{"BitErrorCount", prbsEntry.BitErrorCount}
    prbsEntry.EntityData.Leafs["found-count"] = types.YLeaf{"FoundCount", prbsEntry.FoundCount}
    prbsEntry.EntityData.Leafs["lost-count"] = types.YLeaf{"LostCount", prbsEntry.LostCount}
    prbsEntry.EntityData.Leafs["found-at"] = types.YLeaf{"FoundAt", prbsEntry.FoundAt}
    prbsEntry.EntityData.Leafs["lost-at"] = types.YLeaf{"LostAt", prbsEntry.LostAt}
    return &(prbsEntry.EntityData)
}

// Dwdm_Ports_Port_Prbs_FifteenMinutesBucket
// Port 15-minute PRBS statistics table
type Dwdm_Ports_Port_Prbs_FifteenMinutesBucket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port 15-minute PRBS statistics data.
    FifteenMinutesStatistics Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics
}

func (fifteenMinutesBucket *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket) GetEntityData() *types.CommonEntityData {
    fifteenMinutesBucket.EntityData.YFilter = fifteenMinutesBucket.YFilter
    fifteenMinutesBucket.EntityData.YangName = "fifteen-minutes-bucket"
    fifteenMinutesBucket.EntityData.BundleName = "cisco_ios_xr"
    fifteenMinutesBucket.EntityData.ParentYangName = "prbs"
    fifteenMinutesBucket.EntityData.SegmentPath = "fifteen-minutes-bucket"
    fifteenMinutesBucket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fifteenMinutesBucket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fifteenMinutesBucket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fifteenMinutesBucket.EntityData.Children = make(map[string]types.YChild)
    fifteenMinutesBucket.EntityData.Children["fifteen-minutes-statistics"] = types.YChild{"FifteenMinutesStatistics", &fifteenMinutesBucket.FifteenMinutesStatistics}
    fifteenMinutesBucket.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fifteenMinutesBucket.EntityData)
}

// Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics
// Port 15-minute PRBS statistics data
type Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 'True' if PRBS is enabled 'False' otherwise. The type is bool.
    IsPrbsEnabled interface{}

    // Configured mode of PRBS test. The type is G709prbsMode.
    PrbsConfigMode interface{}

    // History consists of 15-minute/24-hour intervals. The type is slice of
    // Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry.
    PrbsEntry []Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry
}

func (fifteenMinutesStatistics *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics) GetEntityData() *types.CommonEntityData {
    fifteenMinutesStatistics.EntityData.YFilter = fifteenMinutesStatistics.YFilter
    fifteenMinutesStatistics.EntityData.YangName = "fifteen-minutes-statistics"
    fifteenMinutesStatistics.EntityData.BundleName = "cisco_ios_xr"
    fifteenMinutesStatistics.EntityData.ParentYangName = "fifteen-minutes-bucket"
    fifteenMinutesStatistics.EntityData.SegmentPath = "fifteen-minutes-statistics"
    fifteenMinutesStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fifteenMinutesStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fifteenMinutesStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fifteenMinutesStatistics.EntityData.Children = make(map[string]types.YChild)
    fifteenMinutesStatistics.EntityData.Children["prbs-entry"] = types.YChild{"PrbsEntry", nil}
    for i := range fifteenMinutesStatistics.PrbsEntry {
        fifteenMinutesStatistics.EntityData.Children[types.GetSegmentPath(&fifteenMinutesStatistics.PrbsEntry[i])] = types.YChild{"PrbsEntry", &fifteenMinutesStatistics.PrbsEntry[i]}
    }
    fifteenMinutesStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    fifteenMinutesStatistics.EntityData.Leafs["is-prbs-enabled"] = types.YLeaf{"IsPrbsEnabled", fifteenMinutesStatistics.IsPrbsEnabled}
    fifteenMinutesStatistics.EntityData.Leafs["prbs-config-mode"] = types.YLeaf{"PrbsConfigMode", fifteenMinutesStatistics.PrbsConfigMode}
    return &(fifteenMinutesStatistics.EntityData)
}

// Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry
// History consists of 15-minute/24-hour intervals
type Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index of bucket, current and previous. The type is G709prbsInterval.
    IntervalIndex interface{}

    // Configured pattern of PRBS test. The type is G709prbsPattern.
    ConfiguredPattern interface{}

    // Interval start timestamp. The type is string with length: 0..64.
    StartAt interface{}

    // Interval stop timestamp. The type is string with length: 0..64.
    StopAt interface{}

    // Received Pattern of PRBS Test. The type is G709prbsPattern.
    ReceivedPattern interface{}

    // Bit Error Count. The type is interface{} with range:
    // 0..18446744073709551615.
    BitErrorCount interface{}

    // Count of pattern found in interval. The type is interface{} with range:
    // 0..18446744073709551615.
    FoundCount interface{}

    // Count of pattern lost in interval. The type is interface{} with range:
    // 0..18446744073709551615.
    LostCount interface{}

    // Pattern first found at timestamp. The type is string with length: 0..64.
    FoundAt interface{}

    // Pattern first lost at timestamp. The type is string with length: 0..64.
    LostAt interface{}
}

func (prbsEntry *Dwdm_Ports_Port_Prbs_FifteenMinutesBucket_FifteenMinutesStatistics_PrbsEntry) GetEntityData() *types.CommonEntityData {
    prbsEntry.EntityData.YFilter = prbsEntry.YFilter
    prbsEntry.EntityData.YangName = "prbs-entry"
    prbsEntry.EntityData.BundleName = "cisco_ios_xr"
    prbsEntry.EntityData.ParentYangName = "fifteen-minutes-statistics"
    prbsEntry.EntityData.SegmentPath = "prbs-entry"
    prbsEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prbsEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prbsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prbsEntry.EntityData.Children = make(map[string]types.YChild)
    prbsEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    prbsEntry.EntityData.Leafs["interval-index"] = types.YLeaf{"IntervalIndex", prbsEntry.IntervalIndex}
    prbsEntry.EntityData.Leafs["configured-pattern"] = types.YLeaf{"ConfiguredPattern", prbsEntry.ConfiguredPattern}
    prbsEntry.EntityData.Leafs["start-at"] = types.YLeaf{"StartAt", prbsEntry.StartAt}
    prbsEntry.EntityData.Leafs["stop-at"] = types.YLeaf{"StopAt", prbsEntry.StopAt}
    prbsEntry.EntityData.Leafs["received-pattern"] = types.YLeaf{"ReceivedPattern", prbsEntry.ReceivedPattern}
    prbsEntry.EntityData.Leafs["bit-error-count"] = types.YLeaf{"BitErrorCount", prbsEntry.BitErrorCount}
    prbsEntry.EntityData.Leafs["found-count"] = types.YLeaf{"FoundCount", prbsEntry.FoundCount}
    prbsEntry.EntityData.Leafs["lost-count"] = types.YLeaf{"LostCount", prbsEntry.LostCount}
    prbsEntry.EntityData.Leafs["found-at"] = types.YLeaf{"FoundAt", prbsEntry.FoundAt}
    prbsEntry.EntityData.Leafs["lost-at"] = types.YLeaf{"LostAt", prbsEntry.LostAt}
    return &(prbsEntry.EntityData)
}

// Dwdm_Ports_Port_Optics
// DWDM Port optics operational data
type Dwdm_Ports_Port_Optics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DWDM port wavelength information data.
    WaveInfo Dwdm_Ports_Port_Optics_WaveInfo
}

func (optics *Dwdm_Ports_Port_Optics) GetEntityData() *types.CommonEntityData {
    optics.EntityData.YFilter = optics.YFilter
    optics.EntityData.YangName = "optics"
    optics.EntityData.BundleName = "cisco_ios_xr"
    optics.EntityData.ParentYangName = "port"
    optics.EntityData.SegmentPath = "optics"
    optics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optics.EntityData.Children = make(map[string]types.YChild)
    optics.EntityData.Children["wave-info"] = types.YChild{"WaveInfo", &optics.WaveInfo}
    optics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(optics.EntityData)
}

// Dwdm_Ports_Port_Optics_WaveInfo
// DWDM port wavelength information data
type Dwdm_Ports_Port_Optics_WaveInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Wavelength band. The type is interface{} with range: 0..4294967295.
    WaveBand interface{}

    // Lowest ITU wavelength channel number supported. The type is interface{}
    // with range: 0..4294967295.
    WaveChannelMin interface{}

    // Highest ITU wavelength channel number supported. The type is interface{}
    // with range: 0..4294967295.
    WaveChannelMax interface{}
}

func (waveInfo *Dwdm_Ports_Port_Optics_WaveInfo) GetEntityData() *types.CommonEntityData {
    waveInfo.EntityData.YFilter = waveInfo.YFilter
    waveInfo.EntityData.YangName = "wave-info"
    waveInfo.EntityData.BundleName = "cisco_ios_xr"
    waveInfo.EntityData.ParentYangName = "optics"
    waveInfo.EntityData.SegmentPath = "wave-info"
    waveInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    waveInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    waveInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    waveInfo.EntityData.Children = make(map[string]types.YChild)
    waveInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    waveInfo.EntityData.Leafs["wave-band"] = types.YLeaf{"WaveBand", waveInfo.WaveBand}
    waveInfo.EntityData.Leafs["wave-channel-min"] = types.YLeaf{"WaveChannelMin", waveInfo.WaveChannelMin}
    waveInfo.EntityData.Leafs["wave-channel-max"] = types.YLeaf{"WaveChannelMax", waveInfo.WaveChannelMax}
    return &(waveInfo.EntityData)
}

// Dwdm_Ports_Port_Info
// DWDM port operational data
type Dwdm_Ports_Port_Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DWDM controller state: Up, Down or Administratively Down. The type is
    // DwdmControllerState.
    ControllerState interface{}

    // DWDM controller TAS state: IS, OOS, OOS-MT or IS-CFG. The type is
    // DwdmtasState.
    TransportAdminState interface{}

    // DWDM port slice state Up/Down. The type is bool.
    SliceState interface{}

    // G709 operational information.
    G709Info Dwdm_Ports_Port_Info_G709Info

    // Optics operational information.
    OpticsInfo Dwdm_Ports_Port_Info_OpticsInfo

    // TDC operational information.
    TdcInfo Dwdm_Ports_Port_Info_TdcInfo

    // Network SRLG information.
    NetworkSrlgInfo Dwdm_Ports_Port_Info_NetworkSrlgInfo

    // Proactive protection information.
    Proactive Dwdm_Ports_Port_Info_Proactive

    // Signal log information.
    SignalLog Dwdm_Ports_Port_Info_SignalLog
}

func (info *Dwdm_Ports_Port_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "port"
    info.EntityData.SegmentPath = "info"
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = make(map[string]types.YChild)
    info.EntityData.Children["g709-info"] = types.YChild{"G709Info", &info.G709Info}
    info.EntityData.Children["optics-info"] = types.YChild{"OpticsInfo", &info.OpticsInfo}
    info.EntityData.Children["tdc-info"] = types.YChild{"TdcInfo", &info.TdcInfo}
    info.EntityData.Children["network-srlg-info"] = types.YChild{"NetworkSrlgInfo", &info.NetworkSrlgInfo}
    info.EntityData.Children["proactive"] = types.YChild{"Proactive", &info.Proactive}
    info.EntityData.Children["signal-log"] = types.YChild{"SignalLog", &info.SignalLog}
    info.EntityData.Leafs = make(map[string]types.YLeaf)
    info.EntityData.Leafs["controller-state"] = types.YLeaf{"ControllerState", info.ControllerState}
    info.EntityData.Leafs["transport-admin-state"] = types.YLeaf{"TransportAdminState", info.TransportAdminState}
    info.EntityData.Leafs["slice-state"] = types.YLeaf{"SliceState", info.SliceState}
    return &(info.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info
// G709 operational information
type Dwdm_Ports_Port_Info_G709Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is G709 framing enabled. The type is bool.
    IsG709Enabled interface{}

    // Is Operating FEC Mode Default. The type is bool.
    IsFecModeDefault interface{}

    // FEC information. The type is interface{} with range:
    // -2147483648..2147483647.
    FecMode interface{}

    // Remote FEC information. The type is interface{} with range:
    // -2147483648..2147483647.
    RemoteFecMode interface{}

    // EFEC information. The type is G709efecMode.
    EfecMode interface{}

    // Loopback information. The type is interface{} with range:
    // -2147483648..2147483647.
    LoopbackMode interface{}

    // Corrected bit error counter . The type is interface{} with range:
    // 0..18446744073709551615.
    Ec interface{}

    // FEC Corrected bit error accumulated counter. The type is interface{} with
    // range: 0..18446744073709551615.
    EcAccum interface{}

    // FEC Uncorrected words counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc interface{}

    // pre fec ber calculated. The type is interface{} with range:
    // 0..18446744073709551615.
    FecBer interface{}

    // pre fec ber calculated. The type is interface{} with range:
    // -2147483648..2147483647.
    FecBerMan interface{}

    // q value calculated. The type is interface{} with range:
    // 0..18446744073709551615.
    Q interface{}

    // q margin calculated. The type is interface{} with range:
    // 0..18446744073709551615.
    QMargin interface{}

    // FEC BER String . The type is string with length: 0..64.
    FeCstr interface{}

    // Q String . The type is string with length: 0..64.
    Qstr interface{}

    // QMargin String. The type is string with length: 0..64.
    QmarginStr interface{}

    // Network port ID. The type is string with length: 0..65.
    NetworkPortId interface{}

    // Network connection ID. The type is string with length: 0..65.
    NetworkConnId interface{}

    // 'true' if Prbs is enabled 'false' otherwise. The type is bool.
    IsPrbsEnabled interface{}

    // Configured mode of PRBS Test. The type is G709prbsMode.
    G709PrbsMode interface{}

    // Pattern of PRBS Test. The type is G709prbsPattern.
    G709PrbsPattern interface{}

    // Time stamp for prbs configuration. The type is interface{} with range:
    // 0..18446744073709551615.
    PrbsTimeStamp interface{}

    // FEC mismatch alarm.
    FecMismatch Dwdm_Ports_Port_Info_G709Info_FecMismatch

    // FEC Corrected bits TCA information.
    EcTca Dwdm_Ports_Port_Info_G709Info_EcTca

    // FEC uncorrected words TCA information.
    UcTca Dwdm_Ports_Port_Info_G709Info_UcTca

    // OTU layer information.
    OtuInfo Dwdm_Ports_Port_Info_G709Info_OtuInfo

    // ODU layer Information.
    OduInfo Dwdm_Ports_Port_Info_G709Info_OduInfo
}

func (g709Info *Dwdm_Ports_Port_Info_G709Info) GetEntityData() *types.CommonEntityData {
    g709Info.EntityData.YFilter = g709Info.YFilter
    g709Info.EntityData.YangName = "g709-info"
    g709Info.EntityData.BundleName = "cisco_ios_xr"
    g709Info.EntityData.ParentYangName = "info"
    g709Info.EntityData.SegmentPath = "g709-info"
    g709Info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    g709Info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    g709Info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    g709Info.EntityData.Children = make(map[string]types.YChild)
    g709Info.EntityData.Children["fec-mismatch"] = types.YChild{"FecMismatch", &g709Info.FecMismatch}
    g709Info.EntityData.Children["ec-tca"] = types.YChild{"EcTca", &g709Info.EcTca}
    g709Info.EntityData.Children["uc-tca"] = types.YChild{"UcTca", &g709Info.UcTca}
    g709Info.EntityData.Children["otu-info"] = types.YChild{"OtuInfo", &g709Info.OtuInfo}
    g709Info.EntityData.Children["odu-info"] = types.YChild{"OduInfo", &g709Info.OduInfo}
    g709Info.EntityData.Leafs = make(map[string]types.YLeaf)
    g709Info.EntityData.Leafs["is-g709-enabled"] = types.YLeaf{"IsG709Enabled", g709Info.IsG709Enabled}
    g709Info.EntityData.Leafs["is-fec-mode-default"] = types.YLeaf{"IsFecModeDefault", g709Info.IsFecModeDefault}
    g709Info.EntityData.Leafs["fec-mode"] = types.YLeaf{"FecMode", g709Info.FecMode}
    g709Info.EntityData.Leafs["remote-fec-mode"] = types.YLeaf{"RemoteFecMode", g709Info.RemoteFecMode}
    g709Info.EntityData.Leafs["efec-mode"] = types.YLeaf{"EfecMode", g709Info.EfecMode}
    g709Info.EntityData.Leafs["loopback-mode"] = types.YLeaf{"LoopbackMode", g709Info.LoopbackMode}
    g709Info.EntityData.Leafs["ec"] = types.YLeaf{"Ec", g709Info.Ec}
    g709Info.EntityData.Leafs["ec-accum"] = types.YLeaf{"EcAccum", g709Info.EcAccum}
    g709Info.EntityData.Leafs["uc"] = types.YLeaf{"Uc", g709Info.Uc}
    g709Info.EntityData.Leafs["fec-ber"] = types.YLeaf{"FecBer", g709Info.FecBer}
    g709Info.EntityData.Leafs["fec-ber-man"] = types.YLeaf{"FecBerMan", g709Info.FecBerMan}
    g709Info.EntityData.Leafs["q"] = types.YLeaf{"Q", g709Info.Q}
    g709Info.EntityData.Leafs["q-margin"] = types.YLeaf{"QMargin", g709Info.QMargin}
    g709Info.EntityData.Leafs["fe-cstr"] = types.YLeaf{"FeCstr", g709Info.FeCstr}
    g709Info.EntityData.Leafs["qstr"] = types.YLeaf{"Qstr", g709Info.Qstr}
    g709Info.EntityData.Leafs["qmargin-str"] = types.YLeaf{"QmarginStr", g709Info.QmarginStr}
    g709Info.EntityData.Leafs["network-port-id"] = types.YLeaf{"NetworkPortId", g709Info.NetworkPortId}
    g709Info.EntityData.Leafs["network-conn-id"] = types.YLeaf{"NetworkConnId", g709Info.NetworkConnId}
    g709Info.EntityData.Leafs["is-prbs-enabled"] = types.YLeaf{"IsPrbsEnabled", g709Info.IsPrbsEnabled}
    g709Info.EntityData.Leafs["g709-prbs-mode"] = types.YLeaf{"G709PrbsMode", g709Info.G709PrbsMode}
    g709Info.EntityData.Leafs["g709-prbs-pattern"] = types.YLeaf{"G709PrbsPattern", g709Info.G709PrbsPattern}
    g709Info.EntityData.Leafs["prbs-time-stamp"] = types.YLeaf{"PrbsTimeStamp", g709Info.PrbsTimeStamp}
    return &(g709Info.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_FecMismatch
// FEC mismatch alarm
type Dwdm_Ports_Port_Info_G709Info_FecMismatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (fecMismatch *Dwdm_Ports_Port_Info_G709Info_FecMismatch) GetEntityData() *types.CommonEntityData {
    fecMismatch.EntityData.YFilter = fecMismatch.YFilter
    fecMismatch.EntityData.YangName = "fec-mismatch"
    fecMismatch.EntityData.BundleName = "cisco_ios_xr"
    fecMismatch.EntityData.ParentYangName = "g709-info"
    fecMismatch.EntityData.SegmentPath = "fec-mismatch"
    fecMismatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fecMismatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fecMismatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fecMismatch.EntityData.Children = make(map[string]types.YChild)
    fecMismatch.EntityData.Leafs = make(map[string]types.YLeaf)
    fecMismatch.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", fecMismatch.ReportingEnabled}
    fecMismatch.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", fecMismatch.IsDetected}
    fecMismatch.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", fecMismatch.IsAsserted}
    fecMismatch.EntityData.Leafs["counter"] = types.YLeaf{"Counter", fecMismatch.Counter}
    return &(fecMismatch.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_EcTca
// FEC Corrected bits TCA information
type Dwdm_Ports_Port_Info_G709Info_EcTca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ecTca *Dwdm_Ports_Port_Info_G709Info_EcTca) GetEntityData() *types.CommonEntityData {
    ecTca.EntityData.YFilter = ecTca.YFilter
    ecTca.EntityData.YangName = "ec-tca"
    ecTca.EntityData.BundleName = "cisco_ios_xr"
    ecTca.EntityData.ParentYangName = "g709-info"
    ecTca.EntityData.SegmentPath = "ec-tca"
    ecTca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ecTca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ecTca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ecTca.EntityData.Children = make(map[string]types.YChild)
    ecTca.EntityData.Leafs = make(map[string]types.YLeaf)
    ecTca.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", ecTca.ReportingEnabled}
    ecTca.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", ecTca.IsDetected}
    ecTca.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", ecTca.IsAsserted}
    ecTca.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", ecTca.Threshold}
    ecTca.EntityData.Leafs["counter"] = types.YLeaf{"Counter", ecTca.Counter}
    return &(ecTca.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_UcTca
// FEC uncorrected words TCA information
type Dwdm_Ports_Port_Info_G709Info_UcTca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ucTca *Dwdm_Ports_Port_Info_G709Info_UcTca) GetEntityData() *types.CommonEntityData {
    ucTca.EntityData.YFilter = ucTca.YFilter
    ucTca.EntityData.YangName = "uc-tca"
    ucTca.EntityData.BundleName = "cisco_ios_xr"
    ucTca.EntityData.ParentYangName = "g709-info"
    ucTca.EntityData.SegmentPath = "uc-tca"
    ucTca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ucTca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ucTca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ucTca.EntityData.Children = make(map[string]types.YChild)
    ucTca.EntityData.Leafs = make(map[string]types.YLeaf)
    ucTca.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", ucTca.ReportingEnabled}
    ucTca.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", ucTca.IsDetected}
    ucTca.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", ucTca.IsAsserted}
    ucTca.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", ucTca.Threshold}
    ucTca.EntityData.Leafs["counter"] = types.YLeaf{"Counter", ucTca.Counter}
    return &(ucTca.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo
// OTU layer information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backward Error Indication counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Bei interface{}

    // Bit Interleave Parity(BIP) counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Bip interface{}

    // Loss of Signal information.
    Los Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los

    // Loss of Frame information.
    Lof Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof

    // Loss of MultiFrame information.
    Lom Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom

    // Out of Frame information.
    Oof Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof

    // Out of MultiFrame information.
    Oom Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom

    // Alarm Indication Signal information.
    Ais Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais

    // Incoming Alignment Error information.
    Iae Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae

    // Backward Defect Indication information.
    Bdi Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi

    // Trace Identifier Mismatch information.
    Tim Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim

    // GCC End of Channel information.
    Eoc Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc

    // Signal Fail  BER information.
    SfBer Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer

    // Signal Degrade BER information.
    SdBer Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer

    // Prefec Signal Fail BER information.
    PrefecSfBer Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer

    // Prefec Signal Degrade BER information.
    PrefecSdBer Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer

    // Backgound Block Error TCA information.
    BbeTca Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca

    // Errored Seconds TCA information.
    EsTca Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca

    // Backgound Block Error information.
    Bbe Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe

    // Errored Seconds information .
    Es Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es

    // Severly Errored Seconds information.
    Ses Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses

    // Unavailability Seconds information.
    Uas Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas

    // Failure Count information.
    Fc Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc

    // Backgound Block Error Rate information.
    Bber Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber

    // Errored Seconds Rate information.
    Esr Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr

    // Severly Errored Seconds Rate information.
    Sesr Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr

    // Trail Trace Identifier information.
    Tti Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti
}

func (otuInfo *Dwdm_Ports_Port_Info_G709Info_OtuInfo) GetEntityData() *types.CommonEntityData {
    otuInfo.EntityData.YFilter = otuInfo.YFilter
    otuInfo.EntityData.YangName = "otu-info"
    otuInfo.EntityData.BundleName = "cisco_ios_xr"
    otuInfo.EntityData.ParentYangName = "g709-info"
    otuInfo.EntityData.SegmentPath = "otu-info"
    otuInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otuInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otuInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otuInfo.EntityData.Children = make(map[string]types.YChild)
    otuInfo.EntityData.Children["los"] = types.YChild{"Los", &otuInfo.Los}
    otuInfo.EntityData.Children["lof"] = types.YChild{"Lof", &otuInfo.Lof}
    otuInfo.EntityData.Children["lom"] = types.YChild{"Lom", &otuInfo.Lom}
    otuInfo.EntityData.Children["oof"] = types.YChild{"Oof", &otuInfo.Oof}
    otuInfo.EntityData.Children["oom"] = types.YChild{"Oom", &otuInfo.Oom}
    otuInfo.EntityData.Children["ais"] = types.YChild{"Ais", &otuInfo.Ais}
    otuInfo.EntityData.Children["iae"] = types.YChild{"Iae", &otuInfo.Iae}
    otuInfo.EntityData.Children["bdi"] = types.YChild{"Bdi", &otuInfo.Bdi}
    otuInfo.EntityData.Children["tim"] = types.YChild{"Tim", &otuInfo.Tim}
    otuInfo.EntityData.Children["eoc"] = types.YChild{"Eoc", &otuInfo.Eoc}
    otuInfo.EntityData.Children["sf-ber"] = types.YChild{"SfBer", &otuInfo.SfBer}
    otuInfo.EntityData.Children["sd-ber"] = types.YChild{"SdBer", &otuInfo.SdBer}
    otuInfo.EntityData.Children["prefec-sf-ber"] = types.YChild{"PrefecSfBer", &otuInfo.PrefecSfBer}
    otuInfo.EntityData.Children["prefec-sd-ber"] = types.YChild{"PrefecSdBer", &otuInfo.PrefecSdBer}
    otuInfo.EntityData.Children["bbe-tca"] = types.YChild{"BbeTca", &otuInfo.BbeTca}
    otuInfo.EntityData.Children["es-tca"] = types.YChild{"EsTca", &otuInfo.EsTca}
    otuInfo.EntityData.Children["bbe"] = types.YChild{"Bbe", &otuInfo.Bbe}
    otuInfo.EntityData.Children["es"] = types.YChild{"Es", &otuInfo.Es}
    otuInfo.EntityData.Children["ses"] = types.YChild{"Ses", &otuInfo.Ses}
    otuInfo.EntityData.Children["uas"] = types.YChild{"Uas", &otuInfo.Uas}
    otuInfo.EntityData.Children["fc"] = types.YChild{"Fc", &otuInfo.Fc}
    otuInfo.EntityData.Children["bber"] = types.YChild{"Bber", &otuInfo.Bber}
    otuInfo.EntityData.Children["esr"] = types.YChild{"Esr", &otuInfo.Esr}
    otuInfo.EntityData.Children["sesr"] = types.YChild{"Sesr", &otuInfo.Sesr}
    otuInfo.EntityData.Children["tti"] = types.YChild{"Tti", &otuInfo.Tti}
    otuInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    otuInfo.EntityData.Leafs["bei"] = types.YLeaf{"Bei", otuInfo.Bei}
    otuInfo.EntityData.Leafs["bip"] = types.YLeaf{"Bip", otuInfo.Bip}
    return &(otuInfo.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los
// Loss of Signal information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (los *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Los) GetEntityData() *types.CommonEntityData {
    los.EntityData.YFilter = los.YFilter
    los.EntityData.YangName = "los"
    los.EntityData.BundleName = "cisco_ios_xr"
    los.EntityData.ParentYangName = "otu-info"
    los.EntityData.SegmentPath = "los"
    los.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    los.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    los.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    los.EntityData.Children = make(map[string]types.YChild)
    los.EntityData.Leafs = make(map[string]types.YLeaf)
    los.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", los.ReportingEnabled}
    los.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", los.IsDetected}
    los.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", los.IsAsserted}
    los.EntityData.Leafs["counter"] = types.YLeaf{"Counter", los.Counter}
    return &(los.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof
// Loss of Frame information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (lof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lof) GetEntityData() *types.CommonEntityData {
    lof.EntityData.YFilter = lof.YFilter
    lof.EntityData.YangName = "lof"
    lof.EntityData.BundleName = "cisco_ios_xr"
    lof.EntityData.ParentYangName = "otu-info"
    lof.EntityData.SegmentPath = "lof"
    lof.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lof.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lof.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lof.EntityData.Children = make(map[string]types.YChild)
    lof.EntityData.Leafs = make(map[string]types.YLeaf)
    lof.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", lof.ReportingEnabled}
    lof.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", lof.IsDetected}
    lof.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", lof.IsAsserted}
    lof.EntityData.Leafs["counter"] = types.YLeaf{"Counter", lof.Counter}
    return &(lof.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom
// Loss of MultiFrame information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (lom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Lom) GetEntityData() *types.CommonEntityData {
    lom.EntityData.YFilter = lom.YFilter
    lom.EntityData.YangName = "lom"
    lom.EntityData.BundleName = "cisco_ios_xr"
    lom.EntityData.ParentYangName = "otu-info"
    lom.EntityData.SegmentPath = "lom"
    lom.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lom.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lom.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lom.EntityData.Children = make(map[string]types.YChild)
    lom.EntityData.Leafs = make(map[string]types.YLeaf)
    lom.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", lom.ReportingEnabled}
    lom.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", lom.IsDetected}
    lom.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", lom.IsAsserted}
    lom.EntityData.Leafs["counter"] = types.YLeaf{"Counter", lom.Counter}
    return &(lom.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof
// Out of Frame information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (oof *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oof) GetEntityData() *types.CommonEntityData {
    oof.EntityData.YFilter = oof.YFilter
    oof.EntityData.YangName = "oof"
    oof.EntityData.BundleName = "cisco_ios_xr"
    oof.EntityData.ParentYangName = "otu-info"
    oof.EntityData.SegmentPath = "oof"
    oof.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oof.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oof.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oof.EntityData.Children = make(map[string]types.YChild)
    oof.EntityData.Leafs = make(map[string]types.YLeaf)
    oof.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", oof.ReportingEnabled}
    oof.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", oof.IsDetected}
    oof.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", oof.IsAsserted}
    oof.EntityData.Leafs["counter"] = types.YLeaf{"Counter", oof.Counter}
    return &(oof.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom
// Out of MultiFrame information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (oom *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Oom) GetEntityData() *types.CommonEntityData {
    oom.EntityData.YFilter = oom.YFilter
    oom.EntityData.YangName = "oom"
    oom.EntityData.BundleName = "cisco_ios_xr"
    oom.EntityData.ParentYangName = "otu-info"
    oom.EntityData.SegmentPath = "oom"
    oom.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oom.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oom.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oom.EntityData.Children = make(map[string]types.YChild)
    oom.EntityData.Leafs = make(map[string]types.YLeaf)
    oom.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", oom.ReportingEnabled}
    oom.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", oom.IsDetected}
    oom.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", oom.IsAsserted}
    oom.EntityData.Leafs["counter"] = types.YLeaf{"Counter", oom.Counter}
    return &(oom.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais
// Alarm Indication Signal information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ais) GetEntityData() *types.CommonEntityData {
    ais.EntityData.YFilter = ais.YFilter
    ais.EntityData.YangName = "ais"
    ais.EntityData.BundleName = "cisco_ios_xr"
    ais.EntityData.ParentYangName = "otu-info"
    ais.EntityData.SegmentPath = "ais"
    ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ais.EntityData.Children = make(map[string]types.YChild)
    ais.EntityData.Leafs = make(map[string]types.YLeaf)
    ais.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", ais.ReportingEnabled}
    ais.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", ais.IsDetected}
    ais.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", ais.IsAsserted}
    ais.EntityData.Leafs["counter"] = types.YLeaf{"Counter", ais.Counter}
    return &(ais.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae
// Incoming Alignment Error information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (iae *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Iae) GetEntityData() *types.CommonEntityData {
    iae.EntityData.YFilter = iae.YFilter
    iae.EntityData.YangName = "iae"
    iae.EntityData.BundleName = "cisco_ios_xr"
    iae.EntityData.ParentYangName = "otu-info"
    iae.EntityData.SegmentPath = "iae"
    iae.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iae.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iae.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iae.EntityData.Children = make(map[string]types.YChild)
    iae.EntityData.Leafs = make(map[string]types.YLeaf)
    iae.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", iae.ReportingEnabled}
    iae.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", iae.IsDetected}
    iae.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", iae.IsAsserted}
    iae.EntityData.Leafs["counter"] = types.YLeaf{"Counter", iae.Counter}
    return &(iae.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi
// Backward Defect Indication information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bdi) GetEntityData() *types.CommonEntityData {
    bdi.EntityData.YFilter = bdi.YFilter
    bdi.EntityData.YangName = "bdi"
    bdi.EntityData.BundleName = "cisco_ios_xr"
    bdi.EntityData.ParentYangName = "otu-info"
    bdi.EntityData.SegmentPath = "bdi"
    bdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdi.EntityData.Children = make(map[string]types.YChild)
    bdi.EntityData.Leafs = make(map[string]types.YLeaf)
    bdi.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", bdi.ReportingEnabled}
    bdi.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", bdi.IsDetected}
    bdi.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", bdi.IsAsserted}
    bdi.EntityData.Leafs["counter"] = types.YLeaf{"Counter", bdi.Counter}
    return &(bdi.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim
// Trace Identifier Mismatch information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tim) GetEntityData() *types.CommonEntityData {
    tim.EntityData.YFilter = tim.YFilter
    tim.EntityData.YangName = "tim"
    tim.EntityData.BundleName = "cisco_ios_xr"
    tim.EntityData.ParentYangName = "otu-info"
    tim.EntityData.SegmentPath = "tim"
    tim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tim.EntityData.Children = make(map[string]types.YChild)
    tim.EntityData.Leafs = make(map[string]types.YLeaf)
    tim.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", tim.ReportingEnabled}
    tim.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", tim.IsDetected}
    tim.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", tim.IsAsserted}
    tim.EntityData.Leafs["counter"] = types.YLeaf{"Counter", tim.Counter}
    return &(tim.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc
// GCC End of Channel information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Eoc) GetEntityData() *types.CommonEntityData {
    eoc.EntityData.YFilter = eoc.YFilter
    eoc.EntityData.YangName = "eoc"
    eoc.EntityData.BundleName = "cisco_ios_xr"
    eoc.EntityData.ParentYangName = "otu-info"
    eoc.EntityData.SegmentPath = "eoc"
    eoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eoc.EntityData.Children = make(map[string]types.YChild)
    eoc.EntityData.Leafs = make(map[string]types.YLeaf)
    eoc.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", eoc.ReportingEnabled}
    eoc.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", eoc.IsDetected}
    eoc.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", eoc.IsAsserted}
    eoc.EntityData.Leafs["counter"] = types.YLeaf{"Counter", eoc.Counter}
    return &(eoc.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer
// Signal Fail  BER information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SfBer) GetEntityData() *types.CommonEntityData {
    sfBer.EntityData.YFilter = sfBer.YFilter
    sfBer.EntityData.YangName = "sf-ber"
    sfBer.EntityData.BundleName = "cisco_ios_xr"
    sfBer.EntityData.ParentYangName = "otu-info"
    sfBer.EntityData.SegmentPath = "sf-ber"
    sfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfBer.EntityData.Children = make(map[string]types.YChild)
    sfBer.EntityData.Leafs = make(map[string]types.YLeaf)
    sfBer.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", sfBer.ReportingEnabled}
    sfBer.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", sfBer.IsDetected}
    sfBer.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", sfBer.IsAsserted}
    sfBer.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", sfBer.Threshold}
    sfBer.EntityData.Leafs["counter"] = types.YLeaf{"Counter", sfBer.Counter}
    return &(sfBer.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer
// Signal Degrade BER information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_SdBer) GetEntityData() *types.CommonEntityData {
    sdBer.EntityData.YFilter = sdBer.YFilter
    sdBer.EntityData.YangName = "sd-ber"
    sdBer.EntityData.BundleName = "cisco_ios_xr"
    sdBer.EntityData.ParentYangName = "otu-info"
    sdBer.EntityData.SegmentPath = "sd-ber"
    sdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdBer.EntityData.Children = make(map[string]types.YChild)
    sdBer.EntityData.Leafs = make(map[string]types.YLeaf)
    sdBer.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", sdBer.ReportingEnabled}
    sdBer.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", sdBer.IsDetected}
    sdBer.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", sdBer.IsAsserted}
    sdBer.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", sdBer.Threshold}
    sdBer.EntityData.Leafs["counter"] = types.YLeaf{"Counter", sdBer.Counter}
    return &(sdBer.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer
// Prefec Signal Fail BER information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (prefecSfBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSfBer) GetEntityData() *types.CommonEntityData {
    prefecSfBer.EntityData.YFilter = prefecSfBer.YFilter
    prefecSfBer.EntityData.YangName = "prefec-sf-ber"
    prefecSfBer.EntityData.BundleName = "cisco_ios_xr"
    prefecSfBer.EntityData.ParentYangName = "otu-info"
    prefecSfBer.EntityData.SegmentPath = "prefec-sf-ber"
    prefecSfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefecSfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefecSfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefecSfBer.EntityData.Children = make(map[string]types.YChild)
    prefecSfBer.EntityData.Leafs = make(map[string]types.YLeaf)
    prefecSfBer.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", prefecSfBer.ReportingEnabled}
    prefecSfBer.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", prefecSfBer.IsDetected}
    prefecSfBer.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", prefecSfBer.IsAsserted}
    prefecSfBer.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", prefecSfBer.Threshold}
    prefecSfBer.EntityData.Leafs["counter"] = types.YLeaf{"Counter", prefecSfBer.Counter}
    return &(prefecSfBer.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer
// Prefec Signal Degrade BER information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (prefecSdBer *Dwdm_Ports_Port_Info_G709Info_OtuInfo_PrefecSdBer) GetEntityData() *types.CommonEntityData {
    prefecSdBer.EntityData.YFilter = prefecSdBer.YFilter
    prefecSdBer.EntityData.YangName = "prefec-sd-ber"
    prefecSdBer.EntityData.BundleName = "cisco_ios_xr"
    prefecSdBer.EntityData.ParentYangName = "otu-info"
    prefecSdBer.EntityData.SegmentPath = "prefec-sd-ber"
    prefecSdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefecSdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefecSdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefecSdBer.EntityData.Children = make(map[string]types.YChild)
    prefecSdBer.EntityData.Leafs = make(map[string]types.YLeaf)
    prefecSdBer.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", prefecSdBer.ReportingEnabled}
    prefecSdBer.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", prefecSdBer.IsDetected}
    prefecSdBer.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", prefecSdBer.IsAsserted}
    prefecSdBer.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", prefecSdBer.Threshold}
    prefecSdBer.EntityData.Leafs["counter"] = types.YLeaf{"Counter", prefecSdBer.Counter}
    return &(prefecSdBer.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca
//  Backgound Block Error TCA information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_BbeTca) GetEntityData() *types.CommonEntityData {
    bbeTca.EntityData.YFilter = bbeTca.YFilter
    bbeTca.EntityData.YangName = "bbe-tca"
    bbeTca.EntityData.BundleName = "cisco_ios_xr"
    bbeTca.EntityData.ParentYangName = "otu-info"
    bbeTca.EntityData.SegmentPath = "bbe-tca"
    bbeTca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bbeTca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bbeTca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bbeTca.EntityData.Children = make(map[string]types.YChild)
    bbeTca.EntityData.Leafs = make(map[string]types.YLeaf)
    bbeTca.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", bbeTca.ReportingEnabled}
    bbeTca.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", bbeTca.IsDetected}
    bbeTca.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", bbeTca.IsAsserted}
    bbeTca.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", bbeTca.Threshold}
    bbeTca.EntityData.Leafs["counter"] = types.YLeaf{"Counter", bbeTca.Counter}
    return &(bbeTca.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca
// Errored Seconds TCA information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OtuInfo_EsTca) GetEntityData() *types.CommonEntityData {
    esTca.EntityData.YFilter = esTca.YFilter
    esTca.EntityData.YangName = "es-tca"
    esTca.EntityData.BundleName = "cisco_ios_xr"
    esTca.EntityData.ParentYangName = "otu-info"
    esTca.EntityData.SegmentPath = "es-tca"
    esTca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esTca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esTca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esTca.EntityData.Children = make(map[string]types.YChild)
    esTca.EntityData.Leafs = make(map[string]types.YLeaf)
    esTca.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", esTca.ReportingEnabled}
    esTca.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", esTca.IsDetected}
    esTca.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", esTca.IsAsserted}
    esTca.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", esTca.Threshold}
    esTca.EntityData.Leafs["counter"] = types.YLeaf{"Counter", esTca.Counter}
    return &(esTca.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe
// Backgound Block Error information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bbe) GetEntityData() *types.CommonEntityData {
    bbe.EntityData.YFilter = bbe.YFilter
    bbe.EntityData.YangName = "bbe"
    bbe.EntityData.BundleName = "cisco_ios_xr"
    bbe.EntityData.ParentYangName = "otu-info"
    bbe.EntityData.SegmentPath = "bbe"
    bbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bbe.EntityData.Children = make(map[string]types.YChild)
    bbe.EntityData.Leafs = make(map[string]types.YLeaf)
    bbe.EntityData.Leafs["counter"] = types.YLeaf{"Counter", bbe.Counter}
    return &(bbe.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es
// Errored Seconds information 
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (es *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Es) GetEntityData() *types.CommonEntityData {
    es.EntityData.YFilter = es.YFilter
    es.EntityData.YangName = "es"
    es.EntityData.BundleName = "cisco_ios_xr"
    es.EntityData.ParentYangName = "otu-info"
    es.EntityData.SegmentPath = "es"
    es.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    es.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    es.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    es.EntityData.Children = make(map[string]types.YChild)
    es.EntityData.Leafs = make(map[string]types.YLeaf)
    es.EntityData.Leafs["counter"] = types.YLeaf{"Counter", es.Counter}
    return &(es.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses
// Severly Errored Seconds information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Ses) GetEntityData() *types.CommonEntityData {
    ses.EntityData.YFilter = ses.YFilter
    ses.EntityData.YangName = "ses"
    ses.EntityData.BundleName = "cisco_ios_xr"
    ses.EntityData.ParentYangName = "otu-info"
    ses.EntityData.SegmentPath = "ses"
    ses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ses.EntityData.Children = make(map[string]types.YChild)
    ses.EntityData.Leafs = make(map[string]types.YLeaf)
    ses.EntityData.Leafs["counter"] = types.YLeaf{"Counter", ses.Counter}
    return &(ses.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas
// Unavailability Seconds information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Uas) GetEntityData() *types.CommonEntityData {
    uas.EntityData.YFilter = uas.YFilter
    uas.EntityData.YangName = "uas"
    uas.EntityData.BundleName = "cisco_ios_xr"
    uas.EntityData.ParentYangName = "otu-info"
    uas.EntityData.SegmentPath = "uas"
    uas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    uas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    uas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    uas.EntityData.Children = make(map[string]types.YChild)
    uas.EntityData.Leafs = make(map[string]types.YLeaf)
    uas.EntityData.Leafs["counter"] = types.YLeaf{"Counter", uas.Counter}
    return &(uas.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc
// Failure Count information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Fc) GetEntityData() *types.CommonEntityData {
    fc.EntityData.YFilter = fc.YFilter
    fc.EntityData.YangName = "fc"
    fc.EntityData.BundleName = "cisco_ios_xr"
    fc.EntityData.ParentYangName = "otu-info"
    fc.EntityData.SegmentPath = "fc"
    fc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fc.EntityData.Children = make(map[string]types.YChild)
    fc.EntityData.Leafs = make(map[string]types.YLeaf)
    fc.EntityData.Leafs["counter"] = types.YLeaf{"Counter", fc.Counter}
    return &(fc.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber
// Backgound Block Error Rate information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Bber) GetEntityData() *types.CommonEntityData {
    bber.EntityData.YFilter = bber.YFilter
    bber.EntityData.YangName = "bber"
    bber.EntityData.BundleName = "cisco_ios_xr"
    bber.EntityData.ParentYangName = "otu-info"
    bber.EntityData.SegmentPath = "bber"
    bber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bber.EntityData.Children = make(map[string]types.YChild)
    bber.EntityData.Leafs = make(map[string]types.YLeaf)
    bber.EntityData.Leafs["counter"] = types.YLeaf{"Counter", bber.Counter}
    return &(bber.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr
// Errored Seconds Rate information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Esr) GetEntityData() *types.CommonEntityData {
    esr.EntityData.YFilter = esr.YFilter
    esr.EntityData.YangName = "esr"
    esr.EntityData.BundleName = "cisco_ios_xr"
    esr.EntityData.ParentYangName = "otu-info"
    esr.EntityData.SegmentPath = "esr"
    esr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esr.EntityData.Children = make(map[string]types.YChild)
    esr.EntityData.Leafs = make(map[string]types.YLeaf)
    esr.EntityData.Leafs["counter"] = types.YLeaf{"Counter", esr.Counter}
    return &(esr.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr
// Severly Errored Seconds Rate information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Sesr) GetEntityData() *types.CommonEntityData {
    sesr.EntityData.YFilter = sesr.YFilter
    sesr.EntityData.YangName = "sesr"
    sesr.EntityData.BundleName = "cisco_ios_xr"
    sesr.EntityData.ParentYangName = "otu-info"
    sesr.EntityData.SegmentPath = "sesr"
    sesr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sesr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sesr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sesr.EntityData.Children = make(map[string]types.YChild)
    sesr.EntityData.Leafs = make(map[string]types.YLeaf)
    sesr.EntityData.Leafs["counter"] = types.YLeaf{"Counter", sesr.Counter}
    return &(sesr.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti
// Trail Trace Identifier information
type Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of String. The type is interface{} with range: 0..4294967295.
    TxStringType interface{}

    // Type of String. The type is interface{} with range: 0..4294967295.
    ExpectedStringType interface{}

    // Type of String. The type is interface{} with range: 0..4294967295.
    RxStringType interface{}

    // Tx TTI String . The type is string with length: 0..129.
    TxTti interface{}

    // Tx SAPI[0] Field. The type is string with length: 0..5.
    TxSapi0 interface{}

    // Tx SAPI[1-15] Field. The type is string with length: 0..16.
    TxSapi interface{}

    // Tx SAPI Range String. The type is string with length: 0..6.
    TxSapiRange interface{}

    // Tx DAPI[0] Field. The type is string with length: 0..5.
    TxDapi0 interface{}

    // Tx DAPI[1-15] Field. The type is string with length: 0..16.
    TxDapi interface{}

    // Tx DAPI Range String. The type is string with length: 0..6.
    TxDapiRange interface{}

    // Tx Operator Specific Field. The type is string with length: 0..33.
    TxOperSpec interface{}

    // Tx Operator Specific Field Range String. The type is string with length:
    // 0..6.
    TxOperSpecRange interface{}

    // Rx TTI String . The type is string with length: 0..129.
    RxTti interface{}

    // Rx SAPI[0] Field. The type is string with length: 0..5.
    RxSapi0 interface{}

    // Rx SAPI[1-15] Field. The type is string with length: 0..16.
    RxSapi interface{}

    // Rx SAPI Range String. The type is string with length: 0..6.
    RxSapiRange interface{}

    // Rx DAPI[0] Field. The type is string with length: 0..5.
    RxDapi0 interface{}

    // Rx DAPI[1-15] Field. The type is string with length: 0..16.
    RxDapi interface{}

    // Rx DAPI Range String. The type is string with length: 0..6.
    RxDapiRange interface{}

    // Rx Operator Specific Field Range String. The type is string with length:
    // 0..6.
    RxOperSpecRange interface{}

    // Rx Operator Specific Field. The type is string with length: 0..33.
    RxOperSpec interface{}

    // Expected TTI String. The type is string with length: 0..129.
    ExpectedTti interface{}

    // Expected SAPI[0] Field. The type is string with length: 0..5.
    ExpectedSapi0 interface{}

    // Expected SAPI[1-15] Field. The type is string with length: 0..16.
    ExpectedSapi interface{}

    // Expected SAPI Range String. The type is string with length: 0..6.
    ExpSapiRange interface{}

    // Expected DAPI[0] Field. The type is string with length: 0..5.
    ExpectedDapi0 interface{}

    // Expected DAPI[1-15] Field. The type is string with length: 0..16.
    ExpectedDapi interface{}

    // Expected DAPI Range String. The type is string with length: 0..6.
    ExpDapiRange interface{}

    // Expected Operator Specific Field. The type is string with length: 0..33.
    ExpectedOperSpec interface{}

    // Expected Operator Specific Field Range String. The type is string with
    // length: 0..6.
    ExpOperSpecRange interface{}
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OtuInfo_Tti) GetEntityData() *types.CommonEntityData {
    tti.EntityData.YFilter = tti.YFilter
    tti.EntityData.YangName = "tti"
    tti.EntityData.BundleName = "cisco_ios_xr"
    tti.EntityData.ParentYangName = "otu-info"
    tti.EntityData.SegmentPath = "tti"
    tti.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tti.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tti.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tti.EntityData.Children = make(map[string]types.YChild)
    tti.EntityData.Leafs = make(map[string]types.YLeaf)
    tti.EntityData.Leafs["tx-string-type"] = types.YLeaf{"TxStringType", tti.TxStringType}
    tti.EntityData.Leafs["expected-string-type"] = types.YLeaf{"ExpectedStringType", tti.ExpectedStringType}
    tti.EntityData.Leafs["rx-string-type"] = types.YLeaf{"RxStringType", tti.RxStringType}
    tti.EntityData.Leafs["tx-tti"] = types.YLeaf{"TxTti", tti.TxTti}
    tti.EntityData.Leafs["tx-sapi0"] = types.YLeaf{"TxSapi0", tti.TxSapi0}
    tti.EntityData.Leafs["tx-sapi"] = types.YLeaf{"TxSapi", tti.TxSapi}
    tti.EntityData.Leafs["tx-sapi-range"] = types.YLeaf{"TxSapiRange", tti.TxSapiRange}
    tti.EntityData.Leafs["tx-dapi0"] = types.YLeaf{"TxDapi0", tti.TxDapi0}
    tti.EntityData.Leafs["tx-dapi"] = types.YLeaf{"TxDapi", tti.TxDapi}
    tti.EntityData.Leafs["tx-dapi-range"] = types.YLeaf{"TxDapiRange", tti.TxDapiRange}
    tti.EntityData.Leafs["tx-oper-spec"] = types.YLeaf{"TxOperSpec", tti.TxOperSpec}
    tti.EntityData.Leafs["tx-oper-spec-range"] = types.YLeaf{"TxOperSpecRange", tti.TxOperSpecRange}
    tti.EntityData.Leafs["rx-tti"] = types.YLeaf{"RxTti", tti.RxTti}
    tti.EntityData.Leafs["rx-sapi0"] = types.YLeaf{"RxSapi0", tti.RxSapi0}
    tti.EntityData.Leafs["rx-sapi"] = types.YLeaf{"RxSapi", tti.RxSapi}
    tti.EntityData.Leafs["rx-sapi-range"] = types.YLeaf{"RxSapiRange", tti.RxSapiRange}
    tti.EntityData.Leafs["rx-dapi0"] = types.YLeaf{"RxDapi0", tti.RxDapi0}
    tti.EntityData.Leafs["rx-dapi"] = types.YLeaf{"RxDapi", tti.RxDapi}
    tti.EntityData.Leafs["rx-dapi-range"] = types.YLeaf{"RxDapiRange", tti.RxDapiRange}
    tti.EntityData.Leafs["rx-oper-spec-range"] = types.YLeaf{"RxOperSpecRange", tti.RxOperSpecRange}
    tti.EntityData.Leafs["rx-oper-spec"] = types.YLeaf{"RxOperSpec", tti.RxOperSpec}
    tti.EntityData.Leafs["expected-tti"] = types.YLeaf{"ExpectedTti", tti.ExpectedTti}
    tti.EntityData.Leafs["expected-sapi0"] = types.YLeaf{"ExpectedSapi0", tti.ExpectedSapi0}
    tti.EntityData.Leafs["expected-sapi"] = types.YLeaf{"ExpectedSapi", tti.ExpectedSapi}
    tti.EntityData.Leafs["exp-sapi-range"] = types.YLeaf{"ExpSapiRange", tti.ExpSapiRange}
    tti.EntityData.Leafs["expected-dapi0"] = types.YLeaf{"ExpectedDapi0", tti.ExpectedDapi0}
    tti.EntityData.Leafs["expected-dapi"] = types.YLeaf{"ExpectedDapi", tti.ExpectedDapi}
    tti.EntityData.Leafs["exp-dapi-range"] = types.YLeaf{"ExpDapiRange", tti.ExpDapiRange}
    tti.EntityData.Leafs["expected-oper-spec"] = types.YLeaf{"ExpectedOperSpec", tti.ExpectedOperSpec}
    tti.EntityData.Leafs["exp-oper-spec-range"] = types.YLeaf{"ExpOperSpecRange", tti.ExpOperSpecRange}
    return &(tti.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo
// ODU layer Information
type Dwdm_Ports_Port_Info_G709Info_OduInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bit Interleave Parity(BIP) counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Bip interface{}

    // Backward Error Indication counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Bei interface{}

    // Open Connection Indiction information.
    Oci Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci

    // Alarm Indication Signal information.
    Ais Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais

    // Upstream Connection Locked information.
    Lck Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck

    // Backward Defect Indication information.
    Bdi Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi

    // GCC End of Channel information.
    Eoc Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc

    // Payload Type Identifier Mismatch information.
    Ptim Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim

    // Trace Identifier Mismatch information.
    Tim Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim

    // Signal Fail  BER information.
    SfBer Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer

    // Signal Degrade BER information.
    SdBer Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer

    // Background Block Error TCA information.
    BbeTca Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca

    // Errored Seconds TCA information.
    EsTca Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca

    // Background Block Error information.
    Bbe Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe

    // Errored Seconds information.
    Es Dwdm_Ports_Port_Info_G709Info_OduInfo_Es

    // Severly Errored Seconds information.
    Ses Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses

    // Unavailability Seconds information.
    Uas Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas

    // Failure count information.
    Fc Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc

    // Background Block Error Rate count information.
    Bber Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber

    // Errored Seconds Rate information.
    Esr Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr

    // Severly Errored Seconds Rate information.
    Sesr Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr

    // Trail Trace Identifier information.
    Tti Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti
}

func (oduInfo *Dwdm_Ports_Port_Info_G709Info_OduInfo) GetEntityData() *types.CommonEntityData {
    oduInfo.EntityData.YFilter = oduInfo.YFilter
    oduInfo.EntityData.YangName = "odu-info"
    oduInfo.EntityData.BundleName = "cisco_ios_xr"
    oduInfo.EntityData.ParentYangName = "g709-info"
    oduInfo.EntityData.SegmentPath = "odu-info"
    oduInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oduInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oduInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oduInfo.EntityData.Children = make(map[string]types.YChild)
    oduInfo.EntityData.Children["oci"] = types.YChild{"Oci", &oduInfo.Oci}
    oduInfo.EntityData.Children["ais"] = types.YChild{"Ais", &oduInfo.Ais}
    oduInfo.EntityData.Children["lck"] = types.YChild{"Lck", &oduInfo.Lck}
    oduInfo.EntityData.Children["bdi"] = types.YChild{"Bdi", &oduInfo.Bdi}
    oduInfo.EntityData.Children["eoc"] = types.YChild{"Eoc", &oduInfo.Eoc}
    oduInfo.EntityData.Children["ptim"] = types.YChild{"Ptim", &oduInfo.Ptim}
    oduInfo.EntityData.Children["tim"] = types.YChild{"Tim", &oduInfo.Tim}
    oduInfo.EntityData.Children["sf-ber"] = types.YChild{"SfBer", &oduInfo.SfBer}
    oduInfo.EntityData.Children["sd-ber"] = types.YChild{"SdBer", &oduInfo.SdBer}
    oduInfo.EntityData.Children["bbe-tca"] = types.YChild{"BbeTca", &oduInfo.BbeTca}
    oduInfo.EntityData.Children["es-tca"] = types.YChild{"EsTca", &oduInfo.EsTca}
    oduInfo.EntityData.Children["bbe"] = types.YChild{"Bbe", &oduInfo.Bbe}
    oduInfo.EntityData.Children["es"] = types.YChild{"Es", &oduInfo.Es}
    oduInfo.EntityData.Children["ses"] = types.YChild{"Ses", &oduInfo.Ses}
    oduInfo.EntityData.Children["uas"] = types.YChild{"Uas", &oduInfo.Uas}
    oduInfo.EntityData.Children["fc"] = types.YChild{"Fc", &oduInfo.Fc}
    oduInfo.EntityData.Children["bber"] = types.YChild{"Bber", &oduInfo.Bber}
    oduInfo.EntityData.Children["esr"] = types.YChild{"Esr", &oduInfo.Esr}
    oduInfo.EntityData.Children["sesr"] = types.YChild{"Sesr", &oduInfo.Sesr}
    oduInfo.EntityData.Children["tti"] = types.YChild{"Tti", &oduInfo.Tti}
    oduInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    oduInfo.EntityData.Leafs["bip"] = types.YLeaf{"Bip", oduInfo.Bip}
    oduInfo.EntityData.Leafs["bei"] = types.YLeaf{"Bei", oduInfo.Bei}
    return &(oduInfo.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci
// Open Connection Indiction information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (oci *Dwdm_Ports_Port_Info_G709Info_OduInfo_Oci) GetEntityData() *types.CommonEntityData {
    oci.EntityData.YFilter = oci.YFilter
    oci.EntityData.YangName = "oci"
    oci.EntityData.BundleName = "cisco_ios_xr"
    oci.EntityData.ParentYangName = "odu-info"
    oci.EntityData.SegmentPath = "oci"
    oci.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oci.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oci.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oci.EntityData.Children = make(map[string]types.YChild)
    oci.EntityData.Leafs = make(map[string]types.YLeaf)
    oci.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", oci.ReportingEnabled}
    oci.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", oci.IsDetected}
    oci.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", oci.IsAsserted}
    oci.EntityData.Leafs["counter"] = types.YLeaf{"Counter", oci.Counter}
    return &(oci.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais
// Alarm Indication Signal information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ais *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ais) GetEntityData() *types.CommonEntityData {
    ais.EntityData.YFilter = ais.YFilter
    ais.EntityData.YangName = "ais"
    ais.EntityData.BundleName = "cisco_ios_xr"
    ais.EntityData.ParentYangName = "odu-info"
    ais.EntityData.SegmentPath = "ais"
    ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ais.EntityData.Children = make(map[string]types.YChild)
    ais.EntityData.Leafs = make(map[string]types.YLeaf)
    ais.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", ais.ReportingEnabled}
    ais.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", ais.IsDetected}
    ais.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", ais.IsAsserted}
    ais.EntityData.Leafs["counter"] = types.YLeaf{"Counter", ais.Counter}
    return &(ais.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck
// Upstream Connection Locked information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (lck *Dwdm_Ports_Port_Info_G709Info_OduInfo_Lck) GetEntityData() *types.CommonEntityData {
    lck.EntityData.YFilter = lck.YFilter
    lck.EntityData.YangName = "lck"
    lck.EntityData.BundleName = "cisco_ios_xr"
    lck.EntityData.ParentYangName = "odu-info"
    lck.EntityData.SegmentPath = "lck"
    lck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lck.EntityData.Children = make(map[string]types.YChild)
    lck.EntityData.Leafs = make(map[string]types.YLeaf)
    lck.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", lck.ReportingEnabled}
    lck.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", lck.IsDetected}
    lck.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", lck.IsAsserted}
    lck.EntityData.Leafs["counter"] = types.YLeaf{"Counter", lck.Counter}
    return &(lck.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi
// Backward Defect Indication information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (bdi *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bdi) GetEntityData() *types.CommonEntityData {
    bdi.EntityData.YFilter = bdi.YFilter
    bdi.EntityData.YangName = "bdi"
    bdi.EntityData.BundleName = "cisco_ios_xr"
    bdi.EntityData.ParentYangName = "odu-info"
    bdi.EntityData.SegmentPath = "bdi"
    bdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdi.EntityData.Children = make(map[string]types.YChild)
    bdi.EntityData.Leafs = make(map[string]types.YLeaf)
    bdi.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", bdi.ReportingEnabled}
    bdi.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", bdi.IsDetected}
    bdi.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", bdi.IsAsserted}
    bdi.EntityData.Leafs["counter"] = types.YLeaf{"Counter", bdi.Counter}
    return &(bdi.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc
// GCC End of Channel information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (eoc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Eoc) GetEntityData() *types.CommonEntityData {
    eoc.EntityData.YFilter = eoc.YFilter
    eoc.EntityData.YangName = "eoc"
    eoc.EntityData.BundleName = "cisco_ios_xr"
    eoc.EntityData.ParentYangName = "odu-info"
    eoc.EntityData.SegmentPath = "eoc"
    eoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eoc.EntityData.Children = make(map[string]types.YChild)
    eoc.EntityData.Leafs = make(map[string]types.YLeaf)
    eoc.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", eoc.ReportingEnabled}
    eoc.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", eoc.IsDetected}
    eoc.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", eoc.IsAsserted}
    eoc.EntityData.Leafs["counter"] = types.YLeaf{"Counter", eoc.Counter}
    return &(eoc.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim
// Payload Type Identifier Mismatch information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (ptim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ptim) GetEntityData() *types.CommonEntityData {
    ptim.EntityData.YFilter = ptim.YFilter
    ptim.EntityData.YangName = "ptim"
    ptim.EntityData.BundleName = "cisco_ios_xr"
    ptim.EntityData.ParentYangName = "odu-info"
    ptim.EntityData.SegmentPath = "ptim"
    ptim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ptim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ptim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ptim.EntityData.Children = make(map[string]types.YChild)
    ptim.EntityData.Leafs = make(map[string]types.YLeaf)
    ptim.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", ptim.ReportingEnabled}
    ptim.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", ptim.IsDetected}
    ptim.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", ptim.IsAsserted}
    ptim.EntityData.Leafs["counter"] = types.YLeaf{"Counter", ptim.Counter}
    return &(ptim.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim
// Trace Identifier Mismatch information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Alarm counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (tim *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tim) GetEntityData() *types.CommonEntityData {
    tim.EntityData.YFilter = tim.YFilter
    tim.EntityData.YangName = "tim"
    tim.EntityData.BundleName = "cisco_ios_xr"
    tim.EntityData.ParentYangName = "odu-info"
    tim.EntityData.SegmentPath = "tim"
    tim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tim.EntityData.Children = make(map[string]types.YChild)
    tim.EntityData.Leafs = make(map[string]types.YLeaf)
    tim.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", tim.ReportingEnabled}
    tim.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", tim.IsDetected}
    tim.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", tim.IsAsserted}
    tim.EntityData.Leafs["counter"] = types.YLeaf{"Counter", tim.Counter}
    return &(tim.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer
// Signal Fail  BER information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sfBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SfBer) GetEntityData() *types.CommonEntityData {
    sfBer.EntityData.YFilter = sfBer.YFilter
    sfBer.EntityData.YangName = "sf-ber"
    sfBer.EntityData.BundleName = "cisco_ios_xr"
    sfBer.EntityData.ParentYangName = "odu-info"
    sfBer.EntityData.SegmentPath = "sf-ber"
    sfBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sfBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sfBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sfBer.EntityData.Children = make(map[string]types.YChild)
    sfBer.EntityData.Leafs = make(map[string]types.YLeaf)
    sfBer.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", sfBer.ReportingEnabled}
    sfBer.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", sfBer.IsDetected}
    sfBer.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", sfBer.IsAsserted}
    sfBer.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", sfBer.Threshold}
    sfBer.EntityData.Leafs["counter"] = types.YLeaf{"Counter", sfBer.Counter}
    return &(sfBer.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer
// Signal Degrade BER information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (sdBer *Dwdm_Ports_Port_Info_G709Info_OduInfo_SdBer) GetEntityData() *types.CommonEntityData {
    sdBer.EntityData.YFilter = sdBer.YFilter
    sdBer.EntityData.YangName = "sd-ber"
    sdBer.EntityData.BundleName = "cisco_ios_xr"
    sdBer.EntityData.ParentYangName = "odu-info"
    sdBer.EntityData.SegmentPath = "sd-ber"
    sdBer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdBer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdBer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdBer.EntityData.Children = make(map[string]types.YChild)
    sdBer.EntityData.Leafs = make(map[string]types.YLeaf)
    sdBer.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", sdBer.ReportingEnabled}
    sdBer.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", sdBer.IsDetected}
    sdBer.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", sdBer.IsAsserted}
    sdBer.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", sdBer.Threshold}
    sdBer.EntityData.Leafs["counter"] = types.YLeaf{"Counter", sdBer.Counter}
    return &(sdBer.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca
// Background Block Error TCA information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (bbeTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_BbeTca) GetEntityData() *types.CommonEntityData {
    bbeTca.EntityData.YFilter = bbeTca.YFilter
    bbeTca.EntityData.YangName = "bbe-tca"
    bbeTca.EntityData.BundleName = "cisco_ios_xr"
    bbeTca.EntityData.ParentYangName = "odu-info"
    bbeTca.EntityData.SegmentPath = "bbe-tca"
    bbeTca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bbeTca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bbeTca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bbeTca.EntityData.Children = make(map[string]types.YChild)
    bbeTca.EntityData.Leafs = make(map[string]types.YLeaf)
    bbeTca.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", bbeTca.ReportingEnabled}
    bbeTca.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", bbeTca.IsDetected}
    bbeTca.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", bbeTca.IsAsserted}
    bbeTca.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", bbeTca.Threshold}
    bbeTca.EntityData.Leafs["counter"] = types.YLeaf{"Counter", bbeTca.Counter}
    return &(bbeTca.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca
// Errored Seconds TCA information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is reporting enabled?. The type is bool.
    ReportingEnabled interface{}

    // Is defect detected?. The type is bool.
    IsDetected interface{}

    // Is defect delared?. The type is bool.
    IsAsserted interface{}

    // Error threshold power. The type is interface{} with range:
    // -2147483648..2147483647.
    Threshold interface{}

    // Error counter. The type is interface{} with range: 0..18446744073709551615.
    Counter interface{}
}

func (esTca *Dwdm_Ports_Port_Info_G709Info_OduInfo_EsTca) GetEntityData() *types.CommonEntityData {
    esTca.EntityData.YFilter = esTca.YFilter
    esTca.EntityData.YangName = "es-tca"
    esTca.EntityData.BundleName = "cisco_ios_xr"
    esTca.EntityData.ParentYangName = "odu-info"
    esTca.EntityData.SegmentPath = "es-tca"
    esTca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esTca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esTca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esTca.EntityData.Children = make(map[string]types.YChild)
    esTca.EntityData.Leafs = make(map[string]types.YLeaf)
    esTca.EntityData.Leafs["reporting-enabled"] = types.YLeaf{"ReportingEnabled", esTca.ReportingEnabled}
    esTca.EntityData.Leafs["is-detected"] = types.YLeaf{"IsDetected", esTca.IsDetected}
    esTca.EntityData.Leafs["is-asserted"] = types.YLeaf{"IsAsserted", esTca.IsAsserted}
    esTca.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", esTca.Threshold}
    esTca.EntityData.Leafs["counter"] = types.YLeaf{"Counter", esTca.Counter}
    return &(esTca.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe
// Background Block Error information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (bbe *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bbe) GetEntityData() *types.CommonEntityData {
    bbe.EntityData.YFilter = bbe.YFilter
    bbe.EntityData.YangName = "bbe"
    bbe.EntityData.BundleName = "cisco_ios_xr"
    bbe.EntityData.ParentYangName = "odu-info"
    bbe.EntityData.SegmentPath = "bbe"
    bbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bbe.EntityData.Children = make(map[string]types.YChild)
    bbe.EntityData.Leafs = make(map[string]types.YLeaf)
    bbe.EntityData.Leafs["counter"] = types.YLeaf{"Counter", bbe.Counter}
    return &(bbe.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Es
// Errored Seconds information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Es struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (es *Dwdm_Ports_Port_Info_G709Info_OduInfo_Es) GetEntityData() *types.CommonEntityData {
    es.EntityData.YFilter = es.YFilter
    es.EntityData.YangName = "es"
    es.EntityData.BundleName = "cisco_ios_xr"
    es.EntityData.ParentYangName = "odu-info"
    es.EntityData.SegmentPath = "es"
    es.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    es.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    es.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    es.EntityData.Children = make(map[string]types.YChild)
    es.EntityData.Leafs = make(map[string]types.YLeaf)
    es.EntityData.Leafs["counter"] = types.YLeaf{"Counter", es.Counter}
    return &(es.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses
// Severly Errored Seconds information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (ses *Dwdm_Ports_Port_Info_G709Info_OduInfo_Ses) GetEntityData() *types.CommonEntityData {
    ses.EntityData.YFilter = ses.YFilter
    ses.EntityData.YangName = "ses"
    ses.EntityData.BundleName = "cisco_ios_xr"
    ses.EntityData.ParentYangName = "odu-info"
    ses.EntityData.SegmentPath = "ses"
    ses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ses.EntityData.Children = make(map[string]types.YChild)
    ses.EntityData.Leafs = make(map[string]types.YLeaf)
    ses.EntityData.Leafs["counter"] = types.YLeaf{"Counter", ses.Counter}
    return &(ses.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas
// Unavailability Seconds information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (uas *Dwdm_Ports_Port_Info_G709Info_OduInfo_Uas) GetEntityData() *types.CommonEntityData {
    uas.EntityData.YFilter = uas.YFilter
    uas.EntityData.YangName = "uas"
    uas.EntityData.BundleName = "cisco_ios_xr"
    uas.EntityData.ParentYangName = "odu-info"
    uas.EntityData.SegmentPath = "uas"
    uas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    uas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    uas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    uas.EntityData.Children = make(map[string]types.YChild)
    uas.EntityData.Leafs = make(map[string]types.YLeaf)
    uas.EntityData.Leafs["counter"] = types.YLeaf{"Counter", uas.Counter}
    return &(uas.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc
// Failure count information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (fc *Dwdm_Ports_Port_Info_G709Info_OduInfo_Fc) GetEntityData() *types.CommonEntityData {
    fc.EntityData.YFilter = fc.YFilter
    fc.EntityData.YangName = "fc"
    fc.EntityData.BundleName = "cisco_ios_xr"
    fc.EntityData.ParentYangName = "odu-info"
    fc.EntityData.SegmentPath = "fc"
    fc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fc.EntityData.Children = make(map[string]types.YChild)
    fc.EntityData.Leafs = make(map[string]types.YLeaf)
    fc.EntityData.Leafs["counter"] = types.YLeaf{"Counter", fc.Counter}
    return &(fc.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber
// Background Block Error Rate count information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (bber *Dwdm_Ports_Port_Info_G709Info_OduInfo_Bber) GetEntityData() *types.CommonEntityData {
    bber.EntityData.YFilter = bber.YFilter
    bber.EntityData.YangName = "bber"
    bber.EntityData.BundleName = "cisco_ios_xr"
    bber.EntityData.ParentYangName = "odu-info"
    bber.EntityData.SegmentPath = "bber"
    bber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bber.EntityData.Children = make(map[string]types.YChild)
    bber.EntityData.Leafs = make(map[string]types.YLeaf)
    bber.EntityData.Leafs["counter"] = types.YLeaf{"Counter", bber.Counter}
    return &(bber.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr
// Errored Seconds Rate information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (esr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Esr) GetEntityData() *types.CommonEntityData {
    esr.EntityData.YFilter = esr.YFilter
    esr.EntityData.YangName = "esr"
    esr.EntityData.BundleName = "cisco_ios_xr"
    esr.EntityData.ParentYangName = "odu-info"
    esr.EntityData.SegmentPath = "esr"
    esr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esr.EntityData.Children = make(map[string]types.YChild)
    esr.EntityData.Leafs = make(map[string]types.YLeaf)
    esr.EntityData.Leafs["counter"] = types.YLeaf{"Counter", esr.Counter}
    return &(esr.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr
// Severly Errored Seconds Rate information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Performance Monitoring counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Counter interface{}
}

func (sesr *Dwdm_Ports_Port_Info_G709Info_OduInfo_Sesr) GetEntityData() *types.CommonEntityData {
    sesr.EntityData.YFilter = sesr.YFilter
    sesr.EntityData.YangName = "sesr"
    sesr.EntityData.BundleName = "cisco_ios_xr"
    sesr.EntityData.ParentYangName = "odu-info"
    sesr.EntityData.SegmentPath = "sesr"
    sesr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sesr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sesr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sesr.EntityData.Children = make(map[string]types.YChild)
    sesr.EntityData.Leafs = make(map[string]types.YLeaf)
    sesr.EntityData.Leafs["counter"] = types.YLeaf{"Counter", sesr.Counter}
    return &(sesr.EntityData)
}

// Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti
// Trail Trace Identifier information
type Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of String. The type is interface{} with range: 0..4294967295.
    TxStringType interface{}

    // Type of String. The type is interface{} with range: 0..4294967295.
    ExpectedStringType interface{}

    // Type of String. The type is interface{} with range: 0..4294967295.
    RxStringType interface{}

    // Tx TTI String . The type is string with length: 0..129.
    TxTti interface{}

    // Tx SAPI[0] Field. The type is string with length: 0..5.
    TxSapi0 interface{}

    // Tx SAPI[1-15] Field. The type is string with length: 0..16.
    TxSapi interface{}

    // Tx SAPI Range String. The type is string with length: 0..6.
    TxSapiRange interface{}

    // Tx DAPI[0] Field. The type is string with length: 0..5.
    TxDapi0 interface{}

    // Tx DAPI[1-15] Field. The type is string with length: 0..16.
    TxDapi interface{}

    // Tx DAPI Range String. The type is string with length: 0..6.
    TxDapiRange interface{}

    // Tx Operator Specific Field. The type is string with length: 0..33.
    TxOperSpec interface{}

    // Tx Operator Specific Field Range String. The type is string with length:
    // 0..6.
    TxOperSpecRange interface{}

    // Rx TTI String . The type is string with length: 0..129.
    RxTti interface{}

    // Rx SAPI[0] Field. The type is string with length: 0..5.
    RxSapi0 interface{}

    // Rx SAPI[1-15] Field. The type is string with length: 0..16.
    RxSapi interface{}

    // Rx SAPI Range String. The type is string with length: 0..6.
    RxSapiRange interface{}

    // Rx DAPI[0] Field. The type is string with length: 0..5.
    RxDapi0 interface{}

    // Rx DAPI[1-15] Field. The type is string with length: 0..16.
    RxDapi interface{}

    // Rx DAPI Range String. The type is string with length: 0..6.
    RxDapiRange interface{}

    // Rx Operator Specific Field Range String. The type is string with length:
    // 0..6.
    RxOperSpecRange interface{}

    // Rx Operator Specific Field. The type is string with length: 0..33.
    RxOperSpec interface{}

    // Expected TTI String. The type is string with length: 0..129.
    ExpectedTti interface{}

    // Expected SAPI[0] Field. The type is string with length: 0..5.
    ExpectedSapi0 interface{}

    // Expected SAPI[1-15] Field. The type is string with length: 0..16.
    ExpectedSapi interface{}

    // Expected SAPI Range String. The type is string with length: 0..6.
    ExpSapiRange interface{}

    // Expected DAPI[0] Field. The type is string with length: 0..5.
    ExpectedDapi0 interface{}

    // Expected DAPI[1-15] Field. The type is string with length: 0..16.
    ExpectedDapi interface{}

    // Expected DAPI Range String. The type is string with length: 0..6.
    ExpDapiRange interface{}

    // Expected Operator Specific Field. The type is string with length: 0..33.
    ExpectedOperSpec interface{}

    // Expected Operator Specific Field Range String. The type is string with
    // length: 0..6.
    ExpOperSpecRange interface{}
}

func (tti *Dwdm_Ports_Port_Info_G709Info_OduInfo_Tti) GetEntityData() *types.CommonEntityData {
    tti.EntityData.YFilter = tti.YFilter
    tti.EntityData.YangName = "tti"
    tti.EntityData.BundleName = "cisco_ios_xr"
    tti.EntityData.ParentYangName = "odu-info"
    tti.EntityData.SegmentPath = "tti"
    tti.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tti.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tti.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tti.EntityData.Children = make(map[string]types.YChild)
    tti.EntityData.Leafs = make(map[string]types.YLeaf)
    tti.EntityData.Leafs["tx-string-type"] = types.YLeaf{"TxStringType", tti.TxStringType}
    tti.EntityData.Leafs["expected-string-type"] = types.YLeaf{"ExpectedStringType", tti.ExpectedStringType}
    tti.EntityData.Leafs["rx-string-type"] = types.YLeaf{"RxStringType", tti.RxStringType}
    tti.EntityData.Leafs["tx-tti"] = types.YLeaf{"TxTti", tti.TxTti}
    tti.EntityData.Leafs["tx-sapi0"] = types.YLeaf{"TxSapi0", tti.TxSapi0}
    tti.EntityData.Leafs["tx-sapi"] = types.YLeaf{"TxSapi", tti.TxSapi}
    tti.EntityData.Leafs["tx-sapi-range"] = types.YLeaf{"TxSapiRange", tti.TxSapiRange}
    tti.EntityData.Leafs["tx-dapi0"] = types.YLeaf{"TxDapi0", tti.TxDapi0}
    tti.EntityData.Leafs["tx-dapi"] = types.YLeaf{"TxDapi", tti.TxDapi}
    tti.EntityData.Leafs["tx-dapi-range"] = types.YLeaf{"TxDapiRange", tti.TxDapiRange}
    tti.EntityData.Leafs["tx-oper-spec"] = types.YLeaf{"TxOperSpec", tti.TxOperSpec}
    tti.EntityData.Leafs["tx-oper-spec-range"] = types.YLeaf{"TxOperSpecRange", tti.TxOperSpecRange}
    tti.EntityData.Leafs["rx-tti"] = types.YLeaf{"RxTti", tti.RxTti}
    tti.EntityData.Leafs["rx-sapi0"] = types.YLeaf{"RxSapi0", tti.RxSapi0}
    tti.EntityData.Leafs["rx-sapi"] = types.YLeaf{"RxSapi", tti.RxSapi}
    tti.EntityData.Leafs["rx-sapi-range"] = types.YLeaf{"RxSapiRange", tti.RxSapiRange}
    tti.EntityData.Leafs["rx-dapi0"] = types.YLeaf{"RxDapi0", tti.RxDapi0}
    tti.EntityData.Leafs["rx-dapi"] = types.YLeaf{"RxDapi", tti.RxDapi}
    tti.EntityData.Leafs["rx-dapi-range"] = types.YLeaf{"RxDapiRange", tti.RxDapiRange}
    tti.EntityData.Leafs["rx-oper-spec-range"] = types.YLeaf{"RxOperSpecRange", tti.RxOperSpecRange}
    tti.EntityData.Leafs["rx-oper-spec"] = types.YLeaf{"RxOperSpec", tti.RxOperSpec}
    tti.EntityData.Leafs["expected-tti"] = types.YLeaf{"ExpectedTti", tti.ExpectedTti}
    tti.EntityData.Leafs["expected-sapi0"] = types.YLeaf{"ExpectedSapi0", tti.ExpectedSapi0}
    tti.EntityData.Leafs["expected-sapi"] = types.YLeaf{"ExpectedSapi", tti.ExpectedSapi}
    tti.EntityData.Leafs["exp-sapi-range"] = types.YLeaf{"ExpSapiRange", tti.ExpSapiRange}
    tti.EntityData.Leafs["expected-dapi0"] = types.YLeaf{"ExpectedDapi0", tti.ExpectedDapi0}
    tti.EntityData.Leafs["expected-dapi"] = types.YLeaf{"ExpectedDapi", tti.ExpectedDapi}
    tti.EntityData.Leafs["exp-dapi-range"] = types.YLeaf{"ExpDapiRange", tti.ExpDapiRange}
    tti.EntityData.Leafs["expected-oper-spec"] = types.YLeaf{"ExpectedOperSpec", tti.ExpectedOperSpec}
    tti.EntityData.Leafs["exp-oper-spec-range"] = types.YLeaf{"ExpOperSpecRange", tti.ExpOperSpecRange}
    return &(tti.EntityData)
}

// Dwdm_Ports_Port_Info_OpticsInfo
// Optics operational information
type Dwdm_Ports_Port_Info_OpticsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Optics type name. The type is string with length: 0..64.
    OpticsType interface{}

    // Actual transmit clock source. The type is interface{} with range: 0..255.
    ClockSource interface{}

    // Wave Frequency Information for Progressive Frequencies. The type is string
    // with length: 0..64.
    WaveFrequencyProgressiveString interface{}

    // Wavelength Information for Progressive Frequencies. The type is string with
    // length: 0..64.
    WavelengthProgressiveString interface{}

    // True if Progressive Frequency is supported by hw. The type is bool.
    IsWaveFrequencyProgressiveValid interface{}

    // Wavelength Information for Progressive Frequencies. The type is interface{}
    // with range: 0..4294967295.
    WavelengthProgressive interface{}

    // Wavelength band information. The type is interface{} with range:
    // 0..4294967295.
    WaveBand interface{}

    // Current ITU wavelength channel number. The type is interface{} with range:
    // 0..4294967295.
    WaveChannel interface{}

    // wavelenght frequency read from hw in the uint 0 .01nm. The type is
    // interface{} with range: 0..4294967295.
    WaveFrequency interface{}

    // True if hw supported wavelength frequency readback. The type is bool.
    IsWaveFrequencyValid interface{}

    // Owner of current wavelength. The type is DwdmWaveChannelOwner.
    WaveChannelOwner interface{}

    // Wavelength channel set by GMPLS. The type is interface{} with range:
    // 0..65535.
    GmplsSetWaveChannel interface{}

    // Wavelength channel set from configuration. The type is interface{} with
    // range: 0..65535.
    ConfiguredWaveChannel interface{}

    // Wavelength channel default from hardware. The type is interface{} with
    // range: 0..65535.
    DefaultWaveChannel interface{}

    // Transmit power in the unit of 0.01dbm. The type is interface{} with range:
    // -2147483648..2147483647.
    TransmitPower interface{}

    // Transmit power threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    TransmitPowerThreshold interface{}

    // Laser current bias value. The type is interface{} with range:
    // -2147483648..2147483647.
    LaserCurrentBias interface{}

    // Laser Current Bias threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    LaserCurrentBiasThreshold interface{}

    // Transponder receive power. The type is interface{} with range:
    // -2147483648..2147483647.
    ReceivePower interface{}

    // TRUE if  Rx LOS thresold configurable. The type is bool.
    IsRxLosThresholdSupported interface{}

    // Rx LOS threshold value. The type is interface{} with range:
    // -2147483648..2147483647.
    RxLosThreshold interface{}

    // Transmit  power mininum value in the interval time. The type is interface{}
    // with range: -2147483648..2147483647.
    TransmitPowerMin interface{}

    // Transmit power maximum value in the interval time. The type is interface{}
    // with range: -2147483648..2147483647.
    TransmitPowerMax interface{}

    // Transmit optical average value in the interval time. The type is
    // interface{} with range: -2147483648..2147483647.
    TransmitPowerAvg interface{}

    // Recieve power mininum value in the interval time. The type is interface{}
    // with range: -2147483648..2147483647.
    ReceivePowerMin interface{}

    // Receive power maximum value in the interval time. The type is interface{}
    // with range: -2147483648..2147483647.
    ReceivePowerMax interface{}

    // Recieve power average value in the interval time. The type is interface{}
    // with range: -2147483648..2147483647.
    ReceivePowerAvg interface{}

    // Laser bias current minimum value in the interval time. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserBiasCurrentMin interface{}

    // Laser bias current maxinum value in the interval time. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserBiasCurrentMax interface{}

    // Laser bias current average value in the interval time. The type is
    // interface{} with range: -2147483648..2147483647.
    LaserBiasCurrentAvg interface{}

    // Current chromatic dispersion. The type is interface{} with range:
    // -2147483648..2147483647.
    ChromaticDispersion interface{}

    // Current differential group Delay. The type is interface{} with range:
    // -2147483648..2147483647.
    DifferentialGroupDelay interface{}

    // Current polarization mode dispersion. The type is interface{} with range:
    // -2147483648..2147483647.
    PolarizationModeDispersion interface{}

    // Current optical signal to noise ratio. The type is interface{} with range:
    // -2147483648..2147483647.
    SignalToNoiseRatio interface{}

    // Current Polarization Dependent loss. The type is interface{} with range:
    // -2147483648..2147483647.
    PolarizationDependentLoss interface{}

    // Current Polarization change rate. The type is interface{} with range:
    // 0..4294967295.
    PolarizationChangeRate interface{}

    // Current Phase Noise. The type is interface{} with range: 0..4294967295.
    PhaseNoise interface{}

    // Transmit power failure(above/belowe a threshold) count. The type is
    // interface{} with range: 0..4294967295.
    OutputPowerFail interface{}

    // Receive power failure(above/belowe a threshold) count. The type is
    // interface{} with range: 0..4294967295.
    InputPowerFail interface{}
}

func (opticsInfo *Dwdm_Ports_Port_Info_OpticsInfo) GetEntityData() *types.CommonEntityData {
    opticsInfo.EntityData.YFilter = opticsInfo.YFilter
    opticsInfo.EntityData.YangName = "optics-info"
    opticsInfo.EntityData.BundleName = "cisco_ios_xr"
    opticsInfo.EntityData.ParentYangName = "info"
    opticsInfo.EntityData.SegmentPath = "optics-info"
    opticsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsInfo.EntityData.Children = make(map[string]types.YChild)
    opticsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    opticsInfo.EntityData.Leafs["optics-type"] = types.YLeaf{"OpticsType", opticsInfo.OpticsType}
    opticsInfo.EntityData.Leafs["clock-source"] = types.YLeaf{"ClockSource", opticsInfo.ClockSource}
    opticsInfo.EntityData.Leafs["wave-frequency-progressive-string"] = types.YLeaf{"WaveFrequencyProgressiveString", opticsInfo.WaveFrequencyProgressiveString}
    opticsInfo.EntityData.Leafs["wavelength-progressive-string"] = types.YLeaf{"WavelengthProgressiveString", opticsInfo.WavelengthProgressiveString}
    opticsInfo.EntityData.Leafs["is-wave-frequency-progressive-valid"] = types.YLeaf{"IsWaveFrequencyProgressiveValid", opticsInfo.IsWaveFrequencyProgressiveValid}
    opticsInfo.EntityData.Leafs["wavelength-progressive"] = types.YLeaf{"WavelengthProgressive", opticsInfo.WavelengthProgressive}
    opticsInfo.EntityData.Leafs["wave-band"] = types.YLeaf{"WaveBand", opticsInfo.WaveBand}
    opticsInfo.EntityData.Leafs["wave-channel"] = types.YLeaf{"WaveChannel", opticsInfo.WaveChannel}
    opticsInfo.EntityData.Leafs["wave-frequency"] = types.YLeaf{"WaveFrequency", opticsInfo.WaveFrequency}
    opticsInfo.EntityData.Leafs["is-wave-frequency-valid"] = types.YLeaf{"IsWaveFrequencyValid", opticsInfo.IsWaveFrequencyValid}
    opticsInfo.EntityData.Leafs["wave-channel-owner"] = types.YLeaf{"WaveChannelOwner", opticsInfo.WaveChannelOwner}
    opticsInfo.EntityData.Leafs["gmpls-set-wave-channel"] = types.YLeaf{"GmplsSetWaveChannel", opticsInfo.GmplsSetWaveChannel}
    opticsInfo.EntityData.Leafs["configured-wave-channel"] = types.YLeaf{"ConfiguredWaveChannel", opticsInfo.ConfiguredWaveChannel}
    opticsInfo.EntityData.Leafs["default-wave-channel"] = types.YLeaf{"DefaultWaveChannel", opticsInfo.DefaultWaveChannel}
    opticsInfo.EntityData.Leafs["transmit-power"] = types.YLeaf{"TransmitPower", opticsInfo.TransmitPower}
    opticsInfo.EntityData.Leafs["transmit-power-threshold"] = types.YLeaf{"TransmitPowerThreshold", opticsInfo.TransmitPowerThreshold}
    opticsInfo.EntityData.Leafs["laser-current-bias"] = types.YLeaf{"LaserCurrentBias", opticsInfo.LaserCurrentBias}
    opticsInfo.EntityData.Leafs["laser-current-bias-threshold"] = types.YLeaf{"LaserCurrentBiasThreshold", opticsInfo.LaserCurrentBiasThreshold}
    opticsInfo.EntityData.Leafs["receive-power"] = types.YLeaf{"ReceivePower", opticsInfo.ReceivePower}
    opticsInfo.EntityData.Leafs["is-rx-los-threshold-supported"] = types.YLeaf{"IsRxLosThresholdSupported", opticsInfo.IsRxLosThresholdSupported}
    opticsInfo.EntityData.Leafs["rx-los-threshold"] = types.YLeaf{"RxLosThreshold", opticsInfo.RxLosThreshold}
    opticsInfo.EntityData.Leafs["transmit-power-min"] = types.YLeaf{"TransmitPowerMin", opticsInfo.TransmitPowerMin}
    opticsInfo.EntityData.Leafs["transmit-power-max"] = types.YLeaf{"TransmitPowerMax", opticsInfo.TransmitPowerMax}
    opticsInfo.EntityData.Leafs["transmit-power-avg"] = types.YLeaf{"TransmitPowerAvg", opticsInfo.TransmitPowerAvg}
    opticsInfo.EntityData.Leafs["receive-power-min"] = types.YLeaf{"ReceivePowerMin", opticsInfo.ReceivePowerMin}
    opticsInfo.EntityData.Leafs["receive-power-max"] = types.YLeaf{"ReceivePowerMax", opticsInfo.ReceivePowerMax}
    opticsInfo.EntityData.Leafs["receive-power-avg"] = types.YLeaf{"ReceivePowerAvg", opticsInfo.ReceivePowerAvg}
    opticsInfo.EntityData.Leafs["laser-bias-current-min"] = types.YLeaf{"LaserBiasCurrentMin", opticsInfo.LaserBiasCurrentMin}
    opticsInfo.EntityData.Leafs["laser-bias-current-max"] = types.YLeaf{"LaserBiasCurrentMax", opticsInfo.LaserBiasCurrentMax}
    opticsInfo.EntityData.Leafs["laser-bias-current-avg"] = types.YLeaf{"LaserBiasCurrentAvg", opticsInfo.LaserBiasCurrentAvg}
    opticsInfo.EntityData.Leafs["chromatic-dispersion"] = types.YLeaf{"ChromaticDispersion", opticsInfo.ChromaticDispersion}
    opticsInfo.EntityData.Leafs["differential-group-delay"] = types.YLeaf{"DifferentialGroupDelay", opticsInfo.DifferentialGroupDelay}
    opticsInfo.EntityData.Leafs["polarization-mode-dispersion"] = types.YLeaf{"PolarizationModeDispersion", opticsInfo.PolarizationModeDispersion}
    opticsInfo.EntityData.Leafs["signal-to-noise-ratio"] = types.YLeaf{"SignalToNoiseRatio", opticsInfo.SignalToNoiseRatio}
    opticsInfo.EntityData.Leafs["polarization-dependent-loss"] = types.YLeaf{"PolarizationDependentLoss", opticsInfo.PolarizationDependentLoss}
    opticsInfo.EntityData.Leafs["polarization-change-rate"] = types.YLeaf{"PolarizationChangeRate", opticsInfo.PolarizationChangeRate}
    opticsInfo.EntityData.Leafs["phase-noise"] = types.YLeaf{"PhaseNoise", opticsInfo.PhaseNoise}
    opticsInfo.EntityData.Leafs["output-power-fail"] = types.YLeaf{"OutputPowerFail", opticsInfo.OutputPowerFail}
    opticsInfo.EntityData.Leafs["input-power-fail"] = types.YLeaf{"InputPowerFail", opticsInfo.InputPowerFail}
    return &(opticsInfo.EntityData)
}

// Dwdm_Ports_Port_Info_TdcInfo
// TDC operational information
type Dwdm_Ports_Port_Info_TdcInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE for Valid else Invalid. The type is bool.
    TdcValid interface{}

    // TRUE for Alarm condition else FALSE. The type is bool.
    MajorAlarm interface{}

    // TRUE for MANUAL else AUTO. The type is bool.
    OperationMode interface{}

    // TRUE if TDC Aquiring else Locked. The type is bool.
    TdcStatus interface{}

    // TDC Dispersion Offset. The type is interface{} with range:
    // -2147483648..2147483647.
    DispersionOffset interface{}

    // Reroute BER. The type is interface{} with range: -2147483648..2147483647.
    RerouteBer interface{}

    // TRUE for ENABLED else DISABLED. The type is bool.
    IsRerouteControlEnabled interface{}
}

func (tdcInfo *Dwdm_Ports_Port_Info_TdcInfo) GetEntityData() *types.CommonEntityData {
    tdcInfo.EntityData.YFilter = tdcInfo.YFilter
    tdcInfo.EntityData.YangName = "tdc-info"
    tdcInfo.EntityData.BundleName = "cisco_ios_xr"
    tdcInfo.EntityData.ParentYangName = "info"
    tdcInfo.EntityData.SegmentPath = "tdc-info"
    tdcInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tdcInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tdcInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tdcInfo.EntityData.Children = make(map[string]types.YChild)
    tdcInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    tdcInfo.EntityData.Leafs["tdc-valid"] = types.YLeaf{"TdcValid", tdcInfo.TdcValid}
    tdcInfo.EntityData.Leafs["major-alarm"] = types.YLeaf{"MajorAlarm", tdcInfo.MajorAlarm}
    tdcInfo.EntityData.Leafs["operation-mode"] = types.YLeaf{"OperationMode", tdcInfo.OperationMode}
    tdcInfo.EntityData.Leafs["tdc-status"] = types.YLeaf{"TdcStatus", tdcInfo.TdcStatus}
    tdcInfo.EntityData.Leafs["dispersion-offset"] = types.YLeaf{"DispersionOffset", tdcInfo.DispersionOffset}
    tdcInfo.EntityData.Leafs["reroute-ber"] = types.YLeaf{"RerouteBer", tdcInfo.RerouteBer}
    tdcInfo.EntityData.Leafs["is-reroute-control-enabled"] = types.YLeaf{"IsRerouteControlEnabled", tdcInfo.IsRerouteControlEnabled}
    return &(tdcInfo.EntityData)
}

// Dwdm_Ports_Port_Info_NetworkSrlgInfo
// Network SRLG information
type Dwdm_Ports_Port_Info_NetworkSrlgInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network Srlg. The type is slice of interface{} with range: 0..4294967295.
    NetworkSrlg []interface{}
}

func (networkSrlgInfo *Dwdm_Ports_Port_Info_NetworkSrlgInfo) GetEntityData() *types.CommonEntityData {
    networkSrlgInfo.EntityData.YFilter = networkSrlgInfo.YFilter
    networkSrlgInfo.EntityData.YangName = "network-srlg-info"
    networkSrlgInfo.EntityData.BundleName = "cisco_ios_xr"
    networkSrlgInfo.EntityData.ParentYangName = "info"
    networkSrlgInfo.EntityData.SegmentPath = "network-srlg-info"
    networkSrlgInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkSrlgInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkSrlgInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkSrlgInfo.EntityData.Children = make(map[string]types.YChild)
    networkSrlgInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    networkSrlgInfo.EntityData.Leafs["network-srlg"] = types.YLeaf{"NetworkSrlg", networkSrlgInfo.NetworkSrlg}
    return &(networkSrlgInfo.EntityData)
}

// Dwdm_Ports_Port_Info_Proactive
// Proactive protection information
type Dwdm_Ports_Port_Info_Proactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Feature Support. The type is bool.
    ProactiveFeature interface{}

    // Proactive Mode. The type is G709ppfsmMode.
    ProactiveMode interface{}

    // Proactive FSM State. The type is G709ppfsmState.
    ProactiveFsmState interface{}

    // Proactive FSM IF State. The type is G709ppintfState.
    ProactiveFsmIfState interface{}

    // TAS State. The type is DwdmtasState.
    TasState interface{}

    // Trigger threshold coefficient. The type is interface{} with range: 0..255.
    TrigThreshCoeff interface{}

    // Trigger threshold power. The type is interface{} with range: 0..255.
    TrigThreshPower interface{}

    // Revert threshold coefficient. The type is interface{} with range: 0..255.
    RvrtThreshCoeff interface{}

    // Revert threshold power. The type is interface{} with range: 0..255.
    RvrtThreshPower interface{}

    // Default Trigger threshold coefficient. The type is interface{} with range:
    // 0..255.
    DefaultTrigThreshCoeff interface{}

    // Default Trigger threshold power. The type is interface{} with range:
    // 0..255.
    DefaultTrigThreshPower interface{}

    // Default Revert threshold coefficient. The type is interface{} with range:
    // 0..255.
    DefaultRvrtThreshCoeff interface{}

    // Default Revert threshold power. The type is interface{} with range: 0..255.
    DefaultRvrtThreshPower interface{}

    // Required Trigger Samples. The type is interface{} with range: 0..255.
    TrigSamples interface{}

    // Required Revert Samples. The type is interface{} with range: 0..255.
    RvrtSamples interface{}

    // Trigger Integration window. The type is interface{} with range:
    // 0..4294967295.
    TriggerWindow interface{}

    // Revert Integration Window. The type is interface{} with range:
    // 0..4294967295.
    RevertWindow interface{}

    // Protection Trigger State. The type is bool.
    ProtectionTrigger interface{}

    // Proactive Interface Triffer. The type is bool.
    InterfaceTrigger interface{}

    // Transmitted APS Byte. The type is interface{} with range: 0..255.
    TxAps interface{}

    // Tx APS Description. The type is G709apsByte.
    TxApsDescr interface{}

    // Received APS byte. The type is interface{} with range: 0..255.
    RxAps interface{}

    // Rx APS Description. The type is G709apsByte.
    RxApsDescr interface{}

    // AlarmState. The type is bool.
    AlarmState interface{}

    // Trigger EC Cnt. The type is interface{} with range: 0..4294967295.
    TrigEcCnt interface{}

    // Revert EC Cnt. The type is interface{} with range: 0..4294967295.
    RvrtEcCnt interface{}

    // Prefec Trigger Thresh Crossed. The type is bool.
    PrefecThreshCrossed interface{}
}

func (proactive *Dwdm_Ports_Port_Info_Proactive) GetEntityData() *types.CommonEntityData {
    proactive.EntityData.YFilter = proactive.YFilter
    proactive.EntityData.YangName = "proactive"
    proactive.EntityData.BundleName = "cisco_ios_xr"
    proactive.EntityData.ParentYangName = "info"
    proactive.EntityData.SegmentPath = "proactive"
    proactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proactive.EntityData.Children = make(map[string]types.YChild)
    proactive.EntityData.Leafs = make(map[string]types.YLeaf)
    proactive.EntityData.Leafs["proactive-feature"] = types.YLeaf{"ProactiveFeature", proactive.ProactiveFeature}
    proactive.EntityData.Leafs["proactive-mode"] = types.YLeaf{"ProactiveMode", proactive.ProactiveMode}
    proactive.EntityData.Leafs["proactive-fsm-state"] = types.YLeaf{"ProactiveFsmState", proactive.ProactiveFsmState}
    proactive.EntityData.Leafs["proactive-fsm-if-state"] = types.YLeaf{"ProactiveFsmIfState", proactive.ProactiveFsmIfState}
    proactive.EntityData.Leafs["tas-state"] = types.YLeaf{"TasState", proactive.TasState}
    proactive.EntityData.Leafs["trig-thresh-coeff"] = types.YLeaf{"TrigThreshCoeff", proactive.TrigThreshCoeff}
    proactive.EntityData.Leafs["trig-thresh-power"] = types.YLeaf{"TrigThreshPower", proactive.TrigThreshPower}
    proactive.EntityData.Leafs["rvrt-thresh-coeff"] = types.YLeaf{"RvrtThreshCoeff", proactive.RvrtThreshCoeff}
    proactive.EntityData.Leafs["rvrt-thresh-power"] = types.YLeaf{"RvrtThreshPower", proactive.RvrtThreshPower}
    proactive.EntityData.Leafs["default-trig-thresh-coeff"] = types.YLeaf{"DefaultTrigThreshCoeff", proactive.DefaultTrigThreshCoeff}
    proactive.EntityData.Leafs["default-trig-thresh-power"] = types.YLeaf{"DefaultTrigThreshPower", proactive.DefaultTrigThreshPower}
    proactive.EntityData.Leafs["default-rvrt-thresh-coeff"] = types.YLeaf{"DefaultRvrtThreshCoeff", proactive.DefaultRvrtThreshCoeff}
    proactive.EntityData.Leafs["default-rvrt-thresh-power"] = types.YLeaf{"DefaultRvrtThreshPower", proactive.DefaultRvrtThreshPower}
    proactive.EntityData.Leafs["trig-samples"] = types.YLeaf{"TrigSamples", proactive.TrigSamples}
    proactive.EntityData.Leafs["rvrt-samples"] = types.YLeaf{"RvrtSamples", proactive.RvrtSamples}
    proactive.EntityData.Leafs["trigger-window"] = types.YLeaf{"TriggerWindow", proactive.TriggerWindow}
    proactive.EntityData.Leafs["revert-window"] = types.YLeaf{"RevertWindow", proactive.RevertWindow}
    proactive.EntityData.Leafs["protection-trigger"] = types.YLeaf{"ProtectionTrigger", proactive.ProtectionTrigger}
    proactive.EntityData.Leafs["interface-trigger"] = types.YLeaf{"InterfaceTrigger", proactive.InterfaceTrigger}
    proactive.EntityData.Leafs["tx-aps"] = types.YLeaf{"TxAps", proactive.TxAps}
    proactive.EntityData.Leafs["tx-aps-descr"] = types.YLeaf{"TxApsDescr", proactive.TxApsDescr}
    proactive.EntityData.Leafs["rx-aps"] = types.YLeaf{"RxAps", proactive.RxAps}
    proactive.EntityData.Leafs["rx-aps-descr"] = types.YLeaf{"RxApsDescr", proactive.RxApsDescr}
    proactive.EntityData.Leafs["alarm-state"] = types.YLeaf{"AlarmState", proactive.AlarmState}
    proactive.EntityData.Leafs["trig-ec-cnt"] = types.YLeaf{"TrigEcCnt", proactive.TrigEcCnt}
    proactive.EntityData.Leafs["rvrt-ec-cnt"] = types.YLeaf{"RvrtEcCnt", proactive.RvrtEcCnt}
    proactive.EntityData.Leafs["prefec-thresh-crossed"] = types.YLeaf{"PrefecThreshCrossed", proactive.PrefecThreshCrossed}
    return &(proactive.EntityData)
}

// Dwdm_Ports_Port_Info_SignalLog
// Signal log information
type Dwdm_Ports_Port_Info_SignalLog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 'true' if signal log is enabled 'false' otherwise. The type is bool.
    IsLogEnabled interface{}

    // Log file name . The type is string with length: 0..64.
    LogFilename interface{}
}

func (signalLog *Dwdm_Ports_Port_Info_SignalLog) GetEntityData() *types.CommonEntityData {
    signalLog.EntityData.YFilter = signalLog.YFilter
    signalLog.EntityData.YangName = "signal-log"
    signalLog.EntityData.BundleName = "cisco_ios_xr"
    signalLog.EntityData.ParentYangName = "info"
    signalLog.EntityData.SegmentPath = "signal-log"
    signalLog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signalLog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signalLog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signalLog.EntityData.Children = make(map[string]types.YChild)
    signalLog.EntityData.Leafs = make(map[string]types.YLeaf)
    signalLog.EntityData.Leafs["is-log-enabled"] = types.YLeaf{"IsLogEnabled", signalLog.IsLogEnabled}
    signalLog.EntityData.Leafs["log-filename"] = types.YLeaf{"LogFilename", signalLog.LogFilename}
    return &(signalLog.EntityData)
}

// Vtxp
// vtxp
type Vtxp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DWDM operational data.
    DwdmVtxp Vtxp_DwdmVtxp
}

func (vtxp *Vtxp) GetEntityData() *types.CommonEntityData {
    vtxp.EntityData.YFilter = vtxp.YFilter
    vtxp.EntityData.YangName = "vtxp"
    vtxp.EntityData.BundleName = "cisco_ios_xr"
    vtxp.EntityData.ParentYangName = "Cisco-IOS-XR-dwdm-ui-oper"
    vtxp.EntityData.SegmentPath = "Cisco-IOS-XR-dwdm-ui-oper:vtxp"
    vtxp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vtxp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vtxp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vtxp.EntityData.Children = make(map[string]types.YChild)
    vtxp.EntityData.Children["dwdm-vtxp"] = types.YChild{"DwdmVtxp", &vtxp.DwdmVtxp}
    vtxp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vtxp.EntityData)
}

// Vtxp_DwdmVtxp
// DWDM operational data
type Vtxp_DwdmVtxp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All DWDM Port operational data.
    PortVtxps Vtxp_DwdmVtxp_PortVtxps
}

func (dwdmVtxp *Vtxp_DwdmVtxp) GetEntityData() *types.CommonEntityData {
    dwdmVtxp.EntityData.YFilter = dwdmVtxp.YFilter
    dwdmVtxp.EntityData.YangName = "dwdm-vtxp"
    dwdmVtxp.EntityData.BundleName = "cisco_ios_xr"
    dwdmVtxp.EntityData.ParentYangName = "vtxp"
    dwdmVtxp.EntityData.SegmentPath = "dwdm-vtxp"
    dwdmVtxp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dwdmVtxp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dwdmVtxp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dwdmVtxp.EntityData.Children = make(map[string]types.YChild)
    dwdmVtxp.EntityData.Children["port-vtxps"] = types.YChild{"PortVtxps", &dwdmVtxp.PortVtxps}
    dwdmVtxp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dwdmVtxp.EntityData)
}

// Vtxp_DwdmVtxp_PortVtxps
// All DWDM Port operational data
type Vtxp_DwdmVtxp_PortVtxps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DWDM Port operational data. The type is slice of
    // Vtxp_DwdmVtxp_PortVtxps_PortVtxp.
    PortVtxp []Vtxp_DwdmVtxp_PortVtxps_PortVtxp
}

func (portVtxps *Vtxp_DwdmVtxp_PortVtxps) GetEntityData() *types.CommonEntityData {
    portVtxps.EntityData.YFilter = portVtxps.YFilter
    portVtxps.EntityData.YangName = "port-vtxps"
    portVtxps.EntityData.BundleName = "cisco_ios_xr"
    portVtxps.EntityData.ParentYangName = "dwdm-vtxp"
    portVtxps.EntityData.SegmentPath = "port-vtxps"
    portVtxps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portVtxps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portVtxps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portVtxps.EntityData.Children = make(map[string]types.YChild)
    portVtxps.EntityData.Children["port-vtxp"] = types.YChild{"PortVtxp", nil}
    for i := range portVtxps.PortVtxp {
        portVtxps.EntityData.Children[types.GetSegmentPath(&portVtxps.PortVtxp[i])] = types.YChild{"PortVtxp", &portVtxps.PortVtxp[i]}
    }
    portVtxps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(portVtxps.EntityData)
}

// Vtxp_DwdmVtxp_PortVtxps_PortVtxp
// DWDM Port operational data
type Vtxp_DwdmVtxp_PortVtxps_PortVtxp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // DWDM port operational data.
    Info Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info
}

func (portVtxp *Vtxp_DwdmVtxp_PortVtxps_PortVtxp) GetEntityData() *types.CommonEntityData {
    portVtxp.EntityData.YFilter = portVtxp.YFilter
    portVtxp.EntityData.YangName = "port-vtxp"
    portVtxp.EntityData.BundleName = "cisco_ios_xr"
    portVtxp.EntityData.ParentYangName = "port-vtxps"
    portVtxp.EntityData.SegmentPath = "port-vtxp" + "[name='" + fmt.Sprintf("%v", portVtxp.Name) + "']"
    portVtxp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portVtxp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portVtxp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portVtxp.EntityData.Children = make(map[string]types.YChild)
    portVtxp.EntityData.Children["info"] = types.YChild{"Info", &portVtxp.Info}
    portVtxp.EntityData.Leafs = make(map[string]types.YLeaf)
    portVtxp.EntityData.Leafs["name"] = types.YLeaf{"Name", portVtxp.Name}
    return &(portVtxp.EntityData)
}

// Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info
// DWDM port operational data
type Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is VTXP attribute enabled. The type is bool.
    VtxpEnable interface{}
}

func (info *Vtxp_DwdmVtxp_PortVtxps_PortVtxp_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "port-vtxp"
    info.EntityData.SegmentPath = "info"
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = make(map[string]types.YChild)
    info.EntityData.Leafs = make(map[string]types.YLeaf)
    info.EntityData.Leafs["vtxp-enable"] = types.YLeaf{"VtxpEnable", info.VtxpEnable}
    return &(info.EntityData)
}

