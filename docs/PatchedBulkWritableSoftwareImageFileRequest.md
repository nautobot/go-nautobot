# PatchedBulkWritableSoftwareImageFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ImageFileName** | Pointer to **string** |  | [optional] 
**ImageFileChecksum** | Pointer to **string** |  | [optional] 
**HashingAlgorithm** | Pointer to [**BulkWritableSoftwareImageFileRequestHashingAlgorithm**](BulkWritableSoftwareImageFileRequestHashingAlgorithm.md) |  | [optional] 
**ImageFileSize** | Pointer to **NullableInt64** | Image file size in bytes | [optional] 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**DefaultImage** | Pointer to **bool** | Is the default image for this software version | [optional] 
**SoftwareVersion** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableSoftwareImageFileRequest

`func NewPatchedBulkWritableSoftwareImageFileRequest(id string, ) *PatchedBulkWritableSoftwareImageFileRequest`

NewPatchedBulkWritableSoftwareImageFileRequest instantiates a new PatchedBulkWritableSoftwareImageFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableSoftwareImageFileRequestWithDefaults

`func NewPatchedBulkWritableSoftwareImageFileRequestWithDefaults() *PatchedBulkWritableSoftwareImageFileRequest`

NewPatchedBulkWritableSoftwareImageFileRequestWithDefaults instantiates a new PatchedBulkWritableSoftwareImageFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetId(v string)`

SetId sets Id field to given value.


### GetImageFileName

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetImageFileName() string`

GetImageFileName returns the ImageFileName field if non-nil, zero value otherwise.

### GetImageFileNameOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetImageFileNameOk() (*string, bool)`

GetImageFileNameOk returns a tuple with the ImageFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileName

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetImageFileName(v string)`

SetImageFileName sets ImageFileName field to given value.

### HasImageFileName

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasImageFileName() bool`

HasImageFileName returns a boolean if a field has been set.

### GetImageFileChecksum

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetImageFileChecksum() string`

GetImageFileChecksum returns the ImageFileChecksum field if non-nil, zero value otherwise.

### GetImageFileChecksumOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetImageFileChecksumOk() (*string, bool)`

GetImageFileChecksumOk returns a tuple with the ImageFileChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileChecksum

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetImageFileChecksum(v string)`

SetImageFileChecksum sets ImageFileChecksum field to given value.

### HasImageFileChecksum

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasImageFileChecksum() bool`

HasImageFileChecksum returns a boolean if a field has been set.

### GetHashingAlgorithm

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetHashingAlgorithm() BulkWritableSoftwareImageFileRequestHashingAlgorithm`

GetHashingAlgorithm returns the HashingAlgorithm field if non-nil, zero value otherwise.

### GetHashingAlgorithmOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetHashingAlgorithmOk() (*BulkWritableSoftwareImageFileRequestHashingAlgorithm, bool)`

GetHashingAlgorithmOk returns a tuple with the HashingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashingAlgorithm

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetHashingAlgorithm(v BulkWritableSoftwareImageFileRequestHashingAlgorithm)`

SetHashingAlgorithm sets HashingAlgorithm field to given value.

### HasHashingAlgorithm

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasHashingAlgorithm() bool`

HasHashingAlgorithm returns a boolean if a field has been set.

### GetImageFileSize

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetImageFileSize() int64`

GetImageFileSize returns the ImageFileSize field if non-nil, zero value otherwise.

### GetImageFileSizeOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetImageFileSizeOk() (*int64, bool)`

GetImageFileSizeOk returns a tuple with the ImageFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileSize

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetImageFileSize(v int64)`

SetImageFileSize sets ImageFileSize field to given value.

### HasImageFileSize

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasImageFileSize() bool`

HasImageFileSize returns a boolean if a field has been set.

### SetImageFileSizeNil

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetImageFileSizeNil(b bool)`

 SetImageFileSizeNil sets the value for ImageFileSize to be an explicit nil

### UnsetImageFileSize
`func (o *PatchedBulkWritableSoftwareImageFileRequest) UnsetImageFileSize()`

UnsetImageFileSize ensures that no value is present for ImageFileSize, not even an explicit nil
### GetDownloadUrl

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetDefaultImage

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetDefaultImage() bool`

GetDefaultImage returns the DefaultImage field if non-nil, zero value otherwise.

### GetDefaultImageOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetDefaultImageOk() (*bool, bool)`

GetDefaultImageOk returns a tuple with the DefaultImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImage

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetDefaultImage(v bool)`

SetDefaultImage sets DefaultImage field to given value.

### HasDefaultImage

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasDefaultImage() bool`

HasDefaultImage returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetSoftwareVersion() BulkWritableCableRequestStatus`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetSoftwareVersionOk() (*BulkWritableCableRequestStatus, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetSoftwareVersion(v BulkWritableCableRequestStatus)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableSoftwareImageFileRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableSoftwareImageFileRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableSoftwareImageFileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


