# PatchedRegularExpressionValidationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ErrorMessage** | Pointer to **string** | Optional error message to display when validation fails. | [optional] 
**RegularExpression** | Pointer to **string** |  | [optional] 
**ContextProcessing** | Pointer to **bool** | When enabled, the regular expression value is first processed as a Jinja2 template with access to the object being validated in a variable named &lt;code&gt;obj&lt;/code&gt;. | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedRegularExpressionValidationRuleRequest

`func NewPatchedRegularExpressionValidationRuleRequest() *PatchedRegularExpressionValidationRuleRequest`

NewPatchedRegularExpressionValidationRuleRequest instantiates a new PatchedRegularExpressionValidationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRegularExpressionValidationRuleRequestWithDefaults

`func NewPatchedRegularExpressionValidationRuleRequestWithDefaults() *PatchedRegularExpressionValidationRuleRequest`

NewPatchedRegularExpressionValidationRuleRequestWithDefaults instantiates a new PatchedRegularExpressionValidationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedRegularExpressionValidationRuleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedRegularExpressionValidationRuleRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedRegularExpressionValidationRuleRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentType

`func (o *PatchedRegularExpressionValidationRuleRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedRegularExpressionValidationRuleRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedRegularExpressionValidationRuleRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetName

`func (o *PatchedRegularExpressionValidationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRegularExpressionValidationRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRegularExpressionValidationRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetField

`func (o *PatchedRegularExpressionValidationRuleRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *PatchedRegularExpressionValidationRuleRequest) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *PatchedRegularExpressionValidationRuleRequest) HasField() bool`

HasField returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedRegularExpressionValidationRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedRegularExpressionValidationRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedRegularExpressionValidationRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PatchedRegularExpressionValidationRuleRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PatchedRegularExpressionValidationRuleRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PatchedRegularExpressionValidationRuleRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetRegularExpression

`func (o *PatchedRegularExpressionValidationRuleRequest) GetRegularExpression() string`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetRegularExpressionOk() (*string, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *PatchedRegularExpressionValidationRuleRequest) SetRegularExpression(v string)`

SetRegularExpression sets RegularExpression field to given value.

### HasRegularExpression

`func (o *PatchedRegularExpressionValidationRuleRequest) HasRegularExpression() bool`

HasRegularExpression returns a boolean if a field has been set.

### GetContextProcessing

`func (o *PatchedRegularExpressionValidationRuleRequest) GetContextProcessing() bool`

GetContextProcessing returns the ContextProcessing field if non-nil, zero value otherwise.

### GetContextProcessingOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetContextProcessingOk() (*bool, bool)`

GetContextProcessingOk returns a tuple with the ContextProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextProcessing

`func (o *PatchedRegularExpressionValidationRuleRequest) SetContextProcessing(v bool)`

SetContextProcessing sets ContextProcessing field to given value.

### HasContextProcessing

`func (o *PatchedRegularExpressionValidationRuleRequest) HasContextProcessing() bool`

HasContextProcessing returns a boolean if a field has been set.

### GetTags

`func (o *PatchedRegularExpressionValidationRuleRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedRegularExpressionValidationRuleRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedRegularExpressionValidationRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedRegularExpressionValidationRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedRegularExpressionValidationRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedRegularExpressionValidationRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedRegularExpressionValidationRuleRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedRegularExpressionValidationRuleRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedRegularExpressionValidationRuleRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedRegularExpressionValidationRuleRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


