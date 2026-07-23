# BulkWritableDataComplianceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewBulkWritableDataComplianceRequest

`func NewBulkWritableDataComplianceRequest(id string, complianceClassName string, objectId string, valid bool, contentType ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *BulkWritableDataComplianceRequest`

NewBulkWritableDataComplianceRequest instantiates a new BulkWritableDataComplianceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableDataComplianceRequestWithDefaults

`func NewBulkWritableDataComplianceRequestWithDefaults() *BulkWritableDataComplianceRequest`

NewBulkWritableDataComplianceRequestWithDefaults instantiates a new BulkWritableDataComplianceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableDataComplianceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableDataComplianceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableDataComplianceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetComplianceClassName

`func (o *BulkWritableDataComplianceRequest) GetComplianceClassName() string`

GetComplianceClassName returns the ComplianceClassName field if non-nil, zero value otherwise.

### GetComplianceClassNameOk

`func (o *BulkWritableDataComplianceRequest) GetComplianceClassNameOk() (*string, bool)`

GetComplianceClassNameOk returns a tuple with the ComplianceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceClassName

`func (o *BulkWritableDataComplianceRequest) SetComplianceClassName(v string)`

SetComplianceClassName sets ComplianceClassName field to given value.


### GetObjectId

`func (o *BulkWritableDataComplianceRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BulkWritableDataComplianceRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BulkWritableDataComplianceRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetValidatedObjectStr

`func (o *BulkWritableDataComplianceRequest) GetValidatedObjectStr() string`

GetValidatedObjectStr returns the ValidatedObjectStr field if non-nil, zero value otherwise.

### GetValidatedObjectStrOk

`func (o *BulkWritableDataComplianceRequest) GetValidatedObjectStrOk() (*string, bool)`

GetValidatedObjectStrOk returns a tuple with the ValidatedObjectStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedObjectStr

`func (o *BulkWritableDataComplianceRequest) SetValidatedObjectStr(v string)`

SetValidatedObjectStr sets ValidatedObjectStr field to given value.

### HasValidatedObjectStr

`func (o *BulkWritableDataComplianceRequest) HasValidatedObjectStr() bool`

HasValidatedObjectStr returns a boolean if a field has been set.

### GetValidatedAttribute

`func (o *BulkWritableDataComplianceRequest) GetValidatedAttribute() string`

GetValidatedAttribute returns the ValidatedAttribute field if non-nil, zero value otherwise.

### GetValidatedAttributeOk

`func (o *BulkWritableDataComplianceRequest) GetValidatedAttributeOk() (*string, bool)`

GetValidatedAttributeOk returns a tuple with the ValidatedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAttribute

`func (o *BulkWritableDataComplianceRequest) SetValidatedAttribute(v string)`

SetValidatedAttribute sets ValidatedAttribute field to given value.

### HasValidatedAttribute

`func (o *BulkWritableDataComplianceRequest) HasValidatedAttribute() bool`

HasValidatedAttribute returns a boolean if a field has been set.

### GetValidatedAttributeValue

`func (o *BulkWritableDataComplianceRequest) GetValidatedAttributeValue() string`

GetValidatedAttributeValue returns the ValidatedAttributeValue field if non-nil, zero value otherwise.

### GetValidatedAttributeValueOk

`func (o *BulkWritableDataComplianceRequest) GetValidatedAttributeValueOk() (*string, bool)`

GetValidatedAttributeValueOk returns a tuple with the ValidatedAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAttributeValue

`func (o *BulkWritableDataComplianceRequest) SetValidatedAttributeValue(v string)`

SetValidatedAttributeValue sets ValidatedAttributeValue field to given value.

### HasValidatedAttributeValue

`func (o *BulkWritableDataComplianceRequest) HasValidatedAttributeValue() bool`

HasValidatedAttributeValue returns a boolean if a field has been set.

### GetValid

`func (o *BulkWritableDataComplianceRequest) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *BulkWritableDataComplianceRequest) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *BulkWritableDataComplianceRequest) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetMessage

`func (o *BulkWritableDataComplianceRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BulkWritableDataComplianceRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BulkWritableDataComplianceRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BulkWritableDataComplianceRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetContentType

`func (o *BulkWritableDataComplianceRequest) GetContentType() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkWritableDataComplianceRequest) GetContentTypeOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkWritableDataComplianceRequest) SetContentType(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetContentType sets ContentType field to given value.


### GetCustomFields

`func (o *BulkWritableDataComplianceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableDataComplianceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableDataComplianceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableDataComplianceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableDataComplianceRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableDataComplianceRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableDataComplianceRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableDataComplianceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


