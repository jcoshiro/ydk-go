package sysadmin_fpd_infra_cli_fpdserv_ctrace

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_fpd_infra_cli_fpdserv_ctrace"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpdserv-ctrace fpdserv}", reflect.TypeOf(Fpdserv{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpdserv-ctrace:fpdserv", reflect.TypeOf(Fpdserv{}))
}

// Fpdserv
type Fpdserv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Fpdserv_Trace.
    Trace []Fpdserv_Trace
}

func (fpdserv *Fpdserv) GetEntityData() *types.CommonEntityData {
    fpdserv.EntityData.YFilter = fpdserv.YFilter
    fpdserv.EntityData.YangName = "fpdserv"
    fpdserv.EntityData.BundleName = "cisco_ios_xr"
    fpdserv.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpdserv-ctrace"
    fpdserv.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpdserv-ctrace:fpdserv"
    fpdserv.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpdserv.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpdserv.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpdserv.EntityData.Children = make(map[string]types.YChild)
    fpdserv.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range fpdserv.Trace {
        fpdserv.EntityData.Children[types.GetSegmentPath(&fpdserv.Trace[i])] = types.YChild{"Trace", &fpdserv.Trace[i]}
    }
    fpdserv.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fpdserv.EntityData)
}

// Fpdserv_Trace
// show traceable processes
type Fpdserv_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Fpdserv_Trace_Location.
    Location []Fpdserv_Trace_Location
}

func (trace *Fpdserv_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "fpdserv"
    trace.EntityData.SegmentPath = "trace" + "[buffer='" + fmt.Sprintf("%v", trace.Buffer) + "']"
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = make(map[string]types.YChild)
    trace.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range trace.Location {
        trace.EntityData.Children[types.GetSegmentPath(&trace.Location[i])] = types.YChild{"Location", &trace.Location[i]}
    }
    trace.EntityData.Leafs = make(map[string]types.YLeaf)
    trace.EntityData.Leafs["buffer"] = types.YLeaf{"Buffer", trace.Buffer}
    return &(trace.EntityData)
}

// Fpdserv_Trace_Location
type Fpdserv_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Fpdserv_Trace_Location_AllOptions.
    AllOptions []Fpdserv_Trace_Location_AllOptions
}

func (location *Fpdserv_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + "[location_name='" + fmt.Sprintf("%v", location.LocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["all-options"] = types.YChild{"AllOptions", nil}
    for i := range location.AllOptions {
        location.EntityData.Children[types.GetSegmentPath(&location.AllOptions[i])] = types.YChild{"AllOptions", &location.AllOptions[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location_name"] = types.YLeaf{"LocationName", location.LocationName}
    return &(location.EntityData)
}

// Fpdserv_Trace_Location_AllOptions
type Fpdserv_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Fpdserv_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []Fpdserv_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Fpdserv_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + "[option='" + fmt.Sprintf("%v", allOptions.Option) + "']"
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = make(map[string]types.YChild)
    allOptions.EntityData.Children["trace-blocks"] = types.YChild{"TraceBlocks", nil}
    for i := range allOptions.TraceBlocks {
        allOptions.EntityData.Children[types.GetSegmentPath(&allOptions.TraceBlocks[i])] = types.YChild{"TraceBlocks", &allOptions.TraceBlocks[i]}
    }
    allOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    allOptions.EntityData.Leafs["option"] = types.YLeaf{"Option", allOptions.Option}
    return &(allOptions.EntityData)
}

// Fpdserv_Trace_Location_AllOptions_TraceBlocks
type Fpdserv_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Fpdserv_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks"
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = make(map[string]types.YChild)
    traceBlocks.EntityData.Leafs = make(map[string]types.YLeaf)
    traceBlocks.EntityData.Leafs["data"] = types.YLeaf{"Data", traceBlocks.Data}
    return &(traceBlocks.EntityData)
}

