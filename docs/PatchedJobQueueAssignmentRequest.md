# PatchedJobQueueAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Job** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**JobQueue** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewPatchedJobQueueAssignmentRequest

`func NewPatchedJobQueueAssignmentRequest() *PatchedJobQueueAssignmentRequest`

NewPatchedJobQueueAssignmentRequest instantiates a new PatchedJobQueueAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedJobQueueAssignmentRequestWithDefaults

`func NewPatchedJobQueueAssignmentRequestWithDefaults() *PatchedJobQueueAssignmentRequest`

NewPatchedJobQueueAssignmentRequestWithDefaults instantiates a new PatchedJobQueueAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedJobQueueAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedJobQueueAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedJobQueueAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedJobQueueAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJob

`func (o *PatchedJobQueueAssignmentRequest) GetJob() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *PatchedJobQueueAssignmentRequest) GetJobOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *PatchedJobQueueAssignmentRequest) SetJob(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetJob sets Job field to given value.

### HasJob

`func (o *PatchedJobQueueAssignmentRequest) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobQueue

`func (o *PatchedJobQueueAssignmentRequest) GetJobQueue() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetJobQueue returns the JobQueue field if non-nil, zero value otherwise.

### GetJobQueueOk

`func (o *PatchedJobQueueAssignmentRequest) GetJobQueueOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetJobQueueOk returns a tuple with the JobQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQueue

`func (o *PatchedJobQueueAssignmentRequest) SetJobQueue(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetJobQueue sets JobQueue field to given value.

### HasJobQueue

`func (o *PatchedJobQueueAssignmentRequest) HasJobQueue() bool`

HasJobQueue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


