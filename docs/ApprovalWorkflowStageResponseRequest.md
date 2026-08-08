# ApprovalWorkflowStageResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** | User comments to explain the decision that he/she made | [optional] 
**State** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | User response to this approval workflow stage instance. Eligible values are: Pending, Comment, Approved, Denied. | [optional] 
**ApprovalWorkflowStage** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**User** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewApprovalWorkflowStageResponseRequest

`func NewApprovalWorkflowStageResponseRequest(approvalWorkflowStage ApprovalWorkflowStageResponseApprovalWorkflowStage, user ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *ApprovalWorkflowStageResponseRequest`

NewApprovalWorkflowStageResponseRequest instantiates a new ApprovalWorkflowStageResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWorkflowStageResponseRequestWithDefaults

`func NewApprovalWorkflowStageResponseRequestWithDefaults() *ApprovalWorkflowStageResponseRequest`

NewApprovalWorkflowStageResponseRequestWithDefaults instantiates a new ApprovalWorkflowStageResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalWorkflowStageResponseRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalWorkflowStageResponseRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalWorkflowStageResponseRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalWorkflowStageResponseRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComments

`func (o *ApprovalWorkflowStageResponseRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ApprovalWorkflowStageResponseRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ApprovalWorkflowStageResponseRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ApprovalWorkflowStageResponseRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetState

`func (o *ApprovalWorkflowStageResponseRequest) GetState() ApprovalWorkflowStateChoices`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApprovalWorkflowStageResponseRequest) GetStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApprovalWorkflowStageResponseRequest) SetState(v ApprovalWorkflowStateChoices)`

SetState sets State field to given value.

### HasState

`func (o *ApprovalWorkflowStageResponseRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApprovalWorkflowStage

`func (o *ApprovalWorkflowStageResponseRequest) GetApprovalWorkflowStage() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetApprovalWorkflowStage returns the ApprovalWorkflowStage field if non-nil, zero value otherwise.

### GetApprovalWorkflowStageOk

`func (o *ApprovalWorkflowStageResponseRequest) GetApprovalWorkflowStageOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetApprovalWorkflowStageOk returns a tuple with the ApprovalWorkflowStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowStage

`func (o *ApprovalWorkflowStageResponseRequest) SetApprovalWorkflowStage(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetApprovalWorkflowStage sets ApprovalWorkflowStage field to given value.


### GetUser

`func (o *ApprovalWorkflowStageResponseRequest) GetUser() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApprovalWorkflowStageResponseRequest) GetUserOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApprovalWorkflowStageResponseRequest) SetUser(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


