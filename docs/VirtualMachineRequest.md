# VirtualMachineRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalConfigContextData** | [**map[string]Object**](.md) |  | [optional] [default to null]
**LocalConfigContextDataOwnerObjectId** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Vcpus** | **int32** |  | [optional] [default to null]
**Memory** | **int32** |  | [optional] [default to null]
**Disk** | **int32** |  | [optional] [default to null]
**Comments** | **string** |  | [optional] [default to null]
**LocalConfigContextSchema** | [***BulkWritableConfigContextRequestConfigContextSchema**](BulkWritableConfigContextRequest_config_context_schema.md) |  | [optional] [default to null]
**LocalConfigContextDataOwnerContentType** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Cluster** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Tenant** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Platform** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Role** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**PrimaryIp4** | [***PrimaryIpv4**](Primary IPv4.md) |  | [optional] [default to null]
**PrimaryIp6** | [***PrimaryIpv6**](Primary IPv6.md) |  | [optional] [default to null]
**SoftwareVersion** | [***BulkWritableVirtualMachineRequestSoftwareVersion**](BulkWritableVirtualMachineRequest_software_version.md) |  | [optional] [default to null]
**SoftwareImageFiles** | [**[]SoftwareImageFiles**](Software Image Files.md) | Override the software image files associated with the software version for this virtual machine | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

