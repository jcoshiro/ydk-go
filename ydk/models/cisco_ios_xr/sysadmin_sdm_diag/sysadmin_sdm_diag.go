package sysadmin_sdm_diag

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_sdm_diag"))
    ydk.RegisterEntity("{http://cisco.com/calvados/diagnostics diagnostic}", reflect.TypeOf(Diagnostic{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-sdm-diag:diagnostic", reflect.TypeOf(Diagnostic{}))
}

// Diagnostic
type Diagnostic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Monitor Diagnostic_Monitor

    
    Schedule Diagnostic_Schedule

    
    Status Diagnostic_Status

    
    DiagStart Diagnostic_DiagStart

    
    DiagStop Diagnostic_DiagStop

    
    Content Diagnostic_Content

    
    Result Diagnostic_Result
}

func (diagnostic *Diagnostic) GetEntityData() *types.CommonEntityData {
    diagnostic.EntityData.YFilter = diagnostic.YFilter
    diagnostic.EntityData.YangName = "diagnostic"
    diagnostic.EntityData.BundleName = "cisco_ios_xr"
    diagnostic.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-sdm-diag"
    diagnostic.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-sdm-diag:diagnostic"
    diagnostic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diagnostic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diagnostic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diagnostic.EntityData.Children = make(map[string]types.YChild)
    diagnostic.EntityData.Children["monitor"] = types.YChild{"Monitor", &diagnostic.Monitor}
    diagnostic.EntityData.Children["schedule"] = types.YChild{"Schedule", &diagnostic.Schedule}
    diagnostic.EntityData.Children["status"] = types.YChild{"Status", &diagnostic.Status}
    diagnostic.EntityData.Children["diag_start"] = types.YChild{"DiagStart", &diagnostic.DiagStart}
    diagnostic.EntityData.Children["diag_stop"] = types.YChild{"DiagStop", &diagnostic.DiagStop}
    diagnostic.EntityData.Children["content"] = types.YChild{"Content", &diagnostic.Content}
    diagnostic.EntityData.Children["result"] = types.YChild{"Result", &diagnostic.Result}
    diagnostic.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diagnostic.EntityData)
}

// Diagnostic_Monitor
type Diagnostic_Monitor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rejected Diagnostic_Monitor_Rejected

    
    Interval Diagnostic_Monitor_Interval

    
    Threshold Diagnostic_Monitor_Threshold
}

func (monitor *Diagnostic_Monitor) GetEntityData() *types.CommonEntityData {
    monitor.EntityData.YFilter = monitor.YFilter
    monitor.EntityData.YangName = "monitor"
    monitor.EntityData.BundleName = "cisco_ios_xr"
    monitor.EntityData.ParentYangName = "diagnostic"
    monitor.EntityData.SegmentPath = "monitor"
    monitor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitor.EntityData.Children = make(map[string]types.YChild)
    monitor.EntityData.Children["rejected"] = types.YChild{"Rejected", &monitor.Rejected}
    monitor.EntityData.Children["interval"] = types.YChild{"Interval", &monitor.Interval}
    monitor.EntityData.Children["threshold"] = types.YChild{"Threshold", &monitor.Threshold}
    monitor.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(monitor.EntityData)
}

// Diagnostic_Monitor_Rejected
type Diagnostic_Monitor_Rejected struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diagnostic_Monitor_Rejected_Location.
    Location []Diagnostic_Monitor_Rejected_Location
}

func (rejected *Diagnostic_Monitor_Rejected) GetEntityData() *types.CommonEntityData {
    rejected.EntityData.YFilter = rejected.YFilter
    rejected.EntityData.YangName = "rejected"
    rejected.EntityData.BundleName = "cisco_ios_xr"
    rejected.EntityData.ParentYangName = "monitor"
    rejected.EntityData.SegmentPath = "rejected"
    rejected.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rejected.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rejected.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rejected.EntityData.Children = make(map[string]types.YChild)
    rejected.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rejected.Location {
        rejected.EntityData.Children[types.GetSegmentPath(&rejected.Location[i])] = types.YChild{"Location", &rejected.Location[i]}
    }
    rejected.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rejected.EntityData)
}

// Diagnostic_Monitor_Rejected_Location
type Diagnostic_Monitor_Rejected_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Loc interface{}

    // The type is slice of Diagnostic_Monitor_Rejected_Location_Test.
    Test []Diagnostic_Monitor_Rejected_Location_Test
}

