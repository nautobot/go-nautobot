# JobLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**LogLevel** | Pointer to [**LogLevelEnum**](LogLevelEnum.md) |  | [optional] 
**Grouping** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**LogObject** | Pointer to **string** |  | [optional] 
**AbsoluteUrl** | Pointer to **string** |  | [optional] 
**JobResult** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Created** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewJobLogEntry

`func NewJobLogEntry(id string, objectType string, display string, url string, naturalSlug string, jobResult BulkWritableCableRequestStatus, ) *JobLogEntry`

NewJobLogEntry instantiates a new JobLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobLogEntryWithDefaults

`func NewJobLogEntryWithDefaults() *JobLogEntry`

NewJobLogEntryWithDefaults instantiates a new JobLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobLogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobLogEntry) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *JobLogEntry) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *JobLogEntry) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *JobLogEntry) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *JobLogEntry) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *JobLogEntry) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *JobLogEntry) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *JobLogEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JobLogEntry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JobLogEntry) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *JobLogEntry) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *JobLogEntry) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *JobLogEntry) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetLogLevel

`func (o *JobLogEntry) GetLogLevel() LogLevelEnum`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *JobLogEntry) GetLogLevelOk() (*LogLevelEnum, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *JobLogEntry) SetLogLevel(v LogLevelEnum)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *JobLogEntry) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetGrouping

`func (o *JobLogEntry) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *JobLogEntry) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *JobLogEntry) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *JobLogEntry) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetMessage

`func (o *JobLogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JobLogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JobLogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *JobLogEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLogObject

`func (o *JobLogEntry) GetLogObject() string`

GetLogObject returns the LogObject field if non-nil, zero value otherwise.

### GetLogObjectOk

`func (o *JobLogEntry) GetLogObjectOk() (*string, bool)`

GetLogObjectOk returns a tuple with the LogObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogObject

`func (o *JobLogEntry) SetLogObject(v string)`

SetLogObject sets LogObject field to given value.

### HasLogObject

`func (o *JobLogEntry) HasLogObject() bool`

HasLogObject returns a boolean if a field has been set.

### GetAbsoluteUrl

`func (o *JobLogEntry) GetAbsoluteUrl() string`

GetAbsoluteUrl returns the AbsoluteUrl field if non-nil, zero value otherwise.

### GetAbsoluteUrlOk

`func (o *JobLogEntry) GetAbsoluteUrlOk() (*string, bool)`

GetAbsoluteUrlOk returns a tuple with the AbsoluteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteUrl

`func (o *JobLogEntry) SetAbsoluteUrl(v string)`

SetAbsoluteUrl sets AbsoluteUrl field to given value.

### HasAbsoluteUrl

`func (o *JobLogEntry) HasAbsoluteUrl() bool`

HasAbsoluteUrl returns a boolean if a field has been set.

### GetJobResult

`func (o *JobLogEntry) GetJobResult() BulkWritableCableRequestStatus`

GetJobResult returns the JobResult field if non-nil, zero value otherwise.

### GetJobResultOk

`func (o *JobLogEntry) GetJobResultOk() (*BulkWritableCableRequestStatus, bool)`

GetJobResultOk returns a tuple with the JobResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResult

`func (o *JobLogEntry) SetJobResult(v BulkWritableCableRequestStatus)`

SetJobResult sets JobResult field to given value.


### GetCreated

`func (o *JobLogEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JobLogEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JobLogEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *JobLogEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


