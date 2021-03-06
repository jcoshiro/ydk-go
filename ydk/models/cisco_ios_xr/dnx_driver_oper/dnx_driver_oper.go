// This module contains a collection of YANG definitions
// for Cisco IOS-XR dnx-driver package operational data.
// 
// This module contains definitions
// for the following management objects:
//   fia: FIA driver operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package dnx_driver_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dnx_driver_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dnx-driver-oper fia}", reflect.TypeOf(Fia{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dnx-driver-oper:fia", reflect.TypeOf(Fia{}))
}

// AsicInitMethod represents Asic init method
type AsicInitMethod string

const (
    // asic init method unset
    AsicInitMethod_asic_init_method_unset AsicInitMethod = "asic-init-method-unset"

    // asic init method no reset
    AsicInitMethod_asic_init_method_no_reset AsicInitMethod = "asic-init-method-no-reset"

    // asic init method pon reset
    AsicInitMethod_asic_init_method_pon_reset AsicInitMethod = "asic-init-method-pon-reset"

    // asic init method pon reset on intr
    AsicInitMethod_asic_init_method_pon_reset_on_intr AsicInitMethod = "asic-init-method-pon-reset-on-intr"

    // asic init method hard reset
    AsicInitMethod_asic_init_method_hard_reset AsicInitMethod = "asic-init-method-hard-reset"

    // asic init method warmboot
    AsicInitMethod_asic_init_method_warmboot AsicInitMethod = "asic-init-method-warmboot"

    // asic init method issu wb
    AsicInitMethod_asic_init_method_issu_wb AsicInitMethod = "asic-init-method-issu-wb"

    // asic init method pci shutdown
    AsicInitMethod_asic_init_method_pci_shutdown AsicInitMethod = "asic-init-method-pci-shutdown"

    // asic init method quiesce
    AsicInitMethod_asic_init_method_quiesce AsicInitMethod = "asic-init-method-quiesce"

    // asic init method issu started
    AsicInitMethod_asic_init_method_issu_started AsicInitMethod = "asic-init-method-issu-started"

    // asic init method issu rollback
    AsicInitMethod_asic_init_method_issu_rollback AsicInitMethod = "asic-init-method-issu-rollback"

    // asic init method issu abort
    AsicInitMethod_asic_init_method_issu_abort AsicInitMethod = "asic-init-method-issu-abort"

    // asic init method slice cleanup
    AsicInitMethod_asic_init_method_slice_cleanup AsicInitMethod = "asic-init-method-slice-cleanup"

    // asic init method lc remove
    AsicInitMethod_asic_init_method_lc_remove AsicInitMethod = "asic-init-method-lc-remove"

    // asic init method node down
    AsicInitMethod_asic_init_method_node_down AsicInitMethod = "asic-init-method-node-down"

    // asic init method intr
    AsicInitMethod_asic_init_method_intr AsicInitMethod = "asic-init-method-intr"

    // asic init method board reload
    AsicInitMethod_asic_init_method_board_reload AsicInitMethod = "asic-init-method-board-reload"

    // asic init method max
    AsicInitMethod_asic_init_method_max AsicInitMethod = "asic-init-method-max"
)

// AsicAccessState represents Asic access state
type AsicAccessState string

const (
    // asic state unset
    AsicAccessState_asic_state_unset AsicAccessState = "asic-state-unset"

    // asic state none
    AsicAccessState_asic_state_none AsicAccessState = "asic-state-none"

    // asic state device off line
    AsicAccessState_asic_state_device_off_line AsicAccessState = "asic-state-device-off-line"

    // asic state device created
    AsicAccessState_asic_state_device_created AsicAccessState = "asic-state-device-created"

    // asic state device online
    AsicAccessState_asic_state_device_online AsicAccessState = "asic-state-device-online"

    // asic state warmboot
    AsicAccessState_asic_state_warmboot AsicAccessState = "asic-state-warmboot"

    // asic state de init start
    AsicAccessState_asic_state_de_init_start AsicAccessState = "asic-state-de-init-start"

    // asic state intr de init
    AsicAccessState_asic_state_intr_de_init AsicAccessState = "asic-state-intr-de-init"

    // asic state bcm detach
    AsicAccessState_asic_state_bcm_detach AsicAccessState = "asic-state-bcm-detach"

    // asic state soc de init
    AsicAccessState_asic_state_soc_de_init AsicAccessState = "asic-state-soc-de-init"

    // asic state de init done
    AsicAccessState_asic_state_de_init_done AsicAccessState = "asic-state-de-init-done"

    // asic state soc init
    AsicAccessState_asic_state_soc_init AsicAccessState = "asic-state-soc-init"

    // asic state bcm init
    AsicAccessState_asic_state_bcm_init AsicAccessState = "asic-state-bcm-init"

    // asic state intr init
    AsicAccessState_asic_state_intr_init AsicAccessState = "asic-state-intr-init"

    // asic state soc init start
    AsicAccessState_asic_state_soc_init_start AsicAccessState = "asic-state-soc-init-start"

    // asic state bcm init start
    AsicAccessState_asic_state_bcm_init_start AsicAccessState = "asic-state-bcm-init-start"

    // asic state intr init start
    AsicAccessState_asic_state_intr_init_start AsicAccessState = "asic-state-intr-init-start"

    // asic state hard reset
    AsicAccessState_asic_state_hard_reset AsicAccessState = "asic-state-hard-reset"

    // asic state normal
    AsicAccessState_asic_state_normal AsicAccessState = "asic-state-normal"

    // asic state exception
    AsicAccessState_asic_state_exception AsicAccessState = "asic-state-exception"

    // asic state hp attached
    AsicAccessState_asic_state_hp_attached AsicAccessState = "asic-state-hp-attached"

    // asic state quiesce
    AsicAccessState_asic_state_quiesce AsicAccessState = "asic-state-quiesce"

    // asic state issu started
    AsicAccessState_asic_state_issu_started AsicAccessState = "asic-state-issu-started"

    // asic state issu started nn
    AsicAccessState_asic_state_issu_started_nn AsicAccessState = "asic-state-issu-started-nn"

    // asic state issu abort
    AsicAccessState_asic_state_issu_abort AsicAccessState = "asic-state-issu-abort"

    // asic state max
    AsicAccessState_asic_state_max AsicAccessState = "asic-state-max"
)

// AsicOperState represents Asic oper state
type AsicOperState string

const (
    // asic oper unset
    AsicOperState_asic_oper_unset AsicOperState = "asic-oper-unset"

    // asic oper unknown
    AsicOperState_asic_oper_unknown AsicOperState = "asic-oper-unknown"

    // asic oper up
    AsicOperState_asic_oper_up AsicOperState = "asic-oper-up"

    // asic oper down
    AsicOperState_asic_oper_down AsicOperState = "asic-oper-down"

    // asic card down
    AsicOperState_asic_card_down AsicOperState = "asic-card-down"
)

// SliceState represents Slice state
type SliceState string

const (
    // slice oper unset
    SliceState_slice_oper_unset SliceState = "slice-oper-unset"

    // slice oper down
    SliceState_slice_oper_down SliceState = "slice-oper-down"

    // slice oper up
    SliceState_slice_oper_up SliceState = "slice-oper-up"

    // slice oper na
    SliceState_slice_oper_na SliceState = "slice-oper-na"
)

// FcMode represents Fc mode
type FcMode string

const (
    // fc mode unset
    FcMode_fc_mode_unset FcMode = "fc-mode-unset"

    // fc mode unavail
    FcMode_fc_mode_unavail FcMode = "fc-mode-unavail"

    // fc mode inband
    FcMode_fc_mode_inband FcMode = "fc-mode-inband"

    // fc mode oob
    FcMode_fc_mode_oob FcMode = "fc-mode-oob"
)

// LinkErrorState represents Link error state
type LinkErrorState string

const (
    // link error unset
    LinkErrorState_link_error_unset LinkErrorState = "link-error-unset"

    // link error none
    LinkErrorState_link_error_none LinkErrorState = "link-error-none"

    // link error shut
    LinkErrorState_link_error_shut LinkErrorState = "link-error-shut"

    // link error max
    LinkErrorState_link_error_max LinkErrorState = "link-error-max"
)

// OperState represents Oper state
type OperState string

const (
    // oper unset
    OperState_oper_unset OperState = "oper-unset"

    // oper unknown
    OperState_oper_unknown OperState = "oper-unknown"

    // oper up
    OperState_oper_up OperState = "oper-up"

    // oper down
    OperState_oper_down OperState = "oper-down"

    // card down
    OperState_card_down OperState = "card-down"
)

// AdminState represents Admin state
type AdminState string

const (
    // admin unset
    AdminState_admin_unset AdminState = "admin-unset"

    // admin up
    AdminState_admin_up AdminState = "admin-up"

    // admin down
    AdminState_admin_down AdminState = "admin-down"
)

// LinkStage represents Link stage
type LinkStage string

const (
    // link stage unset
    LinkStage_link_stage_unset LinkStage = "link-stage-unset"

    // link stage unused
    LinkStage_link_stage_unused LinkStage = "link-stage-unused"

    // link stage fia
    LinkStage_link_stage_fia LinkStage = "link-stage-fia"

    // link stage s1
    LinkStage_link_stage_s1 LinkStage = "link-stage-s1"

    // link stage s2
    LinkStage_link_stage_s2 LinkStage = "link-stage-s2"

    // link stage s3
    LinkStage_link_stage_s3 LinkStage = "link-stage-s3"

    // link stage unknown
    LinkStage_link_stage_unknown LinkStage = "link-stage-unknown"
)

// Link represents Link
type Link string

const (
    // link type unset
    Link_link_type_unset Link = "link-type-unset"

    // link type unavail
    Link_link_type_unavail Link = "link-type-unavail"

    // link type tx
    Link_link_type_tx Link = "link-type-tx"

    // link type rx
    Link_link_type_rx Link = "link-type-rx"
)

// Asic represents Asic
type Asic string

const (
    // asic unset
    Asic_asic_unset Asic = "asic-unset"

    // asic unavail
    Asic_asic_unavail Asic = "asic-unavail"

    // asic fia
    Asic_asic_fia Asic = "asic-fia"

    // asic s123
    Asic_asic_s123 Asic = "asic-s123"

    // asic s13
    Asic_asic_s13 Asic = "asic-s13"

    // asic s2
    Asic_asic_s2 Asic = "asic-s2"

    // asic b2b
    Asic_asic_b2b Asic = "asic-b2b"

    // asic type unknown
    Asic_asic_type_unknown Asic = "asic-type-unknown"
)

// Rack represents Rack
type Rack string

const (
    // rack type unset
    Rack_rack_type_unset Rack = "rack-type-unset"

    // rack type lcc
    Rack_rack_type_lcc Rack = "rack-type-lcc"

    // rack type fcc
    Rack_rack_type_fcc Rack = "rack-type-fcc"
)

// Fia
// FIA driver operational data
type Fia struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FIA driver operational data for available nodes.
    Nodes Fia_Nodes
}

func (fia *Fia) GetEntityData() *types.CommonEntityData {
    fia.EntityData.YFilter = fia.YFilter
    fia.EntityData.YangName = "fia"
    fia.EntityData.BundleName = "cisco_ios_xr"
    fia.EntityData.ParentYangName = "Cisco-IOS-XR-dnx-driver-oper"
    fia.EntityData.SegmentPath = "Cisco-IOS-XR-dnx-driver-oper:fia"
    fia.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fia.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fia.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fia.EntityData.Children = make(map[string]types.YChild)
    fia.EntityData.Children["nodes"] = types.YChild{"Nodes", &fia.Nodes}
    fia.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fia.EntityData)
}

// Fia_Nodes
// FIA driver operational data for available nodes
type Fia_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FIA operational data for a particular node. The type is slice of
    // Fia_Nodes_Node.
    Node []Fia_Nodes_Node
}

func (nodes *Fia_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "fia"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Fia_Nodes_Node
// FIA operational data for a particular node
type Fia_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // FIA link rx information.
    RxLinkInformation Fia_Nodes_Node_RxLinkInformation

    // FIA driver information.
    DriverInformation Fia_Nodes_Node_DriverInformation

    // Clear statistics information.
    ClearStatistics Fia_Nodes_Node_ClearStatistics

    // FIA link TX information.
    TxLinkInformation Fia_Nodes_Node_TxLinkInformation

    // FIA diag shell information.
    DiagShell Fia_Nodes_Node_DiagShell

    // FIA operational data of oir history.
    OirHistory Fia_Nodes_Node_OirHistory

    // FIA asic statistics information.
    AsicStatistics Fia_Nodes_Node_AsicStatistics
}

