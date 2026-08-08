# VPNTunnelEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**SourceFqdn** | Pointer to **string** | Mutually Exclusive with Source IP Address | [optional] 
**Device** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**SourceInterface** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**SourceIpaddress** | Pointer to [**NullableSourceIPAddress**](SourceIPAddress.md) |  | [optional] 
**TunnelInterface** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**VpnProfile** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewVPNTunnelEndpoint

`func NewVPNTunnelEndpoint(objectType string, display string, url string, naturalSlug string, name string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *VPNTunnelEndpoint`

NewVPNTunnelEndpoint instantiates a new VPNTunnelEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNTunnelEndpointWithDefaults

`func NewVPNTunnelEndpointWithDefaults() *VPNTunnelEndpoint`

NewVPNTunnelEndpointWithDefaults instantiates a new VPNTunnelEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNTunnelEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNTunnelEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNTunnelEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNTunnelEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VPNTunnelEndpoint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VPNTunnelEndpoint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VPNTunnelEndpoint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VPNTunnelEndpoint) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VPNTunnelEndpoint) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VPNTunnelEndpoint) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VPNTunnelEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VPNTunnelEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VPNTunnelEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VPNTunnelEndpoint) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VPNTunnelEndpoint) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VPNTunnelEndpoint) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetName

`func (o *VPNTunnelEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPNTunnelEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPNTunnelEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetSourceFqdn

`func (o *VPNTunnelEndpoint) GetSourceFqdn() string`

GetSourceFqdn returns the SourceFqdn field if non-nil, zero value otherwise.

### GetSourceFqdnOk

`func (o *VPNTunnelEndpoint) GetSourceFqdnOk() (*string, bool)`

GetSourceFqdnOk returns a tuple with the SourceFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFqdn

`func (o *VPNTunnelEndpoint) SetSourceFqdn(v string)`

SetSourceFqdn sets SourceFqdn field to given value.

### HasSourceFqdn

`func (o *VPNTunnelEndpoint) HasSourceFqdn() bool`

HasSourceFqdn returns a boolean if a field has been set.

### GetDevice

`func (o *VPNTunnelEndpoint) GetDevice() ApprovalWorkflowUser`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VPNTunnelEndpoint) GetDeviceOk() (*ApprovalWorkflowUser, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VPNTunnelEndpoint) SetDevice(v ApprovalWorkflowUser)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VPNTunnelEndpoint) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *VPNTunnelEndpoint) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *VPNTunnelEndpoint) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetSourceInterface

`func (o *VPNTunnelEndpoint) GetSourceInterface() ApprovalWorkflowUser`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *VPNTunnelEndpoint) GetSourceInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *VPNTunnelEndpoint) SetSourceInterface(v ApprovalWorkflowUser)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *VPNTunnelEndpoint) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.

### SetSourceInterfaceNil

`func (o *VPNTunnelEndpoint) SetSourceInterfaceNil(b bool)`

 SetSourceInterfaceNil sets the value for SourceInterface to be an explicit nil

### UnsetSourceInterface
`func (o *VPNTunnelEndpoint) UnsetSourceInterface()`

UnsetSourceInterface ensures that no value is present for SourceInterface, not even an explicit nil
### GetSourceIpaddress

`func (o *VPNTunnelEndpoint) GetSourceIpaddress() SourceIPAddress`

GetSourceIpaddress returns the SourceIpaddress field if non-nil, zero value otherwise.

### GetSourceIpaddressOk

`func (o *VPNTunnelEndpoint) GetSourceIpaddressOk() (*SourceIPAddress, bool)`

GetSourceIpaddressOk returns a tuple with the SourceIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpaddress

`func (o *VPNTunnelEndpoint) SetSourceIpaddress(v SourceIPAddress)`

SetSourceIpaddress sets SourceIpaddress field to given value.

### HasSourceIpaddress

`func (o *VPNTunnelEndpoint) HasSourceIpaddress() bool`

HasSourceIpaddress returns a boolean if a field has been set.

### SetSourceIpaddressNil

`func (o *VPNTunnelEndpoint) SetSourceIpaddressNil(b bool)`

 SetSourceIpaddressNil sets the value for SourceIpaddress to be an explicit nil

### UnsetSourceIpaddress
`func (o *VPNTunnelEndpoint) UnsetSourceIpaddress()`

UnsetSourceIpaddress ensures that no value is present for SourceIpaddress, not even an explicit nil
### GetTunnelInterface

`func (o *VPNTunnelEndpoint) GetTunnelInterface() ApprovalWorkflowUser`

GetTunnelInterface returns the TunnelInterface field if non-nil, zero value otherwise.

### GetTunnelInterfaceOk

`func (o *VPNTunnelEndpoint) GetTunnelInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetTunnelInterfaceOk returns a tuple with the TunnelInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelInterface

`func (o *VPNTunnelEndpoint) SetTunnelInterface(v ApprovalWorkflowUser)`

SetTunnelInterface sets TunnelInterface field to given value.

### HasTunnelInterface

`func (o *VPNTunnelEndpoint) HasTunnelInterface() bool`

HasTunnelInterface returns a boolean if a field has been set.

### SetTunnelInterfaceNil

`func (o *VPNTunnelEndpoint) SetTunnelInterfaceNil(b bool)`

 SetTunnelInterfaceNil sets the value for TunnelInterface to be an explicit nil

### UnsetTunnelInterface
`func (o *VPNTunnelEndpoint) UnsetTunnelInterface()`

UnsetTunnelInterface ensures that no value is present for TunnelInterface, not even an explicit nil
### GetVpnProfile

`func (o *VPNTunnelEndpoint) GetVpnProfile() ApprovalWorkflowUser`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *VPNTunnelEndpoint) GetVpnProfileOk() (*ApprovalWorkflowUser, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *VPNTunnelEndpoint) SetVpnProfile(v ApprovalWorkflowUser)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *VPNTunnelEndpoint) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### SetVpnProfileNil

`func (o *VPNTunnelEndpoint) SetVpnProfileNil(b bool)`

 SetVpnProfileNil sets the value for VpnProfile to be an explicit nil

### UnsetVpnProfile
`func (o *VPNTunnelEndpoint) UnsetVpnProfile()`

UnsetVpnProfile ensures that no value is present for VpnProfile, not even an explicit nil
### GetRole

`func (o *VPNTunnelEndpoint) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VPNTunnelEndpoint) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VPNTunnelEndpoint) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *VPNTunnelEndpoint) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VPNTunnelEndpoint) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VPNTunnelEndpoint) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *VPNTunnelEndpoint) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VPNTunnelEndpoint) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VPNTunnelEndpoint) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VPNTunnelEndpoint) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VPNTunnelEndpoint) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VPNTunnelEndpoint) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCreated

`func (o *VPNTunnelEndpoint) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VPNTunnelEndpoint) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VPNTunnelEndpoint) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VPNTunnelEndpoint) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VPNTunnelEndpoint) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VPNTunnelEndpoint) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VPNTunnelEndpoint) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VPNTunnelEndpoint) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VPNTunnelEndpoint) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VPNTunnelEndpoint) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *VPNTunnelEndpoint) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VPNTunnelEndpoint) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VPNTunnelEndpoint) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *VPNTunnelEndpoint) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNTunnelEndpoint) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNTunnelEndpoint) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNTunnelEndpoint) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *VPNTunnelEndpoint) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VPNTunnelEndpoint) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VPNTunnelEndpoint) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VPNTunnelEndpoint) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