func (location *Diagnostic_Monitor_Rejected_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rejected"
    location.EntityData.SegmentPath = "location" + "[loc='" + fmt.Sprintf("%v", location.Loc) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["test"] = types.YChild{"Test", nil}
    for i := range location.Test {
        location.EntityData.Children[types.GetSegmentPath(&location.Test[i])] = types.YChild{"Test", &location.Test[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["loc"] = types.YLeaf{"Loc", location.Loc}
    return &(location.EntityData)
}

// Diagnostic_Monitor_Rejected_Location_Test
type Diagnostic_Monitor_Rejected_Location_Test struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    TestId interface{}

    // This attribute is a key. The type is Run.
    Run interface{}
}

func (test *Diagnostic_Monitor_Rejected_Location_Test) GetEntityData() *types.CommonEntityData {
    test.EntityData.YFilter = test.YFilter
    test.EntityData.YangName = "test"
    test.EntityData.BundleName = "cisco_ios_xr"
    test.EntityData.ParentYangName = "location"
    test.EntityData.SegmentPath = "test" + "[test_id='" + fmt.Sprintf("%v", test.TestId) + "']" + "[run='" + fmt.Sprintf("%v", test.Run) + "']"
    test.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    test.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    test.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    test.EntityData.Children = make(map[string]types.YChild)
    test.EntityData.Leafs = make(map[string]types.YLeaf)
    test.EntityData.Leafs["test_id"] = types.YLeaf{"TestId", test.TestId}
    test.EntityData.Leafs["run"] = types.YLeaf{"Run", test.Run}
    return &(test.EntityData)
}

// Diagnostic_Monitor_Rejected_Location_Test_Run
type Diagnostic_Monitor_Rejected_Location_Test_Run string

const (
    Diagnostic_Monitor_Rejected_Location_Test_Run_disable Diagnostic_Monitor_Rejected_Location_Test_Run = "disable"
)

// Diagnostic_Monitor_Interval
type Diagnostic_Monitor_Interval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diagnostic_Monitor_Interval_Location.
    Location []Diagnostic_Monitor_Interval_Location
}

func (interval *Diagnostic_Monitor_Interval) GetEntityData() *types.CommonEntityData {
    interval.EntityData.YFilter = interval.YFilter
    interval.EntityData.YangName = "interval"
    interval.EntityData.BundleName = "cisco_ios_xr"
    interval.EntityData.ParentYangName = "monitor"
    interval.EntityData.SegmentPath = "interval"
    interval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interval.EntityData.Children = make(map[string]types.YChild)
    interval.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range interval.Location {
        interval.EntityData.Children[types.GetSegmentPath(&interval.Location[i])] = types.YChild{"Location", &interval.Location[i]}
    }
    interval.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interval.EntityData)
}

// Diagnostic_Monitor_Interval_Location
type Diagnostic_Monitor_Interval_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Loc interface{}

    // The type is slice of Diagnostic_Monitor_Interval_Location_Test.
    Test []Diagnostic_Monitor_Interval_Location_Test
}

func (location *Diagnostic_Monitor_Interval_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "interval"
    location.EntityData.SegmentPath = "location" + "[loc='" + fmt.Sprintf("%v", location.Loc) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["test"] = types.YChild{"Test", nil}
    for i := range location.Test {
        location.EntityData.Children[types.GetSegmentPath(&location.Test[i])] = types.YChild{"Test", &location.Test[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["loc"] = types.YLeaf{"Loc", location.Loc}
    return &(location.EntityData)
}

// Diagnostic_Monitor_Interval_Location_Test
type Diagnostic_Monitor_Interval_Location_Test struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    TestId interface{}

    // The type is interface{} with range: 0..20. This attribute is mandatory.
    Days interface{}

    // The type is string with pattern:
    // b'([01]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]'. This attribute is mandatory.
    Time interface{}
}

func (test *Diagnostic_Monitor_Interval_Location_Test) GetEntityData() *types.CommonEntityData {
    test.EntityData.YFilter = test.YFilter
    test.EntityData.YangName = "test"
    test.EntityData.BundleName = "cisco_ios_xr"
    test.EntityData.ParentYangName = "location"
    test.EntityData.SegmentPath = "test" + "[test_id='" + fmt.Sprintf("%v", test.TestId) + "']"
    test.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    test.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    test.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    test.EntityData.Children = make(map[string]types.YChild)
    test.EntityData.Leafs = make(map[string]types.YLeaf)
    test.EntityData.Leafs["test_id"] = types.YLeaf{"TestId", test.TestId}
    test.EntityData.Leafs["days"] = types.YLeaf{"Days", test.Days}
    test.EntityData.Leafs["time"] = types.YLeaf{"Time", test.Time}
    return &(test.EntityData)
}

// Diagnostic_Monitor_Threshold
type Diagnostic_Monitor_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diagnostic_Monitor_Threshold_Location.
    Location []Diagnostic_Monitor_Threshold_Location
}

func (threshold *Diagnostic_Monitor_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "monitor"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range threshold.Location {
        threshold.EntityData.Children[types.GetSegmentPath(&threshold.Location[i])] = types.YChild{"Location", &threshold.Location[i]}
    }
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(threshold.EntityData)
}

// Diagnostic_Monitor_Threshold_Location
type Diagnostic_Monitor_Threshold_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Loc interface{}

    // The type is slice of Diagnostic_Monitor_Threshold_Location_Test.
    Test []Diagnostic_Monitor_Threshold_Location_Test
}