func (node *Fia_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["rx-link-information"] = types.YChild{"RxLinkInformation", &node.RxLinkInformation}
    node.EntityData.Children["driver-information"] = types.YChild{"DriverInformation", &node.DriverInformation}
    node.EntityData.Children["clear-statistics"] = types.YChild{"ClearStatistics", &node.ClearStatistics}
    node.EntityData.Children["tx-link-information"] = types.YChild{"TxLinkInformation", &node.TxLinkInformation}
    node.EntityData.Children["diag-shell"] = types.YChild{"DiagShell", &node.DiagShell}
    node.EntityData.Children["oir-history"] = types.YChild{"OirHistory", &node.OirHistory}
    node.EntityData.Children["asic-statistics"] = types.YChild{"AsicStatistics", &node.AsicStatistics}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation
// FIA link rx information
type Fia_Nodes_Node_RxLinkInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Option table for link rx information.
    LinkOptions Fia_Nodes_Node_RxLinkInformation_LinkOptions
}

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetEntityData() *types.CommonEntityData {
    rxLinkInformation.EntityData.YFilter = rxLinkInformation.YFilter
    rxLinkInformation.EntityData.YangName = "rx-link-information"
    rxLinkInformation.EntityData.BundleName = "cisco_ios_xr"
    rxLinkInformation.EntityData.ParentYangName = "node"
    rxLinkInformation.EntityData.SegmentPath = "rx-link-information"
    rxLinkInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLinkInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLinkInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLinkInformation.EntityData.Children = make(map[string]types.YChild)
    rxLinkInformation.EntityData.Children["link-options"] = types.YChild{"LinkOptions", &rxLinkInformation.LinkOptions}
    rxLinkInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rxLinkInformation.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions
// Option table for link rx information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Option : topo , flag , stats. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption.
    LinkOption []Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption
}

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetEntityData() *types.CommonEntityData {
    linkOptions.EntityData.YFilter = linkOptions.YFilter
    linkOptions.EntityData.YangName = "link-options"
    linkOptions.EntityData.BundleName = "cisco_ios_xr"
    linkOptions.EntityData.ParentYangName = "rx-link-information"
    linkOptions.EntityData.SegmentPath = "link-options"
    linkOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkOptions.EntityData.Children = make(map[string]types.YChild)
    linkOptions.EntityData.Children["link-option"] = types.YChild{"LinkOption", nil}
    for i := range linkOptions.LinkOption {
        linkOptions.EntityData.Children[types.GetSegmentPath(&linkOptions.LinkOption[i])] = types.YChild{"LinkOption", &linkOptions.LinkOption[i]}
    }
    linkOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(linkOptions.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption
// Option : topo , flag , stats
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link option. The type is string with pattern:
    // b'(flap)|(topo)'.
    Option interface{}

    // Instance table for rx information.
    RxAsicInstances Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances
}

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetEntityData() *types.CommonEntityData {
    linkOption.EntityData.YFilter = linkOption.YFilter
    linkOption.EntityData.YangName = "link-option"
    linkOption.EntityData.BundleName = "cisco_ios_xr"
    linkOption.EntityData.ParentYangName = "link-options"
    linkOption.EntityData.SegmentPath = "link-option" + "[option='" + fmt.Sprintf("%v", linkOption.Option) + "']"
    linkOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkOption.EntityData.Children = make(map[string]types.YChild)
    linkOption.EntityData.Children["rx-asic-instances"] = types.YChild{"RxAsicInstances", &linkOption.RxAsicInstances}
    linkOption.EntityData.Leafs = make(map[string]types.YLeaf)
    linkOption.EntityData.Leafs["option"] = types.YLeaf{"Option", linkOption.Option}
    return &(linkOption.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances
// Instance table for rx information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance number for rx link information. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance.
    RxAsicInstance []Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance
}

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetEntityData() *types.CommonEntityData {
    rxAsicInstances.EntityData.YFilter = rxAsicInstances.YFilter
    rxAsicInstances.EntityData.YangName = "rx-asic-instances"
    rxAsicInstances.EntityData.BundleName = "cisco_ios_xr"
    rxAsicInstances.EntityData.ParentYangName = "link-option"
    rxAsicInstances.EntityData.SegmentPath = "rx-asic-instances"
    rxAsicInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxAsicInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxAsicInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxAsicInstances.EntityData.Children = make(map[string]types.YChild)
    rxAsicInstances.EntityData.Children["rx-asic-instance"] = types.YChild{"RxAsicInstance", nil}
    for i := range rxAsicInstances.RxAsicInstance {
        rxAsicInstances.EntityData.Children[types.GetSegmentPath(&rxAsicInstances.RxAsicInstance[i])] = types.YChild{"RxAsicInstance", &rxAsicInstances.RxAsicInstance[i]}
    }
    rxAsicInstances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rxAsicInstances.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance
// Instance number for rx link information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Receive instance. The type is interface{} with
    // range: 0..255.
    Instance interface{}

    // Link table class for rx information.
    RxLinks Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks
}

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetEntityData() *types.CommonEntityData {
    rxAsicInstance.EntityData.YFilter = rxAsicInstance.YFilter
    rxAsicInstance.EntityData.YangName = "rx-asic-instance"
    rxAsicInstance.EntityData.BundleName = "cisco_ios_xr"
    rxAsicInstance.EntityData.ParentYangName = "rx-asic-instances"
    rxAsicInstance.EntityData.SegmentPath = "rx-asic-instance" + "[instance='" + fmt.Sprintf("%v", rxAsicInstance.Instance) + "']"
    rxAsicInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxAsicInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxAsicInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxAsicInstance.EntityData.Children = make(map[string]types.YChild)
    rxAsicInstance.EntityData.Children["rx-links"] = types.YChild{"RxLinks", &rxAsicInstance.RxLinks}
    rxAsicInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    rxAsicInstance.EntityData.Leafs["instance"] = types.YLeaf{"Instance", rxAsicInstance.Instance}
    return &(rxAsicInstance.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks
// Link table class for rx information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link number for rx link information. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink.
    RxLink []Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink
}

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetEntityData() *types.CommonEntityData {
    rxLinks.EntityData.YFilter = rxLinks.YFilter
    rxLinks.EntityData.YangName = "rx-links"
    rxLinks.EntityData.BundleName = "cisco_ios_xr"
    rxLinks.EntityData.ParentYangName = "rx-asic-instance"
    rxLinks.EntityData.SegmentPath = "rx-links"
    rxLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLinks.EntityData.Children = make(map[string]types.YChild)
    rxLinks.EntityData.Children["rx-link"] = types.YChild{"RxLink", nil}
    for i := range rxLinks.RxLink {
        rxLinks.EntityData.Children[types.GetSegmentPath(&rxLinks.RxLink[i])] = types.YChild{"RxLink", &rxLinks.RxLink[i]}
    }
    rxLinks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rxLinks.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink
// Link number for rx link information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start number. The type is interface{} with range: 0..47.
    StartNumber interface{}

    // End number. The type is interface{} with range: 0..47.
    EndNumber interface{}

    // RX link status option. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    StatusOption interface{}

    // Single link information. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink.
    RxLink []Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetEntityData() *types.CommonEntityData {
    rxLink.EntityData.YFilter = rxLink.YFilter
    rxLink.EntityData.YangName = "rx-link"
    rxLink.EntityData.BundleName = "cisco_ios_xr"
    rxLink.EntityData.ParentYangName = "rx-links"
    rxLink.EntityData.SegmentPath = "rx-link"
    rxLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLink.EntityData.Children = make(map[string]types.YChild)
    rxLink.EntityData.Children["rx-link"] = types.YChild{"RxLink", nil}
    for i := range rxLink.RxLink {
        rxLink.EntityData.Children[types.GetSegmentPath(&rxLink.RxLink[i])] = types.YChild{"RxLink", &rxLink.RxLink[i]}
    }
    rxLink.EntityData.Leafs = make(map[string]types.YLeaf)
    rxLink.EntityData.Leafs["start-number"] = types.YLeaf{"StartNumber", rxLink.StartNumber}
    rxLink.EntityData.Leafs["end-number"] = types.YLeaf{"EndNumber", rxLink.EndNumber}
    rxLink.EntityData.Leafs["status-option"] = types.YLeaf{"StatusOption", rxLink.StatusOption}
    return &(rxLink.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_
// Single link information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Single link. The type is interface{} with range:
    // -2147483648..2147483647.
    Link interface{}

    // speed. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Stage. The type is LinkStage.
    Stage interface{}

    // is link valid. The type is bool.
    IsLinkValid interface{}

    // is conf pending. The type is bool.
    IsConfPending interface{}

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is OperState.
    OperState interface{}

    // Error State. The type is LinkErrorState.
    ErrorState interface{}

    // flags. The type is string.
    Flags interface{}

    // flap cnt. The type is interface{} with range: 0..4294967295.
    FlapCnt interface{}

    // num admin shuts. The type is interface{} with range: 0..4294967295.
    NumAdminShuts interface{}

    // correctable errors. The type is interface{} with range:
    // 0..18446744073709551615.
    CorrectableErrors interface{}

    // uncorrectable errors. The type is interface{} with range:
    // 0..18446744073709551615.
    UncorrectableErrors interface{}

    // this link.
    ThisLink Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__ThisLink

    // far end link.
    FarEndLink Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLink

    // far end link in hw.
    FarEndLinkInHw Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLinkInHw

    // history.
    History Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__History
}

func (rxLink_ *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_) GetEntityData() *types.CommonEntityData {
    rxLink_.EntityData.YFilter = rxLink_.YFilter
    rxLink_.EntityData.YangName = "rx-link"
    rxLink_.EntityData.BundleName = "cisco_ios_xr"
    rxLink_.EntityData.ParentYangName = "rx-link"
    rxLink_.EntityData.SegmentPath = "rx-link" + "[link='" + fmt.Sprintf("%v", rxLink_.Link) + "']"
    rxLink_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLink_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLink_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLink_.EntityData.Children = make(map[string]types.YChild)
    rxLink_.EntityData.Children["this-link"] = types.YChild{"ThisLink", &rxLink_.ThisLink}
    rxLink_.EntityData.Children["far-end-link"] = types.YChild{"FarEndLink", &rxLink_.FarEndLink}
    rxLink_.EntityData.Children["far-end-link-in-hw"] = types.YChild{"FarEndLinkInHw", &rxLink_.FarEndLinkInHw}
    rxLink_.EntityData.Children["history"] = types.YChild{"History", &rxLink_.History}
    rxLink_.EntityData.Leafs = make(map[string]types.YLeaf)
    rxLink_.EntityData.Leafs["link"] = types.YLeaf{"Link", rxLink_.Link}
    rxLink_.EntityData.Leafs["speed"] = types.YLeaf{"Speed", rxLink_.Speed}
    rxLink_.EntityData.Leafs["stage"] = types.YLeaf{"Stage", rxLink_.Stage}
    rxLink_.EntityData.Leafs["is-link-valid"] = types.YLeaf{"IsLinkValid", rxLink_.IsLinkValid}
    rxLink_.EntityData.Leafs["is-conf-pending"] = types.YLeaf{"IsConfPending", rxLink_.IsConfPending}
    rxLink_.EntityData.Leafs["admin-state"] = types.YLeaf{"AdminState", rxLink_.AdminState}
    rxLink_.EntityData.Leafs["oper-state"] = types.YLeaf{"OperState", rxLink_.OperState}
    rxLink_.EntityData.Leafs["error-state"] = types.YLeaf{"ErrorState", rxLink_.ErrorState}
    rxLink_.EntityData.Leafs["flags"] = types.YLeaf{"Flags", rxLink_.Flags}
    rxLink_.EntityData.Leafs["flap-cnt"] = types.YLeaf{"FlapCnt", rxLink_.FlapCnt}
    rxLink_.EntityData.Leafs["num-admin-shuts"] = types.YLeaf{"NumAdminShuts", rxLink_.NumAdminShuts}
    rxLink_.EntityData.Leafs["correctable-errors"] = types.YLeaf{"CorrectableErrors", rxLink_.CorrectableErrors}
    rxLink_.EntityData.Leafs["uncorrectable-errors"] = types.YLeaf{"UncorrectableErrors", rxLink_.UncorrectableErrors}
    return &(rxLink_.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__ThisLink
// this link
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__ThisLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link Type. The type is Link.
    LinkType interface{}

    // Link Stage. The type is LinkStage.
    LinkStage interface{}

    // link num. The type is interface{} with range: 0..4294967295.
    LinkNum interface{}

    // phy link num. The type is interface{} with range: 0..4294967295.
    PhyLinkNum interface{}

    // asic id.
    AsicId Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__ThisLink_AsicId
}

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__ThisLink) GetEntityData() *types.CommonEntityData {
    thisLink.EntityData.YFilter = thisLink.YFilter
    thisLink.EntityData.YangName = "this-link"
    thisLink.EntityData.BundleName = "cisco_ios_xr"
    thisLink.EntityData.ParentYangName = "rx-link"
    thisLink.EntityData.SegmentPath = "this-link"
    thisLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thisLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thisLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thisLink.EntityData.Children = make(map[string]types.YChild)
    thisLink.EntityData.Children["asic-id"] = types.YChild{"AsicId", &thisLink.AsicId}
    thisLink.EntityData.Leafs = make(map[string]types.YLeaf)
    thisLink.EntityData.Leafs["link-type"] = types.YLeaf{"LinkType", thisLink.LinkType}
    thisLink.EntityData.Leafs["link-stage"] = types.YLeaf{"LinkStage", thisLink.LinkStage}
    thisLink.EntityData.Leafs["link-num"] = types.YLeaf{"LinkNum", thisLink.LinkNum}
    thisLink.EntityData.Leafs["phy-link-num"] = types.YLeaf{"PhyLinkNum", thisLink.PhyLinkNum}
    return &(thisLink.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__ThisLink_AsicId
// asic id
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__ThisLink_AsicId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__ThisLink_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "this-link"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = make(map[string]types.YChild)
    asicId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicId.EntityData.Leafs["rack-type"] = types.YLeaf{"RackType", asicId.RackType}
    asicId.EntityData.Leafs["asic-type"] = types.YLeaf{"AsicType", asicId.AsicType}
    asicId.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", asicId.RackNum}
    asicId.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", asicId.SlotNum}
    asicId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicId.AsicInstance}
    return &(asicId.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLink
// far end link
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link Type. The type is Link.
    LinkType interface{}

    // Link Stage. The type is LinkStage.
    LinkStage interface{}

    // link num. The type is interface{} with range: 0..4294967295.
    LinkNum interface{}

    // phy link num. The type is interface{} with range: 0..4294967295.
    PhyLinkNum interface{}

    // asic id.
    AsicId Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLink_AsicId
}

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLink) GetEntityData() *types.CommonEntityData {
    farEndLink.EntityData.YFilter = farEndLink.YFilter
    farEndLink.EntityData.YangName = "far-end-link"
    farEndLink.EntityData.BundleName = "cisco_ios_xr"
    farEndLink.EntityData.ParentYangName = "rx-link"
    farEndLink.EntityData.SegmentPath = "far-end-link"
    farEndLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    farEndLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    farEndLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    farEndLink.EntityData.Children = make(map[string]types.YChild)
    farEndLink.EntityData.Children["asic-id"] = types.YChild{"AsicId", &farEndLink.AsicId}
    farEndLink.EntityData.Leafs = make(map[string]types.YLeaf)
    farEndLink.EntityData.Leafs["link-type"] = types.YLeaf{"LinkType", farEndLink.LinkType}
    farEndLink.EntityData.Leafs["link-stage"] = types.YLeaf{"LinkStage", farEndLink.LinkStage}
    farEndLink.EntityData.Leafs["link-num"] = types.YLeaf{"LinkNum", farEndLink.LinkNum}
    farEndLink.EntityData.Leafs["phy-link-num"] = types.YLeaf{"PhyLinkNum", farEndLink.PhyLinkNum}
    return &(farEndLink.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLink_AsicId
// asic id
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLink_AsicId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLink_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "far-end-link"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = make(map[string]types.YChild)
    asicId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicId.EntityData.Leafs["rack-type"] = types.YLeaf{"RackType", asicId.RackType}
    asicId.EntityData.Leafs["asic-type"] = types.YLeaf{"AsicType", asicId.AsicType}
    asicId.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", asicId.RackNum}
    asicId.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", asicId.SlotNum}
    asicId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicId.AsicInstance}
    return &(asicId.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLinkInHw
// far end link in hw
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLinkInHw struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link Type. The type is Link.
    LinkType interface{}

    // Link Stage. The type is LinkStage.
    LinkStage interface{}

    // link num. The type is interface{} with range: 0..4294967295.
    LinkNum interface{}

    // phy link num. The type is interface{} with range: 0..4294967295.
    PhyLinkNum interface{}

    // asic id.
    AsicId Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLinkInHw_AsicId
}

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLinkInHw) GetEntityData() *types.CommonEntityData {
    farEndLinkInHw.EntityData.YFilter = farEndLinkInHw.YFilter
    farEndLinkInHw.EntityData.YangName = "far-end-link-in-hw"
    farEndLinkInHw.EntityData.BundleName = "cisco_ios_xr"
    farEndLinkInHw.EntityData.ParentYangName = "rx-link"
    farEndLinkInHw.EntityData.SegmentPath = "far-end-link-in-hw"
    farEndLinkInHw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    farEndLinkInHw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    farEndLinkInHw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    farEndLinkInHw.EntityData.Children = make(map[string]types.YChild)
    farEndLinkInHw.EntityData.Children["asic-id"] = types.YChild{"AsicId", &farEndLinkInHw.AsicId}
    farEndLinkInHw.EntityData.Leafs = make(map[string]types.YLeaf)
    farEndLinkInHw.EntityData.Leafs["link-type"] = types.YLeaf{"LinkType", farEndLinkInHw.LinkType}
    farEndLinkInHw.EntityData.Leafs["link-stage"] = types.YLeaf{"LinkStage", farEndLinkInHw.LinkStage}
    farEndLinkInHw.EntityData.Leafs["link-num"] = types.YLeaf{"LinkNum", farEndLinkInHw.LinkNum}
    farEndLinkInHw.EntityData.Leafs["phy-link-num"] = types.YLeaf{"PhyLinkNum", farEndLinkInHw.PhyLinkNum}
    return &(farEndLinkInHw.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLinkInHw_AsicId
// asic id
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLinkInHw_AsicId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__FarEndLinkInHw_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "far-end-link-in-hw"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = make(map[string]types.YChild)
    asicId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicId.EntityData.Leafs["rack-type"] = types.YLeaf{"RackType", asicId.RackType}
    asicId.EntityData.Leafs["asic-type"] = types.YLeaf{"AsicType", asicId.AsicType}
    asicId.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", asicId.RackNum}
    asicId.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", asicId.SlotNum}
    asicId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicId.AsicInstance}
    return &(asicId.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__History
// history
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // histnum. The type is interface{} with range: 0..255.
    Histnum interface{}

    // start index. The type is interface{} with range: 0..255.
    StartIndex interface{}

    // hist. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__History_Hist.
    Hist []Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__History_Hist
}

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "rx-link"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Children["hist"] = types.YChild{"Hist", nil}
    for i := range history.Hist {
        history.EntityData.Children[types.GetSegmentPath(&history.Hist[i])] = types.YChild{"Hist", &history.Hist[i]}
    }
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    history.EntityData.Leafs["histnum"] = types.YLeaf{"Histnum", history.Histnum}
    history.EntityData.Leafs["start-index"] = types.YLeaf{"StartIndex", history.StartIndex}
    return &(history.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__History_Hist
// hist
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__History_Hist struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is OperState.
    OperState interface{}

    // Error State. The type is LinkErrorState.
    ErrorState interface{}

    // timestamp. The type is interface{} with range: 0..18446744073709551615.
    Timestamp interface{}

    // reasons. The type is string.
    Reasons interface{}
}

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink__History_Hist) GetEntityData() *types.CommonEntityData {
    hist.EntityData.YFilter = hist.YFilter
    hist.EntityData.YangName = "hist"
    hist.EntityData.BundleName = "cisco_ios_xr"
    hist.EntityData.ParentYangName = "history"
    hist.EntityData.SegmentPath = "hist"
    hist.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hist.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hist.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hist.EntityData.Children = make(map[string]types.YChild)
    hist.EntityData.Leafs = make(map[string]types.YLeaf)
    hist.EntityData.Leafs["admin-state"] = types.YLeaf{"AdminState", hist.AdminState}
    hist.EntityData.Leafs["oper-state"] = types.YLeaf{"OperState", hist.OperState}
    hist.EntityData.Leafs["error-state"] = types.YLeaf{"ErrorState", hist.ErrorState}
    hist.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", hist.Timestamp}
    hist.EntityData.Leafs["reasons"] = types.YLeaf{"Reasons", hist.Reasons}
    return &(hist.EntityData)
}

// Fia_Nodes_Node_DriverInformation
// FIA driver information
type Fia_Nodes_Node_DriverInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // drv version. The type is interface{} with range: 0..4294967295.
    DrvVersion interface{}

    // coeff major rev. The type is interface{} with range: 0..4294967295.
    CoeffMajorRev interface{}

    // coeff minor rev. The type is interface{} with range: 0..4294967295.
    CoeffMinorRev interface{}

    // functional role. The type is interface{} with range: 0..255.
    FunctionalRole interface{}

    // issu role. The type is interface{} with range: 0..255.
    IssuRole interface{}

    // node id. The type is string.
    NodeId interface{}

    // rack type. The type is interface{} with range: -2147483648..2147483647.
    RackType interface{}

    // rack num. The type is interface{} with range: 0..255.
    RackNum interface{}

    // is driver ready. The type is bool.
    IsDriverReady interface{}

    // card avail mask. The type is interface{} with range: 0..4294967295.
    CardAvailMask interface{}

    // asic avail mask. The type is interface{} with range:
    // 0..18446744073709551615.
    AsicAvailMask interface{}

    // exp asic avail mask. The type is interface{} with range:
    // 0..18446744073709551615.
    ExpAsicAvailMask interface{}

    // ucmc ratio. The type is interface{} with range: 0..4294967295.
    UcmcRatio interface{}

    // asic oper notify to fsdb pending bmap. The type is interface{} with range:
    // 0..18446744073709551615.
    AsicOperNotifyToFsdbPendingBmap interface{}

    // is full fgid download req. The type is bool.
    IsFullFgidDownloadReq interface{}

    // is fgid download in progress. The type is bool.
    IsFgidDownloadInProgress interface{}

    // is fgid download completed. The type is bool.
    IsFgidDownloadCompleted interface{}

    // fsdb conn active. The type is bool.
    FsdbConnActive interface{}

    // fgid conn active. The type is bool.
    FgidConnActive interface{}

    // issu mgr conn active. The type is bool.
    IssuMgrConnActive interface{}

    // fsdb reg active. The type is bool.
    FsdbRegActive interface{}

    // fgid reg active. The type is bool.
    FgidRegActive interface{}

    // issu mgr reg active. The type is bool.
    IssuMgrRegActive interface{}

    // num pm conn reqs. The type is interface{} with range: 0..255.
    NumPmConnReqs interface{}

    // num fsdb conn reqs. The type is interface{} with range: 0..255.
    NumFsdbConnReqs interface{}

    // num fgid conn reqs. The type is interface{} with range: 0..255.
    NumFgidConnReqs interface{}

    // num fstats conn reqs. The type is interface{} with range: 0..255.
    NumFstatsConnReqs interface{}

    // num cm conn reqs. The type is interface{} with range: 0..255.
    NumCmConnReqs interface{}

    // num issu mgr conn reqs. The type is interface{} with range: 0..255.
    NumIssuMgrConnReqs interface{}

    // num peer fia conn reqs. The type is interface{} with range: 0..255.
    NumPeerFiaConnReqs interface{}

    // is gaspp registered. The type is bool.
    IsGasppRegistered interface{}

    // is cih registered. The type is bool.
    IsCihRegistered interface{}

    // drvr initial startup timestamp. The type is string.
    DrvrInitialStartupTimestamp interface{}

    // drvr current startup timestamp. The type is string.
    DrvrCurrentStartupTimestamp interface{}

    // num intf ports. The type is interface{} with range: 0..4294967295.
    NumIntfPorts interface{}

    // uc weight. The type is interface{} with range: 0..255.
    UcWeight interface{}

    // respawn count. The type is interface{} with range: 0..255.
    RespawnCount interface{}

    // total asics. The type is interface{} with range: 0..255.
    TotalAsics interface{}

    // issu ready ntfy pending. The type is bool.
    IssuReadyNtfyPending interface{}

    // issu abort sent. The type is bool.
    IssuAbortSent interface{}

    // issu abort rcvd. The type is bool.
    IssuAbortRcvd interface{}

    // fabric mode. The type is interface{} with range: 0..255.
    FabricMode interface{}

    // FC Mode. The type is FcMode.
    FcMode interface{}

    // board rev id. The type is interface{} with range: 0..4294967295.
    BoardRevId interface{}

    // device info. The type is slice of
    // Fia_Nodes_Node_DriverInformation_DeviceInfo.
    DeviceInfo []Fia_Nodes_Node_DriverInformation_DeviceInfo

    // card info. The type is slice of Fia_Nodes_Node_DriverInformation_CardInfo.
    CardInfo []Fia_Nodes_Node_DriverInformation_CardInfo
}

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetEntityData() *types.CommonEntityData {
    driverInformation.EntityData.YFilter = driverInformation.YFilter
    driverInformation.EntityData.YangName = "driver-information"
    driverInformation.EntityData.BundleName = "cisco_ios_xr"
    driverInformation.EntityData.ParentYangName = "node"
    driverInformation.EntityData.SegmentPath = "driver-information"
    driverInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    driverInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    driverInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    driverInformation.EntityData.Children = make(map[string]types.YChild)
    driverInformation.EntityData.Children["device-info"] = types.YChild{"DeviceInfo", nil}
    for i := range driverInformation.DeviceInfo {
        driverInformation.EntityData.Children[types.GetSegmentPath(&driverInformation.DeviceInfo[i])] = types.YChild{"DeviceInfo", &driverInformation.DeviceInfo[i]}
    }
    driverInformation.EntityData.Children["card-info"] = types.YChild{"CardInfo", nil}
    for i := range driverInformation.CardInfo {
        driverInformation.EntityData.Children[types.GetSegmentPath(&driverInformation.CardInfo[i])] = types.YChild{"CardInfo", &driverInformation.CardInfo[i]}
    }
    driverInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    driverInformation.EntityData.Leafs["drv-version"] = types.YLeaf{"DrvVersion", driverInformation.DrvVersion}
    driverInformation.EntityData.Leafs["coeff-major-rev"] = types.YLeaf{"CoeffMajorRev", driverInformation.CoeffMajorRev}
    driverInformation.EntityData.Leafs["coeff-minor-rev"] = types.YLeaf{"CoeffMinorRev", driverInformation.CoeffMinorRev}
    driverInformation.EntityData.Leafs["functional-role"] = types.YLeaf{"FunctionalRole", driverInformation.FunctionalRole}
    driverInformation.EntityData.Leafs["issu-role"] = types.YLeaf{"IssuRole", driverInformation.IssuRole}
    driverInformation.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", driverInformation.NodeId}
    driverInformation.EntityData.Leafs["rack-type"] = types.YLeaf{"RackType", driverInformation.RackType}
    driverInformation.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", driverInformation.RackNum}
    driverInformation.EntityData.Leafs["is-driver-ready"] = types.YLeaf{"IsDriverReady", driverInformation.IsDriverReady}
    driverInformation.EntityData.Leafs["card-avail-mask"] = types.YLeaf{"CardAvailMask", driverInformation.CardAvailMask}
    driverInformation.EntityData.Leafs["asic-avail-mask"] = types.YLeaf{"AsicAvailMask", driverInformation.AsicAvailMask}
    driverInformation.EntityData.Leafs["exp-asic-avail-mask"] = types.YLeaf{"ExpAsicAvailMask", driverInformation.ExpAsicAvailMask}
    driverInformation.EntityData.Leafs["ucmc-ratio"] = types.YLeaf{"UcmcRatio", driverInformation.UcmcRatio}
    driverInformation.EntityData.Leafs["asic-oper-notify-to-fsdb-pending-bmap"] = types.YLeaf{"AsicOperNotifyToFsdbPendingBmap", driverInformation.AsicOperNotifyToFsdbPendingBmap}
    driverInformation.EntityData.Leafs["is-full-fgid-download-req"] = types.YLeaf{"IsFullFgidDownloadReq", driverInformation.IsFullFgidDownloadReq}
    driverInformation.EntityData.Leafs["is-fgid-download-in-progress"] = types.YLeaf{"IsFgidDownloadInProgress", driverInformation.IsFgidDownloadInProgress}
    driverInformation.EntityData.Leafs["is-fgid-download-completed"] = types.YLeaf{"IsFgidDownloadCompleted", driverInformation.IsFgidDownloadCompleted}
    driverInformation.EntityData.Leafs["fsdb-conn-active"] = types.YLeaf{"FsdbConnActive", driverInformation.FsdbConnActive}
    driverInformation.EntityData.Leafs["fgid-conn-active"] = types.YLeaf{"FgidConnActive", driverInformation.FgidConnActive}
    driverInformation.EntityData.Leafs["issu-mgr-conn-active"] = types.YLeaf{"IssuMgrConnActive", driverInformation.IssuMgrConnActive}
    driverInformation.EntityData.Leafs["fsdb-reg-active"] = types.YLeaf{"FsdbRegActive", driverInformation.FsdbRegActive}
    driverInformation.EntityData.Leafs["fgid-reg-active"] = types.YLeaf{"FgidRegActive", driverInformation.FgidRegActive}
    driverInformation.EntityData.Leafs["issu-mgr-reg-active"] = types.YLeaf{"IssuMgrRegActive", driverInformation.IssuMgrRegActive}
    driverInformation.EntityData.Leafs["num-pm-conn-reqs"] = types.YLeaf{"NumPmConnReqs", driverInformation.NumPmConnReqs}
    driverInformation.EntityData.Leafs["num-fsdb-conn-reqs"] = types.YLeaf{"NumFsdbConnReqs", driverInformation.NumFsdbConnReqs}
    driverInformation.EntityData.Leafs["num-fgid-conn-reqs"] = types.YLeaf{"NumFgidConnReqs", driverInformation.NumFgidConnReqs}
    driverInformation.EntityData.Leafs["num-fstats-conn-reqs"] = types.YLeaf{"NumFstatsConnReqs", driverInformation.NumFstatsConnReqs}
    driverInformation.EntityData.Leafs["num-cm-conn-reqs"] = types.YLeaf{"NumCmConnReqs", driverInformation.NumCmConnReqs}
    driverInformation.EntityData.Leafs["num-issu-mgr-conn-reqs"] = types.YLeaf{"NumIssuMgrConnReqs", driverInformation.NumIssuMgrConnReqs}
    driverInformation.EntityData.Leafs["num-peer-fia-conn-reqs"] = types.YLeaf{"NumPeerFiaConnReqs", driverInformation.NumPeerFiaConnReqs}
    driverInformation.EntityData.Leafs["is-gaspp-registered"] = types.YLeaf{"IsGasppRegistered", driverInformation.IsGasppRegistered}
    driverInformation.EntityData.Leafs["is-cih-registered"] = types.YLeaf{"IsCihRegistered", driverInformation.IsCihRegistered}
    driverInformation.EntityData.Leafs["drvr-initial-startup-timestamp"] = types.YLeaf{"DrvrInitialStartupTimestamp", driverInformation.DrvrInitialStartupTimestamp}
    driverInformation.EntityData.Leafs["drvr-current-startup-timestamp"] = types.YLeaf{"DrvrCurrentStartupTimestamp", driverInformation.DrvrCurrentStartupTimestamp}
    driverInformation.EntityData.Leafs["num-intf-ports"] = types.YLeaf{"NumIntfPorts", driverInformation.NumIntfPorts}
    driverInformation.EntityData.Leafs["uc-weight"] = types.YLeaf{"UcWeight", driverInformation.UcWeight}
    driverInformation.EntityData.Leafs["respawn-count"] = types.YLeaf{"RespawnCount", driverInformation.RespawnCount}
    driverInformation.EntityData.Leafs["total-asics"] = types.YLeaf{"TotalAsics", driverInformation.TotalAsics}
    driverInformation.EntityData.Leafs["issu-ready-ntfy-pending"] = types.YLeaf{"IssuReadyNtfyPending", driverInformation.IssuReadyNtfyPending}
    driverInformation.EntityData.Leafs["issu-abort-sent"] = types.YLeaf{"IssuAbortSent", driverInformation.IssuAbortSent}
    driverInformation.EntityData.Leafs["issu-abort-rcvd"] = types.YLeaf{"IssuAbortRcvd", driverInformation.IssuAbortRcvd}
    driverInformation.EntityData.Leafs["fabric-mode"] = types.YLeaf{"FabricMode", driverInformation.FabricMode}
    driverInformation.EntityData.Leafs["fc-mode"] = types.YLeaf{"FcMode", driverInformation.FcMode}
    driverInformation.EntityData.Leafs["board-rev-id"] = types.YLeaf{"BoardRevId", driverInformation.BoardRevId}
    return &(driverInformation.EntityData)
}

// Fia_Nodes_Node_DriverInformation_DeviceInfo
// device info
type Fia_Nodes_Node_DriverInformation_DeviceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // is valid. The type is bool.
    IsValid interface{}

    // fapid. The type is interface{} with range: 0..4294967295.
    Fapid interface{}

    // hotplug event. The type is interface{} with range: 0..4294967295.
    HotplugEvent interface{}

    // Slice State. The type is SliceState.
    SliceState interface{}

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is AsicOperState.
    OperState interface{}

    // Asic State. The type is AsicAccessState.
    AsicState interface{}

    // last init cause. The type is AsicInitMethod.
    LastInitCause interface{}

    // num pon resets. The type is interface{} with range: 0..4294967295.
    NumPonResets interface{}

    // num hard resets. The type is interface{} with range: 0..4294967295.
    NumHardResets interface{}

    // local switch state. The type is bool.
    LocalSwitchState interface{}

    // asic id.
    AsicId Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId
}

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetEntityData() *types.CommonEntityData {
    deviceInfo.EntityData.YFilter = deviceInfo.YFilter
    deviceInfo.EntityData.YangName = "device-info"
    deviceInfo.EntityData.BundleName = "cisco_ios_xr"
    deviceInfo.EntityData.ParentYangName = "driver-information"
    deviceInfo.EntityData.SegmentPath = "device-info"
    deviceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    deviceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    deviceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    deviceInfo.EntityData.Children = make(map[string]types.YChild)
    deviceInfo.EntityData.Children["asic-id"] = types.YChild{"AsicId", &deviceInfo.AsicId}
    deviceInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    deviceInfo.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", deviceInfo.IsValid}
    deviceInfo.EntityData.Leafs["fapid"] = types.YLeaf{"Fapid", deviceInfo.Fapid}
    deviceInfo.EntityData.Leafs["hotplug-event"] = types.YLeaf{"HotplugEvent", deviceInfo.HotplugEvent}
    deviceInfo.EntityData.Leafs["slice-state"] = types.YLeaf{"SliceState", deviceInfo.SliceState}
    deviceInfo.EntityData.Leafs["admin-state"] = types.YLeaf{"AdminState", deviceInfo.AdminState}
    deviceInfo.EntityData.Leafs["oper-state"] = types.YLeaf{"OperState", deviceInfo.OperState}
    deviceInfo.EntityData.Leafs["asic-state"] = types.YLeaf{"AsicState", deviceInfo.AsicState}
    deviceInfo.EntityData.Leafs["last-init-cause"] = types.YLeaf{"LastInitCause", deviceInfo.LastInitCause}
    deviceInfo.EntityData.Leafs["num-pon-resets"] = types.YLeaf{"NumPonResets", deviceInfo.NumPonResets}
    deviceInfo.EntityData.Leafs["num-hard-resets"] = types.YLeaf{"NumHardResets", deviceInfo.NumHardResets}
    deviceInfo.EntityData.Leafs["local-switch-state"] = types.YLeaf{"LocalSwitchState", deviceInfo.LocalSwitchState}
    return &(deviceInfo.EntityData)
}

// Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId
// asic id
type Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "device-info"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = make(map[string]types.YChild)
    asicId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicId.EntityData.Leafs["rack-type"] = types.YLeaf{"RackType", asicId.RackType}
    asicId.EntityData.Leafs["asic-type"] = types.YLeaf{"AsicType", asicId.AsicType}
    asicId.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", asicId.RackNum}
    asicId.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", asicId.SlotNum}
    asicId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicId.AsicInstance}
    return &(asicId.EntityData)
}

// Fia_Nodes_Node_DriverInformation_CardInfo
// card info
type Fia_Nodes_Node_DriverInformation_CardInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card type. The type is interface{} with range: -2147483648..2147483647.
    CardType interface{}

    // card name. The type is string.
    CardName interface{}

    // slot no. The type is interface{} with range: -2147483648..2147483647.
    SlotNo interface{}

    // card flag. The type is interface{} with range: -2147483648..2147483647.
    CardFlag interface{}

    // evt flag. The type is interface{} with range: -2147483648..2147483647.
    EvtFlag interface{}

    // reg flag. The type is interface{} with range: -2147483648..2147483647.
    RegFlag interface{}

    // instance. The type is interface{} with range: -2147483648..2147483647.
    Instance interface{}

    // card state. The type is interface{} with range: 0..255.
    CardState interface{}

    // exp num asics. The type is interface{} with range: 0..4294967295.
    ExpNumAsics interface{}

    // exp num asics per fsdb. The type is interface{} with range: 0..4294967295.
    ExpNumAsicsPerFsdb interface{}

    // is powered. The type is bool.
    IsPowered interface{}

    // cxp avail bitmap. The type is interface{} with range:
    // 0..18446744073709551615.
    CxpAvailBitmap interface{}

    // num ilkns per asic. The type is interface{} with range: 0..4294967295.
    NumIlknsPerAsic interface{}

    // num local ports per ilkn. The type is interface{} with range:
    // 0..4294967295.
    NumLocalPortsPerIlkn interface{}

    // num cos per port. The type is interface{} with range: 0..255.
    NumCosPerPort interface{}

    // oir circular buffer.
    OirCircularBuffer Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer
}

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetEntityData() *types.CommonEntityData {
    cardInfo.EntityData.YFilter = cardInfo.YFilter
    cardInfo.EntityData.YangName = "card-info"
    cardInfo.EntityData.BundleName = "cisco_ios_xr"
    cardInfo.EntityData.ParentYangName = "driver-information"
    cardInfo.EntityData.SegmentPath = "card-info"
    cardInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardInfo.EntityData.Children = make(map[string]types.YChild)
    cardInfo.EntityData.Children["oir-circular-buffer"] = types.YChild{"OirCircularBuffer", &cardInfo.OirCircularBuffer}
    cardInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    cardInfo.EntityData.Leafs["card-type"] = types.YLeaf{"CardType", cardInfo.CardType}
    cardInfo.EntityData.Leafs["card-name"] = types.YLeaf{"CardName", cardInfo.CardName}
    cardInfo.EntityData.Leafs["slot-no"] = types.YLeaf{"SlotNo", cardInfo.SlotNo}
    cardInfo.EntityData.Leafs["card-flag"] = types.YLeaf{"CardFlag", cardInfo.CardFlag}
    cardInfo.EntityData.Leafs["evt-flag"] = types.YLeaf{"EvtFlag", cardInfo.EvtFlag}
    cardInfo.EntityData.Leafs["reg-flag"] = types.YLeaf{"RegFlag", cardInfo.RegFlag}
    cardInfo.EntityData.Leafs["instance"] = types.YLeaf{"Instance", cardInfo.Instance}
    cardInfo.EntityData.Leafs["card-state"] = types.YLeaf{"CardState", cardInfo.CardState}
    cardInfo.EntityData.Leafs["exp-num-asics"] = types.YLeaf{"ExpNumAsics", cardInfo.ExpNumAsics}
    cardInfo.EntityData.Leafs["exp-num-asics-per-fsdb"] = types.YLeaf{"ExpNumAsicsPerFsdb", cardInfo.ExpNumAsicsPerFsdb}
    cardInfo.EntityData.Leafs["is-powered"] = types.YLeaf{"IsPowered", cardInfo.IsPowered}
    cardInfo.EntityData.Leafs["cxp-avail-bitmap"] = types.YLeaf{"CxpAvailBitmap", cardInfo.CxpAvailBitmap}
    cardInfo.EntityData.Leafs["num-ilkns-per-asic"] = types.YLeaf{"NumIlknsPerAsic", cardInfo.NumIlknsPerAsic}
    cardInfo.EntityData.Leafs["num-local-ports-per-ilkn"] = types.YLeaf{"NumLocalPortsPerIlkn", cardInfo.NumLocalPortsPerIlkn}
    cardInfo.EntityData.Leafs["num-cos-per-port"] = types.YLeaf{"NumCosPerPort", cardInfo.NumCosPerPort}
    return &(cardInfo.EntityData)
}

// Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer
// oir circular buffer
type Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // count. The type is interface{} with range: -2147483648..2147483647.
    Count interface{}

    // start. The type is interface{} with range: -2147483648..2147483647.
    Start interface{}

    // end. The type is interface{} with range: -2147483648..2147483647.
    End interface{}

    // fia oir info. The type is slice of
    // Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo.
    FiaOirInfo []Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo
}

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetEntityData() *types.CommonEntityData {
    oirCircularBuffer.EntityData.YFilter = oirCircularBuffer.YFilter
    oirCircularBuffer.EntityData.YangName = "oir-circular-buffer"
    oirCircularBuffer.EntityData.BundleName = "cisco_ios_xr"
    oirCircularBuffer.EntityData.ParentYangName = "card-info"
    oirCircularBuffer.EntityData.SegmentPath = "oir-circular-buffer"
    oirCircularBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oirCircularBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oirCircularBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oirCircularBuffer.EntityData.Children = make(map[string]types.YChild)
    oirCircularBuffer.EntityData.Children["fia-oir-info"] = types.YChild{"FiaOirInfo", nil}
    for i := range oirCircularBuffer.FiaOirInfo {
        oirCircularBuffer.EntityData.Children[types.GetSegmentPath(&oirCircularBuffer.FiaOirInfo[i])] = types.YChild{"FiaOirInfo", &oirCircularBuffer.FiaOirInfo[i]}
    }
    oirCircularBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    oirCircularBuffer.EntityData.Leafs["count"] = types.YLeaf{"Count", oirCircularBuffer.Count}
    oirCircularBuffer.EntityData.Leafs["start"] = types.YLeaf{"Start", oirCircularBuffer.Start}
    oirCircularBuffer.EntityData.Leafs["end"] = types.YLeaf{"End", oirCircularBuffer.End}
    return &(oirCircularBuffer.EntityData)
}

// Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo
// fia oir info
type Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card flag. The type is interface{} with range: -2147483648..2147483647.
    CardFlag interface{}

    // card type. The type is interface{} with range: -2147483648..2147483647.
    CardType interface{}

    // reg flag. The type is interface{} with range: -2147483648..2147483647.
    RegFlag interface{}

    // evt flag. The type is interface{} with range: -2147483648..2147483647.
    EvtFlag interface{}

    // rack num. The type is interface{} with range: -2147483648..2147483647.
    RackNum interface{}

    // instance. The type is interface{} with range: -2147483648..2147483647.
    Instance interface{}

    // cur card state. The type is interface{} with range:
    // -2147483648..2147483647.
    CurCardState interface{}
}

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetEntityData() *types.CommonEntityData {
    fiaOirInfo.EntityData.YFilter = fiaOirInfo.YFilter
    fiaOirInfo.EntityData.YangName = "fia-oir-info"
    fiaOirInfo.EntityData.BundleName = "cisco_ios_xr"
    fiaOirInfo.EntityData.ParentYangName = "oir-circular-buffer"
    fiaOirInfo.EntityData.SegmentPath = "fia-oir-info"
    fiaOirInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fiaOirInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fiaOirInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fiaOirInfo.EntityData.Children = make(map[string]types.YChild)
    fiaOirInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fiaOirInfo.EntityData.Leafs["card-flag"] = types.YLeaf{"CardFlag", fiaOirInfo.CardFlag}
    fiaOirInfo.EntityData.Leafs["card-type"] = types.YLeaf{"CardType", fiaOirInfo.CardType}
    fiaOirInfo.EntityData.Leafs["reg-flag"] = types.YLeaf{"RegFlag", fiaOirInfo.RegFlag}
    fiaOirInfo.EntityData.Leafs["evt-flag"] = types.YLeaf{"EvtFlag", fiaOirInfo.EvtFlag}
    fiaOirInfo.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", fiaOirInfo.RackNum}
    fiaOirInfo.EntityData.Leafs["instance"] = types.YLeaf{"Instance", fiaOirInfo.Instance}
    fiaOirInfo.EntityData.Leafs["cur-card-state"] = types.YLeaf{"CurCardState", fiaOirInfo.CurCardState}
    return &(fiaOirInfo.EntityData)
}

// Fia_Nodes_Node_ClearStatistics
// Clear statistics information
type Fia_Nodes_Node_ClearStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance table for clear statistics information.
    AsicInstances Fia_Nodes_Node_ClearStatistics_AsicInstances
}

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetEntityData() *types.CommonEntityData {
    clearStatistics.EntityData.YFilter = clearStatistics.YFilter
    clearStatistics.EntityData.YangName = "clear-statistics"
    clearStatistics.EntityData.BundleName = "cisco_ios_xr"
    clearStatistics.EntityData.ParentYangName = "node"
    clearStatistics.EntityData.SegmentPath = "clear-statistics"
    clearStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearStatistics.EntityData.Children = make(map[string]types.YChild)
    clearStatistics.EntityData.Children["asic-instances"] = types.YChild{"AsicInstances", &clearStatistics.AsicInstances}
    clearStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clearStatistics.EntityData)
}

// Fia_Nodes_Node_ClearStatistics_AsicInstances
// Instance table for clear statistics
// information
type Fia_Nodes_Node_ClearStatistics_AsicInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Asic instance to be cleared. The type is slice of
    // Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance.
    AsicInstance []Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance
}

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetEntityData() *types.CommonEntityData {
    asicInstances.EntityData.YFilter = asicInstances.YFilter
    asicInstances.EntityData.YangName = "asic-instances"
    asicInstances.EntityData.BundleName = "cisco_ios_xr"
    asicInstances.EntityData.ParentYangName = "clear-statistics"
    asicInstances.EntityData.SegmentPath = "asic-instances"
    asicInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicInstances.EntityData.Children = make(map[string]types.YChild)
    asicInstances.EntityData.Children["asic-instance"] = types.YChild{"AsicInstance", nil}
    for i := range asicInstances.AsicInstance {
        asicInstances.EntityData.Children[types.GetSegmentPath(&asicInstances.AsicInstance[i])] = types.YChild{"AsicInstance", &asicInstances.AsicInstance[i]}
    }
    asicInstances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicInstances.EntityData)
}

// Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance
// Asic instance to be cleared
type Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Asic instance. The type is interface{} with range:
    // 0..255.
    AsicInstance interface{}

    // Clear value. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    Instance interface{}
}

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetEntityData() *types.CommonEntityData {
    asicInstance.EntityData.YFilter = asicInstance.YFilter
    asicInstance.EntityData.YangName = "asic-instance"
    asicInstance.EntityData.BundleName = "cisco_ios_xr"
    asicInstance.EntityData.ParentYangName = "asic-instances"
    asicInstance.EntityData.SegmentPath = "asic-instance" + "[asic-instance='" + fmt.Sprintf("%v", asicInstance.AsicInstance) + "']"
    asicInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicInstance.EntityData.Children = make(map[string]types.YChild)
    asicInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    asicInstance.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicInstance.AsicInstance}
    asicInstance.EntityData.Leafs["instance"] = types.YLeaf{"Instance", asicInstance.Instance}
    return &(asicInstance.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation
// FIA link TX information
type Fia_Nodes_Node_TxLinkInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link table for tx information.
    TxStatusOptionTable Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable
}

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetEntityData() *types.CommonEntityData {
    txLinkInformation.EntityData.YFilter = txLinkInformation.YFilter
    txLinkInformation.EntityData.YangName = "tx-link-information"
    txLinkInformation.EntityData.BundleName = "cisco_ios_xr"
    txLinkInformation.EntityData.ParentYangName = "node"
    txLinkInformation.EntityData.SegmentPath = "tx-link-information"
    txLinkInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txLinkInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txLinkInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txLinkInformation.EntityData.Children = make(map[string]types.YChild)
    txLinkInformation.EntityData.Children["tx-status-option-table"] = types.YChild{"TxStatusOptionTable", &txLinkInformation.TxStatusOptionTable}
    txLinkInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(txLinkInformation.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable
// Link table for tx information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Option: data, ctrl, all- for now none.
    TxStatusOption Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption
}

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetEntityData() *types.CommonEntityData {
    txStatusOptionTable.EntityData.YFilter = txStatusOptionTable.YFilter
    txStatusOptionTable.EntityData.YangName = "tx-status-option-table"
    txStatusOptionTable.EntityData.BundleName = "cisco_ios_xr"
    txStatusOptionTable.EntityData.ParentYangName = "tx-link-information"
    txStatusOptionTable.EntityData.SegmentPath = "tx-status-option-table"
    txStatusOptionTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txStatusOptionTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txStatusOptionTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txStatusOptionTable.EntityData.Children = make(map[string]types.YChild)
    txStatusOptionTable.EntityData.Children["tx-status-option"] = types.YChild{"TxStatusOption", &txStatusOptionTable.TxStatusOption}
    txStatusOptionTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(txStatusOptionTable.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption
// Option: data, ctrl, all- for now none
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance table for tx information.
    TxAsicInstances Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances
}

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetEntityData() *types.CommonEntityData {
    txStatusOption.EntityData.YFilter = txStatusOption.YFilter
    txStatusOption.EntityData.YangName = "tx-status-option"
    txStatusOption.EntityData.BundleName = "cisco_ios_xr"
    txStatusOption.EntityData.ParentYangName = "tx-status-option-table"
    txStatusOption.EntityData.SegmentPath = "tx-status-option"
    txStatusOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txStatusOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txStatusOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txStatusOption.EntityData.Children = make(map[string]types.YChild)
    txStatusOption.EntityData.Children["tx-asic-instances"] = types.YChild{"TxAsicInstances", &txStatusOption.TxAsicInstances}
    txStatusOption.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(txStatusOption.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances
// Instance table for tx information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance number for tx link information. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance.
    TxAsicInstance []Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance
}

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetEntityData() *types.CommonEntityData {
    txAsicInstances.EntityData.YFilter = txAsicInstances.YFilter
    txAsicInstances.EntityData.YangName = "tx-asic-instances"
    txAsicInstances.EntityData.BundleName = "cisco_ios_xr"
    txAsicInstances.EntityData.ParentYangName = "tx-status-option"
    txAsicInstances.EntityData.SegmentPath = "tx-asic-instances"
    txAsicInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txAsicInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txAsicInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txAsicInstances.EntityData.Children = make(map[string]types.YChild)
    txAsicInstances.EntityData.Children["tx-asic-instance"] = types.YChild{"TxAsicInstance", nil}
    for i := range txAsicInstances.TxAsicInstance {
        txAsicInstances.EntityData.Children[types.GetSegmentPath(&txAsicInstances.TxAsicInstance[i])] = types.YChild{"TxAsicInstance", &txAsicInstances.TxAsicInstance[i]}
    }
    txAsicInstances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(txAsicInstances.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance
// Instance number for tx link information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Transmit instance. The type is interface{} with
    // range: 0..255.
    Instance interface{}

    // Link table for tx information.
    TxLinks Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks
}

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetEntityData() *types.CommonEntityData {
    txAsicInstance.EntityData.YFilter = txAsicInstance.YFilter
    txAsicInstance.EntityData.YangName = "tx-asic-instance"
    txAsicInstance.EntityData.BundleName = "cisco_ios_xr"
    txAsicInstance.EntityData.ParentYangName = "tx-asic-instances"
    txAsicInstance.EntityData.SegmentPath = "tx-asic-instance" + "[instance='" + fmt.Sprintf("%v", txAsicInstance.Instance) + "']"
    txAsicInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txAsicInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txAsicInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txAsicInstance.EntityData.Children = make(map[string]types.YChild)
    txAsicInstance.EntityData.Children["tx-links"] = types.YChild{"TxLinks", &txAsicInstance.TxLinks}
    txAsicInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    txAsicInstance.EntityData.Leafs["instance"] = types.YLeaf{"Instance", txAsicInstance.Instance}
    return &(txAsicInstance.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks
// Link table for tx information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link number for tx link information. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink.
    TxLink []Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink
}

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetEntityData() *types.CommonEntityData {
    txLinks.EntityData.YFilter = txLinks.YFilter
    txLinks.EntityData.YangName = "tx-links"
    txLinks.EntityData.BundleName = "cisco_ios_xr"
    txLinks.EntityData.ParentYangName = "tx-asic-instance"
    txLinks.EntityData.SegmentPath = "tx-links"
    txLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txLinks.EntityData.Children = make(map[string]types.YChild)
    txLinks.EntityData.Children["tx-link"] = types.YChild{"TxLink", nil}
    for i := range txLinks.TxLink {
        txLinks.EntityData.Children[types.GetSegmentPath(&txLinks.TxLink[i])] = types.YChild{"TxLink", &txLinks.TxLink[i]}
    }
    txLinks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(txLinks.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink
// Link number for tx link information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start number. The type is interface{} with range: 0..47.
    StartNumber interface{}

    // End number. The type is interface{} with range: 0..47.
    EndNumber interface{}

    // Single link information. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink.
    TxLink []Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetEntityData() *types.CommonEntityData {
    txLink.EntityData.YFilter = txLink.YFilter
    txLink.EntityData.YangName = "tx-link"
    txLink.EntityData.BundleName = "cisco_ios_xr"
    txLink.EntityData.ParentYangName = "tx-links"
    txLink.EntityData.SegmentPath = "tx-link"
    txLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txLink.EntityData.Children = make(map[string]types.YChild)
    txLink.EntityData.Children["tx-link"] = types.YChild{"TxLink", nil}
    for i := range txLink.TxLink {
        txLink.EntityData.Children[types.GetSegmentPath(&txLink.TxLink[i])] = types.YChild{"TxLink", &txLink.TxLink[i]}
    }
    txLink.EntityData.Leafs = make(map[string]types.YLeaf)
    txLink.EntityData.Leafs["start-number"] = types.YLeaf{"StartNumber", txLink.StartNumber}
    txLink.EntityData.Leafs["end-number"] = types.YLeaf{"EndNumber", txLink.EndNumber}
    return &(txLink.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_
// Single link information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Single Link. The type is interface{} with range:
    // -2147483648..2147483647.
    Link interface{}

    // speed. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // stage. The type is interface{} with range: 0..255.
    Stage interface{}

    // is link valid. The type is bool.
    IsLinkValid interface{}

    // is conf pending. The type is bool.
    IsConfPending interface{}

    // is power enabled. The type is bool.
    IsPowerEnabled interface{}

    // coeff1. The type is interface{} with range: 0..4294967295.
    Coeff1 interface{}

    // coeff2. The type is interface{} with range: 0..4294967295.
    Coeff2 interface{}

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is OperState.
    OperState interface{}

    // Error State. The type is LinkErrorState.
    ErrorState interface{}

    // num admin shuts. The type is interface{} with range: 0..4294967295.
    NumAdminShuts interface{}

    // this link.
    ThisLink Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__ThisLink

    // far end link.
    FarEndLink Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__FarEndLink

    // stats.
    Stats Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__Stats

    // history.
    History Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__History
}

func (txLink_ *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_) GetEntityData() *types.CommonEntityData {
    txLink_.EntityData.YFilter = txLink_.YFilter
    txLink_.EntityData.YangName = "tx-link"
    txLink_.EntityData.BundleName = "cisco_ios_xr"
    txLink_.EntityData.ParentYangName = "tx-link"
    txLink_.EntityData.SegmentPath = "tx-link" + "[link='" + fmt.Sprintf("%v", txLink_.Link) + "']"
    txLink_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txLink_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txLink_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txLink_.EntityData.Children = make(map[string]types.YChild)
    txLink_.EntityData.Children["this-link"] = types.YChild{"ThisLink", &txLink_.ThisLink}
    txLink_.EntityData.Children["far-end-link"] = types.YChild{"FarEndLink", &txLink_.FarEndLink}
    txLink_.EntityData.Children["stats"] = types.YChild{"Stats", &txLink_.Stats}
    txLink_.EntityData.Children["history"] = types.YChild{"History", &txLink_.History}
    txLink_.EntityData.Leafs = make(map[string]types.YLeaf)
    txLink_.EntityData.Leafs["link"] = types.YLeaf{"Link", txLink_.Link}
    txLink_.EntityData.Leafs["speed"] = types.YLeaf{"Speed", txLink_.Speed}
    txLink_.EntityData.Leafs["stage"] = types.YLeaf{"Stage", txLink_.Stage}
    txLink_.EntityData.Leafs["is-link-valid"] = types.YLeaf{"IsLinkValid", txLink_.IsLinkValid}
    txLink_.EntityData.Leafs["is-conf-pending"] = types.YLeaf{"IsConfPending", txLink_.IsConfPending}
    txLink_.EntityData.Leafs["is-power-enabled"] = types.YLeaf{"IsPowerEnabled", txLink_.IsPowerEnabled}
    txLink_.EntityData.Leafs["coeff1"] = types.YLeaf{"Coeff1", txLink_.Coeff1}
    txLink_.EntityData.Leafs["coeff2"] = types.YLeaf{"Coeff2", txLink_.Coeff2}
    txLink_.EntityData.Leafs["admin-state"] = types.YLeaf{"AdminState", txLink_.AdminState}
    txLink_.EntityData.Leafs["oper-state"] = types.YLeaf{"OperState", txLink_.OperState}
    txLink_.EntityData.Leafs["error-state"] = types.YLeaf{"ErrorState", txLink_.ErrorState}
    txLink_.EntityData.Leafs["num-admin-shuts"] = types.YLeaf{"NumAdminShuts", txLink_.NumAdminShuts}
    return &(txLink_.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__ThisLink
// this link
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__ThisLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link Type. The type is Link.
    LinkType interface{}

    // Link Stage. The type is LinkStage.
    LinkStage interface{}

    // link num. The type is interface{} with range: 0..4294967295.
    LinkNum interface{}

    // phy link num. The type is interface{} with range: 0..4294967295.
    PhyLinkNum interface{}

    // asic id.
    AsicId Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__ThisLink_AsicId
}

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__ThisLink) GetEntityData() *types.CommonEntityData {
    thisLink.EntityData.YFilter = thisLink.YFilter
    thisLink.EntityData.YangName = "this-link"
    thisLink.EntityData.BundleName = "cisco_ios_xr"
    thisLink.EntityData.ParentYangName = "tx-link"
    thisLink.EntityData.SegmentPath = "this-link"
    thisLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thisLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thisLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thisLink.EntityData.Children = make(map[string]types.YChild)
    thisLink.EntityData.Children["asic-id"] = types.YChild{"AsicId", &thisLink.AsicId}
    thisLink.EntityData.Leafs = make(map[string]types.YLeaf)
    thisLink.EntityData.Leafs["link-type"] = types.YLeaf{"LinkType", thisLink.LinkType}
    thisLink.EntityData.Leafs["link-stage"] = types.YLeaf{"LinkStage", thisLink.LinkStage}
    thisLink.EntityData.Leafs["link-num"] = types.YLeaf{"LinkNum", thisLink.LinkNum}
    thisLink.EntityData.Leafs["phy-link-num"] = types.YLeaf{"PhyLinkNum", thisLink.PhyLinkNum}
    return &(thisLink.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__ThisLink_AsicId
// asic id
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__ThisLink_AsicId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__ThisLink_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "this-link"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = make(map[string]types.YChild)
    asicId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicId.EntityData.Leafs["rack-type"] = types.YLeaf{"RackType", asicId.RackType}
    asicId.EntityData.Leafs["asic-type"] = types.YLeaf{"AsicType", asicId.AsicType}
    asicId.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", asicId.RackNum}
    asicId.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", asicId.SlotNum}
    asicId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicId.AsicInstance}
    return &(asicId.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__FarEndLink
// far end link
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__FarEndLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link Type. The type is Link.
    LinkType interface{}

    // Link Stage. The type is LinkStage.
    LinkStage interface{}

    // link num. The type is interface{} with range: 0..4294967295.
    LinkNum interface{}

    // phy link num. The type is interface{} with range: 0..4294967295.
    PhyLinkNum interface{}

    // asic id.
    AsicId Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__FarEndLink_AsicId
}

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__FarEndLink) GetEntityData() *types.CommonEntityData {
    farEndLink.EntityData.YFilter = farEndLink.YFilter
    farEndLink.EntityData.YangName = "far-end-link"
    farEndLink.EntityData.BundleName = "cisco_ios_xr"
    farEndLink.EntityData.ParentYangName = "tx-link"
    farEndLink.EntityData.SegmentPath = "far-end-link"
    farEndLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    farEndLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    farEndLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    farEndLink.EntityData.Children = make(map[string]types.YChild)
    farEndLink.EntityData.Children["asic-id"] = types.YChild{"AsicId", &farEndLink.AsicId}
    farEndLink.EntityData.Leafs = make(map[string]types.YLeaf)
    farEndLink.EntityData.Leafs["link-type"] = types.YLeaf{"LinkType", farEndLink.LinkType}
    farEndLink.EntityData.Leafs["link-stage"] = types.YLeaf{"LinkStage", farEndLink.LinkStage}
    farEndLink.EntityData.Leafs["link-num"] = types.YLeaf{"LinkNum", farEndLink.LinkNum}
    farEndLink.EntityData.Leafs["phy-link-num"] = types.YLeaf{"PhyLinkNum", farEndLink.PhyLinkNum}
    return &(farEndLink.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__FarEndLink_AsicId
// asic id
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__FarEndLink_AsicId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__FarEndLink_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "far-end-link"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = make(map[string]types.YChild)
    asicId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicId.EntityData.Leafs["rack-type"] = types.YLeaf{"RackType", asicId.RackType}
    asicId.EntityData.Leafs["asic-type"] = types.YLeaf{"AsicType", asicId.AsicType}
    asicId.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", asicId.RackNum}
    asicId.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", asicId.SlotNum}
    asicId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicId.AsicInstance}
    return &(asicId.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__Stats
// stats
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dummy. The type is interface{} with range: 0..4294967295.
    Dummy interface{}
}

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "tx-link"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["dummy"] = types.YLeaf{"Dummy", stats.Dummy}
    return &(stats.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__History
// history
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // histnum. The type is interface{} with range: 0..255.
    Histnum interface{}

    // start index. The type is interface{} with range: 0..255.
    StartIndex interface{}

    // hist. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__History_Hist.
    Hist []Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__History_Hist
}

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "tx-link"
    history.EntityData.SegmentPath = "history"
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = make(map[string]types.YChild)
    history.EntityData.Children["hist"] = types.YChild{"Hist", nil}
    for i := range history.Hist {
        history.EntityData.Children[types.GetSegmentPath(&history.Hist[i])] = types.YChild{"Hist", &history.Hist[i]}
    }
    history.EntityData.Leafs = make(map[string]types.YLeaf)
    history.EntityData.Leafs["histnum"] = types.YLeaf{"Histnum", history.Histnum}
    history.EntityData.Leafs["start-index"] = types.YLeaf{"StartIndex", history.StartIndex}
    return &(history.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__History_Hist
// hist
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__History_Hist struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is OperState.
    OperState interface{}

    // Error State. The type is LinkErrorState.
    ErrorState interface{}

    // timestamp. The type is interface{} with range: 0..18446744073709551615.
    Timestamp interface{}

    // reasons. The type is string.
    Reasons interface{}
}

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink__History_Hist) GetEntityData() *types.CommonEntityData {
    hist.EntityData.YFilter = hist.YFilter
    hist.EntityData.YangName = "hist"
    hist.EntityData.BundleName = "cisco_ios_xr"
    hist.EntityData.ParentYangName = "history"
    hist.EntityData.SegmentPath = "hist"
    hist.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hist.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hist.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hist.EntityData.Children = make(map[string]types.YChild)
    hist.EntityData.Leafs = make(map[string]types.YLeaf)
    hist.EntityData.Leafs["admin-state"] = types.YLeaf{"AdminState", hist.AdminState}
    hist.EntityData.Leafs["oper-state"] = types.YLeaf{"OperState", hist.OperState}
    hist.EntityData.Leafs["error-state"] = types.YLeaf{"ErrorState", hist.ErrorState}
    hist.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", hist.Timestamp}
    hist.EntityData.Leafs["reasons"] = types.YLeaf{"Reasons", hist.Reasons}
    return &(hist.EntityData)
}

// Fia_Nodes_Node_DiagShell
// FIA diag shell information
type Fia_Nodes_Node_DiagShell struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unit table for diag shell.
    DiagShellUnits Fia_Nodes_Node_DiagShell_DiagShellUnits
}

func (diagShell *Fia_Nodes_Node_DiagShell) GetEntityData() *types.CommonEntityData {
    diagShell.EntityData.YFilter = diagShell.YFilter
    diagShell.EntityData.YangName = "diag-shell"
    diagShell.EntityData.BundleName = "cisco_ios_xr"
    diagShell.EntityData.ParentYangName = "node"
    diagShell.EntityData.SegmentPath = "diag-shell"
    diagShell.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diagShell.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diagShell.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diagShell.EntityData.Children = make(map[string]types.YChild)
    diagShell.EntityData.Children["diag-shell-units"] = types.YChild{"DiagShellUnits", &diagShell.DiagShellUnits}
    diagShell.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diagShell.EntityData)
}

// Fia_Nodes_Node_DiagShell_DiagShellUnits
// Unit table for diag shell
type Fia_Nodes_Node_DiagShell_DiagShellUnits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unit number for diag shell statistics. The type is slice of
    // Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit.
    DiagShellUnit []Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit
}

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetEntityData() *types.CommonEntityData {
    diagShellUnits.EntityData.YFilter = diagShellUnits.YFilter
    diagShellUnits.EntityData.YangName = "diag-shell-units"
    diagShellUnits.EntityData.BundleName = "cisco_ios_xr"
    diagShellUnits.EntityData.ParentYangName = "diag-shell"
    diagShellUnits.EntityData.SegmentPath = "diag-shell-units"
    diagShellUnits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diagShellUnits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diagShellUnits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diagShellUnits.EntityData.Children = make(map[string]types.YChild)
    diagShellUnits.EntityData.Children["diag-shell-unit"] = types.YChild{"DiagShellUnit", nil}
    for i := range diagShellUnits.DiagShellUnit {
        diagShellUnits.EntityData.Children[types.GetSegmentPath(&diagShellUnits.DiagShellUnit[i])] = types.YChild{"DiagShellUnit", &diagShellUnits.DiagShellUnit[i]}
    }
    diagShellUnits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diagShellUnits.EntityData)
}

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit
// Unit number for diag shell statistics
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Unit number. The type is interface{} with range:
    // 0..63.
    Unit interface{}

    // Command table for diag shell.
    Commands Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands
}

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetEntityData() *types.CommonEntityData {
    diagShellUnit.EntityData.YFilter = diagShellUnit.YFilter
    diagShellUnit.EntityData.YangName = "diag-shell-unit"
    diagShellUnit.EntityData.BundleName = "cisco_ios_xr"
    diagShellUnit.EntityData.ParentYangName = "diag-shell-units"
    diagShellUnit.EntityData.SegmentPath = "diag-shell-unit" + "[unit='" + fmt.Sprintf("%v", diagShellUnit.Unit) + "']"
    diagShellUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diagShellUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diagShellUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diagShellUnit.EntityData.Children = make(map[string]types.YChild)
    diagShellUnit.EntityData.Children["commands"] = types.YChild{"Commands", &diagShellUnit.Commands}
    diagShellUnit.EntityData.Leafs = make(map[string]types.YLeaf)
    diagShellUnit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", diagShellUnit.Unit}
    return &(diagShellUnit.EntityData)
}

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands
// Command table for diag shell
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Command for diag shell statistics. The type is slice of
    // Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command.
    Command []Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command
}

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetEntityData() *types.CommonEntityData {
    commands.EntityData.YFilter = commands.YFilter
    commands.EntityData.YangName = "commands"
    commands.EntityData.BundleName = "cisco_ios_xr"
    commands.EntityData.ParentYangName = "diag-shell-unit"
    commands.EntityData.SegmentPath = "commands"
    commands.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commands.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commands.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commands.EntityData.Children = make(map[string]types.YChild)
    commands.EntityData.Children["command"] = types.YChild{"Command", nil}
    for i := range commands.Command {
        commands.EntityData.Children[types.GetSegmentPath(&commands.Command[i])] = types.YChild{"Command", &commands.Command[i]}
    }
    commands.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(commands.EntityData)
}

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command
// Command for diag shell statistics
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Shell command. The type is string.
    Cmd interface{}

    // Added to support datalist. The type is slice of
    // Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output.
    Output []Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output
}

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetEntityData() *types.CommonEntityData {
    command.EntityData.YFilter = command.YFilter
    command.EntityData.YangName = "command"
    command.EntityData.BundleName = "cisco_ios_xr"
    command.EntityData.ParentYangName = "commands"
    command.EntityData.SegmentPath = "command" + "[cmd='" + fmt.Sprintf("%v", command.Cmd) + "']"
    command.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    command.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    command.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    command.EntityData.Children = make(map[string]types.YChild)
    command.EntityData.Children["output"] = types.YChild{"Output", nil}
    for i := range command.Output {
        command.EntityData.Children[types.GetSegmentPath(&command.Output[i])] = types.YChild{"Output", &command.Output[i]}
    }
    command.EntityData.Leafs = make(map[string]types.YLeaf)
    command.EntityData.Leafs["cmd"] = types.YLeaf{"Cmd", command.Cmd}
    return &(command.EntityData)
}

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output
// Added to support datalist
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. First line. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Output interface{}

    // output xr. The type is string.
    OutputXr interface{}
}

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "command"
    output.EntityData.SegmentPath = "output" + "[output='" + fmt.Sprintf("%v", output.Output) + "']"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["output"] = types.YLeaf{"Output", output.Output}
    output.EntityData.Leafs["output-xr"] = types.YLeaf{"OutputXr", output.OutputXr}
    return &(output.EntityData)
}

