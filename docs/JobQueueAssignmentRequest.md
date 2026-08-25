# JobQueueAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Job** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**JobQueue** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewJobQueueAssignmentRequest

`func NewJobQueueAssignmentRequest(job BulkWritableCableRequestStatus, jobQueue BulkWritableCableRequestStatus, ) *JobQueueAssignmentRequest`

NewJobQueueAssignmentRequest instantiates a new JobQueueAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobQueueAssignmentRequestWithDefaults

`func NewJobQueueAssignmentRequestWithDefaults() *JobQueueAssignmentRequest`

NewJobQueueAssignmentRequestWithDefaults instantiates a new JobQueueAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobQueueAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobQueueAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobQueueAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobQueueAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJob

`func (o *JobQueueAssignmentRequest) GetJob() BulkWritableCableRequestStatus`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *JobQueueAssignmentRequest) GetJobOk() (*BulkWritableCableRequestStatus, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *JobQueueAssignmentRequest) SetJob(v BulkWritableCableRequestStatus)`

SetJob sets Job field to given value.


### GetJobQueue

`func (o *JobQueueAssignmentRequest) GetJobQueue() BulkWritableCableRequestStatus`

GetJobQueue returns the JobQueue field if non-nil, zero value otherwise.

### GetJobQueueOk

`func (o *JobQueueAssignmentRequest) GetJobQueueOk() (*BulkWritableCableRequestStatus, bool)`

GetJobQueueOk returns a tuple with the JobQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQueue

`func (o *JobQueueAssignmentRequest) SetJobQueue(v BulkWritableCableRequestStatus)`

SetJobQueue sets JobQueue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


