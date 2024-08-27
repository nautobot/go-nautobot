# BulkWritableCloudResourceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentTypes** | **[]string** |  | 
**Name** | **string** | Type of cloud objects | 
**Description** | Pointer to **string** |  | [optional] 
**ConfigSchema** | Pointer to **interface{}** |  | [optional] 
**Provider** | [**BulkWritableCloudAccountRequestProvider**](BulkWritableCloudAccountRequestProvider.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableCloudResourceTypeRequest

`func NewBulkWritableCloudResourceTypeRequest(id string, contentTypes []string, name string, provider BulkWritableCloudAccountRequestProvider, ) *BulkWritableCloudResourceTypeRequest`

NewBulkWritableCloudResourceTypeRequest instantiates a new BulkWritableCloudResourceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableCloudResourceTypeRequestWithDefaults

`func NewBulkWritableCloudResourceTypeRequestWithDefaults() *BulkWritableCloudResourceTypeRequest`

NewBulkWritableCloudResourceTypeRequestWithDefaults instantiates a new BulkWritableCloudResourceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableCloudResourceTypeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableCloudResourceTypeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableCloudResourceTypeRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentTypes

`func (o *BulkWritableCloudResourceTypeRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *BulkWritableCloudResourceTypeRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *BulkWritableCloudResourceTypeRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.


### GetName

`func (o *BulkWritableCloudResourceTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableCloudResourceTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableCloudResourceTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableCloudResourceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableCloudResourceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableCloudResourceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableCloudResourceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigSchema

`func (o *BulkWritableCloudResourceTypeRequest) GetConfigSchema() interface{}`

GetConfigSchema returns the ConfigSchema field if non-nil, zero value otherwise.

### GetConfigSchemaOk

`func (o *BulkWritableCloudResourceTypeRequest) GetConfigSchemaOk() (*interface{}, bool)`

GetConfigSchemaOk returns a tuple with the ConfigSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSchema

`func (o *BulkWritableCloudResourceTypeRequest) SetConfigSchema(v interface{})`

SetConfigSchema sets ConfigSchema field to given value.

### HasConfigSchema

`func (o *BulkWritableCloudResourceTypeRequest) HasConfigSchema() bool`

HasConfigSchema returns a boolean if a field has been set.

### SetConfigSchemaNil

`func (o *BulkWritableCloudResourceTypeRequest) SetConfigSchemaNil(b bool)`

 SetConfigSchemaNil sets the value for ConfigSchema to be an explicit nil

### UnsetConfigSchema
`func (o *BulkWritableCloudResourceTypeRequest) UnsetConfigSchema()`

UnsetConfigSchema ensures that no value is present for ConfigSchema, not even an explicit nil
### GetProvider

`func (o *BulkWritableCloudResourceTypeRequest) GetProvider() BulkWritableCloudAccountRequestProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BulkWritableCloudResourceTypeRequest) GetProviderOk() (*BulkWritableCloudAccountRequestProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BulkWritableCloudResourceTypeRequest) SetProvider(v BulkWritableCloudAccountRequestProvider)`

SetProvider sets Provider field to given value.


### GetCustomFields

`func (o *BulkWritableCloudResourceTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableCloudResourceTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableCloudResourceTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableCloudResourceTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableCloudResourceTypeRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableCloudResourceTypeRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableCloudResourceTypeRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableCloudResourceTypeRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableCloudResourceTypeRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableCloudResourceTypeRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableCloudResourceTypeRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableCloudResourceTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


