// This module contains a collection of YANG definitions
// for Cisco IOS-XR ethernet-cfm package operational data.
// 
// This module contains definitions
// for the following management objects:
//   cfm: CFM operational data
// 
// This YANG module augments the
//   Cisco-IOS-XR-infra-sla-oper
// module with state data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_cfm_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_cfm_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ethernet-cfm-oper cfm}", reflect.TypeOf(Cfm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ethernet-cfm-oper:cfm", reflect.TypeOf(Cfm{}))
}

// CfmAisDir represents Cfm ais dir
type CfmAisDir string

const (
    // Packets sent inward
    CfmAisDir_up CfmAisDir = "up"

    // Packets sent outward
    CfmAisDir_down CfmAisDir = "down"
)

// CfmBagSmanFmt represents Short MA Name format
type CfmBagSmanFmt string

const (
    // Short MA Name is a 12-bit VLAN-ID
    CfmBagSmanFmt_sman_vlan_id CfmBagSmanFmt = "sman-vlan-id"

    // Short MA Name is a character string
    CfmBagSmanFmt_sman_string CfmBagSmanFmt = "sman-string"

    // Short MA Name is a 16-bit unsigned integer
    CfmBagSmanFmt_sman_uint16 CfmBagSmanFmt = "sman-uint16"

    // Short MA Name is a global VPN identifier
    CfmBagSmanFmt_sman_vpn_id CfmBagSmanFmt = "sman-vpn-id"

    // Short MA Name uses the ICC-based format
    CfmBagSmanFmt_sman_icc CfmBagSmanFmt = "sman-icc"

    // Unknown Short MA Name format
    CfmBagSmanFmt_sman_unknown CfmBagSmanFmt = "sman-unknown"
)

// CfmBagMdidFmt represents CFM MDID format
type CfmBagMdidFmt string

const (
    // MDID is explicity NULL
    CfmBagMdidFmt_mdid_null CfmBagMdidFmt = "mdid-null"

    // MDID is based on a DNS name
    CfmBagMdidFmt_mdid_dns_like CfmBagMdidFmt = "mdid-dns-like"

    // MDID is a (MAC address, integer) pair
    CfmBagMdidFmt_mdid_mac_address CfmBagMdidFmt = "mdid-mac-address"

    // MDID is a character string
    CfmBagMdidFmt_mdid_string CfmBagMdidFmt = "mdid-string"

    // Unknown MDID format
    CfmBagMdidFmt_mdid_unknown CfmBagMdidFmt = "mdid-unknown"
)

// CfmBagCcmOffload represents Offload status of CCM processing
type CfmBagCcmOffload string

const (
    // CCM processing has not been offloaded
    CfmBagCcmOffload_offload_none CfmBagCcmOffload = "offload-none"

    // CCM processing has been offloaded to software
    CfmBagCcmOffload_offload_software CfmBagCcmOffload = "offload-software"

    // CCM processing has been offloaded to hardware
    CfmBagCcmOffload_offload_hardware CfmBagCcmOffload = "offload-hardware"
)

// CfmBagCcmInterval represents CFM CCM intervals
type CfmBagCcmInterval string

const (
    // Invalid CCM interval
    CfmBagCcmInterval_interval_none CfmBagCcmInterval = "interval-none"

    // Interval of 3.3ms
    CfmBagCcmInterval_interval3_3ms CfmBagCcmInterval = "interval3-3ms"

    // Interval of 10ms
    CfmBagCcmInterval_interval10ms CfmBagCcmInterval = "interval10ms"

    // Interval of 100ms
    CfmBagCcmInterval_interval100ms CfmBagCcmInterval = "interval100ms"

    // Interval of 1s
    CfmBagCcmInterval_interval1s CfmBagCcmInterval = "interval1s"

    // Interval of 10s
    CfmBagCcmInterval_interval10s CfmBagCcmInterval = "interval10s"

    // Interval of 1 min
    CfmBagCcmInterval_interval1m CfmBagCcmInterval = "interval1m"

    // Interval of 10 mins
    CfmBagCcmInterval_interval10m CfmBagCcmInterval = "interval10m"
)

// CfmBagBdidFmt represents Bridge domain identifier format
type CfmBagBdidFmt string

const (
    // Invalid BDID identifier format
    CfmBagBdidFmt_invalid CfmBagBdidFmt = "invalid"

    // Identifier is a bridge domain ID
    CfmBagBdidFmt_bd_id CfmBagBdidFmt = "bd-id"

    // Identifier is a P2P cross-connect ID
    CfmBagBdidFmt_xc_p2p_id CfmBagBdidFmt = "xc-p2p-id"

    // Identifier is a MP2MP cross-connect ID
    CfmBagBdidFmt_xc_mp2mp_id CfmBagBdidFmt = "xc-mp2mp-id"

    // Identifier is a VLAN-aware flexible
    // cross-connect ID
    CfmBagBdidFmt_fxc_vlan_aware_id CfmBagBdidFmt = "fxc-vlan-aware-id"

    // Identifier is a VLAN-unaware flexible
    // cross-connect ID
    CfmBagBdidFmt_fxc_vlan_unaware_id CfmBagBdidFmt = "fxc-vlan-unaware-id"

    // Identifier is a maintenance association name
    CfmBagBdidFmt_down_only CfmBagBdidFmt = "down-only"
)

// CfmMaMpVariety represents CFM MA Maintenance Point varieties
type CfmMaMpVariety string

const (
    // MIP
    CfmMaMpVariety_mip CfmMaMpVariety = "mip"

    // Up MEP
    CfmMaMpVariety_up_mep CfmMaMpVariety = "up-mep"

    // Down MEP
    CfmMaMpVariety_downmep CfmMaMpVariety = "downmep"

    // Unknown MEP
    CfmMaMpVariety_unknown_mep CfmMaMpVariety = "unknown-mep"
)

// CfmBagIssuRole represents CFM ISSU role
type CfmBagIssuRole string

const (
    // Unknown
    CfmBagIssuRole_unknown CfmBagIssuRole = "unknown"

    // Primary
    CfmBagIssuRole_primary CfmBagIssuRole = "primary"

    // Secondary
    CfmBagIssuRole_secondary CfmBagIssuRole = "secondary"
)

// CfmBagOpcode represents CFM Opcode
type CfmBagOpcode string

const (
    // Reserved
    CfmBagOpcode_reserved CfmBagOpcode = "reserved"

    // Continuity Check
    CfmBagOpcode_ccm CfmBagOpcode = "ccm"

    // Loopback Reply
    CfmBagOpcode_lbr CfmBagOpcode = "lbr"

    // Loopback Message
    CfmBagOpcode_lbm CfmBagOpcode = "lbm"

    // Linktrace Reply
    CfmBagOpcode_ltr CfmBagOpcode = "ltr"

    // Linktrace Message
    CfmBagOpcode_ltm CfmBagOpcode = "ltm"
)

// CfmBagAisInterval represents CFM AIS intervals
type CfmBagAisInterval string

const (
    // Invalid AIS interval
    CfmBagAisInterval_ais_interval_none CfmBagAisInterval = "ais-interval-none"

    // Interval of 1s
    CfmBagAisInterval_ais_interval1s CfmBagAisInterval = "ais-interval1s"

    // Interval of 1 min
    CfmBagAisInterval_ais_interval1m CfmBagAisInterval = "ais-interval1m"
)

// CfmBagMdLevel represents CFM level
type CfmBagMdLevel string

const (
    // CFM level 0
    CfmBagMdLevel_level0 CfmBagMdLevel = "level0"

    // CFM level 1
    CfmBagMdLevel_level1 CfmBagMdLevel = "level1"

    // CFM level 2
    CfmBagMdLevel_level2 CfmBagMdLevel = "level2"

    // CFM level 3
    CfmBagMdLevel_level3 CfmBagMdLevel = "level3"

    // CFM level 4
    CfmBagMdLevel_level4 CfmBagMdLevel = "level4"

    // CFM level 5
    CfmBagMdLevel_level5 CfmBagMdLevel = "level5"

    // CFM level 6
    CfmBagMdLevel_level6 CfmBagMdLevel = "level6"

    // CFM level 7
    CfmBagMdLevel_level7 CfmBagMdLevel = "level7"

    // Invalid CFM level
    CfmBagMdLevel_level_invalid CfmBagMdLevel = "level-invalid"
)

// CfmBagDirection represents MEP direction
type CfmBagDirection string

const (
    // Up
    CfmBagDirection_direction_up CfmBagDirection = "direction-up"

    // Down
    CfmBagDirection_direction_down CfmBagDirection = "direction-down"

    // Invalid direction
    CfmBagDirection_direction_invalid CfmBagDirection = "direction-invalid"
)

// CfmBagStpState represents CFM STP state
type CfmBagStpState string

const (
    // Interface is UP
    CfmBagStpState_stp_up CfmBagStpState = "stp-up"

    // Interface is STP-blocked
    CfmBagStpState_stp_blocked CfmBagStpState = "stp-blocked"

    // Unknown Interface STP state
    CfmBagStpState_stp_unknown CfmBagStpState = "stp-unknown"
)

// CfmBagIwState represents CFM Interworking state
type CfmBagIwState string

const (
    // Interface is UP
    CfmBagIwState_interworking_up CfmBagIwState = "interworking-up"

    // Interface is in TEST mode
    CfmBagIwState_interworking_test CfmBagIwState = "interworking-test"
)

// CfmPmAddlIntfStatus represents Additional interface status
type CfmPmAddlIntfStatus string

const (
    // Additional interface status unknown
    CfmPmAddlIntfStatus_unknown CfmPmAddlIntfStatus = "unknown"

    // Interface is explicitly shutdown in
    // configuration
    CfmPmAddlIntfStatus_administratively_down CfmPmAddlIntfStatus = "administratively-down"

    // Remote interface has exceeded its 802.3 Link
    // OAM error threshold
    CfmPmAddlIntfStatus_remote_excessive_errors CfmPmAddlIntfStatus = "remote-excessive-errors"

    // Local interface has exceeded its 802.3 Link OAM
    // error threshold
    CfmPmAddlIntfStatus_local_excessive_errors CfmPmAddlIntfStatus = "local-excessive-errors"
)

// CfmPmIntfStatus represents Interface status
type CfmPmIntfStatus string

const (
    // Interface is up
    CfmPmIntfStatus_interface_status_up CfmPmIntfStatus = "interface-status-up"

    // Interface is down
    CfmPmIntfStatus_interface_status_down CfmPmIntfStatus = "interface-status-down"

    // Interface is in testing mode
    CfmPmIntfStatus_interface_status_testing CfmPmIntfStatus = "interface-status-testing"

    // Unknown interface status
    CfmPmIntfStatus_interface_status_unknown CfmPmIntfStatus = "interface-status-unknown"

    // Interface is dormant
    CfmPmIntfStatus_interface_status_dormant CfmPmIntfStatus = "interface-status-dormant"

    // Interface status not found
    CfmPmIntfStatus_interface_status_not_present CfmPmIntfStatus = "interface-status-not-present"

    // Lower layer is down
    CfmPmIntfStatus_interface_status_lower_layer_down CfmPmIntfStatus = "interface-status-lower-layer-down"
)

// CfmPmPortStatus represents Port status
type CfmPmPortStatus string

const (
    // Port is STP blocked
    CfmPmPortStatus_port_status_blocked CfmPmPortStatus = "port-status-blocked"

    // Port is up
    CfmPmPortStatus_port_status_up CfmPmPortStatus = "port-status-up"

    // Unknown port status
    CfmPmPortStatus_port_status_unknown CfmPmPortStatus = "port-status-unknown"
)

// CfmPmRmepState represents State of the Peer MEP state machine
type CfmPmRmepState string

const (
    // Momentary state during reset
    CfmPmRmepState_peer_mep_idle CfmPmRmepState = "peer-mep-idle"

    // Loss timer not expired since reset, but no
    // valid CCM received
    CfmPmRmepState_peer_mep_start CfmPmRmepState = "peer-mep-start"

    // Loss timer has expired
    CfmPmRmepState_peer_mep_failed CfmPmRmepState = "peer-mep-failed"

    // Loss timer has not expired since last valid CCM
    CfmPmRmepState_peer_mep_ok CfmPmRmepState = "peer-mep-ok"
)

// CfmPmRmepXcState represents Cross-check state of a peer MEP
type CfmPmRmepXcState string

const (
    // Cross-check OK
    CfmPmRmepXcState_cross_check_ok CfmPmRmepXcState = "cross-check-ok"

    // No CCMs received within loss time from peer MEP
    CfmPmRmepXcState_cross_check_missing CfmPmRmepXcState = "cross-check-missing"

    // CCMs received from peer MEP not marked for
    // cross-check
    CfmPmRmepXcState_cross_check_extra CfmPmRmepXcState = "cross-check-extra"
)

// CfmPmAisReceive represents signal, directly or via AIS or LCK messages.
type CfmPmAisReceive string

const (
    // No signal received
    CfmPmAisReceive_receive_none CfmPmAisReceive = "receive-none"

    // Receiving AIS messages
    CfmPmAisReceive_receive_ais CfmPmAisReceive = "receive-ais"

    // Receiving LCK messages
    CfmPmAisReceive_receive_lck CfmPmAisReceive = "receive-lck"

    // Receiving AIS directly from another MEP on the
    // same interface
    CfmPmAisReceive_receive_direct CfmPmAisReceive = "receive-direct"
)

// CfmPmAisTransmit represents via a MIP or directly to a higher MEP
type CfmPmAisTransmit string

const (
    // AIS not transmitted
    CfmPmAisTransmit_transmit_none CfmPmAisTransmit = "transmit-none"

    // AIS transmitted via MIP
    CfmPmAisTransmit_transmit_ais CfmPmAisTransmit = "transmit-ais"

    // AIS signal passed directly to a higher MEP
    CfmPmAisTransmit_transmit_ais_direct CfmPmAisTransmit = "transmit-ais-direct"
)

// CfmPmMepDefect represents Defects that can be reported by a MEP
type CfmPmMepDefect string

const (
    // No defect reported
    CfmPmMepDefect_defect_none CfmPmMepDefect = "defect-none"

    // Some Peer MEP's CCM has the RDI bit set
    CfmPmMepDefect_defect_rdi_ccm CfmPmMepDefect = "defect-rdi-ccm"

    // A Peer MEP port or interface status error has
    // been reported
    CfmPmMepDefect_defect_ma_cstatus CfmPmMepDefect = "defect-ma-cstatus"

    // Not receiving valid CCMs from at least one Peer
    // MEP
    CfmPmMepDefect_defect_remote_ccm CfmPmMepDefect = "defect-remote-ccm"

    // Currently receiving invalid CCMs from at least
    // one Peer MEP
    CfmPmMepDefect_defect_error_ccm CfmPmMepDefect = "defect-error-ccm"

    // Currently receiving CCMs from an incorrect
    // service (MA)
    CfmPmMepDefect_defect_cross_connect_ccm CfmPmMepDefect = "defect-cross-connect-ccm"
)

// CfmPmMepFngState represents states
type CfmPmMepFngState string

const (
    // FNG in reset state
    CfmPmMepFngState_fng_reset CfmPmMepFngState = "fng-reset"

    // FNG has detected but not yet reported a defect
    CfmPmMepFngState_fng_defect CfmPmMepFngState = "fng-defect"

    // FNG is in the process of reporting a defect
    CfmPmMepFngState_fng_report_defect CfmPmMepFngState = "fng-report-defect"

    // FNG has reported a defect
    CfmPmMepFngState_fng_defect_reported CfmPmMepFngState = "fng-defect-reported"

    // No defect present, but the reset timer has not
    // yet expired
    CfmPmMepFngState_fng_defect_clearing CfmPmMepFngState = "fng-defect-clearing"
)

// CfmPmElrEgressAction represents ELR Egress action
type CfmPmElrEgressAction string

const (
    // OK
    CfmPmElrEgressAction_elr_egress_ok CfmPmElrEgressAction = "elr-egress-ok"

    // Down
    CfmPmElrEgressAction_elr_egress_down CfmPmElrEgressAction = "elr-egress-down"

    // STP Blocked
    CfmPmElrEgressAction_elr_egress_blocked CfmPmElrEgressAction = "elr-egress-blocked"

    // VID Blocked
    CfmPmElrEgressAction_elr_egress_vid CfmPmElrEgressAction = "elr-egress-vid"

    // MAC Pruned
    CfmPmElrEgressAction_elr_egress_mac CfmPmElrEgressAction = "elr-egress-mac"
)

// CfmPmElrIngressAction represents ELR Ingress action
type CfmPmElrIngressAction string

const (
    // OK
    CfmPmElrIngressAction_elr_ingress_ok CfmPmElrIngressAction = "elr-ingress-ok"

    // Down
    CfmPmElrIngressAction_elr_ingress_down CfmPmElrIngressAction = "elr-ingress-down"

    // STP Blocked
    CfmPmElrIngressAction_elr_ingress_blocked CfmPmElrIngressAction = "elr-ingress-blocked"

    // VID Blocked
    CfmPmElrIngressAction_elr_ingress_vid CfmPmElrIngressAction = "elr-ingress-vid"
)

// CfmPmElrRelayAction represents ELR relay action
type CfmPmElrRelayAction string

const (
    // Target Hit
    CfmPmElrRelayAction_elr_relay_hit CfmPmElrRelayAction = "elr-relay-hit"

    // Filtering database
    CfmPmElrRelayAction_elr_relay_fdb CfmPmElrRelayAction = "elr-relay-fdb"

    // Flood forwarded
    CfmPmElrRelayAction_elr_relay_flood CfmPmElrRelayAction = "elr-relay-flood"

    // Dropped
    CfmPmElrRelayAction_elr_relay_drop CfmPmElrRelayAction = "elr-relay-drop"
)

// CfmPmLastHopFmt represents Last hop identifier format
type CfmPmLastHopFmt string

const (
    // No last hop identifier
    CfmPmLastHopFmt_last_hop_none CfmPmLastHopFmt = "last-hop-none"

    // Last hop identifier is a hostname
    CfmPmLastHopFmt_last_hop_host_name CfmPmLastHopFmt = "last-hop-host-name"

    // Last hop identifier is an egress ID
    CfmPmLastHopFmt_last_hop_egress_id CfmPmLastHopFmt = "last-hop-egress-id"
)

// CfmPmEgressAction represents Egress action
type CfmPmEgressAction string

const (
    // OK
    CfmPmEgressAction_egress_ok CfmPmEgressAction = "egress-ok"

    // Down
    CfmPmEgressAction_egress_down CfmPmEgressAction = "egress-down"

    // STP Blocked
    CfmPmEgressAction_egress_blocked CfmPmEgressAction = "egress-blocked"

    // VID Blocked
    CfmPmEgressAction_egress_vid CfmPmEgressAction = "egress-vid"
)

// CfmPmPortIdFmt represents Port ID format
type CfmPmPortIdFmt string

const (
    // Port ID is an interface alias
    CfmPmPortIdFmt_port_id_interface_alias CfmPmPortIdFmt = "port-id-interface-alias"

    // Port ID is a component name
    CfmPmPortIdFmt_port_id_port_component CfmPmPortIdFmt = "port-id-port-component"

    // Port ID is a MAC address
    CfmPmPortIdFmt_port_id_mac_address CfmPmPortIdFmt = "port-id-mac-address"

    // Port ID is a network address
    CfmPmPortIdFmt_port_id_network_address CfmPmPortIdFmt = "port-id-network-address"

    // Port ID is an interface name
    CfmPmPortIdFmt_port_id_interface_name CfmPmPortIdFmt = "port-id-interface-name"

    // Port ID is an agent name
    CfmPmPortIdFmt_port_id_agent_circuit_id CfmPmPortIdFmt = "port-id-agent-circuit-id"

    // Port ID is a local name
    CfmPmPortIdFmt_port_id_local CfmPmPortIdFmt = "port-id-local"

    // Port ID format unknown
    CfmPmPortIdFmt_port_id_unknown CfmPmPortIdFmt = "port-id-unknown"
)

// CfmPmIngressAction represents Ingress action
type CfmPmIngressAction string

const (
    // OK
    CfmPmIngressAction_ingress_ok CfmPmIngressAction = "ingress-ok"

    // Down
    CfmPmIngressAction_ingress_down CfmPmIngressAction = "ingress-down"

    // STP Blocked
    CfmPmIngressAction_ingress_blocked CfmPmIngressAction = "ingress-blocked"

    // VID Blocked
    CfmPmIngressAction_ingress_vid CfmPmIngressAction = "ingress-vid"
)

// CfmPmIdFmt represents ID format
type CfmPmIdFmt string

const (
    // ID format is a string
    CfmPmIdFmt_id_format_is_string CfmPmIdFmt = "id-format-is-string"

    // ID format is a MAC address
    CfmPmIdFmt_id_format_is_mac_address CfmPmIdFmt = "id-format-is-mac-address"

    // ID format is raw hex
    CfmPmIdFmt_id_format_is_raw_hex CfmPmIdFmt = "id-format-is-raw-hex"
)

// CfmPmChassisIdFmt represents Chassis ID type
type CfmPmChassisIdFmt string

const (
    // Chassis ID is a component name
    CfmPmChassisIdFmt_chassis_id_chassis_component CfmPmChassisIdFmt = "chassis-id-chassis-component"

    // Chassis ID is an interface alias
    CfmPmChassisIdFmt_chassis_id_interface_alias CfmPmChassisIdFmt = "chassis-id-interface-alias"

    // Chassis ID is a port component name
    CfmPmChassisIdFmt_chassis_id_port_component CfmPmChassisIdFmt = "chassis-id-port-component"

    // Chassis ID is a MAC address
    CfmPmChassisIdFmt_chassis_id_mac_address CfmPmChassisIdFmt = "chassis-id-mac-address"

    // Chassis ID is a network address
    CfmPmChassisIdFmt_chassis_id_network_address CfmPmChassisIdFmt = "chassis-id-network-address"

    // Chassis ID is an interface name
    CfmPmChassisIdFmt_chassis_id_interface_name CfmPmChassisIdFmt = "chassis-id-interface-name"

    // Chassis ID is a local name
    CfmPmChassisIdFmt_chassis_id_local CfmPmChassisIdFmt = "chassis-id-local"

    // Unknown Chassis ID type
    CfmPmChassisIdFmt_chassis_id_unknown_type CfmPmChassisIdFmt = "chassis-id-unknown-type"
)

// CfmPmRelayAction represents LTR relay action
type CfmPmRelayAction string

const (
    // Target Hit
    CfmPmRelayAction_relay_hit CfmPmRelayAction = "relay-hit"

    // Filtering database
    CfmPmRelayAction_relay_fdb CfmPmRelayAction = "relay-fdb"

    // CCM Learning database
    CfmPmRelayAction_relay_mpdb CfmPmRelayAction = "relay-mpdb"
)

// CfmPmElmReplyFilter represents operations
type CfmPmElmReplyFilter string

const (
    // Reply Filter not present
    CfmPmElmReplyFilter_reply_filter_not_present CfmPmElmReplyFilter = "reply-filter-not-present"

    // Reply from ports which are not MAC-pruned,
    // VID-pruned, or STP-blocked
    CfmPmElmReplyFilter_reply_filter_default CfmPmElmReplyFilter = "reply-filter-default"

    // Reply from ports which are not VID-pruned or
    // STP-blocked
    CfmPmElmReplyFilter_reply_filter_vlan_topology CfmPmElmReplyFilter = "reply-filter-vlan-topology"

    // Reply from ports which are not STP-blocked
    CfmPmElmReplyFilter_reply_filter_spanning_tree CfmPmElmReplyFilter = "reply-filter-spanning-tree"

    // Reply from all ports
    CfmPmElmReplyFilter_reply_filter_all_ports CfmPmElmReplyFilter = "reply-filter-all-ports"
)

// CfmPmEltDelayModel represents operations
type CfmPmEltDelayModel string

const (
    // Not a valid delay model
    CfmPmEltDelayModel_delay_model_invalid CfmPmEltDelayModel = "delay-model-invalid"

    // Reply using logarithmic delay model
    CfmPmEltDelayModel_delay_model_logarithmic CfmPmEltDelayModel = "delay-model-logarithmic"

    // Reply using constant delay model
    CfmPmEltDelayModel_delay_model_constant CfmPmEltDelayModel = "delay-model-constant"
)

// CfmPmLtMode represents Type of Linktrace operation
type CfmPmLtMode string

const (
    // Basic IEEE 802.1ag Linktrace
    CfmPmLtMode_cfm_pm_lt_mode_basic CfmPmLtMode = "cfm-pm-lt-mode-basic"

    // Cisco Exploratory Linktrace
    CfmPmLtMode_cfm_pm_lt_mode_exploratory CfmPmLtMode = "cfm-pm-lt-mode-exploratory"
)

// CfmPmPktAction represents Action taken for received packet
type CfmPmPktAction string

