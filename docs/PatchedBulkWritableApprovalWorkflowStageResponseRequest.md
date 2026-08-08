# PatchedBulkWritableApprovalWorkflowStageResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Comments** | Pointer to **string** | User comments to explain the decision that he/she made | [optional] 
**State** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | User response to this approval workflow stage instance. Eligible values are: Pending, Comment, Approved, Denied. | [optional] 
**ApprovalWorkflowStage** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**User** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableApprovalWorkflowStageResponseRequest

`func NewPatchedBulkWritableApprovalWorkflowStageResponseRequest(id string, ) *PatchedBulkWritableApprovalWorkflowStageResponseRequest`

NewPatchedBulkWritableApprovalWorkflowStageResponseRequest instantiates a new PatchedBulkWritableApprovalWorkflowStageResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableApprovalWorkflowStageResponseRequestWithDefaults

`func NewPatchedBulkWritableApprovalWorkflowStageResponseRequestWithDefaults() *PatchedBulkWritableApprovalWorkflowStageResponseRequest`

NewPatchedBulkWritableApprovalWorkflowStageResponseRequestWithDefaults instantiates a new PatchedBulkWritableApprovalWorkflowStageResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) SetId(v string)`

SetId sets Id field to given value.


### GetComments

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetState

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) GetState() ApprovalWorkflowStateChoices`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) GetStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) SetState(v ApprovalWorkflowStateChoices)`

SetState sets State field to given value.

### HasState

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApprovalWorkflowStage

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) GetApprovalWorkflowStage() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetApprovalWorkflowStage returns the ApprovalWorkflowStage field if non-nil, zero value otherwise.

### GetApprovalWorkflowStageOk

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) GetApprovalWorkflowStageOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetApprovalWorkflowStageOk returns a tuple with the ApprovalWorkflowStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowStage

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) SetApprovalWorkflowStage(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetApprovalWorkflowStage sets ApprovalWorkflowStage field to given value.

### HasApprovalWorkflowStage

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) HasApprovalWorkflowStage() bool`

HasApprovalWorkflowStage returns a boolean if a field has been set.

### GetUser

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) GetUser() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) GetUserOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) SetUser(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedBulkWritableApprovalWorkflowStageResponseRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


