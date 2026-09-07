# JobResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**TaskName** | Pointer to **NullableString** | Registered name of the Celery task for this job. Internal use only. | [optional] 
**DateStarted** | Pointer to **NullableTime** |  | [optional] 
**DateDone** | Pointer to **NullableTime** |  | [optional] 
**Worker** | Pointer to **NullableString** |  | [optional] 
**TaskArgs** | Pointer to **interface{}** |  | [optional] 
**TaskKwargs** | Pointer to **interface{}** |  | [optional] 
**CeleryKwargs** | Pointer to **interface{}** |  | [optional] 
**Traceback** | Pointer to **NullableString** |  | [optional] 
**CancelType** | Pointer to [**JobResultCancelType**](JobResultCancelType.md) |  | [optional] 
**DateCanceled** | Pointer to **NullableTime** | Timestamp at which the job was canceled | [optional] 
**JobModel** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**User** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ScheduledJob** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CanceledBy** | Pointer to [**NullableJobResultCanceledBy**](JobResultCanceledBy.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Custom field data for this object, keyed by each applicable Custom Field&#39;s &#x60;key&#x60;. Value types vary with the custom field&#39;s type (text, integer, boolean, date, URL, JSON, select, multi-select); undefined values are &#x60;null&#x60;. On write, the payload is merged with existing values (PATCH-style: keys omitted from the payload are left untouched), and keys that do not correspond to a defined custom field are ignored. | [optional] 

## Methods

### NewJobResultRequest

`func NewJobResultRequest(name string, ) *JobResultRequest`

NewJobResultRequest instantiates a new JobResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResultRequestWithDefaults

`func NewJobResultRequestWithDefaults() *JobResultRequest`

NewJobResultRequestWithDefaults instantiates a new JobResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobResultRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobResultRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobResultRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobResultRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *JobResultRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobResultRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobResultRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTaskName

`func (o *JobResultRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *JobResultRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *JobResultRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *JobResultRequest) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### SetTaskNameNil

`func (o *JobResultRequest) SetTaskNameNil(b bool)`

 SetTaskNameNil sets the value for TaskName to be an explicit nil

### UnsetTaskName
`func (o *JobResultRequest) UnsetTaskName()`

UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
### GetDateStarted

`func (o *JobResultRequest) GetDateStarted() time.Time`

GetDateStarted returns the DateStarted field if non-nil, zero value otherwise.

### GetDateStartedOk

`func (o *JobResultRequest) GetDateStartedOk() (*time.Time, bool)`

GetDateStartedOk returns a tuple with the DateStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStarted

`func (o *JobResultRequest) SetDateStarted(v time.Time)`

SetDateStarted sets DateStarted field to given value.

### HasDateStarted

`func (o *JobResultRequest) HasDateStarted() bool`

HasDateStarted returns a boolean if a field has been set.

### SetDateStartedNil

`func (o *JobResultRequest) SetDateStartedNil(b bool)`

 SetDateStartedNil sets the value for DateStarted to be an explicit nil

### UnsetDateStarted
`func (o *JobResultRequest) UnsetDateStarted()`

UnsetDateStarted ensures that no value is present for DateStarted, not even an explicit nil
### GetDateDone

`func (o *JobResultRequest) GetDateDone() time.Time`

GetDateDone returns the DateDone field if non-nil, zero value otherwise.

### GetDateDoneOk

`func (o *JobResultRequest) GetDateDoneOk() (*time.Time, bool)`

GetDateDoneOk returns a tuple with the DateDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDone

`func (o *JobResultRequest) SetDateDone(v time.Time)`

SetDateDone sets DateDone field to given value.

### HasDateDone

`func (o *JobResultRequest) HasDateDone() bool`

HasDateDone returns a boolean if a field has been set.

### SetDateDoneNil

`func (o *JobResultRequest) SetDateDoneNil(b bool)`

 SetDateDoneNil sets the value for DateDone to be an explicit nil

### UnsetDateDone
`func (o *JobResultRequest) UnsetDateDone()`

UnsetDateDone ensures that no value is present for DateDone, not even an explicit nil
### GetWorker

`func (o *JobResultRequest) GetWorker() string`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *JobResultRequest) GetWorkerOk() (*string, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *JobResultRequest) SetWorker(v string)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *JobResultRequest) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### SetWorkerNil

`func (o *JobResultRequest) SetWorkerNil(b bool)`

 SetWorkerNil sets the value for Worker to be an explicit nil

### UnsetWorker
`func (o *JobResultRequest) UnsetWorker()`

UnsetWorker ensures that no value is present for Worker, not even an explicit nil
### GetTaskArgs

`func (o *JobResultRequest) GetTaskArgs() interface{}`

GetTaskArgs returns the TaskArgs field if non-nil, zero value otherwise.

### GetTaskArgsOk

`func (o *JobResultRequest) GetTaskArgsOk() (*interface{}, bool)`

GetTaskArgsOk returns a tuple with the TaskArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArgs

`func (o *JobResultRequest) SetTaskArgs(v interface{})`

SetTaskArgs sets TaskArgs field to given value.

### HasTaskArgs

`func (o *JobResultRequest) HasTaskArgs() bool`

HasTaskArgs returns a boolean if a field has been set.

### SetTaskArgsNil

`func (o *JobResultRequest) SetTaskArgsNil(b bool)`

 SetTaskArgsNil sets the value for TaskArgs to be an explicit nil

### UnsetTaskArgs
`func (o *JobResultRequest) UnsetTaskArgs()`

UnsetTaskArgs ensures that no value is present for TaskArgs, not even an explicit nil
### GetTaskKwargs

`func (o *JobResultRequest) GetTaskKwargs() interface{}`

GetTaskKwargs returns the TaskKwargs field if non-nil, zero value otherwise.

### GetTaskKwargsOk

`func (o *JobResultRequest) GetTaskKwargsOk() (*interface{}, bool)`

GetTaskKwargsOk returns a tuple with the TaskKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskKwargs

`func (o *JobResultRequest) SetTaskKwargs(v interface{})`

SetTaskKwargs sets TaskKwargs field to given value.

### HasTaskKwargs

`func (o *JobResultRequest) HasTaskKwargs() bool`

HasTaskKwargs returns a boolean if a field has been set.

### SetTaskKwargsNil

`func (o *JobResultRequest) SetTaskKwargsNil(b bool)`

 SetTaskKwargsNil sets the value for TaskKwargs to be an explicit nil

### UnsetTaskKwargs
`func (o *JobResultRequest) UnsetTaskKwargs()`

UnsetTaskKwargs ensures that no value is present for TaskKwargs, not even an explicit nil
### GetCeleryKwargs

`func (o *JobResultRequest) GetCeleryKwargs() interface{}`

GetCeleryKwargs returns the CeleryKwargs field if non-nil, zero value otherwise.

### GetCeleryKwargsOk

`func (o *JobResultRequest) GetCeleryKwargsOk() (*interface{}, bool)`

GetCeleryKwargsOk returns a tuple with the CeleryKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeleryKwargs

`func (o *JobResultRequest) SetCeleryKwargs(v interface{})`

SetCeleryKwargs sets CeleryKwargs field to given value.

### HasCeleryKwargs

`func (o *JobResultRequest) HasCeleryKwargs() bool`

HasCeleryKwargs returns a boolean if a field has been set.

### SetCeleryKwargsNil

`func (o *JobResultRequest) SetCeleryKwargsNil(b bool)`

 SetCeleryKwargsNil sets the value for CeleryKwargs to be an explicit nil

### UnsetCeleryKwargs
`func (o *JobResultRequest) UnsetCeleryKwargs()`

UnsetCeleryKwargs ensures that no value is present for CeleryKwargs, not even an explicit nil
### GetTraceback

`func (o *JobResultRequest) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *JobResultRequest) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *JobResultRequest) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *JobResultRequest) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.

