# BulkWritableVPNTunnelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewBulkWritableVPNTunnelRequest

`func NewBulkWritableVPNTunnelRequest(id string, name string, status ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *BulkWritableVPNTunnelRequest`

NewBulkWritableVPNTunnelRequest instantiates a new BulkWritableVPNTunnelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVPNTunnelRequestWithDefaults

`func NewBulkWritableVPNTunnelRequestWithDefaults() *BulkWritableVPNTunnelRequest`

NewBulkWritableVPNTunnelRequestWithDefaults instantiates a new BulkWritableVPNTunnelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVPNTunnelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVPNTunnelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVPNTunnelRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableVPNTunnelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableVPNTunnelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableVPNTunnelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableVPNTunnelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableVPNTunnelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableVPNTunnelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableVPNTunnelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTunnelId

`func (o *BulkWritableVPNTunnelRequest) GetTunnelId() string`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *BulkWritableVPNTunnelRequest) GetTunnelIdOk() (*string, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *BulkWritableVPNTunnelRequest) SetTunnelId(v string)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *BulkWritableVPNTunnelRequest) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### GetEncapsulation

`func (o *BulkWritableVPNTunnelRequest) GetEncapsulation() BulkWritableVPNTunnelRequestEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *BulkWritableVPNTunnelRequest) GetEncapsulationOk() (*BulkWritableVPNTunnelRequestEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *BulkWritableVPNTunnelRequest) SetEncapsulation(v BulkWritableVPNTunnelRequestEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.

### HasEncapsulation

`func (o *BulkWritableVPNTunnelRequest) HasEncapsulation() bool`

HasEncapsulation returns a boolean if a field has been set.

### GetVpnProfile

`func (o *BulkWritableVPNTunnelRequest) GetVpnProfile() ApprovalWorkflowUser`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *BulkWritableVPNTunnelRequest) GetVpnProfileOk() (*ApprovalWorkflowUser, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *BulkWritableVPNTunnelRequest) SetVpnProfile(v ApprovalWorkflowUser)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *BulkWritableVPNTunnelRequest) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### SetVpnProfileNil

`func (o *BulkWritableVPNTunnelRequest) SetVpnProfileNil(b bool)`

 SetVpnProfileNil sets the value for VpnProfile to be an explicit nil

### UnsetVpnProfile
`func (o *BulkWritableVPNTunnelRequest) UnsetVpnProfile()`

UnsetVpnProfile ensures that no value is present for VpnProfile, not even an explicit nil
### GetVpn

`func (o *BulkWritableVPNTunnelRequest) GetVpn() BulkWritableVPNTunnelRequestVpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *BulkWritableVPNTunnelRequest) GetVpnOk() (*BulkWritableVPNTunnelRequestVpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *BulkWritableVPNTunnelRequest) SetVpn(v BulkWritableVPNTunnelRequestVpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *BulkWritableVPNTunnelRequest) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### SetVpnNil

`func (o *BulkWritableVPNTunnelRequest) SetVpnNil(b bool)`

 SetVpnNil sets the value for Vpn to be an explicit nil

### UnsetVpn
`func (o *BulkWritableVPNTunnelRequest) UnsetVpn()`

UnsetVpn ensures that no value is present for Vpn, not even an explicit nil
### GetRole

`func (o *BulkWritableVPNTunnelRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkWritableVPNTunnelRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkWritableVPNTunnelRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkWritableVPNTunnelRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkWritableVPNTunnelRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkWritableVPNTunnelRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetStatus

`func (o *BulkWritableVPNTunnelRequest) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableVPNTunnelRequest) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableVPNTunnelRequest) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.


### GetSecretsGroup

`func (o *BulkWritableVPNTunnelRequest) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *BulkWritableVPNTunnelRequest) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *BulkWritableVPNTunnelRequest) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *BulkWritableVPNTunnelRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *BulkWritableVPNTunnelRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *BulkWritableVPNTunnelRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetEndpointA

`func (o *BulkWritableVPNTunnelRequest) GetEndpointA() BulkWritableVPNTunnelRequestEndpointA`

GetEndpointA returns the EndpointA field if non-nil, zero value otherwise.

### GetEndpointAOk

`func (o *BulkWritableVPNTunnelRequest) GetEndpointAOk() (*BulkWritableVPNTunnelRequestEndpointA, bool)`

GetEndpointAOk returns a tuple with the EndpointA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointA

`func (o *BulkWritableVPNTunnelRequest) SetEndpointA(v BulkWritableVPNTunnelRequestEndpointA)`

SetEndpointA sets EndpointA field to given value.

### HasEndpointA

`func (o *BulkWritableVPNTunnelRequest) HasEndpointA() bool`

HasEndpointA returns a boolean if a field has been set.

### SetEndpointANil

`func (o *BulkWritableVPNTunnelRequest) SetEndpointANil(b bool)`

 SetEndpointANil sets the value for EndpointA to be an explicit nil

### UnsetEndpointA
`func (o *BulkWritableVPNTunnelRequest) UnsetEndpointA()`

UnsetEndpointA ensures that no value is present for EndpointA, not even an explicit nil
### GetEndpointZ

`func (o *BulkWritableVPNTunnelRequest) GetEndpointZ() BulkWritableVPNTunnelRequestEndpointZ`

GetEndpointZ returns the EndpointZ field if non-nil, zero value otherwise.

### GetEndpointZOk

`func (o *BulkWritableVPNTunnelRequest) GetEndpointZOk() (*BulkWritableVPNTunnelRequestEndpointZ, bool)`

GetEndpointZOk returns a tuple with the EndpointZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointZ

`func (o *BulkWritableVPNTunnelRequest) SetEndpointZ(v BulkWritableVPNTunnelRequestEndpointZ)`

SetEndpointZ sets EndpointZ field to given value.

### HasEndpointZ

`func (o *BulkWritableVPNTunnelRequest) HasEndpointZ() bool`

HasEndpointZ returns a boolean if a field has been set.

### SetEndpointZNil

`func (o *BulkWritableVPNTunnelRequest) SetEndpointZNil(b bool)`

 SetEndpointZNil sets the value for EndpointZ to be an explicit nil

### UnsetEndpointZ
`func (o *BulkWritableVPNTunnelRequest) UnsetEndpointZ()`

UnsetEndpointZ ensures that no value is present for EndpointZ, not even an explicit nil
### GetTenant

`func (o *BulkWritableVPNTunnelRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableVPNTunnelRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableVPNTunnelRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableVPNTunnelRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableVPNTunnelRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableVPNTunnelRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableVPNTunnelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableVPNTunnelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableVPNTunnelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableVPNTunnelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableVPNTunnelRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableVPNTunnelRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableVPNTunnelRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableVPNTunnelRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableVPNTunnelRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableVPNTunnelRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableVPNTunnelRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableVPNTunnelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


