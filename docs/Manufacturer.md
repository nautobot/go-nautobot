# Manufacturer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**CloudAccountCount** | Pointer to **int32** |  | [optional] [readonly] 
**DeviceTypeCount** | **int32** |  | [readonly] 
**InventoryItemCount** | **int32** |  | [readonly] 
**PlatformCount** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewManufacturer

`func NewManufacturer(id string, objectType string, display string, url string, naturalSlug string, deviceTypeCount int32, inventoryItemCount int32, platformCount int32, name string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Manufacturer`

NewManufacturer instantiates a new Manufacturer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManufacturerWithDefaults

`func NewManufacturerWithDefaults() *Manufacturer`

NewManufacturerWithDefaults instantiates a new Manufacturer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Manufacturer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Manufacturer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Manufacturer) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Manufacturer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Manufacturer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Manufacturer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Manufacturer) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Manufacturer) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Manufacturer) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Manufacturer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Manufacturer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Manufacturer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Manufacturer) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Manufacturer) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Manufacturer) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCloudAccountCount

`func (o *Manufacturer) GetCloudAccountCount() int32`

GetCloudAccountCount returns the CloudAccountCount field if non-nil, zero value otherwise.

### GetCloudAccountCountOk

`func (o *Manufacturer) GetCloudAccountCountOk() (*int32, bool)`

GetCloudAccountCountOk returns a tuple with the CloudAccountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountCount

`func (o *Manufacturer) SetCloudAccountCount(v int32)`

SetCloudAccountCount sets CloudAccountCount field to given value.

### HasCloudAccountCount

`func (o *Manufacturer) HasCloudAccountCount() bool`

HasCloudAccountCount returns a boolean if a field has been set.

### GetDeviceTypeCount

`func (o *Manufacturer) GetDeviceTypeCount() int32`

GetDeviceTypeCount returns the DeviceTypeCount field if non-nil, zero value otherwise.

### GetDeviceTypeCountOk

`func (o *Manufacturer) GetDeviceTypeCountOk() (*int32, bool)`

GetDeviceTypeCountOk returns a tuple with the DeviceTypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeCount

`func (o *Manufacturer) SetDeviceTypeCount(v int32)`

SetDeviceTypeCount sets DeviceTypeCount field to given value.


### GetInventoryItemCount

`func (o *Manufacturer) GetInventoryItemCount() int32`

GetInventoryItemCount returns the InventoryItemCount field if non-nil, zero value otherwise.

### GetInventoryItemCountOk

`func (o *Manufacturer) GetInventoryItemCountOk() (*int32, bool)`

GetInventoryItemCountOk returns a tuple with the InventoryItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemCount

`func (o *Manufacturer) SetInventoryItemCount(v int32)`

SetInventoryItemCount sets InventoryItemCount field to given value.


### GetPlatformCount

`func (o *Manufacturer) GetPlatformCount() int32`

GetPlatformCount returns the PlatformCount field if non-nil, zero value otherwise.

### GetPlatformCountOk

`func (o *Manufacturer) GetPlatformCountOk() (*int32, bool)`

GetPlatformCountOk returns a tuple with the PlatformCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCount

`func (o *Manufacturer) SetPlatformCount(v int32)`

SetPlatformCount sets PlatformCount field to given value.


### GetName

`func (o *Manufacturer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Manufacturer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Manufacturer) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Manufacturer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Manufacturer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Manufacturer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Manufacturer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *Manufacturer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Manufacturer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Manufacturer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Manufacturer) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Manufacturer) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Manufacturer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Manufacturer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Manufacturer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Manufacturer) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Manufacturer) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *Manufacturer) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Manufacturer) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Manufacturer) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *Manufacturer) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Manufacturer) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Manufacturer) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Manufacturer) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


