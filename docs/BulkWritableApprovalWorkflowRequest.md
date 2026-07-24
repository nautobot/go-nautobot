# BulkWritableApprovalWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ObjectUnderReviewContentType** | **string** |  | 
**ObjectUnderReviewObjectId** | **string** |  | 
**CurrentState** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | Current state of the approval workflow. Eligible values are: Pending, Approved, Denied, Canceled. | [optional] 
**ApprovalWorkflowDefinition** | [**ApprovalWorkflowApprovalWorkflowDefinition**](ApprovalWorkflowApprovalWorkflowDefinition.md) |  | 
**User** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableApprovalWorkflowRequest

`func NewBulkWritableApprovalWorkflowRequest(id string, objectUnderReviewContentType string, objectUnderReviewObjectId string, approvalWorkflowDefinition ApprovalWorkflowApprovalWorkflowDefinition, ) *BulkWritableApprovalWorkflowRequest`

NewBulkWritableApprovalWorkflowRequest instantiates a new BulkWritableApprovalWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableApprovalWorkflowRequestWithDefaults

`func NewBulkWritableApprovalWorkflowRequestWithDefaults() *BulkWritableApprovalWorkflowRequest`

NewBulkWritableApprovalWorkflowRequestWithDefaults instantiates a new BulkWritableApprovalWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableApprovalWorkflowRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableApprovalWorkflowRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableApprovalWorkflowRequest) SetId(v string)`

SetId sets Id field to given value.


### GetObjectUnderReviewContentType

`func (o *BulkWritableApprovalWorkflowRequest) GetObjectUnderReviewContentType() string`

GetObjectUnderReviewContentType returns the ObjectUnderReviewContentType field if non-nil, zero value otherwise.

### GetObjectUnderReviewContentTypeOk

`func (o *BulkWritableApprovalWorkflowRequest) GetObjectUnderReviewContentTypeOk() (*string, bool)`

GetObjectUnderReviewContentTypeOk returns a tuple with the ObjectUnderReviewContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUnderReviewContentType

`func (o *BulkWritableApprovalWorkflowRequest) SetObjectUnderReviewContentType(v string)`

SetObjectUnderReviewContentType sets ObjectUnderReviewContentType field to given value.


### GetObjectUnderReviewObjectId

`func (o *BulkWritableApprovalWorkflowRequest) GetObjectUnderReviewObjectId() string`

GetObjectUnderReviewObjectId returns the ObjectUnderReviewObjectId field if non-nil, zero value otherwise.

### GetObjectUnderReviewObjectIdOk

`func (o *BulkWritableApprovalWorkflowRequest) GetObjectUnderReviewObjectIdOk() (*string, bool)`

GetObjectUnderReviewObjectIdOk returns a tuple with the ObjectUnderReviewObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUnderReviewObjectId

`func (o *BulkWritableApprovalWorkflowRequest) SetObjectUnderReviewObjectId(v string)`

SetObjectUnderReviewObjectId sets ObjectUnderReviewObjectId field to given value.


### GetCurrentState

`func (o *BulkWritableApprovalWorkflowRequest) GetCurrentState() ApprovalWorkflowStateChoices`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *BulkWritableApprovalWorkflowRequest) GetCurrentStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *BulkWritableApprovalWorkflowRequest) SetCurrentState(v ApprovalWorkflowStateChoices)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *BulkWritableApprovalWorkflowRequest) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetApprovalWorkflowDefinition

`func (o *BulkWritableApprovalWorkflowRequest) GetApprovalWorkflowDefinition() ApprovalWorkflowApprovalWorkflowDefinition`

GetApprovalWorkflowDefinition returns the ApprovalWorkflowDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowDefinitionOk

`func (o *BulkWritableApprovalWorkflowRequest) GetApprovalWorkflowDefinitionOk() (*ApprovalWorkflowApprovalWorkflowDefinition, bool)`

GetApprovalWorkflowDefinitionOk returns a tuple with the ApprovalWorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowDefinition

`func (o *BulkWritableApprovalWorkflowRequest) SetApprovalWorkflowDefinition(v ApprovalWorkflowApprovalWorkflowDefinition)`

SetApprovalWorkflowDefinition sets ApprovalWorkflowDefinition field to given value.


### GetUser

`func (o *BulkWritableApprovalWorkflowRequest) GetUser() ApprovalWorkflowUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BulkWritableApprovalWorkflowRequest) GetUserOk() (*ApprovalWorkflowUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BulkWritableApprovalWorkflowRequest) SetUser(v ApprovalWorkflowUser)`

SetUser sets User field to given value.

### HasUser

`func (o *BulkWritableApprovalWorkflowRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *BulkWritableApprovalWorkflowRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *BulkWritableApprovalWorkflowRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableApprovalWorkflowRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableApprovalWorkflowRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableApprovalWorkflowRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableApprovalWorkflowRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableApprovalWorkflowRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableApprovalWorkflowRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableApprovalWorkflowRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableApprovalWorkflowRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


