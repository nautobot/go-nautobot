# CustomFieldChoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**CustomField** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewCustomFieldChoiceRequest

`func NewCustomFieldChoiceRequest(value string, customField BulkWritableCableRequestStatus, ) *CustomFieldChoiceRequest`

NewCustomFieldChoiceRequest instantiates a new CustomFieldChoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldChoiceRequestWithDefaults

`func NewCustomFieldChoiceRequestWithDefaults() *CustomFieldChoiceRequest`

NewCustomFieldChoiceRequestWithDefaults instantiates a new CustomFieldChoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CustomFieldChoiceRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldChoiceRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldChoiceRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetWeight

`func (o *CustomFieldChoiceRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CustomFieldChoiceRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CustomFieldChoiceRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CustomFieldChoiceRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCustomField

`func (o *CustomFieldChoiceRequest) GetCustomField() BulkWritableCableRequestStatus`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *CustomFieldChoiceRequest) GetCustomFieldOk() (*BulkWritableCableRequestStatus, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *CustomFieldChoiceRequest) SetCustomField(v BulkWritableCableRequestStatus)`

SetCustomField sets CustomField field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


