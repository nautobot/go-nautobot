# PatchedBulkWritableJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Grouping** | Pointer to **string** | Human-readable grouping that this job belongs to | [optional] 
**Name** | Pointer to **string** | Human-readable name of this job | [optional] 
**Description** | Pointer to **string** | Markdown formatting and a limited subset of HTML are supported | [optional] 
**Enabled** | Pointer to **bool** | Whether this job can be executed by users | [optional] 
**HasSensitiveVariables** | Pointer to **bool** | Whether this job contains sensitive variables | [optional] 
**IsSingleton** | Pointer to **bool** | Whether this job should fail to run if another instance of this job is already running | [optional] 
**Hidden** | Pointer to **bool** | Whether the job defaults to not being shown in the UI | [optional] 
**DryrunDefault** | Pointer to **bool** | Whether the job defaults to running with dryrun argument set to true | [optional] 
**SoftTimeLimit** | Pointer to **float64** | Maximum runtime in seconds before the job will receive a &lt;code&gt;SoftTimeLimitExceeded&lt;/code&gt; exception.&lt;br&gt;Set to 0 to use Nautobot system default | [optional] 
**TimeLimit** | Pointer to **float64** | Maximum runtime in seconds before the job will be forcibly terminated.&lt;br&gt;Set to 0 to use Nautobot system default | [optional] 
**GroupingOverride** | Pointer to **bool** | If set, the configured grouping will remain even if the underlying Job source code changes | [optional] 
**NameOverride** | Pointer to **bool** | If set, the configured name will remain even if the underlying Job source code changes | [optional] 
**DescriptionOverride** | Pointer to **bool** | If set, the configured description will remain even if the underlying Job source code changes | [optional] 
**DryrunDefaultOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**HiddenOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**SoftTimeLimitOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**TimeLimitOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**HasSensitiveVariablesOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**JobQueuesOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**DefaultJobQueueOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**IsSingletonOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**DefaultJobQueue** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableJobRequest

`func NewPatchedBulkWritableJobRequest(id string, ) *PatchedBulkWritableJobRequest`

NewPatchedBulkWritableJobRequest instantiates a new PatchedBulkWritableJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableJobRequestWithDefaults

`func NewPatchedBulkWritableJobRequestWithDefaults() *PatchedBulkWritableJobRequest`

NewPatchedBulkWritableJobRequestWithDefaults instantiates a new PatchedBulkWritableJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableJobRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableJobRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableJobRequest) SetId(v string)`

SetId sets Id field to given value.


### GetGrouping

`func (o *PatchedBulkWritableJobRequest) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *PatchedBulkWritableJobRequest) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *PatchedBulkWritableJobRequest) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *PatchedBulkWritableJobRequest) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableJobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableJobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableJobRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableJobRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableJobRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableJobRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableJobRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableJobRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkWritableJobRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkWritableJobRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkWritableJobRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkWritableJobRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHasSensitiveVariables

`func (o *PatchedBulkWritableJobRequest) GetHasSensitiveVariables() bool`

GetHasSensitiveVariables returns the HasSensitiveVariables field if non-nil, zero value otherwise.

### GetHasSensitiveVariablesOk

`func (o *PatchedBulkWritableJobRequest) GetHasSensitiveVariablesOk() (*bool, bool)`

GetHasSensitiveVariablesOk returns a tuple with the HasSensitiveVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSensitiveVariables

`func (o *PatchedBulkWritableJobRequest) SetHasSensitiveVariables(v bool)`

SetHasSensitiveVariables sets HasSensitiveVariables field to given value.

### HasHasSensitiveVariables

`func (o *PatchedBulkWritableJobRequest) HasHasSensitiveVariables() bool`

HasHasSensitiveVariables returns a boolean if a field has been set.

### GetIsSingleton

`func (o *PatchedBulkWritableJobRequest) GetIsSingleton() bool`

GetIsSingleton returns the IsSingleton field if non-nil, zero value otherwise.

### GetIsSingletonOk

