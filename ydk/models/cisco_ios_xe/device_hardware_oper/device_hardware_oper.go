// This module contains a collection of YANG definitions
// for Device Hardware operational data.
// Copyright (c) 2017 by Cisco Systems, Inc.
// All rights reserved.
package device_hardware_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package device_hardware_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-device-hardware-oper device-hardware-data}", reflect.TypeOf(DeviceHardwareData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-device-hardware-oper:device-hardware-data", reflect.TypeOf(DeviceHardwareData{}))
}

// HwType represents The broad type of hardware device
type HwType string

const (
    // Unknown Hardware Type
    HwType_hw_type_unknown HwType = "hw-type-unknown"

    // Chassis
    HwType_hw_type_chassis HwType = "hw-type-chassis"

    // Central Processing Unit
    HwType_hw_type_cpu HwType = "hw-type-cpu"

    // Dynamic Random Access Memory
    HwType_hw_type_dram HwType = "hw-type-dram"

    // Flash
    HwType_hw_type_flash HwType = "hw-type-flash"

    // embedded Memory Controller
    HwType_hw_type_emmc HwType = "hw-type-emmc"

    // SD-Card
    HwType_hw_type_sdcard HwType = "hw-type-sdcard"

    // Universal Serial Bus
    HwType_hw_type_usb HwType = "hw-type-usb"

    // Pluggable interface model
    HwType_hw_type_pim HwType = "hw-type-pim"

    // Transceiver
    HwType_hw_type_transceiver HwType = "hw-type-transceiver"

    // Fan tray
    HwType_hw_type_fantray HwType = "hw-type-fantray"

    // Power Entry Module
    HwType_hw_type_pem HwType = "hw-type-pem"
)

// AlarmSeverity represents Alarm severity
type AlarmSeverity string

const (
    // Critical Alarms
    AlarmSeverity_alarm_severity_critical AlarmSeverity = "alarm-severity-critical"

    // Major Alarms
    AlarmSeverity_alarm_severity_major AlarmSeverity = "alarm-severity-major"

    // Minor Alarms
    AlarmSeverity_alarm_severity_minor AlarmSeverity = "alarm-severity-minor"
)

// DeviceHardwareData
// Device Hardware
type DeviceHardwareData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The device hardware model.
    DeviceHardware DeviceHardwareData_DeviceHardware
}

func (deviceHardwareData *DeviceHardwareData) GetEntityData() *types.CommonEntityData {
    deviceHardwareData.EntityData.YFilter = deviceHardwareData.YFilter
    deviceHardwareData.EntityData.YangName = "device-hardware-data"
    deviceHardwareData.EntityData.BundleName = "cisco_ios_xe"
    deviceHardwareData.EntityData.ParentYangName = "Cisco-IOS-XE-device-hardware-oper"
    deviceHardwareData.EntityData.SegmentPath = "Cisco-IOS-XE-device-hardware-oper:device-hardware-data"
    deviceHardwareData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    deviceHardwareData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    deviceHardwareData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    deviceHardwareData.EntityData.Children = make(map[string]types.YChild)
    deviceHardwareData.EntityData.Children["device-hardware"] = types.YChild{"DeviceHardware", &deviceHardwareData.DeviceHardware}
    deviceHardwareData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(deviceHardwareData.EntityData)
}

// DeviceHardwareData_DeviceHardware
// The device hardware model
// This type is a presence type.
type DeviceHardwareData_DeviceHardware struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All the inventory in the hardware. The type is slice of
    // DeviceHardwareData_DeviceHardware_DeviceInventory.
    DeviceInventory []DeviceHardwareData_DeviceHardware_DeviceInventory

    // The current alarms. The type is slice of
    // DeviceHardwareData_DeviceHardware_DeviceAlarm.
    DeviceAlarm []DeviceHardwareData_DeviceHardware_DeviceAlarm

    // The current device system data.
    DeviceSystemData DeviceHardwareData_DeviceHardware_DeviceSystemData
}

