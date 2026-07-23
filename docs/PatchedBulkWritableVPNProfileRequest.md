# PatchedBulkWritableVPNProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**KeepaliveEnabled** | Pointer to **bool** |  | [optional] 
**KeepaliveInterval** | Pointer to **NullableInt32** |  | [optional] 
**KeepaliveRetries** | Pointer to **NullableInt32** |  | [optional] 
**NatTraversal** | Pointer to **bool** |  | [optional] 
**ExtraOptions** | Pointer to **interface{}** | Additional options specific to the VPN technology and/or implementation | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**SecretsGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVPNProfileRequest

`func NewPatchedBulkWritableVPNProfileRequest(id string, ) *PatchedBulkWritableVPNProfileRequest`

NewPatchedBulkWritableVPNProfileRequest instantiates a new PatchedBulkWritableVPNProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVPNProfileRequestWithDefaults

`func NewPatchedBulkWritableVPNProfileRequestWithDefaults() *PatchedBulkWritableVPNProfileRequest`

NewPatchedBulkWritableVPNProfileRequestWithDefaults instantiates a new PatchedBulkWritableVPNProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVPNProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVPNProfileRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableVPNProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableVPNProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableVPNProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableVPNProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableVPNProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableVPNProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKeepaliveEnabled

`func (o *PatchedBulkWritableVPNProfileRequest) GetKeepaliveEnabled() bool`

GetKeepaliveEnabled returns the KeepaliveEnabled field if non-nil, zero value otherwise.

### GetKeepaliveEnabledOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetKeepaliveEnabledOk() (*bool, bool)`

GetKeepaliveEnabledOk returns a tuple with the KeepaliveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveEnabled

`func (o *PatchedBulkWritableVPNProfileRequest) SetKeepaliveEnabled(v bool)`

SetKeepaliveEnabled sets KeepaliveEnabled field to given value.

### HasKeepaliveEnabled

`func (o *PatchedBulkWritableVPNProfileRequest) HasKeepaliveEnabled() bool`

HasKeepaliveEnabled returns a boolean if a field has been set.

### GetKeepaliveInterval

`func (o *PatchedBulkWritableVPNProfileRequest) GetKeepaliveInterval() int32`

GetKeepaliveInterval returns the KeepaliveInterval field if non-nil, zero value otherwise.

### GetKeepaliveIntervalOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetKeepaliveIntervalOk() (*int32, bool)`

GetKeepaliveIntervalOk returns a tuple with the KeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveInterval

`func (o *PatchedBulkWritableVPNProfileRequest) SetKeepaliveInterval(v int32)`

SetKeepaliveInterval sets KeepaliveInterval field to given value.

### HasKeepaliveInterval

`func (o *PatchedBulkWritableVPNProfileRequest) HasKeepaliveInterval() bool`

HasKeepaliveInterval returns a boolean if a field has been set.

### SetKeepaliveIntervalNil

`func (o *PatchedBulkWritableVPNProfileRequest) SetKeepaliveIntervalNil(b bool)`

 SetKeepaliveIntervalNil sets the value for KeepaliveInterval to be an explicit nil

### UnsetKeepaliveInterval
`func (o *PatchedBulkWritableVPNProfileRequest) UnsetKeepaliveInterval()`

UnsetKeepaliveInterval ensures that no value is present for KeepaliveInterval, not even an explicit nil
### GetKeepaliveRetries

`func (o *PatchedBulkWritableVPNProfileRequest) GetKeepaliveRetries() int32`

GetKeepaliveRetries returns the KeepaliveRetries field if non-nil, zero value otherwise.

### GetKeepaliveRetriesOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetKeepaliveRetriesOk() (*int32, bool)`

GetKeepaliveRetriesOk returns a tuple with the KeepaliveRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveRetries

`func (o *PatchedBulkWritableVPNProfileRequest) SetKeepaliveRetries(v int32)`

SetKeepaliveRetries sets KeepaliveRetries field to given value.

### HasKeepaliveRetries

`func (o *PatchedBulkWritableVPNProfileRequest) HasKeepaliveRetries() bool`

HasKeepaliveRetries returns a boolean if a field has been set.

### SetKeepaliveRetriesNil

`func (o *PatchedBulkWritableVPNProfileRequest) SetKeepaliveRetriesNil(b bool)`

 SetKeepaliveRetriesNil sets the value for KeepaliveRetries to be an explicit nil

### UnsetKeepaliveRetries
`func (o *PatchedBulkWritableVPNProfileRequest) UnsetKeepaliveRetries()`

UnsetKeepaliveRetries ensures that no value is present for KeepaliveRetries, not even an explicit nil
### GetNatTraversal

`func (o *PatchedBulkWritableVPNProfileRequest) GetNatTraversal() bool`

GetNatTraversal returns the NatTraversal field if non-nil, zero value otherwise.

### GetNatTraversalOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetNatTraversalOk() (*bool, bool)`

GetNatTraversalOk returns a tuple with the NatTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTraversal

`func (o *PatchedBulkWritableVPNProfileRequest) SetNatTraversal(v bool)`

SetNatTraversal sets NatTraversal field to given value.

### HasNatTraversal

`func (o *PatchedBulkWritableVPNProfileRequest) HasNatTraversal() bool`

HasNatTraversal returns a boolean if a field has been set.

### GetExtraOptions

`func (o *PatchedBulkWritableVPNProfileRequest) GetExtraOptions() interface{}`

GetExtraOptions returns the ExtraOptions field if non-nil, zero value otherwise.

### GetExtraOptionsOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetExtraOptionsOk() (*interface{}, bool)`

GetExtraOptionsOk returns a tuple with the ExtraOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraOptions

`func (o *PatchedBulkWritableVPNProfileRequest) SetExtraOptions(v interface{})`

SetExtraOptions sets ExtraOptions field to given value.

### HasExtraOptions

`func (o *PatchedBulkWritableVPNProfileRequest) HasExtraOptions() bool`

HasExtraOptions returns a boolean if a field has been set.

### SetExtraOptionsNil

`func (o *PatchedBulkWritableVPNProfileRequest) SetExtraOptionsNil(b bool)`

 SetExtraOptionsNil sets the value for ExtraOptions to be an explicit nil

### UnsetExtraOptions
`func (o *PatchedBulkWritableVPNProfileRequest) UnsetExtraOptions()`

UnsetExtraOptions ensures that no value is present for ExtraOptions, not even an explicit nil
### GetRole

`func (o *PatchedBulkWritableVPNProfileRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkWritableVPNProfileRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkWritableVPNProfileRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedBulkWritableVPNProfileRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedBulkWritableVPNProfileRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSecretsGroup

`func (o *PatchedBulkWritableVPNProfileRequest) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedBulkWritableVPNProfileRequest) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedBulkWritableVPNProfileRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedBulkWritableVPNProfileRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedBulkWritableVPNProfileRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetTenant

`func (o *PatchedBulkWritableVPNProfileRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableVPNProfileRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableVPNProfileRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableVPNProfileRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableVPNProfileRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableVPNProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableVPNProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableVPNProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableVPNProfileRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableVPNProfileRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableVPNProfileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableVPNProfileRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableVPNProfileRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableVPNProfileRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableVPNProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


