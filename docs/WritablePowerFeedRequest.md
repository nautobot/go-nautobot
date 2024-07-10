# WritablePowerFeedRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**Type_** | [***PowerFeedTypeChoices**](PowerFeedTypeChoices.md) |  | [optional] [default to null]
**Supply** | [***SupplyEnum**](SupplyEnum.md) |  | [optional] [default to null]
**Phase** | [***PhaseEnum**](PhaseEnum.md) |  | [optional] [default to null]
**Voltage** | **int32** |  | [optional] [default to null]
**Amperage** | **int32** |  | [optional] [default to null]
**MaxUtilization** | **int32** | Maximum permissible draw (percentage) | [optional] [default to null]
**Comments** | **string** |  | [optional] [default to null]
**Cable** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**PowerPanel** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Rack** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

