// This module contains a collection of YANG definitions for
// monitoring memory in a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package memory_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package memory_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-memory-oper memory-statistics}", reflect.TypeOf(MemoryStatistics{}))
    ydk.RegisterEntity("Cisco-IOS-XE-memory-oper:memory-statistics", reflect.TypeOf(MemoryStatistics{}))
}

// MemoryStatistics
// Data nodes for All Memory Pool Statistics.
type MemoryStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of software memory pools in the system. The type is slice of
    // MemoryStatistics_MemoryStatistic.
    MemoryStatistic []MemoryStatistics_MemoryStatistic
}

func (memoryStatistics *MemoryStatistics) GetEntityData() *types.CommonEntityData {
    memoryStatistics.EntityData.YFilter = memoryStatistics.YFilter
    memoryStatistics.EntityData.YangName = "memory-statistics"
    memoryStatistics.EntityData.BundleName = "cisco_ios_xe"
    memoryStatistics.EntityData.ParentYangName = "Cisco-IOS-XE-memory-oper"
    memoryStatistics.EntityData.SegmentPath = "Cisco-IOS-XE-memory-oper:memory-statistics"
    memoryStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    memoryStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    memoryStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    memoryStatistics.EntityData.Children = make(map[string]types.YChild)
    memoryStatistics.EntityData.Children["memory-statistic"] = types.YChild{"MemoryStatistic", nil}
    for i := range memoryStatistics.MemoryStatistic {
        memoryStatistics.EntityData.Children[types.GetSegmentPath(&memoryStatistics.MemoryStatistic[i])] = types.YChild{"MemoryStatistic", &memoryStatistics.MemoryStatistic[i]}
    }
    memoryStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memoryStatistics.EntityData)
}

// MemoryStatistics_MemoryStatistic
// The list of software memory pools in the system.
type MemoryStatistics_MemoryStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the memory pool. The type is string.
    Name interface{}

    // Total memory in the pool (bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    TotalMemory interface{}

    // Total used memory in the pool (bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    UsedMemory interface{}

    // Total free memory in the pool (bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    FreeMemory interface{}

    // Historical lowest memory usage (bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    LowestUsage interface{}

    // Historical highest memory usage (bytes). The type is interface{} with
    // range: 0..18446744073709551615. Units are bytes.
    HighestUsage interface{}
}

func (memoryStatistic *MemoryStatistics_MemoryStatistic) GetEntityData() *types.CommonEntityData {
    memoryStatistic.EntityData.YFilter = memoryStatistic.YFilter
    memoryStatistic.EntityData.YangName = "memory-statistic"
    memoryStatistic.EntityData.BundleName = "cisco_ios_xe"
    memoryStatistic.EntityData.ParentYangName = "memory-statistics"
    memoryStatistic.EntityData.SegmentPath = "memory-statistic" + "[name='" + fmt.Sprintf("%v", memoryStatistic.Name) + "']"
    memoryStatistic.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    memoryStatistic.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    memoryStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    memoryStatistic.EntityData.Children = make(map[string]types.YChild)
    memoryStatistic.EntityData.Leafs = make(map[string]types.YLeaf)
    memoryStatistic.EntityData.Leafs["name"] = types.YLeaf{"Name", memoryStatistic.Name}
    memoryStatistic.EntityData.Leafs["total-memory"] = types.YLeaf{"TotalMemory", memoryStatistic.TotalMemory}
    memoryStatistic.EntityData.Leafs["used-memory"] = types.YLeaf{"UsedMemory", memoryStatistic.UsedMemory}
    memoryStatistic.EntityData.Leafs["free-memory"] = types.YLeaf{"FreeMemory", memoryStatistic.FreeMemory}
    memoryStatistic.EntityData.Leafs["lowest-usage"] = types.YLeaf{"LowestUsage", memoryStatistic.LowestUsage}
    memoryStatistic.EntityData.Leafs["highest-usage"] = types.YLeaf{"HighestUsage", memoryStatistic.HighestUsage}
    return &(memoryStatistic.EntityData)
}

