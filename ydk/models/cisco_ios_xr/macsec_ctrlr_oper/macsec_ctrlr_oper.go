// This module contains a collection of YANG definitions
// for Cisco IOS-XR macsec-ctrlr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   macsec-ctrlr-oper: Macsec controller data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package macsec_ctrlr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package macsec_ctrlr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-macsec-ctrlr-oper macsec-ctrlr-oper}", reflect.TypeOf(MacsecCtrlrOper{}))
    ydk.RegisterEntity("Cisco-IOS-XR-macsec-ctrlr-oper:macsec-ctrlr-oper", reflect.TypeOf(MacsecCtrlrOper{}))
}

// MacsecCtrlrCiphersuit represents Macsec ctrlr ciphersuit
type MacsecCtrlrCiphersuit string

const (
    // GCM AES 256
    MacsecCtrlrCiphersuit_gcm_aes_256 MacsecCtrlrCiphersuit = "gcm-aes-256"

    // GCM AES 128
    MacsecCtrlrCiphersuit_gcm_aes_128 MacsecCtrlrCiphersuit = "gcm-aes-128"

    // GCM AES XPN 256
    MacsecCtrlrCiphersuit_gcm_aes_xpn_256 MacsecCtrlrCiphersuit = "gcm-aes-xpn-256"
)

// MacsecCtrlrState represents Macsec ctrlr state
type MacsecCtrlrState string

const (
    // Up
    MacsecCtrlrState_macsec_ctrlr_state_up MacsecCtrlrState = "macsec-ctrlr-state-up"

    // Down
    MacsecCtrlrState_macsec_ctrlr_state_down MacsecCtrlrState = "macsec-ctrlr-state-down"

    // Administratively Down
    MacsecCtrlrState_macsec_ctrlr_state_admin_down MacsecCtrlrState = "macsec-ctrlr-state-admin-down"
)

// MacsecCtrlrOper
// Macsec controller data
type MacsecCtrlrOper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All Macsec Controller Port operational data.
    MacsecCtrlrPorts MacsecCtrlrOper_MacsecCtrlrPorts
}

func (macsecCtrlrOper *MacsecCtrlrOper) GetEntityData() *types.CommonEntityData {
    macsecCtrlrOper.EntityData.YFilter = macsecCtrlrOper.YFilter
    macsecCtrlrOper.EntityData.YangName = "macsec-ctrlr-oper"
    macsecCtrlrOper.EntityData.BundleName = "cisco_ios_xr"
    macsecCtrlrOper.EntityData.ParentYangName = "Cisco-IOS-XR-macsec-ctrlr-oper"
    macsecCtrlrOper.EntityData.SegmentPath = "Cisco-IOS-XR-macsec-ctrlr-oper:macsec-ctrlr-oper"
    macsecCtrlrOper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecCtrlrOper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecCtrlrOper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecCtrlrOper.EntityData.Children = make(map[string]types.YChild)
    macsecCtrlrOper.EntityData.Children["macsec-ctrlr-ports"] = types.YChild{"MacsecCtrlrPorts", &macsecCtrlrOper.MacsecCtrlrPorts}
    macsecCtrlrOper.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macsecCtrlrOper.EntityData)
}

