# PatchedJobRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grouping** | **string** | Human-readable grouping that this job belongs to | [optional] [default to null]
**Name** | **string** | Human-readable name of this job | [optional] [default to null]
**Description** | **string** | Markdown formatting and a limited subset of HTML are supported | [optional] [default to null]
**Enabled** | **bool** | Whether this job can be executed by users | [optional] [default to null]
**HasSensitiveVariables** | **bool** | Whether this job contains sensitive variables | [optional] [default to null]
**ApprovalRequired** | **bool** | Whether the job requires approval from another user before running | [optional] [default to null]
**Hidden** | **bool** | Whether the job defaults to not being shown in the UI | [optional] [default to null]
**DryrunDefault** | **bool** | Whether the job defaults to running with dryrun argument set to true | [optional] [default to null]
**SoftTimeLimit** | **float64** | Maximum runtime in seconds before the job will receive a &lt;code&gt;SoftTimeLimitExceeded&lt;/code&gt; exception.&lt;br&gt;Set to 0 to use Nautobot system default | [optional] [default to null]
**TimeLimit** | **float64** | Maximum runtime in seconds before the job will be forcibly terminated.&lt;br&gt;Set to 0 to use Nautobot system default | [optional] [default to null]
**TaskQueues** | [**map[string]Object**](.md) | Comma separated list of task queues that this job can run on. A blank list will use the default queue | [optional] [default to null]
**GroupingOverride** | **bool** | If set, the configured grouping will remain even if the underlying Job source code changes | [optional] [default to null]
**NameOverride** | **bool** | If set, the configured name will remain even if the underlying Job source code changes | [optional] [default to null]
**DescriptionOverride** | **bool** | If set, the configured description will remain even if the underlying Job source code changes | [optional] [default to null]
**ApprovalRequiredOverride** | **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] [default to null]
**DryrunDefaultOverride** | **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] [default to null]
**HiddenOverride** | **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] [default to null]
**SoftTimeLimitOverride** | **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] [default to null]
**TimeLimitOverride** | **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] [default to null]
**HasSensitiveVariablesOverride** | **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] [default to null]
**TaskQueuesOverride** | **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

