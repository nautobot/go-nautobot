# BulkWritableConfigContextSchemaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OwnerContentType** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DataSchema** | **map[string]interface{}** | A JSON Schema document which is used to validate a config context object. | 
**OwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableConfigContextSchemaRequest

`func NewBulkWritableConfigContextSchemaRequest(id string, name string, dataSchema map[string]interface{}, ) *BulkWritableConfigContextSchemaRequest`

NewBulkWritableConfigContextSchemaRequest instantiates a new BulkWritableConfigContextSchemaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableConfigContextSchemaRequestWithDefaults

`func NewBulkWritableConfigContextSchemaRequestWithDefaults() *BulkWritableConfigContextSchemaRequest`

NewBulkWritableConfigContextSchemaRequestWithDefaults instantiates a new BulkWritableConfigContextSchemaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableConfigContextSchemaRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableConfigContextSchemaRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableConfigContextSchemaRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOwnerContentType

`func (o *BulkWritableConfigContextSchemaRequest) GetOwnerContentType() string`

GetOwnerContentType returns the OwnerContentType field if non-nil, zero value otherwise.

### GetOwnerContentTypeOk

`func (o *BulkWritableConfigContextSchemaRequest) GetOwnerContentTypeOk() (*string, bool)`

GetOwnerContentTypeOk returns a tuple with the OwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContentType

`func (o *BulkWritableConfigContextSchemaRequest) SetOwnerContentType(v string)`

SetOwnerContentType sets OwnerContentType field to given value.

### HasOwnerContentType

`func (o *BulkWritableConfigContextSchemaRequest) HasOwnerContentType() bool`

HasOwnerContentType returns a boolean if a field has been set.

### SetOwnerContentTypeNil

`func (o *BulkWritableConfigContextSchemaRequest) SetOwnerContentTypeNil(b bool)`

 SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil

### UnsetOwnerContentType
`func (o *BulkWritableConfigContextSchemaRequest) UnsetOwnerContentType()`

UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
### GetName

`func (o *BulkWritableConfigContextSchemaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableConfigContextSchemaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableConfigContextSchemaRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableConfigContextSchemaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableConfigContextSchemaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableConfigContextSchemaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableConfigContextSchemaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDataSchema

`func (o *BulkWritableConfigContextSchemaRequest) GetDataSchema() map[string]interface{}`

GetDataSchema returns the DataSchema field if non-nil, zero value otherwise.

### GetDataSchemaOk

`func (o *BulkWritableConfigContextSchemaRequest) GetDataSchemaOk() (*map[string]interface{}, bool)`

GetDataSchemaOk returns a tuple with the DataSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSchema

`func (o *BulkWritableConfigContextSchemaRequest) SetDataSchema(v map[string]interface{})`

SetDataSchema sets DataSchema field to given value.


### GetOwnerObjectId

`func (o *BulkWritableConfigContextSchemaRequest) GetOwnerObjectId() string`

GetOwnerObjectId returns the OwnerObjectId field if non-nil, zero value otherwise.

### GetOwnerObjectIdOk

`func (o *BulkWritableConfigContextSchemaRequest) GetOwnerObjectIdOk() (*string, bool)`

GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerObjectId

`func (o *BulkWritableConfigContextSchemaRequest) SetOwnerObjectId(v string)`

SetOwnerObjectId sets OwnerObjectId field to given value.

### HasOwnerObjectId

`func (o *BulkWritableConfigContextSchemaRequest) HasOwnerObjectId() bool`

HasOwnerObjectId returns a boolean if a field has been set.

### SetOwnerObjectIdNil

`func (o *BulkWritableConfigContextSchemaRequest) SetOwnerObjectIdNil(b bool)`

 SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil

### UnsetOwnerObjectId
`func (o *BulkWritableConfigContextSchemaRequest) UnsetOwnerObjectId()`

UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableConfigContextSchemaRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableConfigContextSchemaRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableConfigContextSchemaRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableConfigContextSchemaRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableConfigContextSchemaRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableConfigContextSchemaRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableConfigContextSchemaRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableConfigContextSchemaRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


