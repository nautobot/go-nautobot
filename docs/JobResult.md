# JobResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Status** | [**JobResultStatus**](JobResultStatus.md) |  | 
**Name** | **string** |  | 
**TaskName** | Pointer to **NullableString** | Registered name of the Celery task for this job. Internal use only. | [optional] 
**DateCreated** | **time.Time** |  | [readonly] 
**DateStarted** | Pointer to **NullableTime** |  | [optional] 
**DateDone** | Pointer to **NullableTime** |  | [optional] 
**Result** | **interface{}** | The data returned by the task | [readonly] 
**Worker** | Pointer to **NullableString** |  | [optional] 
**TaskArgs** | Pointer to **interface{}** |  | [optional] 
**TaskKwargs** | Pointer to **interface{}** |  | [optional] 
**CeleryKwargs** | Pointer to **interface{}** |  | [optional] 
**Traceback** | Pointer to **NullableString** |  | [optional] 
**Meta** | **interface{}** |  | [readonly] 
**DebugLogCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**SuccessLogCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**InfoLogCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**WarningLogCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**ErrorLogCount** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**JobModel** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**User** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ScheduledJob** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ComputedFields** | **map[string]interface{}** |  | [readonly] 

## Methods

### NewJobResult

`func NewJobResult(objectType string, display string, url string, naturalSlug string, status JobResultStatus, name string, dateCreated time.Time, result interface{}, meta interface{}, computedFields map[string]interface{}, ) *JobResult`

NewJobResult instantiates a new JobResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResultWithDefaults

`func NewJobResultWithDefaults() *JobResult`

NewJobResultWithDefaults instantiates a new JobResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *JobResult) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *JobResult) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *JobResult) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *JobResult) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *JobResult) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *JobResult) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *JobResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JobResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JobResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *JobResult) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *JobResult) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *JobResult) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetStatus

`func (o *JobResult) GetStatus() JobResultStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobResult) GetStatusOk() (*JobResultStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobResult) SetStatus(v JobResultStatus)`

SetStatus sets Status field to given value.


### GetName

`func (o *JobResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobResult) SetName(v string)`

SetName sets Name field to given value.


### GetTaskName

`func (o *JobResult) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *JobResult) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *JobResult) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *JobResult) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### SetTaskNameNil

`func (o *JobResult) SetTaskNameNil(b bool)`

 SetTaskNameNil sets the value for TaskName to be an explicit nil

### UnsetTaskName
`func (o *JobResult) UnsetTaskName()`

UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
### GetDateCreated

`func (o *JobResult) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *JobResult) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *JobResult) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetDateStarted

`func (o *JobResult) GetDateStarted() time.Time`

GetDateStarted returns the DateStarted field if non-nil, zero value otherwise.

### GetDateStartedOk

`func (o *JobResult) GetDateStartedOk() (*time.Time, bool)`

GetDateStartedOk returns a tuple with the DateStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStarted

`func (o *JobResult) SetDateStarted(v time.Time)`

SetDateStarted sets DateStarted field to given value.

### HasDateStarted

`func (o *JobResult) HasDateStarted() bool`

HasDateStarted returns a boolean if a field has been set.

### SetDateStartedNil

`func (o *JobResult) SetDateStartedNil(b bool)`

 SetDateStartedNil sets the value for DateStarted to be an explicit nil

### UnsetDateStarted
`func (o *JobResult) UnsetDateStarted()`

UnsetDateStarted ensures that no value is present for DateStarted, not even an explicit nil
### GetDateDone

`func (o *JobResult) GetDateDone() time.Time`

GetDateDone returns the DateDone field if non-nil, zero value otherwise.

### GetDateDoneOk

`func (o *JobResult) GetDateDoneOk() (*time.Time, bool)`

GetDateDoneOk returns a tuple with the DateDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDone

`func (o *JobResult) SetDateDone(v time.Time)`

SetDateDone sets DateDone field to given value.

### HasDateDone

`func (o *JobResult) HasDateDone() bool`

HasDateDone returns a boolean if a field has been set.