func (location *Diagnostic_Monitor_Threshold_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "threshold"
    location.EntityData.SegmentPath = "location" + "[loc='" + fmt.Sprintf("%v", location.Loc) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["test"] = types.YChild{"Test", nil}
    for i := range location.Test {
        location.EntityData.Children[types.GetSegmentPath(&location.Test[i])] = types.YChild{"Test", &location.Test[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["loc"] = types.YLeaf{"Loc", location.Loc}
    return &(location.EntityData)
}

// Diagnostic_Monitor_Threshold_Location_Test
type Diagnostic_Monitor_Threshold_Location_Test struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    TestId interface{}

    // The type is interface{} with range: 1..99. This attribute is mandatory.
    FailureCount interface{}
}

func (test *Diagnostic_Monitor_Threshold_Location_Test) GetEntityData() *types.CommonEntityData {
    test.EntityData.YFilter = test.YFilter
    test.EntityData.YangName = "test"
    test.EntityData.BundleName = "cisco_ios_xr"
    test.EntityData.ParentYangName = "location"
    test.EntityData.SegmentPath = "test" + "[test_id='" + fmt.Sprintf("%v", test.TestId) + "']"
    test.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    test.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    test.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    test.EntityData.Children = make(map[string]types.YChild)
    test.EntityData.Leafs = make(map[string]types.YLeaf)
    test.EntityData.Leafs["test_id"] = types.YLeaf{"TestId", test.TestId}
    test.EntityData.Leafs["failure-count"] = types.YLeaf{"FailureCount", test.FailureCount}
    return &(test.EntityData)
}

// Diagnostic_Schedule
type Diagnostic_Schedule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Start Diagnostic_Schedule_Start
}

func (schedule *Diagnostic_Schedule) GetEntityData() *types.CommonEntityData {
    schedule.EntityData.YFilter = schedule.YFilter
    schedule.EntityData.YangName = "schedule"
    schedule.EntityData.BundleName = "cisco_ios_xr"
    schedule.EntityData.ParentYangName = "diagnostic"
    schedule.EntityData.SegmentPath = "schedule"
    schedule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    schedule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    schedule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    schedule.EntityData.Children = make(map[string]types.YChild)
    schedule.EntityData.Children["start"] = types.YChild{"Start", &schedule.Start}
    schedule.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(schedule.EntityData)
}

// Diagnostic_Schedule_Start
type Diagnostic_Schedule_Start struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diagnostic_Schedule_Start_Location.
    Location []Diagnostic_Schedule_Start_Location
}

func (start *Diagnostic_Schedule_Start) GetEntityData() *types.CommonEntityData {
    start.EntityData.YFilter = start.YFilter
    start.EntityData.YangName = "start"
    start.EntityData.BundleName = "cisco_ios_xr"
    start.EntityData.ParentYangName = "schedule"
    start.EntityData.SegmentPath = "start"
    start.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    start.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    start.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    start.EntityData.Children = make(map[string]types.YChild)
    start.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range start.Location {
        start.EntityData.Children[types.GetSegmentPath(&start.Location[i])] = types.YChild{"Location", &start.Location[i]}
    }
    start.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(start.EntityData)
}

// Diagnostic_Schedule_Start_Location
type Diagnostic_Schedule_Start_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Loc interface{}

    // The type is slice of Diagnostic_Schedule_Start_Location_Test.
    Test []Diagnostic_Schedule_Start_Location_Test
}

func (location *Diagnostic_Schedule_Start_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "start"
    location.EntityData.SegmentPath = "location" + "[loc='" + fmt.Sprintf("%v", location.Loc) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["test"] = types.YChild{"Test", nil}
    for i := range location.Test {
        location.EntityData.Children[types.GetSegmentPath(&location.Test[i])] = types.YChild{"Test", &location.Test[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["loc"] = types.YLeaf{"Loc", location.Loc}
    return &(location.EntityData)
}

// Diagnostic_Schedule_Start_Location_Test
type Diagnostic_Schedule_Start_Location_Test struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    TestId interface{}

    // The type is slice of Diagnostic_Schedule_Start_Location_Test_Daily.
    Daily []Diagnostic_Schedule_Start_Location_Test_Daily

    // The type is slice of Diagnostic_Schedule_Start_Location_Test_On.
    On []Diagnostic_Schedule_Start_Location_Test_On

    // The type is slice of Diagnostic_Schedule_Start_Location_Test_Weekly.
    Weekly []Diagnostic_Schedule_Start_Location_Test_Weekly
}

func (test *Diagnostic_Schedule_Start_Location_Test) GetEntityData() *types.CommonEntityData {
    test.EntityData.YFilter = test.YFilter
    test.EntityData.YangName = "test"
    test.EntityData.BundleName = "cisco_ios_xr"
    test.EntityData.ParentYangName = "location"
    test.EntityData.SegmentPath = "test" + "[test_id='" + fmt.Sprintf("%v", test.TestId) + "']"
    test.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    test.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    test.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    test.EntityData.Children = make(map[string]types.YChild)
    test.EntityData.Children["daily"] = types.YChild{"Daily", nil}
    for i := range test.Daily {
        test.EntityData.Children[types.GetSegmentPath(&test.Daily[i])] = types.YChild{"Daily", &test.Daily[i]}
    }
    test.EntityData.Children["on"] = types.YChild{"On", nil}
    for i := range test.On {
        test.EntityData.Children[types.GetSegmentPath(&test.On[i])] = types.YChild{"On", &test.On[i]}
    }
    test.EntityData.Children["weekly"] = types.YChild{"Weekly", nil}
    for i := range test.Weekly {
        test.EntityData.Children[types.GetSegmentPath(&test.Weekly[i])] = types.YChild{"Weekly", &test.Weekly[i]}
    }
    test.EntityData.Leafs = make(map[string]types.YLeaf)
    test.EntityData.Leafs["test_id"] = types.YLeaf{"TestId", test.TestId}
    return &(test.EntityData)
}

// Diagnostic_Schedule_Start_Location_Test_Daily
type Diagnostic_Schedule_Start_Location_Test_Daily struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'([01]?[0-9]|2[0-3]):[0-5][0-9]'.
    HourMin interface{}
}

func (daily *Diagnostic_Schedule_Start_Location_Test_Daily) GetEntityData() *types.CommonEntityData {
    daily.EntityData.YFilter = daily.YFilter
    daily.EntityData.YangName = "daily"
    daily.EntityData.BundleName = "cisco_ios_xr"
    daily.EntityData.ParentYangName = "test"
    daily.EntityData.SegmentPath = "daily" + "[hour_min='" + fmt.Sprintf("%v", daily.HourMin) + "']"
    daily.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    daily.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    daily.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    daily.EntityData.Children = make(map[string]types.YChild)
    daily.EntityData.Leafs = make(map[string]types.YLeaf)
    daily.EntityData.Leafs["hour_min"] = types.YLeaf{"HourMin", daily.HourMin}
    return &(daily.EntityData)
}

// Diagnostic_Schedule_Start_Location_Test_On
type Diagnostic_Schedule_Start_Location_Test_On struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is Month.
    Month interface{}

    // This attribute is a key. The type is interface{} with range: 1..31.
    DayOfMonth interface{}

    // This attribute is a key. The type is interface{} with range: 2013..2099.
    Year interface{}

    // This attribute is a key. The type is string with pattern:
    // b'([01]?[0-9]|2[0-3]):[0-5][0-9]'.
    HourMin interface{}
}

