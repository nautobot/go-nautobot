# BulkWritablePowerFeedRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**Type_** | [***AllOfBulkWritablePowerFeedRequestType_**](AllOfBulkWritablePowerFeedRequestType_.md) |  | [optional] [default to {"value":"primary","label":"Primary"}]
**Supply** | [***AllOfBulkWritablePowerFeedRequestSupply**](AllOfBulkWritablePowerFeedRequestSupply.md) |  | [optional] [default to {"value":"ac","label":"AC"}]
**Phase** | [***AllOfBulkWritablePowerFeedRequestPhase**](AllOfBulkWritablePowerFeedRequestPhase.md) |  | [optional] [default to {"value":"single-phase","label":"Single phase"}]
**Name** | **string** |  | [default to null]
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

