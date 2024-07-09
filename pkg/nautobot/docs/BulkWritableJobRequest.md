# BulkWritableJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Grouping** | **string** | Human-readable grouping that this job belongs to | 
**Name** | **string** | Human-readable name of this job | 
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

### NewBulkWritableJobRequest

`func NewBulkWritableJobRequest(id string, grouping string, name string, ) *BulkWritableJobRequest`

NewBulkWritableJobRequest instantiates a new BulkWritableJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableJobRequestWithDefaults

`func NewBulkWritableJobRequestWithDefaults() *BulkWritableJobRequest`

NewBulkWritableJobRequestWithDefaults instantiates a new BulkWritableJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableJobRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableJobRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableJobRequest) SetId(v string)`

SetId sets Id field to given value.


### GetGrouping

`func (o *BulkWritableJobRequest) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *BulkWritableJobRequest) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *BulkWritableJobRequest) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.


### GetName

`func (o *BulkWritableJobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableJobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableJobRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableJobRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableJobRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableJobRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableJobRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BulkWritableJobRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkWritableJobRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkWritableJobRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkWritableJobRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHasSensitiveVariables

`func (o *BulkWritableJobRequest) GetHasSensitiveVariables() bool`

GetHasSensitiveVariables returns the HasSensitiveVariables field if non-nil, zero value otherwise.

### GetHasSensitiveVariablesOk

`func (o *BulkWritableJobRequest) GetHasSensitiveVariablesOk() (*bool, bool)`

GetHasSensitiveVariablesOk returns a tuple with the HasSensitiveVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSensitiveVariables

`func (o *BulkWritableJobRequest) SetHasSensitiveVariables(v bool)`

SetHasSensitiveVariables sets HasSensitiveVariables field to given value.

### HasHasSensitiveVariables

`func (o *BulkWritableJobRequest) HasHasSensitiveVariables() bool`

HasHasSensitiveVariables returns a boolean if a field has been set.

### GetApprovalRequired

`func (o *BulkWritableJobRequest) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *BulkWritableJobRequest) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *BulkWritableJobRequest) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.

### HasApprovalRequired

`func (o *BulkWritableJobRequest) HasApprovalRequired() bool`

HasApprovalRequired returns a boolean if a field has been set.

### GetHidden

