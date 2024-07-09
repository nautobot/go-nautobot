# BulkWritableDeviceRedundancyGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FailoverStrategy** | Pointer to [**BulkWritableDeviceRedundancyGroupRequestFailoverStrategy**](BulkWritableDeviceRedundancyGroupRequestFailoverStrategy.md) |  | [optional] [default to null]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**SecretsGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableDeviceRedundancyGroupRequest

`func NewBulkWritableDeviceRedundancyGroupRequest(id string, name string, status BulkWritableCableRequestStatus, ) *BulkWritableDeviceRedundancyGroupRequest`

NewBulkWritableDeviceRedundancyGroupRequest instantiates a new BulkWritableDeviceRedundancyGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableDeviceRedundancyGroupRequestWithDefaults

`func NewBulkWritableDeviceRedundancyGroupRequestWithDefaults() *BulkWritableDeviceRedundancyGroupRequest`

NewBulkWritableDeviceRedundancyGroupRequestWithDefaults instantiates a new BulkWritableDeviceRedundancyGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetId(v string)`

SetId sets Id field to given value.


### GetFailoverStrategy

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetFailoverStrategy() BulkWritableDeviceRedundancyGroupRequestFailoverStrategy`

GetFailoverStrategy returns the FailoverStrategy field if non-nil, zero value otherwise.

### GetFailoverStrategyOk

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetFailoverStrategyOk() (*BulkWritableDeviceRedundancyGroupRequestFailoverStrategy, bool)`

GetFailoverStrategyOk returns a tuple with the FailoverStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverStrategy

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetFailoverStrategy(v BulkWritableDeviceRedundancyGroupRequestFailoverStrategy)`

SetFailoverStrategy sets FailoverStrategy field to given value.

### HasFailoverStrategy

`func (o *BulkWritableDeviceRedundancyGroupRequest) HasFailoverStrategy() bool`

HasFailoverStrategy returns a boolean if a field has been set.

### GetName

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableDeviceRedundancyGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWritableDeviceRedundancyGroupRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetStatus

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetSecretsGroup

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *BulkWritableDeviceRedundancyGroupRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *BulkWritableDeviceRedundancyGroupRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetTags

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableDeviceRedundancyGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableDeviceRedundancyGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableDeviceRedundancyGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableDeviceRedundancyGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableDeviceRedundancyGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