func (on *Diagnostic_Schedule_Start_Location_Test_On) GetEntityData() *types.CommonEntityData {
    on.EntityData.YFilter = on.YFilter
    on.EntityData.YangName = "on"
    on.EntityData.BundleName = "cisco_ios_xr"
    on.EntityData.ParentYangName = "test"
    on.EntityData.SegmentPath = "on" + "[month='" + fmt.Sprintf("%v", on.Month) + "']" + "[day_of_month='" + fmt.Sprintf("%v", on.DayOfMonth) + "']" + "[year='" + fmt.Sprintf("%v", on.Year) + "']" + "[hour_min='" + fmt.Sprintf("%v", on.HourMin) + "']"
    on.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    on.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    on.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    on.EntityData.Children = make(map[string]types.YChild)
    on.EntityData.Leafs = make(map[string]types.YLeaf)
    on.EntityData.Leafs["month"] = types.YLeaf{"Month", on.Month}
    on.EntityData.Leafs["day_of_month"] = types.YLeaf{"DayOfMonth", on.DayOfMonth}
    on.EntityData.Leafs["year"] = types.YLeaf{"Year", on.Year}
    on.EntityData.Leafs["hour_min"] = types.YLeaf{"HourMin", on.HourMin}
    return &(on.EntityData)
}

// Diagnostic_Schedule_Start_Location_Test_On_Month
type Diagnostic_Schedule_Start_Location_Test_On_Month string

const (
    Diagnostic_Schedule_Start_Location_Test_On_Month_JAN Diagnostic_Schedule_Start_Location_Test_On_Month = "JAN"

    Diagnostic_Schedule_Start_Location_Test_On_Month_FEB Diagnostic_Schedule_Start_Location_Test_On_Month = "FEB"

    Diagnostic_Schedule_Start_Location_Test_On_Month_MAR Diagnostic_Schedule_Start_Location_Test_On_Month = "MAR"

    Diagnostic_Schedule_Start_Location_Test_On_Month_APR Diagnostic_Schedule_Start_Location_Test_On_Month = "APR"

    Diagnostic_Schedule_Start_Location_Test_On_Month_MAY Diagnostic_Schedule_Start_Location_Test_On_Month = "MAY"

    Diagnostic_Schedule_Start_Location_Test_On_Month_JUN Diagnostic_Schedule_Start_Location_Test_On_Month = "JUN"

    Diagnostic_Schedule_Start_Location_Test_On_Month_JUL Diagnostic_Schedule_Start_Location_Test_On_Month = "JUL"

    Diagnostic_Schedule_Start_Location_Test_On_Month_AUG Diagnostic_Schedule_Start_Location_Test_On_Month = "AUG"

    Diagnostic_Schedule_Start_Location_Test_On_Month_SEP Diagnostic_Schedule_Start_Location_Test_On_Month = "SEP"

    Diagnostic_Schedule_Start_Location_Test_On_Month_OCT Diagnostic_Schedule_Start_Location_Test_On_Month = "OCT"

    Diagnostic_Schedule_Start_Location_Test_On_Month_NOV Diagnostic_Schedule_Start_Location_Test_On_Month = "NOV"

    Diagnostic_Schedule_Start_Location_Test_On_Month_DEC Diagnostic_Schedule_Start_Location_Test_On_Month = "DEC"
)