// MacsecCtrlrOper_MacsecCtrlrPorts
// All Macsec Controller Port operational data
type MacsecCtrlrOper_MacsecCtrlrPorts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controller name. The type is slice of
    // MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort.
    MacsecCtrlrPort []MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort
}

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetEntityData() *types.CommonEntityData {
    macsecCtrlrPorts.EntityData.YFilter = macsecCtrlrPorts.YFilter
    macsecCtrlrPorts.EntityData.YangName = "macsec-ctrlr-ports"
    macsecCtrlrPorts.EntityData.BundleName = "cisco_ios_xr"
    macsecCtrlrPorts.EntityData.ParentYangName = "macsec-ctrlr-oper"
    macsecCtrlrPorts.EntityData.SegmentPath = "macsec-ctrlr-ports"
    macsecCtrlrPorts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecCtrlrPorts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecCtrlrPorts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecCtrlrPorts.EntityData.Children = make(map[string]types.YChild)
    macsecCtrlrPorts.EntityData.Children["macsec-ctrlr-port"] = types.YChild{"MacsecCtrlrPort", nil}
    for i := range macsecCtrlrPorts.MacsecCtrlrPort {
        macsecCtrlrPorts.EntityData.Children[types.GetSegmentPath(&macsecCtrlrPorts.MacsecCtrlrPort[i])] = types.YChild{"MacsecCtrlrPort", &macsecCtrlrPorts.MacsecCtrlrPort[i]}
    }
    macsecCtrlrPorts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macsecCtrlrPorts.EntityData)
}

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort
// Controller name
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // Macsec Controller operational data.
    MacsecCtrlrInfo MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo
}

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetEntityData() *types.CommonEntityData {
    macsecCtrlrPort.EntityData.YFilter = macsecCtrlrPort.YFilter
    macsecCtrlrPort.EntityData.YangName = "macsec-ctrlr-port"
    macsecCtrlrPort.EntityData.BundleName = "cisco_ios_xr"
    macsecCtrlrPort.EntityData.ParentYangName = "macsec-ctrlr-ports"
    macsecCtrlrPort.EntityData.SegmentPath = "macsec-ctrlr-port" + "[name='" + fmt.Sprintf("%v", macsecCtrlrPort.Name) + "']"
    macsecCtrlrPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecCtrlrPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecCtrlrPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecCtrlrPort.EntityData.Children = make(map[string]types.YChild)
    macsecCtrlrPort.EntityData.Children["macsec-ctrlr-info"] = types.YChild{"MacsecCtrlrInfo", &macsecCtrlrPort.MacsecCtrlrInfo}
    macsecCtrlrPort.EntityData.Leafs = make(map[string]types.YLeaf)
    macsecCtrlrPort.EntityData.Leafs["name"] = types.YLeaf{"Name", macsecCtrlrPort.Name}
    return &(macsecCtrlrPort.EntityData)
}

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo
// Macsec Controller operational data
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State. The type is MacsecCtrlrState.
    State interface{}

    // Replay Window Size. The type is interface{} with range: 0..4294967295.
    ReplayWindowSize interface{}

    // Must Secure. The type is bool.
    MustSecure interface{}

    // Secure Mode. The type is interface{} with range: 0..4294967295.
    SecureMode interface{}

    // Encrypt Secure Channel Status.
    EncryptScStatus MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus

    // Decrypt Secure Channel Status.
    DecryptScStatus MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus
}

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetEntityData() *types.CommonEntityData {
    macsecCtrlrInfo.EntityData.YFilter = macsecCtrlrInfo.YFilter
    macsecCtrlrInfo.EntityData.YangName = "macsec-ctrlr-info"
    macsecCtrlrInfo.EntityData.BundleName = "cisco_ios_xr"
    macsecCtrlrInfo.EntityData.ParentYangName = "macsec-ctrlr-port"
    macsecCtrlrInfo.EntityData.SegmentPath = "macsec-ctrlr-info"
    macsecCtrlrInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecCtrlrInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecCtrlrInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecCtrlrInfo.EntityData.Children = make(map[string]types.YChild)
    macsecCtrlrInfo.EntityData.Children["encrypt-sc-status"] = types.YChild{"EncryptScStatus", &macsecCtrlrInfo.EncryptScStatus}
    macsecCtrlrInfo.EntityData.Children["decrypt-sc-status"] = types.YChild{"DecryptScStatus", &macsecCtrlrInfo.DecryptScStatus}
    macsecCtrlrInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    macsecCtrlrInfo.EntityData.Leafs["state"] = types.YLeaf{"State", macsecCtrlrInfo.State}
    macsecCtrlrInfo.EntityData.Leafs["replay-window-size"] = types.YLeaf{"ReplayWindowSize", macsecCtrlrInfo.ReplayWindowSize}
    macsecCtrlrInfo.EntityData.Leafs["must-secure"] = types.YLeaf{"MustSecure", macsecCtrlrInfo.MustSecure}
    macsecCtrlrInfo.EntityData.Leafs["secure-mode"] = types.YLeaf{"SecureMode", macsecCtrlrInfo.SecureMode}
    return &(macsecCtrlrInfo.EntityData)
}

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus
// Encrypt Secure Channel Status
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protection Enabled. The type is bool.
    ProtectionEnabled interface{}

    // Secure Channel Id. The type is string.
    SecureChannelId interface{}

    // Confidentiality offset. The type is interface{} with range: 0..4294967295.
    ConfidentialityOffset interface{}

    // Cipher Suite. The type is MacsecCtrlrCiphersuit.
    CipherSuite interface{}

    // Max packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPacketNumber interface{}

    // Recent Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    RecentPacketNumber interface{}

    // Active Associations. The type is slice of
    // MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation.
    ActiveAssociation []MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation
}

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetEntityData() *types.CommonEntityData {
    encryptScStatus.EntityData.YFilter = encryptScStatus.YFilter
    encryptScStatus.EntityData.YangName = "encrypt-sc-status"
    encryptScStatus.EntityData.BundleName = "cisco_ios_xr"
    encryptScStatus.EntityData.ParentYangName = "macsec-ctrlr-info"
    encryptScStatus.EntityData.SegmentPath = "encrypt-sc-status"
    encryptScStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptScStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptScStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptScStatus.EntityData.Children = make(map[string]types.YChild)
    encryptScStatus.EntityData.Children["active-association"] = types.YChild{"ActiveAssociation", nil}
    for i := range encryptScStatus.ActiveAssociation {
        encryptScStatus.EntityData.Children[types.GetSegmentPath(&encryptScStatus.ActiveAssociation[i])] = types.YChild{"ActiveAssociation", &encryptScStatus.ActiveAssociation[i]}
    }
    encryptScStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    encryptScStatus.EntityData.Leafs["protection-enabled"] = types.YLeaf{"ProtectionEnabled", encryptScStatus.ProtectionEnabled}
    encryptScStatus.EntityData.Leafs["secure-channel-id"] = types.YLeaf{"SecureChannelId", encryptScStatus.SecureChannelId}
    encryptScStatus.EntityData.Leafs["confidentiality-offset"] = types.YLeaf{"ConfidentialityOffset", encryptScStatus.ConfidentialityOffset}
    encryptScStatus.EntityData.Leafs["cipher-suite"] = types.YLeaf{"CipherSuite", encryptScStatus.CipherSuite}
    encryptScStatus.EntityData.Leafs["max-packet-number"] = types.YLeaf{"MaxPacketNumber", encryptScStatus.MaxPacketNumber}
    encryptScStatus.EntityData.Leafs["recent-packet-number"] = types.YLeaf{"RecentPacketNumber", encryptScStatus.RecentPacketNumber}
    return &(encryptScStatus.EntityData)
}

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation
// Active Associations
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Association Number. The type is interface{} with range: 0..255.
    AssociationNumber interface{}

    // Short secure channel id. The type is interface{} with range: 0..4294967295.
    ShortSecureChannelId interface{}
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetEntityData() *types.CommonEntityData {
    activeAssociation.EntityData.YFilter = activeAssociation.YFilter
    activeAssociation.EntityData.YangName = "active-association"
    activeAssociation.EntityData.BundleName = "cisco_ios_xr"
    activeAssociation.EntityData.ParentYangName = "encrypt-sc-status"
    activeAssociation.EntityData.SegmentPath = "active-association"
    activeAssociation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeAssociation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeAssociation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeAssociation.EntityData.Children = make(map[string]types.YChild)
    activeAssociation.EntityData.Leafs = make(map[string]types.YLeaf)
    activeAssociation.EntityData.Leafs["association-number"] = types.YLeaf{"AssociationNumber", activeAssociation.AssociationNumber}
    activeAssociation.EntityData.Leafs["short-secure-channel-id"] = types.YLeaf{"ShortSecureChannelId", activeAssociation.ShortSecureChannelId}
    return &(activeAssociation.EntityData)
}

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus
// Decrypt Secure Channel Status
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protection Enabled. The type is bool.
    ProtectionEnabled interface{}

    // Secure Channel Id. The type is string.
    SecureChannelId interface{}

    // Confidentiality offset. The type is interface{} with range: 0..4294967295.
    ConfidentialityOffset interface{}

    // Cipher Suite. The type is MacsecCtrlrCiphersuit.
    CipherSuite interface{}

    // Max packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPacketNumber interface{}

    // Recent Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    RecentPacketNumber interface{}

    // Active Associations. The type is slice of
    // MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation.
    ActiveAssociation []MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation
}

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetEntityData() *types.CommonEntityData {
    decryptScStatus.EntityData.YFilter = decryptScStatus.YFilter
    decryptScStatus.EntityData.YangName = "decrypt-sc-status"
    decryptScStatus.EntityData.BundleName = "cisco_ios_xr"
    decryptScStatus.EntityData.ParentYangName = "macsec-ctrlr-info"
    decryptScStatus.EntityData.SegmentPath = "decrypt-sc-status"
    decryptScStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    decryptScStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    decryptScStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    decryptScStatus.EntityData.Children = make(map[string]types.YChild)
    decryptScStatus.EntityData.Children["active-association"] = types.YChild{"ActiveAssociation", nil}
    for i := range decryptScStatus.ActiveAssociation {
        decryptScStatus.EntityData.Children[types.GetSegmentPath(&decryptScStatus.ActiveAssociation[i])] = types.YChild{"ActiveAssociation", &decryptScStatus.ActiveAssociation[i]}
    }
    decryptScStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    decryptScStatus.EntityData.Leafs["protection-enabled"] = types.YLeaf{"ProtectionEnabled", decryptScStatus.ProtectionEnabled}
    decryptScStatus.EntityData.Leafs["secure-channel-id"] = types.YLeaf{"SecureChannelId", decryptScStatus.SecureChannelId}
    decryptScStatus.EntityData.Leafs["confidentiality-offset"] = types.YLeaf{"ConfidentialityOffset", decryptScStatus.ConfidentialityOffset}
    decryptScStatus.EntityData.Leafs["cipher-suite"] = types.YLeaf{"CipherSuite", decryptScStatus.CipherSuite}
    decryptScStatus.EntityData.Leafs["max-packet-number"] = types.YLeaf{"MaxPacketNumber", decryptScStatus.MaxPacketNumber}
    decryptScStatus.EntityData.Leafs["recent-packet-number"] = types.YLeaf{"RecentPacketNumber", decryptScStatus.RecentPacketNumber}
    return &(decryptScStatus.EntityData)
}

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation
// Active Associations
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Association Number. The type is interface{} with range: 0..255.
    AssociationNumber interface{}

    // Short secure channel id. The type is interface{} with range: 0..4294967295.
    ShortSecureChannelId interface{}
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetEntityData() *types.CommonEntityData {
    activeAssociation.EntityData.YFilter = activeAssociation.YFilter
    activeAssociation.EntityData.YangName = "active-association"
    activeAssociation.EntityData.BundleName = "cisco_ios_xr"
    activeAssociation.EntityData.ParentYangName = "decrypt-sc-status"
    activeAssociation.EntityData.SegmentPath = "active-association"
    activeAssociation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeAssociation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeAssociation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeAssociation.EntityData.Children = make(map[string]types.YChild)
    activeAssociation.EntityData.Leafs = make(map[string]types.YLeaf)
    activeAssociation.EntityData.Leafs["association-number"] = types.YLeaf{"AssociationNumber", activeAssociation.AssociationNumber}
    activeAssociation.EntityData.Leafs["short-secure-channel-id"] = types.YLeaf{"ShortSecureChannelId", activeAssociation.ShortSecureChannelId}
    return &(activeAssociation.EntityData)
}

