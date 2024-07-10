# PatchedBulkWritablePowerPortTemplateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**Type_** | [***PowerPortTypeChoices**](PowerPortTypeChoices.md) |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Label** | **string** | Physical label | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**MaximumDraw** | **int32** | Maximum power draw (watts) | [optional] [default to null]
**AllocatedDraw** | **int32** | Allocated power draw (watts) | [optional] [default to null]
**DeviceType** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

