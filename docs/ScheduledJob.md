# ScheduledJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Name** | **string** | Human-readable description of this scheduled task | 
**Task** | **string** | The name of the Celery task that should be run. (Example: \&quot;proj.tasks.import_contacts\&quot;) | 
**Interval** | [**JobExecutionTypeIntervalChoices**](JobExecutionTypeIntervalChoices.md) |  | 
**Args** | Pointer to **interface{}** |  | [optional] 
**Kwargs** | Pointer to **interface{}** |  | [optional] 
**CeleryKwargs** | Pointer to **interface{}** |  | [optional] 
**Queue** | Pointer to **string** | Queue defined in CELERY_TASK_QUEUES. Leave empty for default queuing. | [optional] 
**OneOff** | Pointer to **bool** | If True, the schedule will only run the task a single time | [optional] 
**StartTime** | **time.Time** | Datetime when the schedule should begin triggering the task to run | 
**Enabled** | Pointer to **bool** | Set to False to disable the schedule | [optional] 
**LastRunAt** | **NullableTime** | Datetime that the schedule last triggered the task to run. Reset to None if enabled is set to False. | [readonly] 
**TotalRunCount** | Pointer to **int32** | Running count of how many times the schedule has triggered the task | [optional] [readonly] 
**DateChanged** | **time.Time** | Datetime that this scheduled job was last modified | [readonly] 
**Description** | Pointer to **string** | Detailed description about the details of this scheduled job | [optional] 
**ApprovalRequired** | Pointer to **bool** |  | [optional] 
**ApprovedAt** | **NullableTime** | Datetime that the schedule was approved | [readonly] 
**Crontab** | Pointer to **string** | Cronjob syntax string for custom scheduling | [optional] 
**JobModel** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**User** | Pointer to [**NullableScheduledJobUser**](ScheduledJobUser.md) |  | [optional] 
**ApprovedByUser** | Pointer to [**NullableScheduledJobApprovedByUser**](ScheduledJobApprovedByUser.md) |  | [optional] 

## Methods

### NewScheduledJob

`func NewScheduledJob(id string, objectType string, display string, url string, naturalSlug string, name string, task string, interval JobExecutionTypeIntervalChoices, startTime time.Time, lastRunAt NullableTime, dateChanged time.Time, approvedAt NullableTime, ) *ScheduledJob`

NewScheduledJob instantiates a new ScheduledJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledJobWithDefaults

`func NewScheduledJobWithDefaults() *ScheduledJob`

NewScheduledJobWithDefaults instantiates a new ScheduledJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledJob) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *ScheduledJob) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ScheduledJob) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ScheduledJob) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ScheduledJob) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ScheduledJob) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ScheduledJob) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ScheduledJob) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ScheduledJob) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ScheduledJob) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ScheduledJob) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ScheduledJob) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ScheduledJob) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetName

`func (o *ScheduledJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduledJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduledJob) SetName(v string)`

SetName sets Name field to given value.


### GetTask

`func (o *ScheduledJob) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ScheduledJob) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ScheduledJob) SetTask(v string)`

SetTask sets Task field to given value.


### GetInterval

`func (o *ScheduledJob) GetInterval() JobExecutionTypeIntervalChoices`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ScheduledJob) GetIntervalOk() (*JobExecutionTypeIntervalChoices, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ScheduledJob) SetInterval(v JobExecutionTypeIntervalChoices)`

SetInterval sets Interval field to given value.


### GetArgs

`func (o *ScheduledJob) GetArgs() interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ScheduledJob) GetArgsOk() (*interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ScheduledJob) SetArgs(v interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ScheduledJob) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *ScheduledJob) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *ScheduledJob) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetKwargs

`func (o *ScheduledJob) GetKwargs() interface{}`

GetKwargs returns the Kwargs field if non-nil, zero value otherwise.

### GetKwargsOk

`func (o *ScheduledJob) GetKwargsOk() (*interface{}, bool)`

