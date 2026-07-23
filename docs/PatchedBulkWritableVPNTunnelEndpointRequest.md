# PatchedBulkWritableVPNTunnelEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SourceFqdn** | Pointer to **string** | Mutually Exclusive with Source IP Address | [optional] 
**Device** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**SourceInterface** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**SourceIpaddress** | Pointer to [**NullableSourceIPAddress**](SourceIPAddress.md) |  | [optional] 
**TunnelInterface** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**VpnProfile** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ProtectedPrefixes** | Pointer to [**[]ProtectedPrefixes**](ProtectedPrefixes.md) |  | [optional] 
**ProtectedPrefixesDg** | Pointer to [**[]ProtectedPrefixesDynamicGroup**](ProtectedPrefixesDynamicGroup.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVPNTunnelEndpointRequest

`func NewPatchedBulkWritableVPNTunnelEndpointRequest(id string, ) *PatchedBulkWritableVPNTunnelEndpointRequest`

NewPatchedBulkWritableVPNTunnelEndpointRequest instantiates a new PatchedBulkWritableVPNTunnelEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVPNTunnelEndpointRequestWithDefaults

`func NewPatchedBulkWritableVPNTunnelEndpointRequestWithDefaults() *PatchedBulkWritableVPNTunnelEndpointRequest`

NewPatchedBulkWritableVPNTunnelEndpointRequestWithDefaults instantiates a new PatchedBulkWritableVPNTunnelEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetId(v string)`

SetId sets Id field to given value.


### GetSourceFqdn

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetSourceFqdn() string`

GetSourceFqdn returns the SourceFqdn field if non-nil, zero value otherwise.

### GetSourceFqdnOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetSourceFqdnOk() (*string, bool)`

GetSourceFqdnOk returns a tuple with the SourceFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFqdn

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetSourceFqdn(v string)`

SetSourceFqdn sets SourceFqdn field to given value.

### HasSourceFqdn

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasSourceFqdn() bool`

HasSourceFqdn returns a boolean if a field has been set.

### GetDevice

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetDevice() ApprovalWorkflowUser`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetDeviceOk() (*ApprovalWorkflowUser, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetDevice(v ApprovalWorkflowUser)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetSourceInterface

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetSourceInterface() ApprovalWorkflowUser`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetSourceInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetSourceInterface(v ApprovalWorkflowUser)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.

### SetSourceInterfaceNil

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetSourceInterfaceNil(b bool)`

 SetSourceInterfaceNil sets the value for SourceInterface to be an explicit nil

### UnsetSourceInterface
`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) UnsetSourceInterface()`

UnsetSourceInterface ensures that no value is present for SourceInterface, not even an explicit nil
### GetSourceIpaddress

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetSourceIpaddress() SourceIPAddress`

GetSourceIpaddress returns the SourceIpaddress field if non-nil, zero value otherwise.

### GetSourceIpaddressOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetSourceIpaddressOk() (*SourceIPAddress, bool)`

GetSourceIpaddressOk returns a tuple with the SourceIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpaddress

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetSourceIpaddress(v SourceIPAddress)`

SetSourceIpaddress sets SourceIpaddress field to given value.

### HasSourceIpaddress

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasSourceIpaddress() bool`

HasSourceIpaddress returns a boolean if a field has been set.

### SetSourceIpaddressNil

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetSourceIpaddressNil(b bool)`

 SetSourceIpaddressNil sets the value for SourceIpaddress to be an explicit nil

### UnsetSourceIpaddress
`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) UnsetSourceIpaddress()`

UnsetSourceIpaddress ensures that no value is present for SourceIpaddress, not even an explicit nil
### GetTunnelInterface

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetTunnelInterface() ApprovalWorkflowUser`

GetTunnelInterface returns the TunnelInterface field if non-nil, zero value otherwise.

### GetTunnelInterfaceOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetTunnelInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetTunnelInterfaceOk returns a tuple with the TunnelInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelInterface

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetTunnelInterface(v ApprovalWorkflowUser)`

SetTunnelInterface sets TunnelInterface field to given value.

### HasTunnelInterface

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasTunnelInterface() bool`

HasTunnelInterface returns a boolean if a field has been set.

### SetTunnelInterfaceNil

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetTunnelInterfaceNil(b bool)`

 SetTunnelInterfaceNil sets the value for TunnelInterface to be an explicit nil

### UnsetTunnelInterface
`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) UnsetTunnelInterface()`

UnsetTunnelInterface ensures that no value is present for TunnelInterface, not even an explicit nil
### GetVpnProfile

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetVpnProfile() ApprovalWorkflowUser`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetVpnProfileOk() (*ApprovalWorkflowUser, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetVpnProfile(v ApprovalWorkflowUser)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### SetVpnProfileNil

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetVpnProfileNil(b bool)`

 SetVpnProfileNil sets the value for VpnProfile to be an explicit nil

### UnsetVpnProfile
`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) UnsetVpnProfile()`

UnsetVpnProfile ensures that no value is present for VpnProfile, not even an explicit nil
### GetRole

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetProtectedPrefixes

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetProtectedPrefixes() []ProtectedPrefixes`

GetProtectedPrefixes returns the ProtectedPrefixes field if non-nil, zero value otherwise.

### GetProtectedPrefixesOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetProtectedPrefixesOk() (*[]ProtectedPrefixes, bool)`

GetProtectedPrefixesOk returns a tuple with the ProtectedPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedPrefixes

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetProtectedPrefixes(v []ProtectedPrefixes)`

SetProtectedPrefixes sets ProtectedPrefixes field to given value.

### HasProtectedPrefixes

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasProtectedPrefixes() bool`

HasProtectedPrefixes returns a boolean if a field has been set.

### GetProtectedPrefixesDg

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetProtectedPrefixesDg() []ProtectedPrefixesDynamicGroup`

GetProtectedPrefixesDg returns the ProtectedPrefixesDg field if non-nil, zero value otherwise.

### GetProtectedPrefixesDgOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetProtectedPrefixesDgOk() (*[]ProtectedPrefixesDynamicGroup, bool)`

GetProtectedPrefixesDgOk returns a tuple with the ProtectedPrefixesDg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedPrefixesDg

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetProtectedPrefixesDg(v []ProtectedPrefixesDynamicGroup)`

SetProtectedPrefixesDg sets ProtectedPrefixesDg field to given value.

### HasProtectedPrefixesDg

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasProtectedPrefixesDg() bool`

HasProtectedPrefixesDg returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableVPNTunnelEndpointRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


