# PatchedBulkWritableCustomFieldChoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Value** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**CustomField** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableCustomFieldChoiceRequest

`func NewPatchedBulkWritableCustomFieldChoiceRequest(id string, ) *PatchedBulkWritableCustomFieldChoiceRequest`

NewPatchedBulkWritableCustomFieldChoiceRequest instantiates a new PatchedBulkWritableCustomFieldChoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableCustomFieldChoiceRequestWithDefaults

`func NewPatchedBulkWritableCustomFieldChoiceRequestWithDefaults() *PatchedBulkWritableCustomFieldChoiceRequest`

NewPatchedBulkWritableCustomFieldChoiceRequestWithDefaults instantiates a new PatchedBulkWritableCustomFieldChoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCustomField

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) GetCustomField() BulkWritableCableRequestStatus`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) GetCustomFieldOk() (*BulkWritableCableRequestStatus, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) SetCustomField(v BulkWritableCableRequestStatus)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *PatchedBulkWritableCustomFieldChoiceRequest) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


