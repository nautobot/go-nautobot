# BulkWritableServiceRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**Protocol** | [***ServiceProtocolChoices**](ServiceProtocolChoices.md) |  | [optional] [default to null]
**Ports** | **[]int32** |  | [default to null]
**Name** | **string** |  | [default to null]
**Description** | **string** |  | [optional] [default to null]
**Device** | [***BulkWritableServiceRequestDevice**](BulkWritableServiceRequest_device.md) |  | [optional] [default to null]
**VirtualMachine** | [***BulkWritableServiceRequestVirtualMachine**](BulkWritableServiceRequest_virtual_machine.md) |  | [optional] [default to null]
**IpAddresses** | [**[]IpAddresses**](IP addresses.md) |  | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

