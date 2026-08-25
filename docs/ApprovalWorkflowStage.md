# ApprovalWorkflowStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**DecisionDate** | **NullableTime** |  | [readonly] 
**Responses** | [**[]ApprovalWorkflowStageResponse**](ApprovalWorkflowStageResponse.md) |  | [readonly] 
**State** | Pointer to [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | State of the approval workflow stage instance. Eligible values are: Pending, Approved, Denied. | [optional] 
**ApprovalWorkflow** | [**ApprovalWorkflowStageApprovalWorkflow**](ApprovalWorkflowStageApprovalWorkflow.md) |  | 
**ApprovalWorkflowStageDefinition** | Pointer to [**NullableApprovalWorkflowStageApprovalWorkflowStageDefinition**](ApprovalWorkflowStageApprovalWorkflowStageDefinition.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApprovalWorkflowStage

`func NewApprovalWorkflowStage(objectType string, display string, url string, naturalSlug string, decisionDate NullableTime, responses []ApprovalWorkflowStageResponse, approvalWorkflow ApprovalWorkflowStageApprovalWorkflow, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *ApprovalWorkflowStage`

NewApprovalWorkflowStage instantiates a new ApprovalWorkflowStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWorkflowStageWithDefaults

`func NewApprovalWorkflowStageWithDefaults() *ApprovalWorkflowStage`

NewApprovalWorkflowStageWithDefaults instantiates a new ApprovalWorkflowStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalWorkflowStage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalWorkflowStage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalWorkflowStage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalWorkflowStage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *ApprovalWorkflowStage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApprovalWorkflowStage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApprovalWorkflowStage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ApprovalWorkflowStage) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ApprovalWorkflowStage) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ApprovalWorkflowStage) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ApprovalWorkflowStage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApprovalWorkflowStage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApprovalWorkflowStage) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ApprovalWorkflowStage) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ApprovalWorkflowStage) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ApprovalWorkflowStage) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetDecisionDate

`func (o *ApprovalWorkflowStage) GetDecisionDate() time.Time`

GetDecisionDate returns the DecisionDate field if non-nil, zero value otherwise.

### GetDecisionDateOk

`func (o *ApprovalWorkflowStage) GetDecisionDateOk() (*time.Time, bool)`

GetDecisionDateOk returns a tuple with the DecisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDate

`func (o *ApprovalWorkflowStage) SetDecisionDate(v time.Time)`

SetDecisionDate sets DecisionDate field to given value.


### SetDecisionDateNil

`func (o *ApprovalWorkflowStage) SetDecisionDateNil(b bool)`

 SetDecisionDateNil sets the value for DecisionDate to be an explicit nil

### UnsetDecisionDate
`func (o *ApprovalWorkflowStage) UnsetDecisionDate()`

UnsetDecisionDate ensures that no value is present for DecisionDate, not even an explicit nil
### GetResponses

`func (o *ApprovalWorkflowStage) GetResponses() []ApprovalWorkflowStageResponse`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *ApprovalWorkflowStage) GetResponsesOk() (*[]ApprovalWorkflowStageResponse, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *ApprovalWorkflowStage) SetResponses(v []ApprovalWorkflowStageResponse)`

SetResponses sets Responses field to given value.


### GetState

`func (o *ApprovalWorkflowStage) GetState() ApprovalWorkflowStateChoices`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApprovalWorkflowStage) GetStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApprovalWorkflowStage) SetState(v ApprovalWorkflowStateChoices)`

SetState sets State field to given value.

### HasState

`func (o *ApprovalWorkflowStage) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApprovalWorkflow

`func (o *ApprovalWorkflowStage) GetApprovalWorkflow() ApprovalWorkflowStageApprovalWorkflow`

GetApprovalWorkflow returns the ApprovalWorkflow field if non-nil, zero value otherwise.

### GetApprovalWorkflowOk

`func (o *ApprovalWorkflowStage) GetApprovalWorkflowOk() (*ApprovalWorkflowStageApprovalWorkflow, bool)`

GetApprovalWorkflowOk returns a tuple with the ApprovalWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflow

`func (o *ApprovalWorkflowStage) SetApprovalWorkflow(v ApprovalWorkflowStageApprovalWorkflow)`

SetApprovalWorkflow sets ApprovalWorkflow field to given value.


### GetApprovalWorkflowStageDefinition

`func (o *ApprovalWorkflowStage) GetApprovalWorkflowStageDefinition() ApprovalWorkflowStageApprovalWorkflowStageDefinition`

GetApprovalWorkflowStageDefinition returns the ApprovalWorkflowStageDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowStageDefinitionOk

`func (o *ApprovalWorkflowStage) GetApprovalWorkflowStageDefinitionOk() (*ApprovalWorkflowStageApprovalWorkflowStageDefinition, bool)`

GetApprovalWorkflowStageDefinitionOk returns a tuple with the ApprovalWorkflowStageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowStageDefinition

`func (o *ApprovalWorkflowStage) SetApprovalWorkflowStageDefinition(v ApprovalWorkflowStageApprovalWorkflowStageDefinition)`

SetApprovalWorkflowStageDefinition sets ApprovalWorkflowStageDefinition field to given value.

### HasApprovalWorkflowStageDefinition

`func (o *ApprovalWorkflowStage) HasApprovalWorkflowStageDefinition() bool`

HasApprovalWorkflowStageDefinition returns a boolean if a field has been set.

### SetApprovalWorkflowStageDefinitionNil

`func (o *ApprovalWorkflowStage) SetApprovalWorkflowStageDefinitionNil(b bool)`

 SetApprovalWorkflowStageDefinitionNil sets the value for ApprovalWorkflowStageDefinition to be an explicit nil

### UnsetApprovalWorkflowStageDefinition
`func (o *ApprovalWorkflowStage) UnsetApprovalWorkflowStageDefinition()`

UnsetApprovalWorkflowStageDefinition ensures that no value is present for ApprovalWorkflowStageDefinition, not even an explicit nil
### GetCreated

`func (o *ApprovalWorkflowStage) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApprovalWorkflowStage) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApprovalWorkflowStage) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ApprovalWorkflowStage) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ApprovalWorkflowStage) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ApprovalWorkflowStage) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ApprovalWorkflowStage) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ApprovalWorkflowStage) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ApprovalWorkflowStage) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ApprovalWorkflowStage) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *ApprovalWorkflowStage) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *ApprovalWorkflowStage) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *ApprovalWorkflowStage) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *ApprovalWorkflowStage) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ApprovalWorkflowStage) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ApprovalWorkflowStage) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ApprovalWorkflowStage) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


