# PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DeviceType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**SoftwareImageFile** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableDeviceTypeToSoftwareImageFileRequest

`func NewPatchedBulkWritableDeviceTypeToSoftwareImageFileRequest(id string, ) *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest`

NewPatchedBulkWritableDeviceTypeToSoftwareImageFileRequest instantiates a new PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableDeviceTypeToSoftwareImageFileRequestWithDefaults

`func NewPatchedBulkWritableDeviceTypeToSoftwareImageFileRequestWithDefaults() *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest`

NewPatchedBulkWritableDeviceTypeToSoftwareImageFileRequestWithDefaults instantiates a new PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceType

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetSoftwareImageFile

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetSoftwareImageFile() BulkWritableCableRequestStatus`

GetSoftwareImageFile returns the SoftwareImageFile field if non-nil, zero value otherwise.

### GetSoftwareImageFileOk

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) GetSoftwareImageFileOk() (*BulkWritableCableRequestStatus, bool)`

GetSoftwareImageFileOk returns a tuple with the SoftwareImageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageFile

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) SetSoftwareImageFile(v BulkWritableCableRequestStatus)`

SetSoftwareImageFile sets SoftwareImageFile field to given value.

### HasSoftwareImageFile

`func (o *PatchedBulkWritableDeviceTypeToSoftwareImageFileRequest) HasSoftwareImageFile() bool`

HasSoftwareImageFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


