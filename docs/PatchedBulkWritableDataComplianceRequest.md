# PatchedBulkWritableDataComplianceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ComplianceClassName** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ValidatedObjectStr** | Pointer to **string** |  | [optional] 
**ValidatedAttribute** | Pointer to **string** |  | [optional] [default to ""]
**ValidatedAttributeValue** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableDataComplianceRequest

`func NewPatchedBulkWritableDataComplianceRequest(id string, ) *PatchedBulkWritableDataComplianceRequest`

NewPatchedBulkWritableDataComplianceRequest instantiates a new PatchedBulkWritableDataComplianceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableDataComplianceRequestWithDefaults

`func NewPatchedBulkWritableDataComplianceRequestWithDefaults() *PatchedBulkWritableDataComplianceRequest`

NewPatchedBulkWritableDataComplianceRequestWithDefaults instantiates a new PatchedBulkWritableDataComplianceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableDataComplianceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableDataComplianceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetComplianceClassName

`func (o *PatchedBulkWritableDataComplianceRequest) GetComplianceClassName() string`

GetComplianceClassName returns the ComplianceClassName field if non-nil, zero value otherwise.

### GetComplianceClassNameOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetComplianceClassNameOk() (*string, bool)`

GetComplianceClassNameOk returns a tuple with the ComplianceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceClassName

`func (o *PatchedBulkWritableDataComplianceRequest) SetComplianceClassName(v string)`

SetComplianceClassName sets ComplianceClassName field to given value.

### HasComplianceClassName

`func (o *PatchedBulkWritableDataComplianceRequest) HasComplianceClassName() bool`

HasComplianceClassName returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedBulkWritableDataComplianceRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedBulkWritableDataComplianceRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedBulkWritableDataComplianceRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetValidatedObjectStr

`func (o *PatchedBulkWritableDataComplianceRequest) GetValidatedObjectStr() string`

GetValidatedObjectStr returns the ValidatedObjectStr field if non-nil, zero value otherwise.

### GetValidatedObjectStrOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetValidatedObjectStrOk() (*string, bool)`

GetValidatedObjectStrOk returns a tuple with the ValidatedObjectStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedObjectStr

`func (o *PatchedBulkWritableDataComplianceRequest) SetValidatedObjectStr(v string)`

SetValidatedObjectStr sets ValidatedObjectStr field to given value.

### HasValidatedObjectStr

`func (o *PatchedBulkWritableDataComplianceRequest) HasValidatedObjectStr() bool`

HasValidatedObjectStr returns a boolean if a field has been set.

### GetValidatedAttribute

`func (o *PatchedBulkWritableDataComplianceRequest) GetValidatedAttribute() string`

GetValidatedAttribute returns the ValidatedAttribute field if non-nil, zero value otherwise.

### GetValidatedAttributeOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetValidatedAttributeOk() (*string, bool)`

GetValidatedAttributeOk returns a tuple with the ValidatedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAttribute

`func (o *PatchedBulkWritableDataComplianceRequest) SetValidatedAttribute(v string)`

SetValidatedAttribute sets ValidatedAttribute field to given value.

### HasValidatedAttribute

`func (o *PatchedBulkWritableDataComplianceRequest) HasValidatedAttribute() bool`

HasValidatedAttribute returns a boolean if a field has been set.

### GetValidatedAttributeValue

`func (o *PatchedBulkWritableDataComplianceRequest) GetValidatedAttributeValue() string`

GetValidatedAttributeValue returns the ValidatedAttributeValue field if non-nil, zero value otherwise.

### GetValidatedAttributeValueOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetValidatedAttributeValueOk() (*string, bool)`

GetValidatedAttributeValueOk returns a tuple with the ValidatedAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAttributeValue

`func (o *PatchedBulkWritableDataComplianceRequest) SetValidatedAttributeValue(v string)`

SetValidatedAttributeValue sets ValidatedAttributeValue field to given value.

### HasValidatedAttributeValue

`func (o *PatchedBulkWritableDataComplianceRequest) HasValidatedAttributeValue() bool`

HasValidatedAttributeValue returns a boolean if a field has been set.

### GetValid

`func (o *PatchedBulkWritableDataComplianceRequest) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *PatchedBulkWritableDataComplianceRequest) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *PatchedBulkWritableDataComplianceRequest) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetMessage

`func (o *PatchedBulkWritableDataComplianceRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PatchedBulkWritableDataComplianceRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PatchedBulkWritableDataComplianceRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetContentType

`func (o *PatchedBulkWritableDataComplianceRequest) GetContentType() BulkWritableCableRequestStatus`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetContentTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedBulkWritableDataComplianceRequest) SetContentType(v BulkWritableCableRequestStatus)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedBulkWritableDataComplianceRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableDataComplianceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableDataComplianceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableDataComplianceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableDataComplianceRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableDataComplianceRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableDataComplianceRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableDataComplianceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


