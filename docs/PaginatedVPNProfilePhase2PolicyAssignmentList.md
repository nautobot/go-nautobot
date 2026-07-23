# PaginatedVPNProfilePhase2PolicyAssignmentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]VPNProfilePhase2PolicyAssignment**](VPNProfilePhase2PolicyAssignment.md) |  | 

## Methods

### NewPaginatedVPNProfilePhase2PolicyAssignmentList

`func NewPaginatedVPNProfilePhase2PolicyAssignmentList(count int32, results []VPNProfilePhase2PolicyAssignment, ) *PaginatedVPNProfilePhase2PolicyAssignmentList`

NewPaginatedVPNProfilePhase2PolicyAssignmentList instantiates a new PaginatedVPNProfilePhase2PolicyAssignmentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedVPNProfilePhase2PolicyAssignmentListWithDefaults

`func NewPaginatedVPNProfilePhase2PolicyAssignmentListWithDefaults() *PaginatedVPNProfilePhase2PolicyAssignmentList`

NewPaginatedVPNProfilePhase2PolicyAssignmentListWithDefaults instantiates a new PaginatedVPNProfilePhase2PolicyAssignmentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) GetResults() []VPNProfilePhase2PolicyAssignment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) GetResultsOk() (*[]VPNProfilePhase2PolicyAssignment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedVPNProfilePhase2PolicyAssignmentList) SetResults(v []VPNProfilePhase2PolicyAssignment)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


