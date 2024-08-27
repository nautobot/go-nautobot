# PatchedBulkWritableDeviceRedundancyGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FailoverStrategy** | Pointer to [**BulkWritableDeviceRedundancyGroupRequestFailoverStrategy**](BulkWritableDeviceRedundancyGroupRequestFailoverStrategy.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**SecretsGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableDeviceRedundancyGroupRequest

`func NewPatchedBulkWritableDeviceRedundancyGroupRequest(id string, ) *PatchedBulkWritableDeviceRedundancyGroupRequest`

NewPatchedBulkWritableDeviceRedundancyGroupRequest instantiates a new PatchedBulkWritableDeviceRedundancyGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableDeviceRedundancyGroupRequestWithDefaults

`func NewPatchedBulkWritableDeviceRedundancyGroupRequestWithDefaults() *PatchedBulkWritableDeviceRedundancyGroupRequest`

NewPatchedBulkWritableDeviceRedundancyGroupRequestWithDefaults instantiates a new PatchedBulkWritableDeviceRedundancyGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetId(v string)`

SetId sets Id field to given value.


### GetFailoverStrategy

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetFailoverStrategy() BulkWritableDeviceRedundancyGroupRequestFailoverStrategy`

GetFailoverStrategy returns the FailoverStrategy field if non-nil, zero value otherwise.

### GetFailoverStrategyOk

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetFailoverStrategyOk() (*BulkWritableDeviceRedundancyGroupRequestFailoverStrategy, bool)`

GetFailoverStrategyOk returns a tuple with the FailoverStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverStrategy

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetFailoverStrategy(v BulkWritableDeviceRedundancyGroupRequestFailoverStrategy)`

SetFailoverStrategy sets FailoverStrategy field to given value.

### HasFailoverStrategy

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) HasFailoverStrategy() bool`

HasFailoverStrategy returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetTags

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableDeviceRedundancyGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