GetKwargsOk returns a tuple with the Kwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKwargs

`func (o *ScheduledJob) SetKwargs(v interface{})`

SetKwargs sets Kwargs field to given value.

### HasKwargs

`func (o *ScheduledJob) HasKwargs() bool`

HasKwargs returns a boolean if a field has been set.

### SetKwargsNil

`func (o *ScheduledJob) SetKwargsNil(b bool)`

 SetKwargsNil sets the value for Kwargs to be an explicit nil

### UnsetKwargs
`func (o *ScheduledJob) UnsetKwargs()`

UnsetKwargs ensures that no value is present for Kwargs, not even an explicit nil
### GetCeleryKwargs

`func (o *ScheduledJob) GetCeleryKwargs() interface{}`

GetCeleryKwargs returns the CeleryKwargs field if non-nil, zero value otherwise.

### GetCeleryKwargsOk

`func (o *ScheduledJob) GetCeleryKwargsOk() (*interface{}, bool)`

GetCeleryKwargsOk returns a tuple with the CeleryKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeleryKwargs

`func (o *ScheduledJob) SetCeleryKwargs(v interface{})`

SetCeleryKwargs sets CeleryKwargs field to given value.

### HasCeleryKwargs

`func (o *ScheduledJob) HasCeleryKwargs() bool`

HasCeleryKwargs returns a boolean if a field has been set.

### SetCeleryKwargsNil

`func (o *ScheduledJob) SetCeleryKwargsNil(b bool)`

 SetCeleryKwargsNil sets the value for CeleryKwargs to be an explicit nil

### UnsetCeleryKwargs
`func (o *ScheduledJob) UnsetCeleryKwargs()`

UnsetCeleryKwargs ensures that no value is present for CeleryKwargs, not even an explicit nil
### GetQueue

`func (o *ScheduledJob) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *ScheduledJob) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *ScheduledJob) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *ScheduledJob) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetOneOff

`func (o *ScheduledJob) GetOneOff() bool`

GetOneOff returns the OneOff field if non-nil, zero value otherwise.

### GetOneOffOk

`func (o *ScheduledJob) GetOneOffOk() (*bool, bool)`

GetOneOffOk returns a tuple with the OneOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneOff

`func (o *ScheduledJob) SetOneOff(v bool)`

SetOneOff sets OneOff field to given value.

### HasOneOff

`func (o *ScheduledJob) HasOneOff() bool`

HasOneOff returns a boolean if a field has been set.

### GetStartTime

`func (o *ScheduledJob) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScheduledJob) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScheduledJob) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEnabled

`func (o *ScheduledJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScheduledJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScheduledJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ScheduledJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastRunAt

`func (o *ScheduledJob) GetLastRunAt() time.Time`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *ScheduledJob) GetLastRunAtOk() (*time.Time, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *ScheduledJob) SetLastRunAt(v time.Time)`

SetLastRunAt sets LastRunAt field to given value.


### SetLastRunAtNil

`func (o *ScheduledJob) SetLastRunAtNil(b bool)`

 SetLastRunAtNil sets the value for LastRunAt to be an explicit nil

### UnsetLastRunAt
`func (o *ScheduledJob) UnsetLastRunAt()`

UnsetLastRunAt ensures that no value is present for LastRunAt, not even an explicit nil
### GetTotalRunCount

`func (o *ScheduledJob) GetTotalRunCount() int32`

GetTotalRunCount returns the TotalRunCount field if non-nil, zero value otherwise.

### GetTotalRunCountOk

`func (o *ScheduledJob) GetTotalRunCountOk() (*int32, bool)`

GetTotalRunCountOk returns a tuple with the TotalRunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRunCount

`func (o *ScheduledJob) SetTotalRunCount(v int32)`

SetTotalRunCount sets TotalRunCount field to given value.

### HasTotalRunCount

`func (o *ScheduledJob) HasTotalRunCount() bool`

HasTotalRunCount returns a boolean if a field has been set.