`func (o *PatchedBulkWritableJobRequest) GetIsSingletonOk() (*bool, bool)`

GetIsSingletonOk returns a tuple with the IsSingleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleton

`func (o *PatchedBulkWritableJobRequest) SetIsSingleton(v bool)`

SetIsSingleton sets IsSingleton field to given value.

### HasIsSingleton

`func (o *PatchedBulkWritableJobRequest) HasIsSingleton() bool`

HasIsSingleton returns a boolean if a field has been set.

### GetHidden

`func (o *PatchedBulkWritableJobRequest) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PatchedBulkWritableJobRequest) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PatchedBulkWritableJobRequest) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PatchedBulkWritableJobRequest) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetDryrunDefault

`func (o *PatchedBulkWritableJobRequest) GetDryrunDefault() bool`

GetDryrunDefault returns the DryrunDefault field if non-nil, zero value otherwise.

### GetDryrunDefaultOk

`func (o *PatchedBulkWritableJobRequest) GetDryrunDefaultOk() (*bool, bool)`

GetDryrunDefaultOk returns a tuple with the DryrunDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryrunDefault

`func (o *PatchedBulkWritableJobRequest) SetDryrunDefault(v bool)`

SetDryrunDefault sets DryrunDefault field to given value.

### HasDryrunDefault

`func (o *PatchedBulkWritableJobRequest) HasDryrunDefault() bool`

HasDryrunDefault returns a boolean if a field has been set.

### GetSoftTimeLimit

`func (o *PatchedBulkWritableJobRequest) GetSoftTimeLimit() float64`

GetSoftTimeLimit returns the SoftTimeLimit field if non-nil, zero value otherwise.

### GetSoftTimeLimitOk

`func (o *PatchedBulkWritableJobRequest) GetSoftTimeLimitOk() (*float64, bool)`

GetSoftTimeLimitOk returns a tuple with the SoftTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftTimeLimit

`func (o *PatchedBulkWritableJobRequest) SetSoftTimeLimit(v float64)`

SetSoftTimeLimit sets SoftTimeLimit field to given value.

### HasSoftTimeLimit

`func (o *PatchedBulkWritableJobRequest) HasSoftTimeLimit() bool`

HasSoftTimeLimit returns a boolean if a field has been set.

### GetTimeLimit

`func (o *PatchedBulkWritableJobRequest) GetTimeLimit() float64`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *PatchedBulkWritableJobRequest) GetTimeLimitOk() (*float64, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *PatchedBulkWritableJobRequest) SetTimeLimit(v float64)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *PatchedBulkWritableJobRequest) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetGroupingOverride

`func (o *PatchedBulkWritableJobRequest) GetGroupingOverride() bool`

GetGroupingOverride returns the GroupingOverride field if non-nil, zero value otherwise.

### GetGroupingOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetGroupingOverrideOk() (*bool, bool)`

GetGroupingOverrideOk returns a tuple with the GroupingOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingOverride

`func (o *PatchedBulkWritableJobRequest) SetGroupingOverride(v bool)`

SetGroupingOverride sets GroupingOverride field to given value.

### HasGroupingOverride

`func (o *PatchedBulkWritableJobRequest) HasGroupingOverride() bool`

HasGroupingOverride returns a boolean if a field has been set.

### GetNameOverride

`func (o *PatchedBulkWritableJobRequest) GetNameOverride() bool`

GetNameOverride returns the NameOverride field if non-nil, zero value otherwise.

### GetNameOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetNameOverrideOk() (*bool, bool)`

GetNameOverrideOk returns a tuple with the NameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOverride

`func (o *PatchedBulkWritableJobRequest) SetNameOverride(v bool)`

SetNameOverride sets NameOverride field to given value.

### HasNameOverride

`func (o *PatchedBulkWritableJobRequest) HasNameOverride() bool`

HasNameOverride returns a boolean if a field has been set.

### GetDescriptionOverride

`func (o *PatchedBulkWritableJobRequest) GetDescriptionOverride() bool`

