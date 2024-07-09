# WritableDeviceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrontImage** | Pointer to **Nullable*os.File** |  | [optional] 
**RearImage** | Pointer to **Nullable*os.File** |  | [optional] 
**Model** | **string** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**UHeight** | Pointer to **int32** |  | [optional] 
**IsFullDepth** | Pointer to **bool** | Device consumes both front and rear rack faces | [optional] 
**SubdeviceRole** | Pointer to [**ParentChildStatus**](ParentChildStatus.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Manufacturer** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**DeviceFamily** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewWritableDeviceTypeRequest

`func NewWritableDeviceTypeRequest(model string, manufacturer BulkWritableCableRequestStatus, ) *WritableDeviceTypeRequest`

NewWritableDeviceTypeRequest instantiates a new WritableDeviceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableDeviceTypeRequestWithDefaults

`func NewWritableDeviceTypeRequestWithDefaults() *WritableDeviceTypeRequest`

NewWritableDeviceTypeRequestWithDefaults instantiates a new WritableDeviceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrontImage

`func (o *WritableDeviceTypeRequest) GetFrontImage() *os.File`

GetFrontImage returns the FrontImage field if non-nil, zero value otherwise.

### GetFrontImageOk

`func (o *WritableDeviceTypeRequest) GetFrontImageOk() (**os.File, bool)`

GetFrontImageOk returns a tuple with the FrontImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImage

`func (o *WritableDeviceTypeRequest) SetFrontImage(v *os.File)`

SetFrontImage sets FrontImage field to given value.

### HasFrontImage

`func (o *WritableDeviceTypeRequest) HasFrontImage() bool`

HasFrontImage returns a boolean if a field has been set.

### SetFrontImageNil

`func (o *WritableDeviceTypeRequest) SetFrontImageNil(b bool)`

 SetFrontImageNil sets the value for FrontImage to be an explicit nil

### UnsetFrontImage
`func (o *WritableDeviceTypeRequest) UnsetFrontImage()`

UnsetFrontImage ensures that no value is present for FrontImage, not even an explicit nil
### GetRearImage

`func (o *WritableDeviceTypeRequest) GetRearImage() *os.File`

GetRearImage returns the RearImage field if non-nil, zero value otherwise.

### GetRearImageOk

`func (o *WritableDeviceTypeRequest) GetRearImageOk() (**os.File, bool)`

GetRearImageOk returns a tuple with the RearImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearImage

`func (o *WritableDeviceTypeRequest) SetRearImage(v *os.File)`

SetRearImage sets RearImage field to given value.

### HasRearImage

`func (o *WritableDeviceTypeRequest) HasRearImage() bool`

HasRearImage returns a boolean if a field has been set.

### SetRearImageNil

`func (o *WritableDeviceTypeRequest) SetRearImageNil(b bool)`

 SetRearImageNil sets the value for RearImage to be an explicit nil

### UnsetRearImage
`func (o *WritableDeviceTypeRequest) UnsetRearImage()`

UnsetRearImage ensures that no value is present for RearImage, not even an explicit nil
### GetModel

`func (o *WritableDeviceTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *WritableDeviceTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *WritableDeviceTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetPartNumber

`func (o *WritableDeviceTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *WritableDeviceTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *WritableDeviceTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *WritableDeviceTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetUHeight

`func (o *WritableDeviceTypeRequest) GetUHeight() int32`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *WritableDeviceTypeRequest) GetUHeightOk() (*int32, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *WritableDeviceTypeRequest) SetUHeight(v int32)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *WritableDeviceTypeRequest) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetIsFullDepth

`func (o *WritableDeviceTypeRequest) GetIsFullDepth() bool`

GetIsFullDepth returns the IsFullDepth field if non-nil, zero value otherwise.

### GetIsFullDepthOk

`func (o *WritableDeviceTypeRequest) GetIsFullDepthOk() (*bool, bool)`

GetIsFullDepthOk returns a tuple with the IsFullDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullDepth

`func (o *WritableDeviceTypeRequest) SetIsFullDepth(v bool)`

SetIsFullDepth sets IsFullDepth field to given value.

### HasIsFullDepth

`func (o *WritableDeviceTypeRequest) HasIsFullDepth() bool`

HasIsFullDepth returns a boolean if a field has been set.

### GetSubdeviceRole

`func (o *WritableDeviceTypeRequest) GetSubdeviceRole() ParentChildStatus`

GetSubdeviceRole returns the SubdeviceRole field if non-nil, zero value otherwise.

### GetSubdeviceRoleOk

`func (o *WritableDeviceTypeRequest) GetSubdeviceRoleOk() (*ParentChildStatus, bool)`

GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdeviceRole

`func (o *WritableDeviceTypeRequest) SetSubdeviceRole(v ParentChildStatus)`

SetSubdeviceRole sets SubdeviceRole field to given value.

### HasSubdeviceRole

`func (o *WritableDeviceTypeRequest) HasSubdeviceRole() bool`

HasSubdeviceRole returns a boolean if a field has been set.

### GetComments

`func (o *WritableDeviceTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableDeviceTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableDeviceTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableDeviceTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetManufacturer

`func (o *WritableDeviceTypeRequest) GetManufacturer() BulkWritableCableRequestStatus`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *WritableDeviceTypeRequest) GetManufacturerOk() (*BulkWritableCableRequestStatus, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *WritableDeviceTypeRequest) SetManufacturer(v BulkWritableCableRequestStatus)`

SetManufacturer sets Manufacturer field to given value.


### GetDeviceFamily

`func (o *WritableDeviceTypeRequest) GetDeviceFamily() BulkWritableCircuitRequestTenant`

GetDeviceFamily returns the DeviceFamily field if non-nil, zero value otherwise.

### GetDeviceFamilyOk

`func (o *WritableDeviceTypeRequest) GetDeviceFamilyOk() (*BulkWritableCircuitRequestTenant, bool)`

GetDeviceFamilyOk returns a tuple with the DeviceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFamily

`func (o *WritableDeviceTypeRequest) SetDeviceFamily(v BulkWritableCircuitRequestTenant)`

SetDeviceFamily sets DeviceFamily field to given value.

### HasDeviceFamily

`func (o *WritableDeviceTypeRequest) HasDeviceFamily() bool`

HasDeviceFamily returns a boolean if a field has been set.

### SetDeviceFamilyNil

`func (o *WritableDeviceTypeRequest) SetDeviceFamilyNil(b bool)`

 SetDeviceFamilyNil sets the value for DeviceFamily to be an explicit nil

### UnsetDeviceFamily
`func (o *WritableDeviceTypeRequest) UnsetDeviceFamily()`

UnsetDeviceFamily ensures that no value is present for DeviceFamily, not even an explicit nil
### GetTags

`func (o *WritableDeviceTypeRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableDeviceTypeRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableDeviceTypeRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableDeviceTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableDeviceTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableDeviceTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableDeviceTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableDeviceTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritableDeviceTypeRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritableDeviceTypeRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritableDeviceTypeRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritableDeviceTypeRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


