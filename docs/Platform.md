# Platform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**NetworkDriverMappings** | **interface{}** |  | [readonly] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] 
**VirtualMachineCount** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**NetworkDriver** | Pointer to **string** | The normalized network driver to use when interacting with devices, e.g. cisco_ios, arista_eos, etc. Library-specific driver names will be derived from this setting as appropriate | [optional] 
**NapalmDriver** | Pointer to **string** | The name of the NAPALM driver to use when Nautobot internals interact with devices | [optional] 
**NapalmArgs** | Pointer to **interface{}** | Additional arguments to pass when initiating the NAPALM driver (JSON format) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to [**NullableBulkWritablePlatformRequestManufacturer**](BulkWritablePlatformRequestManufacturer.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPlatform

`func NewPlatform(id string, objectType string, display string, url string, naturalSlug string, networkDriverMappings interface{}, name string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Platform`

NewPlatform instantiates a new Platform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformWithDefaults

`func NewPlatformWithDefaults() *Platform`

NewPlatformWithDefaults instantiates a new Platform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Platform) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Platform) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Platform) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Platform) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Platform) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Platform) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Platform) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Platform) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Platform) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Platform) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Platform) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Platform) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Platform) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Platform) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Platform) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetNetworkDriverMappings

`func (o *Platform) GetNetworkDriverMappings() interface{}`

GetNetworkDriverMappings returns the NetworkDriverMappings field if non-nil, zero value otherwise.

### GetNetworkDriverMappingsOk

`func (o *Platform) GetNetworkDriverMappingsOk() (*interface{}, bool)`

GetNetworkDriverMappingsOk returns a tuple with the NetworkDriverMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriverMappings

`func (o *Platform) SetNetworkDriverMappings(v interface{})`

SetNetworkDriverMappings sets NetworkDriverMappings field to given value.


### SetNetworkDriverMappingsNil

`func (o *Platform) SetNetworkDriverMappingsNil(b bool)`

 SetNetworkDriverMappingsNil sets the value for NetworkDriverMappings to be an explicit nil

### UnsetNetworkDriverMappings
`func (o *Platform) UnsetNetworkDriverMappings()`

UnsetNetworkDriverMappings ensures that no value is present for NetworkDriverMappings, not even an explicit nil
### GetDeviceCount

`func (o *Platform) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *Platform) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *Platform) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *Platform) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetVirtualMachineCount

`func (o *Platform) GetVirtualMachineCount() int32`

GetVirtualMachineCount returns the VirtualMachineCount field if non-nil, zero value otherwise.

### GetVirtualMachineCountOk

`func (o *Platform) GetVirtualMachineCountOk() (*int32, bool)`

GetVirtualMachineCountOk returns a tuple with the VirtualMachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineCount

`func (o *Platform) SetVirtualMachineCount(v int32)`

SetVirtualMachineCount sets VirtualMachineCount field to given value.

### HasVirtualMachineCount

`func (o *Platform) HasVirtualMachineCount() bool`

HasVirtualMachineCount returns a boolean if a field has been set.

### GetName

`func (o *Platform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Platform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Platform) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkDriver

`func (o *Platform) GetNetworkDriver() string`

GetNetworkDriver returns the NetworkDriver field if non-nil, zero value otherwise.

### GetNetworkDriverOk

`func (o *Platform) GetNetworkDriverOk() (*string, bool)`

GetNetworkDriverOk returns a tuple with the NetworkDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriver

`func (o *Platform) SetNetworkDriver(v string)`

SetNetworkDriver sets NetworkDriver field to given value.

### HasNetworkDriver

`func (o *Platform) HasNetworkDriver() bool`

HasNetworkDriver returns a boolean if a field has been set.

### GetNapalmDriver

`func (o *Platform) GetNapalmDriver() string`

GetNapalmDriver returns the NapalmDriver field if non-nil, zero value otherwise.

### GetNapalmDriverOk

`func (o *Platform) GetNapalmDriverOk() (*string, bool)`

GetNapalmDriverOk returns a tuple with the NapalmDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNapalmDriver

`func (o *Platform) SetNapalmDriver(v string)`

SetNapalmDriver sets NapalmDriver field to given value.

### HasNapalmDriver

`func (o *Platform) HasNapalmDriver() bool`

HasNapalmDriver returns a boolean if a field has been set.

### GetNapalmArgs

`func (o *Platform) GetNapalmArgs() interface{}`

GetNapalmArgs returns the NapalmArgs field if non-nil, zero value otherwise.

### GetNapalmArgsOk

`func (o *Platform) GetNapalmArgsOk() (*interface{}, bool)`

GetNapalmArgsOk returns a tuple with the NapalmArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNapalmArgs

`func (o *Platform) SetNapalmArgs(v interface{})`

SetNapalmArgs sets NapalmArgs field to given value.

### HasNapalmArgs

`func (o *Platform) HasNapalmArgs() bool`

HasNapalmArgs returns a boolean if a field has been set.

### SetNapalmArgsNil

`func (o *Platform) SetNapalmArgsNil(b bool)`

 SetNapalmArgsNil sets the value for NapalmArgs to be an explicit nil

### UnsetNapalmArgs
`func (o *Platform) UnsetNapalmArgs()`

UnsetNapalmArgs ensures that no value is present for NapalmArgs, not even an explicit nil
### GetDescription

`func (o *Platform) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Platform) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Platform) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Platform) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManufacturer

`func (o *Platform) GetManufacturer() BulkWritablePlatformRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *Platform) GetManufacturerOk() (*BulkWritablePlatformRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *Platform) SetManufacturer(v BulkWritablePlatformRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *Platform) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *Platform) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *Platform) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetCreated

`func (o *Platform) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Platform) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Platform) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Platform) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Platform) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Platform) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Platform) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Platform) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Platform) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Platform) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *Platform) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Platform) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Platform) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *Platform) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Platform) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Platform) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Platform) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


