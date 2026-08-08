# PatchedBulkWritableVPNTunnelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TunnelId** | Pointer to **string** |  | [optional] 
**Encapsulation** | Pointer to [**BulkWritableVPNTunnelRequestEncapsulation**](BulkWritableVPNTunnelRequestEncapsulation.md) |  | [optional] 
**VpnProfile** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Vpn** | Pointer to [**NullableBulkWritableVPNTunnelRequestVpn**](BulkWritableVPNTunnelRequestVpn.md) |  | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Status** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**SecretsGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**EndpointA** | Pointer to [**NullableBulkWritableVPNTunnelRequestEndpointA**](BulkWritableVPNTunnelRequestEndpointA.md) |  | [optional] 
**EndpointZ** | Pointer to [**NullableBulkWritableVPNTunnelRequestEndpointZ**](BulkWritableVPNTunnelRequestEndpointZ.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVPNTunnelRequest

`func NewPatchedBulkWritableVPNTunnelRequest(id string, ) *PatchedBulkWritableVPNTunnelRequest`

NewPatchedBulkWritableVPNTunnelRequest instantiates a new PatchedBulkWritableVPNTunnelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVPNTunnelRequestWithDefaults

`func NewPatchedBulkWritableVPNTunnelRequestWithDefaults() *PatchedBulkWritableVPNTunnelRequest`

NewPatchedBulkWritableVPNTunnelRequestWithDefaults instantiates a new PatchedBulkWritableVPNTunnelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVPNTunnelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVPNTunnelRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableVPNTunnelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableVPNTunnelRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableVPNTunnelRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableVPNTunnelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableVPNTunnelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableVPNTunnelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTunnelId

`func (o *PatchedBulkWritableVPNTunnelRequest) GetTunnelId() string`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetTunnelIdOk() (*string, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *PatchedBulkWritableVPNTunnelRequest) SetTunnelId(v string)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *PatchedBulkWritableVPNTunnelRequest) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### GetEncapsulation

`func (o *PatchedBulkWritableVPNTunnelRequest) GetEncapsulation() BulkWritableVPNTunnelRequestEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetEncapsulationOk() (*BulkWritableVPNTunnelRequestEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *PatchedBulkWritableVPNTunnelRequest) SetEncapsulation(v BulkWritableVPNTunnelRequestEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.

### HasEncapsulation

`func (o *PatchedBulkWritableVPNTunnelRequest) HasEncapsulation() bool`

HasEncapsulation returns a boolean if a field has been set.

### GetVpnProfile

`func (o *PatchedBulkWritableVPNTunnelRequest) GetVpnProfile() ApprovalWorkflowUser`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetVpnProfileOk() (*ApprovalWorkflowUser, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *PatchedBulkWritableVPNTunnelRequest) SetVpnProfile(v ApprovalWorkflowUser)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *PatchedBulkWritableVPNTunnelRequest) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### SetVpnProfileNil

`func (o *PatchedBulkWritableVPNTunnelRequest) SetVpnProfileNil(b bool)`

 SetVpnProfileNil sets the value for VpnProfile to be an explicit nil

### UnsetVpnProfile
`func (o *PatchedBulkWritableVPNTunnelRequest) UnsetVpnProfile()`

UnsetVpnProfile ensures that no value is present for VpnProfile, not even an explicit nil
### GetVpn

`func (o *PatchedBulkWritableVPNTunnelRequest) GetVpn() BulkWritableVPNTunnelRequestVpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetVpnOk() (*BulkWritableVPNTunnelRequestVpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *PatchedBulkWritableVPNTunnelRequest) SetVpn(v BulkWritableVPNTunnelRequestVpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *PatchedBulkWritableVPNTunnelRequest) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### SetVpnNil

`func (o *PatchedBulkWritableVPNTunnelRequest) SetVpnNil(b bool)`

 SetVpnNil sets the value for Vpn to be an explicit nil

### UnsetVpn
`func (o *PatchedBulkWritableVPNTunnelRequest) UnsetVpn()`

UnsetVpn ensures that no value is present for Vpn, not even an explicit nil
### GetRole

`func (o *PatchedBulkWritableVPNTunnelRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkWritableVPNTunnelRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkWritableVPNTunnelRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedBulkWritableVPNTunnelRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedBulkWritableVPNTunnelRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetStatus

`func (o *PatchedBulkWritableVPNTunnelRequest) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableVPNTunnelRequest) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableVPNTunnelRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *PatchedBulkWritableVPNTunnelRequest) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedBulkWritableVPNTunnelRequest) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedBulkWritableVPNTunnelRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedBulkWritableVPNTunnelRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedBulkWritableVPNTunnelRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetEndpointA

`func (o *PatchedBulkWritableVPNTunnelRequest) GetEndpointA() BulkWritableVPNTunnelRequestEndpointA`

GetEndpointA returns the EndpointA field if non-nil, zero value otherwise.

### GetEndpointAOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetEndpointAOk() (*BulkWritableVPNTunnelRequestEndpointA, bool)`

GetEndpointAOk returns a tuple with the EndpointA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointA

`func (o *PatchedBulkWritableVPNTunnelRequest) SetEndpointA(v BulkWritableVPNTunnelRequestEndpointA)`

SetEndpointA sets EndpointA field to given value.

### HasEndpointA

`func (o *PatchedBulkWritableVPNTunnelRequest) HasEndpointA() bool`

HasEndpointA returns a boolean if a field has been set.

### SetEndpointANil

`func (o *PatchedBulkWritableVPNTunnelRequest) SetEndpointANil(b bool)`

 SetEndpointANil sets the value for EndpointA to be an explicit nil

### UnsetEndpointA
`func (o *PatchedBulkWritableVPNTunnelRequest) UnsetEndpointA()`

UnsetEndpointA ensures that no value is present for EndpointA, not even an explicit nil
### GetEndpointZ

`func (o *PatchedBulkWritableVPNTunnelRequest) GetEndpointZ() BulkWritableVPNTunnelRequestEndpointZ`

GetEndpointZ returns the EndpointZ field if non-nil, zero value otherwise.

### GetEndpointZOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetEndpointZOk() (*BulkWritableVPNTunnelRequestEndpointZ, bool)`

GetEndpointZOk returns a tuple with the EndpointZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointZ

`func (o *PatchedBulkWritableVPNTunnelRequest) SetEndpointZ(v BulkWritableVPNTunnelRequestEndpointZ)`

SetEndpointZ sets EndpointZ field to given value.

### HasEndpointZ

`func (o *PatchedBulkWritableVPNTunnelRequest) HasEndpointZ() bool`

HasEndpointZ returns a boolean if a field has been set.

### SetEndpointZNil

`func (o *PatchedBulkWritableVPNTunnelRequest) SetEndpointZNil(b bool)`

 SetEndpointZNil sets the value for EndpointZ to be an explicit nil

### UnsetEndpointZ
`func (o *PatchedBulkWritableVPNTunnelRequest) UnsetEndpointZ()`

UnsetEndpointZ ensures that no value is present for EndpointZ, not even an explicit nil
### GetTenant

`func (o *PatchedBulkWritableVPNTunnelRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableVPNTunnelRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableVPNTunnelRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableVPNTunnelRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableVPNTunnelRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableVPNTunnelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableVPNTunnelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableVPNTunnelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableVPNTunnelRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableVPNTunnelRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableVPNTunnelRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableVPNTunnelRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableVPNTunnelRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableVPNTunnelRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableVPNTunnelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


