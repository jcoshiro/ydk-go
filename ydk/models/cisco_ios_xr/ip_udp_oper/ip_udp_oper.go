// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-udp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   udp: IP UDP Operational Data
//   udp-connection: udp connection
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_udp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_udp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-udp-oper udp}", reflect.TypeOf(Udp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-udp-oper:udp", reflect.TypeOf(Udp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-udp-oper udp-connection}", reflect.TypeOf(UdpConnection{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-udp-oper:udp-connection", reflect.TypeOf(UdpConnection{}))
}

// LptsPcbQuery represents Lpts pcb query
type LptsPcbQuery string

const (
    // No filter
    LptsPcbQuery_all LptsPcbQuery = "all"

    // Static policy filter
    LptsPcbQuery_static_policy LptsPcbQuery = "static-policy"

    // Interface filter
    LptsPcbQuery_interface_ LptsPcbQuery = "interface"

    // Packet type filter
    LptsPcbQuery_packet LptsPcbQuery = "packet"
)

// MessageTypeIgmp represents LPTS IGMP message types
type MessageTypeIgmp string

const (
    // IGMP Packet type: Membership query
    MessageTypeIgmp_membership_query MessageTypeIgmp = "membership-query"

    // IGMP Packet type: V1 membership report
    MessageTypeIgmp_v1_membership_report MessageTypeIgmp = "v1-membership-report"

    // IGMP Packet type: DVMRP
    MessageTypeIgmp_dvmrp MessageTypeIgmp = "dvmrp"

    // IGMP Packet type: PIM version 1
    MessageTypeIgmp_pi_mv1 MessageTypeIgmp = "pi-mv1"

    // IGMP Packet type: Cisco Trace Messages
    MessageTypeIgmp_cisco_trace_messages MessageTypeIgmp = "cisco-trace-messages"

    // IGMP Packet type: V2 membership report
    MessageTypeIgmp_v2_membership_report MessageTypeIgmp = "v2-membership-report"

    // IGMP Packet type: V2 leave group
    MessageTypeIgmp_v2_leave_group MessageTypeIgmp = "v2-leave-group"

    // IGMP Packet type: Multicast traceroute response
    MessageTypeIgmp_multicast_traceroute_response MessageTypeIgmp = "multicast-traceroute-response"

    // IGMP Packet type: MulticastTraceroute
    MessageTypeIgmp_multicast_traceroute MessageTypeIgmp = "multicast-traceroute"

    // IGMP Packet type: V3 membership report
    MessageTypeIgmp_v3_membership_report MessageTypeIgmp = "v3-membership-report"

    // IGMP Packet type: Multicast router
    // advertisement
    MessageTypeIgmp_multicast_router_advertisement MessageTypeIgmp = "multicast-router-advertisement"

    // IGMP Packet type: Multicast router solicitation
    MessageTypeIgmp_multicast_router_solicitation MessageTypeIgmp = "multicast-router-solicitation"

    // IGMP Packet type: Multicast router termination
    MessageTypeIgmp_multicast_router_termination MessageTypeIgmp = "multicast-router-termination"
)

// MessageTypeIgmp_ represents LPTS IGMP message types
type MessageTypeIgmp_ string

const (
    // IGMP Packet type: Membership query
    MessageTypeIgmp__membership_query MessageTypeIgmp_ = "membership-query"

    // IGMP Packet type: V1 membership report
    MessageTypeIgmp__v1_membership_report MessageTypeIgmp_ = "v1-membership-report"

    // IGMP Packet type: DVMRP
    MessageTypeIgmp__dvmrp MessageTypeIgmp_ = "dvmrp"

    // IGMP Packet type: PIM version 1
    MessageTypeIgmp__pi_mv1 MessageTypeIgmp_ = "pi-mv1"

    // IGMP Packet type: Cisco Trace Messages
    MessageTypeIgmp__cisco_trace_messages MessageTypeIgmp_ = "cisco-trace-messages"

    // IGMP Packet type: V2 membership report
    MessageTypeIgmp__v2_membership_report MessageTypeIgmp_ = "v2-membership-report"

    // IGMP Packet type: V2 leave group
    MessageTypeIgmp__v2_leave_group MessageTypeIgmp_ = "v2-leave-group"

    // IGMP Packet type: Multicast traceroute response
    MessageTypeIgmp__multicast_traceroute_response MessageTypeIgmp_ = "multicast-traceroute-response"

    // IGMP Packet type: MulticastTraceroute
    MessageTypeIgmp__multicast_traceroute MessageTypeIgmp_ = "multicast-traceroute"

    // IGMP Packet type: V3 membership report
    MessageTypeIgmp__v3_membership_report MessageTypeIgmp_ = "v3-membership-report"

    // IGMP Packet type: Multicast router
    // advertisement
    MessageTypeIgmp__multicast_router_advertisement MessageTypeIgmp_ = "multicast-router-advertisement"

    // IGMP Packet type: Multicast router solicitation
    MessageTypeIgmp__multicast_router_solicitation MessageTypeIgmp_ = "multicast-router-solicitation"

    // IGMP Packet type: Multicast router termination
    MessageTypeIgmp__multicast_router_termination MessageTypeIgmp_ = "multicast-router-termination"
)

// MessageTypeIcmpv6 represents LPTS ICMPv6 message types
type MessageTypeIcmpv6 string

const (
    // ICMPv6 Packet type: Destination unreachable
    MessageTypeIcmpv6_destination_unreachable MessageTypeIcmpv6 = "destination-unreachable"

    // ICMPv6 Packet type: packet too big
    MessageTypeIcmpv6_packet_too_big MessageTypeIcmpv6 = "packet-too-big"

    // ICMPv6 Packet type: Time exceeded
    MessageTypeIcmpv6_time_exceeded MessageTypeIcmpv6 = "time-exceeded"

    // ICMPv6 Packet type: Parameter problem
    MessageTypeIcmpv6_parameter_problem MessageTypeIcmpv6 = "parameter-problem"

    // ICMPv6 Packet type: Echo request
    MessageTypeIcmpv6_echo_request MessageTypeIcmpv6 = "echo-request"

    // ICMPv6 Packet type: Echo reply
    MessageTypeIcmpv6_echo_reply MessageTypeIcmpv6 = "echo-reply"

    // ICMPv6 Packet type: Multicast listener query
    MessageTypeIcmpv6_multicast_listener_query MessageTypeIcmpv6 = "multicast-listener-query"

    // ICMPv6 Packet type: Multicast listener report
    MessageTypeIcmpv6_multicast_listener_report MessageTypeIcmpv6 = "multicast-listener-report"

    // ICMPv6 Packet type: Multicast listener done
    MessageTypeIcmpv6_multicast_listener_done MessageTypeIcmpv6 = "multicast-listener-done"

    // ICMPv6 Packet type: Router solicitation
    MessageTypeIcmpv6_router_solicitation MessageTypeIcmpv6 = "router-solicitation"

    // ICMPv6 Packet type: Router advertisement
    MessageTypeIcmpv6_router_advertisement MessageTypeIcmpv6 = "router-advertisement"

    // ICMPv6 Packet type: Neighbor solicitation
    MessageTypeIcmpv6_neighbor_solicitation MessageTypeIcmpv6 = "neighbor-solicitation"

    // ICMPv6 Packet type: Neighbor advertisement
    MessageTypeIcmpv6_neighbor_advertisement MessageTypeIcmpv6 = "neighbor-advertisement"

    // ICMPv6 Packet type: Redirect message
    MessageTypeIcmpv6_redirect_message MessageTypeIcmpv6 = "redirect-message"

    // ICMPv6 Packet type: Router renumbering
    MessageTypeIcmpv6_router_renumbering MessageTypeIcmpv6 = "router-renumbering"

    // ICMPv6 Packet type: Node information query
    MessageTypeIcmpv6_node_information_query MessageTypeIcmpv6 = "node-information-query"

    // ICMPv6 Packet type: Node information reply
    MessageTypeIcmpv6_node_information_reply MessageTypeIcmpv6 = "node-information-reply"

    // ICMPv6 Packet type: Inverse neighbor discovery
    // solicitation message
    MessageTypeIcmpv6_inverse_neighbor_discovery_solicitaion MessageTypeIcmpv6 = "inverse-neighbor-discovery-solicitaion"

    // ICMPv6 Packet type: Inverse neighbor discovery
    // advertisement message
    MessageTypeIcmpv6_inverse_neighbor_discover_advertisement MessageTypeIcmpv6 = "inverse-neighbor-discover-advertisement"

    // ICMPv6 Packet type: Version 2 multicast
    // listener report
    MessageTypeIcmpv6_v2_multicast_listener_report MessageTypeIcmpv6 = "v2-multicast-listener-report"

    // ICMPv6 Packet type: Home agent address
    // discovery request message
    MessageTypeIcmpv6_home_agent_address_discovery_request MessageTypeIcmpv6 = "home-agent-address-discovery-request"

    // ICMPv6 Packet type: Home agent address
    // discovery reply message
    MessageTypeIcmpv6_home_agent_address_discovery_reply MessageTypeIcmpv6 = "home-agent-address-discovery-reply"

    // ICMPv6 Packet type: Mobile prefix solicitation
    MessageTypeIcmpv6_mobile_prefix_solicitation MessageTypeIcmpv6 = "mobile-prefix-solicitation"

    // ICMPv6 Packet type: Mobile prefix advertisement
    MessageTypeIcmpv6_mobile_prefix_advertisement MessageTypeIcmpv6 = "mobile-prefix-advertisement"

    // ICMPv6 Packet type: Certification path
    // solicitation message
    MessageTypeIcmpv6_certification_path_solicitation_message MessageTypeIcmpv6 = "certification-path-solicitation-message"

    // ICMPv6 Packet type: Certification path
    // advertisement message
    MessageTypeIcmpv6_certification_path_advertisement_message MessageTypeIcmpv6 = "certification-path-advertisement-message"

    // ICMPv6 Packet type: ICMP messages utilized by
    // experimental mobility  protocols such as
    // seamoby
    MessageTypeIcmpv6_experimental_mobility_protocols MessageTypeIcmpv6 = "experimental-mobility-protocols"

    // ICMPv6 Packet type: Multicast router
    // advertisement
    MessageTypeIcmpv6_multicast_router_advertisement MessageTypeIcmpv6 = "multicast-router-advertisement"

    // ICMPv6 Packet type: Multicast router
    // solicitation
    MessageTypeIcmpv6_multicast_router_solicitation MessageTypeIcmpv6 = "multicast-router-solicitation"

    // ICMPv6 Packet type: Multicast router
    // termination
    MessageTypeIcmpv6_multicast_router_termination MessageTypeIcmpv6 = "multicast-router-termination"

    // ICMPv6 Packet type: FMIPv6 messages
    MessageTypeIcmpv6_fmipv6_messages MessageTypeIcmpv6 = "fmipv6-messages"
)

