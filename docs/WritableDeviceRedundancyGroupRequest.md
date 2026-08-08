# WritableDeviceRedundancyGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**FailoverStrategy** | Pointer to [**BulkWritableDeviceRedundancyGroupRequestFailoverStrategy**](BulkWritableDeviceRedundancyGroupRequestFailoverStrategy.md) |  | [optional] [default to BULKWRITABLEDEVICEREDUNDANCYGROUPREQUESTFAILOVERSTRATEGY_EMPTY]
**Comments** | Pointer to **string** |  | [optional] 
**Status** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**SecretsGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewWritableDeviceRedundancyGroupRequest

`func NewWritableDeviceRedundancyGroupRequest(name string, status ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *WritableDeviceRedundancyGroupRequest`

NewWritableDeviceRedundancyGroupRequest instantiates a new WritableDeviceRedundancyGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableDeviceRedundancyGroupRequestWithDefaults

`func NewWritableDeviceRedundancyGroupRequestWithDefaults() *WritableDeviceRedundancyGroupRequest`

NewWritableDeviceRedundancyGroupRequestWithDefaults instantiates a new WritableDeviceRedundancyGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableDeviceRedundancyGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableDeviceRedundancyGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableDeviceRedundancyGroupRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WritableDeviceRedundancyGroupRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WritableDeviceRedundancyGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableDeviceRedundancyGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableDeviceRedundancyGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WritableDeviceRedundancyGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableDeviceRedundancyGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableDeviceRedundancyGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableDeviceRedundancyGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFailoverStrategy

`func (o *WritableDeviceRedundancyGroupRequest) GetFailoverStrategy() BulkWritableDeviceRedundancyGroupRequestFailoverStrategy`

GetFailoverStrategy returns the FailoverStrategy field if non-nil, zero value otherwise.

### GetFailoverStrategyOk

`func (o *WritableDeviceRedundancyGroupRequest) GetFailoverStrategyOk() (*BulkWritableDeviceRedundancyGroupRequestFailoverStrategy, bool)`

GetFailoverStrategyOk returns a tuple with the FailoverStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverStrategy

`func (o *WritableDeviceRedundancyGroupRequest) SetFailoverStrategy(v BulkWritableDeviceRedundancyGroupRequestFailoverStrategy)`

SetFailoverStrategy sets FailoverStrategy field to given value.

### HasFailoverStrategy

`func (o *WritableDeviceRedundancyGroupRequest) HasFailoverStrategy() bool`

HasFailoverStrategy returns a boolean if a field has been set.

### GetComments

`func (o *WritableDeviceRedundancyGroupRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableDeviceRedundancyGroupRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableDeviceRedundancyGroupRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableDeviceRedundancyGroupRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetStatus

`func (o *WritableDeviceRedundancyGroupRequest) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableDeviceRedundancyGroupRequest) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableDeviceRedundancyGroupRequest) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.


### GetSecretsGroup

`func (o *WritableDeviceRedundancyGroupRequest) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *WritableDeviceRedundancyGroupRequest) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *WritableDeviceRedundancyGroupRequest) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *WritableDeviceRedundancyGroupRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *WritableDeviceRedundancyGroupRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *WritableDeviceRedundancyGroupRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCustomFields

`func (o *WritableDeviceRedundancyGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableDeviceRedundancyGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableDeviceRedundancyGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableDeviceRedundancyGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritableDeviceRedundancyGroupRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritableDeviceRedundancyGroupRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritableDeviceRedundancyGroupRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritableDeviceRedundancyGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *WritableDeviceRedundancyGroupRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableDeviceRedundancyGroupRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableDeviceRedundancyGroupRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableDeviceRedundancyGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