func (deviceHardware *DeviceHardwareData_DeviceHardware) GetEntityData() *types.CommonEntityData {
    deviceHardware.EntityData.YFilter = deviceHardware.YFilter
    deviceHardware.EntityData.YangName = "device-hardware"
    deviceHardware.EntityData.BundleName = "cisco_ios_xe"
    deviceHardware.EntityData.ParentYangName = "device-hardware-data"
    deviceHardware.EntityData.SegmentPath = "device-hardware"
    deviceHardware.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    deviceHardware.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    deviceHardware.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    deviceHardware.EntityData.Children = make(map[string]types.YChild)
    deviceHardware.EntityData.Children["device-inventory"] = types.YChild{"DeviceInventory", nil}
    for i := range deviceHardware.DeviceInventory {
        deviceHardware.EntityData.Children[types.GetSegmentPath(&deviceHardware.DeviceInventory[i])] = types.YChild{"DeviceInventory", &deviceHardware.DeviceInventory[i]}
    }
    deviceHardware.EntityData.Children["device-alarm"] = types.YChild{"DeviceAlarm", nil}
    for i := range deviceHardware.DeviceAlarm {
        deviceHardware.EntityData.Children[types.GetSegmentPath(&deviceHardware.DeviceAlarm[i])] = types.YChild{"DeviceAlarm", &deviceHardware.DeviceAlarm[i]}
    }
    deviceHardware.EntityData.Children["device-system-data"] = types.YChild{"DeviceSystemData", &deviceHardware.DeviceSystemData}
    deviceHardware.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(deviceHardware.EntityData)
}

// DeviceHardwareData_DeviceHardware_DeviceInventory
// All the inventory in the hardware
type DeviceHardwareData_DeviceHardware_DeviceInventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Type of the hardware being represented. The type
    // is HwType.
    HwType interface{}

    // This attribute is a key. The physical index of the inventory item. The type
    // is interface{} with range: 0..4294967295.
    HwDevIndex interface{}

    // Version of this device. The type is string.
    Version interface{}

    // The part number of this device. This  is only valid if the device is a
    // field replaceable unit. The type is string.
    PartNumber interface{}

    // The serial number of the device. This is only valid if the device is
    // individually trackable. The type is string.
    SerialNumber interface{}

    // Description of the device. The type is string.
    HwDescription interface{}
}

func (deviceInventory *DeviceHardwareData_DeviceHardware_DeviceInventory) GetEntityData() *types.CommonEntityData {
    deviceInventory.EntityData.YFilter = deviceInventory.YFilter
    deviceInventory.EntityData.YangName = "device-inventory"
    deviceInventory.EntityData.BundleName = "cisco_ios_xe"
    deviceInventory.EntityData.ParentYangName = "device-hardware"
    deviceInventory.EntityData.SegmentPath = "device-inventory" + "[hw-type='" + fmt.Sprintf("%v", deviceInventory.HwType) + "']" + "[hw-dev-index='" + fmt.Sprintf("%v", deviceInventory.HwDevIndex) + "']"
    deviceInventory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    deviceInventory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    deviceInventory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    deviceInventory.EntityData.Children = make(map[string]types.YChild)
    deviceInventory.EntityData.Leafs = make(map[string]types.YLeaf)
    deviceInventory.EntityData.Leafs["hw-type"] = types.YLeaf{"HwType", deviceInventory.HwType}
    deviceInventory.EntityData.Leafs["hw-dev-index"] = types.YLeaf{"HwDevIndex", deviceInventory.HwDevIndex}
    deviceInventory.EntityData.Leafs["version"] = types.YLeaf{"Version", deviceInventory.Version}
    deviceInventory.EntityData.Leafs["part-number"] = types.YLeaf{"PartNumber", deviceInventory.PartNumber}
    deviceInventory.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", deviceInventory.SerialNumber}
    deviceInventory.EntityData.Leafs["hw-description"] = types.YLeaf{"HwDescription", deviceInventory.HwDescription}
    return &(deviceInventory.EntityData)
}

// DeviceHardwareData_DeviceHardware_DeviceAlarm
// The current alarms
type DeviceHardwareData_DeviceHardware_DeviceAlarm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Alarm Identifier. The type is interface{} with
    // range: 0..4294967295.
    AlarmId interface{}

    // This attribute is a key. The instance of this alarm. This corresponds to
    // the  entity index. The type is interface{} with range: 0..4294967295.
    AlarmInstance interface{}

    // The name of the alarm. The type is string.
    AlarmName interface{}

    // The catagory (or severity) of the alarm that is being asserted. The type is
    // AlarmSeverity.
    AlarmCategory interface{}

    // Time the alarm was raised. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    AlarmTime interface{}

    // Description of the alarm. The type is string.
    AlarmDescription interface{}
}

