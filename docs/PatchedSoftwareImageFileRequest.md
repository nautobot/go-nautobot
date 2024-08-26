# PatchedSoftwareImageFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageFileName** | Pointer to **string** |  | [optional] 
**ImageFileChecksum** | Pointer to **string** |  | [optional] 
**HashingAlgorithm** | Pointer to [**BulkWritableSoftwareImageFileRequestHashingAlgorithm**](BulkWritableSoftwareImageFileRequestHashingAlgorithm.md) |  | [optional] 
**ImageFileSize** | Pointer to **NullableInt64** | Image file size in bytes | [optional] 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**DefaultImage** | Pointer to **bool** | Is the default image for this software version | [optional] 
**SoftwareVersion** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedSoftwareImageFileRequest

`func NewPatchedSoftwareImageFileRequest() *PatchedSoftwareImageFileRequest`

NewPatchedSoftwareImageFileRequest instantiates a new PatchedSoftwareImageFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSoftwareImageFileRequestWithDefaults

`func NewPatchedSoftwareImageFileRequestWithDefaults() *PatchedSoftwareImageFileRequest`

NewPatchedSoftwareImageFileRequestWithDefaults instantiates a new PatchedSoftwareImageFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageFileName

`func (o *PatchedSoftwareImageFileRequest) GetImageFileName() string`

GetImageFileName returns the ImageFileName field if non-nil, zero value otherwise.

### GetImageFileNameOk

`func (o *PatchedSoftwareImageFileRequest) GetImageFileNameOk() (*string, bool)`

GetImageFileNameOk returns a tuple with the ImageFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileName

`func (o *PatchedSoftwareImageFileRequest) SetImageFileName(v string)`

SetImageFileName sets ImageFileName field to given value.

### HasImageFileName

`func (o *PatchedSoftwareImageFileRequest) HasImageFileName() bool`

HasImageFileName returns a boolean if a field has been set.

### GetImageFileChecksum

`func (o *PatchedSoftwareImageFileRequest) GetImageFileChecksum() string`

GetImageFileChecksum returns the ImageFileChecksum field if non-nil, zero value otherwise.

### GetImageFileChecksumOk

`func (o *PatchedSoftwareImageFileRequest) GetImageFileChecksumOk() (*string, bool)`

GetImageFileChecksumOk returns a tuple with the ImageFileChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileChecksum

`func (o *PatchedSoftwareImageFileRequest) SetImageFileChecksum(v string)`

SetImageFileChecksum sets ImageFileChecksum field to given value.

### HasImageFileChecksum

`func (o *PatchedSoftwareImageFileRequest) HasImageFileChecksum() bool`

HasImageFileChecksum returns a boolean if a field has been set.

### GetHashingAlgorithm

`func (o *PatchedSoftwareImageFileRequest) GetHashingAlgorithm() BulkWritableSoftwareImageFileRequestHashingAlgorithm`

GetHashingAlgorithm returns the HashingAlgorithm field if non-nil, zero value otherwise.

### GetHashingAlgorithmOk

`func (o *PatchedSoftwareImageFileRequest) GetHashingAlgorithmOk() (*BulkWritableSoftwareImageFileRequestHashingAlgorithm, bool)`

GetHashingAlgorithmOk returns a tuple with the HashingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashingAlgorithm

`func (o *PatchedSoftwareImageFileRequest) SetHashingAlgorithm(v BulkWritableSoftwareImageFileRequestHashingAlgorithm)`

SetHashingAlgorithm sets HashingAlgorithm field to given value.

### HasHashingAlgorithm

`func (o *PatchedSoftwareImageFileRequest) HasHashingAlgorithm() bool`

HasHashingAlgorithm returns a boolean if a field has been set.

### GetImageFileSize

`func (o *PatchedSoftwareImageFileRequest) GetImageFileSize() int64`

GetImageFileSize returns the ImageFileSize field if non-nil, zero value otherwise.

### GetImageFileSizeOk

`func (o *PatchedSoftwareImageFileRequest) GetImageFileSizeOk() (*int64, bool)`

GetImageFileSizeOk returns a tuple with the ImageFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileSize

`func (o *PatchedSoftwareImageFileRequest) SetImageFileSize(v int64)`

SetImageFileSize sets ImageFileSize field to given value.

### HasImageFileSize

`func (o *PatchedSoftwareImageFileRequest) HasImageFileSize() bool`

HasImageFileSize returns a boolean if a field has been set.

### SetImageFileSizeNil

`func (o *PatchedSoftwareImageFileRequest) SetImageFileSizeNil(b bool)`

 SetImageFileSizeNil sets the value for ImageFileSize to be an explicit nil

### UnsetImageFileSize
`func (o *PatchedSoftwareImageFileRequest) UnsetImageFileSize()`

UnsetImageFileSize ensures that no value is present for ImageFileSize, not even an explicit nil
### GetDownloadUrl

`func (o *PatchedSoftwareImageFileRequest) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *PatchedSoftwareImageFileRequest) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *PatchedSoftwareImageFileRequest) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *PatchedSoftwareImageFileRequest) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetDefaultImage

`func (o *PatchedSoftwareImageFileRequest) GetDefaultImage() bool`

GetDefaultImage returns the DefaultImage field if non-nil, zero value otherwise.

### GetDefaultImageOk

`func (o *PatchedSoftwareImageFileRequest) GetDefaultImageOk() (*bool, bool)`

GetDefaultImageOk returns a tuple with the DefaultImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImage

`func (o *PatchedSoftwareImageFileRequest) SetDefaultImage(v bool)`

SetDefaultImage sets DefaultImage field to given value.

### HasDefaultImage

`func (o *PatchedSoftwareImageFileRequest) HasDefaultImage() bool`

HasDefaultImage returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *PatchedSoftwareImageFileRequest) GetSoftwareVersion() BulkWritableCableRequestStatus`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *PatchedSoftwareImageFileRequest) GetSoftwareVersionOk() (*BulkWritableCableRequestStatus, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *PatchedSoftwareImageFileRequest) SetSoftwareVersion(v BulkWritableCableRequestStatus)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *PatchedSoftwareImageFileRequest) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedSoftwareImageFileRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedSoftwareImageFileRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedSoftwareImageFileRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedSoftwareImageFileRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PatchedSoftwareImageFileRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedSoftwareImageFileRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedSoftwareImageFileRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedSoftwareImageFileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedSoftwareImageFileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedSoftwareImageFileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedSoftwareImageFileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedSoftwareImageFileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedSoftwareImageFileRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedSoftwareImageFileRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedSoftwareImageFileRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedSoftwareImageFileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