// MessageTypeIcmpv6_ represents LPTS ICMPv6 message types
type MessageTypeIcmpv6_ string

const (
    // ICMPv6 Packet type: Destination unreachable
    MessageTypeIcmpv6__destination_unreachable MessageTypeIcmpv6_ = "destination-unreachable"

    // ICMPv6 Packet type: packet too big
    MessageTypeIcmpv6__packet_too_big MessageTypeIcmpv6_ = "packet-too-big"

    // ICMPv6 Packet type: Time exceeded
    MessageTypeIcmpv6__time_exceeded MessageTypeIcmpv6_ = "time-exceeded"

    // ICMPv6 Packet type: Parameter problem
    MessageTypeIcmpv6__parameter_problem MessageTypeIcmpv6_ = "parameter-problem"

    // ICMPv6 Packet type: Echo request
    MessageTypeIcmpv6__echo_request MessageTypeIcmpv6_ = "echo-request"

    // ICMPv6 Packet type: Echo reply
    MessageTypeIcmpv6__echo_reply MessageTypeIcmpv6_ = "echo-reply"

    // ICMPv6 Packet type: Multicast listener query
    MessageTypeIcmpv6__multicast_listener_query MessageTypeIcmpv6_ = "multicast-listener-query"

    // ICMPv6 Packet type: Multicast listener report
    MessageTypeIcmpv6__multicast_listener_report MessageTypeIcmpv6_ = "multicast-listener-report"

    // ICMPv6 Packet type: Multicast listener done
    MessageTypeIcmpv6__multicast_listener_done MessageTypeIcmpv6_ = "multicast-listener-done"

    // ICMPv6 Packet type: Router solicitation
    MessageTypeIcmpv6__router_solicitation MessageTypeIcmpv6_ = "router-solicitation"

    // ICMPv6 Packet type: Router advertisement
    MessageTypeIcmpv6__router_advertisement MessageTypeIcmpv6_ = "router-advertisement"

    // ICMPv6 Packet type: Neighbor solicitation
    MessageTypeIcmpv6__neighbor_solicitation MessageTypeIcmpv6_ = "neighbor-solicitation"

    // ICMPv6 Packet type: Neighbor advertisement
    MessageTypeIcmpv6__neighbor_advertisement MessageTypeIcmpv6_ = "neighbor-advertisement"

    // ICMPv6 Packet type: Redirect message
    MessageTypeIcmpv6__redirect_message MessageTypeIcmpv6_ = "redirect-message"

    // ICMPv6 Packet type: Router renumbering
    MessageTypeIcmpv6__router_renumbering MessageTypeIcmpv6_ = "router-renumbering"

    // ICMPv6 Packet type: Node information query
    MessageTypeIcmpv6__node_information_query MessageTypeIcmpv6_ = "node-information-query"

    // ICMPv6 Packet type: Node information reply
    MessageTypeIcmpv6__node_information_reply MessageTypeIcmpv6_ = "node-information-reply"

    // ICMPv6 Packet type: Inverse neighbor discovery
    // solicitation message
    MessageTypeIcmpv6__inverse_neighbor_discovery_solicitaion MessageTypeIcmpv6_ = "inverse-neighbor-discovery-solicitaion"

    // ICMPv6 Packet type: Inverse neighbor discovery
    // advertisement message
    MessageTypeIcmpv6__inverse_neighbor_discover_advertisement MessageTypeIcmpv6_ = "inverse-neighbor-discover-advertisement"

    // ICMPv6 Packet type: Version 2 multicast
    // listener report
    MessageTypeIcmpv6__v2_multicast_listener_report MessageTypeIcmpv6_ = "v2-multicast-listener-report"

    // ICMPv6 Packet type: Home agent address
    // discovery request message
    MessageTypeIcmpv6__home_agent_address_discovery_request MessageTypeIcmpv6_ = "home-agent-address-discovery-request"

    // ICMPv6 Packet type: Home agent address
    // discovery reply message
    MessageTypeIcmpv6__home_agent_address_discovery_reply MessageTypeIcmpv6_ = "home-agent-address-discovery-reply"

    // ICMPv6 Packet type: Mobile prefix solicitation
    MessageTypeIcmpv6__mobile_prefix_solicitation MessageTypeIcmpv6_ = "mobile-prefix-solicitation"

    // ICMPv6 Packet type: Mobile prefix advertisement
    MessageTypeIcmpv6__mobile_prefix_advertisement MessageTypeIcmpv6_ = "mobile-prefix-advertisement"

    // ICMPv6 Packet type: Certification path
    // solicitation message
    MessageTypeIcmpv6__certification_path_solicitation_message MessageTypeIcmpv6_ = "certification-path-solicitation-message"

    // ICMPv6 Packet type: Certification path
    // advertisement message
    MessageTypeIcmpv6__certification_path_advertisement_message MessageTypeIcmpv6_ = "certification-path-advertisement-message"

    // ICMPv6 Packet type: ICMP messages utilized by
    // experimental mobility  protocols such as
    // seamoby
    MessageTypeIcmpv6__experimental_mobility_protocols MessageTypeIcmpv6_ = "experimental-mobility-protocols"

    // ICMPv6 Packet type: Multicast router
    // advertisement
    MessageTypeIcmpv6__multicast_router_advertisement MessageTypeIcmpv6_ = "multicast-router-advertisement"

    // ICMPv6 Packet type: Multicast router
    // solicitation
    MessageTypeIcmpv6__multicast_router_solicitation MessageTypeIcmpv6_ = "multicast-router-solicitation"

    // ICMPv6 Packet type: Multicast router
    // termination
    MessageTypeIcmpv6__multicast_router_termination MessageTypeIcmpv6_ = "multicast-router-termination"

    // ICMPv6 Packet type: FMIPv6 messages
    MessageTypeIcmpv6__fmipv6_messages MessageTypeIcmpv6_ = "fmipv6-messages"
)

// MessageTypeIcmp represents LPTS ICMP message types
type MessageTypeIcmp string

const (
    // ICMP Packet type: Echo reply
    MessageTypeIcmp_echo_reply MessageTypeIcmp = "echo-reply"

    // ICMP Packet type: Destination unreachable
    MessageTypeIcmp_destination_unreachable MessageTypeIcmp = "destination-unreachable"

    // ICMP Packet type: Source quench
    MessageTypeIcmp_source_quench MessageTypeIcmp = "source-quench"

    // ICMP Packet type: Redirect
    MessageTypeIcmp_redirect MessageTypeIcmp = "redirect"

    // ICMP Packet type: Alternate host address
    MessageTypeIcmp_alternate_host_address MessageTypeIcmp = "alternate-host-address"

    // ICMP Packet type: Echo
    MessageTypeIcmp_echo MessageTypeIcmp = "echo"

    // ICMP Packet type: Router advertisement
    MessageTypeIcmp_router_advertisement MessageTypeIcmp = "router-advertisement"

    // ICMP Packet type: Router selection
    MessageTypeIcmp_router_selection MessageTypeIcmp = "router-selection"

    // ICMP Packet type: Time exceeded
    MessageTypeIcmp_time_exceeded MessageTypeIcmp = "time-exceeded"

    // ICMP Packet type: Parameter problem
    MessageTypeIcmp_parameter_problem MessageTypeIcmp = "parameter-problem"

    // ICMP Packet type: Time stamp
    MessageTypeIcmp_time_stamp MessageTypeIcmp = "time-stamp"

    // ICMP Packet type: Time stamp reply
    MessageTypeIcmp_time_stamp_reply MessageTypeIcmp = "time-stamp-reply"

    // ICMP Packet type: Information request
    MessageTypeIcmp_information_request MessageTypeIcmp = "information-request"

    // ICMP Packet type: Information reply
    MessageTypeIcmp_information_reply MessageTypeIcmp = "information-reply"

    // ICMP Packet type: Address mask request
    MessageTypeIcmp_address_mask_request MessageTypeIcmp = "address-mask-request"

    // ICMP Packet type: Address mask reply
    MessageTypeIcmp_address_mask_reply MessageTypeIcmp = "address-mask-reply"

    // ICMP Packet type: Trace route
    MessageTypeIcmp_trace_route MessageTypeIcmp = "trace-route"

    // ICMP Packet type: Datagram Conversion error
    MessageTypeIcmp_datagram_conversion_error MessageTypeIcmp = "datagram-conversion-error"

    // ICMP Packet type: Mobile host redirect
    MessageTypeIcmp_mobile_host_redirect MessageTypeIcmp = "mobile-host-redirect"

    // ICMP Packet type: IPv6 where-are-you
    MessageTypeIcmp_where_are_you MessageTypeIcmp = "where-are-you"

    // ICMP Packet type: IPv6 i-am-here
    MessageTypeIcmp_iam_here MessageTypeIcmp = "iam-here"

    // ICMP Packet type: Mobile registration request
    MessageTypeIcmp_mobile_registration_request MessageTypeIcmp = "mobile-registration-request"

    // ICMP Packet type: Mobile registration reply
    MessageTypeIcmp_mobile_registration_reply MessageTypeIcmp = "mobile-registration-reply"

    // ICMP Packet type: Domain name request
    MessageTypeIcmp_domain_name_request MessageTypeIcmp = "domain-name-request"
)