const (
    // Packet processed successfully
    CfmPmPktAction_packet_processed CfmPmPktAction = "packet-processed"

    // Packet forwarded
    CfmPmPktAction_packet_forwarded CfmPmPktAction = "packet-forwarded"

    // Packet dropped at a MEP due to unknown opcode
    CfmPmPktAction_unknown_opcode CfmPmPktAction = "unknown-opcode"

    // Packet dropped due to level/opcode filtering at
    // a MEP
    CfmPmPktAction_filter_level CfmPmPktAction = "filter-level"

    // Packet dropped because interface is STP blocked
    CfmPmPktAction_filter_blocked CfmPmPktAction = "filter-blocked"

    // Packet dropped due to local destination MAC
    CfmPmPktAction_filter_local_mac CfmPmPktAction = "filter-local-mac"

    // CCM too short or too long
    CfmPmPktAction_malformed_ccm_size CfmPmPktAction = "malformed-ccm-size"

    // Invalid MEP-ID
    CfmPmPktAction_malformed_ccm_mep_id CfmPmPktAction = "malformed-ccm-mep-id"

    // Packet too short
    CfmPmPktAction_malformed_too_short CfmPmPktAction = "malformed-too-short"

    // Destination MAC address does not match
    // interface
    CfmPmPktAction_malformed_destination_mac_unicast CfmPmPktAction = "malformed-destination-mac-unicast"

    // Invalid multicast destination MAC address
    CfmPmPktAction_malformed_destination_mac_multicast CfmPmPktAction = "malformed-destination-mac-multicast"

    // TLV offset too short or beyond the end of the
    // packet
    CfmPmPktAction_malformed_tlv_offset CfmPmPktAction = "malformed-tlv-offset"

    // Invalid source MAC address for LBM
    CfmPmPktAction_malformed_lbm_source_mac CfmPmPktAction = "malformed-lbm-source-mac"

    // Unknown LTR relay action
    CfmPmPktAction_malformed_ltr_relay_action CfmPmPktAction = "malformed-ltr-relay-action"

    // LTR has neither reply-ingress or reply-egress
    CfmPmPktAction_malformed_ltr_reply_tlv CfmPmPktAction = "malformed-ltr-reply-tlv"

    // Invalid Linktrace Message origin MAC address
    CfmPmPktAction_malformed_lt_origin CfmPmPktAction = "malformed-lt-origin"

    // Invalid LTM target MAC address
    CfmPmPktAction_malformed_ltm_target CfmPmPktAction = "malformed-ltm-target"

    // Invalid source MAC address
    CfmPmPktAction_malformed_source_mac CfmPmPktAction = "malformed-source-mac"

    // Packet too short for CFM header
    CfmPmPktAction_malformed_header_too_short CfmPmPktAction = "malformed-header-too-short"

    // TLV header extends beyond the end of the packet
    CfmPmPktAction_malformed_tlv_header_overrun CfmPmPktAction = "malformed-tlv-header-overrun"

    // TLV extends beyond the end of the packet
    CfmPmPktAction_malformed_tlv_overrun CfmPmPktAction = "malformed-tlv-overrun"

    // Multiple Sender-ID TLVs found
    CfmPmPktAction_malformed_duplicate_sender_id CfmPmPktAction = "malformed-duplicate-sender-id"

    // Multiple Port-status TLVs found
    CfmPmPktAction_malformed_duplicate_port_status CfmPmPktAction = "malformed-duplicate-port-status"

    // Multiple Interface-state TLVs found
    CfmPmPktAction_malformed_duplicate_interface_status CfmPmPktAction = "malformed-duplicate-interface-status"

    // Invalid TLV for this type of packet found
    CfmPmPktAction_malformed_wrong_tlv CfmPmPktAction = "malformed-wrong-tlv"

    // Multiple Data TLVs found
    CfmPmPktAction_malformed_duplicate_data CfmPmPktAction = "malformed-duplicate-data"

    // Multiple LTR-Egress-ID TLVs found
    CfmPmPktAction_malformed_duplicate_ltr_egress_id CfmPmPktAction = "malformed-duplicate-ltr-egress-id"

    // Multiple Reply-ingress TLVs found
    CfmPmPktAction_malformed_duplicate_reply_ingress CfmPmPktAction = "malformed-duplicate-reply-ingress"

    // Multiple Reply-egress TLVs found
    CfmPmPktAction_malformed_duplicate_reply_egress CfmPmPktAction = "malformed-duplicate-reply-egress"

    // Multiple LTM-Egress-ID TLVs found
    CfmPmPktAction_malformed_duplicate_ltm_egress_id CfmPmPktAction = "malformed-duplicate-ltm-egress-id"

    // Sender-ID TLV is too short
    CfmPmPktAction_malformed_sender_id_size CfmPmPktAction = "malformed-sender-id-size"

    // Sender-ID TLV is too short to contain the
    // Chassis ID
    CfmPmPktAction_malformed_chassis_id_size CfmPmPktAction = "malformed-chassis-id-size"

    // Sender-ID TLV is too short to contain the
    // management address domain
    CfmPmPktAction_malformed_mgmt_address_domain_size CfmPmPktAction = "malformed-mgmt-address-domain-size"

    // Sender-ID TLV is too short to contain the
    // management address
    CfmPmPktAction_malformed_mgmt_address_size CfmPmPktAction = "malformed-mgmt-address-size"

    // Port-status TLV is too short
    CfmPmPktAction_malformed_port_status_size CfmPmPktAction = "malformed-port-status-size"

    // Invalid Port status value
    CfmPmPktAction_malformed_port_status CfmPmPktAction = "malformed-port-status"

    // Interface-status TLV is too short
    CfmPmPktAction_malformed_interface_status_size CfmPmPktAction = "malformed-interface-status-size"

    // Invalid Interface status value
    CfmPmPktAction_malformed_interface_status CfmPmPktAction = "malformed-interface-status"

    // Organization-specific TLV is too short
    CfmPmPktAction_malformed_organization_specific_tlv_size CfmPmPktAction = "malformed-organization-specific-tlv-size"

    // Multiple MEP-name TLVs found
    CfmPmPktAction_malformed_duplicate_mep_name CfmPmPktAction = "malformed-duplicate-mep-name"

    // Multiple additional-interface-status TLVs found
    CfmPmPktAction_malformed_duplicate_additional_interface_status CfmPmPktAction = "malformed-duplicate-additional-interface-status"

    // LTR-Egress-ID TLV is too short
    CfmPmPktAction_malformed_ltr_egress_id_size CfmPmPktAction = "malformed-ltr-egress-id-size"

    // Reply-ingress TLV is too short
    CfmPmPktAction_malformed_reply_ingress_size CfmPmPktAction = "malformed-reply-ingress-size"

    // Invalid ingress-action value
    CfmPmPktAction_malformed_ingress_action CfmPmPktAction = "malformed-ingress-action"

    // Reply-ingress TLV has invalid MAC address
    CfmPmPktAction_malformed_reply_ingress_mac CfmPmPktAction = "malformed-reply-ingress-mac"

    // Reply-ingress TLV is too short to contain the
    // Port ID type
    CfmPmPktAction_malformed_ingress_port_length_size CfmPmPktAction = "malformed-ingress-port-length-size"

    // Reply-ingress TLV has a zero Port ID length
    CfmPmPktAction_malformed_ingress_port_id_length CfmPmPktAction = "malformed-ingress-port-id-length"

    // Reply-ingress TLV is too short to contain the
    // Port ID
    CfmPmPktAction_malformed_ingress_port_id_size CfmPmPktAction = "malformed-ingress-port-id-size"

    // Reply-egress TLV is too short
    CfmPmPktAction_malformed_reply_egress_size CfmPmPktAction = "malformed-reply-egress-size"

    // Invalid egress-action value
    CfmPmPktAction_malformed_egress_action CfmPmPktAction = "malformed-egress-action"

    // Reply-egress TLV has invalid MAC address
    CfmPmPktAction_malformed_reply_egress_mac CfmPmPktAction = "malformed-reply-egress-mac"

    // Reply-egress TLV is too short to contain the
    // Port ID type
    CfmPmPktAction_malformed_egress_port_length_size CfmPmPktAction = "malformed-egress-port-length-size"

    // Reply-egress TLV has a zero Port ID length
    CfmPmPktAction_malformed_egress_port_id_length CfmPmPktAction = "malformed-egress-port-id-length"

    // Reply-egress TLV is too short to contain the
    // Port ID
    CfmPmPktAction_malformed_egress_port_id_size CfmPmPktAction = "malformed-egress-port-id-size"

    // LTM-Egress_ID TLV is too short
    CfmPmPktAction_malformed_ltm_egress_id_size CfmPmPktAction = "malformed-ltm-egress-id-size"

    // MEP-name TLV is too short
    CfmPmPktAction_malformed_mep_name_size CfmPmPktAction = "malformed-mep-name-size"

    // MEP-name TLV is too short to contain a MEP name
    CfmPmPktAction_malformed_mep_name_name_length CfmPmPktAction = "malformed-mep-name-name-length"

    // Additional-interface-status is too short
    CfmPmPktAction_malformed_additional_interface_status_size CfmPmPktAction = "malformed-additional-interface-status-size"

    // Invalid additional interface status
    CfmPmPktAction_malformed_additional_interface_status CfmPmPktAction = "malformed-additional-interface-status"

    // CCM has a zero CCM interval
    CfmPmPktAction_malformed_ccm_interval CfmPmPktAction = "malformed-ccm-interval"

    // CCM has a MAC-address MDID but the MDID is the
    // wrong length
    CfmPmPktAction_malformed_mdid_mac_address_length CfmPmPktAction = "malformed-mdid-mac-address-length"

    // CCM has an invalid MDID length
    CfmPmPktAction_malformed_mdid_length CfmPmPktAction = "malformed-mdid-length"

    // CCM has an invalid Short MA Name length
    CfmPmPktAction_malformed_sman_length CfmPmPktAction = "malformed-sman-length"

    // CCM has a VID or 16-bit Short MA Name but a
    // mismatched length
    CfmPmPktAction_malformed_sman2_byte_length CfmPmPktAction = "malformed-sman2-byte-length"

    // CCM has a VPNID Short MA Name but a mismatched
    // length
    CfmPmPktAction_malformed_sman_vpn_id_length CfmPmPktAction = "malformed-sman-vpn-id-length"

    // ELR has no ELR Reply TLVs
    CfmPmPktAction_malformed_elr_no_reply_tlv CfmPmPktAction = "malformed-elr-no-reply-tlv"

    // ELR Reply Egress TLVs not all adjacent
    CfmPmPktAction_malformed_separate_elr_reply_egress CfmPmPktAction = "malformed-separate-elr-reply-egress"

    // DCM has a multicast destination MAC
    CfmPmPktAction_malformed_dcm_destination_multicast CfmPmPktAction = "malformed-dcm-destination-multicast"

    // DCM is too short to contain an Embedded PDU
    CfmPmPktAction_malformed_dcm_embed_length CfmPmPktAction = "malformed-dcm-embed-length"

    // DCM Embedded PDU level does not match DCM level
    CfmPmPktAction_malformed_dcm_embed_level CfmPmPktAction = "malformed-dcm-embed-level"

    // DCM Embedded PDU version does not match DCM
    // version
    CfmPmPktAction_malformed_dcm_embed_version CfmPmPktAction = "malformed-dcm-embed-version"

    // Unknown ELR relay action
    CfmPmPktAction_malformed_elr_relay_action CfmPmPktAction = "malformed-elr-relay-action"

    // Reply Ingress TTL is not one greater than Reply
    // Egress TTL
    CfmPmPktAction_malformed_elr_tt_ls CfmPmPktAction = "malformed-elr-tt-ls"

    // Reply Ingress TTL present without ELR Reply
    // Ingress TLV
    CfmPmPktAction_malformed_elr_ttl_ingress CfmPmPktAction = "malformed-elr-ttl-ingress"

    // Reply Egress TTL present without ELR Reply
    // Egress TLV
    CfmPmPktAction_malformed_elr_ttl_egress CfmPmPktAction = "malformed-elr-ttl-egress"

    // ELM Destination MAC must not be unicast
    CfmPmPktAction_malformed_elm_destination_unicast CfmPmPktAction = "malformed-elm-destination-unicast"

    // ELM has no LTM Egress ID TLV
    CfmPmPktAction_malformed_elm_egress_id CfmPmPktAction = "malformed-elm-egress-id"

    // Embedded DCM OUI unrecognized
    CfmPmPktAction_malformed_dcm_embed_oui CfmPmPktAction = "malformed-dcm-embed-oui"

    // Embedded DCM Opcode is not ELM
    CfmPmPktAction_malformed_dcm_embed_opcode CfmPmPktAction = "malformed-dcm-embed-opcode"

    // ELM Constant Factor is zero
    CfmPmPktAction_malformed_elm_constant_zero CfmPmPktAction = "malformed-elm-constant-zero"

    // ELR Next-Hop Timeout is zero
    CfmPmPktAction_malformed_elr_timeout_zero CfmPmPktAction = "malformed-elr-timeout-zero"

    // Multiple Test TLVs found
    CfmPmPktAction_malformed_duplicate_test CfmPmPktAction = "malformed-duplicate-test"

    // Invalid source MAC address for DMM
    CfmPmPktAction_malformed_dmm_source_mac CfmPmPktAction = "malformed-dmm-source-mac"

    // Test TLV is too short
    CfmPmPktAction_malformed_test_size CfmPmPktAction = "malformed-test-size"

    // DMR has exactly one of its Rxf and Txb
    // timestamps unspecified
    CfmPmPktAction_malformed_dmr_time_stamps CfmPmPktAction = "malformed-dmr-time-stamps"

    // The format of one or more timestamps is invalid
    CfmPmPktAction_malformed_dm_time_stamp_fmt CfmPmPktAction = "malformed-dm-time-stamp-fmt"

    // AIS/LCK has invalid interval value (not 1
    // second or 1 minute)
    CfmPmPktAction_malformed_ais_interval CfmPmPktAction = "malformed-ais-interval"

    // Packet dropped due to interface being down
    CfmPmPktAction_filter_interface_down CfmPmPktAction = "filter-interface-down"

    // Packet dropped - not forwarded because
    // interface is in standby mode
    CfmPmPktAction_filter_forward_standby CfmPmPktAction = "filter-forward-standby"

    // CCM has an ICC-based format Short MA Name but a
    // mismatched length
    CfmPmPktAction_malformed_sman_icc_based_length CfmPmPktAction = "malformed-sman-icc-based-length"

    // Packet dropped - not forwarded in secondary HA
    // role
    CfmPmPktAction_filter_foward_issu_secondary CfmPmPktAction = "filter-foward-issu-secondary"

    // Packet dropped - not responded to because
    // interface is in standby mode
    CfmPmPktAction_filter_response_standby CfmPmPktAction = "filter-response-standby"

    // Packet dropped - not responded to in secondary
    // HA role
    CfmPmPktAction_filter_response_issu_secondary CfmPmPktAction = "filter-response-issu-secondary"
)

// SlaOperPacketPriority represents Priority scheme for packet priority
type SlaOperPacketPriority string

const (
    // Packet does not use any specified priority.
    SlaOperPacketPriority_priority_none SlaOperPacketPriority = "priority-none"

    // Packet uses a specified 3-bit COS priority
    // value.
    SlaOperPacketPriority_priority_cos SlaOperPacketPriority = "priority-cos"
)

// SlaOperTestPatternScheme represents Test pattern scheme for packet padding
type SlaOperTestPatternScheme string

const (
    // Packet is padded with a user-specified string
    SlaOperTestPatternScheme_hex SlaOperTestPatternScheme = "hex"

    // Packet is padded with a pseudo-random bit
    // sequence
    SlaOperTestPatternScheme_pseudo_random SlaOperTestPatternScheme = "pseudo-random"
)

// SlaOperBucket represents Type of SLA metric bucket
type SlaOperBucket string

const (
    // SLA metric bin
    SlaOperBucket_bucket_type_bins SlaOperBucket = "bucket-type-bins"

    // SLA metric sample
    SlaOperBucket_bucket_type_samples SlaOperBucket = "bucket-type-samples"
)

// SlaBucketSize represents Type of configuration of a bucket size
type SlaBucketSize string

const (
    // Bucket size is configured as buckets per probe
    SlaBucketSize_buckets_per_probe SlaBucketSize = "buckets-per-probe"

    // Bucket size is configured as probes per bucket
    SlaBucketSize_probes_per_bucket SlaBucketSize = "probes-per-bucket"
)

// SlaRecordableMetric represents Types of metrics that can be recorded by probes
type SlaRecordableMetric string

const (
    // Not a valid metric type
    SlaRecordableMetric_metric_invalid SlaRecordableMetric = "metric-invalid"

    // Round-trip Delay
    SlaRecordableMetric_metric_round_trip_delay SlaRecordableMetric = "metric-round-trip-delay"

    // One-way Delay (Source->Destination)
    SlaRecordableMetric_metric_one_way_delay_sd SlaRecordableMetric = "metric-one-way-delay-sd"

    // One-way Delay (Destination->Source)
    SlaRecordableMetric_metric_one_way_delay_ds SlaRecordableMetric = "metric-one-way-delay-ds"

    // Round-trip Jitter
    SlaRecordableMetric_metric_round_trip_jitter SlaRecordableMetric = "metric-round-trip-jitter"

    // One-way Jitter (Source->Destination)
    SlaRecordableMetric_metric_one_way_jitter_sd SlaRecordableMetric = "metric-one-way-jitter-sd"

    // One-way Jitter (Destination->Source)
    SlaRecordableMetric_metric_one_way_jitter_ds SlaRecordableMetric = "metric-one-way-jitter-ds"

    // One-way Frame Loss Ratio (Source->Destination)
    SlaRecordableMetric_metric_one_way_flr_sd SlaRecordableMetric = "metric-one-way-flr-sd"

    // One-way Frame Loss Ratio (Destination->Source)
    SlaRecordableMetric_metric_one_way_flr_ds SlaRecordableMetric = "metric-one-way-flr-ds"
)

// SlaOperOperation represents Type of SLA operation
type SlaOperOperation string

const (
    // Configured SLA operation
    SlaOperOperation_operation_type_configured SlaOperOperation = "operation-type-configured"

    // On-demand SLA operation
    SlaOperOperation_operation_type_ondemand SlaOperOperation = "operation-type-ondemand"
)

// Cfm
// CFM operational data
type Cfm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node table for node-specific operational data.
    Nodes Cfm_Nodes

    // Global operational data.
    Global Cfm_Global
}

func (cfm *Cfm) GetEntityData() *types.CommonEntityData {
    cfm.EntityData.YFilter = cfm.YFilter
    cfm.EntityData.YangName = "cfm"
    cfm.EntityData.BundleName = "cisco_ios_xr"
    cfm.EntityData.ParentYangName = "Cisco-IOS-XR-ethernet-cfm-oper"
    cfm.EntityData.SegmentPath = "Cisco-IOS-XR-ethernet-cfm-oper:cfm"
    cfm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cfm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cfm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cfm.EntityData.Children = make(map[string]types.YChild)
    cfm.EntityData.Children["nodes"] = types.YChild{"Nodes", &cfm.Nodes}
    cfm.EntityData.Children["global"] = types.YChild{"Global", &cfm.Global}
    cfm.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cfm.EntityData)
}

// Cfm_Nodes
// Node table for node-specific operational data
type Cfm_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // Cfm_Nodes_Node.
    Node []Cfm_Nodes_Node
}

func (nodes *Cfm_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "cfm"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Cfm_Nodes_Node
// Node-specific data for a particular node
type Cfm_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Interface AIS table.
    InterfaceAises Cfm_Nodes_Node_InterfaceAises

    // Interface Statistics table.
    InterfaceStatistics Cfm_Nodes_Node_InterfaceStatistics

    // Summary.
    Summary Cfm_Nodes_Node_Summary

    // CCMLearningDatabase table.
    CcmLearningDatabases Cfm_Nodes_Node_CcmLearningDatabases
}

func (node *Cfm_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["interface-aises"] = types.YChild{"InterfaceAises", &node.InterfaceAises}
    node.EntityData.Children["interface-statistics"] = types.YChild{"InterfaceStatistics", &node.InterfaceStatistics}
    node.EntityData.Children["summary"] = types.YChild{"Summary", &node.Summary}
    node.EntityData.Children["ccm-learning-databases"] = types.YChild{"CcmLearningDatabases", &node.CcmLearningDatabases}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node"] = types.YLeaf{"Node", node.Node}
    return &(node.EntityData)
}

// Cfm_Nodes_Node_InterfaceAises
// Interface AIS table
type Cfm_Nodes_Node_InterfaceAises struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AIS statistics for a particular interface. The type is slice of
    // Cfm_Nodes_Node_InterfaceAises_InterfaceAis.
    InterfaceAis []Cfm_Nodes_Node_InterfaceAises_InterfaceAis
}

func (interfaceAises *Cfm_Nodes_Node_InterfaceAises) GetEntityData() *types.CommonEntityData {
    interfaceAises.EntityData.YFilter = interfaceAises.YFilter
    interfaceAises.EntityData.YangName = "interface-aises"
    interfaceAises.EntityData.BundleName = "cisco_ios_xr"
    interfaceAises.EntityData.ParentYangName = "node"
    interfaceAises.EntityData.SegmentPath = "interface-aises"
    interfaceAises.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAises.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAises.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAises.EntityData.Children = make(map[string]types.YChild)
    interfaceAises.EntityData.Children["interface-ais"] = types.YChild{"InterfaceAis", nil}
    for i := range interfaceAises.InterfaceAis {
        interfaceAises.EntityData.Children[types.GetSegmentPath(&interfaceAises.InterfaceAis[i])] = types.YChild{"InterfaceAis", &interfaceAises.InterfaceAis[i]}
    }
    interfaceAises.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceAises.EntityData)
}

// Cfm_Nodes_Node_InterfaceAises_InterfaceAis
// AIS statistics for a particular interface
type Cfm_Nodes_Node_InterfaceAises_InterfaceAis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // This attribute is a key. AIS Direction. The type is CfmAisDir.
    Direction interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // IM Interface state. The type is string.
    InterfaceState interface{}

    // Interface interworking state. The type is CfmBagIwState.
    InterworkingState interface{}

    // STP state. The type is CfmBagStpState.
    StpState interface{}

    // AIS statistics.
    Statistics Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics
}

func (interfaceAis *Cfm_Nodes_Node_InterfaceAises_InterfaceAis) GetEntityData() *types.CommonEntityData {
    interfaceAis.EntityData.YFilter = interfaceAis.YFilter
    interfaceAis.EntityData.YangName = "interface-ais"
    interfaceAis.EntityData.BundleName = "cisco_ios_xr"
    interfaceAis.EntityData.ParentYangName = "interface-aises"
    interfaceAis.EntityData.SegmentPath = "interface-ais" + "[interface-name='" + fmt.Sprintf("%v", interfaceAis.InterfaceName) + "']" + "[direction='" + fmt.Sprintf("%v", interfaceAis.Direction) + "']"
    interfaceAis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAis.EntityData.Children = make(map[string]types.YChild)
    interfaceAis.EntityData.Children["statistics"] = types.YChild{"Statistics", &interfaceAis.Statistics}
    interfaceAis.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceAis.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceAis.InterfaceName}
    interfaceAis.EntityData.Leafs["direction"] = types.YLeaf{"Direction", interfaceAis.Direction}
    interfaceAis.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", interfaceAis.Interface_}
    interfaceAis.EntityData.Leafs["interface-state"] = types.YLeaf{"InterfaceState", interfaceAis.InterfaceState}
    interfaceAis.EntityData.Leafs["interworking-state"] = types.YLeaf{"InterworkingState", interfaceAis.InterworkingState}
    interfaceAis.EntityData.Leafs["stp-state"] = types.YLeaf{"StpState", interfaceAis.StpState}
    return &(interfaceAis.EntityData)
}

// Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics
// AIS statistics
type Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Direction of AIS packets. The type is CfmBagDirection.
    Direction interface{}

    // Level of the lowest MEP transmitting AIS. The type is CfmBagMdLevel.
    LowestLevel interface{}

    // Level that AIS packets are transmitted on. The type is CfmBagMdLevel.
    TransmissionLevel interface{}

    // Interval at which AIS packets are transmitted. The type is
    // CfmBagAisInterval.
    TransmissionInterval interface{}

    // Total number of packets sent by the transmitting MEP. The type is
    // interface{} with range: 0..4294967295.
    SentPackets interface{}

    // Levels of other MEPs receiving AIS. The type is slice of CfmBagMdLevel.
    ViaLevel []interface{}

    // Defects detected.
    Defects Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects

    // Time elapsed since sending last started.
    LastStarted Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted
}

func (statistics *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "interface-ais"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["defects"] = types.YChild{"Defects", &statistics.Defects}
    statistics.EntityData.Children["last-started"] = types.YChild{"LastStarted", &statistics.LastStarted}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["direction"] = types.YLeaf{"Direction", statistics.Direction}
    statistics.EntityData.Leafs["lowest-level"] = types.YLeaf{"LowestLevel", statistics.LowestLevel}
    statistics.EntityData.Leafs["transmission-level"] = types.YLeaf{"TransmissionLevel", statistics.TransmissionLevel}
    statistics.EntityData.Leafs["transmission-interval"] = types.YLeaf{"TransmissionInterval", statistics.TransmissionInterval}
    statistics.EntityData.Leafs["sent-packets"] = types.YLeaf{"SentPackets", statistics.SentPackets}
    statistics.EntityData.Leafs["via-level"] = types.YLeaf{"ViaLevel", statistics.ViaLevel}
    return &(statistics.EntityData)
}

// Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects
// Defects detected
type Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AIS or LCK received. The type is bool.
    AisReceived interface{}

    // Number of peer MEPs that have timed out. The type is interface{} with
    // range: 0..4294967295.
    PeerMepsThatTimedOut interface{}

    // Number of missing peer MEPs. The type is interface{} with range:
    // 0..4294967295.
    Missing interface{}

    // Number of missing auto cross-check MEPs. The type is interface{} with
    // range: 0..4294967295.
    AutoMissing interface{}

    // Number of unexpected peer MEPs. The type is interface{} with range:
    // 0..4294967295.
    Unexpected interface{}

    // The local port or interface is down. The type is bool.
    LocalPortStatus interface{}

    // A peer port or interface is down. The type is bool.
    PeerPortStatus interface{}

    // Defects detected from remote MEPs.
    RemoteMepsDefects Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects
}

func (defects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects) GetEntityData() *types.CommonEntityData {
    defects.EntityData.YFilter = defects.YFilter
    defects.EntityData.YangName = "defects"
    defects.EntityData.BundleName = "cisco_ios_xr"
    defects.EntityData.ParentYangName = "statistics"
    defects.EntityData.SegmentPath = "defects"
    defects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defects.EntityData.Children = make(map[string]types.YChild)
    defects.EntityData.Children["remote-meps-defects"] = types.YChild{"RemoteMepsDefects", &defects.RemoteMepsDefects}
    defects.EntityData.Leafs = make(map[string]types.YLeaf)
    defects.EntityData.Leafs["ais-received"] = types.YLeaf{"AisReceived", defects.AisReceived}
    defects.EntityData.Leafs["peer-meps-that-timed-out"] = types.YLeaf{"PeerMepsThatTimedOut", defects.PeerMepsThatTimedOut}
    defects.EntityData.Leafs["missing"] = types.YLeaf{"Missing", defects.Missing}
    defects.EntityData.Leafs["auto-missing"] = types.YLeaf{"AutoMissing", defects.AutoMissing}
    defects.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", defects.Unexpected}
    defects.EntityData.Leafs["local-port-status"] = types.YLeaf{"LocalPortStatus", defects.LocalPortStatus}
    defects.EntityData.Leafs["peer-port-status"] = types.YLeaf{"PeerPortStatus", defects.PeerPortStatus}
    return &(defects.EntityData)
}

// Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects
// Defects detected from remote MEPs
type Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timed out (loss threshold exceeded). The type is bool.
    LossThresholdExceeded interface{}

    // Invalid level. The type is bool.
    InvalidLevel interface{}

    // Invalid MAID. The type is bool.
    InvalidMaid interface{}

    // Invalid CCM interval. The type is bool.
    InvalidCcmInterval interface{}

    // Loop detected (our MAC address received). The type is bool.
    ReceivedOurMac interface{}

    // Configuration Error (our MEP ID received). The type is bool.
    ReceivedOurMepId interface{}

    // Remote defection indication received. The type is bool.
    ReceivedRdi interface{}
}

func (remoteMepsDefects *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_Defects_RemoteMepsDefects) GetEntityData() *types.CommonEntityData {
    remoteMepsDefects.EntityData.YFilter = remoteMepsDefects.YFilter
    remoteMepsDefects.EntityData.YangName = "remote-meps-defects"
    remoteMepsDefects.EntityData.BundleName = "cisco_ios_xr"
    remoteMepsDefects.EntityData.ParentYangName = "defects"
    remoteMepsDefects.EntityData.SegmentPath = "remote-meps-defects"
    remoteMepsDefects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteMepsDefects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteMepsDefects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteMepsDefects.EntityData.Children = make(map[string]types.YChild)
    remoteMepsDefects.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteMepsDefects.EntityData.Leafs["loss-threshold-exceeded"] = types.YLeaf{"LossThresholdExceeded", remoteMepsDefects.LossThresholdExceeded}
    remoteMepsDefects.EntityData.Leafs["invalid-level"] = types.YLeaf{"InvalidLevel", remoteMepsDefects.InvalidLevel}
    remoteMepsDefects.EntityData.Leafs["invalid-maid"] = types.YLeaf{"InvalidMaid", remoteMepsDefects.InvalidMaid}
    remoteMepsDefects.EntityData.Leafs["invalid-ccm-interval"] = types.YLeaf{"InvalidCcmInterval", remoteMepsDefects.InvalidCcmInterval}
    remoteMepsDefects.EntityData.Leafs["received-our-mac"] = types.YLeaf{"ReceivedOurMac", remoteMepsDefects.ReceivedOurMac}
    remoteMepsDefects.EntityData.Leafs["received-our-mep-id"] = types.YLeaf{"ReceivedOurMepId", remoteMepsDefects.ReceivedOurMepId}
    remoteMepsDefects.EntityData.Leafs["received-rdi"] = types.YLeaf{"ReceivedRdi", remoteMepsDefects.ReceivedRdi}
    return &(remoteMepsDefects.EntityData)
}

// Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted
// Time elapsed since sending last started
type Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastStarted *Cfm_Nodes_Node_InterfaceAises_InterfaceAis_Statistics_LastStarted) GetEntityData() *types.CommonEntityData {
    lastStarted.EntityData.YFilter = lastStarted.YFilter
    lastStarted.EntityData.YangName = "last-started"
    lastStarted.EntityData.BundleName = "cisco_ios_xr"
    lastStarted.EntityData.ParentYangName = "statistics"
    lastStarted.EntityData.SegmentPath = "last-started"
    lastStarted.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastStarted.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastStarted.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastStarted.EntityData.Children = make(map[string]types.YChild)
    lastStarted.EntityData.Leafs = make(map[string]types.YLeaf)
    lastStarted.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", lastStarted.Seconds}
    lastStarted.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", lastStarted.Nanoseconds}
    return &(lastStarted.EntityData)
}

// Cfm_Nodes_Node_InterfaceStatistics
// Interface Statistics table
type Cfm_Nodes_Node_InterfaceStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Counters for a particular interface. The type is slice of
    // Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic.
    InterfaceStatistic []Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic
}

func (interfaceStatistics *Cfm_Nodes_Node_InterfaceStatistics) GetEntityData() *types.CommonEntityData {
    interfaceStatistics.EntityData.YFilter = interfaceStatistics.YFilter
    interfaceStatistics.EntityData.YangName = "interface-statistics"
    interfaceStatistics.EntityData.BundleName = "cisco_ios_xr"
    interfaceStatistics.EntityData.ParentYangName = "node"
    interfaceStatistics.EntityData.SegmentPath = "interface-statistics"
    interfaceStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStatistics.EntityData.Children = make(map[string]types.YChild)
    interfaceStatistics.EntityData.Children["interface-statistic"] = types.YChild{"InterfaceStatistic", nil}
    for i := range interfaceStatistics.InterfaceStatistic {
        interfaceStatistics.EntityData.Children[types.GetSegmentPath(&interfaceStatistics.InterfaceStatistic[i])] = types.YChild{"InterfaceStatistic", &interfaceStatistics.InterfaceStatistic[i]}
    }
    interfaceStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceStatistics.EntityData)
}

// Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic
// Counters for a particular interface
type Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceXr interface{}

    // EFP statistics.
    Statistics Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics
}

func (interfaceStatistic *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic) GetEntityData() *types.CommonEntityData {
    interfaceStatistic.EntityData.YFilter = interfaceStatistic.YFilter
    interfaceStatistic.EntityData.YangName = "interface-statistic"
    interfaceStatistic.EntityData.BundleName = "cisco_ios_xr"
    interfaceStatistic.EntityData.ParentYangName = "interface-statistics"
    interfaceStatistic.EntityData.SegmentPath = "interface-statistic" + "[interface='" + fmt.Sprintf("%v", interfaceStatistic.Interface_) + "']"
    interfaceStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStatistic.EntityData.Children = make(map[string]types.YChild)
    interfaceStatistic.EntityData.Children["statistics"] = types.YChild{"Statistics", &interfaceStatistic.Statistics}
    interfaceStatistic.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceStatistic.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", interfaceStatistic.Interface_}
    interfaceStatistic.EntityData.Leafs["interface-xr"] = types.YLeaf{"InterfaceXr", interfaceStatistic.InterfaceXr}
    return &(interfaceStatistic.EntityData)
}

// Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics
// EFP statistics
type Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of malformed packets received at this EFP. The type is interface{}
    // with range: 0..18446744073709551615.
    MalformedPackets interface{}

    // Number of packets dropped at this EFP. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}

    // Opcode for last malformed packet. The type is CfmBagOpcode.
    LastMalformedOpcode interface{}

    // Reason last malformed packet was malformed. The type is CfmPmPktAction.
    LastMalformedReason interface{}
}

