# BulkWritableApprovalWorkflowStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | State of the approval workflow stage instance. Eligible values are: Pending, Approved, Denied. | [optional] 
**ApprovalWorkflow** | [**ApprovalWorkflowStageApprovalWorkflow**](ApprovalWorkflowStageApprovalWorkflow.md) |  | 
**ApprovalWorkflowStageDefinition** | Pointer to [**NullableApprovalWorkflowStageApprovalWorkflowStageDefinition**](ApprovalWorkflowStageApprovalWorkflowStageDefinition.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableApprovalWorkflowStageRequest

`func NewBulkWritableApprovalWorkflowStageRequest(id string, approvalWorkflow ApprovalWorkflowStageApprovalWorkflow, ) *BulkWritableApprovalWorkflowStageRequest`

NewBulkWritableApprovalWorkflowStageRequest instantiates a new BulkWritableApprovalWorkflowStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableApprovalWorkflowStageRequestWithDefaults

`func NewBulkWritableApprovalWorkflowStageRequestWithDefaults() *BulkWritableApprovalWorkflowStageRequest`

NewBulkWritableApprovalWorkflowStageRequestWithDefaults instantiates a new BulkWritableApprovalWorkflowStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableApprovalWorkflowStageRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableApprovalWorkflowStageRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableApprovalWorkflowStageRequest) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *BulkWritableApprovalWorkflowStageRequest) GetState() ApprovalWorkflowStateChoices`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkWritableApprovalWorkflowStageRequest) GetStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkWritableApprovalWorkflowStageRequest) SetState(v ApprovalWorkflowStateChoices)`

SetState sets State field to given value.

### HasState

`func (o *BulkWritableApprovalWorkflowStageRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApprovalWorkflow

`func (o *BulkWritableApprovalWorkflowStageRequest) GetApprovalWorkflow() ApprovalWorkflowStageApprovalWorkflow`

GetApprovalWorkflow returns the ApprovalWorkflow field if non-nil, zero value otherwise.

### GetApprovalWorkflowOk

`func (o *BulkWritableApprovalWorkflowStageRequest) GetApprovalWorkflowOk() (*ApprovalWorkflowStageApprovalWorkflow, bool)`

GetApprovalWorkflowOk returns a tuple with the ApprovalWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflow

`func (o *BulkWritableApprovalWorkflowStageRequest) SetApprovalWorkflow(v ApprovalWorkflowStageApprovalWorkflow)`

SetApprovalWorkflow sets ApprovalWorkflow field to given value.


### GetApprovalWorkflowStageDefinition

`func (o *BulkWritableApprovalWorkflowStageRequest) GetApprovalWorkflowStageDefinition() ApprovalWorkflowStageApprovalWorkflowStageDefinition`

GetApprovalWorkflowStageDefinition returns the ApprovalWorkflowStageDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowStageDefinitionOk

`func (o *BulkWritableApprovalWorkflowStageRequest) GetApprovalWorkflowStageDefinitionOk() (*ApprovalWorkflowStageApprovalWorkflowStageDefinition, bool)`

GetApprovalWorkflowStageDefinitionOk returns a tuple with the ApprovalWorkflowStageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowStageDefinition

`func (o *BulkWritableApprovalWorkflowStageRequest) SetApprovalWorkflowStageDefinition(v ApprovalWorkflowStageApprovalWorkflowStageDefinition)`

SetApprovalWorkflowStageDefinition sets ApprovalWorkflowStageDefinition field to given value.

### HasApprovalWorkflowStageDefinition

`func (o *BulkWritableApprovalWorkflowStageRequest) HasApprovalWorkflowStageDefinition() bool`

HasApprovalWorkflowStageDefinition returns a boolean if a field has been set.

### SetApprovalWorkflowStageDefinitionNil

`func (o *BulkWritableApprovalWorkflowStageRequest) SetApprovalWorkflowStageDefinitionNil(b bool)`

 SetApprovalWorkflowStageDefinitionNil sets the value for ApprovalWorkflowStageDefinition to be an explicit nil

### UnsetApprovalWorkflowStageDefinition
`func (o *BulkWritableApprovalWorkflowStageRequest) UnsetApprovalWorkflowStageDefinition()`

UnsetApprovalWorkflowStageDefinition ensures that no value is present for ApprovalWorkflowStageDefinition, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableApprovalWorkflowStageRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableApprovalWorkflowStageRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableApprovalWorkflowStageRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableApprovalWorkflowStageRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableApprovalWorkflowStageRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableApprovalWorkflowStageRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableApprovalWorkflowStageRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableApprovalWorkflowStageRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


