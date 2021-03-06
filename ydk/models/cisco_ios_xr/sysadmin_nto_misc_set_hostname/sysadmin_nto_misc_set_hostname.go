// This module contains a collection of YANG definitions
// for Cisco IOS-XR syadmin hostname configuration and cli.
// 
// This module contains definitions
// for the following management objects:
// hostname cli and configuration data
// 
// Copyright (c) 2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_nto_misc_set_hostname

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_nto_misc_set_hostname"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-nto-misc-set-hostname hostname}", reflect.TypeOf(Hostname{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-nto-misc-set-hostname:hostname", reflect.TypeOf(Hostname{}))
}

// Hostname
// Set system`s network name
type Hostname struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern: b'[a-zA-Z0-9_.{}+-]+'.
    Name interface{}
}

func (hostname *Hostname) GetEntityData() *types.CommonEntityData {
    hostname.EntityData.YFilter = hostname.YFilter
    hostname.EntityData.YangName = "hostname"
    hostname.EntityData.BundleName = "cisco_ios_xr"
    hostname.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-nto-misc-set-hostname"
    hostname.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-nto-misc-set-hostname:hostname"
    hostname.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostname.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostname.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostname.EntityData.Children = make(map[string]types.YChild)
    hostname.EntityData.Leafs = make(map[string]types.YLeaf)
    hostname.EntityData.Leafs["name"] = types.YLeaf{"Name", hostname.Name}
    return &(hostname.EntityData)
}