func (statistics *Cfm_Nodes_Node_InterfaceStatistics_InterfaceStatistic_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "interface-statistic"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["malformed-packets"] = types.YLeaf{"MalformedPackets", statistics.MalformedPackets}
    statistics.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", statistics.DroppedPackets}
    statistics.EntityData.Leafs["last-malformed-opcode"] = types.YLeaf{"LastMalformedOpcode", statistics.LastMalformedOpcode}
    statistics.EntityData.Leafs["last-malformed-reason"] = types.YLeaf{"LastMalformedReason", statistics.LastMalformedReason}
    return &(statistics.EntityData)
}

// Cfm_Nodes_Node_Summary
// Summary
type Cfm_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of domains in the CFM database. The type is interface{} with
    // range: 0..4294967295.
    Domains interface{}

    // The number of services in the CFM database. The type is interface{} with
    // range: 0..4294967295.
    Services interface{}

    // The combined rate of CCMs on this card. The type is interface{} with range:
    // 0..4294967295.
    CcmRate interface{}

    // The number of local MEPs in the CFM database. The type is interface{} with
    // range: 0..4294967295.
    LocalMeps interface{}

    // The number of operational local MEPs. The type is interface{} with range:
    // 0..4294967295.
    OperationalLocalMeps interface{}

    // The number of down-MEPs. The type is interface{} with range: 0..4294967295.
    DownMeps interface{}

    // The number of up-MEPs. The type is interface{} with range: 0..4294967295.
    UpMeps interface{}

    // The number of MEPs for which CCM processing has been offloaded. The type is
    // interface{} with range: 0..4294967295.
    Offloaded interface{}

    // The number of MEPs offloaded with CCMs at 3.3ms intervals. The type is
    // interface{} with range: 0..4294967295.
    OffloadedAt33Ms interface{}

    // The number of MEPs offloaded with CCMs at 10ms intervals. The type is
    // interface{} with range: 0..4294967295.
    OffloadedAt10Ms interface{}

    // The number of local MEPs disabled due to configuration errors. The type is
    // interface{} with range: 0..4294967295.
    DisabledMisconfigured interface{}

    // The number of local MEPs disabled due to lack of resources. The type is
    // interface{} with range: 0..4294967295.
    DisabledOutOfResources interface{}

    // The number of local MEPs disabled due to operational errors. The type is
    // interface{} with range: 0..4294967295.
    DisabledOperationalError interface{}

    // The number of peer MEPs. The type is interface{} with range: 0..4294967295.
    PeerMeps interface{}

    // The number of operational peer MEPs recorded in the CFM database. The type
    // is interface{} with range: 0..4294967295.
    OperationalPeerMeps interface{}

    // The number of peer MEPs with defects. The type is interface{} with range:
    // 0..4294967295.
    PeerMepsWithDefects interface{}

    // The number of peer MEPs without defects. The type is interface{} with
    // range: 0..4294967295.
    PeerMepsWithoutDefects interface{}

    // The number of peer MEPs that have timed out. The type is interface{} with
    // range: 0..4294967295.
    PeerMepsTimedOut interface{}

    // The number of MIPs. The type is interface{} with range: 0..4294967295.
    Mips interface{}

    // The number of interfaces running CFM. The type is interface{} with range:
    // 0..4294967295.
    Interfaces interface{}

    // Number or bridge domains and crossconnects. The type is interface{} with
    // range: 0..4294967295.
    BridgeDomainsAndXconnects interface{}

    // Number of traceroute cache entries. The type is interface{} with range:
    // 0..4294967295.
    TracerouteCacheEntries interface{}

    // Number of traceroute cache replies. The type is interface{} with range:
    // 0..4294967295.
    TracerouteCacheReplies interface{}

    // Number of entries in the CCM learning database. The type is interface{}
    // with range: 0..4294967295.
    CcmLearningDbEntries interface{}

    // ISSU Role of CFM-D, if any. The type is CfmBagIssuRole.
    IssuRole interface{}

    // Number of BNM Enabled Links. The type is interface{} with range:
    // 0..4294967295.
    BnmEnabledLinks interface{}
}

func (summary *Cfm_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["domains"] = types.YLeaf{"Domains", summary.Domains}
    summary.EntityData.Leafs["services"] = types.YLeaf{"Services", summary.Services}
    summary.EntityData.Leafs["ccm-rate"] = types.YLeaf{"CcmRate", summary.CcmRate}
    summary.EntityData.Leafs["local-meps"] = types.YLeaf{"LocalMeps", summary.LocalMeps}
    summary.EntityData.Leafs["operational-local-meps"] = types.YLeaf{"OperationalLocalMeps", summary.OperationalLocalMeps}
    summary.EntityData.Leafs["down-meps"] = types.YLeaf{"DownMeps", summary.DownMeps}
    summary.EntityData.Leafs["up-meps"] = types.YLeaf{"UpMeps", summary.UpMeps}
    summary.EntityData.Leafs["offloaded"] = types.YLeaf{"Offloaded", summary.Offloaded}
    summary.EntityData.Leafs["offloaded-at3-3ms"] = types.YLeaf{"OffloadedAt33Ms", summary.OffloadedAt33Ms}
    summary.EntityData.Leafs["offloaded-at10ms"] = types.YLeaf{"OffloadedAt10Ms", summary.OffloadedAt10Ms}
    summary.EntityData.Leafs["disabled-misconfigured"] = types.YLeaf{"DisabledMisconfigured", summary.DisabledMisconfigured}
    summary.EntityData.Leafs["disabled-out-of-resources"] = types.YLeaf{"DisabledOutOfResources", summary.DisabledOutOfResources}
    summary.EntityData.Leafs["disabled-operational-error"] = types.YLeaf{"DisabledOperationalError", summary.DisabledOperationalError}
    summary.EntityData.Leafs["peer-meps"] = types.YLeaf{"PeerMeps", summary.PeerMeps}
    summary.EntityData.Leafs["operational-peer-meps"] = types.YLeaf{"OperationalPeerMeps", summary.OperationalPeerMeps}
    summary.EntityData.Leafs["peer-meps-with-defects"] = types.YLeaf{"PeerMepsWithDefects", summary.PeerMepsWithDefects}
    summary.EntityData.Leafs["peer-meps-without-defects"] = types.YLeaf{"PeerMepsWithoutDefects", summary.PeerMepsWithoutDefects}
    summary.EntityData.Leafs["peer-meps-timed-out"] = types.YLeaf{"PeerMepsTimedOut", summary.PeerMepsTimedOut}
    summary.EntityData.Leafs["mips"] = types.YLeaf{"Mips", summary.Mips}
    summary.EntityData.Leafs["interfaces"] = types.YLeaf{"Interfaces", summary.Interfaces}
    summary.EntityData.Leafs["bridge-domains-and-xconnects"] = types.YLeaf{"BridgeDomainsAndXconnects", summary.BridgeDomainsAndXconnects}
    summary.EntityData.Leafs["traceroute-cache-entries"] = types.YLeaf{"TracerouteCacheEntries", summary.TracerouteCacheEntries}
    summary.EntityData.Leafs["traceroute-cache-replies"] = types.YLeaf{"TracerouteCacheReplies", summary.TracerouteCacheReplies}
    summary.EntityData.Leafs["ccm-learning-db-entries"] = types.YLeaf{"CcmLearningDbEntries", summary.CcmLearningDbEntries}
    summary.EntityData.Leafs["issu-role"] = types.YLeaf{"IssuRole", summary.IssuRole}
    summary.EntityData.Leafs["bnm-enabled-links"] = types.YLeaf{"BnmEnabledLinks", summary.BnmEnabledLinks}
    return &(summary.EntityData)
}

// Cfm_Nodes_Node_CcmLearningDatabases
// CCMLearningDatabase table
type Cfm_Nodes_Node_CcmLearningDatabases struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CCM Learning Database entry. The type is slice of
    // Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase.
    CcmLearningDatabase []Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase
}

func (ccmLearningDatabases *Cfm_Nodes_Node_CcmLearningDatabases) GetEntityData() *types.CommonEntityData {
    ccmLearningDatabases.EntityData.YFilter = ccmLearningDatabases.YFilter
    ccmLearningDatabases.EntityData.YangName = "ccm-learning-databases"
    ccmLearningDatabases.EntityData.BundleName = "cisco_ios_xr"
    ccmLearningDatabases.EntityData.ParentYangName = "node"
    ccmLearningDatabases.EntityData.SegmentPath = "ccm-learning-databases"
    ccmLearningDatabases.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ccmLearningDatabases.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ccmLearningDatabases.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ccmLearningDatabases.EntityData.Children = make(map[string]types.YChild)
    ccmLearningDatabases.EntityData.Children["ccm-learning-database"] = types.YChild{"CcmLearningDatabase", nil}
    for i := range ccmLearningDatabases.CcmLearningDatabase {
        ccmLearningDatabases.EntityData.Children[types.GetSegmentPath(&ccmLearningDatabases.CcmLearningDatabase[i])] = types.YChild{"CcmLearningDatabase", &ccmLearningDatabases.CcmLearningDatabase[i]}
    }
    ccmLearningDatabases.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ccmLearningDatabases.EntityData)
}

// Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase
// CCM Learning Database entry
type Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Maintenance domain name. The type is string.
    DomainXr interface{}

    // Maintenance level. The type is CfmBagMdLevel.
    Level interface{}

    // Maintenance association name. The type is string.
    ServiceXr interface{}

    // Source MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SourceMacAddress interface{}

    // The XID of the ingress interface for the CCM. The type is interface{} with
    // range: 0..4294967295.
    IngressInterface interface{}

    // The XID is stale and may have been reused for a different interface. The
    // type is bool.
    Stale interface{}

    // String representation of the Bridge Domain or Cross-Connect associated with
    // the ingress XID. The type is string.
    IngressInterfaceString interface{}
}

func (ccmLearningDatabase *Cfm_Nodes_Node_CcmLearningDatabases_CcmLearningDatabase) GetEntityData() *types.CommonEntityData {
    ccmLearningDatabase.EntityData.YFilter = ccmLearningDatabase.YFilter
    ccmLearningDatabase.EntityData.YangName = "ccm-learning-database"
    ccmLearningDatabase.EntityData.BundleName = "cisco_ios_xr"
    ccmLearningDatabase.EntityData.ParentYangName = "ccm-learning-databases"
    ccmLearningDatabase.EntityData.SegmentPath = "ccm-learning-database" + "[domain='" + fmt.Sprintf("%v", ccmLearningDatabase.Domain) + "']" + "[service='" + fmt.Sprintf("%v", ccmLearningDatabase.Service) + "']" + "[mac-address='" + fmt.Sprintf("%v", ccmLearningDatabase.MacAddress) + "']"
    ccmLearningDatabase.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ccmLearningDatabase.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ccmLearningDatabase.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ccmLearningDatabase.EntityData.Children = make(map[string]types.YChild)
    ccmLearningDatabase.EntityData.Leafs = make(map[string]types.YLeaf)
    ccmLearningDatabase.EntityData.Leafs["domain"] = types.YLeaf{"Domain", ccmLearningDatabase.Domain}
    ccmLearningDatabase.EntityData.Leafs["service"] = types.YLeaf{"Service", ccmLearningDatabase.Service}
    ccmLearningDatabase.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", ccmLearningDatabase.MacAddress}
    ccmLearningDatabase.EntityData.Leafs["domain-xr"] = types.YLeaf{"DomainXr", ccmLearningDatabase.DomainXr}
    ccmLearningDatabase.EntityData.Leafs["level"] = types.YLeaf{"Level", ccmLearningDatabase.Level}
    ccmLearningDatabase.EntityData.Leafs["service-xr"] = types.YLeaf{"ServiceXr", ccmLearningDatabase.ServiceXr}
    ccmLearningDatabase.EntityData.Leafs["source-mac-address"] = types.YLeaf{"SourceMacAddress", ccmLearningDatabase.SourceMacAddress}
    ccmLearningDatabase.EntityData.Leafs["ingress-interface"] = types.YLeaf{"IngressInterface", ccmLearningDatabase.IngressInterface}
    ccmLearningDatabase.EntityData.Leafs["stale"] = types.YLeaf{"Stale", ccmLearningDatabase.Stale}
    ccmLearningDatabase.EntityData.Leafs["ingress-interface-string"] = types.YLeaf{"IngressInterfaceString", ccmLearningDatabase.IngressInterfaceString}
    return &(ccmLearningDatabase.EntityData)
}

// Cfm_Global
// Global operational data
type Cfm_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Incomplete Traceroute table.
    IncompleteTraceroutes Cfm_Global_IncompleteTraceroutes

    // Maintenance Points table.
    MaintenancePoints Cfm_Global_MaintenancePoints

    // Global configuration errors table.
    GlobalConfigurationErrors Cfm_Global_GlobalConfigurationErrors

    // MEP configuration errors table.
    MepConfigurationErrors Cfm_Global_MepConfigurationErrors

    // Traceroute Cache table.
    TracerouteCaches Cfm_Global_TracerouteCaches

    // Local MEPs table.
    LocalMeps Cfm_Global_LocalMeps

    // Peer MEPs table Version 2.
    PeerMePv2S Cfm_Global_PeerMePv2S
}

func (global *Cfm_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "cfm"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["incomplete-traceroutes"] = types.YChild{"IncompleteTraceroutes", &global.IncompleteTraceroutes}
    global.EntityData.Children["maintenance-points"] = types.YChild{"MaintenancePoints", &global.MaintenancePoints}
    global.EntityData.Children["global-configuration-errors"] = types.YChild{"GlobalConfigurationErrors", &global.GlobalConfigurationErrors}
    global.EntityData.Children["mep-configuration-errors"] = types.YChild{"MepConfigurationErrors", &global.MepConfigurationErrors}
    global.EntityData.Children["traceroute-caches"] = types.YChild{"TracerouteCaches", &global.TracerouteCaches}
    global.EntityData.Children["local-meps"] = types.YChild{"LocalMeps", &global.LocalMeps}
    global.EntityData.Children["peer-me-pv2s"] = types.YChild{"PeerMePv2S", &global.PeerMePv2S}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(global.EntityData)
}

// Cfm_Global_IncompleteTraceroutes
// Incomplete Traceroute table
type Cfm_Global_IncompleteTraceroutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a traceroute operation that has not yet timed out. The
    // type is slice of Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute.
    IncompleteTraceroute []Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute
}

func (incompleteTraceroutes *Cfm_Global_IncompleteTraceroutes) GetEntityData() *types.CommonEntityData {
    incompleteTraceroutes.EntityData.YFilter = incompleteTraceroutes.YFilter
    incompleteTraceroutes.EntityData.YangName = "incomplete-traceroutes"
    incompleteTraceroutes.EntityData.BundleName = "cisco_ios_xr"
    incompleteTraceroutes.EntityData.ParentYangName = "global"
    incompleteTraceroutes.EntityData.SegmentPath = "incomplete-traceroutes"
    incompleteTraceroutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incompleteTraceroutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incompleteTraceroutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incompleteTraceroutes.EntityData.Children = make(map[string]types.YChild)
    incompleteTraceroutes.EntityData.Children["incomplete-traceroute"] = types.YChild{"IncompleteTraceroute", nil}
    for i := range incompleteTraceroutes.IncompleteTraceroute {
        incompleteTraceroutes.EntityData.Children[types.GetSegmentPath(&incompleteTraceroutes.IncompleteTraceroute[i])] = types.YChild{"IncompleteTraceroute", &incompleteTraceroutes.IncompleteTraceroute[i]}
    }
    incompleteTraceroutes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(incompleteTraceroutes.EntityData)
}

// Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute
// Information about a traceroute operation that
// has not yet timed out
type Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. MEP ID. The type is interface{} with range:
    // 1..8191.
    MepId interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // This attribute is a key. Transaction ID. The type is interface{} with
    // range: -2147483648..2147483647.
    TransactionId interface{}

    // Time (in seconds) before the traceroute completes. The type is interface{}
    // with range: 0..18446744073709551615. Units are second.
    TimeLeft interface{}

    // Information about the traceroute operation.
    TracerouteInformation Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation
}

func (incompleteTraceroute *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute) GetEntityData() *types.CommonEntityData {
    incompleteTraceroute.EntityData.YFilter = incompleteTraceroute.YFilter
    incompleteTraceroute.EntityData.YangName = "incomplete-traceroute"
    incompleteTraceroute.EntityData.BundleName = "cisco_ios_xr"
    incompleteTraceroute.EntityData.ParentYangName = "incomplete-traceroutes"
    incompleteTraceroute.EntityData.SegmentPath = "incomplete-traceroute" + "[domain='" + fmt.Sprintf("%v", incompleteTraceroute.Domain) + "']" + "[service='" + fmt.Sprintf("%v", incompleteTraceroute.Service) + "']" + "[mep-id='" + fmt.Sprintf("%v", incompleteTraceroute.MepId) + "']" + "[interface='" + fmt.Sprintf("%v", incompleteTraceroute.Interface_) + "']" + "[transaction-id='" + fmt.Sprintf("%v", incompleteTraceroute.TransactionId) + "']"
    incompleteTraceroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incompleteTraceroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incompleteTraceroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incompleteTraceroute.EntityData.Children = make(map[string]types.YChild)
    incompleteTraceroute.EntityData.Children["traceroute-information"] = types.YChild{"TracerouteInformation", &incompleteTraceroute.TracerouteInformation}
    incompleteTraceroute.EntityData.Leafs = make(map[string]types.YLeaf)
    incompleteTraceroute.EntityData.Leafs["domain"] = types.YLeaf{"Domain", incompleteTraceroute.Domain}
    incompleteTraceroute.EntityData.Leafs["service"] = types.YLeaf{"Service", incompleteTraceroute.Service}
    incompleteTraceroute.EntityData.Leafs["mep-id"] = types.YLeaf{"MepId", incompleteTraceroute.MepId}
    incompleteTraceroute.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", incompleteTraceroute.Interface_}
    incompleteTraceroute.EntityData.Leafs["transaction-id"] = types.YLeaf{"TransactionId", incompleteTraceroute.TransactionId}
    incompleteTraceroute.EntityData.Leafs["time-left"] = types.YLeaf{"TimeLeft", incompleteTraceroute.TimeLeft}
    return &(incompleteTraceroute.EntityData)
}

// Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation
// Information about the traceroute operation
type Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maintenance domain name. The type is string.
    Domain interface{}

    // Service name. The type is string.
    Service interface{}

    // Maintenance level. The type is CfmBagMdLevel.
    Level interface{}

    // Source MEP ID. The type is interface{} with range: 0..65535.
    SourceMepId interface{}

    // Source interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SourceInterface interface{}

    // Source MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SourceMacAddress interface{}

    // Target MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    TargetMacAddress interface{}

    // Directed MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    DirectedMacAddress interface{}

    // Target MEP ID. The type is interface{} with range: 0..65535.
    TargetMepId interface{}

    // Timestamp of initiation time (seconds). The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    Timestamp interface{}

    // Time to live. The type is interface{} with range: 0..255.
    Ttl interface{}

    // Transaction ID. The type is interface{} with range: 0..4294967295.
    TransactionId interface{}

    // Options affecting traceroute behavior.
    Options Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options
}

func (tracerouteInformation *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation) GetEntityData() *types.CommonEntityData {
    tracerouteInformation.EntityData.YFilter = tracerouteInformation.YFilter
    tracerouteInformation.EntityData.YangName = "traceroute-information"
    tracerouteInformation.EntityData.BundleName = "cisco_ios_xr"
    tracerouteInformation.EntityData.ParentYangName = "incomplete-traceroute"
    tracerouteInformation.EntityData.SegmentPath = "traceroute-information"
    tracerouteInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracerouteInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracerouteInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracerouteInformation.EntityData.Children = make(map[string]types.YChild)
    tracerouteInformation.EntityData.Children["options"] = types.YChild{"Options", &tracerouteInformation.Options}
    tracerouteInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    tracerouteInformation.EntityData.Leafs["domain"] = types.YLeaf{"Domain", tracerouteInformation.Domain}
    tracerouteInformation.EntityData.Leafs["service"] = types.YLeaf{"Service", tracerouteInformation.Service}
    tracerouteInformation.EntityData.Leafs["level"] = types.YLeaf{"Level", tracerouteInformation.Level}
    tracerouteInformation.EntityData.Leafs["source-mep-id"] = types.YLeaf{"SourceMepId", tracerouteInformation.SourceMepId}
    tracerouteInformation.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", tracerouteInformation.SourceInterface}
    tracerouteInformation.EntityData.Leafs["source-mac-address"] = types.YLeaf{"SourceMacAddress", tracerouteInformation.SourceMacAddress}
    tracerouteInformation.EntityData.Leafs["target-mac-address"] = types.YLeaf{"TargetMacAddress", tracerouteInformation.TargetMacAddress}
    tracerouteInformation.EntityData.Leafs["directed-mac-address"] = types.YLeaf{"DirectedMacAddress", tracerouteInformation.DirectedMacAddress}
    tracerouteInformation.EntityData.Leafs["target-mep-id"] = types.YLeaf{"TargetMepId", tracerouteInformation.TargetMepId}
    tracerouteInformation.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", tracerouteInformation.Timestamp}
    tracerouteInformation.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", tracerouteInformation.Ttl}
    tracerouteInformation.EntityData.Leafs["transaction-id"] = types.YLeaf{"TransactionId", tracerouteInformation.TransactionId}
    return &(tracerouteInformation.EntityData)
}

// Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options
// Options affecting traceroute behavior
type Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mode. The type is CfmPmLtMode.
    Mode interface{}

    // Options for a basic IEEE 802.1ag Linktrace.
    BasicOptions Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions

    // Options for an Exploratory Linktrace.
    ExploratoryOptions Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions
}

func (options *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options) GetEntityData() *types.CommonEntityData {
    options.EntityData.YFilter = options.YFilter
    options.EntityData.YangName = "options"
    options.EntityData.BundleName = "cisco_ios_xr"
    options.EntityData.ParentYangName = "traceroute-information"
    options.EntityData.SegmentPath = "options"
    options.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    options.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    options.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    options.EntityData.Children = make(map[string]types.YChild)
    options.EntityData.Children["basic-options"] = types.YChild{"BasicOptions", &options.BasicOptions}
    options.EntityData.Children["exploratory-options"] = types.YChild{"ExploratoryOptions", &options.ExploratoryOptions}
    options.EntityData.Leafs = make(map[string]types.YLeaf)
    options.EntityData.Leafs["mode"] = types.YLeaf{"Mode", options.Mode}
    return &(options.EntityData)
}

// Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions
// Options for a basic IEEE 802.1ag Linktrace
type Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traceroute was initiated automatically. The type is bool.
    IsAuto interface{}

    // Only use the Filtering Database for forwarding lookups. The type is bool.
    FdbOnly interface{}
}

func (basicOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_BasicOptions) GetEntityData() *types.CommonEntityData {
    basicOptions.EntityData.YFilter = basicOptions.YFilter
    basicOptions.EntityData.YangName = "basic-options"
    basicOptions.EntityData.BundleName = "cisco_ios_xr"
    basicOptions.EntityData.ParentYangName = "options"
    basicOptions.EntityData.SegmentPath = "basic-options"
    basicOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicOptions.EntityData.Children = make(map[string]types.YChild)
    basicOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    basicOptions.EntityData.Leafs["is-auto"] = types.YLeaf{"IsAuto", basicOptions.IsAuto}
    basicOptions.EntityData.Leafs["fdb-only"] = types.YLeaf{"FdbOnly", basicOptions.FdbOnly}
    return &(basicOptions.EntityData)
}

// Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions
// Options for an Exploratory Linktrace
type Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay model for delay calculations. The type is CfmPmEltDelayModel.
    DelayModel interface{}

    // Constant Factor for delay calculations. The type is interface{} with range:
    // 0..4294967295.
    DelayConstantFactor interface{}

    // Reply Filtering mode used by responders. The type is CfmPmElmReplyFilter.
    ReplyFilter interface{}
}

func (exploratoryOptions *Cfm_Global_IncompleteTraceroutes_IncompleteTraceroute_TracerouteInformation_Options_ExploratoryOptions) GetEntityData() *types.CommonEntityData {
    exploratoryOptions.EntityData.YFilter = exploratoryOptions.YFilter
    exploratoryOptions.EntityData.YangName = "exploratory-options"
    exploratoryOptions.EntityData.BundleName = "cisco_ios_xr"
    exploratoryOptions.EntityData.ParentYangName = "options"
    exploratoryOptions.EntityData.SegmentPath = "exploratory-options"
    exploratoryOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exploratoryOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exploratoryOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exploratoryOptions.EntityData.Children = make(map[string]types.YChild)
    exploratoryOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    exploratoryOptions.EntityData.Leafs["delay-model"] = types.YLeaf{"DelayModel", exploratoryOptions.DelayModel}
    exploratoryOptions.EntityData.Leafs["delay-constant-factor"] = types.YLeaf{"DelayConstantFactor", exploratoryOptions.DelayConstantFactor}
    exploratoryOptions.EntityData.Leafs["reply-filter"] = types.YLeaf{"ReplyFilter", exploratoryOptions.ReplyFilter}
    return &(exploratoryOptions.EntityData)
}

// Cfm_Global_MaintenancePoints
// Maintenance Points table
type Cfm_Global_MaintenancePoints struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular Maintenance Point. The type is slice of
    // Cfm_Global_MaintenancePoints_MaintenancePoint.
    MaintenancePoint []Cfm_Global_MaintenancePoints_MaintenancePoint
}

func (maintenancePoints *Cfm_Global_MaintenancePoints) GetEntityData() *types.CommonEntityData {
    maintenancePoints.EntityData.YFilter = maintenancePoints.YFilter
    maintenancePoints.EntityData.YangName = "maintenance-points"
    maintenancePoints.EntityData.BundleName = "cisco_ios_xr"
    maintenancePoints.EntityData.ParentYangName = "global"
    maintenancePoints.EntityData.SegmentPath = "maintenance-points"
    maintenancePoints.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maintenancePoints.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maintenancePoints.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maintenancePoints.EntityData.Children = make(map[string]types.YChild)
    maintenancePoints.EntityData.Children["maintenance-point"] = types.YChild{"MaintenancePoint", nil}
    for i := range maintenancePoints.MaintenancePoint {
        maintenancePoints.EntityData.Children[types.GetSegmentPath(&maintenancePoints.MaintenancePoint[i])] = types.YChild{"MaintenancePoint", &maintenancePoints.MaintenancePoint[i]}
    }
    maintenancePoints.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(maintenancePoints.EntityData)
}

// Cfm_Global_MaintenancePoints_MaintenancePoint
// Information about a particular Maintenance
// Point
type Cfm_Global_MaintenancePoints_MaintenancePoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // MEP error flag. The type is bool.
    MepHasError interface{}

    // MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Maintenance Point.
    MaintenancePoint Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint_
}

func (maintenancePoint *Cfm_Global_MaintenancePoints_MaintenancePoint) GetEntityData() *types.CommonEntityData {
    maintenancePoint.EntityData.YFilter = maintenancePoint.YFilter
    maintenancePoint.EntityData.YangName = "maintenance-point"
    maintenancePoint.EntityData.BundleName = "cisco_ios_xr"
    maintenancePoint.EntityData.ParentYangName = "maintenance-points"
    maintenancePoint.EntityData.SegmentPath = "maintenance-point" + "[domain='" + fmt.Sprintf("%v", maintenancePoint.Domain) + "']" + "[service='" + fmt.Sprintf("%v", maintenancePoint.Service) + "']" + "[interface='" + fmt.Sprintf("%v", maintenancePoint.Interface_) + "']"
    maintenancePoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maintenancePoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maintenancePoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maintenancePoint.EntityData.Children = make(map[string]types.YChild)
    maintenancePoint.EntityData.Children["maintenance-point"] = types.YChild{"MaintenancePoint", &maintenancePoint.MaintenancePoint}
    maintenancePoint.EntityData.Leafs = make(map[string]types.YLeaf)
    maintenancePoint.EntityData.Leafs["domain"] = types.YLeaf{"Domain", maintenancePoint.Domain}
    maintenancePoint.EntityData.Leafs["service"] = types.YLeaf{"Service", maintenancePoint.Service}
    maintenancePoint.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", maintenancePoint.Interface_}
    maintenancePoint.EntityData.Leafs["mep-has-error"] = types.YLeaf{"MepHasError", maintenancePoint.MepHasError}
    maintenancePoint.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", maintenancePoint.MacAddress}
    return &(maintenancePoint.EntityData)
}

// Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint_
// Maintenance Point
type Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain name. The type is string.
    DomainName interface{}

    // Domain level. The type is CfmBagMdLevel.
    Level interface{}

    // Service name. The type is string.
    ServiceName interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Type of Maintenance Point. The type is CfmMaMpVariety.
    MaintenancePointType interface{}

    // MEP ID. The type is interface{} with range: 0..65535.
    MepId interface{}
}

