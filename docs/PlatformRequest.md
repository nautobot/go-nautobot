# PlatformRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**NetworkDriver** | **string** | The normalized network driver to use when interacting with devices, e.g. cisco_ios, arista_eos, etc. Library-specific driver names will be derived from this setting as appropriate | [optional] [default to null]
**NapalmDriver** | **string** | The name of the NAPALM driver to use when Nautobot internals interact with devices | [optional] [default to null]
**NapalmArgs** | [**map[string]Object**](.md) | Additional arguments to pass when initiating the NAPALM driver (JSON format) | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Manufacturer** | [***BulkWritablePlatformRequestManufacturer**](BulkWritablePlatformRequest_manufacturer.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

