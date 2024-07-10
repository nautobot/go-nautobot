# WritablePowerOutletTemplateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**Label** | **string** | Physical label | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Type_** | [***OneOfWritablePowerOutletTemplateRequestType_**](OneOfWritablePowerOutletTemplateRequestType_.md) |  | [optional] [default to null]
**FeedLeg** | [***OneOfWritablePowerOutletTemplateRequestFeedLeg**](OneOfWritablePowerOutletTemplateRequestFeedLeg.md) | Phase (for three-phase feeds) | [optional] [default to null]
**DeviceType** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**PowerPortTemplate** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

