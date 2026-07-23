# PatchedVirtualServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Protocol** | Pointer to [**BulkWritableVirtualServerRequestProtocol**](BulkWritableVirtualServerRequestProtocol.md) |  | [optional] 
**SourceNatType** | Pointer to [**BulkWritableVirtualServerRequestSourceNatType**](BulkWritableVirtualServerRequestSourceNatType.md) |  | [optional] 
**LoadBalancerType** | Pointer to [**BulkWritableVirtualServerRequestLoadBalancerType**](BulkWritableVirtualServerRequestLoadBalancerType.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SslOffload** | Pointer to **bool** |  | [optional] 
**Vip** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
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

### NewPatchedVirtualServerRequest

`func NewPatchedVirtualServerRequest() *PatchedVirtualServerRequest`

NewPatchedVirtualServerRequest instantiates a new PatchedVirtualServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVirtualServerRequestWithDefaults

`func NewPatchedVirtualServerRequestWithDefaults() *PatchedVirtualServerRequest`

NewPatchedVirtualServerRequestWithDefaults instantiates a new PatchedVirtualServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedVirtualServerRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedVirtualServerRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedVirtualServerRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedVirtualServerRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedVirtualServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedVirtualServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedVirtualServerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedVirtualServerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *PatchedVirtualServerRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedVirtualServerRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedVirtualServerRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedVirtualServerRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *PatchedVirtualServerRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *PatchedVirtualServerRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProtocol

`func (o *PatchedVirtualServerRequest) GetProtocol() BulkWritableVirtualServerRequestProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedVirtualServerRequest) GetProtocolOk() (*BulkWritableVirtualServerRequestProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedVirtualServerRequest) SetProtocol(v BulkWritableVirtualServerRequestProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedVirtualServerRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSourceNatType

`func (o *PatchedVirtualServerRequest) GetSourceNatType() BulkWritableVirtualServerRequestSourceNatType`

GetSourceNatType returns the SourceNatType field if non-nil, zero value otherwise.

### GetSourceNatTypeOk

`func (o *PatchedVirtualServerRequest) GetSourceNatTypeOk() (*BulkWritableVirtualServerRequestSourceNatType, bool)`

GetSourceNatTypeOk returns a tuple with the SourceNatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNatType

`func (o *PatchedVirtualServerRequest) SetSourceNatType(v BulkWritableVirtualServerRequestSourceNatType)`

SetSourceNatType sets SourceNatType field to given value.

### HasSourceNatType

`func (o *PatchedVirtualServerRequest) HasSourceNatType() bool`

HasSourceNatType returns a boolean if a field has been set.

### GetLoadBalancerType

`func (o *PatchedVirtualServerRequest) GetLoadBalancerType() BulkWritableVirtualServerRequestLoadBalancerType`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *PatchedVirtualServerRequest) GetLoadBalancerTypeOk() (*BulkWritableVirtualServerRequestLoadBalancerType, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *PatchedVirtualServerRequest) SetLoadBalancerType(v BulkWritableVirtualServerRequestLoadBalancerType)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *PatchedVirtualServerRequest) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedVirtualServerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedVirtualServerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedVirtualServerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedVirtualServerRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSslOffload

`func (o *PatchedVirtualServerRequest) GetSslOffload() bool`

GetSslOffload returns the SslOffload field if non-nil, zero value otherwise.

### GetSslOffloadOk

`func (o *PatchedVirtualServerRequest) GetSslOffloadOk() (*bool, bool)`

GetSslOffloadOk returns a tuple with the SslOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslOffload

`func (o *PatchedVirtualServerRequest) SetSslOffload(v bool)`

SetSslOffload sets SslOffload field to given value.

### HasSslOffload

`func (o *PatchedVirtualServerRequest) HasSslOffload() bool`

HasSslOffload returns a boolean if a field has been set.

### GetVip

`func (o *PatchedVirtualServerRequest) GetVip() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *PatchedVirtualServerRequest) GetVipOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *PatchedVirtualServerRequest) SetVip(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVip sets Vip field to given value.

### HasVip

`func (o *PatchedVirtualServerRequest) HasVip() bool`

HasVip returns a boolean if a field has been set.

### GetSourceNatPool

`func (o *PatchedVirtualServerRequest) GetSourceNatPool() ApprovalWorkflowUser`

GetSourceNatPool returns the SourceNatPool field if non-nil, zero value otherwise.

### GetSourceNatPoolOk

`func (o *PatchedVirtualServerRequest) GetSourceNatPoolOk() (*ApprovalWorkflowUser, bool)`

GetSourceNatPoolOk returns a tuple with the SourceNatPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNatPool

`func (o *PatchedVirtualServerRequest) SetSourceNatPool(v ApprovalWorkflowUser)`

SetSourceNatPool sets SourceNatPool field to given value.

### HasSourceNatPool

`func (o *PatchedVirtualServerRequest) HasSourceNatPool() bool`

HasSourceNatPool returns a boolean if a field has been set.

### SetSourceNatPoolNil

`func (o *PatchedVirtualServerRequest) SetSourceNatPoolNil(b bool)`

 SetSourceNatPoolNil sets the value for SourceNatPool to be an explicit nil

### UnsetSourceNatPool
`func (o *PatchedVirtualServerRequest) UnsetSourceNatPool()`

UnsetSourceNatPool ensures that no value is present for SourceNatPool, not even an explicit nil
### GetDevice

`func (o *PatchedVirtualServerRequest) GetDevice() ApprovalWorkflowUser`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedVirtualServerRequest) GetDeviceOk() (*ApprovalWorkflowUser, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedVirtualServerRequest) SetDevice(v ApprovalWorkflowUser)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedVirtualServerRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *PatchedVirtualServerRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *PatchedVirtualServerRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetDeviceRedundancyGroup

`func (o *PatchedVirtualServerRequest) GetDeviceRedundancyGroup() ApprovalWorkflowUser`

GetDeviceRedundancyGroup returns the DeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupOk

`func (o *PatchedVirtualServerRequest) GetDeviceRedundancyGroupOk() (*ApprovalWorkflowUser, bool)`

GetDeviceRedundancyGroupOk returns a tuple with the DeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroup

`func (o *PatchedVirtualServerRequest) SetDeviceRedundancyGroup(v ApprovalWorkflowUser)`

SetDeviceRedundancyGroup sets DeviceRedundancyGroup field to given value.

### HasDeviceRedundancyGroup

`func (o *PatchedVirtualServerRequest) HasDeviceRedundancyGroup() bool`

HasDeviceRedundancyGroup returns a boolean if a field has been set.

### SetDeviceRedundancyGroupNil

`func (o *PatchedVirtualServerRequest) SetDeviceRedundancyGroupNil(b bool)`

 SetDeviceRedundancyGroupNil sets the value for DeviceRedundancyGroup to be an explicit nil

### UnsetDeviceRedundancyGroup
`func (o *PatchedVirtualServerRequest) UnsetDeviceRedundancyGroup()`

UnsetDeviceRedundancyGroup ensures that no value is present for DeviceRedundancyGroup, not even an explicit nil
### GetCloudService

`func (o *PatchedVirtualServerRequest) GetCloudService() ApprovalWorkflowUser`

GetCloudService returns the CloudService field if non-nil, zero value otherwise.

### GetCloudServiceOk

`func (o *PatchedVirtualServerRequest) GetCloudServiceOk() (*ApprovalWorkflowUser, bool)`

GetCloudServiceOk returns a tuple with the CloudService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudService

`func (o *PatchedVirtualServerRequest) SetCloudService(v ApprovalWorkflowUser)`

SetCloudService sets CloudService field to given value.

### HasCloudService

`func (o *PatchedVirtualServerRequest) HasCloudService() bool`

HasCloudService returns a boolean if a field has been set.

### SetCloudServiceNil