### GetDateChanged

`func (o *ScheduledJob) GetDateChanged() time.Time`

GetDateChanged returns the DateChanged field if non-nil, zero value otherwise.

### GetDateChangedOk

`func (o *ScheduledJob) GetDateChangedOk() (*time.Time, bool)`

GetDateChangedOk returns a tuple with the DateChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateChanged

`func (o *ScheduledJob) SetDateChanged(v time.Time)`

SetDateChanged sets DateChanged field to given value.


### GetDescription

`func (o *ScheduledJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduledJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduledJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduledJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApprovalRequired

`func (o *ScheduledJob) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *ScheduledJob) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *ScheduledJob) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.

### HasApprovalRequired

`func (o *ScheduledJob) HasApprovalRequired() bool`

HasApprovalRequired returns a boolean if a field has been set.

### GetApprovedAt

`func (o *ScheduledJob) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *ScheduledJob) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *ScheduledJob) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.


### SetApprovedAtNil

`func (o *ScheduledJob) SetApprovedAtNil(b bool)`

 SetApprovedAtNil sets the value for ApprovedAt to be an explicit nil

### UnsetApprovedAt
`func (o *ScheduledJob) UnsetApprovedAt()`

UnsetApprovedAt ensures that no value is present for ApprovedAt, not even an explicit nil
### GetCrontab

`func (o *ScheduledJob) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *ScheduledJob) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *ScheduledJob) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.

### HasCrontab

`func (o *ScheduledJob) HasCrontab() bool`

HasCrontab returns a boolean if a field has been set.

### GetJobModel

`func (o *ScheduledJob) GetJobModel() BulkWritableCircuitRequestTenant`

GetJobModel returns the JobModel field if non-nil, zero value otherwise.

### GetJobModelOk

`func (o *ScheduledJob) GetJobModelOk() (*BulkWritableCircuitRequestTenant, bool)`

GetJobModelOk returns a tuple with the JobModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobModel

`func (o *ScheduledJob) SetJobModel(v BulkWritableCircuitRequestTenant)`

SetJobModel sets JobModel field to given value.

### HasJobModel

`func (o *ScheduledJob) HasJobModel() bool`

HasJobModel returns a boolean if a field has been set.

### SetJobModelNil

`func (o *ScheduledJob) SetJobModelNil(b bool)`

 SetJobModelNil sets the value for JobModel to be an explicit nil

### UnsetJobModel
`func (o *ScheduledJob) UnsetJobModel()`

UnsetJobModel ensures that no value is present for JobModel, not even an explicit nil
### GetUser

`func (o *ScheduledJob) GetUser() ScheduledJobUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ScheduledJob) GetUserOk() (*ScheduledJobUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ScheduledJob) SetUser(v ScheduledJobUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ScheduledJob) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ScheduledJob) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ScheduledJob) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetApprovedByUser

`func (o *ScheduledJob) GetApprovedByUser() ScheduledJobApprovedByUser`

GetApprovedByUser returns the ApprovedByUser field if non-nil, zero value otherwise.

### GetApprovedByUserOk

`func (o *ScheduledJob) GetApprovedByUserOk() (*ScheduledJobApprovedByUser, bool)`

GetApprovedByUserOk returns a tuple with the ApprovedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUser

`func (o *ScheduledJob) SetApprovedByUser(v ScheduledJobApprovedByUser)`

SetApprovedByUser sets ApprovedByUser field to given value.

### HasApprovedByUser

`func (o *ScheduledJob) HasApprovedByUser() bool`

HasApprovedByUser returns a boolean if a field has been set.

### SetApprovedByUserNil

`func (o *ScheduledJob) SetApprovedByUserNil(b bool)`

 SetApprovedByUserNil sets the value for ApprovedByUser to be an explicit nil

### UnsetApprovedByUser
`func (o *ScheduledJob) UnsetApprovedByUser()`

UnsetApprovedByUser ensures that no value is present for ApprovedByUser, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


