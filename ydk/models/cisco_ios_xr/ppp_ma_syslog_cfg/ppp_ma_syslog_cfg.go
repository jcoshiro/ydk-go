// This module contains a collection of YANG definitions
// for Cisco IOS-XR ppp-ma-syslog package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ppp: PPP configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ppp_ma_syslog_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ppp_ma_syslog_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ppp-ma-syslog-cfg ppp}", reflect.TypeOf(Ppp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ppp-ma-syslog-cfg:ppp", reflect.TypeOf(Ppp{}))
}

// Ppp
// PPP configuration
type Ppp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // syslog option for session status.
    Syslog Ppp_Syslog
}

func (ppp *Ppp) GetEntityData() *types.CommonEntityData {
    ppp.EntityData.YFilter = ppp.YFilter
    ppp.EntityData.YangName = "ppp"
    ppp.EntityData.BundleName = "cisco_ios_xr"
    ppp.EntityData.ParentYangName = "Cisco-IOS-XR-ppp-ma-syslog-cfg"
    ppp.EntityData.SegmentPath = "Cisco-IOS-XR-ppp-ma-syslog-cfg:ppp"
    ppp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ppp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ppp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ppp.EntityData.Children = make(map[string]types.YChild)
    ppp.EntityData.Children["syslog"] = types.YChild{"Syslog", &ppp.Syslog}
    ppp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ppp.EntityData)
}

// Ppp_Syslog
// syslog option for session status
type Ppp_Syslog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable syslog for ppp session status. The type is interface{}.
    EnableSessionStatus interface{}
}

func (syslog *Ppp_Syslog) GetEntityData() *types.CommonEntityData {
    syslog.EntityData.YFilter = syslog.YFilter
    syslog.EntityData.YangName = "syslog"
    syslog.EntityData.BundleName = "cisco_ios_xr"
    syslog.EntityData.ParentYangName = "ppp"
    syslog.EntityData.SegmentPath = "syslog"
    syslog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslog.EntityData.Children = make(map[string]types.YChild)
    syslog.EntityData.Leafs = make(map[string]types.YLeaf)
    syslog.EntityData.Leafs["enable-session-status"] = types.YLeaf{"EnableSessionStatus", syslog.EnableSessionStatus}
    return &(syslog.EntityData)
}

