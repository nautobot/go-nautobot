# PatchedWritablePowerOutletRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [optional] [default to null]
**Label** | **string** | Physical label | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Type_** | [***OneOfPatchedWritablePowerOutletRequestType_**](OneOfPatchedWritablePowerOutletRequestType_.md) | Physical port type | [optional] [default to null]
**FeedLeg** | [***OneOfPatchedWritablePowerOutletRequestFeedLeg**](OneOfPatchedWritablePowerOutletRequestFeedLeg.md) | Phase (for three-phase feeds) | [optional] [default to null]
**Device** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**PowerPort** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

