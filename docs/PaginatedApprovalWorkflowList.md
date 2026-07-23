# PaginatedApprovalWorkflowList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]ApprovalWorkflow**](ApprovalWorkflow.md) |  | 

## Methods

### NewPaginatedApprovalWorkflowList

`func NewPaginatedApprovalWorkflowList(count int32, results []ApprovalWorkflow, ) *PaginatedApprovalWorkflowList`

NewPaginatedApprovalWorkflowList instantiates a new PaginatedApprovalWorkflowList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedApprovalWorkflowListWithDefaults

`func NewPaginatedApprovalWorkflowListWithDefaults() *PaginatedApprovalWorkflowList`

NewPaginatedApprovalWorkflowListWithDefaults instantiates a new PaginatedApprovalWorkflowList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedApprovalWorkflowList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedApprovalWorkflowList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedApprovalWorkflowList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedApprovalWorkflowList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedApprovalWorkflowList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedApprovalWorkflowList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedApprovalWorkflowList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedApprovalWorkflowList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedApprovalWorkflowList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedApprovalWorkflowList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedApprovalWorkflowList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedApprovalWorkflowList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedApprovalWorkflowList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedApprovalWorkflowList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedApprovalWorkflowList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedApprovalWorkflowList) GetResults() []ApprovalWorkflow`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedApprovalWorkflowList) GetResultsOk() (*[]ApprovalWorkflow, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedApprovalWorkflowList) SetResults(v []ApprovalWorkflow)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


