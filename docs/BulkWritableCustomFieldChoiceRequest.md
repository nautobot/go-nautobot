# BulkWritableCustomFieldChoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Value** | **string** |  | 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**CustomField** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewBulkWritableCustomFieldChoiceRequest

`func NewBulkWritableCustomFieldChoiceRequest(id string, value string, customField BulkWritableCableRequestStatus, ) *BulkWritableCustomFieldChoiceRequest`

NewBulkWritableCustomFieldChoiceRequest instantiates a new BulkWritableCustomFieldChoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableCustomFieldChoiceRequestWithDefaults

`func NewBulkWritableCustomFieldChoiceRequestWithDefaults() *BulkWritableCustomFieldChoiceRequest`

NewBulkWritableCustomFieldChoiceRequestWithDefaults instantiates a new BulkWritableCustomFieldChoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableCustomFieldChoiceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableCustomFieldChoiceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableCustomFieldChoiceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *BulkWritableCustomFieldChoiceRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BulkWritableCustomFieldChoiceRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BulkWritableCustomFieldChoiceRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetWeight

`func (o *BulkWritableCustomFieldChoiceRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkWritableCustomFieldChoiceRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkWritableCustomFieldChoiceRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BulkWritableCustomFieldChoiceRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCustomField

`func (o *BulkWritableCustomFieldChoiceRequest) GetCustomField() BulkWritableCableRequestStatus`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *BulkWritableCustomFieldChoiceRequest) GetCustomFieldOk() (*BulkWritableCableRequestStatus, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *BulkWritableCustomFieldChoiceRequest) SetCustomField(v BulkWritableCableRequestStatus)`

SetCustomField sets CustomField field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


