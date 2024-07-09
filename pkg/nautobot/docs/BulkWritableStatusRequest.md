# BulkWritableStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Color** | Pointer to **string** | RGB color in hexadecimal (e.g. 00ff00) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableStatusRequest

`func NewBulkWritableStatusRequest(id string, contentTypes []string, name string, ) *BulkWritableStatusRequest`

NewBulkWritableStatusRequest instantiates a new BulkWritableStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableStatusRequestWithDefaults

`func NewBulkWritableStatusRequestWithDefaults() *BulkWritableStatusRequest`

NewBulkWritableStatusRequestWithDefaults instantiates a new BulkWritableStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableStatusRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableStatusRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableStatusRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentTypes

`func (o *BulkWritableStatusRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *BulkWritableStatusRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *BulkWritableStatusRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.


### GetName

`func (o *BulkWritableStatusRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableStatusRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableStatusRequest) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *BulkWritableStatusRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BulkWritableStatusRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BulkWritableStatusRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *BulkWritableStatusRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *BulkWritableStatusRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableStatusRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableStatusRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableStatusRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableStatusRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableStatusRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableStatusRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableStatusRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableStatusRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableStatusRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableStatusRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableStatusRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


