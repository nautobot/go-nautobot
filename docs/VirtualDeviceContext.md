# VirtualDeviceContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Identifier** | Pointer to **NullableInt32** | Unique identifier provided by the platform being virtualized (Example: Nexus VDC Identifier) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Device** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Status** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NullablePrimaryIPv4**](PrimaryIPv4.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullablePrimaryIPv6**](PrimaryIPv6.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewVirtualDeviceContext

`func NewVirtualDeviceContext(objectType string, display string, url string, naturalSlug string, name string, device ApprovalWorkflowStageResponseApprovalWorkflowStage, status ApprovalWorkflowStageResponseApprovalWorkflowStage, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *VirtualDeviceContext`

NewVirtualDeviceContext instantiates a new VirtualDeviceContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDeviceContextWithDefaults

`func NewVirtualDeviceContextWithDefaults() *VirtualDeviceContext`

NewVirtualDeviceContextWithDefaults instantiates a new VirtualDeviceContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualDeviceContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualDeviceContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualDeviceContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualDeviceContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VirtualDeviceContext) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualDeviceContext) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualDeviceContext) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VirtualDeviceContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VirtualDeviceContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VirtualDeviceContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VirtualDeviceContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualDeviceContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualDeviceContext) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VirtualDeviceContext) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VirtualDeviceContext) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VirtualDeviceContext) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetName

`func (o *VirtualDeviceContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualDeviceContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualDeviceContext) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *VirtualDeviceContext) GetIdentifier() int32`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VirtualDeviceContext) GetIdentifierOk() (*int32, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VirtualDeviceContext) SetIdentifier(v int32)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VirtualDeviceContext) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *VirtualDeviceContext) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *VirtualDeviceContext) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetDescription

`func (o *VirtualDeviceContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualDeviceContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualDeviceContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualDeviceContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevice

`func (o *VirtualDeviceContext) GetDevice() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VirtualDeviceContext) GetDeviceOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VirtualDeviceContext) SetDevice(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDevice sets Device field to given value.


### GetStatus

`func (o *VirtualDeviceContext) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualDeviceContext) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualDeviceContext) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.


### GetRole

`func (o *VirtualDeviceContext) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VirtualDeviceContext) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VirtualDeviceContext) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *VirtualDeviceContext) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VirtualDeviceContext) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VirtualDeviceContext) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPrimaryIp4

`func (o *VirtualDeviceContext) GetPrimaryIp4() PrimaryIPv4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *VirtualDeviceContext) GetPrimaryIp4Ok() (*PrimaryIPv4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *VirtualDeviceContext) SetPrimaryIp4(v PrimaryIPv4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *VirtualDeviceContext) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *VirtualDeviceContext) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *VirtualDeviceContext) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *VirtualDeviceContext) GetPrimaryIp6() PrimaryIPv6`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *VirtualDeviceContext) GetPrimaryIp6Ok() (*PrimaryIPv6, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *VirtualDeviceContext) SetPrimaryIp6(v PrimaryIPv6)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *VirtualDeviceContext) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *VirtualDeviceContext) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *VirtualDeviceContext) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetTenant

`func (o *VirtualDeviceContext) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VirtualDeviceContext) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VirtualDeviceContext) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VirtualDeviceContext) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VirtualDeviceContext) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VirtualDeviceContext) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCreated

`func (o *VirtualDeviceContext) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VirtualDeviceContext) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VirtualDeviceContext) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VirtualDeviceContext) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VirtualDeviceContext) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VirtualDeviceContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VirtualDeviceContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VirtualDeviceContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VirtualDeviceContext) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VirtualDeviceContext) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *VirtualDeviceContext) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VirtualDeviceContext) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VirtualDeviceContext) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *VirtualDeviceContext) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualDeviceContext) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualDeviceContext) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualDeviceContext) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *VirtualDeviceContext) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualDeviceContext) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualDeviceContext) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualDeviceContext) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


