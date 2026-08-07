# VirtualServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Protocol** | [**BulkWritableVirtualServerRequestProtocol**](BulkWritableVirtualServerRequestProtocol.md) |  | 
**SourceNatType** | Pointer to [**BulkWritableVirtualServerRequestSourceNatType**](BulkWritableVirtualServerRequestSourceNatType.md) |  | [optional] 
**LoadBalancerType** | Pointer to [**BulkWritableVirtualServerRequestLoadBalancerType**](BulkWritableVirtualServerRequestLoadBalancerType.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SslOffload** | Pointer to **bool** |  | [optional] 
**Vip** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**SourceNatPool** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Device** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**DeviceRedundancyGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CloudService** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**VirtualChassis** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**LoadBalancerPool** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**HealthCheckMonitor** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewVirtualServer

`func NewVirtualServer(objectType string, display string, url string, naturalSlug string, name string, protocol BulkWritableVirtualServerRequestProtocol, vip BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *VirtualServer`

NewVirtualServer instantiates a new VirtualServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualServerWithDefaults

`func NewVirtualServerWithDefaults() *VirtualServer`

NewVirtualServerWithDefaults instantiates a new VirtualServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VirtualServer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualServer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualServer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VirtualServer) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VirtualServer) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VirtualServer) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VirtualServer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualServer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualServer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VirtualServer) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VirtualServer) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VirtualServer) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetName

`func (o *VirtualServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualServer) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *VirtualServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VirtualServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VirtualServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VirtualServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *VirtualServer) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *VirtualServer) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProtocol

`func (o *VirtualServer) GetProtocol() BulkWritableVirtualServerRequestProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VirtualServer) GetProtocolOk() (*BulkWritableVirtualServerRequestProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VirtualServer) SetProtocol(v BulkWritableVirtualServerRequestProtocol)`

SetProtocol sets Protocol field to given value.


### GetSourceNatType

`func (o *VirtualServer) GetSourceNatType() BulkWritableVirtualServerRequestSourceNatType`

GetSourceNatType returns the SourceNatType field if non-nil, zero value otherwise.

### GetSourceNatTypeOk

`func (o *VirtualServer) GetSourceNatTypeOk() (*BulkWritableVirtualServerRequestSourceNatType, bool)`

GetSourceNatTypeOk returns a tuple with the SourceNatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNatType

`func (o *VirtualServer) SetSourceNatType(v BulkWritableVirtualServerRequestSourceNatType)`

SetSourceNatType sets SourceNatType field to given value.

### HasSourceNatType

`func (o *VirtualServer) HasSourceNatType() bool`

HasSourceNatType returns a boolean if a field has been set.

### GetLoadBalancerType

`func (o *VirtualServer) GetLoadBalancerType() BulkWritableVirtualServerRequestLoadBalancerType`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *VirtualServer) GetLoadBalancerTypeOk() (*BulkWritableVirtualServerRequestLoadBalancerType, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *VirtualServer) SetLoadBalancerType(v BulkWritableVirtualServerRequestLoadBalancerType)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *VirtualServer) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### GetEnabled

`func (o *VirtualServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VirtualServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VirtualServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VirtualServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSslOffload

`func (o *VirtualServer) GetSslOffload() bool`

GetSslOffload returns the SslOffload field if non-nil, zero value otherwise.

### GetSslOffloadOk

`func (o *VirtualServer) GetSslOffloadOk() (*bool, bool)`

GetSslOffloadOk returns a tuple with the SslOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslOffload

`func (o *VirtualServer) SetSslOffload(v bool)`

SetSslOffload sets SslOffload field to given value.

### HasSslOffload

`func (o *VirtualServer) HasSslOffload() bool`

HasSslOffload returns a boolean if a field has been set.

### GetVip

`func (o *VirtualServer) GetVip() BulkWritableCableRequestStatus`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *VirtualServer) GetVipOk() (*BulkWritableCableRequestStatus, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *VirtualServer) SetVip(v BulkWritableCableRequestStatus)`

SetVip sets Vip field to given value.


### GetSourceNatPool

`func (o *VirtualServer) GetSourceNatPool() ApprovalWorkflowUser`

GetSourceNatPool returns the SourceNatPool field if non-nil, zero value otherwise.

### GetSourceNatPoolOk

`func (o *VirtualServer) GetSourceNatPoolOk() (*ApprovalWorkflowUser, bool)`

GetSourceNatPoolOk returns a tuple with the SourceNatPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNatPool

`func (o *VirtualServer) SetSourceNatPool(v ApprovalWorkflowUser)`

SetSourceNatPool sets SourceNatPool field to given value.

### HasSourceNatPool

`func (o *VirtualServer) HasSourceNatPool() bool`

HasSourceNatPool returns a boolean if a field has been set.

### SetSourceNatPoolNil

`func (o *VirtualServer) SetSourceNatPoolNil(b bool)`

 SetSourceNatPoolNil sets the value for SourceNatPool to be an explicit nil

### UnsetSourceNatPool
`func (o *VirtualServer) UnsetSourceNatPool()`

UnsetSourceNatPool ensures that no value is present for SourceNatPool, not even an explicit nil
### GetDevice

`func (o *VirtualServer) GetDevice() ApprovalWorkflowUser`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VirtualServer) GetDeviceOk() (*ApprovalWorkflowUser, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VirtualServer) SetDevice(v ApprovalWorkflowUser)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VirtualServer) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *VirtualServer) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *VirtualServer) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetDeviceRedundancyGroup

`func (o *VirtualServer) GetDeviceRedundancyGroup() ApprovalWorkflowUser`

GetDeviceRedundancyGroup returns the DeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetDeviceRedundancyGroupOk

`func (o *VirtualServer) GetDeviceRedundancyGroupOk() (*ApprovalWorkflowUser, bool)`

GetDeviceRedundancyGroupOk returns a tuple with the DeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRedundancyGroup

`func (o *VirtualServer) SetDeviceRedundancyGroup(v ApprovalWorkflowUser)`

SetDeviceRedundancyGroup sets DeviceRedundancyGroup field to given value.

### HasDeviceRedundancyGroup

`func (o *VirtualServer) HasDeviceRedundancyGroup() bool`

HasDeviceRedundancyGroup returns a boolean if a field has been set.

### SetDeviceRedundancyGroupNil

`func (o *VirtualServer) SetDeviceRedundancyGroupNil(b bool)`

 SetDeviceRedundancyGroupNil sets the value for DeviceRedundancyGroup to be an explicit nil

### UnsetDeviceRedundancyGroup
`func (o *VirtualServer) UnsetDeviceRedundancyGroup()`

UnsetDeviceRedundancyGroup ensures that no value is present for DeviceRedundancyGroup, not even an explicit nil
### GetCloudService

`func (o *VirtualServer) GetCloudService() ApprovalWorkflowUser`

GetCloudService returns the CloudService field if non-nil, zero value otherwise.

### GetCloudServiceOk

`func (o *VirtualServer) GetCloudServiceOk() (*ApprovalWorkflowUser, bool)`

GetCloudServiceOk returns a tuple with the CloudService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudService

`func (o *VirtualServer) SetCloudService(v ApprovalWorkflowUser)`

SetCloudService sets CloudService field to given value.

### HasCloudService

`func (o *VirtualServer) HasCloudService() bool`

HasCloudService returns a boolean if a field has been set.

### SetCloudServiceNil

`func (o *VirtualServer) SetCloudServiceNil(b bool)`

 SetCloudServiceNil sets the value for CloudService to be an explicit nil

### UnsetCloudService
`func (o *VirtualServer) UnsetCloudService()`

UnsetCloudService ensures that no value is present for CloudService, not even an explicit nil
### GetVirtualChassis

`func (o *VirtualServer) GetVirtualChassis() ApprovalWorkflowUser`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *VirtualServer) GetVirtualChassisOk() (*ApprovalWorkflowUser, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *VirtualServer) SetVirtualChassis(v ApprovalWorkflowUser)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *VirtualServer) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *VirtualServer) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *VirtualServer) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetTenant

`func (o *VirtualServer) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VirtualServer) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VirtualServer) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VirtualServer) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VirtualServer) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VirtualServer) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetLoadBalancerPool

