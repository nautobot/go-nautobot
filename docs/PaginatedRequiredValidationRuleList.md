# PaginatedRequiredValidationRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]RequiredValidationRule**](RequiredValidationRule.md) |  | 

## Methods

### NewPaginatedRequiredValidationRuleList

`func NewPaginatedRequiredValidationRuleList(count int32, results []RequiredValidationRule, ) *PaginatedRequiredValidationRuleList`

NewPaginatedRequiredValidationRuleList instantiates a new PaginatedRequiredValidationRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedRequiredValidationRuleListWithDefaults

`func NewPaginatedRequiredValidationRuleListWithDefaults() *PaginatedRequiredValidationRuleList`

NewPaginatedRequiredValidationRuleListWithDefaults instantiates a new PaginatedRequiredValidationRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedRequiredValidationRuleList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedRequiredValidationRuleList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedRequiredValidationRuleList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedRequiredValidationRuleList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedRequiredValidationRuleList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedRequiredValidationRuleList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedRequiredValidationRuleList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedRequiredValidationRuleList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedRequiredValidationRuleList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedRequiredValidationRuleList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedRequiredValidationRuleList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedRequiredValidationRuleList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedRequiredValidationRuleList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedRequiredValidationRuleList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedRequiredValidationRuleList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedRequiredValidationRuleList) GetResults() []RequiredValidationRule`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedRequiredValidationRuleList) GetResultsOk() (*[]RequiredValidationRule, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedRequiredValidationRuleList) SetResults(v []RequiredValidationRule)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


