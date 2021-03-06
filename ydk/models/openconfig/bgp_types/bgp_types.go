// This module contains general data definitions for use in BGP
// policy. It can be imported by modules that make use of BGP
// attributes
package bgp_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bgp_types"))
}

type BGPCAPABILITY struct {
}

func (id BGPCAPABILITY) String() string {
	return "openconfig-bgp-types:BGP_CAPABILITY"
}

type MPBGP struct {
}

func (id MPBGP) String() string {
	return "openconfig-bgp-types:MPBGP"
}

type ROUTEREFRESH struct {
}

func (id ROUTEREFRESH) String() string {
	return "openconfig-bgp-types:ROUTE_REFRESH"
}

type ASN32 struct {
}

func (id ASN32) String() string {
	return "openconfig-bgp-types:ASN32"
}

type GRACEFULRESTART struct {
}

func (id GRACEFULRESTART) String() string {
	return "openconfig-bgp-types:GRACEFUL_RESTART"
}

type ADDPATHS struct {
}

func (id ADDPATHS) String() string {
	return "openconfig-bgp-types:ADD_PATHS"
}

type AFISAFITYPE struct {
}

func (id AFISAFITYPE) String() string {
	return "openconfig-bgp-types:AFI_SAFI_TYPE"
}

type IPV4UNICAST struct {
}

func (id IPV4UNICAST) String() string {
	return "openconfig-bgp-types:IPV4_UNICAST"
}

type IPV6UNICAST struct {
}

func (id IPV6UNICAST) String() string {
	return "openconfig-bgp-types:IPV6_UNICAST"
}

type IPV4LABELEDUNICAST struct {
}

func (id IPV4LABELEDUNICAST) String() string {
	return "openconfig-bgp-types:IPV4_LABELED_UNICAST"
}

type IPV6LABELEDUNICAST struct {
}

func (id IPV6LABELEDUNICAST) String() string {
	return "openconfig-bgp-types:IPV6_LABELED_UNICAST"
}

type L3VPNIPV4UNICAST struct {
}

func (id L3VPNIPV4UNICAST) String() string {
	return "openconfig-bgp-types:L3VPN_IPV4_UNICAST"
}

type L3VPNIPV6UNICAST struct {
}

func (id L3VPNIPV6UNICAST) String() string {
	return "openconfig-bgp-types:L3VPN_IPV6_UNICAST"
}

type L3VPNIPV4MULTICAST struct {
}

func (id L3VPNIPV4MULTICAST) String() string {
	return "openconfig-bgp-types:L3VPN_IPV4_MULTICAST"
}

type L3VPNIPV6MULTICAST struct {
}

func (id L3VPNIPV6MULTICAST) String() string {
	return "openconfig-bgp-types:L3VPN_IPV6_MULTICAST"
}

type L2VPNVPLS struct {
}

func (id L2VPNVPLS) String() string {
	return "openconfig-bgp-types:L2VPN_VPLS"
}

type L2VPNEVPN struct {
}

func (id L2VPNEVPN) String() string {
	return "openconfig-bgp-types:L2VPN_EVPN"
}

type BGPWELLKNOWNSTDCOMMUNITY struct {
}

func (id BGPWELLKNOWNSTDCOMMUNITY) String() string {
	return "openconfig-bgp-types:BGP_WELL_KNOWN_STD_COMMUNITY"
}

type NOEXPORT struct {
}

func (id NOEXPORT) String() string {
	return "openconfig-bgp-types:NO_EXPORT"
}

type NOADVERTISE struct {
}

func (id NOADVERTISE) String() string {
	return "openconfig-bgp-types:NO_ADVERTISE"
}

type NOEXPORTSUBCONFED struct {
}

func (id NOEXPORTSUBCONFED) String() string {
	return "openconfig-bgp-types:NO_EXPORT_SUBCONFED"
}

type NOPEER struct {
}

func (id NOPEER) String() string {
	return "openconfig-bgp-types:NOPEER"
}

type REMOVEPRIVATEASOPTION struct {
}

func (id REMOVEPRIVATEASOPTION) String() string {
	return "openconfig-bgp-types:REMOVE_PRIVATE_AS_OPTION"
}

type PRIVATEASREMOVEALL struct {
}

func (id PRIVATEASREMOVEALL) String() string {
	return "openconfig-bgp-types:PRIVATE_AS_REMOVE_ALL"
}

type PRIVATEASREPLACEALL struct {
}

func (id PRIVATEASREPLACEALL) String() string {
	return "openconfig-bgp-types:PRIVATE_AS_REPLACE_ALL"
}

// BgpSessionDirection represents Type to describe the direction of NLRI transmission
type BgpSessionDirection string

const (
    // Refers to all NLRI received from the BGP peer
    BgpSessionDirection_INBOUND BgpSessionDirection = "INBOUND"

    // Refers to all NLRI advertised to the BGP peer
    BgpSessionDirection_OUTBOUND BgpSessionDirection = "OUTBOUND"
)

// BgpOriginAttrType represents Type definition for standard BGP origin attribute
type BgpOriginAttrType string

const (
    // Origin of the NLRI is internal
    BgpOriginAttrType_IGP BgpOriginAttrType = "IGP"

    // Origin of the NLRI is EGP
    BgpOriginAttrType_EGP BgpOriginAttrType = "EGP"

    // Origin of the NLRI is neither IGP or EGP
    BgpOriginAttrType_INCOMPLETE BgpOriginAttrType = "INCOMPLETE"
)

// PeerType represents external
type PeerType string

const (
    // internal (iBGP) peer
    PeerType_INTERNAL PeerType = "INTERNAL"

    // external (eBGP) peer
    PeerType_EXTERNAL PeerType = "EXTERNAL"
)

// CommunityType represents BOTH: both standard and extended community
type CommunityType string

const (
    // send only standard communities
    CommunityType_STANDARD CommunityType = "STANDARD"

    // send only extended communities
    CommunityType_EXTENDED CommunityType = "EXTENDED"

    // send both standard and extended communities
    CommunityType_BOTH CommunityType = "BOTH"

    // do not send any community attribute
    CommunityType_NONE CommunityType = "NONE"
)

