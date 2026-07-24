# BulkWritableVirtualServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Protocol** | [**BulkWritableVirtualServerRequestProtocol**](BulkWritableVirtualServerRequestProtocol.md) |  | 
**SourceNatType** | Pointer to [**BulkWritableVirtualServerRequestSourceNatType**](BulkWritableVirtualServerRequestSourceNatType.md) |  | [optional] 
**LoadBalancerType** | Pointer to [**BulkWritableVirtualServerRequestLoadBalancerType**](BulkWritableVirtualServerRequestLoadBalancerType.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SslOffload** | Pointer to **bool** |  | [optional] 
**Vip** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**SourceNatPool** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Device** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**DeviceRedundancyGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CloudService** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**VirtualChassis** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**LoadBalancerPool** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**HealthCheckMonitor** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewBulkWritableVirtualServerRequest

`func NewBulkWritableVirtualServerRequest(id string, name string, protocol BulkWritableVirtualServerRequestProtocol, vip ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *BulkWritableVirtualServerRequest`

NewBulkWritableVirtualServerRequest instantiates a new BulkWritableVirtualServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVirtualServerRequestWithDefaults

`func NewBulkWritableVirtualServerRequestWithDefaults() *BulkWritableVirtualServerRequest`

NewBulkWritableVirtualServerRequestWithDefaults instantiates a new BulkWritableVirtualServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVirtualServerRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVirtualServerRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVirtualServerRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableVirtualServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableVirtualServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableVirtualServerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *BulkWritableVirtualServerRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BulkWritableVirtualServerRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BulkWritableVirtualServerRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *BulkWritableVirtualServerRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *BulkWritableVirtualServerRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *BulkWritableVirtualServerRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProtocol

`func (o *BulkWritableVirtualServerRequest) GetProtocol() BulkWritableVirtualServerRequestProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BulkWritableVirtualServerRequest) GetProtocolOk() (*BulkWritableVirtualServerRequestProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BulkWritableVirtualServerRequest) SetProtocol(v BulkWritableVirtualServerRequestProtocol)`

SetProtocol sets Protocol field to given value.


### GetSourceNatType

`func (o *BulkWritableVirtualServerRequest) GetSourceNatType() BulkWritableVirtualServerRequestSourceNatType`

GetSourceNatType returns the SourceNatType field if non-nil, zero value otherwise.

### GetSourceNatTypeOk

`func (o *BulkWritableVirtualServerRequest) GetSourceNatTypeOk() (*BulkWritableVirtualServerRequestSourceNatType, bool)`

GetSourceNatTypeOk returns a tuple with the SourceNatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNatType

`func (o *BulkWritableVirtualServerRequest) SetSourceNatType(v BulkWritableVirtualServerRequestSourceNatType)`

SetSourceNatType sets SourceNatType field to given value.

### HasSourceNatType

`func (o *BulkWritableVirtualServerRequest) HasSourceNatType() bool`

HasSourceNatType returns a boolean if a field has been set.

### GetLoadBalancerType

`func (o *BulkWritableVirtualServerRequest) GetLoadBalancerType() BulkWritableVirtualServerRequestLoadBalancerType`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *BulkWritableVirtualServerRequest) GetLoadBalancerTypeOk() (*BulkWritableVirtualServerRequestLoadBalancerType, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *BulkWritableVirtualServerRequest) SetLoadBalancerType(v BulkWritableVirtualServerRequestLoadBalancerType)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *BulkWritableVirtualServerRequest) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### GetEnabled

