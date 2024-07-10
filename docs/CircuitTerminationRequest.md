# CircuitTerminationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermSide** | [***AllOfCircuitTerminationRequestTermSide**](AllOfCircuitTerminationRequestTermSide.md) |  | [default to null]
**PortSpeed** | **int32** |  | [optional] [default to null]
**UpstreamSpeed** | **int32** | Upstream speed, if different from port speed | [optional] [default to null]
**XconnectId** | **string** |  | [optional] [default to null]
**PpInfo** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Circuit** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Location** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**ProviderNetwork** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

