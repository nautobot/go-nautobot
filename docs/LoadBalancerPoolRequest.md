# LoadBalancerPoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**LoadBalancingAlgorithm** | [**LoadBalancingAlgorithmEnum**](LoadBalancingAlgorithmEnum.md) |  | 
**HealthCheckMonitor** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewLoadBalancerPoolRequest

`func NewLoadBalancerPoolRequest(name string, loadBalancingAlgorithm LoadBalancingAlgorithmEnum, ) *LoadBalancerPoolRequest`

NewLoadBalancerPoolRequest instantiates a new LoadBalancerPoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerPoolRequestWithDefaults

`func NewLoadBalancerPoolRequestWithDefaults() *LoadBalancerPoolRequest`

NewLoadBalancerPoolRequestWithDefaults instantiates a new LoadBalancerPoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerPoolRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerPoolRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerPoolRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancerPoolRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LoadBalancerPoolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerPoolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerPoolRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLoadBalancingAlgorithm

`func (o *LoadBalancerPoolRequest) GetLoadBalancingAlgorithm() LoadBalancingAlgorithmEnum`

GetLoadBalancingAlgorithm returns the LoadBalancingAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancingAlgorithmOk

`func (o *LoadBalancerPoolRequest) GetLoadBalancingAlgorithmOk() (*LoadBalancingAlgorithmEnum, bool)`

GetLoadBalancingAlgorithmOk returns a tuple with the LoadBalancingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingAlgorithm

`func (o *LoadBalancerPoolRequest) SetLoadBalancingAlgorithm(v LoadBalancingAlgorithmEnum)`

SetLoadBalancingAlgorithm sets LoadBalancingAlgorithm field to given value.


### GetHealthCheckMonitor

`func (o *LoadBalancerPoolRequest) GetHealthCheckMonitor() ApprovalWorkflowUser`

GetHealthCheckMonitor returns the HealthCheckMonitor field if non-nil, zero value otherwise.

### GetHealthCheckMonitorOk

`func (o *LoadBalancerPoolRequest) GetHealthCheckMonitorOk() (*ApprovalWorkflowUser, bool)`

GetHealthCheckMonitorOk returns a tuple with the HealthCheckMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckMonitor

`func (o *LoadBalancerPoolRequest) SetHealthCheckMonitor(v ApprovalWorkflowUser)`

SetHealthCheckMonitor sets HealthCheckMonitor field to given value.

### HasHealthCheckMonitor

`func (o *LoadBalancerPoolRequest) HasHealthCheckMonitor() bool`

HasHealthCheckMonitor returns a boolean if a field has been set.

### SetHealthCheckMonitorNil

`func (o *LoadBalancerPoolRequest) SetHealthCheckMonitorNil(b bool)`

 SetHealthCheckMonitorNil sets the value for HealthCheckMonitor to be an explicit nil

### UnsetHealthCheckMonitor
`func (o *LoadBalancerPoolRequest) UnsetHealthCheckMonitor()`

UnsetHealthCheckMonitor ensures that no value is present for HealthCheckMonitor, not even an explicit nil
### GetTenant

`func (o *LoadBalancerPoolRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *LoadBalancerPoolRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *LoadBalancerPoolRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *LoadBalancerPoolRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *LoadBalancerPoolRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *LoadBalancerPoolRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *LoadBalancerPoolRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *LoadBalancerPoolRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *LoadBalancerPoolRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *LoadBalancerPoolRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *LoadBalancerPoolRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LoadBalancerPoolRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LoadBalancerPoolRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LoadBalancerPoolRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *LoadBalancerPoolRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LoadBalancerPoolRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LoadBalancerPoolRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LoadBalancerPoolRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


