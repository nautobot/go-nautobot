# VPNTunnelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**TunnelId** | Pointer to **string** |  | [optional] 
**Encapsulation** | Pointer to [**BulkWritableVPNTunnelRequestEncapsulation**](BulkWritableVPNTunnelRequestEncapsulation.md) |  | [optional] 
**VpnProfile** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Vpn** | Pointer to [**NullableBulkWritableVPNTunnelRequestVpn**](BulkWritableVPNTunnelRequestVpn.md) |  | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Status** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**SecretsGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**EndpointA** | Pointer to [**NullableBulkWritableVPNTunnelRequestEndpointA**](BulkWritableVPNTunnelRequestEndpointA.md) |  | [optional] 
**EndpointZ** | Pointer to [**NullableBulkWritableVPNTunnelRequestEndpointZ**](BulkWritableVPNTunnelRequestEndpointZ.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewVPNTunnelRequest

`func NewVPNTunnelRequest(name string, status ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *VPNTunnelRequest`

NewVPNTunnelRequest instantiates a new VPNTunnelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNTunnelRequestWithDefaults

`func NewVPNTunnelRequestWithDefaults() *VPNTunnelRequest`

NewVPNTunnelRequestWithDefaults instantiates a new VPNTunnelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNTunnelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNTunnelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNTunnelRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNTunnelRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VPNTunnelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPNTunnelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPNTunnelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VPNTunnelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VPNTunnelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VPNTunnelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VPNTunnelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTunnelId

`func (o *VPNTunnelRequest) GetTunnelId() string`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *VPNTunnelRequest) GetTunnelIdOk() (*string, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *VPNTunnelRequest) SetTunnelId(v string)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *VPNTunnelRequest) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### GetEncapsulation

`func (o *VPNTunnelRequest) GetEncapsulation() BulkWritableVPNTunnelRequestEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *VPNTunnelRequest) GetEncapsulationOk() (*BulkWritableVPNTunnelRequestEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *VPNTunnelRequest) SetEncapsulation(v BulkWritableVPNTunnelRequestEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.

### HasEncapsulation

`func (o *VPNTunnelRequest) HasEncapsulation() bool`

HasEncapsulation returns a boolean if a field has been set.

### GetVpnProfile

`func (o *VPNTunnelRequest) GetVpnProfile() ApprovalWorkflowUser`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *VPNTunnelRequest) GetVpnProfileOk() (*ApprovalWorkflowUser, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *VPNTunnelRequest) SetVpnProfile(v ApprovalWorkflowUser)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *VPNTunnelRequest) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### SetVpnProfileNil

`func (o *VPNTunnelRequest) SetVpnProfileNil(b bool)`

 SetVpnProfileNil sets the value for VpnProfile to be an explicit nil

### UnsetVpnProfile
`func (o *VPNTunnelRequest) UnsetVpnProfile()`

UnsetVpnProfile ensures that no value is present for VpnProfile, not even an explicit nil
### GetVpn

`func (o *VPNTunnelRequest) GetVpn() BulkWritableVPNTunnelRequestVpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *VPNTunnelRequest) GetVpnOk() (*BulkWritableVPNTunnelRequestVpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *VPNTunnelRequest) SetVpn(v BulkWritableVPNTunnelRequestVpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *VPNTunnelRequest) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### SetVpnNil

`func (o *VPNTunnelRequest) SetVpnNil(b bool)`

 SetVpnNil sets the value for Vpn to be an explicit nil

### UnsetVpn
`func (o *VPNTunnelRequest) UnsetVpn()`

UnsetVpn ensures that no value is present for Vpn, not even an explicit nil
### GetRole

`func (o *VPNTunnelRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VPNTunnelRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VPNTunnelRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *VPNTunnelRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VPNTunnelRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VPNTunnelRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetStatus

`func (o *VPNTunnelRequest) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VPNTunnelRequest) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VPNTunnelRequest) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.


### GetSecretsGroup

`func (o *VPNTunnelRequest) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *VPNTunnelRequest) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *VPNTunnelRequest) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *VPNTunnelRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *VPNTunnelRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *VPNTunnelRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetEndpointA

`func (o *VPNTunnelRequest) GetEndpointA() BulkWritableVPNTunnelRequestEndpointA`

GetEndpointA returns the EndpointA field if non-nil, zero value otherwise.

### GetEndpointAOk

`func (o *VPNTunnelRequest) GetEndpointAOk() (*BulkWritableVPNTunnelRequestEndpointA, bool)`

GetEndpointAOk returns a tuple with the EndpointA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointA

`func (o *VPNTunnelRequest) SetEndpointA(v BulkWritableVPNTunnelRequestEndpointA)`

SetEndpointA sets EndpointA field to given value.

### HasEndpointA

`func (o *VPNTunnelRequest) HasEndpointA() bool`

HasEndpointA returns a boolean if a field has been set.

### SetEndpointANil

`func (o *VPNTunnelRequest) SetEndpointANil(b bool)`

 SetEndpointANil sets the value for EndpointA to be an explicit nil

### UnsetEndpointA
`func (o *VPNTunnelRequest) UnsetEndpointA()`

UnsetEndpointA ensures that no value is present for EndpointA, not even an explicit nil
### GetEndpointZ

`func (o *VPNTunnelRequest) GetEndpointZ() BulkWritableVPNTunnelRequestEndpointZ`

GetEndpointZ returns the EndpointZ field if non-nil, zero value otherwise.

### GetEndpointZOk

`func (o *VPNTunnelRequest) GetEndpointZOk() (*BulkWritableVPNTunnelRequestEndpointZ, bool)`

GetEndpointZOk returns a tuple with the EndpointZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointZ

`func (o *VPNTunnelRequest) SetEndpointZ(v BulkWritableVPNTunnelRequestEndpointZ)`

SetEndpointZ sets EndpointZ field to given value.

### HasEndpointZ

`func (o *VPNTunnelRequest) HasEndpointZ() bool`

HasEndpointZ returns a boolean if a field has been set.

### SetEndpointZNil

`func (o *VPNTunnelRequest) SetEndpointZNil(b bool)`

 SetEndpointZNil sets the value for EndpointZ to be an explicit nil

### UnsetEndpointZ
`func (o *VPNTunnelRequest) UnsetEndpointZ()`

UnsetEndpointZ ensures that no value is present for EndpointZ, not even an explicit nil
### GetTenant

`func (o *VPNTunnelRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VPNTunnelRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VPNTunnelRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VPNTunnelRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VPNTunnelRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VPNTunnelRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *VPNTunnelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNTunnelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNTunnelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNTunnelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *VPNTunnelRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *VPNTunnelRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *VPNTunnelRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *VPNTunnelRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *VPNTunnelRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VPNTunnelRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VPNTunnelRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VPNTunnelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


