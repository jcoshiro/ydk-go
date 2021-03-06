// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-dumper package configuration.
// 
// This module contains definitions
// for the following management objects:
//   exception: Core dump configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_dumper_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_dumper_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-dumper-cfg exception}", reflect.TypeOf(Exception{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-dumper-cfg:exception", reflect.TypeOf(Exception{}))
}

// Copy represents Copy
type Copy string

const (
    // copy not configured
    Copy_default_ Copy = "default"

    // Dump to local memory: for time critical
    // processes
    Copy_copy Copy = "copy"
)

// Context represents Context
type Context string

const (
    // context not configured
    Context_default_ Context = "default"

    // Dump context info only: Overrides other options
    // except for 'no-core'
    Context_context Context = "context"
)

// Sharedmemory represents Sharedmemory
type Sharedmemory string

const (
    // shared memory not configured
    Sharedmemory_default_ Sharedmemory = "default"

    // Dump shared memory of the target process
    Sharedmemory_shared_memory Sharedmemory = "shared-memory"
)

// Sparse represents Sparse
type Sparse string

const (
    // sparse not configured
    Sparse_default_ Sparse = "default"

    // Dump memory relevant to stack trace only: for
    // time critical processes
    Sparse_sparse Sparse = "sparse"
)

// Nocore represents Nocore
type Nocore string

const (
    // no core not configured
    Nocore_default_ Nocore = "default"

    // Disable core dump for the target process:
    // Overrides other options
    Nocore_no_core Nocore = "no-core"
)

// Packetmemory represents Packetmemory
type Packetmemory string

const (
    // packet memory not configured
    Packetmemory_default_ Packetmemory = "default"

    // Dump packet memory of the target process
    Packetmemory_packet_memory Packetmemory = "packet-memory"
)

// Skipcpuinfo represents Skipcpuinfo
type Skipcpuinfo string

const (
    // skip-cpu-info not configured
    Skipcpuinfo_default_ Skipcpuinfo = "default"

    // Skip CPU usage snapshot: for time critical
    // processes
    Skipcpuinfo_skip_cpu_info Skipcpuinfo = "skip-cpu-info"
)

// Mainmemory represents Mainmemory
type Mainmemory string

const (
    // main memory not configured
    Mainmemory_default_ Mainmemory = "default"

    // Dump main memory of the target process
    Mainmemory_main_memory Mainmemory = "main-memory"
)

// Exception
// Core dump configuration commands
type Exception struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify 'true' to enable sparse core dump, 'false' to disable sparse core
    // dump. The type is bool.
    Sparse interface{}

    // Disable core file verification. The type is bool.
    CoreVerification interface{}

    // Only print out stack trace and create no core file beyond this size. The
    // type is interface{} with range: 1..4095.
    CoreSize interface{}

    // Enable kernel debugger. The type is interface{}.
    KernelDebugger interface{}

    // Specify 'true' to dump packet memory for all process, 'false' to disable
    // dump of packet memory. The type is bool.
    PacketMemory interface{}

    // Switch to sparse core dump at this size. The type is interface{} with
    // range: 1..4095.
    SparseSize interface{}

    // Give up core dump if specified free memory can not be secured. The type is
    // interface{} with range: 3..40. Units are percentage.
    MemoryThreshold interface{}

    // Preference of the dump location.
    Choice1 Exception_Choice1

    // Preference of the dump location.
    Choice3 Exception_Choice3

    // Specify per process configuration.
    ProcessNames Exception_ProcessNames

    // Preference of the dump location.
    Choice2 Exception_Choice2
}

