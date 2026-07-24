# ApprovalWorkflowStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | State of the approval workflow stage instance. Eligible values are: Pending, Approved, Denied. | [optional] 
**ApprovalWorkflow** | [**ApprovalWorkflowStageApprovalWorkflow**](ApprovalWorkflowStageApprovalWorkflow.md) |  | 
**ApprovalWorkflowStageDefinition** | [**ApprovalWorkflowStageApprovalWorkflowStageDefinition**](ApprovalWorkflowStageApprovalWorkflowStageDefinition.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewApprovalWorkflowStageRequest

`func NewApprovalWorkflowStageRequest(approvalWorkflow ApprovalWorkflowStageApprovalWorkflow, approvalWorkflowStageDefinition ApprovalWorkflowStageApprovalWorkflowStageDefinition, ) *ApprovalWorkflowStageRequest`

NewApprovalWorkflowStageRequest instantiates a new ApprovalWorkflowStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWorkflowStageRequestWithDefaults

`func NewApprovalWorkflowStageRequestWithDefaults() *ApprovalWorkflowStageRequest`

NewApprovalWorkflowStageRequestWithDefaults instantiates a new ApprovalWorkflowStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalWorkflowStageRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalWorkflowStageRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalWorkflowStageRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalWorkflowStageRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *ApprovalWorkflowStageRequest) GetState() ApprovalWorkflowStateChoices`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApprovalWorkflowStageRequest) GetStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApprovalWorkflowStageRequest) SetState(v ApprovalWorkflowStateChoices)`

SetState sets State field to given value.

### HasState

`func (o *ApprovalWorkflowStageRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApprovalWorkflow

`func (o *ApprovalWorkflowStageRequest) GetApprovalWorkflow() ApprovalWorkflowStageApprovalWorkflow`

GetApprovalWorkflow returns the ApprovalWorkflow field if non-nil, zero value otherwise.

### GetApprovalWorkflowOk

`func (o *ApprovalWorkflowStageRequest) GetApprovalWorkflowOk() (*ApprovalWorkflowStageApprovalWorkflow, bool)`

GetApprovalWorkflowOk returns a tuple with the ApprovalWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflow

`func (o *ApprovalWorkflowStageRequest) SetApprovalWorkflow(v ApprovalWorkflowStageApprovalWorkflow)`

SetApprovalWorkflow sets ApprovalWorkflow field to given value.


### GetApprovalWorkflowStageDefinition

`func (o *ApprovalWorkflowStageRequest) GetApprovalWorkflowStageDefinition() ApprovalWorkflowStageApprovalWorkflowStageDefinition`

GetApprovalWorkflowStageDefinition returns the ApprovalWorkflowStageDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowStageDefinitionOk

`func (o *ApprovalWorkflowStageRequest) GetApprovalWorkflowStageDefinitionOk() (*ApprovalWorkflowStageApprovalWorkflowStageDefinition, bool)`

GetApprovalWorkflowStageDefinitionOk returns a tuple with the ApprovalWorkflowStageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowStageDefinition

`func (o *ApprovalWorkflowStageRequest) SetApprovalWorkflowStageDefinition(v ApprovalWorkflowStageApprovalWorkflowStageDefinition)`

SetApprovalWorkflowStageDefinition sets ApprovalWorkflowStageDefinition field to given value.


### GetCustomFields

`func (o *ApprovalWorkflowStageRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ApprovalWorkflowStageRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ApprovalWorkflowStageRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ApprovalWorkflowStageRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ApprovalWorkflowStageRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ApprovalWorkflowStageRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ApprovalWorkflowStageRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ApprovalWorkflowStageRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


