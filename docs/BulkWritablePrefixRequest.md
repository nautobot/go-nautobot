# BulkWritablePrefixRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**Prefix** | **string** |  | [default to null]
**Type_** | [***AllOfBulkWritablePrefixRequestType_**](AllOfBulkWritablePrefixRequestType_.md) |  | [optional] [default to {"value":"network","label":"Network"}]
**Location** | [***BulkWritablePrefixRequestLocation**](BulkWritablePrefixRequest_location.md) |  | [optional] [default to null]
**DateAllocated** | [**time.Time**](time.Time.md) | Date this prefix was allocated to an RIR, reserved in IPAM, etc. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Role** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Parent** | [***BulkWritablePrefixRequestParent**](BulkWritablePrefixRequest_parent.md) |  | [optional] [default to null]
**Namespace** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**Tenant** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Vlan** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Rir** | [***BulkWritablePrefixRequestRir**](BulkWritablePrefixRequest_rir.md) |  | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