func (exception *Exception) GetEntityData() *types.CommonEntityData {
    exception.EntityData.YFilter = exception.YFilter
    exception.EntityData.YangName = "exception"
    exception.EntityData.BundleName = "cisco_ios_xr"
    exception.EntityData.ParentYangName = "Cisco-IOS-XR-infra-dumper-cfg"
    exception.EntityData.SegmentPath = "Cisco-IOS-XR-infra-dumper-cfg:exception"
    exception.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exception.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exception.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exception.EntityData.Children = make(map[string]types.YChild)
    exception.EntityData.Children["choice1"] = types.YChild{"Choice1", &exception.Choice1}
    exception.EntityData.Children["choice3"] = types.YChild{"Choice3", &exception.Choice3}
    exception.EntityData.Children["process-names"] = types.YChild{"ProcessNames", &exception.ProcessNames}
    exception.EntityData.Children["choice2"] = types.YChild{"Choice2", &exception.Choice2}
    exception.EntityData.Leafs = make(map[string]types.YLeaf)
    exception.EntityData.Leafs["sparse"] = types.YLeaf{"Sparse", exception.Sparse}
    exception.EntityData.Leafs["core-verification"] = types.YLeaf{"CoreVerification", exception.CoreVerification}
    exception.EntityData.Leafs["core-size"] = types.YLeaf{"CoreSize", exception.CoreSize}
    exception.EntityData.Leafs["kernel-debugger"] = types.YLeaf{"KernelDebugger", exception.KernelDebugger}
    exception.EntityData.Leafs["packet-memory"] = types.YLeaf{"PacketMemory", exception.PacketMemory}
    exception.EntityData.Leafs["sparse-size"] = types.YLeaf{"SparseSize", exception.SparseSize}
    exception.EntityData.Leafs["memory-threshold"] = types.YLeaf{"MemoryThreshold", exception.MemoryThreshold}
    return &(exception.EntityData)
}

// Exception_Choice1
// Preference of the dump location
type Exception_Choice1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify 'true' to compress core files dumped on this path, 'false' to not
    // compress. The type is bool.
    Compress interface{}

    // Lower limit.  This is required if Filename is specified. The type is
    // interface{} with range: 0..4.
    LowerLimit interface{}

    // Higher limit.  This is required if Filename is specified. The type is
    // interface{} with range: 5..64.
    HigherLimit interface{}

    // Protocol and directory. The type is string.
    FilePath interface{}

    // Dump filename. The type is string.
    Filename interface{}
}

func (choice1 *Exception_Choice1) GetEntityData() *types.CommonEntityData {
    choice1.EntityData.YFilter = choice1.YFilter
    choice1.EntityData.YangName = "choice1"
    choice1.EntityData.BundleName = "cisco_ios_xr"
    choice1.EntityData.ParentYangName = "exception"
    choice1.EntityData.SegmentPath = "choice1"
    choice1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    choice1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    choice1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    choice1.EntityData.Children = make(map[string]types.YChild)
    choice1.EntityData.Leafs = make(map[string]types.YLeaf)
    choice1.EntityData.Leafs["compress"] = types.YLeaf{"Compress", choice1.Compress}
    choice1.EntityData.Leafs["lower-limit"] = types.YLeaf{"LowerLimit", choice1.LowerLimit}
    choice1.EntityData.Leafs["higher-limit"] = types.YLeaf{"HigherLimit", choice1.HigherLimit}
    choice1.EntityData.Leafs["file-path"] = types.YLeaf{"FilePath", choice1.FilePath}
    choice1.EntityData.Leafs["filename"] = types.YLeaf{"Filename", choice1.Filename}
    return &(choice1.EntityData)
}

// Exception_Choice3
// Preference of the dump location
type Exception_Choice3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify 'true' to compress core files dumped on this path, 'false' to not
    // compress. The type is bool.
    Compress interface{}

    // Lower limit.  This is required if Filename is specified. The type is
    // interface{} with range: 0..4.
    LowerLimit interface{}

    // Higher limit.  This is required if Filename is specified. The type is
    // interface{} with range: 5..64.
    HigherLimit interface{}

    // Protocol and directory. The type is string.
    FilePath interface{}

    // Dump filename. The type is string.
    Filename interface{}
}

