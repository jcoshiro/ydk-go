// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-ipsla package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipsla: IPSLA operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package man_ipsla_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package man_ipsla_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-ipsla-oper ipsla}", reflect.TypeOf(Ipsla{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-ipsla-oper:ipsla", reflect.TypeOf(Ipsla{}))
}

// OpTypeEnum represents IPSLA Operation Types
type OpTypeEnum string

const (
    // icmp echo
    OpTypeEnum_icmp_echo OpTypeEnum = "icmp-echo"

    // icmp path jitter
    OpTypeEnum_icmp_path_jitter OpTypeEnum = "icmp-path-jitter"

    // icmp path echo
    OpTypeEnum_icmp_path_echo OpTypeEnum = "icmp-path-echo"

    // udp jitter
    OpTypeEnum_udp_jitter OpTypeEnum = "udp-jitter"

    // udp echo
    OpTypeEnum_udp_echo OpTypeEnum = "udp-echo"

    // mpls lsp ping
    OpTypeEnum_mpls_lsp_ping OpTypeEnum = "mpls-lsp-ping"

    // mpls lsp trace
    OpTypeEnum_mpls_lsp_trace OpTypeEnum = "mpls-lsp-trace"

    // mpls lsp group
    OpTypeEnum_mpls_lsp_group OpTypeEnum = "mpls-lsp-group"
)

// IpslaRetCode represents Ipsla ret code
type IpslaRetCode string

const (
    // ipsla ret code unknown
    IpslaRetCode_ipsla_ret_code_unknown IpslaRetCode = "ipsla-ret-code-unknown"

    // ipsla ret code ok
    IpslaRetCode_ipsla_ret_code_ok IpslaRetCode = "ipsla-ret-code-ok"

    // ipsla ret code disconnect
    IpslaRetCode_ipsla_ret_code_disconnect IpslaRetCode = "ipsla-ret-code-disconnect"

    // ipsla ret code over threshold
    IpslaRetCode_ipsla_ret_code_over_threshold IpslaRetCode = "ipsla-ret-code-over-threshold"

    // ipsla ret code timeout
    IpslaRetCode_ipsla_ret_code_timeout IpslaRetCode = "ipsla-ret-code-timeout"

    // ipsla ret code busy
    IpslaRetCode_ipsla_ret_code_busy IpslaRetCode = "ipsla-ret-code-busy"

    // ipsla ret code no connection
    IpslaRetCode_ipsla_ret_code_no_connection IpslaRetCode = "ipsla-ret-code-no-connection"

    // ipsla ret code dropped
    IpslaRetCode_ipsla_ret_code_dropped IpslaRetCode = "ipsla-ret-code-dropped"

    // ipsla ret code sequence error
    IpslaRetCode_ipsla_ret_code_sequence_error IpslaRetCode = "ipsla-ret-code-sequence-error"

    // ipsla ret code verify error
    IpslaRetCode_ipsla_ret_code_verify_error IpslaRetCode = "ipsla-ret-code-verify-error"

    // ipsla ret code application specific
    IpslaRetCode_ipsla_ret_code_application_specific IpslaRetCode = "ipsla-ret-code-application-specific"

    // ipsla ret code dns server timeout
    IpslaRetCode_ipsla_ret_code_dns_server_timeout IpslaRetCode = "ipsla-ret-code-dns-server-timeout"

    // ipsla ret code tcp connect timeout
    IpslaRetCode_ipsla_ret_code_tcp_connect_timeout IpslaRetCode = "ipsla-ret-code-tcp-connect-timeout"

    // ipsla ret code http transaction timeout
    IpslaRetCode_ipsla_ret_code_http_transaction_timeout IpslaRetCode = "ipsla-ret-code-http-transaction-timeout"

    // ipsla ret code dns query error
    IpslaRetCode_ipsla_ret_code_dns_query_error IpslaRetCode = "ipsla-ret-code-dns-query-error"

    // ipsla ret code http error
    IpslaRetCode_ipsla_ret_code_http_error IpslaRetCode = "ipsla-ret-code-http-error"

    // ipsla ret code internal error
    IpslaRetCode_ipsla_ret_code_internal_error IpslaRetCode = "ipsla-ret-code-internal-error"

    // ipsla ret code mpls lsp echo tx error
    IpslaRetCode_ipsla_ret_code_mpls_lsp_echo_tx_error IpslaRetCode = "ipsla-ret-code-mpls-lsp-echo-tx-error"

    // ipsla ret code mpls lsp unreachable
    IpslaRetCode_ipsla_ret_code_mpls_lsp_unreachable IpslaRetCode = "ipsla-ret-code-mpls-lsp-unreachable"

    // ipsla ret code mpls lsp malformed request
    IpslaRetCode_ipsla_ret_code_mpls_lsp_malformed_request IpslaRetCode = "ipsla-ret-code-mpls-lsp-malformed-request"

    // ipsla ret code mpls lsp reachable but not fec
    IpslaRetCode_ipsla_ret_code_mpls_lsp_reachable_but_not_fec IpslaRetCode = "ipsla-ret-code-mpls-lsp-reachable-but-not-fec"

    // ipsla ret code mpls lsp ds map mismatch
    IpslaRetCode_ipsla_ret_code_mpls_lsp_ds_map_mismatch IpslaRetCode = "ipsla-ret-code-mpls-lsp-ds-map-mismatch"

    // ipsla ret code mpls lsp duplicate
    IpslaRetCode_ipsla_ret_code_mpls_lsp_duplicate IpslaRetCode = "ipsla-ret-code-mpls-lsp-duplicate"

    // ipsla ret code failure
    IpslaRetCode_ipsla_ret_code_failure IpslaRetCode = "ipsla-ret-code-failure"

    // ipsla ret code malloc failure
    IpslaRetCode_ipsla_ret_code_malloc_failure IpslaRetCode = "ipsla-ret-code-malloc-failure"

    // ipsla ret code sock open error
    IpslaRetCode_ipsla_ret_code_sock_open_error IpslaRetCode = "ipsla-ret-code-sock-open-error"

    // ipsla ret code sock bind error
    IpslaRetCode_ipsla_ret_code_sock_bind_error IpslaRetCode = "ipsla-ret-code-sock-bind-error"

    // ipsla ret code sock send error
    IpslaRetCode_ipsla_ret_code_sock_send_error IpslaRetCode = "ipsla-ret-code-sock-send-error"

    // ipsla ret code sock recv error
    IpslaRetCode_ipsla_ret_code_sock_recv_error IpslaRetCode = "ipsla-ret-code-sock-recv-error"

    // ipsla ret code sock connect error
    IpslaRetCode_ipsla_ret_code_sock_connect_error IpslaRetCode = "ipsla-ret-code-sock-connect-error"

    // ipsla ret code sock set option error
    IpslaRetCode_ipsla_ret_code_sock_set_option_error IpslaRetCode = "ipsla-ret-code-sock-set-option-error"

    // ipsla ret code sock attach error
    IpslaRetCode_ipsla_ret_code_sock_attach_error IpslaRetCode = "ipsla-ret-code-sock-attach-error"

    // ipsla ret code ctrl msg error
    IpslaRetCode_ipsla_ret_code_ctrl_msg_error IpslaRetCode = "ipsla-ret-code-ctrl-msg-error"

    // ipsla ret code no key chain
    IpslaRetCode_ipsla_ret_code_no_key_chain IpslaRetCode = "ipsla-ret-code-no-key-chain"

    // ipsla ret code key chain lib fail
    IpslaRetCode_ipsla_ret_code_key_chain_lib_fail IpslaRetCode = "ipsla-ret-code-key-chain-lib-fail"

    // ipsla ret code no key id
    IpslaRetCode_ipsla_ret_code_no_key_id IpslaRetCode = "ipsla-ret-code-no-key-id"

    // ipsla ret code invalid key id
    IpslaRetCode_ipsla_ret_code_invalid_key_id IpslaRetCode = "ipsla-ret-code-invalid-key-id"

    // ipsla ret code entry exist
    IpslaRetCode_ipsla_ret_code_entry_exist IpslaRetCode = "ipsla-ret-code-entry-exist"

    // ipsla ret code entry not found
    IpslaRetCode_ipsla_ret_code_entry_not_found IpslaRetCode = "ipsla-ret-code-entry-not-found"

    // ipsla ret code hop over max
    IpslaRetCode_ipsla_ret_code_hop_over_max IpslaRetCode = "ipsla-ret-code-hop-over-max"

    // ipsla ret code hop dup address
    IpslaRetCode_ipsla_ret_code_hop_dup_address IpslaRetCode = "ipsla-ret-code-hop-dup-address"

    // ipsla ret code vrf name error
    IpslaRetCode_ipsla_ret_code_vrf_name_error IpslaRetCode = "ipsla-ret-code-vrf-name-error"

    // ipsla ret code resp failure
    IpslaRetCode_ipsla_ret_code_resp_failure IpslaRetCode = "ipsla-ret-code-resp-failure"

    // ipsla ret code auth failure
    IpslaRetCode_ipsla_ret_code_auth_failure IpslaRetCode = "ipsla-ret-code-auth-failure"

    // ipsla ret code format failure
    IpslaRetCode_ipsla_ret_code_format_failure IpslaRetCode = "ipsla-ret-code-format-failure"

    // ipsla ret code port in use
    IpslaRetCode_ipsla_ret_code_port_in_use IpslaRetCode = "ipsla-ret-code-port-in-use"

    // ipsla ret code no route
    IpslaRetCode_ipsla_ret_code_no_route IpslaRetCode = "ipsla-ret-code-no-route"

    // ipsla ret code pending
    IpslaRetCode_ipsla_ret_code_pending IpslaRetCode = "ipsla-ret-code-pending"

    // ipsla ret code invalid address
    IpslaRetCode_ipsla_ret_code_invalid_address IpslaRetCode = "ipsla-ret-code-invalid-address"

    // ipsla ret code max
    IpslaRetCode_ipsla_ret_code_max IpslaRetCode = "ipsla-ret-code-max"
)

// IpslaMplsLpdDiscoveryModeEnum represents Ipsla mpls lpd discovery mode enum
type IpslaMplsLpdDiscoveryModeEnum string

const (
    // ipsla mpls lpd unknown
    IpslaMplsLpdDiscoveryModeEnum_ipsla_mpls_lpd_unknown IpslaMplsLpdDiscoveryModeEnum = "ipsla-mpls-lpd-unknown"

    // ipsla mpls lpd initial running
    IpslaMplsLpdDiscoveryModeEnum_ipsla_mpls_lpd_initial_running IpslaMplsLpdDiscoveryModeEnum = "ipsla-mpls-lpd-initial-running"

    // ipsla mpls lpd initial complete
    IpslaMplsLpdDiscoveryModeEnum_ipsla_mpls_lpd_initial_complete IpslaMplsLpdDiscoveryModeEnum = "ipsla-mpls-lpd-initial-complete"

    // ipsla mpls lpd rediscovery running
    IpslaMplsLpdDiscoveryModeEnum_ipsla_mpls_lpd_rediscovery_running IpslaMplsLpdDiscoveryModeEnum = "ipsla-mpls-lpd-rediscovery-running"

    // ipsla mpls lpd rediscovery complete
    IpslaMplsLpdDiscoveryModeEnum_ipsla_mpls_lpd_rediscovery_complete IpslaMplsLpdDiscoveryModeEnum = "ipsla-mpls-lpd-rediscovery-complete"
)

// IpslaMplsLpdPathDiscoveryStatus represents Ipsla mpls lpd path discovery status
type IpslaMplsLpdPathDiscoveryStatus string

const (
    // ipsla mpls lpd path discovery unknown
    IpslaMplsLpdPathDiscoveryStatus_ipsla_mpls_lpd_path_discovery_unknown IpslaMplsLpdPathDiscoveryStatus = "ipsla-mpls-lpd-path-discovery-unknown"

    // ipsla mpls lpd path discovery ok
    IpslaMplsLpdPathDiscoveryStatus_ipsla_mpls_lpd_path_discovery_ok IpslaMplsLpdPathDiscoveryStatus = "ipsla-mpls-lpd-path-discovery-ok"

    // ipsla mpls lpd path discovery broken
    IpslaMplsLpdPathDiscoveryStatus_ipsla_mpls_lpd_path_discovery_broken IpslaMplsLpdPathDiscoveryStatus = "ipsla-mpls-lpd-path-discovery-broken"

    // ipsla mpls lpd path discovery unexplorable
    IpslaMplsLpdPathDiscoveryStatus_ipsla_mpls_lpd_path_discovery_unexplorable IpslaMplsLpdPathDiscoveryStatus = "ipsla-mpls-lpd-path-discovery-unexplorable"
)

// IpslaMplsLpdRetCode represents Ipsla mpls lpd ret code
type IpslaMplsLpdRetCode string

const (
    // ipsla mpls lpd ret code unknown
    IpslaMplsLpdRetCode_ipsla_mpls_lpd_ret_code_unknown IpslaMplsLpdRetCode = "ipsla-mpls-lpd-ret-code-unknown"

    // ipsla mpls lpd ret code no path
    IpslaMplsLpdRetCode_ipsla_mpls_lpd_ret_code_no_path IpslaMplsLpdRetCode = "ipsla-mpls-lpd-ret-code-no-path"

    // ipsla mpls lpd ret code all path broken
    IpslaMplsLpdRetCode_ipsla_mpls_lpd_ret_code_all_path_broken IpslaMplsLpdRetCode = "ipsla-mpls-lpd-ret-code-all-path-broken"

    // ipsla mpls lpd ret code all path unexplorable
    IpslaMplsLpdRetCode_ipsla_mpls_lpd_ret_code_all_path_unexplorable IpslaMplsLpdRetCode = "ipsla-mpls-lpd-ret-code-all-path-unexplorable"

    // ipsla mpls lpd ret code all path broken or
    // unexplorable
    IpslaMplsLpdRetCode_ipsla_mpls_lpd_ret_code_all_path_broken_or_unexplorable IpslaMplsLpdRetCode = "ipsla-mpls-lpd-ret-code-all-path-broken-or-unexplorable"

    // ipsla mpls lpd ret code timeout
    IpslaMplsLpdRetCode_ipsla_mpls_lpd_ret_code_timeout IpslaMplsLpdRetCode = "ipsla-mpls-lpd-ret-code-timeout"

    // ipsla mpls lpd ret code error
    IpslaMplsLpdRetCode_ipsla_mpls_lpd_ret_code_error IpslaMplsLpdRetCode = "ipsla-mpls-lpd-ret-code-error"

    // ipsla mpls lpd ret code ok
    IpslaMplsLpdRetCode_ipsla_mpls_lpd_ret_code_ok IpslaMplsLpdRetCode = "ipsla-mpls-lpd-ret-code-ok"
)

// IpslaTargetTypeEnum represents IPSLA Target Types
type IpslaTargetTypeEnum string

const (
    // IPv4 address
    IpslaTargetTypeEnum_ipv4_address_target_type IpslaTargetTypeEnum = "ipv4-address-target-type"

    // IPv4 prefix
    IpslaTargetTypeEnum_ipv4_prefix_target_type IpslaTargetTypeEnum = "ipv4-prefix-target-type"

    // Tunnel ID
    IpslaTargetTypeEnum_tunnel_id_target_type IpslaTargetTypeEnum = "tunnel-id-target-type"

    // IPv4 pseudowire
    IpslaTargetTypeEnum_ipv4_pseudowire_target_type IpslaTargetTypeEnum = "ipv4-pseudowire-target-type"

    // IPv6 address
    IpslaTargetTypeEnum_ipv6_address_target_type IpslaTargetTypeEnum = "ipv6-address-target-type"
)

// IpslaOperStateEnum represents Ipsla oper state enum
type IpslaOperStateEnum string

const (
    // ipsla oper state inactive
    IpslaOperStateEnum_ipsla_oper_state_inactive IpslaOperStateEnum = "ipsla-oper-state-inactive"

    // ipsla oper state pending
    IpslaOperStateEnum_ipsla_oper_state_pending IpslaOperStateEnum = "ipsla-oper-state-pending"

    // ipsla oper state active
    IpslaOperStateEnum_ipsla_oper_state_active IpslaOperStateEnum = "ipsla-oper-state-active"
)

// IpslaMplsAddDeleteEnum represents Ipsla mpls add delete enum
type IpslaMplsAddDeleteEnum string

const (
    // ipsla mpls add delete add q
    IpslaMplsAddDeleteEnum_ipsla_mpls_add_delete_add_q IpslaMplsAddDeleteEnum = "ipsla-mpls-add-delete-add-q"

    // ipsla mpls add delete delete q
    IpslaMplsAddDeleteEnum_ipsla_mpls_add_delete_delete_q IpslaMplsAddDeleteEnum = "ipsla-mpls-add-delete-delete-q"
)

// IpslaLspGrpPathStatusEnum represents Ipsla lsp grp path status enum
type IpslaLspGrpPathStatusEnum string

const (
    // ipsla lsp grp path status unknown
    IpslaLspGrpPathStatusEnum_ipsla_lsp_grp_path_status_unknown IpslaLspGrpPathStatusEnum = "ipsla-lsp-grp-path-status-unknown"

    // ipsla lsp grp path status up
    IpslaLspGrpPathStatusEnum_ipsla_lsp_grp_path_status_up IpslaLspGrpPathStatusEnum = "ipsla-lsp-grp-path-status-up"

    // ipsla lsp grp path status down
    IpslaLspGrpPathStatusEnum_ipsla_lsp_grp_path_status_down IpslaLspGrpPathStatusEnum = "ipsla-lsp-grp-path-status-down"

    // ipsla lsp grp path status retry
    IpslaLspGrpPathStatusEnum_ipsla_lsp_grp_path_status_retry IpslaLspGrpPathStatusEnum = "ipsla-lsp-grp-path-status-retry"

    // ipsla lsp grp path status pending
    IpslaLspGrpPathStatusEnum_ipsla_lsp_grp_path_status_pending IpslaLspGrpPathStatusEnum = "ipsla-lsp-grp-path-status-pending"
)

// IpslaLspGrpStatusEnum represents Ipsla lsp grp status enum
type IpslaLspGrpStatusEnum string

const (
    // ipsla lsp grp status unknown
    IpslaLspGrpStatusEnum_ipsla_lsp_grp_status_unknown IpslaLspGrpStatusEnum = "ipsla-lsp-grp-status-unknown"

    // ipsla lsp grp status up
    IpslaLspGrpStatusEnum_ipsla_lsp_grp_status_up IpslaLspGrpStatusEnum = "ipsla-lsp-grp-status-up"

    // ipsla lsp grp status partial
    IpslaLspGrpStatusEnum_ipsla_lsp_grp_status_partial IpslaLspGrpStatusEnum = "ipsla-lsp-grp-status-partial"

    // ipsla lsp grp status down
    IpslaLspGrpStatusEnum_ipsla_lsp_grp_status_down IpslaLspGrpStatusEnum = "ipsla-lsp-grp-status-down"

    // ipsla lsp grp status pending
    IpslaLspGrpStatusEnum_ipsla_lsp_grp_status_pending IpslaLspGrpStatusEnum = "ipsla-lsp-grp-status-pending"
)

// SlaOpTypes represents IPSLA Operation Types
type SlaOpTypes string

const (
    // ICMP Echo
    SlaOpTypes_oper_icmp_echo SlaOpTypes = "oper-icmp-echo"

    // ICMP PathJitter
    SlaOpTypes_oper_icmp_path_jitter SlaOpTypes = "oper-icmp-path-jitter"

    // ICMP Path Echo
    SlaOpTypes_oper_icmp_path_echo SlaOpTypes = "oper-icmp-path-echo"

    // UDP Jitter
    SlaOpTypes_oper_udp_jitter SlaOpTypes = "oper-udp-jitter"

    // UDP Echo
    SlaOpTypes_oper_udp_echo SlaOpTypes = "oper-udp-echo"

    // MPLS LSP Ping
    SlaOpTypes_oper_mpls_lsp_ping SlaOpTypes = "oper-mpls-lsp-ping"

    // MPLS LSP Trace
    SlaOpTypes_oper_mpls_lsp_trace SlaOpTypes = "oper-mpls-lsp-trace"

    // MPLS LSP Group
    SlaOpTypes_oper_mpls_lsp_group SlaOpTypes = "oper-mpls-lsp-group"
)

// Ipsla
// IPSLA operational data
type Ipsla struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS operational data.
    MplsData Ipsla_MplsData

    // Data from responder probe handling.
    Responder Ipsla_Responder

    // Operations data.
    OperationData Ipsla_OperationData

    // IPSLA application information.
    ApplicationInfo Ipsla_ApplicationInfo
}

func (ipsla *Ipsla) GetEntityData() *types.CommonEntityData {
    ipsla.EntityData.YFilter = ipsla.YFilter
    ipsla.EntityData.YangName = "ipsla"
    ipsla.EntityData.BundleName = "cisco_ios_xr"
    ipsla.EntityData.ParentYangName = "Cisco-IOS-XR-man-ipsla-oper"
    ipsla.EntityData.SegmentPath = "Cisco-IOS-XR-man-ipsla-oper:ipsla"
    ipsla.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipsla.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipsla.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipsla.EntityData.Children = make(map[string]types.YChild)
    ipsla.EntityData.Children["mpls-data"] = types.YChild{"MplsData", &ipsla.MplsData}
    ipsla.EntityData.Children["responder"] = types.YChild{"Responder", &ipsla.Responder}
    ipsla.EntityData.Children["operation-data"] = types.YChild{"OperationData", &ipsla.OperationData}
    ipsla.EntityData.Children["application-info"] = types.YChild{"ApplicationInfo", &ipsla.ApplicationInfo}
    ipsla.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipsla.EntityData)
}

// Ipsla_MplsData
// MPLS operational data
type Ipsla_MplsData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of MPLS LSP Monitor instances.
    LspMonitors Ipsla_MplsData_LspMonitors

    // Provider Edge(PE) discovery operational data.
    Discovery Ipsla_MplsData_Discovery
}

func (mplsData *Ipsla_MplsData) GetEntityData() *types.CommonEntityData {
    mplsData.EntityData.YFilter = mplsData.YFilter
    mplsData.EntityData.YangName = "mpls-data"
    mplsData.EntityData.BundleName = "cisco_ios_xr"
    mplsData.EntityData.ParentYangName = "ipsla"
    mplsData.EntityData.SegmentPath = "mpls-data"
    mplsData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsData.EntityData.Children = make(map[string]types.YChild)
    mplsData.EntityData.Children["lsp-monitors"] = types.YChild{"LspMonitors", &mplsData.LspMonitors}
    mplsData.EntityData.Children["discovery"] = types.YChild{"Discovery", &mplsData.Discovery}
    mplsData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsData.EntityData)
}

// Ipsla_MplsData_LspMonitors
// List of MPLS LSP Monitor instances
type Ipsla_MplsData_LspMonitors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for MPLS LSP Monitor. The type is slice of
    // Ipsla_MplsData_LspMonitors_LspMonitor.
    LspMonitor []Ipsla_MplsData_LspMonitors_LspMonitor
}

func (lspMonitors *Ipsla_MplsData_LspMonitors) GetEntityData() *types.CommonEntityData {
    lspMonitors.EntityData.YFilter = lspMonitors.YFilter
    lspMonitors.EntityData.YangName = "lsp-monitors"
    lspMonitors.EntityData.BundleName = "cisco_ios_xr"
    lspMonitors.EntityData.ParentYangName = "mpls-data"
    lspMonitors.EntityData.SegmentPath = "lsp-monitors"
    lspMonitors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspMonitors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspMonitors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspMonitors.EntityData.Children = make(map[string]types.YChild)
    lspMonitors.EntityData.Children["lsp-monitor"] = types.YChild{"LspMonitor", nil}
    for i := range lspMonitors.LspMonitor {
        lspMonitors.EntityData.Children[types.GetSegmentPath(&lspMonitors.LspMonitor[i])] = types.YChild{"LspMonitor", &lspMonitors.LspMonitor[i]}
    }
    lspMonitors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspMonitors.EntityData)
}

// Ipsla_MplsData_LspMonitors_LspMonitor
// Operational data for MPLS LSP Monitor
type Ipsla_MplsData_LspMonitors_LspMonitor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Monitor ID. The type is interface{} with range:
    // -2147483648..2147483647.
    MonitorId interface{}

    // Operational state of MPLS LSP Monitor.
    State Ipsla_MplsData_LspMonitors_LspMonitor_State

    // List of operations in MPLS LSP Monitor.
    Operations Ipsla_MplsData_LspMonitors_LspMonitor_Operations

    // List of Scan Queue entries in MPLS LSP Monitor.
    ScanQueues Ipsla_MplsData_LspMonitors_LspMonitor_ScanQueues
}

func (lspMonitor *Ipsla_MplsData_LspMonitors_LspMonitor) GetEntityData() *types.CommonEntityData {
    lspMonitor.EntityData.YFilter = lspMonitor.YFilter
    lspMonitor.EntityData.YangName = "lsp-monitor"
    lspMonitor.EntityData.BundleName = "cisco_ios_xr"
    lspMonitor.EntityData.ParentYangName = "lsp-monitors"
    lspMonitor.EntityData.SegmentPath = "lsp-monitor" + "[monitor-id='" + fmt.Sprintf("%v", lspMonitor.MonitorId) + "']"
    lspMonitor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspMonitor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspMonitor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspMonitor.EntityData.Children = make(map[string]types.YChild)
    lspMonitor.EntityData.Children["state"] = types.YChild{"State", &lspMonitor.State}
    lspMonitor.EntityData.Children["operations"] = types.YChild{"Operations", &lspMonitor.Operations}
    lspMonitor.EntityData.Children["scan-queues"] = types.YChild{"ScanQueues", &lspMonitor.ScanQueues}
    lspMonitor.EntityData.Leafs = make(map[string]types.YLeaf)
    lspMonitor.EntityData.Leafs["monitor-id"] = types.YLeaf{"MonitorId", lspMonitor.MonitorId}
    return &(lspMonitor.EntityData)
}

// Ipsla_MplsData_LspMonitors_LspMonitor_State
// Operational state of MPLS LSP Monitor
type Ipsla_MplsData_LspMonitors_LspMonitor_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds left before next scan for addition (0xffffffff means the
    // timer is not running). The type is interface{} with range: 0..4294967295.
    // Units are second.
    ScanRemaining interface{}

    // Number of seconds left before next scan for deletion  (0xffffffff means the
    // timer is not running). The type is interface{} with range: 0..4294967295.
    // Units are second.
    DeleteScanRemaining interface{}

    // Number of seconds left before next path discovery (0xffffffff means the
    // timer is not running). The type is interface{} with range: 0..4294967295.
    // Units are second.
    RediscoveryRemaining interface{}

    // LPD completion time (seconds) for the entire set of PEs which are
    // discovered in this MPLSLM instance (0xffffffff means LPD is never completed
    // yet). The type is interface{} with range: 0..4294967295. Units are second.
    LpdCompeletionTime interface{}
}

func (state *Ipsla_MplsData_LspMonitors_LspMonitor_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "lsp-monitor"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["scan-remaining"] = types.YLeaf{"ScanRemaining", state.ScanRemaining}
    state.EntityData.Leafs["delete-scan-remaining"] = types.YLeaf{"DeleteScanRemaining", state.DeleteScanRemaining}
    state.EntityData.Leafs["rediscovery-remaining"] = types.YLeaf{"RediscoveryRemaining", state.RediscoveryRemaining}
    state.EntityData.Leafs["lpd-compeletion-time"] = types.YLeaf{"LpdCompeletionTime", state.LpdCompeletionTime}
    return &(state.EntityData)
}

// Ipsla_MplsData_LspMonitors_LspMonitor_Operations
// List of operations in MPLS LSP Monitor
type Ipsla_MplsData_LspMonitors_LspMonitor_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation created in MPLS LSP Monitor. The type is slice of
    // Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation.
    Operation []Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation
}

func (operations *Ipsla_MplsData_LspMonitors_LspMonitor_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xr"
    operations.EntityData.ParentYangName = "lsp-monitor"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["operation"] = types.YChild{"Operation", nil}
    for i := range operations.Operation {
        operations.EntityData.Children[types.GetSegmentPath(&operations.Operation[i])] = types.YChild{"Operation", &operations.Operation[i]}
    }
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(operations.EntityData)
}

// Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation
// Operation created in MPLS LSP Monitor
type Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Operation ID. The type is interface{} with range:
    // -2147483648..2147483647.
    OperationId interface{}

    // Operational state of the created operation.
    State Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_State

    // List of LPD paths in MPLS LPD group operation.
    LpdPaths Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths
}

func (operation *Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation) GetEntityData() *types.CommonEntityData {
    operation.EntityData.YFilter = operation.YFilter
    operation.EntityData.YangName = "operation"
    operation.EntityData.BundleName = "cisco_ios_xr"
    operation.EntityData.ParentYangName = "operations"
    operation.EntityData.SegmentPath = "operation" + "[operation-id='" + fmt.Sprintf("%v", operation.OperationId) + "']"
    operation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operation.EntityData.Children = make(map[string]types.YChild)
    operation.EntityData.Children["state"] = types.YChild{"State", &operation.State}
    operation.EntityData.Children["lpd-paths"] = types.YChild{"LpdPaths", &operation.LpdPaths}
    operation.EntityData.Leafs = make(map[string]types.YLeaf)
    operation.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", operation.OperationId}
    return &(operation.EntityData)
}

// Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_State
// Operational state of the created operation
type Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PE target address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TargetAddress interface{}

    // PE target mask length. The type is interface{} with range: 0..4294967295.
    TargetMask interface{}

    // Latest LSP group status. The type is IpslaLspGrpStatusEnum.
    GroupStatus interface{}

    // Latest operation time. The type is interface{} with range:
    // 0..18446744073709551615.
    OperationTime interface{}
}

func (state *Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "operation"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["target-address"] = types.YLeaf{"TargetAddress", state.TargetAddress}
    state.EntityData.Leafs["target-mask"] = types.YLeaf{"TargetMask", state.TargetMask}
    state.EntityData.Leafs["group-status"] = types.YLeaf{"GroupStatus", state.GroupStatus}
    state.EntityData.Leafs["operation-time"] = types.YLeaf{"OperationTime", state.OperationTime}
    return &(state.EntityData)
}

// Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths
// List of LPD paths in MPLS LPD group
// operation
type Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational state of LPD path in MPLS LSP Group operation. The type is
    // slice of
    // Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths_LpdPath.
    LpdPath []Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths_LpdPath
}

func (lpdPaths *Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths) GetEntityData() *types.CommonEntityData {
    lpdPaths.EntityData.YFilter = lpdPaths.YFilter
    lpdPaths.EntityData.YangName = "lpd-paths"
    lpdPaths.EntityData.BundleName = "cisco_ios_xr"
    lpdPaths.EntityData.ParentYangName = "operation"
    lpdPaths.EntityData.SegmentPath = "lpd-paths"
    lpdPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpdPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpdPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpdPaths.EntityData.Children = make(map[string]types.YChild)
    lpdPaths.EntityData.Children["lpd-path"] = types.YChild{"LpdPath", nil}
    for i := range lpdPaths.LpdPath {
        lpdPaths.EntityData.Children[types.GetSegmentPath(&lpdPaths.LpdPath[i])] = types.YChild{"LpdPath", &lpdPaths.LpdPath[i]}
    }
    lpdPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lpdPaths.EntityData)
}

// Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths_LpdPath
// Operational state of LPD path in MPLS LSP
// Group operation
type Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths_LpdPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPD path index. The type is interface{} with
    // range: -2147483648..2147483647.
    PathIndex interface{}

    // Latest path status. The type is IpslaLspGrpPathStatusEnum.
    PathStatus interface{}

    // Latest operation time. The type is interface{} with range:
    // 0..18446744073709551615.
    OperationTime interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Number of path successes. The type is interface{} with range:
    // 0..4294967295.
    SuccessCount interface{}

    // Number of path failures. The type is interface{} with range: 0..4294967295.
    FailureCount interface{}

    // LPD path identifier.
    PathId Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths_LpdPath_PathId
}

func (lpdPath *Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths_LpdPath) GetEntityData() *types.CommonEntityData {
    lpdPath.EntityData.YFilter = lpdPath.YFilter
    lpdPath.EntityData.YangName = "lpd-path"
    lpdPath.EntityData.BundleName = "cisco_ios_xr"
    lpdPath.EntityData.ParentYangName = "lpd-paths"
    lpdPath.EntityData.SegmentPath = "lpd-path" + "[path-index='" + fmt.Sprintf("%v", lpdPath.PathIndex) + "']"
    lpdPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpdPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpdPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpdPath.EntityData.Children = make(map[string]types.YChild)
    lpdPath.EntityData.Children["path-id"] = types.YChild{"PathId", &lpdPath.PathId}
    lpdPath.EntityData.Leafs = make(map[string]types.YLeaf)
    lpdPath.EntityData.Leafs["path-index"] = types.YLeaf{"PathIndex", lpdPath.PathIndex}
    lpdPath.EntityData.Leafs["path-status"] = types.YLeaf{"PathStatus", lpdPath.PathStatus}
    lpdPath.EntityData.Leafs["operation-time"] = types.YLeaf{"OperationTime", lpdPath.OperationTime}
    lpdPath.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", lpdPath.ResponseTime}
    lpdPath.EntityData.Leafs["success-count"] = types.YLeaf{"SuccessCount", lpdPath.SuccessCount}
    lpdPath.EntityData.Leafs["failure-count"] = types.YLeaf{"FailureCount", lpdPath.FailureCount}
    return &(lpdPath.EntityData)
}

// Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths_LpdPath_PathId
// LPD path identifier
type Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths_LpdPath_PathId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP selector. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LspSelector interface{}

    // Output interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    OutputInterface interface{}

    // Nexthop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}

    // Downstream label stacks. The type is slice of interface{} with range:
    // 0..4294967295.
    DownstreamLabel []interface{}
}

func (pathId *Ipsla_MplsData_LspMonitors_LspMonitor_Operations_Operation_LpdPaths_LpdPath_PathId) GetEntityData() *types.CommonEntityData {
    pathId.EntityData.YFilter = pathId.YFilter
    pathId.EntityData.YangName = "path-id"
    pathId.EntityData.BundleName = "cisco_ios_xr"
    pathId.EntityData.ParentYangName = "lpd-path"
    pathId.EntityData.SegmentPath = "path-id"
    pathId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathId.EntityData.Children = make(map[string]types.YChild)
    pathId.EntityData.Leafs = make(map[string]types.YLeaf)
    pathId.EntityData.Leafs["lsp-selector"] = types.YLeaf{"LspSelector", pathId.LspSelector}
    pathId.EntityData.Leafs["output-interface"] = types.YLeaf{"OutputInterface", pathId.OutputInterface}
    pathId.EntityData.Leafs["nexthop-address"] = types.YLeaf{"NexthopAddress", pathId.NexthopAddress}
    pathId.EntityData.Leafs["downstream-label"] = types.YLeaf{"DownstreamLabel", pathId.DownstreamLabel}
    return &(pathId.EntityData)
}

// Ipsla_MplsData_LspMonitors_LspMonitor_ScanQueues
// List of Scan Queue entries in MPLS LSP
// Monitor
type Ipsla_MplsData_LspMonitors_LspMonitor_ScanQueues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Provider Edge(PE) addition or deletion requests in Scan Queue. The type is
    // slice of Ipsla_MplsData_LspMonitors_LspMonitor_ScanQueues_ScanQueue.
    ScanQueue []Ipsla_MplsData_LspMonitors_LspMonitor_ScanQueues_ScanQueue
}

func (scanQueues *Ipsla_MplsData_LspMonitors_LspMonitor_ScanQueues) GetEntityData() *types.CommonEntityData {
    scanQueues.EntityData.YFilter = scanQueues.YFilter
    scanQueues.EntityData.YangName = "scan-queues"
    scanQueues.EntityData.BundleName = "cisco_ios_xr"
    scanQueues.EntityData.ParentYangName = "lsp-monitor"
    scanQueues.EntityData.SegmentPath = "scan-queues"
    scanQueues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scanQueues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scanQueues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scanQueues.EntityData.Children = make(map[string]types.YChild)
    scanQueues.EntityData.Children["scan-queue"] = types.YChild{"ScanQueue", nil}
    for i := range scanQueues.ScanQueue {
        scanQueues.EntityData.Children[types.GetSegmentPath(&scanQueues.ScanQueue[i])] = types.YChild{"ScanQueue", &scanQueues.ScanQueue[i]}
    }
    scanQueues.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(scanQueues.EntityData)
}

