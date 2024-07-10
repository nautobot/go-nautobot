# ScheduledJob

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**Name** | **string** | Human-readable description of this scheduled task | [default to null]
**Task** | **string** | The name of the Celery task that should be run. (Example: \&quot;proj.tasks.import_contacts\&quot;) | [default to null]
**Interval** | [***JobExecutionTypeIntervalChoices**](JobExecutionTypeIntervalChoices.md) |  | [default to null]
**Args** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Kwargs** | [**map[string]Object**](.md) |  | [optional] [default to null]
**CeleryKwargs** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Queue** | **string** | Queue defined in CELERY_TASK_QUEUES. Leave empty for default queuing. | [optional] [default to null]
**OneOff** | **bool** | If True, the schedule will only run the task a single time | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | Datetime when the schedule should begin triggering the task to run | [default to null]
**Enabled** | **bool** | Set to False to disable the schedule | [optional] [default to null]
**LastRunAt** | [**time.Time**](time.Time.md) | Datetime that the schedule last triggered the task to run. Reset to None if enabled is set to False. | [default to null]
**TotalRunCount** | **int32** | Running count of how many times the schedule has triggered the task | [default to null]
**DateChanged** | [**time.Time**](time.Time.md) | Datetime that this scheduled job was last modified | [default to null]
**Description** | **string** | Detailed description about the details of this scheduled job | [optional] [default to null]
**ApprovalRequired** | **bool** |  | [optional] [default to null]
**ApprovedAt** | [**time.Time**](time.Time.md) | Datetime that the schedule was approved | [default to null]
**Crontab** | **string** | Cronjob syntax string for custom scheduling | [optional] [default to null]
**JobModel** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**User** | [***ScheduledJobUser**](ScheduledJob_user.md) |  | [optional] [default to null]
**ApprovedByUser** | [***ScheduledJobApprovedByUser**](ScheduledJob_approved_by_user.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