### SetDateDoneNil

`func (o *JobResult) SetDateDoneNil(b bool)`

 SetDateDoneNil sets the value for DateDone to be an explicit nil

### UnsetDateDone
`func (o *JobResult) UnsetDateDone()`

UnsetDateDone ensures that no value is present for DateDone, not even an explicit nil
### GetResult

`func (o *JobResult) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *JobResult) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *JobResult) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *JobResult) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *JobResult) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetWorker

`func (o *JobResult) GetWorker() string`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *JobResult) GetWorkerOk() (*string, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *JobResult) SetWorker(v string)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *JobResult) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### SetWorkerNil

`func (o *JobResult) SetWorkerNil(b bool)`

 SetWorkerNil sets the value for Worker to be an explicit nil

### UnsetWorker
`func (o *JobResult) UnsetWorker()`

UnsetWorker ensures that no value is present for Worker, not even an explicit nil
### GetTaskArgs

`func (o *JobResult) GetTaskArgs() interface{}`

GetTaskArgs returns the TaskArgs field if non-nil, zero value otherwise.

### GetTaskArgsOk

`func (o *JobResult) GetTaskArgsOk() (*interface{}, bool)`

GetTaskArgsOk returns a tuple with the TaskArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArgs

`func (o *JobResult) SetTaskArgs(v interface{})`

SetTaskArgs sets TaskArgs field to given value.

### HasTaskArgs

`func (o *JobResult) HasTaskArgs() bool`

HasTaskArgs returns a boolean if a field has been set.

### SetTaskArgsNil

`func (o *JobResult) SetTaskArgsNil(b bool)`

 SetTaskArgsNil sets the value for TaskArgs to be an explicit nil

### UnsetTaskArgs
`func (o *JobResult) UnsetTaskArgs()`

UnsetTaskArgs ensures that no value is present for TaskArgs, not even an explicit nil
### GetTaskKwargs

`func (o *JobResult) GetTaskKwargs() interface{}`

GetTaskKwargs returns the TaskKwargs field if non-nil, zero value otherwise.

### GetTaskKwargsOk

`func (o *JobResult) GetTaskKwargsOk() (*interface{}, bool)`

GetTaskKwargsOk returns a tuple with the TaskKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskKwargs

`func (o *JobResult) SetTaskKwargs(v interface{})`

SetTaskKwargs sets TaskKwargs field to given value.

### HasTaskKwargs

`func (o *JobResult) HasTaskKwargs() bool`

HasTaskKwargs returns a boolean if a field has been set.

### SetTaskKwargsNil

`func (o *JobResult) SetTaskKwargsNil(b bool)`

 SetTaskKwargsNil sets the value for TaskKwargs to be an explicit nil

### UnsetTaskKwargs
`func (o *JobResult) UnsetTaskKwargs()`

UnsetTaskKwargs ensures that no value is present for TaskKwargs, not even an explicit nil
### GetCeleryKwargs

`func (o *JobResult) GetCeleryKwargs() interface{}`

GetCeleryKwargs returns the CeleryKwargs field if non-nil, zero value otherwise.

### GetCeleryKwargsOk

`func (o *JobResult) GetCeleryKwargsOk() (*interface{}, bool)`

GetCeleryKwargsOk returns a tuple with the CeleryKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeleryKwargs

`func (o *JobResult) SetCeleryKwargs(v interface{})`

SetCeleryKwargs sets CeleryKwargs field to given value.

### HasCeleryKwargs

`func (o *JobResult) HasCeleryKwargs() bool`

HasCeleryKwargs returns a boolean if a field has been set.

### SetCeleryKwargsNil

`func (o *JobResult) SetCeleryKwargsNil(b bool)`

 SetCeleryKwargsNil sets the value for CeleryKwargs to be an explicit nil

### UnsetCeleryKwargs
`func (o *JobResult) UnsetCeleryKwargs()`

UnsetCeleryKwargs ensures that no value is present for CeleryKwargs, not even an explicit nil
### GetTraceback

`func (o *JobResult) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *JobResult) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *JobResult) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *JobResult) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.