func (maintenancePoint_ *Cfm_Global_MaintenancePoints_MaintenancePoint_MaintenancePoint_) GetEntityData() *types.CommonEntityData {
    maintenancePoint_.EntityData.YFilter = maintenancePoint_.YFilter
    maintenancePoint_.EntityData.YangName = "maintenance-point"
    maintenancePoint_.EntityData.BundleName = "cisco_ios_xr"
    maintenancePoint_.EntityData.ParentYangName = "maintenance-point"
    maintenancePoint_.EntityData.SegmentPath = "maintenance-point"
    maintenancePoint_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maintenancePoint_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maintenancePoint_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maintenancePoint_.EntityData.Children = make(map[string]types.YChild)
    maintenancePoint_.EntityData.Leafs = make(map[string]types.YLeaf)
    maintenancePoint_.EntityData.Leafs["domain-name"] = types.YLeaf{"DomainName", maintenancePoint_.DomainName}
    maintenancePoint_.EntityData.Leafs["level"] = types.YLeaf{"Level", maintenancePoint_.Level}
    maintenancePoint_.EntityData.Leafs["service-name"] = types.YLeaf{"ServiceName", maintenancePoint_.ServiceName}
    maintenancePoint_.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", maintenancePoint_.Interface_}
    maintenancePoint_.EntityData.Leafs["maintenance-point-type"] = types.YLeaf{"MaintenancePointType", maintenancePoint_.MaintenancePointType}
    maintenancePoint_.EntityData.Leafs["mep-id"] = types.YLeaf{"MepId", maintenancePoint_.MepId}
    return &(maintenancePoint_.EntityData)
}

// Cfm_Global_GlobalConfigurationErrors
// Global configuration errors table
type Cfm_Global_GlobalConfigurationErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular configuration error. The type is slice of
    // Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError.
    GlobalConfigurationError []Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError
}

func (globalConfigurationErrors *Cfm_Global_GlobalConfigurationErrors) GetEntityData() *types.CommonEntityData {
    globalConfigurationErrors.EntityData.YFilter = globalConfigurationErrors.YFilter
    globalConfigurationErrors.EntityData.YangName = "global-configuration-errors"
    globalConfigurationErrors.EntityData.BundleName = "cisco_ios_xr"
    globalConfigurationErrors.EntityData.ParentYangName = "global"
    globalConfigurationErrors.EntityData.SegmentPath = "global-configuration-errors"
    globalConfigurationErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalConfigurationErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalConfigurationErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalConfigurationErrors.EntityData.Children = make(map[string]types.YChild)
    globalConfigurationErrors.EntityData.Children["global-configuration-error"] = types.YChild{"GlobalConfigurationError", nil}
    for i := range globalConfigurationErrors.GlobalConfigurationError {
        globalConfigurationErrors.EntityData.Children[types.GetSegmentPath(&globalConfigurationErrors.GlobalConfigurationError[i])] = types.YChild{"GlobalConfigurationError", &globalConfigurationErrors.GlobalConfigurationError[i]}
    }
    globalConfigurationErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(globalConfigurationErrors.EntityData)
}

// Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError
// Information about a particular configuration
// error
type Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Level. The type is CfmBagMdLevel.
    Level interface{}

    // Service name. The type is string.
    ServiceName interface{}

    // The BD/XC is configured globally. The type is bool.
    BridgeDomainIsConfigured interface{}

    // The BD/XC could not be downloaded to L2FIB. The type is bool.
    L2FibDownloadError interface{}

    // BD/XC ID, or Service name if the Service is 'down-only'.
    BridgeDomainId Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId
}

func (globalConfigurationError *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError) GetEntityData() *types.CommonEntityData {
    globalConfigurationError.EntityData.YFilter = globalConfigurationError.YFilter
    globalConfigurationError.EntityData.YangName = "global-configuration-error"
    globalConfigurationError.EntityData.BundleName = "cisco_ios_xr"
    globalConfigurationError.EntityData.ParentYangName = "global-configuration-errors"
    globalConfigurationError.EntityData.SegmentPath = "global-configuration-error" + "[domain='" + fmt.Sprintf("%v", globalConfigurationError.Domain) + "']" + "[service='" + fmt.Sprintf("%v", globalConfigurationError.Service) + "']"
    globalConfigurationError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalConfigurationError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalConfigurationError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalConfigurationError.EntityData.Children = make(map[string]types.YChild)
    globalConfigurationError.EntityData.Children["bridge-domain-id"] = types.YChild{"BridgeDomainId", &globalConfigurationError.BridgeDomainId}
    globalConfigurationError.EntityData.Leafs = make(map[string]types.YLeaf)
    globalConfigurationError.EntityData.Leafs["domain"] = types.YLeaf{"Domain", globalConfigurationError.Domain}
    globalConfigurationError.EntityData.Leafs["service"] = types.YLeaf{"Service", globalConfigurationError.Service}
    globalConfigurationError.EntityData.Leafs["domain-name"] = types.YLeaf{"DomainName", globalConfigurationError.DomainName}
    globalConfigurationError.EntityData.Leafs["level"] = types.YLeaf{"Level", globalConfigurationError.Level}
    globalConfigurationError.EntityData.Leafs["service-name"] = types.YLeaf{"ServiceName", globalConfigurationError.ServiceName}
    globalConfigurationError.EntityData.Leafs["bridge-domain-is-configured"] = types.YLeaf{"BridgeDomainIsConfigured", globalConfigurationError.BridgeDomainIsConfigured}
    globalConfigurationError.EntityData.Leafs["l2-fib-download-error"] = types.YLeaf{"L2FibDownloadError", globalConfigurationError.L2FibDownloadError}
    return &(globalConfigurationError.EntityData)
}

// Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId
// BD/XC ID, or Service name if the Service is
// 'down-only'
type Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge domain identifier format. The type is CfmBagBdidFmt.
    BridgeDomainIdFormat interface{}

    // Name of the Bridge/XConnect Group. The type is string.
    Group interface{}

    // Name of the Bridge Domain/XConnect. The type is string.
    Name interface{}

    // Local Customer Edge Identifier (CE-ID). The type is interface{} with range:
    // 0..65535.
    CeId interface{}

    // Remote Customer Edge Identifier (CE-ID). The type is interface{} with
    // range: 0..65535.
    RemoteCeId interface{}

    // EVPN ID for VLAN-aware flexible cross-connects. The type is interface{}
    // with range: 0..4294967295.
    Evi interface{}
}

func (bridgeDomainId *Cfm_Global_GlobalConfigurationErrors_GlobalConfigurationError_BridgeDomainId) GetEntityData() *types.CommonEntityData {
    bridgeDomainId.EntityData.YFilter = bridgeDomainId.YFilter
    bridgeDomainId.EntityData.YangName = "bridge-domain-id"
    bridgeDomainId.EntityData.BundleName = "cisco_ios_xr"
    bridgeDomainId.EntityData.ParentYangName = "global-configuration-error"
    bridgeDomainId.EntityData.SegmentPath = "bridge-domain-id"
    bridgeDomainId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridgeDomainId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridgeDomainId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridgeDomainId.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainId.EntityData.Leafs = make(map[string]types.YLeaf)
    bridgeDomainId.EntityData.Leafs["bridge-domain-id-format"] = types.YLeaf{"BridgeDomainIdFormat", bridgeDomainId.BridgeDomainIdFormat}
    bridgeDomainId.EntityData.Leafs["group"] = types.YLeaf{"Group", bridgeDomainId.Group}
    bridgeDomainId.EntityData.Leafs["name"] = types.YLeaf{"Name", bridgeDomainId.Name}
    bridgeDomainId.EntityData.Leafs["ce-id"] = types.YLeaf{"CeId", bridgeDomainId.CeId}
    bridgeDomainId.EntityData.Leafs["remote-ce-id"] = types.YLeaf{"RemoteCeId", bridgeDomainId.RemoteCeId}
    bridgeDomainId.EntityData.Leafs["evi"] = types.YLeaf{"Evi", bridgeDomainId.Evi}
    return &(bridgeDomainId.EntityData)
}

// Cfm_Global_MepConfigurationErrors
// MEP configuration errors table
type Cfm_Global_MepConfigurationErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular configuration error. The type is slice of
    // Cfm_Global_MepConfigurationErrors_MepConfigurationError.
    MepConfigurationError []Cfm_Global_MepConfigurationErrors_MepConfigurationError
}

func (mepConfigurationErrors *Cfm_Global_MepConfigurationErrors) GetEntityData() *types.CommonEntityData {
    mepConfigurationErrors.EntityData.YFilter = mepConfigurationErrors.YFilter
    mepConfigurationErrors.EntityData.YangName = "mep-configuration-errors"
    mepConfigurationErrors.EntityData.BundleName = "cisco_ios_xr"
    mepConfigurationErrors.EntityData.ParentYangName = "global"
    mepConfigurationErrors.EntityData.SegmentPath = "mep-configuration-errors"
    mepConfigurationErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mepConfigurationErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mepConfigurationErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mepConfigurationErrors.EntityData.Children = make(map[string]types.YChild)
    mepConfigurationErrors.EntityData.Children["mep-configuration-error"] = types.YChild{"MepConfigurationError", nil}
    for i := range mepConfigurationErrors.MepConfigurationError {
        mepConfigurationErrors.EntityData.Children[types.GetSegmentPath(&mepConfigurationErrors.MepConfigurationError[i])] = types.YChild{"MepConfigurationError", &mepConfigurationErrors.MepConfigurationError[i]}
    }
    mepConfigurationErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mepConfigurationErrors.EntityData)
}

// Cfm_Global_MepConfigurationErrors_MepConfigurationError
// Information about a particular configuration
// error
type Cfm_Global_MepConfigurationErrors_MepConfigurationError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Interval between CCMs sent on this MEP. The type is CfmBagCcmInterval.
    CcmInterval interface{}

    // The MEP's Domain is not configured. The type is bool.
    NoDomain interface{}

    // The MEP's Service is not configured. The type is bool.
    NoService interface{}

    // The MEP's EFP is not in the Service's Bridge Domain. The type is bool.
    BridgeDomainMismatch interface{}

    // Another MEP facing in the same direction is at the same Maintenance Level.
    // The type is bool.
    LevelConflict interface{}

    // CCM Interval is less than minimum interval supported by hardware. The type
    // is bool.
    CcmIntervalNotSupported interface{}

    // Offload resource limits have been exceeded. The type is bool.
    OffloadOutOfResources interface{}

    // Multiple offloaded MEPs on the same interface are not supported. The type
    // is bool.
    OffloadMultipleLocalMep interface{}

    // The MEP should be offloaded but crosscheck has not been configured. The
    // type is bool.
    OffloadNoCrossCheck interface{}

    // The MEP should be offloaded but multiple crosscheck MEPs have been
    // configured, and this is not supported. The type is bool.
    OffloadMultiplePeerMeps interface{}

    // The MEP direction does not support offload. The type is bool.
    OffloadMepDirectionNotSupported interface{}

    // AIS is configured on the same interface as the down MEP. The type is bool.
    AisConfigured interface{}

    // The MEP is configured in a domain at level 0, on a bundle interface or
    // sub-interface.  This is not supported. The type is bool.
    BundleLevel0 interface{}

    // A BD/XC specified in the MEG config, but it does not exist globally. The
    // type is bool.
    BridgeDomainNotInBdInfra interface{}

    // The configured MAID format is not supported for hardware offload. The type
    // is bool.
    MaidFormatNotSupported interface{}

    // The platform returned a fatal error when passed the offload session. The
    // type is bool.
    FatalOffloadError interface{}

    // A satellite limitation is preventing MEP being offloaded to satellite. The
    // type is bool.
    SatelliteLimitation interface{}

    // In-progress Ethernet SLA loopback operations are disabled due to satellite
    // having loopback responder-only capabilities. The type is bool.
    SlaLoopbackOperationsDisabled interface{}

    // In-progress Ethernet SLA synthetic loss measurement operations are disabled
    // due to satellite having synthetic loss measurement responder-only
    // capabilities. The type is bool.
    SlaSyntheticLossOperationsDisabled interface{}

    // In-progress Ethernet SLA delay measurement operations are disabled due to
    // satellite having delay measurement responder-only capabilities. The type is
    // bool.
    SlaDelayMeasurementOperationsDisabled interface{}

    // The EFP doesn't have a valid MAC address yet. This will also get set if the
    // MAC address we have is a multicast address. The type is bool.
    NoValidMacAddress interface{}

    // We haven't yet been able to look up the interface type to find whether the
    // interface is a bundle. The type is bool.
    NoInterfaceType interface{}

    // The EFP has been deleted from IM. The type is bool.
    NotInIm interface{}

    // The EFP is a bundle and the mLACP mode is not yet known. The type is bool.
    NoMlacp interface{}

    // Error string returned from satellite. The type is string.
    SatelliteErrorString interface{}

    // ID of the satellite. The type is interface{} with range: 0..65535.
    SatelliteId interface{}

    // The MEP that has errors.
    Mep Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep

    // BD/XC ID for the MEP's Service, or Service name if the Service is
    // 'down-only'.
    ServiceBridgeDomain Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain

    // ID of the BD/XC that the MEP's EFP is in, if any.
    InterfaceBridgeDomain Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain

    // Satellite Capabilities.
    SatelliteCapabilities Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities
}

func (mepConfigurationError *Cfm_Global_MepConfigurationErrors_MepConfigurationError) GetEntityData() *types.CommonEntityData {
    mepConfigurationError.EntityData.YFilter = mepConfigurationError.YFilter
    mepConfigurationError.EntityData.YangName = "mep-configuration-error"
    mepConfigurationError.EntityData.BundleName = "cisco_ios_xr"
    mepConfigurationError.EntityData.ParentYangName = "mep-configuration-errors"
    mepConfigurationError.EntityData.SegmentPath = "mep-configuration-error" + "[domain='" + fmt.Sprintf("%v", mepConfigurationError.Domain) + "']" + "[service='" + fmt.Sprintf("%v", mepConfigurationError.Service) + "']" + "[interface='" + fmt.Sprintf("%v", mepConfigurationError.Interface_) + "']"
    mepConfigurationError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mepConfigurationError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mepConfigurationError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mepConfigurationError.EntityData.Children = make(map[string]types.YChild)
    mepConfigurationError.EntityData.Children["mep"] = types.YChild{"Mep", &mepConfigurationError.Mep}
    mepConfigurationError.EntityData.Children["service-bridge-domain"] = types.YChild{"ServiceBridgeDomain", &mepConfigurationError.ServiceBridgeDomain}
    mepConfigurationError.EntityData.Children["interface-bridge-domain"] = types.YChild{"InterfaceBridgeDomain", &mepConfigurationError.InterfaceBridgeDomain}
    mepConfigurationError.EntityData.Children["satellite-capabilities"] = types.YChild{"SatelliteCapabilities", &mepConfigurationError.SatelliteCapabilities}
    mepConfigurationError.EntityData.Leafs = make(map[string]types.YLeaf)
    mepConfigurationError.EntityData.Leafs["domain"] = types.YLeaf{"Domain", mepConfigurationError.Domain}
    mepConfigurationError.EntityData.Leafs["service"] = types.YLeaf{"Service", mepConfigurationError.Service}
    mepConfigurationError.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", mepConfigurationError.Interface_}
    mepConfigurationError.EntityData.Leafs["ccm-interval"] = types.YLeaf{"CcmInterval", mepConfigurationError.CcmInterval}
    mepConfigurationError.EntityData.Leafs["no-domain"] = types.YLeaf{"NoDomain", mepConfigurationError.NoDomain}
    mepConfigurationError.EntityData.Leafs["no-service"] = types.YLeaf{"NoService", mepConfigurationError.NoService}
    mepConfigurationError.EntityData.Leafs["bridge-domain-mismatch"] = types.YLeaf{"BridgeDomainMismatch", mepConfigurationError.BridgeDomainMismatch}
    mepConfigurationError.EntityData.Leafs["level-conflict"] = types.YLeaf{"LevelConflict", mepConfigurationError.LevelConflict}
    mepConfigurationError.EntityData.Leafs["ccm-interval-not-supported"] = types.YLeaf{"CcmIntervalNotSupported", mepConfigurationError.CcmIntervalNotSupported}
    mepConfigurationError.EntityData.Leafs["offload-out-of-resources"] = types.YLeaf{"OffloadOutOfResources", mepConfigurationError.OffloadOutOfResources}
    mepConfigurationError.EntityData.Leafs["offload-multiple-local-mep"] = types.YLeaf{"OffloadMultipleLocalMep", mepConfigurationError.OffloadMultipleLocalMep}
    mepConfigurationError.EntityData.Leafs["offload-no-cross-check"] = types.YLeaf{"OffloadNoCrossCheck", mepConfigurationError.OffloadNoCrossCheck}
    mepConfigurationError.EntityData.Leafs["offload-multiple-peer-meps"] = types.YLeaf{"OffloadMultiplePeerMeps", mepConfigurationError.OffloadMultiplePeerMeps}
    mepConfigurationError.EntityData.Leafs["offload-mep-direction-not-supported"] = types.YLeaf{"OffloadMepDirectionNotSupported", mepConfigurationError.OffloadMepDirectionNotSupported}
    mepConfigurationError.EntityData.Leafs["ais-configured"] = types.YLeaf{"AisConfigured", mepConfigurationError.AisConfigured}
    mepConfigurationError.EntityData.Leafs["bundle-level0"] = types.YLeaf{"BundleLevel0", mepConfigurationError.BundleLevel0}
    mepConfigurationError.EntityData.Leafs["bridge-domain-not-in-bd-infra"] = types.YLeaf{"BridgeDomainNotInBdInfra", mepConfigurationError.BridgeDomainNotInBdInfra}
    mepConfigurationError.EntityData.Leafs["maid-format-not-supported"] = types.YLeaf{"MaidFormatNotSupported", mepConfigurationError.MaidFormatNotSupported}
    mepConfigurationError.EntityData.Leafs["fatal-offload-error"] = types.YLeaf{"FatalOffloadError", mepConfigurationError.FatalOffloadError}
    mepConfigurationError.EntityData.Leafs["satellite-limitation"] = types.YLeaf{"SatelliteLimitation", mepConfigurationError.SatelliteLimitation}
    mepConfigurationError.EntityData.Leafs["sla-loopback-operations-disabled"] = types.YLeaf{"SlaLoopbackOperationsDisabled", mepConfigurationError.SlaLoopbackOperationsDisabled}
    mepConfigurationError.EntityData.Leafs["sla-synthetic-loss-operations-disabled"] = types.YLeaf{"SlaSyntheticLossOperationsDisabled", mepConfigurationError.SlaSyntheticLossOperationsDisabled}
    mepConfigurationError.EntityData.Leafs["sla-delay-measurement-operations-disabled"] = types.YLeaf{"SlaDelayMeasurementOperationsDisabled", mepConfigurationError.SlaDelayMeasurementOperationsDisabled}
    mepConfigurationError.EntityData.Leafs["no-valid-mac-address"] = types.YLeaf{"NoValidMacAddress", mepConfigurationError.NoValidMacAddress}
    mepConfigurationError.EntityData.Leafs["no-interface-type"] = types.YLeaf{"NoInterfaceType", mepConfigurationError.NoInterfaceType}
    mepConfigurationError.EntityData.Leafs["not-in-im"] = types.YLeaf{"NotInIm", mepConfigurationError.NotInIm}
    mepConfigurationError.EntityData.Leafs["no-mlacp"] = types.YLeaf{"NoMlacp", mepConfigurationError.NoMlacp}
    mepConfigurationError.EntityData.Leafs["satellite-error-string"] = types.YLeaf{"SatelliteErrorString", mepConfigurationError.SatelliteErrorString}
    mepConfigurationError.EntityData.Leafs["satellite-id"] = types.YLeaf{"SatelliteId", mepConfigurationError.SatelliteId}
    return &(mepConfigurationError.EntityData)
}

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep
// The MEP that has errors
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain name. The type is string.
    DomainName interface{}

    // Domain level. The type is CfmBagMdLevel.
    Level interface{}

    // Service name. The type is string.
    ServiceName interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Type of Maintenance Point. The type is CfmMaMpVariety.
    MaintenancePointType interface{}

    // MEP ID. The type is interface{} with range: 0..65535.
    MepId interface{}
}

func (mep *Cfm_Global_MepConfigurationErrors_MepConfigurationError_Mep) GetEntityData() *types.CommonEntityData {
    mep.EntityData.YFilter = mep.YFilter
    mep.EntityData.YangName = "mep"
    mep.EntityData.BundleName = "cisco_ios_xr"
    mep.EntityData.ParentYangName = "mep-configuration-error"
    mep.EntityData.SegmentPath = "mep"
    mep.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mep.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mep.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mep.EntityData.Children = make(map[string]types.YChild)
    mep.EntityData.Leafs = make(map[string]types.YLeaf)
    mep.EntityData.Leafs["domain-name"] = types.YLeaf{"DomainName", mep.DomainName}
    mep.EntityData.Leafs["level"] = types.YLeaf{"Level", mep.Level}
    mep.EntityData.Leafs["service-name"] = types.YLeaf{"ServiceName", mep.ServiceName}
    mep.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", mep.Interface_}
    mep.EntityData.Leafs["maintenance-point-type"] = types.YLeaf{"MaintenancePointType", mep.MaintenancePointType}
    mep.EntityData.Leafs["mep-id"] = types.YLeaf{"MepId", mep.MepId}
    return &(mep.EntityData)
}

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain
// BD/XC ID for the MEP's Service, or Service name
// if the Service is 'down-only'
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge domain identifier format. The type is CfmBagBdidFmt.
    BridgeDomainIdFormat interface{}

    // Name of the Bridge/XConnect Group. The type is string.
    Group interface{}

    // Name of the Bridge Domain/XConnect. The type is string.
    Name interface{}

    // Local Customer Edge Identifier (CE-ID). The type is interface{} with range:
    // 0..65535.
    CeId interface{}

    // Remote Customer Edge Identifier (CE-ID). The type is interface{} with
    // range: 0..65535.
    RemoteCeId interface{}

    // EVPN ID for VLAN-aware flexible cross-connects. The type is interface{}
    // with range: 0..4294967295.
    Evi interface{}
}

func (serviceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_ServiceBridgeDomain) GetEntityData() *types.CommonEntityData {
    serviceBridgeDomain.EntityData.YFilter = serviceBridgeDomain.YFilter
    serviceBridgeDomain.EntityData.YangName = "service-bridge-domain"
    serviceBridgeDomain.EntityData.BundleName = "cisco_ios_xr"
    serviceBridgeDomain.EntityData.ParentYangName = "mep-configuration-error"
    serviceBridgeDomain.EntityData.SegmentPath = "service-bridge-domain"
    serviceBridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceBridgeDomain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceBridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceBridgeDomain.EntityData.Children = make(map[string]types.YChild)
    serviceBridgeDomain.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceBridgeDomain.EntityData.Leafs["bridge-domain-id-format"] = types.YLeaf{"BridgeDomainIdFormat", serviceBridgeDomain.BridgeDomainIdFormat}
    serviceBridgeDomain.EntityData.Leafs["group"] = types.YLeaf{"Group", serviceBridgeDomain.Group}
    serviceBridgeDomain.EntityData.Leafs["name"] = types.YLeaf{"Name", serviceBridgeDomain.Name}
    serviceBridgeDomain.EntityData.Leafs["ce-id"] = types.YLeaf{"CeId", serviceBridgeDomain.CeId}
    serviceBridgeDomain.EntityData.Leafs["remote-ce-id"] = types.YLeaf{"RemoteCeId", serviceBridgeDomain.RemoteCeId}
    serviceBridgeDomain.EntityData.Leafs["evi"] = types.YLeaf{"Evi", serviceBridgeDomain.Evi}
    return &(serviceBridgeDomain.EntityData)
}

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain
// ID of the BD/XC that the MEP's EFP is in, if any
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge domain identifier format. The type is CfmBagBdidFmt.
    BridgeDomainIdFormat interface{}

    // Name of the Bridge/XConnect Group. The type is string.
    Group interface{}

    // Name of the Bridge Domain/XConnect. The type is string.
    Name interface{}

    // Local Customer Edge Identifier (CE-ID). The type is interface{} with range:
    // 0..65535.
    CeId interface{}

    // Remote Customer Edge Identifier (CE-ID). The type is interface{} with
    // range: 0..65535.
    RemoteCeId interface{}

    // EVPN ID for VLAN-aware flexible cross-connects. The type is interface{}
    // with range: 0..4294967295.
    Evi interface{}
}

func (interfaceBridgeDomain *Cfm_Global_MepConfigurationErrors_MepConfigurationError_InterfaceBridgeDomain) GetEntityData() *types.CommonEntityData {
    interfaceBridgeDomain.EntityData.YFilter = interfaceBridgeDomain.YFilter
    interfaceBridgeDomain.EntityData.YangName = "interface-bridge-domain"
    interfaceBridgeDomain.EntityData.BundleName = "cisco_ios_xr"
    interfaceBridgeDomain.EntityData.ParentYangName = "mep-configuration-error"
    interfaceBridgeDomain.EntityData.SegmentPath = "interface-bridge-domain"
    interfaceBridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceBridgeDomain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceBridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceBridgeDomain.EntityData.Children = make(map[string]types.YChild)
    interfaceBridgeDomain.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceBridgeDomain.EntityData.Leafs["bridge-domain-id-format"] = types.YLeaf{"BridgeDomainIdFormat", interfaceBridgeDomain.BridgeDomainIdFormat}
    interfaceBridgeDomain.EntityData.Leafs["group"] = types.YLeaf{"Group", interfaceBridgeDomain.Group}
    interfaceBridgeDomain.EntityData.Leafs["name"] = types.YLeaf{"Name", interfaceBridgeDomain.Name}
    interfaceBridgeDomain.EntityData.Leafs["ce-id"] = types.YLeaf{"CeId", interfaceBridgeDomain.CeId}
    interfaceBridgeDomain.EntityData.Leafs["remote-ce-id"] = types.YLeaf{"RemoteCeId", interfaceBridgeDomain.RemoteCeId}
    interfaceBridgeDomain.EntityData.Leafs["evi"] = types.YLeaf{"Evi", interfaceBridgeDomain.Evi}
    return &(interfaceBridgeDomain.EntityData)
}

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities
// Satellite Capabilities
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Loopback.
    Loopback Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback

    // Delay Measurement.
    DelayMeasurement Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement

    // Synthetic Loss Measurement.
    SyntheticLossMeasurement Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement
}

func (satelliteCapabilities *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities) GetEntityData() *types.CommonEntityData {
    satelliteCapabilities.EntityData.YFilter = satelliteCapabilities.YFilter
    satelliteCapabilities.EntityData.YangName = "satellite-capabilities"
    satelliteCapabilities.EntityData.BundleName = "cisco_ios_xr"
    satelliteCapabilities.EntityData.ParentYangName = "mep-configuration-error"
    satelliteCapabilities.EntityData.SegmentPath = "satellite-capabilities"
    satelliteCapabilities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    satelliteCapabilities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    satelliteCapabilities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    satelliteCapabilities.EntityData.Children = make(map[string]types.YChild)
    satelliteCapabilities.EntityData.Children["loopback"] = types.YChild{"Loopback", &satelliteCapabilities.Loopback}
    satelliteCapabilities.EntityData.Children["delay-measurement"] = types.YChild{"DelayMeasurement", &satelliteCapabilities.DelayMeasurement}
    satelliteCapabilities.EntityData.Children["synthetic-loss-measurement"] = types.YChild{"SyntheticLossMeasurement", &satelliteCapabilities.SyntheticLossMeasurement}
    satelliteCapabilities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(satelliteCapabilities.EntityData)
}

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback
// Loopback
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responder. The type is bool.
    Responder interface{}

    // Controller. The type is bool.
    Controller interface{}
}

func (loopback *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_Loopback) GetEntityData() *types.CommonEntityData {
    loopback.EntityData.YFilter = loopback.YFilter
    loopback.EntityData.YangName = "loopback"
    loopback.EntityData.BundleName = "cisco_ios_xr"
    loopback.EntityData.ParentYangName = "satellite-capabilities"
    loopback.EntityData.SegmentPath = "loopback"
    loopback.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loopback.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loopback.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loopback.EntityData.Children = make(map[string]types.YChild)
    loopback.EntityData.Leafs = make(map[string]types.YLeaf)
    loopback.EntityData.Leafs["responder"] = types.YLeaf{"Responder", loopback.Responder}
    loopback.EntityData.Leafs["controller"] = types.YLeaf{"Controller", loopback.Controller}
    return &(loopback.EntityData)
}

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement
// Delay Measurement
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responder. The type is bool.
    Responder interface{}

    // Controller. The type is bool.
    Controller interface{}
}

func (delayMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_DelayMeasurement) GetEntityData() *types.CommonEntityData {
    delayMeasurement.EntityData.YFilter = delayMeasurement.YFilter
    delayMeasurement.EntityData.YangName = "delay-measurement"
    delayMeasurement.EntityData.BundleName = "cisco_ios_xr"
    delayMeasurement.EntityData.ParentYangName = "satellite-capabilities"
    delayMeasurement.EntityData.SegmentPath = "delay-measurement"
    delayMeasurement.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayMeasurement.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayMeasurement.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayMeasurement.EntityData.Children = make(map[string]types.YChild)
    delayMeasurement.EntityData.Leafs = make(map[string]types.YLeaf)
    delayMeasurement.EntityData.Leafs["responder"] = types.YLeaf{"Responder", delayMeasurement.Responder}
    delayMeasurement.EntityData.Leafs["controller"] = types.YLeaf{"Controller", delayMeasurement.Controller}
    return &(delayMeasurement.EntityData)
}

// Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement
// Synthetic Loss Measurement
type Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responder. The type is bool.
    Responder interface{}

    // Controller. The type is bool.
    Controller interface{}
}

