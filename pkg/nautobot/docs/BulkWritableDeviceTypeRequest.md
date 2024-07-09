# BulkWritableDeviceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SubdeviceRole** | Pointer to [**SubdeviceRoleEnum**](SubdeviceRoleEnum.md) |  | [optional] 
**FrontImage** | Pointer to **Nullable*os.File** |  | [optional] 
**RearImage** | Pointer to **Nullable*os.File** |  | [optional] 
**Model** | **string** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**UHeight** | Pointer to **int32** |  | [optional] 
**IsFullDepth** | Pointer to **bool** | Device consumes both front and rear rack faces | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Manufacturer** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**DeviceFamily** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableDeviceTypeRequest

`func NewBulkWritableDeviceTypeRequest(id string, model string, manufacturer BulkWritableCableRequestStatus, ) *BulkWritableDeviceTypeRequest`

NewBulkWritableDeviceTypeRequest instantiates a new BulkWritableDeviceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableDeviceTypeRequestWithDefaults

`func NewBulkWritableDeviceTypeRequestWithDefaults() *BulkWritableDeviceTypeRequest`

NewBulkWritableDeviceTypeRequestWithDefaults instantiates a new BulkWritableDeviceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableDeviceTypeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableDeviceTypeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableDeviceTypeRequest) SetId(v string)`

SetId sets Id field to given value.


### GetSubdeviceRole

`func (o *BulkWritableDeviceTypeRequest) GetSubdeviceRole() SubdeviceRoleEnum`

GetSubdeviceRole returns the SubdeviceRole field if non-nil, zero value otherwise.

### GetSubdeviceRoleOk

`func (o *BulkWritableDeviceTypeRequest) GetSubdeviceRoleOk() (*SubdeviceRoleEnum, bool)`

GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdeviceRole

`func (o *BulkWritableDeviceTypeRequest) SetSubdeviceRole(v SubdeviceRoleEnum)`

SetSubdeviceRole sets SubdeviceRole field to given value.

### HasSubdeviceRole

`func (o *BulkWritableDeviceTypeRequest) HasSubdeviceRole() bool`

HasSubdeviceRole returns a boolean if a field has been set.

### GetFrontImage

`func (o *BulkWritableDeviceTypeRequest) GetFrontImage() *os.File`

GetFrontImage returns the FrontImage field if non-nil, zero value otherwise.

### GetFrontImageOk

`func (o *BulkWritableDeviceTypeRequest) GetFrontImageOk() (**os.File, bool)`

GetFrontImageOk returns a tuple with the FrontImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImage

`func (o *BulkWritableDeviceTypeRequest) SetFrontImage(v *os.File)`

SetFrontImage sets FrontImage field to given value.

### HasFrontImage

`func (o *BulkWritableDeviceTypeRequest) HasFrontImage() bool`

HasFrontImage returns a boolean if a field has been set.

### SetFrontImageNil

`func (o *BulkWritableDeviceTypeRequest) SetFrontImageNil(b bool)`

 SetFrontImageNil sets the value for FrontImage to be an explicit nil

### UnsetFrontImage
`func (o *BulkWritableDeviceTypeRequest) UnsetFrontImage()`

UnsetFrontImage ensures that no value is present for FrontImage, not even an explicit nil
### GetRearImage

`func (o *BulkWritableDeviceTypeRequest) GetRearImage() *os.File`

GetRearImage returns the RearImage field if non-nil, zero value otherwise.

### GetRearImageOk

`func (o *BulkWritableDeviceTypeRequest) GetRearImageOk() (**os.File, bool)`

GetRearImageOk returns a tuple with the RearImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearImage

`func (o *BulkWritableDeviceTypeRequest) SetRearImage(v *os.File)`

SetRearImage sets RearImage field to given value.

### HasRearImage

`func (o *BulkWritableDeviceTypeRequest) HasRearImage() bool`

HasRearImage returns a boolean if a field has been set.

### SetRearImageNil

`func (o *BulkWritableDeviceTypeRequest) SetRearImageNil(b bool)`

 SetRearImageNil sets the value for RearImage to be an explicit nil

### UnsetRearImage
`func (o *BulkWritableDeviceTypeRequest) UnsetRearImage()`

UnsetRearImage ensures that no value is present for RearImage, not even an explicit nil
### GetModel

`func (o *BulkWritableDeviceTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BulkWritableDeviceTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BulkWritableDeviceTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetPartNumber

`func (o *BulkWritableDeviceTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BulkWritableDeviceTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BulkWritableDeviceTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BulkWritableDeviceTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetUHeight

`func (o *BulkWritableDeviceTypeRequest) GetUHeight() int32`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *BulkWritableDeviceTypeRequest) GetUHeightOk() (*int32, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *BulkWritableDeviceTypeRequest) SetUHeight(v int32)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *BulkWritableDeviceTypeRequest) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetIsFullDepth

`func (o *BulkWritableDeviceTypeRequest) GetIsFullDepth() bool`

GetIsFullDepth returns the IsFullDepth field if non-nil, zero value otherwise.

### GetIsFullDepthOk

`func (o *BulkWritableDeviceTypeRequest) GetIsFullDepthOk() (*bool, bool)`

GetIsFullDepthOk returns a tuple with the IsFullDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullDepth

`func (o *BulkWritableDeviceTypeRequest) SetIsFullDepth(v bool)`

SetIsFullDepth sets IsFullDepth field to given value.

### HasIsFullDepth

`func (o *BulkWritableDeviceTypeRequest) HasIsFullDepth() bool`

HasIsFullDepth returns a boolean if a field has been set.

### GetComments

`func (o *BulkWritableDeviceTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWritableDeviceTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWritableDeviceTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWritableDeviceTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetManufacturer

`func (o *BulkWritableDeviceTypeRequest) GetManufacturer() BulkWritableCableRequestStatus`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BulkWritableDeviceTypeRequest) GetManufacturerOk() (*BulkWritableCableRequestStatus, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BulkWritableDeviceTypeRequest) SetManufacturer(v BulkWritableCableRequestStatus)`

SetManufacturer sets Manufacturer field to given value.


### GetDeviceFamily

`func (o *BulkWritableDeviceTypeRequest) GetDeviceFamily() BulkWritableCircuitRequestTenant`

GetDeviceFamily returns the DeviceFamily field if non-nil, zero value otherwise.

### GetDeviceFamilyOk

`func (o *BulkWritableDeviceTypeRequest) GetDeviceFamilyOk() (*BulkWritableCircuitRequestTenant, bool)`

GetDeviceFamilyOk returns a tuple with the DeviceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFamily

`func (o *BulkWritableDeviceTypeRequest) SetDeviceFamily(v BulkWritableCircuitRequestTenant)`

SetDeviceFamily sets DeviceFamily field to given value.

### HasDeviceFamily

`func (o *BulkWritableDeviceTypeRequest) HasDeviceFamily() bool`

HasDeviceFamily returns a boolean if a field has been set.

### SetDeviceFamilyNil

`func (o *BulkWritableDeviceTypeRequest) SetDeviceFamilyNil(b bool)`

 SetDeviceFamilyNil sets the value for DeviceFamily to be an explicit nil

### UnsetDeviceFamily
`func (o *BulkWritableDeviceTypeRequest) UnsetDeviceFamily()`

UnsetDeviceFamily ensures that no value is present for DeviceFamily, not even an explicit nil
### GetTags

`func (o *BulkWritableDeviceTypeRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableDeviceTypeRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableDeviceTypeRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableDeviceTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableDeviceTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableDeviceTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableDeviceTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableDeviceTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableDeviceTypeRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableDeviceTypeRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableDeviceTypeRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableDeviceTypeRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


