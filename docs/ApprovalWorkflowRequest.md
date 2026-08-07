# ApprovalWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectUnderReviewContentType** | **string** |  | 
**ObjectUnderReviewObjectId** | **string** |  | 
**CurrentState** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | Current state of the approval workflow. Eligible values are: Pending, Approved, Denied, Canceled. | [optional] 
**ApprovalWorkflowDefinition** | Pointer to [**NullableApprovalWorkflowApprovalWorkflowDefinition**](ApprovalWorkflowApprovalWorkflowDefinition.md) |  | [optional] 
**User** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewApprovalWorkflowRequest

`func NewApprovalWorkflowRequest(objectUnderReviewContentType string, objectUnderReviewObjectId string, ) *ApprovalWorkflowRequest`

NewApprovalWorkflowRequest instantiates a new ApprovalWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWorkflowRequestWithDefaults

`func NewApprovalWorkflowRequestWithDefaults() *ApprovalWorkflowRequest`

NewApprovalWorkflowRequestWithDefaults instantiates a new ApprovalWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalWorkflowRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalWorkflowRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalWorkflowRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalWorkflowRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectUnderReviewContentType

`func (o *ApprovalWorkflowRequest) GetObjectUnderReviewContentType() string`

GetObjectUnderReviewContentType returns the ObjectUnderReviewContentType field if non-nil, zero value otherwise.

### GetObjectUnderReviewContentTypeOk

`func (o *ApprovalWorkflowRequest) GetObjectUnderReviewContentTypeOk() (*string, bool)`

GetObjectUnderReviewContentTypeOk returns a tuple with the ObjectUnderReviewContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUnderReviewContentType

`func (o *ApprovalWorkflowRequest) SetObjectUnderReviewContentType(v string)`

SetObjectUnderReviewContentType sets ObjectUnderReviewContentType field to given value.


### GetObjectUnderReviewObjectId

`func (o *ApprovalWorkflowRequest) GetObjectUnderReviewObjectId() string`

GetObjectUnderReviewObjectId returns the ObjectUnderReviewObjectId field if non-nil, zero value otherwise.

### GetObjectUnderReviewObjectIdOk

`func (o *ApprovalWorkflowRequest) GetObjectUnderReviewObjectIdOk() (*string, bool)`

GetObjectUnderReviewObjectIdOk returns a tuple with the ObjectUnderReviewObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUnderReviewObjectId

`func (o *ApprovalWorkflowRequest) SetObjectUnderReviewObjectId(v string)`

SetObjectUnderReviewObjectId sets ObjectUnderReviewObjectId field to given value.


### GetCurrentState

`func (o *ApprovalWorkflowRequest) GetCurrentState() ApprovalWorkflowStateChoices`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *ApprovalWorkflowRequest) GetCurrentStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *ApprovalWorkflowRequest) SetCurrentState(v ApprovalWorkflowStateChoices)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *ApprovalWorkflowRequest) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetApprovalWorkflowDefinition

`func (o *ApprovalWorkflowRequest) GetApprovalWorkflowDefinition() ApprovalWorkflowApprovalWorkflowDefinition`

GetApprovalWorkflowDefinition returns the ApprovalWorkflowDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowDefinitionOk

`func (o *ApprovalWorkflowRequest) GetApprovalWorkflowDefinitionOk() (*ApprovalWorkflowApprovalWorkflowDefinition, bool)`

GetApprovalWorkflowDefinitionOk returns a tuple with the ApprovalWorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowDefinition

`func (o *ApprovalWorkflowRequest) SetApprovalWorkflowDefinition(v ApprovalWorkflowApprovalWorkflowDefinition)`

SetApprovalWorkflowDefinition sets ApprovalWorkflowDefinition field to given value.

### HasApprovalWorkflowDefinition

`func (o *ApprovalWorkflowRequest) HasApprovalWorkflowDefinition() bool`

HasApprovalWorkflowDefinition returns a boolean if a field has been set.

### SetApprovalWorkflowDefinitionNil

`func (o *ApprovalWorkflowRequest) SetApprovalWorkflowDefinitionNil(b bool)`

 SetApprovalWorkflowDefinitionNil sets the value for ApprovalWorkflowDefinition to be an explicit nil

### UnsetApprovalWorkflowDefinition
`func (o *ApprovalWorkflowRequest) UnsetApprovalWorkflowDefinition()`

UnsetApprovalWorkflowDefinition ensures that no value is present for ApprovalWorkflowDefinition, not even an explicit nil
### GetUser

`func (o *ApprovalWorkflowRequest) GetUser() ApprovalWorkflowUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApprovalWorkflowRequest) GetUserOk() (*ApprovalWorkflowUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApprovalWorkflowRequest) SetUser(v ApprovalWorkflowUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ApprovalWorkflowRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ApprovalWorkflowRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ApprovalWorkflowRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetCustomFields

`func (o *ApprovalWorkflowRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ApprovalWorkflowRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ApprovalWorkflowRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ApprovalWorkflowRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ApprovalWorkflowRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ApprovalWorkflowRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ApprovalWorkflowRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ApprovalWorkflowRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