func (syntheticLossMeasurement *Cfm_Global_MepConfigurationErrors_MepConfigurationError_SatelliteCapabilities_SyntheticLossMeasurement) GetEntityData() *types.CommonEntityData {
    syntheticLossMeasurement.EntityData.YFilter = syntheticLossMeasurement.YFilter
    syntheticLossMeasurement.EntityData.YangName = "synthetic-loss-measurement"
    syntheticLossMeasurement.EntityData.BundleName = "cisco_ios_xr"
    syntheticLossMeasurement.EntityData.ParentYangName = "satellite-capabilities"
    syntheticLossMeasurement.EntityData.SegmentPath = "synthetic-loss-measurement"
    syntheticLossMeasurement.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syntheticLossMeasurement.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syntheticLossMeasurement.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syntheticLossMeasurement.EntityData.Children = make(map[string]types.YChild)
    syntheticLossMeasurement.EntityData.Leafs = make(map[string]types.YLeaf)
    syntheticLossMeasurement.EntityData.Leafs["responder"] = types.YLeaf{"Responder", syntheticLossMeasurement.Responder}
    syntheticLossMeasurement.EntityData.Leafs["controller"] = types.YLeaf{"Controller", syntheticLossMeasurement.Controller}
    return &(syntheticLossMeasurement.EntityData)
}

// Cfm_Global_TracerouteCaches
// Traceroute Cache table
type Cfm_Global_TracerouteCaches struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular traceroute operation. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache.
    TracerouteCache []Cfm_Global_TracerouteCaches_TracerouteCache
}

func (tracerouteCaches *Cfm_Global_TracerouteCaches) GetEntityData() *types.CommonEntityData {
    tracerouteCaches.EntityData.YFilter = tracerouteCaches.YFilter
    tracerouteCaches.EntityData.YangName = "traceroute-caches"
    tracerouteCaches.EntityData.BundleName = "cisco_ios_xr"
    tracerouteCaches.EntityData.ParentYangName = "global"
    tracerouteCaches.EntityData.SegmentPath = "traceroute-caches"
    tracerouteCaches.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracerouteCaches.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracerouteCaches.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracerouteCaches.EntityData.Children = make(map[string]types.YChild)
    tracerouteCaches.EntityData.Children["traceroute-cache"] = types.YChild{"TracerouteCache", nil}
    for i := range tracerouteCaches.TracerouteCache {
        tracerouteCaches.EntityData.Children[types.GetSegmentPath(&tracerouteCaches.TracerouteCache[i])] = types.YChild{"TracerouteCache", &tracerouteCaches.TracerouteCache[i]}
    }
    tracerouteCaches.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tracerouteCaches.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache
// Information about a particular traceroute
// operation
type Cfm_Global_TracerouteCaches_TracerouteCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. MEP ID. The type is interface{} with range:
    // 1..8191.
    MepId interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // This attribute is a key. Transaction ID. The type is interface{} with
    // range: -2147483648..2147483647.
    TransactionId interface{}

    // Count of ignored replies for this request. The type is interface{} with
    // range: 0..4294967295.
    RepliesDropped interface{}

    // Information about the traceroute operation.
    TracerouteInformation Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation

    // Received linktrace replies. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply.
    LinktraceReply []Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply

    // Received exploratory linktrace replies. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply.
    ExploratoryLinktraceReply []Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply
}

