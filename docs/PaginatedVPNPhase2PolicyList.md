# PaginatedVPNPhase2PolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]VPNPhase2Policy**](VPNPhase2Policy.md) |  | 

## Methods

### NewPaginatedVPNPhase2PolicyList

`func NewPaginatedVPNPhase2PolicyList(count int32, results []VPNPhase2Policy, ) *PaginatedVPNPhase2PolicyList`

NewPaginatedVPNPhase2PolicyList instantiates a new PaginatedVPNPhase2PolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedVPNPhase2PolicyListWithDefaults

`func NewPaginatedVPNPhase2PolicyListWithDefaults() *PaginatedVPNPhase2PolicyList`

NewPaginatedVPNPhase2PolicyListWithDefaults instantiates a new PaginatedVPNPhase2PolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedVPNPhase2PolicyList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedVPNPhase2PolicyList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedVPNPhase2PolicyList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedVPNPhase2PolicyList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedVPNPhase2PolicyList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedVPNPhase2PolicyList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedVPNPhase2PolicyList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedVPNPhase2PolicyList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedVPNPhase2PolicyList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedVPNPhase2PolicyList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedVPNPhase2PolicyList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedVPNPhase2PolicyList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedVPNPhase2PolicyList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedVPNPhase2PolicyList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedVPNPhase2PolicyList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedVPNPhase2PolicyList) GetResults() []VPNPhase2Policy`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedVPNPhase2PolicyList) GetResultsOk() (*[]VPNPhase2Policy, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedVPNPhase2PolicyList) SetResults(v []VPNPhase2Policy)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