func (choice3 *Exception_Choice3) GetEntityData() *types.CommonEntityData {
    choice3.EntityData.YFilter = choice3.YFilter
    choice3.EntityData.YangName = "choice3"
    choice3.EntityData.BundleName = "cisco_ios_xr"
    choice3.EntityData.ParentYangName = "exception"
    choice3.EntityData.SegmentPath = "choice3"
    choice3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    choice3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    choice3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    choice3.EntityData.Children = make(map[string]types.YChild)
    choice3.EntityData.Leafs = make(map[string]types.YLeaf)
    choice3.EntityData.Leafs["compress"] = types.YLeaf{"Compress", choice3.Compress}
    choice3.EntityData.Leafs["lower-limit"] = types.YLeaf{"LowerLimit", choice3.LowerLimit}
    choice3.EntityData.Leafs["higher-limit"] = types.YLeaf{"HigherLimit", choice3.HigherLimit}
    choice3.EntityData.Leafs["file-path"] = types.YLeaf{"FilePath", choice3.FilePath}
    choice3.EntityData.Leafs["filename"] = types.YLeaf{"Filename", choice3.Filename}
    return &(choice3.EntityData)
}

// Exception_ProcessNames
// Specify per process configuration
type Exception_ProcessNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify per process configuration. The type is slice of
    // Exception_ProcessNames_ProcessName.
    ProcessName []Exception_ProcessNames_ProcessName
}

func (processNames *Exception_ProcessNames) GetEntityData() *types.CommonEntityData {
    processNames.EntityData.YFilter = processNames.YFilter
    processNames.EntityData.YangName = "process-names"
    processNames.EntityData.BundleName = "cisco_ios_xr"
    processNames.EntityData.ParentYangName = "exception"
    processNames.EntityData.SegmentPath = "process-names"
    processNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNames.EntityData.Children = make(map[string]types.YChild)
    processNames.EntityData.Children["process-name"] = types.YChild{"ProcessName", nil}
    for i := range processNames.ProcessName {
        processNames.EntityData.Children[types.GetSegmentPath(&processNames.ProcessName[i])] = types.YChild{"ProcessName", &processNames.ProcessName[i]}
    }
    processNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(processNames.EntityData)
}

// Exception_ProcessNames_ProcessName
// Specify per process configuration
type Exception_ProcessNames_ProcessName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specify per process configuration. The type is
    // string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Processname interface{}

    // Specify per process core option.
    CoreOption Exception_ProcessNames_ProcessName_CoreOption
}

func (processName *Exception_ProcessNames_ProcessName) GetEntityData() *types.CommonEntityData {
    processName.EntityData.YFilter = processName.YFilter
    processName.EntityData.YangName = "process-name"
    processName.EntityData.BundleName = "cisco_ios_xr"
    processName.EntityData.ParentYangName = "process-names"
    processName.EntityData.SegmentPath = "process-name" + "[processname='" + fmt.Sprintf("%v", processName.Processname) + "']"
    processName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processName.EntityData.Children = make(map[string]types.YChild)
    processName.EntityData.Children["core-option"] = types.YChild{"CoreOption", &processName.CoreOption}
    processName.EntityData.Leafs = make(map[string]types.YLeaf)
    processName.EntityData.Leafs["processname"] = types.YLeaf{"Processname", processName.Processname}
    return &(processName.EntityData)
}

