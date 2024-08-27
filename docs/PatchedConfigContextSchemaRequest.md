# PatchedConfigContextSchemaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerContentType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DataSchema** | Pointer to **interface{}** | A JSON Schema document which is used to validate a config context object. | [optional] 
**OwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedConfigContextSchemaRequest

`func NewPatchedConfigContextSchemaRequest() *PatchedConfigContextSchemaRequest`

NewPatchedConfigContextSchemaRequest instantiates a new PatchedConfigContextSchemaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedConfigContextSchemaRequestWithDefaults

`func NewPatchedConfigContextSchemaRequestWithDefaults() *PatchedConfigContextSchemaRequest`

NewPatchedConfigContextSchemaRequestWithDefaults instantiates a new PatchedConfigContextSchemaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerContentType

`func (o *PatchedConfigContextSchemaRequest) GetOwnerContentType() string`

GetOwnerContentType returns the OwnerContentType field if non-nil, zero value otherwise.

### GetOwnerContentTypeOk

`func (o *PatchedConfigContextSchemaRequest) GetOwnerContentTypeOk() (*string, bool)`

GetOwnerContentTypeOk returns a tuple with the OwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContentType

`func (o *PatchedConfigContextSchemaRequest) SetOwnerContentType(v string)`

SetOwnerContentType sets OwnerContentType field to given value.

### HasOwnerContentType

`func (o *PatchedConfigContextSchemaRequest) HasOwnerContentType() bool`

HasOwnerContentType returns a boolean if a field has been set.

### SetOwnerContentTypeNil

`func (o *PatchedConfigContextSchemaRequest) SetOwnerContentTypeNil(b bool)`

 SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil

### UnsetOwnerContentType
`func (o *PatchedConfigContextSchemaRequest) UnsetOwnerContentType()`

UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
### GetName

`func (o *PatchedConfigContextSchemaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedConfigContextSchemaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedConfigContextSchemaRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedConfigContextSchemaRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedConfigContextSchemaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedConfigContextSchemaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedConfigContextSchemaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedConfigContextSchemaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDataSchema

`func (o *PatchedConfigContextSchemaRequest) GetDataSchema() interface{}`

GetDataSchema returns the DataSchema field if non-nil, zero value otherwise.

### GetDataSchemaOk

`func (o *PatchedConfigContextSchemaRequest) GetDataSchemaOk() (*interface{}, bool)`

GetDataSchemaOk returns a tuple with the DataSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSchema

`func (o *PatchedConfigContextSchemaRequest) SetDataSchema(v interface{})`

SetDataSchema sets DataSchema field to given value.

### HasDataSchema

`func (o *PatchedConfigContextSchemaRequest) HasDataSchema() bool`

HasDataSchema returns a boolean if a field has been set.

### SetDataSchemaNil

`func (o *PatchedConfigContextSchemaRequest) SetDataSchemaNil(b bool)`

 SetDataSchemaNil sets the value for DataSchema to be an explicit nil

### UnsetDataSchema
`func (o *PatchedConfigContextSchemaRequest) UnsetDataSchema()`

UnsetDataSchema ensures that no value is present for DataSchema, not even an explicit nil
### GetOwnerObjectId

`func (o *PatchedConfigContextSchemaRequest) GetOwnerObjectId() string`

GetOwnerObjectId returns the OwnerObjectId field if non-nil, zero value otherwise.

### GetOwnerObjectIdOk

`func (o *PatchedConfigContextSchemaRequest) GetOwnerObjectIdOk() (*string, bool)`

GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerObjectId

`func (o *PatchedConfigContextSchemaRequest) SetOwnerObjectId(v string)`

SetOwnerObjectId sets OwnerObjectId field to given value.

### HasOwnerObjectId

`func (o *PatchedConfigContextSchemaRequest) HasOwnerObjectId() bool`

HasOwnerObjectId returns a boolean if a field has been set.

### SetOwnerObjectIdNil

`func (o *PatchedConfigContextSchemaRequest) SetOwnerObjectIdNil(b bool)`

 SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil

### UnsetOwnerObjectId
`func (o *PatchedConfigContextSchemaRequest) UnsetOwnerObjectId()`

UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
### GetCustomFields

`func (o *PatchedConfigContextSchemaRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedConfigContextSchemaRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedConfigContextSchemaRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedConfigContextSchemaRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedConfigContextSchemaRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedConfigContextSchemaRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedConfigContextSchemaRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedConfigContextSchemaRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


