// This module contains a collection of YANG definitions
// for Cisco IOS-XR spirit-corehelper package configuration.
// 
// This module contains definitions
// for the following management objects:
//   exception: Core dump configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package spirit_corehelper_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package spirit_corehelper_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-corehelper-cfg exception}", reflect.TypeOf(Exception{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-corehelper-cfg:exception", reflect.TypeOf(Exception{}))
}

// Exception
// Core dump configuration commands
type Exception struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Container for the order of preference.
    File Exception_File
}

func (exception *Exception) GetEntityData() *types.CommonEntityData {
    exception.EntityData.YFilter = exception.YFilter
    exception.EntityData.YangName = "exception"
    exception.EntityData.BundleName = "cisco_ios_xr"
    exception.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-corehelper-cfg"
    exception.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-corehelper-cfg:exception"
    exception.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exception.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exception.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exception.EntityData.Children = make(map[string]types.YChild)
    exception.EntityData.Children["file"] = types.YChild{"File", &exception.File}
    exception.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(exception.EntityData)
}

// Exception_File
// Container for the order of preference
type Exception_File struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Preference of the dump location. The type is string.
    Choice2 interface{}

    // Preference of the dump location. The type is string.
    Choice1 interface{}

    // Preference of the dump location. The type is string.
    Choice3 interface{}
}

func (file *Exception_File) GetEntityData() *types.CommonEntityData {
    file.EntityData.YFilter = file.YFilter
    file.EntityData.YangName = "file"
    file.EntityData.BundleName = "cisco_ios_xr"
    file.EntityData.ParentYangName = "exception"
    file.EntityData.SegmentPath = "file"
    file.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    file.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    file.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    file.EntityData.Children = make(map[string]types.YChild)
    file.EntityData.Leafs = make(map[string]types.YLeaf)
    file.EntityData.Leafs["choice2"] = types.YLeaf{"Choice2", file.Choice2}
    file.EntityData.Leafs["choice1"] = types.YLeaf{"Choice1", file.Choice1}
    file.EntityData.Leafs["choice3"] = types.YLeaf{"Choice3", file.Choice3}
    return &(file.EntityData)
}