GetDescriptionOverride returns the DescriptionOverride field if non-nil, zero value otherwise.

### GetDescriptionOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetDescriptionOverrideOk() (*bool, bool)`

GetDescriptionOverrideOk returns a tuple with the DescriptionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionOverride

`func (o *PatchedBulkWritableJobRequest) SetDescriptionOverride(v bool)`

SetDescriptionOverride sets DescriptionOverride field to given value.

### HasDescriptionOverride

`func (o *PatchedBulkWritableJobRequest) HasDescriptionOverride() bool`

HasDescriptionOverride returns a boolean if a field has been set.

### GetDryrunDefaultOverride

`func (o *PatchedBulkWritableJobRequest) GetDryrunDefaultOverride() bool`

GetDryrunDefaultOverride returns the DryrunDefaultOverride field if non-nil, zero value otherwise.

### GetDryrunDefaultOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetDryrunDefaultOverrideOk() (*bool, bool)`

GetDryrunDefaultOverrideOk returns a tuple with the DryrunDefaultOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryrunDefaultOverride

`func (o *PatchedBulkWritableJobRequest) SetDryrunDefaultOverride(v bool)`

SetDryrunDefaultOverride sets DryrunDefaultOverride field to given value.

### HasDryrunDefaultOverride

`func (o *PatchedBulkWritableJobRequest) HasDryrunDefaultOverride() bool`

HasDryrunDefaultOverride returns a boolean if a field has been set.

### GetHiddenOverride

`func (o *PatchedBulkWritableJobRequest) GetHiddenOverride() bool`

GetHiddenOverride returns the HiddenOverride field if non-nil, zero value otherwise.

### GetHiddenOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetHiddenOverrideOk() (*bool, bool)`

GetHiddenOverrideOk returns a tuple with the HiddenOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenOverride

`func (o *PatchedBulkWritableJobRequest) SetHiddenOverride(v bool)`

SetHiddenOverride sets HiddenOverride field to given value.

### HasHiddenOverride

`func (o *PatchedBulkWritableJobRequest) HasHiddenOverride() bool`

HasHiddenOverride returns a boolean if a field has been set.

### GetSoftTimeLimitOverride

`func (o *PatchedBulkWritableJobRequest) GetSoftTimeLimitOverride() bool`

GetSoftTimeLimitOverride returns the SoftTimeLimitOverride field if non-nil, zero value otherwise.

### GetSoftTimeLimitOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetSoftTimeLimitOverrideOk() (*bool, bool)`

GetSoftTimeLimitOverrideOk returns a tuple with the SoftTimeLimitOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftTimeLimitOverride

`func (o *PatchedBulkWritableJobRequest) SetSoftTimeLimitOverride(v bool)`

SetSoftTimeLimitOverride sets SoftTimeLimitOverride field to given value.

### HasSoftTimeLimitOverride

`func (o *PatchedBulkWritableJobRequest) HasSoftTimeLimitOverride() bool`

HasSoftTimeLimitOverride returns a boolean if a field has been set.

### GetTimeLimitOverride

`func (o *PatchedBulkWritableJobRequest) GetTimeLimitOverride() bool`

GetTimeLimitOverride returns the TimeLimitOverride field if non-nil, zero value otherwise.

### GetTimeLimitOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetTimeLimitOverrideOk() (*bool, bool)`

GetTimeLimitOverrideOk returns a tuple with the TimeLimitOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimitOverride

`func (o *PatchedBulkWritableJobRequest) SetTimeLimitOverride(v bool)`

SetTimeLimitOverride sets TimeLimitOverride field to given value.

### HasTimeLimitOverride

`func (o *PatchedBulkWritableJobRequest) HasTimeLimitOverride() bool`

HasTimeLimitOverride returns a boolean if a field has been set.

### GetHasSensitiveVariablesOverride

`func (o *PatchedBulkWritableJobRequest) GetHasSensitiveVariablesOverride() bool`

GetHasSensitiveVariablesOverride returns the HasSensitiveVariablesOverride field if non-nil, zero value otherwise.

### GetHasSensitiveVariablesOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetHasSensitiveVariablesOverrideOk() (*bool, bool)`

GetHasSensitiveVariablesOverrideOk returns a tuple with the HasSensitiveVariablesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSensitiveVariablesOverride

`func (o *PatchedBulkWritableJobRequest) SetHasSensitiveVariablesOverride(v bool)`

SetHasSensitiveVariablesOverride sets HasSensitiveVariablesOverride field to given value.

### HasHasSensitiveVariablesOverride

`func (o *PatchedBulkWritableJobRequest) HasHasSensitiveVariablesOverride() bool`

HasHasSensitiveVariablesOverride returns a boolean if a field has been set.

### GetJobQueuesOverride

`func (o *PatchedBulkWritableJobRequest) GetJobQueuesOverride() bool`

GetJobQueuesOverride returns the JobQueuesOverride field if non-nil, zero value otherwise.

### GetJobQueuesOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetJobQueuesOverrideOk() (*bool, bool)`

GetJobQueuesOverrideOk returns a tuple with the JobQueuesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobQueuesOverride

`func (o *PatchedBulkWritableJobRequest) SetJobQueuesOverride(v bool)`

SetJobQueuesOverride sets JobQueuesOverride field to given value.

### HasJobQueuesOverride

`func (o *PatchedBulkWritableJobRequest) HasJobQueuesOverride() bool`

HasJobQueuesOverride returns a boolean if a field has been set.

### GetDefaultJobQueueOverride

`func (o *PatchedBulkWritableJobRequest) GetDefaultJobQueueOverride() bool`

GetDefaultJobQueueOverride returns the DefaultJobQueueOverride field if non-nil, zero value otherwise.

### GetDefaultJobQueueOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetDefaultJobQueueOverrideOk() (*bool, bool)`

GetDefaultJobQueueOverrideOk returns a tuple with the DefaultJobQueueOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultJobQueueOverride

`func (o *PatchedBulkWritableJobRequest) SetDefaultJobQueueOverride(v bool)`

SetDefaultJobQueueOverride sets DefaultJobQueueOverride field to given value.

### HasDefaultJobQueueOverride

`func (o *PatchedBulkWritableJobRequest) HasDefaultJobQueueOverride() bool`

HasDefaultJobQueueOverride returns a boolean if a field has been set.

### GetIsSingletonOverride

`func (o *PatchedBulkWritableJobRequest) GetIsSingletonOverride() bool`

GetIsSingletonOverride returns the IsSingletonOverride field if non-nil, zero value otherwise.

### GetIsSingletonOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetIsSingletonOverrideOk() (*bool, bool)`

GetIsSingletonOverrideOk returns a tuple with the IsSingletonOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingletonOverride

`func (o *PatchedBulkWritableJobRequest) SetIsSingletonOverride(v bool)`

SetIsSingletonOverride sets IsSingletonOverride field to given value.

### HasIsSingletonOverride

`func (o *PatchedBulkWritableJobRequest) HasIsSingletonOverride() bool`

HasIsSingletonOverride returns a boolean if a field has been set.

### GetDefaultJobQueue

`func (o *PatchedBulkWritableJobRequest) GetDefaultJobQueue() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDefaultJobQueue returns the DefaultJobQueue field if non-nil, zero value otherwise.

### GetDefaultJobQueueOk

`func (o *PatchedBulkWritableJobRequest) GetDefaultJobQueueOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDefaultJobQueueOk returns a tuple with the DefaultJobQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultJobQueue

`func (o *PatchedBulkWritableJobRequest) SetDefaultJobQueue(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDefaultJobQueue sets DefaultJobQueue field to given value.

### HasDefaultJobQueue

`func (o *PatchedBulkWritableJobRequest) HasDefaultJobQueue() bool`

HasDefaultJobQueue returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableJobRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableJobRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableJobRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableJobRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableJobRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableJobRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableJobRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableJobRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableJobRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableJobRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableJobRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableJobRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


