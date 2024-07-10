# Prefix

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**Prefix** | **string** |  | [default to null]
**Type_** | [***PrefixType**](Prefix_type.md) |  | [optional] [default to null]
**Network** | **string** | IPv4 or IPv6 network address | [default to null]
**Broadcast** | **string** | IPv4 or IPv6 broadcast address | [default to null]
**PrefixLength** | **int32** | Length of the Network prefix, in bits. | [default to null]
**IpVersion** | **int32** |  | [default to null]
**DateAllocated** | [**time.Time**](time.Time.md) | Date this prefix was allocated to an RIR, reserved in IPAM, etc. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Role** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Parent** | [***BulkWritablePrefixRequestParent**](BulkWritablePrefixRequest_parent.md) |  | [optional] [default to null]
**Namespace** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**Tenant** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Vlan** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Rir** | [***BulkWritablePrefixRequestRir**](BulkWritablePrefixRequest_rir.md) |  | [optional] [default to null]
**Locations** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

