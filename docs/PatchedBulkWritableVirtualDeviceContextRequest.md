# PatchedBulkWritableVirtualDeviceContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **NullableInt32** | Unique identifier provided by the platform being virtualized (Example: Nexus VDC Identifier) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Device** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**Status** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NullablePrimaryIPv4**](PrimaryIPv4.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullablePrimaryIPv6**](PrimaryIPv6.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVirtualDeviceContextRequest

`func NewPatchedBulkWritableVirtualDeviceContextRequest(id string, ) *PatchedBulkWritableVirtualDeviceContextRequest`

NewPatchedBulkWritableVirtualDeviceContextRequest instantiates a new PatchedBulkWritableVirtualDeviceContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVirtualDeviceContextRequestWithDefaults

`func NewPatchedBulkWritableVirtualDeviceContextRequestWithDefaults() *PatchedBulkWritableVirtualDeviceContextRequest`

NewPatchedBulkWritableVirtualDeviceContextRequestWithDefaults instantiates a new PatchedBulkWritableVirtualDeviceContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdentifier

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetIdentifier() int32`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetIdentifierOk() (*int32, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetIdentifier(v int32)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *PatchedBulkWritableVirtualDeviceContextRequest) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetDescription

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevice

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetDevice() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetDeviceOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetDevice(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedBulkWritableVirtualDeviceContextRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPrimaryIp4

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetPrimaryIp4() PrimaryIPv4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetPrimaryIp4Ok() (*PrimaryIPv4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetPrimaryIp4(v PrimaryIPv4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *PatchedBulkWritableVirtualDeviceContextRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetPrimaryIp6() PrimaryIPv6`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetPrimaryIp6Ok() (*PrimaryIPv6, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetPrimaryIp6(v PrimaryIPv6)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *PatchedBulkWritableVirtualDeviceContextRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetTenant

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableVirtualDeviceContextRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableVirtualDeviceContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