### SetTracebackNil

`func (o *JobResultRequest) SetTracebackNil(b bool)`

 SetTracebackNil sets the value for Traceback to be an explicit nil

### UnsetTraceback
`func (o *JobResultRequest) UnsetTraceback()`

UnsetTraceback ensures that no value is present for Traceback, not even an explicit nil
### GetCancelType

`func (o *JobResultRequest) GetCancelType() JobResultCancelType`

GetCancelType returns the CancelType field if non-nil, zero value otherwise.

### GetCancelTypeOk

`func (o *JobResultRequest) GetCancelTypeOk() (*JobResultCancelType, bool)`

GetCancelTypeOk returns a tuple with the CancelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelType

`func (o *JobResultRequest) SetCancelType(v JobResultCancelType)`

SetCancelType sets CancelType field to given value.

### HasCancelType

`func (o *JobResultRequest) HasCancelType() bool`

HasCancelType returns a boolean if a field has been set.

### GetDateCanceled

`func (o *JobResultRequest) GetDateCanceled() time.Time`

GetDateCanceled returns the DateCanceled field if non-nil, zero value otherwise.

### GetDateCanceledOk

`func (o *JobResultRequest) GetDateCanceledOk() (*time.Time, bool)`

GetDateCanceledOk returns a tuple with the DateCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCanceled

