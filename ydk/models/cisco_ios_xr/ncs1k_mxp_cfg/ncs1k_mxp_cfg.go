// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-mxp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: NCS1k HW module config
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs1k_mxp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1k_mxp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1k-mxp-cfg hardware-module}", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1k-mxp-cfg:hardware-module", reflect.TypeOf(HardwareModule{}))
}

// TrunkDataRate represents Trunk data rate
type TrunkDataRate string

const (
    // HundredGig
    TrunkDataRate_hundred_gig TrunkDataRate = "hundred-gig"

    // TwoHundredGig
    TrunkDataRate_two_hundred_gig TrunkDataRate = "two-hundred-gig"

    // TwoHundredFiftyGig
    TrunkDataRate_two_hundred_fifty_gig TrunkDataRate = "two-hundred-fifty-gig"
)

// ClientDataRate represents Client data rate
type ClientDataRate string

const (
    // TenGig
    ClientDataRate_ten_gig ClientDataRate = "ten-gig"

    // FortyGig
    ClientDataRate_forty_gig ClientDataRate = "forty-gig"

    // HundredGig
    ClientDataRate_hundred_gig ClientDataRate = "hundred-gig"

    // TenAndHundredGig
    ClientDataRate_ten_and_hundred_gig ClientDataRate = "ten-and-hundred-gig"
)

// Fec represents Fec
type Fec string

const (
    // SoftDecision7
    Fec_sd7 Fec = "sd7"

    // SoftDecision20
    Fec_sd20 Fec = "sd20"
)

// HardwareModule
// NCS1k HW module config
type HardwareModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is slice of HardwareModule_Node.
    Node []HardwareModule_Node
}

func (hardwareModule *HardwareModule) GetEntityData() *types.CommonEntityData {
    hardwareModule.EntityData.YFilter = hardwareModule.YFilter
    hardwareModule.EntityData.YangName = "hardware-module"
    hardwareModule.EntityData.BundleName = "cisco_ios_xr"
    hardwareModule.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1k-mxp-cfg"
    hardwareModule.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1k-mxp-cfg:hardware-module"
    hardwareModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModule.EntityData.Children = make(map[string]types.YChild)
    hardwareModule.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range hardwareModule.Node {
        hardwareModule.EntityData.Children[types.GetSegmentPath(&hardwareModule.Node[i])] = types.YChild{"Node", &hardwareModule.Node[i]}
    }
    hardwareModule.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareModule.EntityData)
}

// HardwareModule_Node
// Node
type HardwareModule_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fully qualified line card specification. The type
    // is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Location interface{}

    // Slice to be Provisioned. The type is slice of HardwareModule_Node_Slice.
    Slice []HardwareModule_Node_Slice
}

func (node *HardwareModule_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "hardware-module"
    node.EntityData.SegmentPath = "node" + "[location='" + fmt.Sprintf("%v", node.Location) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["slice"] = types.YChild{"Slice", nil}
    for i := range node.Slice {
        node.EntityData.Children[types.GetSegmentPath(&node.Slice[i])] = types.YChild{"Slice", &node.Slice[i]}
    }
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["location"] = types.YLeaf{"Location", node.Location}
    return &(node.EntityData)
}

// HardwareModule_Node_Slice
// Slice to be Provisioned
type HardwareModule_Node_Slice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set Slice. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SliceId interface{}

    // Drop LLDP Packets. The type is bool.
    Lldp interface{}

    // Data rates & FEC.
    Values HardwareModule_Node_Slice_Values
}

func (slice *HardwareModule_Node_Slice) GetEntityData() *types.CommonEntityData {
    slice.EntityData.YFilter = slice.YFilter
    slice.EntityData.YangName = "slice"
    slice.EntityData.BundleName = "cisco_ios_xr"
    slice.EntityData.ParentYangName = "node"
    slice.EntityData.SegmentPath = "slice" + "[slice-id='" + fmt.Sprintf("%v", slice.SliceId) + "']"
    slice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slice.EntityData.Children = make(map[string]types.YChild)
    slice.EntityData.Children["values"] = types.YChild{"Values", &slice.Values}
    slice.EntityData.Leafs = make(map[string]types.YLeaf)
    slice.EntityData.Leafs["slice-id"] = types.YLeaf{"SliceId", slice.SliceId}
    slice.EntityData.Leafs["lldp"] = types.YLeaf{"Lldp", slice.Lldp}
    return &(slice.EntityData)
}

// HardwareModule_Node_Slice_Values
// Data rates & FEC
type HardwareModule_Node_Slice_Values struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client Rate. The type is ClientDataRate.
    ClientRate interface{}

    // TrunkRate. The type is TrunkDataRate.
    TrunkRate interface{}

    // FEC. The type is Fec.
    Fec interface{}

    // Encrypted. The type is bool. The default value is false.
    Encrypted interface{}
}

func (values *HardwareModule_Node_Slice_Values) GetEntityData() *types.CommonEntityData {
    values.EntityData.YFilter = values.YFilter
    values.EntityData.YangName = "values"
    values.EntityData.BundleName = "cisco_ios_xr"
    values.EntityData.ParentYangName = "slice"
    values.EntityData.SegmentPath = "values"
    values.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    values.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    values.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    values.EntityData.Children = make(map[string]types.YChild)
    values.EntityData.Leafs = make(map[string]types.YLeaf)
    values.EntityData.Leafs["client-rate"] = types.YLeaf{"ClientRate", values.ClientRate}
    values.EntityData.Leafs["trunk-rate"] = types.YLeaf{"TrunkRate", values.TrunkRate}
    values.EntityData.Leafs["fec"] = types.YLeaf{"Fec", values.Fec}
    values.EntityData.Leafs["encrypted"] = types.YLeaf{"Encrypted", values.Encrypted}
    return &(values.EntityData)
}

