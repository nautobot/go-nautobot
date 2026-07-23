# BulkWritableModuleTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FrontImage** | Pointer to ***os.File** |  | [optional] 
**RearImage** | Pointer to ***os.File** |  | [optional] 
**Model** | **string** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Manufacturer** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**ModuleFamily** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableModuleTypeRequest

`func NewBulkWritableModuleTypeRequest(id string, model string, manufacturer BulkWritableCableRequestStatus, ) *BulkWritableModuleTypeRequest`

NewBulkWritableModuleTypeRequest instantiates a new BulkWritableModuleTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableModuleTypeRequestWithDefaults

`func NewBulkWritableModuleTypeRequestWithDefaults() *BulkWritableModuleTypeRequest`

NewBulkWritableModuleTypeRequestWithDefaults instantiates a new BulkWritableModuleTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableModuleTypeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableModuleTypeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableModuleTypeRequest) SetId(v string)`

SetId sets Id field to given value.


### GetFrontImage

`func (o *BulkWritableModuleTypeRequest) GetFrontImage() *os.File`

GetFrontImage returns the FrontImage field if non-nil, zero value otherwise.

### GetFrontImageOk

`func (o *BulkWritableModuleTypeRequest) GetFrontImageOk() (**os.File, bool)`

GetFrontImageOk returns a tuple with the FrontImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImage

`func (o *BulkWritableModuleTypeRequest) SetFrontImage(v *os.File)`

SetFrontImage sets FrontImage field to given value.

### HasFrontImage

`func (o *BulkWritableModuleTypeRequest) HasFrontImage() bool`

HasFrontImage returns a boolean if a field has been set.

### GetRearImage

`func (o *BulkWritableModuleTypeRequest) GetRearImage() *os.File`

GetRearImage returns the RearImage field if non-nil, zero value otherwise.

### GetRearImageOk

`func (o *BulkWritableModuleTypeRequest) GetRearImageOk() (**os.File, bool)`

GetRearImageOk returns a tuple with the RearImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearImage

`func (o *BulkWritableModuleTypeRequest) SetRearImage(v *os.File)`

SetRearImage sets RearImage field to given value.

### HasRearImage

`func (o *BulkWritableModuleTypeRequest) HasRearImage() bool`

HasRearImage returns a boolean if a field has been set.

### GetModel

`func (o *BulkWritableModuleTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BulkWritableModuleTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BulkWritableModuleTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetPartNumber

`func (o *BulkWritableModuleTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BulkWritableModuleTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BulkWritableModuleTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BulkWritableModuleTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetComments

`func (o *BulkWritableModuleTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWritableModuleTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWritableModuleTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWritableModuleTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetManufacturer

`func (o *BulkWritableModuleTypeRequest) GetManufacturer() BulkWritableCableRequestStatus`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BulkWritableModuleTypeRequest) GetManufacturerOk() (*BulkWritableCableRequestStatus, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BulkWritableModuleTypeRequest) SetManufacturer(v BulkWritableCableRequestStatus)`

SetManufacturer sets Manufacturer field to given value.


### GetModuleFamily

`func (o *BulkWritableModuleTypeRequest) GetModuleFamily() ApprovalWorkflowUser`

GetModuleFamily returns the ModuleFamily field if non-nil, zero value otherwise.

### GetModuleFamilyOk

`func (o *BulkWritableModuleTypeRequest) GetModuleFamilyOk() (*ApprovalWorkflowUser, bool)`

GetModuleFamilyOk returns a tuple with the ModuleFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleFamily

`func (o *BulkWritableModuleTypeRequest) SetModuleFamily(v ApprovalWorkflowUser)`

SetModuleFamily sets ModuleFamily field to given value.

### HasModuleFamily

`func (o *BulkWritableModuleTypeRequest) HasModuleFamily() bool`

HasModuleFamily returns a boolean if a field has been set.

### SetModuleFamilyNil

`func (o *BulkWritableModuleTypeRequest) SetModuleFamilyNil(b bool)`

 SetModuleFamilyNil sets the value for ModuleFamily to be an explicit nil

### UnsetModuleFamily
`func (o *BulkWritableModuleTypeRequest) UnsetModuleFamily()`

UnsetModuleFamily ensures that no value is present for ModuleFamily, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableModuleTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableModuleTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableModuleTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableModuleTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableModuleTypeRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableModuleTypeRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableModuleTypeRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableModuleTypeRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableModuleTypeRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableModuleTypeRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableModuleTypeRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableModuleTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