// Exception_ProcessNames_ProcessName_CoreOption
// Specify per process core option
type Exception_ProcessNames_ProcessName_CoreOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dump main memory of the target process. The type is Mainmemory. The default
    // value is default.
    MainMemoryval interface{}

    // Dump shared memory of the target process. The type is Sharedmemory. The
    // default value is default.
    SharedMemoryval interface{}

    // Dump packet memory of the target process. The type is Packetmemory. The
    // default value is default.
    PacketMemoryval interface{}

    // Dump to local memory: for time critical processes. The type is Copy. The
    // default value is default.
    Copyval interface{}

    // Dump memory relevant to stack trace only: for time critical processes. The
    // type is Sparse. The default value is default.
    Sparseval interface{}

    // Skip CPU usage snapshot: for time critical processes. The type is
    // Skipcpuinfo. The default value is default.
    Skipcpuinfoval interface{}

    // Dump context info only: Overrides other options except for 'no-core'. The
    // type is Context. The default value is default.
    Contextval interface{}

    // Disable core dump for the target process: Overrides other options. The type
    // is Nocore. The default value is default.
    Nocoreval interface{}
}

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetEntityData() *types.CommonEntityData {
    coreOption.EntityData.YFilter = coreOption.YFilter
    coreOption.EntityData.YangName = "core-option"
    coreOption.EntityData.BundleName = "cisco_ios_xr"
    coreOption.EntityData.ParentYangName = "process-name"
    coreOption.EntityData.SegmentPath = "core-option"
    coreOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coreOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coreOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coreOption.EntityData.Children = make(map[string]types.YChild)
    coreOption.EntityData.Leafs = make(map[string]types.YLeaf)
    coreOption.EntityData.Leafs["main-memoryval"] = types.YLeaf{"MainMemoryval", coreOption.MainMemoryval}
    coreOption.EntityData.Leafs["shared-memoryval"] = types.YLeaf{"SharedMemoryval", coreOption.SharedMemoryval}
    coreOption.EntityData.Leafs["packet-memoryval"] = types.YLeaf{"PacketMemoryval", coreOption.PacketMemoryval}
    coreOption.EntityData.Leafs["copyval"] = types.YLeaf{"Copyval", coreOption.Copyval}
    coreOption.EntityData.Leafs["sparseval"] = types.YLeaf{"Sparseval", coreOption.Sparseval}
    coreOption.EntityData.Leafs["skipcpuinfoval"] = types.YLeaf{"Skipcpuinfoval", coreOption.Skipcpuinfoval}
    coreOption.EntityData.Leafs["contextval"] = types.YLeaf{"Contextval", coreOption.Contextval}
    coreOption.EntityData.Leafs["nocoreval"] = types.YLeaf{"Nocoreval", coreOption.Nocoreval}
    return &(coreOption.EntityData)
}

// Exception_Choice2
// Preference of the dump location
type Exception_Choice2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify 'true' to compress core files dumped on this path, 'false' to not
    // compress. The type is bool.
    Compress interface{}

    // Lower limit.  This is required if Filename is specified. The type is
    // interface{} with range: 0..4.
    LowerLimit interface{}

    // Higher limit.  This is required if Filename is specified. The type is
    // interface{} with range: 5..64.
    HigherLimit interface{}

    // Protocol and directory. The type is string.
    FilePath interface{}

    // Dump filename. The type is string.
    Filename interface{}
}

func (choice2 *Exception_Choice2) GetEntityData() *types.CommonEntityData {
    choice2.EntityData.YFilter = choice2.YFilter
    choice2.EntityData.YangName = "choice2"
    choice2.EntityData.BundleName = "cisco_ios_xr"
    choice2.EntityData.ParentYangName = "exception"
    choice2.EntityData.SegmentPath = "choice2"
    choice2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    choice2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    choice2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    choice2.EntityData.Children = make(map[string]types.YChild)
    choice2.EntityData.Leafs = make(map[string]types.YLeaf)
    choice2.EntityData.Leafs["compress"] = types.YLeaf{"Compress", choice2.Compress}
    choice2.EntityData.Leafs["lower-limit"] = types.YLeaf{"LowerLimit", choice2.LowerLimit}
    choice2.EntityData.Leafs["higher-limit"] = types.YLeaf{"HigherLimit", choice2.HigherLimit}
    choice2.EntityData.Leafs["file-path"] = types.YLeaf{"FilePath", choice2.FilePath}
    choice2.EntityData.Leafs["filename"] = types.YLeaf{"Filename", choice2.Filename}
    return &(choice2.EntityData)
}

