# GitRepositorySyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | [readonly] 
**JobResult** | [**JobResult**](JobResult.md) |  | [readonly] 

## Methods

### NewGitRepositorySyncResponse

`func NewGitRepositorySyncResponse(message string, jobResult JobResult, ) *GitRepositorySyncResponse`

NewGitRepositorySyncResponse instantiates a new GitRepositorySyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRepositorySyncResponseWithDefaults

`func NewGitRepositorySyncResponseWithDefaults() *GitRepositorySyncResponse`

NewGitRepositorySyncResponseWithDefaults instantiates a new GitRepositorySyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GitRepositorySyncResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitRepositorySyncResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitRepositorySyncResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetJobResult

`func (o *GitRepositorySyncResponse) GetJobResult() JobResult`

GetJobResult returns the JobResult field if non-nil, zero value otherwise.

### GetJobResultOk

`func (o *GitRepositorySyncResponse) GetJobResultOk() (*JobResult, bool)`

GetJobResultOk returns a tuple with the JobResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResult

`func (o *GitRepositorySyncResponse) SetJobResult(v JobResult)`

SetJobResult sets JobResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