// Diagnostic_Schedule_Start_Location_Test_Weekly
type Diagnostic_Schedule_Start_Location_Test_Weekly struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is Daysofweek.
    Daysofweek interface{}

    // This attribute is a key. The type is string with pattern:
    // b'([01]?[0-9]|2[0-3]):[0-5][0-9]'.
    HourMin interface{}
}

func (weekly *Diagnostic_Schedule_Start_Location_Test_Weekly) GetEntityData() *types.CommonEntityData {
    weekly.EntityData.YFilter = weekly.YFilter
    weekly.EntityData.YangName = "weekly"
    weekly.EntityData.BundleName = "cisco_ios_xr"
    weekly.EntityData.ParentYangName = "test"
    weekly.EntityData.SegmentPath = "weekly" + "[daysofweek='" + fmt.Sprintf("%v", weekly.Daysofweek) + "']" + "[hour_min='" + fmt.Sprintf("%v", weekly.HourMin) + "']"
    weekly.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    weekly.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    weekly.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    weekly.EntityData.Children = make(map[string]types.YChild)
    weekly.EntityData.Leafs = make(map[string]types.YLeaf)
    weekly.EntityData.Leafs["daysofweek"] = types.YLeaf{"Daysofweek", weekly.Daysofweek}
    weekly.EntityData.Leafs["hour_min"] = types.YLeaf{"HourMin", weekly.HourMin}
    return &(weekly.EntityData)
}

// Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek
type Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek string

const (
    Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek_SUN Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek = "SUN"

    Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek_MON Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek = "MON"

    Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek_TUE Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek = "TUE"

    Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek_WED Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek = "WED"

    Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek_THR Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek = "THR"

    Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek_FRI Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek = "FRI"

    Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek_SAT Diagnostic_Schedule_Start_Location_Test_Weekly_Daysofweek = "SAT"
)

// Diagnostic_Status
type Diagnostic_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diagnostic_Status_LocationIndex.
    LocationIndex []Diagnostic_Status_LocationIndex
}

func (status *Diagnostic_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xr"
    status.EntityData.ParentYangName = "diagnostic"
    status.EntityData.SegmentPath = "status"
    status.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    status.EntityData.Children = make(map[string]types.YChild)
    status.EntityData.Children["location_index"] = types.YChild{"LocationIndex", nil}
    for i := range status.LocationIndex {
        status.EntityData.Children[types.GetSegmentPath(&status.LocationIndex[i])] = types.YChild{"LocationIndex", &status.LocationIndex[i]}
    }
    status.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(status.EntityData)
}

// Diagnostic_Status_LocationIndex
type Diagnostic_Status_LocationIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    DataIdx interface{}

    // The type is string.
    Description interface{}

    // The type is string.
    CurrRunningTstRunby interface{}
}

func (locationIndex *Diagnostic_Status_LocationIndex) GetEntityData() *types.CommonEntityData {
    locationIndex.EntityData.YFilter = locationIndex.YFilter
    locationIndex.EntityData.YangName = "location_index"
    locationIndex.EntityData.BundleName = "cisco_ios_xr"
    locationIndex.EntityData.ParentYangName = "status"
    locationIndex.EntityData.SegmentPath = "location_index" + "[data_idx='" + fmt.Sprintf("%v", locationIndex.DataIdx) + "']"
    locationIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationIndex.EntityData.Children = make(map[string]types.YChild)
    locationIndex.EntityData.Leafs = make(map[string]types.YLeaf)
    locationIndex.EntityData.Leafs["data_idx"] = types.YLeaf{"DataIdx", locationIndex.DataIdx}
    locationIndex.EntityData.Leafs["description"] = types.YLeaf{"Description", locationIndex.Description}
    locationIndex.EntityData.Leafs["curr_running_tst_runby"] = types.YLeaf{"CurrRunningTstRunby", locationIndex.CurrRunningTstRunby}
    return &(locationIndex.EntityData)
}

// Diagnostic_DiagStart
type Diagnostic_DiagStart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diagnostic_DiagStart_Location.
    Location []Diagnostic_DiagStart_Location
}

func (diagStart *Diagnostic_DiagStart) GetEntityData() *types.CommonEntityData {
    diagStart.EntityData.YFilter = diagStart.YFilter
    diagStart.EntityData.YangName = "diag_start"
    diagStart.EntityData.BundleName = "cisco_ios_xr"
    diagStart.EntityData.ParentYangName = "diagnostic"
    diagStart.EntityData.SegmentPath = "diag_start"
    diagStart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diagStart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diagStart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diagStart.EntityData.Children = make(map[string]types.YChild)
    diagStart.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range diagStart.Location {
        diagStart.EntityData.Children[types.GetSegmentPath(&diagStart.Location[i])] = types.YChild{"Location", &diagStart.Location[i]}
    }
    diagStart.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diagStart.EntityData)
}

// Diagnostic_DiagStart_Location
type Diagnostic_DiagStart_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Loc interface{}

    // The type is string. The default value is port..
    Description interface{}

    // The type is slice of Diagnostic_DiagStart_Location_Test.
    Test []Diagnostic_DiagStart_Location_Test
}