`func (o *BulkWritableVirtualServerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkWritableVirtualServerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkWritableVirtualServerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkWritableVirtualServerRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSslOffload

`func (o *BulkWritableVirtualServerRequest) GetSslOffload() bool`

GetSslOffload returns the SslOffload field if non-nil, zero value otherwise.

### GetSslOffloadOk

`func (o *BulkWritableVirtualServerRequest) GetSslOffloadOk() (*bool, bool)`

GetSslOffloadOk returns a tuple with the SslOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslOffload

`func (o *BulkWritableVirtualServerRequest) SetSslOffload(v bool)`

SetSslOffload sets SslOffload field to given value.

### HasSslOffload

`func (o *BulkWritableVirtualServerRequest) HasSslOffload() bool`

HasSslOffload returns a boolean if a field has been set.

### GetVip

`func (o *BulkWritableVirtualServerRequest) GetVip() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *BulkWritableVirtualServerRequest) GetVipOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *BulkWritableVirtualServerRequest) SetVip(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVip sets Vip field to given value.


### GetSourceNatPool

`func (o *BulkWritableVirtualServerRequest) GetSourceNatPool() ApprovalWorkflowUser`

GetSourceNatPool returns the SourceNatPool field if non-nil, zero value otherwise.

### GetSourceNatPoolOk

`func (o *BulkWritableVirtualServerRequest) GetSourceNatPoolOk() (*ApprovalWorkflowUser, bool)`

GetSourceNatPoolOk returns a tuple with the SourceNatPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNatPool

`func (o *BulkWritableVirtualServerRequest) SetSourceNatPool(v ApprovalWorkflowUser)`

SetSourceNatPool sets SourceNatPool field to given value.

### HasSourceNatPool

`func (o *BulkWritableVirtualServerRequest) HasSourceNatPool() bool`

HasSourceNatPool returns a boolean if a field has been set.

### SetSourceNatPoolNil

`func (o *BulkWritableVirtualServerRequest) SetSourceNatPoolNil(b bool)`

 SetSourceNatPoolNil sets the value for SourceNatPool to be an explicit nil

### UnsetSourceNatPool
`func (o *BulkWritableVirtualServerRequest) UnsetSourceNatPool()`

UnsetSourceNatPool ensures that no value is present for SourceNatPool, not even an explicit nil
### GetDevice

`func (o *BulkWritableVirtualServerRequest) GetDevice() ApprovalWorkflowUser`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkWritableVirtualServerRequest) GetDeviceOk() (*ApprovalWorkflowUser, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkWritableVirtualServerRequest) SetDevice(v ApprovalWorkflowUser)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *BulkWritableVirtualServerRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *BulkWritableVirtualServerRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *BulkWritableVirtualServerRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetDeviceRedundancyGroup

`func (o *BulkWritableVirtualServerRequest) GetDeviceRedundancyGroup() ApprovalWorkflowUser`

GetDeviceRedundancyGroup returns the DeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupOk

`func (o *BulkWritableVirtualServerRequest) GetDeviceRedundancyGroupOk() (*ApprovalWorkflowUser, bool)`

GetDeviceRedundancyGroupOk returns a tuple with the DeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroup

`func (o *BulkWritableVirtualServerRequest) SetDeviceRedundancyGroup(v ApprovalWorkflowUser)`

SetDeviceRedundancyGroup sets DeviceRedundancyGroup field to given value.

### HasDeviceRedundancyGroup

`func (o *BulkWritableVirtualServerRequest) HasDeviceRedundancyGroup() bool`

HasDeviceRedundancyGroup returns a boolean if a field has been set.

### SetDeviceRedundancyGroupNil

`func (o *BulkWritableVirtualServerRequest) SetDeviceRedundancyGroupNil(b bool)`

 SetDeviceRedundancyGroupNil sets the value for DeviceRedundancyGroup to be an explicit nil

### UnsetDeviceRedundancyGroup
`func (o *BulkWritableVirtualServerRequest) UnsetDeviceRedundancyGroup()`

UnsetDeviceRedundancyGroup ensures that no value is present for DeviceRedundancyGroup, not even an explicit nil
### GetCloudService

`func (o *BulkWritableVirtualServerRequest) GetCloudService() ApprovalWorkflowUser`

GetCloudService returns the CloudService field if non-nil, zero value otherwise.

### GetCloudServiceOk

`func (o *BulkWritableVirtualServerRequest) GetCloudServiceOk() (*ApprovalWorkflowUser, bool)`

GetCloudServiceOk returns a tuple with the CloudService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudService

`func (o *BulkWritableVirtualServerRequest) SetCloudService(v ApprovalWorkflowUser)`

SetCloudService sets CloudService field to given value.

### HasCloudService

`func (o *BulkWritableVirtualServerRequest) HasCloudService() bool`

