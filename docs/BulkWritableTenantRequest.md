# BulkWritableTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**TenantGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableTenantRequest

`func NewBulkWritableTenantRequest(id string, name string, ) *BulkWritableTenantRequest`

NewBulkWritableTenantRequest instantiates a new BulkWritableTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableTenantRequestWithDefaults

`func NewBulkWritableTenantRequestWithDefaults() *BulkWritableTenantRequest`

NewBulkWritableTenantRequestWithDefaults instantiates a new BulkWritableTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableTenantRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableTenantRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableTenantRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableTenantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableTenantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableTenantRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableTenantRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableTenantRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableTenantRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableTenantRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *BulkWritableTenantRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWritableTenantRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWritableTenantRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWritableTenantRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTenantGroup

`func (o *BulkWritableTenantRequest) GetTenantGroup() BulkWritableCircuitRequestTenant`

GetTenantGroup returns the TenantGroup field if non-nil, zero value otherwise.

### GetTenantGroupOk

`func (o *BulkWritableTenantRequest) GetTenantGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantGroupOk returns a tuple with the TenantGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroup

`func (o *BulkWritableTenantRequest) SetTenantGroup(v BulkWritableCircuitRequestTenant)`

SetTenantGroup sets TenantGroup field to given value.

### HasTenantGroup

`func (o *BulkWritableTenantRequest) HasTenantGroup() bool`

HasTenantGroup returns a boolean if a field has been set.

### SetTenantGroupNil

`func (o *BulkWritableTenantRequest) SetTenantGroupNil(b bool)`

 SetTenantGroupNil sets the value for TenantGroup to be an explicit nil

### UnsetTenantGroup
`func (o *BulkWritableTenantRequest) UnsetTenantGroup()`

UnsetTenantGroup ensures that no value is present for TenantGroup, not even an explicit nil
### GetTags

`func (o *BulkWritableTenantRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableTenantRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableTenantRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableTenantRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableTenantRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableTenantRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableTenantRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableTenantRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableTenantRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableTenantRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableTenantRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableTenantRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