`func (o *JobResultRequest) SetDateCanceled(v time.Time)`

SetDateCanceled sets DateCanceled field to given value.

### HasDateCanceled

`func (o *JobResultRequest) HasDateCanceled() bool`

HasDateCanceled returns a boolean if a field has been set.

### SetDateCanceledNil

`func (o *JobResultRequest) SetDateCanceledNil(b bool)`

 SetDateCanceledNil sets the value for DateCanceled to be an explicit nil

### UnsetDateCanceled
`func (o *JobResultRequest) UnsetDateCanceled()`

UnsetDateCanceled ensures that no value is present for DateCanceled, not even an explicit nil
### GetJobModel

`func (o *JobResultRequest) GetJobModel() ApprovalWorkflowUser`

GetJobModel returns the JobModel field if non-nil, zero value otherwise.

### GetJobModelOk

`func (o *JobResultRequest) GetJobModelOk() (*ApprovalWorkflowUser, bool)`

GetJobModelOk returns a tuple with the JobModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobModel

`func (o *JobResultRequest) SetJobModel(v ApprovalWorkflowUser)`

SetJobModel sets JobModel field to given value.

### HasJobModel

`func (o *JobResultRequest) HasJobModel() bool`

HasJobModel returns a boolean if a field has been set.

### SetJobModelNil

`func (o *JobResultRequest) SetJobModelNil(b bool)`

 SetJobModelNil sets the value for JobModel to be an explicit nil

### UnsetJobModel
`func (o *JobResultRequest) UnsetJobModel()`

UnsetJobModel ensures that no value is present for JobModel, not even an explicit nil
### GetUser

`func (o *JobResultRequest) GetUser() ApprovalWorkflowUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *JobResultRequest) GetUserOk() (*ApprovalWorkflowUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *JobResultRequest) SetUser(v ApprovalWorkflowUser)`

SetUser sets User field to given value.

### HasUser

`func (o *JobResultRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *JobResultRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *JobResultRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetScheduledJob

`func (o *JobResultRequest) GetScheduledJob() ApprovalWorkflowUser`

GetScheduledJob returns the ScheduledJob field if non-nil, zero value otherwise.

### GetScheduledJobOk

`func (o *JobResultRequest) GetScheduledJobOk() (*ApprovalWorkflowUser, bool)`

GetScheduledJobOk returns a tuple with the ScheduledJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledJob

`func (o *JobResultRequest) SetScheduledJob(v ApprovalWorkflowUser)`

SetScheduledJob sets ScheduledJob field to given value.

### HasScheduledJob

`func (o *JobResultRequest) HasScheduledJob() bool`

HasScheduledJob returns a boolean if a field has been set.

### SetScheduledJobNil

`func (o *JobResultRequest) SetScheduledJobNil(b bool)`

 SetScheduledJobNil sets the value for ScheduledJob to be an explicit nil

### UnsetScheduledJob
`func (o *JobResultRequest) UnsetScheduledJob()`

UnsetScheduledJob ensures that no value is present for ScheduledJob, not even an explicit nil
### GetCanceledBy

`func (o *JobResultRequest) GetCanceledBy() JobResultCanceledBy`

GetCanceledBy returns the CanceledBy field if non-nil, zero value otherwise.

### GetCanceledByOk

`func (o *JobResultRequest) GetCanceledByOk() (*JobResultCanceledBy, bool)`

GetCanceledByOk returns a tuple with the CanceledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledBy

`func (o *JobResultRequest) SetCanceledBy(v JobResultCanceledBy)`

SetCanceledBy sets CanceledBy field to given value.

### HasCanceledBy

`func (o *JobResultRequest) HasCanceledBy() bool`

HasCanceledBy returns a boolean if a field has been set.

### SetCanceledByNil

`func (o *JobResultRequest) SetCanceledByNil(b bool)`

 SetCanceledByNil sets the value for CanceledBy to be an explicit nil

### UnsetCanceledBy
`func (o *JobResultRequest) UnsetCanceledBy()`

UnsetCanceledBy ensures that no value is present for CanceledBy, not even an explicit nil
### GetCustomFields

`func (o *JobResultRequest) GetCustomFields() map[string]*interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *JobResultRequest) GetCustomFieldsOk() (*map[string]*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *JobResultRequest) SetCustomFields(v map[string]*interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *JobResultRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


