# BulkWritableLoadBalancerPoolMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Label** | Pointer to **string** | Optional label for the load balancer pool member. | [optional] 
**Port** | **int32** |  | 
**SslOffload** | Pointer to **bool** |  | [optional] 
**IpAddress** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**LoadBalancerPool** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**HealthCheckMonitor** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableLoadBalancerPoolMemberRequest

`func NewBulkWritableLoadBalancerPoolMemberRequest(id string, port int32, ipAddress BulkWritableCableRequestStatus, loadBalancerPool BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, ) *BulkWritableLoadBalancerPoolMemberRequest`

NewBulkWritableLoadBalancerPoolMemberRequest instantiates a new BulkWritableLoadBalancerPoolMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableLoadBalancerPoolMemberRequestWithDefaults

`func NewBulkWritableLoadBalancerPoolMemberRequestWithDefaults() *BulkWritableLoadBalancerPoolMemberRequest`

NewBulkWritableLoadBalancerPoolMemberRequestWithDefaults instantiates a new BulkWritableLoadBalancerPoolMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkWritableLoadBalancerPoolMemberRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPort

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSslOffload

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetSslOffload() bool`

GetSslOffload returns the SslOffload field if non-nil, zero value otherwise.

### GetSslOffloadOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetSslOffloadOk() (*bool, bool)`

GetSslOffloadOk returns a tuple with the SslOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslOffload

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetSslOffload(v bool)`

SetSslOffload sets SslOffload field to given value.

### HasSslOffload

`func (o *BulkWritableLoadBalancerPoolMemberRequest) HasSslOffload() bool`

HasSslOffload returns a boolean if a field has been set.

### GetIpAddress

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetIpAddress() BulkWritableCableRequestStatus`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetIpAddressOk() (*BulkWritableCableRequestStatus, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetIpAddress(v BulkWritableCableRequestStatus)`

SetIpAddress sets IpAddress field to given value.


### GetLoadBalancerPool

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetLoadBalancerPool() BulkWritableCableRequestStatus`

GetLoadBalancerPool returns the LoadBalancerPool field if non-nil, zero value otherwise.

### GetLoadBalancerPoolOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetLoadBalancerPoolOk() (*BulkWritableCableRequestStatus, bool)`

GetLoadBalancerPoolOk returns a tuple with the LoadBalancerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPool

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetLoadBalancerPool(v BulkWritableCableRequestStatus)`

SetLoadBalancerPool sets LoadBalancerPool field to given value.


### GetHealthCheckMonitor

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetHealthCheckMonitor() ApprovalWorkflowUser`

GetHealthCheckMonitor returns the HealthCheckMonitor field if non-nil, zero value otherwise.

### GetHealthCheckMonitorOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetHealthCheckMonitorOk() (*ApprovalWorkflowUser, bool)`

GetHealthCheckMonitorOk returns a tuple with the HealthCheckMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckMonitor

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetHealthCheckMonitor(v ApprovalWorkflowUser)`

SetHealthCheckMonitor sets HealthCheckMonitor field to given value.

### HasHealthCheckMonitor

`func (o *BulkWritableLoadBalancerPoolMemberRequest) HasHealthCheckMonitor() bool`

HasHealthCheckMonitor returns a boolean if a field has been set.

### SetHealthCheckMonitorNil

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetHealthCheckMonitorNil(b bool)`

 SetHealthCheckMonitorNil sets the value for HealthCheckMonitor to be an explicit nil

### UnsetHealthCheckMonitor
`func (o *BulkWritableLoadBalancerPoolMemberRequest) UnsetHealthCheckMonitor()`

UnsetHealthCheckMonitor ensures that no value is present for HealthCheckMonitor, not even an explicit nil
### GetTenant

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableLoadBalancerPoolMemberRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableLoadBalancerPoolMemberRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetCustomFields

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableLoadBalancerPoolMemberRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableLoadBalancerPoolMemberRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableLoadBalancerPoolMemberRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableLoadBalancerPoolMemberRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableLoadBalancerPoolMemberRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