func (location *Diagnostic_DiagStart_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "diag_start"
    location.EntityData.SegmentPath = "location" + "[loc='" + fmt.Sprintf("%v", location.Loc) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["test"] = types.YChild{"Test", nil}
    for i := range location.Test {
        location.EntityData.Children[types.GetSegmentPath(&location.Test[i])] = types.YChild{"Test", &location.Test[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["loc"] = types.YLeaf{"Loc", location.Loc}
    location.EntityData.Leafs["description"] = types.YLeaf{"Description", location.Description}
    return &(location.EntityData)
}

// Diagnostic_DiagStart_Location_Test
type Diagnostic_DiagStart_Location_Test struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    TestType interface{}

    // The type is string. The default value is Test..
    Description interface{}
}

func (test *Diagnostic_DiagStart_Location_Test) GetEntityData() *types.CommonEntityData {
    test.EntityData.YFilter = test.YFilter
    test.EntityData.YangName = "test"
    test.EntityData.BundleName = "cisco_ios_xr"
    test.EntityData.ParentYangName = "location"
    test.EntityData.SegmentPath = "test" + "[test_type='" + fmt.Sprintf("%v", test.TestType) + "']"
    test.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    test.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    test.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    test.EntityData.Children = make(map[string]types.YChild)
    test.EntityData.Leafs = make(map[string]types.YLeaf)
    test.EntityData.Leafs["test_type"] = types.YLeaf{"TestType", test.TestType}
    test.EntityData.Leafs["description"] = types.YLeaf{"Description", test.Description}
    return &(test.EntityData)
}

// Diagnostic_DiagStop
type Diagnostic_DiagStop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diagnostic_DiagStop_Location.
    Location []Diagnostic_DiagStop_Location
}

func (diagStop *Diagnostic_DiagStop) GetEntityData() *types.CommonEntityData {
    diagStop.EntityData.YFilter = diagStop.YFilter
    diagStop.EntityData.YangName = "diag_stop"
    diagStop.EntityData.BundleName = "cisco_ios_xr"
    diagStop.EntityData.ParentYangName = "diagnostic"
    diagStop.EntityData.SegmentPath = "diag_stop"
    diagStop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diagStop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diagStop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diagStop.EntityData.Children = make(map[string]types.YChild)
    diagStop.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range diagStop.Location {
        diagStop.EntityData.Children[types.GetSegmentPath(&diagStop.Location[i])] = types.YChild{"Location", &diagStop.Location[i]}
    }
    diagStop.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diagStop.EntityData)
}

// Diagnostic_DiagStop_Location
type Diagnostic_DiagStop_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Loc interface{}

    // The type is string. The default value is port..
    Description interface{}

    // The type is slice of Diagnostic_DiagStop_Location_Test.
    Test []Diagnostic_DiagStop_Location_Test
}

func (location *Diagnostic_DiagStop_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "diag_stop"
    location.EntityData.SegmentPath = "location" + "[loc='" + fmt.Sprintf("%v", location.Loc) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["test"] = types.YChild{"Test", nil}
    for i := range location.Test {
        location.EntityData.Children[types.GetSegmentPath(&location.Test[i])] = types.YChild{"Test", &location.Test[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["loc"] = types.YLeaf{"Loc", location.Loc}
    location.EntityData.Leafs["description"] = types.YLeaf{"Description", location.Description}
    return &(location.EntityData)
}

// Diagnostic_DiagStop_Location_Test
type Diagnostic_DiagStop_Location_Test struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    TestType interface{}

    // The type is string. The default value is Test..
    Description interface{}
}

func (test *Diagnostic_DiagStop_Location_Test) GetEntityData() *types.CommonEntityData {
    test.EntityData.YFilter = test.YFilter
    test.EntityData.YangName = "test"
    test.EntityData.BundleName = "cisco_ios_xr"
    test.EntityData.ParentYangName = "location"
    test.EntityData.SegmentPath = "test" + "[test_type='" + fmt.Sprintf("%v", test.TestType) + "']"
    test.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    test.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    test.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    test.EntityData.Children = make(map[string]types.YChild)
    test.EntityData.Leafs = make(map[string]types.YLeaf)
    test.EntityData.Leafs["test_type"] = types.YLeaf{"TestType", test.TestType}
    test.EntityData.Leafs["description"] = types.YLeaf{"Description", test.Description}
    return &(test.EntityData)
}

// Diagnostic_Content
type Diagnostic_Content struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diagnostic_Content_Location.
    Location []Diagnostic_Content_Location
}

func (content *Diagnostic_Content) GetEntityData() *types.CommonEntityData {
    content.EntityData.YFilter = content.YFilter
    content.EntityData.YangName = "content"
    content.EntityData.BundleName = "cisco_ios_xr"
    content.EntityData.ParentYangName = "diagnostic"
    content.EntityData.SegmentPath = "content"
    content.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    content.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    content.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    content.EntityData.Children = make(map[string]types.YChild)
    content.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range content.Location {
        content.EntityData.Children[types.GetSegmentPath(&content.Location[i])] = types.YChild{"Location", &content.Location[i]}
    }
    content.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(content.EntityData)
}

// Diagnostic_Content_Location
type Diagnostic_Content_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Loc interface{}

    // The type is string. The default value is port..
    Description interface{}

    // The type is slice of Diagnostic_Content_Location_DataList.
    DataList []Diagnostic_Content_Location_DataList
}

