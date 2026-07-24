# LoadBalancerPoolMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Optional label for the load balancer pool member. | [optional] 
**Port** | **int32** |  | 
**SslOffload** | Pointer to **bool** |  | [optional] 
**IpAddress** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**LoadBalancerPool** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**HealthCheckMonitor** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Status** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewLoadBalancerPoolMemberRequest

`func NewLoadBalancerPoolMemberRequest(port int32, ipAddress ApprovalWorkflowStageResponseApprovalWorkflowStage, loadBalancerPool ApprovalWorkflowStageResponseApprovalWorkflowStage, status ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *LoadBalancerPoolMemberRequest`

NewLoadBalancerPoolMemberRequest instantiates a new LoadBalancerPoolMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerPoolMemberRequestWithDefaults

`func NewLoadBalancerPoolMemberRequestWithDefaults() *LoadBalancerPoolMemberRequest`

NewLoadBalancerPoolMemberRequestWithDefaults instantiates a new LoadBalancerPoolMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerPoolMemberRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerPoolMemberRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerPoolMemberRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancerPoolMemberRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *LoadBalancerPoolMemberRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LoadBalancerPoolMemberRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LoadBalancerPoolMemberRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LoadBalancerPoolMemberRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPort

`func (o *LoadBalancerPoolMemberRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerPoolMemberRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerPoolMemberRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSslOffload

`func (o *LoadBalancerPoolMemberRequest) GetSslOffload() bool`

GetSslOffload returns the SslOffload field if non-nil, zero value otherwise.

### GetSslOffloadOk

`func (o *LoadBalancerPoolMemberRequest) GetSslOffloadOk() (*bool, bool)`

GetSslOffloadOk returns a tuple with the SslOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslOffload

`func (o *LoadBalancerPoolMemberRequest) SetSslOffload(v bool)`

SetSslOffload sets SslOffload field to given value.

### HasSslOffload

`func (o *LoadBalancerPoolMemberRequest) HasSslOffload() bool`

HasSslOffload returns a boolean if a field has been set.

### GetIpAddress

`func (o *LoadBalancerPoolMemberRequest) GetIpAddress() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LoadBalancerPoolMemberRequest) GetIpAddressOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LoadBalancerPoolMemberRequest) SetIpAddress(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetIpAddress sets IpAddress field to given value.


### GetLoadBalancerPool

`func (o *LoadBalancerPoolMemberRequest) GetLoadBalancerPool() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLoadBalancerPool returns the LoadBalancerPool field if non-nil, zero value otherwise.

### GetLoadBalancerPoolOk

`func (o *LoadBalancerPoolMemberRequest) GetLoadBalancerPoolOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLoadBalancerPoolOk returns a tuple with the LoadBalancerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPool

`func (o *LoadBalancerPoolMemberRequest) SetLoadBalancerPool(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLoadBalancerPool sets LoadBalancerPool field to given value.


### GetHealthCheckMonitor

`func (o *LoadBalancerPoolMemberRequest) GetHealthCheckMonitor() ApprovalWorkflowUser`

GetHealthCheckMonitor returns the HealthCheckMonitor field if non-nil, zero value otherwise.

### GetHealthCheckMonitorOk

`func (o *LoadBalancerPoolMemberRequest) GetHealthCheckMonitorOk() (*ApprovalWorkflowUser, bool)`

GetHealthCheckMonitorOk returns a tuple with the HealthCheckMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckMonitor

`func (o *LoadBalancerPoolMemberRequest) SetHealthCheckMonitor(v ApprovalWorkflowUser)`

SetHealthCheckMonitor sets HealthCheckMonitor field to given value.

### HasHealthCheckMonitor

`func (o *LoadBalancerPoolMemberRequest) HasHealthCheckMonitor() bool`

HasHealthCheckMonitor returns a boolean if a field has been set.

### SetHealthCheckMonitorNil

`func (o *LoadBalancerPoolMemberRequest) SetHealthCheckMonitorNil(b bool)`

 SetHealthCheckMonitorNil sets the value for HealthCheckMonitor to be an explicit nil

### UnsetHealthCheckMonitor
`func (o *LoadBalancerPoolMemberRequest) UnsetHealthCheckMonitor()`

UnsetHealthCheckMonitor ensures that no value is present for HealthCheckMonitor, not even an explicit nil
### GetTenant

`func (o *LoadBalancerPoolMemberRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *LoadBalancerPoolMemberRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *LoadBalancerPoolMemberRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *LoadBalancerPoolMemberRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *LoadBalancerPoolMemberRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *LoadBalancerPoolMemberRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *LoadBalancerPoolMemberRequest) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoadBalancerPoolMemberRequest) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoadBalancerPoolMemberRequest) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.


### GetCustomFields

`func (o *LoadBalancerPoolMemberRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *LoadBalancerPoolMemberRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *LoadBalancerPoolMemberRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *LoadBalancerPoolMemberRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *LoadBalancerPoolMemberRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LoadBalancerPoolMemberRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LoadBalancerPoolMemberRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LoadBalancerPoolMemberRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *LoadBalancerPoolMemberRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LoadBalancerPoolMemberRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LoadBalancerPoolMemberRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LoadBalancerPoolMemberRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


