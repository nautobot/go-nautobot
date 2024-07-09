# PaginatedVRFPrefixAssignmentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]VRFPrefixAssignment**](VRFPrefixAssignment.md) |  | 

## Methods

### NewPaginatedVRFPrefixAssignmentList

`func NewPaginatedVRFPrefixAssignmentList(count int32, results []VRFPrefixAssignment, ) *PaginatedVRFPrefixAssignmentList`

NewPaginatedVRFPrefixAssignmentList instantiates a new PaginatedVRFPrefixAssignmentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedVRFPrefixAssignmentListWithDefaults

`func NewPaginatedVRFPrefixAssignmentListWithDefaults() *PaginatedVRFPrefixAssignmentList`

NewPaginatedVRFPrefixAssignmentListWithDefaults instantiates a new PaginatedVRFPrefixAssignmentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedVRFPrefixAssignmentList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedVRFPrefixAssignmentList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedVRFPrefixAssignmentList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedVRFPrefixAssignmentList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedVRFPrefixAssignmentList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedVRFPrefixAssignmentList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedVRFPrefixAssignmentList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedVRFPrefixAssignmentList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedVRFPrefixAssignmentList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedVRFPrefixAssignmentList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedVRFPrefixAssignmentList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedVRFPrefixAssignmentList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedVRFPrefixAssignmentList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedVRFPrefixAssignmentList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedVRFPrefixAssignmentList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedVRFPrefixAssignmentList) GetResults() []VRFPrefixAssignment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedVRFPrefixAssignmentList) GetResultsOk() (*[]VRFPrefixAssignment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedVRFPrefixAssignmentList) SetResults(v []VRFPrefixAssignment)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