// MessageTypeIcmp_ represents LPTS ICMP message types
type MessageTypeIcmp_ string

const (
    // ICMP Packet type: Echo reply
    MessageTypeIcmp__echo_reply MessageTypeIcmp_ = "echo-reply"

    // ICMP Packet type: Destination unreachable
    MessageTypeIcmp__destination_unreachable MessageTypeIcmp_ = "destination-unreachable"

    // ICMP Packet type: Source quench
    MessageTypeIcmp__source_quench MessageTypeIcmp_ = "source-quench"

    // ICMP Packet type: Redirect
    MessageTypeIcmp__redirect MessageTypeIcmp_ = "redirect"

    // ICMP Packet type: Alternate host address
    MessageTypeIcmp__alternate_host_address MessageTypeIcmp_ = "alternate-host-address"

    // ICMP Packet type: Echo
    MessageTypeIcmp__echo MessageTypeIcmp_ = "echo"

    // ICMP Packet type: Router advertisement
    MessageTypeIcmp__router_advertisement MessageTypeIcmp_ = "router-advertisement"

    // ICMP Packet type: Router selection
    MessageTypeIcmp__router_selection MessageTypeIcmp_ = "router-selection"

    // ICMP Packet type: Time exceeded
    MessageTypeIcmp__time_exceeded MessageTypeIcmp_ = "time-exceeded"

    // ICMP Packet type: Parameter problem
    MessageTypeIcmp__parameter_problem MessageTypeIcmp_ = "parameter-problem"

    // ICMP Packet type: Time stamp
    MessageTypeIcmp__time_stamp MessageTypeIcmp_ = "time-stamp"

    // ICMP Packet type: Time stamp reply
    MessageTypeIcmp__time_stamp_reply MessageTypeIcmp_ = "time-stamp-reply"

    // ICMP Packet type: Information request
    MessageTypeIcmp__information_request MessageTypeIcmp_ = "information-request"

    // ICMP Packet type: Information reply
    MessageTypeIcmp__information_reply MessageTypeIcmp_ = "information-reply"

    // ICMP Packet type: Address mask request
    MessageTypeIcmp__address_mask_request MessageTypeIcmp_ = "address-mask-request"

    // ICMP Packet type: Address mask reply
    MessageTypeIcmp__address_mask_reply MessageTypeIcmp_ = "address-mask-reply"

    // ICMP Packet type: Trace route
    MessageTypeIcmp__trace_route MessageTypeIcmp_ = "trace-route"

    // ICMP Packet type: Datagram Conversion error
    MessageTypeIcmp__datagram_conversion_error MessageTypeIcmp_ = "datagram-conversion-error"

    // ICMP Packet type: Mobile host redirect
    MessageTypeIcmp__mobile_host_redirect MessageTypeIcmp_ = "mobile-host-redirect"

    // ICMP Packet type: IPv6 where-are-you
    MessageTypeIcmp__where_are_you MessageTypeIcmp_ = "where-are-you"

    // ICMP Packet type: IPv6 i-am-here
    MessageTypeIcmp__iam_here MessageTypeIcmp_ = "iam-here"

    // ICMP Packet type: Mobile registration request
    MessageTypeIcmp__mobile_registration_request MessageTypeIcmp_ = "mobile-registration-request"

    // ICMP Packet type: Mobile registration reply
    MessageTypeIcmp__mobile_registration_reply MessageTypeIcmp_ = "mobile-registration-reply"

    // ICMP Packet type: Domain name request
    MessageTypeIcmp__domain_name_request MessageTypeIcmp_ = "domain-name-request"
)

// Packet represents Packet type
type Packet string

const (
    // ICMP packet type
    Packet_icmp Packet = "icmp"

    // ICMPv6 packet type
    Packet_icm_pv6 Packet = "icm-pv6"

    // IGMP packet type
    Packet_igmp Packet = "igmp"

    // Packet type unknown
    Packet_unknown Packet = "unknown"
)

// AddrFamily represents Address Family Types
type AddrFamily string

const (
    // Unspecified
    AddrFamily_unspecified AddrFamily = "unspecified"

    // Local to host (pipes, portals)
    AddrFamily_local AddrFamily = "local"

    // Internetwork: UDP, TCP, etc.
    AddrFamily_inet AddrFamily = "inet"

    // arpanet imp addresses
    AddrFamily_implink AddrFamily = "implink"

    // Pup protocols: e.g. BSP
    AddrFamily_pup AddrFamily = "pup"

    // mit CHAOS protocols
    AddrFamily_chaos AddrFamily = "chaos"

    // XEROX NS protocols
    AddrFamily_ns AddrFamily = "ns"

    // ISO protocols
    AddrFamily_iso AddrFamily = "iso"

    // European computer manufacturers
    AddrFamily_ecma AddrFamily = "ecma"

    // Datakit protocols
    AddrFamily_data_kit AddrFamily = "data-kit"

    // CCITT protocols, X.25 etc
    AddrFamily_ccitt AddrFamily = "ccitt"

    // IBM SNA
    AddrFamily_sna AddrFamily = "sna"

    // DECnet
    AddrFamily_de_cnet AddrFamily = "de-cnet"

    // DEC Direct data link interface
    AddrFamily_dli AddrFamily = "dli"

    // LAT
    AddrFamily_lat AddrFamily = "lat"

    // NSC Hyperchannel
    AddrFamily_hylink AddrFamily = "hylink"

    // Apple Talk
    AddrFamily_appletalk AddrFamily = "appletalk"

    // Internal Routing Protocol
    AddrFamily_route AddrFamily = "route"

    // Link layer interface
    AddrFamily_link AddrFamily = "link"

    // eXpress Transfer Protocol (no AF)
    AddrFamily_pseudo_xtp AddrFamily = "pseudo-xtp"

    // Connection-oriented IP, aka ST II
    AddrFamily_coip AddrFamily = "coip"

    // Computer Network Technology
    AddrFamily_cnt AddrFamily = "cnt"

    // Help Identify RTIP packets
    AddrFamily_pseudo_rtip AddrFamily = "pseudo-rtip"

    // Novell Internet Protocol
    AddrFamily_ipx AddrFamily = "ipx"

    // Simple Internet Protocol
    AddrFamily_sip AddrFamily = "sip"

    // Help Identify PIP packets
    AddrFamily_pseudo_pip AddrFamily = "pseudo-pip"

    // IP version 6
    AddrFamily_inet6 AddrFamily = "inet6"

    // 802.2 SNAP sockets
    AddrFamily_snap AddrFamily = "snap"

    // SAP_CLNS + nlpid encaps
    AddrFamily_clnl AddrFamily = "clnl"

    // cisco HDLC on serial
    AddrFamily_chdlc AddrFamily = "chdlc"

    // PPP sockets
    AddrFamily_ppp AddrFamily = "ppp"

    // Host-based CAS signaling
    AddrFamily_host_cas AddrFamily = "host-cas"

    // DSP messaging
    AddrFamily_dsp AddrFamily = "dsp"

    // SAP Sockets
    AddrFamily_sap AddrFamily = "sap"

    // ATM Sockets
    AddrFamily_atm AddrFamily = "atm"

    // Frame Relay sockets
    AddrFamily_fr AddrFamily = "fr"

    // Voice Media Stream Sockets
    AddrFamily_mso AddrFamily = "mso"

    // ISDN D Channel Sockets
    AddrFamily_dchan AddrFamily = "dchan"

    // Trunk Framer media IF Sockets
    AddrFamily_cas AddrFamily = "cas"

    // Network Address Translation Sockets
    AddrFamily_nat AddrFamily = "nat"

    // Generic Ethernet Sockets
    AddrFamily_ether AddrFamily = "ether"

    // Spatial Reuse Protocol Sockets
    AddrFamily_srp AddrFamily = "srp"
)

// UdpAddressFamily represents Address Family Type
type UdpAddressFamily string

const (
    // IPv4
    UdpAddressFamily_ipv4 UdpAddressFamily = "ipv4"

    // IPv6
    UdpAddressFamily_ipv6 UdpAddressFamily = "ipv6"
)

// Udp
// IP UDP Operational Data
type Udp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific UDP operational data.
    Nodes Udp_Nodes
}

func (udp *Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xr"
    udp.EntityData.ParentYangName = "Cisco-IOS-XR-ip-udp-oper"
    udp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-udp-oper:udp"
    udp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udp.EntityData.Children = make(map[string]types.YChild)
    udp.EntityData.Children["nodes"] = types.YChild{"Nodes", &udp.Nodes}
    udp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(udp.EntityData)
}

// Udp_Nodes
// Node-specific UDP operational data
type Udp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UDP operational data for a particular node. The type is slice of
    // Udp_Nodes_Node.
    Node []Udp_Nodes_Node
}

func (nodes *Udp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "udp"
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

// Udp_Nodes_Node
// UDP operational data for a particular node
type Udp_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Statistical UDP operational data for a node.
    Statistics Udp_Nodes_Node_Statistics
}

func (node *Udp_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["statistics"] = types.YChild{"Statistics", &node.Statistics}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Udp_Nodes_Node_Statistics
// Statistical UDP operational data for a node
type Udp_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UDP Traffic statistics for IPv4.
    Ipv4Traffic Udp_Nodes_Node_Statistics_Ipv4Traffic

    // UDP Traffic statistics for IPv6.
    Ipv6Traffic Udp_Nodes_Node_Statistics_Ipv6Traffic
}

