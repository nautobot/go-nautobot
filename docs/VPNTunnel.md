# VPNTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
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
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewVPNTunnel

`func NewVPNTunnel(objectType string, display string, url string, naturalSlug string, name string, status ApprovalWorkflowStageResponseApprovalWorkflowStage, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *VPNTunnel`

NewVPNTunnel instantiates a new VPNTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNTunnelWithDefaults

`func NewVPNTunnelWithDefaults() *VPNTunnel`

NewVPNTunnelWithDefaults instantiates a new VPNTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNTunnel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNTunnel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNTunnel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNTunnel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VPNTunnel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VPNTunnel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VPNTunnel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VPNTunnel) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VPNTunnel) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VPNTunnel) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VPNTunnel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VPNTunnel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VPNTunnel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VPNTunnel) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VPNTunnel) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VPNTunnel) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetName

`func (o *VPNTunnel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPNTunnel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPNTunnel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VPNTunnel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VPNTunnel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VPNTunnel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VPNTunnel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTunnelId

`func (o *VPNTunnel) GetTunnelId() string`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *VPNTunnel) GetTunnelIdOk() (*string, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *VPNTunnel) SetTunnelId(v string)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *VPNTunnel) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### GetEncapsulation

`func (o *VPNTunnel) GetEncapsulation() BulkWritableVPNTunnelRequestEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *VPNTunnel) GetEncapsulationOk() (*BulkWritableVPNTunnelRequestEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *VPNTunnel) SetEncapsulation(v BulkWritableVPNTunnelRequestEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.

### HasEncapsulation

`func (o *VPNTunnel) HasEncapsulation() bool`

HasEncapsulation returns a boolean if a field has been set.

### GetVpnProfile

`func (o *VPNTunnel) GetVpnProfile() ApprovalWorkflowUser`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *VPNTunnel) GetVpnProfileOk() (*ApprovalWorkflowUser, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *VPNTunnel) SetVpnProfile(v ApprovalWorkflowUser)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *VPNTunnel) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### SetVpnProfileNil

`func (o *VPNTunnel) SetVpnProfileNil(b bool)`

 SetVpnProfileNil sets the value for VpnProfile to be an explicit nil

### UnsetVpnProfile
`func (o *VPNTunnel) UnsetVpnProfile()`

UnsetVpnProfile ensures that no value is present for VpnProfile, not even an explicit nil
### GetVpn

`func (o *VPNTunnel) GetVpn() BulkWritableVPNTunnelRequestVpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *VPNTunnel) GetVpnOk() (*BulkWritableVPNTunnelRequestVpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *VPNTunnel) SetVpn(v BulkWritableVPNTunnelRequestVpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *VPNTunnel) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### SetVpnNil

`func (o *VPNTunnel) SetVpnNil(b bool)`

 SetVpnNil sets the value for Vpn to be an explicit nil

### UnsetVpn
`func (o *VPNTunnel) UnsetVpn()`

UnsetVpn ensures that no value is present for Vpn, not even an explicit nil
### GetRole

`func (o *VPNTunnel) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VPNTunnel) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VPNTunnel) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *VPNTunnel) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VPNTunnel) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VPNTunnel) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetStatus

`func (o *VPNTunnel) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VPNTunnel) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VPNTunnel) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.


### GetSecretsGroup

`func (o *VPNTunnel) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *VPNTunnel) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *VPNTunnel) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *VPNTunnel) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *VPNTunnel) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *VPNTunnel) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetEndpointA

`func (o *VPNTunnel) GetEndpointA() BulkWritableVPNTunnelRequestEndpointA`

GetEndpointA returns the EndpointA field if non-nil, zero value otherwise.

### GetEndpointAOk

`func (o *VPNTunnel) GetEndpointAOk() (*BulkWritableVPNTunnelRequestEndpointA, bool)`

GetEndpointAOk returns a tuple with the EndpointA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointA

`func (o *VPNTunnel) SetEndpointA(v BulkWritableVPNTunnelRequestEndpointA)`

SetEndpointA sets EndpointA field to given value.

### HasEndpointA

`func (o *VPNTunnel) HasEndpointA() bool`

HasEndpointA returns a boolean if a field has been set.

### SetEndpointANil

`func (o *VPNTunnel) SetEndpointANil(b bool)`

 SetEndpointANil sets the value for EndpointA to be an explicit nil

### UnsetEndpointA
`func (o *VPNTunnel) UnsetEndpointA()`

UnsetEndpointA ensures that no value is present for EndpointA, not even an explicit nil
### GetEndpointZ

`func (o *VPNTunnel) GetEndpointZ() BulkWritableVPNTunnelRequestEndpointZ`

GetEndpointZ returns the EndpointZ field if non-nil, zero value otherwise.

### GetEndpointZOk

`func (o *VPNTunnel) GetEndpointZOk() (*BulkWritableVPNTunnelRequestEndpointZ, bool)`

GetEndpointZOk returns a tuple with the EndpointZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointZ

`func (o *VPNTunnel) SetEndpointZ(v BulkWritableVPNTunnelRequestEndpointZ)`

SetEndpointZ sets EndpointZ field to given value.

### HasEndpointZ

`func (o *VPNTunnel) HasEndpointZ() bool`

HasEndpointZ returns a boolean if a field has been set.

### SetEndpointZNil

`func (o *VPNTunnel) SetEndpointZNil(b bool)`

 SetEndpointZNil sets the value for EndpointZ to be an explicit nil

### UnsetEndpointZ
`func (o *VPNTunnel) UnsetEndpointZ()`

UnsetEndpointZ ensures that no value is present for EndpointZ, not even an explicit nil
### GetTenant

`func (o *VPNTunnel) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VPNTunnel) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VPNTunnel) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VPNTunnel) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VPNTunnel) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VPNTunnel) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCreated

`func (o *VPNTunnel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VPNTunnel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VPNTunnel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VPNTunnel) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VPNTunnel) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VPNTunnel) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VPNTunnel) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VPNTunnel) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VPNTunnel) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VPNTunnel) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *VPNTunnel) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VPNTunnel) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VPNTunnel) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *VPNTunnel) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNTunnel) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNTunnel) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNTunnel) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *VPNTunnel) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VPNTunnel) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VPNTunnel) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VPNTunnel) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


