# SoftwareImageFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageFileName** | **string** |  | 
**ImageFileChecksum** | Pointer to **string** |  | [optional] 
**HashingAlgorithm** | Pointer to [**BulkWritableSoftwareImageFileRequestHashingAlgorithm**](BulkWritableSoftwareImageFileRequestHashingAlgorithm.md) |  | [optional] 
**ImageFileSize** | Pointer to **NullableInt64** | Image file size in bytes | [optional] 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**DefaultImage** | Pointer to **bool** | Is the default image for this software version | [optional] 
**SoftwareVersion** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewSoftwareImageFileRequest

`func NewSoftwareImageFileRequest(imageFileName string, softwareVersion BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, ) *SoftwareImageFileRequest`

NewSoftwareImageFileRequest instantiates a new SoftwareImageFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareImageFileRequestWithDefaults

`func NewSoftwareImageFileRequestWithDefaults() *SoftwareImageFileRequest`

NewSoftwareImageFileRequestWithDefaults instantiates a new SoftwareImageFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageFileName

`func (o *SoftwareImageFileRequest) GetImageFileName() string`

GetImageFileName returns the ImageFileName field if non-nil, zero value otherwise.

### GetImageFileNameOk

`func (o *SoftwareImageFileRequest) GetImageFileNameOk() (*string, bool)`

GetImageFileNameOk returns a tuple with the ImageFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileName

`func (o *SoftwareImageFileRequest) SetImageFileName(v string)`

SetImageFileName sets ImageFileName field to given value.


### GetImageFileChecksum

`func (o *SoftwareImageFileRequest) GetImageFileChecksum() string`

GetImageFileChecksum returns the ImageFileChecksum field if non-nil, zero value otherwise.

### GetImageFileChecksumOk

`func (o *SoftwareImageFileRequest) GetImageFileChecksumOk() (*string, bool)`

GetImageFileChecksumOk returns a tuple with the ImageFileChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileChecksum

`func (o *SoftwareImageFileRequest) SetImageFileChecksum(v string)`

SetImageFileChecksum sets ImageFileChecksum field to given value.

### HasImageFileChecksum

`func (o *SoftwareImageFileRequest) HasImageFileChecksum() bool`

HasImageFileChecksum returns a boolean if a field has been set.

### GetHashingAlgorithm

`func (o *SoftwareImageFileRequest) GetHashingAlgorithm() BulkWritableSoftwareImageFileRequestHashingAlgorithm`

GetHashingAlgorithm returns the HashingAlgorithm field if non-nil, zero value otherwise.

### GetHashingAlgorithmOk

`func (o *SoftwareImageFileRequest) GetHashingAlgorithmOk() (*BulkWritableSoftwareImageFileRequestHashingAlgorithm, bool)`

GetHashingAlgorithmOk returns a tuple with the HashingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashingAlgorithm

`func (o *SoftwareImageFileRequest) SetHashingAlgorithm(v BulkWritableSoftwareImageFileRequestHashingAlgorithm)`

SetHashingAlgorithm sets HashingAlgorithm field to given value.

### HasHashingAlgorithm

`func (o *SoftwareImageFileRequest) HasHashingAlgorithm() bool`

HasHashingAlgorithm returns a boolean if a field has been set.

### GetImageFileSize

`func (o *SoftwareImageFileRequest) GetImageFileSize() int64`

GetImageFileSize returns the ImageFileSize field if non-nil, zero value otherwise.

### GetImageFileSizeOk

`func (o *SoftwareImageFileRequest) GetImageFileSizeOk() (*int64, bool)`

GetImageFileSizeOk returns a tuple with the ImageFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileSize

`func (o *SoftwareImageFileRequest) SetImageFileSize(v int64)`

SetImageFileSize sets ImageFileSize field to given value.

### HasImageFileSize

`func (o *SoftwareImageFileRequest) HasImageFileSize() bool`

HasImageFileSize returns a boolean if a field has been set.

### SetImageFileSizeNil

`func (o *SoftwareImageFileRequest) SetImageFileSizeNil(b bool)`

 SetImageFileSizeNil sets the value for ImageFileSize to be an explicit nil

### UnsetImageFileSize
`func (o *SoftwareImageFileRequest) UnsetImageFileSize()`

UnsetImageFileSize ensures that no value is present for ImageFileSize, not even an explicit nil
### GetDownloadUrl

`func (o *SoftwareImageFileRequest) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *SoftwareImageFileRequest) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *SoftwareImageFileRequest) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *SoftwareImageFileRequest) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetDefaultImage

`func (o *SoftwareImageFileRequest) GetDefaultImage() bool`

GetDefaultImage returns the DefaultImage field if non-nil, zero value otherwise.

### GetDefaultImageOk

`func (o *SoftwareImageFileRequest) GetDefaultImageOk() (*bool, bool)`

GetDefaultImageOk returns a tuple with the DefaultImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImage

`func (o *SoftwareImageFileRequest) SetDefaultImage(v bool)`

SetDefaultImage sets DefaultImage field to given value.

### HasDefaultImage

`func (o *SoftwareImageFileRequest) HasDefaultImage() bool`

HasDefaultImage returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *SoftwareImageFileRequest) GetSoftwareVersion() BulkWritableCableRequestStatus`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *SoftwareImageFileRequest) GetSoftwareVersionOk() (*BulkWritableCableRequestStatus, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *SoftwareImageFileRequest) SetSoftwareVersion(v BulkWritableCableRequestStatus)`

SetSoftwareVersion sets SoftwareVersion field to given value.


### GetStatus

`func (o *SoftwareImageFileRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SoftwareImageFileRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SoftwareImageFileRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *SoftwareImageFileRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SoftwareImageFileRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SoftwareImageFileRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SoftwareImageFileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *SoftwareImageFileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SoftwareImageFileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SoftwareImageFileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SoftwareImageFileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *SoftwareImageFileRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SoftwareImageFileRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SoftwareImageFileRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SoftwareImageFileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


