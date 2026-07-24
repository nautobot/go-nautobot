# BulkWritablePowerPanelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PanelType** | Pointer to [**PanelTypeEnum**](PanelTypeEnum.md) |  | [optional] 
**PowerPath** | Pointer to [**PowerPathEnum**](PowerPathEnum.md) |  | [optional] 
**Name** | **string** |  | 
**BreakerPositionCount** | Pointer to **NullableInt32** | Total number of breaker positions in the panel (e.g., 42) | [optional] 
**Location** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**RackGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewBulkWritablePowerPanelRequest

`func NewBulkWritablePowerPanelRequest(id string, name string, location ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *BulkWritablePowerPanelRequest`

NewBulkWritablePowerPanelRequest instantiates a new BulkWritablePowerPanelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritablePowerPanelRequestWithDefaults

`func NewBulkWritablePowerPanelRequestWithDefaults() *BulkWritablePowerPanelRequest`

NewBulkWritablePowerPanelRequestWithDefaults instantiates a new BulkWritablePowerPanelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritablePowerPanelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritablePowerPanelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritablePowerPanelRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPanelType

`func (o *BulkWritablePowerPanelRequest) GetPanelType() PanelTypeEnum`

GetPanelType returns the PanelType field if non-nil, zero value otherwise.

### GetPanelTypeOk

`func (o *BulkWritablePowerPanelRequest) GetPanelTypeOk() (*PanelTypeEnum, bool)`

GetPanelTypeOk returns a tuple with the PanelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelType

`func (o *BulkWritablePowerPanelRequest) SetPanelType(v PanelTypeEnum)`

SetPanelType sets PanelType field to given value.

### HasPanelType

`func (o *BulkWritablePowerPanelRequest) HasPanelType() bool`

HasPanelType returns a boolean if a field has been set.

### GetPowerPath

`func (o *BulkWritablePowerPanelRequest) GetPowerPath() PowerPathEnum`

GetPowerPath returns the PowerPath field if non-nil, zero value otherwise.

### GetPowerPathOk

`func (o *BulkWritablePowerPanelRequest) GetPowerPathOk() (*PowerPathEnum, bool)`

GetPowerPathOk returns a tuple with the PowerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPath

`func (o *BulkWritablePowerPanelRequest) SetPowerPath(v PowerPathEnum)`

SetPowerPath sets PowerPath field to given value.

### HasPowerPath

`func (o *BulkWritablePowerPanelRequest) HasPowerPath() bool`

HasPowerPath returns a boolean if a field has been set.

### GetName

`func (o *BulkWritablePowerPanelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritablePowerPanelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritablePowerPanelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBreakerPositionCount

`func (o *BulkWritablePowerPanelRequest) GetBreakerPositionCount() int32`

GetBreakerPositionCount returns the BreakerPositionCount field if non-nil, zero value otherwise.

### GetBreakerPositionCountOk

`func (o *BulkWritablePowerPanelRequest) GetBreakerPositionCountOk() (*int32, bool)`

GetBreakerPositionCountOk returns a tuple with the BreakerPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakerPositionCount

`func (o *BulkWritablePowerPanelRequest) SetBreakerPositionCount(v int32)`

SetBreakerPositionCount sets BreakerPositionCount field to given value.

### HasBreakerPositionCount

`func (o *BulkWritablePowerPanelRequest) HasBreakerPositionCount() bool`

HasBreakerPositionCount returns a boolean if a field has been set.

### SetBreakerPositionCountNil

`func (o *BulkWritablePowerPanelRequest) SetBreakerPositionCountNil(b bool)`

 SetBreakerPositionCountNil sets the value for BreakerPositionCount to be an explicit nil

### UnsetBreakerPositionCount
`func (o *BulkWritablePowerPanelRequest) UnsetBreakerPositionCount()`

UnsetBreakerPositionCount ensures that no value is present for BreakerPositionCount, not even an explicit nil
### GetLocation

`func (o *BulkWritablePowerPanelRequest) GetLocation() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BulkWritablePowerPanelRequest) GetLocationOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BulkWritablePowerPanelRequest) SetLocation(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLocation sets Location field to given value.


### GetRackGroup

`func (o *BulkWritablePowerPanelRequest) GetRackGroup() ApprovalWorkflowUser`

GetRackGroup returns the RackGroup field if non-nil, zero value otherwise.

### GetRackGroupOk

`func (o *BulkWritablePowerPanelRequest) GetRackGroupOk() (*ApprovalWorkflowUser, bool)`

GetRackGroupOk returns a tuple with the RackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackGroup

`func (o *BulkWritablePowerPanelRequest) SetRackGroup(v ApprovalWorkflowUser)`

SetRackGroup sets RackGroup field to given value.

### HasRackGroup

`func (o *BulkWritablePowerPanelRequest) HasRackGroup() bool`

HasRackGroup returns a boolean if a field has been set.

### SetRackGroupNil

`func (o *BulkWritablePowerPanelRequest) SetRackGroupNil(b bool)`

 SetRackGroupNil sets the value for RackGroup to be an explicit nil

### UnsetRackGroup
`func (o *BulkWritablePowerPanelRequest) UnsetRackGroup()`

UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
### GetCustomFields

`func (o *BulkWritablePowerPanelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritablePowerPanelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritablePowerPanelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritablePowerPanelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritablePowerPanelRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritablePowerPanelRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritablePowerPanelRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritablePowerPanelRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritablePowerPanelRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritablePowerPanelRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritablePowerPanelRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritablePowerPanelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


