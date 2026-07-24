# PatchedBulkWritableMinMaxValidationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ErrorMessage** | Pointer to **string** | Optional error message to display when validation fails. | [optional] 
**Min** | Pointer to **NullableFloat64** | When set, apply a minimum value contraint to the value of the model field. | [optional] 
**Max** | Pointer to **NullableFloat64** | When set, apply a maximum value contraint to the value of the model field. | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableMinMaxValidationRuleRequest

`func NewPatchedBulkWritableMinMaxValidationRuleRequest(id string, ) *PatchedBulkWritableMinMaxValidationRuleRequest`

NewPatchedBulkWritableMinMaxValidationRuleRequest instantiates a new PatchedBulkWritableMinMaxValidationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableMinMaxValidationRuleRequestWithDefaults

`func NewPatchedBulkWritableMinMaxValidationRuleRequestWithDefaults() *PatchedBulkWritableMinMaxValidationRuleRequest`

NewPatchedBulkWritableMinMaxValidationRuleRequestWithDefaults instantiates a new PatchedBulkWritableMinMaxValidationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentType

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetField

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) HasField() bool`

HasField returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetMin

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetTags

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableMinMaxValidationRuleRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


