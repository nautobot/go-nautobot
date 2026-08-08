# WritablePowerPortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PatchedWritablePowerPortRequestType**](PatchedWritablePowerPortRequestType.md) |  | [optional] 
**MaximumDraw** | Pointer to **NullableInt32** | Maximum power draw (watts) | [optional] 
**AllocatedDraw** | Pointer to **NullableInt32** | Allocated power draw (watts) | [optional] 
**PowerFactor** | Pointer to **float64** | Power factor (0.01-1.00) for converting between watts (W) and volt-amps (VA). Defaults to 0.95. | [optional] 
**Device** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Module** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewWritablePowerPortRequest

`func NewWritablePowerPortRequest(name string, ) *WritablePowerPortRequest`

NewWritablePowerPortRequest instantiates a new WritablePowerPortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePowerPortRequestWithDefaults

`func NewWritablePowerPortRequestWithDefaults() *WritablePowerPortRequest`

NewWritablePowerPortRequestWithDefaults instantiates a new WritablePowerPortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritablePowerPortRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritablePowerPortRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritablePowerPortRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WritablePowerPortRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WritablePowerPortRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePowerPortRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePowerPortRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritablePowerPortRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritablePowerPortRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritablePowerPortRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritablePowerPortRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *WritablePowerPortRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritablePowerPortRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritablePowerPortRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritablePowerPortRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *WritablePowerPortRequest) GetType() PatchedWritablePowerPortRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritablePowerPortRequest) GetTypeOk() (*PatchedWritablePowerPortRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritablePowerPortRequest) SetType(v PatchedWritablePowerPortRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *WritablePowerPortRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMaximumDraw

`func (o *WritablePowerPortRequest) GetMaximumDraw() int32`

GetMaximumDraw returns the MaximumDraw field if non-nil, zero value otherwise.

### GetMaximumDrawOk

`func (o *WritablePowerPortRequest) GetMaximumDrawOk() (*int32, bool)`

GetMaximumDrawOk returns a tuple with the MaximumDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDraw

`func (o *WritablePowerPortRequest) SetMaximumDraw(v int32)`

SetMaximumDraw sets MaximumDraw field to given value.

### HasMaximumDraw

`func (o *WritablePowerPortRequest) HasMaximumDraw() bool`

HasMaximumDraw returns a boolean if a field has been set.

### SetMaximumDrawNil

`func (o *WritablePowerPortRequest) SetMaximumDrawNil(b bool)`

 SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil

### UnsetMaximumDraw
`func (o *WritablePowerPortRequest) UnsetMaximumDraw()`

UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
### GetAllocatedDraw

`func (o *WritablePowerPortRequest) GetAllocatedDraw() int32`

GetAllocatedDraw returns the AllocatedDraw field if non-nil, zero value otherwise.

### GetAllocatedDrawOk

`func (o *WritablePowerPortRequest) GetAllocatedDrawOk() (*int32, bool)`

GetAllocatedDrawOk returns a tuple with the AllocatedDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedDraw

`func (o *WritablePowerPortRequest) SetAllocatedDraw(v int32)`

SetAllocatedDraw sets AllocatedDraw field to given value.

### HasAllocatedDraw

`func (o *WritablePowerPortRequest) HasAllocatedDraw() bool`

HasAllocatedDraw returns a boolean if a field has been set.

### SetAllocatedDrawNil

`func (o *WritablePowerPortRequest) SetAllocatedDrawNil(b bool)`

 SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil

### UnsetAllocatedDraw
`func (o *WritablePowerPortRequest) UnsetAllocatedDraw()`

UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
### GetPowerFactor

`func (o *WritablePowerPortRequest) GetPowerFactor() float64`

GetPowerFactor returns the PowerFactor field if non-nil, zero value otherwise.

### GetPowerFactorOk

`func (o *WritablePowerPortRequest) GetPowerFactorOk() (*float64, bool)`

GetPowerFactorOk returns a tuple with the PowerFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFactor

`func (o *WritablePowerPortRequest) SetPowerFactor(v float64)`

SetPowerFactor sets PowerFactor field to given value.

### HasPowerFactor

`func (o *WritablePowerPortRequest) HasPowerFactor() bool`

HasPowerFactor returns a boolean if a field has been set.

### GetDevice

`func (o *WritablePowerPortRequest) GetDevice() ApprovalWorkflowUser`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritablePowerPortRequest) GetDeviceOk() (*ApprovalWorkflowUser, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritablePowerPortRequest) SetDevice(v ApprovalWorkflowUser)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *WritablePowerPortRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *WritablePowerPortRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *WritablePowerPortRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetModule

`func (o *WritablePowerPortRequest) GetModule() ApprovalWorkflowUser`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WritablePowerPortRequest) GetModuleOk() (*ApprovalWorkflowUser, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WritablePowerPortRequest) SetModule(v ApprovalWorkflowUser)`

SetModule sets Module field to given value.

### HasModule

`func (o *WritablePowerPortRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *WritablePowerPortRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *WritablePowerPortRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetCustomFields

`func (o *WritablePowerPortRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePowerPortRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePowerPortRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePowerPortRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritablePowerPortRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritablePowerPortRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritablePowerPortRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritablePowerPortRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *WritablePowerPortRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritablePowerPortRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritablePowerPortRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritablePowerPortRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


