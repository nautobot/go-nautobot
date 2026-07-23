# PaginatedVPNProfilePhase1PolicyAssignmentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]VPNProfilePhase1PolicyAssignment**](VPNProfilePhase1PolicyAssignment.md) |  | 

## Methods

### NewPaginatedVPNProfilePhase1PolicyAssignmentList

`func NewPaginatedVPNProfilePhase1PolicyAssignmentList(count int32, results []VPNProfilePhase1PolicyAssignment, ) *PaginatedVPNProfilePhase1PolicyAssignmentList`

NewPaginatedVPNProfilePhase1PolicyAssignmentList instantiates a new PaginatedVPNProfilePhase1PolicyAssignmentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedVPNProfilePhase1PolicyAssignmentListWithDefaults

`func NewPaginatedVPNProfilePhase1PolicyAssignmentListWithDefaults() *PaginatedVPNProfilePhase1PolicyAssignmentList`

NewPaginatedVPNProfilePhase1PolicyAssignmentListWithDefaults instantiates a new PaginatedVPNProfilePhase1PolicyAssignmentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) GetResults() []VPNProfilePhase1PolicyAssignment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) GetResultsOk() (*[]VPNProfilePhase1PolicyAssignment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedVPNProfilePhase1PolicyAssignmentList) SetResults(v []VPNProfilePhase1PolicyAssignment)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