HasCloudService returns a boolean if a field has been set.

### SetCloudServiceNil

`func (o *BulkWritableVirtualServerRequest) SetCloudServiceNil(b bool)`

 SetCloudServiceNil sets the value for CloudService to be an explicit nil

### UnsetCloudService
`func (o *BulkWritableVirtualServerRequest) UnsetCloudService()`

UnsetCloudService ensures that no value is present for CloudService, not even an explicit nil
### GetVirtualChassis

`func (o *BulkWritableVirtualServerRequest) GetVirtualChassis() ApprovalWorkflowUser`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *BulkWritableVirtualServerRequest) GetVirtualChassisOk() (*ApprovalWorkflowUser, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *BulkWritableVirtualServerRequest) SetVirtualChassis(v ApprovalWorkflowUser)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *BulkWritableVirtualServerRequest) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *BulkWritableVirtualServerRequest) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *BulkWritableVirtualServerRequest) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetTenant

`func (o *BulkWritableVirtualServerRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableVirtualServerRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableVirtualServerRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableVirtualServerRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableVirtualServerRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableVirtualServerRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetLoadBalancerPool

`func (o *BulkWritableVirtualServerRequest) GetLoadBalancerPool() ApprovalWorkflowUser`

GetLoadBalancerPool returns the LoadBalancerPool field if non-nil, zero value otherwise.

### GetLoadBalancerPoolOk

`func (o *BulkWritableVirtualServerRequest) GetLoadBalancerPoolOk() (*ApprovalWorkflowUser, bool)`

GetLoadBalancerPoolOk returns a tuple with the LoadBalancerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPool

`func (o *BulkWritableVirtualServerRequest) SetLoadBalancerPool(v ApprovalWorkflowUser)`

SetLoadBalancerPool sets LoadBalancerPool field to given value.

### HasLoadBalancerPool

`func (o *BulkWritableVirtualServerRequest) HasLoadBalancerPool() bool`

HasLoadBalancerPool returns a boolean if a field has been set.

### SetLoadBalancerPoolNil

`func (o *BulkWritableVirtualServerRequest) SetLoadBalancerPoolNil(b bool)`

 SetLoadBalancerPoolNil sets the value for LoadBalancerPool to be an explicit nil

### UnsetLoadBalancerPool
`func (o *BulkWritableVirtualServerRequest) UnsetLoadBalancerPool()`

UnsetLoadBalancerPool ensures that no value is present for LoadBalancerPool, not even an explicit nil
### GetHealthCheckMonitor

`func (o *BulkWritableVirtualServerRequest) GetHealthCheckMonitor() ApprovalWorkflowUser`

GetHealthCheckMonitor returns the HealthCheckMonitor field if non-nil, zero value otherwise.

### GetHealthCheckMonitorOk

`func (o *BulkWritableVirtualServerRequest) GetHealthCheckMonitorOk() (*ApprovalWorkflowUser, bool)`

GetHealthCheckMonitorOk returns a tuple with the HealthCheckMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckMonitor

`func (o *BulkWritableVirtualServerRequest) SetHealthCheckMonitor(v ApprovalWorkflowUser)`

SetHealthCheckMonitor sets HealthCheckMonitor field to given value.

### HasHealthCheckMonitor

`func (o *BulkWritableVirtualServerRequest) HasHealthCheckMonitor() bool`

HasHealthCheckMonitor returns a boolean if a field has been set.

### SetHealthCheckMonitorNil

`func (o *BulkWritableVirtualServerRequest) SetHealthCheckMonitorNil(b bool)`

 SetHealthCheckMonitorNil sets the value for HealthCheckMonitor to be an explicit nil

### UnsetHealthCheckMonitor
`func (o *BulkWritableVirtualServerRequest) UnsetHealthCheckMonitor()`

UnsetHealthCheckMonitor ensures that no value is present for HealthCheckMonitor, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableVirtualServerRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableVirtualServerRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableVirtualServerRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableVirtualServerRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableVirtualServerRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableVirtualServerRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableVirtualServerRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableVirtualServerRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableVirtualServerRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableVirtualServerRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableVirtualServerRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableVirtualServerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