func (statistics *Udp_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["ipv4-traffic"] = types.YChild{"Ipv4Traffic", &statistics.Ipv4Traffic}
    statistics.EntityData.Children["ipv6-traffic"] = types.YChild{"Ipv6Traffic", &statistics.Ipv6Traffic}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Udp_Nodes_Node_Statistics_Ipv4Traffic
// UDP Traffic statistics for IPv4
type Udp_Nodes_Node_Statistics_Ipv4Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UDP Received. The type is interface{} with range: 0..4294967295.
    UdpInputPackets interface{}

    // UDP Checksum Errors. The type is interface{} with range: 0..4294967295.
    UdpChecksumErrorPackets interface{}

    // UDP No Port. The type is interface{} with range: 0..4294967295.
    UdpNoPortPackets interface{}

    // UDP bad length. The type is interface{} with range: 0..4294967295.
    UdpBadLengthPackets interface{}

    // UDP Transmitted. The type is interface{} with range: 0..4294967295.
    UdpOutputPackets interface{}

    // UDP drop for other reason. The type is interface{} with range:
    // 0..4294967295.
    UdpDroppedPackets interface{}
}

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetEntityData() *types.CommonEntityData {
    ipv4Traffic.EntityData.YFilter = ipv4Traffic.YFilter
    ipv4Traffic.EntityData.YangName = "ipv4-traffic"
    ipv4Traffic.EntityData.BundleName = "cisco_ios_xr"
    ipv4Traffic.EntityData.ParentYangName = "statistics"
    ipv4Traffic.EntityData.SegmentPath = "ipv4-traffic"
    ipv4Traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Traffic.EntityData.Children = make(map[string]types.YChild)
    ipv4Traffic.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Traffic.EntityData.Leafs["udp-input-packets"] = types.YLeaf{"UdpInputPackets", ipv4Traffic.UdpInputPackets}
    ipv4Traffic.EntityData.Leafs["udp-checksum-error-packets"] = types.YLeaf{"UdpChecksumErrorPackets", ipv4Traffic.UdpChecksumErrorPackets}
    ipv4Traffic.EntityData.Leafs["udp-no-port-packets"] = types.YLeaf{"UdpNoPortPackets", ipv4Traffic.UdpNoPortPackets}
    ipv4Traffic.EntityData.Leafs["udp-bad-length-packets"] = types.YLeaf{"UdpBadLengthPackets", ipv4Traffic.UdpBadLengthPackets}
    ipv4Traffic.EntityData.Leafs["udp-output-packets"] = types.YLeaf{"UdpOutputPackets", ipv4Traffic.UdpOutputPackets}
    ipv4Traffic.EntityData.Leafs["udp-dropped-packets"] = types.YLeaf{"UdpDroppedPackets", ipv4Traffic.UdpDroppedPackets}
    return &(ipv4Traffic.EntityData)
}

// Udp_Nodes_Node_Statistics_Ipv6Traffic
// UDP Traffic statistics for IPv6
type Udp_Nodes_Node_Statistics_Ipv6Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UDP Received. The type is interface{} with range: 0..4294967295.
    UdpInputPackets interface{}

    // UDP Checksum Errors. The type is interface{} with range: 0..4294967295.
    UdpChecksumErrorPackets interface{}

    // UDP No Port. The type is interface{} with range: 0..4294967295.
    UdpNoPortPackets interface{}

    // UDP bad length. The type is interface{} with range: 0..4294967295.
    UdpBadLengthPackets interface{}

    // UDP Transmitted. The type is interface{} with range: 0..4294967295.
    UdpOutputPackets interface{}

    // UDP drop for other reason. The type is interface{} with range:
    // 0..4294967295.
    UdpDroppedPackets interface{}
}

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetEntityData() *types.CommonEntityData {
    ipv6Traffic.EntityData.YFilter = ipv6Traffic.YFilter
    ipv6Traffic.EntityData.YangName = "ipv6-traffic"
    ipv6Traffic.EntityData.BundleName = "cisco_ios_xr"
    ipv6Traffic.EntityData.ParentYangName = "statistics"
    ipv6Traffic.EntityData.SegmentPath = "ipv6-traffic"
    ipv6Traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Traffic.EntityData.Children = make(map[string]types.YChild)
    ipv6Traffic.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Traffic.EntityData.Leafs["udp-input-packets"] = types.YLeaf{"UdpInputPackets", ipv6Traffic.UdpInputPackets}
    ipv6Traffic.EntityData.Leafs["udp-checksum-error-packets"] = types.YLeaf{"UdpChecksumErrorPackets", ipv6Traffic.UdpChecksumErrorPackets}
    ipv6Traffic.EntityData.Leafs["udp-no-port-packets"] = types.YLeaf{"UdpNoPortPackets", ipv6Traffic.UdpNoPortPackets}
    ipv6Traffic.EntityData.Leafs["udp-bad-length-packets"] = types.YLeaf{"UdpBadLengthPackets", ipv6Traffic.UdpBadLengthPackets}
    ipv6Traffic.EntityData.Leafs["udp-output-packets"] = types.YLeaf{"UdpOutputPackets", ipv6Traffic.UdpOutputPackets}
    ipv6Traffic.EntityData.Leafs["udp-dropped-packets"] = types.YLeaf{"UdpDroppedPackets", ipv6Traffic.UdpDroppedPackets}
    return &(ipv6Traffic.EntityData)
}

// UdpConnection
// udp connection
type UdpConnection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of UDP connections nodes.
    Nodes UdpConnection_Nodes
}

func (udpConnection *UdpConnection) GetEntityData() *types.CommonEntityData {
    udpConnection.EntityData.YFilter = udpConnection.YFilter
    udpConnection.EntityData.YangName = "udp-connection"
    udpConnection.EntityData.BundleName = "cisco_ios_xr"
    udpConnection.EntityData.ParentYangName = "Cisco-IOS-XR-ip-udp-oper"
    udpConnection.EntityData.SegmentPath = "Cisco-IOS-XR-ip-udp-oper:udp-connection"
    udpConnection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udpConnection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udpConnection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udpConnection.EntityData.Children = make(map[string]types.YChild)
    udpConnection.EntityData.Children["nodes"] = types.YChild{"Nodes", &udpConnection.Nodes}
    udpConnection.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(udpConnection.EntityData)
}

// UdpConnection_Nodes
// List of UDP connections nodes
type UdpConnection_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular node. The type is slice of
    // UdpConnection_Nodes_Node.
    Node []UdpConnection_Nodes_Node
}

func (nodes *UdpConnection_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "udp-connection"
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

// UdpConnection_Nodes_Node
// Information about a particular node
type UdpConnection_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Statistics of UDP connections.
    Statistics UdpConnection_Nodes_Node_Statistics

    // LPTS statistical data.
    Lpts UdpConnection_Nodes_Node_Lpts

    // Detail information for list of UDP connections .
    PcbDetails UdpConnection_Nodes_Node_PcbDetails

    // Brief information for list of UDP connections.
    PcbBriefs UdpConnection_Nodes_Node_PcbBriefs
}

func (node *UdpConnection_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["statistics"] = types.YChild{"Statistics", &node.Statistics}
    node.EntityData.Children["lpts"] = types.YChild{"Lpts", &node.Lpts}
    node.EntityData.Children["pcb-details"] = types.YChild{"PcbDetails", &node.PcbDetails}
    node.EntityData.Children["pcb-briefs"] = types.YChild{"PcbBriefs", &node.PcbBriefs}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// UdpConnection_Nodes_Node_Statistics
// Statistics of UDP connections
type UdpConnection_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table listing clients.
    Clients UdpConnection_Nodes_Node_Statistics_Clients

    // Summary statistics across all UDP connections.
    Summary UdpConnection_Nodes_Node_Statistics_Summary

    // Table listing the UDP connections for which statistics are provided.
    PcbStatistics UdpConnection_Nodes_Node_Statistics_PcbStatistics
}

func (statistics *UdpConnection_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["clients"] = types.YChild{"Clients", &statistics.Clients}
    statistics.EntityData.Children["summary"] = types.YChild{"Summary", &statistics.Summary}
    statistics.EntityData.Children["pcb-statistics"] = types.YChild{"PcbStatistics", &statistics.PcbStatistics}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// UdpConnection_Nodes_Node_Statistics_Clients
// Table listing clients
type UdpConnection_Nodes_Node_Statistics_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describing Client ID. The type is slice of
    // UdpConnection_Nodes_Node_Statistics_Clients_Client.
    Client []UdpConnection_Nodes_Node_Statistics_Clients_Client
}

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "statistics"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = make(map[string]types.YChild)
    clients.EntityData.Children["client"] = types.YChild{"Client", nil}
    for i := range clients.Client {
        clients.EntityData.Children[types.GetSegmentPath(&clients.Client[i])] = types.YChild{"Client", &clients.Client[i]}
    }
    clients.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clients.EntityData)
}

