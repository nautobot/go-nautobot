# BulkWritableApprovalWorkflowStageResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Comments** | Pointer to **string** | User comments to explain the decision that he/she made | [optional] 
**State** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | User response to this approval workflow stage instance. Eligible values are: Pending, Comment, Approved, Denied. | [optional] 
**ApprovalWorkflowStage** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**User** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewBulkWritableApprovalWorkflowStageResponseRequest

`func NewBulkWritableApprovalWorkflowStageResponseRequest(id string, approvalWorkflowStage ApprovalWorkflowStageResponseApprovalWorkflowStage, user ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *BulkWritableApprovalWorkflowStageResponseRequest`

NewBulkWritableApprovalWorkflowStageResponseRequest instantiates a new BulkWritableApprovalWorkflowStageResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableApprovalWorkflowStageResponseRequestWithDefaults

`func NewBulkWritableApprovalWorkflowStageResponseRequestWithDefaults() *BulkWritableApprovalWorkflowStageResponseRequest`

NewBulkWritableApprovalWorkflowStageResponseRequestWithDefaults instantiates a new BulkWritableApprovalWorkflowStageResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) SetId(v string)`

SetId sets Id field to given value.


### GetComments

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetState

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) GetState() ApprovalWorkflowStateChoices`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) GetStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) SetState(v ApprovalWorkflowStateChoices)`

SetState sets State field to given value.

### HasState

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApprovalWorkflowStage

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) GetApprovalWorkflowStage() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetApprovalWorkflowStage returns the ApprovalWorkflowStage field if non-nil, zero value otherwise.

### GetApprovalWorkflowStageOk

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) GetApprovalWorkflowStageOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetApprovalWorkflowStageOk returns a tuple with the ApprovalWorkflowStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowStage

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) SetApprovalWorkflowStage(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetApprovalWorkflowStage sets ApprovalWorkflowStage field to given value.


### GetUser

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) GetUser() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) GetUserOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BulkWritableApprovalWorkflowStageResponseRequest) SetUser(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


