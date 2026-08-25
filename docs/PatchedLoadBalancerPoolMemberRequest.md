# PatchedLoadBalancerPoolMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Optional label for the load balancer pool member. | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**SslOffload** | Pointer to **bool** |  | [optional] 
**IpAddress** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**LoadBalancerPool** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**HealthCheckMonitor** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedLoadBalancerPoolMemberRequest

`func NewPatchedLoadBalancerPoolMemberRequest() *PatchedLoadBalancerPoolMemberRequest`

NewPatchedLoadBalancerPoolMemberRequest instantiates a new PatchedLoadBalancerPoolMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLoadBalancerPoolMemberRequestWithDefaults

`func NewPatchedLoadBalancerPoolMemberRequestWithDefaults() *PatchedLoadBalancerPoolMemberRequest`

NewPatchedLoadBalancerPoolMemberRequestWithDefaults instantiates a new PatchedLoadBalancerPoolMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedLoadBalancerPoolMemberRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedLoadBalancerPoolMemberRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedLoadBalancerPoolMemberRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedLoadBalancerPoolMemberRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedLoadBalancerPoolMemberRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedLoadBalancerPoolMemberRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPort

`func (o *PatchedLoadBalancerPoolMemberRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedLoadBalancerPoolMemberRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedLoadBalancerPoolMemberRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSslOffload

`func (o *PatchedLoadBalancerPoolMemberRequest) GetSslOffload() bool`

GetSslOffload returns the SslOffload field if non-nil, zero value otherwise.

### GetSslOffloadOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetSslOffloadOk() (*bool, bool)`

GetSslOffloadOk returns a tuple with the SslOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslOffload

`func (o *PatchedLoadBalancerPoolMemberRequest) SetSslOffload(v bool)`

SetSslOffload sets SslOffload field to given value.

### HasSslOffload

`func (o *PatchedLoadBalancerPoolMemberRequest) HasSslOffload() bool`

HasSslOffload returns a boolean if a field has been set.

### GetIpAddress

`func (o *PatchedLoadBalancerPoolMemberRequest) GetIpAddress() BulkWritableCableRequestStatus`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetIpAddressOk() (*BulkWritableCableRequestStatus, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PatchedLoadBalancerPoolMemberRequest) SetIpAddress(v BulkWritableCableRequestStatus)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *PatchedLoadBalancerPoolMemberRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLoadBalancerPool

`func (o *PatchedLoadBalancerPoolMemberRequest) GetLoadBalancerPool() BulkWritableCableRequestStatus`

GetLoadBalancerPool returns the LoadBalancerPool field if non-nil, zero value otherwise.

### GetLoadBalancerPoolOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetLoadBalancerPoolOk() (*BulkWritableCableRequestStatus, bool)`

GetLoadBalancerPoolOk returns a tuple with the LoadBalancerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPool

`func (o *PatchedLoadBalancerPoolMemberRequest) SetLoadBalancerPool(v BulkWritableCableRequestStatus)`

SetLoadBalancerPool sets LoadBalancerPool field to given value.

### HasLoadBalancerPool

`func (o *PatchedLoadBalancerPoolMemberRequest) HasLoadBalancerPool() bool`

HasLoadBalancerPool returns a boolean if a field has been set.

### GetHealthCheckMonitor

`func (o *PatchedLoadBalancerPoolMemberRequest) GetHealthCheckMonitor() ApprovalWorkflowUser`

GetHealthCheckMonitor returns the HealthCheckMonitor field if non-nil, zero value otherwise.

### GetHealthCheckMonitorOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetHealthCheckMonitorOk() (*ApprovalWorkflowUser, bool)`

GetHealthCheckMonitorOk returns a tuple with the HealthCheckMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckMonitor

`func (o *PatchedLoadBalancerPoolMemberRequest) SetHealthCheckMonitor(v ApprovalWorkflowUser)`

SetHealthCheckMonitor sets HealthCheckMonitor field to given value.

### HasHealthCheckMonitor

`func (o *PatchedLoadBalancerPoolMemberRequest) HasHealthCheckMonitor() bool`

HasHealthCheckMonitor returns a boolean if a field has been set.

### SetHealthCheckMonitorNil

`func (o *PatchedLoadBalancerPoolMemberRequest) SetHealthCheckMonitorNil(b bool)`

 SetHealthCheckMonitorNil sets the value for HealthCheckMonitor to be an explicit nil

### UnsetHealthCheckMonitor
`func (o *PatchedLoadBalancerPoolMemberRequest) UnsetHealthCheckMonitor()`

UnsetHealthCheckMonitor ensures that no value is present for HealthCheckMonitor, not even an explicit nil
### GetTenant

`func (o *PatchedLoadBalancerPoolMemberRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedLoadBalancerPoolMemberRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedLoadBalancerPoolMemberRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedLoadBalancerPoolMemberRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedLoadBalancerPoolMemberRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *PatchedLoadBalancerPoolMemberRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedLoadBalancerPoolMemberRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedLoadBalancerPoolMemberRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedLoadBalancerPoolMemberRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedLoadBalancerPoolMemberRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedLoadBalancerPoolMemberRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedLoadBalancerPoolMemberRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedLoadBalancerPoolMemberRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedLoadBalancerPoolMemberRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedLoadBalancerPoolMemberRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedLoadBalancerPoolMemberRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedLoadBalancerPoolMemberRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedLoadBalancerPoolMemberRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


