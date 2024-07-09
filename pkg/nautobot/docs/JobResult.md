# JobResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Status** | [**JobResultStatus**](JobResultStatus.md) |  | 
**Name** | **string** |  | 
**TaskName** | Pointer to **NullableString** | Registered name of the Celery task for this job. Internal use only. | [optional] 
**DateCreated** | **time.Time** |  | [readonly] 
**DateDone** | Pointer to **NullableTime** |  | [optional] 
**Result** | **map[string]interface{}** | The data returned by the task | [readonly] 
**Worker** | Pointer to **NullableString** |  | [optional] 
**TaskArgs** | Pointer to **map[string]interface{}** |  | [optional] 
**TaskKwargs** | Pointer to **map[string]interface{}** |  | [optional] 
**CeleryKwargs** | Pointer to **map[string]interface{}** |  | [optional] 
**Traceback** | Pointer to **NullableString** |  | [optional] 
**Meta** | **map[string]interface{}** |  | [readonly] 
**JobModel** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**User** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**ScheduledJob** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Files** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [readonly] 

## Methods

### NewJobResult

`func NewJobResult(id string, objectType string, display string, url string, naturalSlug string, status JobResultStatus, name string, dateCreated time.Time, result map[string]interface{}, meta map[string]interface{}, files []BulkWritableCableRequestStatus, ) *JobResult`

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

`func (o *JobResult) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *JobResult) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *JobResult) SetResult(v map[string]interface{})`

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

`func (o *JobResult) GetTaskArgs() map[string]interface{}`

GetTaskArgs returns the TaskArgs field if non-nil, zero value otherwise.

### GetTaskArgsOk

`func (o *JobResult) GetTaskArgsOk() (*map[string]interface{}, bool)`

GetTaskArgsOk returns a tuple with the TaskArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArgs

`func (o *JobResult) SetTaskArgs(v map[string]interface{})`

SetTaskArgs sets TaskArgs field to given value.

### HasTaskArgs

`func (o *JobResult) HasTaskArgs() bool`

HasTaskArgs returns a boolean if a field has been set.

### GetTaskKwargs

`func (o *JobResult) GetTaskKwargs() map[string]interface{}`

GetTaskKwargs returns the TaskKwargs field if non-nil, zero value otherwise.

### GetTaskKwargsOk

`func (o *JobResult) GetTaskKwargsOk() (*map[string]interface{}, bool)`

GetTaskKwargsOk returns a tuple with the TaskKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskKwargs

`func (o *JobResult) SetTaskKwargs(v map[string]interface{})`

SetTaskKwargs sets TaskKwargs field to given value.

### HasTaskKwargs

`func (o *JobResult) HasTaskKwargs() bool`

HasTaskKwargs returns a boolean if a field has been set.

### GetCeleryKwargs

`func (o *JobResult) GetCeleryKwargs() map[string]interface{}`

GetCeleryKwargs returns the CeleryKwargs field if non-nil, zero value otherwise.

### GetCeleryKwargsOk

`func (o *JobResult) GetCeleryKwargsOk() (*map[string]interface{}, bool)`

GetCeleryKwargsOk returns a tuple with the CeleryKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeleryKwargs

`func (o *JobResult) SetCeleryKwargs(v map[string]interface{})`

SetCeleryKwargs sets CeleryKwargs field to given value.

### HasCeleryKwargs

`func (o *JobResult) HasCeleryKwargs() bool`

HasCeleryKwargs returns a boolean if a field has been set.

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

`func (o *JobResult) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JobResult) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JobResult) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### SetMetaNil

`func (o *JobResult) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *JobResult) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetJobModel

`func (o *JobResult) GetJobModel() BulkWritableCircuitRequestTenant`

GetJobModel returns the JobModel field if non-nil, zero value otherwise.

### GetJobModelOk

`func (o *JobResult) GetJobModelOk() (*BulkWritableCircuitRequestTenant, bool)`

GetJobModelOk returns a tuple with the JobModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobModel

`func (o *JobResult) SetJobModel(v BulkWritableCircuitRequestTenant)`

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

`func (o *JobResult) GetUser() BulkWritableCircuitRequestTenant`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *JobResult) GetUserOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *JobResult) SetUser(v BulkWritableCircuitRequestTenant)`

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

`func (o *JobResult) GetScheduledJob() BulkWritableCircuitRequestTenant`

GetScheduledJob returns the ScheduledJob field if non-nil, zero value otherwise.

### GetScheduledJobOk

`func (o *JobResult) GetScheduledJobOk() (*BulkWritableCircuitRequestTenant, bool)`

GetScheduledJobOk returns a tuple with the ScheduledJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledJob

`func (o *JobResult) SetScheduledJob(v BulkWritableCircuitRequestTenant)`

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

### GetFiles

`func (o *JobResult) GetFiles() []BulkWritableCableRequestStatus`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *JobResult) GetFilesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *JobResult) SetFiles(v []BulkWritableCableRequestStatus)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


