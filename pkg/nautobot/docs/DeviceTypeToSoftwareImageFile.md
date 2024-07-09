# DeviceTypeToSoftwareImageFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**DeviceType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**SoftwareImageFile** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewDeviceTypeToSoftwareImageFile

`func NewDeviceTypeToSoftwareImageFile(id string, objectType string, display string, url string, naturalSlug string, deviceType BulkWritableCableRequestStatus, softwareImageFile BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, ) *DeviceTypeToSoftwareImageFile`

NewDeviceTypeToSoftwareImageFile instantiates a new DeviceTypeToSoftwareImageFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTypeToSoftwareImageFileWithDefaults

`func NewDeviceTypeToSoftwareImageFileWithDefaults() *DeviceTypeToSoftwareImageFile`

NewDeviceTypeToSoftwareImageFileWithDefaults instantiates a new DeviceTypeToSoftwareImageFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceTypeToSoftwareImageFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceTypeToSoftwareImageFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceTypeToSoftwareImageFile) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *DeviceTypeToSoftwareImageFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DeviceTypeToSoftwareImageFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DeviceTypeToSoftwareImageFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *DeviceTypeToSoftwareImageFile) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DeviceTypeToSoftwareImageFile) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DeviceTypeToSoftwareImageFile) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *DeviceTypeToSoftwareImageFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeviceTypeToSoftwareImageFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeviceTypeToSoftwareImageFile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *DeviceTypeToSoftwareImageFile) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *DeviceTypeToSoftwareImageFile) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *DeviceTypeToSoftwareImageFile) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetDeviceType

`func (o *DeviceTypeToSoftwareImageFile) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceTypeToSoftwareImageFile) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceTypeToSoftwareImageFile) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.


### GetSoftwareImageFile

`func (o *DeviceTypeToSoftwareImageFile) GetSoftwareImageFile() BulkWritableCableRequestStatus`

GetSoftwareImageFile returns the SoftwareImageFile field if non-nil, zero value otherwise.

### GetSoftwareImageFileOk

`func (o *DeviceTypeToSoftwareImageFile) GetSoftwareImageFileOk() (*BulkWritableCableRequestStatus, bool)`

GetSoftwareImageFileOk returns a tuple with the SoftwareImageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageFile

`func (o *DeviceTypeToSoftwareImageFile) SetSoftwareImageFile(v BulkWritableCableRequestStatus)`

SetSoftwareImageFile sets SoftwareImageFile field to given value.


### GetCreated

`func (o *DeviceTypeToSoftwareImageFile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceTypeToSoftwareImageFile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceTypeToSoftwareImageFile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *DeviceTypeToSoftwareImageFile) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *DeviceTypeToSoftwareImageFile) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *DeviceTypeToSoftwareImageFile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceTypeToSoftwareImageFile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceTypeToSoftwareImageFile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *DeviceTypeToSoftwareImageFile) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *DeviceTypeToSoftwareImageFile) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