// Fia_Nodes_Node_OirHistory
// FIA operational data of oir history
type Fia_Nodes_Node_OirHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flag table for history.
    Flags Fia_Nodes_Node_OirHistory_Flags
}

func (oirHistory *Fia_Nodes_Node_OirHistory) GetEntityData() *types.CommonEntityData {
    oirHistory.EntityData.YFilter = oirHistory.YFilter
    oirHistory.EntityData.YangName = "oir-history"
    oirHistory.EntityData.BundleName = "cisco_ios_xr"
    oirHistory.EntityData.ParentYangName = "node"
    oirHistory.EntityData.SegmentPath = "oir-history"
    oirHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oirHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oirHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oirHistory.EntityData.Children = make(map[string]types.YChild)
    oirHistory.EntityData.Children["flags"] = types.YChild{"Flags", &oirHistory.Flags}
    oirHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oirHistory.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags
// Flag table for history
type Fia_Nodes_Node_OirHistory_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flag value for physical location. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag.
    Flag []Fia_Nodes_Node_OirHistory_Flags_Flag
}

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xr"
    flags.EntityData.ParentYangName = "oir-history"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flags.EntityData.Children = make(map[string]types.YChild)
    flags.EntityData.Children["flag"] = types.YChild{"Flag", nil}
    for i := range flags.Flag {
        flags.EntityData.Children[types.GetSegmentPath(&flags.Flag[i])] = types.YChild{"Flag", &flags.Flag[i]}
    }
    flags.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flags.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag
// Flag value for physical location
type Fia_Nodes_Node_OirHistory_Flags_Flag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Flag value. The type is interface{} with range:
    // -2147483648..2147483647.
    Flag interface{}

    // Slot table for history.
    Slots Fia_Nodes_Node_OirHistory_Flags_Flag_Slots
}

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetEntityData() *types.CommonEntityData {
    flag.EntityData.YFilter = flag.YFilter
    flag.EntityData.YangName = "flag"
    flag.EntityData.BundleName = "cisco_ios_xr"
    flag.EntityData.ParentYangName = "flags"
    flag.EntityData.SegmentPath = "flag" + "[flag='" + fmt.Sprintf("%v", flag.Flag) + "']"
    flag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flag.EntityData.Children = make(map[string]types.YChild)
    flag.EntityData.Children["slots"] = types.YChild{"Slots", &flag.Slots}
    flag.EntityData.Leafs = make(map[string]types.YLeaf)
    flag.EntityData.Leafs["flag"] = types.YLeaf{"Flag", flag.Flag}
    return &(flag.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots
// Slot table for history
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot number for getting history. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot.
    Slot []Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot
}

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "flag"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = make(map[string]types.YChild)
    slots.EntityData.Children["slot"] = types.YChild{"Slot", nil}
    for i := range slots.Slot {
        slots.EntityData.Children[types.GetSegmentPath(&slots.Slot[i])] = types.YChild{"Slot", &slots.Slot[i]}
    }
    slots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slots.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot
// Slot number for getting history
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slot number. The type is interface{} with range:
    // -2147483648..2147483647.
    Slot interface{}

    // drv version. The type is interface{} with range: 0..4294967295.
    DrvVersion interface{}

    // coeff major rev. The type is interface{} with range: 0..4294967295.
    CoeffMajorRev interface{}

    // coeff minor rev. The type is interface{} with range: 0..4294967295.
    CoeffMinorRev interface{}

    // functional role. The type is interface{} with range: 0..255.
    FunctionalRole interface{}

    // issu role. The type is interface{} with range: 0..255.
    IssuRole interface{}

    // node id. The type is string.
    NodeId interface{}

    // rack type. The type is interface{} with range: -2147483648..2147483647.
    RackType interface{}

    // rack num. The type is interface{} with range: 0..255.
    RackNum interface{}

    // is driver ready. The type is bool.
    IsDriverReady interface{}

    // card avail mask. The type is interface{} with range: 0..4294967295.
    CardAvailMask interface{}

    // asic avail mask. The type is interface{} with range:
    // 0..18446744073709551615.
    AsicAvailMask interface{}

    // exp asic avail mask. The type is interface{} with range:
    // 0..18446744073709551615.
    ExpAsicAvailMask interface{}

    // ucmc ratio. The type is interface{} with range: 0..4294967295.
    UcmcRatio interface{}

    // asic oper notify to fsdb pending bmap. The type is interface{} with range:
    // 0..18446744073709551615.
    AsicOperNotifyToFsdbPendingBmap interface{}

    // is full fgid download req. The type is bool.
    IsFullFgidDownloadReq interface{}

    // is fgid download in progress. The type is bool.
    IsFgidDownloadInProgress interface{}

    // is fgid download completed. The type is bool.
    IsFgidDownloadCompleted interface{}

    // fsdb conn active. The type is bool.
    FsdbConnActive interface{}

    // fgid conn active. The type is bool.
    FgidConnActive interface{}

    // issu mgr conn active. The type is bool.
    IssuMgrConnActive interface{}

    // fsdb reg active. The type is bool.
    FsdbRegActive interface{}

    // fgid reg active. The type is bool.
    FgidRegActive interface{}

    // issu mgr reg active. The type is bool.
    IssuMgrRegActive interface{}

    // num pm conn reqs. The type is interface{} with range: 0..255.
    NumPmConnReqs interface{}

    // num fsdb conn reqs. The type is interface{} with range: 0..255.
    NumFsdbConnReqs interface{}

    // num fgid conn reqs. The type is interface{} with range: 0..255.
    NumFgidConnReqs interface{}

    // num fstats conn reqs. The type is interface{} with range: 0..255.
    NumFstatsConnReqs interface{}

    // num cm conn reqs. The type is interface{} with range: 0..255.
    NumCmConnReqs interface{}

    // num issu mgr conn reqs. The type is interface{} with range: 0..255.
    NumIssuMgrConnReqs interface{}

    // num peer fia conn reqs. The type is interface{} with range: 0..255.
    NumPeerFiaConnReqs interface{}

    // is gaspp registered. The type is bool.
    IsGasppRegistered interface{}

    // is cih registered. The type is bool.
    IsCihRegistered interface{}

    // drvr initial startup timestamp. The type is string.
    DrvrInitialStartupTimestamp interface{}

    // drvr current startup timestamp. The type is string.
    DrvrCurrentStartupTimestamp interface{}

    // num intf ports. The type is interface{} with range: 0..4294967295.
    NumIntfPorts interface{}

    // uc weight. The type is interface{} with range: 0..255.
    UcWeight interface{}

    // respawn count. The type is interface{} with range: 0..255.
    RespawnCount interface{}

    // total asics. The type is interface{} with range: 0..255.
    TotalAsics interface{}

    // issu ready ntfy pending. The type is bool.
    IssuReadyNtfyPending interface{}

    // issu abort sent. The type is bool.
    IssuAbortSent interface{}

    // issu abort rcvd. The type is bool.
    IssuAbortRcvd interface{}

    // fabric mode. The type is interface{} with range: 0..255.
    FabricMode interface{}

    // FC Mode. The type is FcMode.
    FcMode interface{}

    // board rev id. The type is interface{} with range: 0..4294967295.
    BoardRevId interface{}

    // device info. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo.
    DeviceInfo []Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo

    // card info. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo.
    CardInfo []Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo
}

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = make(map[string]types.YChild)
    slot.EntityData.Children["device-info"] = types.YChild{"DeviceInfo", nil}
    for i := range slot.DeviceInfo {
        slot.EntityData.Children[types.GetSegmentPath(&slot.DeviceInfo[i])] = types.YChild{"DeviceInfo", &slot.DeviceInfo[i]}
    }
    slot.EntityData.Children["card-info"] = types.YChild{"CardInfo", nil}
    for i := range slot.CardInfo {
        slot.EntityData.Children[types.GetSegmentPath(&slot.CardInfo[i])] = types.YChild{"CardInfo", &slot.CardInfo[i]}
    }
    slot.EntityData.Leafs = make(map[string]types.YLeaf)
    slot.EntityData.Leafs["slot"] = types.YLeaf{"Slot", slot.Slot}
    slot.EntityData.Leafs["drv-version"] = types.YLeaf{"DrvVersion", slot.DrvVersion}
    slot.EntityData.Leafs["coeff-major-rev"] = types.YLeaf{"CoeffMajorRev", slot.CoeffMajorRev}
    slot.EntityData.Leafs["coeff-minor-rev"] = types.YLeaf{"CoeffMinorRev", slot.CoeffMinorRev}
    slot.EntityData.Leafs["functional-role"] = types.YLeaf{"FunctionalRole", slot.FunctionalRole}
    slot.EntityData.Leafs["issu-role"] = types.YLeaf{"IssuRole", slot.IssuRole}
    slot.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", slot.NodeId}
    slot.EntityData.Leafs["rack-type"] = types.YLeaf{"RackType", slot.RackType}
    slot.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", slot.RackNum}
    slot.EntityData.Leafs["is-driver-ready"] = types.YLeaf{"IsDriverReady", slot.IsDriverReady}
    slot.EntityData.Leafs["card-avail-mask"] = types.YLeaf{"CardAvailMask", slot.CardAvailMask}
    slot.EntityData.Leafs["asic-avail-mask"] = types.YLeaf{"AsicAvailMask", slot.AsicAvailMask}
    slot.EntityData.Leafs["exp-asic-avail-mask"] = types.YLeaf{"ExpAsicAvailMask", slot.ExpAsicAvailMask}
    slot.EntityData.Leafs["ucmc-ratio"] = types.YLeaf{"UcmcRatio", slot.UcmcRatio}
    slot.EntityData.Leafs["asic-oper-notify-to-fsdb-pending-bmap"] = types.YLeaf{"AsicOperNotifyToFsdbPendingBmap", slot.AsicOperNotifyToFsdbPendingBmap}
    slot.EntityData.Leafs["is-full-fgid-download-req"] = types.YLeaf{"IsFullFgidDownloadReq", slot.IsFullFgidDownloadReq}
    slot.EntityData.Leafs["is-fgid-download-in-progress"] = types.YLeaf{"IsFgidDownloadInProgress", slot.IsFgidDownloadInProgress}
    slot.EntityData.Leafs["is-fgid-download-completed"] = types.YLeaf{"IsFgidDownloadCompleted", slot.IsFgidDownloadCompleted}
    slot.EntityData.Leafs["fsdb-conn-active"] = types.YLeaf{"FsdbConnActive", slot.FsdbConnActive}
    slot.EntityData.Leafs["fgid-conn-active"] = types.YLeaf{"FgidConnActive", slot.FgidConnActive}
    slot.EntityData.Leafs["issu-mgr-conn-active"] = types.YLeaf{"IssuMgrConnActive", slot.IssuMgrConnActive}
    slot.EntityData.Leafs["fsdb-reg-active"] = types.YLeaf{"FsdbRegActive", slot.FsdbRegActive}
    slot.EntityData.Leafs["fgid-reg-active"] = types.YLeaf{"FgidRegActive", slot.FgidRegActive}
    slot.EntityData.Leafs["issu-mgr-reg-active"] = types.YLeaf{"IssuMgrRegActive", slot.IssuMgrRegActive}
    slot.EntityData.Leafs["num-pm-conn-reqs"] = types.YLeaf{"NumPmConnReqs", slot.NumPmConnReqs}
    slot.EntityData.Leafs["num-fsdb-conn-reqs"] = types.YLeaf{"NumFsdbConnReqs", slot.NumFsdbConnReqs}
    slot.EntityData.Leafs["num-fgid-conn-reqs"] = types.YLeaf{"NumFgidConnReqs", slot.NumFgidConnReqs}
    slot.EntityData.Leafs["num-fstats-conn-reqs"] = types.YLeaf{"NumFstatsConnReqs", slot.NumFstatsConnReqs}
    slot.EntityData.Leafs["num-cm-conn-reqs"] = types.YLeaf{"NumCmConnReqs", slot.NumCmConnReqs}
    slot.EntityData.Leafs["num-issu-mgr-conn-reqs"] = types.YLeaf{"NumIssuMgrConnReqs", slot.NumIssuMgrConnReqs}
    slot.EntityData.Leafs["num-peer-fia-conn-reqs"] = types.YLeaf{"NumPeerFiaConnReqs", slot.NumPeerFiaConnReqs}
    slot.EntityData.Leafs["is-gaspp-registered"] = types.YLeaf{"IsGasppRegistered", slot.IsGasppRegistered}
    slot.EntityData.Leafs["is-cih-registered"] = types.YLeaf{"IsCihRegistered", slot.IsCihRegistered}
    slot.EntityData.Leafs["drvr-initial-startup-timestamp"] = types.YLeaf{"DrvrInitialStartupTimestamp", slot.DrvrInitialStartupTimestamp}
    slot.EntityData.Leafs["drvr-current-startup-timestamp"] = types.YLeaf{"DrvrCurrentStartupTimestamp", slot.DrvrCurrentStartupTimestamp}
    slot.EntityData.Leafs["num-intf-ports"] = types.YLeaf{"NumIntfPorts", slot.NumIntfPorts}
    slot.EntityData.Leafs["uc-weight"] = types.YLeaf{"UcWeight", slot.UcWeight}
    slot.EntityData.Leafs["respawn-count"] = types.YLeaf{"RespawnCount", slot.RespawnCount}
    slot.EntityData.Leafs["total-asics"] = types.YLeaf{"TotalAsics", slot.TotalAsics}
    slot.EntityData.Leafs["issu-ready-ntfy-pending"] = types.YLeaf{"IssuReadyNtfyPending", slot.IssuReadyNtfyPending}
    slot.EntityData.Leafs["issu-abort-sent"] = types.YLeaf{"IssuAbortSent", slot.IssuAbortSent}
    slot.EntityData.Leafs["issu-abort-rcvd"] = types.YLeaf{"IssuAbortRcvd", slot.IssuAbortRcvd}
    slot.EntityData.Leafs["fabric-mode"] = types.YLeaf{"FabricMode", slot.FabricMode}
    slot.EntityData.Leafs["fc-mode"] = types.YLeaf{"FcMode", slot.FcMode}
    slot.EntityData.Leafs["board-rev-id"] = types.YLeaf{"BoardRevId", slot.BoardRevId}
    return &(slot.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo
// device info
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // is valid. The type is bool.
    IsValid interface{}

    // fapid. The type is interface{} with range: 0..4294967295.
    Fapid interface{}

    // hotplug event. The type is interface{} with range: 0..4294967295.
    HotplugEvent interface{}

    // Slice State. The type is SliceState.
    SliceState interface{}

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is AsicOperState.
    OperState interface{}

    // Asic State. The type is AsicAccessState.
    AsicState interface{}

    // last init cause. The type is AsicInitMethod.
    LastInitCause interface{}

    // num pon resets. The type is interface{} with range: 0..4294967295.
    NumPonResets interface{}

    // num hard resets. The type is interface{} with range: 0..4294967295.
    NumHardResets interface{}

    // local switch state. The type is bool.
    LocalSwitchState interface{}

    // asic id.
    AsicId Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId
}

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetEntityData() *types.CommonEntityData {
    deviceInfo.EntityData.YFilter = deviceInfo.YFilter
    deviceInfo.EntityData.YangName = "device-info"
    deviceInfo.EntityData.BundleName = "cisco_ios_xr"
    deviceInfo.EntityData.ParentYangName = "slot"
    deviceInfo.EntityData.SegmentPath = "device-info"
    deviceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    deviceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    deviceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    deviceInfo.EntityData.Children = make(map[string]types.YChild)
    deviceInfo.EntityData.Children["asic-id"] = types.YChild{"AsicId", &deviceInfo.AsicId}
    deviceInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    deviceInfo.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", deviceInfo.IsValid}
    deviceInfo.EntityData.Leafs["fapid"] = types.YLeaf{"Fapid", deviceInfo.Fapid}
    deviceInfo.EntityData.Leafs["hotplug-event"] = types.YLeaf{"HotplugEvent", deviceInfo.HotplugEvent}
    deviceInfo.EntityData.Leafs["slice-state"] = types.YLeaf{"SliceState", deviceInfo.SliceState}
    deviceInfo.EntityData.Leafs["admin-state"] = types.YLeaf{"AdminState", deviceInfo.AdminState}
    deviceInfo.EntityData.Leafs["oper-state"] = types.YLeaf{"OperState", deviceInfo.OperState}
    deviceInfo.EntityData.Leafs["asic-state"] = types.YLeaf{"AsicState", deviceInfo.AsicState}
    deviceInfo.EntityData.Leafs["last-init-cause"] = types.YLeaf{"LastInitCause", deviceInfo.LastInitCause}
    deviceInfo.EntityData.Leafs["num-pon-resets"] = types.YLeaf{"NumPonResets", deviceInfo.NumPonResets}
    deviceInfo.EntityData.Leafs["num-hard-resets"] = types.YLeaf{"NumHardResets", deviceInfo.NumHardResets}
    deviceInfo.EntityData.Leafs["local-switch-state"] = types.YLeaf{"LocalSwitchState", deviceInfo.LocalSwitchState}
    return &(deviceInfo.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId
// asic id
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "device-info"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = make(map[string]types.YChild)
    asicId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicId.EntityData.Leafs["rack-type"] = types.YLeaf{"RackType", asicId.RackType}
    asicId.EntityData.Leafs["asic-type"] = types.YLeaf{"AsicType", asicId.AsicType}
    asicId.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", asicId.RackNum}
    asicId.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", asicId.SlotNum}
    asicId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicId.AsicInstance}
    return &(asicId.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo
// card info
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card type. The type is interface{} with range: -2147483648..2147483647.
    CardType interface{}

    // card name. The type is string.
    CardName interface{}

    // slot no. The type is interface{} with range: -2147483648..2147483647.
    SlotNo interface{}

    // card flag. The type is interface{} with range: -2147483648..2147483647.
    CardFlag interface{}

    // evt flag. The type is interface{} with range: -2147483648..2147483647.
    EvtFlag interface{}

    // reg flag. The type is interface{} with range: -2147483648..2147483647.
    RegFlag interface{}

    // instance. The type is interface{} with range: -2147483648..2147483647.
    Instance interface{}

    // card state. The type is interface{} with range: 0..255.
    CardState interface{}

    // exp num asics. The type is interface{} with range: 0..4294967295.
    ExpNumAsics interface{}

    // exp num asics per fsdb. The type is interface{} with range: 0..4294967295.
    ExpNumAsicsPerFsdb interface{}

    // is powered. The type is bool.
    IsPowered interface{}

    // cxp avail bitmap. The type is interface{} with range:
    // 0..18446744073709551615.
    CxpAvailBitmap interface{}

    // num ilkns per asic. The type is interface{} with range: 0..4294967295.
    NumIlknsPerAsic interface{}

    // num local ports per ilkn. The type is interface{} with range:
    // 0..4294967295.
    NumLocalPortsPerIlkn interface{}

    // num cos per port. The type is interface{} with range: 0..255.
    NumCosPerPort interface{}

    // oir circular buffer.
    OirCircularBuffer Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer
}

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetEntityData() *types.CommonEntityData {
    cardInfo.EntityData.YFilter = cardInfo.YFilter
    cardInfo.EntityData.YangName = "card-info"
    cardInfo.EntityData.BundleName = "cisco_ios_xr"
    cardInfo.EntityData.ParentYangName = "slot"
    cardInfo.EntityData.SegmentPath = "card-info"
    cardInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardInfo.EntityData.Children = make(map[string]types.YChild)
    cardInfo.EntityData.Children["oir-circular-buffer"] = types.YChild{"OirCircularBuffer", &cardInfo.OirCircularBuffer}
    cardInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    cardInfo.EntityData.Leafs["card-type"] = types.YLeaf{"CardType", cardInfo.CardType}
    cardInfo.EntityData.Leafs["card-name"] = types.YLeaf{"CardName", cardInfo.CardName}
    cardInfo.EntityData.Leafs["slot-no"] = types.YLeaf{"SlotNo", cardInfo.SlotNo}
    cardInfo.EntityData.Leafs["card-flag"] = types.YLeaf{"CardFlag", cardInfo.CardFlag}
    cardInfo.EntityData.Leafs["evt-flag"] = types.YLeaf{"EvtFlag", cardInfo.EvtFlag}
    cardInfo.EntityData.Leafs["reg-flag"] = types.YLeaf{"RegFlag", cardInfo.RegFlag}
    cardInfo.EntityData.Leafs["instance"] = types.YLeaf{"Instance", cardInfo.Instance}
    cardInfo.EntityData.Leafs["card-state"] = types.YLeaf{"CardState", cardInfo.CardState}
    cardInfo.EntityData.Leafs["exp-num-asics"] = types.YLeaf{"ExpNumAsics", cardInfo.ExpNumAsics}
    cardInfo.EntityData.Leafs["exp-num-asics-per-fsdb"] = types.YLeaf{"ExpNumAsicsPerFsdb", cardInfo.ExpNumAsicsPerFsdb}
    cardInfo.EntityData.Leafs["is-powered"] = types.YLeaf{"IsPowered", cardInfo.IsPowered}
    cardInfo.EntityData.Leafs["cxp-avail-bitmap"] = types.YLeaf{"CxpAvailBitmap", cardInfo.CxpAvailBitmap}
    cardInfo.EntityData.Leafs["num-ilkns-per-asic"] = types.YLeaf{"NumIlknsPerAsic", cardInfo.NumIlknsPerAsic}
    cardInfo.EntityData.Leafs["num-local-ports-per-ilkn"] = types.YLeaf{"NumLocalPortsPerIlkn", cardInfo.NumLocalPortsPerIlkn}
    cardInfo.EntityData.Leafs["num-cos-per-port"] = types.YLeaf{"NumCosPerPort", cardInfo.NumCosPerPort}
    return &(cardInfo.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer
// oir circular buffer
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // count. The type is interface{} with range: -2147483648..2147483647.
    Count interface{}

    // start. The type is interface{} with range: -2147483648..2147483647.
    Start interface{}

    // end. The type is interface{} with range: -2147483648..2147483647.
    End interface{}

    // fia oir info. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo.
    FiaOirInfo []Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo
}

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetEntityData() *types.CommonEntityData {
    oirCircularBuffer.EntityData.YFilter = oirCircularBuffer.YFilter
    oirCircularBuffer.EntityData.YangName = "oir-circular-buffer"
    oirCircularBuffer.EntityData.BundleName = "cisco_ios_xr"
    oirCircularBuffer.EntityData.ParentYangName = "card-info"
    oirCircularBuffer.EntityData.SegmentPath = "oir-circular-buffer"
    oirCircularBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oirCircularBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oirCircularBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oirCircularBuffer.EntityData.Children = make(map[string]types.YChild)
    oirCircularBuffer.EntityData.Children["fia-oir-info"] = types.YChild{"FiaOirInfo", nil}
    for i := range oirCircularBuffer.FiaOirInfo {
        oirCircularBuffer.EntityData.Children[types.GetSegmentPath(&oirCircularBuffer.FiaOirInfo[i])] = types.YChild{"FiaOirInfo", &oirCircularBuffer.FiaOirInfo[i]}
    }
    oirCircularBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    oirCircularBuffer.EntityData.Leafs["count"] = types.YLeaf{"Count", oirCircularBuffer.Count}
    oirCircularBuffer.EntityData.Leafs["start"] = types.YLeaf{"Start", oirCircularBuffer.Start}
    oirCircularBuffer.EntityData.Leafs["end"] = types.YLeaf{"End", oirCircularBuffer.End}
    return &(oirCircularBuffer.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo
// fia oir info
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // card flag. The type is interface{} with range: -2147483648..2147483647.
    CardFlag interface{}

    // card type. The type is interface{} with range: -2147483648..2147483647.
    CardType interface{}

    // reg flag. The type is interface{} with range: -2147483648..2147483647.
    RegFlag interface{}

    // evt flag. The type is interface{} with range: -2147483648..2147483647.
    EvtFlag interface{}

    // rack num. The type is interface{} with range: -2147483648..2147483647.
    RackNum interface{}

    // instance. The type is interface{} with range: -2147483648..2147483647.
    Instance interface{}

    // cur card state. The type is interface{} with range:
    // -2147483648..2147483647.
    CurCardState interface{}
}

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetEntityData() *types.CommonEntityData {
    fiaOirInfo.EntityData.YFilter = fiaOirInfo.YFilter
    fiaOirInfo.EntityData.YangName = "fia-oir-info"
    fiaOirInfo.EntityData.BundleName = "cisco_ios_xr"
    fiaOirInfo.EntityData.ParentYangName = "oir-circular-buffer"
    fiaOirInfo.EntityData.SegmentPath = "fia-oir-info"
    fiaOirInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fiaOirInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fiaOirInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fiaOirInfo.EntityData.Children = make(map[string]types.YChild)
    fiaOirInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fiaOirInfo.EntityData.Leafs["card-flag"] = types.YLeaf{"CardFlag", fiaOirInfo.CardFlag}
    fiaOirInfo.EntityData.Leafs["card-type"] = types.YLeaf{"CardType", fiaOirInfo.CardType}
    fiaOirInfo.EntityData.Leafs["reg-flag"] = types.YLeaf{"RegFlag", fiaOirInfo.RegFlag}
    fiaOirInfo.EntityData.Leafs["evt-flag"] = types.YLeaf{"EvtFlag", fiaOirInfo.EvtFlag}
    fiaOirInfo.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", fiaOirInfo.RackNum}
    fiaOirInfo.EntityData.Leafs["instance"] = types.YLeaf{"Instance", fiaOirInfo.Instance}
    fiaOirInfo.EntityData.Leafs["cur-card-state"] = types.YLeaf{"CurCardState", fiaOirInfo.CurCardState}
    return &(fiaOirInfo.EntityData)
}

// Fia_Nodes_Node_AsicStatistics
// FIA asic statistics information
type Fia_Nodes_Node_AsicStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance table for statistics.
    StatisticsAsicInstances Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances
}

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetEntityData() *types.CommonEntityData {
    asicStatistics.EntityData.YFilter = asicStatistics.YFilter
    asicStatistics.EntityData.YangName = "asic-statistics"
    asicStatistics.EntityData.BundleName = "cisco_ios_xr"
    asicStatistics.EntityData.ParentYangName = "node"
    asicStatistics.EntityData.SegmentPath = "asic-statistics"
    asicStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatistics.EntityData.Children = make(map[string]types.YChild)
    asicStatistics.EntityData.Children["statistics-asic-instances"] = types.YChild{"StatisticsAsicInstances", &asicStatistics.StatisticsAsicInstances}
    asicStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicStatistics.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances
// Instance table for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Asic instance for statistics. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance.
    StatisticsAsicInstance []Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance
}

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetEntityData() *types.CommonEntityData {
    statisticsAsicInstances.EntityData.YFilter = statisticsAsicInstances.YFilter
    statisticsAsicInstances.EntityData.YangName = "statistics-asic-instances"
    statisticsAsicInstances.EntityData.BundleName = "cisco_ios_xr"
    statisticsAsicInstances.EntityData.ParentYangName = "asic-statistics"
    statisticsAsicInstances.EntityData.SegmentPath = "statistics-asic-instances"
    statisticsAsicInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsAsicInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsAsicInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsAsicInstances.EntityData.Children = make(map[string]types.YChild)
    statisticsAsicInstances.EntityData.Children["statistics-asic-instance"] = types.YChild{"StatisticsAsicInstance", nil}
    for i := range statisticsAsicInstances.StatisticsAsicInstance {
        statisticsAsicInstances.EntityData.Children[types.GetSegmentPath(&statisticsAsicInstances.StatisticsAsicInstance[i])] = types.YChild{"StatisticsAsicInstance", &statisticsAsicInstances.StatisticsAsicInstance[i]}
    }
    statisticsAsicInstances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statisticsAsicInstances.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance
// Asic instance for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Asic instance. The type is interface{} with range:
    // 0..255.
    Instance interface{}

    // Packet Byte Counter for a Asic.
    PbcStatistics Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics

    // Statistics of FMAC.
    FmacStatistics Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics
}

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetEntityData() *types.CommonEntityData {
    statisticsAsicInstance.EntityData.YFilter = statisticsAsicInstance.YFilter
    statisticsAsicInstance.EntityData.YangName = "statistics-asic-instance"
    statisticsAsicInstance.EntityData.BundleName = "cisco_ios_xr"
    statisticsAsicInstance.EntityData.ParentYangName = "statistics-asic-instances"
    statisticsAsicInstance.EntityData.SegmentPath = "statistics-asic-instance" + "[instance='" + fmt.Sprintf("%v", statisticsAsicInstance.Instance) + "']"
    statisticsAsicInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsAsicInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsAsicInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsAsicInstance.EntityData.Children = make(map[string]types.YChild)
    statisticsAsicInstance.EntityData.Children["pbc-statistics"] = types.YChild{"PbcStatistics", &statisticsAsicInstance.PbcStatistics}
    statisticsAsicInstance.EntityData.Children["fmac-statistics"] = types.YChild{"FmacStatistics", &statisticsAsicInstance.FmacStatistics}
    statisticsAsicInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    statisticsAsicInstance.EntityData.Leafs["instance"] = types.YLeaf{"Instance", statisticsAsicInstance.Instance}
    return &(statisticsAsicInstance.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics
// Packet Byte Counter for a Asic
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBC stats bag.
    PbcStats Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats
}

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetEntityData() *types.CommonEntityData {
    pbcStatistics.EntityData.YFilter = pbcStatistics.YFilter
    pbcStatistics.EntityData.YangName = "pbc-statistics"
    pbcStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbcStatistics.EntityData.ParentYangName = "statistics-asic-instance"
    pbcStatistics.EntityData.SegmentPath = "pbc-statistics"
    pbcStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbcStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbcStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbcStatistics.EntityData.Children = make(map[string]types.YChild)
    pbcStatistics.EntityData.Children["pbc-stats"] = types.YChild{"PbcStats", &pbcStatistics.PbcStats}
    pbcStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pbcStatistics.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats
// PBC stats bag
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // valid. The type is bool.
    Valid interface{}

    // rack no. The type is interface{} with range: 0..4294967295.
    RackNo interface{}

    // slot no. The type is interface{} with range: 0..4294967295.
    SlotNo interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}

    // chip ver. The type is interface{} with range: 0..65535.
    ChipVer interface{}

    // stats info.
    StatsInfo Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo
}

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetEntityData() *types.CommonEntityData {
    pbcStats.EntityData.YFilter = pbcStats.YFilter
    pbcStats.EntityData.YangName = "pbc-stats"
    pbcStats.EntityData.BundleName = "cisco_ios_xr"
    pbcStats.EntityData.ParentYangName = "pbc-statistics"
    pbcStats.EntityData.SegmentPath = "pbc-stats"
    pbcStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbcStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbcStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbcStats.EntityData.Children = make(map[string]types.YChild)
    pbcStats.EntityData.Children["stats-info"] = types.YChild{"StatsInfo", &pbcStats.StatsInfo}
    pbcStats.EntityData.Leafs = make(map[string]types.YLeaf)
    pbcStats.EntityData.Leafs["valid"] = types.YLeaf{"Valid", pbcStats.Valid}
    pbcStats.EntityData.Leafs["rack-no"] = types.YLeaf{"RackNo", pbcStats.RackNo}
    pbcStats.EntityData.Leafs["slot-no"] = types.YLeaf{"SlotNo", pbcStats.SlotNo}
    pbcStats.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", pbcStats.AsicInstance}
    pbcStats.EntityData.Leafs["chip-ver"] = types.YLeaf{"ChipVer", pbcStats.ChipVer}
    return &(pbcStats.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo
// stats info
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Num Blocks. The type is interface{} with range: 0..255.
    NumBlocks interface{}

    // block info. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo.
    BlockInfo []Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo
}

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetEntityData() *types.CommonEntityData {
    statsInfo.EntityData.YFilter = statsInfo.YFilter
    statsInfo.EntityData.YangName = "stats-info"
    statsInfo.EntityData.BundleName = "cisco_ios_xr"
    statsInfo.EntityData.ParentYangName = "pbc-stats"
    statsInfo.EntityData.SegmentPath = "stats-info"
    statsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsInfo.EntityData.Children = make(map[string]types.YChild)
    statsInfo.EntityData.Children["block-info"] = types.YChild{"BlockInfo", nil}
    for i := range statsInfo.BlockInfo {
        statsInfo.EntityData.Children[types.GetSegmentPath(&statsInfo.BlockInfo[i])] = types.YChild{"BlockInfo", &statsInfo.BlockInfo[i]}
    }
    statsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    statsInfo.EntityData.Leafs["num-blocks"] = types.YLeaf{"NumBlocks", statsInfo.NumBlocks}
    return &(statsInfo.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo
// block info
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Block Name. The type is string with length: 0..10.
    BlockName interface{}

    // Num Fields. The type is interface{} with range: 0..255.
    NumFields interface{}

    // field info. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo.
    FieldInfo []Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo
}

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetEntityData() *types.CommonEntityData {
    blockInfo.EntityData.YFilter = blockInfo.YFilter
    blockInfo.EntityData.YangName = "block-info"
    blockInfo.EntityData.BundleName = "cisco_ios_xr"
    blockInfo.EntityData.ParentYangName = "stats-info"
    blockInfo.EntityData.SegmentPath = "block-info"
    blockInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockInfo.EntityData.Children = make(map[string]types.YChild)
    blockInfo.EntityData.Children["field-info"] = types.YChild{"FieldInfo", nil}
    for i := range blockInfo.FieldInfo {
        blockInfo.EntityData.Children[types.GetSegmentPath(&blockInfo.FieldInfo[i])] = types.YChild{"FieldInfo", &blockInfo.FieldInfo[i]}
    }
    blockInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    blockInfo.EntityData.Leafs["block-name"] = types.YLeaf{"BlockName", blockInfo.BlockName}
    blockInfo.EntityData.Leafs["num-fields"] = types.YLeaf{"NumFields", blockInfo.NumFields}
    return &(blockInfo.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo
// field info
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Field Name. The type is string with length: 0..80.
    FieldName interface{}

    // Field Value. The type is interface{} with range: 0..18446744073709551615.
    FieldValue interface{}

    // Is Ovf. The type is bool.
    IsOvf interface{}
}

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetEntityData() *types.CommonEntityData {
    fieldInfo.EntityData.YFilter = fieldInfo.YFilter
    fieldInfo.EntityData.YangName = "field-info"
    fieldInfo.EntityData.BundleName = "cisco_ios_xr"
    fieldInfo.EntityData.ParentYangName = "block-info"
    fieldInfo.EntityData.SegmentPath = "field-info"
    fieldInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fieldInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fieldInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fieldInfo.EntityData.Children = make(map[string]types.YChild)
    fieldInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fieldInfo.EntityData.Leafs["field-name"] = types.YLeaf{"FieldName", fieldInfo.FieldName}
    fieldInfo.EntityData.Leafs["field-value"] = types.YLeaf{"FieldValue", fieldInfo.FieldValue}
    fieldInfo.EntityData.Leafs["is-ovf"] = types.YLeaf{"IsOvf", fieldInfo.IsOvf}
    return &(fieldInfo.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics
// Statistics of FMAC
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link table for statistics.
    FmacLinks Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks
}

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetEntityData() *types.CommonEntityData {
    fmacStatistics.EntityData.YFilter = fmacStatistics.YFilter
    fmacStatistics.EntityData.YangName = "fmac-statistics"
    fmacStatistics.EntityData.BundleName = "cisco_ios_xr"
    fmacStatistics.EntityData.ParentYangName = "statistics-asic-instance"
    fmacStatistics.EntityData.SegmentPath = "fmac-statistics"
    fmacStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmacStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmacStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmacStatistics.EntityData.Children = make(map[string]types.YChild)
    fmacStatistics.EntityData.Children["fmac-links"] = types.YChild{"FmacLinks", &fmacStatistics.FmacLinks}
    fmacStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fmacStatistics.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks
// Link table for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link number for statistics. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink.
    FmacLink []Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink
}

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetEntityData() *types.CommonEntityData {
    fmacLinks.EntityData.YFilter = fmacLinks.YFilter
    fmacLinks.EntityData.YangName = "fmac-links"
    fmacLinks.EntityData.BundleName = "cisco_ios_xr"
    fmacLinks.EntityData.ParentYangName = "fmac-statistics"
    fmacLinks.EntityData.SegmentPath = "fmac-links"
    fmacLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmacLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmacLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmacLinks.EntityData.Children = make(map[string]types.YChild)
    fmacLinks.EntityData.Children["fmac-link"] = types.YChild{"FmacLink", nil}
    for i := range fmacLinks.FmacLink {
        fmacLinks.EntityData.Children[types.GetSegmentPath(&fmacLinks.FmacLink[i])] = types.YChild{"FmacLink", &fmacLinks.FmacLink[i]}
    }
    fmacLinks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fmacLinks.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink
// Link number for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link number. The type is interface{} with range:
    // -2147483648..2147483647.
    Link interface{}

    // Single aisc information. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic.
    FmacAsic []Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic
}

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetEntityData() *types.CommonEntityData {
    fmacLink.EntityData.YFilter = fmacLink.YFilter
    fmacLink.EntityData.YangName = "fmac-link"
    fmacLink.EntityData.BundleName = "cisco_ios_xr"
    fmacLink.EntityData.ParentYangName = "fmac-links"
    fmacLink.EntityData.SegmentPath = "fmac-link" + "[link='" + fmt.Sprintf("%v", fmacLink.Link) + "']"
    fmacLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmacLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmacLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmacLink.EntityData.Children = make(map[string]types.YChild)
    fmacLink.EntityData.Children["fmac-asic"] = types.YChild{"FmacAsic", nil}
    for i := range fmacLink.FmacAsic {
        fmacLink.EntityData.Children[types.GetSegmentPath(&fmacLink.FmacAsic[i])] = types.YChild{"FmacAsic", &fmacLink.FmacAsic[i]}
    }
    fmacLink.EntityData.Leafs = make(map[string]types.YLeaf)
    fmacLink.EntityData.Leafs["link"] = types.YLeaf{"Link", fmacLink.Link}
    return &(fmacLink.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic
// Single aisc information
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Single asic. The type is interface{} with range:
    // -2147483648..2147483647.
    Asic interface{}

    // valid. The type is bool.
    Valid interface{}

    // rack no. The type is interface{} with range: 0..4294967295.
    RackNo interface{}

    // slot no. The type is interface{} with range: 0..4294967295.
    SlotNo interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}

    // link no. The type is interface{} with range: 0..4294967295.
    LinkNo interface{}

    // link valid. The type is bool.
    LinkValid interface{}

    // aggr stats.
    AggrStats Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats

    // incr stats.
    IncrStats Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats
}

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetEntityData() *types.CommonEntityData {
    fmacAsic.EntityData.YFilter = fmacAsic.YFilter
    fmacAsic.EntityData.YangName = "fmac-asic"
    fmacAsic.EntityData.BundleName = "cisco_ios_xr"
    fmacAsic.EntityData.ParentYangName = "fmac-link"
    fmacAsic.EntityData.SegmentPath = "fmac-asic" + "[asic='" + fmt.Sprintf("%v", fmacAsic.Asic) + "']"
    fmacAsic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmacAsic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmacAsic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmacAsic.EntityData.Children = make(map[string]types.YChild)
    fmacAsic.EntityData.Children["aggr-stats"] = types.YChild{"AggrStats", &fmacAsic.AggrStats}
    fmacAsic.EntityData.Children["incr-stats"] = types.YChild{"IncrStats", &fmacAsic.IncrStats}
    fmacAsic.EntityData.Leafs = make(map[string]types.YLeaf)
    fmacAsic.EntityData.Leafs["asic"] = types.YLeaf{"Asic", fmacAsic.Asic}
    fmacAsic.EntityData.Leafs["valid"] = types.YLeaf{"Valid", fmacAsic.Valid}
    fmacAsic.EntityData.Leafs["rack-no"] = types.YLeaf{"RackNo", fmacAsic.RackNo}
    fmacAsic.EntityData.Leafs["slot-no"] = types.YLeaf{"SlotNo", fmacAsic.SlotNo}
    fmacAsic.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", fmacAsic.AsicInstance}
    fmacAsic.EntityData.Leafs["link-no"] = types.YLeaf{"LinkNo", fmacAsic.LinkNo}
    fmacAsic.EntityData.Leafs["link-valid"] = types.YLeaf{"LinkValid", fmacAsic.LinkValid}
    return &(fmacAsic.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats
// aggr stats
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // link error status.
    LinkErrorStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus

    // link counters.
    LinkCounters Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters

    // ovf status.
    OvfStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus
}

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetEntityData() *types.CommonEntityData {
    aggrStats.EntityData.YFilter = aggrStats.YFilter
    aggrStats.EntityData.YangName = "aggr-stats"
    aggrStats.EntityData.BundleName = "cisco_ios_xr"
    aggrStats.EntityData.ParentYangName = "fmac-asic"
    aggrStats.EntityData.SegmentPath = "aggr-stats"
    aggrStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggrStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggrStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggrStats.EntityData.Children = make(map[string]types.YChild)
    aggrStats.EntityData.Children["link-error-status"] = types.YChild{"LinkErrorStatus", &aggrStats.LinkErrorStatus}
    aggrStats.EntityData.Children["link-counters"] = types.YChild{"LinkCounters", &aggrStats.LinkCounters}
    aggrStats.EntityData.Children["ovf-status"] = types.YChild{"OvfStatus", &aggrStats.OvfStatus}
    aggrStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aggrStats.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus
// link error status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // link crc error. The type is interface{} with range: 0..4294967295.
    LinkCrcError interface{}

    // link size error. The type is interface{} with range: 0..4294967295.
    LinkSizeError interface{}

    // link mis align error. The type is interface{} with range: 0..4294967295.
    LinkMisAlignError interface{}

    // link code group error. The type is interface{} with range: 0..4294967295.
    LinkCodeGroupError interface{}

    // link no sig lock error. The type is interface{} with range: 0..4294967295.
    LinkNoSigLockError interface{}

    // link no sig accept error. The type is interface{} with range:
    // 0..4294967295.
    LinkNoSigAcceptError interface{}

    // link tokens error. The type is interface{} with range: 0..4294967295.
    LinkTokensError interface{}

    // error token count. The type is interface{} with range: 0..4294967295.
    ErrorTokenCount interface{}
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetEntityData() *types.CommonEntityData {
    linkErrorStatus.EntityData.YFilter = linkErrorStatus.YFilter
    linkErrorStatus.EntityData.YangName = "link-error-status"
    linkErrorStatus.EntityData.BundleName = "cisco_ios_xr"
    linkErrorStatus.EntityData.ParentYangName = "aggr-stats"
    linkErrorStatus.EntityData.SegmentPath = "link-error-status"
    linkErrorStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkErrorStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkErrorStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkErrorStatus.EntityData.Children = make(map[string]types.YChild)
    linkErrorStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    linkErrorStatus.EntityData.Leafs["link-crc-error"] = types.YLeaf{"LinkCrcError", linkErrorStatus.LinkCrcError}
    linkErrorStatus.EntityData.Leafs["link-size-error"] = types.YLeaf{"LinkSizeError", linkErrorStatus.LinkSizeError}
    linkErrorStatus.EntityData.Leafs["link-mis-align-error"] = types.YLeaf{"LinkMisAlignError", linkErrorStatus.LinkMisAlignError}
    linkErrorStatus.EntityData.Leafs["link-code-group-error"] = types.YLeaf{"LinkCodeGroupError", linkErrorStatus.LinkCodeGroupError}
    linkErrorStatus.EntityData.Leafs["link-no-sig-lock-error"] = types.YLeaf{"LinkNoSigLockError", linkErrorStatus.LinkNoSigLockError}
    linkErrorStatus.EntityData.Leafs["link-no-sig-accept-error"] = types.YLeaf{"LinkNoSigAcceptError", linkErrorStatus.LinkNoSigAcceptError}
    linkErrorStatus.EntityData.Leafs["link-tokens-error"] = types.YLeaf{"LinkTokensError", linkErrorStatus.LinkTokensError}
    linkErrorStatus.EntityData.Leafs["error-token-count"] = types.YLeaf{"ErrorTokenCount", linkErrorStatus.ErrorTokenCount}
    return &(linkErrorStatus.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters
// link counters
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TX Control cells counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxControlCellsCounter interface{}

    // TX Data cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxDataCellCounter interface{}

    // TX Data byte counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxDataByteCounter interface{}

    // RX CRC errors counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxCrcErrorsCounter interface{}

    // RX LFEC FEC correctable error. The type is interface{} with range:
    // 0..18446744073709551615.
    RxLfecFecCorrectableError interface{}

    // RX 8b 10b disparity errors. The type is interface{} with range:
    // 0..18446744073709551615.
    Rx8B10BDisparityErrors interface{}

    // RX Control cells counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxControlCellsCounter interface{}

    // RX Data cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDataCellCounter interface{}

    // RX Data byte counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDataByteCounter interface{}

    // RX dropped retransmitted control. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDroppedRetransmittedControl interface{}

    // TX Asyn fifo rate. The type is interface{} with range:
    // 0..18446744073709551615.
    TxAsynFifoRate interface{}

    // RX Asyn fifo rate. The type is interface{} with range:
    // 0..18446744073709551615.
    RxAsynFifoRate interface{}

    // RX LFEC FEC uncorrectable errors. The type is interface{} with range:
    // 0..18446744073709551615.
    RxLfecFecUncorrectableErrors interface{}

    // RX 8b 10b code errors. The type is interface{} with range:
    // 0..18446744073709551615.
    Rx8B10BCodeErrors interface{}
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetEntityData() *types.CommonEntityData {
    linkCounters.EntityData.YFilter = linkCounters.YFilter
    linkCounters.EntityData.YangName = "link-counters"
    linkCounters.EntityData.BundleName = "cisco_ios_xr"
    linkCounters.EntityData.ParentYangName = "aggr-stats"
    linkCounters.EntityData.SegmentPath = "link-counters"
    linkCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkCounters.EntityData.Children = make(map[string]types.YChild)
    linkCounters.EntityData.Leafs = make(map[string]types.YLeaf)
    linkCounters.EntityData.Leafs["tx-control-cells-counter"] = types.YLeaf{"TxControlCellsCounter", linkCounters.TxControlCellsCounter}
    linkCounters.EntityData.Leafs["tx-data-cell-counter"] = types.YLeaf{"TxDataCellCounter", linkCounters.TxDataCellCounter}
    linkCounters.EntityData.Leafs["tx-data-byte-counter"] = types.YLeaf{"TxDataByteCounter", linkCounters.TxDataByteCounter}
    linkCounters.EntityData.Leafs["rx-crc-errors-counter"] = types.YLeaf{"RxCrcErrorsCounter", linkCounters.RxCrcErrorsCounter}
    linkCounters.EntityData.Leafs["rx-lfec-fec-correctable-error"] = types.YLeaf{"RxLfecFecCorrectableError", linkCounters.RxLfecFecCorrectableError}
    linkCounters.EntityData.Leafs["rx-8b-10b-disparity-errors"] = types.YLeaf{"Rx8B10BDisparityErrors", linkCounters.Rx8B10BDisparityErrors}
    linkCounters.EntityData.Leafs["rx-control-cells-counter"] = types.YLeaf{"RxControlCellsCounter", linkCounters.RxControlCellsCounter}
    linkCounters.EntityData.Leafs["rx-data-cell-counter"] = types.YLeaf{"RxDataCellCounter", linkCounters.RxDataCellCounter}
    linkCounters.EntityData.Leafs["rx-data-byte-counter"] = types.YLeaf{"RxDataByteCounter", linkCounters.RxDataByteCounter}
    linkCounters.EntityData.Leafs["rx-dropped-retransmitted-control"] = types.YLeaf{"RxDroppedRetransmittedControl", linkCounters.RxDroppedRetransmittedControl}
    linkCounters.EntityData.Leafs["tx-asyn-fifo-rate"] = types.YLeaf{"TxAsynFifoRate", linkCounters.TxAsynFifoRate}
    linkCounters.EntityData.Leafs["rx-asyn-fifo-rate"] = types.YLeaf{"RxAsynFifoRate", linkCounters.RxAsynFifoRate}
    linkCounters.EntityData.Leafs["rx-lfec-fec-uncorrectable-errors"] = types.YLeaf{"RxLfecFecUncorrectableErrors", linkCounters.RxLfecFecUncorrectableErrors}
    linkCounters.EntityData.Leafs["rx-8b-10b-code-errors"] = types.YLeaf{"Rx8B10BCodeErrors", linkCounters.Rx8B10BCodeErrors}
    return &(linkCounters.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus
// ovf status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TX Control cells counter. The type is string with length: 0..6.
    TxControlCellsCounter interface{}

    // TX Data cell counter. The type is string with length: 0..6.
    TxDataCellCounter interface{}

    // TX Data byte counter. The type is string with length: 0..6.
    TxDataByteCounter interface{}

    // RX CRC errors counter. The type is string with length: 0..6.
    RxCrcErrorsCounter interface{}

    // RX LFEC FEC correctable error. The type is string with length: 0..6.
    RxLfecFecCorrectableError interface{}

    // RX 8b 10b disparity errors. The type is string with length: 0..6.
    Rx8B10BDisparityErrors interface{}

    // RX Control cells counter. The type is string with length: 0..6.
    RxControlCellsCounter interface{}

    // RX Data cell counter. The type is string with length: 0..6.
    RxDataCellCounter interface{}

    // RX Data byte counter. The type is string with length: 0..6.
    RxDataByteCounter interface{}

    // RX dropped retransmitted control. The type is string with length: 0..6.
    RxDroppedRetransmittedControl interface{}

    // TX Asyn fifo rate. The type is string with length: 0..6.
    TxAsynFifoRate interface{}

    // RX Asyn fifo rate. The type is string with length: 0..6.
    RxAsynFifoRate interface{}

    // RX LFEC FEC uncorrectable errors. The type is string with length: 0..6.
    RxLfecFecUncorrectableErrors interface{}

    // RX 8b 10b code errors. The type is string with length: 0..6.
    Rx8B10BCodeErrors interface{}
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetEntityData() *types.CommonEntityData {
    ovfStatus.EntityData.YFilter = ovfStatus.YFilter
    ovfStatus.EntityData.YangName = "ovf-status"
    ovfStatus.EntityData.BundleName = "cisco_ios_xr"
    ovfStatus.EntityData.ParentYangName = "aggr-stats"
    ovfStatus.EntityData.SegmentPath = "ovf-status"
    ovfStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ovfStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ovfStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ovfStatus.EntityData.Children = make(map[string]types.YChild)
    ovfStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    ovfStatus.EntityData.Leafs["tx-control-cells-counter"] = types.YLeaf{"TxControlCellsCounter", ovfStatus.TxControlCellsCounter}
    ovfStatus.EntityData.Leafs["tx-data-cell-counter"] = types.YLeaf{"TxDataCellCounter", ovfStatus.TxDataCellCounter}
    ovfStatus.EntityData.Leafs["tx-data-byte-counter"] = types.YLeaf{"TxDataByteCounter", ovfStatus.TxDataByteCounter}
    ovfStatus.EntityData.Leafs["rx-crc-errors-counter"] = types.YLeaf{"RxCrcErrorsCounter", ovfStatus.RxCrcErrorsCounter}
    ovfStatus.EntityData.Leafs["rx-lfec-fec-correctable-error"] = types.YLeaf{"RxLfecFecCorrectableError", ovfStatus.RxLfecFecCorrectableError}
    ovfStatus.EntityData.Leafs["rx-8b-10b-disparity-errors"] = types.YLeaf{"Rx8B10BDisparityErrors", ovfStatus.Rx8B10BDisparityErrors}
    ovfStatus.EntityData.Leafs["rx-control-cells-counter"] = types.YLeaf{"RxControlCellsCounter", ovfStatus.RxControlCellsCounter}
    ovfStatus.EntityData.Leafs["rx-data-cell-counter"] = types.YLeaf{"RxDataCellCounter", ovfStatus.RxDataCellCounter}
    ovfStatus.EntityData.Leafs["rx-data-byte-counter"] = types.YLeaf{"RxDataByteCounter", ovfStatus.RxDataByteCounter}
    ovfStatus.EntityData.Leafs["rx-dropped-retransmitted-control"] = types.YLeaf{"RxDroppedRetransmittedControl", ovfStatus.RxDroppedRetransmittedControl}
    ovfStatus.EntityData.Leafs["tx-asyn-fifo-rate"] = types.YLeaf{"TxAsynFifoRate", ovfStatus.TxAsynFifoRate}
    ovfStatus.EntityData.Leafs["rx-asyn-fifo-rate"] = types.YLeaf{"RxAsynFifoRate", ovfStatus.RxAsynFifoRate}
    ovfStatus.EntityData.Leafs["rx-lfec-fec-uncorrectable-errors"] = types.YLeaf{"RxLfecFecUncorrectableErrors", ovfStatus.RxLfecFecUncorrectableErrors}
    ovfStatus.EntityData.Leafs["rx-8b-10b-code-errors"] = types.YLeaf{"Rx8B10BCodeErrors", ovfStatus.Rx8B10BCodeErrors}
    return &(ovfStatus.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats
// incr stats
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // link error status.
    LinkErrorStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus

    // link counters.
    LinkCounters Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters

    // ovf status.
    OvfStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus
}

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetEntityData() *types.CommonEntityData {
    incrStats.EntityData.YFilter = incrStats.YFilter
    incrStats.EntityData.YangName = "incr-stats"
    incrStats.EntityData.BundleName = "cisco_ios_xr"
    incrStats.EntityData.ParentYangName = "fmac-asic"
    incrStats.EntityData.SegmentPath = "incr-stats"
    incrStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incrStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incrStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incrStats.EntityData.Children = make(map[string]types.YChild)
    incrStats.EntityData.Children["link-error-status"] = types.YChild{"LinkErrorStatus", &incrStats.LinkErrorStatus}
    incrStats.EntityData.Children["link-counters"] = types.YChild{"LinkCounters", &incrStats.LinkCounters}
    incrStats.EntityData.Children["ovf-status"] = types.YChild{"OvfStatus", &incrStats.OvfStatus}
    incrStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(incrStats.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus
// link error status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // link crc error. The type is interface{} with range: 0..4294967295.
    LinkCrcError interface{}

    // link size error. The type is interface{} with range: 0..4294967295.
    LinkSizeError interface{}

    // link mis align error. The type is interface{} with range: 0..4294967295.
    LinkMisAlignError interface{}

    // link code group error. The type is interface{} with range: 0..4294967295.
    LinkCodeGroupError interface{}

    // link no sig lock error. The type is interface{} with range: 0..4294967295.
    LinkNoSigLockError interface{}

    // link no sig accept error. The type is interface{} with range:
    // 0..4294967295.
    LinkNoSigAcceptError interface{}

    // link tokens error. The type is interface{} with range: 0..4294967295.
    LinkTokensError interface{}

    // error token count. The type is interface{} with range: 0..4294967295.
    ErrorTokenCount interface{}
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetEntityData() *types.CommonEntityData {
    linkErrorStatus.EntityData.YFilter = linkErrorStatus.YFilter
    linkErrorStatus.EntityData.YangName = "link-error-status"
    linkErrorStatus.EntityData.BundleName = "cisco_ios_xr"
    linkErrorStatus.EntityData.ParentYangName = "incr-stats"
    linkErrorStatus.EntityData.SegmentPath = "link-error-status"
    linkErrorStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkErrorStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkErrorStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkErrorStatus.EntityData.Children = make(map[string]types.YChild)
    linkErrorStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    linkErrorStatus.EntityData.Leafs["link-crc-error"] = types.YLeaf{"LinkCrcError", linkErrorStatus.LinkCrcError}
    linkErrorStatus.EntityData.Leafs["link-size-error"] = types.YLeaf{"LinkSizeError", linkErrorStatus.LinkSizeError}
    linkErrorStatus.EntityData.Leafs["link-mis-align-error"] = types.YLeaf{"LinkMisAlignError", linkErrorStatus.LinkMisAlignError}
    linkErrorStatus.EntityData.Leafs["link-code-group-error"] = types.YLeaf{"LinkCodeGroupError", linkErrorStatus.LinkCodeGroupError}
    linkErrorStatus.EntityData.Leafs["link-no-sig-lock-error"] = types.YLeaf{"LinkNoSigLockError", linkErrorStatus.LinkNoSigLockError}
    linkErrorStatus.EntityData.Leafs["link-no-sig-accept-error"] = types.YLeaf{"LinkNoSigAcceptError", linkErrorStatus.LinkNoSigAcceptError}
    linkErrorStatus.EntityData.Leafs["link-tokens-error"] = types.YLeaf{"LinkTokensError", linkErrorStatus.LinkTokensError}
    linkErrorStatus.EntityData.Leafs["error-token-count"] = types.YLeaf{"ErrorTokenCount", linkErrorStatus.ErrorTokenCount}
    return &(linkErrorStatus.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters
// link counters
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TX Control cells counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxControlCellsCounter interface{}

    // TX Data cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxDataCellCounter interface{}

    // TX Data byte counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxDataByteCounter interface{}

    // RX CRC errors counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxCrcErrorsCounter interface{}

    // RX LFEC FEC correctable error. The type is interface{} with range:
    // 0..18446744073709551615.
    RxLfecFecCorrectableError interface{}

    // RX 8b 10b disparity errors. The type is interface{} with range:
    // 0..18446744073709551615.
    Rx8B10BDisparityErrors interface{}

    // RX Control cells counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxControlCellsCounter interface{}

    // RX Data cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDataCellCounter interface{}

    // RX Data byte counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDataByteCounter interface{}

    // RX dropped retransmitted control. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDroppedRetransmittedControl interface{}

    // TX Asyn fifo rate. The type is interface{} with range:
    // 0..18446744073709551615.
    TxAsynFifoRate interface{}

    // RX Asyn fifo rate. The type is interface{} with range:
    // 0..18446744073709551615.
    RxAsynFifoRate interface{}

    // RX LFEC FEC uncorrectable errors. The type is interface{} with range:
    // 0..18446744073709551615.
    RxLfecFecUncorrectableErrors interface{}

    // RX 8b 10b code errors. The type is interface{} with range:
    // 0..18446744073709551615.
    Rx8B10BCodeErrors interface{}
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetEntityData() *types.CommonEntityData {
    linkCounters.EntityData.YFilter = linkCounters.YFilter
    linkCounters.EntityData.YangName = "link-counters"
    linkCounters.EntityData.BundleName = "cisco_ios_xr"
    linkCounters.EntityData.ParentYangName = "incr-stats"
    linkCounters.EntityData.SegmentPath = "link-counters"
    linkCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkCounters.EntityData.Children = make(map[string]types.YChild)
    linkCounters.EntityData.Leafs = make(map[string]types.YLeaf)
    linkCounters.EntityData.Leafs["tx-control-cells-counter"] = types.YLeaf{"TxControlCellsCounter", linkCounters.TxControlCellsCounter}
    linkCounters.EntityData.Leafs["tx-data-cell-counter"] = types.YLeaf{"TxDataCellCounter", linkCounters.TxDataCellCounter}
    linkCounters.EntityData.Leafs["tx-data-byte-counter"] = types.YLeaf{"TxDataByteCounter", linkCounters.TxDataByteCounter}
    linkCounters.EntityData.Leafs["rx-crc-errors-counter"] = types.YLeaf{"RxCrcErrorsCounter", linkCounters.RxCrcErrorsCounter}
    linkCounters.EntityData.Leafs["rx-lfec-fec-correctable-error"] = types.YLeaf{"RxLfecFecCorrectableError", linkCounters.RxLfecFecCorrectableError}
    linkCounters.EntityData.Leafs["rx-8b-10b-disparity-errors"] = types.YLeaf{"Rx8B10BDisparityErrors", linkCounters.Rx8B10BDisparityErrors}
    linkCounters.EntityData.Leafs["rx-control-cells-counter"] = types.YLeaf{"RxControlCellsCounter", linkCounters.RxControlCellsCounter}
    linkCounters.EntityData.Leafs["rx-data-cell-counter"] = types.YLeaf{"RxDataCellCounter", linkCounters.RxDataCellCounter}
    linkCounters.EntityData.Leafs["rx-data-byte-counter"] = types.YLeaf{"RxDataByteCounter", linkCounters.RxDataByteCounter}
    linkCounters.EntityData.Leafs["rx-dropped-retransmitted-control"] = types.YLeaf{"RxDroppedRetransmittedControl", linkCounters.RxDroppedRetransmittedControl}
    linkCounters.EntityData.Leafs["tx-asyn-fifo-rate"] = types.YLeaf{"TxAsynFifoRate", linkCounters.TxAsynFifoRate}
    linkCounters.EntityData.Leafs["rx-asyn-fifo-rate"] = types.YLeaf{"RxAsynFifoRate", linkCounters.RxAsynFifoRate}
    linkCounters.EntityData.Leafs["rx-lfec-fec-uncorrectable-errors"] = types.YLeaf{"RxLfecFecUncorrectableErrors", linkCounters.RxLfecFecUncorrectableErrors}
    linkCounters.EntityData.Leafs["rx-8b-10b-code-errors"] = types.YLeaf{"Rx8B10BCodeErrors", linkCounters.Rx8B10BCodeErrors}
    return &(linkCounters.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus
// ovf status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TX Control cells counter. The type is string with length: 0..6.
    TxControlCellsCounter interface{}

    // TX Data cell counter. The type is string with length: 0..6.
    TxDataCellCounter interface{}

    // TX Data byte counter. The type is string with length: 0..6.
    TxDataByteCounter interface{}

    // RX CRC errors counter. The type is string with length: 0..6.
    RxCrcErrorsCounter interface{}

    // RX LFEC FEC correctable error. The type is string with length: 0..6.
    RxLfecFecCorrectableError interface{}

    // RX 8b 10b disparity errors. The type is string with length: 0..6.
    Rx8B10BDisparityErrors interface{}

    // RX Control cells counter. The type is string with length: 0..6.
    RxControlCellsCounter interface{}

    // RX Data cell counter. The type is string with length: 0..6.
    RxDataCellCounter interface{}

    // RX Data byte counter. The type is string with length: 0..6.
    RxDataByteCounter interface{}

    // RX dropped retransmitted control. The type is string with length: 0..6.
    RxDroppedRetransmittedControl interface{}

    // TX Asyn fifo rate. The type is string with length: 0..6.
    TxAsynFifoRate interface{}

    // RX Asyn fifo rate. The type is string with length: 0..6.
    RxAsynFifoRate interface{}

    // RX LFEC FEC uncorrectable errors. The type is string with length: 0..6.
    RxLfecFecUncorrectableErrors interface{}

    // RX 8b 10b code errors. The type is string with length: 0..6.
    Rx8B10BCodeErrors interface{}
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetEntityData() *types.CommonEntityData {
    ovfStatus.EntityData.YFilter = ovfStatus.YFilter
    ovfStatus.EntityData.YangName = "ovf-status"
    ovfStatus.EntityData.BundleName = "cisco_ios_xr"
    ovfStatus.EntityData.ParentYangName = "incr-stats"
    ovfStatus.EntityData.SegmentPath = "ovf-status"
    ovfStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ovfStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ovfStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ovfStatus.EntityData.Children = make(map[string]types.YChild)
    ovfStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    ovfStatus.EntityData.Leafs["tx-control-cells-counter"] = types.YLeaf{"TxControlCellsCounter", ovfStatus.TxControlCellsCounter}
    ovfStatus.EntityData.Leafs["tx-data-cell-counter"] = types.YLeaf{"TxDataCellCounter", ovfStatus.TxDataCellCounter}
    ovfStatus.EntityData.Leafs["tx-data-byte-counter"] = types.YLeaf{"TxDataByteCounter", ovfStatus.TxDataByteCounter}
    ovfStatus.EntityData.Leafs["rx-crc-errors-counter"] = types.YLeaf{"RxCrcErrorsCounter", ovfStatus.RxCrcErrorsCounter}
    ovfStatus.EntityData.Leafs["rx-lfec-fec-correctable-error"] = types.YLeaf{"RxLfecFecCorrectableError", ovfStatus.RxLfecFecCorrectableError}
    ovfStatus.EntityData.Leafs["rx-8b-10b-disparity-errors"] = types.YLeaf{"Rx8B10BDisparityErrors", ovfStatus.Rx8B10BDisparityErrors}
    ovfStatus.EntityData.Leafs["rx-control-cells-counter"] = types.YLeaf{"RxControlCellsCounter", ovfStatus.RxControlCellsCounter}
    ovfStatus.EntityData.Leafs["rx-data-cell-counter"] = types.YLeaf{"RxDataCellCounter", ovfStatus.RxDataCellCounter}
    ovfStatus.EntityData.Leafs["rx-data-byte-counter"] = types.YLeaf{"RxDataByteCounter", ovfStatus.RxDataByteCounter}
    ovfStatus.EntityData.Leafs["rx-dropped-retransmitted-control"] = types.YLeaf{"RxDroppedRetransmittedControl", ovfStatus.RxDroppedRetransmittedControl}
    ovfStatus.EntityData.Leafs["tx-asyn-fifo-rate"] = types.YLeaf{"TxAsynFifoRate", ovfStatus.TxAsynFifoRate}
    ovfStatus.EntityData.Leafs["rx-asyn-fifo-rate"] = types.YLeaf{"RxAsynFifoRate", ovfStatus.RxAsynFifoRate}
    ovfStatus.EntityData.Leafs["rx-lfec-fec-uncorrectable-errors"] = types.YLeaf{"RxLfecFecUncorrectableErrors", ovfStatus.RxLfecFecUncorrectableErrors}
    ovfStatus.EntityData.Leafs["rx-8b-10b-code-errors"] = types.YLeaf{"Rx8B10BCodeErrors", ovfStatus.Rx8B10BCodeErrors}
    return &(ovfStatus.EntityData)
}

