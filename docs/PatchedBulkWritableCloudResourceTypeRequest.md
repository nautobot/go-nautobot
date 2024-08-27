# PatchedBulkWritableCloudResourceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Type of cloud objects | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ConfigSchema** | Pointer to **interface{}** |  | [optional] 
**Provider** | Pointer to [**BulkWritableCloudAccountRequestProvider**](BulkWritableCloudAccountRequestProvider.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableCloudResourceTypeRequest

`func NewPatchedBulkWritableCloudResourceTypeRequest(id string, ) *PatchedBulkWritableCloudResourceTypeRequest`

NewPatchedBulkWritableCloudResourceTypeRequest instantiates a new PatchedBulkWritableCloudResourceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableCloudResourceTypeRequestWithDefaults

`func NewPatchedBulkWritableCloudResourceTypeRequestWithDefaults() *PatchedBulkWritableCloudResourceTypeRequest`

NewPatchedBulkWritableCloudResourceTypeRequestWithDefaults instantiates a new PatchedBulkWritableCloudResourceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableCloudResourceTypeRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentTypes

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *PatchedBulkWritableCloudResourceTypeRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *PatchedBulkWritableCloudResourceTypeRequest) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableCloudResourceTypeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableCloudResourceTypeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableCloudResourceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableCloudResourceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigSchema

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetConfigSchema() interface{}`

GetConfigSchema returns the ConfigSchema field if non-nil, zero value otherwise.

### GetConfigSchemaOk

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetConfigSchemaOk() (*interface{}, bool)`

GetConfigSchemaOk returns a tuple with the ConfigSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSchema

`func (o *PatchedBulkWritableCloudResourceTypeRequest) SetConfigSchema(v interface{})`

SetConfigSchema sets ConfigSchema field to given value.

### HasConfigSchema

`func (o *PatchedBulkWritableCloudResourceTypeRequest) HasConfigSchema() bool`

HasConfigSchema returns a boolean if a field has been set.

### SetConfigSchemaNil

`func (o *PatchedBulkWritableCloudResourceTypeRequest) SetConfigSchemaNil(b bool)`

 SetConfigSchemaNil sets the value for ConfigSchema to be an explicit nil

### UnsetConfigSchema
`func (o *PatchedBulkWritableCloudResourceTypeRequest) UnsetConfigSchema()`

UnsetConfigSchema ensures that no value is present for ConfigSchema, not even an explicit nil
### GetProvider

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetProvider() BulkWritableCloudAccountRequestProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetProviderOk() (*BulkWritableCloudAccountRequestProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PatchedBulkWritableCloudResourceTypeRequest) SetProvider(v BulkWritableCloudAccountRequestProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PatchedBulkWritableCloudResourceTypeRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableCloudResourceTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableCloudResourceTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableCloudResourceTypeRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableCloudResourceTypeRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableCloudResourceTypeRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableCloudResourceTypeRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableCloudResourceTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