// UdpConnection_Nodes_Node_Statistics_Clients_Client
// Describing Client ID
type UdpConnection_Nodes_Node_Statistics_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Displaying client's aggregated statistics. The
    // type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // Job ID of the transport client. The type is interface{} with range:
    // -2147483648..2147483647.
    ClientJid interface{}

    // Transport client name. The type is string with length: 0..21.
    ClientName interface{}

    // Total IPv4 packets received from client. The type is interface{} with
    // range: 0..4294967295.
    Ipv4ReceivedPackets interface{}

    // Total IPv4 packets sent to client. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SentPackets interface{}

    // Total IPv6 packets received from app. The type is interface{} with range:
    // 0..4294967295.
    Ipv6ReceivedPackets interface{}

    // Total IPv6 packets sent to app. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SentPackets interface{}
}

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["client-id"] = types.YLeaf{"ClientId", client.ClientId}
    client.EntityData.Leafs["client-jid"] = types.YLeaf{"ClientJid", client.ClientJid}
    client.EntityData.Leafs["client-name"] = types.YLeaf{"ClientName", client.ClientName}
    client.EntityData.Leafs["ipv4-received-packets"] = types.YLeaf{"Ipv4ReceivedPackets", client.Ipv4ReceivedPackets}
    client.EntityData.Leafs["ipv4-sent-packets"] = types.YLeaf{"Ipv4SentPackets", client.Ipv4SentPackets}
    client.EntityData.Leafs["ipv6-received-packets"] = types.YLeaf{"Ipv6ReceivedPackets", client.Ipv6ReceivedPackets}
    client.EntityData.Leafs["ipv6-sent-packets"] = types.YLeaf{"Ipv6SentPackets", client.Ipv6SentPackets}
    return &(client.EntityData)
}

// UdpConnection_Nodes_Node_Statistics_Summary
// Summary statistics across all UDP connections
type UdpConnection_Nodes_Node_Statistics_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total packets received. The type is interface{} with range: 0..4294967295.
    ReceivedTotalPackets interface{}

    // Packets received when no wild listener. The type is interface{} with range:
    // 0..4294967295.
    ReceivedNoPortPackets interface{}

    // Packets received has bad checksum. The type is interface{} with range:
    // 0..4294967295.
    ReceivedBadChecksumPackets interface{}

    // Packets received is too short. The type is interface{} with range:
    // 0..4294967295.
    ReceivedTooShortPackets interface{}

    // Packets dropped for other reasons. The type is interface{} with range:
    // 0..4294967295.
    ReceivedDropPackets interface{}

    // Total packets sent. The type is interface{} with range: 0..4294967295.
    SentTotalPackets interface{}

    // Total send erorr packets. The type is interface{} with range:
    // 0..4294967295.
    SentErrorPackets interface{}

    // Total forwarding broadcast packets. The type is interface{} with range:
    // 0..4294967295.
    ForwardBroadcastPackets interface{}

    // Total cloned packets. The type is interface{} with range: 0..4294967295.
    ClonedPackets interface{}

    // Total failed cloned packets. The type is interface{} with range:
    // 0..4294967295.
    FailedClonePackets interface{}
}

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "statistics"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["received-total-packets"] = types.YLeaf{"ReceivedTotalPackets", summary.ReceivedTotalPackets}
    summary.EntityData.Leafs["received-no-port-packets"] = types.YLeaf{"ReceivedNoPortPackets", summary.ReceivedNoPortPackets}
    summary.EntityData.Leafs["received-bad-checksum-packets"] = types.YLeaf{"ReceivedBadChecksumPackets", summary.ReceivedBadChecksumPackets}
    summary.EntityData.Leafs["received-too-short-packets"] = types.YLeaf{"ReceivedTooShortPackets", summary.ReceivedTooShortPackets}
    summary.EntityData.Leafs["received-drop-packets"] = types.YLeaf{"ReceivedDropPackets", summary.ReceivedDropPackets}
    summary.EntityData.Leafs["sent-total-packets"] = types.YLeaf{"SentTotalPackets", summary.SentTotalPackets}
    summary.EntityData.Leafs["sent-error-packets"] = types.YLeaf{"SentErrorPackets", summary.SentErrorPackets}
    summary.EntityData.Leafs["forward-broadcast-packets"] = types.YLeaf{"ForwardBroadcastPackets", summary.ForwardBroadcastPackets}
    summary.EntityData.Leafs["cloned-packets"] = types.YLeaf{"ClonedPackets", summary.ClonedPackets}
    summary.EntityData.Leafs["failed-clone-packets"] = types.YLeaf{"FailedClonePackets", summary.FailedClonePackets}
    return &(summary.EntityData)
}

// UdpConnection_Nodes_Node_Statistics_PcbStatistics
// Table listing the UDP connections for which
// statistics are provided
type UdpConnection_Nodes_Node_Statistics_PcbStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satistics associated with a particular PCB. The type is slice of
    // UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic.
    PcbStatistic []UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic
}

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetEntityData() *types.CommonEntityData {
    pcbStatistics.EntityData.YFilter = pcbStatistics.YFilter
    pcbStatistics.EntityData.YangName = "pcb-statistics"
    pcbStatistics.EntityData.BundleName = "cisco_ios_xr"
    pcbStatistics.EntityData.ParentYangName = "statistics"
    pcbStatistics.EntityData.SegmentPath = "pcb-statistics"
    pcbStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcbStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcbStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcbStatistics.EntityData.Children = make(map[string]types.YChild)
    pcbStatistics.EntityData.Children["pcb-statistic"] = types.YChild{"PcbStatistic", nil}
    for i := range pcbStatistics.PcbStatistic {
        pcbStatistics.EntityData.Children[types.GetSegmentPath(&pcbStatistics.PcbStatistic[i])] = types.YChild{"PcbStatistic", &pcbStatistics.PcbStatistic[i]}
    }
    pcbStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pcbStatistics.EntityData)
}

// UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic
// Satistics associated with a particular PCB
type UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Protocol Control Block address. The type is
    // interface{} with range: 0..4294967295.
    PcbAddress interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // True if paw socket. The type is bool.
    IsPawSocket interface{}

    // UDP send statistics.
    Send UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send

    // UDP receive statistics.
    Receive UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive
}

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetEntityData() *types.CommonEntityData {
    pcbStatistic.EntityData.YFilter = pcbStatistic.YFilter
    pcbStatistic.EntityData.YangName = "pcb-statistic"
    pcbStatistic.EntityData.BundleName = "cisco_ios_xr"
    pcbStatistic.EntityData.ParentYangName = "pcb-statistics"
    pcbStatistic.EntityData.SegmentPath = "pcb-statistic" + "[pcb-address='" + fmt.Sprintf("%v", pcbStatistic.PcbAddress) + "']"
    pcbStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcbStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcbStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcbStatistic.EntityData.Children = make(map[string]types.YChild)
    pcbStatistic.EntityData.Children["send"] = types.YChild{"Send", &pcbStatistic.Send}
    pcbStatistic.EntityData.Children["receive"] = types.YChild{"Receive", &pcbStatistic.Receive}
    pcbStatistic.EntityData.Leafs = make(map[string]types.YLeaf)
    pcbStatistic.EntityData.Leafs["pcb-address"] = types.YLeaf{"PcbAddress", pcbStatistic.PcbAddress}
    pcbStatistic.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", pcbStatistic.VrfId}
    pcbStatistic.EntityData.Leafs["is-paw-socket"] = types.YLeaf{"IsPawSocket", pcbStatistic.IsPawSocket}
    return &(pcbStatistic.EntityData)
}

// UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send
// UDP send statistics
type UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bytes received from application. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ReceivedApplicationBytes interface{}

    // XIPC pulses received from application. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedXipcPulses interface{}

    // Packets sent to network (v4/v6 IO). The type is interface{} with range:
    // 0..18446744073709551615.
    SentNetworkPackets interface{}

    // Packets sent to network (NetIO). The type is interface{} with range:
    // 0..18446744073709551615.
    SentNetIoPackets interface{}

    // Packets failed getting queued to network (v4/v6 IO). The type is
    // interface{} with range: 0..4294967295.
    FailedQueuedNetworkPackets interface{}

    // Packets failed getting queued to network (NetIO). The type is interface{}
    // with range: 0..4294967295.
    FailedQueuedNetIoPackets interface{}
}

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetEntityData() *types.CommonEntityData {
    send.EntityData.YFilter = send.YFilter
    send.EntityData.YangName = "send"
    send.EntityData.BundleName = "cisco_ios_xr"
    send.EntityData.ParentYangName = "pcb-statistic"
    send.EntityData.SegmentPath = "send"
    send.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    send.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    send.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    send.EntityData.Children = make(map[string]types.YChild)
    send.EntityData.Leafs = make(map[string]types.YLeaf)
    send.EntityData.Leafs["received-application-bytes"] = types.YLeaf{"ReceivedApplicationBytes", send.ReceivedApplicationBytes}
    send.EntityData.Leafs["received-xipc-pulses"] = types.YLeaf{"ReceivedXipcPulses", send.ReceivedXipcPulses}
    send.EntityData.Leafs["sent-network-packets"] = types.YLeaf{"SentNetworkPackets", send.SentNetworkPackets}
    send.EntityData.Leafs["sent-net-io-packets"] = types.YLeaf{"SentNetIoPackets", send.SentNetIoPackets}
    send.EntityData.Leafs["failed-queued-network-packets"] = types.YLeaf{"FailedQueuedNetworkPackets", send.FailedQueuedNetworkPackets}
    send.EntityData.Leafs["failed-queued-net-io-packets"] = types.YLeaf{"FailedQueuedNetIoPackets", send.FailedQueuedNetIoPackets}
    return &(send.EntityData)
}

// UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive
// UDP receive statistics
type UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received from network. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedNetworkPackets interface{}

    // Packets failed queued to application. The type is interface{} with range:
    // 0..4294967295.
    FailedQueuedApplicationPackets interface{}

    // Packets queued to application. The type is interface{} with range:
    // 0..18446744073709551615.
    QueuedApplicationPackets interface{}

    // Packet that couldn't be queued to application.on socket. The type is
    // interface{} with range: 0..4294967295.
    FailedQueuedApplicationSocketPackets interface{}

    // Packets queued to application on socket. The type is interface{} with
    // range: 0..18446744073709551615.
    QueuedApplicationSocketPackets interface{}
}

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetEntityData() *types.CommonEntityData {
    receive.EntityData.YFilter = receive.YFilter
    receive.EntityData.YangName = "receive"
    receive.EntityData.BundleName = "cisco_ios_xr"
    receive.EntityData.ParentYangName = "pcb-statistic"
    receive.EntityData.SegmentPath = "receive"
    receive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receive.EntityData.Children = make(map[string]types.YChild)
    receive.EntityData.Leafs = make(map[string]types.YLeaf)
    receive.EntityData.Leafs["received-network-packets"] = types.YLeaf{"ReceivedNetworkPackets", receive.ReceivedNetworkPackets}
    receive.EntityData.Leafs["failed-queued-application-packets"] = types.YLeaf{"FailedQueuedApplicationPackets", receive.FailedQueuedApplicationPackets}
    receive.EntityData.Leafs["queued-application-packets"] = types.YLeaf{"QueuedApplicationPackets", receive.QueuedApplicationPackets}
    receive.EntityData.Leafs["failed-queued-application-socket-packets"] = types.YLeaf{"FailedQueuedApplicationSocketPackets", receive.FailedQueuedApplicationSocketPackets}
    receive.EntityData.Leafs["queued-application-socket-packets"] = types.YLeaf{"QueuedApplicationSocketPackets", receive.QueuedApplicationSocketPackets}
    return &(receive.EntityData)
}

// UdpConnection_Nodes_Node_Lpts
// LPTS statistical data
type UdpConnection_Nodes_Node_Lpts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of query options.
    Queries UdpConnection_Nodes_Node_Lpts_Queries
}

