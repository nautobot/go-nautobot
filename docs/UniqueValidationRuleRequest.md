# UniqueValidationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContentType** | **string** |  | 
**Name** | **string** |  | 
**Field** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**ErrorMessage** | Pointer to **string** | Optional error message to display when validation fails. | [optional] 
**MaxInstances** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewUniqueValidationRuleRequest

`func NewUniqueValidationRuleRequest(contentType string, name string, field string, ) *UniqueValidationRuleRequest`

NewUniqueValidationRuleRequest instantiates a new UniqueValidationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniqueValidationRuleRequestWithDefaults

`func NewUniqueValidationRuleRequestWithDefaults() *UniqueValidationRuleRequest`

NewUniqueValidationRuleRequestWithDefaults instantiates a new UniqueValidationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UniqueValidationRuleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UniqueValidationRuleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UniqueValidationRuleRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UniqueValidationRuleRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentType

`func (o *UniqueValidationRuleRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *UniqueValidationRuleRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *UniqueValidationRuleRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetName

`func (o *UniqueValidationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniqueValidationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniqueValidationRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetField

`func (o *UniqueValidationRuleRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *UniqueValidationRuleRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *UniqueValidationRuleRequest) SetField(v string)`

SetField sets Field field to given value.


### GetEnabled

`func (o *UniqueValidationRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UniqueValidationRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UniqueValidationRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UniqueValidationRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorMessage

`func (o *UniqueValidationRuleRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *UniqueValidationRuleRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *UniqueValidationRuleRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *UniqueValidationRuleRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetMaxInstances

`func (o *UniqueValidationRuleRequest) GetMaxInstances() int32`

GetMaxInstances returns the MaxInstances field if non-nil, zero value otherwise.

### GetMaxInstancesOk

`func (o *UniqueValidationRuleRequest) GetMaxInstancesOk() (*int32, bool)`

GetMaxInstancesOk returns a tuple with the MaxInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstances

`func (o *UniqueValidationRuleRequest) SetMaxInstances(v int32)`

SetMaxInstances sets MaxInstances field to given value.

### HasMaxInstances

`func (o *UniqueValidationRuleRequest) HasMaxInstances() bool`

HasMaxInstances returns a boolean if a field has been set.

### GetTags

`func (o *UniqueValidationRuleRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UniqueValidationRuleRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UniqueValidationRuleRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UniqueValidationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *UniqueValidationRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UniqueValidationRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UniqueValidationRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UniqueValidationRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *UniqueValidationRuleRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UniqueValidationRuleRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UniqueValidationRuleRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *UniqueValidationRuleRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


