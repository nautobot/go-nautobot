# PatchedVPNTunnelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TunnelId** | Pointer to **string** |  | [optional] 
**Encapsulation** | Pointer to [**BulkWritableVPNTunnelRequestEncapsulation**](BulkWritableVPNTunnelRequestEncapsulation.md) |  | [optional] 
**VpnProfile** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Vpn** | Pointer to [**NullableBulkWritableVPNTunnelRequestVpn**](BulkWritableVPNTunnelRequestVpn.md) |  | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**SecretsGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**EndpointA** | Pointer to [**NullableBulkWritableVPNTunnelRequestEndpointA**](BulkWritableVPNTunnelRequestEndpointA.md) |  | [optional] 
**EndpointZ** | Pointer to [**NullableBulkWritableVPNTunnelRequestEndpointZ**](BulkWritableVPNTunnelRequestEndpointZ.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedVPNTunnelRequest

`func NewPatchedVPNTunnelRequest() *PatchedVPNTunnelRequest`

NewPatchedVPNTunnelRequest instantiates a new PatchedVPNTunnelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVPNTunnelRequestWithDefaults

`func NewPatchedVPNTunnelRequestWithDefaults() *PatchedVPNTunnelRequest`

NewPatchedVPNTunnelRequestWithDefaults instantiates a new PatchedVPNTunnelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedVPNTunnelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedVPNTunnelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedVPNTunnelRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedVPNTunnelRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedVPNTunnelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedVPNTunnelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedVPNTunnelRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedVPNTunnelRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedVPNTunnelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedVPNTunnelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedVPNTunnelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedVPNTunnelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTunnelId

`func (o *PatchedVPNTunnelRequest) GetTunnelId() string`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *PatchedVPNTunnelRequest) GetTunnelIdOk() (*string, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *PatchedVPNTunnelRequest) SetTunnelId(v string)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *PatchedVPNTunnelRequest) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### GetEncapsulation

`func (o *PatchedVPNTunnelRequest) GetEncapsulation() BulkWritableVPNTunnelRequestEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *PatchedVPNTunnelRequest) GetEncapsulationOk() (*BulkWritableVPNTunnelRequestEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *PatchedVPNTunnelRequest) SetEncapsulation(v BulkWritableVPNTunnelRequestEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.

### HasEncapsulation

`func (o *PatchedVPNTunnelRequest) HasEncapsulation() bool`

HasEncapsulation returns a boolean if a field has been set.

### GetVpnProfile

`func (o *PatchedVPNTunnelRequest) GetVpnProfile() ApprovalWorkflowUser`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *PatchedVPNTunnelRequest) GetVpnProfileOk() (*ApprovalWorkflowUser, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *PatchedVPNTunnelRequest) SetVpnProfile(v ApprovalWorkflowUser)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *PatchedVPNTunnelRequest) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### SetVpnProfileNil

`func (o *PatchedVPNTunnelRequest) SetVpnProfileNil(b bool)`

 SetVpnProfileNil sets the value for VpnProfile to be an explicit nil

### UnsetVpnProfile
`func (o *PatchedVPNTunnelRequest) UnsetVpnProfile()`

UnsetVpnProfile ensures that no value is present for VpnProfile, not even an explicit nil
### GetVpn

`func (o *PatchedVPNTunnelRequest) GetVpn() BulkWritableVPNTunnelRequestVpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *PatchedVPNTunnelRequest) GetVpnOk() (*BulkWritableVPNTunnelRequestVpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *PatchedVPNTunnelRequest) SetVpn(v BulkWritableVPNTunnelRequestVpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *PatchedVPNTunnelRequest) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### SetVpnNil

`func (o *PatchedVPNTunnelRequest) SetVpnNil(b bool)`

 SetVpnNil sets the value for Vpn to be an explicit nil

### UnsetVpn
`func (o *PatchedVPNTunnelRequest) UnsetVpn()`

UnsetVpn ensures that no value is present for Vpn, not even an explicit nil
### GetRole

`func (o *PatchedVPNTunnelRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedVPNTunnelRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedVPNTunnelRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedVPNTunnelRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedVPNTunnelRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedVPNTunnelRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetStatus

`func (o *PatchedVPNTunnelRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedVPNTunnelRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedVPNTunnelRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedVPNTunnelRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *PatchedVPNTunnelRequest) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedVPNTunnelRequest) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedVPNTunnelRequest) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedVPNTunnelRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedVPNTunnelRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedVPNTunnelRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetEndpointA

`func (o *PatchedVPNTunnelRequest) GetEndpointA() BulkWritableVPNTunnelRequestEndpointA`

GetEndpointA returns the EndpointA field if non-nil, zero value otherwise.

### GetEndpointAOk

`func (o *PatchedVPNTunnelRequest) GetEndpointAOk() (*BulkWritableVPNTunnelRequestEndpointA, bool)`

GetEndpointAOk returns a tuple with the EndpointA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointA

`func (o *PatchedVPNTunnelRequest) SetEndpointA(v BulkWritableVPNTunnelRequestEndpointA)`

SetEndpointA sets EndpointA field to given value.

### HasEndpointA

`func (o *PatchedVPNTunnelRequest) HasEndpointA() bool`

HasEndpointA returns a boolean if a field has been set.

### SetEndpointANil

`func (o *PatchedVPNTunnelRequest) SetEndpointANil(b bool)`

 SetEndpointANil sets the value for EndpointA to be an explicit nil

### UnsetEndpointA
`func (o *PatchedVPNTunnelRequest) UnsetEndpointA()`

UnsetEndpointA ensures that no value is present for EndpointA, not even an explicit nil
### GetEndpointZ

`func (o *PatchedVPNTunnelRequest) GetEndpointZ() BulkWritableVPNTunnelRequestEndpointZ`

GetEndpointZ returns the EndpointZ field if non-nil, zero value otherwise.

### GetEndpointZOk

`func (o *PatchedVPNTunnelRequest) GetEndpointZOk() (*BulkWritableVPNTunnelRequestEndpointZ, bool)`

GetEndpointZOk returns a tuple with the EndpointZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointZ

`func (o *PatchedVPNTunnelRequest) SetEndpointZ(v BulkWritableVPNTunnelRequestEndpointZ)`

SetEndpointZ sets EndpointZ field to given value.

### HasEndpointZ

`func (o *PatchedVPNTunnelRequest) HasEndpointZ() bool`

HasEndpointZ returns a boolean if a field has been set.

### SetEndpointZNil

`func (o *PatchedVPNTunnelRequest) SetEndpointZNil(b bool)`

 SetEndpointZNil sets the value for EndpointZ to be an explicit nil

### UnsetEndpointZ
`func (o *PatchedVPNTunnelRequest) UnsetEndpointZ()`

UnsetEndpointZ ensures that no value is present for EndpointZ, not even an explicit nil
### GetTenant

`func (o *PatchedVPNTunnelRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedVPNTunnelRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedVPNTunnelRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedVPNTunnelRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedVPNTunnelRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedVPNTunnelRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedVPNTunnelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedVPNTunnelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedVPNTunnelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedVPNTunnelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedVPNTunnelRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedVPNTunnelRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedVPNTunnelRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedVPNTunnelRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedVPNTunnelRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedVPNTunnelRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedVPNTunnelRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedVPNTunnelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


