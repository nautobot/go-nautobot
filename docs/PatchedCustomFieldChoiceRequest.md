# PatchedCustomFieldChoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**CustomField** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedCustomFieldChoiceRequest

`func NewPatchedCustomFieldChoiceRequest() *PatchedCustomFieldChoiceRequest`

NewPatchedCustomFieldChoiceRequest instantiates a new PatchedCustomFieldChoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCustomFieldChoiceRequestWithDefaults

`func NewPatchedCustomFieldChoiceRequestWithDefaults() *PatchedCustomFieldChoiceRequest`

NewPatchedCustomFieldChoiceRequestWithDefaults instantiates a new PatchedCustomFieldChoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PatchedCustomFieldChoiceRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchedCustomFieldChoiceRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchedCustomFieldChoiceRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchedCustomFieldChoiceRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedCustomFieldChoiceRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedCustomFieldChoiceRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedCustomFieldChoiceRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedCustomFieldChoiceRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCustomField

`func (o *PatchedCustomFieldChoiceRequest) GetCustomField() BulkWritableCableRequestStatus`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *PatchedCustomFieldChoiceRequest) GetCustomFieldOk() (*BulkWritableCableRequestStatus, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *PatchedCustomFieldChoiceRequest) SetCustomField(v BulkWritableCableRequestStatus)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *PatchedCustomFieldChoiceRequest) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


