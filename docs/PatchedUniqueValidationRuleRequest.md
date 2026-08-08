# PatchedUniqueValidationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ErrorMessage** | Pointer to **string** | Optional error message to display when validation fails. | [optional] 
**MaxInstances** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedUniqueValidationRuleRequest

`func NewPatchedUniqueValidationRuleRequest() *PatchedUniqueValidationRuleRequest`

NewPatchedUniqueValidationRuleRequest instantiates a new PatchedUniqueValidationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUniqueValidationRuleRequestWithDefaults

`func NewPatchedUniqueValidationRuleRequestWithDefaults() *PatchedUniqueValidationRuleRequest`

NewPatchedUniqueValidationRuleRequestWithDefaults instantiates a new PatchedUniqueValidationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedUniqueValidationRuleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedUniqueValidationRuleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedUniqueValidationRuleRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedUniqueValidationRuleRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentType

`func (o *PatchedUniqueValidationRuleRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedUniqueValidationRuleRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedUniqueValidationRuleRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedUniqueValidationRuleRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetName

`func (o *PatchedUniqueValidationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedUniqueValidationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedUniqueValidationRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedUniqueValidationRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetField

`func (o *PatchedUniqueValidationRuleRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *PatchedUniqueValidationRuleRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *PatchedUniqueValidationRuleRequest) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *PatchedUniqueValidationRuleRequest) HasField() bool`

HasField returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedUniqueValidationRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedUniqueValidationRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedUniqueValidationRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedUniqueValidationRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PatchedUniqueValidationRuleRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PatchedUniqueValidationRuleRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PatchedUniqueValidationRuleRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PatchedUniqueValidationRuleRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetMaxInstances

`func (o *PatchedUniqueValidationRuleRequest) GetMaxInstances() int32`

GetMaxInstances returns the MaxInstances field if non-nil, zero value otherwise.

### GetMaxInstancesOk

`func (o *PatchedUniqueValidationRuleRequest) GetMaxInstancesOk() (*int32, bool)`

GetMaxInstancesOk returns a tuple with the MaxInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstances

`func (o *PatchedUniqueValidationRuleRequest) SetMaxInstances(v int32)`

SetMaxInstances sets MaxInstances field to given value.

### HasMaxInstances

`func (o *PatchedUniqueValidationRuleRequest) HasMaxInstances() bool`

HasMaxInstances returns a boolean if a field has been set.

### GetTags

`func (o *PatchedUniqueValidationRuleRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedUniqueValidationRuleRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedUniqueValidationRuleRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedUniqueValidationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedUniqueValidationRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedUniqueValidationRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedUniqueValidationRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedUniqueValidationRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedUniqueValidationRuleRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedUniqueValidationRuleRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedUniqueValidationRuleRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedUniqueValidationRuleRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


