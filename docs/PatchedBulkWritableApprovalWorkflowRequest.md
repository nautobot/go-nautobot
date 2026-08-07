# PatchedBulkWritableApprovalWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ObjectUnderReviewContentType** | Pointer to **string** |  | [optional] 
**ObjectUnderReviewObjectId** | Pointer to **string** |  | [optional] 
**CurrentState** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | Current state of the approval workflow. Eligible values are: Pending, Approved, Denied, Canceled. | [optional] 
**ApprovalWorkflowDefinition** | Pointer to [**NullableApprovalWorkflowApprovalWorkflowDefinition**](ApprovalWorkflowApprovalWorkflowDefinition.md) |  | [optional] 
**User** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableApprovalWorkflowRequest

`func NewPatchedBulkWritableApprovalWorkflowRequest(id string, ) *PatchedBulkWritableApprovalWorkflowRequest`

NewPatchedBulkWritableApprovalWorkflowRequest instantiates a new PatchedBulkWritableApprovalWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableApprovalWorkflowRequestWithDefaults

`func NewPatchedBulkWritableApprovalWorkflowRequestWithDefaults() *PatchedBulkWritableApprovalWorkflowRequest`

NewPatchedBulkWritableApprovalWorkflowRequestWithDefaults instantiates a new PatchedBulkWritableApprovalWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableApprovalWorkflowRequest) SetId(v string)`

SetId sets Id field to given value.


### GetObjectUnderReviewContentType

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetObjectUnderReviewContentType() string`

GetObjectUnderReviewContentType returns the ObjectUnderReviewContentType field if non-nil, zero value otherwise.

### GetObjectUnderReviewContentTypeOk

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetObjectUnderReviewContentTypeOk() (*string, bool)`

GetObjectUnderReviewContentTypeOk returns a tuple with the ObjectUnderReviewContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUnderReviewContentType

`func (o *PatchedBulkWritableApprovalWorkflowRequest) SetObjectUnderReviewContentType(v string)`

SetObjectUnderReviewContentType sets ObjectUnderReviewContentType field to given value.

### HasObjectUnderReviewContentType

`func (o *PatchedBulkWritableApprovalWorkflowRequest) HasObjectUnderReviewContentType() bool`

HasObjectUnderReviewContentType returns a boolean if a field has been set.

### GetObjectUnderReviewObjectId

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetObjectUnderReviewObjectId() string`

GetObjectUnderReviewObjectId returns the ObjectUnderReviewObjectId field if non-nil, zero value otherwise.

### GetObjectUnderReviewObjectIdOk

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetObjectUnderReviewObjectIdOk() (*string, bool)`

GetObjectUnderReviewObjectIdOk returns a tuple with the ObjectUnderReviewObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUnderReviewObjectId

`func (o *PatchedBulkWritableApprovalWorkflowRequest) SetObjectUnderReviewObjectId(v string)`

SetObjectUnderReviewObjectId sets ObjectUnderReviewObjectId field to given value.

### HasObjectUnderReviewObjectId

`func (o *PatchedBulkWritableApprovalWorkflowRequest) HasObjectUnderReviewObjectId() bool`

HasObjectUnderReviewObjectId returns a boolean if a field has been set.

### GetCurrentState

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetCurrentState() ApprovalWorkflowStateChoices`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetCurrentStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *PatchedBulkWritableApprovalWorkflowRequest) SetCurrentState(v ApprovalWorkflowStateChoices)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *PatchedBulkWritableApprovalWorkflowRequest) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetApprovalWorkflowDefinition

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetApprovalWorkflowDefinition() ApprovalWorkflowApprovalWorkflowDefinition`

GetApprovalWorkflowDefinition returns the ApprovalWorkflowDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowDefinitionOk

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetApprovalWorkflowDefinitionOk() (*ApprovalWorkflowApprovalWorkflowDefinition, bool)`

GetApprovalWorkflowDefinitionOk returns a tuple with the ApprovalWorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowDefinition

`func (o *PatchedBulkWritableApprovalWorkflowRequest) SetApprovalWorkflowDefinition(v ApprovalWorkflowApprovalWorkflowDefinition)`

SetApprovalWorkflowDefinition sets ApprovalWorkflowDefinition field to given value.

### HasApprovalWorkflowDefinition

`func (o *PatchedBulkWritableApprovalWorkflowRequest) HasApprovalWorkflowDefinition() bool`

HasApprovalWorkflowDefinition returns a boolean if a field has been set.

### SetApprovalWorkflowDefinitionNil

`func (o *PatchedBulkWritableApprovalWorkflowRequest) SetApprovalWorkflowDefinitionNil(b bool)`

 SetApprovalWorkflowDefinitionNil sets the value for ApprovalWorkflowDefinition to be an explicit nil

### UnsetApprovalWorkflowDefinition
`func (o *PatchedBulkWritableApprovalWorkflowRequest) UnsetApprovalWorkflowDefinition()`

UnsetApprovalWorkflowDefinition ensures that no value is present for ApprovalWorkflowDefinition, not even an explicit nil
### GetUser

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetUser() ApprovalWorkflowUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetUserOk() (*ApprovalWorkflowUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedBulkWritableApprovalWorkflowRequest) SetUser(v ApprovalWorkflowUser)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedBulkWritableApprovalWorkflowRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedBulkWritableApprovalWorkflowRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedBulkWritableApprovalWorkflowRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableApprovalWorkflowRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableApprovalWorkflowRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableApprovalWorkflowRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableApprovalWorkflowRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableApprovalWorkflowRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


