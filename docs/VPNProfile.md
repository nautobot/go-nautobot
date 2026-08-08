# VPNProfile

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
**KeepaliveEnabled** | Pointer to **bool** |  | [optional] 
**KeepaliveInterval** | Pointer to **NullableInt32** |  | [optional] 
**KeepaliveRetries** | Pointer to **NullableInt32** |  | [optional] 
**NatTraversal** | Pointer to **bool** |  | [optional] 
**ExtraOptions** | Pointer to **interface{}** | Additional options specific to the VPN technology and/or implementation | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**SecretsGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewVPNProfile

`func NewVPNProfile(objectType string, display string, url string, naturalSlug string, name string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *VPNProfile`

NewVPNProfile instantiates a new VPNProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNProfileWithDefaults

`func NewVPNProfileWithDefaults() *VPNProfile`

NewVPNProfileWithDefaults instantiates a new VPNProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VPNProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VPNProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VPNProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VPNProfile) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VPNProfile) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VPNProfile) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VPNProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VPNProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VPNProfile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VPNProfile) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VPNProfile) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VPNProfile) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetName

`func (o *VPNProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPNProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPNProfile) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VPNProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VPNProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VPNProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VPNProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKeepaliveEnabled

`func (o *VPNProfile) GetKeepaliveEnabled() bool`

GetKeepaliveEnabled returns the KeepaliveEnabled field if non-nil, zero value otherwise.

### GetKeepaliveEnabledOk

`func (o *VPNProfile) GetKeepaliveEnabledOk() (*bool, bool)`

GetKeepaliveEnabledOk returns a tuple with the KeepaliveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveEnabled

`func (o *VPNProfile) SetKeepaliveEnabled(v bool)`

SetKeepaliveEnabled sets KeepaliveEnabled field to given value.

### HasKeepaliveEnabled

`func (o *VPNProfile) HasKeepaliveEnabled() bool`

HasKeepaliveEnabled returns a boolean if a field has been set.

### GetKeepaliveInterval

`func (o *VPNProfile) GetKeepaliveInterval() int32`

GetKeepaliveInterval returns the KeepaliveInterval field if non-nil, zero value otherwise.

### GetKeepaliveIntervalOk

`func (o *VPNProfile) GetKeepaliveIntervalOk() (*int32, bool)`

GetKeepaliveIntervalOk returns a tuple with the KeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveInterval

`func (o *VPNProfile) SetKeepaliveInterval(v int32)`

SetKeepaliveInterval sets KeepaliveInterval field to given value.

### HasKeepaliveInterval

`func (o *VPNProfile) HasKeepaliveInterval() bool`

HasKeepaliveInterval returns a boolean if a field has been set.

### SetKeepaliveIntervalNil

`func (o *VPNProfile) SetKeepaliveIntervalNil(b bool)`

 SetKeepaliveIntervalNil sets the value for KeepaliveInterval to be an explicit nil

### UnsetKeepaliveInterval
`func (o *VPNProfile) UnsetKeepaliveInterval()`

UnsetKeepaliveInterval ensures that no value is present for KeepaliveInterval, not even an explicit nil
### GetKeepaliveRetries

`func (o *VPNProfile) GetKeepaliveRetries() int32`

GetKeepaliveRetries returns the KeepaliveRetries field if non-nil, zero value otherwise.

### GetKeepaliveRetriesOk

`func (o *VPNProfile) GetKeepaliveRetriesOk() (*int32, bool)`

GetKeepaliveRetriesOk returns a tuple with the KeepaliveRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveRetries

`func (o *VPNProfile) SetKeepaliveRetries(v int32)`

SetKeepaliveRetries sets KeepaliveRetries field to given value.

### HasKeepaliveRetries

`func (o *VPNProfile) HasKeepaliveRetries() bool`

HasKeepaliveRetries returns a boolean if a field has been set.

### SetKeepaliveRetriesNil

`func (o *VPNProfile) SetKeepaliveRetriesNil(b bool)`

 SetKeepaliveRetriesNil sets the value for KeepaliveRetries to be an explicit nil

### UnsetKeepaliveRetries
`func (o *VPNProfile) UnsetKeepaliveRetries()`

UnsetKeepaliveRetries ensures that no value is present for KeepaliveRetries, not even an explicit nil
### GetNatTraversal

`func (o *VPNProfile) GetNatTraversal() bool`

GetNatTraversal returns the NatTraversal field if non-nil, zero value otherwise.

### GetNatTraversalOk

`func (o *VPNProfile) GetNatTraversalOk() (*bool, bool)`

GetNatTraversalOk returns a tuple with the NatTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatTraversal

`func (o *VPNProfile) SetNatTraversal(v bool)`

SetNatTraversal sets NatTraversal field to given value.

### HasNatTraversal

`func (o *VPNProfile) HasNatTraversal() bool`

HasNatTraversal returns a boolean if a field has been set.

### GetExtraOptions

`func (o *VPNProfile) GetExtraOptions() interface{}`

GetExtraOptions returns the ExtraOptions field if non-nil, zero value otherwise.

### GetExtraOptionsOk

`func (o *VPNProfile) GetExtraOptionsOk() (*interface{}, bool)`

GetExtraOptionsOk returns a tuple with the ExtraOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraOptions

`func (o *VPNProfile) SetExtraOptions(v interface{})`

SetExtraOptions sets ExtraOptions field to given value.

### HasExtraOptions

`func (o *VPNProfile) HasExtraOptions() bool`

HasExtraOptions returns a boolean if a field has been set.

### SetExtraOptionsNil

`func (o *VPNProfile) SetExtraOptionsNil(b bool)`

 SetExtraOptionsNil sets the value for ExtraOptions to be an explicit nil

### UnsetExtraOptions
`func (o *VPNProfile) UnsetExtraOptions()`

UnsetExtraOptions ensures that no value is present for ExtraOptions, not even an explicit nil
### GetRole

`func (o *VPNProfile) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VPNProfile) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VPNProfile) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *VPNProfile) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VPNProfile) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VPNProfile) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSecretsGroup

`func (o *VPNProfile) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *VPNProfile) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *VPNProfile) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *VPNProfile) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *VPNProfile) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *VPNProfile) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetTenant

`func (o *VPNProfile) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VPNProfile) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VPNProfile) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VPNProfile) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VPNProfile) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VPNProfile) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCreated

`func (o *VPNProfile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VPNProfile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VPNProfile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VPNProfile) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VPNProfile) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VPNProfile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VPNProfile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VPNProfile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VPNProfile) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VPNProfile) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *VPNProfile) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VPNProfile) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VPNProfile) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *VPNProfile) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNProfile) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNProfile) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNProfile) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *VPNProfile) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VPNProfile) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VPNProfile) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VPNProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


