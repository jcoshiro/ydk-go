// This module holds OBFL data.
package sysadmin_obfl

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_obfl"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-obfl obfl}", reflect.TypeOf(Obfl{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-obfl:obfl", reflect.TypeOf(Obfl{}))
}

// Obfl
type Obfl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    ObflMgr Obfl_ObflMgr

    
    ObflShow Obfl_ObflShow
}

func (obfl *Obfl) GetEntityData() *types.CommonEntityData {
    obfl.EntityData.YFilter = obfl.YFilter
    obfl.EntityData.YangName = "obfl"
    obfl.EntityData.BundleName = "cisco_ios_xr"
    obfl.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-obfl"
    obfl.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-obfl:obfl"
    obfl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    obfl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    obfl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    obfl.EntityData.Children = make(map[string]types.YChild)
    obfl.EntityData.Children["obfl_mgr"] = types.YChild{"ObflMgr", &obfl.ObflMgr}
    obfl.EntityData.Children["obfl_show"] = types.YChild{"ObflShow", &obfl.ObflShow}
    obfl.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(obfl.EntityData)
}

// Obfl_ObflMgr
type Obfl_ObflMgr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Obfl_ObflMgr_Trace.
    Trace []Obfl_ObflMgr_Trace
}

func (obflMgr *Obfl_ObflMgr) GetEntityData() *types.CommonEntityData {
    obflMgr.EntityData.YFilter = obflMgr.YFilter
    obflMgr.EntityData.YangName = "obfl_mgr"
    obflMgr.EntityData.BundleName = "cisco_ios_xr"
    obflMgr.EntityData.ParentYangName = "obfl"
    obflMgr.EntityData.SegmentPath = "obfl_mgr"
    obflMgr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    obflMgr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    obflMgr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    obflMgr.EntityData.Children = make(map[string]types.YChild)
    obflMgr.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range obflMgr.Trace {
        obflMgr.EntityData.Children[types.GetSegmentPath(&obflMgr.Trace[i])] = types.YChild{"Trace", &obflMgr.Trace[i]}
    }
    obflMgr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(obflMgr.EntityData)
}

// Obfl_ObflMgr_Trace
// show traceable processes
type Obfl_ObflMgr_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Obfl_ObflMgr_Trace_Location.
    Location []Obfl_ObflMgr_Trace_Location
}

func (trace *Obfl_ObflMgr_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "obfl_mgr"
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

// Obfl_ObflMgr_Trace_Location
type Obfl_ObflMgr_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Obfl_ObflMgr_Trace_Location_AllOptions.
    AllOptions []Obfl_ObflMgr_Trace_Location_AllOptions
}

func (location *Obfl_ObflMgr_Trace_Location) GetEntityData() *types.CommonEntityData {
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

// Obfl_ObflMgr_Trace_Location_AllOptions
type Obfl_ObflMgr_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Obfl_ObflMgr_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []Obfl_ObflMgr_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Obfl_ObflMgr_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
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

// Obfl_ObflMgr_Trace_Location_AllOptions_TraceBlocks
type Obfl_ObflMgr_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Obfl_ObflMgr_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
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

// Obfl_ObflShow
type Obfl_ObflShow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Obfl_ObflShow_Trace.
    Trace []Obfl_ObflShow_Trace
}

func (obflShow *Obfl_ObflShow) GetEntityData() *types.CommonEntityData {
    obflShow.EntityData.YFilter = obflShow.YFilter
    obflShow.EntityData.YangName = "obfl_show"
    obflShow.EntityData.BundleName = "cisco_ios_xr"
    obflShow.EntityData.ParentYangName = "obfl"
    obflShow.EntityData.SegmentPath = "obfl_show"
    obflShow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    obflShow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    obflShow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    obflShow.EntityData.Children = make(map[string]types.YChild)
    obflShow.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range obflShow.Trace {
        obflShow.EntityData.Children[types.GetSegmentPath(&obflShow.Trace[i])] = types.YChild{"Trace", &obflShow.Trace[i]}
    }
    obflShow.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(obflShow.EntityData)
}

// Obfl_ObflShow_Trace
// show traceable processes
type Obfl_ObflShow_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Obfl_ObflShow_Trace_Location.
    Location []Obfl_ObflShow_Trace_Location
}

func (trace *Obfl_ObflShow_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "obfl_show"
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

// Obfl_ObflShow_Trace_Location
type Obfl_ObflShow_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Obfl_ObflShow_Trace_Location_AllOptions.
    AllOptions []Obfl_ObflShow_Trace_Location_AllOptions
}

func (location *Obfl_ObflShow_Trace_Location) GetEntityData() *types.CommonEntityData {
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

// Obfl_ObflShow_Trace_Location_AllOptions
type Obfl_ObflShow_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Obfl_ObflShow_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []Obfl_ObflShow_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Obfl_ObflShow_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
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

// Obfl_ObflShow_Trace_Location_AllOptions_TraceBlocks
type Obfl_ObflShow_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Obfl_ObflShow_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
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