// Ipsla_MplsData_LspMonitors_LspMonitor_ScanQueues_ScanQueue
// Provider Edge(PE) addition or deletion
// requests in Scan Queue
type Ipsla_MplsData_LspMonitors_LspMonitor_ScanQueues_ScanQueue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Nexthop Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // PE target address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TargetAddress interface{}

    // PE target mask length. The type is interface{} with range: 0..4294967295.
    TargetMask interface{}

    // PE addition or deletion. The type is IpslaMplsAddDeleteEnum.
    Entry interface{}
}

func (scanQueue *Ipsla_MplsData_LspMonitors_LspMonitor_ScanQueues_ScanQueue) GetEntityData() *types.CommonEntityData {
    scanQueue.EntityData.YFilter = scanQueue.YFilter
    scanQueue.EntityData.YangName = "scan-queue"
    scanQueue.EntityData.BundleName = "cisco_ios_xr"
    scanQueue.EntityData.ParentYangName = "scan-queues"
    scanQueue.EntityData.SegmentPath = "scan-queue" + "[address='" + fmt.Sprintf("%v", scanQueue.Address) + "']"
    scanQueue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scanQueue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scanQueue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scanQueue.EntityData.Children = make(map[string]types.YChild)
    scanQueue.EntityData.Leafs = make(map[string]types.YLeaf)
    scanQueue.EntityData.Leafs["address"] = types.YLeaf{"Address", scanQueue.Address}
    scanQueue.EntityData.Leafs["target-address"] = types.YLeaf{"TargetAddress", scanQueue.TargetAddress}
    scanQueue.EntityData.Leafs["target-mask"] = types.YLeaf{"TargetMask", scanQueue.TargetMask}
    scanQueue.EntityData.Leafs["entry"] = types.YLeaf{"Entry", scanQueue.Entry}
    return &(scanQueue.EntityData)
}

// Ipsla_MplsData_Discovery
// Provider Edge(PE) discovery operational data
type Ipsla_MplsData_Discovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L3 VPN PE discovery operational data.
    Vpn Ipsla_MplsData_Discovery_Vpn
}

func (discovery *Ipsla_MplsData_Discovery) GetEntityData() *types.CommonEntityData {
    discovery.EntityData.YFilter = discovery.YFilter
    discovery.EntityData.YangName = "discovery"
    discovery.EntityData.BundleName = "cisco_ios_xr"
    discovery.EntityData.ParentYangName = "mpls-data"
    discovery.EntityData.SegmentPath = "discovery"
    discovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    discovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    discovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    discovery.EntityData.Children = make(map[string]types.YChild)
    discovery.EntityData.Children["vpn"] = types.YChild{"Vpn", &discovery.Vpn}
    discovery.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(discovery.EntityData)
}

// Ipsla_MplsData_Discovery_Vpn
// L3 VPN PE discovery operational data
type Ipsla_MplsData_Discovery_Vpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational state of PE discovery.
    State Ipsla_MplsData_Discovery_Vpn_State

    // List of nexthop addresses for remote PE routers.
    Nexthops Ipsla_MplsData_Discovery_Vpn_Nexthops
}

func (vpn *Ipsla_MplsData_Discovery_Vpn) GetEntityData() *types.CommonEntityData {
    vpn.EntityData.YFilter = vpn.YFilter
    vpn.EntityData.YangName = "vpn"
    vpn.EntityData.BundleName = "cisco_ios_xr"
    vpn.EntityData.ParentYangName = "discovery"
    vpn.EntityData.SegmentPath = "vpn"
    vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpn.EntityData.Children = make(map[string]types.YChild)
    vpn.EntityData.Children["state"] = types.YChild{"State", &vpn.State}
    vpn.EntityData.Children["nexthops"] = types.YChild{"Nexthops", &vpn.Nexthops}
    vpn.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vpn.EntityData)
}

// Ipsla_MplsData_Discovery_Vpn_State
// Operational state of PE discovery
type Ipsla_MplsData_Discovery_Vpn_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds left before next refresh. The type is interface{} with
    // range: 0..4294967295. Units are second.
    RefreshRemaining interface{}
}

func (state *Ipsla_MplsData_Discovery_Vpn_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "vpn"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["refresh-remaining"] = types.YLeaf{"RefreshRemaining", state.RefreshRemaining}
    return &(state.EntityData)
}

// Ipsla_MplsData_Discovery_Vpn_Nexthops
// List of nexthop addresses for remote PE
// routers
type Ipsla_MplsData_Discovery_Vpn_Nexthops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nexthop address for remote PE router. The type is slice of
    // Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop.
    Nexthop []Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop
}

func (nexthops *Ipsla_MplsData_Discovery_Vpn_Nexthops) GetEntityData() *types.CommonEntityData {
    nexthops.EntityData.YFilter = nexthops.YFilter
    nexthops.EntityData.YangName = "nexthops"
    nexthops.EntityData.BundleName = "cisco_ios_xr"
    nexthops.EntityData.ParentYangName = "vpn"
    nexthops.EntityData.SegmentPath = "nexthops"
    nexthops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nexthops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nexthops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nexthops.EntityData.Children = make(map[string]types.YChild)
    nexthops.EntityData.Children["nexthop"] = types.YChild{"Nexthop", nil}
    for i := range nexthops.Nexthop {
        nexthops.EntityData.Children[types.GetSegmentPath(&nexthops.Nexthop[i])] = types.YChild{"Nexthop", &nexthops.Nexthop[i]}
    }
    nexthops.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nexthops.EntityData)
}

// Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop
// Nexthop address for remote PE router
type Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Nexthop Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // List of VRFs for the nexthop address.
    Vrfs Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Vrfs

    // Prefix of the nexthop address.
    Prefix Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Prefix
}

func (nexthop *Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop) GetEntityData() *types.CommonEntityData {
    nexthop.EntityData.YFilter = nexthop.YFilter
    nexthop.EntityData.YangName = "nexthop"
    nexthop.EntityData.BundleName = "cisco_ios_xr"
    nexthop.EntityData.ParentYangName = "nexthops"
    nexthop.EntityData.SegmentPath = "nexthop" + "[address='" + fmt.Sprintf("%v", nexthop.Address) + "']"
    nexthop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nexthop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nexthop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nexthop.EntityData.Children = make(map[string]types.YChild)
    nexthop.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &nexthop.Vrfs}
    nexthop.EntityData.Children["prefix"] = types.YChild{"Prefix", &nexthop.Prefix}
    nexthop.EntityData.Leafs = make(map[string]types.YLeaf)
    nexthop.EntityData.Leafs["address"] = types.YLeaf{"Address", nexthop.Address}
    return &(nexthop.EntityData)
}

// Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Vrfs
// List of VRFs for the nexthop address
type Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF information of the nexthop address. The type is slice of
    // Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Vrfs_Vrf.
    Vrf []Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Vrfs_Vrf
}

func (vrfs *Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "nexthop"
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

// Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Vrfs_Vrf
// VRF information of the nexthop address
type Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with length: 1..32.
    VrfName interface{}

    // Number of prefixes in VRF. The type is interface{} with range:
    // 0..4294967295.
    PrefixCount interface{}
}

func (vrf *Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    vrf.EntityData.Leafs["prefix-count"] = types.YLeaf{"PrefixCount", vrf.PrefixCount}
    return &(vrf.EntityData)
}

// Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Prefix
// Prefix of the nexthop address
type Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PE target address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TargetAddress interface{}

    // PE target mask length. The type is interface{} with range: 0..4294967295.
    TargetMask interface{}
}

func (prefix *Ipsla_MplsData_Discovery_Vpn_Nexthops_Nexthop_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "nexthop"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["target-address"] = types.YLeaf{"TargetAddress", prefix.TargetAddress}
    prefix.EntityData.Leafs["target-mask"] = types.YLeaf{"TargetMask", prefix.TargetMask}
    return &(prefix.EntityData)
}

// Ipsla_Responder
// Data from responder probe handling
type Ipsla_Responder struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics maintained by responder.
    Ports Ipsla_Responder_Ports
}

func (responder *Ipsla_Responder) GetEntityData() *types.CommonEntityData {
    responder.EntityData.YFilter = responder.YFilter
    responder.EntityData.YangName = "responder"
    responder.EntityData.BundleName = "cisco_ios_xr"
    responder.EntityData.ParentYangName = "ipsla"
    responder.EntityData.SegmentPath = "responder"
    responder.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    responder.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    responder.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    responder.EntityData.Children = make(map[string]types.YChild)
    responder.EntityData.Children["ports"] = types.YChild{"Ports", &responder.Ports}
    responder.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(responder.EntityData)
}

// Ipsla_Responder_Ports
// Statistics maintained by responder
type Ipsla_Responder_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port data. The type is slice of Ipsla_Responder_Ports_Port.
    Port []Ipsla_Responder_Ports_Port
}

func (ports *Ipsla_Responder_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "cisco_ios_xr"
    ports.EntityData.ParentYangName = "responder"
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

// Ipsla_Responder_Ports_Port
// Port data
type Ipsla_Responder_Ports_Port struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port. The type is interface{} with range:
    // 0..65535.
    Port interface{}

    // Port on which Responder is listening. The type is interface{} with range:
    // 0..65535.
    PortXr interface{}

    // IP address of Responder. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddress interface{}

    // Number of probes received from remote end. The type is interface{} with
    // range: 0..4294967295.
    NumProbes interface{}

    // Number of control probes received from remote end. The type is interface{}
    // with range: 0..4294967295.
    CtrlProbes interface{}

    // Port type if this is permanent or dynamic port. The type is bool.
    Permanent interface{}

    // Current discard socket option flag for the port. The type is bool.
    DiscardOn interface{}

    // PD Timestamp failure. The type is bool.
    PdTimeStampFailed interface{}

    // IPSLA or TWAMP protocol. The type is bool.
    IsIpsla interface{}

    // Drop counter for the Responder port. The type is interface{} with range:
    // 0..4294967295.
    DropCounter interface{}

    // Socket. The type is interface{} with range: -2147483648..2147483647.
    Socket interface{}

    // List of senders. The type is slice of Ipsla_Responder_Ports_Port_Sender.
    Sender []Ipsla_Responder_Ports_Port_Sender
}

func (port *Ipsla_Responder_Ports_Port) GetEntityData() *types.CommonEntityData {
    port.EntityData.YFilter = port.YFilter
    port.EntityData.YangName = "port"
    port.EntityData.BundleName = "cisco_ios_xr"
    port.EntityData.ParentYangName = "ports"
    port.EntityData.SegmentPath = "port" + "[port='" + fmt.Sprintf("%v", port.Port) + "']"
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = make(map[string]types.YChild)
    port.EntityData.Children["sender"] = types.YChild{"Sender", nil}
    for i := range port.Sender {
        port.EntityData.Children[types.GetSegmentPath(&port.Sender[i])] = types.YChild{"Sender", &port.Sender[i]}
    }
    port.EntityData.Leafs = make(map[string]types.YLeaf)
    port.EntityData.Leafs["port"] = types.YLeaf{"Port", port.Port}
    port.EntityData.Leafs["port-xr"] = types.YLeaf{"PortXr", port.PortXr}
    port.EntityData.Leafs["local-address"] = types.YLeaf{"LocalAddress", port.LocalAddress}
    port.EntityData.Leafs["num-probes"] = types.YLeaf{"NumProbes", port.NumProbes}
    port.EntityData.Leafs["ctrl-probes"] = types.YLeaf{"CtrlProbes", port.CtrlProbes}
    port.EntityData.Leafs["permanent"] = types.YLeaf{"Permanent", port.Permanent}
    port.EntityData.Leafs["discard-on"] = types.YLeaf{"DiscardOn", port.DiscardOn}
    port.EntityData.Leafs["pd-time-stamp-failed"] = types.YLeaf{"PdTimeStampFailed", port.PdTimeStampFailed}
    port.EntityData.Leafs["is-ipsla"] = types.YLeaf{"IsIpsla", port.IsIpsla}
    port.EntityData.Leafs["drop-counter"] = types.YLeaf{"DropCounter", port.DropCounter}
    port.EntityData.Leafs["socket"] = types.YLeaf{"Socket", port.Socket}
    return &(port.EntityData)
}

// Ipsla_Responder_Ports_Port_Sender
// List of senders
type Ipsla_Responder_Ports_Port_Sender struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of Sender. The type is interface{} with range: 0..4294967295.
    IpAddress interface{}

    // Port on which Sender is sending. The type is interface{} with range:
    // 0..65535.
    Port interface{}

    // Last received time. The type is interface{} with range:
    // 0..18446744073709551615.
    LastRecvTime interface{}
}

func (sender *Ipsla_Responder_Ports_Port_Sender) GetEntityData() *types.CommonEntityData {
    sender.EntityData.YFilter = sender.YFilter
    sender.EntityData.YangName = "sender"
    sender.EntityData.BundleName = "cisco_ios_xr"
    sender.EntityData.ParentYangName = "port"
    sender.EntityData.SegmentPath = "sender"
    sender.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sender.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sender.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sender.EntityData.Children = make(map[string]types.YChild)
    sender.EntityData.Leafs = make(map[string]types.YLeaf)
    sender.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", sender.IpAddress}
    sender.EntityData.Leafs["port"] = types.YLeaf{"Port", sender.Port}
    sender.EntityData.Leafs["last-recv-time"] = types.YLeaf{"LastRecvTime", sender.LastRecvTime}
    return &(sender.EntityData)
}

// Ipsla_OperationData
// Operations data
type Ipsla_OperationData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured operations.
    Operations Ipsla_OperationData_Operations
}

func (operationData *Ipsla_OperationData) GetEntityData() *types.CommonEntityData {
    operationData.EntityData.YFilter = operationData.YFilter
    operationData.EntityData.YangName = "operation-data"
    operationData.EntityData.BundleName = "cisco_ios_xr"
    operationData.EntityData.ParentYangName = "ipsla"
    operationData.EntityData.SegmentPath = "operation-data"
    operationData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationData.EntityData.Children = make(map[string]types.YChild)
    operationData.EntityData.Children["operations"] = types.YChild{"Operations", &operationData.Operations}
    operationData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(operationData.EntityData)
}

// Ipsla_OperationData_Operations
// Configured operations
type Ipsla_OperationData_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for an operation. The type is slice of
    // Ipsla_OperationData_Operations_Operation.
    Operation []Ipsla_OperationData_Operations_Operation
}

func (operations *Ipsla_OperationData_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xr"
    operations.EntityData.ParentYangName = "operation-data"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operations.EntityData.Children = make(map[string]types.YChild)
    operations.EntityData.Children["operation"] = types.YChild{"Operation", nil}
    for i := range operations.Operation {
        operations.EntityData.Children[types.GetSegmentPath(&operations.Operation[i])] = types.YChild{"Operation", &operations.Operation[i]}
    }
    operations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(operations.EntityData)
}

// Ipsla_OperationData_Operations_Operation
// Operational data for an operation
type Ipsla_OperationData_Operations_Operation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Operation ID. The type is interface{} with range:
    // -2147483648..2147483647.
    OperationId interface{}

    // Common data for all operation types.
    Common Ipsla_OperationData_Operations_Operation_Common

    // LPD operational data of MPLS LSP group operation.
    Lpd Ipsla_OperationData_Operations_Operation_Lpd

    // Historical data for an operation.
    History Ipsla_OperationData_Operations_Operation_History

    // Statistics collected for an operation.
    Statistics Ipsla_OperationData_Operations_Operation_Statistics
}

