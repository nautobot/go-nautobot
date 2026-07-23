# WritablePowerPanelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**PanelType** | Pointer to [**PatchedWritablePowerPanelRequestPanelType**](PatchedWritablePowerPanelRequestPanelType.md) |  | [optional] 
**BreakerPositionCount** | Pointer to **NullableInt32** | Total number of breaker positions in the panel (e.g., 42) | [optional] 
**PowerPath** | Pointer to [**PatchedWritablePowerFeedRequestPowerPath**](PatchedWritablePowerFeedRequestPowerPath.md) |  | [optional] 
**Location** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**RackGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewWritablePowerPanelRequest

`func NewWritablePowerPanelRequest(name string, location BulkWritableCableRequestStatus, ) *WritablePowerPanelRequest`

NewWritablePowerPanelRequest instantiates a new WritablePowerPanelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePowerPanelRequestWithDefaults

`func NewWritablePowerPanelRequestWithDefaults() *WritablePowerPanelRequest`

NewWritablePowerPanelRequestWithDefaults instantiates a new WritablePowerPanelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritablePowerPanelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritablePowerPanelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritablePowerPanelRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WritablePowerPanelRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WritablePowerPanelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePowerPanelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePowerPanelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPanelType

`func (o *WritablePowerPanelRequest) GetPanelType() PatchedWritablePowerPanelRequestPanelType`

GetPanelType returns the PanelType field if non-nil, zero value otherwise.

### GetPanelTypeOk

`func (o *WritablePowerPanelRequest) GetPanelTypeOk() (*PatchedWritablePowerPanelRequestPanelType, bool)`

GetPanelTypeOk returns a tuple with the PanelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelType

`func (o *WritablePowerPanelRequest) SetPanelType(v PatchedWritablePowerPanelRequestPanelType)`

SetPanelType sets PanelType field to given value.

### HasPanelType

`func (o *WritablePowerPanelRequest) HasPanelType() bool`

HasPanelType returns a boolean if a field has been set.

### GetBreakerPositionCount

`func (o *WritablePowerPanelRequest) GetBreakerPositionCount() int32`

GetBreakerPositionCount returns the BreakerPositionCount field if non-nil, zero value otherwise.

### GetBreakerPositionCountOk

`func (o *WritablePowerPanelRequest) GetBreakerPositionCountOk() (*int32, bool)`

GetBreakerPositionCountOk returns a tuple with the BreakerPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakerPositionCount

`func (o *WritablePowerPanelRequest) SetBreakerPositionCount(v int32)`

SetBreakerPositionCount sets BreakerPositionCount field to given value.

### HasBreakerPositionCount

`func (o *WritablePowerPanelRequest) HasBreakerPositionCount() bool`

HasBreakerPositionCount returns a boolean if a field has been set.

### SetBreakerPositionCountNil

`func (o *WritablePowerPanelRequest) SetBreakerPositionCountNil(b bool)`

 SetBreakerPositionCountNil sets the value for BreakerPositionCount to be an explicit nil

### UnsetBreakerPositionCount
`func (o *WritablePowerPanelRequest) UnsetBreakerPositionCount()`

UnsetBreakerPositionCount ensures that no value is present for BreakerPositionCount, not even an explicit nil
### GetPowerPath

`func (o *WritablePowerPanelRequest) GetPowerPath() PatchedWritablePowerFeedRequestPowerPath`

GetPowerPath returns the PowerPath field if non-nil, zero value otherwise.

### GetPowerPathOk

`func (o *WritablePowerPanelRequest) GetPowerPathOk() (*PatchedWritablePowerFeedRequestPowerPath, bool)`

GetPowerPathOk returns a tuple with the PowerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPath

`func (o *WritablePowerPanelRequest) SetPowerPath(v PatchedWritablePowerFeedRequestPowerPath)`

SetPowerPath sets PowerPath field to given value.

### HasPowerPath

`func (o *WritablePowerPanelRequest) HasPowerPath() bool`

HasPowerPath returns a boolean if a field has been set.

### GetLocation

`func (o *WritablePowerPanelRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WritablePowerPanelRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WritablePowerPanelRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.


### GetRackGroup

`func (o *WritablePowerPanelRequest) GetRackGroup() ApprovalWorkflowUser`

GetRackGroup returns the RackGroup field if non-nil, zero value otherwise.

### GetRackGroupOk

`func (o *WritablePowerPanelRequest) GetRackGroupOk() (*ApprovalWorkflowUser, bool)`

GetRackGroupOk returns a tuple with the RackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackGroup

`func (o *WritablePowerPanelRequest) SetRackGroup(v ApprovalWorkflowUser)`

SetRackGroup sets RackGroup field to given value.

### HasRackGroup

`func (o *WritablePowerPanelRequest) HasRackGroup() bool`

HasRackGroup returns a boolean if a field has been set.

### SetRackGroupNil

`func (o *WritablePowerPanelRequest) SetRackGroupNil(b bool)`

 SetRackGroupNil sets the value for RackGroup to be an explicit nil

### UnsetRackGroup
`func (o *WritablePowerPanelRequest) UnsetRackGroup()`

UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
### GetCustomFields

`func (o *WritablePowerPanelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePowerPanelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePowerPanelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePowerPanelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritablePowerPanelRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritablePowerPanelRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritablePowerPanelRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritablePowerPanelRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *WritablePowerPanelRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritablePowerPanelRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritablePowerPanelRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritablePowerPanelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


