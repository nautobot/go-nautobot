# PatchedBulkWritablePowerPanelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PanelType** | Pointer to [**PanelTypeEnum**](PanelTypeEnum.md) |  | [optional] 
**PowerPath** | Pointer to [**PowerPathEnum**](PowerPathEnum.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BreakerPositionCount** | Pointer to **NullableInt32** | Total number of breaker positions in the panel (e.g., 42) | [optional] 
**Location** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**RackGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewPatchedBulkWritablePowerPanelRequest

`func NewPatchedBulkWritablePowerPanelRequest(id string, ) *PatchedBulkWritablePowerPanelRequest`

NewPatchedBulkWritablePowerPanelRequest instantiates a new PatchedBulkWritablePowerPanelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritablePowerPanelRequestWithDefaults

`func NewPatchedBulkWritablePowerPanelRequestWithDefaults() *PatchedBulkWritablePowerPanelRequest`

NewPatchedBulkWritablePowerPanelRequestWithDefaults instantiates a new PatchedBulkWritablePowerPanelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritablePowerPanelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritablePowerPanelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritablePowerPanelRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPanelType

`func (o *PatchedBulkWritablePowerPanelRequest) GetPanelType() PanelTypeEnum`

GetPanelType returns the PanelType field if non-nil, zero value otherwise.

### GetPanelTypeOk

`func (o *PatchedBulkWritablePowerPanelRequest) GetPanelTypeOk() (*PanelTypeEnum, bool)`

GetPanelTypeOk returns a tuple with the PanelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelType

`func (o *PatchedBulkWritablePowerPanelRequest) SetPanelType(v PanelTypeEnum)`

SetPanelType sets PanelType field to given value.

### HasPanelType

`func (o *PatchedBulkWritablePowerPanelRequest) HasPanelType() bool`

HasPanelType returns a boolean if a field has been set.

### GetPowerPath

`func (o *PatchedBulkWritablePowerPanelRequest) GetPowerPath() PowerPathEnum`

GetPowerPath returns the PowerPath field if non-nil, zero value otherwise.

### GetPowerPathOk

`func (o *PatchedBulkWritablePowerPanelRequest) GetPowerPathOk() (*PowerPathEnum, bool)`

GetPowerPathOk returns a tuple with the PowerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPath

`func (o *PatchedBulkWritablePowerPanelRequest) SetPowerPath(v PowerPathEnum)`

SetPowerPath sets PowerPath field to given value.

### HasPowerPath

`func (o *PatchedBulkWritablePowerPanelRequest) HasPowerPath() bool`

HasPowerPath returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritablePowerPanelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritablePowerPanelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritablePowerPanelRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritablePowerPanelRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBreakerPositionCount

`func (o *PatchedBulkWritablePowerPanelRequest) GetBreakerPositionCount() int32`

GetBreakerPositionCount returns the BreakerPositionCount field if non-nil, zero value otherwise.

### GetBreakerPositionCountOk

`func (o *PatchedBulkWritablePowerPanelRequest) GetBreakerPositionCountOk() (*int32, bool)`

GetBreakerPositionCountOk returns a tuple with the BreakerPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakerPositionCount

`func (o *PatchedBulkWritablePowerPanelRequest) SetBreakerPositionCount(v int32)`

SetBreakerPositionCount sets BreakerPositionCount field to given value.

### HasBreakerPositionCount

`func (o *PatchedBulkWritablePowerPanelRequest) HasBreakerPositionCount() bool`

HasBreakerPositionCount returns a boolean if a field has been set.

### SetBreakerPositionCountNil

`func (o *PatchedBulkWritablePowerPanelRequest) SetBreakerPositionCountNil(b bool)`

 SetBreakerPositionCountNil sets the value for BreakerPositionCount to be an explicit nil

### UnsetBreakerPositionCount
`func (o *PatchedBulkWritablePowerPanelRequest) UnsetBreakerPositionCount()`

UnsetBreakerPositionCount ensures that no value is present for BreakerPositionCount, not even an explicit nil
### GetLocation

`func (o *PatchedBulkWritablePowerPanelRequest) GetLocation() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedBulkWritablePowerPanelRequest) GetLocationOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedBulkWritablePowerPanelRequest) SetLocation(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedBulkWritablePowerPanelRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRackGroup

`func (o *PatchedBulkWritablePowerPanelRequest) GetRackGroup() ApprovalWorkflowUser`

GetRackGroup returns the RackGroup field if non-nil, zero value otherwise.

### GetRackGroupOk

`func (o *PatchedBulkWritablePowerPanelRequest) GetRackGroupOk() (*ApprovalWorkflowUser, bool)`

GetRackGroupOk returns a tuple with the RackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackGroup

`func (o *PatchedBulkWritablePowerPanelRequest) SetRackGroup(v ApprovalWorkflowUser)`

SetRackGroup sets RackGroup field to given value.

### HasRackGroup

`func (o *PatchedBulkWritablePowerPanelRequest) HasRackGroup() bool`

HasRackGroup returns a boolean if a field has been set.

### SetRackGroupNil

`func (o *PatchedBulkWritablePowerPanelRequest) SetRackGroupNil(b bool)`

 SetRackGroupNil sets the value for RackGroup to be an explicit nil

### UnsetRackGroup
`func (o *PatchedBulkWritablePowerPanelRequest) UnsetRackGroup()`

UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritablePowerPanelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritablePowerPanelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritablePowerPanelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritablePowerPanelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritablePowerPanelRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritablePowerPanelRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritablePowerPanelRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritablePowerPanelRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritablePowerPanelRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritablePowerPanelRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritablePowerPanelRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritablePowerPanelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


