# ApprovalWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ObjectUnderReviewContentType** | **string** |  | 
**DecisionDate** | **NullableTime** |  | [readonly] 
**ObjectUnderReviewObjectId** | **string** |  | 
**CurrentState** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | Current state of the approval workflow. Eligible values are: Pending, Approved, Denied, Canceled. | [optional] 
**UserName** | **string** |  | [readonly] 
**ApprovalWorkflowDefinition** | Pointer to [**NullableApprovalWorkflowApprovalWorkflowDefinition**](ApprovalWorkflowApprovalWorkflowDefinition.md) |  | [optional] 
**User** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApprovalWorkflow

`func NewApprovalWorkflow(objectType string, display string, url string, naturalSlug string, objectUnderReviewContentType string, decisionDate NullableTime, objectUnderReviewObjectId string, userName string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *ApprovalWorkflow`

NewApprovalWorkflow instantiates a new ApprovalWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWorkflowWithDefaults

`func NewApprovalWorkflowWithDefaults() *ApprovalWorkflow`

NewApprovalWorkflowWithDefaults instantiates a new ApprovalWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalWorkflow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalWorkflow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalWorkflow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalWorkflow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *ApprovalWorkflow) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApprovalWorkflow) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApprovalWorkflow) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ApprovalWorkflow) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ApprovalWorkflow) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ApprovalWorkflow) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ApprovalWorkflow) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApprovalWorkflow) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApprovalWorkflow) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ApprovalWorkflow) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ApprovalWorkflow) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ApprovalWorkflow) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetObjectUnderReviewContentType

`func (o *ApprovalWorkflow) GetObjectUnderReviewContentType() string`

GetObjectUnderReviewContentType returns the ObjectUnderReviewContentType field if non-nil, zero value otherwise.

### GetObjectUnderReviewContentTypeOk

`func (o *ApprovalWorkflow) GetObjectUnderReviewContentTypeOk() (*string, bool)`

GetObjectUnderReviewContentTypeOk returns a tuple with the ObjectUnderReviewContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUnderReviewContentType

`func (o *ApprovalWorkflow) SetObjectUnderReviewContentType(v string)`

SetObjectUnderReviewContentType sets ObjectUnderReviewContentType field to given value.


### GetDecisionDate

`func (o *ApprovalWorkflow) GetDecisionDate() time.Time`

GetDecisionDate returns the DecisionDate field if non-nil, zero value otherwise.

### GetDecisionDateOk

`func (o *ApprovalWorkflow) GetDecisionDateOk() (*time.Time, bool)`

GetDecisionDateOk returns a tuple with the DecisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDate

`func (o *ApprovalWorkflow) SetDecisionDate(v time.Time)`

SetDecisionDate sets DecisionDate field to given value.


### SetDecisionDateNil

`func (o *ApprovalWorkflow) SetDecisionDateNil(b bool)`

 SetDecisionDateNil sets the value for DecisionDate to be an explicit nil

### UnsetDecisionDate
`func (o *ApprovalWorkflow) UnsetDecisionDate()`

UnsetDecisionDate ensures that no value is present for DecisionDate, not even an explicit nil
### GetObjectUnderReviewObjectId

`func (o *ApprovalWorkflow) GetObjectUnderReviewObjectId() string`

GetObjectUnderReviewObjectId returns the ObjectUnderReviewObjectId field if non-nil, zero value otherwise.

### GetObjectUnderReviewObjectIdOk

`func (o *ApprovalWorkflow) GetObjectUnderReviewObjectIdOk() (*string, bool)`

GetObjectUnderReviewObjectIdOk returns a tuple with the ObjectUnderReviewObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUnderReviewObjectId

`func (o *ApprovalWorkflow) SetObjectUnderReviewObjectId(v string)`

SetObjectUnderReviewObjectId sets ObjectUnderReviewObjectId field to given value.


### GetCurrentState

`func (o *ApprovalWorkflow) GetCurrentState() ApprovalWorkflowStateChoices`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *ApprovalWorkflow) GetCurrentStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *ApprovalWorkflow) SetCurrentState(v ApprovalWorkflowStateChoices)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *ApprovalWorkflow) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetUserName

`func (o *ApprovalWorkflow) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ApprovalWorkflow) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ApprovalWorkflow) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetApprovalWorkflowDefinition

`func (o *ApprovalWorkflow) GetApprovalWorkflowDefinition() ApprovalWorkflowApprovalWorkflowDefinition`

GetApprovalWorkflowDefinition returns the ApprovalWorkflowDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowDefinitionOk

`func (o *ApprovalWorkflow) GetApprovalWorkflowDefinitionOk() (*ApprovalWorkflowApprovalWorkflowDefinition, bool)`

GetApprovalWorkflowDefinitionOk returns a tuple with the ApprovalWorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowDefinition

`func (o *ApprovalWorkflow) SetApprovalWorkflowDefinition(v ApprovalWorkflowApprovalWorkflowDefinition)`

SetApprovalWorkflowDefinition sets ApprovalWorkflowDefinition field to given value.

### HasApprovalWorkflowDefinition

`func (o *ApprovalWorkflow) HasApprovalWorkflowDefinition() bool`

HasApprovalWorkflowDefinition returns a boolean if a field has been set.

### SetApprovalWorkflowDefinitionNil

`func (o *ApprovalWorkflow) SetApprovalWorkflowDefinitionNil(b bool)`

 SetApprovalWorkflowDefinitionNil sets the value for ApprovalWorkflowDefinition to be an explicit nil

### UnsetApprovalWorkflowDefinition
`func (o *ApprovalWorkflow) UnsetApprovalWorkflowDefinition()`

UnsetApprovalWorkflowDefinition ensures that no value is present for ApprovalWorkflowDefinition, not even an explicit nil
### GetUser

`func (o *ApprovalWorkflow) GetUser() ApprovalWorkflowUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApprovalWorkflow) GetUserOk() (*ApprovalWorkflowUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApprovalWorkflow) SetUser(v ApprovalWorkflowUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ApprovalWorkflow) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ApprovalWorkflow) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ApprovalWorkflow) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetCreated

`func (o *ApprovalWorkflow) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApprovalWorkflow) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApprovalWorkflow) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ApprovalWorkflow) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ApprovalWorkflow) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ApprovalWorkflow) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ApprovalWorkflow) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ApprovalWorkflow) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ApprovalWorkflow) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ApprovalWorkflow) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *ApprovalWorkflow) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *ApprovalWorkflow) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *ApprovalWorkflow) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *ApprovalWorkflow) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ApprovalWorkflow) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ApprovalWorkflow) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ApprovalWorkflow) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


