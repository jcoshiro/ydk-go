// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the top level container for
// all 'console' commands for Sysadmin.
// 
// Copyright(c) 2015-2016 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_console

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_console"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-console console}", reflect.TypeOf(Console{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-console:console", reflect.TypeOf(Console{}))
}

// Console
type Console struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Config Console_Config
}

func (console *Console) GetEntityData() *types.CommonEntityData {
    console.EntityData.YFilter = console.YFilter
    console.EntityData.YangName = "console"
    console.EntityData.BundleName = "cisco_ios_xr"
    console.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-console"
    console.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-console:console"
    console.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    console.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    console.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    console.EntityData.Children = make(map[string]types.YChild)
    console.EntityData.Children["config"] = types.YChild{"Config", &console.Config}
    console.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(console.EntityData)
}

// Console_Config
type Console_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    AttachSdr Console_Config_AttachSdr
}

func (config *Console_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "console"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Children["attach-sdr"] = types.YChild{"AttachSdr", &config.AttachSdr}
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(config.EntityData)
}

// Console_Config_AttachSdr
type Console_Config_AttachSdr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Runtime Console_Config_AttachSdr_Runtime

    
    Boot Console_Config_AttachSdr_Boot
}

func (attachSdr *Console_Config_AttachSdr) GetEntityData() *types.CommonEntityData {
    attachSdr.EntityData.YFilter = attachSdr.YFilter
    attachSdr.EntityData.YangName = "attach-sdr"
    attachSdr.EntityData.BundleName = "cisco_ios_xr"
    attachSdr.EntityData.ParentYangName = "config"
    attachSdr.EntityData.SegmentPath = "attach-sdr"
    attachSdr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachSdr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachSdr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachSdr.EntityData.Children = make(map[string]types.YChild)
    attachSdr.EntityData.Children["runtime"] = types.YChild{"Runtime", &attachSdr.Runtime}
    attachSdr.EntityData.Children["boot"] = types.YChild{"Boot", &attachSdr.Boot}
    attachSdr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attachSdr.EntityData)
}

// Console_Config_AttachSdr_Runtime
type Console_Config_AttachSdr_Runtime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Console_Config_AttachSdr_Runtime_Location.
    Location []Console_Config_AttachSdr_Runtime_Location
}

func (runtime *Console_Config_AttachSdr_Runtime) GetEntityData() *types.CommonEntityData {
    runtime.EntityData.YFilter = runtime.YFilter
    runtime.EntityData.YangName = "runtime"
    runtime.EntityData.BundleName = "cisco_ios_xr"
    runtime.EntityData.ParentYangName = "attach-sdr"
    runtime.EntityData.SegmentPath = "runtime"
    runtime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    runtime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    runtime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    runtime.EntityData.Children = make(map[string]types.YChild)
    runtime.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range runtime.Location {
        runtime.EntityData.Children[types.GetSegmentPath(&runtime.Location[i])] = types.YChild{"Location", &runtime.Location[i]}
    }
    runtime.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(runtime.EntityData)
}

// Console_Config_AttachSdr_Runtime_Location
type Console_Config_AttachSdr_Runtime_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((0?[0-9]|1[1-5])/([rR][pP]\\d{1,2}))(/[cC][pP][uU]0)?'.
    LocationRp interface{}

    // The type is slice of Console_Config_AttachSdr_Runtime_Location_TtyName.
    TtyName []Console_Config_AttachSdr_Runtime_Location_TtyName
}

func (location *Console_Config_AttachSdr_Runtime_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "runtime"
    location.EntityData.SegmentPath = "location" + "[location-rp='" + fmt.Sprintf("%v", location.LocationRp) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["tty-name"] = types.YChild{"TtyName", nil}
    for i := range location.TtyName {
        location.EntityData.Children[types.GetSegmentPath(&location.TtyName[i])] = types.YChild{"TtyName", &location.TtyName[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location-rp"] = types.YLeaf{"LocationRp", location.LocationRp}
    return &(location.EntityData)
}

// Console_Config_AttachSdr_Runtime_Location_TtyName
type Console_Config_AttachSdr_Runtime_Location_TtyName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'console1|console2'.
    Ttyname interface{}

    // The type is string with pattern: b'[a-zA-Z0-9_.{}+-]+'.
    SdrName interface{}
}

func (ttyName *Console_Config_AttachSdr_Runtime_Location_TtyName) GetEntityData() *types.CommonEntityData {
    ttyName.EntityData.YFilter = ttyName.YFilter
    ttyName.EntityData.YangName = "tty-name"
    ttyName.EntityData.BundleName = "cisco_ios_xr"
    ttyName.EntityData.ParentYangName = "location"
    ttyName.EntityData.SegmentPath = "tty-name" + "[ttyname='" + fmt.Sprintf("%v", ttyName.Ttyname) + "']"
    ttyName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ttyName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ttyName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ttyName.EntityData.Children = make(map[string]types.YChild)
    ttyName.EntityData.Leafs = make(map[string]types.YLeaf)
    ttyName.EntityData.Leafs["ttyname"] = types.YLeaf{"Ttyname", ttyName.Ttyname}
    ttyName.EntityData.Leafs["sdr-name"] = types.YLeaf{"SdrName", ttyName.SdrName}
    return &(ttyName.EntityData)
}

// Console_Config_AttachSdr_Boot
type Console_Config_AttachSdr_Boot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Console_Config_AttachSdr_Boot_Location.
    Location []Console_Config_AttachSdr_Boot_Location
}

func (boot *Console_Config_AttachSdr_Boot) GetEntityData() *types.CommonEntityData {
    boot.EntityData.YFilter = boot.YFilter
    boot.EntityData.YangName = "boot"
    boot.EntityData.BundleName = "cisco_ios_xr"
    boot.EntityData.ParentYangName = "attach-sdr"
    boot.EntityData.SegmentPath = "boot"
    boot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    boot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    boot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    boot.EntityData.Children = make(map[string]types.YChild)
    boot.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range boot.Location {
        boot.EntityData.Children[types.GetSegmentPath(&boot.Location[i])] = types.YChild{"Location", &boot.Location[i]}
    }
    boot.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(boot.EntityData)
}

// Console_Config_AttachSdr_Boot_Location
type Console_Config_AttachSdr_Boot_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((0?[0-9]|1[1-5])/([rR][pP]\\d{1,2}))(/[cC][pP][uU]0)?'.
    LocationRp interface{}

    // The type is string with pattern: b'[a-zA-Z0-9_.{}+-]+'.
    SdrName interface{}
}

func (location *Console_Config_AttachSdr_Boot_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "boot"
    location.EntityData.SegmentPath = "location" + "[location-rp='" + fmt.Sprintf("%v", location.LocationRp) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location-rp"] = types.YLeaf{"LocationRp", location.LocationRp}
    location.EntityData.Leafs["sdr-name"] = types.YLeaf{"SdrName", location.SdrName}
    return &(location.EntityData)
}

