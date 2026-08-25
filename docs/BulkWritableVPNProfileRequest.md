# BulkWritableVPNProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewBulkWritableVPNProfileRequest

`func NewBulkWritableVPNProfileRequest(id string, name string, ) *BulkWritableVPNProfileRequest`

NewBulkWritableVPNProfileRequest instantiates a new BulkWritableVPNProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVPNProfileRequestWithDefaults

`func NewBulkWritableVPNProfileRequestWithDefaults() *BulkWritableVPNProfileRequest`

NewBulkWritableVPNProfileRequestWithDefaults instantiates a new BulkWritableVPNProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVPNProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVPNProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVPNProfileRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableVPNProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableVPNProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableVPNProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableVPNProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableVPNProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableVPNProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableVPNProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKeepaliveEnabled

`func (o *BulkWritableVPNProfileRequest) GetKeepaliveEnabled() bool`

GetKeepaliveEnabled returns the KeepaliveEnabled field if non-nil, zero value otherwise.

### GetKeepaliveEnabledOk

`func (o *BulkWritableVPNProfileRequest) GetKeepaliveEnabledOk() (*bool, bool)`

GetKeepaliveEnabledOk returns a tuple with the KeepaliveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveEnabled

`func (o *BulkWritableVPNProfileRequest) SetKeepaliveEnabled(v bool)`

SetKeepaliveEnabled sets KeepaliveEnabled field to given value.

### HasKeepaliveEnabled

`func (o *BulkWritableVPNProfileRequest) HasKeepaliveEnabled() bool`

HasKeepaliveEnabled returns a boolean if a field has been set.

### GetKeepaliveInterval

`func (o *BulkWritableVPNProfileRequest) GetKeepaliveInterval() int32`

GetKeepaliveInterval returns the KeepaliveInterval field if non-nil, zero value otherwise.

### GetKeepaliveIntervalOk

`func (o *BulkWritableVPNProfileRequest) GetKeepaliveIntervalOk() (*int32, bool)`

GetKeepaliveIntervalOk returns a tuple with the KeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveInterval

`func (o *BulkWritableVPNProfileRequest) SetKeepaliveInterval(v int32)`

SetKeepaliveInterval sets KeepaliveInterval field to given value.

### HasKeepaliveInterval

`func (o *BulkWritableVPNProfileRequest) HasKeepaliveInterval() bool`

HasKeepaliveInterval returns a boolean if a field has been set.

### SetKeepaliveIntervalNil

`func (o *BulkWritableVPNProfileRequest) SetKeepaliveIntervalNil(b bool)`

 SetKeepaliveIntervalNil sets the value for KeepaliveInterval to be an explicit nil

### UnsetKeepaliveInterval
`func (o *BulkWritableVPNProfileRequest) UnsetKeepaliveInterval()`

UnsetKeepaliveInterval ensures that no value is present for KeepaliveInterval, not even an explicit nil
### GetKeepaliveRetries

`func (o *BulkWritableVPNProfileRequest) GetKeepaliveRetries() int32`

GetKeepaliveRetries returns the KeepaliveRetries field if non-nil, zero value otherwise.

### GetKeepaliveRetriesOk

`func (o *BulkWritableVPNProfileRequest) GetKeepaliveRetriesOk() (*int32, bool)`

GetKeepaliveRetriesOk returns a tuple with the KeepaliveRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveRetries

`func (o *BulkWritableVPNProfileRequest) SetKeepaliveRetries(v int32)`

SetKeepaliveRetries sets KeepaliveRetries field to given value.

### HasKeepaliveRetries

`func (o *BulkWritableVPNProfileRequest) HasKeepaliveRetries() bool`

HasKeepaliveRetries returns a boolean if a field has been set.

### SetKeepaliveRetriesNil

`func (o *BulkWritableVPNProfileRequest) SetKeepaliveRetriesNil(b bool)`

 SetKeepaliveRetriesNil sets the value for KeepaliveRetries to be an explicit nil

### UnsetKeepaliveRetries
`func (o *BulkWritableVPNProfileRequest) UnsetKeepaliveRetries()`

UnsetKeepaliveRetries ensures that no value is present for KeepaliveRetries, not even an explicit nil
### GetNatTraversal

`func (o *BulkWritableVPNProfileRequest) GetNatTraversal() bool`

GetNatTraversal returns the NatTraversal field if non-nil, zero value otherwise.

### GetNatTraversalOk

`func (o *BulkWritableVPNProfileRequest) GetNatTraversalOk() (*bool, bool)`

GetNatTraversalOk returns a tuple with the NatTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTraversal

`func (o *BulkWritableVPNProfileRequest) SetNatTraversal(v bool)`

SetNatTraversal sets NatTraversal field to given value.

### HasNatTraversal

`func (o *BulkWritableVPNProfileRequest) HasNatTraversal() bool`

HasNatTraversal returns a boolean if a field has been set.

### GetExtraOptions

`func (o *BulkWritableVPNProfileRequest) GetExtraOptions() interface{}`

GetExtraOptions returns the ExtraOptions field if non-nil, zero value otherwise.

### GetExtraOptionsOk

`func (o *BulkWritableVPNProfileRequest) GetExtraOptionsOk() (*interface{}, bool)`

GetExtraOptionsOk returns a tuple with the ExtraOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraOptions

`func (o *BulkWritableVPNProfileRequest) SetExtraOptions(v interface{})`

SetExtraOptions sets ExtraOptions field to given value.

### HasExtraOptions

`func (o *BulkWritableVPNProfileRequest) HasExtraOptions() bool`

HasExtraOptions returns a boolean if a field has been set.

### SetExtraOptionsNil

`func (o *BulkWritableVPNProfileRequest) SetExtraOptionsNil(b bool)`

 SetExtraOptionsNil sets the value for ExtraOptions to be an explicit nil

### UnsetExtraOptions
`func (o *BulkWritableVPNProfileRequest) UnsetExtraOptions()`

UnsetExtraOptions ensures that no value is present for ExtraOptions, not even an explicit nil
### GetRole

`func (o *BulkWritableVPNProfileRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkWritableVPNProfileRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkWritableVPNProfileRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkWritableVPNProfileRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkWritableVPNProfileRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkWritableVPNProfileRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSecretsGroup

`func (o *BulkWritableVPNProfileRequest) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *BulkWritableVPNProfileRequest) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *BulkWritableVPNProfileRequest) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *BulkWritableVPNProfileRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *BulkWritableVPNProfileRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *BulkWritableVPNProfileRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetTenant

`func (o *BulkWritableVPNProfileRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableVPNProfileRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableVPNProfileRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableVPNProfileRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableVPNProfileRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableVPNProfileRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableVPNProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableVPNProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableVPNProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableVPNProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableVPNProfileRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableVPNProfileRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableVPNProfileRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableVPNProfileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableVPNProfileRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableVPNProfileRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableVPNProfileRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableVPNProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


