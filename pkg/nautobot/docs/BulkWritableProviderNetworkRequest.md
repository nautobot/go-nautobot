# BulkWritableProviderNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Provider** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableProviderNetworkRequest

`func NewBulkWritableProviderNetworkRequest(id string, name string, provider BulkWritableCableRequestStatus, ) *BulkWritableProviderNetworkRequest`

NewBulkWritableProviderNetworkRequest instantiates a new BulkWritableProviderNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableProviderNetworkRequestWithDefaults

`func NewBulkWritableProviderNetworkRequestWithDefaults() *BulkWritableProviderNetworkRequest`

NewBulkWritableProviderNetworkRequestWithDefaults instantiates a new BulkWritableProviderNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableProviderNetworkRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableProviderNetworkRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableProviderNetworkRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableProviderNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableProviderNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableProviderNetworkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableProviderNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableProviderNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableProviderNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableProviderNetworkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *BulkWritableProviderNetworkRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWritableProviderNetworkRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWritableProviderNetworkRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWritableProviderNetworkRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetProvider

`func (o *BulkWritableProviderNetworkRequest) GetProvider() BulkWritableCableRequestStatus`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BulkWritableProviderNetworkRequest) GetProviderOk() (*BulkWritableCableRequestStatus, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BulkWritableProviderNetworkRequest) SetProvider(v BulkWritableCableRequestStatus)`

SetProvider sets Provider field to given value.


### GetTags

`func (o *BulkWritableProviderNetworkRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableProviderNetworkRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableProviderNetworkRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableProviderNetworkRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableProviderNetworkRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableProviderNetworkRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableProviderNetworkRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableProviderNetworkRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableProviderNetworkRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableProviderNetworkRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableProviderNetworkRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableProviderNetworkRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


