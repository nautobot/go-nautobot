# PatchedIpAddressRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | [optional] [default to null]
**Namespace** | [***BulkWritableIpAddressRequestNamespace**](BulkWritableIPAddressRequest_namespace.md) |  | [optional] [default to null]
**Type_** | [***IpAddressTypeChoices**](IPAddressTypeChoices.md) |  | [optional] [default to null]
**DnsName** | **string** | Hostname or FQDN (not case-sensitive) | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**Role** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Parent** | [***BulkWritableIpAddressRequestParent**](BulkWritableIPAddressRequest_parent.md) |  | [optional] [default to null]
**Tenant** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**NatInside** | [***NatInside**](NAT Inside.md) |  | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

