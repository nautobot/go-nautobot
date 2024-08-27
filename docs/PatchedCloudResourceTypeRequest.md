# PatchedCloudResourceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | Type of cloud objects | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ConfigSchema** | Pointer to **interface{}** |  | [optional] 
**Provider** | Pointer to [**BulkWritableCloudAccountRequestProvider**](BulkWritableCloudAccountRequestProvider.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedCloudResourceTypeRequest

`func NewPatchedCloudResourceTypeRequest() *PatchedCloudResourceTypeRequest`

NewPatchedCloudResourceTypeRequest instantiates a new PatchedCloudResourceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCloudResourceTypeRequestWithDefaults

`func NewPatchedCloudResourceTypeRequestWithDefaults() *PatchedCloudResourceTypeRequest`

NewPatchedCloudResourceTypeRequestWithDefaults instantiates a new PatchedCloudResourceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTypes

`func (o *PatchedCloudResourceTypeRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *PatchedCloudResourceTypeRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *PatchedCloudResourceTypeRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *PatchedCloudResourceTypeRequest) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedCloudResourceTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCloudResourceTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCloudResourceTypeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCloudResourceTypeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedCloudResourceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedCloudResourceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedCloudResourceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedCloudResourceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigSchema

`func (o *PatchedCloudResourceTypeRequest) GetConfigSchema() interface{}`

GetConfigSchema returns the ConfigSchema field if non-nil, zero value otherwise.

### GetConfigSchemaOk

`func (o *PatchedCloudResourceTypeRequest) GetConfigSchemaOk() (*interface{}, bool)`

GetConfigSchemaOk returns a tuple with the ConfigSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSchema

`func (o *PatchedCloudResourceTypeRequest) SetConfigSchema(v interface{})`

SetConfigSchema sets ConfigSchema field to given value.

### HasConfigSchema

`func (o *PatchedCloudResourceTypeRequest) HasConfigSchema() bool`

HasConfigSchema returns a boolean if a field has been set.

### SetConfigSchemaNil

`func (o *PatchedCloudResourceTypeRequest) SetConfigSchemaNil(b bool)`

 SetConfigSchemaNil sets the value for ConfigSchema to be an explicit nil

### UnsetConfigSchema
`func (o *PatchedCloudResourceTypeRequest) UnsetConfigSchema()`

UnsetConfigSchema ensures that no value is present for ConfigSchema, not even an explicit nil
### GetProvider

`func (o *PatchedCloudResourceTypeRequest) GetProvider() BulkWritableCloudAccountRequestProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PatchedCloudResourceTypeRequest) GetProviderOk() (*BulkWritableCloudAccountRequestProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PatchedCloudResourceTypeRequest) SetProvider(v BulkWritableCloudAccountRequestProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PatchedCloudResourceTypeRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedCloudResourceTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedCloudResourceTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedCloudResourceTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedCloudResourceTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedCloudResourceTypeRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedCloudResourceTypeRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedCloudResourceTypeRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedCloudResourceTypeRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedCloudResourceTypeRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedCloudResourceTypeRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedCloudResourceTypeRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedCloudResourceTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


