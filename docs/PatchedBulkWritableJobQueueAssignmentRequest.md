# PatchedBulkWritableJobQueueAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Job** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**JobQueue** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableJobQueueAssignmentRequest

`func NewPatchedBulkWritableJobQueueAssignmentRequest(id string, ) *PatchedBulkWritableJobQueueAssignmentRequest`

NewPatchedBulkWritableJobQueueAssignmentRequest instantiates a new PatchedBulkWritableJobQueueAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableJobQueueAssignmentRequestWithDefaults

`func NewPatchedBulkWritableJobQueueAssignmentRequestWithDefaults() *PatchedBulkWritableJobQueueAssignmentRequest`

NewPatchedBulkWritableJobQueueAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableJobQueueAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetJob

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) GetJob() BulkWritableCableRequestStatus`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) GetJobOk() (*BulkWritableCableRequestStatus, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) SetJob(v BulkWritableCableRequestStatus)`

SetJob sets Job field to given value.

### HasJob

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobQueue

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) GetJobQueue() BulkWritableCableRequestStatus`

GetJobQueue returns the JobQueue field if non-nil, zero value otherwise.

### GetJobQueueOk

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) GetJobQueueOk() (*BulkWritableCableRequestStatus, bool)`

GetJobQueueOk returns a tuple with the JobQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQueue

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) SetJobQueue(v BulkWritableCableRequestStatus)`

SetJobQueue sets JobQueue field to given value.

### HasJobQueue

`func (o *PatchedBulkWritableJobQueueAssignmentRequest) HasJobQueue() bool`

HasJobQueue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


