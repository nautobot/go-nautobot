# BulkWritableSoftwareImageFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewBulkWritableSoftwareImageFileRequest

`func NewBulkWritableSoftwareImageFileRequest(id string, imageFileName string, softwareVersion BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, ) *BulkWritableSoftwareImageFileRequest`

NewBulkWritableSoftwareImageFileRequest instantiates a new BulkWritableSoftwareImageFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableSoftwareImageFileRequestWithDefaults

`func NewBulkWritableSoftwareImageFileRequestWithDefaults() *BulkWritableSoftwareImageFileRequest`

NewBulkWritableSoftwareImageFileRequestWithDefaults instantiates a new BulkWritableSoftwareImageFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableSoftwareImageFileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableSoftwareImageFileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableSoftwareImageFileRequest) SetId(v string)`

SetId sets Id field to given value.


### GetImageFileName

`func (o *BulkWritableSoftwareImageFileRequest) GetImageFileName() string`

GetImageFileName returns the ImageFileName field if non-nil, zero value otherwise.

### GetImageFileNameOk

`func (o *BulkWritableSoftwareImageFileRequest) GetImageFileNameOk() (*string, bool)`

GetImageFileNameOk returns a tuple with the ImageFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileName

`func (o *BulkWritableSoftwareImageFileRequest) SetImageFileName(v string)`

SetImageFileName sets ImageFileName field to given value.


### GetImageFileChecksum

`func (o *BulkWritableSoftwareImageFileRequest) GetImageFileChecksum() string`

GetImageFileChecksum returns the ImageFileChecksum field if non-nil, zero value otherwise.

### GetImageFileChecksumOk

`func (o *BulkWritableSoftwareImageFileRequest) GetImageFileChecksumOk() (*string, bool)`

GetImageFileChecksumOk returns a tuple with the ImageFileChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileChecksum

`func (o *BulkWritableSoftwareImageFileRequest) SetImageFileChecksum(v string)`

SetImageFileChecksum sets ImageFileChecksum field to given value.

### HasImageFileChecksum

`func (o *BulkWritableSoftwareImageFileRequest) HasImageFileChecksum() bool`

HasImageFileChecksum returns a boolean if a field has been set.

### GetHashingAlgorithm

`func (o *BulkWritableSoftwareImageFileRequest) GetHashingAlgorithm() BulkWritableSoftwareImageFileRequestHashingAlgorithm`

GetHashingAlgorithm returns the HashingAlgorithm field if non-nil, zero value otherwise.

### GetHashingAlgorithmOk

`func (o *BulkWritableSoftwareImageFileRequest) GetHashingAlgorithmOk() (*BulkWritableSoftwareImageFileRequestHashingAlgorithm, bool)`

GetHashingAlgorithmOk returns a tuple with the HashingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashingAlgorithm

`func (o *BulkWritableSoftwareImageFileRequest) SetHashingAlgorithm(v BulkWritableSoftwareImageFileRequestHashingAlgorithm)`

SetHashingAlgorithm sets HashingAlgorithm field to given value.

### HasHashingAlgorithm

`func (o *BulkWritableSoftwareImageFileRequest) HasHashingAlgorithm() bool`

HasHashingAlgorithm returns a boolean if a field has been set.

### GetImageFileSize

`func (o *BulkWritableSoftwareImageFileRequest) GetImageFileSize() int64`

GetImageFileSize returns the ImageFileSize field if non-nil, zero value otherwise.

### GetImageFileSizeOk

`func (o *BulkWritableSoftwareImageFileRequest) GetImageFileSizeOk() (*int64, bool)`

GetImageFileSizeOk returns a tuple with the ImageFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileSize

`func (o *BulkWritableSoftwareImageFileRequest) SetImageFileSize(v int64)`

SetImageFileSize sets ImageFileSize field to given value.

### HasImageFileSize

`func (o *BulkWritableSoftwareImageFileRequest) HasImageFileSize() bool`

HasImageFileSize returns a boolean if a field has been set.

### SetImageFileSizeNil

`func (o *BulkWritableSoftwareImageFileRequest) SetImageFileSizeNil(b bool)`

 SetImageFileSizeNil sets the value for ImageFileSize to be an explicit nil

### UnsetImageFileSize
`func (o *BulkWritableSoftwareImageFileRequest) UnsetImageFileSize()`

UnsetImageFileSize ensures that no value is present for ImageFileSize, not even an explicit nil
### GetDownloadUrl

`func (o *BulkWritableSoftwareImageFileRequest) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *BulkWritableSoftwareImageFileRequest) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *BulkWritableSoftwareImageFileRequest) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *BulkWritableSoftwareImageFileRequest) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetDefaultImage

`func (o *BulkWritableSoftwareImageFileRequest) GetDefaultImage() bool`

GetDefaultImage returns the DefaultImage field if non-nil, zero value otherwise.

### GetDefaultImageOk

`func (o *BulkWritableSoftwareImageFileRequest) GetDefaultImageOk() (*bool, bool)`

GetDefaultImageOk returns a tuple with the DefaultImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImage

`func (o *BulkWritableSoftwareImageFileRequest) SetDefaultImage(v bool)`

SetDefaultImage sets DefaultImage field to given value.

### HasDefaultImage

`func (o *BulkWritableSoftwareImageFileRequest) HasDefaultImage() bool`

HasDefaultImage returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *BulkWritableSoftwareImageFileRequest) GetSoftwareVersion() BulkWritableCableRequestStatus`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *BulkWritableSoftwareImageFileRequest) GetSoftwareVersionOk() (*BulkWritableCableRequestStatus, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *BulkWritableSoftwareImageFileRequest) SetSoftwareVersion(v BulkWritableCableRequestStatus)`

SetSoftwareVersion sets SoftwareVersion field to given value.


### GetStatus

`func (o *BulkWritableSoftwareImageFileRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableSoftwareImageFileRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableSoftwareImageFileRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *BulkWritableSoftwareImageFileRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableSoftwareImageFileRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableSoftwareImageFileRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableSoftwareImageFileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableSoftwareImageFileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableSoftwareImageFileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableSoftwareImageFileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableSoftwareImageFileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableSoftwareImageFileRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableSoftwareImageFileRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableSoftwareImageFileRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableSoftwareImageFileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


