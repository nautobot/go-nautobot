# PatchedBulkWritableJobHookRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ContentTypes** | **[]string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**TypeCreate** | **bool** | Call this job hook when a matching object is created. | [optional] [default to null]
**TypeDelete** | **bool** | Call this job hook when a matching object is deleted. | [optional] [default to null]
**TypeUpdate** | **bool** | Call this job hook when a matching object is updated. | [optional] [default to null]
**Job** | [***BulkWritableJobHookRequestJob**](BulkWritableJobHookRequest_job.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

