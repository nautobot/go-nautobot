# DataComplianceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ComplianceClassName** | **string** |  | 
**ObjectId** | **string** |  | 
**ValidatedObjectStr** | Pointer to **string** |  | [optional] 
**ValidatedAttribute** | Pointer to **string** |  | [optional] [default to ""]
**ValidatedAttributeValue** | Pointer to **string** |  | [optional] 
**Valid** | **bool** |  | 
**Message** | Pointer to **string** |  | [optional] 
**ContentType** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewDataComplianceRequest

`func NewDataComplianceRequest(complianceClassName string, objectId string, valid bool, contentType ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *DataComplianceRequest`

NewDataComplianceRequest instantiates a new DataComplianceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataComplianceRequestWithDefaults

`func NewDataComplianceRequestWithDefaults() *DataComplianceRequest`

NewDataComplianceRequestWithDefaults instantiates a new DataComplianceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataComplianceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataComplianceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataComplianceRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataComplianceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComplianceClassName

`func (o *DataComplianceRequest) GetComplianceClassName() string`

GetComplianceClassName returns the ComplianceClassName field if non-nil, zero value otherwise.

### GetComplianceClassNameOk

`func (o *DataComplianceRequest) GetComplianceClassNameOk() (*string, bool)`

GetComplianceClassNameOk returns a tuple with the ComplianceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceClassName

`func (o *DataComplianceRequest) SetComplianceClassName(v string)`

SetComplianceClassName sets ComplianceClassName field to given value.


### GetObjectId

`func (o *DataComplianceRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *DataComplianceRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *DataComplianceRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetValidatedObjectStr

`func (o *DataComplianceRequest) GetValidatedObjectStr() string`

GetValidatedObjectStr returns the ValidatedObjectStr field if non-nil, zero value otherwise.

### GetValidatedObjectStrOk

`func (o *DataComplianceRequest) GetValidatedObjectStrOk() (*string, bool)`

GetValidatedObjectStrOk returns a tuple with the ValidatedObjectStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedObjectStr

`func (o *DataComplianceRequest) SetValidatedObjectStr(v string)`

SetValidatedObjectStr sets ValidatedObjectStr field to given value.

### HasValidatedObjectStr

`func (o *DataComplianceRequest) HasValidatedObjectStr() bool`

HasValidatedObjectStr returns a boolean if a field has been set.

### GetValidatedAttribute

`func (o *DataComplianceRequest) GetValidatedAttribute() string`

GetValidatedAttribute returns the ValidatedAttribute field if non-nil, zero value otherwise.

### GetValidatedAttributeOk

`func (o *DataComplianceRequest) GetValidatedAttributeOk() (*string, bool)`

GetValidatedAttributeOk returns a tuple with the ValidatedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAttribute

`func (o *DataComplianceRequest) SetValidatedAttribute(v string)`

SetValidatedAttribute sets ValidatedAttribute field to given value.

### HasValidatedAttribute

`func (o *DataComplianceRequest) HasValidatedAttribute() bool`

HasValidatedAttribute returns a boolean if a field has been set.

### GetValidatedAttributeValue

`func (o *DataComplianceRequest) GetValidatedAttributeValue() string`

GetValidatedAttributeValue returns the ValidatedAttributeValue field if non-nil, zero value otherwise.

### GetValidatedAttributeValueOk

`func (o *DataComplianceRequest) GetValidatedAttributeValueOk() (*string, bool)`

GetValidatedAttributeValueOk returns a tuple with the ValidatedAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAttributeValue

`func (o *DataComplianceRequest) SetValidatedAttributeValue(v string)`

SetValidatedAttributeValue sets ValidatedAttributeValue field to given value.

### HasValidatedAttributeValue

`func (o *DataComplianceRequest) HasValidatedAttributeValue() bool`

HasValidatedAttributeValue returns a boolean if a field has been set.

### GetValid

`func (o *DataComplianceRequest) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DataComplianceRequest) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DataComplianceRequest) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetMessage

`func (o *DataComplianceRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DataComplianceRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DataComplianceRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DataComplianceRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetContentType

`func (o *DataComplianceRequest) GetContentType() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DataComplianceRequest) GetContentTypeOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DataComplianceRequest) SetContentType(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetContentType sets ContentType field to given value.


### GetCustomFields

`func (o *DataComplianceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DataComplianceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DataComplianceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DataComplianceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *DataComplianceRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *DataComplianceRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *DataComplianceRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *DataComplianceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