### SetTracebackNil

`func (o *JobResult) SetTracebackNil(b bool)`

 SetTracebackNil sets the value for Traceback to be an explicit nil

### UnsetTraceback
`func (o *JobResult) UnsetTraceback()`

UnsetTraceback ensures that no value is present for Traceback, not even an explicit nil
### GetMeta

`func (o *JobResult) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JobResult) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JobResult) SetMeta(v interface{})`

SetMeta sets Meta field to given value.


### SetMetaNil

`func (o *JobResult) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *JobResult) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetDebugLogCount

`func (o *JobResult) GetDebugLogCount() int32`

GetDebugLogCount returns the DebugLogCount field if non-nil, zero value otherwise.

### GetDebugLogCountOk

`func (o *JobResult) GetDebugLogCountOk() (*int32, bool)`

GetDebugLogCountOk returns a tuple with the DebugLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLogCount

`func (o *JobResult) SetDebugLogCount(v int32)`

SetDebugLogCount sets DebugLogCount field to given value.

### HasDebugLogCount

`func (o *JobResult) HasDebugLogCount() bool`

HasDebugLogCount returns a boolean if a field has been set.

### SetDebugLogCountNil

`func (o *JobResult) SetDebugLogCountNil(b bool)`

 SetDebugLogCountNil sets the value for DebugLogCount to be an explicit nil

### UnsetDebugLogCount
`func (o *JobResult) UnsetDebugLogCount()`

UnsetDebugLogCount ensures that no value is present for DebugLogCount, not even an explicit nil
### GetSuccessLogCount

`func (o *JobResult) GetSuccessLogCount() int32`

GetSuccessLogCount returns the SuccessLogCount field if non-nil, zero value otherwise.

### GetSuccessLogCountOk

`func (o *JobResult) GetSuccessLogCountOk() (*int32, bool)`

GetSuccessLogCountOk returns a tuple with the SuccessLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessLogCount

`func (o *JobResult) SetSuccessLogCount(v int32)`

SetSuccessLogCount sets SuccessLogCount field to given value.

### HasSuccessLogCount

`func (o *JobResult) HasSuccessLogCount() bool`

HasSuccessLogCount returns a boolean if a field has been set.

### SetSuccessLogCountNil

`func (o *JobResult) SetSuccessLogCountNil(b bool)`

 SetSuccessLogCountNil sets the value for SuccessLogCount to be an explicit nil

### UnsetSuccessLogCount
`func (o *JobResult) UnsetSuccessLogCount()`

UnsetSuccessLogCount ensures that no value is present for SuccessLogCount, not even an explicit nil
### GetInfoLogCount

`func (o *JobResult) GetInfoLogCount() int32`

GetInfoLogCount returns the InfoLogCount field if non-nil, zero value otherwise.

### GetInfoLogCountOk

`func (o *JobResult) GetInfoLogCountOk() (*int32, bool)`

GetInfoLogCountOk returns a tuple with the InfoLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLogCount

`func (o *JobResult) SetInfoLogCount(v int32)`

SetInfoLogCount sets InfoLogCount field to given value.

### HasInfoLogCount

`func (o *JobResult) HasInfoLogCount() bool`

HasInfoLogCount returns a boolean if a field has been set.

### SetInfoLogCountNil

`func (o *JobResult) SetInfoLogCountNil(b bool)`

 SetInfoLogCountNil sets the value for InfoLogCount to be an explicit nil

### UnsetInfoLogCount
`func (o *JobResult) UnsetInfoLogCount()`

UnsetInfoLogCount ensures that no value is present for InfoLogCount, not even an explicit nil
### GetWarningLogCount

`func (o *JobResult) GetWarningLogCount() int32`

GetWarningLogCount returns the WarningLogCount field if non-nil, zero value otherwise.

### GetWarningLogCountOk

`func (o *JobResult) GetWarningLogCountOk() (*int32, bool)`

GetWarningLogCountOk returns a tuple with the WarningLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningLogCount

`func (o *JobResult) SetWarningLogCount(v int32)`

SetWarningLogCount sets WarningLogCount field to given value.

### HasWarningLogCount

`func (o *JobResult) HasWarningLogCount() bool`

HasWarningLogCount returns a boolean if a field has been set.

### SetWarningLogCountNil

`func (o *JobResult) SetWarningLogCountNil(b bool)`

 SetWarningLogCountNil sets the value for WarningLogCount to be an explicit nil

### UnsetWarningLogCount
`func (o *JobResult) UnsetWarningLogCount()`

UnsetWarningLogCount ensures that no value is present for WarningLogCount, not even an explicit nil
### GetErrorLogCount

`func (o *JobResult) GetErrorLogCount() int32`

GetErrorLogCount returns the ErrorLogCount field if non-nil, zero value otherwise.

### GetErrorLogCountOk

`func (o *JobResult) GetErrorLogCountOk() (*int32, bool)`

GetErrorLogCountOk returns a tuple with the ErrorLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLogCount

`func (o *JobResult) SetErrorLogCount(v int32)`

SetErrorLogCount sets ErrorLogCount field to given value.

### HasErrorLogCount

`func (o *JobResult) HasErrorLogCount() bool`

HasErrorLogCount returns a boolean if a field has been set.

### SetErrorLogCountNil

`func (o *JobResult) SetErrorLogCountNil(b bool)`

 SetErrorLogCountNil sets the value for ErrorLogCount to be an explicit nil

### UnsetErrorLogCount
`func (o *JobResult) UnsetErrorLogCount()`

UnsetErrorLogCount ensures that no value is present for ErrorLogCount, not even an explicit nil
### GetJobModel

`func (o *JobResult) GetJobModel() ApprovalWorkflowUser`

GetJobModel returns the JobModel field if non-nil, zero value otherwise.

### GetJobModelOk

`func (o *JobResult) GetJobModelOk() (*ApprovalWorkflowUser, bool)`

GetJobModelOk returns a tuple with the JobModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobModel

`func (o *JobResult) SetJobModel(v ApprovalWorkflowUser)`

SetJobModel sets JobModel field to given value.

### HasJobModel

`func (o *JobResult) HasJobModel() bool`

HasJobModel returns a boolean if a field has been set.

### SetJobModelNil

`func (o *JobResult) SetJobModelNil(b bool)`

 SetJobModelNil sets the value for JobModel to be an explicit nil

### UnsetJobModel
`func (o *JobResult) UnsetJobModel()`

UnsetJobModel ensures that no value is present for JobModel, not even an explicit nil
### GetUser

`func (o *JobResult) GetUser() ApprovalWorkflowUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *JobResult) GetUserOk() (*ApprovalWorkflowUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *JobResult) SetUser(v ApprovalWorkflowUser)`

SetUser sets User field to given value.

### HasUser

`func (o *JobResult) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *JobResult) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *JobResult) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetScheduledJob

`func (o *JobResult) GetScheduledJob() ApprovalWorkflowUser`

GetScheduledJob returns the ScheduledJob field if non-nil, zero value otherwise.

### GetScheduledJobOk

`func (o *JobResult) GetScheduledJobOk() (*ApprovalWorkflowUser, bool)`

GetScheduledJobOk returns a tuple with the ScheduledJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledJob

`func (o *JobResult) SetScheduledJob(v ApprovalWorkflowUser)`

SetScheduledJob sets ScheduledJob field to given value.

### HasScheduledJob

`func (o *JobResult) HasScheduledJob() bool`

HasScheduledJob returns a boolean if a field has been set.

### SetScheduledJobNil

`func (o *JobResult) SetScheduledJobNil(b bool)`

 SetScheduledJobNil sets the value for ScheduledJob to be an explicit nil

### UnsetScheduledJob
`func (o *JobResult) UnsetScheduledJob()`

UnsetScheduledJob ensures that no value is present for ScheduledJob, not even an explicit nil
### GetCustomFields

`func (o *JobResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *JobResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *JobResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *JobResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComputedFields

`func (o *JobResult) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *JobResult) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *JobResult) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


