# PatchedBulkWritableApprovalWorkflowStageDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ApproverGroup** | Pointer to **string** | The group that will be assigned to approve this stage. | [optional] 
**Sequence** | Pointer to **int32** | The sequence dictates the order in which this stage will need to be approved. The lower the number, the earlier it will be. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MinApprovers** | Pointer to **int32** | Minimum number of approvers required to approve this stage. | [optional] 
**DenialMessage** | Pointer to **string** | Message to show when the stage is denied. | [optional] 
**ApprovalWorkflowDefinition** | Pointer to [**ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition**](ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableApprovalWorkflowStageDefinitionRequest

`func NewPatchedBulkWritableApprovalWorkflowStageDefinitionRequest(id string, ) *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest`

NewPatchedBulkWritableApprovalWorkflowStageDefinitionRequest instantiates a new PatchedBulkWritableApprovalWorkflowStageDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableApprovalWorkflowStageDefinitionRequestWithDefaults

`func NewPatchedBulkWritableApprovalWorkflowStageDefinitionRequestWithDefaults() *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest`

NewPatchedBulkWritableApprovalWorkflowStageDefinitionRequestWithDefaults instantiates a new PatchedBulkWritableApprovalWorkflowStageDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetApproverGroup

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetApproverGroup() string`

GetApproverGroup returns the ApproverGroup field if non-nil, zero value otherwise.

### GetApproverGroupOk

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetApproverGroupOk() (*string, bool)`

GetApproverGroupOk returns a tuple with the ApproverGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverGroup

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) SetApproverGroup(v string)`

SetApproverGroup sets ApproverGroup field to given value.

### HasApproverGroup

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) HasApproverGroup() bool`

HasApproverGroup returns a boolean if a field has been set.

### GetSequence

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMinApprovers

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetMinApprovers() int32`

GetMinApprovers returns the MinApprovers field if non-nil, zero value otherwise.

### GetMinApproversOk

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetMinApproversOk() (*int32, bool)`

GetMinApproversOk returns a tuple with the MinApprovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinApprovers

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) SetMinApprovers(v int32)`

SetMinApprovers sets MinApprovers field to given value.

### HasMinApprovers

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) HasMinApprovers() bool`

HasMinApprovers returns a boolean if a field has been set.

### GetDenialMessage

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetDenialMessage() string`

GetDenialMessage returns the DenialMessage field if non-nil, zero value otherwise.

### GetDenialMessageOk

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetDenialMessageOk() (*string, bool)`

GetDenialMessageOk returns a tuple with the DenialMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialMessage

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) SetDenialMessage(v string)`

SetDenialMessage sets DenialMessage field to given value.

### HasDenialMessage

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) HasDenialMessage() bool`

HasDenialMessage returns a boolean if a field has been set.

### GetApprovalWorkflowDefinition

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetApprovalWorkflowDefinition() ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition`

GetApprovalWorkflowDefinition returns the ApprovalWorkflowDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowDefinitionOk

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetApprovalWorkflowDefinitionOk() (*ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition, bool)`

GetApprovalWorkflowDefinitionOk returns a tuple with the ApprovalWorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowDefinition

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) SetApprovalWorkflowDefinition(v ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition)`

SetApprovalWorkflowDefinition sets ApprovalWorkflowDefinition field to given value.

### HasApprovalWorkflowDefinition

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) HasApprovalWorkflowDefinition() bool`

HasApprovalWorkflowDefinition returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableApprovalWorkflowStageDefinitionRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


