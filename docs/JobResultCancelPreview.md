# JobResultCancelPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Confirmation prompt to display to the user. | 
**JobStatus** | [**JobStatusEnum**](JobStatusEnum.md) | For unready jobs: RUNNING, NOT RUNNING, or UNKNOWN. For ready jobs: the terminal state. | 
**Irreversible** | Pointer to **string** | Warning that the action cannot be undone. Omitted when the job is already finished. | [optional] 
**Timestamp** | **time.Time** | Server time when this preview was generated. | 

## Methods

### NewJobResultCancelPreview

`func NewJobResultCancelPreview(message string, jobStatus JobStatusEnum, timestamp time.Time, ) *JobResultCancelPreview`

NewJobResultCancelPreview instantiates a new JobResultCancelPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResultCancelPreviewWithDefaults

`func NewJobResultCancelPreviewWithDefaults() *JobResultCancelPreview`

NewJobResultCancelPreviewWithDefaults instantiates a new JobResultCancelPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *JobResultCancelPreview) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JobResultCancelPreview) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JobResultCancelPreview) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetJobStatus

`func (o *JobResultCancelPreview) GetJobStatus() JobStatusEnum`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *JobResultCancelPreview) GetJobStatusOk() (*JobStatusEnum, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *JobResultCancelPreview) SetJobStatus(v JobStatusEnum)`

SetJobStatus sets JobStatus field to given value.


### GetIrreversible

`func (o *JobResultCancelPreview) GetIrreversible() string`

GetIrreversible returns the Irreversible field if non-nil, zero value otherwise.

### GetIrreversibleOk

`func (o *JobResultCancelPreview) GetIrreversibleOk() (*string, bool)`

GetIrreversibleOk returns a tuple with the Irreversible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrreversible

`func (o *JobResultCancelPreview) SetIrreversible(v string)`

SetIrreversible sets Irreversible field to given value.

### HasIrreversible

`func (o *JobResultCancelPreview) HasIrreversible() bool`

HasIrreversible returns a boolean if a field has been set.

### GetTimestamp

`func (o *JobResultCancelPreview) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JobResultCancelPreview) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JobResultCancelPreview) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


