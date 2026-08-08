# ControllerManagedDeviceGroupWirelessNetworkAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ControllerManagedDeviceGroup** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**WirelessNetwork** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Vlan** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 

## Methods

### NewControllerManagedDeviceGroupWirelessNetworkAssignment

`func NewControllerManagedDeviceGroupWirelessNetworkAssignment(objectType string, display string, url string, naturalSlug string, controllerManagedDeviceGroup ApprovalWorkflowStageResponseApprovalWorkflowStage, wirelessNetwork ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *ControllerManagedDeviceGroupWirelessNetworkAssignment`

NewControllerManagedDeviceGroupWirelessNetworkAssignment instantiates a new ControllerManagedDeviceGroupWirelessNetworkAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerManagedDeviceGroupWirelessNetworkAssignmentWithDefaults

`func NewControllerManagedDeviceGroupWirelessNetworkAssignmentWithDefaults() *ControllerManagedDeviceGroupWirelessNetworkAssignment`

NewControllerManagedDeviceGroupWirelessNetworkAssignmentWithDefaults instantiates a new ControllerManagedDeviceGroupWirelessNetworkAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetControllerManagedDeviceGroup

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetControllerManagedDeviceGroup() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetControllerManagedDeviceGroup returns the ControllerManagedDeviceGroup field if non-nil, zero value otherwise.

### GetControllerManagedDeviceGroupOk

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetControllerManagedDeviceGroupOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetControllerManagedDeviceGroupOk returns a tuple with the ControllerManagedDeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerManagedDeviceGroup

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) SetControllerManagedDeviceGroup(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetControllerManagedDeviceGroup sets ControllerManagedDeviceGroup field to given value.


### GetWirelessNetwork

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetWirelessNetwork() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetWirelessNetwork returns the WirelessNetwork field if non-nil, zero value otherwise.

### GetWirelessNetworkOk

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetWirelessNetworkOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetWirelessNetworkOk returns a tuple with the WirelessNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessNetwork

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) SetWirelessNetwork(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetWirelessNetwork sets WirelessNetwork field to given value.


### GetVlan

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetVlan() ApprovalWorkflowUser`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) GetVlanOk() (*ApprovalWorkflowUser, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) SetVlan(v ApprovalWorkflowUser)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *ControllerManagedDeviceGroupWirelessNetworkAssignment) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


