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
**ApprovalRequired** | Pointer to **bool** | Whether the job requires approval from another user before running | [optional] 
**Hidden** | Pointer to **bool** | Whether the job defaults to not being shown in the UI | [optional] 
**DryrunDefault** | Pointer to **bool** | Whether the job defaults to running with dryrun argument set to true | [optional] 
**SoftTimeLimit** | Pointer to **float64** | Maximum runtime in seconds before the job will receive a &lt;code&gt;SoftTimeLimitExceeded&lt;/code&gt; exception.&lt;br&gt;Set to 0 to use Nautobot system default | [optional] 
**TimeLimit** | Pointer to **float64** | Maximum runtime in seconds before the job will be forcibly terminated.&lt;br&gt;Set to 0 to use Nautobot system default | [optional] 
**TaskQueues** | Pointer to **map[string]interface{}** | Comma separated list of task queues that this job can run on. A blank list will use the default queue | [optional] 
**GroupingOverride** | Pointer to **bool** | If set, the configured grouping will remain even if the underlying Job source code changes | [optional] 
**NameOverride** | Pointer to **bool** | If set, the configured name will remain even if the underlying Job source code changes | [optional] 
**DescriptionOverride** | Pointer to **bool** | If set, the configured description will remain even if the underlying Job source code changes | [optional] 
**ApprovalRequiredOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**DryrunDefaultOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**HiddenOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**SoftTimeLimitOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**TimeLimitOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**HasSensitiveVariablesOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**TaskQueuesOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

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

### GetApprovalRequired

`func (o *PatchedBulkWritableJobRequest) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *PatchedBulkWritableJobRequest) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *PatchedBulkWritableJobRequest) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.

### HasApprovalRequired

`func (o *PatchedBulkWritableJobRequest) HasApprovalRequired() bool`

HasApprovalRequired returns a boolean if a field has been set.

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

### GetTaskQueues

`func (o *PatchedBulkWritableJobRequest) GetTaskQueues() map[string]interface{}`

GetTaskQueues returns the TaskQueues field if non-nil, zero value otherwise.

### GetTaskQueuesOk

`func (o *PatchedBulkWritableJobRequest) GetTaskQueuesOk() (*map[string]interface{}, bool)`

GetTaskQueuesOk returns a tuple with the TaskQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueues

`func (o *PatchedBulkWritableJobRequest) SetTaskQueues(v map[string]interface{})`

SetTaskQueues sets TaskQueues field to given value.

### HasTaskQueues

`func (o *PatchedBulkWritableJobRequest) HasTaskQueues() bool`

HasTaskQueues returns a boolean if a field has been set.

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

### GetApprovalRequiredOverride

`func (o *PatchedBulkWritableJobRequest) GetApprovalRequiredOverride() bool`

GetApprovalRequiredOverride returns the ApprovalRequiredOverride field if non-nil, zero value otherwise.

### GetApprovalRequiredOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetApprovalRequiredOverrideOk() (*bool, bool)`

GetApprovalRequiredOverrideOk returns a tuple with the ApprovalRequiredOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequiredOverride

`func (o *PatchedBulkWritableJobRequest) SetApprovalRequiredOverride(v bool)`

SetApprovalRequiredOverride sets ApprovalRequiredOverride field to given value.

### HasApprovalRequiredOverride

`func (o *PatchedBulkWritableJobRequest) HasApprovalRequiredOverride() bool`

HasApprovalRequiredOverride returns a boolean if a field has been set.

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

### GetTaskQueuesOverride

`func (o *PatchedBulkWritableJobRequest) GetTaskQueuesOverride() bool`

GetTaskQueuesOverride returns the TaskQueuesOverride field if non-nil, zero value otherwise.

### GetTaskQueuesOverrideOk

`func (o *PatchedBulkWritableJobRequest) GetTaskQueuesOverrideOk() (*bool, bool)`

GetTaskQueuesOverrideOk returns a tuple with the TaskQueuesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueuesOverride

`func (o *PatchedBulkWritableJobRequest) SetTaskQueuesOverride(v bool)`

SetTaskQueuesOverride sets TaskQueuesOverride field to given value.

### HasTaskQueuesOverride

`func (o *PatchedBulkWritableJobRequest) HasTaskQueuesOverride() bool`

HasTaskQueuesOverride returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableJobRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableJobRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableJobRequest) SetTags(v []BulkWritableCableRequestStatus)`

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

`func (o *PatchedBulkWritableJobRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableJobRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableJobRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableJobRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