func (location *Diagnostic_Content_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "content"
    location.EntityData.SegmentPath = "location" + "[loc='" + fmt.Sprintf("%v", location.Loc) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["data_list"] = types.YChild{"DataList", nil}
    for i := range location.DataList {
        location.EntityData.Children[types.GetSegmentPath(&location.DataList[i])] = types.YChild{"DataList", &location.DataList[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["loc"] = types.YLeaf{"Loc", location.Loc}
    location.EntityData.Leafs["description"] = types.YLeaf{"Description", location.Description}
    return &(location.EntityData)
}

// Diagnostic_Content_Location_DataList
type Diagnostic_Content_Location_DataList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    DataIdx interface{}

    // To display the test information. The type is interface{} with range:
    // -2147483648..2147483647.
    Id interface{}

    // The type is string.
    TestName interface{}

    // The type is string.
    Attribute interface{}

    // The type is string.
    Timeinfo interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    Threshold interface{}
}

func (dataList *Diagnostic_Content_Location_DataList) GetEntityData() *types.CommonEntityData {
    dataList.EntityData.YFilter = dataList.YFilter
    dataList.EntityData.YangName = "data_list"
    dataList.EntityData.BundleName = "cisco_ios_xr"
    dataList.EntityData.ParentYangName = "location"
    dataList.EntityData.SegmentPath = "data_list" + "[data_idx='" + fmt.Sprintf("%v", dataList.DataIdx) + "']"
    dataList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataList.EntityData.Children = make(map[string]types.YChild)
    dataList.EntityData.Leafs = make(map[string]types.YLeaf)
    dataList.EntityData.Leafs["data_idx"] = types.YLeaf{"DataIdx", dataList.DataIdx}
    dataList.EntityData.Leafs["id"] = types.YLeaf{"Id", dataList.Id}
    dataList.EntityData.Leafs["test_name"] = types.YLeaf{"TestName", dataList.TestName}
    dataList.EntityData.Leafs["attribute"] = types.YLeaf{"Attribute", dataList.Attribute}
    dataList.EntityData.Leafs["timeinfo"] = types.YLeaf{"Timeinfo", dataList.Timeinfo}
    dataList.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", dataList.Threshold}
    return &(dataList.EntityData)
}

// Diagnostic_Result
type Diagnostic_Result struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Diagnostic_Result_Location.
    Location []Diagnostic_Result_Location
}

func (result *Diagnostic_Result) GetEntityData() *types.CommonEntityData {
    result.EntityData.YFilter = result.YFilter
    result.EntityData.YangName = "result"
    result.EntityData.BundleName = "cisco_ios_xr"
    result.EntityData.ParentYangName = "diagnostic"
    result.EntityData.SegmentPath = "result"
    result.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    result.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    result.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    result.EntityData.Children = make(map[string]types.YChild)
    result.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range result.Location {
        result.EntityData.Children[types.GetSegmentPath(&result.Location[i])] = types.YChild{"Location", &result.Location[i]}
    }
    result.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(result.EntityData)
}

// Diagnostic_Result_Location
type Diagnostic_Result_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Loc interface{}

    // The type is string. The default value is port..
    Description interface{}

    // The type is slice of Diagnostic_Result_Location_Test.
    Test []Diagnostic_Result_Location_Test
}