func (tracerouteCache *Cfm_Global_TracerouteCaches_TracerouteCache) GetEntityData() *types.CommonEntityData {
    tracerouteCache.EntityData.YFilter = tracerouteCache.YFilter
    tracerouteCache.EntityData.YangName = "traceroute-cache"
    tracerouteCache.EntityData.BundleName = "cisco_ios_xr"
    tracerouteCache.EntityData.ParentYangName = "traceroute-caches"
    tracerouteCache.EntityData.SegmentPath = "traceroute-cache" + "[domain='" + fmt.Sprintf("%v", tracerouteCache.Domain) + "']" + "[service='" + fmt.Sprintf("%v", tracerouteCache.Service) + "']" + "[mep-id='" + fmt.Sprintf("%v", tracerouteCache.MepId) + "']" + "[interface='" + fmt.Sprintf("%v", tracerouteCache.Interface_) + "']" + "[transaction-id='" + fmt.Sprintf("%v", tracerouteCache.TransactionId) + "']"
    tracerouteCache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracerouteCache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracerouteCache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracerouteCache.EntityData.Children = make(map[string]types.YChild)
    tracerouteCache.EntityData.Children["traceroute-information"] = types.YChild{"TracerouteInformation", &tracerouteCache.TracerouteInformation}
    tracerouteCache.EntityData.Children["linktrace-reply"] = types.YChild{"LinktraceReply", nil}
    for i := range tracerouteCache.LinktraceReply {
        tracerouteCache.EntityData.Children[types.GetSegmentPath(&tracerouteCache.LinktraceReply[i])] = types.YChild{"LinktraceReply", &tracerouteCache.LinktraceReply[i]}
    }
    tracerouteCache.EntityData.Children["exploratory-linktrace-reply"] = types.YChild{"ExploratoryLinktraceReply", nil}
    for i := range tracerouteCache.ExploratoryLinktraceReply {
        tracerouteCache.EntityData.Children[types.GetSegmentPath(&tracerouteCache.ExploratoryLinktraceReply[i])] = types.YChild{"ExploratoryLinktraceReply", &tracerouteCache.ExploratoryLinktraceReply[i]}
    }
    tracerouteCache.EntityData.Leafs = make(map[string]types.YLeaf)
    tracerouteCache.EntityData.Leafs["domain"] = types.YLeaf{"Domain", tracerouteCache.Domain}
    tracerouteCache.EntityData.Leafs["service"] = types.YLeaf{"Service", tracerouteCache.Service}
    tracerouteCache.EntityData.Leafs["mep-id"] = types.YLeaf{"MepId", tracerouteCache.MepId}
    tracerouteCache.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", tracerouteCache.Interface_}
    tracerouteCache.EntityData.Leafs["transaction-id"] = types.YLeaf{"TransactionId", tracerouteCache.TransactionId}
    tracerouteCache.EntityData.Leafs["replies-dropped"] = types.YLeaf{"RepliesDropped", tracerouteCache.RepliesDropped}
    return &(tracerouteCache.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation
// Information about the traceroute operation
type Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maintenance domain name. The type is string.
    Domain interface{}

    // Service name. The type is string.
    Service interface{}

    // Maintenance level. The type is CfmBagMdLevel.
    Level interface{}

    // Source MEP ID. The type is interface{} with range: 0..65535.
    SourceMepId interface{}

    // Source interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SourceInterface interface{}

    // Source MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SourceMacAddress interface{}

    // Target MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    TargetMacAddress interface{}

    // Directed MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    DirectedMacAddress interface{}

    // Target MEP ID. The type is interface{} with range: 0..65535.
    TargetMepId interface{}

    // Timestamp of initiation time (seconds). The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    Timestamp interface{}

    // Time to live. The type is interface{} with range: 0..255.
    Ttl interface{}

    // Transaction ID. The type is interface{} with range: 0..4294967295.
    TransactionId interface{}

    // Options affecting traceroute behavior.
    Options Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options
}

func (tracerouteInformation *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation) GetEntityData() *types.CommonEntityData {
    tracerouteInformation.EntityData.YFilter = tracerouteInformation.YFilter
    tracerouteInformation.EntityData.YangName = "traceroute-information"
    tracerouteInformation.EntityData.BundleName = "cisco_ios_xr"
    tracerouteInformation.EntityData.ParentYangName = "traceroute-cache"
    tracerouteInformation.EntityData.SegmentPath = "traceroute-information"
    tracerouteInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracerouteInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracerouteInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracerouteInformation.EntityData.Children = make(map[string]types.YChild)
    tracerouteInformation.EntityData.Children["options"] = types.YChild{"Options", &tracerouteInformation.Options}
    tracerouteInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    tracerouteInformation.EntityData.Leafs["domain"] = types.YLeaf{"Domain", tracerouteInformation.Domain}
    tracerouteInformation.EntityData.Leafs["service"] = types.YLeaf{"Service", tracerouteInformation.Service}
    tracerouteInformation.EntityData.Leafs["level"] = types.YLeaf{"Level", tracerouteInformation.Level}
    tracerouteInformation.EntityData.Leafs["source-mep-id"] = types.YLeaf{"SourceMepId", tracerouteInformation.SourceMepId}
    tracerouteInformation.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", tracerouteInformation.SourceInterface}
    tracerouteInformation.EntityData.Leafs["source-mac-address"] = types.YLeaf{"SourceMacAddress", tracerouteInformation.SourceMacAddress}
    tracerouteInformation.EntityData.Leafs["target-mac-address"] = types.YLeaf{"TargetMacAddress", tracerouteInformation.TargetMacAddress}
    tracerouteInformation.EntityData.Leafs["directed-mac-address"] = types.YLeaf{"DirectedMacAddress", tracerouteInformation.DirectedMacAddress}
    tracerouteInformation.EntityData.Leafs["target-mep-id"] = types.YLeaf{"TargetMepId", tracerouteInformation.TargetMepId}
    tracerouteInformation.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", tracerouteInformation.Timestamp}
    tracerouteInformation.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", tracerouteInformation.Ttl}
    tracerouteInformation.EntityData.Leafs["transaction-id"] = types.YLeaf{"TransactionId", tracerouteInformation.TransactionId}
    return &(tracerouteInformation.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options
// Options affecting traceroute behavior
type Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mode. The type is CfmPmLtMode.
    Mode interface{}

    // Options for a basic IEEE 802.1ag Linktrace.
    BasicOptions Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions

    // Options for an Exploratory Linktrace.
    ExploratoryOptions Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions
}

func (options *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options) GetEntityData() *types.CommonEntityData {
    options.EntityData.YFilter = options.YFilter
    options.EntityData.YangName = "options"
    options.EntityData.BundleName = "cisco_ios_xr"
    options.EntityData.ParentYangName = "traceroute-information"
    options.EntityData.SegmentPath = "options"
    options.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    options.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    options.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    options.EntityData.Children = make(map[string]types.YChild)
    options.EntityData.Children["basic-options"] = types.YChild{"BasicOptions", &options.BasicOptions}
    options.EntityData.Children["exploratory-options"] = types.YChild{"ExploratoryOptions", &options.ExploratoryOptions}
    options.EntityData.Leafs = make(map[string]types.YLeaf)
    options.EntityData.Leafs["mode"] = types.YLeaf{"Mode", options.Mode}
    return &(options.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions
// Options for a basic IEEE 802.1ag Linktrace
type Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traceroute was initiated automatically. The type is bool.
    IsAuto interface{}

    // Only use the Filtering Database for forwarding lookups. The type is bool.
    FdbOnly interface{}
}

func (basicOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_BasicOptions) GetEntityData() *types.CommonEntityData {
    basicOptions.EntityData.YFilter = basicOptions.YFilter
    basicOptions.EntityData.YangName = "basic-options"
    basicOptions.EntityData.BundleName = "cisco_ios_xr"
    basicOptions.EntityData.ParentYangName = "options"
    basicOptions.EntityData.SegmentPath = "basic-options"
    basicOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicOptions.EntityData.Children = make(map[string]types.YChild)
    basicOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    basicOptions.EntityData.Leafs["is-auto"] = types.YLeaf{"IsAuto", basicOptions.IsAuto}
    basicOptions.EntityData.Leafs["fdb-only"] = types.YLeaf{"FdbOnly", basicOptions.FdbOnly}
    return &(basicOptions.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions
// Options for an Exploratory Linktrace
type Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay model for delay calculations. The type is CfmPmEltDelayModel.
    DelayModel interface{}

    // Constant Factor for delay calculations. The type is interface{} with range:
    // 0..4294967295.
    DelayConstantFactor interface{}

    // Reply Filtering mode used by responders. The type is CfmPmElmReplyFilter.
    ReplyFilter interface{}
}

func (exploratoryOptions *Cfm_Global_TracerouteCaches_TracerouteCache_TracerouteInformation_Options_ExploratoryOptions) GetEntityData() *types.CommonEntityData {
    exploratoryOptions.EntityData.YFilter = exploratoryOptions.YFilter
    exploratoryOptions.EntityData.YangName = "exploratory-options"
    exploratoryOptions.EntityData.BundleName = "cisco_ios_xr"
    exploratoryOptions.EntityData.ParentYangName = "options"
    exploratoryOptions.EntityData.SegmentPath = "exploratory-options"
    exploratoryOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exploratoryOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exploratoryOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exploratoryOptions.EntityData.Children = make(map[string]types.YChild)
    exploratoryOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    exploratoryOptions.EntityData.Leafs["delay-model"] = types.YLeaf{"DelayModel", exploratoryOptions.DelayModel}
    exploratoryOptions.EntityData.Leafs["delay-constant-factor"] = types.YLeaf{"DelayConstantFactor", exploratoryOptions.DelayConstantFactor}
    exploratoryOptions.EntityData.Leafs["reply-filter"] = types.YLeaf{"ReplyFilter", exploratoryOptions.ReplyFilter}
    return &(exploratoryOptions.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply
// Received linktrace replies
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Undecoded frame. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    RawData interface{}

    // Frame header.
    Header Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header

    // Sender ID TLV.
    SenderId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId

    // Egress ID TLV.
    EgressId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId

    // Reply ingress TLV.
    ReplyIngress Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress

    // Reply egress TLV.
    ReplyEgress Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress

    // Last hop ID.
    LastHop Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop

    // Organizational-specific TLVs. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv.
    OrganizationSpecificTlv []Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv

    // Unknown TLVs. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv.
    UnknownTlv []Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv
}

func (linktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply) GetEntityData() *types.CommonEntityData {
    linktraceReply.EntityData.YFilter = linktraceReply.YFilter
    linktraceReply.EntityData.YangName = "linktrace-reply"
    linktraceReply.EntityData.BundleName = "cisco_ios_xr"
    linktraceReply.EntityData.ParentYangName = "traceroute-cache"
    linktraceReply.EntityData.SegmentPath = "linktrace-reply"
    linktraceReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linktraceReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linktraceReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linktraceReply.EntityData.Children = make(map[string]types.YChild)
    linktraceReply.EntityData.Children["header"] = types.YChild{"Header", &linktraceReply.Header}
    linktraceReply.EntityData.Children["sender-id"] = types.YChild{"SenderId", &linktraceReply.SenderId}
    linktraceReply.EntityData.Children["egress-id"] = types.YChild{"EgressId", &linktraceReply.EgressId}
    linktraceReply.EntityData.Children["reply-ingress"] = types.YChild{"ReplyIngress", &linktraceReply.ReplyIngress}
    linktraceReply.EntityData.Children["reply-egress"] = types.YChild{"ReplyEgress", &linktraceReply.ReplyEgress}
    linktraceReply.EntityData.Children["last-hop"] = types.YChild{"LastHop", &linktraceReply.LastHop}
    linktraceReply.EntityData.Children["organization-specific-tlv"] = types.YChild{"OrganizationSpecificTlv", nil}
    for i := range linktraceReply.OrganizationSpecificTlv {
        linktraceReply.EntityData.Children[types.GetSegmentPath(&linktraceReply.OrganizationSpecificTlv[i])] = types.YChild{"OrganizationSpecificTlv", &linktraceReply.OrganizationSpecificTlv[i]}
    }
    linktraceReply.EntityData.Children["unknown-tlv"] = types.YChild{"UnknownTlv", nil}
    for i := range linktraceReply.UnknownTlv {
        linktraceReply.EntityData.Children[types.GetSegmentPath(&linktraceReply.UnknownTlv[i])] = types.YChild{"UnknownTlv", &linktraceReply.UnknownTlv[i]}
    }
    linktraceReply.EntityData.Leafs = make(map[string]types.YLeaf)
    linktraceReply.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", linktraceReply.RawData}
    return &(linktraceReply.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header
// Frame header
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MD level. The type is CfmBagMdLevel.
    Level interface{}

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Use filtering DB only. The type is bool.
    UseFdbOnly interface{}

    // LTR was forwarded. The type is bool.
    Forwarded interface{}

    // Terminal MEP reached. The type is bool.
    TerminalMep interface{}

    // Transaction ID. The type is interface{} with range: 0..4294967295.
    TransactionId interface{}

    // TTL. The type is interface{} with range: 0..255.
    Ttl interface{}

    // Relay action. The type is CfmPmRelayAction.
    RelayAction interface{}
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "linktrace-reply"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["level"] = types.YLeaf{"Level", header.Level}
    header.EntityData.Leafs["version"] = types.YLeaf{"Version", header.Version}
    header.EntityData.Leafs["use-fdb-only"] = types.YLeaf{"UseFdbOnly", header.UseFdbOnly}
    header.EntityData.Leafs["forwarded"] = types.YLeaf{"Forwarded", header.Forwarded}
    header.EntityData.Leafs["terminal-mep"] = types.YLeaf{"TerminalMep", header.TerminalMep}
    header.EntityData.Leafs["transaction-id"] = types.YLeaf{"TransactionId", header.TransactionId}
    header.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", header.Ttl}
    header.EntityData.Leafs["relay-action"] = types.YLeaf{"RelayAction", header.RelayAction}
    return &(header.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId
// Sender ID TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Management address domain. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ManagementAddressDomain interface{}

    // Management address. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ManagementAddress interface{}

    // Chassis ID.
    ChassisId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId) GetEntityData() *types.CommonEntityData {
    senderId.EntityData.YFilter = senderId.YFilter
    senderId.EntityData.YangName = "sender-id"
    senderId.EntityData.BundleName = "cisco_ios_xr"
    senderId.EntityData.ParentYangName = "linktrace-reply"
    senderId.EntityData.SegmentPath = "sender-id"
    senderId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    senderId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    senderId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    senderId.EntityData.Children = make(map[string]types.YChild)
    senderId.EntityData.Children["chassis-id"] = types.YChild{"ChassisId", &senderId.ChassisId}
    senderId.EntityData.Leafs = make(map[string]types.YLeaf)
    senderId.EntityData.Leafs["management-address-domain"] = types.YLeaf{"ManagementAddressDomain", senderId.ManagementAddressDomain}
    senderId.EntityData.Leafs["management-address"] = types.YLeaf{"ManagementAddress", senderId.ManagementAddress}
    return &(senderId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId
// Chassis ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Chassis ID Type. The type is CfmPmChassisIdFmt.
    ChassisIdType interface{}

    // Chassis ID Type. The type is interface{} with range: 0..255.
    ChassisIdTypeValue interface{}

    // Chassis ID (Deprecated). The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ChassisId interface{}

    // Chassis ID (Current).
    ChassisIdValue Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId) GetEntityData() *types.CommonEntityData {
    chassisId.EntityData.YFilter = chassisId.YFilter
    chassisId.EntityData.YangName = "chassis-id"
    chassisId.EntityData.BundleName = "cisco_ios_xr"
    chassisId.EntityData.ParentYangName = "sender-id"
    chassisId.EntityData.SegmentPath = "chassis-id"
    chassisId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisId.EntityData.Children = make(map[string]types.YChild)
    chassisId.EntityData.Children["chassis-id-value"] = types.YChild{"ChassisIdValue", &chassisId.ChassisIdValue}
    chassisId.EntityData.Leafs = make(map[string]types.YLeaf)
    chassisId.EntityData.Leafs["chassis-id-type"] = types.YLeaf{"ChassisIdType", chassisId.ChassisIdType}
    chassisId.EntityData.Leafs["chassis-id-type-value"] = types.YLeaf{"ChassisIdTypeValue", chassisId.ChassisIdTypeValue}
    chassisId.EntityData.Leafs["chassis-id"] = types.YLeaf{"ChassisId", chassisId.ChassisId}
    return &(chassisId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue
// Chassis ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ChassisIDFormat. The type is CfmPmIdFmt.
    ChassisIdFormat interface{}

    // Chassis ID String. The type is string.
    ChassisIdString interface{}

    // Chassis ID MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    ChassisIdMac interface{}

    // Raw Chassis ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ChassisIdRaw interface{}
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_SenderId_ChassisId_ChassisIdValue) GetEntityData() *types.CommonEntityData {
    chassisIdValue.EntityData.YFilter = chassisIdValue.YFilter
    chassisIdValue.EntityData.YangName = "chassis-id-value"
    chassisIdValue.EntityData.BundleName = "cisco_ios_xr"
    chassisIdValue.EntityData.ParentYangName = "chassis-id"
    chassisIdValue.EntityData.SegmentPath = "chassis-id-value"
    chassisIdValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisIdValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisIdValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisIdValue.EntityData.Children = make(map[string]types.YChild)
    chassisIdValue.EntityData.Leafs = make(map[string]types.YLeaf)
    chassisIdValue.EntityData.Leafs["chassis-id-format"] = types.YLeaf{"ChassisIdFormat", chassisIdValue.ChassisIdFormat}
    chassisIdValue.EntityData.Leafs["chassis-id-string"] = types.YLeaf{"ChassisIdString", chassisIdValue.ChassisIdString}
    chassisIdValue.EntityData.Leafs["chassis-id-mac"] = types.YLeaf{"ChassisIdMac", chassisIdValue.ChassisIdMac}
    chassisIdValue.EntityData.Leafs["chassis-id-raw"] = types.YLeaf{"ChassisIdRaw", chassisIdValue.ChassisIdRaw}
    return &(chassisIdValue.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId
// Egress ID TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Last egress ID.
    LastEgressId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId

    // Next egress ID.
    NextEgressId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId) GetEntityData() *types.CommonEntityData {
    egressId.EntityData.YFilter = egressId.YFilter
    egressId.EntityData.YangName = "egress-id"
    egressId.EntityData.BundleName = "cisco_ios_xr"
    egressId.EntityData.ParentYangName = "linktrace-reply"
    egressId.EntityData.SegmentPath = "egress-id"
    egressId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    egressId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    egressId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    egressId.EntityData.Children = make(map[string]types.YChild)
    egressId.EntityData.Children["last-egress-id"] = types.YChild{"LastEgressId", &egressId.LastEgressId}
    egressId.EntityData.Children["next-egress-id"] = types.YChild{"NextEgressId", &egressId.NextEgressId}
    egressId.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(egressId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId
// Last egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_LastEgressId) GetEntityData() *types.CommonEntityData {
    lastEgressId.EntityData.YFilter = lastEgressId.YFilter
    lastEgressId.EntityData.YangName = "last-egress-id"
    lastEgressId.EntityData.BundleName = "cisco_ios_xr"
    lastEgressId.EntityData.ParentYangName = "egress-id"
    lastEgressId.EntityData.SegmentPath = "last-egress-id"
    lastEgressId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastEgressId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastEgressId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastEgressId.EntityData.Children = make(map[string]types.YChild)
    lastEgressId.EntityData.Leafs = make(map[string]types.YLeaf)
    lastEgressId.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", lastEgressId.UniqueId}
    lastEgressId.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", lastEgressId.MacAddress}
    return &(lastEgressId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId
// Next egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_EgressId_NextEgressId) GetEntityData() *types.CommonEntityData {
    nextEgressId.EntityData.YFilter = nextEgressId.YFilter
    nextEgressId.EntityData.YangName = "next-egress-id"
    nextEgressId.EntityData.BundleName = "cisco_ios_xr"
    nextEgressId.EntityData.ParentYangName = "egress-id"
    nextEgressId.EntityData.SegmentPath = "next-egress-id"
    nextEgressId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextEgressId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextEgressId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextEgressId.EntityData.Children = make(map[string]types.YChild)
    nextEgressId.EntityData.Leafs = make(map[string]types.YLeaf)
    nextEgressId.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", nextEgressId.UniqueId}
    nextEgressId.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", nextEgressId.MacAddress}
    return &(nextEgressId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress
// Reply ingress TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reply ingress action. The type is CfmPmIngressAction.
    Action interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Port ID.
    PortId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress) GetEntityData() *types.CommonEntityData {
    replyIngress.EntityData.YFilter = replyIngress.YFilter
    replyIngress.EntityData.YangName = "reply-ingress"
    replyIngress.EntityData.BundleName = "cisco_ios_xr"
    replyIngress.EntityData.ParentYangName = "linktrace-reply"
    replyIngress.EntityData.SegmentPath = "reply-ingress"
    replyIngress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replyIngress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replyIngress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replyIngress.EntityData.Children = make(map[string]types.YChild)
    replyIngress.EntityData.Children["port-id"] = types.YChild{"PortId", &replyIngress.PortId}
    replyIngress.EntityData.Leafs = make(map[string]types.YLeaf)
    replyIngress.EntityData.Leafs["action"] = types.YLeaf{"Action", replyIngress.Action}
    replyIngress.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", replyIngress.MacAddress}
    return &(replyIngress.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId
// Port ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port ID type. The type is CfmPmPortIdFmt.
    PortIdType interface{}

    // Port ID type value. The type is interface{} with range: 0..255.
    PortIdTypeValue interface{}

    // Port ID (Deprecated). The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    PortId interface{}

    // Port ID (Current).
    PortIdValue Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId) GetEntityData() *types.CommonEntityData {
    portId.EntityData.YFilter = portId.YFilter
    portId.EntityData.YangName = "port-id"
    portId.EntityData.BundleName = "cisco_ios_xr"
    portId.EntityData.ParentYangName = "reply-ingress"
    portId.EntityData.SegmentPath = "port-id"
    portId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portId.EntityData.Children = make(map[string]types.YChild)
    portId.EntityData.Children["port-id-value"] = types.YChild{"PortIdValue", &portId.PortIdValue}
    portId.EntityData.Leafs = make(map[string]types.YLeaf)
    portId.EntityData.Leafs["port-id-type"] = types.YLeaf{"PortIdType", portId.PortIdType}
    portId.EntityData.Leafs["port-id-type-value"] = types.YLeaf{"PortIdTypeValue", portId.PortIdTypeValue}
    portId.EntityData.Leafs["port-id"] = types.YLeaf{"PortId", portId.PortId}
    return &(portId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue
// Port ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PortIDFormat. The type is CfmPmIdFmt.
    PortIdFormat interface{}

    // Port ID String. The type is string.
    PortIdString interface{}

    // Port ID MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    PortIdMac interface{}

    // Raw Port ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    PortIdRaw interface{}
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyIngress_PortId_PortIdValue) GetEntityData() *types.CommonEntityData {
    portIdValue.EntityData.YFilter = portIdValue.YFilter
    portIdValue.EntityData.YangName = "port-id-value"
    portIdValue.EntityData.BundleName = "cisco_ios_xr"
    portIdValue.EntityData.ParentYangName = "port-id"
    portIdValue.EntityData.SegmentPath = "port-id-value"
    portIdValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIdValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIdValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIdValue.EntityData.Children = make(map[string]types.YChild)
    portIdValue.EntityData.Leafs = make(map[string]types.YLeaf)
    portIdValue.EntityData.Leafs["port-id-format"] = types.YLeaf{"PortIdFormat", portIdValue.PortIdFormat}
    portIdValue.EntityData.Leafs["port-id-string"] = types.YLeaf{"PortIdString", portIdValue.PortIdString}
    portIdValue.EntityData.Leafs["port-id-mac"] = types.YLeaf{"PortIdMac", portIdValue.PortIdMac}
    portIdValue.EntityData.Leafs["port-id-raw"] = types.YLeaf{"PortIdRaw", portIdValue.PortIdRaw}
    return &(portIdValue.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress
// Reply egress TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reply egress action. The type is CfmPmEgressAction.
    Action interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Port ID.
    PortId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress) GetEntityData() *types.CommonEntityData {
    replyEgress.EntityData.YFilter = replyEgress.YFilter
    replyEgress.EntityData.YangName = "reply-egress"
    replyEgress.EntityData.BundleName = "cisco_ios_xr"
    replyEgress.EntityData.ParentYangName = "linktrace-reply"
    replyEgress.EntityData.SegmentPath = "reply-egress"
    replyEgress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replyEgress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replyEgress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replyEgress.EntityData.Children = make(map[string]types.YChild)
    replyEgress.EntityData.Children["port-id"] = types.YChild{"PortId", &replyEgress.PortId}
    replyEgress.EntityData.Leafs = make(map[string]types.YLeaf)
    replyEgress.EntityData.Leafs["action"] = types.YLeaf{"Action", replyEgress.Action}
    replyEgress.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", replyEgress.MacAddress}
    return &(replyEgress.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId
// Port ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port ID type. The type is CfmPmPortIdFmt.
    PortIdType interface{}

    // Port ID type value. The type is interface{} with range: 0..255.
    PortIdTypeValue interface{}

    // Port ID (Deprecated). The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    PortId interface{}

    // Port ID (Current).
    PortIdValue Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId) GetEntityData() *types.CommonEntityData {
    portId.EntityData.YFilter = portId.YFilter
    portId.EntityData.YangName = "port-id"
    portId.EntityData.BundleName = "cisco_ios_xr"
    portId.EntityData.ParentYangName = "reply-egress"
    portId.EntityData.SegmentPath = "port-id"
    portId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portId.EntityData.Children = make(map[string]types.YChild)
    portId.EntityData.Children["port-id-value"] = types.YChild{"PortIdValue", &portId.PortIdValue}
    portId.EntityData.Leafs = make(map[string]types.YLeaf)
    portId.EntityData.Leafs["port-id-type"] = types.YLeaf{"PortIdType", portId.PortIdType}
    portId.EntityData.Leafs["port-id-type-value"] = types.YLeaf{"PortIdTypeValue", portId.PortIdTypeValue}
    portId.EntityData.Leafs["port-id"] = types.YLeaf{"PortId", portId.PortId}
    return &(portId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue
// Port ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PortIDFormat. The type is CfmPmIdFmt.
    PortIdFormat interface{}

    // Port ID String. The type is string.
    PortIdString interface{}

    // Port ID MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    PortIdMac interface{}

    // Raw Port ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    PortIdRaw interface{}
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_ReplyEgress_PortId_PortIdValue) GetEntityData() *types.CommonEntityData {
    portIdValue.EntityData.YFilter = portIdValue.YFilter
    portIdValue.EntityData.YangName = "port-id-value"
    portIdValue.EntityData.BundleName = "cisco_ios_xr"
    portIdValue.EntityData.ParentYangName = "port-id"
    portIdValue.EntityData.SegmentPath = "port-id-value"
    portIdValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIdValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIdValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIdValue.EntityData.Children = make(map[string]types.YChild)
    portIdValue.EntityData.Leafs = make(map[string]types.YLeaf)
    portIdValue.EntityData.Leafs["port-id-format"] = types.YLeaf{"PortIdFormat", portIdValue.PortIdFormat}
    portIdValue.EntityData.Leafs["port-id-string"] = types.YLeaf{"PortIdString", portIdValue.PortIdString}
    portIdValue.EntityData.Leafs["port-id-mac"] = types.YLeaf{"PortIdMac", portIdValue.PortIdMac}
    portIdValue.EntityData.Leafs["port-id-raw"] = types.YLeaf{"PortIdRaw", portIdValue.PortIdRaw}
    return &(portIdValue.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop
// Last hop ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LastHopFormat. The type is CfmPmLastHopFmt.
    LastHopFormat interface{}

    // Hostname. The type is string.
    HostName interface{}

    // Egress ID.
    EgressId Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop) GetEntityData() *types.CommonEntityData {
    lastHop.EntityData.YFilter = lastHop.YFilter
    lastHop.EntityData.YangName = "last-hop"
    lastHop.EntityData.BundleName = "cisco_ios_xr"
    lastHop.EntityData.ParentYangName = "linktrace-reply"
    lastHop.EntityData.SegmentPath = "last-hop"
    lastHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastHop.EntityData.Children = make(map[string]types.YChild)
    lastHop.EntityData.Children["egress-id"] = types.YChild{"EgressId", &lastHop.EgressId}
    lastHop.EntityData.Leafs = make(map[string]types.YLeaf)
    lastHop.EntityData.Leafs["last-hop-format"] = types.YLeaf{"LastHopFormat", lastHop.LastHopFormat}
    lastHop.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", lastHop.HostName}
    return &(lastHop.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId
// Egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_LastHop_EgressId) GetEntityData() *types.CommonEntityData {
    egressId.EntityData.YFilter = egressId.YFilter
    egressId.EntityData.YangName = "egress-id"
    egressId.EntityData.BundleName = "cisco_ios_xr"
    egressId.EntityData.ParentYangName = "last-hop"
    egressId.EntityData.SegmentPath = "egress-id"
    egressId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    egressId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    egressId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    egressId.EntityData.Children = make(map[string]types.YChild)
    egressId.EntityData.Leafs = make(map[string]types.YLeaf)
    egressId.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", egressId.UniqueId}
    egressId.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", egressId.MacAddress}
    return &(egressId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv
// Organizational-specific TLVs
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Organizationally-unique ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Oui interface{}

    // Subtype of TLV. The type is interface{} with range: 0..255.
    Subtype interface{}

    // Binary data in TLV. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_OrganizationSpecificTlv) GetEntityData() *types.CommonEntityData {
    organizationSpecificTlv.EntityData.YFilter = organizationSpecificTlv.YFilter
    organizationSpecificTlv.EntityData.YangName = "organization-specific-tlv"
    organizationSpecificTlv.EntityData.BundleName = "cisco_ios_xr"
    organizationSpecificTlv.EntityData.ParentYangName = "linktrace-reply"
    organizationSpecificTlv.EntityData.SegmentPath = "organization-specific-tlv"
    organizationSpecificTlv.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    organizationSpecificTlv.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    organizationSpecificTlv.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    organizationSpecificTlv.EntityData.Children = make(map[string]types.YChild)
    organizationSpecificTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    organizationSpecificTlv.EntityData.Leafs["oui"] = types.YLeaf{"Oui", organizationSpecificTlv.Oui}
    organizationSpecificTlv.EntityData.Leafs["subtype"] = types.YLeaf{"Subtype", organizationSpecificTlv.Subtype}
    organizationSpecificTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", organizationSpecificTlv.Value}
    return &(organizationSpecificTlv.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv
// Unknown TLVs
type Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type code of TLV. The type is interface{} with range: 0..255.
    Typecode interface{}

    // Binary data in TLV. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_LinktraceReply_UnknownTlv) GetEntityData() *types.CommonEntityData {
    unknownTlv.EntityData.YFilter = unknownTlv.YFilter
    unknownTlv.EntityData.YangName = "unknown-tlv"
    unknownTlv.EntityData.BundleName = "cisco_ios_xr"
    unknownTlv.EntityData.ParentYangName = "linktrace-reply"
    unknownTlv.EntityData.SegmentPath = "unknown-tlv"
    unknownTlv.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownTlv.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownTlv.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownTlv.EntityData.Children = make(map[string]types.YChild)
    unknownTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownTlv.EntityData.Leafs["typecode"] = types.YLeaf{"Typecode", unknownTlv.Typecode}
    unknownTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownTlv.Value}
    return &(unknownTlv.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply
// Received exploratory linktrace replies
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Undecoded frame. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    RawData interface{}

    // Frame header.
    Header Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header

    // Sender ID TLV.
    SenderId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId

    // Reply ingress TLV.
    ReplyIngress Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress

    // Reply egress TLV.
    ReplyEgress Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress

    // Last hop ID.
    LastHop Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop

    // Organizational-specific TLVs. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv.
    OrganizationSpecificTlv []Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv

    // Unknown TLVs. The type is slice of
    // Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv.
    UnknownTlv []Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv
}

func (exploratoryLinktraceReply *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply) GetEntityData() *types.CommonEntityData {
    exploratoryLinktraceReply.EntityData.YFilter = exploratoryLinktraceReply.YFilter
    exploratoryLinktraceReply.EntityData.YangName = "exploratory-linktrace-reply"
    exploratoryLinktraceReply.EntityData.BundleName = "cisco_ios_xr"
    exploratoryLinktraceReply.EntityData.ParentYangName = "traceroute-cache"
    exploratoryLinktraceReply.EntityData.SegmentPath = "exploratory-linktrace-reply"
    exploratoryLinktraceReply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exploratoryLinktraceReply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exploratoryLinktraceReply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exploratoryLinktraceReply.EntityData.Children = make(map[string]types.YChild)
    exploratoryLinktraceReply.EntityData.Children["header"] = types.YChild{"Header", &exploratoryLinktraceReply.Header}
    exploratoryLinktraceReply.EntityData.Children["sender-id"] = types.YChild{"SenderId", &exploratoryLinktraceReply.SenderId}
    exploratoryLinktraceReply.EntityData.Children["reply-ingress"] = types.YChild{"ReplyIngress", &exploratoryLinktraceReply.ReplyIngress}
    exploratoryLinktraceReply.EntityData.Children["reply-egress"] = types.YChild{"ReplyEgress", &exploratoryLinktraceReply.ReplyEgress}
    exploratoryLinktraceReply.EntityData.Children["last-hop"] = types.YChild{"LastHop", &exploratoryLinktraceReply.LastHop}
    exploratoryLinktraceReply.EntityData.Children["organization-specific-tlv"] = types.YChild{"OrganizationSpecificTlv", nil}
    for i := range exploratoryLinktraceReply.OrganizationSpecificTlv {
        exploratoryLinktraceReply.EntityData.Children[types.GetSegmentPath(&exploratoryLinktraceReply.OrganizationSpecificTlv[i])] = types.YChild{"OrganizationSpecificTlv", &exploratoryLinktraceReply.OrganizationSpecificTlv[i]}
    }
    exploratoryLinktraceReply.EntityData.Children["unknown-tlv"] = types.YChild{"UnknownTlv", nil}
    for i := range exploratoryLinktraceReply.UnknownTlv {
        exploratoryLinktraceReply.EntityData.Children[types.GetSegmentPath(&exploratoryLinktraceReply.UnknownTlv[i])] = types.YChild{"UnknownTlv", &exploratoryLinktraceReply.UnknownTlv[i]}
    }
    exploratoryLinktraceReply.EntityData.Leafs = make(map[string]types.YLeaf)
    exploratoryLinktraceReply.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", exploratoryLinktraceReply.RawData}
    return &(exploratoryLinktraceReply.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header
// Frame header
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MD level. The type is CfmBagMdLevel.
    Level interface{}

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // ELR was forwarded. The type is bool.
    Forwarded interface{}

    // Terminal MEP reached. The type is bool.
    TerminalMep interface{}

    // Reply Filter unrecognized. The type is bool.
    ReplyFilterUnknown interface{}

    // Transaction ID. The type is interface{} with range: 0..4294967295.
    TransactionId interface{}

    // TTL. The type is interface{} with range: 0..255.
    Ttl interface{}

    // Relay action. The type is CfmPmElrRelayAction.
    RelayAction interface{}

    // Next Hop Timeout, in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    NextHopTimeout interface{}

    // Delay Model. The type is CfmPmEltDelayModel.
    DelayModel interface{}
}

func (header *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "exploratory-linktrace-reply"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["level"] = types.YLeaf{"Level", header.Level}
    header.EntityData.Leafs["version"] = types.YLeaf{"Version", header.Version}
    header.EntityData.Leafs["forwarded"] = types.YLeaf{"Forwarded", header.Forwarded}
    header.EntityData.Leafs["terminal-mep"] = types.YLeaf{"TerminalMep", header.TerminalMep}
    header.EntityData.Leafs["reply-filter-unknown"] = types.YLeaf{"ReplyFilterUnknown", header.ReplyFilterUnknown}
    header.EntityData.Leafs["transaction-id"] = types.YLeaf{"TransactionId", header.TransactionId}
    header.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", header.Ttl}
    header.EntityData.Leafs["relay-action"] = types.YLeaf{"RelayAction", header.RelayAction}
    header.EntityData.Leafs["next-hop-timeout"] = types.YLeaf{"NextHopTimeout", header.NextHopTimeout}
    header.EntityData.Leafs["delay-model"] = types.YLeaf{"DelayModel", header.DelayModel}
    return &(header.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId
// Sender ID TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Management address domain. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ManagementAddressDomain interface{}

    // Management address. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ManagementAddress interface{}

    // Chassis ID.
    ChassisId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId
}

func (senderId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId) GetEntityData() *types.CommonEntityData {
    senderId.EntityData.YFilter = senderId.YFilter
    senderId.EntityData.YangName = "sender-id"
    senderId.EntityData.BundleName = "cisco_ios_xr"
    senderId.EntityData.ParentYangName = "exploratory-linktrace-reply"
    senderId.EntityData.SegmentPath = "sender-id"
    senderId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    senderId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    senderId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    senderId.EntityData.Children = make(map[string]types.YChild)
    senderId.EntityData.Children["chassis-id"] = types.YChild{"ChassisId", &senderId.ChassisId}
    senderId.EntityData.Leafs = make(map[string]types.YLeaf)
    senderId.EntityData.Leafs["management-address-domain"] = types.YLeaf{"ManagementAddressDomain", senderId.ManagementAddressDomain}
    senderId.EntityData.Leafs["management-address"] = types.YLeaf{"ManagementAddress", senderId.ManagementAddress}
    return &(senderId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId
// Chassis ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Chassis ID Type. The type is CfmPmChassisIdFmt.
    ChassisIdType interface{}

    // Chassis ID Type. The type is interface{} with range: 0..255.
    ChassisIdTypeValue interface{}

    // Chassis ID (Deprecated). The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ChassisId interface{}

    // Chassis ID (Current).
    ChassisIdValue Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue
}

func (chassisId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId) GetEntityData() *types.CommonEntityData {
    chassisId.EntityData.YFilter = chassisId.YFilter
    chassisId.EntityData.YangName = "chassis-id"
    chassisId.EntityData.BundleName = "cisco_ios_xr"
    chassisId.EntityData.ParentYangName = "sender-id"
    chassisId.EntityData.SegmentPath = "chassis-id"
    chassisId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisId.EntityData.Children = make(map[string]types.YChild)
    chassisId.EntityData.Children["chassis-id-value"] = types.YChild{"ChassisIdValue", &chassisId.ChassisIdValue}
    chassisId.EntityData.Leafs = make(map[string]types.YLeaf)
    chassisId.EntityData.Leafs["chassis-id-type"] = types.YLeaf{"ChassisIdType", chassisId.ChassisIdType}
    chassisId.EntityData.Leafs["chassis-id-type-value"] = types.YLeaf{"ChassisIdTypeValue", chassisId.ChassisIdTypeValue}
    chassisId.EntityData.Leafs["chassis-id"] = types.YLeaf{"ChassisId", chassisId.ChassisId}
    return &(chassisId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue
// Chassis ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ChassisIDFormat. The type is CfmPmIdFmt.
    ChassisIdFormat interface{}

    // Chassis ID String. The type is string.
    ChassisIdString interface{}

    // Chassis ID MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    ChassisIdMac interface{}

    // Raw Chassis ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ChassisIdRaw interface{}
}

func (chassisIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_SenderId_ChassisId_ChassisIdValue) GetEntityData() *types.CommonEntityData {
    chassisIdValue.EntityData.YFilter = chassisIdValue.YFilter
    chassisIdValue.EntityData.YangName = "chassis-id-value"
    chassisIdValue.EntityData.BundleName = "cisco_ios_xr"
    chassisIdValue.EntityData.ParentYangName = "chassis-id"
    chassisIdValue.EntityData.SegmentPath = "chassis-id-value"
    chassisIdValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisIdValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisIdValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisIdValue.EntityData.Children = make(map[string]types.YChild)
    chassisIdValue.EntityData.Leafs = make(map[string]types.YLeaf)
    chassisIdValue.EntityData.Leafs["chassis-id-format"] = types.YLeaf{"ChassisIdFormat", chassisIdValue.ChassisIdFormat}
    chassisIdValue.EntityData.Leafs["chassis-id-string"] = types.YLeaf{"ChassisIdString", chassisIdValue.ChassisIdString}
    chassisIdValue.EntityData.Leafs["chassis-id-mac"] = types.YLeaf{"ChassisIdMac", chassisIdValue.ChassisIdMac}
    chassisIdValue.EntityData.Leafs["chassis-id-raw"] = types.YLeaf{"ChassisIdRaw", chassisIdValue.ChassisIdRaw}
    return &(chassisIdValue.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress
// Reply ingress TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ELR Reply ingress action. The type is CfmPmElrIngressAction.
    Action interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Last egress ID.
    LastEgressId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId

    // Next egress ID.
    NextEgressId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId

    // Port ID.
    PortId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId
}

func (replyIngress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress) GetEntityData() *types.CommonEntityData {
    replyIngress.EntityData.YFilter = replyIngress.YFilter
    replyIngress.EntityData.YangName = "reply-ingress"
    replyIngress.EntityData.BundleName = "cisco_ios_xr"
    replyIngress.EntityData.ParentYangName = "exploratory-linktrace-reply"
    replyIngress.EntityData.SegmentPath = "reply-ingress"
    replyIngress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replyIngress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replyIngress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replyIngress.EntityData.Children = make(map[string]types.YChild)
    replyIngress.EntityData.Children["last-egress-id"] = types.YChild{"LastEgressId", &replyIngress.LastEgressId}
    replyIngress.EntityData.Children["next-egress-id"] = types.YChild{"NextEgressId", &replyIngress.NextEgressId}
    replyIngress.EntityData.Children["port-id"] = types.YChild{"PortId", &replyIngress.PortId}
    replyIngress.EntityData.Leafs = make(map[string]types.YLeaf)
    replyIngress.EntityData.Leafs["action"] = types.YLeaf{"Action", replyIngress.Action}
    replyIngress.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", replyIngress.MacAddress}
    return &(replyIngress.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId
// Last egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_LastEgressId) GetEntityData() *types.CommonEntityData {
    lastEgressId.EntityData.YFilter = lastEgressId.YFilter
    lastEgressId.EntityData.YangName = "last-egress-id"
    lastEgressId.EntityData.BundleName = "cisco_ios_xr"
    lastEgressId.EntityData.ParentYangName = "reply-ingress"
    lastEgressId.EntityData.SegmentPath = "last-egress-id"
    lastEgressId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastEgressId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastEgressId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastEgressId.EntityData.Children = make(map[string]types.YChild)
    lastEgressId.EntityData.Leafs = make(map[string]types.YLeaf)
    lastEgressId.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", lastEgressId.UniqueId}
    lastEgressId.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", lastEgressId.MacAddress}
    return &(lastEgressId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId
// Next egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_NextEgressId) GetEntityData() *types.CommonEntityData {
    nextEgressId.EntityData.YFilter = nextEgressId.YFilter
    nextEgressId.EntityData.YangName = "next-egress-id"
    nextEgressId.EntityData.BundleName = "cisco_ios_xr"
    nextEgressId.EntityData.ParentYangName = "reply-ingress"
    nextEgressId.EntityData.SegmentPath = "next-egress-id"
    nextEgressId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextEgressId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextEgressId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextEgressId.EntityData.Children = make(map[string]types.YChild)
    nextEgressId.EntityData.Leafs = make(map[string]types.YLeaf)
    nextEgressId.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", nextEgressId.UniqueId}
    nextEgressId.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", nextEgressId.MacAddress}
    return &(nextEgressId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId
// Port ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port ID type. The type is CfmPmPortIdFmt.
    PortIdType interface{}

    // Port ID type value. The type is interface{} with range: 0..255.
    PortIdTypeValue interface{}

    // Port ID (Deprecated). The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    PortId interface{}

    // Port ID (Current).
    PortIdValue Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId) GetEntityData() *types.CommonEntityData {
    portId.EntityData.YFilter = portId.YFilter
    portId.EntityData.YangName = "port-id"
    portId.EntityData.BundleName = "cisco_ios_xr"
    portId.EntityData.ParentYangName = "reply-ingress"
    portId.EntityData.SegmentPath = "port-id"
    portId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portId.EntityData.Children = make(map[string]types.YChild)
    portId.EntityData.Children["port-id-value"] = types.YChild{"PortIdValue", &portId.PortIdValue}
    portId.EntityData.Leafs = make(map[string]types.YLeaf)
    portId.EntityData.Leafs["port-id-type"] = types.YLeaf{"PortIdType", portId.PortIdType}
    portId.EntityData.Leafs["port-id-type-value"] = types.YLeaf{"PortIdTypeValue", portId.PortIdTypeValue}
    portId.EntityData.Leafs["port-id"] = types.YLeaf{"PortId", portId.PortId}
    return &(portId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue
// Port ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PortIDFormat. The type is CfmPmIdFmt.
    PortIdFormat interface{}

    // Port ID String. The type is string.
    PortIdString interface{}

    // Port ID MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    PortIdMac interface{}

    // Raw Port ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    PortIdRaw interface{}
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyIngress_PortId_PortIdValue) GetEntityData() *types.CommonEntityData {
    portIdValue.EntityData.YFilter = portIdValue.YFilter
    portIdValue.EntityData.YangName = "port-id-value"
    portIdValue.EntityData.BundleName = "cisco_ios_xr"
    portIdValue.EntityData.ParentYangName = "port-id"
    portIdValue.EntityData.SegmentPath = "port-id-value"
    portIdValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIdValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIdValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIdValue.EntityData.Children = make(map[string]types.YChild)
    portIdValue.EntityData.Leafs = make(map[string]types.YLeaf)
    portIdValue.EntityData.Leafs["port-id-format"] = types.YLeaf{"PortIdFormat", portIdValue.PortIdFormat}
    portIdValue.EntityData.Leafs["port-id-string"] = types.YLeaf{"PortIdString", portIdValue.PortIdString}
    portIdValue.EntityData.Leafs["port-id-mac"] = types.YLeaf{"PortIdMac", portIdValue.PortIdMac}
    portIdValue.EntityData.Leafs["port-id-raw"] = types.YLeaf{"PortIdRaw", portIdValue.PortIdRaw}
    return &(portIdValue.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress
// Reply egress TLV
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reply egress action. The type is CfmPmElrEgressAction.
    Action interface{}

    // MAC address of egress interface. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Last Egress ID.
    LastEgressId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId

    // Next Egress ID.
    NextEgressId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId

    // Port ID.
    PortId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId
}

func (replyEgress *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress) GetEntityData() *types.CommonEntityData {
    replyEgress.EntityData.YFilter = replyEgress.YFilter
    replyEgress.EntityData.YangName = "reply-egress"
    replyEgress.EntityData.BundleName = "cisco_ios_xr"
    replyEgress.EntityData.ParentYangName = "exploratory-linktrace-reply"
    replyEgress.EntityData.SegmentPath = "reply-egress"
    replyEgress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replyEgress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replyEgress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replyEgress.EntityData.Children = make(map[string]types.YChild)
    replyEgress.EntityData.Children["last-egress-id"] = types.YChild{"LastEgressId", &replyEgress.LastEgressId}
    replyEgress.EntityData.Children["next-egress-id"] = types.YChild{"NextEgressId", &replyEgress.NextEgressId}
    replyEgress.EntityData.Children["port-id"] = types.YChild{"PortId", &replyEgress.PortId}
    replyEgress.EntityData.Leafs = make(map[string]types.YLeaf)
    replyEgress.EntityData.Leafs["action"] = types.YLeaf{"Action", replyEgress.Action}
    replyEgress.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", replyEgress.MacAddress}
    return &(replyEgress.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId
// Last Egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}
}

func (lastEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_LastEgressId) GetEntityData() *types.CommonEntityData {
    lastEgressId.EntityData.YFilter = lastEgressId.YFilter
    lastEgressId.EntityData.YangName = "last-egress-id"
    lastEgressId.EntityData.BundleName = "cisco_ios_xr"
    lastEgressId.EntityData.ParentYangName = "reply-egress"
    lastEgressId.EntityData.SegmentPath = "last-egress-id"
    lastEgressId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastEgressId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastEgressId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastEgressId.EntityData.Children = make(map[string]types.YChild)
    lastEgressId.EntityData.Leafs = make(map[string]types.YLeaf)
    lastEgressId.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", lastEgressId.UniqueId}
    lastEgressId.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", lastEgressId.MacAddress}
    return &(lastEgressId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId
// Next Egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}
}

func (nextEgressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_NextEgressId) GetEntityData() *types.CommonEntityData {
    nextEgressId.EntityData.YFilter = nextEgressId.YFilter
    nextEgressId.EntityData.YangName = "next-egress-id"
    nextEgressId.EntityData.BundleName = "cisco_ios_xr"
    nextEgressId.EntityData.ParentYangName = "reply-egress"
    nextEgressId.EntityData.SegmentPath = "next-egress-id"
    nextEgressId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextEgressId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextEgressId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextEgressId.EntityData.Children = make(map[string]types.YChild)
    nextEgressId.EntityData.Leafs = make(map[string]types.YLeaf)
    nextEgressId.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", nextEgressId.UniqueId}
    nextEgressId.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", nextEgressId.MacAddress}
    return &(nextEgressId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId
// Port ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port ID type. The type is CfmPmPortIdFmt.
    PortIdType interface{}

    // Port ID type value. The type is interface{} with range: 0..255.
    PortIdTypeValue interface{}

    // Port ID (Deprecated). The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    PortId interface{}

    // Port ID (Current).
    PortIdValue Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue
}

func (portId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId) GetEntityData() *types.CommonEntityData {
    portId.EntityData.YFilter = portId.YFilter
    portId.EntityData.YangName = "port-id"
    portId.EntityData.BundleName = "cisco_ios_xr"
    portId.EntityData.ParentYangName = "reply-egress"
    portId.EntityData.SegmentPath = "port-id"
    portId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portId.EntityData.Children = make(map[string]types.YChild)
    portId.EntityData.Children["port-id-value"] = types.YChild{"PortIdValue", &portId.PortIdValue}
    portId.EntityData.Leafs = make(map[string]types.YLeaf)
    portId.EntityData.Leafs["port-id-type"] = types.YLeaf{"PortIdType", portId.PortIdType}
    portId.EntityData.Leafs["port-id-type-value"] = types.YLeaf{"PortIdTypeValue", portId.PortIdTypeValue}
    portId.EntityData.Leafs["port-id"] = types.YLeaf{"PortId", portId.PortId}
    return &(portId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue
// Port ID (Current)
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PortIDFormat. The type is CfmPmIdFmt.
    PortIdFormat interface{}

    // Port ID String. The type is string.
    PortIdString interface{}

    // Port ID MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    PortIdMac interface{}

    // Raw Port ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    PortIdRaw interface{}
}

func (portIdValue *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_ReplyEgress_PortId_PortIdValue) GetEntityData() *types.CommonEntityData {
    portIdValue.EntityData.YFilter = portIdValue.YFilter
    portIdValue.EntityData.YangName = "port-id-value"
    portIdValue.EntityData.BundleName = "cisco_ios_xr"
    portIdValue.EntityData.ParentYangName = "port-id"
    portIdValue.EntityData.SegmentPath = "port-id-value"
    portIdValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIdValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIdValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIdValue.EntityData.Children = make(map[string]types.YChild)
    portIdValue.EntityData.Leafs = make(map[string]types.YLeaf)
    portIdValue.EntityData.Leafs["port-id-format"] = types.YLeaf{"PortIdFormat", portIdValue.PortIdFormat}
    portIdValue.EntityData.Leafs["port-id-string"] = types.YLeaf{"PortIdString", portIdValue.PortIdString}
    portIdValue.EntityData.Leafs["port-id-mac"] = types.YLeaf{"PortIdMac", portIdValue.PortIdMac}
    portIdValue.EntityData.Leafs["port-id-raw"] = types.YLeaf{"PortIdRaw", portIdValue.PortIdRaw}
    return &(portIdValue.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop
// Last hop ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LastHopFormat. The type is CfmPmLastHopFmt.
    LastHopFormat interface{}

    // Hostname. The type is string.
    HostName interface{}

    // Egress ID.
    EgressId Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId
}

func (lastHop *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop) GetEntityData() *types.CommonEntityData {
    lastHop.EntityData.YFilter = lastHop.YFilter
    lastHop.EntityData.YangName = "last-hop"
    lastHop.EntityData.BundleName = "cisco_ios_xr"
    lastHop.EntityData.ParentYangName = "exploratory-linktrace-reply"
    lastHop.EntityData.SegmentPath = "last-hop"
    lastHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastHop.EntityData.Children = make(map[string]types.YChild)
    lastHop.EntityData.Children["egress-id"] = types.YChild{"EgressId", &lastHop.EgressId}
    lastHop.EntityData.Leafs = make(map[string]types.YLeaf)
    lastHop.EntityData.Leafs["last-hop-format"] = types.YLeaf{"LastHopFormat", lastHop.LastHopFormat}
    lastHop.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", lastHop.HostName}
    return &(lastHop.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId
// Egress ID
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique ID. The type is interface{} with range: 0..65535.
    UniqueId interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}
}

func (egressId *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_LastHop_EgressId) GetEntityData() *types.CommonEntityData {
    egressId.EntityData.YFilter = egressId.YFilter
    egressId.EntityData.YangName = "egress-id"
    egressId.EntityData.BundleName = "cisco_ios_xr"
    egressId.EntityData.ParentYangName = "last-hop"
    egressId.EntityData.SegmentPath = "egress-id"
    egressId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    egressId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    egressId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    egressId.EntityData.Children = make(map[string]types.YChild)
    egressId.EntityData.Leafs = make(map[string]types.YLeaf)
    egressId.EntityData.Leafs["unique-id"] = types.YLeaf{"UniqueId", egressId.UniqueId}
    egressId.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", egressId.MacAddress}
    return &(egressId.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv
// Organizational-specific TLVs
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Organizationally-unique ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Oui interface{}

    // Subtype of TLV. The type is interface{} with range: 0..255.
    Subtype interface{}

    // Binary data in TLV. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (organizationSpecificTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_OrganizationSpecificTlv) GetEntityData() *types.CommonEntityData {
    organizationSpecificTlv.EntityData.YFilter = organizationSpecificTlv.YFilter
    organizationSpecificTlv.EntityData.YangName = "organization-specific-tlv"
    organizationSpecificTlv.EntityData.BundleName = "cisco_ios_xr"
    organizationSpecificTlv.EntityData.ParentYangName = "exploratory-linktrace-reply"
    organizationSpecificTlv.EntityData.SegmentPath = "organization-specific-tlv"
    organizationSpecificTlv.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    organizationSpecificTlv.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    organizationSpecificTlv.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    organizationSpecificTlv.EntityData.Children = make(map[string]types.YChild)
    organizationSpecificTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    organizationSpecificTlv.EntityData.Leafs["oui"] = types.YLeaf{"Oui", organizationSpecificTlv.Oui}
    organizationSpecificTlv.EntityData.Leafs["subtype"] = types.YLeaf{"Subtype", organizationSpecificTlv.Subtype}
    organizationSpecificTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", organizationSpecificTlv.Value}
    return &(organizationSpecificTlv.EntityData)
}

// Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv
// Unknown TLVs
type Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type code of TLV. The type is interface{} with range: 0..255.
    Typecode interface{}

    // Binary data in TLV. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (unknownTlv *Cfm_Global_TracerouteCaches_TracerouteCache_ExploratoryLinktraceReply_UnknownTlv) GetEntityData() *types.CommonEntityData {
    unknownTlv.EntityData.YFilter = unknownTlv.YFilter
    unknownTlv.EntityData.YangName = "unknown-tlv"
    unknownTlv.EntityData.BundleName = "cisco_ios_xr"
    unknownTlv.EntityData.ParentYangName = "exploratory-linktrace-reply"
    unknownTlv.EntityData.SegmentPath = "unknown-tlv"
    unknownTlv.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownTlv.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownTlv.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownTlv.EntityData.Children = make(map[string]types.YChild)
    unknownTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownTlv.EntityData.Leafs["typecode"] = types.YLeaf{"Typecode", unknownTlv.Typecode}
    unknownTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownTlv.Value}
    return &(unknownTlv.EntityData)
}

// Cfm_Global_LocalMeps
// Local MEPs table
type Cfm_Global_LocalMeps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular local MEP. The type is slice of
    // Cfm_Global_LocalMeps_LocalMep.
    LocalMep []Cfm_Global_LocalMeps_LocalMep
}

func (localMeps *Cfm_Global_LocalMeps) GetEntityData() *types.CommonEntityData {
    localMeps.EntityData.YFilter = localMeps.YFilter
    localMeps.EntityData.YangName = "local-meps"
    localMeps.EntityData.BundleName = "cisco_ios_xr"
    localMeps.EntityData.ParentYangName = "global"
    localMeps.EntityData.SegmentPath = "local-meps"
    localMeps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localMeps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localMeps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localMeps.EntityData.Children = make(map[string]types.YChild)
    localMeps.EntityData.Children["local-mep"] = types.YChild{"LocalMep", nil}
    for i := range localMeps.LocalMep {
        localMeps.EntityData.Children[types.GetSegmentPath(&localMeps.LocalMep[i])] = types.YChild{"LocalMep", &localMeps.LocalMep[i]}
    }
    localMeps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(localMeps.EntityData)
}

// Cfm_Global_LocalMeps_LocalMep
// Information about a particular local MEP
type Cfm_Global_LocalMeps_LocalMep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. MEP ID. The type is interface{} with range:
    // 1..8191.
    MepId interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Maintenance domain name. The type is string.
    DomainXr interface{}

    // Service name. The type is string.
    ServiceXr interface{}

    // Maintenance level. The type is CfmBagMdLevel.
    Level interface{}

    // MEP ID. The type is interface{} with range: 0..65535.
    MepIdXr interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceXr interface{}

    // IM Interface state. The type is string.
    InterfaceState interface{}

    // Interface interworking state. The type is CfmBagIwState.
    InterworkingState interface{}

    // STP state. The type is CfmBagStpState.
    StpState interface{}

    // MEP facing direction. The type is CfmBagDirection.
    MepDirection interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Number of peer MEPs detected. The type is interface{} with range:
    // 0..4294967295.
    PeerMepsDetected interface{}

    // Number of peer MEPs detected with errors. The type is interface{} with
    // range: 0..4294967295.
    PeerMepsWithErrorsDetected interface{}

    // Remote defect indicated. The type is bool.
    RemoteDefect interface{}

    // Fault Notification Generation state. The type is CfmPmMepFngState.
    FaultNotificationState interface{}

    // CCM generation enabled. The type is bool.
    CcmGenerationEnabled interface{}

    // The interval between CCMs. The type is CfmBagCcmInterval.
    CcmInterval interface{}

    // Offload status of CCM processing. The type is CfmBagCcmOffload.
    CcmOffload interface{}

    // Highest-priority defect present since last FNG reset. The type is
    // CfmPmMepDefect.
    HighestDefect interface{}

    // A peer MEP RDI defect has been reported. The type is bool.
    RdiDefect interface{}

    // A peer MEP port or interface status error has been reported. The type is
    // bool.
    MacStatusDefect interface{}

    // A peer MEP CCM error has been reported. The type is bool.
    PeerMepCcmDefect interface{}

    // A CCM error has been reported. The type is bool.
    ErrorCcmDefect interface{}

    // A cross-connect CCM error has been reported. The type is bool.
    CrossConnectCcmDefect interface{}

    // Next Transaction ID to be sent in a Loopback Message. The type is
    // interface{} with range: 0..4294967295.
    NextLbmId interface{}

    // Next Transaction ID to be sent in a Linktrace Message. The type is
    // interface{} with range: 0..4294967295.
    NextLtmId interface{}

    // CoS bits the MEP will use for sent packets, if configured. The type is
    // interface{} with range: 0..255.
    Cos interface{}

    // EFD triggered on the interface. The type is bool.
    EfdTriggered interface{}

    // The local MEP is on an interface in standby mode. The type is bool.
    Standby interface{}

    // MEP is on a sub-interface in the same bridge-domain and on the same trunk
    // interface as another sub-interface. The type is bool.
    Hairpin interface{}

    // Defects present but ignored due to 'report defects' configuration. The type
    // is bool.
    DefectsIgnored interface{}

    // MEP statistics.
    Statistics Cfm_Global_LocalMeps_LocalMep_Statistics

    // MEP AIS statistics.
    AisStatistics Cfm_Global_LocalMeps_LocalMep_AisStatistics

    // Defects detected from peer MEPs.
    Defects Cfm_Global_LocalMeps_LocalMep_Defects
}

func (localMep *Cfm_Global_LocalMeps_LocalMep) GetEntityData() *types.CommonEntityData {
    localMep.EntityData.YFilter = localMep.YFilter
    localMep.EntityData.YangName = "local-mep"
    localMep.EntityData.BundleName = "cisco_ios_xr"
    localMep.EntityData.ParentYangName = "local-meps"
    localMep.EntityData.SegmentPath = "local-mep" + "[domain='" + fmt.Sprintf("%v", localMep.Domain) + "']" + "[service='" + fmt.Sprintf("%v", localMep.Service) + "']" + "[mep-id='" + fmt.Sprintf("%v", localMep.MepId) + "']" + "[interface='" + fmt.Sprintf("%v", localMep.Interface_) + "']"
    localMep.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localMep.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localMep.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localMep.EntityData.Children = make(map[string]types.YChild)
    localMep.EntityData.Children["statistics"] = types.YChild{"Statistics", &localMep.Statistics}
    localMep.EntityData.Children["ais-statistics"] = types.YChild{"AisStatistics", &localMep.AisStatistics}
    localMep.EntityData.Children["defects"] = types.YChild{"Defects", &localMep.Defects}
    localMep.EntityData.Leafs = make(map[string]types.YLeaf)
    localMep.EntityData.Leafs["domain"] = types.YLeaf{"Domain", localMep.Domain}
    localMep.EntityData.Leafs["service"] = types.YLeaf{"Service", localMep.Service}
    localMep.EntityData.Leafs["mep-id"] = types.YLeaf{"MepId", localMep.MepId}
    localMep.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", localMep.Interface_}
    localMep.EntityData.Leafs["domain-xr"] = types.YLeaf{"DomainXr", localMep.DomainXr}
    localMep.EntityData.Leafs["service-xr"] = types.YLeaf{"ServiceXr", localMep.ServiceXr}
    localMep.EntityData.Leafs["level"] = types.YLeaf{"Level", localMep.Level}
    localMep.EntityData.Leafs["mep-id-xr"] = types.YLeaf{"MepIdXr", localMep.MepIdXr}
    localMep.EntityData.Leafs["interface-xr"] = types.YLeaf{"InterfaceXr", localMep.InterfaceXr}
    localMep.EntityData.Leafs["interface-state"] = types.YLeaf{"InterfaceState", localMep.InterfaceState}
    localMep.EntityData.Leafs["interworking-state"] = types.YLeaf{"InterworkingState", localMep.InterworkingState}
    localMep.EntityData.Leafs["stp-state"] = types.YLeaf{"StpState", localMep.StpState}
    localMep.EntityData.Leafs["mep-direction"] = types.YLeaf{"MepDirection", localMep.MepDirection}
    localMep.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", localMep.MacAddress}
    localMep.EntityData.Leafs["peer-meps-detected"] = types.YLeaf{"PeerMepsDetected", localMep.PeerMepsDetected}
    localMep.EntityData.Leafs["peer-meps-with-errors-detected"] = types.YLeaf{"PeerMepsWithErrorsDetected", localMep.PeerMepsWithErrorsDetected}
    localMep.EntityData.Leafs["remote-defect"] = types.YLeaf{"RemoteDefect", localMep.RemoteDefect}
    localMep.EntityData.Leafs["fault-notification-state"] = types.YLeaf{"FaultNotificationState", localMep.FaultNotificationState}
    localMep.EntityData.Leafs["ccm-generation-enabled"] = types.YLeaf{"CcmGenerationEnabled", localMep.CcmGenerationEnabled}
    localMep.EntityData.Leafs["ccm-interval"] = types.YLeaf{"CcmInterval", localMep.CcmInterval}
    localMep.EntityData.Leafs["ccm-offload"] = types.YLeaf{"CcmOffload", localMep.CcmOffload}
    localMep.EntityData.Leafs["highest-defect"] = types.YLeaf{"HighestDefect", localMep.HighestDefect}
    localMep.EntityData.Leafs["rdi-defect"] = types.YLeaf{"RdiDefect", localMep.RdiDefect}
    localMep.EntityData.Leafs["mac-status-defect"] = types.YLeaf{"MacStatusDefect", localMep.MacStatusDefect}
    localMep.EntityData.Leafs["peer-mep-ccm-defect"] = types.YLeaf{"PeerMepCcmDefect", localMep.PeerMepCcmDefect}
    localMep.EntityData.Leafs["error-ccm-defect"] = types.YLeaf{"ErrorCcmDefect", localMep.ErrorCcmDefect}
    localMep.EntityData.Leafs["cross-connect-ccm-defect"] = types.YLeaf{"CrossConnectCcmDefect", localMep.CrossConnectCcmDefect}
    localMep.EntityData.Leafs["next-lbm-id"] = types.YLeaf{"NextLbmId", localMep.NextLbmId}
    localMep.EntityData.Leafs["next-ltm-id"] = types.YLeaf{"NextLtmId", localMep.NextLtmId}
    localMep.EntityData.Leafs["cos"] = types.YLeaf{"Cos", localMep.Cos}
    localMep.EntityData.Leafs["efd-triggered"] = types.YLeaf{"EfdTriggered", localMep.EfdTriggered}
    localMep.EntityData.Leafs["standby"] = types.YLeaf{"Standby", localMep.Standby}
    localMep.EntityData.Leafs["hairpin"] = types.YLeaf{"Hairpin", localMep.Hairpin}
    localMep.EntityData.Leafs["defects-ignored"] = types.YLeaf{"DefectsIgnored", localMep.DefectsIgnored}
    return &(localMep.EntityData)
}

// Cfm_Global_LocalMeps_LocalMep_Statistics
// MEP statistics
type Cfm_Global_LocalMeps_LocalMep_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of CCMs sent. The type is interface{} with range:
    // 0..18446744073709551615.
    CcmsSent interface{}

    // Number of CCMs received. The type is interface{} with range:
    // 0..18446744073709551615.
    CcmsReceived interface{}

    // Number of CCMs received out-of-sequence. The type is interface{} with
    // range: 0..18446744073709551615.
    CcmsOutOfSequence interface{}

    // Number of CCMs discarded because maximum MEPs limit was reached. The type
    // is interface{} with range: 0..18446744073709551615.
    CcmsDiscarded interface{}

    // Number of LBMs sent. The type is interface{} with range:
    // 0..18446744073709551615.
    LbMsSent interface{}

    // Number of LBRs sent. The type is interface{} with range:
    // 0..18446744073709551615.
    LbRsSent interface{}

    // Number of LBRs received. The type is interface{} with range:
    // 0..18446744073709551615.
    LbRsReceived interface{}

    // Number of LBRs received out-of-sequence. The type is interface{} with
    // range: 0..18446744073709551615.
    LbRsOutOfSequence interface{}

    // Number of LBRs received with non-matching user-specified data. The type is
    // interface{} with range: 0..18446744073709551615.
    LbRsBadData interface{}

    // Number of LBMs received. The type is interface{} with range:
    // 0..18446744073709551615.
    LbMsReceived interface{}

    // Number of unexpected LTRs received. The type is interface{} with range:
    // 0..18446744073709551615.
    LtRsReceivedUnexpected interface{}

    // Number of AIS messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AiSsSent interface{}

    // Number of AIS messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    AiSsReceived interface{}

    // Number of LCK messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    LcKsReceived interface{}

    // Number of DMM messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    DmMsSent interface{}

    // Number of DMM messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    DmMsReceived interface{}

    // Number of DMR messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    DmRsSent interface{}

    // Number of DMR messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    DmRsReceived interface{}

    // Number of SLM messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    SlMsSent interface{}

    // Number of SLM messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    SlMsReceived interface{}

    // Number of SLR messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    SlRsSent interface{}

    // Number of SLR messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    SlRsReceived interface{}

    // Number of LMM messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    LmMsSent interface{}

    // Number of LMM messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    LmMsReceived interface{}

    // Number of LMR messages sent. The type is interface{} with range:
    // 0..18446744073709551615.
    LmRsSent interface{}

    // Number of LMR messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    LmRsReceived interface{}

    // Number of BNM messages received. The type is interface{} with range:
    // 0..18446744073709551615.
    BnMsReceived interface{}

    // Number of BNM messages discarded. The type is interface{} with range:
    // 0..18446744073709551615.
    BnMsDiscarded interface{}
}

func (statistics *Cfm_Global_LocalMeps_LocalMep_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "local-mep"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["ccms-sent"] = types.YLeaf{"CcmsSent", statistics.CcmsSent}
    statistics.EntityData.Leafs["ccms-received"] = types.YLeaf{"CcmsReceived", statistics.CcmsReceived}
    statistics.EntityData.Leafs["ccms-out-of-sequence"] = types.YLeaf{"CcmsOutOfSequence", statistics.CcmsOutOfSequence}
    statistics.EntityData.Leafs["ccms-discarded"] = types.YLeaf{"CcmsDiscarded", statistics.CcmsDiscarded}
    statistics.EntityData.Leafs["lb-ms-sent"] = types.YLeaf{"LbMsSent", statistics.LbMsSent}
    statistics.EntityData.Leafs["lb-rs-sent"] = types.YLeaf{"LbRsSent", statistics.LbRsSent}
    statistics.EntityData.Leafs["lb-rs-received"] = types.YLeaf{"LbRsReceived", statistics.LbRsReceived}
    statistics.EntityData.Leafs["lb-rs-out-of-sequence"] = types.YLeaf{"LbRsOutOfSequence", statistics.LbRsOutOfSequence}
    statistics.EntityData.Leafs["lb-rs-bad-data"] = types.YLeaf{"LbRsBadData", statistics.LbRsBadData}
    statistics.EntityData.Leafs["lb-ms-received"] = types.YLeaf{"LbMsReceived", statistics.LbMsReceived}
    statistics.EntityData.Leafs["lt-rs-received-unexpected"] = types.YLeaf{"LtRsReceivedUnexpected", statistics.LtRsReceivedUnexpected}
    statistics.EntityData.Leafs["ai-ss-sent"] = types.YLeaf{"AiSsSent", statistics.AiSsSent}
    statistics.EntityData.Leafs["ai-ss-received"] = types.YLeaf{"AiSsReceived", statistics.AiSsReceived}
    statistics.EntityData.Leafs["lc-ks-received"] = types.YLeaf{"LcKsReceived", statistics.LcKsReceived}
    statistics.EntityData.Leafs["dm-ms-sent"] = types.YLeaf{"DmMsSent", statistics.DmMsSent}
    statistics.EntityData.Leafs["dm-ms-received"] = types.YLeaf{"DmMsReceived", statistics.DmMsReceived}
    statistics.EntityData.Leafs["dm-rs-sent"] = types.YLeaf{"DmRsSent", statistics.DmRsSent}
    statistics.EntityData.Leafs["dm-rs-received"] = types.YLeaf{"DmRsReceived", statistics.DmRsReceived}
    statistics.EntityData.Leafs["sl-ms-sent"] = types.YLeaf{"SlMsSent", statistics.SlMsSent}
    statistics.EntityData.Leafs["sl-ms-received"] = types.YLeaf{"SlMsReceived", statistics.SlMsReceived}
    statistics.EntityData.Leafs["sl-rs-sent"] = types.YLeaf{"SlRsSent", statistics.SlRsSent}
    statistics.EntityData.Leafs["sl-rs-received"] = types.YLeaf{"SlRsReceived", statistics.SlRsReceived}
    statistics.EntityData.Leafs["lm-ms-sent"] = types.YLeaf{"LmMsSent", statistics.LmMsSent}
    statistics.EntityData.Leafs["lm-ms-received"] = types.YLeaf{"LmMsReceived", statistics.LmMsReceived}
    statistics.EntityData.Leafs["lm-rs-sent"] = types.YLeaf{"LmRsSent", statistics.LmRsSent}
    statistics.EntityData.Leafs["lm-rs-received"] = types.YLeaf{"LmRsReceived", statistics.LmRsReceived}
    statistics.EntityData.Leafs["bn-ms-received"] = types.YLeaf{"BnMsReceived", statistics.BnMsReceived}
    statistics.EntityData.Leafs["bn-ms-discarded"] = types.YLeaf{"BnMsDiscarded", statistics.BnMsDiscarded}
    return &(statistics.EntityData)
}

// Cfm_Global_LocalMeps_LocalMep_AisStatistics
// MEP AIS statistics
type Cfm_Global_LocalMeps_LocalMep_AisStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AIS transmission level. The type is CfmBagMdLevel.
    Level interface{}

    // AIS transmission interval. The type is CfmBagAisInterval.
    Interval interface{}

    // Details of how AIS is being transmitted. The type is CfmPmAisTransmit.
    SendingAis interface{}

    // Details of how the signal is being received. The type is CfmPmAisReceive.
    ReceivingAis interface{}

    // The interval of the last received AIS packet. The type is
    // CfmBagAisInterval.
    LastInterval interface{}

    // Source MAC address of the last received AIS packet. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    LastMacAddress interface{}

    // Time elapsed since AIS sending started.
    SendingStart Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart

    // Time elapsed since AIS receiving started.
    ReceivingStart Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart
}

func (aisStatistics *Cfm_Global_LocalMeps_LocalMep_AisStatistics) GetEntityData() *types.CommonEntityData {
    aisStatistics.EntityData.YFilter = aisStatistics.YFilter
    aisStatistics.EntityData.YangName = "ais-statistics"
    aisStatistics.EntityData.BundleName = "cisco_ios_xr"
    aisStatistics.EntityData.ParentYangName = "local-mep"
    aisStatistics.EntityData.SegmentPath = "ais-statistics"
    aisStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aisStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aisStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aisStatistics.EntityData.Children = make(map[string]types.YChild)
    aisStatistics.EntityData.Children["sending-start"] = types.YChild{"SendingStart", &aisStatistics.SendingStart}
    aisStatistics.EntityData.Children["receiving-start"] = types.YChild{"ReceivingStart", &aisStatistics.ReceivingStart}
    aisStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    aisStatistics.EntityData.Leafs["level"] = types.YLeaf{"Level", aisStatistics.Level}
    aisStatistics.EntityData.Leafs["interval"] = types.YLeaf{"Interval", aisStatistics.Interval}
    aisStatistics.EntityData.Leafs["sending-ais"] = types.YLeaf{"SendingAis", aisStatistics.SendingAis}
    aisStatistics.EntityData.Leafs["receiving-ais"] = types.YLeaf{"ReceivingAis", aisStatistics.ReceivingAis}
    aisStatistics.EntityData.Leafs["last-interval"] = types.YLeaf{"LastInterval", aisStatistics.LastInterval}
    aisStatistics.EntityData.Leafs["last-mac-address"] = types.YLeaf{"LastMacAddress", aisStatistics.LastMacAddress}
    return &(aisStatistics.EntityData)
}

// Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart
// Time elapsed since AIS sending started
type Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (sendingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_SendingStart) GetEntityData() *types.CommonEntityData {
    sendingStart.EntityData.YFilter = sendingStart.YFilter
    sendingStart.EntityData.YangName = "sending-start"
    sendingStart.EntityData.BundleName = "cisco_ios_xr"
    sendingStart.EntityData.ParentYangName = "ais-statistics"
    sendingStart.EntityData.SegmentPath = "sending-start"
    sendingStart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendingStart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendingStart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendingStart.EntityData.Children = make(map[string]types.YChild)
    sendingStart.EntityData.Leafs = make(map[string]types.YLeaf)
    sendingStart.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", sendingStart.Seconds}
    sendingStart.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", sendingStart.Nanoseconds}
    return &(sendingStart.EntityData)
}

// Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart
// Time elapsed since AIS receiving started
type Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (receivingStart *Cfm_Global_LocalMeps_LocalMep_AisStatistics_ReceivingStart) GetEntityData() *types.CommonEntityData {
    receivingStart.EntityData.YFilter = receivingStart.YFilter
    receivingStart.EntityData.YangName = "receiving-start"
    receivingStart.EntityData.BundleName = "cisco_ios_xr"
    receivingStart.EntityData.ParentYangName = "ais-statistics"
    receivingStart.EntityData.SegmentPath = "receiving-start"
    receivingStart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivingStart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivingStart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivingStart.EntityData.Children = make(map[string]types.YChild)
    receivingStart.EntityData.Leafs = make(map[string]types.YLeaf)
    receivingStart.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", receivingStart.Seconds}
    receivingStart.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", receivingStart.Nanoseconds}
    return &(receivingStart.EntityData)
}

// Cfm_Global_LocalMeps_LocalMep_Defects
// Defects detected from peer MEPs
type Cfm_Global_LocalMeps_LocalMep_Defects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AIS or LCK received. The type is bool.
    AisReceived interface{}

    // Number of peer MEPs that have timed out. The type is interface{} with
    // range: 0..4294967295.
    PeerMepsThatTimedOut interface{}

    // Number of missing peer MEPs. The type is interface{} with range:
    // 0..4294967295.
    Missing interface{}

    // Number of missing auto cross-check MEPs. The type is interface{} with
    // range: 0..4294967295.
    AutoMissing interface{}

    // Number of unexpected peer MEPs. The type is interface{} with range:
    // 0..4294967295.
    Unexpected interface{}

    // The local port or interface is down. The type is bool.
    LocalPortStatus interface{}

    // A peer port or interface is down. The type is bool.
    PeerPortStatus interface{}

    // Defects detected from remote MEPs.
    RemoteMepsDefects Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects
}

func (defects *Cfm_Global_LocalMeps_LocalMep_Defects) GetEntityData() *types.CommonEntityData {
    defects.EntityData.YFilter = defects.YFilter
    defects.EntityData.YangName = "defects"
    defects.EntityData.BundleName = "cisco_ios_xr"
    defects.EntityData.ParentYangName = "local-mep"
    defects.EntityData.SegmentPath = "defects"
    defects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defects.EntityData.Children = make(map[string]types.YChild)
    defects.EntityData.Children["remote-meps-defects"] = types.YChild{"RemoteMepsDefects", &defects.RemoteMepsDefects}
    defects.EntityData.Leafs = make(map[string]types.YLeaf)
    defects.EntityData.Leafs["ais-received"] = types.YLeaf{"AisReceived", defects.AisReceived}
    defects.EntityData.Leafs["peer-meps-that-timed-out"] = types.YLeaf{"PeerMepsThatTimedOut", defects.PeerMepsThatTimedOut}
    defects.EntityData.Leafs["missing"] = types.YLeaf{"Missing", defects.Missing}
    defects.EntityData.Leafs["auto-missing"] = types.YLeaf{"AutoMissing", defects.AutoMissing}
    defects.EntityData.Leafs["unexpected"] = types.YLeaf{"Unexpected", defects.Unexpected}
    defects.EntityData.Leafs["local-port-status"] = types.YLeaf{"LocalPortStatus", defects.LocalPortStatus}
    defects.EntityData.Leafs["peer-port-status"] = types.YLeaf{"PeerPortStatus", defects.PeerPortStatus}
    return &(defects.EntityData)
}

// Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects
// Defects detected from remote MEPs
type Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timed out (loss threshold exceeded). The type is bool.
    LossThresholdExceeded interface{}

    // Invalid level. The type is bool.
    InvalidLevel interface{}

    // Invalid MAID. The type is bool.
    InvalidMaid interface{}

    // Invalid CCM interval. The type is bool.
    InvalidCcmInterval interface{}

    // Loop detected (our MAC address received). The type is bool.
    ReceivedOurMac interface{}

    // Configuration Error (our MEP ID received). The type is bool.
    ReceivedOurMepId interface{}

    // Remote defection indication received. The type is bool.
    ReceivedRdi interface{}
}

func (remoteMepsDefects *Cfm_Global_LocalMeps_LocalMep_Defects_RemoteMepsDefects) GetEntityData() *types.CommonEntityData {
    remoteMepsDefects.EntityData.YFilter = remoteMepsDefects.YFilter
    remoteMepsDefects.EntityData.YangName = "remote-meps-defects"
    remoteMepsDefects.EntityData.BundleName = "cisco_ios_xr"
    remoteMepsDefects.EntityData.ParentYangName = "defects"
    remoteMepsDefects.EntityData.SegmentPath = "remote-meps-defects"
    remoteMepsDefects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteMepsDefects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteMepsDefects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteMepsDefects.EntityData.Children = make(map[string]types.YChild)
    remoteMepsDefects.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteMepsDefects.EntityData.Leafs["loss-threshold-exceeded"] = types.YLeaf{"LossThresholdExceeded", remoteMepsDefects.LossThresholdExceeded}
    remoteMepsDefects.EntityData.Leafs["invalid-level"] = types.YLeaf{"InvalidLevel", remoteMepsDefects.InvalidLevel}
    remoteMepsDefects.EntityData.Leafs["invalid-maid"] = types.YLeaf{"InvalidMaid", remoteMepsDefects.InvalidMaid}
    remoteMepsDefects.EntityData.Leafs["invalid-ccm-interval"] = types.YLeaf{"InvalidCcmInterval", remoteMepsDefects.InvalidCcmInterval}
    remoteMepsDefects.EntityData.Leafs["received-our-mac"] = types.YLeaf{"ReceivedOurMac", remoteMepsDefects.ReceivedOurMac}
    remoteMepsDefects.EntityData.Leafs["received-our-mep-id"] = types.YLeaf{"ReceivedOurMepId", remoteMepsDefects.ReceivedOurMepId}
    remoteMepsDefects.EntityData.Leafs["received-rdi"] = types.YLeaf{"ReceivedRdi", remoteMepsDefects.ReceivedRdi}
    return &(remoteMepsDefects.EntityData)
}

// Cfm_Global_PeerMePv2S
// Peer MEPs table Version 2
type Cfm_Global_PeerMePv2S struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a peer MEP for a particular local MEP. The type is slice
    // of Cfm_Global_PeerMePv2S_PeerMePv2.
    PeerMePv2 []Cfm_Global_PeerMePv2S_PeerMePv2
}

func (peerMePv2S *Cfm_Global_PeerMePv2S) GetEntityData() *types.CommonEntityData {
    peerMePv2S.EntityData.YFilter = peerMePv2S.YFilter
    peerMePv2S.EntityData.YangName = "peer-me-pv2s"
    peerMePv2S.EntityData.BundleName = "cisco_ios_xr"
    peerMePv2S.EntityData.ParentYangName = "global"
    peerMePv2S.EntityData.SegmentPath = "peer-me-pv2s"
    peerMePv2S.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerMePv2S.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerMePv2S.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerMePv2S.EntityData.Children = make(map[string]types.YChild)
    peerMePv2S.EntityData.Children["peer-me-pv2"] = types.YChild{"PeerMePv2", nil}
    for i := range peerMePv2S.PeerMePv2 {
        peerMePv2S.EntityData.Children[types.GetSegmentPath(&peerMePv2S.PeerMePv2[i])] = types.YChild{"PeerMePv2", &peerMePv2S.PeerMePv2[i]}
    }
    peerMePv2S.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peerMePv2S.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2
// Information about a peer MEP for a particular
// local MEP
type Cfm_Global_PeerMePv2S_PeerMePv2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // This attribute is a key. MEP ID of Local MEP. The type is interface{} with
    // range: 1..8191.
    LocalMepId interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // This attribute is a key. MEP ID of Peer MEP. The type is interface{} with
    // range: 1..8191.
    PeerMepId interface{}

    // This attribute is a key. Peer MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    PeerMacAddress interface{}

    // Maintenance domain name. The type is string.
    DomainXr interface{}

    // Service name. The type is string.
    ServiceXr interface{}

    // Maintenance level. The type is CfmBagMdLevel.
    Level interface{}

    // MEP ID. The type is interface{} with range: 0..65535.
    MepId interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceXr interface{}

    // MEP facing direction. The type is CfmBagDirection.
    MepDirection interface{}

    // The local MEP is on an interface in standby mode. The type is bool.
    Standby interface{}

    // Peer MEP.
    PeerMep Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep
}

func (peerMePv2 *Cfm_Global_PeerMePv2S_PeerMePv2) GetEntityData() *types.CommonEntityData {
    peerMePv2.EntityData.YFilter = peerMePv2.YFilter
    peerMePv2.EntityData.YangName = "peer-me-pv2"
    peerMePv2.EntityData.BundleName = "cisco_ios_xr"
    peerMePv2.EntityData.ParentYangName = "peer-me-pv2s"
    peerMePv2.EntityData.SegmentPath = "peer-me-pv2" + "[domain='" + fmt.Sprintf("%v", peerMePv2.Domain) + "']" + "[service='" + fmt.Sprintf("%v", peerMePv2.Service) + "']" + "[local-mep-id='" + fmt.Sprintf("%v", peerMePv2.LocalMepId) + "']" + "[interface='" + fmt.Sprintf("%v", peerMePv2.Interface_) + "']" + "[peer-mep-id='" + fmt.Sprintf("%v", peerMePv2.PeerMepId) + "']" + "[peer-mac-address='" + fmt.Sprintf("%v", peerMePv2.PeerMacAddress) + "']"
    peerMePv2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerMePv2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerMePv2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerMePv2.EntityData.Children = make(map[string]types.YChild)
    peerMePv2.EntityData.Children["peer-mep"] = types.YChild{"PeerMep", &peerMePv2.PeerMep}
    peerMePv2.EntityData.Leafs = make(map[string]types.YLeaf)
    peerMePv2.EntityData.Leafs["domain"] = types.YLeaf{"Domain", peerMePv2.Domain}
    peerMePv2.EntityData.Leafs["service"] = types.YLeaf{"Service", peerMePv2.Service}
    peerMePv2.EntityData.Leafs["local-mep-id"] = types.YLeaf{"LocalMepId", peerMePv2.LocalMepId}
    peerMePv2.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", peerMePv2.Interface_}
    peerMePv2.EntityData.Leafs["peer-mep-id"] = types.YLeaf{"PeerMepId", peerMePv2.PeerMepId}
    peerMePv2.EntityData.Leafs["peer-mac-address"] = types.YLeaf{"PeerMacAddress", peerMePv2.PeerMacAddress}
    peerMePv2.EntityData.Leafs["domain-xr"] = types.YLeaf{"DomainXr", peerMePv2.DomainXr}
    peerMePv2.EntityData.Leafs["service-xr"] = types.YLeaf{"ServiceXr", peerMePv2.ServiceXr}
    peerMePv2.EntityData.Leafs["level"] = types.YLeaf{"Level", peerMePv2.Level}
    peerMePv2.EntityData.Leafs["mep-id"] = types.YLeaf{"MepId", peerMePv2.MepId}
    peerMePv2.EntityData.Leafs["interface-xr"] = types.YLeaf{"InterfaceXr", peerMePv2.InterfaceXr}
    peerMePv2.EntityData.Leafs["mep-direction"] = types.YLeaf{"MepDirection", peerMePv2.MepDirection}
    peerMePv2.EntityData.Leafs["standby"] = types.YLeaf{"Standby", peerMePv2.Standby}
    return &(peerMePv2.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep
// Peer MEP
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MEP ID. The type is interface{} with range: 0..65535.
    MepId interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Cross-check state. The type is CfmPmRmepXcState.
    CrossCheckState interface{}

    // State of the peer MEP state machine. The type is CfmPmRmepState.
    PeerMepState interface{}

    // Offload status of received CCM handling. The type is CfmBagCcmOffload.
    CcmOffload interface{}

    // Error state.
    ErrorState Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState

    // Elapsed time since peer MEP became active or timed out.
    LastUpDownTime Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime

    // Last CCM received from the peer MEP.
    LastCcmReceived Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived

    // Peer MEP statistics.
    Statistics Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics
}

func (peerMep *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep) GetEntityData() *types.CommonEntityData {
    peerMep.EntityData.YFilter = peerMep.YFilter
    peerMep.EntityData.YangName = "peer-mep"
    peerMep.EntityData.BundleName = "cisco_ios_xr"
    peerMep.EntityData.ParentYangName = "peer-me-pv2"
    peerMep.EntityData.SegmentPath = "peer-mep"
    peerMep.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerMep.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerMep.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerMep.EntityData.Children = make(map[string]types.YChild)
    peerMep.EntityData.Children["error-state"] = types.YChild{"ErrorState", &peerMep.ErrorState}
    peerMep.EntityData.Children["last-up-down-time"] = types.YChild{"LastUpDownTime", &peerMep.LastUpDownTime}
    peerMep.EntityData.Children["last-ccm-received"] = types.YChild{"LastCcmReceived", &peerMep.LastCcmReceived}
    peerMep.EntityData.Children["statistics"] = types.YChild{"Statistics", &peerMep.Statistics}
    peerMep.EntityData.Leafs = make(map[string]types.YLeaf)
    peerMep.EntityData.Leafs["mep-id"] = types.YLeaf{"MepId", peerMep.MepId}
    peerMep.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", peerMep.MacAddress}
    peerMep.EntityData.Leafs["cross-check-state"] = types.YLeaf{"CrossCheckState", peerMep.CrossCheckState}
    peerMep.EntityData.Leafs["peer-mep-state"] = types.YLeaf{"PeerMepState", peerMep.PeerMepState}
    peerMep.EntityData.Leafs["ccm-offload"] = types.YLeaf{"CcmOffload", peerMep.CcmOffload}
    return &(peerMep.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState
// Error state
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timed out (loss threshold exceeded). The type is bool.
    LossThresholdExceeded interface{}

    // Invalid level. The type is bool.
    InvalidLevel interface{}

    // Invalid MAID. The type is bool.
    InvalidMaid interface{}

    // Invalid CCM interval. The type is bool.
    InvalidCcmInterval interface{}

    // Loop detected (our MAC address received). The type is bool.
    ReceivedOurMac interface{}

    // Configuration Error (our MEP ID received). The type is bool.
    ReceivedOurMepId interface{}

    // Remote defection indication received. The type is bool.
    ReceivedRdi interface{}
}

func (errorState *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_ErrorState) GetEntityData() *types.CommonEntityData {
    errorState.EntityData.YFilter = errorState.YFilter
    errorState.EntityData.YangName = "error-state"
    errorState.EntityData.BundleName = "cisco_ios_xr"
    errorState.EntityData.ParentYangName = "peer-mep"
    errorState.EntityData.SegmentPath = "error-state"
    errorState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorState.EntityData.Children = make(map[string]types.YChild)
    errorState.EntityData.Leafs = make(map[string]types.YLeaf)
    errorState.EntityData.Leafs["loss-threshold-exceeded"] = types.YLeaf{"LossThresholdExceeded", errorState.LossThresholdExceeded}
    errorState.EntityData.Leafs["invalid-level"] = types.YLeaf{"InvalidLevel", errorState.InvalidLevel}
    errorState.EntityData.Leafs["invalid-maid"] = types.YLeaf{"InvalidMaid", errorState.InvalidMaid}
    errorState.EntityData.Leafs["invalid-ccm-interval"] = types.YLeaf{"InvalidCcmInterval", errorState.InvalidCcmInterval}
    errorState.EntityData.Leafs["received-our-mac"] = types.YLeaf{"ReceivedOurMac", errorState.ReceivedOurMac}
    errorState.EntityData.Leafs["received-our-mep-id"] = types.YLeaf{"ReceivedOurMepId", errorState.ReceivedOurMepId}
    errorState.EntityData.Leafs["received-rdi"] = types.YLeaf{"ReceivedRdi", errorState.ReceivedRdi}
    return &(errorState.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime
// Elapsed time since peer MEP became active or
// timed out
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastUpDownTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastUpDownTime) GetEntityData() *types.CommonEntityData {
    lastUpDownTime.EntityData.YFilter = lastUpDownTime.YFilter
    lastUpDownTime.EntityData.YangName = "last-up-down-time"
    lastUpDownTime.EntityData.BundleName = "cisco_ios_xr"
    lastUpDownTime.EntityData.ParentYangName = "peer-mep"
    lastUpDownTime.EntityData.SegmentPath = "last-up-down-time"
    lastUpDownTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpDownTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpDownTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpDownTime.EntityData.Children = make(map[string]types.YChild)
    lastUpDownTime.EntityData.Leafs = make(map[string]types.YLeaf)
    lastUpDownTime.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", lastUpDownTime.Seconds}
    lastUpDownTime.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", lastUpDownTime.Nanoseconds}
    return &(lastUpDownTime.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived
// Last CCM received from the peer MEP
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port status. The type is CfmPmPortStatus.
    PortStatus interface{}

    // Interface status. The type is CfmPmIntfStatus.
    InterfaceStatus interface{}

    // Additional interface status. The type is CfmPmAddlIntfStatus.
    AdditionalInterfaceStatus interface{}

    // Undecoded frame. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    RawData interface{}

    // Frame header.
    Header Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header

    // Sender ID TLV.
    SenderId Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId

    // MEP name.
    MepName Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName

    // Organizational-specific TLVs. The type is slice of
    // Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv.
    OrganizationSpecificTlv []Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv

    // Unknown TLVs. The type is slice of
    // Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv.
    UnknownTlv []Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv
}

func (lastCcmReceived *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived) GetEntityData() *types.CommonEntityData {
    lastCcmReceived.EntityData.YFilter = lastCcmReceived.YFilter
    lastCcmReceived.EntityData.YangName = "last-ccm-received"
    lastCcmReceived.EntityData.BundleName = "cisco_ios_xr"
    lastCcmReceived.EntityData.ParentYangName = "peer-mep"
    lastCcmReceived.EntityData.SegmentPath = "last-ccm-received"
    lastCcmReceived.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastCcmReceived.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastCcmReceived.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastCcmReceived.EntityData.Children = make(map[string]types.YChild)
    lastCcmReceived.EntityData.Children["header"] = types.YChild{"Header", &lastCcmReceived.Header}
    lastCcmReceived.EntityData.Children["sender-id"] = types.YChild{"SenderId", &lastCcmReceived.SenderId}
    lastCcmReceived.EntityData.Children["mep-name"] = types.YChild{"MepName", &lastCcmReceived.MepName}
    lastCcmReceived.EntityData.Children["organization-specific-tlv"] = types.YChild{"OrganizationSpecificTlv", nil}
    for i := range lastCcmReceived.OrganizationSpecificTlv {
        lastCcmReceived.EntityData.Children[types.GetSegmentPath(&lastCcmReceived.OrganizationSpecificTlv[i])] = types.YChild{"OrganizationSpecificTlv", &lastCcmReceived.OrganizationSpecificTlv[i]}
    }
    lastCcmReceived.EntityData.Children["unknown-tlv"] = types.YChild{"UnknownTlv", nil}
    for i := range lastCcmReceived.UnknownTlv {
        lastCcmReceived.EntityData.Children[types.GetSegmentPath(&lastCcmReceived.UnknownTlv[i])] = types.YChild{"UnknownTlv", &lastCcmReceived.UnknownTlv[i]}
    }
    lastCcmReceived.EntityData.Leafs = make(map[string]types.YLeaf)
    lastCcmReceived.EntityData.Leafs["port-status"] = types.YLeaf{"PortStatus", lastCcmReceived.PortStatus}
    lastCcmReceived.EntityData.Leafs["interface-status"] = types.YLeaf{"InterfaceStatus", lastCcmReceived.InterfaceStatus}
    lastCcmReceived.EntityData.Leafs["additional-interface-status"] = types.YLeaf{"AdditionalInterfaceStatus", lastCcmReceived.AdditionalInterfaceStatus}
    lastCcmReceived.EntityData.Leafs["raw-data"] = types.YLeaf{"RawData", lastCcmReceived.RawData}
    return &(lastCcmReceived.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header
// Frame header
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MD level. The type is CfmBagMdLevel.
    Level interface{}

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // CCM interval. The type is CfmBagCcmInterval.
    Interval interface{}

    // Remote defect indicated. The type is bool.
    Rdi interface{}

    // CCM sequence number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // MEP ID. The type is interface{} with range: 0..65535.
    MepId interface{}

    // MDID Format. The type is interface{} with range: 0..255.
    MdidFormat interface{}

    // Short MA Name format. The type is interface{} with range: 0..255.
    ShortMaNameFormat interface{}

    // MDID.
    Mdid Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid

    // Short MA Name.
    ShortMaName Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName
}

func (header *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "last-ccm-received"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Children["mdid"] = types.YChild{"Mdid", &header.Mdid}
    header.EntityData.Children["short-ma-name"] = types.YChild{"ShortMaName", &header.ShortMaName}
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["level"] = types.YLeaf{"Level", header.Level}
    header.EntityData.Leafs["version"] = types.YLeaf{"Version", header.Version}
    header.EntityData.Leafs["interval"] = types.YLeaf{"Interval", header.Interval}
    header.EntityData.Leafs["rdi"] = types.YLeaf{"Rdi", header.Rdi}
    header.EntityData.Leafs["sequence-number"] = types.YLeaf{"SequenceNumber", header.SequenceNumber}
    header.EntityData.Leafs["mep-id"] = types.YLeaf{"MepId", header.MepId}
    header.EntityData.Leafs["mdid-format"] = types.YLeaf{"MdidFormat", header.MdidFormat}
    header.EntityData.Leafs["short-ma-name-format"] = types.YLeaf{"ShortMaNameFormat", header.ShortMaNameFormat}
    return &(header.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid
// MDID
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MDIDFormatValue. The type is CfmBagMdidFmt.
    MdidFormatValue interface{}

    // DNS-like name. The type is string.
    DnsLikeName interface{}

    // String name. The type is string.
    StringName interface{}

    // Hex data. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    MdidData interface{}

    // MAC address name.
    MacName Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName
}

func (mdid *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid) GetEntityData() *types.CommonEntityData {
    mdid.EntityData.YFilter = mdid.YFilter
    mdid.EntityData.YangName = "mdid"
    mdid.EntityData.BundleName = "cisco_ios_xr"
    mdid.EntityData.ParentYangName = "header"
    mdid.EntityData.SegmentPath = "mdid"
    mdid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mdid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mdid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mdid.EntityData.Children = make(map[string]types.YChild)
    mdid.EntityData.Children["mac-name"] = types.YChild{"MacName", &mdid.MacName}
    mdid.EntityData.Leafs = make(map[string]types.YLeaf)
    mdid.EntityData.Leafs["mdid-format-value"] = types.YLeaf{"MdidFormatValue", mdid.MdidFormatValue}
    mdid.EntityData.Leafs["dns-like-name"] = types.YLeaf{"DnsLikeName", mdid.DnsLikeName}
    mdid.EntityData.Leafs["string-name"] = types.YLeaf{"StringName", mdid.StringName}
    mdid.EntityData.Leafs["mdid-data"] = types.YLeaf{"MdidData", mdid.MdidData}
    return &(mdid.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName
// MAC address name
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Integer. The type is interface{} with range: 0..65535.
    Integer interface{}
}

func (macName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_Mdid_MacName) GetEntityData() *types.CommonEntityData {
    macName.EntityData.YFilter = macName.YFilter
    macName.EntityData.YangName = "mac-name"
    macName.EntityData.BundleName = "cisco_ios_xr"
    macName.EntityData.ParentYangName = "mdid"
    macName.EntityData.SegmentPath = "mac-name"
    macName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macName.EntityData.Children = make(map[string]types.YChild)
    macName.EntityData.Leafs = make(map[string]types.YLeaf)
    macName.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", macName.MacAddress}
    macName.EntityData.Leafs["integer"] = types.YLeaf{"Integer", macName.Integer}
    return &(macName.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName
// Short MA Name
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ShortMANameFormatValue. The type is CfmBagSmanFmt.
    ShortMaNameFormatValue interface{}

    // VLAN ID name. The type is interface{} with range: 0..65535.
    VlanIdName interface{}

    // String name. The type is string.
    StringName interface{}

    // Unsigned integer name. The type is interface{} with range: 0..65535.
    IntegerName interface{}

    // ICC-based format. The type is string.
    IccBased interface{}

    // Hex data. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ShortMaNameData interface{}

    // VPN ID name.
    VpnIdName Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName
}

func (shortMaName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName) GetEntityData() *types.CommonEntityData {
    shortMaName.EntityData.YFilter = shortMaName.YFilter
    shortMaName.EntityData.YangName = "short-ma-name"
    shortMaName.EntityData.BundleName = "cisco_ios_xr"
    shortMaName.EntityData.ParentYangName = "header"
    shortMaName.EntityData.SegmentPath = "short-ma-name"
    shortMaName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shortMaName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shortMaName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shortMaName.EntityData.Children = make(map[string]types.YChild)
    shortMaName.EntityData.Children["vpn-id-name"] = types.YChild{"VpnIdName", &shortMaName.VpnIdName}
    shortMaName.EntityData.Leafs = make(map[string]types.YLeaf)
    shortMaName.EntityData.Leafs["short-ma-name-format-value"] = types.YLeaf{"ShortMaNameFormatValue", shortMaName.ShortMaNameFormatValue}
    shortMaName.EntityData.Leafs["vlan-id-name"] = types.YLeaf{"VlanIdName", shortMaName.VlanIdName}
    shortMaName.EntityData.Leafs["string-name"] = types.YLeaf{"StringName", shortMaName.StringName}
    shortMaName.EntityData.Leafs["integer-name"] = types.YLeaf{"IntegerName", shortMaName.IntegerName}
    shortMaName.EntityData.Leafs["icc-based"] = types.YLeaf{"IccBased", shortMaName.IccBased}
    shortMaName.EntityData.Leafs["short-ma-name-data"] = types.YLeaf{"ShortMaNameData", shortMaName.ShortMaNameData}
    return &(shortMaName.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName
// VPN ID name
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPN authority organizationally-unique ID. The type is interface{} with
    // range: 0..4294967295.
    Oui interface{}

    // VPN index. The type is interface{} with range: 0..4294967295.
    Index interface{}
}

func (vpnIdName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_Header_ShortMaName_VpnIdName) GetEntityData() *types.CommonEntityData {
    vpnIdName.EntityData.YFilter = vpnIdName.YFilter
    vpnIdName.EntityData.YangName = "vpn-id-name"
    vpnIdName.EntityData.BundleName = "cisco_ios_xr"
    vpnIdName.EntityData.ParentYangName = "short-ma-name"
    vpnIdName.EntityData.SegmentPath = "vpn-id-name"
    vpnIdName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpnIdName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpnIdName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpnIdName.EntityData.Children = make(map[string]types.YChild)
    vpnIdName.EntityData.Leafs = make(map[string]types.YLeaf)
    vpnIdName.EntityData.Leafs["oui"] = types.YLeaf{"Oui", vpnIdName.Oui}
    vpnIdName.EntityData.Leafs["index"] = types.YLeaf{"Index", vpnIdName.Index}
    return &(vpnIdName.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId
// Sender ID TLV
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Management address domain. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ManagementAddressDomain interface{}

    // Management address. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ManagementAddress interface{}

    // Chassis ID.
    ChassisId Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId
}

func (senderId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId) GetEntityData() *types.CommonEntityData {
    senderId.EntityData.YFilter = senderId.YFilter
    senderId.EntityData.YangName = "sender-id"
    senderId.EntityData.BundleName = "cisco_ios_xr"
    senderId.EntityData.ParentYangName = "last-ccm-received"
    senderId.EntityData.SegmentPath = "sender-id"
    senderId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    senderId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    senderId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    senderId.EntityData.Children = make(map[string]types.YChild)
    senderId.EntityData.Children["chassis-id"] = types.YChild{"ChassisId", &senderId.ChassisId}
    senderId.EntityData.Leafs = make(map[string]types.YLeaf)
    senderId.EntityData.Leafs["management-address-domain"] = types.YLeaf{"ManagementAddressDomain", senderId.ManagementAddressDomain}
    senderId.EntityData.Leafs["management-address"] = types.YLeaf{"ManagementAddress", senderId.ManagementAddress}
    return &(senderId.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId
// Chassis ID
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Chassis ID Type. The type is CfmPmChassisIdFmt.
    ChassisIdType interface{}

    // Chassis ID Type. The type is interface{} with range: 0..255.
    ChassisIdTypeValue interface{}

    // Chassis ID (Deprecated). The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ChassisId interface{}

    // Chassis ID (Current).
    ChassisIdValue Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue
}

func (chassisId *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId) GetEntityData() *types.CommonEntityData {
    chassisId.EntityData.YFilter = chassisId.YFilter
    chassisId.EntityData.YangName = "chassis-id"
    chassisId.EntityData.BundleName = "cisco_ios_xr"
    chassisId.EntityData.ParentYangName = "sender-id"
    chassisId.EntityData.SegmentPath = "chassis-id"
    chassisId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisId.EntityData.Children = make(map[string]types.YChild)
    chassisId.EntityData.Children["chassis-id-value"] = types.YChild{"ChassisIdValue", &chassisId.ChassisIdValue}
    chassisId.EntityData.Leafs = make(map[string]types.YLeaf)
    chassisId.EntityData.Leafs["chassis-id-type"] = types.YLeaf{"ChassisIdType", chassisId.ChassisIdType}
    chassisId.EntityData.Leafs["chassis-id-type-value"] = types.YLeaf{"ChassisIdTypeValue", chassisId.ChassisIdTypeValue}
    chassisId.EntityData.Leafs["chassis-id"] = types.YLeaf{"ChassisId", chassisId.ChassisId}
    return &(chassisId.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue
// Chassis ID (Current)
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ChassisIDFormat. The type is CfmPmIdFmt.
    ChassisIdFormat interface{}

    // Chassis ID String. The type is string.
    ChassisIdString interface{}

    // Chassis ID MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    ChassisIdMac interface{}

    // Raw Chassis ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ChassisIdRaw interface{}
}

func (chassisIdValue *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_SenderId_ChassisId_ChassisIdValue) GetEntityData() *types.CommonEntityData {
    chassisIdValue.EntityData.YFilter = chassisIdValue.YFilter
    chassisIdValue.EntityData.YangName = "chassis-id-value"
    chassisIdValue.EntityData.BundleName = "cisco_ios_xr"
    chassisIdValue.EntityData.ParentYangName = "chassis-id"
    chassisIdValue.EntityData.SegmentPath = "chassis-id-value"
    chassisIdValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisIdValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisIdValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisIdValue.EntityData.Children = make(map[string]types.YChild)
    chassisIdValue.EntityData.Leafs = make(map[string]types.YLeaf)
    chassisIdValue.EntityData.Leafs["chassis-id-format"] = types.YLeaf{"ChassisIdFormat", chassisIdValue.ChassisIdFormat}
    chassisIdValue.EntityData.Leafs["chassis-id-string"] = types.YLeaf{"ChassisIdString", chassisIdValue.ChassisIdString}
    chassisIdValue.EntityData.Leafs["chassis-id-mac"] = types.YLeaf{"ChassisIdMac", chassisIdValue.ChassisIdMac}
    chassisIdValue.EntityData.Leafs["chassis-id-raw"] = types.YLeaf{"ChassisIdRaw", chassisIdValue.ChassisIdRaw}
    return &(chassisIdValue.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName
// MEP name
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MEP name. The type is string.
    Name interface{}
}

func (mepName *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_MepName) GetEntityData() *types.CommonEntityData {
    mepName.EntityData.YFilter = mepName.YFilter
    mepName.EntityData.YangName = "mep-name"
    mepName.EntityData.BundleName = "cisco_ios_xr"
    mepName.EntityData.ParentYangName = "last-ccm-received"
    mepName.EntityData.SegmentPath = "mep-name"
    mepName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mepName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mepName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mepName.EntityData.Children = make(map[string]types.YChild)
    mepName.EntityData.Leafs = make(map[string]types.YLeaf)
    mepName.EntityData.Leafs["name"] = types.YLeaf{"Name", mepName.Name}
    return &(mepName.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv
// Organizational-specific TLVs
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Organizationally-unique ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Oui interface{}

    // Subtype of TLV. The type is interface{} with range: 0..255.
    Subtype interface{}

    // Binary data in TLV. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (organizationSpecificTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_OrganizationSpecificTlv) GetEntityData() *types.CommonEntityData {
    organizationSpecificTlv.EntityData.YFilter = organizationSpecificTlv.YFilter
    organizationSpecificTlv.EntityData.YangName = "organization-specific-tlv"
    organizationSpecificTlv.EntityData.BundleName = "cisco_ios_xr"
    organizationSpecificTlv.EntityData.ParentYangName = "last-ccm-received"
    organizationSpecificTlv.EntityData.SegmentPath = "organization-specific-tlv"
    organizationSpecificTlv.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    organizationSpecificTlv.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    organizationSpecificTlv.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    organizationSpecificTlv.EntityData.Children = make(map[string]types.YChild)
    organizationSpecificTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    organizationSpecificTlv.EntityData.Leafs["oui"] = types.YLeaf{"Oui", organizationSpecificTlv.Oui}
    organizationSpecificTlv.EntityData.Leafs["subtype"] = types.YLeaf{"Subtype", organizationSpecificTlv.Subtype}
    organizationSpecificTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", organizationSpecificTlv.Value}
    return &(organizationSpecificTlv.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv
// Unknown TLVs
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type code of TLV. The type is interface{} with range: 0..255.
    Typecode interface{}

    // Binary data in TLV. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Value interface{}
}

func (unknownTlv *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_LastCcmReceived_UnknownTlv) GetEntityData() *types.CommonEntityData {
    unknownTlv.EntityData.YFilter = unknownTlv.YFilter
    unknownTlv.EntityData.YangName = "unknown-tlv"
    unknownTlv.EntityData.BundleName = "cisco_ios_xr"
    unknownTlv.EntityData.ParentYangName = "last-ccm-received"
    unknownTlv.EntityData.SegmentPath = "unknown-tlv"
    unknownTlv.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownTlv.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownTlv.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownTlv.EntityData.Children = make(map[string]types.YChild)
    unknownTlv.EntityData.Leafs = make(map[string]types.YLeaf)
    unknownTlv.EntityData.Leafs["typecode"] = types.YLeaf{"Typecode", unknownTlv.Typecode}
    unknownTlv.EntityData.Leafs["value"] = types.YLeaf{"Value", unknownTlv.Value}
    return &(unknownTlv.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics
// Peer MEP statistics
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of CCMs received. The type is interface{} with range:
    // 0..18446744073709551615.
    CcmsReceived interface{}

    // Number of CCMs received with an invalid level. The type is interface{} with
    // range: 0..18446744073709551615.
    CcmsWrongLevel interface{}

    // Number of CCMs received with an invalid MAID. The type is interface{} with
    // range: 0..18446744073709551615.
    CcmsInvalidMaid interface{}

    // Number of CCMs received with an invalid interval. The type is interface{}
    // with range: 0..18446744073709551615.
    CcmsInvalidInterval interface{}

    // Number of CCMs received with an invalid source MAC address. The type is
    // interface{} with range: 0..18446744073709551615.
    CcmsInvalidSourceMacAddress interface{}

    // Number of CCMs received with our MEP ID. The type is interface{} with
    // range: 0..18446744073709551615.
    CcmsOurMepId interface{}

    // Number of CCMs received with the Remote Defect Indication bit set. The type
    // is interface{} with range: 0..18446744073709551615.
    CcmsRdi interface{}

    // Number of CCMs received out-of-sequence. The type is interface{} with
    // range: 0..18446744073709551615.
    CcmsOutOfSequence interface{}

    // Sequence number of last CCM received. The type is interface{} with range:
    // 0..4294967295.
    LastCcmSequenceNumber interface{}

    // Elapsed time since last CCM received.
    LastCcmReceivedTime Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime
}

func (statistics *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "peer-mep"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["last-ccm-received-time"] = types.YChild{"LastCcmReceivedTime", &statistics.LastCcmReceivedTime}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["ccms-received"] = types.YLeaf{"CcmsReceived", statistics.CcmsReceived}
    statistics.EntityData.Leafs["ccms-wrong-level"] = types.YLeaf{"CcmsWrongLevel", statistics.CcmsWrongLevel}
    statistics.EntityData.Leafs["ccms-invalid-maid"] = types.YLeaf{"CcmsInvalidMaid", statistics.CcmsInvalidMaid}
    statistics.EntityData.Leafs["ccms-invalid-interval"] = types.YLeaf{"CcmsInvalidInterval", statistics.CcmsInvalidInterval}
    statistics.EntityData.Leafs["ccms-invalid-source-mac-address"] = types.YLeaf{"CcmsInvalidSourceMacAddress", statistics.CcmsInvalidSourceMacAddress}
    statistics.EntityData.Leafs["ccms-our-mep-id"] = types.YLeaf{"CcmsOurMepId", statistics.CcmsOurMepId}
    statistics.EntityData.Leafs["ccms-rdi"] = types.YLeaf{"CcmsRdi", statistics.CcmsRdi}
    statistics.EntityData.Leafs["ccms-out-of-sequence"] = types.YLeaf{"CcmsOutOfSequence", statistics.CcmsOutOfSequence}
    statistics.EntityData.Leafs["last-ccm-sequence-number"] = types.YLeaf{"LastCcmSequenceNumber", statistics.LastCcmSequenceNumber}
    return &(statistics.EntityData)
}

// Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime
// Elapsed time since last CCM received
type Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastCcmReceivedTime *Cfm_Global_PeerMePv2S_PeerMePv2_PeerMep_Statistics_LastCcmReceivedTime) GetEntityData() *types.CommonEntityData {
    lastCcmReceivedTime.EntityData.YFilter = lastCcmReceivedTime.YFilter
    lastCcmReceivedTime.EntityData.YangName = "last-ccm-received-time"
    lastCcmReceivedTime.EntityData.BundleName = "cisco_ios_xr"
    lastCcmReceivedTime.EntityData.ParentYangName = "statistics"
    lastCcmReceivedTime.EntityData.SegmentPath = "last-ccm-received-time"
    lastCcmReceivedTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastCcmReceivedTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastCcmReceivedTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastCcmReceivedTime.EntityData.Children = make(map[string]types.YChild)
    lastCcmReceivedTime.EntityData.Leafs = make(map[string]types.YLeaf)
    lastCcmReceivedTime.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", lastCcmReceivedTime.Seconds}
    lastCcmReceivedTime.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", lastCcmReceivedTime.Nanoseconds}
    return &(lastCcmReceivedTime.EntityData)
}

