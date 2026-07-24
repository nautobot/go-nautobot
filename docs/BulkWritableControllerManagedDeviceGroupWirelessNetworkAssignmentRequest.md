# BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ControllerManagedDeviceGroup** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**WirelessNetwork** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Vlan** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 

## Methods

### NewBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest

`func NewBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest(id string, controllerManagedDeviceGroup ApprovalWorkflowStageResponseApprovalWorkflowStage, wirelessNetwork ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest`

NewBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest instantiates a new BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequestWithDefaults

`func NewBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequestWithDefaults() *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest`

NewBulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequestWithDefaults instantiates a new BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetControllerManagedDeviceGroup

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetControllerManagedDeviceGroup() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetControllerManagedDeviceGroup returns the ControllerManagedDeviceGroup field if non-nil, zero value otherwise.

### GetControllerManagedDeviceGroupOk

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetControllerManagedDeviceGroupOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetControllerManagedDeviceGroupOk returns a tuple with the ControllerManagedDeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerManagedDeviceGroup

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetControllerManagedDeviceGroup(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetControllerManagedDeviceGroup sets ControllerManagedDeviceGroup field to given value.


### GetWirelessNetwork

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetWirelessNetwork() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetWirelessNetwork returns the WirelessNetwork field if non-nil, zero value otherwise.

### GetWirelessNetworkOk

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetWirelessNetworkOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetWirelessNetworkOk returns a tuple with the WirelessNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessNetwork

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetWirelessNetwork(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetWirelessNetwork sets WirelessNetwork field to given value.


### GetVlan

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetVlan() ApprovalWorkflowUser`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) GetVlanOk() (*ApprovalWorkflowUser, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetVlan(v ApprovalWorkflowUser)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *BulkWritableControllerManagedDeviceGroupWirelessNetworkAssignmentRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


