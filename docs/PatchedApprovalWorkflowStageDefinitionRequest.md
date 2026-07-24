# PatchedApprovalWorkflowStageDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ApproverGroup** | Pointer to **string** | The group that will be assigned to approve this stage. | [optional] 
**Sequence** | Pointer to **int32** | The sequence dictates the order in which this stage will need to be approved. The lower the number, the earlier it will be. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MinApprovers** | Pointer to **int32** | Minimum number of approvers required to approve this stage. | [optional] 
**DenialMessage** | Pointer to **string** | Message to show when the stage is denied. | [optional] 
**ApprovalWorkflowDefinition** | Pointer to [**ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition**](ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedApprovalWorkflowStageDefinitionRequest

`func NewPatchedApprovalWorkflowStageDefinitionRequest() *PatchedApprovalWorkflowStageDefinitionRequest`

NewPatchedApprovalWorkflowStageDefinitionRequest instantiates a new PatchedApprovalWorkflowStageDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedApprovalWorkflowStageDefinitionRequestWithDefaults

`func NewPatchedApprovalWorkflowStageDefinitionRequestWithDefaults() *PatchedApprovalWorkflowStageDefinitionRequest`

NewPatchedApprovalWorkflowStageDefinitionRequestWithDefaults instantiates a new PatchedApprovalWorkflowStageDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApproverGroup

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetApproverGroup() string`

GetApproverGroup returns the ApproverGroup field if non-nil, zero value otherwise.

### GetApproverGroupOk

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetApproverGroupOk() (*string, bool)`

GetApproverGroupOk returns a tuple with the ApproverGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverGroup

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) SetApproverGroup(v string)`

SetApproverGroup sets ApproverGroup field to given value.

### HasApproverGroup

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) HasApproverGroup() bool`

HasApproverGroup returns a boolean if a field has been set.

### GetSequence

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetName

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMinApprovers

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetMinApprovers() int32`

GetMinApprovers returns the MinApprovers field if non-nil, zero value otherwise.

### GetMinApproversOk

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetMinApproversOk() (*int32, bool)`

GetMinApproversOk returns a tuple with the MinApprovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinApprovers

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) SetMinApprovers(v int32)`

SetMinApprovers sets MinApprovers field to given value.

### HasMinApprovers

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) HasMinApprovers() bool`

HasMinApprovers returns a boolean if a field has been set.

### GetDenialMessage

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetDenialMessage() string`

GetDenialMessage returns the DenialMessage field if non-nil, zero value otherwise.

### GetDenialMessageOk

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetDenialMessageOk() (*string, bool)`

GetDenialMessageOk returns a tuple with the DenialMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialMessage

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) SetDenialMessage(v string)`

SetDenialMessage sets DenialMessage field to given value.

### HasDenialMessage

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) HasDenialMessage() bool`

HasDenialMessage returns a boolean if a field has been set.

### GetApprovalWorkflowDefinition

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetApprovalWorkflowDefinition() ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition`

GetApprovalWorkflowDefinition returns the ApprovalWorkflowDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowDefinitionOk

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetApprovalWorkflowDefinitionOk() (*ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition, bool)`

GetApprovalWorkflowDefinitionOk returns a tuple with the ApprovalWorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowDefinition

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) SetApprovalWorkflowDefinition(v ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition)`

SetApprovalWorkflowDefinition sets ApprovalWorkflowDefinition field to given value.

### HasApprovalWorkflowDefinition

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) HasApprovalWorkflowDefinition() bool`

HasApprovalWorkflowDefinition returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedApprovalWorkflowStageDefinitionRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


