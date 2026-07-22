# PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]ControllerManagedDeviceGroupRadioProfileAssignment**](ControllerManagedDeviceGroupRadioProfileAssignment.md) |  | 

## Methods

### NewPaginatedControllerManagedDeviceGroupRadioProfileAssignmentList

`func NewPaginatedControllerManagedDeviceGroupRadioProfileAssignmentList(count int32, results []ControllerManagedDeviceGroupRadioProfileAssignment, ) *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList`

NewPaginatedControllerManagedDeviceGroupRadioProfileAssignmentList instantiates a new PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedControllerManagedDeviceGroupRadioProfileAssignmentListWithDefaults

`func NewPaginatedControllerManagedDeviceGroupRadioProfileAssignmentListWithDefaults() *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList`

NewPaginatedControllerManagedDeviceGroupRadioProfileAssignmentListWithDefaults instantiates a new PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) GetResults() []ControllerManagedDeviceGroupRadioProfileAssignment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) GetResultsOk() (*[]ControllerManagedDeviceGroupRadioProfileAssignment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedControllerManagedDeviceGroupRadioProfileAssignmentList) SetResults(v []ControllerManagedDeviceGroupRadioProfileAssignment)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