func (lpts *UdpConnection_Nodes_Node_Lpts) GetEntityData() *types.CommonEntityData {
    lpts.EntityData.YFilter = lpts.YFilter
    lpts.EntityData.YangName = "lpts"
    lpts.EntityData.BundleName = "cisco_ios_xr"
    lpts.EntityData.ParentYangName = "node"
    lpts.EntityData.SegmentPath = "lpts"
    lpts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lpts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lpts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lpts.EntityData.Children = make(map[string]types.YChild)
    lpts.EntityData.Children["queries"] = types.YChild{"Queries", &lpts.Queries}
    lpts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lpts.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries
// List of query options
type UdpConnection_Nodes_Node_Lpts_Queries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Query option. The type is slice of
    // UdpConnection_Nodes_Node_Lpts_Queries_Query.
    Query []UdpConnection_Nodes_Node_Lpts_Queries_Query
}

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetEntityData() *types.CommonEntityData {
    queries.EntityData.YFilter = queries.YFilter
    queries.EntityData.YangName = "queries"
    queries.EntityData.BundleName = "cisco_ios_xr"
    queries.EntityData.ParentYangName = "lpts"
    queries.EntityData.SegmentPath = "queries"
    queries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queries.EntityData.Children = make(map[string]types.YChild)
    queries.EntityData.Children["query"] = types.YChild{"Query", nil}
    for i := range queries.Query {
        queries.EntityData.Children[types.GetSegmentPath(&queries.Query[i])] = types.YChild{"Query", &queries.Query[i]}
    }
    queries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(queries.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query
// Query option
type UdpConnection_Nodes_Node_Lpts_Queries_Query struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Query option. The type is LptsPcbQuery.
    QueryName interface{}

    // List of PCBs.
    Pcbs UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs
}

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetEntityData() *types.CommonEntityData {
    query.EntityData.YFilter = query.YFilter
    query.EntityData.YangName = "query"
    query.EntityData.BundleName = "cisco_ios_xr"
    query.EntityData.ParentYangName = "queries"
    query.EntityData.SegmentPath = "query" + "[query-name='" + fmt.Sprintf("%v", query.QueryName) + "']"
    query.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    query.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    query.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    query.EntityData.Children = make(map[string]types.YChild)
    query.EntityData.Children["pcbs"] = types.YChild{"Pcbs", &query.Pcbs}
    query.EntityData.Leafs = make(map[string]types.YLeaf)
    query.EntityData.Leafs["query-name"] = types.YLeaf{"QueryName", query.QueryName}
    return &(query.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs
// List of PCBs
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A PCB information. The type is slice of
    // UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb.
    Pcb []UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb
}

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetEntityData() *types.CommonEntityData {
    pcbs.EntityData.YFilter = pcbs.YFilter
    pcbs.EntityData.YangName = "pcbs"
    pcbs.EntityData.BundleName = "cisco_ios_xr"
    pcbs.EntityData.ParentYangName = "query"
    pcbs.EntityData.SegmentPath = "pcbs"
    pcbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcbs.EntityData.Children = make(map[string]types.YChild)
    pcbs.EntityData.Children["pcb"] = types.YChild{"Pcb", nil}
    for i := range pcbs.Pcb {
        pcbs.EntityData.Children[types.GetSegmentPath(&pcbs.Pcb[i])] = types.YChild{"Pcb", &pcbs.Pcb[i]}
    }
    pcbs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pcbs.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb
// A PCB information
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. PCB address. The type is interface{} with range:
    // 0..4294967295.
    PcbAddress interface{}

    // Layer 4 protocol. The type is interface{} with range: 0..4294967295.
    L4Protocol interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Remote port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // Local IP address.
    LocalAddress UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress

    // Remote IP address.
    ForeignAddress UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress

    // Common PCB information.
    Common UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common
}

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetEntityData() *types.CommonEntityData {
    pcb.EntityData.YFilter = pcb.YFilter
    pcb.EntityData.YangName = "pcb"
    pcb.EntityData.BundleName = "cisco_ios_xr"
    pcb.EntityData.ParentYangName = "pcbs"
    pcb.EntityData.SegmentPath = "pcb" + "[pcb-address='" + fmt.Sprintf("%v", pcb.PcbAddress) + "']"
    pcb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcb.EntityData.Children = make(map[string]types.YChild)
    pcb.EntityData.Children["local-address"] = types.YChild{"LocalAddress", &pcb.LocalAddress}
    pcb.EntityData.Children["foreign-address"] = types.YChild{"ForeignAddress", &pcb.ForeignAddress}
    pcb.EntityData.Children["common"] = types.YChild{"Common", &pcb.Common}
    pcb.EntityData.Leafs = make(map[string]types.YLeaf)
    pcb.EntityData.Leafs["pcb-address"] = types.YLeaf{"PcbAddress", pcb.PcbAddress}
    pcb.EntityData.Leafs["l4-protocol"] = types.YLeaf{"L4Protocol", pcb.L4Protocol}
    pcb.EntityData.Leafs["local-port"] = types.YLeaf{"LocalPort", pcb.LocalPort}
    pcb.EntityData.Leafs["foreign-port"] = types.YLeaf{"ForeignPort", pcb.ForeignPort}
    return &(pcb.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress
// Local IP address
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is AddrFamily.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "pcb"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = make(map[string]types.YChild)
    localAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    localAddress.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", localAddress.AfName}
    localAddress.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", localAddress.Ipv4Address}
    localAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", localAddress.Ipv6Address}
    return &(localAddress.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress
// Remote IP address
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is AddrFamily.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetEntityData() *types.CommonEntityData {
    foreignAddress.EntityData.YFilter = foreignAddress.YFilter
    foreignAddress.EntityData.YangName = "foreign-address"
    foreignAddress.EntityData.BundleName = "cisco_ios_xr"
    foreignAddress.EntityData.ParentYangName = "pcb"
    foreignAddress.EntityData.SegmentPath = "foreign-address"
    foreignAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignAddress.EntityData.Children = make(map[string]types.YChild)
    foreignAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    foreignAddress.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", foreignAddress.AfName}
    foreignAddress.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", foreignAddress.Ipv4Address}
    foreignAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", foreignAddress.Ipv6Address}
    return &(foreignAddress.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common
// Common PCB information
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family. The type is AddrFamily.
    AfName interface{}

    // LPTS PCB information.
    LptsPcb UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb
}

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xr"
    common.EntityData.ParentYangName = "pcb"
    common.EntityData.SegmentPath = "common"
    common.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    common.EntityData.Children = make(map[string]types.YChild)
    common.EntityData.Children["lpts-pcb"] = types.YChild{"LptsPcb", &common.LptsPcb}
    common.EntityData.Leafs = make(map[string]types.YLeaf)
    common.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", common.AfName}
    return &(common.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb
// LPTS PCB information
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum TTL. The type is interface{} with range: 0..255.
    Ttl interface{}

    // flow information. The type is interface{} with range: 0..4294967295.
    FlowTypesInfo interface{}

    // Receive options.
    Options UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options

    // LPTS flags.
    LptsFlags UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags

    // AcceptMask.
    AcceptMask UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask

    // Interface Filters. The type is slice of
    // UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter.
    Filter []UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter
}

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetEntityData() *types.CommonEntityData {
    lptsPcb.EntityData.YFilter = lptsPcb.YFilter
    lptsPcb.EntityData.YangName = "lpts-pcb"
    lptsPcb.EntityData.BundleName = "cisco_ios_xr"
    lptsPcb.EntityData.ParentYangName = "common"
    lptsPcb.EntityData.SegmentPath = "lpts-pcb"
    lptsPcb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lptsPcb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lptsPcb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lptsPcb.EntityData.Children = make(map[string]types.YChild)
    lptsPcb.EntityData.Children["options"] = types.YChild{"Options", &lptsPcb.Options}
    lptsPcb.EntityData.Children["lpts-flags"] = types.YChild{"LptsFlags", &lptsPcb.LptsFlags}
    lptsPcb.EntityData.Children["accept-mask"] = types.YChild{"AcceptMask", &lptsPcb.AcceptMask}
    lptsPcb.EntityData.Children["filter"] = types.YChild{"Filter", nil}
    for i := range lptsPcb.Filter {
        lptsPcb.EntityData.Children[types.GetSegmentPath(&lptsPcb.Filter[i])] = types.YChild{"Filter", &lptsPcb.Filter[i]}
    }
    lptsPcb.EntityData.Leafs = make(map[string]types.YLeaf)
    lptsPcb.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", lptsPcb.Ttl}
    lptsPcb.EntityData.Leafs["flow-types-info"] = types.YLeaf{"FlowTypesInfo", lptsPcb.FlowTypesInfo}
    return &(lptsPcb.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options
// Receive options
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Receive filter enabled. The type is bool.
    IsReceiveFilter interface{}

    // IP SLA. The type is bool.
    IsIpSla interface{}
}

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetEntityData() *types.CommonEntityData {
    options.EntityData.YFilter = options.YFilter
    options.EntityData.YangName = "options"
    options.EntityData.BundleName = "cisco_ios_xr"
    options.EntityData.ParentYangName = "lpts-pcb"
    options.EntityData.SegmentPath = "options"
    options.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    options.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    options.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    options.EntityData.Children = make(map[string]types.YChild)
    options.EntityData.Leafs = make(map[string]types.YLeaf)
    options.EntityData.Leafs["is-receive-filter"] = types.YLeaf{"IsReceiveFilter", options.IsReceiveFilter}
    options.EntityData.Leafs["is-ip-sla"] = types.YLeaf{"IsIpSla", options.IsIpSla}
    return &(options.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags
// LPTS flags
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCB bound. The type is bool.
    IsPcbBound interface{}

    // Sent drop packets. The type is bool.
    IsLocalAddressIgnore interface{}

    // Ignore VRF Filter. The type is bool.
    IsIgnoreVrfFilter interface{}
}

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetEntityData() *types.CommonEntityData {
    lptsFlags.EntityData.YFilter = lptsFlags.YFilter
    lptsFlags.EntityData.YangName = "lpts-flags"
    lptsFlags.EntityData.BundleName = "cisco_ios_xr"
    lptsFlags.EntityData.ParentYangName = "lpts-pcb"
    lptsFlags.EntityData.SegmentPath = "lpts-flags"
    lptsFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lptsFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lptsFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lptsFlags.EntityData.Children = make(map[string]types.YChild)
    lptsFlags.EntityData.Leafs = make(map[string]types.YLeaf)
    lptsFlags.EntityData.Leafs["is-pcb-bound"] = types.YLeaf{"IsPcbBound", lptsFlags.IsPcbBound}
    lptsFlags.EntityData.Leafs["is-local-address-ignore"] = types.YLeaf{"IsLocalAddressIgnore", lptsFlags.IsLocalAddressIgnore}
    lptsFlags.EntityData.Leafs["is-ignore-vrf-filter"] = types.YLeaf{"IsIgnoreVrfFilter", lptsFlags.IsIgnoreVrfFilter}
    return &(lptsFlags.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask
// AcceptMask
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set interface. The type is bool.
    IsInterface interface{}

    // Set packet type. The type is bool.
    IsPacketType interface{}

    // Set Remote address. The type is bool.
    IsRemoteAddress interface{}

    // Set Remote Port. The type is bool.
    IsRemotePort interface{}

    // Set Local Address. The type is bool.
    IsLocalAddress interface{}

    // Set Local Port. The type is bool.
    IsLocalPort interface{}
}

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetEntityData() *types.CommonEntityData {
    acceptMask.EntityData.YFilter = acceptMask.YFilter
    acceptMask.EntityData.YangName = "accept-mask"
    acceptMask.EntityData.BundleName = "cisco_ios_xr"
    acceptMask.EntityData.ParentYangName = "lpts-pcb"
    acceptMask.EntityData.SegmentPath = "accept-mask"
    acceptMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acceptMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acceptMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acceptMask.EntityData.Children = make(map[string]types.YChild)
    acceptMask.EntityData.Leafs = make(map[string]types.YLeaf)
    acceptMask.EntityData.Leafs["is-interface"] = types.YLeaf{"IsInterface", acceptMask.IsInterface}
    acceptMask.EntityData.Leafs["is-packet-type"] = types.YLeaf{"IsPacketType", acceptMask.IsPacketType}
    acceptMask.EntityData.Leafs["is-remote-address"] = types.YLeaf{"IsRemoteAddress", acceptMask.IsRemoteAddress}
    acceptMask.EntityData.Leafs["is-remote-port"] = types.YLeaf{"IsRemotePort", acceptMask.IsRemotePort}
    acceptMask.EntityData.Leafs["is-local-address"] = types.YLeaf{"IsLocalAddress", acceptMask.IsLocalAddress}
    acceptMask.EntityData.Leafs["is-local-port"] = types.YLeaf{"IsLocalPort", acceptMask.IsLocalPort}
    return &(acceptMask.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter
// Interface Filters
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Remote address length. The type is interface{} with range: 0..65535.
    RemoteLength interface{}

    // Local address length. The type is interface{} with range: 0..65535.
    LocalLength interface{}

    // Receive Remote port. The type is interface{} with range: 0..65535.
    ReceiveRemotePort interface{}

    // Receive Local port. The type is interface{} with range: 0..65535.
    ReceiveLocalPort interface{}

    // Priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Minimum TTL. The type is interface{} with range: 0..255.
    Ttl interface{}

    // flow information. The type is interface{} with range: 0..4294967295.
    FlowTypesInfo interface{}

    // Protocol-specific packet type.
    PacketType UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType

    // Remote address.
    RemoteAddress UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress

    // Local address.
    LocalAddress UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress
}

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetEntityData() *types.CommonEntityData {
    filter.EntityData.YFilter = filter.YFilter
    filter.EntityData.YangName = "filter"
    filter.EntityData.BundleName = "cisco_ios_xr"
    filter.EntityData.ParentYangName = "lpts-pcb"
    filter.EntityData.SegmentPath = "filter"
    filter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filter.EntityData.Children = make(map[string]types.YChild)
    filter.EntityData.Children["packet-type"] = types.YChild{"PacketType", &filter.PacketType}
    filter.EntityData.Children["remote-address"] = types.YChild{"RemoteAddress", &filter.RemoteAddress}
    filter.EntityData.Children["local-address"] = types.YChild{"LocalAddress", &filter.LocalAddress}
    filter.EntityData.Leafs = make(map[string]types.YLeaf)
    filter.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", filter.InterfaceName}
    filter.EntityData.Leafs["remote-length"] = types.YLeaf{"RemoteLength", filter.RemoteLength}
    filter.EntityData.Leafs["local-length"] = types.YLeaf{"LocalLength", filter.LocalLength}
    filter.EntityData.Leafs["receive-remote-port"] = types.YLeaf{"ReceiveRemotePort", filter.ReceiveRemotePort}
    filter.EntityData.Leafs["receive-local-port"] = types.YLeaf{"ReceiveLocalPort", filter.ReceiveLocalPort}
    filter.EntityData.Leafs["priority"] = types.YLeaf{"Priority", filter.Priority}
    filter.EntityData.Leafs["ttl"] = types.YLeaf{"Ttl", filter.Ttl}
    filter.EntityData.Leafs["flow-types-info"] = types.YLeaf{"FlowTypesInfo", filter.FlowTypesInfo}
    return &(filter.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType
// Protocol-specific packet type
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is Packet.
    Type_ interface{}

    // ICMP message type. The type is MessageTypeIcmp_.
    IcmpMessageType interface{}

    // ICMPv6 message type. The type is MessageTypeIcmpv6_.
    IcmPv6MessageType interface{}

    // IGMP message type. The type is MessageTypeIgmp_.
    IgmpMessageType interface{}

    // Message type in number. The type is interface{} with range: 0..4294967295.
    MessageId interface{}
}

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetEntityData() *types.CommonEntityData {
    packetType.EntityData.YFilter = packetType.YFilter
    packetType.EntityData.YangName = "packet-type"
    packetType.EntityData.BundleName = "cisco_ios_xr"
    packetType.EntityData.ParentYangName = "filter"
    packetType.EntityData.SegmentPath = "packet-type"
    packetType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetType.EntityData.Children = make(map[string]types.YChild)
    packetType.EntityData.Leafs = make(map[string]types.YLeaf)
    packetType.EntityData.Leafs["type"] = types.YLeaf{"Type_", packetType.Type_}
    packetType.EntityData.Leafs["icmp-message-type"] = types.YLeaf{"IcmpMessageType", packetType.IcmpMessageType}
    packetType.EntityData.Leafs["icm-pv6-message-type"] = types.YLeaf{"IcmPv6MessageType", packetType.IcmPv6MessageType}
    packetType.EntityData.Leafs["igmp-message-type"] = types.YLeaf{"IgmpMessageType", packetType.IgmpMessageType}
    packetType.EntityData.Leafs["message-id"] = types.YLeaf{"MessageId", packetType.MessageId}
    return &(packetType.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress
// Remote address
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is AddrFamily.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetEntityData() *types.CommonEntityData {
    remoteAddress.EntityData.YFilter = remoteAddress.YFilter
    remoteAddress.EntityData.YangName = "remote-address"
    remoteAddress.EntityData.BundleName = "cisco_ios_xr"
    remoteAddress.EntityData.ParentYangName = "filter"
    remoteAddress.EntityData.SegmentPath = "remote-address"
    remoteAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAddress.EntityData.Children = make(map[string]types.YChild)
    remoteAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteAddress.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", remoteAddress.AfName}
    remoteAddress.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", remoteAddress.Ipv4Address}
    remoteAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", remoteAddress.Ipv6Address}
    return &(remoteAddress.EntityData)
}

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress
// Local address
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is AddrFamily.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "filter"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = make(map[string]types.YChild)
    localAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    localAddress.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", localAddress.AfName}
    localAddress.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", localAddress.Ipv4Address}
    localAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", localAddress.Ipv6Address}
    return &(localAddress.EntityData)
}

// UdpConnection_Nodes_Node_PcbDetails
// Detail information for list of UDP connections
// .
type UdpConnection_Nodes_Node_PcbDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detail information about a UDP connection. The type is slice of
    // UdpConnection_Nodes_Node_PcbDetails_PcbDetail.
    PcbDetail []UdpConnection_Nodes_Node_PcbDetails_PcbDetail
}

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetEntityData() *types.CommonEntityData {
    pcbDetails.EntityData.YFilter = pcbDetails.YFilter
    pcbDetails.EntityData.YangName = "pcb-details"
    pcbDetails.EntityData.BundleName = "cisco_ios_xr"
    pcbDetails.EntityData.ParentYangName = "node"
    pcbDetails.EntityData.SegmentPath = "pcb-details"
    pcbDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcbDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcbDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcbDetails.EntityData.Children = make(map[string]types.YChild)
    pcbDetails.EntityData.Children["pcb-detail"] = types.YChild{"PcbDetail", nil}
    for i := range pcbDetails.PcbDetail {
        pcbDetails.EntityData.Children[types.GetSegmentPath(&pcbDetails.PcbDetail[i])] = types.YChild{"PcbDetail", &pcbDetails.PcbDetail[i]}
    }
    pcbDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pcbDetails.EntityData)
}

// UdpConnection_Nodes_Node_PcbDetails_PcbDetail
// Detail information about a UDP connection
type UdpConnection_Nodes_Node_PcbDetails_PcbDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Protocol Control Block address. The type is
    // interface{} with range: 0..4294967295.
    PcbAddress interface{}

    // Address family. The type is UdpAddressFamily.
    AfName interface{}

    // ID of local process. The type is interface{} with range: 0..4294967295.
    LocalProcessId interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Foreign port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // Receive queue count. The type is interface{} with range: 0..4294967295.
    ReceiveQueue interface{}

    // Send queue count. The type is interface{} with range: 0..4294967295.
    SendQueue interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Local address.
    LocalAddress UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress

    // Foreign address.
    ForeignAddress UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress
}

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetEntityData() *types.CommonEntityData {
    pcbDetail.EntityData.YFilter = pcbDetail.YFilter
    pcbDetail.EntityData.YangName = "pcb-detail"
    pcbDetail.EntityData.BundleName = "cisco_ios_xr"
    pcbDetail.EntityData.ParentYangName = "pcb-details"
    pcbDetail.EntityData.SegmentPath = "pcb-detail" + "[pcb-address='" + fmt.Sprintf("%v", pcbDetail.PcbAddress) + "']"
    pcbDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcbDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcbDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcbDetail.EntityData.Children = make(map[string]types.YChild)
    pcbDetail.EntityData.Children["local-address"] = types.YChild{"LocalAddress", &pcbDetail.LocalAddress}
    pcbDetail.EntityData.Children["foreign-address"] = types.YChild{"ForeignAddress", &pcbDetail.ForeignAddress}
    pcbDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    pcbDetail.EntityData.Leafs["pcb-address"] = types.YLeaf{"PcbAddress", pcbDetail.PcbAddress}
    pcbDetail.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", pcbDetail.AfName}
    pcbDetail.EntityData.Leafs["local-process-id"] = types.YLeaf{"LocalProcessId", pcbDetail.LocalProcessId}
    pcbDetail.EntityData.Leafs["local-port"] = types.YLeaf{"LocalPort", pcbDetail.LocalPort}
    pcbDetail.EntityData.Leafs["foreign-port"] = types.YLeaf{"ForeignPort", pcbDetail.ForeignPort}
    pcbDetail.EntityData.Leafs["receive-queue"] = types.YLeaf{"ReceiveQueue", pcbDetail.ReceiveQueue}
    pcbDetail.EntityData.Leafs["send-queue"] = types.YLeaf{"SendQueue", pcbDetail.SendQueue}
    pcbDetail.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", pcbDetail.VrfId}
    return &(pcbDetail.EntityData)
}

// UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress
// Local address
type UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is UdpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "pcb-detail"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = make(map[string]types.YChild)
    localAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    localAddress.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", localAddress.AfName}
    localAddress.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", localAddress.Ipv4Address}
    localAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", localAddress.Ipv6Address}
    return &(localAddress.EntityData)
}

// UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress
// Foreign address
type UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is UdpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetEntityData() *types.CommonEntityData {
    foreignAddress.EntityData.YFilter = foreignAddress.YFilter
    foreignAddress.EntityData.YangName = "foreign-address"
    foreignAddress.EntityData.BundleName = "cisco_ios_xr"
    foreignAddress.EntityData.ParentYangName = "pcb-detail"
    foreignAddress.EntityData.SegmentPath = "foreign-address"
    foreignAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignAddress.EntityData.Children = make(map[string]types.YChild)
    foreignAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    foreignAddress.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", foreignAddress.AfName}
    foreignAddress.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", foreignAddress.Ipv4Address}
    foreignAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", foreignAddress.Ipv6Address}
    return &(foreignAddress.EntityData)
}

// UdpConnection_Nodes_Node_PcbBriefs
// Brief information for list of UDP connections.
type UdpConnection_Nodes_Node_PcbBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information about a UDP connection. The type is slice of
    // UdpConnection_Nodes_Node_PcbBriefs_PcbBrief.
    PcbBrief []UdpConnection_Nodes_Node_PcbBriefs_PcbBrief
}

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetEntityData() *types.CommonEntityData {
    pcbBriefs.EntityData.YFilter = pcbBriefs.YFilter
    pcbBriefs.EntityData.YangName = "pcb-briefs"
    pcbBriefs.EntityData.BundleName = "cisco_ios_xr"
    pcbBriefs.EntityData.ParentYangName = "node"
    pcbBriefs.EntityData.SegmentPath = "pcb-briefs"
    pcbBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcbBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcbBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcbBriefs.EntityData.Children = make(map[string]types.YChild)
    pcbBriefs.EntityData.Children["pcb-brief"] = types.YChild{"PcbBrief", nil}
    for i := range pcbBriefs.PcbBrief {
        pcbBriefs.EntityData.Children[types.GetSegmentPath(&pcbBriefs.PcbBrief[i])] = types.YChild{"PcbBrief", &pcbBriefs.PcbBrief[i]}
    }
    pcbBriefs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pcbBriefs.EntityData)
}

// UdpConnection_Nodes_Node_PcbBriefs_PcbBrief
// Brief information about a UDP connection
type UdpConnection_Nodes_Node_PcbBriefs_PcbBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Protocol Control Block address. The type is
    // interface{} with range: 0..4294967295.
    PcbAddress interface{}

    // Address family. The type is UdpAddressFamily.
    AfName interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Foreign port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // Receive queue count. The type is interface{} with range: 0..4294967295.
    ReceiveQueue interface{}

    // Send queue count. The type is interface{} with range: 0..4294967295.
    SendQueue interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Local address.
    LocalAddress UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress

    // Foreign address.
    ForeignAddress UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress
}

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetEntityData() *types.CommonEntityData {
    pcbBrief.EntityData.YFilter = pcbBrief.YFilter
    pcbBrief.EntityData.YangName = "pcb-brief"
    pcbBrief.EntityData.BundleName = "cisco_ios_xr"
    pcbBrief.EntityData.ParentYangName = "pcb-briefs"
    pcbBrief.EntityData.SegmentPath = "pcb-brief" + "[pcb-address='" + fmt.Sprintf("%v", pcbBrief.PcbAddress) + "']"
    pcbBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcbBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcbBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcbBrief.EntityData.Children = make(map[string]types.YChild)
    pcbBrief.EntityData.Children["local-address"] = types.YChild{"LocalAddress", &pcbBrief.LocalAddress}
    pcbBrief.EntityData.Children["foreign-address"] = types.YChild{"ForeignAddress", &pcbBrief.ForeignAddress}
    pcbBrief.EntityData.Leafs = make(map[string]types.YLeaf)
    pcbBrief.EntityData.Leafs["pcb-address"] = types.YLeaf{"PcbAddress", pcbBrief.PcbAddress}
    pcbBrief.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", pcbBrief.AfName}
    pcbBrief.EntityData.Leafs["local-port"] = types.YLeaf{"LocalPort", pcbBrief.LocalPort}
    pcbBrief.EntityData.Leafs["foreign-port"] = types.YLeaf{"ForeignPort", pcbBrief.ForeignPort}
    pcbBrief.EntityData.Leafs["receive-queue"] = types.YLeaf{"ReceiveQueue", pcbBrief.ReceiveQueue}
    pcbBrief.EntityData.Leafs["send-queue"] = types.YLeaf{"SendQueue", pcbBrief.SendQueue}
    pcbBrief.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", pcbBrief.VrfId}
    return &(pcbBrief.EntityData)
}

// UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress
// Local address
type UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is UdpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "pcb-brief"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = make(map[string]types.YChild)
    localAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    localAddress.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", localAddress.AfName}
    localAddress.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", localAddress.Ipv4Address}
    localAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", localAddress.Ipv6Address}
    return &(localAddress.EntityData)
}

// UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress
// Foreign address
type UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is UdpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetEntityData() *types.CommonEntityData {
    foreignAddress.EntityData.YFilter = foreignAddress.YFilter
    foreignAddress.EntityData.YangName = "foreign-address"
    foreignAddress.EntityData.BundleName = "cisco_ios_xr"
    foreignAddress.EntityData.ParentYangName = "pcb-brief"
    foreignAddress.EntityData.SegmentPath = "foreign-address"
    foreignAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignAddress.EntityData.Children = make(map[string]types.YChild)
    foreignAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    foreignAddress.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", foreignAddress.AfName}
    foreignAddress.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", foreignAddress.Ipv4Address}
    foreignAddress.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", foreignAddress.Ipv6Address}
    return &(foreignAddress.EntityData)
}

