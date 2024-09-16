# SoftwareImageFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ImageFileName** | **string** |  | 
**ImageFileChecksum** | Pointer to **string** |  | [optional] 
**HashingAlgorithm** | Pointer to [**BulkWritableSoftwareImageFileRequestHashingAlgorithm**](BulkWritableSoftwareImageFileRequestHashingAlgorithm.md) |  | [optional] 
**ImageFileSize** | Pointer to **NullableInt64** | Image file size in bytes | [optional] 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**DefaultImage** | Pointer to **bool** | Is the default image for this software version | [optional] 
**SoftwareVersion** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewSoftwareImageFile

`func NewSoftwareImageFile(id string, objectType string, display string, url string, naturalSlug string, imageFileName string, softwareVersion BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *SoftwareImageFile`

NewSoftwareImageFile instantiates a new SoftwareImageFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareImageFileWithDefaults

`func NewSoftwareImageFileWithDefaults() *SoftwareImageFile`

NewSoftwareImageFileWithDefaults instantiates a new SoftwareImageFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SoftwareImageFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SoftwareImageFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SoftwareImageFile) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *SoftwareImageFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareImageFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareImageFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *SoftwareImageFile) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SoftwareImageFile) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SoftwareImageFile) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *SoftwareImageFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SoftwareImageFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SoftwareImageFile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *SoftwareImageFile) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *SoftwareImageFile) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *SoftwareImageFile) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetImageFileName

`func (o *SoftwareImageFile) GetImageFileName() string`

GetImageFileName returns the ImageFileName field if non-nil, zero value otherwise.

### GetImageFileNameOk

`func (o *SoftwareImageFile) GetImageFileNameOk() (*string, bool)`

GetImageFileNameOk returns a tuple with the ImageFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileName

`func (o *SoftwareImageFile) SetImageFileName(v string)`

SetImageFileName sets ImageFileName field to given value.


### GetImageFileChecksum

`func (o *SoftwareImageFile) GetImageFileChecksum() string`

GetImageFileChecksum returns the ImageFileChecksum field if non-nil, zero value otherwise.

### GetImageFileChecksumOk

`func (o *SoftwareImageFile) GetImageFileChecksumOk() (*string, bool)`

GetImageFileChecksumOk returns a tuple with the ImageFileChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileChecksum

`func (o *SoftwareImageFile) SetImageFileChecksum(v string)`

SetImageFileChecksum sets ImageFileChecksum field to given value.

### HasImageFileChecksum

`func (o *SoftwareImageFile) HasImageFileChecksum() bool`

HasImageFileChecksum returns a boolean if a field has been set.

### GetHashingAlgorithm

`func (o *SoftwareImageFile) GetHashingAlgorithm() BulkWritableSoftwareImageFileRequestHashingAlgorithm`

GetHashingAlgorithm returns the HashingAlgorithm field if non-nil, zero value otherwise.

### GetHashingAlgorithmOk

`func (o *SoftwareImageFile) GetHashingAlgorithmOk() (*BulkWritableSoftwareImageFileRequestHashingAlgorithm, bool)`

GetHashingAlgorithmOk returns a tuple with the HashingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashingAlgorithm

`func (o *SoftwareImageFile) SetHashingAlgorithm(v BulkWritableSoftwareImageFileRequestHashingAlgorithm)`

SetHashingAlgorithm sets HashingAlgorithm field to given value.

### HasHashingAlgorithm

`func (o *SoftwareImageFile) HasHashingAlgorithm() bool`

HasHashingAlgorithm returns a boolean if a field has been set.

### GetImageFileSize

`func (o *SoftwareImageFile) GetImageFileSize() int64`

GetImageFileSize returns the ImageFileSize field if non-nil, zero value otherwise.

### GetImageFileSizeOk

`func (o *SoftwareImageFile) GetImageFileSizeOk() (*int64, bool)`

GetImageFileSizeOk returns a tuple with the ImageFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFileSize

`func (o *SoftwareImageFile) SetImageFileSize(v int64)`

SetImageFileSize sets ImageFileSize field to given value.

### HasImageFileSize

`func (o *SoftwareImageFile) HasImageFileSize() bool`

HasImageFileSize returns a boolean if a field has been set.

### SetImageFileSizeNil

`func (o *SoftwareImageFile) SetImageFileSizeNil(b bool)`

 SetImageFileSizeNil sets the value for ImageFileSize to be an explicit nil

### UnsetImageFileSize
`func (o *SoftwareImageFile) UnsetImageFileSize()`

UnsetImageFileSize ensures that no value is present for ImageFileSize, not even an explicit nil
### GetDownloadUrl

`func (o *SoftwareImageFile) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *SoftwareImageFile) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *SoftwareImageFile) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *SoftwareImageFile) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetDefaultImage

`func (o *SoftwareImageFile) GetDefaultImage() bool`

GetDefaultImage returns the DefaultImage field if non-nil, zero value otherwise.

### GetDefaultImageOk

`func (o *SoftwareImageFile) GetDefaultImageOk() (*bool, bool)`

GetDefaultImageOk returns a tuple with the DefaultImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImage

`func (o *SoftwareImageFile) SetDefaultImage(v bool)`

SetDefaultImage sets DefaultImage field to given value.

### HasDefaultImage

`func (o *SoftwareImageFile) HasDefaultImage() bool`

HasDefaultImage returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *SoftwareImageFile) GetSoftwareVersion() BulkWritableCableRequestStatus`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *SoftwareImageFile) GetSoftwareVersionOk() (*BulkWritableCableRequestStatus, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *SoftwareImageFile) SetSoftwareVersion(v BulkWritableCableRequestStatus)`

SetSoftwareVersion sets SoftwareVersion field to given value.


### GetStatus

`func (o *SoftwareImageFile) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SoftwareImageFile) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SoftwareImageFile) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *SoftwareImageFile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SoftwareImageFile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SoftwareImageFile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *SoftwareImageFile) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SoftwareImageFile) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *SoftwareImageFile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SoftwareImageFile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SoftwareImageFile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *SoftwareImageFile) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *SoftwareImageFile) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *SoftwareImageFile) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *SoftwareImageFile) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *SoftwareImageFile) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *SoftwareImageFile) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SoftwareImageFile) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SoftwareImageFile) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SoftwareImageFile) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *SoftwareImageFile) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SoftwareImageFile) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SoftwareImageFile) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SoftwareImageFile) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


