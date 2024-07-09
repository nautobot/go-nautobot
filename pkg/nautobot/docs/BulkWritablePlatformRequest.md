# BulkWritablePlatformRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**NetworkDriver** | Pointer to **string** | The normalized network driver to use when interacting with devices, e.g. cisco_ios, arista_eos, etc. Library-specific driver names will be derived from this setting as appropriate | [optional] 
**NapalmDriver** | Pointer to **string** | The name of the NAPALM driver to use when Nautobot internals interact with devices | [optional] 
**NapalmArgs** | Pointer to **map[string]interface{}** | Additional arguments to pass when initiating the NAPALM driver (JSON format) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to [**NullableBulkWritablePlatformRequestManufacturer**](BulkWritablePlatformRequestManufacturer.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritablePlatformRequest

`func NewBulkWritablePlatformRequest(id string, name string, ) *BulkWritablePlatformRequest`

NewBulkWritablePlatformRequest instantiates a new BulkWritablePlatformRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritablePlatformRequestWithDefaults

`func NewBulkWritablePlatformRequestWithDefaults() *BulkWritablePlatformRequest`

NewBulkWritablePlatformRequestWithDefaults instantiates a new BulkWritablePlatformRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritablePlatformRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritablePlatformRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritablePlatformRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritablePlatformRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritablePlatformRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritablePlatformRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkDriver

`func (o *BulkWritablePlatformRequest) GetNetworkDriver() string`

GetNetworkDriver returns the NetworkDriver field if non-nil, zero value otherwise.

### GetNetworkDriverOk

`func (o *BulkWritablePlatformRequest) GetNetworkDriverOk() (*string, bool)`

GetNetworkDriverOk returns a tuple with the NetworkDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDriver

`func (o *BulkWritablePlatformRequest) SetNetworkDriver(v string)`

SetNetworkDriver sets NetworkDriver field to given value.

### HasNetworkDriver

`func (o *BulkWritablePlatformRequest) HasNetworkDriver() bool`

HasNetworkDriver returns a boolean if a field has been set.

### GetNapalmDriver

`func (o *BulkWritablePlatformRequest) GetNapalmDriver() string`

GetNapalmDriver returns the NapalmDriver field if non-nil, zero value otherwise.

### GetNapalmDriverOk

`func (o *BulkWritablePlatformRequest) GetNapalmDriverOk() (*string, bool)`

GetNapalmDriverOk returns a tuple with the NapalmDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNapalmDriver

`func (o *BulkWritablePlatformRequest) SetNapalmDriver(v string)`

SetNapalmDriver sets NapalmDriver field to given value.

### HasNapalmDriver

`func (o *BulkWritablePlatformRequest) HasNapalmDriver() bool`

HasNapalmDriver returns a boolean if a field has been set.

### GetNapalmArgs

`func (o *BulkWritablePlatformRequest) GetNapalmArgs() map[string]interface{}`

GetNapalmArgs returns the NapalmArgs field if non-nil, zero value otherwise.

### GetNapalmArgsOk

`func (o *BulkWritablePlatformRequest) GetNapalmArgsOk() (*map[string]interface{}, bool)`

GetNapalmArgsOk returns a tuple with the NapalmArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNapalmArgs

`func (o *BulkWritablePlatformRequest) SetNapalmArgs(v map[string]interface{})`

SetNapalmArgs sets NapalmArgs field to given value.

### HasNapalmArgs

`func (o *BulkWritablePlatformRequest) HasNapalmArgs() bool`

HasNapalmArgs returns a boolean if a field has been set.

### SetNapalmArgsNil

`func (o *BulkWritablePlatformRequest) SetNapalmArgsNil(b bool)`

 SetNapalmArgsNil sets the value for NapalmArgs to be an explicit nil

### UnsetNapalmArgs
`func (o *BulkWritablePlatformRequest) UnsetNapalmArgs()`

UnsetNapalmArgs ensures that no value is present for NapalmArgs, not even an explicit nil
### GetDescription

`func (o *BulkWritablePlatformRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritablePlatformRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritablePlatformRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritablePlatformRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManufacturer

`func (o *BulkWritablePlatformRequest) GetManufacturer() BulkWritablePlatformRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BulkWritablePlatformRequest) GetManufacturerOk() (*BulkWritablePlatformRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BulkWritablePlatformRequest) SetManufacturer(v BulkWritablePlatformRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *BulkWritablePlatformRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *BulkWritablePlatformRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *BulkWritablePlatformRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetCustomFields

`func (o *BulkWritablePlatformRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritablePlatformRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritablePlatformRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritablePlatformRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritablePlatformRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritablePlatformRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritablePlatformRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritablePlatformRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


