# JobResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**Status** | [***JobResultStatus**](JobResult_status.md) |  | [default to null]
**Name** | **string** |  | [default to null]
**TaskName** | **string** | Registered name of the Celery task for this job. Internal use only. | [optional] [default to null]
**DateCreated** | [**time.Time**](time.Time.md) |  | [default to null]
**DateDone** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Result** | [**map[string]Object**](.md) | The data returned by the task | [default to null]
**Worker** | **string** |  | [optional] [default to null]
**TaskArgs** | [**map[string]Object**](.md) |  | [optional] [default to null]
**TaskKwargs** | [**map[string]Object**](.md) |  | [optional] [default to null]
**CeleryKwargs** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Traceback** | **string** |  | [optional] [default to null]
**Meta** | [**map[string]Object**](.md) |  | [default to null]
**JobModel** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**User** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**ScheduledJob** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Files** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

