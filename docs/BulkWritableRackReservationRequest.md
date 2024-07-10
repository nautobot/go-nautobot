# BulkWritableRackReservationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**Units** | [**map[string]Object**](.md) | List of rack unit numbers to reserve | [default to null]
**Description** | **string** |  | [default to null]
**Rack** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Tenant** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**User** | [***BulkWritableRackReservationRequestUser**](BulkWritableRackReservationRequest_user.md) |  | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

