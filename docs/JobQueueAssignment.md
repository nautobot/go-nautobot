# JobQueueAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Job** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**JobQueue** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewJobQueueAssignment

`func NewJobQueueAssignment(objectType string, display string, url string, naturalSlug string, job ApprovalWorkflowStageResponseApprovalWorkflowStage, jobQueue ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *JobQueueAssignment`

NewJobQueueAssignment instantiates a new JobQueueAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobQueueAssignmentWithDefaults

`func NewJobQueueAssignmentWithDefaults() *JobQueueAssignment`

NewJobQueueAssignmentWithDefaults instantiates a new JobQueueAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobQueueAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobQueueAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobQueueAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobQueueAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *JobQueueAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *JobQueueAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *JobQueueAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *JobQueueAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *JobQueueAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *JobQueueAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *JobQueueAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JobQueueAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JobQueueAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *JobQueueAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *JobQueueAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *JobQueueAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetJob

`func (o *JobQueueAssignment) GetJob() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *JobQueueAssignment) GetJobOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *JobQueueAssignment) SetJob(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetJob sets Job field to given value.


### GetJobQueue

`func (o *JobQueueAssignment) GetJobQueue() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetJobQueue returns the JobQueue field if non-nil, zero value otherwise.

### GetJobQueueOk

`func (o *JobQueueAssignment) GetJobQueueOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetJobQueueOk returns a tuple with the JobQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQueue

`func (o *JobQueueAssignment) SetJobQueue(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetJobQueue sets JobQueue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


