# Platform

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**NetworkDriverMappings** | [**map[string]Object**](.md) |  | [default to null]
**DeviceCount** | **int32** |  | [default to null]
**VirtualMachineCount** | **int32** |  | [default to null]
**Name** | **string** |  | [default to null]
**NetworkDriver** | **string** | The normalized network driver to use when interacting with devices, e.g. cisco_ios, arista_eos, etc. Library-specific driver names will be derived from this setting as appropriate | [optional] [default to null]
**NapalmDriver** | **string** | The name of the NAPALM driver to use when Nautobot internals interact with devices | [optional] [default to null]
**NapalmArgs** | [**map[string]Object**](.md) | Additional arguments to pass when initiating the NAPALM driver (JSON format) | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Manufacturer** | [***BulkWritablePlatformRequestManufacturer**](BulkWritablePlatformRequest_manufacturer.md) |  | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

