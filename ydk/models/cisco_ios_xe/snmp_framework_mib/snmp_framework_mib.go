// The SNMP Management Architecture MIB
// 
// Copyright (C) The Internet Society (2002). This
// version of this MIB module is part of RFC 3411;
// see the RFC itself for full legal notices.
package snmp_framework_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_framework_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:SNMP-FRAMEWORK-MIB SNMP-FRAMEWORK-MIB}", reflect.TypeOf(SNMPFRAMEWORKMIB{}))
    ydk.RegisterEntity("SNMP-FRAMEWORK-MIB:SNMP-FRAMEWORK-MIB", reflect.TypeOf(SNMPFRAMEWORKMIB{}))
}

type Snmpauthprotocols struct {
}

func (id Snmpauthprotocols) String() string {
	return "SNMP-FRAMEWORK-MIB:snmpAuthProtocols"
}

type Snmpprivprotocols struct {
}

func (id Snmpprivprotocols) String() string {
	return "SNMP-FRAMEWORK-MIB:snmpPrivProtocols"
}

// SnmpSecurityLevel represents authNoPriv is less than authPriv.
type SnmpSecurityLevel string

const (
    SnmpSecurityLevel_noAuthNoPriv SnmpSecurityLevel = "noAuthNoPriv"

    SnmpSecurityLevel_authNoPriv SnmpSecurityLevel = "authNoPriv"

    SnmpSecurityLevel_authPriv SnmpSecurityLevel = "authPriv"
)

// SNMPFRAMEWORKMIB
type SNMPFRAMEWORKMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Snmpengine SNMPFRAMEWORKMIB_Snmpengine
}

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetEntityData() *types.CommonEntityData {
    sNMPFRAMEWORKMIB.EntityData.YFilter = sNMPFRAMEWORKMIB.YFilter
    sNMPFRAMEWORKMIB.EntityData.YangName = "SNMP-FRAMEWORK-MIB"
    sNMPFRAMEWORKMIB.EntityData.BundleName = "cisco_ios_xe"
    sNMPFRAMEWORKMIB.EntityData.ParentYangName = "SNMP-FRAMEWORK-MIB"
    sNMPFRAMEWORKMIB.EntityData.SegmentPath = "SNMP-FRAMEWORK-MIB:SNMP-FRAMEWORK-MIB"
    sNMPFRAMEWORKMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sNMPFRAMEWORKMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sNMPFRAMEWORKMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sNMPFRAMEWORKMIB.EntityData.Children = make(map[string]types.YChild)
    sNMPFRAMEWORKMIB.EntityData.Children["snmpEngine"] = types.YChild{"Snmpengine", &sNMPFRAMEWORKMIB.Snmpengine}
    sNMPFRAMEWORKMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sNMPFRAMEWORKMIB.EntityData)
}

// SNMPFRAMEWORKMIB_Snmpengine
type SNMPFRAMEWORKMIB_Snmpengine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An SNMP engine's administratively-unique identifier.  This information
    // SHOULD be stored in non-volatile storage so that it remains constant across
    // re-initializations of the SNMP engine. The type is string with length:
    // 5..32.
    Snmpengineid interface{}

    // The number of times that the SNMP engine has (re-)initialized itself since
    // snmpEngineID was last configured. The type is interface{} with range:
    // 1..2147483647.
    Snmpengineboots interface{}

    // The number of seconds since the value of the snmpEngineBoots object last
    // changed. When incrementing this object's value would cause it to exceed its
    // maximum, snmpEngineBoots is incremented as if a re-initialization had
    // occurred, and this object's value consequently reverts to zero. The type is
    // interface{} with range: 0..2147483647. Units are seconds.
    Snmpenginetime interface{}

    // The maximum length in octets of an SNMP message which this SNMP engine can
    // send or receive and process, determined as the minimum of the maximum
    // message size values supported among all of the transports available to and
    // supported by the engine. The type is interface{} with range:
    // 484..2147483647.
    Snmpenginemaxmessagesize interface{}
}

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetEntityData() *types.CommonEntityData {
    snmpengine.EntityData.YFilter = snmpengine.YFilter
    snmpengine.EntityData.YangName = "snmpEngine"
    snmpengine.EntityData.BundleName = "cisco_ios_xe"
    snmpengine.EntityData.ParentYangName = "SNMP-FRAMEWORK-MIB"
    snmpengine.EntityData.SegmentPath = "snmpEngine"
    snmpengine.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpengine.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpengine.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpengine.EntityData.Children = make(map[string]types.YChild)
    snmpengine.EntityData.Leafs = make(map[string]types.YLeaf)
    snmpengine.EntityData.Leafs["snmpEngineID"] = types.YLeaf{"Snmpengineid", snmpengine.Snmpengineid}
    snmpengine.EntityData.Leafs["snmpEngineBoots"] = types.YLeaf{"Snmpengineboots", snmpengine.Snmpengineboots}
    snmpengine.EntityData.Leafs["snmpEngineTime"] = types.YLeaf{"Snmpenginetime", snmpengine.Snmpenginetime}
    snmpengine.EntityData.Leafs["snmpEngineMaxMessageSize"] = types.YLeaf{"Snmpenginemaxmessagesize", snmpengine.Snmpenginemaxmessagesize}
    return &(snmpengine.EntityData)
}

