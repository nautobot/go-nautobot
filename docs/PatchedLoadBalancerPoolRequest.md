# PatchedLoadBalancerPoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LoadBalancingAlgorithm** | Pointer to [**LoadBalancingAlgorithmEnum**](LoadBalancingAlgorithmEnum.md) |  | [optional] 
**HealthCheckMonitor** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewPatchedLoadBalancerPoolRequest

`func NewPatchedLoadBalancerPoolRequest() *PatchedLoadBalancerPoolRequest`

NewPatchedLoadBalancerPoolRequest instantiates a new PatchedLoadBalancerPoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLoadBalancerPoolRequestWithDefaults

`func NewPatchedLoadBalancerPoolRequestWithDefaults() *PatchedLoadBalancerPoolRequest`

NewPatchedLoadBalancerPoolRequestWithDefaults instantiates a new PatchedLoadBalancerPoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedLoadBalancerPoolRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedLoadBalancerPoolRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedLoadBalancerPoolRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedLoadBalancerPoolRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedLoadBalancerPoolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedLoadBalancerPoolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedLoadBalancerPoolRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedLoadBalancerPoolRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLoadBalancingAlgorithm

`func (o *PatchedLoadBalancerPoolRequest) GetLoadBalancingAlgorithm() LoadBalancingAlgorithmEnum`

GetLoadBalancingAlgorithm returns the LoadBalancingAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancingAlgorithmOk

`func (o *PatchedLoadBalancerPoolRequest) GetLoadBalancingAlgorithmOk() (*LoadBalancingAlgorithmEnum, bool)`

GetLoadBalancingAlgorithmOk returns a tuple with the LoadBalancingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingAlgorithm

`func (o *PatchedLoadBalancerPoolRequest) SetLoadBalancingAlgorithm(v LoadBalancingAlgorithmEnum)`

SetLoadBalancingAlgorithm sets LoadBalancingAlgorithm field to given value.

### HasLoadBalancingAlgorithm

`func (o *PatchedLoadBalancerPoolRequest) HasLoadBalancingAlgorithm() bool`

HasLoadBalancingAlgorithm returns a boolean if a field has been set.

### GetHealthCheckMonitor

`func (o *PatchedLoadBalancerPoolRequest) GetHealthCheckMonitor() ApprovalWorkflowUser`

GetHealthCheckMonitor returns the HealthCheckMonitor field if non-nil, zero value otherwise.

### GetHealthCheckMonitorOk

`func (o *PatchedLoadBalancerPoolRequest) GetHealthCheckMonitorOk() (*ApprovalWorkflowUser, bool)`

GetHealthCheckMonitorOk returns a tuple with the HealthCheckMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckMonitor

`func (o *PatchedLoadBalancerPoolRequest) SetHealthCheckMonitor(v ApprovalWorkflowUser)`

SetHealthCheckMonitor sets HealthCheckMonitor field to given value.

### HasHealthCheckMonitor

`func (o *PatchedLoadBalancerPoolRequest) HasHealthCheckMonitor() bool`

HasHealthCheckMonitor returns a boolean if a field has been set.

### SetHealthCheckMonitorNil

`func (o *PatchedLoadBalancerPoolRequest) SetHealthCheckMonitorNil(b bool)`

 SetHealthCheckMonitorNil sets the value for HealthCheckMonitor to be an explicit nil

### UnsetHealthCheckMonitor
`func (o *PatchedLoadBalancerPoolRequest) UnsetHealthCheckMonitor()`

UnsetHealthCheckMonitor ensures that no value is present for HealthCheckMonitor, not even an explicit nil
### GetTenant

`func (o *PatchedLoadBalancerPoolRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedLoadBalancerPoolRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedLoadBalancerPoolRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedLoadBalancerPoolRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedLoadBalancerPoolRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedLoadBalancerPoolRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedLoadBalancerPoolRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedLoadBalancerPoolRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedLoadBalancerPoolRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedLoadBalancerPoolRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedLoadBalancerPoolRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedLoadBalancerPoolRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedLoadBalancerPoolRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedLoadBalancerPoolRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedLoadBalancerPoolRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedLoadBalancerPoolRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedLoadBalancerPoolRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedLoadBalancerPoolRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


