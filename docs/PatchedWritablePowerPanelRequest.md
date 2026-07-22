# PatchedWritablePowerPanelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PanelType** | Pointer to [**PatchedWritablePowerPanelRequestPanelType**](PatchedWritablePowerPanelRequestPanelType.md) |  | [optional] 
**BreakerPositionCount** | Pointer to **NullableInt32** | Total number of breaker positions in the panel (e.g., 42) | [optional] 
**PowerPath** | Pointer to [**PatchedWritablePowerFeedRequestPowerPath**](PatchedWritablePowerFeedRequestPowerPath.md) |  | [optional] 
**Location** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**RackGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedWritablePowerPanelRequest

`func NewPatchedWritablePowerPanelRequest() *PatchedWritablePowerPanelRequest`

NewPatchedWritablePowerPanelRequest instantiates a new PatchedWritablePowerPanelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritablePowerPanelRequestWithDefaults

`func NewPatchedWritablePowerPanelRequestWithDefaults() *PatchedWritablePowerPanelRequest`

NewPatchedWritablePowerPanelRequestWithDefaults instantiates a new PatchedWritablePowerPanelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritablePowerPanelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritablePowerPanelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritablePowerPanelRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritablePowerPanelRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritablePowerPanelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritablePowerPanelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritablePowerPanelRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritablePowerPanelRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPanelType

`func (o *PatchedWritablePowerPanelRequest) GetPanelType() PatchedWritablePowerPanelRequestPanelType`

GetPanelType returns the PanelType field if non-nil, zero value otherwise.

### GetPanelTypeOk

`func (o *PatchedWritablePowerPanelRequest) GetPanelTypeOk() (*PatchedWritablePowerPanelRequestPanelType, bool)`

GetPanelTypeOk returns a tuple with the PanelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelType

`func (o *PatchedWritablePowerPanelRequest) SetPanelType(v PatchedWritablePowerPanelRequestPanelType)`

SetPanelType sets PanelType field to given value.

### HasPanelType

`func (o *PatchedWritablePowerPanelRequest) HasPanelType() bool`

HasPanelType returns a boolean if a field has been set.

### GetBreakerPositionCount

`func (o *PatchedWritablePowerPanelRequest) GetBreakerPositionCount() int32`

GetBreakerPositionCount returns the BreakerPositionCount field if non-nil, zero value otherwise.

### GetBreakerPositionCountOk

`func (o *PatchedWritablePowerPanelRequest) GetBreakerPositionCountOk() (*int32, bool)`

GetBreakerPositionCountOk returns a tuple with the BreakerPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakerPositionCount

`func (o *PatchedWritablePowerPanelRequest) SetBreakerPositionCount(v int32)`

SetBreakerPositionCount sets BreakerPositionCount field to given value.

### HasBreakerPositionCount

`func (o *PatchedWritablePowerPanelRequest) HasBreakerPositionCount() bool`

HasBreakerPositionCount returns a boolean if a field has been set.

### SetBreakerPositionCountNil

`func (o *PatchedWritablePowerPanelRequest) SetBreakerPositionCountNil(b bool)`

 SetBreakerPositionCountNil sets the value for BreakerPositionCount to be an explicit nil

### UnsetBreakerPositionCount
`func (o *PatchedWritablePowerPanelRequest) UnsetBreakerPositionCount()`

UnsetBreakerPositionCount ensures that no value is present for BreakerPositionCount, not even an explicit nil
### GetPowerPath

`func (o *PatchedWritablePowerPanelRequest) GetPowerPath() PatchedWritablePowerFeedRequestPowerPath`

GetPowerPath returns the PowerPath field if non-nil, zero value otherwise.

### GetPowerPathOk

`func (o *PatchedWritablePowerPanelRequest) GetPowerPathOk() (*PatchedWritablePowerFeedRequestPowerPath, bool)`

GetPowerPathOk returns a tuple with the PowerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPath

`func (o *PatchedWritablePowerPanelRequest) SetPowerPath(v PatchedWritablePowerFeedRequestPowerPath)`

SetPowerPath sets PowerPath field to given value.

### HasPowerPath

`func (o *PatchedWritablePowerPanelRequest) HasPowerPath() bool`

HasPowerPath returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedWritablePowerPanelRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedWritablePowerPanelRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedWritablePowerPanelRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedWritablePowerPanelRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRackGroup

`func (o *PatchedWritablePowerPanelRequest) GetRackGroup() BulkWritableCircuitRequestTenant`

GetRackGroup returns the RackGroup field if non-nil, zero value otherwise.

### GetRackGroupOk

`func (o *PatchedWritablePowerPanelRequest) GetRackGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRackGroupOk returns a tuple with the RackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackGroup

`func (o *PatchedWritablePowerPanelRequest) SetRackGroup(v BulkWritableCircuitRequestTenant)`

SetRackGroup sets RackGroup field to given value.

### HasRackGroup

`func (o *PatchedWritablePowerPanelRequest) HasRackGroup() bool`

HasRackGroup returns a boolean if a field has been set.

### SetRackGroupNil

`func (o *PatchedWritablePowerPanelRequest) SetRackGroupNil(b bool)`

 SetRackGroupNil sets the value for RackGroup to be an explicit nil

### UnsetRackGroup
`func (o *PatchedWritablePowerPanelRequest) UnsetRackGroup()`

UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
### GetCustomFields

`func (o *PatchedWritablePowerPanelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritablePowerPanelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritablePowerPanelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritablePowerPanelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritablePowerPanelRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritablePowerPanelRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritablePowerPanelRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritablePowerPanelRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritablePowerPanelRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritablePowerPanelRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritablePowerPanelRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritablePowerPanelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


