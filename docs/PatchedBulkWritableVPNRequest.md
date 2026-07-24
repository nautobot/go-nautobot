# PatchedBulkWritableVPNRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**VpnId** | Pointer to **string** |  | [optional] 
**VpnProfile** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVPNRequest

`func NewPatchedBulkWritableVPNRequest(id string, ) *PatchedBulkWritableVPNRequest`

NewPatchedBulkWritableVPNRequest instantiates a new PatchedBulkWritableVPNRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVPNRequestWithDefaults

`func NewPatchedBulkWritableVPNRequestWithDefaults() *PatchedBulkWritableVPNRequest`

NewPatchedBulkWritableVPNRequestWithDefaults instantiates a new PatchedBulkWritableVPNRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVPNRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVPNRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVPNRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableVPNRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableVPNRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableVPNRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableVPNRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableVPNRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableVPNRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableVPNRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableVPNRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVpnId

`func (o *PatchedBulkWritableVPNRequest) GetVpnId() string`

GetVpnId returns the VpnId field if non-nil, zero value otherwise.

### GetVpnIdOk

`func (o *PatchedBulkWritableVPNRequest) GetVpnIdOk() (*string, bool)`

GetVpnIdOk returns a tuple with the VpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnId

`func (o *PatchedBulkWritableVPNRequest) SetVpnId(v string)`

SetVpnId sets VpnId field to given value.

### HasVpnId

`func (o *PatchedBulkWritableVPNRequest) HasVpnId() bool`

HasVpnId returns a boolean if a field has been set.

### GetVpnProfile

`func (o *PatchedBulkWritableVPNRequest) GetVpnProfile() ApprovalWorkflowUser`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *PatchedBulkWritableVPNRequest) GetVpnProfileOk() (*ApprovalWorkflowUser, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *PatchedBulkWritableVPNRequest) SetVpnProfile(v ApprovalWorkflowUser)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *PatchedBulkWritableVPNRequest) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### SetVpnProfileNil

`func (o *PatchedBulkWritableVPNRequest) SetVpnProfileNil(b bool)`

 SetVpnProfileNil sets the value for VpnProfile to be an explicit nil

### UnsetVpnProfile
`func (o *PatchedBulkWritableVPNRequest) UnsetVpnProfile()`

UnsetVpnProfile ensures that no value is present for VpnProfile, not even an explicit nil
### GetRole

`func (o *PatchedBulkWritableVPNRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkWritableVPNRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkWritableVPNRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkWritableVPNRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedBulkWritableVPNRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedBulkWritableVPNRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *PatchedBulkWritableVPNRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableVPNRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableVPNRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableVPNRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableVPNRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableVPNRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableVPNRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableVPNRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableVPNRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableVPNRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableVPNRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableVPNRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableVPNRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableVPNRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableVPNRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableVPNRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableVPNRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableVPNRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