func (deviceAlarm *DeviceHardwareData_DeviceHardware_DeviceAlarm) GetEntityData() *types.CommonEntityData {
    deviceAlarm.EntityData.YFilter = deviceAlarm.YFilter
    deviceAlarm.EntityData.YangName = "device-alarm"
    deviceAlarm.EntityData.BundleName = "cisco_ios_xe"
    deviceAlarm.EntityData.ParentYangName = "device-hardware"
    deviceAlarm.EntityData.SegmentPath = "device-alarm" + "[alarm-id='" + fmt.Sprintf("%v", deviceAlarm.AlarmId) + "']" + "[alarm-instance='" + fmt.Sprintf("%v", deviceAlarm.AlarmInstance) + "']"
    deviceAlarm.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    deviceAlarm.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    deviceAlarm.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    deviceAlarm.EntityData.Children = make(map[string]types.YChild)
    deviceAlarm.EntityData.Leafs = make(map[string]types.YLeaf)
    deviceAlarm.EntityData.Leafs["alarm-id"] = types.YLeaf{"AlarmId", deviceAlarm.AlarmId}
    deviceAlarm.EntityData.Leafs["alarm-instance"] = types.YLeaf{"AlarmInstance", deviceAlarm.AlarmInstance}
    deviceAlarm.EntityData.Leafs["alarm-name"] = types.YLeaf{"AlarmName", deviceAlarm.AlarmName}
    deviceAlarm.EntityData.Leafs["alarm-category"] = types.YLeaf{"AlarmCategory", deviceAlarm.AlarmCategory}
    deviceAlarm.EntityData.Leafs["alarm-time"] = types.YLeaf{"AlarmTime", deviceAlarm.AlarmTime}
    deviceAlarm.EntityData.Leafs["alarm-description"] = types.YLeaf{"AlarmDescription", deviceAlarm.AlarmDescription}
    return &(deviceAlarm.EntityData)
}

// DeviceHardwareData_DeviceHardware_DeviceSystemData
// The current device system data
// This type is a presence type.
type DeviceHardwareData_DeviceHardware_DeviceSystemData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current time on device in UTC. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    CurrentTime interface{}

    // This timestamp indicates the time that the system was last restarted.  The
    // value is the timestamp in seconds relative to the Unix Epoch (Jan 1, 1970
    // 00:00:00 UTC). The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    BootTime interface{}

    // Software version. The type is string.
    SoftwareVersion interface{}

    // Rommon version. The type is string.
    RommonVersion interface{}

    // Last reboot reason. The type is string.
    LastRebootReason interface{}
}

func (deviceSystemData *DeviceHardwareData_DeviceHardware_DeviceSystemData) GetEntityData() *types.CommonEntityData {
    deviceSystemData.EntityData.YFilter = deviceSystemData.YFilter
    deviceSystemData.EntityData.YangName = "device-system-data"
    deviceSystemData.EntityData.BundleName = "cisco_ios_xe"
    deviceSystemData.EntityData.ParentYangName = "device-hardware"
    deviceSystemData.EntityData.SegmentPath = "device-system-data"
    deviceSystemData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    deviceSystemData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    deviceSystemData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    deviceSystemData.EntityData.Children = make(map[string]types.YChild)
    deviceSystemData.EntityData.Leafs = make(map[string]types.YLeaf)
    deviceSystemData.EntityData.Leafs["current-time"] = types.YLeaf{"CurrentTime", deviceSystemData.CurrentTime}
    deviceSystemData.EntityData.Leafs["boot-time"] = types.YLeaf{"BootTime", deviceSystemData.BootTime}
    deviceSystemData.EntityData.Leafs["software-version"] = types.YLeaf{"SoftwareVersion", deviceSystemData.SoftwareVersion}
    deviceSystemData.EntityData.Leafs["rommon-version"] = types.YLeaf{"RommonVersion", deviceSystemData.RommonVersion}
    deviceSystemData.EntityData.Leafs["last-reboot-reason"] = types.YLeaf{"LastRebootReason", deviceSystemData.LastRebootReason}
    return &(deviceSystemData.EntityData)
}

