# PatchedApprovalWorkflowStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | State of the approval workflow stage instance. Eligible values are: Pending, Approved, Denied. | [optional] 
**ApprovalWorkflow** | Pointer to [**ApprovalWorkflowStageApprovalWorkflow**](ApprovalWorkflowStageApprovalWorkflow.md) |  | [optional] 
**ApprovalWorkflowStageDefinition** | Pointer to [**NullableApprovalWorkflowStageApprovalWorkflowStageDefinition**](ApprovalWorkflowStageApprovalWorkflowStageDefinition.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedApprovalWorkflowStageRequest

`func NewPatchedApprovalWorkflowStageRequest() *PatchedApprovalWorkflowStageRequest`

NewPatchedApprovalWorkflowStageRequest instantiates a new PatchedApprovalWorkflowStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedApprovalWorkflowStageRequestWithDefaults

`func NewPatchedApprovalWorkflowStageRequestWithDefaults() *PatchedApprovalWorkflowStageRequest`

NewPatchedApprovalWorkflowStageRequestWithDefaults instantiates a new PatchedApprovalWorkflowStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedApprovalWorkflowStageRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedApprovalWorkflowStageRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedApprovalWorkflowStageRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedApprovalWorkflowStageRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *PatchedApprovalWorkflowStageRequest) GetState() ApprovalWorkflowStateChoices`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatchedApprovalWorkflowStageRequest) GetStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatchedApprovalWorkflowStageRequest) SetState(v ApprovalWorkflowStateChoices)`

SetState sets State field to given value.

### HasState

`func (o *PatchedApprovalWorkflowStageRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApprovalWorkflow

`func (o *PatchedApprovalWorkflowStageRequest) GetApprovalWorkflow() ApprovalWorkflowStageApprovalWorkflow`

GetApprovalWorkflow returns the ApprovalWorkflow field if non-nil, zero value otherwise.

### GetApprovalWorkflowOk

`func (o *PatchedApprovalWorkflowStageRequest) GetApprovalWorkflowOk() (*ApprovalWorkflowStageApprovalWorkflow, bool)`

GetApprovalWorkflowOk returns a tuple with the ApprovalWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflow

`func (o *PatchedApprovalWorkflowStageRequest) SetApprovalWorkflow(v ApprovalWorkflowStageApprovalWorkflow)`

SetApprovalWorkflow sets ApprovalWorkflow field to given value.

### HasApprovalWorkflow

`func (o *PatchedApprovalWorkflowStageRequest) HasApprovalWorkflow() bool`

HasApprovalWorkflow returns a boolean if a field has been set.

### GetApprovalWorkflowStageDefinition

`func (o *PatchedApprovalWorkflowStageRequest) GetApprovalWorkflowStageDefinition() ApprovalWorkflowStageApprovalWorkflowStageDefinition`

GetApprovalWorkflowStageDefinition returns the ApprovalWorkflowStageDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowStageDefinitionOk

`func (o *PatchedApprovalWorkflowStageRequest) GetApprovalWorkflowStageDefinitionOk() (*ApprovalWorkflowStageApprovalWorkflowStageDefinition, bool)`

GetApprovalWorkflowStageDefinitionOk returns a tuple with the ApprovalWorkflowStageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowStageDefinition

`func (o *PatchedApprovalWorkflowStageRequest) SetApprovalWorkflowStageDefinition(v ApprovalWorkflowStageApprovalWorkflowStageDefinition)`

SetApprovalWorkflowStageDefinition sets ApprovalWorkflowStageDefinition field to given value.

### HasApprovalWorkflowStageDefinition

`func (o *PatchedApprovalWorkflowStageRequest) HasApprovalWorkflowStageDefinition() bool`

HasApprovalWorkflowStageDefinition returns a boolean if a field has been set.

### SetApprovalWorkflowStageDefinitionNil

`func (o *PatchedApprovalWorkflowStageRequest) SetApprovalWorkflowStageDefinitionNil(b bool)`

 SetApprovalWorkflowStageDefinitionNil sets the value for ApprovalWorkflowStageDefinition to be an explicit nil

### UnsetApprovalWorkflowStageDefinition
`func (o *PatchedApprovalWorkflowStageRequest) UnsetApprovalWorkflowStageDefinition()`

UnsetApprovalWorkflowStageDefinition ensures that no value is present for ApprovalWorkflowStageDefinition, not even an explicit nil
### GetCustomFields

`func (o *PatchedApprovalWorkflowStageRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedApprovalWorkflowStageRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedApprovalWorkflowStageRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedApprovalWorkflowStageRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedApprovalWorkflowStageRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedApprovalWorkflowStageRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedApprovalWorkflowStageRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedApprovalWorkflowStageRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