`func (o *BulkWritableJobRequest) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *BulkWritableJobRequest) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *BulkWritableJobRequest) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *BulkWritableJobRequest) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetDryrunDefault

`func (o *BulkWritableJobRequest) GetDryrunDefault() bool`

GetDryrunDefault returns the DryrunDefault field if non-nil, zero value otherwise.

### GetDryrunDefaultOk

`func (o *BulkWritableJobRequest) GetDryrunDefaultOk() (*bool, bool)`

GetDryrunDefaultOk returns a tuple with the DryrunDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryrunDefault

`func (o *BulkWritableJobRequest) SetDryrunDefault(v bool)`

SetDryrunDefault sets DryrunDefault field to given value.

### HasDryrunDefault

`func (o *BulkWritableJobRequest) HasDryrunDefault() bool`

HasDryrunDefault returns a boolean if a field has been set.

### GetSoftTimeLimit

`func (o *BulkWritableJobRequest) GetSoftTimeLimit() float64`

GetSoftTimeLimit returns the SoftTimeLimit field if non-nil, zero value otherwise.

### GetSoftTimeLimitOk

`func (o *BulkWritableJobRequest) GetSoftTimeLimitOk() (*float64, bool)`

GetSoftTimeLimitOk returns a tuple with the SoftTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftTimeLimit

`func (o *BulkWritableJobRequest) SetSoftTimeLimit(v float64)`

SetSoftTimeLimit sets SoftTimeLimit field to given value.

### HasSoftTimeLimit

`func (o *BulkWritableJobRequest) HasSoftTimeLimit() bool`

HasSoftTimeLimit returns a boolean if a field has been set.

### GetTimeLimit

`func (o *BulkWritableJobRequest) GetTimeLimit() float64`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *BulkWritableJobRequest) GetTimeLimitOk() (*float64, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *BulkWritableJobRequest) SetTimeLimit(v float64)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *BulkWritableJobRequest) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTaskQueues

`func (o *BulkWritableJobRequest) GetTaskQueues() map[string]interface{}`

GetTaskQueues returns the TaskQueues field if non-nil, zero value otherwise.

### GetTaskQueuesOk

`func (o *BulkWritableJobRequest) GetTaskQueuesOk() (*map[string]interface{}, bool)`

GetTaskQueuesOk returns a tuple with the TaskQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueues

`func (o *BulkWritableJobRequest) SetTaskQueues(v map[string]interface{})`

SetTaskQueues sets TaskQueues field to given value.

### HasTaskQueues

`func (o *BulkWritableJobRequest) HasTaskQueues() bool`

HasTaskQueues returns a boolean if a field has been set.

### GetGroupingOverride

`func (o *BulkWritableJobRequest) GetGroupingOverride() bool`

GetGroupingOverride returns the GroupingOverride field if non-nil, zero value otherwise.

### GetGroupingOverrideOk

`func (o *BulkWritableJobRequest) GetGroupingOverrideOk() (*bool, bool)`

GetGroupingOverrideOk returns a tuple with the GroupingOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingOverride

`func (o *BulkWritableJobRequest) SetGroupingOverride(v bool)`

SetGroupingOverride sets GroupingOverride field to given value.

### HasGroupingOverride

`func (o *BulkWritableJobRequest) HasGroupingOverride() bool`

HasGroupingOverride returns a boolean if a field has been set.

### GetNameOverride

`func (o *BulkWritableJobRequest) GetNameOverride() bool`

GetNameOverride returns the NameOverride field if non-nil, zero value otherwise.

### GetNameOverrideOk

`func (o *BulkWritableJobRequest) GetNameOverrideOk() (*bool, bool)`

GetNameOverrideOk returns a tuple with the NameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOverride

`func (o *BulkWritableJobRequest) SetNameOverride(v bool)`

SetNameOverride sets NameOverride field to given value.

### HasNameOverride

`func (o *BulkWritableJobRequest) HasNameOverride() bool`

HasNameOverride returns a boolean if a field has been set.

### GetDescriptionOverride

`func (o *BulkWritableJobRequest) GetDescriptionOverride() bool`

GetDescriptionOverride returns the DescriptionOverride field if non-nil, zero value otherwise.

### GetDescriptionOverrideOk

`func (o *BulkWritableJobRequest) GetDescriptionOverrideOk() (*bool, bool)`

GetDescriptionOverrideOk returns a tuple with the DescriptionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionOverride

`func (o *BulkWritableJobRequest) SetDescriptionOverride(v bool)`

SetDescriptionOverride sets DescriptionOverride field to given value.

### HasDescriptionOverride

`func (o *BulkWritableJobRequest) HasDescriptionOverride() bool`

HasDescriptionOverride returns a boolean if a field has been set.

### GetApprovalRequiredOverride

`func (o *BulkWritableJobRequest) GetApprovalRequiredOverride() bool`

GetApprovalRequiredOverride returns the ApprovalRequiredOverride field if non-nil, zero value otherwise.

### GetApprovalRequiredOverrideOk

`func (o *BulkWritableJobRequest) GetApprovalRequiredOverrideOk() (*bool, bool)`

GetApprovalRequiredOverrideOk returns a tuple with the ApprovalRequiredOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequiredOverride

`func (o *BulkWritableJobRequest) SetApprovalRequiredOverride(v bool)`

SetApprovalRequiredOverride sets ApprovalRequiredOverride field to given value.

### HasApprovalRequiredOverride

`func (o *BulkWritableJobRequest) HasApprovalRequiredOverride() bool`

HasApprovalRequiredOverride returns a boolean if a field has been set.

### GetDryrunDefaultOverride

`func (o *BulkWritableJobRequest) GetDryrunDefaultOverride() bool`

GetDryrunDefaultOverride returns the DryrunDefaultOverride field if non-nil, zero value otherwise.

### GetDryrunDefaultOverrideOk

`func (o *BulkWritableJobRequest) GetDryrunDefaultOverrideOk() (*bool, bool)`

GetDryrunDefaultOverrideOk returns a tuple with the DryrunDefaultOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryrunDefaultOverride

`func (o *BulkWritableJobRequest) SetDryrunDefaultOverride(v bool)`

SetDryrunDefaultOverride sets DryrunDefaultOverride field to given value.

### HasDryrunDefaultOverride

`func (o *BulkWritableJobRequest) HasDryrunDefaultOverride() bool`

HasDryrunDefaultOverride returns a boolean if a field has been set.

### GetHiddenOverride

`func (o *BulkWritableJobRequest) GetHiddenOverride() bool`

GetHiddenOverride returns the HiddenOverride field if non-nil, zero value otherwise.

### GetHiddenOverrideOk

`func (o *BulkWritableJobRequest) GetHiddenOverrideOk() (*bool, bool)`

GetHiddenOverrideOk returns a tuple with the HiddenOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenOverride

`func (o *BulkWritableJobRequest) SetHiddenOverride(v bool)`

SetHiddenOverride sets HiddenOverride field to given value.

### HasHiddenOverride

`func (o *BulkWritableJobRequest) HasHiddenOverride() bool`

HasHiddenOverride returns a boolean if a field has been set.

### GetSoftTimeLimitOverride

`func (o *BulkWritableJobRequest) GetSoftTimeLimitOverride() bool`

GetSoftTimeLimitOverride returns the SoftTimeLimitOverride field if non-nil, zero value otherwise.

### GetSoftTimeLimitOverrideOk

`func (o *BulkWritableJobRequest) GetSoftTimeLimitOverrideOk() (*bool, bool)`

GetSoftTimeLimitOverrideOk returns a tuple with the SoftTimeLimitOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftTimeLimitOverride

`func (o *BulkWritableJobRequest) SetSoftTimeLimitOverride(v bool)`

SetSoftTimeLimitOverride sets SoftTimeLimitOverride field to given value.

### HasSoftTimeLimitOverride

`func (o *BulkWritableJobRequest) HasSoftTimeLimitOverride() bool`

HasSoftTimeLimitOverride returns a boolean if a field has been set.

### GetTimeLimitOverride

`func (o *BulkWritableJobRequest) GetTimeLimitOverride() bool`

GetTimeLimitOverride returns the TimeLimitOverride field if non-nil, zero value otherwise.

### GetTimeLimitOverrideOk

`func (o *BulkWritableJobRequest) GetTimeLimitOverrideOk() (*bool, bool)`

GetTimeLimitOverrideOk returns a tuple with the TimeLimitOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimitOverride

`func (o *BulkWritableJobRequest) SetTimeLimitOverride(v bool)`

SetTimeLimitOverride sets TimeLimitOverride field to given value.

### HasTimeLimitOverride

`func (o *BulkWritableJobRequest) HasTimeLimitOverride() bool`

HasTimeLimitOverride returns a boolean if a field has been set.

### GetHasSensitiveVariablesOverride

`func (o *BulkWritableJobRequest) GetHasSensitiveVariablesOverride() bool`

GetHasSensitiveVariablesOverride returns the HasSensitiveVariablesOverride field if non-nil, zero value otherwise.

### GetHasSensitiveVariablesOverrideOk

`func (o *BulkWritableJobRequest) GetHasSensitiveVariablesOverrideOk() (*bool, bool)`

GetHasSensitiveVariablesOverrideOk returns a tuple with the HasSensitiveVariablesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSensitiveVariablesOverride

`func (o *BulkWritableJobRequest) SetHasSensitiveVariablesOverride(v bool)`

SetHasSensitiveVariablesOverride sets HasSensitiveVariablesOverride field to given value.

### HasHasSensitiveVariablesOverride

`func (o *BulkWritableJobRequest) HasHasSensitiveVariablesOverride() bool`

HasHasSensitiveVariablesOverride returns a boolean if a field has been set.

### GetTaskQueuesOverride

`func (o *BulkWritableJobRequest) GetTaskQueuesOverride() bool`

GetTaskQueuesOverride returns the TaskQueuesOverride field if non-nil, zero value otherwise.

### GetTaskQueuesOverrideOk

`func (o *BulkWritableJobRequest) GetTaskQueuesOverrideOk() (*bool, bool)`

GetTaskQueuesOverrideOk returns a tuple with the TaskQueuesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueuesOverride

`func (o *BulkWritableJobRequest) SetTaskQueuesOverride(v bool)`

SetTaskQueuesOverride sets TaskQueuesOverride field to given value.

### HasTaskQueuesOverride

`func (o *BulkWritableJobRequest) HasTaskQueuesOverride() bool`

HasTaskQueuesOverride returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableJobRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableJobRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableJobRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableJobRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableJobRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableJobRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableJobRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableJobRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableJobRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableJobRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableJobRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableJobRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


