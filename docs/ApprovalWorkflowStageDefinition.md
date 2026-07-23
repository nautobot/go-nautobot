# ApprovalWorkflowStageDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ApproverGroup** | **string** | The group that will be assigned to approve this stage. | 
**Sequence** | **int32** | The sequence dictates the order in which this stage will need to be approved. The lower the number, the earlier it will be. | 
**Name** | **string** |  | 
**MinApprovers** | **int32** | Minimum number of approvers required to approve this stage. | 
**DenialMessage** | Pointer to **string** | Message to show when the stage is denied. | [optional] 
**ApprovalWorkflowDefinition** | [**ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition**](ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApprovalWorkflowStageDefinition

`func NewApprovalWorkflowStageDefinition(objectType string, display string, url string, naturalSlug string, approverGroup string, sequence int32, name string, minApprovers int32, approvalWorkflowDefinition ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *ApprovalWorkflowStageDefinition`

NewApprovalWorkflowStageDefinition instantiates a new ApprovalWorkflowStageDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWorkflowStageDefinitionWithDefaults

`func NewApprovalWorkflowStageDefinitionWithDefaults() *ApprovalWorkflowStageDefinition`

NewApprovalWorkflowStageDefinitionWithDefaults instantiates a new ApprovalWorkflowStageDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalWorkflowStageDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalWorkflowStageDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalWorkflowStageDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalWorkflowStageDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *ApprovalWorkflowStageDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApprovalWorkflowStageDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApprovalWorkflowStageDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ApprovalWorkflowStageDefinition) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ApprovalWorkflowStageDefinition) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ApprovalWorkflowStageDefinition) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ApprovalWorkflowStageDefinition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApprovalWorkflowStageDefinition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApprovalWorkflowStageDefinition) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ApprovalWorkflowStageDefinition) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ApprovalWorkflowStageDefinition) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ApprovalWorkflowStageDefinition) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetApproverGroup

`func (o *ApprovalWorkflowStageDefinition) GetApproverGroup() string`

GetApproverGroup returns the ApproverGroup field if non-nil, zero value otherwise.

### GetApproverGroupOk

`func (o *ApprovalWorkflowStageDefinition) GetApproverGroupOk() (*string, bool)`

GetApproverGroupOk returns a tuple with the ApproverGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverGroup

`func (o *ApprovalWorkflowStageDefinition) SetApproverGroup(v string)`

SetApproverGroup sets ApproverGroup field to given value.


### GetSequence

`func (o *ApprovalWorkflowStageDefinition) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ApprovalWorkflowStageDefinition) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ApprovalWorkflowStageDefinition) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetName

`func (o *ApprovalWorkflowStageDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalWorkflowStageDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalWorkflowStageDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetMinApprovers

`func (o *ApprovalWorkflowStageDefinition) GetMinApprovers() int32`

GetMinApprovers returns the MinApprovers field if non-nil, zero value otherwise.

### GetMinApproversOk

`func (o *ApprovalWorkflowStageDefinition) GetMinApproversOk() (*int32, bool)`

GetMinApproversOk returns a tuple with the MinApprovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinApprovers

`func (o *ApprovalWorkflowStageDefinition) SetMinApprovers(v int32)`

SetMinApprovers sets MinApprovers field to given value.


### GetDenialMessage

`func (o *ApprovalWorkflowStageDefinition) GetDenialMessage() string`

GetDenialMessage returns the DenialMessage field if non-nil, zero value otherwise.

### GetDenialMessageOk

`func (o *ApprovalWorkflowStageDefinition) GetDenialMessageOk() (*string, bool)`

GetDenialMessageOk returns a tuple with the DenialMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialMessage

`func (o *ApprovalWorkflowStageDefinition) SetDenialMessage(v string)`

SetDenialMessage sets DenialMessage field to given value.

### HasDenialMessage

`func (o *ApprovalWorkflowStageDefinition) HasDenialMessage() bool`

HasDenialMessage returns a boolean if a field has been set.

### GetApprovalWorkflowDefinition

`func (o *ApprovalWorkflowStageDefinition) GetApprovalWorkflowDefinition() ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition`

GetApprovalWorkflowDefinition returns the ApprovalWorkflowDefinition field if non-nil, zero value otherwise.

### GetApprovalWorkflowDefinitionOk

`func (o *ApprovalWorkflowStageDefinition) GetApprovalWorkflowDefinitionOk() (*ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition, bool)`

GetApprovalWorkflowDefinitionOk returns a tuple with the ApprovalWorkflowDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWorkflowDefinition

`func (o *ApprovalWorkflowStageDefinition) SetApprovalWorkflowDefinition(v ApprovalWorkflowStageDefinitionApprovalWorkflowDefinition)`

SetApprovalWorkflowDefinition sets ApprovalWorkflowDefinition field to given value.


### GetCreated

`func (o *ApprovalWorkflowStageDefinition) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApprovalWorkflowStageDefinition) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApprovalWorkflowStageDefinition) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ApprovalWorkflowStageDefinition) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ApprovalWorkflowStageDefinition) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ApprovalWorkflowStageDefinition) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ApprovalWorkflowStageDefinition) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ApprovalWorkflowStageDefinition) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ApprovalWorkflowStageDefinition) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ApprovalWorkflowStageDefinition) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *ApprovalWorkflowStageDefinition) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *ApprovalWorkflowStageDefinition) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *ApprovalWorkflowStageDefinition) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *ApprovalWorkflowStageDefinition) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ApprovalWorkflowStageDefinition) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ApprovalWorkflowStageDefinition) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ApprovalWorkflowStageDefinition) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