func (location *Diagnostic_Result_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "result"
    location.EntityData.SegmentPath = "location" + "[loc='" + fmt.Sprintf("%v", location.Loc) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["test"] = types.YChild{"Test", nil}
    for i := range location.Test {
        location.EntityData.Children[types.GetSegmentPath(&location.Test[i])] = types.YChild{"Test", &location.Test[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["loc"] = types.YLeaf{"Loc", location.Loc}
    location.EntityData.Leafs["description"] = types.YLeaf{"Description", location.Description}
    return &(location.EntityData)
}

// Diagnostic_Result_Location_Test
type Diagnostic_Result_Location_Test struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. test name|all. The type is string.
    TestType interface{}

    // The type is string. The default value is port..
    Description interface{}

    // The type is slice of Diagnostic_Result_Location_Test_Detail.
    Detail []Diagnostic_Result_Location_Test_Detail
}

func (test *Diagnostic_Result_Location_Test) GetEntityData() *types.CommonEntityData {
    test.EntityData.YFilter = test.YFilter
    test.EntityData.YangName = "test"
    test.EntityData.BundleName = "cisco_ios_xr"
    test.EntityData.ParentYangName = "location"
    test.EntityData.SegmentPath = "test" + "[test_type='" + fmt.Sprintf("%v", test.TestType) + "']"
    test.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    test.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    test.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    test.EntityData.Children = make(map[string]types.YChild)
    test.EntityData.Children["detail"] = types.YChild{"Detail", nil}
    for i := range test.Detail {
        test.EntityData.Children[types.GetSegmentPath(&test.Detail[i])] = types.YChild{"Detail", &test.Detail[i]}
    }
    test.EntityData.Leafs = make(map[string]types.YLeaf)
    test.EntityData.Leafs["test_type"] = types.YLeaf{"TestType", test.TestType}
    test.EntityData.Leafs["description"] = types.YLeaf{"Description", test.Description}
    return &(test.EntityData)
}

// Diagnostic_Result_Location_Test_Detail
type Diagnostic_Result_Location_Test_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Det interface{}

    // The type is slice of Diagnostic_Result_Location_Test_Detail_DataList.
    DataList []Diagnostic_Result_Location_Test_Detail_DataList
}

func (detail *Diagnostic_Result_Location_Test_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "test"
    detail.EntityData.SegmentPath = "detail" + "[det='" + fmt.Sprintf("%v", detail.Det) + "']"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["data_list"] = types.YChild{"DataList", nil}
    for i := range detail.DataList {
        detail.EntityData.Children[types.GetSegmentPath(&detail.DataList[i])] = types.YChild{"DataList", &detail.DataList[i]}
    }
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    detail.EntityData.Leafs["det"] = types.YLeaf{"Det", detail.Det}
    return &(detail.EntityData)
}

// Diagnostic_Result_Location_Test_Detail_DataList
type Diagnostic_Result_Location_Test_Detail_DataList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    DataIdx interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    DetailFlag interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    SlNo interface{}

    // The type is string.
    TestName interface{}

    // The type is string.
    TestResult interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    ErrCode interface{}

    // The type is string.
    ErrDescription interface{}

    // The type is string.
    TestType interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    HmTestCount interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    OdTestCount interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    SchedTestCount interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    RunCnt interface{}

    // The type is string.
    LtExeTime interface{}

    // The type is string.
    FtFailTime interface{}

    // The type is string.
    LtFailTime interface{}

    // The type is string.
    LtPassTime interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    TotalFailCnt interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    ConsFailCnt interface{}

    // The type is string.
    MoreInfo interface{}
}

func (dataList *Diagnostic_Result_Location_Test_Detail_DataList) GetEntityData() *types.CommonEntityData {
    dataList.EntityData.YFilter = dataList.YFilter
    dataList.EntityData.YangName = "data_list"
    dataList.EntityData.BundleName = "cisco_ios_xr"
    dataList.EntityData.ParentYangName = "detail"
    dataList.EntityData.SegmentPath = "data_list" + "[data_idx='" + fmt.Sprintf("%v", dataList.DataIdx) + "']"
    dataList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataList.EntityData.Children = make(map[string]types.YChild)
    dataList.EntityData.Leafs = make(map[string]types.YLeaf)
    dataList.EntityData.Leafs["data_idx"] = types.YLeaf{"DataIdx", dataList.DataIdx}
    dataList.EntityData.Leafs["detail_flag"] = types.YLeaf{"DetailFlag", dataList.DetailFlag}
    dataList.EntityData.Leafs["sl_no"] = types.YLeaf{"SlNo", dataList.SlNo}
    dataList.EntityData.Leafs["test_name"] = types.YLeaf{"TestName", dataList.TestName}
    dataList.EntityData.Leafs["test_result"] = types.YLeaf{"TestResult", dataList.TestResult}
    dataList.EntityData.Leafs["err_code"] = types.YLeaf{"ErrCode", dataList.ErrCode}
    dataList.EntityData.Leafs["err_description"] = types.YLeaf{"ErrDescription", dataList.ErrDescription}
    dataList.EntityData.Leafs["test_type"] = types.YLeaf{"TestType", dataList.TestType}
    dataList.EntityData.Leafs["hm_test_count"] = types.YLeaf{"HmTestCount", dataList.HmTestCount}
    dataList.EntityData.Leafs["od_test_count"] = types.YLeaf{"OdTestCount", dataList.OdTestCount}
    dataList.EntityData.Leafs["sched_test_count"] = types.YLeaf{"SchedTestCount", dataList.SchedTestCount}
    dataList.EntityData.Leafs["run_cnt"] = types.YLeaf{"RunCnt", dataList.RunCnt}
    dataList.EntityData.Leafs["lt_exe_time"] = types.YLeaf{"LtExeTime", dataList.LtExeTime}
    dataList.EntityData.Leafs["ft_fail_time"] = types.YLeaf{"FtFailTime", dataList.FtFailTime}
    dataList.EntityData.Leafs["lt_fail_time"] = types.YLeaf{"LtFailTime", dataList.LtFailTime}
    dataList.EntityData.Leafs["lt_pass_time"] = types.YLeaf{"LtPassTime", dataList.LtPassTime}
    dataList.EntityData.Leafs["total_fail_cnt"] = types.YLeaf{"TotalFailCnt", dataList.TotalFailCnt}
    dataList.EntityData.Leafs["cons_fail_cnt"] = types.YLeaf{"ConsFailCnt", dataList.ConsFailCnt}
    dataList.EntityData.Leafs["more_info"] = types.YLeaf{"MoreInfo", dataList.MoreInfo}
    return &(dataList.EntityData)
}

