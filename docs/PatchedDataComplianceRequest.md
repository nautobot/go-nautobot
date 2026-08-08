# PatchedDataComplianceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ComplianceClassName** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ValidatedObjectStr** | Pointer to **string** |  | [optional] 
**ValidatedAttribute** | Pointer to **string** |  | [optional] [default to ""]
**ValidatedAttributeValue** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedDataComplianceRequest

`func NewPatchedDataComplianceRequest() *PatchedDataComplianceRequest`

NewPatchedDataComplianceRequest instantiates a new PatchedDataComplianceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDataComplianceRequestWithDefaults

`func NewPatchedDataComplianceRequestWithDefaults() *PatchedDataComplianceRequest`

NewPatchedDataComplianceRequestWithDefaults instantiates a new PatchedDataComplianceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDataComplianceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDataComplianceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDataComplianceRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDataComplianceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComplianceClassName

`func (o *PatchedDataComplianceRequest) GetComplianceClassName() string`

GetComplianceClassName returns the ComplianceClassName field if non-nil, zero value otherwise.

### GetComplianceClassNameOk

`func (o *PatchedDataComplianceRequest) GetComplianceClassNameOk() (*string, bool)`

GetComplianceClassNameOk returns a tuple with the ComplianceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceClassName

`func (o *PatchedDataComplianceRequest) SetComplianceClassName(v string)`

SetComplianceClassName sets ComplianceClassName field to given value.

### HasComplianceClassName

`func (o *PatchedDataComplianceRequest) HasComplianceClassName() bool`

HasComplianceClassName returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedDataComplianceRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedDataComplianceRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedDataComplianceRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedDataComplianceRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetValidatedObjectStr

`func (o *PatchedDataComplianceRequest) GetValidatedObjectStr() string`

GetValidatedObjectStr returns the ValidatedObjectStr field if non-nil, zero value otherwise.

### GetValidatedObjectStrOk

`func (o *PatchedDataComplianceRequest) GetValidatedObjectStrOk() (*string, bool)`

GetValidatedObjectStrOk returns a tuple with the ValidatedObjectStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedObjectStr

`func (o *PatchedDataComplianceRequest) SetValidatedObjectStr(v string)`

SetValidatedObjectStr sets ValidatedObjectStr field to given value.

### HasValidatedObjectStr

`func (o *PatchedDataComplianceRequest) HasValidatedObjectStr() bool`

HasValidatedObjectStr returns a boolean if a field has been set.

### GetValidatedAttribute

`func (o *PatchedDataComplianceRequest) GetValidatedAttribute() string`

GetValidatedAttribute returns the ValidatedAttribute field if non-nil, zero value otherwise.

### GetValidatedAttributeOk

`func (o *PatchedDataComplianceRequest) GetValidatedAttributeOk() (*string, bool)`

GetValidatedAttributeOk returns a tuple with the ValidatedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAttribute

`func (o *PatchedDataComplianceRequest) SetValidatedAttribute(v string)`

SetValidatedAttribute sets ValidatedAttribute field to given value.

### HasValidatedAttribute

`func (o *PatchedDataComplianceRequest) HasValidatedAttribute() bool`

HasValidatedAttribute returns a boolean if a field has been set.

### GetValidatedAttributeValue

`func (o *PatchedDataComplianceRequest) GetValidatedAttributeValue() string`

GetValidatedAttributeValue returns the ValidatedAttributeValue field if non-nil, zero value otherwise.

### GetValidatedAttributeValueOk

`func (o *PatchedDataComplianceRequest) GetValidatedAttributeValueOk() (*string, bool)`

GetValidatedAttributeValueOk returns a tuple with the ValidatedAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAttributeValue

`func (o *PatchedDataComplianceRequest) SetValidatedAttributeValue(v string)`

SetValidatedAttributeValue sets ValidatedAttributeValue field to given value.

### HasValidatedAttributeValue

`func (o *PatchedDataComplianceRequest) HasValidatedAttributeValue() bool`

HasValidatedAttributeValue returns a boolean if a field has been set.

### GetValid

`func (o *PatchedDataComplianceRequest) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *PatchedDataComplianceRequest) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *PatchedDataComplianceRequest) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *PatchedDataComplianceRequest) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetMessage

`func (o *PatchedDataComplianceRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PatchedDataComplianceRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PatchedDataComplianceRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PatchedDataComplianceRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetContentType

`func (o *PatchedDataComplianceRequest) GetContentType() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedDataComplianceRequest) GetContentTypeOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedDataComplianceRequest) SetContentType(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedDataComplianceRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedDataComplianceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedDataComplianceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedDataComplianceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedDataComplianceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedDataComplianceRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedDataComplianceRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedDataComplianceRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedDataComplianceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


