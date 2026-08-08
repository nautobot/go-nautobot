# BulkWritableMinMaxValidationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentType** | **string** |  | 
**Name** | **string** |  | 
**Field** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**ErrorMessage** | Pointer to **string** | Optional error message to display when validation fails. | [optional] 
**Min** | Pointer to **NullableFloat64** | When set, apply a minimum value contraint to the value of the model field. | [optional] 
**Max** | Pointer to **NullableFloat64** | When set, apply a maximum value contraint to the value of the model field. | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableMinMaxValidationRuleRequest

`func NewBulkWritableMinMaxValidationRuleRequest(id string, contentType string, name string, field string, ) *BulkWritableMinMaxValidationRuleRequest`

NewBulkWritableMinMaxValidationRuleRequest instantiates a new BulkWritableMinMaxValidationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableMinMaxValidationRuleRequestWithDefaults

`func NewBulkWritableMinMaxValidationRuleRequestWithDefaults() *BulkWritableMinMaxValidationRuleRequest`

NewBulkWritableMinMaxValidationRuleRequestWithDefaults instantiates a new BulkWritableMinMaxValidationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableMinMaxValidationRuleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableMinMaxValidationRuleRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentType

`func (o *BulkWritableMinMaxValidationRuleRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkWritableMinMaxValidationRuleRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetName

`func (o *BulkWritableMinMaxValidationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableMinMaxValidationRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetField

`func (o *BulkWritableMinMaxValidationRuleRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *BulkWritableMinMaxValidationRuleRequest) SetField(v string)`

SetField sets Field field to given value.


### GetEnabled

`func (o *BulkWritableMinMaxValidationRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkWritableMinMaxValidationRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkWritableMinMaxValidationRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorMessage

`func (o *BulkWritableMinMaxValidationRuleRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BulkWritableMinMaxValidationRuleRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BulkWritableMinMaxValidationRuleRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetMin

`func (o *BulkWritableMinMaxValidationRuleRequest) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *BulkWritableMinMaxValidationRuleRequest) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *BulkWritableMinMaxValidationRuleRequest) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *BulkWritableMinMaxValidationRuleRequest) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *BulkWritableMinMaxValidationRuleRequest) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *BulkWritableMinMaxValidationRuleRequest) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *BulkWritableMinMaxValidationRuleRequest) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *BulkWritableMinMaxValidationRuleRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *BulkWritableMinMaxValidationRuleRequest) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *BulkWritableMinMaxValidationRuleRequest) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetTags

`func (o *BulkWritableMinMaxValidationRuleRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableMinMaxValidationRuleRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableMinMaxValidationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableMinMaxValidationRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableMinMaxValidationRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableMinMaxValidationRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableMinMaxValidationRuleRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableMinMaxValidationRuleRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableMinMaxValidationRuleRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableMinMaxValidationRuleRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


