# PaginatedJobQueueAssignmentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]JobQueueAssignment**](JobQueueAssignment.md) |  | 

## Methods

### NewPaginatedJobQueueAssignmentList

`func NewPaginatedJobQueueAssignmentList(count int32, results []JobQueueAssignment, ) *PaginatedJobQueueAssignmentList`

NewPaginatedJobQueueAssignmentList instantiates a new PaginatedJobQueueAssignmentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedJobQueueAssignmentListWithDefaults

`func NewPaginatedJobQueueAssignmentListWithDefaults() *PaginatedJobQueueAssignmentList`

NewPaginatedJobQueueAssignmentListWithDefaults instantiates a new PaginatedJobQueueAssignmentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedJobQueueAssignmentList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedJobQueueAssignmentList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedJobQueueAssignmentList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedJobQueueAssignmentList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedJobQueueAssignmentList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedJobQueueAssignmentList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedJobQueueAssignmentList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedJobQueueAssignmentList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedJobQueueAssignmentList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedJobQueueAssignmentList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedJobQueueAssignmentList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedJobQueueAssignmentList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedJobQueueAssignmentList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedJobQueueAssignmentList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedJobQueueAssignmentList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedJobQueueAssignmentList) GetResults() []JobQueueAssignment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedJobQueueAssignmentList) GetResultsOk() (*[]JobQueueAssignment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedJobQueueAssignmentList) SetResults(v []JobQueueAssignment)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


