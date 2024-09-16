# PrefixType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**PrefixTypeValue**](PrefixTypeValue.md) |  | [optional] [default to PREFIXTYPEVALUE_NETWORK]
**Label** | Pointer to [**PrefixTypeLabel**](PrefixTypeLabel.md) |  | [optional] [default to PREFIXTYPELABEL_NETWORK]

## Methods

### NewPrefixType

`func NewPrefixType() *PrefixType`

NewPrefixType instantiates a new PrefixType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixTypeWithDefaults

`func NewPrefixTypeWithDefaults() *PrefixType`

NewPrefixTypeWithDefaults instantiates a new PrefixType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PrefixType) GetValue() PrefixTypeValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PrefixType) GetValueOk() (*PrefixTypeValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PrefixType) SetValue(v PrefixTypeValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *PrefixType) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *PrefixType) GetLabel() PrefixTypeLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PrefixType) GetLabelOk() (*PrefixTypeLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PrefixType) SetLabel(v PrefixTypeLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PrefixType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


