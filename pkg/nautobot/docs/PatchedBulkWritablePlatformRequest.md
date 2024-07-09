# PatchedBulkWritablePlatformRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**NetworkDriver** | Pointer to **string** | The normalized network driver to use when interacting with devices, e.g. cisco_ios, arista_eos, etc. Library-specific driver names will be derived from this setting as appropriate | [optional] 
**NapalmDriver** | Pointer to **string** | The name of the NAPALM driver to use when Nautobot internals interact with devices | [optional] 
**NapalmArgs** | Pointer to **map[string]interface{}** | Additional arguments to pass when initiating the NAPALM driver (JSON format) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to [**NullableBulkWritablePlatformRequestManufacturer**](BulkWritablePlatformRequestManufacturer.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritablePlatformRequest

`func NewPatchedBulkWritablePlatformRequest(id string, ) *PatchedBulkWritablePlatformRequest`

NewPatchedBulkWritablePlatformRequest instantiates a new PatchedBulkWritablePlatformRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritablePlatformRequestWithDefaults

`func NewPatchedBulkWritablePlatformRequestWithDefaults() *PatchedBulkWritablePlatformRequest`

NewPatchedBulkWritablePlatformRequestWithDefaults instantiates a new PatchedBulkWritablePlatformRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritablePlatformRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritablePlatformRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritablePlatformRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritablePlatformRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritablePlatformRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritablePlatformRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritablePlatformRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkDriver

`func (o *PatchedBulkWritablePlatformRequest) GetNetworkDriver() string`

GetNetworkDriver returns the NetworkDriver field if non-nil, zero value otherwise.

### GetNetworkDriverOk

`func (o *PatchedBulkWritablePlatformRequest) GetNetworkDriverOk() (*string, bool)`

GetNetworkDriverOk returns a tuple with the NetworkDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriver

`func (o *PatchedBulkWritablePlatformRequest) SetNetworkDriver(v string)`

SetNetworkDriver sets NetworkDriver field to given value.

### HasNetworkDriver

`func (o *PatchedBulkWritablePlatformRequest) HasNetworkDriver() bool`

HasNetworkDriver returns a boolean if a field has been set.

### GetNapalmDriver

`func (o *PatchedBulkWritablePlatformRequest) GetNapalmDriver() string`

GetNapalmDriver returns the NapalmDriver field if non-nil, zero value otherwise.

### GetNapalmDriverOk

`func (o *PatchedBulkWritablePlatformRequest) GetNapalmDriverOk() (*string, bool)`

GetNapalmDriverOk returns a tuple with the NapalmDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNapalmDriver

`func (o *PatchedBulkWritablePlatformRequest) SetNapalmDriver(v string)`

SetNapalmDriver sets NapalmDriver field to given value.

### HasNapalmDriver

`func (o *PatchedBulkWritablePlatformRequest) HasNapalmDriver() bool`

HasNapalmDriver returns a boolean if a field has been set.

### GetNapalmArgs

`func (o *PatchedBulkWritablePlatformRequest) GetNapalmArgs() map[string]interface{}`

GetNapalmArgs returns the NapalmArgs field if non-nil, zero value otherwise.

### GetNapalmArgsOk

`func (o *PatchedBulkWritablePlatformRequest) GetNapalmArgsOk() (*map[string]interface{}, bool)`

GetNapalmArgsOk returns a tuple with the NapalmArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNapalmArgs

`func (o *PatchedBulkWritablePlatformRequest) SetNapalmArgs(v map[string]interface{})`

SetNapalmArgs sets NapalmArgs field to given value.

### HasNapalmArgs

`func (o *PatchedBulkWritablePlatformRequest) HasNapalmArgs() bool`

HasNapalmArgs returns a boolean if a field has been set.

### SetNapalmArgsNil

`func (o *PatchedBulkWritablePlatformRequest) SetNapalmArgsNil(b bool)`

 SetNapalmArgsNil sets the value for NapalmArgs to be an explicit nil

### UnsetNapalmArgs
`func (o *PatchedBulkWritablePlatformRequest) UnsetNapalmArgs()`

UnsetNapalmArgs ensures that no value is present for NapalmArgs, not even an explicit nil
### GetDescription

`func (o *PatchedBulkWritablePlatformRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritablePlatformRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritablePlatformRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritablePlatformRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManufacturer

`func (o *PatchedBulkWritablePlatformRequest) GetManufacturer() BulkWritablePlatformRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *PatchedBulkWritablePlatformRequest) GetManufacturerOk() (*BulkWritablePlatformRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *PatchedBulkWritablePlatformRequest) SetManufacturer(v BulkWritablePlatformRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *PatchedBulkWritablePlatformRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *PatchedBulkWritablePlatformRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *PatchedBulkWritablePlatformRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritablePlatformRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritablePlatformRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritablePlatformRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritablePlatformRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritablePlatformRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritablePlatformRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritablePlatformRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritablePlatformRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