`func (o *PatchedVirtualServerRequest) SetCloudServiceNil(b bool)`

 SetCloudServiceNil sets the value for CloudService to be an explicit nil

### UnsetCloudService
`func (o *PatchedVirtualServerRequest) UnsetCloudService()`

UnsetCloudService ensures that no value is present for CloudService, not even an explicit nil
### GetVirtualChassis

`func (o *PatchedVirtualServerRequest) GetVirtualChassis() ApprovalWorkflowUser`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *PatchedVirtualServerRequest) GetVirtualChassisOk() (*ApprovalWorkflowUser, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *PatchedVirtualServerRequest) SetVirtualChassis(v ApprovalWorkflowUser)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *PatchedVirtualServerRequest) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *PatchedVirtualServerRequest) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *PatchedVirtualServerRequest) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetTenant

`func (o *PatchedVirtualServerRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedVirtualServerRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedVirtualServerRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedVirtualServerRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedVirtualServerRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedVirtualServerRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetLoadBalancerPool

`func (o *PatchedVirtualServerRequest) GetLoadBalancerPool() ApprovalWorkflowUser`

GetLoadBalancerPool returns the LoadBalancerPool field if non-nil, zero value otherwise.

### GetLoadBalancerPoolOk

`func (o *PatchedVirtualServerRequest) GetLoadBalancerPoolOk() (*ApprovalWorkflowUser, bool)`

GetLoadBalancerPoolOk returns a tuple with the LoadBalancerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPool

`func (o *PatchedVirtualServerRequest) SetLoadBalancerPool(v ApprovalWorkflowUser)`

SetLoadBalancerPool sets LoadBalancerPool field to given value.

### HasLoadBalancerPool

`func (o *PatchedVirtualServerRequest) HasLoadBalancerPool() bool`

HasLoadBalancerPool returns a boolean if a field has been set.

### SetLoadBalancerPoolNil

`func (o *PatchedVirtualServerRequest) SetLoadBalancerPoolNil(b bool)`

 SetLoadBalancerPoolNil sets the value for LoadBalancerPool to be an explicit nil

### UnsetLoadBalancerPool
`func (o *PatchedVirtualServerRequest) UnsetLoadBalancerPool()`

UnsetLoadBalancerPool ensures that no value is present for LoadBalancerPool, not even an explicit nil
### GetHealthCheckMonitor

`func (o *PatchedVirtualServerRequest) GetHealthCheckMonitor() ApprovalWorkflowUser`

GetHealthCheckMonitor returns the HealthCheckMonitor field if non-nil, zero value otherwise.

### GetHealthCheckMonitorOk

`func (o *PatchedVirtualServerRequest) GetHealthCheckMonitorOk() (*ApprovalWorkflowUser, bool)`

GetHealthCheckMonitorOk returns a tuple with the HealthCheckMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckMonitor

`func (o *PatchedVirtualServerRequest) SetHealthCheckMonitor(v ApprovalWorkflowUser)`

SetHealthCheckMonitor sets HealthCheckMonitor field to given value.

### HasHealthCheckMonitor

`func (o *PatchedVirtualServerRequest) HasHealthCheckMonitor() bool`

HasHealthCheckMonitor returns a boolean if a field has been set.

### SetHealthCheckMonitorNil

`func (o *PatchedVirtualServerRequest) SetHealthCheckMonitorNil(b bool)`

 SetHealthCheckMonitorNil sets the value for HealthCheckMonitor to be an explicit nil

### UnsetHealthCheckMonitor
`func (o *PatchedVirtualServerRequest) UnsetHealthCheckMonitor()`

UnsetHealthCheckMonitor ensures that no value is present for HealthCheckMonitor, not even an explicit nil
### GetCustomFields

`func (o *PatchedVirtualServerRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedVirtualServerRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedVirtualServerRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedVirtualServerRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedVirtualServerRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedVirtualServerRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedVirtualServerRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedVirtualServerRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedVirtualServerRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedVirtualServerRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedVirtualServerRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedVirtualServerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


