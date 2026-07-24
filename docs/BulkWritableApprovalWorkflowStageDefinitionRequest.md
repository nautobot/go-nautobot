# BulkWritableApprovalWorkflowStageDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ApproverGroup** | **string** | The group that will be assigned to approve this stage. | 
**Sequence** | **int32** | The sequence dictates the order in which this stage will need to be approved. The lower the number, the earlier it will be. | 
**Name** | **string** |  | 
**MinApprovers** | **int32** | Minimum number of approvers required to approve this stage. | 
**DenialMessage** | Pointer to **string** | Message to show when the stage is denied. | [optional] 
**ApprovalWorkflowDefinition** | [**ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition**](ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableApprovalWorkflowStageDefinitionRequest

`func NewBulkWritableApprovalWorkflowStageDefinitionRequest(id string, approverGroup string, sequence int32, name string, minApprovers int32, approvalWorkflowDefinition ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition, ) *BulkWritableApprovalWorkflowStageDefinitionRequest`

NewBulkWritableApprovalWorkflowStageDefinitionRequest instantiates a new BulkWritableApprovalWorkflowStageDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableApprovalWorkflowStageDefinitionRequestWithDefaults

`func NewBulkWritableApprovalWorkflowStageDefinitionRequestWithDefaults() *BulkWritableApprovalWorkflowStageDefinitionRequest`

NewBulkWritableApprovalWorkflowStageDefinitionRequestWithDefaults instantiates a new BulkWritableApprovalWorkflowStageDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetApproverGroup

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetApproverGroup() string`

GetApproverGroup returns the ApproverGroup field if non-nil, zero value otherwise.

### GetApproverGroupOk

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetApproverGroupOk() (*string, bool)`

GetApproverGroupOk returns a tuple with the ApproverGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverGroup

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) SetApproverGroup(v string)`

SetApproverGroup sets ApproverGroup field to given value.


### GetSequence

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetName

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMinApprovers

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetMinApprovers() int32`

GetMinApprovers returns the MinApprovers field if non-nil, zero value otherwise.

### GetMinApproversOk

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetMinApproversOk() (*int32, bool)`

GetMinApproversOk returns a tuple with the MinApprovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinApprovers

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) SetMinApprovers(v int32)`

SetMinApprovers sets MinApprovers field to given value.


### GetDenialMessage

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetDenialMessage() string`

GetDenialMessage returns the DenialMessage field if non-nil, zero value otherwise.

### GetDenialMessageOk

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetDenialMessageOk() (*string, bool)`

GetDenialMessageOk returns a tuple with the DenialMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialMessage

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) SetDenialMessage(v string)`

SetDenialMessage sets DenialMessage field to given value.

### HasDenialMessage

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) HasDenialMessage() bool`

HasDenialMessage returns a boolean if a field has been set.

### GetApprovalWorkflowDefinition

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetApprovalWorkflowDefinition() ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition`

GetApprovalWorkflowDefinition returns the ApprovalWorkflowDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowDefinitionOk

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetApprovalWorkflowDefinitionOk() (*ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition, bool)`

GetApprovalWorkflowDefinitionOk returns a tuple with the ApprovalWorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowDefinition

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) SetApprovalWorkflowDefinition(v ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition)`

SetApprovalWorkflowDefinition sets ApprovalWorkflowDefinition field to given value.


### GetCustomFields

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableApprovalWorkflowStageDefinitionRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