`func (o *VirtualServer) GetLoadBalancerPool() ApprovalWorkflowUser`

GetLoadBalancerPool returns the LoadBalancerPool field if non-nil, zero value otherwise.

### GetLoadBalancerPoolOk

`func (o *VirtualServer) GetLoadBalancerPoolOk() (*ApprovalWorkflowUser, bool)`

GetLoadBalancerPoolOk returns a tuple with the LoadBalancerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPool

`func (o *VirtualServer) SetLoadBalancerPool(v ApprovalWorkflowUser)`

SetLoadBalancerPool sets LoadBalancerPool field to given value.

### HasLoadBalancerPool

`func (o *VirtualServer) HasLoadBalancerPool() bool`

HasLoadBalancerPool returns a boolean if a field has been set.

### SetLoadBalancerPoolNil

`func (o *VirtualServer) SetLoadBalancerPoolNil(b bool)`

 SetLoadBalancerPoolNil sets the value for LoadBalancerPool to be an explicit nil

### UnsetLoadBalancerPool
`func (o *VirtualServer) UnsetLoadBalancerPool()`

UnsetLoadBalancerPool ensures that no value is present for LoadBalancerPool, not even an explicit nil
### GetHealthCheckMonitor

`func (o *VirtualServer) GetHealthCheckMonitor() ApprovalWorkflowUser`

GetHealthCheckMonitor returns the HealthCheckMonitor field if non-nil, zero value otherwise.

### GetHealthCheckMonitorOk

`func (o *VirtualServer) GetHealthCheckMonitorOk() (*ApprovalWorkflowUser, bool)`

GetHealthCheckMonitorOk returns a tuple with the HealthCheckMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckMonitor

`func (o *VirtualServer) SetHealthCheckMonitor(v ApprovalWorkflowUser)`

SetHealthCheckMonitor sets HealthCheckMonitor field to given value.

### HasHealthCheckMonitor

`func (o *VirtualServer) HasHealthCheckMonitor() bool`

HasHealthCheckMonitor returns a boolean if a field has been set.

### SetHealthCheckMonitorNil

`func (o *VirtualServer) SetHealthCheckMonitorNil(b bool)`

 SetHealthCheckMonitorNil sets the value for HealthCheckMonitor to be an explicit nil

### UnsetHealthCheckMonitor
`func (o *VirtualServer) UnsetHealthCheckMonitor()`

UnsetHealthCheckMonitor ensures that no value is present for HealthCheckMonitor, not even an explicit nil
### GetCreated

`func (o *VirtualServer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VirtualServer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VirtualServer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VirtualServer) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VirtualServer) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VirtualServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VirtualServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VirtualServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VirtualServer) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VirtualServer) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *VirtualServer) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VirtualServer) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VirtualServer) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *VirtualServer) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualServer) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualServer) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualServer) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *VirtualServer) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualServer) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualServer) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualServer) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