func (operation *Ipsla_OperationData_Operations_Operation) GetEntityData() *types.CommonEntityData {
    operation.EntityData.YFilter = operation.YFilter
    operation.EntityData.YangName = "operation"
    operation.EntityData.BundleName = "cisco_ios_xr"
    operation.EntityData.ParentYangName = "operations"
    operation.EntityData.SegmentPath = "operation" + "[operation-id='" + fmt.Sprintf("%v", operation.OperationId) + "']"
    operation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operation.EntityData.Children = make(map[string]types.YChild)
    operation.EntityData.Children["common"] = types.YChild{"Common", &operation.Common}
    operation.EntityData.Children["lpd"] = types.YChild{"Lpd", &operation.Lpd}
    operation.EntityData.Children["history"] = types.YChild{"History", &operation.History}
    operation.EntityData.Children["statistics"] = types.YChild{"Statistics", &operation.Statistics}
    operation.EntityData.Leafs = make(map[string]types.YLeaf)
    operation.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", operation.OperationId}
    return &(operation.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Common
// Common data for all operation types
type Ipsla_OperationData_Operations_Operation_Common struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational state for an operation.
    OperationalState Ipsla_OperationData_Operations_Operation_Common_OperationalState
}

func (common *Ipsla_OperationData_Operations_Operation_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xr"
    common.EntityData.ParentYangName = "operation"
    common.EntityData.SegmentPath = "common"
    common.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    common.EntityData.Children = make(map[string]types.YChild)
    common.EntityData.Children["operational-state"] = types.YChild{"OperationalState", &common.OperationalState}
    common.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(common.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Common_OperationalState
// Operational state for an operation
type Ipsla_OperationData_Operations_Operation_Common_OperationalState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Last modification time of the operation expressed in msec since 00:00:00
    // UTC, January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615.
    ModificationTime interface{}

    // Last start time of the operation expressedin msec since 00:00:00 UTC,
    // January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615.
    StartTime interface{}

    // Number of data collection attempts. The type is interface{} with range:
    // 0..4294967295.
    AttemptCount interface{}

    // Number of data collection cycles skipped. The type is interface{} with
    // range: 0..4294967295.
    SkippedCount interface{}

    // Number of seconds left in current life. The type is interface{} with range:
    // 0..4294967295. Units are second.
    LifeRemaining interface{}

    // Number of configured frequency Default 60 . The type is interface{} with
    // range: 0..4294967295.
    Frequency interface{}

    // For recurring operation configured. The type is bool.
    Recurring interface{}

    // Operational state. The type is IpslaOperStateEnum.
    OperationalState interface{}

    // Internal flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // Cached local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Unexpected probe pkts punted from LPTS. The type is interface{} with range:
    // 0..4294967295.
    UnexpectedPackets interface{}

    // Unexpected control pkts puntedfrom LPTS. The type is interface{} with
    // range: 0..4294967295.
    UnexpectedControlPackets interface{}

    // Start time of current instance of the operation. The type is interface{}
    // with range: 0..18446744073709551615.
    OperationTime interface{}
}

func (operationalState *Ipsla_OperationData_Operations_Operation_Common_OperationalState) GetEntityData() *types.CommonEntityData {
    operationalState.EntityData.YFilter = operationalState.YFilter
    operationalState.EntityData.YangName = "operational-state"
    operationalState.EntityData.BundleName = "cisco_ios_xr"
    operationalState.EntityData.ParentYangName = "common"
    operationalState.EntityData.SegmentPath = "operational-state"
    operationalState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationalState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationalState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationalState.EntityData.Children = make(map[string]types.YChild)
    operationalState.EntityData.Leafs = make(map[string]types.YLeaf)
    operationalState.EntityData.Leafs["modification-time"] = types.YLeaf{"ModificationTime", operationalState.ModificationTime}
    operationalState.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", operationalState.StartTime}
    operationalState.EntityData.Leafs["attempt-count"] = types.YLeaf{"AttemptCount", operationalState.AttemptCount}
    operationalState.EntityData.Leafs["skipped-count"] = types.YLeaf{"SkippedCount", operationalState.SkippedCount}
    operationalState.EntityData.Leafs["life-remaining"] = types.YLeaf{"LifeRemaining", operationalState.LifeRemaining}
    operationalState.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", operationalState.Frequency}
    operationalState.EntityData.Leafs["recurring"] = types.YLeaf{"Recurring", operationalState.Recurring}
    operationalState.EntityData.Leafs["operational-state"] = types.YLeaf{"OperationalState", operationalState.OperationalState}
    operationalState.EntityData.Leafs["flags"] = types.YLeaf{"Flags", operationalState.Flags}
    operationalState.EntityData.Leafs["local-port"] = types.YLeaf{"LocalPort", operationalState.LocalPort}
    operationalState.EntityData.Leafs["unexpected-packets"] = types.YLeaf{"UnexpectedPackets", operationalState.UnexpectedPackets}
    operationalState.EntityData.Leafs["unexpected-control-packets"] = types.YLeaf{"UnexpectedControlPackets", operationalState.UnexpectedControlPackets}
    operationalState.EntityData.Leafs["operation-time"] = types.YLeaf{"OperationTime", operationalState.OperationTime}
    return &(operationalState.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd
// LPD operational data of MPLS LSP group
// operation
type Ipsla_OperationData_Operations_Operation_Lpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics collected for LPD group.
    Statistics Ipsla_OperationData_Operations_Operation_Lpd_Statistics

    // Operational status of LPD group.
    Status Ipsla_OperationData_Operations_Operation_Lpd_Status
}

func (lpd *Ipsla_OperationData_Operations_Operation_Lpd) GetEntityData() *types.CommonEntityData {
    lpd.EntityData.YFilter = lpd.YFilter
    lpd.EntityData.YangName = "lpd"
    lpd.EntityData.BundleName = "cisco_ios_xr"
    lpd.EntityData.ParentYangName = "operation"
    lpd.EntityData.SegmentPath = "lpd"
    lpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpd.EntityData.Children = make(map[string]types.YChild)
    lpd.EntityData.Children["statistics"] = types.YChild{"Statistics", &lpd.Statistics}
    lpd.EntityData.Children["status"] = types.YChild{"Status", &lpd.Status}
    lpd.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lpd.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics
// Statistics collected for LPD group
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LPD statistics collected during the last sampling cycle.
    Latest Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest

    // Statistics aggregated for LPD group collected over time intervals.
    Aggregated Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated
}

func (statistics *Ipsla_OperationData_Operations_Operation_Lpd_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "lpd"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["latest"] = types.YChild{"Latest", &statistics.Latest}
    statistics.EntityData.Children["aggregated"] = types.YChild{"Aggregated", &statistics.Aggregated}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest
// LPD statistics collected during the last
// sampling cycle
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Latest statistics of LPD group.
    Target Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target
}

func (latest *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest) GetEntityData() *types.CommonEntityData {
    latest.EntityData.YFilter = latest.YFilter
    latest.EntityData.YangName = "latest"
    latest.EntityData.BundleName = "cisco_ios_xr"
    latest.EntityData.ParentYangName = "statistics"
    latest.EntityData.SegmentPath = "latest"
    latest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    latest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    latest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    latest.EntityData.Children = make(map[string]types.YChild)
    latest.EntityData.Children["target"] = types.YChild{"Target", &latest.Target}
    latest.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(latest.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target
// Latest statistics of LPD group
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LPD start time. The type is interface{} with range:
    // 0..18446744073709551615.
    StartTime interface{}

    // LPD return code. The type is IpslaMplsLpdRetCode.
    ReturnCode interface{}

    // Number of CompT samples. The type is interface{} with range: 0..4294967295.
    CompletionTimeCount interface{}

    // LPD Completion time. The type is interface{} with range: 0..4294967295.
    CompletionTime interface{}

    // Minimum CompT. The type is interface{} with range: 0..4294967295.
    MinCompletionTime interface{}

    // Maximum CompT. The type is interface{} with range: 0..4294967295.
    MaxCompletionTime interface{}

    // Sum of CompT. The type is interface{} with range: 0..4294967295.
    SumCompletionTime interface{}

    // Number of paths. The type is interface{} with range: 0..4294967295.
    PathCount interface{}

    // Minimum number of paths. The type is interface{} with range: 0..4294967295.
    MinPathCount interface{}

    // Maximum number of paths. The type is interface{} with range: 0..4294967295.
    MaxPathCount interface{}

    // Number of successes. The type is interface{} with range: 0..4294967295.
    OkCount interface{}

    // Number of failures due to no path. The type is interface{} with range:
    // 0..4294967295.
    NoPathCount interface{}

    // Number of failures due to all paths broken. The type is interface{} with
    // range: 0..4294967295.
    AllPathsBrokenCount interface{}

    // Number of failures due to all paths unexplorable. The type is interface{}
    // with range: 0..4294967295.
    AllPathsUnexplorableCount interface{}

    // Number of failures due to all paths broken or unexplorable. The type is
    // interface{} with range: 0..4294967295.
    AllPathsBrokenOrUnexplorableCount interface{}

    // Number of failures due to timeout. The type is interface{} with range:
    // 0..4294967295.
    TimeoutCount interface{}

    // Number of failures due to internal error. The type is interface{} with
    // range: 0..4294967295.
    InternalErrorCount interface{}

    // Number of failures due to unknown cause. The type is interface{} with
    // range: 0..4294967295.
    UnknownCount interface{}

    // LPD target.
    TargetAddress Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress
}

func (target *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "cisco_ios_xr"
    target.EntityData.ParentYangName = "latest"
    target.EntityData.SegmentPath = "target"
    target.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    target.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    target.EntityData.Children = make(map[string]types.YChild)
    target.EntityData.Children["target-address"] = types.YChild{"TargetAddress", &target.TargetAddress}
    target.EntityData.Leafs = make(map[string]types.YLeaf)
    target.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", target.StartTime}
    target.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", target.ReturnCode}
    target.EntityData.Leafs["completion-time-count"] = types.YLeaf{"CompletionTimeCount", target.CompletionTimeCount}
    target.EntityData.Leafs["completion-time"] = types.YLeaf{"CompletionTime", target.CompletionTime}
    target.EntityData.Leafs["min-completion-time"] = types.YLeaf{"MinCompletionTime", target.MinCompletionTime}
    target.EntityData.Leafs["max-completion-time"] = types.YLeaf{"MaxCompletionTime", target.MaxCompletionTime}
    target.EntityData.Leafs["sum-completion-time"] = types.YLeaf{"SumCompletionTime", target.SumCompletionTime}
    target.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", target.PathCount}
    target.EntityData.Leafs["min-path-count"] = types.YLeaf{"MinPathCount", target.MinPathCount}
    target.EntityData.Leafs["max-path-count"] = types.YLeaf{"MaxPathCount", target.MaxPathCount}
    target.EntityData.Leafs["ok-count"] = types.YLeaf{"OkCount", target.OkCount}
    target.EntityData.Leafs["no-path-count"] = types.YLeaf{"NoPathCount", target.NoPathCount}
    target.EntityData.Leafs["all-paths-broken-count"] = types.YLeaf{"AllPathsBrokenCount", target.AllPathsBrokenCount}
    target.EntityData.Leafs["all-paths-unexplorable-count"] = types.YLeaf{"AllPathsUnexplorableCount", target.AllPathsUnexplorableCount}
    target.EntityData.Leafs["all-paths-broken-or-unexplorable-count"] = types.YLeaf{"AllPathsBrokenOrUnexplorableCount", target.AllPathsBrokenOrUnexplorableCount}
    target.EntityData.Leafs["timeout-count"] = types.YLeaf{"TimeoutCount", target.TimeoutCount}
    target.EntityData.Leafs["internal-error-count"] = types.YLeaf{"InternalErrorCount", target.InternalErrorCount}
    target.EntityData.Leafs["unknown-count"] = types.YLeaf{"UnknownCount", target.UnknownCount}
    return &(target.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress
// LPD target
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TargetType. The type is IpslaTargetTypeEnum.
    TargetType interface{}

    // IPv4 address target. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4AddressTarget interface{}

    // IPv6 address target. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6AddressTarget interface{}

    // IPv4 prefix target.
    Ipv4PrefixTarget Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_Ipv4PrefixTarget

    // Tunnel ID target.
    TunnelIdTarget Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_TunnelIdTarget

    // IPv4 pseudowire target.
    Ipv4PseudowireTarget Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_Ipv4PseudowireTarget
}

func (targetAddress *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress) GetEntityData() *types.CommonEntityData {
    targetAddress.EntityData.YFilter = targetAddress.YFilter
    targetAddress.EntityData.YangName = "target-address"
    targetAddress.EntityData.BundleName = "cisco_ios_xr"
    targetAddress.EntityData.ParentYangName = "target"
    targetAddress.EntityData.SegmentPath = "target-address"
    targetAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetAddress.EntityData.Children = make(map[string]types.YChild)
    targetAddress.EntityData.Children["ipv4-prefix-target"] = types.YChild{"Ipv4PrefixTarget", &targetAddress.Ipv4PrefixTarget}
    targetAddress.EntityData.Children["tunnel-id-target"] = types.YChild{"TunnelIdTarget", &targetAddress.TunnelIdTarget}
    targetAddress.EntityData.Children["ipv4-pseudowire-target"] = types.YChild{"Ipv4PseudowireTarget", &targetAddress.Ipv4PseudowireTarget}
    targetAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    targetAddress.EntityData.Leafs["target-type"] = types.YLeaf{"TargetType", targetAddress.TargetType}
    targetAddress.EntityData.Leafs["ipv4-address-target"] = types.YLeaf{"Ipv4AddressTarget", targetAddress.Ipv4AddressTarget}
    targetAddress.EntityData.Leafs["ipv6-address-target"] = types.YLeaf{"Ipv6AddressTarget", targetAddress.Ipv6AddressTarget}
    return &(targetAddress.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_Ipv4PrefixTarget
// IPv4 prefix target
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_Ipv4PrefixTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Mask length. The type is interface{} with range: 0..255.
    MaskLength interface{}
}

func (ipv4PrefixTarget *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_Ipv4PrefixTarget) GetEntityData() *types.CommonEntityData {
    ipv4PrefixTarget.EntityData.YFilter = ipv4PrefixTarget.YFilter
    ipv4PrefixTarget.EntityData.YangName = "ipv4-prefix-target"
    ipv4PrefixTarget.EntityData.BundleName = "cisco_ios_xr"
    ipv4PrefixTarget.EntityData.ParentYangName = "target-address"
    ipv4PrefixTarget.EntityData.SegmentPath = "ipv4-prefix-target"
    ipv4PrefixTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PrefixTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PrefixTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PrefixTarget.EntityData.Children = make(map[string]types.YChild)
    ipv4PrefixTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4PrefixTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4PrefixTarget.Address}
    ipv4PrefixTarget.EntityData.Leafs["mask-length"] = types.YLeaf{"MaskLength", ipv4PrefixTarget.MaskLength}
    return &(ipv4PrefixTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_TunnelIdTarget
// Tunnel ID target
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_TunnelIdTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}
}

func (tunnelIdTarget *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_TunnelIdTarget) GetEntityData() *types.CommonEntityData {
    tunnelIdTarget.EntityData.YFilter = tunnelIdTarget.YFilter
    tunnelIdTarget.EntityData.YangName = "tunnel-id-target"
    tunnelIdTarget.EntityData.BundleName = "cisco_ios_xr"
    tunnelIdTarget.EntityData.ParentYangName = "target-address"
    tunnelIdTarget.EntityData.SegmentPath = "tunnel-id-target"
    tunnelIdTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelIdTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelIdTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelIdTarget.EntityData.Children = make(map[string]types.YChild)
    tunnelIdTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelIdTarget.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", tunnelIdTarget.TunnelId}
    return &(tunnelIdTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_Ipv4PseudowireTarget
// IPv4 pseudowire target
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_Ipv4PseudowireTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Virtual circuit ID. The type is interface{} with range: 0..4294967295.
    VirtualCircuitId interface{}
}

func (ipv4PseudowireTarget *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Latest_Target_TargetAddress_Ipv4PseudowireTarget) GetEntityData() *types.CommonEntityData {
    ipv4PseudowireTarget.EntityData.YFilter = ipv4PseudowireTarget.YFilter
    ipv4PseudowireTarget.EntityData.YangName = "ipv4-pseudowire-target"
    ipv4PseudowireTarget.EntityData.BundleName = "cisco_ios_xr"
    ipv4PseudowireTarget.EntityData.ParentYangName = "target-address"
    ipv4PseudowireTarget.EntityData.SegmentPath = "ipv4-pseudowire-target"
    ipv4PseudowireTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PseudowireTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PseudowireTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PseudowireTarget.EntityData.Children = make(map[string]types.YChild)
    ipv4PseudowireTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4PseudowireTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4PseudowireTarget.Address}
    ipv4PseudowireTarget.EntityData.Leafs["virtual-circuit-id"] = types.YLeaf{"VirtualCircuitId", ipv4PseudowireTarget.VirtualCircuitId}
    return &(ipv4PseudowireTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated
// Statistics aggregated for LPD group
// collected over time intervals
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of LPD statistics aggregated over 1-hour intervals.
    Hours Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours
}

func (aggregated *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated) GetEntityData() *types.CommonEntityData {
    aggregated.EntityData.YFilter = aggregated.YFilter
    aggregated.EntityData.YangName = "aggregated"
    aggregated.EntityData.BundleName = "cisco_ios_xr"
    aggregated.EntityData.ParentYangName = "statistics"
    aggregated.EntityData.SegmentPath = "aggregated"
    aggregated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregated.EntityData.Children = make(map[string]types.YChild)
    aggregated.EntityData.Children["hours"] = types.YChild{"Hours", &aggregated.Hours}
    aggregated.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aggregated.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours
// Table of LPD statistics aggregated over
// 1-hour intervals
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LPD statistics aggregated for a 1-hour interval. The type is slice of
    // Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour.
    Hour []Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour
}

func (hours *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours) GetEntityData() *types.CommonEntityData {
    hours.EntityData.YFilter = hours.YFilter
    hours.EntityData.YangName = "hours"
    hours.EntityData.BundleName = "cisco_ios_xr"
    hours.EntityData.ParentYangName = "aggregated"
    hours.EntityData.SegmentPath = "hours"
    hours.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hours.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hours.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hours.EntityData.Children = make(map[string]types.YChild)
    hours.EntityData.Children["hour"] = types.YChild{"Hour", nil}
    for i := range hours.Hour {
        hours.EntityData.Children[types.GetSegmentPath(&hours.Hour[i])] = types.YChild{"Hour", &hours.Hour[i]}
    }
    hours.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hours.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour
// LPD statistics aggregated for a 1-hour
// interval
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Hour Index. The type is interface{} with range:
    // -2147483648..2147483647.
    HourIndex interface{}

    // LPD start time. The type is interface{} with range:
    // 0..18446744073709551615.
    StartTime interface{}

    // LPD return code. The type is IpslaMplsLpdRetCode.
    ReturnCode interface{}

    // Number of CompT samples. The type is interface{} with range: 0..4294967295.
    CompletionTimeCount interface{}

    // LPD Completion time. The type is interface{} with range: 0..4294967295.
    CompletionTime interface{}

    // Minimum CompT. The type is interface{} with range: 0..4294967295.
    MinCompletionTime interface{}

    // Maximum CompT. The type is interface{} with range: 0..4294967295.
    MaxCompletionTime interface{}

    // Sum of CompT. The type is interface{} with range: 0..4294967295.
    SumCompletionTime interface{}

    // Number of paths. The type is interface{} with range: 0..4294967295.
    PathCount interface{}

    // Minimum number of paths. The type is interface{} with range: 0..4294967295.
    MinPathCount interface{}

    // Maximum number of paths. The type is interface{} with range: 0..4294967295.
    MaxPathCount interface{}

    // Number of successes. The type is interface{} with range: 0..4294967295.
    OkCount interface{}

    // Number of failures due to no path. The type is interface{} with range:
    // 0..4294967295.
    NoPathCount interface{}

    // Number of failures due to all paths broken. The type is interface{} with
    // range: 0..4294967295.
    AllPathsBrokenCount interface{}

    // Number of failures due to all paths unexplorable. The type is interface{}
    // with range: 0..4294967295.
    AllPathsUnexplorableCount interface{}

    // Number of failures due to all paths broken or unexplorable. The type is
    // interface{} with range: 0..4294967295.
    AllPathsBrokenOrUnexplorableCount interface{}

    // Number of failures due to timeout. The type is interface{} with range:
    // 0..4294967295.
    TimeoutCount interface{}

    // Number of failures due to internal error. The type is interface{} with
    // range: 0..4294967295.
    InternalErrorCount interface{}

    // Number of failures due to unknown cause. The type is interface{} with
    // range: 0..4294967295.
    UnknownCount interface{}

    // LPD target.
    TargetAddress Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress
}

func (hour *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour) GetEntityData() *types.CommonEntityData {
    hour.EntityData.YFilter = hour.YFilter
    hour.EntityData.YangName = "hour"
    hour.EntityData.BundleName = "cisco_ios_xr"
    hour.EntityData.ParentYangName = "hours"
    hour.EntityData.SegmentPath = "hour" + "[hour-index='" + fmt.Sprintf("%v", hour.HourIndex) + "']"
    hour.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hour.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hour.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hour.EntityData.Children = make(map[string]types.YChild)
    hour.EntityData.Children["target-address"] = types.YChild{"TargetAddress", &hour.TargetAddress}
    hour.EntityData.Leafs = make(map[string]types.YLeaf)
    hour.EntityData.Leafs["hour-index"] = types.YLeaf{"HourIndex", hour.HourIndex}
    hour.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", hour.StartTime}
    hour.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", hour.ReturnCode}
    hour.EntityData.Leafs["completion-time-count"] = types.YLeaf{"CompletionTimeCount", hour.CompletionTimeCount}
    hour.EntityData.Leafs["completion-time"] = types.YLeaf{"CompletionTime", hour.CompletionTime}
    hour.EntityData.Leafs["min-completion-time"] = types.YLeaf{"MinCompletionTime", hour.MinCompletionTime}
    hour.EntityData.Leafs["max-completion-time"] = types.YLeaf{"MaxCompletionTime", hour.MaxCompletionTime}
    hour.EntityData.Leafs["sum-completion-time"] = types.YLeaf{"SumCompletionTime", hour.SumCompletionTime}
    hour.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", hour.PathCount}
    hour.EntityData.Leafs["min-path-count"] = types.YLeaf{"MinPathCount", hour.MinPathCount}
    hour.EntityData.Leafs["max-path-count"] = types.YLeaf{"MaxPathCount", hour.MaxPathCount}
    hour.EntityData.Leafs["ok-count"] = types.YLeaf{"OkCount", hour.OkCount}
    hour.EntityData.Leafs["no-path-count"] = types.YLeaf{"NoPathCount", hour.NoPathCount}
    hour.EntityData.Leafs["all-paths-broken-count"] = types.YLeaf{"AllPathsBrokenCount", hour.AllPathsBrokenCount}
    hour.EntityData.Leafs["all-paths-unexplorable-count"] = types.YLeaf{"AllPathsUnexplorableCount", hour.AllPathsUnexplorableCount}
    hour.EntityData.Leafs["all-paths-broken-or-unexplorable-count"] = types.YLeaf{"AllPathsBrokenOrUnexplorableCount", hour.AllPathsBrokenOrUnexplorableCount}
    hour.EntityData.Leafs["timeout-count"] = types.YLeaf{"TimeoutCount", hour.TimeoutCount}
    hour.EntityData.Leafs["internal-error-count"] = types.YLeaf{"InternalErrorCount", hour.InternalErrorCount}
    hour.EntityData.Leafs["unknown-count"] = types.YLeaf{"UnknownCount", hour.UnknownCount}
    return &(hour.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress
// LPD target
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TargetType. The type is IpslaTargetTypeEnum.
    TargetType interface{}

    // IPv4 address target. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4AddressTarget interface{}

    // IPv6 address target. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6AddressTarget interface{}

    // IPv4 prefix target.
    Ipv4PrefixTarget Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_Ipv4PrefixTarget

    // Tunnel ID target.
    TunnelIdTarget Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_TunnelIdTarget

    // IPv4 pseudowire target.
    Ipv4PseudowireTarget Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_Ipv4PseudowireTarget
}

func (targetAddress *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress) GetEntityData() *types.CommonEntityData {
    targetAddress.EntityData.YFilter = targetAddress.YFilter
    targetAddress.EntityData.YangName = "target-address"
    targetAddress.EntityData.BundleName = "cisco_ios_xr"
    targetAddress.EntityData.ParentYangName = "hour"
    targetAddress.EntityData.SegmentPath = "target-address"
    targetAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetAddress.EntityData.Children = make(map[string]types.YChild)
    targetAddress.EntityData.Children["ipv4-prefix-target"] = types.YChild{"Ipv4PrefixTarget", &targetAddress.Ipv4PrefixTarget}
    targetAddress.EntityData.Children["tunnel-id-target"] = types.YChild{"TunnelIdTarget", &targetAddress.TunnelIdTarget}
    targetAddress.EntityData.Children["ipv4-pseudowire-target"] = types.YChild{"Ipv4PseudowireTarget", &targetAddress.Ipv4PseudowireTarget}
    targetAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    targetAddress.EntityData.Leafs["target-type"] = types.YLeaf{"TargetType", targetAddress.TargetType}
    targetAddress.EntityData.Leafs["ipv4-address-target"] = types.YLeaf{"Ipv4AddressTarget", targetAddress.Ipv4AddressTarget}
    targetAddress.EntityData.Leafs["ipv6-address-target"] = types.YLeaf{"Ipv6AddressTarget", targetAddress.Ipv6AddressTarget}
    return &(targetAddress.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_Ipv4PrefixTarget
// IPv4 prefix target
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_Ipv4PrefixTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Mask length. The type is interface{} with range: 0..255.
    MaskLength interface{}
}

func (ipv4PrefixTarget *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_Ipv4PrefixTarget) GetEntityData() *types.CommonEntityData {
    ipv4PrefixTarget.EntityData.YFilter = ipv4PrefixTarget.YFilter
    ipv4PrefixTarget.EntityData.YangName = "ipv4-prefix-target"
    ipv4PrefixTarget.EntityData.BundleName = "cisco_ios_xr"
    ipv4PrefixTarget.EntityData.ParentYangName = "target-address"
    ipv4PrefixTarget.EntityData.SegmentPath = "ipv4-prefix-target"
    ipv4PrefixTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PrefixTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PrefixTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PrefixTarget.EntityData.Children = make(map[string]types.YChild)
    ipv4PrefixTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4PrefixTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4PrefixTarget.Address}
    ipv4PrefixTarget.EntityData.Leafs["mask-length"] = types.YLeaf{"MaskLength", ipv4PrefixTarget.MaskLength}
    return &(ipv4PrefixTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_TunnelIdTarget
// Tunnel ID target
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_TunnelIdTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}
}

func (tunnelIdTarget *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_TunnelIdTarget) GetEntityData() *types.CommonEntityData {
    tunnelIdTarget.EntityData.YFilter = tunnelIdTarget.YFilter
    tunnelIdTarget.EntityData.YangName = "tunnel-id-target"
    tunnelIdTarget.EntityData.BundleName = "cisco_ios_xr"
    tunnelIdTarget.EntityData.ParentYangName = "target-address"
    tunnelIdTarget.EntityData.SegmentPath = "tunnel-id-target"
    tunnelIdTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelIdTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelIdTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelIdTarget.EntityData.Children = make(map[string]types.YChild)
    tunnelIdTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelIdTarget.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", tunnelIdTarget.TunnelId}
    return &(tunnelIdTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_Ipv4PseudowireTarget
// IPv4 pseudowire target
type Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_Ipv4PseudowireTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Virtual circuit ID. The type is interface{} with range: 0..4294967295.
    VirtualCircuitId interface{}
}

func (ipv4PseudowireTarget *Ipsla_OperationData_Operations_Operation_Lpd_Statistics_Aggregated_Hours_Hour_TargetAddress_Ipv4PseudowireTarget) GetEntityData() *types.CommonEntityData {
    ipv4PseudowireTarget.EntityData.YFilter = ipv4PseudowireTarget.YFilter
    ipv4PseudowireTarget.EntityData.YangName = "ipv4-pseudowire-target"
    ipv4PseudowireTarget.EntityData.BundleName = "cisco_ios_xr"
    ipv4PseudowireTarget.EntityData.ParentYangName = "target-address"
    ipv4PseudowireTarget.EntityData.SegmentPath = "ipv4-pseudowire-target"
    ipv4PseudowireTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PseudowireTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PseudowireTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PseudowireTarget.EntityData.Children = make(map[string]types.YChild)
    ipv4PseudowireTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4PseudowireTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4PseudowireTarget.Address}
    ipv4PseudowireTarget.EntityData.Leafs["virtual-circuit-id"] = types.YLeaf{"VirtualCircuitId", ipv4PseudowireTarget.VirtualCircuitId}
    return &(ipv4PseudowireTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Status
// Operational status of LPD group
type Ipsla_OperationData_Operations_Operation_Lpd_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational path state in LPD group.
    LpdPaths Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths

    // Operational status of LPD group.
    State Ipsla_OperationData_Operations_Operation_Lpd_Status_State
}

func (status *Ipsla_OperationData_Operations_Operation_Lpd_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xr"
    status.EntityData.ParentYangName = "lpd"
    status.EntityData.SegmentPath = "status"
    status.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    status.EntityData.Children = make(map[string]types.YChild)
    status.EntityData.Children["lpd-paths"] = types.YChild{"LpdPaths", &status.LpdPaths}
    status.EntityData.Children["state"] = types.YChild{"State", &status.State}
    status.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(status.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths
// Operational path state in LPD group
type Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current operational path state in LPD group. The type is slice of
    // Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths_LpdPath.
    LpdPath []Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths_LpdPath
}

func (lpdPaths *Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths) GetEntityData() *types.CommonEntityData {
    lpdPaths.EntityData.YFilter = lpdPaths.YFilter
    lpdPaths.EntityData.YangName = "lpd-paths"
    lpdPaths.EntityData.BundleName = "cisco_ios_xr"
    lpdPaths.EntityData.ParentYangName = "status"
    lpdPaths.EntityData.SegmentPath = "lpd-paths"
    lpdPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpdPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpdPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpdPaths.EntityData.Children = make(map[string]types.YChild)
    lpdPaths.EntityData.Children["lpd-path"] = types.YChild{"LpdPath", nil}
    for i := range lpdPaths.LpdPath {
        lpdPaths.EntityData.Children[types.GetSegmentPath(&lpdPaths.LpdPath[i])] = types.YChild{"LpdPath", &lpdPaths.LpdPath[i]}
    }
    lpdPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lpdPaths.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths_LpdPath
// Current operational path state in LPD
// group
type Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths_LpdPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPD path index. The type is interface{} with
    // range: -2147483648..2147483647.
    PathIndex interface{}

    // Path status. The type is IpslaMplsLpdPathDiscoveryStatus.
    PathStatus interface{}

    // LPD path identifier.
    PathId Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths_LpdPath_PathId
}

func (lpdPath *Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths_LpdPath) GetEntityData() *types.CommonEntityData {
    lpdPath.EntityData.YFilter = lpdPath.YFilter
    lpdPath.EntityData.YangName = "lpd-path"
    lpdPath.EntityData.BundleName = "cisco_ios_xr"
    lpdPath.EntityData.ParentYangName = "lpd-paths"
    lpdPath.EntityData.SegmentPath = "lpd-path" + "[path-index='" + fmt.Sprintf("%v", lpdPath.PathIndex) + "']"
    lpdPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpdPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpdPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpdPath.EntityData.Children = make(map[string]types.YChild)
    lpdPath.EntityData.Children["path-id"] = types.YChild{"PathId", &lpdPath.PathId}
    lpdPath.EntityData.Leafs = make(map[string]types.YLeaf)
    lpdPath.EntityData.Leafs["path-index"] = types.YLeaf{"PathIndex", lpdPath.PathIndex}
    lpdPath.EntityData.Leafs["path-status"] = types.YLeaf{"PathStatus", lpdPath.PathStatus}
    return &(lpdPath.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths_LpdPath_PathId
// LPD path identifier
type Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths_LpdPath_PathId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP selector. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LspSelector interface{}

    // Output interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    OutputInterface interface{}

    // Nexthop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}

    // Downstream label stacks. The type is slice of interface{} with range:
    // 0..4294967295.
    DownstreamLabel []interface{}
}

func (pathId *Ipsla_OperationData_Operations_Operation_Lpd_Status_LpdPaths_LpdPath_PathId) GetEntityData() *types.CommonEntityData {
    pathId.EntityData.YFilter = pathId.YFilter
    pathId.EntityData.YangName = "path-id"
    pathId.EntityData.BundleName = "cisco_ios_xr"
    pathId.EntityData.ParentYangName = "lpd-path"
    pathId.EntityData.SegmentPath = "path-id"
    pathId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathId.EntityData.Children = make(map[string]types.YChild)
    pathId.EntityData.Leafs = make(map[string]types.YLeaf)
    pathId.EntityData.Leafs["lsp-selector"] = types.YLeaf{"LspSelector", pathId.LspSelector}
    pathId.EntityData.Leafs["output-interface"] = types.YLeaf{"OutputInterface", pathId.OutputInterface}
    pathId.EntityData.Leafs["nexthop-address"] = types.YLeaf{"NexthopAddress", pathId.NexthopAddress}
    pathId.EntityData.Leafs["downstream-label"] = types.YLeaf{"DownstreamLabel", pathId.DownstreamLabel}
    return &(pathId.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Status_State
// Operational status of LPD group
type Ipsla_OperationData_Operations_Operation_Lpd_Status_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLSLM monitor ID. The type is interface{} with range: 0..4294967295.
    MonitorId interface{}

    // Latest LPD mode. The type is IpslaMplsLpdDiscoveryModeEnum.
    DiscoveryMode interface{}

    // Latest start time. The type is interface{} with range:
    // 0..18446744073709551615.
    StartTime interface{}

    // Latest return code. The type is IpslaMplsLpdRetCode.
    ReturnCode interface{}

    // Latest completion time. The type is interface{} with range: 0..4294967295.
    CompletionTime interface{}

    // Number of discovered paths. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Target for LPD.
    TargetAddress Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress
}

func (state *Ipsla_OperationData_Operations_Operation_Lpd_Status_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "status"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["target-address"] = types.YChild{"TargetAddress", &state.TargetAddress}
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["monitor-id"] = types.YLeaf{"MonitorId", state.MonitorId}
    state.EntityData.Leafs["discovery-mode"] = types.YLeaf{"DiscoveryMode", state.DiscoveryMode}
    state.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", state.StartTime}
    state.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", state.ReturnCode}
    state.EntityData.Leafs["completion-time"] = types.YLeaf{"CompletionTime", state.CompletionTime}
    state.EntityData.Leafs["path-count"] = types.YLeaf{"PathCount", state.PathCount}
    return &(state.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress
// Target for LPD
type Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TargetType. The type is IpslaTargetTypeEnum.
    TargetType interface{}

    // IPv4 address target. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4AddressTarget interface{}

    // IPv6 address target. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6AddressTarget interface{}

    // IPv4 prefix target.
    Ipv4PrefixTarget Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_Ipv4PrefixTarget

    // Tunnel ID target.
    TunnelIdTarget Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_TunnelIdTarget

    // IPv4 pseudowire target.
    Ipv4PseudowireTarget Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_Ipv4PseudowireTarget
}

func (targetAddress *Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress) GetEntityData() *types.CommonEntityData {
    targetAddress.EntityData.YFilter = targetAddress.YFilter
    targetAddress.EntityData.YangName = "target-address"
    targetAddress.EntityData.BundleName = "cisco_ios_xr"
    targetAddress.EntityData.ParentYangName = "state"
    targetAddress.EntityData.SegmentPath = "target-address"
    targetAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetAddress.EntityData.Children = make(map[string]types.YChild)
    targetAddress.EntityData.Children["ipv4-prefix-target"] = types.YChild{"Ipv4PrefixTarget", &targetAddress.Ipv4PrefixTarget}
    targetAddress.EntityData.Children["tunnel-id-target"] = types.YChild{"TunnelIdTarget", &targetAddress.TunnelIdTarget}
    targetAddress.EntityData.Children["ipv4-pseudowire-target"] = types.YChild{"Ipv4PseudowireTarget", &targetAddress.Ipv4PseudowireTarget}
    targetAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    targetAddress.EntityData.Leafs["target-type"] = types.YLeaf{"TargetType", targetAddress.TargetType}
    targetAddress.EntityData.Leafs["ipv4-address-target"] = types.YLeaf{"Ipv4AddressTarget", targetAddress.Ipv4AddressTarget}
    targetAddress.EntityData.Leafs["ipv6-address-target"] = types.YLeaf{"Ipv6AddressTarget", targetAddress.Ipv6AddressTarget}
    return &(targetAddress.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_Ipv4PrefixTarget
// IPv4 prefix target
type Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_Ipv4PrefixTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Mask length. The type is interface{} with range: 0..255.
    MaskLength interface{}
}

func (ipv4PrefixTarget *Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_Ipv4PrefixTarget) GetEntityData() *types.CommonEntityData {
    ipv4PrefixTarget.EntityData.YFilter = ipv4PrefixTarget.YFilter
    ipv4PrefixTarget.EntityData.YangName = "ipv4-prefix-target"
    ipv4PrefixTarget.EntityData.BundleName = "cisco_ios_xr"
    ipv4PrefixTarget.EntityData.ParentYangName = "target-address"
    ipv4PrefixTarget.EntityData.SegmentPath = "ipv4-prefix-target"
    ipv4PrefixTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PrefixTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PrefixTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PrefixTarget.EntityData.Children = make(map[string]types.YChild)
    ipv4PrefixTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4PrefixTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4PrefixTarget.Address}
    ipv4PrefixTarget.EntityData.Leafs["mask-length"] = types.YLeaf{"MaskLength", ipv4PrefixTarget.MaskLength}
    return &(ipv4PrefixTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_TunnelIdTarget
// Tunnel ID target
type Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_TunnelIdTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}
}

func (tunnelIdTarget *Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_TunnelIdTarget) GetEntityData() *types.CommonEntityData {
    tunnelIdTarget.EntityData.YFilter = tunnelIdTarget.YFilter
    tunnelIdTarget.EntityData.YangName = "tunnel-id-target"
    tunnelIdTarget.EntityData.BundleName = "cisco_ios_xr"
    tunnelIdTarget.EntityData.ParentYangName = "target-address"
    tunnelIdTarget.EntityData.SegmentPath = "tunnel-id-target"
    tunnelIdTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelIdTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelIdTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelIdTarget.EntityData.Children = make(map[string]types.YChild)
    tunnelIdTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelIdTarget.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", tunnelIdTarget.TunnelId}
    return &(tunnelIdTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_Ipv4PseudowireTarget
// IPv4 pseudowire target
type Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_Ipv4PseudowireTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Virtual circuit ID. The type is interface{} with range: 0..4294967295.
    VirtualCircuitId interface{}
}

func (ipv4PseudowireTarget *Ipsla_OperationData_Operations_Operation_Lpd_Status_State_TargetAddress_Ipv4PseudowireTarget) GetEntityData() *types.CommonEntityData {
    ipv4PseudowireTarget.EntityData.YFilter = ipv4PseudowireTarget.YFilter
    ipv4PseudowireTarget.EntityData.YangName = "ipv4-pseudowire-target"
    ipv4PseudowireTarget.EntityData.BundleName = "cisco_ios_xr"
    ipv4PseudowireTarget.EntityData.ParentYangName = "target-address"
    ipv4PseudowireTarget.EntityData.SegmentPath = "ipv4-pseudowire-target"
    ipv4PseudowireTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PseudowireTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PseudowireTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PseudowireTarget.EntityData.Children = make(map[string]types.YChild)
    ipv4PseudowireTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4PseudowireTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4PseudowireTarget.Address}
    ipv4PseudowireTarget.EntityData.Leafs["virtual-circuit-id"] = types.YLeaf{"VirtualCircuitId", ipv4PseudowireTarget.VirtualCircuitId}
    return &(ipv4PseudowireTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History
// Historical data for an operation
type Ipsla_OperationData_Operations_Operation_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Historical data with multiple hops along the path.
    Path Ipsla_OperationData_Operations_Operation_History_Path

    // Historical data for the destination node.
    Target Ipsla_OperationData_Operations_Operation_History_Target
}

func (history *Ipsla_OperationData_Operations_Operation_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "operation"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Children["path"] = types.YChild{"Path", &history.Path}
    history.EntityData.Children["target"] = types.YChild{"Target", &history.Target}
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(history.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path
// Historical data with multiple hops along the
// path
type Ipsla_OperationData_Operations_Operation_History_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tables of lives for an operation.
    Lifes Ipsla_OperationData_Operations_Operation_History_Path_Lifes
}

func (path *Ipsla_OperationData_Operations_Operation_History_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "history"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Children["lifes"] = types.YChild{"Lifes", &path.Lifes}
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(path.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path_Lifes
// Tables of lives for an operation
type Ipsla_OperationData_Operations_Operation_History_Path_Lifes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // History data for a particular life of the operation. The type is slice of
    // Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life.
    Life []Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life
}

func (lifes *Ipsla_OperationData_Operations_Operation_History_Path_Lifes) GetEntityData() *types.CommonEntityData {
    lifes.EntityData.YFilter = lifes.YFilter
    lifes.EntityData.YangName = "lifes"
    lifes.EntityData.BundleName = "cisco_ios_xr"
    lifes.EntityData.ParentYangName = "path"
    lifes.EntityData.SegmentPath = "lifes"
    lifes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lifes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lifes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lifes.EntityData.Children = make(map[string]types.YChild)
    lifes.EntityData.Children["life"] = types.YChild{"Life", nil}
    for i := range lifes.Life {
        lifes.EntityData.Children[types.GetSegmentPath(&lifes.Life[i])] = types.YChild{"Life", &lifes.Life[i]}
    }
    lifes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lifes.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life
// History data for a particular life of the
// operation
type Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Life Index. The type is interface{} with range:
    // -2147483648..2147483647.
    LifeIndex interface{}

    // Table of history buckets (samples) for a particular operation.
    Buckets Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets
}

func (life *Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life) GetEntityData() *types.CommonEntityData {
    life.EntityData.YFilter = life.YFilter
    life.EntityData.YangName = "life"
    life.EntityData.BundleName = "cisco_ios_xr"
    life.EntityData.ParentYangName = "lifes"
    life.EntityData.SegmentPath = "life" + "[life-index='" + fmt.Sprintf("%v", life.LifeIndex) + "']"
    life.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    life.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    life.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    life.EntityData.Children = make(map[string]types.YChild)
    life.EntityData.Children["buckets"] = types.YChild{"Buckets", &life.Buckets}
    life.EntityData.Leafs = make(map[string]types.YLeaf)
    life.EntityData.Leafs["life-index"] = types.YLeaf{"LifeIndex", life.LifeIndex}
    return &(life.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets
// Table of history buckets (samples) for a
// particular operation
type Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // History bucket for an operation. The type is slice of
    // Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket.
    Bucket []Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket
}

func (buckets *Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets) GetEntityData() *types.CommonEntityData {
    buckets.EntityData.YFilter = buckets.YFilter
    buckets.EntityData.YangName = "buckets"
    buckets.EntityData.BundleName = "cisco_ios_xr"
    buckets.EntityData.ParentYangName = "life"
    buckets.EntityData.SegmentPath = "buckets"
    buckets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    buckets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    buckets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    buckets.EntityData.Children = make(map[string]types.YChild)
    buckets.EntityData.Children["bucket"] = types.YChild{"Bucket", nil}
    for i := range buckets.Bucket {
        buckets.EntityData.Children[types.GetSegmentPath(&buckets.Bucket[i])] = types.YChild{"Bucket", &buckets.Bucket[i]}
    }
    buckets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(buckets.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket
// History bucket for an operation
type Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bucket Index. The type is interface{} with range:
    // -2147483648..2147483647.
    BucketIndex interface{}

    // Table of samples for a particular cycle.
    Samples Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples
}

func (bucket *Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket) GetEntityData() *types.CommonEntityData {
    bucket.EntityData.YFilter = bucket.YFilter
    bucket.EntityData.YangName = "bucket"
    bucket.EntityData.BundleName = "cisco_ios_xr"
    bucket.EntityData.ParentYangName = "buckets"
    bucket.EntityData.SegmentPath = "bucket" + "[bucket-index='" + fmt.Sprintf("%v", bucket.BucketIndex) + "']"
    bucket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bucket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bucket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bucket.EntityData.Children = make(map[string]types.YChild)
    bucket.EntityData.Children["samples"] = types.YChild{"Samples", &bucket.Samples}
    bucket.EntityData.Leafs = make(map[string]types.YLeaf)
    bucket.EntityData.Leafs["bucket-index"] = types.YLeaf{"BucketIndex", bucket.BucketIndex}
    return &(bucket.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples
// Table of samples for a particular cycle
type Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data sample for particular cycle. The type is slice of
    // Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample.
    Sample []Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample
}

func (samples *Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples) GetEntityData() *types.CommonEntityData {
    samples.EntityData.YFilter = samples.YFilter
    samples.EntityData.YangName = "samples"
    samples.EntityData.BundleName = "cisco_ios_xr"
    samples.EntityData.ParentYangName = "bucket"
    samples.EntityData.SegmentPath = "samples"
    samples.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    samples.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    samples.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    samples.EntityData.Children = make(map[string]types.YChild)
    samples.EntityData.Children["sample"] = types.YChild{"Sample", nil}
    for i := range samples.Sample {
        samples.EntityData.Children[types.GetSegmentPath(&samples.Sample[i])] = types.YChild{"Sample", &samples.Sample[i]}
    }
    samples.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(samples.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample
// Data sample for particular cycle
type Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Sample Index. The type is interface{} with range:
    // -2147483648..2147483647.
    SampleIndex interface{}

    // Sample Start Time expressed in msec since00:00 :00 UTC, January 1, 1970.
    // The type is interface{} with range: 0..18446744073709551615.
    StartTime interface{}

    // Round Trip Time (milliseconds). The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    ResponseTime interface{}

    // Response Return Code. The type is IpslaRetCode.
    ReturnCode interface{}

    // Target for the operation.
    TargetAddress Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress
}

func (sample *Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "samples"
    sample.EntityData.SegmentPath = "sample" + "[sample-index='" + fmt.Sprintf("%v", sample.SampleIndex) + "']"
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = make(map[string]types.YChild)
    sample.EntityData.Children["target-address"] = types.YChild{"TargetAddress", &sample.TargetAddress}
    sample.EntityData.Leafs = make(map[string]types.YLeaf)
    sample.EntityData.Leafs["sample-index"] = types.YLeaf{"SampleIndex", sample.SampleIndex}
    sample.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", sample.StartTime}
    sample.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", sample.ResponseTime}
    sample.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", sample.ReturnCode}
    return &(sample.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress
// Target for the operation
type Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TargetType. The type is IpslaTargetTypeEnum.
    TargetType interface{}

    // IPv4 address target. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4AddressTarget interface{}

    // IPv6 address target. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6AddressTarget interface{}

    // IPv4 prefix target.
    Ipv4PrefixTarget Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_Ipv4PrefixTarget

    // Tunnel ID target.
    TunnelIdTarget Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_TunnelIdTarget

    // IPv4 pseudowire target.
    Ipv4PseudowireTarget Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_Ipv4PseudowireTarget
}

func (targetAddress *Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress) GetEntityData() *types.CommonEntityData {
    targetAddress.EntityData.YFilter = targetAddress.YFilter
    targetAddress.EntityData.YangName = "target-address"
    targetAddress.EntityData.BundleName = "cisco_ios_xr"
    targetAddress.EntityData.ParentYangName = "sample"
    targetAddress.EntityData.SegmentPath = "target-address"
    targetAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetAddress.EntityData.Children = make(map[string]types.YChild)
    targetAddress.EntityData.Children["ipv4-prefix-target"] = types.YChild{"Ipv4PrefixTarget", &targetAddress.Ipv4PrefixTarget}
    targetAddress.EntityData.Children["tunnel-id-target"] = types.YChild{"TunnelIdTarget", &targetAddress.TunnelIdTarget}
    targetAddress.EntityData.Children["ipv4-pseudowire-target"] = types.YChild{"Ipv4PseudowireTarget", &targetAddress.Ipv4PseudowireTarget}
    targetAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    targetAddress.EntityData.Leafs["target-type"] = types.YLeaf{"TargetType", targetAddress.TargetType}
    targetAddress.EntityData.Leafs["ipv4-address-target"] = types.YLeaf{"Ipv4AddressTarget", targetAddress.Ipv4AddressTarget}
    targetAddress.EntityData.Leafs["ipv6-address-target"] = types.YLeaf{"Ipv6AddressTarget", targetAddress.Ipv6AddressTarget}
    return &(targetAddress.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_Ipv4PrefixTarget
// IPv4 prefix target
type Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_Ipv4PrefixTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Mask length. The type is interface{} with range: 0..255.
    MaskLength interface{}
}

func (ipv4PrefixTarget *Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_Ipv4PrefixTarget) GetEntityData() *types.CommonEntityData {
    ipv4PrefixTarget.EntityData.YFilter = ipv4PrefixTarget.YFilter
    ipv4PrefixTarget.EntityData.YangName = "ipv4-prefix-target"
    ipv4PrefixTarget.EntityData.BundleName = "cisco_ios_xr"
    ipv4PrefixTarget.EntityData.ParentYangName = "target-address"
    ipv4PrefixTarget.EntityData.SegmentPath = "ipv4-prefix-target"
    ipv4PrefixTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PrefixTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PrefixTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PrefixTarget.EntityData.Children = make(map[string]types.YChild)
    ipv4PrefixTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4PrefixTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4PrefixTarget.Address}
    ipv4PrefixTarget.EntityData.Leafs["mask-length"] = types.YLeaf{"MaskLength", ipv4PrefixTarget.MaskLength}
    return &(ipv4PrefixTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_TunnelIdTarget
// Tunnel ID target
type Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_TunnelIdTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}
}

func (tunnelIdTarget *Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_TunnelIdTarget) GetEntityData() *types.CommonEntityData {
    tunnelIdTarget.EntityData.YFilter = tunnelIdTarget.YFilter
    tunnelIdTarget.EntityData.YangName = "tunnel-id-target"
    tunnelIdTarget.EntityData.BundleName = "cisco_ios_xr"
    tunnelIdTarget.EntityData.ParentYangName = "target-address"
    tunnelIdTarget.EntityData.SegmentPath = "tunnel-id-target"
    tunnelIdTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelIdTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelIdTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelIdTarget.EntityData.Children = make(map[string]types.YChild)
    tunnelIdTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelIdTarget.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", tunnelIdTarget.TunnelId}
    return &(tunnelIdTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_Ipv4PseudowireTarget
// IPv4 pseudowire target
type Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_Ipv4PseudowireTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Virtual circuit ID. The type is interface{} with range: 0..4294967295.
    VirtualCircuitId interface{}
}

func (ipv4PseudowireTarget *Ipsla_OperationData_Operations_Operation_History_Path_Lifes_Life_Buckets_Bucket_Samples_Sample_TargetAddress_Ipv4PseudowireTarget) GetEntityData() *types.CommonEntityData {
    ipv4PseudowireTarget.EntityData.YFilter = ipv4PseudowireTarget.YFilter
    ipv4PseudowireTarget.EntityData.YangName = "ipv4-pseudowire-target"
    ipv4PseudowireTarget.EntityData.BundleName = "cisco_ios_xr"
    ipv4PseudowireTarget.EntityData.ParentYangName = "target-address"
    ipv4PseudowireTarget.EntityData.SegmentPath = "ipv4-pseudowire-target"
    ipv4PseudowireTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PseudowireTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PseudowireTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PseudowireTarget.EntityData.Children = make(map[string]types.YChild)
    ipv4PseudowireTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4PseudowireTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4PseudowireTarget.Address}
    ipv4PseudowireTarget.EntityData.Leafs["virtual-circuit-id"] = types.YLeaf{"VirtualCircuitId", ipv4PseudowireTarget.VirtualCircuitId}
    return &(ipv4PseudowireTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Target
// Historical data for the destination node
type Ipsla_OperationData_Operations_Operation_History_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tables of lives for an operation.
    Lifes Ipsla_OperationData_Operations_Operation_History_Target_Lifes
}

func (target *Ipsla_OperationData_Operations_Operation_History_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "cisco_ios_xr"
    target.EntityData.ParentYangName = "history"
    target.EntityData.SegmentPath = "target"
    target.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    target.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    target.EntityData.Children = make(map[string]types.YChild)
    target.EntityData.Children["lifes"] = types.YChild{"Lifes", &target.Lifes}
    target.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(target.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Target_Lifes
// Tables of lives for an operation
type Ipsla_OperationData_Operations_Operation_History_Target_Lifes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a particular life of the operation. The type is slice
    // of Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life.
    Life []Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life
}

func (lifes *Ipsla_OperationData_Operations_Operation_History_Target_Lifes) GetEntityData() *types.CommonEntityData {
    lifes.EntityData.YFilter = lifes.YFilter
    lifes.EntityData.YangName = "lifes"
    lifes.EntityData.BundleName = "cisco_ios_xr"
    lifes.EntityData.ParentYangName = "target"
    lifes.EntityData.SegmentPath = "lifes"
    lifes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lifes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lifes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lifes.EntityData.Children = make(map[string]types.YChild)
    lifes.EntityData.Children["life"] = types.YChild{"Life", nil}
    for i := range lifes.Life {
        lifes.EntityData.Children[types.GetSegmentPath(&lifes.Life[i])] = types.YChild{"Life", &lifes.Life[i]}
    }
    lifes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lifes.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life
// Operational data for a particular life of
// the operation
type Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Life Index. The type is interface{} with range:
    // -2147483648..2147483647.
    LifeIndex interface{}

    // Table of history buckets (samples) for a particular operation.
    Buckets Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets
}

func (life *Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life) GetEntityData() *types.CommonEntityData {
    life.EntityData.YFilter = life.YFilter
    life.EntityData.YangName = "life"
    life.EntityData.BundleName = "cisco_ios_xr"
    life.EntityData.ParentYangName = "lifes"
    life.EntityData.SegmentPath = "life" + "[life-index='" + fmt.Sprintf("%v", life.LifeIndex) + "']"
    life.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    life.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    life.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    life.EntityData.Children = make(map[string]types.YChild)
    life.EntityData.Children["buckets"] = types.YChild{"Buckets", &life.Buckets}
    life.EntityData.Leafs = make(map[string]types.YLeaf)
    life.EntityData.Leafs["life-index"] = types.YLeaf{"LifeIndex", life.LifeIndex}
    return &(life.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets
// Table of history buckets (samples) for a
// particular operation
type Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // History bucket for an operation. The type is slice of
    // Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket.
    Bucket []Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket
}

func (buckets *Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets) GetEntityData() *types.CommonEntityData {
    buckets.EntityData.YFilter = buckets.YFilter
    buckets.EntityData.YangName = "buckets"
    buckets.EntityData.BundleName = "cisco_ios_xr"
    buckets.EntityData.ParentYangName = "life"
    buckets.EntityData.SegmentPath = "buckets"
    buckets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    buckets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    buckets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    buckets.EntityData.Children = make(map[string]types.YChild)
    buckets.EntityData.Children["bucket"] = types.YChild{"Bucket", nil}
    for i := range buckets.Bucket {
        buckets.EntityData.Children[types.GetSegmentPath(&buckets.Bucket[i])] = types.YChild{"Bucket", &buckets.Bucket[i]}
    }
    buckets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(buckets.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket
// History bucket for an operation
type Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bucket Index. The type is interface{} with range:
    // -2147483648..2147483647.
    BucketIndex interface{}

    // Sample Start Time expressed in msec since00:00 :00 UTC, January 1, 1970.
    // The type is interface{} with range: 0..18446744073709551615.
    StartTime interface{}

    // Round Trip Time (milliseconds). The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    ResponseTime interface{}

    // Response Return Code. The type is IpslaRetCode.
    ReturnCode interface{}

    // Target for the operation.
    TargetAddress Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress
}

func (bucket *Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket) GetEntityData() *types.CommonEntityData {
    bucket.EntityData.YFilter = bucket.YFilter
    bucket.EntityData.YangName = "bucket"
    bucket.EntityData.BundleName = "cisco_ios_xr"
    bucket.EntityData.ParentYangName = "buckets"
    bucket.EntityData.SegmentPath = "bucket" + "[bucket-index='" + fmt.Sprintf("%v", bucket.BucketIndex) + "']"
    bucket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bucket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bucket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bucket.EntityData.Children = make(map[string]types.YChild)
    bucket.EntityData.Children["target-address"] = types.YChild{"TargetAddress", &bucket.TargetAddress}
    bucket.EntityData.Leafs = make(map[string]types.YLeaf)
    bucket.EntityData.Leafs["bucket-index"] = types.YLeaf{"BucketIndex", bucket.BucketIndex}
    bucket.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", bucket.StartTime}
    bucket.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", bucket.ResponseTime}
    bucket.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", bucket.ReturnCode}
    return &(bucket.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress
// Target for the operation
type Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TargetType. The type is IpslaTargetTypeEnum.
    TargetType interface{}

    // IPv4 address target. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4AddressTarget interface{}

    // IPv6 address target. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6AddressTarget interface{}

    // IPv4 prefix target.
    Ipv4PrefixTarget Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_Ipv4PrefixTarget

    // Tunnel ID target.
    TunnelIdTarget Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_TunnelIdTarget

    // IPv4 pseudowire target.
    Ipv4PseudowireTarget Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_Ipv4PseudowireTarget
}

func (targetAddress *Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress) GetEntityData() *types.CommonEntityData {
    targetAddress.EntityData.YFilter = targetAddress.YFilter
    targetAddress.EntityData.YangName = "target-address"
    targetAddress.EntityData.BundleName = "cisco_ios_xr"
    targetAddress.EntityData.ParentYangName = "bucket"
    targetAddress.EntityData.SegmentPath = "target-address"
    targetAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetAddress.EntityData.Children = make(map[string]types.YChild)
    targetAddress.EntityData.Children["ipv4-prefix-target"] = types.YChild{"Ipv4PrefixTarget", &targetAddress.Ipv4PrefixTarget}
    targetAddress.EntityData.Children["tunnel-id-target"] = types.YChild{"TunnelIdTarget", &targetAddress.TunnelIdTarget}
    targetAddress.EntityData.Children["ipv4-pseudowire-target"] = types.YChild{"Ipv4PseudowireTarget", &targetAddress.Ipv4PseudowireTarget}
    targetAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    targetAddress.EntityData.Leafs["target-type"] = types.YLeaf{"TargetType", targetAddress.TargetType}
    targetAddress.EntityData.Leafs["ipv4-address-target"] = types.YLeaf{"Ipv4AddressTarget", targetAddress.Ipv4AddressTarget}
    targetAddress.EntityData.Leafs["ipv6-address-target"] = types.YLeaf{"Ipv6AddressTarget", targetAddress.Ipv6AddressTarget}
    return &(targetAddress.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_Ipv4PrefixTarget
// IPv4 prefix target
type Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_Ipv4PrefixTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Mask length. The type is interface{} with range: 0..255.
    MaskLength interface{}
}

func (ipv4PrefixTarget *Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_Ipv4PrefixTarget) GetEntityData() *types.CommonEntityData {
    ipv4PrefixTarget.EntityData.YFilter = ipv4PrefixTarget.YFilter
    ipv4PrefixTarget.EntityData.YangName = "ipv4-prefix-target"
    ipv4PrefixTarget.EntityData.BundleName = "cisco_ios_xr"
    ipv4PrefixTarget.EntityData.ParentYangName = "target-address"
    ipv4PrefixTarget.EntityData.SegmentPath = "ipv4-prefix-target"
    ipv4PrefixTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PrefixTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PrefixTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PrefixTarget.EntityData.Children = make(map[string]types.YChild)
    ipv4PrefixTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4PrefixTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4PrefixTarget.Address}
    ipv4PrefixTarget.EntityData.Leafs["mask-length"] = types.YLeaf{"MaskLength", ipv4PrefixTarget.MaskLength}
    return &(ipv4PrefixTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_TunnelIdTarget
// Tunnel ID target
type Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_TunnelIdTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}
}

func (tunnelIdTarget *Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_TunnelIdTarget) GetEntityData() *types.CommonEntityData {
    tunnelIdTarget.EntityData.YFilter = tunnelIdTarget.YFilter
    tunnelIdTarget.EntityData.YangName = "tunnel-id-target"
    tunnelIdTarget.EntityData.BundleName = "cisco_ios_xr"
    tunnelIdTarget.EntityData.ParentYangName = "target-address"
    tunnelIdTarget.EntityData.SegmentPath = "tunnel-id-target"
    tunnelIdTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelIdTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelIdTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelIdTarget.EntityData.Children = make(map[string]types.YChild)
    tunnelIdTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelIdTarget.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", tunnelIdTarget.TunnelId}
    return &(tunnelIdTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_Ipv4PseudowireTarget
// IPv4 pseudowire target
type Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_Ipv4PseudowireTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Virtual circuit ID. The type is interface{} with range: 0..4294967295.
    VirtualCircuitId interface{}
}

func (ipv4PseudowireTarget *Ipsla_OperationData_Operations_Operation_History_Target_Lifes_Life_Buckets_Bucket_TargetAddress_Ipv4PseudowireTarget) GetEntityData() *types.CommonEntityData {
    ipv4PseudowireTarget.EntityData.YFilter = ipv4PseudowireTarget.YFilter
    ipv4PseudowireTarget.EntityData.YangName = "ipv4-pseudowire-target"
    ipv4PseudowireTarget.EntityData.BundleName = "cisco_ios_xr"
    ipv4PseudowireTarget.EntityData.ParentYangName = "target-address"
    ipv4PseudowireTarget.EntityData.SegmentPath = "ipv4-pseudowire-target"
    ipv4PseudowireTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PseudowireTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PseudowireTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PseudowireTarget.EntityData.Children = make(map[string]types.YChild)
    ipv4PseudowireTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4PseudowireTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4PseudowireTarget.Address}
    ipv4PseudowireTarget.EntityData.Leafs["virtual-circuit-id"] = types.YLeaf{"VirtualCircuitId", ipv4PseudowireTarget.VirtualCircuitId}
    return &(ipv4PseudowireTarget.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics
// Statistics collected for an operation
type Ipsla_OperationData_Operations_Operation_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics collected during the last sampling cycle of the operation.
    Latest Ipsla_OperationData_Operations_Operation_Statistics_Latest

    // Statistics aggregated for data collected over time intervals.
    Aggregated Ipsla_OperationData_Operations_Operation_Statistics_Aggregated
}

func (statistics *Ipsla_OperationData_Operations_Operation_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "operation"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["latest"] = types.YChild{"Latest", &statistics.Latest}
    statistics.EntityData.Children["aggregated"] = types.YChild{"Aggregated", &statistics.Aggregated}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest
// Statistics collected during the last
// sampling cycle of the operation
type Ipsla_OperationData_Operations_Operation_Statistics_Latest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Latest statistics for the target node.
    Target Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target

    // Latest statistics for hops in a path-enabled operation.
    Hops Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops

    // List of latest LPD paths.
    LpdPaths Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths
}

func (latest *Ipsla_OperationData_Operations_Operation_Statistics_Latest) GetEntityData() *types.CommonEntityData {
    latest.EntityData.YFilter = latest.YFilter
    latest.EntityData.YangName = "latest"
    latest.EntityData.BundleName = "cisco_ios_xr"
    latest.EntityData.ParentYangName = "statistics"
    latest.EntityData.SegmentPath = "latest"
    latest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    latest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    latest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    latest.EntityData.Children = make(map[string]types.YChild)
    latest.EntityData.Children["target"] = types.YChild{"Target", &latest.Target}
    latest.EntityData.Children["hops"] = types.YChild{"Hops", &latest.Hops}
    latest.EntityData.Children["lpd-paths"] = types.YChild{"LpdPaths", &latest.LpdPaths}
    latest.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(latest.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target
// Latest statistics for the target node
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Common Stats.
    CommonStats Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_CommonStats

    // Operation Specific Stats.
    SpecificStats Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats
}

func (target *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "cisco_ios_xr"
    target.EntityData.ParentYangName = "latest"
    target.EntityData.SegmentPath = "target"
    target.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    target.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    target.EntityData.Children = make(map[string]types.YChild)
    target.EntityData.Children["common-stats"] = types.YChild{"CommonStats", &target.CommonStats}
    target.EntityData.Children["specific-stats"] = types.YChild{"SpecificStats", &target.SpecificStats}
    target.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(target.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_CommonStats
// Common Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_CommonStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation Time. The type is interface{} with range:
    // 0..18446744073709551615.
    OperationTime interface{}

    // Return code. The type is IpslaRetCode.
    ReturnCode interface{}

    // Number of RTT samples used for the statistics. The type is interface{} with
    // range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of updates processed. The type is interface{} with range:
    // 0..4294967295.
    UpdateCount interface{}

    // Number of updates with Okay return code. The type is interface{} with
    // range: 0..4294967295.
    OkCount interface{}

    // Number of updates with Disconnected return code. The type is interface{}
    // with range: 0..4294967295.
    DisconnectCount interface{}

    // Number of updates with Timeout return code. The type is interface{} with
    // range: 0..4294967295.
    TimeoutCount interface{}

    // Number of updates with Busy return code. The type is interface{} with
    // range: 0..4294967295.
    BusyCount interface{}

    // Number of updates with NotConnected return code. The type is interface{}
    // with range: 0..4294967295.
    NoConnectionCount interface{}

    // Number of updates with Dropped return code. The type is interface{} with
    // range: 0..4294967295.
    DroppedCount interface{}

    // Number of updates with InternalError return code. The type is interface{}
    // with range: 0..4294967295.
    InternalErrorCount interface{}

    // Number of updates with SeqError return code. The type is interface{} with
    // range: 0..4294967295.
    SequenceErrorCount interface{}

    // Number of updates with VerifyError return code. The type is interface{}
    // with range: 0..4294967295.
    VerifyErrorCount interface{}
}

func (commonStats *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_CommonStats) GetEntityData() *types.CommonEntityData {
    commonStats.EntityData.YFilter = commonStats.YFilter
    commonStats.EntityData.YangName = "common-stats"
    commonStats.EntityData.BundleName = "cisco_ios_xr"
    commonStats.EntityData.ParentYangName = "target"
    commonStats.EntityData.SegmentPath = "common-stats"
    commonStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonStats.EntityData.Children = make(map[string]types.YChild)
    commonStats.EntityData.Leafs = make(map[string]types.YLeaf)
    commonStats.EntityData.Leafs["operation-time"] = types.YLeaf{"OperationTime", commonStats.OperationTime}
    commonStats.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", commonStats.ReturnCode}
    commonStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", commonStats.ResponseTimeCount}
    commonStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", commonStats.ResponseTime}
    commonStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", commonStats.MinResponseTime}
    commonStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", commonStats.MaxResponseTime}
    commonStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", commonStats.SumResponseTime}
    commonStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", commonStats.Sum2ResponseTime}
    commonStats.EntityData.Leafs["update-count"] = types.YLeaf{"UpdateCount", commonStats.UpdateCount}
    commonStats.EntityData.Leafs["ok-count"] = types.YLeaf{"OkCount", commonStats.OkCount}
    commonStats.EntityData.Leafs["disconnect-count"] = types.YLeaf{"DisconnectCount", commonStats.DisconnectCount}
    commonStats.EntityData.Leafs["timeout-count"] = types.YLeaf{"TimeoutCount", commonStats.TimeoutCount}
    commonStats.EntityData.Leafs["busy-count"] = types.YLeaf{"BusyCount", commonStats.BusyCount}
    commonStats.EntityData.Leafs["no-connection-count"] = types.YLeaf{"NoConnectionCount", commonStats.NoConnectionCount}
    commonStats.EntityData.Leafs["dropped-count"] = types.YLeaf{"DroppedCount", commonStats.DroppedCount}
    commonStats.EntityData.Leafs["internal-error-count"] = types.YLeaf{"InternalErrorCount", commonStats.InternalErrorCount}
    commonStats.EntityData.Leafs["sequence-error-count"] = types.YLeaf{"SequenceErrorCount", commonStats.SequenceErrorCount}
    commonStats.EntityData.Leafs["verify-error-count"] = types.YLeaf{"VerifyErrorCount", commonStats.VerifyErrorCount}
    return &(commonStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats
// Operation Specific Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // op type. The type is OpTypeEnum.
    OpType interface{}

    // icmp path jitter stats.
    IcmpPathJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats_IcmpPathJitterStats

    // udp jitter stats.
    UdpJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats_UdpJitterStats
}

func (specificStats *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats) GetEntityData() *types.CommonEntityData {
    specificStats.EntityData.YFilter = specificStats.YFilter
    specificStats.EntityData.YangName = "specific-stats"
    specificStats.EntityData.BundleName = "cisco_ios_xr"
    specificStats.EntityData.ParentYangName = "target"
    specificStats.EntityData.SegmentPath = "specific-stats"
    specificStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificStats.EntityData.Children = make(map[string]types.YChild)
    specificStats.EntityData.Children["icmp-path-jitter-stats"] = types.YChild{"IcmpPathJitterStats", &specificStats.IcmpPathJitterStats}
    specificStats.EntityData.Children["udp-jitter-stats"] = types.YChild{"UdpJitterStats", &specificStats.UdpJitterStats}
    specificStats.EntityData.Leafs = make(map[string]types.YLeaf)
    specificStats.EntityData.Leafs["op-type"] = types.YLeaf{"OpType", specificStats.OpType}
    return &(specificStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats_IcmpPathJitterStats
// icmp path jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats_IcmpPathJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address of the source. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // IP Address of the destination. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestAddress interface{}

    // IP address of the hop in the path. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}

    // Interval between echos in ms. The type is interface{} with range:
    // 0..4294967295.
    PacketInterval interface{}

    // Number of RTT samples  used for the statistics. The type is interface{}
    // with range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of Echo replies received . The type is interface{} with range:
    // 0..4294967295.
    PacketCount interface{}

    // Number of packets lost. The type is interface{} with range: 0..4294967295.
    PacketLossCount interface{}

    // Number of out of sequence packets. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceCount interface{}

    // Number of discarded samples. The type is interface{} with range:
    // 0..4294967295.
    DiscardedSampleCount interface{}

    // Number of packets with data corruption. The type is interface{} with range:
    // 0..4294967295.
    VerifyErrorsCount interface{}

    // Number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedErrorCount interface{}

    // Jitter value for this node in the path. The type is interface{} with range:
    // 0..4294967295.
    Jitter interface{}

    // Sum of positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterSum interface{}

    // Sum of squares of positive jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    PosJitterSum2 interface{}

    // Minimum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMin interface{}

    // Maximum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMax interface{}

    // Number of positive jitter values. The type is interface{} with range:
    // 0..4294967295.
    PosJitterCount interface{}

    // Sum of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterSum interface{}

    // Minimum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMin interface{}

    // Maximum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMax interface{}

    // Sum of squares of negative jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    NegJitterSum2 interface{}

    // Number of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterCount interface{}
}

func (icmpPathJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats_IcmpPathJitterStats) GetEntityData() *types.CommonEntityData {
    icmpPathJitterStats.EntityData.YFilter = icmpPathJitterStats.YFilter
    icmpPathJitterStats.EntityData.YangName = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.BundleName = "cisco_ios_xr"
    icmpPathJitterStats.EntityData.ParentYangName = "specific-stats"
    icmpPathJitterStats.EntityData.SegmentPath = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpPathJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpPathJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpPathJitterStats.EntityData.Children = make(map[string]types.YChild)
    icmpPathJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpPathJitterStats.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", icmpPathJitterStats.SourceAddress}
    icmpPathJitterStats.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", icmpPathJitterStats.DestAddress}
    icmpPathJitterStats.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", icmpPathJitterStats.HopAddress}
    icmpPathJitterStats.EntityData.Leafs["packet-interval"] = types.YLeaf{"PacketInterval", icmpPathJitterStats.PacketInterval}
    icmpPathJitterStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", icmpPathJitterStats.ResponseTimeCount}
    icmpPathJitterStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", icmpPathJitterStats.ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", icmpPathJitterStats.MinResponseTime}
    icmpPathJitterStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", icmpPathJitterStats.MaxResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", icmpPathJitterStats.SumResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", icmpPathJitterStats.Sum2ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["packet-count"] = types.YLeaf{"PacketCount", icmpPathJitterStats.PacketCount}
    icmpPathJitterStats.EntityData.Leafs["packet-loss-count"] = types.YLeaf{"PacketLossCount", icmpPathJitterStats.PacketLossCount}
    icmpPathJitterStats.EntityData.Leafs["out-of-sequence-count"] = types.YLeaf{"OutOfSequenceCount", icmpPathJitterStats.OutOfSequenceCount}
    icmpPathJitterStats.EntityData.Leafs["discarded-sample-count"] = types.YLeaf{"DiscardedSampleCount", icmpPathJitterStats.DiscardedSampleCount}
    icmpPathJitterStats.EntityData.Leafs["verify-errors-count"] = types.YLeaf{"VerifyErrorsCount", icmpPathJitterStats.VerifyErrorsCount}
    icmpPathJitterStats.EntityData.Leafs["dropped-error-count"] = types.YLeaf{"DroppedErrorCount", icmpPathJitterStats.DroppedErrorCount}
    icmpPathJitterStats.EntityData.Leafs["jitter"] = types.YLeaf{"Jitter", icmpPathJitterStats.Jitter}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum"] = types.YLeaf{"PosJitterSum", icmpPathJitterStats.PosJitterSum}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum2"] = types.YLeaf{"PosJitterSum2", icmpPathJitterStats.PosJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-min"] = types.YLeaf{"PosJitterMin", icmpPathJitterStats.PosJitterMin}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-max"] = types.YLeaf{"PosJitterMax", icmpPathJitterStats.PosJitterMax}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-count"] = types.YLeaf{"PosJitterCount", icmpPathJitterStats.PosJitterCount}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum"] = types.YLeaf{"NegJitterSum", icmpPathJitterStats.NegJitterSum}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-min"] = types.YLeaf{"NegJitterMin", icmpPathJitterStats.NegJitterMin}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-max"] = types.YLeaf{"NegJitterMax", icmpPathJitterStats.NegJitterMax}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum2"] = types.YLeaf{"NegJitterSum2", icmpPathJitterStats.NegJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-count"] = types.YLeaf{"NegJitterCount", icmpPathJitterStats.NegJitterCount}
    return &(icmpPathJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats_UdpJitterStats
// udp jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats_UdpJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterIn interface{}

    // Output Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterOut interface{}

    // Packets lost in source to destination (SD) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossSd interface{}

    // Packets lost in destination to source (DS) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossDs interface{}

    // Packets out of sequence. The type is interface{} with range: 0..4294967295.
    PacketOutOfSequence interface{}

    // Packets missing in action (cannot determine if theywere lost in SD or DS
    // direction. The type is interface{} with range: 0..4294967295.
    PacketMia interface{}

    // Packets which are skipped. The type is interface{} with range:
    // 0..4294967295.
    PacketSkipped interface{}

    // Packets arriving late. The type is interface{} with range: 0..4294967295.
    PacketLateArrivals interface{}

    // Packets with bad timestamps. The type is interface{} with range:
    // 0..4294967295.
    PacketInvalidTstamp interface{}

    // Number of internal errors. The type is interface{} with range:
    // 0..4294967295.
    InternalErrorsCount interface{}

    // Number of busies. The type is interface{} with range: 0..4294967295.
    BusiesCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive  packets) in SD direction Measured  in milliseconds. The type
    // is interface{} with range: 0..4294967295. Units are millisecond.
    PositiveSdSum interface{}

    // Sum of squares of positive jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveSdSum2 interface{}

    // Minimum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMin interface{}

    // Maximum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMax interface{}

    // Number of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in SD direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeSdSum interface{}

    // Sum of squares of negative jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeSdSum2 interface{}

    // Minimum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMin interface{}

    // Maximum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMax interface{}

    // Number of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    PositiveDsSum interface{}

    // Sum of squares of positive jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveDsSum2 interface{}

    // Minimum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMin interface{}

    // Maximum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMax interface{}

    // Number of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeDsSum interface{}

    // Sum of squares of negative jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeDsSum2 interface{}

    // Minimum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMin interface{}

    // Maximum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMax interface{}

    // Number of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsCount interface{}

    // Number of probe/probe-response pairs used to compute one-way statistics.
    // The type is interface{} with range: 0..4294967295.
    OneWayCount interface{}

    // Minimum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMin interface{}

    // Maximum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMax interface{}

    // Sum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdSum interface{}

    // Sum of squares of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..18446744073709551615.
    OneWaySdSum2 interface{}

    // Minimum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMin interface{}

    // Maximum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMax interface{}

    // Sum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsSum interface{}

    // Sum of squares of the OneWayMinDS and OneWayMaxDS values (msec). The type
    // is interface{} with range: 0..18446744073709551615.
    OneWayDsSum2 interface{}
}

func (udpJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Target_SpecificStats_UdpJitterStats) GetEntityData() *types.CommonEntityData {
    udpJitterStats.EntityData.YFilter = udpJitterStats.YFilter
    udpJitterStats.EntityData.YangName = "udp-jitter-stats"
    udpJitterStats.EntityData.BundleName = "cisco_ios_xr"
    udpJitterStats.EntityData.ParentYangName = "specific-stats"
    udpJitterStats.EntityData.SegmentPath = "udp-jitter-stats"
    udpJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpJitterStats.EntityData.Children = make(map[string]types.YChild)
    udpJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    udpJitterStats.EntityData.Leafs["jitter-in"] = types.YLeaf{"JitterIn", udpJitterStats.JitterIn}
    udpJitterStats.EntityData.Leafs["jitter-out"] = types.YLeaf{"JitterOut", udpJitterStats.JitterOut}
    udpJitterStats.EntityData.Leafs["packet-loss-sd"] = types.YLeaf{"PacketLossSd", udpJitterStats.PacketLossSd}
    udpJitterStats.EntityData.Leafs["packet-loss-ds"] = types.YLeaf{"PacketLossDs", udpJitterStats.PacketLossDs}
    udpJitterStats.EntityData.Leafs["packet-out-of-sequence"] = types.YLeaf{"PacketOutOfSequence", udpJitterStats.PacketOutOfSequence}
    udpJitterStats.EntityData.Leafs["packet-mia"] = types.YLeaf{"PacketMia", udpJitterStats.PacketMia}
    udpJitterStats.EntityData.Leafs["packet-skipped"] = types.YLeaf{"PacketSkipped", udpJitterStats.PacketSkipped}
    udpJitterStats.EntityData.Leafs["packet-late-arrivals"] = types.YLeaf{"PacketLateArrivals", udpJitterStats.PacketLateArrivals}
    udpJitterStats.EntityData.Leafs["packet-invalid-tstamp"] = types.YLeaf{"PacketInvalidTstamp", udpJitterStats.PacketInvalidTstamp}
    udpJitterStats.EntityData.Leafs["internal-errors-count"] = types.YLeaf{"InternalErrorsCount", udpJitterStats.InternalErrorsCount}
    udpJitterStats.EntityData.Leafs["busies-count"] = types.YLeaf{"BusiesCount", udpJitterStats.BusiesCount}
    udpJitterStats.EntityData.Leafs["positive-sd-sum"] = types.YLeaf{"PositiveSdSum", udpJitterStats.PositiveSdSum}
    udpJitterStats.EntityData.Leafs["positive-sd-sum2"] = types.YLeaf{"PositiveSdSum2", udpJitterStats.PositiveSdSum2}
    udpJitterStats.EntityData.Leafs["positive-sd-min"] = types.YLeaf{"PositiveSdMin", udpJitterStats.PositiveSdMin}
    udpJitterStats.EntityData.Leafs["positive-sd-max"] = types.YLeaf{"PositiveSdMax", udpJitterStats.PositiveSdMax}
    udpJitterStats.EntityData.Leafs["positive-sd-count"] = types.YLeaf{"PositiveSdCount", udpJitterStats.PositiveSdCount}
    udpJitterStats.EntityData.Leafs["negative-sd-sum"] = types.YLeaf{"NegativeSdSum", udpJitterStats.NegativeSdSum}
    udpJitterStats.EntityData.Leafs["negative-sd-sum2"] = types.YLeaf{"NegativeSdSum2", udpJitterStats.NegativeSdSum2}
    udpJitterStats.EntityData.Leafs["negative-sd-min"] = types.YLeaf{"NegativeSdMin", udpJitterStats.NegativeSdMin}
    udpJitterStats.EntityData.Leafs["negative-sd-max"] = types.YLeaf{"NegativeSdMax", udpJitterStats.NegativeSdMax}
    udpJitterStats.EntityData.Leafs["negative-sd-count"] = types.YLeaf{"NegativeSdCount", udpJitterStats.NegativeSdCount}
    udpJitterStats.EntityData.Leafs["positive-ds-sum"] = types.YLeaf{"PositiveDsSum", udpJitterStats.PositiveDsSum}
    udpJitterStats.EntityData.Leafs["positive-ds-sum2"] = types.YLeaf{"PositiveDsSum2", udpJitterStats.PositiveDsSum2}
    udpJitterStats.EntityData.Leafs["positive-ds-min"] = types.YLeaf{"PositiveDsMin", udpJitterStats.PositiveDsMin}
    udpJitterStats.EntityData.Leafs["positive-ds-max"] = types.YLeaf{"PositiveDsMax", udpJitterStats.PositiveDsMax}
    udpJitterStats.EntityData.Leafs["positive-ds-count"] = types.YLeaf{"PositiveDsCount", udpJitterStats.PositiveDsCount}
    udpJitterStats.EntityData.Leafs["negative-ds-sum"] = types.YLeaf{"NegativeDsSum", udpJitterStats.NegativeDsSum}
    udpJitterStats.EntityData.Leafs["negative-ds-sum2"] = types.YLeaf{"NegativeDsSum2", udpJitterStats.NegativeDsSum2}
    udpJitterStats.EntityData.Leafs["negative-ds-min"] = types.YLeaf{"NegativeDsMin", udpJitterStats.NegativeDsMin}
    udpJitterStats.EntityData.Leafs["negative-ds-max"] = types.YLeaf{"NegativeDsMax", udpJitterStats.NegativeDsMax}
    udpJitterStats.EntityData.Leafs["negative-ds-count"] = types.YLeaf{"NegativeDsCount", udpJitterStats.NegativeDsCount}
    udpJitterStats.EntityData.Leafs["one-way-count"] = types.YLeaf{"OneWayCount", udpJitterStats.OneWayCount}
    udpJitterStats.EntityData.Leafs["one-way-sd-min"] = types.YLeaf{"OneWaySdMin", udpJitterStats.OneWaySdMin}
    udpJitterStats.EntityData.Leafs["one-way-sd-max"] = types.YLeaf{"OneWaySdMax", udpJitterStats.OneWaySdMax}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum"] = types.YLeaf{"OneWaySdSum", udpJitterStats.OneWaySdSum}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum2"] = types.YLeaf{"OneWaySdSum2", udpJitterStats.OneWaySdSum2}
    udpJitterStats.EntityData.Leafs["one-way-ds-min"] = types.YLeaf{"OneWayDsMin", udpJitterStats.OneWayDsMin}
    udpJitterStats.EntityData.Leafs["one-way-ds-max"] = types.YLeaf{"OneWayDsMax", udpJitterStats.OneWayDsMax}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum"] = types.YLeaf{"OneWayDsSum", udpJitterStats.OneWayDsSum}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum2"] = types.YLeaf{"OneWayDsSum2", udpJitterStats.OneWayDsSum2}
    return &(udpJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops
// Latest statistics for hops in a
// path-enabled operation
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Latest stats for a hop in a path-enabled operation. The type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop.
    Hop []Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop
}

func (hops *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "latest"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = make(map[string]types.YChild)
    hops.EntityData.Children["hop"] = types.YChild{"Hop", nil}
    for i := range hops.Hop {
        hops.EntityData.Children[types.GetSegmentPath(&hops.Hop[i])] = types.YChild{"Hop", &hops.Hop[i]}
    }
    hops.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hops.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop
// Latest stats for a hop in a path-enabled
// operation
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Hop Index. The type is interface{} with range:
    // -2147483648..2147483647.
    HopIndex interface{}

    // Common Stats.
    CommonStats Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_CommonStats

    // Operation Specific Stats.
    SpecificStats Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats
}

func (hop *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop) GetEntityData() *types.CommonEntityData {
    hop.EntityData.YFilter = hop.YFilter
    hop.EntityData.YangName = "hop"
    hop.EntityData.BundleName = "cisco_ios_xr"
    hop.EntityData.ParentYangName = "hops"
    hop.EntityData.SegmentPath = "hop" + "[hop-index='" + fmt.Sprintf("%v", hop.HopIndex) + "']"
    hop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hop.EntityData.Children = make(map[string]types.YChild)
    hop.EntityData.Children["common-stats"] = types.YChild{"CommonStats", &hop.CommonStats}
    hop.EntityData.Children["specific-stats"] = types.YChild{"SpecificStats", &hop.SpecificStats}
    hop.EntityData.Leafs = make(map[string]types.YLeaf)
    hop.EntityData.Leafs["hop-index"] = types.YLeaf{"HopIndex", hop.HopIndex}
    return &(hop.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_CommonStats
// Common Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_CommonStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation Time. The type is interface{} with range:
    // 0..18446744073709551615.
    OperationTime interface{}

    // Return code. The type is IpslaRetCode.
    ReturnCode interface{}

    // Number of RTT samples used for the statistics. The type is interface{} with
    // range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of updates processed. The type is interface{} with range:
    // 0..4294967295.
    UpdateCount interface{}

    // Number of updates with Okay return code. The type is interface{} with
    // range: 0..4294967295.
    OkCount interface{}

    // Number of updates with Disconnected return code. The type is interface{}
    // with range: 0..4294967295.
    DisconnectCount interface{}

    // Number of updates with Timeout return code. The type is interface{} with
    // range: 0..4294967295.
    TimeoutCount interface{}

    // Number of updates with Busy return code. The type is interface{} with
    // range: 0..4294967295.
    BusyCount interface{}

    // Number of updates with NotConnected return code. The type is interface{}
    // with range: 0..4294967295.
    NoConnectionCount interface{}

    // Number of updates with Dropped return code. The type is interface{} with
    // range: 0..4294967295.
    DroppedCount interface{}

    // Number of updates with InternalError return code. The type is interface{}
    // with range: 0..4294967295.
    InternalErrorCount interface{}

    // Number of updates with SeqError return code. The type is interface{} with
    // range: 0..4294967295.
    SequenceErrorCount interface{}

    // Number of updates with VerifyError return code. The type is interface{}
    // with range: 0..4294967295.
    VerifyErrorCount interface{}
}

func (commonStats *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_CommonStats) GetEntityData() *types.CommonEntityData {
    commonStats.EntityData.YFilter = commonStats.YFilter
    commonStats.EntityData.YangName = "common-stats"
    commonStats.EntityData.BundleName = "cisco_ios_xr"
    commonStats.EntityData.ParentYangName = "hop"
    commonStats.EntityData.SegmentPath = "common-stats"
    commonStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonStats.EntityData.Children = make(map[string]types.YChild)
    commonStats.EntityData.Leafs = make(map[string]types.YLeaf)
    commonStats.EntityData.Leafs["operation-time"] = types.YLeaf{"OperationTime", commonStats.OperationTime}
    commonStats.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", commonStats.ReturnCode}
    commonStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", commonStats.ResponseTimeCount}
    commonStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", commonStats.ResponseTime}
    commonStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", commonStats.MinResponseTime}
    commonStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", commonStats.MaxResponseTime}
    commonStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", commonStats.SumResponseTime}
    commonStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", commonStats.Sum2ResponseTime}
    commonStats.EntityData.Leafs["update-count"] = types.YLeaf{"UpdateCount", commonStats.UpdateCount}
    commonStats.EntityData.Leafs["ok-count"] = types.YLeaf{"OkCount", commonStats.OkCount}
    commonStats.EntityData.Leafs["disconnect-count"] = types.YLeaf{"DisconnectCount", commonStats.DisconnectCount}
    commonStats.EntityData.Leafs["timeout-count"] = types.YLeaf{"TimeoutCount", commonStats.TimeoutCount}
    commonStats.EntityData.Leafs["busy-count"] = types.YLeaf{"BusyCount", commonStats.BusyCount}
    commonStats.EntityData.Leafs["no-connection-count"] = types.YLeaf{"NoConnectionCount", commonStats.NoConnectionCount}
    commonStats.EntityData.Leafs["dropped-count"] = types.YLeaf{"DroppedCount", commonStats.DroppedCount}
    commonStats.EntityData.Leafs["internal-error-count"] = types.YLeaf{"InternalErrorCount", commonStats.InternalErrorCount}
    commonStats.EntityData.Leafs["sequence-error-count"] = types.YLeaf{"SequenceErrorCount", commonStats.SequenceErrorCount}
    commonStats.EntityData.Leafs["verify-error-count"] = types.YLeaf{"VerifyErrorCount", commonStats.VerifyErrorCount}
    return &(commonStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats
// Operation Specific Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // op type. The type is OpTypeEnum.
    OpType interface{}

    // icmp path jitter stats.
    IcmpPathJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats_IcmpPathJitterStats

    // udp jitter stats.
    UdpJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats_UdpJitterStats
}

func (specificStats *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats) GetEntityData() *types.CommonEntityData {
    specificStats.EntityData.YFilter = specificStats.YFilter
    specificStats.EntityData.YangName = "specific-stats"
    specificStats.EntityData.BundleName = "cisco_ios_xr"
    specificStats.EntityData.ParentYangName = "hop"
    specificStats.EntityData.SegmentPath = "specific-stats"
    specificStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificStats.EntityData.Children = make(map[string]types.YChild)
    specificStats.EntityData.Children["icmp-path-jitter-stats"] = types.YChild{"IcmpPathJitterStats", &specificStats.IcmpPathJitterStats}
    specificStats.EntityData.Children["udp-jitter-stats"] = types.YChild{"UdpJitterStats", &specificStats.UdpJitterStats}
    specificStats.EntityData.Leafs = make(map[string]types.YLeaf)
    specificStats.EntityData.Leafs["op-type"] = types.YLeaf{"OpType", specificStats.OpType}
    return &(specificStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats_IcmpPathJitterStats
// icmp path jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats_IcmpPathJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address of the source. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // IP Address of the destination. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestAddress interface{}

    // IP address of the hop in the path. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}

    // Interval between echos in ms. The type is interface{} with range:
    // 0..4294967295.
    PacketInterval interface{}

    // Number of RTT samples  used for the statistics. The type is interface{}
    // with range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of Echo replies received . The type is interface{} with range:
    // 0..4294967295.
    PacketCount interface{}

    // Number of packets lost. The type is interface{} with range: 0..4294967295.
    PacketLossCount interface{}

    // Number of out of sequence packets. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceCount interface{}

    // Number of discarded samples. The type is interface{} with range:
    // 0..4294967295.
    DiscardedSampleCount interface{}

    // Number of packets with data corruption. The type is interface{} with range:
    // 0..4294967295.
    VerifyErrorsCount interface{}

    // Number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedErrorCount interface{}

    // Jitter value for this node in the path. The type is interface{} with range:
    // 0..4294967295.
    Jitter interface{}

    // Sum of positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterSum interface{}

    // Sum of squares of positive jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    PosJitterSum2 interface{}

    // Minimum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMin interface{}

    // Maximum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMax interface{}

    // Number of positive jitter values. The type is interface{} with range:
    // 0..4294967295.
    PosJitterCount interface{}

    // Sum of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterSum interface{}

    // Minimum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMin interface{}

    // Maximum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMax interface{}

    // Sum of squares of negative jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    NegJitterSum2 interface{}

    // Number of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterCount interface{}
}

func (icmpPathJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats_IcmpPathJitterStats) GetEntityData() *types.CommonEntityData {
    icmpPathJitterStats.EntityData.YFilter = icmpPathJitterStats.YFilter
    icmpPathJitterStats.EntityData.YangName = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.BundleName = "cisco_ios_xr"
    icmpPathJitterStats.EntityData.ParentYangName = "specific-stats"
    icmpPathJitterStats.EntityData.SegmentPath = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpPathJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpPathJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpPathJitterStats.EntityData.Children = make(map[string]types.YChild)
    icmpPathJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpPathJitterStats.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", icmpPathJitterStats.SourceAddress}
    icmpPathJitterStats.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", icmpPathJitterStats.DestAddress}
    icmpPathJitterStats.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", icmpPathJitterStats.HopAddress}
    icmpPathJitterStats.EntityData.Leafs["packet-interval"] = types.YLeaf{"PacketInterval", icmpPathJitterStats.PacketInterval}
    icmpPathJitterStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", icmpPathJitterStats.ResponseTimeCount}
    icmpPathJitterStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", icmpPathJitterStats.ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", icmpPathJitterStats.MinResponseTime}
    icmpPathJitterStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", icmpPathJitterStats.MaxResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", icmpPathJitterStats.SumResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", icmpPathJitterStats.Sum2ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["packet-count"] = types.YLeaf{"PacketCount", icmpPathJitterStats.PacketCount}
    icmpPathJitterStats.EntityData.Leafs["packet-loss-count"] = types.YLeaf{"PacketLossCount", icmpPathJitterStats.PacketLossCount}
    icmpPathJitterStats.EntityData.Leafs["out-of-sequence-count"] = types.YLeaf{"OutOfSequenceCount", icmpPathJitterStats.OutOfSequenceCount}
    icmpPathJitterStats.EntityData.Leafs["discarded-sample-count"] = types.YLeaf{"DiscardedSampleCount", icmpPathJitterStats.DiscardedSampleCount}
    icmpPathJitterStats.EntityData.Leafs["verify-errors-count"] = types.YLeaf{"VerifyErrorsCount", icmpPathJitterStats.VerifyErrorsCount}
    icmpPathJitterStats.EntityData.Leafs["dropped-error-count"] = types.YLeaf{"DroppedErrorCount", icmpPathJitterStats.DroppedErrorCount}
    icmpPathJitterStats.EntityData.Leafs["jitter"] = types.YLeaf{"Jitter", icmpPathJitterStats.Jitter}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum"] = types.YLeaf{"PosJitterSum", icmpPathJitterStats.PosJitterSum}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum2"] = types.YLeaf{"PosJitterSum2", icmpPathJitterStats.PosJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-min"] = types.YLeaf{"PosJitterMin", icmpPathJitterStats.PosJitterMin}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-max"] = types.YLeaf{"PosJitterMax", icmpPathJitterStats.PosJitterMax}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-count"] = types.YLeaf{"PosJitterCount", icmpPathJitterStats.PosJitterCount}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum"] = types.YLeaf{"NegJitterSum", icmpPathJitterStats.NegJitterSum}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-min"] = types.YLeaf{"NegJitterMin", icmpPathJitterStats.NegJitterMin}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-max"] = types.YLeaf{"NegJitterMax", icmpPathJitterStats.NegJitterMax}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum2"] = types.YLeaf{"NegJitterSum2", icmpPathJitterStats.NegJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-count"] = types.YLeaf{"NegJitterCount", icmpPathJitterStats.NegJitterCount}
    return &(icmpPathJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats_UdpJitterStats
// udp jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats_UdpJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterIn interface{}

    // Output Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterOut interface{}

    // Packets lost in source to destination (SD) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossSd interface{}

    // Packets lost in destination to source (DS) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossDs interface{}

    // Packets out of sequence. The type is interface{} with range: 0..4294967295.
    PacketOutOfSequence interface{}

    // Packets missing in action (cannot determine if theywere lost in SD or DS
    // direction. The type is interface{} with range: 0..4294967295.
    PacketMia interface{}

    // Packets which are skipped. The type is interface{} with range:
    // 0..4294967295.
    PacketSkipped interface{}

    // Packets arriving late. The type is interface{} with range: 0..4294967295.
    PacketLateArrivals interface{}

    // Packets with bad timestamps. The type is interface{} with range:
    // 0..4294967295.
    PacketInvalidTstamp interface{}

    // Number of internal errors. The type is interface{} with range:
    // 0..4294967295.
    InternalErrorsCount interface{}

    // Number of busies. The type is interface{} with range: 0..4294967295.
    BusiesCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive  packets) in SD direction Measured  in milliseconds. The type
    // is interface{} with range: 0..4294967295. Units are millisecond.
    PositiveSdSum interface{}

    // Sum of squares of positive jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveSdSum2 interface{}

    // Minimum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMin interface{}

    // Maximum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMax interface{}

    // Number of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in SD direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeSdSum interface{}

    // Sum of squares of negative jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeSdSum2 interface{}

    // Minimum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMin interface{}

    // Maximum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMax interface{}

    // Number of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    PositiveDsSum interface{}

    // Sum of squares of positive jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveDsSum2 interface{}

    // Minimum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMin interface{}

    // Maximum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMax interface{}

    // Number of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeDsSum interface{}

    // Sum of squares of negative jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeDsSum2 interface{}

    // Minimum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMin interface{}

    // Maximum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMax interface{}

    // Number of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsCount interface{}

    // Number of probe/probe-response pairs used to compute one-way statistics.
    // The type is interface{} with range: 0..4294967295.
    OneWayCount interface{}

    // Minimum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMin interface{}

    // Maximum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMax interface{}

    // Sum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdSum interface{}

    // Sum of squares of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..18446744073709551615.
    OneWaySdSum2 interface{}

    // Minimum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMin interface{}

    // Maximum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMax interface{}

    // Sum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsSum interface{}

    // Sum of squares of the OneWayMinDS and OneWayMaxDS values (msec). The type
    // is interface{} with range: 0..18446744073709551615.
    OneWayDsSum2 interface{}
}

func (udpJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Latest_Hops_Hop_SpecificStats_UdpJitterStats) GetEntityData() *types.CommonEntityData {
    udpJitterStats.EntityData.YFilter = udpJitterStats.YFilter
    udpJitterStats.EntityData.YangName = "udp-jitter-stats"
    udpJitterStats.EntityData.BundleName = "cisco_ios_xr"
    udpJitterStats.EntityData.ParentYangName = "specific-stats"
    udpJitterStats.EntityData.SegmentPath = "udp-jitter-stats"
    udpJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpJitterStats.EntityData.Children = make(map[string]types.YChild)
    udpJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    udpJitterStats.EntityData.Leafs["jitter-in"] = types.YLeaf{"JitterIn", udpJitterStats.JitterIn}
    udpJitterStats.EntityData.Leafs["jitter-out"] = types.YLeaf{"JitterOut", udpJitterStats.JitterOut}
    udpJitterStats.EntityData.Leafs["packet-loss-sd"] = types.YLeaf{"PacketLossSd", udpJitterStats.PacketLossSd}
    udpJitterStats.EntityData.Leafs["packet-loss-ds"] = types.YLeaf{"PacketLossDs", udpJitterStats.PacketLossDs}
    udpJitterStats.EntityData.Leafs["packet-out-of-sequence"] = types.YLeaf{"PacketOutOfSequence", udpJitterStats.PacketOutOfSequence}
    udpJitterStats.EntityData.Leafs["packet-mia"] = types.YLeaf{"PacketMia", udpJitterStats.PacketMia}
    udpJitterStats.EntityData.Leafs["packet-skipped"] = types.YLeaf{"PacketSkipped", udpJitterStats.PacketSkipped}
    udpJitterStats.EntityData.Leafs["packet-late-arrivals"] = types.YLeaf{"PacketLateArrivals", udpJitterStats.PacketLateArrivals}
    udpJitterStats.EntityData.Leafs["packet-invalid-tstamp"] = types.YLeaf{"PacketInvalidTstamp", udpJitterStats.PacketInvalidTstamp}
    udpJitterStats.EntityData.Leafs["internal-errors-count"] = types.YLeaf{"InternalErrorsCount", udpJitterStats.InternalErrorsCount}
    udpJitterStats.EntityData.Leafs["busies-count"] = types.YLeaf{"BusiesCount", udpJitterStats.BusiesCount}
    udpJitterStats.EntityData.Leafs["positive-sd-sum"] = types.YLeaf{"PositiveSdSum", udpJitterStats.PositiveSdSum}
    udpJitterStats.EntityData.Leafs["positive-sd-sum2"] = types.YLeaf{"PositiveSdSum2", udpJitterStats.PositiveSdSum2}
    udpJitterStats.EntityData.Leafs["positive-sd-min"] = types.YLeaf{"PositiveSdMin", udpJitterStats.PositiveSdMin}
    udpJitterStats.EntityData.Leafs["positive-sd-max"] = types.YLeaf{"PositiveSdMax", udpJitterStats.PositiveSdMax}
    udpJitterStats.EntityData.Leafs["positive-sd-count"] = types.YLeaf{"PositiveSdCount", udpJitterStats.PositiveSdCount}
    udpJitterStats.EntityData.Leafs["negative-sd-sum"] = types.YLeaf{"NegativeSdSum", udpJitterStats.NegativeSdSum}
    udpJitterStats.EntityData.Leafs["negative-sd-sum2"] = types.YLeaf{"NegativeSdSum2", udpJitterStats.NegativeSdSum2}
    udpJitterStats.EntityData.Leafs["negative-sd-min"] = types.YLeaf{"NegativeSdMin", udpJitterStats.NegativeSdMin}
    udpJitterStats.EntityData.Leafs["negative-sd-max"] = types.YLeaf{"NegativeSdMax", udpJitterStats.NegativeSdMax}
    udpJitterStats.EntityData.Leafs["negative-sd-count"] = types.YLeaf{"NegativeSdCount", udpJitterStats.NegativeSdCount}
    udpJitterStats.EntityData.Leafs["positive-ds-sum"] = types.YLeaf{"PositiveDsSum", udpJitterStats.PositiveDsSum}
    udpJitterStats.EntityData.Leafs["positive-ds-sum2"] = types.YLeaf{"PositiveDsSum2", udpJitterStats.PositiveDsSum2}
    udpJitterStats.EntityData.Leafs["positive-ds-min"] = types.YLeaf{"PositiveDsMin", udpJitterStats.PositiveDsMin}
    udpJitterStats.EntityData.Leafs["positive-ds-max"] = types.YLeaf{"PositiveDsMax", udpJitterStats.PositiveDsMax}
    udpJitterStats.EntityData.Leafs["positive-ds-count"] = types.YLeaf{"PositiveDsCount", udpJitterStats.PositiveDsCount}
    udpJitterStats.EntityData.Leafs["negative-ds-sum"] = types.YLeaf{"NegativeDsSum", udpJitterStats.NegativeDsSum}
    udpJitterStats.EntityData.Leafs["negative-ds-sum2"] = types.YLeaf{"NegativeDsSum2", udpJitterStats.NegativeDsSum2}
    udpJitterStats.EntityData.Leafs["negative-ds-min"] = types.YLeaf{"NegativeDsMin", udpJitterStats.NegativeDsMin}
    udpJitterStats.EntityData.Leafs["negative-ds-max"] = types.YLeaf{"NegativeDsMax", udpJitterStats.NegativeDsMax}
    udpJitterStats.EntityData.Leafs["negative-ds-count"] = types.YLeaf{"NegativeDsCount", udpJitterStats.NegativeDsCount}
    udpJitterStats.EntityData.Leafs["one-way-count"] = types.YLeaf{"OneWayCount", udpJitterStats.OneWayCount}
    udpJitterStats.EntityData.Leafs["one-way-sd-min"] = types.YLeaf{"OneWaySdMin", udpJitterStats.OneWaySdMin}
    udpJitterStats.EntityData.Leafs["one-way-sd-max"] = types.YLeaf{"OneWaySdMax", udpJitterStats.OneWaySdMax}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum"] = types.YLeaf{"OneWaySdSum", udpJitterStats.OneWaySdSum}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum2"] = types.YLeaf{"OneWaySdSum2", udpJitterStats.OneWaySdSum2}
    udpJitterStats.EntityData.Leafs["one-way-ds-min"] = types.YLeaf{"OneWayDsMin", udpJitterStats.OneWayDsMin}
    udpJitterStats.EntityData.Leafs["one-way-ds-max"] = types.YLeaf{"OneWayDsMax", udpJitterStats.OneWayDsMax}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum"] = types.YLeaf{"OneWayDsSum", udpJitterStats.OneWayDsSum}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum2"] = types.YLeaf{"OneWayDsSum2", udpJitterStats.OneWayDsSum2}
    return &(udpJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths
// List of latest LPD paths
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Latest path statistics of MPLS LSP group operation. The type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths_LpdPath.
    LpdPath []Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths_LpdPath
}

func (lpdPaths *Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths) GetEntityData() *types.CommonEntityData {
    lpdPaths.EntityData.YFilter = lpdPaths.YFilter
    lpdPaths.EntityData.YangName = "lpd-paths"
    lpdPaths.EntityData.BundleName = "cisco_ios_xr"
    lpdPaths.EntityData.ParentYangName = "latest"
    lpdPaths.EntityData.SegmentPath = "lpd-paths"
    lpdPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpdPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpdPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpdPaths.EntityData.Children = make(map[string]types.YChild)
    lpdPaths.EntityData.Children["lpd-path"] = types.YChild{"LpdPath", nil}
    for i := range lpdPaths.LpdPath {
        lpdPaths.EntityData.Children[types.GetSegmentPath(&lpdPaths.LpdPath[i])] = types.YChild{"LpdPath", &lpdPaths.LpdPath[i]}
    }
    lpdPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lpdPaths.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths_LpdPath
// Latest path statistics of MPLS LSP group
// operation
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths_LpdPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPD path index. The type is interface{} with
    // range: -2147483648..2147483647.
    PathIndex interface{}

    // Path return code. The type is IpslaRetCode.
    ReturnCode interface{}

    // LPD path identifier.
    PathId Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths_LpdPath_PathId
}

func (lpdPath *Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths_LpdPath) GetEntityData() *types.CommonEntityData {
    lpdPath.EntityData.YFilter = lpdPath.YFilter
    lpdPath.EntityData.YangName = "lpd-path"
    lpdPath.EntityData.BundleName = "cisco_ios_xr"
    lpdPath.EntityData.ParentYangName = "lpd-paths"
    lpdPath.EntityData.SegmentPath = "lpd-path" + "[path-index='" + fmt.Sprintf("%v", lpdPath.PathIndex) + "']"
    lpdPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpdPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpdPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpdPath.EntityData.Children = make(map[string]types.YChild)
    lpdPath.EntityData.Children["path-id"] = types.YChild{"PathId", &lpdPath.PathId}
    lpdPath.EntityData.Leafs = make(map[string]types.YLeaf)
    lpdPath.EntityData.Leafs["path-index"] = types.YLeaf{"PathIndex", lpdPath.PathIndex}
    lpdPath.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", lpdPath.ReturnCode}
    return &(lpdPath.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths_LpdPath_PathId
// LPD path identifier
type Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths_LpdPath_PathId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP selector. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LspSelector interface{}

    // Output interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    OutputInterface interface{}

    // Nexthop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}

    // Downstream label stacks. The type is slice of interface{} with range:
    // 0..4294967295.
    DownstreamLabel []interface{}
}

func (pathId *Ipsla_OperationData_Operations_Operation_Statistics_Latest_LpdPaths_LpdPath_PathId) GetEntityData() *types.CommonEntityData {
    pathId.EntityData.YFilter = pathId.YFilter
    pathId.EntityData.YangName = "path-id"
    pathId.EntityData.BundleName = "cisco_ios_xr"
    pathId.EntityData.ParentYangName = "lpd-path"
    pathId.EntityData.SegmentPath = "path-id"
    pathId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathId.EntityData.Children = make(map[string]types.YChild)
    pathId.EntityData.Leafs = make(map[string]types.YLeaf)
    pathId.EntityData.Leafs["lsp-selector"] = types.YLeaf{"LspSelector", pathId.LspSelector}
    pathId.EntityData.Leafs["output-interface"] = types.YLeaf{"OutputInterface", pathId.OutputInterface}
    pathId.EntityData.Leafs["nexthop-address"] = types.YLeaf{"NexthopAddress", pathId.NexthopAddress}
    pathId.EntityData.Leafs["downstream-label"] = types.YLeaf{"DownstreamLabel", pathId.DownstreamLabel}
    return &(pathId.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated
// Statistics aggregated for data collected
// over time intervals
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of statistics aggregated over enhanced intervals.
    EnhancedIntervals Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals

    // Table of statistics aggregated over 1-hour intervals.
    Hours Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours
}

func (aggregated *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated) GetEntityData() *types.CommonEntityData {
    aggregated.EntityData.YFilter = aggregated.YFilter
    aggregated.EntityData.YangName = "aggregated"
    aggregated.EntityData.BundleName = "cisco_ios_xr"
    aggregated.EntityData.ParentYangName = "statistics"
    aggregated.EntityData.SegmentPath = "aggregated"
    aggregated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregated.EntityData.Children = make(map[string]types.YChild)
    aggregated.EntityData.Children["enhanced-intervals"] = types.YChild{"EnhancedIntervals", &aggregated.EnhancedIntervals}
    aggregated.EntityData.Children["hours"] = types.YChild{"Hours", &aggregated.Hours}
    aggregated.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aggregated.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals
// Table of statistics aggregated over
// enhanced intervals
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics aggregated over an interval specified in seconds. Specified
    // interval must be a multiple of the operation frequency. The type is slice
    // of
    // Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval.
    EnhancedInterval []Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval
}

func (enhancedIntervals *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals) GetEntityData() *types.CommonEntityData {
    enhancedIntervals.EntityData.YFilter = enhancedIntervals.YFilter
    enhancedIntervals.EntityData.YangName = "enhanced-intervals"
    enhancedIntervals.EntityData.BundleName = "cisco_ios_xr"
    enhancedIntervals.EntityData.ParentYangName = "aggregated"
    enhancedIntervals.EntityData.SegmentPath = "enhanced-intervals"
    enhancedIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enhancedIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enhancedIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enhancedIntervals.EntityData.Children = make(map[string]types.YChild)
    enhancedIntervals.EntityData.Children["enhanced-interval"] = types.YChild{"EnhancedInterval", nil}
    for i := range enhancedIntervals.EnhancedInterval {
        enhancedIntervals.EntityData.Children[types.GetSegmentPath(&enhancedIntervals.EnhancedInterval[i])] = types.YChild{"EnhancedInterval", &enhancedIntervals.EnhancedInterval[i]}
    }
    enhancedIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enhancedIntervals.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval
// Statistics aggregated over an interval
// specified in seconds. Specified interval
// must be a multiple of the operation
// frequency
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Enhanced Interval in seconds. The type is
    // interface{} with range: -2147483648..2147483647. Units are second.
    EnhancedInterval interface{}

    // Table of start times for the intervals.
    StartTimes Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes
}

func (enhancedInterval *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval) GetEntityData() *types.CommonEntityData {
    enhancedInterval.EntityData.YFilter = enhancedInterval.YFilter
    enhancedInterval.EntityData.YangName = "enhanced-interval"
    enhancedInterval.EntityData.BundleName = "cisco_ios_xr"
    enhancedInterval.EntityData.ParentYangName = "enhanced-intervals"
    enhancedInterval.EntityData.SegmentPath = "enhanced-interval" + "[enhanced-interval='" + fmt.Sprintf("%v", enhancedInterval.EnhancedInterval) + "']"
    enhancedInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enhancedInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enhancedInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enhancedInterval.EntityData.Children = make(map[string]types.YChild)
    enhancedInterval.EntityData.Children["start-times"] = types.YChild{"StartTimes", &enhancedInterval.StartTimes}
    enhancedInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    enhancedInterval.EntityData.Leafs["enhanced-interval"] = types.YLeaf{"EnhancedInterval", enhancedInterval.EnhancedInterval}
    return &(enhancedInterval.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes
// Table of start times for the intervals
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics aggregated over an enhanced interval which starts at a specific
    // time. The type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime.
    StartTime []Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime
}

func (startTimes *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes) GetEntityData() *types.CommonEntityData {
    startTimes.EntityData.YFilter = startTimes.YFilter
    startTimes.EntityData.YangName = "start-times"
    startTimes.EntityData.BundleName = "cisco_ios_xr"
    startTimes.EntityData.ParentYangName = "enhanced-interval"
    startTimes.EntityData.SegmentPath = "start-times"
    startTimes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    startTimes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    startTimes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    startTimes.EntityData.Children = make(map[string]types.YChild)
    startTimes.EntityData.Children["start-time"] = types.YChild{"StartTime", nil}
    for i := range startTimes.StartTime {
        startTimes.EntityData.Children[types.GetSegmentPath(&startTimes.StartTime[i])] = types.YChild{"StartTime", &startTimes.StartTime[i]}
    }
    startTimes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(startTimes.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime
// Statistics aggregated over an enhanced
// interval which starts at a specific time
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interval Start Time. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    IntervalStartTime interface{}

    // Common Stats.
    CommonStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_CommonStats

    // Operation Specific Stats.
    SpecificStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats
}

func (startTime *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime) GetEntityData() *types.CommonEntityData {
    startTime.EntityData.YFilter = startTime.YFilter
    startTime.EntityData.YangName = "start-time"
    startTime.EntityData.BundleName = "cisco_ios_xr"
    startTime.EntityData.ParentYangName = "start-times"
    startTime.EntityData.SegmentPath = "start-time" + "[interval-start-time='" + fmt.Sprintf("%v", startTime.IntervalStartTime) + "']"
    startTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    startTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    startTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    startTime.EntityData.Children = make(map[string]types.YChild)
    startTime.EntityData.Children["common-stats"] = types.YChild{"CommonStats", &startTime.CommonStats}
    startTime.EntityData.Children["specific-stats"] = types.YChild{"SpecificStats", &startTime.SpecificStats}
    startTime.EntityData.Leafs = make(map[string]types.YLeaf)
    startTime.EntityData.Leafs["interval-start-time"] = types.YLeaf{"IntervalStartTime", startTime.IntervalStartTime}
    return &(startTime.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_CommonStats
// Common Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_CommonStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation Time. The type is interface{} with range:
    // 0..18446744073709551615.
    OperationTime interface{}

    // Return code. The type is IpslaRetCode.
    ReturnCode interface{}

    // Number of RTT samples used for the statistics. The type is interface{} with
    // range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of updates processed. The type is interface{} with range:
    // 0..4294967295.
    UpdateCount interface{}

    // Number of updates with Okay return code. The type is interface{} with
    // range: 0..4294967295.
    OkCount interface{}

    // Number of updates with Disconnected return code. The type is interface{}
    // with range: 0..4294967295.
    DisconnectCount interface{}

    // Number of updates with Timeout return code. The type is interface{} with
    // range: 0..4294967295.
    TimeoutCount interface{}

    // Number of updates with Busy return code. The type is interface{} with
    // range: 0..4294967295.
    BusyCount interface{}

    // Number of updates with NotConnected return code. The type is interface{}
    // with range: 0..4294967295.
    NoConnectionCount interface{}

    // Number of updates with Dropped return code. The type is interface{} with
    // range: 0..4294967295.
    DroppedCount interface{}

    // Number of updates with InternalError return code. The type is interface{}
    // with range: 0..4294967295.
    InternalErrorCount interface{}

    // Number of updates with SeqError return code. The type is interface{} with
    // range: 0..4294967295.
    SequenceErrorCount interface{}

    // Number of updates with VerifyError return code. The type is interface{}
    // with range: 0..4294967295.
    VerifyErrorCount interface{}
}

func (commonStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_CommonStats) GetEntityData() *types.CommonEntityData {
    commonStats.EntityData.YFilter = commonStats.YFilter
    commonStats.EntityData.YangName = "common-stats"
    commonStats.EntityData.BundleName = "cisco_ios_xr"
    commonStats.EntityData.ParentYangName = "start-time"
    commonStats.EntityData.SegmentPath = "common-stats"
    commonStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonStats.EntityData.Children = make(map[string]types.YChild)
    commonStats.EntityData.Leafs = make(map[string]types.YLeaf)
    commonStats.EntityData.Leafs["operation-time"] = types.YLeaf{"OperationTime", commonStats.OperationTime}
    commonStats.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", commonStats.ReturnCode}
    commonStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", commonStats.ResponseTimeCount}
    commonStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", commonStats.ResponseTime}
    commonStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", commonStats.MinResponseTime}
    commonStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", commonStats.MaxResponseTime}
    commonStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", commonStats.SumResponseTime}
    commonStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", commonStats.Sum2ResponseTime}
    commonStats.EntityData.Leafs["update-count"] = types.YLeaf{"UpdateCount", commonStats.UpdateCount}
    commonStats.EntityData.Leafs["ok-count"] = types.YLeaf{"OkCount", commonStats.OkCount}
    commonStats.EntityData.Leafs["disconnect-count"] = types.YLeaf{"DisconnectCount", commonStats.DisconnectCount}
    commonStats.EntityData.Leafs["timeout-count"] = types.YLeaf{"TimeoutCount", commonStats.TimeoutCount}
    commonStats.EntityData.Leafs["busy-count"] = types.YLeaf{"BusyCount", commonStats.BusyCount}
    commonStats.EntityData.Leafs["no-connection-count"] = types.YLeaf{"NoConnectionCount", commonStats.NoConnectionCount}
    commonStats.EntityData.Leafs["dropped-count"] = types.YLeaf{"DroppedCount", commonStats.DroppedCount}
    commonStats.EntityData.Leafs["internal-error-count"] = types.YLeaf{"InternalErrorCount", commonStats.InternalErrorCount}
    commonStats.EntityData.Leafs["sequence-error-count"] = types.YLeaf{"SequenceErrorCount", commonStats.SequenceErrorCount}
    commonStats.EntityData.Leafs["verify-error-count"] = types.YLeaf{"VerifyErrorCount", commonStats.VerifyErrorCount}
    return &(commonStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats
// Operation Specific Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // op type. The type is OpTypeEnum.
    OpType interface{}

    // icmp path jitter stats.
    IcmpPathJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats_IcmpPathJitterStats

    // udp jitter stats.
    UdpJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats_UdpJitterStats
}

func (specificStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats) GetEntityData() *types.CommonEntityData {
    specificStats.EntityData.YFilter = specificStats.YFilter
    specificStats.EntityData.YangName = "specific-stats"
    specificStats.EntityData.BundleName = "cisco_ios_xr"
    specificStats.EntityData.ParentYangName = "start-time"
    specificStats.EntityData.SegmentPath = "specific-stats"
    specificStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificStats.EntityData.Children = make(map[string]types.YChild)
    specificStats.EntityData.Children["icmp-path-jitter-stats"] = types.YChild{"IcmpPathJitterStats", &specificStats.IcmpPathJitterStats}
    specificStats.EntityData.Children["udp-jitter-stats"] = types.YChild{"UdpJitterStats", &specificStats.UdpJitterStats}
    specificStats.EntityData.Leafs = make(map[string]types.YLeaf)
    specificStats.EntityData.Leafs["op-type"] = types.YLeaf{"OpType", specificStats.OpType}
    return &(specificStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats_IcmpPathJitterStats
// icmp path jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats_IcmpPathJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address of the source. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // IP Address of the destination. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestAddress interface{}

    // IP address of the hop in the path. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}

    // Interval between echos in ms. The type is interface{} with range:
    // 0..4294967295.
    PacketInterval interface{}

    // Number of RTT samples  used for the statistics. The type is interface{}
    // with range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of Echo replies received . The type is interface{} with range:
    // 0..4294967295.
    PacketCount interface{}

    // Number of packets lost. The type is interface{} with range: 0..4294967295.
    PacketLossCount interface{}

    // Number of out of sequence packets. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceCount interface{}

    // Number of discarded samples. The type is interface{} with range:
    // 0..4294967295.
    DiscardedSampleCount interface{}

    // Number of packets with data corruption. The type is interface{} with range:
    // 0..4294967295.
    VerifyErrorsCount interface{}

    // Number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedErrorCount interface{}

    // Jitter value for this node in the path. The type is interface{} with range:
    // 0..4294967295.
    Jitter interface{}

    // Sum of positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterSum interface{}

    // Sum of squares of positive jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    PosJitterSum2 interface{}

    // Minimum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMin interface{}

    // Maximum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMax interface{}

    // Number of positive jitter values. The type is interface{} with range:
    // 0..4294967295.
    PosJitterCount interface{}

    // Sum of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterSum interface{}

    // Minimum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMin interface{}

    // Maximum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMax interface{}

    // Sum of squares of negative jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    NegJitterSum2 interface{}

    // Number of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterCount interface{}
}

func (icmpPathJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats_IcmpPathJitterStats) GetEntityData() *types.CommonEntityData {
    icmpPathJitterStats.EntityData.YFilter = icmpPathJitterStats.YFilter
    icmpPathJitterStats.EntityData.YangName = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.BundleName = "cisco_ios_xr"
    icmpPathJitterStats.EntityData.ParentYangName = "specific-stats"
    icmpPathJitterStats.EntityData.SegmentPath = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpPathJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpPathJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpPathJitterStats.EntityData.Children = make(map[string]types.YChild)
    icmpPathJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpPathJitterStats.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", icmpPathJitterStats.SourceAddress}
    icmpPathJitterStats.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", icmpPathJitterStats.DestAddress}
    icmpPathJitterStats.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", icmpPathJitterStats.HopAddress}
    icmpPathJitterStats.EntityData.Leafs["packet-interval"] = types.YLeaf{"PacketInterval", icmpPathJitterStats.PacketInterval}
    icmpPathJitterStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", icmpPathJitterStats.ResponseTimeCount}
    icmpPathJitterStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", icmpPathJitterStats.ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", icmpPathJitterStats.MinResponseTime}
    icmpPathJitterStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", icmpPathJitterStats.MaxResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", icmpPathJitterStats.SumResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", icmpPathJitterStats.Sum2ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["packet-count"] = types.YLeaf{"PacketCount", icmpPathJitterStats.PacketCount}
    icmpPathJitterStats.EntityData.Leafs["packet-loss-count"] = types.YLeaf{"PacketLossCount", icmpPathJitterStats.PacketLossCount}
    icmpPathJitterStats.EntityData.Leafs["out-of-sequence-count"] = types.YLeaf{"OutOfSequenceCount", icmpPathJitterStats.OutOfSequenceCount}
    icmpPathJitterStats.EntityData.Leafs["discarded-sample-count"] = types.YLeaf{"DiscardedSampleCount", icmpPathJitterStats.DiscardedSampleCount}
    icmpPathJitterStats.EntityData.Leafs["verify-errors-count"] = types.YLeaf{"VerifyErrorsCount", icmpPathJitterStats.VerifyErrorsCount}
    icmpPathJitterStats.EntityData.Leafs["dropped-error-count"] = types.YLeaf{"DroppedErrorCount", icmpPathJitterStats.DroppedErrorCount}
    icmpPathJitterStats.EntityData.Leafs["jitter"] = types.YLeaf{"Jitter", icmpPathJitterStats.Jitter}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum"] = types.YLeaf{"PosJitterSum", icmpPathJitterStats.PosJitterSum}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum2"] = types.YLeaf{"PosJitterSum2", icmpPathJitterStats.PosJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-min"] = types.YLeaf{"PosJitterMin", icmpPathJitterStats.PosJitterMin}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-max"] = types.YLeaf{"PosJitterMax", icmpPathJitterStats.PosJitterMax}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-count"] = types.YLeaf{"PosJitterCount", icmpPathJitterStats.PosJitterCount}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum"] = types.YLeaf{"NegJitterSum", icmpPathJitterStats.NegJitterSum}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-min"] = types.YLeaf{"NegJitterMin", icmpPathJitterStats.NegJitterMin}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-max"] = types.YLeaf{"NegJitterMax", icmpPathJitterStats.NegJitterMax}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum2"] = types.YLeaf{"NegJitterSum2", icmpPathJitterStats.NegJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-count"] = types.YLeaf{"NegJitterCount", icmpPathJitterStats.NegJitterCount}
    return &(icmpPathJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats_UdpJitterStats
// udp jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats_UdpJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterIn interface{}

    // Output Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterOut interface{}

    // Packets lost in source to destination (SD) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossSd interface{}

    // Packets lost in destination to source (DS) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossDs interface{}

    // Packets out of sequence. The type is interface{} with range: 0..4294967295.
    PacketOutOfSequence interface{}

    // Packets missing in action (cannot determine if theywere lost in SD or DS
    // direction. The type is interface{} with range: 0..4294967295.
    PacketMia interface{}

    // Packets which are skipped. The type is interface{} with range:
    // 0..4294967295.
    PacketSkipped interface{}

    // Packets arriving late. The type is interface{} with range: 0..4294967295.
    PacketLateArrivals interface{}

    // Packets with bad timestamps. The type is interface{} with range:
    // 0..4294967295.
    PacketInvalidTstamp interface{}

    // Number of internal errors. The type is interface{} with range:
    // 0..4294967295.
    InternalErrorsCount interface{}

    // Number of busies. The type is interface{} with range: 0..4294967295.
    BusiesCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive  packets) in SD direction Measured  in milliseconds. The type
    // is interface{} with range: 0..4294967295. Units are millisecond.
    PositiveSdSum interface{}

    // Sum of squares of positive jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveSdSum2 interface{}

    // Minimum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMin interface{}

    // Maximum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMax interface{}

    // Number of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in SD direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeSdSum interface{}

    // Sum of squares of negative jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeSdSum2 interface{}

    // Minimum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMin interface{}

    // Maximum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMax interface{}

    // Number of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    PositiveDsSum interface{}

    // Sum of squares of positive jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveDsSum2 interface{}

    // Minimum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMin interface{}

    // Maximum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMax interface{}

    // Number of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeDsSum interface{}

    // Sum of squares of negative jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeDsSum2 interface{}

    // Minimum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMin interface{}

    // Maximum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMax interface{}

    // Number of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsCount interface{}

    // Number of probe/probe-response pairs used to compute one-way statistics.
    // The type is interface{} with range: 0..4294967295.
    OneWayCount interface{}

    // Minimum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMin interface{}

    // Maximum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMax interface{}

    // Sum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdSum interface{}

    // Sum of squares of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..18446744073709551615.
    OneWaySdSum2 interface{}

    // Minimum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMin interface{}

    // Maximum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMax interface{}

    // Sum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsSum interface{}

    // Sum of squares of the OneWayMinDS and OneWayMaxDS values (msec). The type
    // is interface{} with range: 0..18446744073709551615.
    OneWayDsSum2 interface{}
}

func (udpJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_EnhancedIntervals_EnhancedInterval_StartTimes_StartTime_SpecificStats_UdpJitterStats) GetEntityData() *types.CommonEntityData {
    udpJitterStats.EntityData.YFilter = udpJitterStats.YFilter
    udpJitterStats.EntityData.YangName = "udp-jitter-stats"
    udpJitterStats.EntityData.BundleName = "cisco_ios_xr"
    udpJitterStats.EntityData.ParentYangName = "specific-stats"
    udpJitterStats.EntityData.SegmentPath = "udp-jitter-stats"
    udpJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpJitterStats.EntityData.Children = make(map[string]types.YChild)
    udpJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    udpJitterStats.EntityData.Leafs["jitter-in"] = types.YLeaf{"JitterIn", udpJitterStats.JitterIn}
    udpJitterStats.EntityData.Leafs["jitter-out"] = types.YLeaf{"JitterOut", udpJitterStats.JitterOut}
    udpJitterStats.EntityData.Leafs["packet-loss-sd"] = types.YLeaf{"PacketLossSd", udpJitterStats.PacketLossSd}
    udpJitterStats.EntityData.Leafs["packet-loss-ds"] = types.YLeaf{"PacketLossDs", udpJitterStats.PacketLossDs}
    udpJitterStats.EntityData.Leafs["packet-out-of-sequence"] = types.YLeaf{"PacketOutOfSequence", udpJitterStats.PacketOutOfSequence}
    udpJitterStats.EntityData.Leafs["packet-mia"] = types.YLeaf{"PacketMia", udpJitterStats.PacketMia}
    udpJitterStats.EntityData.Leafs["packet-skipped"] = types.YLeaf{"PacketSkipped", udpJitterStats.PacketSkipped}
    udpJitterStats.EntityData.Leafs["packet-late-arrivals"] = types.YLeaf{"PacketLateArrivals", udpJitterStats.PacketLateArrivals}
    udpJitterStats.EntityData.Leafs["packet-invalid-tstamp"] = types.YLeaf{"PacketInvalidTstamp", udpJitterStats.PacketInvalidTstamp}
    udpJitterStats.EntityData.Leafs["internal-errors-count"] = types.YLeaf{"InternalErrorsCount", udpJitterStats.InternalErrorsCount}
    udpJitterStats.EntityData.Leafs["busies-count"] = types.YLeaf{"BusiesCount", udpJitterStats.BusiesCount}
    udpJitterStats.EntityData.Leafs["positive-sd-sum"] = types.YLeaf{"PositiveSdSum", udpJitterStats.PositiveSdSum}
    udpJitterStats.EntityData.Leafs["positive-sd-sum2"] = types.YLeaf{"PositiveSdSum2", udpJitterStats.PositiveSdSum2}
    udpJitterStats.EntityData.Leafs["positive-sd-min"] = types.YLeaf{"PositiveSdMin", udpJitterStats.PositiveSdMin}
    udpJitterStats.EntityData.Leafs["positive-sd-max"] = types.YLeaf{"PositiveSdMax", udpJitterStats.PositiveSdMax}
    udpJitterStats.EntityData.Leafs["positive-sd-count"] = types.YLeaf{"PositiveSdCount", udpJitterStats.PositiveSdCount}
    udpJitterStats.EntityData.Leafs["negative-sd-sum"] = types.YLeaf{"NegativeSdSum", udpJitterStats.NegativeSdSum}
    udpJitterStats.EntityData.Leafs["negative-sd-sum2"] = types.YLeaf{"NegativeSdSum2", udpJitterStats.NegativeSdSum2}
    udpJitterStats.EntityData.Leafs["negative-sd-min"] = types.YLeaf{"NegativeSdMin", udpJitterStats.NegativeSdMin}
    udpJitterStats.EntityData.Leafs["negative-sd-max"] = types.YLeaf{"NegativeSdMax", udpJitterStats.NegativeSdMax}
    udpJitterStats.EntityData.Leafs["negative-sd-count"] = types.YLeaf{"NegativeSdCount", udpJitterStats.NegativeSdCount}
    udpJitterStats.EntityData.Leafs["positive-ds-sum"] = types.YLeaf{"PositiveDsSum", udpJitterStats.PositiveDsSum}
    udpJitterStats.EntityData.Leafs["positive-ds-sum2"] = types.YLeaf{"PositiveDsSum2", udpJitterStats.PositiveDsSum2}
    udpJitterStats.EntityData.Leafs["positive-ds-min"] = types.YLeaf{"PositiveDsMin", udpJitterStats.PositiveDsMin}
    udpJitterStats.EntityData.Leafs["positive-ds-max"] = types.YLeaf{"PositiveDsMax", udpJitterStats.PositiveDsMax}
    udpJitterStats.EntityData.Leafs["positive-ds-count"] = types.YLeaf{"PositiveDsCount", udpJitterStats.PositiveDsCount}
    udpJitterStats.EntityData.Leafs["negative-ds-sum"] = types.YLeaf{"NegativeDsSum", udpJitterStats.NegativeDsSum}
    udpJitterStats.EntityData.Leafs["negative-ds-sum2"] = types.YLeaf{"NegativeDsSum2", udpJitterStats.NegativeDsSum2}
    udpJitterStats.EntityData.Leafs["negative-ds-min"] = types.YLeaf{"NegativeDsMin", udpJitterStats.NegativeDsMin}
    udpJitterStats.EntityData.Leafs["negative-ds-max"] = types.YLeaf{"NegativeDsMax", udpJitterStats.NegativeDsMax}
    udpJitterStats.EntityData.Leafs["negative-ds-count"] = types.YLeaf{"NegativeDsCount", udpJitterStats.NegativeDsCount}
    udpJitterStats.EntityData.Leafs["one-way-count"] = types.YLeaf{"OneWayCount", udpJitterStats.OneWayCount}
    udpJitterStats.EntityData.Leafs["one-way-sd-min"] = types.YLeaf{"OneWaySdMin", udpJitterStats.OneWaySdMin}
    udpJitterStats.EntityData.Leafs["one-way-sd-max"] = types.YLeaf{"OneWaySdMax", udpJitterStats.OneWaySdMax}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum"] = types.YLeaf{"OneWaySdSum", udpJitterStats.OneWaySdSum}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum2"] = types.YLeaf{"OneWaySdSum2", udpJitterStats.OneWaySdSum2}
    udpJitterStats.EntityData.Leafs["one-way-ds-min"] = types.YLeaf{"OneWayDsMin", udpJitterStats.OneWayDsMin}
    udpJitterStats.EntityData.Leafs["one-way-ds-max"] = types.YLeaf{"OneWayDsMax", udpJitterStats.OneWayDsMax}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum"] = types.YLeaf{"OneWayDsSum", udpJitterStats.OneWayDsSum}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum2"] = types.YLeaf{"OneWayDsSum2", udpJitterStats.OneWayDsSum2}
    return &(udpJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours
// Table of statistics aggregated over 1-hour
// intervals
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics aggregated for a 1-hour interval. The type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour.
    Hour []Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour
}

func (hours *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours) GetEntityData() *types.CommonEntityData {
    hours.EntityData.YFilter = hours.YFilter
    hours.EntityData.YangName = "hours"
    hours.EntityData.BundleName = "cisco_ios_xr"
    hours.EntityData.ParentYangName = "aggregated"
    hours.EntityData.SegmentPath = "hours"
    hours.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hours.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hours.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hours.EntityData.Children = make(map[string]types.YChild)
    hours.EntityData.Children["hour"] = types.YChild{"Hour", nil}
    for i := range hours.Hour {
        hours.EntityData.Children[types.GetSegmentPath(&hours.Hour[i])] = types.YChild{"Hour", &hours.Hour[i]}
    }
    hours.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hours.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour
// Statistics aggregated for a 1-hour
// interval
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Hour Index. The type is interface{} with range:
    // -2147483648..2147483647.
    HourIndex interface{}

    // Statistics aggregated on distribution value intervals for in 1-hour
    // intervals.
    Distributed Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed

    // Statistics aggregated for the total range of values in 1-hour intervals.
    NonDistributed Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed
}

func (hour *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour) GetEntityData() *types.CommonEntityData {
    hour.EntityData.YFilter = hour.YFilter
    hour.EntityData.YangName = "hour"
    hour.EntityData.BundleName = "cisco_ios_xr"
    hour.EntityData.ParentYangName = "hours"
    hour.EntityData.SegmentPath = "hour" + "[hour-index='" + fmt.Sprintf("%v", hour.HourIndex) + "']"
    hour.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hour.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hour.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hour.EntityData.Children = make(map[string]types.YChild)
    hour.EntityData.Children["distributed"] = types.YChild{"Distributed", &hour.Distributed}
    hour.EntityData.Children["non-distributed"] = types.YChild{"NonDistributed", &hour.NonDistributed}
    hour.EntityData.Leafs = make(map[string]types.YLeaf)
    hour.EntityData.Leafs["hour-index"] = types.YLeaf{"HourIndex", hour.HourIndex}
    return &(hour.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed
// Statistics aggregated on distribution
// value intervals for in 1-hour intervals
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of paths identified in the 1-hour interval.
    Paths Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths

    // Distribution statistics for the target node.
    Target Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target
}

func (distributed *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed) GetEntityData() *types.CommonEntityData {
    distributed.EntityData.YFilter = distributed.YFilter
    distributed.EntityData.YangName = "distributed"
    distributed.EntityData.BundleName = "cisco_ios_xr"
    distributed.EntityData.ParentYangName = "hour"
    distributed.EntityData.SegmentPath = "distributed"
    distributed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributed.EntityData.Children = make(map[string]types.YChild)
    distributed.EntityData.Children["paths"] = types.YChild{"Paths", &distributed.Paths}
    distributed.EntityData.Children["target"] = types.YChild{"Target", &distributed.Target}
    distributed.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(distributed.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths
// Table of paths identified in the 1-hour
// interval
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Paths identified in a 1-hour interval. The type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path.
    Path []Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path
}

func (paths *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "distributed"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range paths.Path {
        paths.EntityData.Children[types.GetSegmentPath(&paths.Path[i])] = types.YChild{"Path", &paths.Path[i]}
    }
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(paths.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path
// Paths identified in a 1-hour interval
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path Index. The type is interface{} with range:
    // -2147483648..2147483647.
    PathIndex interface{}

    // Table of hops for a particular path.
    Hops Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops
}

func (path *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + "[path-index='" + fmt.Sprintf("%v", path.PathIndex) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Children["hops"] = types.YChild{"Hops", &path.Hops}
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["path-index"] = types.YLeaf{"PathIndex", path.PathIndex}
    return &(path.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops
// Table of hops for a particular path
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 1-hour aggregated statistics for a hop in a path-enabled operation. The
    // type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop.
    Hop []Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop
}

func (hops *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "path"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = make(map[string]types.YChild)
    hops.EntityData.Children["hop"] = types.YChild{"Hop", nil}
    for i := range hops.Hop {
        hops.EntityData.Children[types.GetSegmentPath(&hops.Hop[i])] = types.YChild{"Hop", &hops.Hop[i]}
    }
    hops.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hops.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop
// 1-hour aggregated statistics for a
// hop in a path-enabled operation
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Hop Index. The type is interface{} with range:
    // -2147483648..2147483647.
    HopIndex interface{}

    // Table of distribution intervals for a particular hop.
    DistributionIntervals Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals
}

func (hop *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop) GetEntityData() *types.CommonEntityData {
    hop.EntityData.YFilter = hop.YFilter
    hop.EntityData.YangName = "hop"
    hop.EntityData.BundleName = "cisco_ios_xr"
    hop.EntityData.ParentYangName = "hops"
    hop.EntityData.SegmentPath = "hop" + "[hop-index='" + fmt.Sprintf("%v", hop.HopIndex) + "']"
    hop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hop.EntityData.Children = make(map[string]types.YChild)
    hop.EntityData.Children["distribution-intervals"] = types.YChild{"DistributionIntervals", &hop.DistributionIntervals}
    hop.EntityData.Leafs = make(map[string]types.YLeaf)
    hop.EntityData.Leafs["hop-index"] = types.YLeaf{"HopIndex", hop.HopIndex}
    return &(hop.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals
// Table of distribution intervals for a particular
// hop
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 1-hour aggregated statistics for a hop in a path-enabled operation. The
    // type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval.
    DistributionInterval []Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval
}

func (distributionIntervals *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals) GetEntityData() *types.CommonEntityData {
    distributionIntervals.EntityData.YFilter = distributionIntervals.YFilter
    distributionIntervals.EntityData.YangName = "distribution-intervals"
    distributionIntervals.EntityData.BundleName = "cisco_ios_xr"
    distributionIntervals.EntityData.ParentYangName = "hop"
    distributionIntervals.EntityData.SegmentPath = "distribution-intervals"
    distributionIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributionIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributionIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributionIntervals.EntityData.Children = make(map[string]types.YChild)
    distributionIntervals.EntityData.Children["distribution-interval"] = types.YChild{"DistributionInterval", nil}
    for i := range distributionIntervals.DistributionInterval {
        distributionIntervals.EntityData.Children[types.GetSegmentPath(&distributionIntervals.DistributionInterval[i])] = types.YChild{"DistributionInterval", &distributionIntervals.DistributionInterval[i]}
    }
    distributionIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(distributionIntervals.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval
// 1-hour aggregated statistics for a hop in a
// path-enabled operation
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Distribution Interval. The type is interface{}
    // with range: -2147483648..2147483647.
    DistributionIndex interface{}

    // Common Stats.
    CommonStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_CommonStats

    // Operation Specific Stats.
    SpecificStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats
}

func (distributionInterval *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval) GetEntityData() *types.CommonEntityData {
    distributionInterval.EntityData.YFilter = distributionInterval.YFilter
    distributionInterval.EntityData.YangName = "distribution-interval"
    distributionInterval.EntityData.BundleName = "cisco_ios_xr"
    distributionInterval.EntityData.ParentYangName = "distribution-intervals"
    distributionInterval.EntityData.SegmentPath = "distribution-interval" + "[distribution-index='" + fmt.Sprintf("%v", distributionInterval.DistributionIndex) + "']"
    distributionInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributionInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributionInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributionInterval.EntityData.Children = make(map[string]types.YChild)
    distributionInterval.EntityData.Children["common-stats"] = types.YChild{"CommonStats", &distributionInterval.CommonStats}
    distributionInterval.EntityData.Children["specific-stats"] = types.YChild{"SpecificStats", &distributionInterval.SpecificStats}
    distributionInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    distributionInterval.EntityData.Leafs["distribution-index"] = types.YLeaf{"DistributionIndex", distributionInterval.DistributionIndex}
    return &(distributionInterval.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_CommonStats
// Common Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_CommonStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation Time. The type is interface{} with range:
    // 0..18446744073709551615.
    OperationTime interface{}

    // Return code. The type is IpslaRetCode.
    ReturnCode interface{}

    // Number of RTT samples used for the statistics. The type is interface{} with
    // range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of updates processed. The type is interface{} with range:
    // 0..4294967295.
    UpdateCount interface{}

    // Number of updates with Okay return code. The type is interface{} with
    // range: 0..4294967295.
    OkCount interface{}

    // Number of updates with Disconnected return code. The type is interface{}
    // with range: 0..4294967295.
    DisconnectCount interface{}

    // Number of updates with Timeout return code. The type is interface{} with
    // range: 0..4294967295.
    TimeoutCount interface{}

    // Number of updates with Busy return code. The type is interface{} with
    // range: 0..4294967295.
    BusyCount interface{}

    // Number of updates with NotConnected return code. The type is interface{}
    // with range: 0..4294967295.
    NoConnectionCount interface{}

    // Number of updates with Dropped return code. The type is interface{} with
    // range: 0..4294967295.
    DroppedCount interface{}

    // Number of updates with InternalError return code. The type is interface{}
    // with range: 0..4294967295.
    InternalErrorCount interface{}

    // Number of updates with SeqError return code. The type is interface{} with
    // range: 0..4294967295.
    SequenceErrorCount interface{}

    // Number of updates with VerifyError return code. The type is interface{}
    // with range: 0..4294967295.
    VerifyErrorCount interface{}
}

func (commonStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_CommonStats) GetEntityData() *types.CommonEntityData {
    commonStats.EntityData.YFilter = commonStats.YFilter
    commonStats.EntityData.YangName = "common-stats"
    commonStats.EntityData.BundleName = "cisco_ios_xr"
    commonStats.EntityData.ParentYangName = "distribution-interval"
    commonStats.EntityData.SegmentPath = "common-stats"
    commonStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonStats.EntityData.Children = make(map[string]types.YChild)
    commonStats.EntityData.Leafs = make(map[string]types.YLeaf)
    commonStats.EntityData.Leafs["operation-time"] = types.YLeaf{"OperationTime", commonStats.OperationTime}
    commonStats.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", commonStats.ReturnCode}
    commonStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", commonStats.ResponseTimeCount}
    commonStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", commonStats.ResponseTime}
    commonStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", commonStats.MinResponseTime}
    commonStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", commonStats.MaxResponseTime}
    commonStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", commonStats.SumResponseTime}
    commonStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", commonStats.Sum2ResponseTime}
    commonStats.EntityData.Leafs["update-count"] = types.YLeaf{"UpdateCount", commonStats.UpdateCount}
    commonStats.EntityData.Leafs["ok-count"] = types.YLeaf{"OkCount", commonStats.OkCount}
    commonStats.EntityData.Leafs["disconnect-count"] = types.YLeaf{"DisconnectCount", commonStats.DisconnectCount}
    commonStats.EntityData.Leafs["timeout-count"] = types.YLeaf{"TimeoutCount", commonStats.TimeoutCount}
    commonStats.EntityData.Leafs["busy-count"] = types.YLeaf{"BusyCount", commonStats.BusyCount}
    commonStats.EntityData.Leafs["no-connection-count"] = types.YLeaf{"NoConnectionCount", commonStats.NoConnectionCount}
    commonStats.EntityData.Leafs["dropped-count"] = types.YLeaf{"DroppedCount", commonStats.DroppedCount}
    commonStats.EntityData.Leafs["internal-error-count"] = types.YLeaf{"InternalErrorCount", commonStats.InternalErrorCount}
    commonStats.EntityData.Leafs["sequence-error-count"] = types.YLeaf{"SequenceErrorCount", commonStats.SequenceErrorCount}
    commonStats.EntityData.Leafs["verify-error-count"] = types.YLeaf{"VerifyErrorCount", commonStats.VerifyErrorCount}
    return &(commonStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats
// Operation Specific Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // op type. The type is OpTypeEnum.
    OpType interface{}

    // icmp path jitter stats.
    IcmpPathJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats_IcmpPathJitterStats

    // udp jitter stats.
    UdpJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats_UdpJitterStats
}

func (specificStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats) GetEntityData() *types.CommonEntityData {
    specificStats.EntityData.YFilter = specificStats.YFilter
    specificStats.EntityData.YangName = "specific-stats"
    specificStats.EntityData.BundleName = "cisco_ios_xr"
    specificStats.EntityData.ParentYangName = "distribution-interval"
    specificStats.EntityData.SegmentPath = "specific-stats"
    specificStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificStats.EntityData.Children = make(map[string]types.YChild)
    specificStats.EntityData.Children["icmp-path-jitter-stats"] = types.YChild{"IcmpPathJitterStats", &specificStats.IcmpPathJitterStats}
    specificStats.EntityData.Children["udp-jitter-stats"] = types.YChild{"UdpJitterStats", &specificStats.UdpJitterStats}
    specificStats.EntityData.Leafs = make(map[string]types.YLeaf)
    specificStats.EntityData.Leafs["op-type"] = types.YLeaf{"OpType", specificStats.OpType}
    return &(specificStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats_IcmpPathJitterStats
// icmp path jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats_IcmpPathJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address of the source. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // IP Address of the destination. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestAddress interface{}

    // IP address of the hop in the path. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}

    // Interval between echos in ms. The type is interface{} with range:
    // 0..4294967295.
    PacketInterval interface{}

    // Number of RTT samples  used for the statistics. The type is interface{}
    // with range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of Echo replies received . The type is interface{} with range:
    // 0..4294967295.
    PacketCount interface{}

    // Number of packets lost. The type is interface{} with range: 0..4294967295.
    PacketLossCount interface{}

    // Number of out of sequence packets. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceCount interface{}

    // Number of discarded samples. The type is interface{} with range:
    // 0..4294967295.
    DiscardedSampleCount interface{}

    // Number of packets with data corruption. The type is interface{} with range:
    // 0..4294967295.
    VerifyErrorsCount interface{}

    // Number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedErrorCount interface{}

    // Jitter value for this node in the path. The type is interface{} with range:
    // 0..4294967295.
    Jitter interface{}

    // Sum of positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterSum interface{}

    // Sum of squares of positive jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    PosJitterSum2 interface{}

    // Minimum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMin interface{}

    // Maximum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMax interface{}

    // Number of positive jitter values. The type is interface{} with range:
    // 0..4294967295.
    PosJitterCount interface{}

    // Sum of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterSum interface{}

    // Minimum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMin interface{}

    // Maximum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMax interface{}

    // Sum of squares of negative jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    NegJitterSum2 interface{}

    // Number of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterCount interface{}
}

func (icmpPathJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats_IcmpPathJitterStats) GetEntityData() *types.CommonEntityData {
    icmpPathJitterStats.EntityData.YFilter = icmpPathJitterStats.YFilter
    icmpPathJitterStats.EntityData.YangName = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.BundleName = "cisco_ios_xr"
    icmpPathJitterStats.EntityData.ParentYangName = "specific-stats"
    icmpPathJitterStats.EntityData.SegmentPath = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpPathJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpPathJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpPathJitterStats.EntityData.Children = make(map[string]types.YChild)
    icmpPathJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpPathJitterStats.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", icmpPathJitterStats.SourceAddress}
    icmpPathJitterStats.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", icmpPathJitterStats.DestAddress}
    icmpPathJitterStats.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", icmpPathJitterStats.HopAddress}
    icmpPathJitterStats.EntityData.Leafs["packet-interval"] = types.YLeaf{"PacketInterval", icmpPathJitterStats.PacketInterval}
    icmpPathJitterStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", icmpPathJitterStats.ResponseTimeCount}
    icmpPathJitterStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", icmpPathJitterStats.ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", icmpPathJitterStats.MinResponseTime}
    icmpPathJitterStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", icmpPathJitterStats.MaxResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", icmpPathJitterStats.SumResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", icmpPathJitterStats.Sum2ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["packet-count"] = types.YLeaf{"PacketCount", icmpPathJitterStats.PacketCount}
    icmpPathJitterStats.EntityData.Leafs["packet-loss-count"] = types.YLeaf{"PacketLossCount", icmpPathJitterStats.PacketLossCount}
    icmpPathJitterStats.EntityData.Leafs["out-of-sequence-count"] = types.YLeaf{"OutOfSequenceCount", icmpPathJitterStats.OutOfSequenceCount}
    icmpPathJitterStats.EntityData.Leafs["discarded-sample-count"] = types.YLeaf{"DiscardedSampleCount", icmpPathJitterStats.DiscardedSampleCount}
    icmpPathJitterStats.EntityData.Leafs["verify-errors-count"] = types.YLeaf{"VerifyErrorsCount", icmpPathJitterStats.VerifyErrorsCount}
    icmpPathJitterStats.EntityData.Leafs["dropped-error-count"] = types.YLeaf{"DroppedErrorCount", icmpPathJitterStats.DroppedErrorCount}
    icmpPathJitterStats.EntityData.Leafs["jitter"] = types.YLeaf{"Jitter", icmpPathJitterStats.Jitter}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum"] = types.YLeaf{"PosJitterSum", icmpPathJitterStats.PosJitterSum}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum2"] = types.YLeaf{"PosJitterSum2", icmpPathJitterStats.PosJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-min"] = types.YLeaf{"PosJitterMin", icmpPathJitterStats.PosJitterMin}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-max"] = types.YLeaf{"PosJitterMax", icmpPathJitterStats.PosJitterMax}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-count"] = types.YLeaf{"PosJitterCount", icmpPathJitterStats.PosJitterCount}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum"] = types.YLeaf{"NegJitterSum", icmpPathJitterStats.NegJitterSum}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-min"] = types.YLeaf{"NegJitterMin", icmpPathJitterStats.NegJitterMin}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-max"] = types.YLeaf{"NegJitterMax", icmpPathJitterStats.NegJitterMax}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum2"] = types.YLeaf{"NegJitterSum2", icmpPathJitterStats.NegJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-count"] = types.YLeaf{"NegJitterCount", icmpPathJitterStats.NegJitterCount}
    return &(icmpPathJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats_UdpJitterStats
// udp jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats_UdpJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterIn interface{}

    // Output Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterOut interface{}

    // Packets lost in source to destination (SD) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossSd interface{}

    // Packets lost in destination to source (DS) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossDs interface{}

    // Packets out of sequence. The type is interface{} with range: 0..4294967295.
    PacketOutOfSequence interface{}

    // Packets missing in action (cannot determine if theywere lost in SD or DS
    // direction. The type is interface{} with range: 0..4294967295.
    PacketMia interface{}

    // Packets which are skipped. The type is interface{} with range:
    // 0..4294967295.
    PacketSkipped interface{}

    // Packets arriving late. The type is interface{} with range: 0..4294967295.
    PacketLateArrivals interface{}

    // Packets with bad timestamps. The type is interface{} with range:
    // 0..4294967295.
    PacketInvalidTstamp interface{}

    // Number of internal errors. The type is interface{} with range:
    // 0..4294967295.
    InternalErrorsCount interface{}

    // Number of busies. The type is interface{} with range: 0..4294967295.
    BusiesCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive  packets) in SD direction Measured  in milliseconds. The type
    // is interface{} with range: 0..4294967295. Units are millisecond.
    PositiveSdSum interface{}

    // Sum of squares of positive jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveSdSum2 interface{}

    // Minimum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMin interface{}

    // Maximum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMax interface{}

    // Number of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in SD direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeSdSum interface{}

    // Sum of squares of negative jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeSdSum2 interface{}

    // Minimum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMin interface{}

    // Maximum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMax interface{}

    // Number of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    PositiveDsSum interface{}

    // Sum of squares of positive jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveDsSum2 interface{}

    // Minimum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMin interface{}

    // Maximum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMax interface{}

    // Number of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeDsSum interface{}

    // Sum of squares of negative jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeDsSum2 interface{}

    // Minimum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMin interface{}

    // Maximum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMax interface{}

    // Number of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsCount interface{}

    // Number of probe/probe-response pairs used to compute one-way statistics.
    // The type is interface{} with range: 0..4294967295.
    OneWayCount interface{}

    // Minimum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMin interface{}

    // Maximum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMax interface{}

    // Sum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdSum interface{}

    // Sum of squares of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..18446744073709551615.
    OneWaySdSum2 interface{}

    // Minimum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMin interface{}

    // Maximum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMax interface{}

    // Sum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsSum interface{}

    // Sum of squares of the OneWayMinDS and OneWayMaxDS values (msec). The type
    // is interface{} with range: 0..18446744073709551615.
    OneWayDsSum2 interface{}
}

func (udpJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Paths_Path_Hops_Hop_DistributionIntervals_DistributionInterval_SpecificStats_UdpJitterStats) GetEntityData() *types.CommonEntityData {
    udpJitterStats.EntityData.YFilter = udpJitterStats.YFilter
    udpJitterStats.EntityData.YangName = "udp-jitter-stats"
    udpJitterStats.EntityData.BundleName = "cisco_ios_xr"
    udpJitterStats.EntityData.ParentYangName = "specific-stats"
    udpJitterStats.EntityData.SegmentPath = "udp-jitter-stats"
    udpJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpJitterStats.EntityData.Children = make(map[string]types.YChild)
    udpJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    udpJitterStats.EntityData.Leafs["jitter-in"] = types.YLeaf{"JitterIn", udpJitterStats.JitterIn}
    udpJitterStats.EntityData.Leafs["jitter-out"] = types.YLeaf{"JitterOut", udpJitterStats.JitterOut}
    udpJitterStats.EntityData.Leafs["packet-loss-sd"] = types.YLeaf{"PacketLossSd", udpJitterStats.PacketLossSd}
    udpJitterStats.EntityData.Leafs["packet-loss-ds"] = types.YLeaf{"PacketLossDs", udpJitterStats.PacketLossDs}
    udpJitterStats.EntityData.Leafs["packet-out-of-sequence"] = types.YLeaf{"PacketOutOfSequence", udpJitterStats.PacketOutOfSequence}
    udpJitterStats.EntityData.Leafs["packet-mia"] = types.YLeaf{"PacketMia", udpJitterStats.PacketMia}
    udpJitterStats.EntityData.Leafs["packet-skipped"] = types.YLeaf{"PacketSkipped", udpJitterStats.PacketSkipped}
    udpJitterStats.EntityData.Leafs["packet-late-arrivals"] = types.YLeaf{"PacketLateArrivals", udpJitterStats.PacketLateArrivals}
    udpJitterStats.EntityData.Leafs["packet-invalid-tstamp"] = types.YLeaf{"PacketInvalidTstamp", udpJitterStats.PacketInvalidTstamp}
    udpJitterStats.EntityData.Leafs["internal-errors-count"] = types.YLeaf{"InternalErrorsCount", udpJitterStats.InternalErrorsCount}
    udpJitterStats.EntityData.Leafs["busies-count"] = types.YLeaf{"BusiesCount", udpJitterStats.BusiesCount}
    udpJitterStats.EntityData.Leafs["positive-sd-sum"] = types.YLeaf{"PositiveSdSum", udpJitterStats.PositiveSdSum}
    udpJitterStats.EntityData.Leafs["positive-sd-sum2"] = types.YLeaf{"PositiveSdSum2", udpJitterStats.PositiveSdSum2}
    udpJitterStats.EntityData.Leafs["positive-sd-min"] = types.YLeaf{"PositiveSdMin", udpJitterStats.PositiveSdMin}
    udpJitterStats.EntityData.Leafs["positive-sd-max"] = types.YLeaf{"PositiveSdMax", udpJitterStats.PositiveSdMax}
    udpJitterStats.EntityData.Leafs["positive-sd-count"] = types.YLeaf{"PositiveSdCount", udpJitterStats.PositiveSdCount}
    udpJitterStats.EntityData.Leafs["negative-sd-sum"] = types.YLeaf{"NegativeSdSum", udpJitterStats.NegativeSdSum}
    udpJitterStats.EntityData.Leafs["negative-sd-sum2"] = types.YLeaf{"NegativeSdSum2", udpJitterStats.NegativeSdSum2}
    udpJitterStats.EntityData.Leafs["negative-sd-min"] = types.YLeaf{"NegativeSdMin", udpJitterStats.NegativeSdMin}
    udpJitterStats.EntityData.Leafs["negative-sd-max"] = types.YLeaf{"NegativeSdMax", udpJitterStats.NegativeSdMax}
    udpJitterStats.EntityData.Leafs["negative-sd-count"] = types.YLeaf{"NegativeSdCount", udpJitterStats.NegativeSdCount}
    udpJitterStats.EntityData.Leafs["positive-ds-sum"] = types.YLeaf{"PositiveDsSum", udpJitterStats.PositiveDsSum}
    udpJitterStats.EntityData.Leafs["positive-ds-sum2"] = types.YLeaf{"PositiveDsSum2", udpJitterStats.PositiveDsSum2}
    udpJitterStats.EntityData.Leafs["positive-ds-min"] = types.YLeaf{"PositiveDsMin", udpJitterStats.PositiveDsMin}
    udpJitterStats.EntityData.Leafs["positive-ds-max"] = types.YLeaf{"PositiveDsMax", udpJitterStats.PositiveDsMax}
    udpJitterStats.EntityData.Leafs["positive-ds-count"] = types.YLeaf{"PositiveDsCount", udpJitterStats.PositiveDsCount}
    udpJitterStats.EntityData.Leafs["negative-ds-sum"] = types.YLeaf{"NegativeDsSum", udpJitterStats.NegativeDsSum}
    udpJitterStats.EntityData.Leafs["negative-ds-sum2"] = types.YLeaf{"NegativeDsSum2", udpJitterStats.NegativeDsSum2}
    udpJitterStats.EntityData.Leafs["negative-ds-min"] = types.YLeaf{"NegativeDsMin", udpJitterStats.NegativeDsMin}
    udpJitterStats.EntityData.Leafs["negative-ds-max"] = types.YLeaf{"NegativeDsMax", udpJitterStats.NegativeDsMax}
    udpJitterStats.EntityData.Leafs["negative-ds-count"] = types.YLeaf{"NegativeDsCount", udpJitterStats.NegativeDsCount}
    udpJitterStats.EntityData.Leafs["one-way-count"] = types.YLeaf{"OneWayCount", udpJitterStats.OneWayCount}
    udpJitterStats.EntityData.Leafs["one-way-sd-min"] = types.YLeaf{"OneWaySdMin", udpJitterStats.OneWaySdMin}
    udpJitterStats.EntityData.Leafs["one-way-sd-max"] = types.YLeaf{"OneWaySdMax", udpJitterStats.OneWaySdMax}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum"] = types.YLeaf{"OneWaySdSum", udpJitterStats.OneWaySdSum}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum2"] = types.YLeaf{"OneWaySdSum2", udpJitterStats.OneWaySdSum2}
    udpJitterStats.EntityData.Leafs["one-way-ds-min"] = types.YLeaf{"OneWayDsMin", udpJitterStats.OneWayDsMin}
    udpJitterStats.EntityData.Leafs["one-way-ds-max"] = types.YLeaf{"OneWayDsMax", udpJitterStats.OneWayDsMax}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum"] = types.YLeaf{"OneWayDsSum", udpJitterStats.OneWayDsSum}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum2"] = types.YLeaf{"OneWayDsSum2", udpJitterStats.OneWayDsSum2}
    return &(udpJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target
// Distribution statistics for the target
// node
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of distribution intervals for a particular hop.
    DistributionIntervals Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals
}

func (target *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "cisco_ios_xr"
    target.EntityData.ParentYangName = "distributed"
    target.EntityData.SegmentPath = "target"
    target.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    target.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    target.EntityData.Children = make(map[string]types.YChild)
    target.EntityData.Children["distribution-intervals"] = types.YChild{"DistributionIntervals", &target.DistributionIntervals}
    target.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(target.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals
// Table of distribution intervals for a particular
// hop
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 1-hour aggregated statistics for a hop in a path-enabled operation. The
    // type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval.
    DistributionInterval []Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval
}

func (distributionIntervals *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals) GetEntityData() *types.CommonEntityData {
    distributionIntervals.EntityData.YFilter = distributionIntervals.YFilter
    distributionIntervals.EntityData.YangName = "distribution-intervals"
    distributionIntervals.EntityData.BundleName = "cisco_ios_xr"
    distributionIntervals.EntityData.ParentYangName = "target"
    distributionIntervals.EntityData.SegmentPath = "distribution-intervals"
    distributionIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributionIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributionIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributionIntervals.EntityData.Children = make(map[string]types.YChild)
    distributionIntervals.EntityData.Children["distribution-interval"] = types.YChild{"DistributionInterval", nil}
    for i := range distributionIntervals.DistributionInterval {
        distributionIntervals.EntityData.Children[types.GetSegmentPath(&distributionIntervals.DistributionInterval[i])] = types.YChild{"DistributionInterval", &distributionIntervals.DistributionInterval[i]}
    }
    distributionIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(distributionIntervals.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval
// 1-hour aggregated statistics for a hop in a
// path-enabled operation
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Distribution Interval. The type is interface{}
    // with range: -2147483648..2147483647.
    DistributionIndex interface{}

    // Common Stats.
    CommonStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_CommonStats

    // Operation Specific Stats.
    SpecificStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats
}

func (distributionInterval *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval) GetEntityData() *types.CommonEntityData {
    distributionInterval.EntityData.YFilter = distributionInterval.YFilter
    distributionInterval.EntityData.YangName = "distribution-interval"
    distributionInterval.EntityData.BundleName = "cisco_ios_xr"
    distributionInterval.EntityData.ParentYangName = "distribution-intervals"
    distributionInterval.EntityData.SegmentPath = "distribution-interval" + "[distribution-index='" + fmt.Sprintf("%v", distributionInterval.DistributionIndex) + "']"
    distributionInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributionInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributionInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributionInterval.EntityData.Children = make(map[string]types.YChild)
    distributionInterval.EntityData.Children["common-stats"] = types.YChild{"CommonStats", &distributionInterval.CommonStats}
    distributionInterval.EntityData.Children["specific-stats"] = types.YChild{"SpecificStats", &distributionInterval.SpecificStats}
    distributionInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    distributionInterval.EntityData.Leafs["distribution-index"] = types.YLeaf{"DistributionIndex", distributionInterval.DistributionIndex}
    return &(distributionInterval.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_CommonStats
// Common Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_CommonStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation Time. The type is interface{} with range:
    // 0..18446744073709551615.
    OperationTime interface{}

    // Return code. The type is IpslaRetCode.
    ReturnCode interface{}

    // Number of RTT samples used for the statistics. The type is interface{} with
    // range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of updates processed. The type is interface{} with range:
    // 0..4294967295.
    UpdateCount interface{}

    // Number of updates with Okay return code. The type is interface{} with
    // range: 0..4294967295.
    OkCount interface{}

    // Number of updates with Disconnected return code. The type is interface{}
    // with range: 0..4294967295.
    DisconnectCount interface{}

    // Number of updates with Timeout return code. The type is interface{} with
    // range: 0..4294967295.
    TimeoutCount interface{}

    // Number of updates with Busy return code. The type is interface{} with
    // range: 0..4294967295.
    BusyCount interface{}

    // Number of updates with NotConnected return code. The type is interface{}
    // with range: 0..4294967295.
    NoConnectionCount interface{}

    // Number of updates with Dropped return code. The type is interface{} with
    // range: 0..4294967295.
    DroppedCount interface{}

    // Number of updates with InternalError return code. The type is interface{}
    // with range: 0..4294967295.
    InternalErrorCount interface{}

    // Number of updates with SeqError return code. The type is interface{} with
    // range: 0..4294967295.
    SequenceErrorCount interface{}

    // Number of updates with VerifyError return code. The type is interface{}
    // with range: 0..4294967295.
    VerifyErrorCount interface{}
}

func (commonStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_CommonStats) GetEntityData() *types.CommonEntityData {
    commonStats.EntityData.YFilter = commonStats.YFilter
    commonStats.EntityData.YangName = "common-stats"
    commonStats.EntityData.BundleName = "cisco_ios_xr"
    commonStats.EntityData.ParentYangName = "distribution-interval"
    commonStats.EntityData.SegmentPath = "common-stats"
    commonStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonStats.EntityData.Children = make(map[string]types.YChild)
    commonStats.EntityData.Leafs = make(map[string]types.YLeaf)
    commonStats.EntityData.Leafs["operation-time"] = types.YLeaf{"OperationTime", commonStats.OperationTime}
    commonStats.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", commonStats.ReturnCode}
    commonStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", commonStats.ResponseTimeCount}
    commonStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", commonStats.ResponseTime}
    commonStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", commonStats.MinResponseTime}
    commonStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", commonStats.MaxResponseTime}
    commonStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", commonStats.SumResponseTime}
    commonStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", commonStats.Sum2ResponseTime}
    commonStats.EntityData.Leafs["update-count"] = types.YLeaf{"UpdateCount", commonStats.UpdateCount}
    commonStats.EntityData.Leafs["ok-count"] = types.YLeaf{"OkCount", commonStats.OkCount}
    commonStats.EntityData.Leafs["disconnect-count"] = types.YLeaf{"DisconnectCount", commonStats.DisconnectCount}
    commonStats.EntityData.Leafs["timeout-count"] = types.YLeaf{"TimeoutCount", commonStats.TimeoutCount}
    commonStats.EntityData.Leafs["busy-count"] = types.YLeaf{"BusyCount", commonStats.BusyCount}
    commonStats.EntityData.Leafs["no-connection-count"] = types.YLeaf{"NoConnectionCount", commonStats.NoConnectionCount}
    commonStats.EntityData.Leafs["dropped-count"] = types.YLeaf{"DroppedCount", commonStats.DroppedCount}
    commonStats.EntityData.Leafs["internal-error-count"] = types.YLeaf{"InternalErrorCount", commonStats.InternalErrorCount}
    commonStats.EntityData.Leafs["sequence-error-count"] = types.YLeaf{"SequenceErrorCount", commonStats.SequenceErrorCount}
    commonStats.EntityData.Leafs["verify-error-count"] = types.YLeaf{"VerifyErrorCount", commonStats.VerifyErrorCount}
    return &(commonStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats
// Operation Specific Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // op type. The type is OpTypeEnum.
    OpType interface{}

    // icmp path jitter stats.
    IcmpPathJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats_IcmpPathJitterStats

    // udp jitter stats.
    UdpJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats_UdpJitterStats
}

func (specificStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats) GetEntityData() *types.CommonEntityData {
    specificStats.EntityData.YFilter = specificStats.YFilter
    specificStats.EntityData.YangName = "specific-stats"
    specificStats.EntityData.BundleName = "cisco_ios_xr"
    specificStats.EntityData.ParentYangName = "distribution-interval"
    specificStats.EntityData.SegmentPath = "specific-stats"
    specificStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificStats.EntityData.Children = make(map[string]types.YChild)
    specificStats.EntityData.Children["icmp-path-jitter-stats"] = types.YChild{"IcmpPathJitterStats", &specificStats.IcmpPathJitterStats}
    specificStats.EntityData.Children["udp-jitter-stats"] = types.YChild{"UdpJitterStats", &specificStats.UdpJitterStats}
    specificStats.EntityData.Leafs = make(map[string]types.YLeaf)
    specificStats.EntityData.Leafs["op-type"] = types.YLeaf{"OpType", specificStats.OpType}
    return &(specificStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats_IcmpPathJitterStats
// icmp path jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats_IcmpPathJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address of the source. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // IP Address of the destination. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestAddress interface{}

    // IP address of the hop in the path. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}

    // Interval between echos in ms. The type is interface{} with range:
    // 0..4294967295.
    PacketInterval interface{}

    // Number of RTT samples  used for the statistics. The type is interface{}
    // with range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of Echo replies received . The type is interface{} with range:
    // 0..4294967295.
    PacketCount interface{}

    // Number of packets lost. The type is interface{} with range: 0..4294967295.
    PacketLossCount interface{}

    // Number of out of sequence packets. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceCount interface{}

    // Number of discarded samples. The type is interface{} with range:
    // 0..4294967295.
    DiscardedSampleCount interface{}

    // Number of packets with data corruption. The type is interface{} with range:
    // 0..4294967295.
    VerifyErrorsCount interface{}

    // Number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedErrorCount interface{}

    // Jitter value for this node in the path. The type is interface{} with range:
    // 0..4294967295.
    Jitter interface{}

    // Sum of positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterSum interface{}

    // Sum of squares of positive jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    PosJitterSum2 interface{}

    // Minimum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMin interface{}

    // Maximum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMax interface{}

    // Number of positive jitter values. The type is interface{} with range:
    // 0..4294967295.
    PosJitterCount interface{}

    // Sum of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterSum interface{}

    // Minimum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMin interface{}

    // Maximum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMax interface{}

    // Sum of squares of negative jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    NegJitterSum2 interface{}

    // Number of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterCount interface{}
}

func (icmpPathJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats_IcmpPathJitterStats) GetEntityData() *types.CommonEntityData {
    icmpPathJitterStats.EntityData.YFilter = icmpPathJitterStats.YFilter
    icmpPathJitterStats.EntityData.YangName = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.BundleName = "cisco_ios_xr"
    icmpPathJitterStats.EntityData.ParentYangName = "specific-stats"
    icmpPathJitterStats.EntityData.SegmentPath = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpPathJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpPathJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpPathJitterStats.EntityData.Children = make(map[string]types.YChild)
    icmpPathJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpPathJitterStats.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", icmpPathJitterStats.SourceAddress}
    icmpPathJitterStats.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", icmpPathJitterStats.DestAddress}
    icmpPathJitterStats.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", icmpPathJitterStats.HopAddress}
    icmpPathJitterStats.EntityData.Leafs["packet-interval"] = types.YLeaf{"PacketInterval", icmpPathJitterStats.PacketInterval}
    icmpPathJitterStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", icmpPathJitterStats.ResponseTimeCount}
    icmpPathJitterStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", icmpPathJitterStats.ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", icmpPathJitterStats.MinResponseTime}
    icmpPathJitterStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", icmpPathJitterStats.MaxResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", icmpPathJitterStats.SumResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", icmpPathJitterStats.Sum2ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["packet-count"] = types.YLeaf{"PacketCount", icmpPathJitterStats.PacketCount}
    icmpPathJitterStats.EntityData.Leafs["packet-loss-count"] = types.YLeaf{"PacketLossCount", icmpPathJitterStats.PacketLossCount}
    icmpPathJitterStats.EntityData.Leafs["out-of-sequence-count"] = types.YLeaf{"OutOfSequenceCount", icmpPathJitterStats.OutOfSequenceCount}
    icmpPathJitterStats.EntityData.Leafs["discarded-sample-count"] = types.YLeaf{"DiscardedSampleCount", icmpPathJitterStats.DiscardedSampleCount}
    icmpPathJitterStats.EntityData.Leafs["verify-errors-count"] = types.YLeaf{"VerifyErrorsCount", icmpPathJitterStats.VerifyErrorsCount}
    icmpPathJitterStats.EntityData.Leafs["dropped-error-count"] = types.YLeaf{"DroppedErrorCount", icmpPathJitterStats.DroppedErrorCount}
    icmpPathJitterStats.EntityData.Leafs["jitter"] = types.YLeaf{"Jitter", icmpPathJitterStats.Jitter}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum"] = types.YLeaf{"PosJitterSum", icmpPathJitterStats.PosJitterSum}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum2"] = types.YLeaf{"PosJitterSum2", icmpPathJitterStats.PosJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-min"] = types.YLeaf{"PosJitterMin", icmpPathJitterStats.PosJitterMin}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-max"] = types.YLeaf{"PosJitterMax", icmpPathJitterStats.PosJitterMax}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-count"] = types.YLeaf{"PosJitterCount", icmpPathJitterStats.PosJitterCount}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum"] = types.YLeaf{"NegJitterSum", icmpPathJitterStats.NegJitterSum}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-min"] = types.YLeaf{"NegJitterMin", icmpPathJitterStats.NegJitterMin}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-max"] = types.YLeaf{"NegJitterMax", icmpPathJitterStats.NegJitterMax}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum2"] = types.YLeaf{"NegJitterSum2", icmpPathJitterStats.NegJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-count"] = types.YLeaf{"NegJitterCount", icmpPathJitterStats.NegJitterCount}
    return &(icmpPathJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats_UdpJitterStats
// udp jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats_UdpJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterIn interface{}

    // Output Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterOut interface{}

    // Packets lost in source to destination (SD) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossSd interface{}

    // Packets lost in destination to source (DS) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossDs interface{}

    // Packets out of sequence. The type is interface{} with range: 0..4294967295.
    PacketOutOfSequence interface{}

    // Packets missing in action (cannot determine if theywere lost in SD or DS
    // direction. The type is interface{} with range: 0..4294967295.
    PacketMia interface{}

    // Packets which are skipped. The type is interface{} with range:
    // 0..4294967295.
    PacketSkipped interface{}

    // Packets arriving late. The type is interface{} with range: 0..4294967295.
    PacketLateArrivals interface{}

    // Packets with bad timestamps. The type is interface{} with range:
    // 0..4294967295.
    PacketInvalidTstamp interface{}

    // Number of internal errors. The type is interface{} with range:
    // 0..4294967295.
    InternalErrorsCount interface{}

    // Number of busies. The type is interface{} with range: 0..4294967295.
    BusiesCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive  packets) in SD direction Measured  in milliseconds. The type
    // is interface{} with range: 0..4294967295. Units are millisecond.
    PositiveSdSum interface{}

    // Sum of squares of positive jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveSdSum2 interface{}

    // Minimum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMin interface{}

    // Maximum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMax interface{}

    // Number of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in SD direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeSdSum interface{}

    // Sum of squares of negative jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeSdSum2 interface{}

    // Minimum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMin interface{}

    // Maximum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMax interface{}

    // Number of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    PositiveDsSum interface{}

    // Sum of squares of positive jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveDsSum2 interface{}

    // Minimum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMin interface{}

    // Maximum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMax interface{}

    // Number of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeDsSum interface{}

    // Sum of squares of negative jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeDsSum2 interface{}

    // Minimum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMin interface{}

    // Maximum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMax interface{}

    // Number of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsCount interface{}

    // Number of probe/probe-response pairs used to compute one-way statistics.
    // The type is interface{} with range: 0..4294967295.
    OneWayCount interface{}

    // Minimum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMin interface{}

    // Maximum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMax interface{}

    // Sum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdSum interface{}

    // Sum of squares of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..18446744073709551615.
    OneWaySdSum2 interface{}

    // Minimum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMin interface{}

    // Maximum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMax interface{}

    // Sum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsSum interface{}

    // Sum of squares of the OneWayMinDS and OneWayMaxDS values (msec). The type
    // is interface{} with range: 0..18446744073709551615.
    OneWayDsSum2 interface{}
}

func (udpJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_Distributed_Target_DistributionIntervals_DistributionInterval_SpecificStats_UdpJitterStats) GetEntityData() *types.CommonEntityData {
    udpJitterStats.EntityData.YFilter = udpJitterStats.YFilter
    udpJitterStats.EntityData.YangName = "udp-jitter-stats"
    udpJitterStats.EntityData.BundleName = "cisco_ios_xr"
    udpJitterStats.EntityData.ParentYangName = "specific-stats"
    udpJitterStats.EntityData.SegmentPath = "udp-jitter-stats"
    udpJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpJitterStats.EntityData.Children = make(map[string]types.YChild)
    udpJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    udpJitterStats.EntityData.Leafs["jitter-in"] = types.YLeaf{"JitterIn", udpJitterStats.JitterIn}
    udpJitterStats.EntityData.Leafs["jitter-out"] = types.YLeaf{"JitterOut", udpJitterStats.JitterOut}
    udpJitterStats.EntityData.Leafs["packet-loss-sd"] = types.YLeaf{"PacketLossSd", udpJitterStats.PacketLossSd}
    udpJitterStats.EntityData.Leafs["packet-loss-ds"] = types.YLeaf{"PacketLossDs", udpJitterStats.PacketLossDs}
    udpJitterStats.EntityData.Leafs["packet-out-of-sequence"] = types.YLeaf{"PacketOutOfSequence", udpJitterStats.PacketOutOfSequence}
    udpJitterStats.EntityData.Leafs["packet-mia"] = types.YLeaf{"PacketMia", udpJitterStats.PacketMia}
    udpJitterStats.EntityData.Leafs["packet-skipped"] = types.YLeaf{"PacketSkipped", udpJitterStats.PacketSkipped}
    udpJitterStats.EntityData.Leafs["packet-late-arrivals"] = types.YLeaf{"PacketLateArrivals", udpJitterStats.PacketLateArrivals}
    udpJitterStats.EntityData.Leafs["packet-invalid-tstamp"] = types.YLeaf{"PacketInvalidTstamp", udpJitterStats.PacketInvalidTstamp}
    udpJitterStats.EntityData.Leafs["internal-errors-count"] = types.YLeaf{"InternalErrorsCount", udpJitterStats.InternalErrorsCount}
    udpJitterStats.EntityData.Leafs["busies-count"] = types.YLeaf{"BusiesCount", udpJitterStats.BusiesCount}
    udpJitterStats.EntityData.Leafs["positive-sd-sum"] = types.YLeaf{"PositiveSdSum", udpJitterStats.PositiveSdSum}
    udpJitterStats.EntityData.Leafs["positive-sd-sum2"] = types.YLeaf{"PositiveSdSum2", udpJitterStats.PositiveSdSum2}
    udpJitterStats.EntityData.Leafs["positive-sd-min"] = types.YLeaf{"PositiveSdMin", udpJitterStats.PositiveSdMin}
    udpJitterStats.EntityData.Leafs["positive-sd-max"] = types.YLeaf{"PositiveSdMax", udpJitterStats.PositiveSdMax}
    udpJitterStats.EntityData.Leafs["positive-sd-count"] = types.YLeaf{"PositiveSdCount", udpJitterStats.PositiveSdCount}
    udpJitterStats.EntityData.Leafs["negative-sd-sum"] = types.YLeaf{"NegativeSdSum", udpJitterStats.NegativeSdSum}
    udpJitterStats.EntityData.Leafs["negative-sd-sum2"] = types.YLeaf{"NegativeSdSum2", udpJitterStats.NegativeSdSum2}
    udpJitterStats.EntityData.Leafs["negative-sd-min"] = types.YLeaf{"NegativeSdMin", udpJitterStats.NegativeSdMin}
    udpJitterStats.EntityData.Leafs["negative-sd-max"] = types.YLeaf{"NegativeSdMax", udpJitterStats.NegativeSdMax}
    udpJitterStats.EntityData.Leafs["negative-sd-count"] = types.YLeaf{"NegativeSdCount", udpJitterStats.NegativeSdCount}
    udpJitterStats.EntityData.Leafs["positive-ds-sum"] = types.YLeaf{"PositiveDsSum", udpJitterStats.PositiveDsSum}
    udpJitterStats.EntityData.Leafs["positive-ds-sum2"] = types.YLeaf{"PositiveDsSum2", udpJitterStats.PositiveDsSum2}
    udpJitterStats.EntityData.Leafs["positive-ds-min"] = types.YLeaf{"PositiveDsMin", udpJitterStats.PositiveDsMin}
    udpJitterStats.EntityData.Leafs["positive-ds-max"] = types.YLeaf{"PositiveDsMax", udpJitterStats.PositiveDsMax}
    udpJitterStats.EntityData.Leafs["positive-ds-count"] = types.YLeaf{"PositiveDsCount", udpJitterStats.PositiveDsCount}
    udpJitterStats.EntityData.Leafs["negative-ds-sum"] = types.YLeaf{"NegativeDsSum", udpJitterStats.NegativeDsSum}
    udpJitterStats.EntityData.Leafs["negative-ds-sum2"] = types.YLeaf{"NegativeDsSum2", udpJitterStats.NegativeDsSum2}
    udpJitterStats.EntityData.Leafs["negative-ds-min"] = types.YLeaf{"NegativeDsMin", udpJitterStats.NegativeDsMin}
    udpJitterStats.EntityData.Leafs["negative-ds-max"] = types.YLeaf{"NegativeDsMax", udpJitterStats.NegativeDsMax}
    udpJitterStats.EntityData.Leafs["negative-ds-count"] = types.YLeaf{"NegativeDsCount", udpJitterStats.NegativeDsCount}
    udpJitterStats.EntityData.Leafs["one-way-count"] = types.YLeaf{"OneWayCount", udpJitterStats.OneWayCount}
    udpJitterStats.EntityData.Leafs["one-way-sd-min"] = types.YLeaf{"OneWaySdMin", udpJitterStats.OneWaySdMin}
    udpJitterStats.EntityData.Leafs["one-way-sd-max"] = types.YLeaf{"OneWaySdMax", udpJitterStats.OneWaySdMax}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum"] = types.YLeaf{"OneWaySdSum", udpJitterStats.OneWaySdSum}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum2"] = types.YLeaf{"OneWaySdSum2", udpJitterStats.OneWaySdSum2}
    udpJitterStats.EntityData.Leafs["one-way-ds-min"] = types.YLeaf{"OneWayDsMin", udpJitterStats.OneWayDsMin}
    udpJitterStats.EntityData.Leafs["one-way-ds-max"] = types.YLeaf{"OneWayDsMax", udpJitterStats.OneWayDsMax}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum"] = types.YLeaf{"OneWayDsSum", udpJitterStats.OneWayDsSum}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum2"] = types.YLeaf{"OneWayDsSum2", udpJitterStats.OneWayDsSum2}
    return &(udpJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed
// Statistics aggregated for the total range
// of values in 1-hour intervals
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total 1-hour aggregated statistics for the target node.
    Target Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target

    // Table of paths identified in the 1-hour interval.
    Paths Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths

    // List of latest LPD paths.
    LpdPaths Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths
}

func (nonDistributed *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed) GetEntityData() *types.CommonEntityData {
    nonDistributed.EntityData.YFilter = nonDistributed.YFilter
    nonDistributed.EntityData.YangName = "non-distributed"
    nonDistributed.EntityData.BundleName = "cisco_ios_xr"
    nonDistributed.EntityData.ParentYangName = "hour"
    nonDistributed.EntityData.SegmentPath = "non-distributed"
    nonDistributed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonDistributed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonDistributed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonDistributed.EntityData.Children = make(map[string]types.YChild)
    nonDistributed.EntityData.Children["target"] = types.YChild{"Target", &nonDistributed.Target}
    nonDistributed.EntityData.Children["paths"] = types.YChild{"Paths", &nonDistributed.Paths}
    nonDistributed.EntityData.Children["lpd-paths"] = types.YChild{"LpdPaths", &nonDistributed.LpdPaths}
    nonDistributed.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nonDistributed.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target
// Total 1-hour aggregated statistics for
// the target node
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Common Stats.
    CommonStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_CommonStats

    // Operation Specific Stats.
    SpecificStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats
}

func (target *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "cisco_ios_xr"
    target.EntityData.ParentYangName = "non-distributed"
    target.EntityData.SegmentPath = "target"
    target.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    target.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    target.EntityData.Children = make(map[string]types.YChild)
    target.EntityData.Children["common-stats"] = types.YChild{"CommonStats", &target.CommonStats}
    target.EntityData.Children["specific-stats"] = types.YChild{"SpecificStats", &target.SpecificStats}
    target.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(target.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_CommonStats
// Common Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_CommonStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation Time. The type is interface{} with range:
    // 0..18446744073709551615.
    OperationTime interface{}

    // Return code. The type is IpslaRetCode.
    ReturnCode interface{}

    // Number of RTT samples used for the statistics. The type is interface{} with
    // range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of updates processed. The type is interface{} with range:
    // 0..4294967295.
    UpdateCount interface{}

    // Number of updates with Okay return code. The type is interface{} with
    // range: 0..4294967295.
    OkCount interface{}

    // Number of updates with Disconnected return code. The type is interface{}
    // with range: 0..4294967295.
    DisconnectCount interface{}

    // Number of updates with Timeout return code. The type is interface{} with
    // range: 0..4294967295.
    TimeoutCount interface{}

    // Number of updates with Busy return code. The type is interface{} with
    // range: 0..4294967295.
    BusyCount interface{}

    // Number of updates with NotConnected return code. The type is interface{}
    // with range: 0..4294967295.
    NoConnectionCount interface{}

    // Number of updates with Dropped return code. The type is interface{} with
    // range: 0..4294967295.
    DroppedCount interface{}

    // Number of updates with InternalError return code. The type is interface{}
    // with range: 0..4294967295.
    InternalErrorCount interface{}

    // Number of updates with SeqError return code. The type is interface{} with
    // range: 0..4294967295.
    SequenceErrorCount interface{}

    // Number of updates with VerifyError return code. The type is interface{}
    // with range: 0..4294967295.
    VerifyErrorCount interface{}
}

func (commonStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_CommonStats) GetEntityData() *types.CommonEntityData {
    commonStats.EntityData.YFilter = commonStats.YFilter
    commonStats.EntityData.YangName = "common-stats"
    commonStats.EntityData.BundleName = "cisco_ios_xr"
    commonStats.EntityData.ParentYangName = "target"
    commonStats.EntityData.SegmentPath = "common-stats"
    commonStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonStats.EntityData.Children = make(map[string]types.YChild)
    commonStats.EntityData.Leafs = make(map[string]types.YLeaf)
    commonStats.EntityData.Leafs["operation-time"] = types.YLeaf{"OperationTime", commonStats.OperationTime}
    commonStats.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", commonStats.ReturnCode}
    commonStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", commonStats.ResponseTimeCount}
    commonStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", commonStats.ResponseTime}
    commonStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", commonStats.MinResponseTime}
    commonStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", commonStats.MaxResponseTime}
    commonStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", commonStats.SumResponseTime}
    commonStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", commonStats.Sum2ResponseTime}
    commonStats.EntityData.Leafs["update-count"] = types.YLeaf{"UpdateCount", commonStats.UpdateCount}
    commonStats.EntityData.Leafs["ok-count"] = types.YLeaf{"OkCount", commonStats.OkCount}
    commonStats.EntityData.Leafs["disconnect-count"] = types.YLeaf{"DisconnectCount", commonStats.DisconnectCount}
    commonStats.EntityData.Leafs["timeout-count"] = types.YLeaf{"TimeoutCount", commonStats.TimeoutCount}
    commonStats.EntityData.Leafs["busy-count"] = types.YLeaf{"BusyCount", commonStats.BusyCount}
    commonStats.EntityData.Leafs["no-connection-count"] = types.YLeaf{"NoConnectionCount", commonStats.NoConnectionCount}
    commonStats.EntityData.Leafs["dropped-count"] = types.YLeaf{"DroppedCount", commonStats.DroppedCount}
    commonStats.EntityData.Leafs["internal-error-count"] = types.YLeaf{"InternalErrorCount", commonStats.InternalErrorCount}
    commonStats.EntityData.Leafs["sequence-error-count"] = types.YLeaf{"SequenceErrorCount", commonStats.SequenceErrorCount}
    commonStats.EntityData.Leafs["verify-error-count"] = types.YLeaf{"VerifyErrorCount", commonStats.VerifyErrorCount}
    return &(commonStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats
// Operation Specific Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // op type. The type is OpTypeEnum.
    OpType interface{}

    // icmp path jitter stats.
    IcmpPathJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats_IcmpPathJitterStats

    // udp jitter stats.
    UdpJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats_UdpJitterStats
}

func (specificStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats) GetEntityData() *types.CommonEntityData {
    specificStats.EntityData.YFilter = specificStats.YFilter
    specificStats.EntityData.YangName = "specific-stats"
    specificStats.EntityData.BundleName = "cisco_ios_xr"
    specificStats.EntityData.ParentYangName = "target"
    specificStats.EntityData.SegmentPath = "specific-stats"
    specificStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificStats.EntityData.Children = make(map[string]types.YChild)
    specificStats.EntityData.Children["icmp-path-jitter-stats"] = types.YChild{"IcmpPathJitterStats", &specificStats.IcmpPathJitterStats}
    specificStats.EntityData.Children["udp-jitter-stats"] = types.YChild{"UdpJitterStats", &specificStats.UdpJitterStats}
    specificStats.EntityData.Leafs = make(map[string]types.YLeaf)
    specificStats.EntityData.Leafs["op-type"] = types.YLeaf{"OpType", specificStats.OpType}
    return &(specificStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats_IcmpPathJitterStats
// icmp path jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats_IcmpPathJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address of the source. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // IP Address of the destination. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestAddress interface{}

    // IP address of the hop in the path. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}

    // Interval between echos in ms. The type is interface{} with range:
    // 0..4294967295.
    PacketInterval interface{}

    // Number of RTT samples  used for the statistics. The type is interface{}
    // with range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of Echo replies received . The type is interface{} with range:
    // 0..4294967295.
    PacketCount interface{}

    // Number of packets lost. The type is interface{} with range: 0..4294967295.
    PacketLossCount interface{}

    // Number of out of sequence packets. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceCount interface{}

    // Number of discarded samples. The type is interface{} with range:
    // 0..4294967295.
    DiscardedSampleCount interface{}

    // Number of packets with data corruption. The type is interface{} with range:
    // 0..4294967295.
    VerifyErrorsCount interface{}

    // Number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedErrorCount interface{}

    // Jitter value for this node in the path. The type is interface{} with range:
    // 0..4294967295.
    Jitter interface{}

    // Sum of positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterSum interface{}

    // Sum of squares of positive jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    PosJitterSum2 interface{}

    // Minimum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMin interface{}

    // Maximum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMax interface{}

    // Number of positive jitter values. The type is interface{} with range:
    // 0..4294967295.
    PosJitterCount interface{}

    // Sum of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterSum interface{}

    // Minimum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMin interface{}

    // Maximum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMax interface{}

    // Sum of squares of negative jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    NegJitterSum2 interface{}

    // Number of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterCount interface{}
}

func (icmpPathJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats_IcmpPathJitterStats) GetEntityData() *types.CommonEntityData {
    icmpPathJitterStats.EntityData.YFilter = icmpPathJitterStats.YFilter
    icmpPathJitterStats.EntityData.YangName = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.BundleName = "cisco_ios_xr"
    icmpPathJitterStats.EntityData.ParentYangName = "specific-stats"
    icmpPathJitterStats.EntityData.SegmentPath = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpPathJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpPathJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpPathJitterStats.EntityData.Children = make(map[string]types.YChild)
    icmpPathJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpPathJitterStats.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", icmpPathJitterStats.SourceAddress}
    icmpPathJitterStats.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", icmpPathJitterStats.DestAddress}
    icmpPathJitterStats.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", icmpPathJitterStats.HopAddress}
    icmpPathJitterStats.EntityData.Leafs["packet-interval"] = types.YLeaf{"PacketInterval", icmpPathJitterStats.PacketInterval}
    icmpPathJitterStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", icmpPathJitterStats.ResponseTimeCount}
    icmpPathJitterStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", icmpPathJitterStats.ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", icmpPathJitterStats.MinResponseTime}
    icmpPathJitterStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", icmpPathJitterStats.MaxResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", icmpPathJitterStats.SumResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", icmpPathJitterStats.Sum2ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["packet-count"] = types.YLeaf{"PacketCount", icmpPathJitterStats.PacketCount}
    icmpPathJitterStats.EntityData.Leafs["packet-loss-count"] = types.YLeaf{"PacketLossCount", icmpPathJitterStats.PacketLossCount}
    icmpPathJitterStats.EntityData.Leafs["out-of-sequence-count"] = types.YLeaf{"OutOfSequenceCount", icmpPathJitterStats.OutOfSequenceCount}
    icmpPathJitterStats.EntityData.Leafs["discarded-sample-count"] = types.YLeaf{"DiscardedSampleCount", icmpPathJitterStats.DiscardedSampleCount}
    icmpPathJitterStats.EntityData.Leafs["verify-errors-count"] = types.YLeaf{"VerifyErrorsCount", icmpPathJitterStats.VerifyErrorsCount}
    icmpPathJitterStats.EntityData.Leafs["dropped-error-count"] = types.YLeaf{"DroppedErrorCount", icmpPathJitterStats.DroppedErrorCount}
    icmpPathJitterStats.EntityData.Leafs["jitter"] = types.YLeaf{"Jitter", icmpPathJitterStats.Jitter}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum"] = types.YLeaf{"PosJitterSum", icmpPathJitterStats.PosJitterSum}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum2"] = types.YLeaf{"PosJitterSum2", icmpPathJitterStats.PosJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-min"] = types.YLeaf{"PosJitterMin", icmpPathJitterStats.PosJitterMin}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-max"] = types.YLeaf{"PosJitterMax", icmpPathJitterStats.PosJitterMax}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-count"] = types.YLeaf{"PosJitterCount", icmpPathJitterStats.PosJitterCount}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum"] = types.YLeaf{"NegJitterSum", icmpPathJitterStats.NegJitterSum}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-min"] = types.YLeaf{"NegJitterMin", icmpPathJitterStats.NegJitterMin}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-max"] = types.YLeaf{"NegJitterMax", icmpPathJitterStats.NegJitterMax}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum2"] = types.YLeaf{"NegJitterSum2", icmpPathJitterStats.NegJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-count"] = types.YLeaf{"NegJitterCount", icmpPathJitterStats.NegJitterCount}
    return &(icmpPathJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats_UdpJitterStats
// udp jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats_UdpJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterIn interface{}

    // Output Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterOut interface{}

    // Packets lost in source to destination (SD) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossSd interface{}

    // Packets lost in destination to source (DS) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossDs interface{}

    // Packets out of sequence. The type is interface{} with range: 0..4294967295.
    PacketOutOfSequence interface{}

    // Packets missing in action (cannot determine if theywere lost in SD or DS
    // direction. The type is interface{} with range: 0..4294967295.
    PacketMia interface{}

    // Packets which are skipped. The type is interface{} with range:
    // 0..4294967295.
    PacketSkipped interface{}

    // Packets arriving late. The type is interface{} with range: 0..4294967295.
    PacketLateArrivals interface{}

    // Packets with bad timestamps. The type is interface{} with range:
    // 0..4294967295.
    PacketInvalidTstamp interface{}

    // Number of internal errors. The type is interface{} with range:
    // 0..4294967295.
    InternalErrorsCount interface{}

    // Number of busies. The type is interface{} with range: 0..4294967295.
    BusiesCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive  packets) in SD direction Measured  in milliseconds. The type
    // is interface{} with range: 0..4294967295. Units are millisecond.
    PositiveSdSum interface{}

    // Sum of squares of positive jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveSdSum2 interface{}

    // Minimum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMin interface{}

    // Maximum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMax interface{}

    // Number of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in SD direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeSdSum interface{}

    // Sum of squares of negative jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeSdSum2 interface{}

    // Minimum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMin interface{}

    // Maximum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMax interface{}

    // Number of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    PositiveDsSum interface{}

    // Sum of squares of positive jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveDsSum2 interface{}

    // Minimum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMin interface{}

    // Maximum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMax interface{}

    // Number of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeDsSum interface{}

    // Sum of squares of negative jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeDsSum2 interface{}

    // Minimum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMin interface{}

    // Maximum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMax interface{}

    // Number of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsCount interface{}

    // Number of probe/probe-response pairs used to compute one-way statistics.
    // The type is interface{} with range: 0..4294967295.
    OneWayCount interface{}

    // Minimum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMin interface{}

    // Maximum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMax interface{}

    // Sum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdSum interface{}

    // Sum of squares of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..18446744073709551615.
    OneWaySdSum2 interface{}

    // Minimum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMin interface{}

    // Maximum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMax interface{}

    // Sum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsSum interface{}

    // Sum of squares of the OneWayMinDS and OneWayMaxDS values (msec). The type
    // is interface{} with range: 0..18446744073709551615.
    OneWayDsSum2 interface{}
}

func (udpJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Target_SpecificStats_UdpJitterStats) GetEntityData() *types.CommonEntityData {
    udpJitterStats.EntityData.YFilter = udpJitterStats.YFilter
    udpJitterStats.EntityData.YangName = "udp-jitter-stats"
    udpJitterStats.EntityData.BundleName = "cisco_ios_xr"
    udpJitterStats.EntityData.ParentYangName = "specific-stats"
    udpJitterStats.EntityData.SegmentPath = "udp-jitter-stats"
    udpJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpJitterStats.EntityData.Children = make(map[string]types.YChild)
    udpJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    udpJitterStats.EntityData.Leafs["jitter-in"] = types.YLeaf{"JitterIn", udpJitterStats.JitterIn}
    udpJitterStats.EntityData.Leafs["jitter-out"] = types.YLeaf{"JitterOut", udpJitterStats.JitterOut}
    udpJitterStats.EntityData.Leafs["packet-loss-sd"] = types.YLeaf{"PacketLossSd", udpJitterStats.PacketLossSd}
    udpJitterStats.EntityData.Leafs["packet-loss-ds"] = types.YLeaf{"PacketLossDs", udpJitterStats.PacketLossDs}
    udpJitterStats.EntityData.Leafs["packet-out-of-sequence"] = types.YLeaf{"PacketOutOfSequence", udpJitterStats.PacketOutOfSequence}
    udpJitterStats.EntityData.Leafs["packet-mia"] = types.YLeaf{"PacketMia", udpJitterStats.PacketMia}
    udpJitterStats.EntityData.Leafs["packet-skipped"] = types.YLeaf{"PacketSkipped", udpJitterStats.PacketSkipped}
    udpJitterStats.EntityData.Leafs["packet-late-arrivals"] = types.YLeaf{"PacketLateArrivals", udpJitterStats.PacketLateArrivals}
    udpJitterStats.EntityData.Leafs["packet-invalid-tstamp"] = types.YLeaf{"PacketInvalidTstamp", udpJitterStats.PacketInvalidTstamp}
    udpJitterStats.EntityData.Leafs["internal-errors-count"] = types.YLeaf{"InternalErrorsCount", udpJitterStats.InternalErrorsCount}
    udpJitterStats.EntityData.Leafs["busies-count"] = types.YLeaf{"BusiesCount", udpJitterStats.BusiesCount}
    udpJitterStats.EntityData.Leafs["positive-sd-sum"] = types.YLeaf{"PositiveSdSum", udpJitterStats.PositiveSdSum}
    udpJitterStats.EntityData.Leafs["positive-sd-sum2"] = types.YLeaf{"PositiveSdSum2", udpJitterStats.PositiveSdSum2}
    udpJitterStats.EntityData.Leafs["positive-sd-min"] = types.YLeaf{"PositiveSdMin", udpJitterStats.PositiveSdMin}
    udpJitterStats.EntityData.Leafs["positive-sd-max"] = types.YLeaf{"PositiveSdMax", udpJitterStats.PositiveSdMax}
    udpJitterStats.EntityData.Leafs["positive-sd-count"] = types.YLeaf{"PositiveSdCount", udpJitterStats.PositiveSdCount}
    udpJitterStats.EntityData.Leafs["negative-sd-sum"] = types.YLeaf{"NegativeSdSum", udpJitterStats.NegativeSdSum}
    udpJitterStats.EntityData.Leafs["negative-sd-sum2"] = types.YLeaf{"NegativeSdSum2", udpJitterStats.NegativeSdSum2}
    udpJitterStats.EntityData.Leafs["negative-sd-min"] = types.YLeaf{"NegativeSdMin", udpJitterStats.NegativeSdMin}
    udpJitterStats.EntityData.Leafs["negative-sd-max"] = types.YLeaf{"NegativeSdMax", udpJitterStats.NegativeSdMax}
    udpJitterStats.EntityData.Leafs["negative-sd-count"] = types.YLeaf{"NegativeSdCount", udpJitterStats.NegativeSdCount}
    udpJitterStats.EntityData.Leafs["positive-ds-sum"] = types.YLeaf{"PositiveDsSum", udpJitterStats.PositiveDsSum}
    udpJitterStats.EntityData.Leafs["positive-ds-sum2"] = types.YLeaf{"PositiveDsSum2", udpJitterStats.PositiveDsSum2}
    udpJitterStats.EntityData.Leafs["positive-ds-min"] = types.YLeaf{"PositiveDsMin", udpJitterStats.PositiveDsMin}
    udpJitterStats.EntityData.Leafs["positive-ds-max"] = types.YLeaf{"PositiveDsMax", udpJitterStats.PositiveDsMax}
    udpJitterStats.EntityData.Leafs["positive-ds-count"] = types.YLeaf{"PositiveDsCount", udpJitterStats.PositiveDsCount}
    udpJitterStats.EntityData.Leafs["negative-ds-sum"] = types.YLeaf{"NegativeDsSum", udpJitterStats.NegativeDsSum}
    udpJitterStats.EntityData.Leafs["negative-ds-sum2"] = types.YLeaf{"NegativeDsSum2", udpJitterStats.NegativeDsSum2}
    udpJitterStats.EntityData.Leafs["negative-ds-min"] = types.YLeaf{"NegativeDsMin", udpJitterStats.NegativeDsMin}
    udpJitterStats.EntityData.Leafs["negative-ds-max"] = types.YLeaf{"NegativeDsMax", udpJitterStats.NegativeDsMax}
    udpJitterStats.EntityData.Leafs["negative-ds-count"] = types.YLeaf{"NegativeDsCount", udpJitterStats.NegativeDsCount}
    udpJitterStats.EntityData.Leafs["one-way-count"] = types.YLeaf{"OneWayCount", udpJitterStats.OneWayCount}
    udpJitterStats.EntityData.Leafs["one-way-sd-min"] = types.YLeaf{"OneWaySdMin", udpJitterStats.OneWaySdMin}
    udpJitterStats.EntityData.Leafs["one-way-sd-max"] = types.YLeaf{"OneWaySdMax", udpJitterStats.OneWaySdMax}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum"] = types.YLeaf{"OneWaySdSum", udpJitterStats.OneWaySdSum}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum2"] = types.YLeaf{"OneWaySdSum2", udpJitterStats.OneWaySdSum2}
    udpJitterStats.EntityData.Leafs["one-way-ds-min"] = types.YLeaf{"OneWayDsMin", udpJitterStats.OneWayDsMin}
    udpJitterStats.EntityData.Leafs["one-way-ds-max"] = types.YLeaf{"OneWayDsMax", udpJitterStats.OneWayDsMax}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum"] = types.YLeaf{"OneWayDsSum", udpJitterStats.OneWayDsSum}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum2"] = types.YLeaf{"OneWayDsSum2", udpJitterStats.OneWayDsSum2}
    return &(udpJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths
// Table of paths identified in the 1-hour
// interval
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Paths identified in a 1-hour interval. The type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path.
    Path []Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path
}

func (paths *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "non-distributed"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range paths.Path {
        paths.EntityData.Children[types.GetSegmentPath(&paths.Path[i])] = types.YChild{"Path", &paths.Path[i]}
    }
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(paths.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path
// Paths identified in a 1-hour interval
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path Index. The type is interface{} with range:
    // -2147483648..2147483647.
    PathIndex interface{}

    // Table of hops for a particular path.
    Hops Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops
}

func (path *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + "[path-index='" + fmt.Sprintf("%v", path.PathIndex) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Children["hops"] = types.YChild{"Hops", &path.Hops}
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["path-index"] = types.YLeaf{"PathIndex", path.PathIndex}
    return &(path.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops
// Table of hops for a particular path
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total 1-hour aggregated statistics for a hop in a path-enabled operation.
    // The type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop.
    Hop []Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop
}

func (hops *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "path"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = make(map[string]types.YChild)
    hops.EntityData.Children["hop"] = types.YChild{"Hop", nil}
    for i := range hops.Hop {
        hops.EntityData.Children[types.GetSegmentPath(&hops.Hop[i])] = types.YChild{"Hop", &hops.Hop[i]}
    }
    hops.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hops.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop
// Total 1-hour aggregated statistics
// for a hop in a path-enabled operation
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Hop Index. The type is interface{} with range:
    // -2147483648..2147483647.
    HopIndex interface{}

    // Common Stats.
    CommonStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_CommonStats

    // Operation Specific Stats.
    SpecificStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats
}

func (hop *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop) GetEntityData() *types.CommonEntityData {
    hop.EntityData.YFilter = hop.YFilter
    hop.EntityData.YangName = "hop"
    hop.EntityData.BundleName = "cisco_ios_xr"
    hop.EntityData.ParentYangName = "hops"
    hop.EntityData.SegmentPath = "hop" + "[hop-index='" + fmt.Sprintf("%v", hop.HopIndex) + "']"
    hop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hop.EntityData.Children = make(map[string]types.YChild)
    hop.EntityData.Children["common-stats"] = types.YChild{"CommonStats", &hop.CommonStats}
    hop.EntityData.Children["specific-stats"] = types.YChild{"SpecificStats", &hop.SpecificStats}
    hop.EntityData.Leafs = make(map[string]types.YLeaf)
    hop.EntityData.Leafs["hop-index"] = types.YLeaf{"HopIndex", hop.HopIndex}
    return &(hop.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_CommonStats
// Common Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_CommonStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation Time. The type is interface{} with range:
    // 0..18446744073709551615.
    OperationTime interface{}

    // Return code. The type is IpslaRetCode.
    ReturnCode interface{}

    // Number of RTT samples used for the statistics. The type is interface{} with
    // range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of updates processed. The type is interface{} with range:
    // 0..4294967295.
    UpdateCount interface{}

    // Number of updates with Okay return code. The type is interface{} with
    // range: 0..4294967295.
    OkCount interface{}

    // Number of updates with Disconnected return code. The type is interface{}
    // with range: 0..4294967295.
    DisconnectCount interface{}

    // Number of updates with Timeout return code. The type is interface{} with
    // range: 0..4294967295.
    TimeoutCount interface{}

    // Number of updates with Busy return code. The type is interface{} with
    // range: 0..4294967295.
    BusyCount interface{}

    // Number of updates with NotConnected return code. The type is interface{}
    // with range: 0..4294967295.
    NoConnectionCount interface{}

    // Number of updates with Dropped return code. The type is interface{} with
    // range: 0..4294967295.
    DroppedCount interface{}

    // Number of updates with InternalError return code. The type is interface{}
    // with range: 0..4294967295.
    InternalErrorCount interface{}

    // Number of updates with SeqError return code. The type is interface{} with
    // range: 0..4294967295.
    SequenceErrorCount interface{}

    // Number of updates with VerifyError return code. The type is interface{}
    // with range: 0..4294967295.
    VerifyErrorCount interface{}
}

func (commonStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_CommonStats) GetEntityData() *types.CommonEntityData {
    commonStats.EntityData.YFilter = commonStats.YFilter
    commonStats.EntityData.YangName = "common-stats"
    commonStats.EntityData.BundleName = "cisco_ios_xr"
    commonStats.EntityData.ParentYangName = "hop"
    commonStats.EntityData.SegmentPath = "common-stats"
    commonStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonStats.EntityData.Children = make(map[string]types.YChild)
    commonStats.EntityData.Leafs = make(map[string]types.YLeaf)
    commonStats.EntityData.Leafs["operation-time"] = types.YLeaf{"OperationTime", commonStats.OperationTime}
    commonStats.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", commonStats.ReturnCode}
    commonStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", commonStats.ResponseTimeCount}
    commonStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", commonStats.ResponseTime}
    commonStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", commonStats.MinResponseTime}
    commonStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", commonStats.MaxResponseTime}
    commonStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", commonStats.SumResponseTime}
    commonStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", commonStats.Sum2ResponseTime}
    commonStats.EntityData.Leafs["update-count"] = types.YLeaf{"UpdateCount", commonStats.UpdateCount}
    commonStats.EntityData.Leafs["ok-count"] = types.YLeaf{"OkCount", commonStats.OkCount}
    commonStats.EntityData.Leafs["disconnect-count"] = types.YLeaf{"DisconnectCount", commonStats.DisconnectCount}
    commonStats.EntityData.Leafs["timeout-count"] = types.YLeaf{"TimeoutCount", commonStats.TimeoutCount}
    commonStats.EntityData.Leafs["busy-count"] = types.YLeaf{"BusyCount", commonStats.BusyCount}
    commonStats.EntityData.Leafs["no-connection-count"] = types.YLeaf{"NoConnectionCount", commonStats.NoConnectionCount}
    commonStats.EntityData.Leafs["dropped-count"] = types.YLeaf{"DroppedCount", commonStats.DroppedCount}
    commonStats.EntityData.Leafs["internal-error-count"] = types.YLeaf{"InternalErrorCount", commonStats.InternalErrorCount}
    commonStats.EntityData.Leafs["sequence-error-count"] = types.YLeaf{"SequenceErrorCount", commonStats.SequenceErrorCount}
    commonStats.EntityData.Leafs["verify-error-count"] = types.YLeaf{"VerifyErrorCount", commonStats.VerifyErrorCount}
    return &(commonStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats
// Operation Specific Stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // op type. The type is OpTypeEnum.
    OpType interface{}

    // icmp path jitter stats.
    IcmpPathJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats_IcmpPathJitterStats

    // udp jitter stats.
    UdpJitterStats Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats_UdpJitterStats
}

func (specificStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats) GetEntityData() *types.CommonEntityData {
    specificStats.EntityData.YFilter = specificStats.YFilter
    specificStats.EntityData.YangName = "specific-stats"
    specificStats.EntityData.BundleName = "cisco_ios_xr"
    specificStats.EntityData.ParentYangName = "hop"
    specificStats.EntityData.SegmentPath = "specific-stats"
    specificStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificStats.EntityData.Children = make(map[string]types.YChild)
    specificStats.EntityData.Children["icmp-path-jitter-stats"] = types.YChild{"IcmpPathJitterStats", &specificStats.IcmpPathJitterStats}
    specificStats.EntityData.Children["udp-jitter-stats"] = types.YChild{"UdpJitterStats", &specificStats.UdpJitterStats}
    specificStats.EntityData.Leafs = make(map[string]types.YLeaf)
    specificStats.EntityData.Leafs["op-type"] = types.YLeaf{"OpType", specificStats.OpType}
    return &(specificStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats_IcmpPathJitterStats
// icmp path jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats_IcmpPathJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address of the source. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // IP Address of the destination. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestAddress interface{}

    // IP address of the hop in the path. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}

    // Interval between echos in ms. The type is interface{} with range:
    // 0..4294967295.
    PacketInterval interface{}

    // Number of RTT samples  used for the statistics. The type is interface{}
    // with range: 0..4294967295.
    ResponseTimeCount interface{}

    // RTT. The type is interface{} with range: 0..4294967295.
    ResponseTime interface{}

    // Minimum RTT. The type is interface{} with range: 0..4294967295.
    MinResponseTime interface{}

    // Maximum RTT. The type is interface{} with range: 0..4294967295.
    MaxResponseTime interface{}

    // Sum of RTT. The type is interface{} with range: 0..4294967295.
    SumResponseTime interface{}

    // Sum of RTT^2. The type is interface{} with range: 0..18446744073709551615.
    Sum2ResponseTime interface{}

    // Number of Echo replies received . The type is interface{} with range:
    // 0..4294967295.
    PacketCount interface{}

    // Number of packets lost. The type is interface{} with range: 0..4294967295.
    PacketLossCount interface{}

    // Number of out of sequence packets. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceCount interface{}

    // Number of discarded samples. The type is interface{} with range:
    // 0..4294967295.
    DiscardedSampleCount interface{}

    // Number of packets with data corruption. The type is interface{} with range:
    // 0..4294967295.
    VerifyErrorsCount interface{}

    // Number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedErrorCount interface{}

    // Jitter value for this node in the path. The type is interface{} with range:
    // 0..4294967295.
    Jitter interface{}

    // Sum of positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterSum interface{}

    // Sum of squares of positive jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    PosJitterSum2 interface{}

    // Minimum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMin interface{}

    // Maximum positive jitter value. The type is interface{} with range:
    // 0..4294967295.
    PosJitterMax interface{}

    // Number of positive jitter values. The type is interface{} with range:
    // 0..4294967295.
    PosJitterCount interface{}

    // Sum of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterSum interface{}

    // Minimum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMin interface{}

    // Maximum negative jitter value. The type is interface{} with range:
    // 0..4294967295.
    NegJitterMax interface{}

    // Sum of squares of negative jitter values. The type is interface{} with
    // range: 0..18446744073709551615.
    NegJitterSum2 interface{}

    // Number of negative jitter values. The type is interface{} with range:
    // 0..4294967295.
    NegJitterCount interface{}
}

func (icmpPathJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats_IcmpPathJitterStats) GetEntityData() *types.CommonEntityData {
    icmpPathJitterStats.EntityData.YFilter = icmpPathJitterStats.YFilter
    icmpPathJitterStats.EntityData.YangName = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.BundleName = "cisco_ios_xr"
    icmpPathJitterStats.EntityData.ParentYangName = "specific-stats"
    icmpPathJitterStats.EntityData.SegmentPath = "icmp-path-jitter-stats"
    icmpPathJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmpPathJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmpPathJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmpPathJitterStats.EntityData.Children = make(map[string]types.YChild)
    icmpPathJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    icmpPathJitterStats.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", icmpPathJitterStats.SourceAddress}
    icmpPathJitterStats.EntityData.Leafs["dest-address"] = types.YLeaf{"DestAddress", icmpPathJitterStats.DestAddress}
    icmpPathJitterStats.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", icmpPathJitterStats.HopAddress}
    icmpPathJitterStats.EntityData.Leafs["packet-interval"] = types.YLeaf{"PacketInterval", icmpPathJitterStats.PacketInterval}
    icmpPathJitterStats.EntityData.Leafs["response-time-count"] = types.YLeaf{"ResponseTimeCount", icmpPathJitterStats.ResponseTimeCount}
    icmpPathJitterStats.EntityData.Leafs["response-time"] = types.YLeaf{"ResponseTime", icmpPathJitterStats.ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["min-response-time"] = types.YLeaf{"MinResponseTime", icmpPathJitterStats.MinResponseTime}
    icmpPathJitterStats.EntityData.Leafs["max-response-time"] = types.YLeaf{"MaxResponseTime", icmpPathJitterStats.MaxResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum-response-time"] = types.YLeaf{"SumResponseTime", icmpPathJitterStats.SumResponseTime}
    icmpPathJitterStats.EntityData.Leafs["sum2-response-time"] = types.YLeaf{"Sum2ResponseTime", icmpPathJitterStats.Sum2ResponseTime}
    icmpPathJitterStats.EntityData.Leafs["packet-count"] = types.YLeaf{"PacketCount", icmpPathJitterStats.PacketCount}
    icmpPathJitterStats.EntityData.Leafs["packet-loss-count"] = types.YLeaf{"PacketLossCount", icmpPathJitterStats.PacketLossCount}
    icmpPathJitterStats.EntityData.Leafs["out-of-sequence-count"] = types.YLeaf{"OutOfSequenceCount", icmpPathJitterStats.OutOfSequenceCount}
    icmpPathJitterStats.EntityData.Leafs["discarded-sample-count"] = types.YLeaf{"DiscardedSampleCount", icmpPathJitterStats.DiscardedSampleCount}
    icmpPathJitterStats.EntityData.Leafs["verify-errors-count"] = types.YLeaf{"VerifyErrorsCount", icmpPathJitterStats.VerifyErrorsCount}
    icmpPathJitterStats.EntityData.Leafs["dropped-error-count"] = types.YLeaf{"DroppedErrorCount", icmpPathJitterStats.DroppedErrorCount}
    icmpPathJitterStats.EntityData.Leafs["jitter"] = types.YLeaf{"Jitter", icmpPathJitterStats.Jitter}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum"] = types.YLeaf{"PosJitterSum", icmpPathJitterStats.PosJitterSum}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-sum2"] = types.YLeaf{"PosJitterSum2", icmpPathJitterStats.PosJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-min"] = types.YLeaf{"PosJitterMin", icmpPathJitterStats.PosJitterMin}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-max"] = types.YLeaf{"PosJitterMax", icmpPathJitterStats.PosJitterMax}
    icmpPathJitterStats.EntityData.Leafs["pos-jitter-count"] = types.YLeaf{"PosJitterCount", icmpPathJitterStats.PosJitterCount}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum"] = types.YLeaf{"NegJitterSum", icmpPathJitterStats.NegJitterSum}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-min"] = types.YLeaf{"NegJitterMin", icmpPathJitterStats.NegJitterMin}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-max"] = types.YLeaf{"NegJitterMax", icmpPathJitterStats.NegJitterMax}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-sum2"] = types.YLeaf{"NegJitterSum2", icmpPathJitterStats.NegJitterSum2}
    icmpPathJitterStats.EntityData.Leafs["neg-jitter-count"] = types.YLeaf{"NegJitterCount", icmpPathJitterStats.NegJitterCount}
    return &(icmpPathJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats_UdpJitterStats
// udp jitter stats
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats_UdpJitterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Input Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterIn interface{}

    // Output Jitter moving average, computed as per RFC1889. The type is
    // interface{} with range: 0..4294967295.
    JitterOut interface{}

    // Packets lost in source to destination (SD) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossSd interface{}

    // Packets lost in destination to source (DS) direction. The type is
    // interface{} with range: 0..4294967295.
    PacketLossDs interface{}

    // Packets out of sequence. The type is interface{} with range: 0..4294967295.
    PacketOutOfSequence interface{}

    // Packets missing in action (cannot determine if theywere lost in SD or DS
    // direction. The type is interface{} with range: 0..4294967295.
    PacketMia interface{}

    // Packets which are skipped. The type is interface{} with range:
    // 0..4294967295.
    PacketSkipped interface{}

    // Packets arriving late. The type is interface{} with range: 0..4294967295.
    PacketLateArrivals interface{}

    // Packets with bad timestamps. The type is interface{} with range:
    // 0..4294967295.
    PacketInvalidTstamp interface{}

    // Number of internal errors. The type is interface{} with range:
    // 0..4294967295.
    InternalErrorsCount interface{}

    // Number of busies. The type is interface{} with range: 0..4294967295.
    BusiesCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive  packets) in SD direction Measured  in milliseconds. The type
    // is interface{} with range: 0..4294967295. Units are millisecond.
    PositiveSdSum interface{}

    // Sum of squares of positive jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveSdSum2 interface{}

    // Minimum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMin interface{}

    // Maximum of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdMax interface{}

    // Number of positive jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveSdCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in SD direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeSdSum interface{}

    // Sum of squares of negative jitter values in SD direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeSdSum2 interface{}

    // Minimum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMin interface{}

    // Maximum of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdMax interface{}

    // Number of negative jitter values in SD direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeSdCount interface{}

    // Sum of positive jitter values (i.e., network latency increases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    PositiveDsSum interface{}

    // Sum of squares of positive jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    PositiveDsSum2 interface{}

    // Minimum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMin interface{}

    // Maximum of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsMax interface{}

    // Number of positive jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    PositiveDsCount interface{}

    // Sum of negative jitter values (i.e., network latency decreases for two
    // consecutive packets) in DS direction Measured  in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    NegativeDsSum interface{}

    // Sum of squares of negative jitter values in DS direction. The type is
    // interface{} with range: 0..18446744073709551615.
    NegativeDsSum2 interface{}

    // Minimum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMin interface{}

    // Maximum of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsMax interface{}

    // Number of negative jitter values in DS direction. The type is interface{}
    // with range: 0..4294967295.
    NegativeDsCount interface{}

    // Number of probe/probe-response pairs used to compute one-way statistics.
    // The type is interface{} with range: 0..4294967295.
    OneWayCount interface{}

    // Minimum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMin interface{}

    // Maximum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdMax interface{}

    // Sum of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWaySdSum interface{}

    // Sum of squares of one-way jitter values in SD direction (msec). The type is
    // interface{} with range: 0..18446744073709551615.
    OneWaySdSum2 interface{}

    // Minimum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMin interface{}

    // Maximum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsMax interface{}

    // Sum of one-way jitter values in DS direction (msec). The type is
    // interface{} with range: 0..4294967295.
    OneWayDsSum interface{}

    // Sum of squares of the OneWayMinDS and OneWayMaxDS values (msec). The type
    // is interface{} with range: 0..18446744073709551615.
    OneWayDsSum2 interface{}
}

func (udpJitterStats *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_Paths_Path_Hops_Hop_SpecificStats_UdpJitterStats) GetEntityData() *types.CommonEntityData {
    udpJitterStats.EntityData.YFilter = udpJitterStats.YFilter
    udpJitterStats.EntityData.YangName = "udp-jitter-stats"
    udpJitterStats.EntityData.BundleName = "cisco_ios_xr"
    udpJitterStats.EntityData.ParentYangName = "specific-stats"
    udpJitterStats.EntityData.SegmentPath = "udp-jitter-stats"
    udpJitterStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpJitterStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpJitterStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpJitterStats.EntityData.Children = make(map[string]types.YChild)
    udpJitterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    udpJitterStats.EntityData.Leafs["jitter-in"] = types.YLeaf{"JitterIn", udpJitterStats.JitterIn}
    udpJitterStats.EntityData.Leafs["jitter-out"] = types.YLeaf{"JitterOut", udpJitterStats.JitterOut}
    udpJitterStats.EntityData.Leafs["packet-loss-sd"] = types.YLeaf{"PacketLossSd", udpJitterStats.PacketLossSd}
    udpJitterStats.EntityData.Leafs["packet-loss-ds"] = types.YLeaf{"PacketLossDs", udpJitterStats.PacketLossDs}
    udpJitterStats.EntityData.Leafs["packet-out-of-sequence"] = types.YLeaf{"PacketOutOfSequence", udpJitterStats.PacketOutOfSequence}
    udpJitterStats.EntityData.Leafs["packet-mia"] = types.YLeaf{"PacketMia", udpJitterStats.PacketMia}
    udpJitterStats.EntityData.Leafs["packet-skipped"] = types.YLeaf{"PacketSkipped", udpJitterStats.PacketSkipped}
    udpJitterStats.EntityData.Leafs["packet-late-arrivals"] = types.YLeaf{"PacketLateArrivals", udpJitterStats.PacketLateArrivals}
    udpJitterStats.EntityData.Leafs["packet-invalid-tstamp"] = types.YLeaf{"PacketInvalidTstamp", udpJitterStats.PacketInvalidTstamp}
    udpJitterStats.EntityData.Leafs["internal-errors-count"] = types.YLeaf{"InternalErrorsCount", udpJitterStats.InternalErrorsCount}
    udpJitterStats.EntityData.Leafs["busies-count"] = types.YLeaf{"BusiesCount", udpJitterStats.BusiesCount}
    udpJitterStats.EntityData.Leafs["positive-sd-sum"] = types.YLeaf{"PositiveSdSum", udpJitterStats.PositiveSdSum}
    udpJitterStats.EntityData.Leafs["positive-sd-sum2"] = types.YLeaf{"PositiveSdSum2", udpJitterStats.PositiveSdSum2}
    udpJitterStats.EntityData.Leafs["positive-sd-min"] = types.YLeaf{"PositiveSdMin", udpJitterStats.PositiveSdMin}
    udpJitterStats.EntityData.Leafs["positive-sd-max"] = types.YLeaf{"PositiveSdMax", udpJitterStats.PositiveSdMax}
    udpJitterStats.EntityData.Leafs["positive-sd-count"] = types.YLeaf{"PositiveSdCount", udpJitterStats.PositiveSdCount}
    udpJitterStats.EntityData.Leafs["negative-sd-sum"] = types.YLeaf{"NegativeSdSum", udpJitterStats.NegativeSdSum}
    udpJitterStats.EntityData.Leafs["negative-sd-sum2"] = types.YLeaf{"NegativeSdSum2", udpJitterStats.NegativeSdSum2}
    udpJitterStats.EntityData.Leafs["negative-sd-min"] = types.YLeaf{"NegativeSdMin", udpJitterStats.NegativeSdMin}
    udpJitterStats.EntityData.Leafs["negative-sd-max"] = types.YLeaf{"NegativeSdMax", udpJitterStats.NegativeSdMax}
    udpJitterStats.EntityData.Leafs["negative-sd-count"] = types.YLeaf{"NegativeSdCount", udpJitterStats.NegativeSdCount}
    udpJitterStats.EntityData.Leafs["positive-ds-sum"] = types.YLeaf{"PositiveDsSum", udpJitterStats.PositiveDsSum}
    udpJitterStats.EntityData.Leafs["positive-ds-sum2"] = types.YLeaf{"PositiveDsSum2", udpJitterStats.PositiveDsSum2}
    udpJitterStats.EntityData.Leafs["positive-ds-min"] = types.YLeaf{"PositiveDsMin", udpJitterStats.PositiveDsMin}
    udpJitterStats.EntityData.Leafs["positive-ds-max"] = types.YLeaf{"PositiveDsMax", udpJitterStats.PositiveDsMax}
    udpJitterStats.EntityData.Leafs["positive-ds-count"] = types.YLeaf{"PositiveDsCount", udpJitterStats.PositiveDsCount}
    udpJitterStats.EntityData.Leafs["negative-ds-sum"] = types.YLeaf{"NegativeDsSum", udpJitterStats.NegativeDsSum}
    udpJitterStats.EntityData.Leafs["negative-ds-sum2"] = types.YLeaf{"NegativeDsSum2", udpJitterStats.NegativeDsSum2}
    udpJitterStats.EntityData.Leafs["negative-ds-min"] = types.YLeaf{"NegativeDsMin", udpJitterStats.NegativeDsMin}
    udpJitterStats.EntityData.Leafs["negative-ds-max"] = types.YLeaf{"NegativeDsMax", udpJitterStats.NegativeDsMax}
    udpJitterStats.EntityData.Leafs["negative-ds-count"] = types.YLeaf{"NegativeDsCount", udpJitterStats.NegativeDsCount}
    udpJitterStats.EntityData.Leafs["one-way-count"] = types.YLeaf{"OneWayCount", udpJitterStats.OneWayCount}
    udpJitterStats.EntityData.Leafs["one-way-sd-min"] = types.YLeaf{"OneWaySdMin", udpJitterStats.OneWaySdMin}
    udpJitterStats.EntityData.Leafs["one-way-sd-max"] = types.YLeaf{"OneWaySdMax", udpJitterStats.OneWaySdMax}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum"] = types.YLeaf{"OneWaySdSum", udpJitterStats.OneWaySdSum}
    udpJitterStats.EntityData.Leafs["one-way-sd-sum2"] = types.YLeaf{"OneWaySdSum2", udpJitterStats.OneWaySdSum2}
    udpJitterStats.EntityData.Leafs["one-way-ds-min"] = types.YLeaf{"OneWayDsMin", udpJitterStats.OneWayDsMin}
    udpJitterStats.EntityData.Leafs["one-way-ds-max"] = types.YLeaf{"OneWayDsMax", udpJitterStats.OneWayDsMax}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum"] = types.YLeaf{"OneWayDsSum", udpJitterStats.OneWayDsSum}
    udpJitterStats.EntityData.Leafs["one-way-ds-sum2"] = types.YLeaf{"OneWayDsSum2", udpJitterStats.OneWayDsSum2}
    return &(udpJitterStats.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths
// List of latest LPD paths
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Latest path statistics of MPLS LSP group operation. The type is slice of
    // Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths_LpdPath.
    LpdPath []Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths_LpdPath
}

func (lpdPaths *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths) GetEntityData() *types.CommonEntityData {
    lpdPaths.EntityData.YFilter = lpdPaths.YFilter
    lpdPaths.EntityData.YangName = "lpd-paths"
    lpdPaths.EntityData.BundleName = "cisco_ios_xr"
    lpdPaths.EntityData.ParentYangName = "non-distributed"
    lpdPaths.EntityData.SegmentPath = "lpd-paths"
    lpdPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpdPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpdPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpdPaths.EntityData.Children = make(map[string]types.YChild)
    lpdPaths.EntityData.Children["lpd-path"] = types.YChild{"LpdPath", nil}
    for i := range lpdPaths.LpdPath {
        lpdPaths.EntityData.Children[types.GetSegmentPath(&lpdPaths.LpdPath[i])] = types.YChild{"LpdPath", &lpdPaths.LpdPath[i]}
    }
    lpdPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lpdPaths.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths_LpdPath
// Latest path statistics of MPLS LSP
// group operation
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths_LpdPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPD path index. The type is interface{} with
    // range: -2147483648..2147483647.
    PathIndex interface{}

    // Path return code. The type is IpslaRetCode.
    ReturnCode interface{}

    // LPD path identifier.
    PathId Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths_LpdPath_PathId
}

func (lpdPath *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths_LpdPath) GetEntityData() *types.CommonEntityData {
    lpdPath.EntityData.YFilter = lpdPath.YFilter
    lpdPath.EntityData.YangName = "lpd-path"
    lpdPath.EntityData.BundleName = "cisco_ios_xr"
    lpdPath.EntityData.ParentYangName = "lpd-paths"
    lpdPath.EntityData.SegmentPath = "lpd-path" + "[path-index='" + fmt.Sprintf("%v", lpdPath.PathIndex) + "']"
    lpdPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpdPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpdPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpdPath.EntityData.Children = make(map[string]types.YChild)
    lpdPath.EntityData.Children["path-id"] = types.YChild{"PathId", &lpdPath.PathId}
    lpdPath.EntityData.Leafs = make(map[string]types.YLeaf)
    lpdPath.EntityData.Leafs["path-index"] = types.YLeaf{"PathIndex", lpdPath.PathIndex}
    lpdPath.EntityData.Leafs["return-code"] = types.YLeaf{"ReturnCode", lpdPath.ReturnCode}
    return &(lpdPath.EntityData)
}

// Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths_LpdPath_PathId
// LPD path identifier
type Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths_LpdPath_PathId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP selector. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LspSelector interface{}

    // Output interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    OutputInterface interface{}

    // Nexthop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}

    // Downstream label stacks. The type is slice of interface{} with range:
    // 0..4294967295.
    DownstreamLabel []interface{}
}

func (pathId *Ipsla_OperationData_Operations_Operation_Statistics_Aggregated_Hours_Hour_NonDistributed_LpdPaths_LpdPath_PathId) GetEntityData() *types.CommonEntityData {
    pathId.EntityData.YFilter = pathId.YFilter
    pathId.EntityData.YangName = "path-id"
    pathId.EntityData.BundleName = "cisco_ios_xr"
    pathId.EntityData.ParentYangName = "lpd-path"
    pathId.EntityData.SegmentPath = "path-id"
    pathId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathId.EntityData.Children = make(map[string]types.YChild)
    pathId.EntityData.Leafs = make(map[string]types.YLeaf)
    pathId.EntityData.Leafs["lsp-selector"] = types.YLeaf{"LspSelector", pathId.LspSelector}
    pathId.EntityData.Leafs["output-interface"] = types.YLeaf{"OutputInterface", pathId.OutputInterface}
    pathId.EntityData.Leafs["nexthop-address"] = types.YLeaf{"NexthopAddress", pathId.NexthopAddress}
    pathId.EntityData.Leafs["downstream-label"] = types.YLeaf{"DownstreamLabel", pathId.DownstreamLabel}
    return &(pathId.EntityData)
}

// Ipsla_ApplicationInfo
// IPSLA application information
type Ipsla_ApplicationInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version of the IPSLA in Version.Release .Patch-level format. The type is
    // string.
    Version interface{}

    // Maximum number of entries. The type is interface{} with range:
    // 0..4294967295.
    MaxEntries interface{}

    // Number of entries configured. The type is interface{} with range:
    // 0..4294967295.
    EntriesConfigured interface{}

    // Number of active entries. The type is interface{} with range:
    // 0..4294967295.
    ActiveEntries interface{}

    // Number of pending entries. The type is interface{} with range:
    // 0..4294967295.
    PendingEntries interface{}

    // Number of inactive entries. The type is interface{} with range:
    // 0..4294967295.
    InactiveEntries interface{}

    // Number of configurable probes. The type is interface{} with range:
    // 0..4294967295.
    ConfigurableProbes interface{}

    // IPSLA low memory watermark in KB. The type is interface{} with range:
    // 0..4294967295.
    MinMemory interface{}

    // IPSLA HW timestamp Disabled flag. The type is bool.
    HwTimestampDisabled interface{}

    // Operation types available in this IPSLA version. The type is slice of
    // SlaOpTypes.
    OperationType []interface{}
}

func (applicationInfo *Ipsla_ApplicationInfo) GetEntityData() *types.CommonEntityData {
    applicationInfo.EntityData.YFilter = applicationInfo.YFilter
    applicationInfo.EntityData.YangName = "application-info"
    applicationInfo.EntityData.BundleName = "cisco_ios_xr"
    applicationInfo.EntityData.ParentYangName = "ipsla"
    applicationInfo.EntityData.SegmentPath = "application-info"
    applicationInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    applicationInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    applicationInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    applicationInfo.EntityData.Children = make(map[string]types.YChild)
    applicationInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    applicationInfo.EntityData.Leafs["version"] = types.YLeaf{"Version", applicationInfo.Version}
    applicationInfo.EntityData.Leafs["max-entries"] = types.YLeaf{"MaxEntries", applicationInfo.MaxEntries}
    applicationInfo.EntityData.Leafs["entries-configured"] = types.YLeaf{"EntriesConfigured", applicationInfo.EntriesConfigured}
    applicationInfo.EntityData.Leafs["active-entries"] = types.YLeaf{"ActiveEntries", applicationInfo.ActiveEntries}
    applicationInfo.EntityData.Leafs["pending-entries"] = types.YLeaf{"PendingEntries", applicationInfo.PendingEntries}
    applicationInfo.EntityData.Leafs["inactive-entries"] = types.YLeaf{"InactiveEntries", applicationInfo.InactiveEntries}
    applicationInfo.EntityData.Leafs["configurable-probes"] = types.YLeaf{"ConfigurableProbes", applicationInfo.ConfigurableProbes}
    applicationInfo.EntityData.Leafs["min-memory"] = types.YLeaf{"MinMemory", applicationInfo.MinMemory}
    applicationInfo.EntityData.Leafs["hw-timestamp-disabled"] = types.YLeaf{"HwTimestampDisabled", applicationInfo.HwTimestampDisabled}
    applicationInfo.EntityData.Leafs["operation-type"] = types.YLeaf{"OperationType", applicationInfo.OperationType}
    return &(applicationInfo.EntityData)
}

