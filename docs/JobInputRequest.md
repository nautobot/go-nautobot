# JobInputRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Schedule** | Pointer to [**JobCreationRequest**](JobCreationRequest.md) |  | [optional] 
**TaskQueue** | Pointer to **string** |  | [optional] 

## Methods

### NewJobInputRequest

`func NewJobInputRequest() *JobInputRequest`

NewJobInputRequest instantiates a new JobInputRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobInputRequestWithDefaults

`func NewJobInputRequestWithDefaults() *JobInputRequest`

NewJobInputRequestWithDefaults instantiates a new JobInputRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JobInputRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobInputRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobInputRequest) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *JobInputRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JobInputRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JobInputRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetSchedule

`func (o *JobInputRequest) GetSchedule() JobCreationRequest`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *JobInputRequest) GetScheduleOk() (*JobCreationRequest, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *JobInputRequest) SetSchedule(v JobCreationRequest)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *JobInputRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTaskQueue

`func (o *JobInputRequest) GetTaskQueue() string`

GetTaskQueue returns the TaskQueue field if non-nil, zero value otherwise.

### GetTaskQueueOk

`func (o *JobInputRequest) GetTaskQueueOk() (*string, bool)`

GetTaskQueueOk returns a tuple with the TaskQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueue

`func (o *JobInputRequest) SetTaskQueue(v string)`

SetTaskQueue sets TaskQueue field to given value.

### HasTaskQueue

`func (o *JobInputRequest) HasTaskQueue() bool`

HasTaskQueue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


