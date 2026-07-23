# VPNProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
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

### NewVPNProfileRequest

`func NewVPNProfileRequest(name string, ) *VPNProfileRequest`

NewVPNProfileRequest instantiates a new VPNProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNProfileRequestWithDefaults

`func NewVPNProfileRequestWithDefaults() *VPNProfileRequest`

NewVPNProfileRequestWithDefaults instantiates a new VPNProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNProfileRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNProfileRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VPNProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPNProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPNProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VPNProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VPNProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VPNProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VPNProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKeepaliveEnabled

`func (o *VPNProfileRequest) GetKeepaliveEnabled() bool`

GetKeepaliveEnabled returns the KeepaliveEnabled field if non-nil, zero value otherwise.

### GetKeepaliveEnabledOk

`func (o *VPNProfileRequest) GetKeepaliveEnabledOk() (*bool, bool)`

GetKeepaliveEnabledOk returns a tuple with the KeepaliveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveEnabled

`func (o *VPNProfileRequest) SetKeepaliveEnabled(v bool)`

SetKeepaliveEnabled sets KeepaliveEnabled field to given value.

### HasKeepaliveEnabled

`func (o *VPNProfileRequest) HasKeepaliveEnabled() bool`

HasKeepaliveEnabled returns a boolean if a field has been set.

### GetKeepaliveInterval

`func (o *VPNProfileRequest) GetKeepaliveInterval() int32`

GetKeepaliveInterval returns the KeepaliveInterval field if non-nil, zero value otherwise.

### GetKeepaliveIntervalOk

`func (o *VPNProfileRequest) GetKeepaliveIntervalOk() (*int32, bool)`

GetKeepaliveIntervalOk returns a tuple with the KeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveInterval

`func (o *VPNProfileRequest) SetKeepaliveInterval(v int32)`

SetKeepaliveInterval sets KeepaliveInterval field to given value.

### HasKeepaliveInterval

`func (o *VPNProfileRequest) HasKeepaliveInterval() bool`

HasKeepaliveInterval returns a boolean if a field has been set.

### SetKeepaliveIntervalNil

`func (o *VPNProfileRequest) SetKeepaliveIntervalNil(b bool)`

 SetKeepaliveIntervalNil sets the value for KeepaliveInterval to be an explicit nil

### UnsetKeepaliveInterval
`func (o *VPNProfileRequest) UnsetKeepaliveInterval()`

UnsetKeepaliveInterval ensures that no value is present for KeepaliveInterval, not even an explicit nil
### GetKeepaliveRetries

`func (o *VPNProfileRequest) GetKeepaliveRetries() int32`

GetKeepaliveRetries returns the KeepaliveRetries field if non-nil, zero value otherwise.

### GetKeepaliveRetriesOk

`func (o *VPNProfileRequest) GetKeepaliveRetriesOk() (*int32, bool)`

GetKeepaliveRetriesOk returns a tuple with the KeepaliveRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveRetries

`func (o *VPNProfileRequest) SetKeepaliveRetries(v int32)`

SetKeepaliveRetries sets KeepaliveRetries field to given value.

### HasKeepaliveRetries

`func (o *VPNProfileRequest) HasKeepaliveRetries() bool`

HasKeepaliveRetries returns a boolean if a field has been set.

### SetKeepaliveRetriesNil

`func (o *VPNProfileRequest) SetKeepaliveRetriesNil(b bool)`

 SetKeepaliveRetriesNil sets the value for KeepaliveRetries to be an explicit nil

### UnsetKeepaliveRetries
`func (o *VPNProfileRequest) UnsetKeepaliveRetries()`

UnsetKeepaliveRetries ensures that no value is present for KeepaliveRetries, not even an explicit nil
### GetNatTraversal

`func (o *VPNProfileRequest) GetNatTraversal() bool`

GetNatTraversal returns the NatTraversal field if non-nil, zero value otherwise.

### GetNatTraversalOk

`func (o *VPNProfileRequest) GetNatTraversalOk() (*bool, bool)`

GetNatTraversalOk returns a tuple with the NatTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTraversal

`func (o *VPNProfileRequest) SetNatTraversal(v bool)`

SetNatTraversal sets NatTraversal field to given value.

### HasNatTraversal

`func (o *VPNProfileRequest) HasNatTraversal() bool`

HasNatTraversal returns a boolean if a field has been set.

### GetExtraOptions

`func (o *VPNProfileRequest) GetExtraOptions() interface{}`

GetExtraOptions returns the ExtraOptions field if non-nil, zero value otherwise.

### GetExtraOptionsOk

`func (o *VPNProfileRequest) GetExtraOptionsOk() (*interface{}, bool)`

GetExtraOptionsOk returns a tuple with the ExtraOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraOptions

`func (o *VPNProfileRequest) SetExtraOptions(v interface{})`

SetExtraOptions sets ExtraOptions field to given value.

### HasExtraOptions

`func (o *VPNProfileRequest) HasExtraOptions() bool`

HasExtraOptions returns a boolean if a field has been set.

### SetExtraOptionsNil

`func (o *VPNProfileRequest) SetExtraOptionsNil(b bool)`

 SetExtraOptionsNil sets the value for ExtraOptions to be an explicit nil

### UnsetExtraOptions
`func (o *VPNProfileRequest) UnsetExtraOptions()`

UnsetExtraOptions ensures that no value is present for ExtraOptions, not even an explicit nil
### GetRole

`func (o *VPNProfileRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VPNProfileRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VPNProfileRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *VPNProfileRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VPNProfileRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VPNProfileRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSecretsGroup

`func (o *VPNProfileRequest) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *VPNProfileRequest) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *VPNProfileRequest) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *VPNProfileRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *VPNProfileRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *VPNProfileRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetTenant

`func (o *VPNProfileRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VPNProfileRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VPNProfileRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VPNProfileRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VPNProfileRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VPNProfileRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *VPNProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *VPNProfileRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *VPNProfileRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *VPNProfileRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *VPNProfileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *VPNProfileRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VPNProfileRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VPNProfileRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VPNProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


