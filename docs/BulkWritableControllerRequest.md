# BulkWritableControllerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Capabilities** | Pointer to [**[]CapabilitiesEnum**](CapabilitiesEnum.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Status** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Location** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Platform** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ExternalIntegration** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ControllerDevice** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ControllerDeviceRedundancyGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewBulkWritableControllerRequest

`func NewBulkWritableControllerRequest(id string, name string, status ApprovalWorkflowStageResponseApprovalWorkflowStage, location ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *BulkWritableControllerRequest`

NewBulkWritableControllerRequest instantiates a new BulkWritableControllerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableControllerRequestWithDefaults

`func NewBulkWritableControllerRequestWithDefaults() *BulkWritableControllerRequest`

NewBulkWritableControllerRequestWithDefaults instantiates a new BulkWritableControllerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableControllerRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableControllerRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableControllerRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCapabilities

`func (o *BulkWritableControllerRequest) GetCapabilities() []CapabilitiesEnum`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *BulkWritableControllerRequest) GetCapabilitiesOk() (*[]CapabilitiesEnum, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *BulkWritableControllerRequest) SetCapabilities(v []CapabilitiesEnum)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *BulkWritableControllerRequest) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *BulkWritableControllerRequest) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *BulkWritableControllerRequest) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetName

`func (o *BulkWritableControllerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableControllerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableControllerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableControllerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableControllerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableControllerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableControllerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *BulkWritableControllerRequest) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableControllerRequest) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableControllerRequest) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.


### GetLocation

`func (o *BulkWritableControllerRequest) GetLocation() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BulkWritableControllerRequest) GetLocationOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BulkWritableControllerRequest) SetLocation(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLocation sets Location field to given value.


### GetPlatform

`func (o *BulkWritableControllerRequest) GetPlatform() ApprovalWorkflowUser`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *BulkWritableControllerRequest) GetPlatformOk() (*ApprovalWorkflowUser, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *BulkWritableControllerRequest) SetPlatform(v ApprovalWorkflowUser)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *BulkWritableControllerRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *BulkWritableControllerRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *BulkWritableControllerRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetRole

`func (o *BulkWritableControllerRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkWritableControllerRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkWritableControllerRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkWritableControllerRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkWritableControllerRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkWritableControllerRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *BulkWritableControllerRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableControllerRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableControllerRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableControllerRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableControllerRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableControllerRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetExternalIntegration

`func (o *BulkWritableControllerRequest) GetExternalIntegration() ApprovalWorkflowUser`

GetExternalIntegration returns the ExternalIntegration field if non-nil, zero value otherwise.

### GetExternalIntegrationOk

`func (o *BulkWritableControllerRequest) GetExternalIntegrationOk() (*ApprovalWorkflowUser, bool)`

GetExternalIntegrationOk returns a tuple with the ExternalIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegration

`func (o *BulkWritableControllerRequest) SetExternalIntegration(v ApprovalWorkflowUser)`

SetExternalIntegration sets ExternalIntegration field to given value.

### HasExternalIntegration

`func (o *BulkWritableControllerRequest) HasExternalIntegration() bool`

HasExternalIntegration returns a boolean if a field has been set.

### SetExternalIntegrationNil

`func (o *BulkWritableControllerRequest) SetExternalIntegrationNil(b bool)`

 SetExternalIntegrationNil sets the value for ExternalIntegration to be an explicit nil

### UnsetExternalIntegration
`func (o *BulkWritableControllerRequest) UnsetExternalIntegration()`

UnsetExternalIntegration ensures that no value is present for ExternalIntegration, not even an explicit nil
### GetControllerDevice

`func (o *BulkWritableControllerRequest) GetControllerDevice() ApprovalWorkflowUser`

GetControllerDevice returns the ControllerDevice field if non-nil, zero value otherwise.

### GetControllerDeviceOk

`func (o *BulkWritableControllerRequest) GetControllerDeviceOk() (*ApprovalWorkflowUser, bool)`

GetControllerDeviceOk returns a tuple with the ControllerDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDevice

`func (o *BulkWritableControllerRequest) SetControllerDevice(v ApprovalWorkflowUser)`

SetControllerDevice sets ControllerDevice field to given value.

### HasControllerDevice

`func (o *BulkWritableControllerRequest) HasControllerDevice() bool`

HasControllerDevice returns a boolean if a field has been set.

### SetControllerDeviceNil

`func (o *BulkWritableControllerRequest) SetControllerDeviceNil(b bool)`

 SetControllerDeviceNil sets the value for ControllerDevice to be an explicit nil

### UnsetControllerDevice
`func (o *BulkWritableControllerRequest) UnsetControllerDevice()`

UnsetControllerDevice ensures that no value is present for ControllerDevice, not even an explicit nil
### GetControllerDeviceRedundancyGroup

`func (o *BulkWritableControllerRequest) GetControllerDeviceRedundancyGroup() ApprovalWorkflowUser`

GetControllerDeviceRedundancyGroup returns the ControllerDeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetControllerDeviceRedundancyGroupOk

`func (o *BulkWritableControllerRequest) GetControllerDeviceRedundancyGroupOk() (*ApprovalWorkflowUser, bool)`

GetControllerDeviceRedundancyGroupOk returns a tuple with the ControllerDeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDeviceRedundancyGroup

`func (o *BulkWritableControllerRequest) SetControllerDeviceRedundancyGroup(v ApprovalWorkflowUser)`

SetControllerDeviceRedundancyGroup sets ControllerDeviceRedundancyGroup field to given value.

### HasControllerDeviceRedundancyGroup

`func (o *BulkWritableControllerRequest) HasControllerDeviceRedundancyGroup() bool`

HasControllerDeviceRedundancyGroup returns a boolean if a field has been set.

### SetControllerDeviceRedundancyGroupNil

`func (o *BulkWritableControllerRequest) SetControllerDeviceRedundancyGroupNil(b bool)`

 SetControllerDeviceRedundancyGroupNil sets the value for ControllerDeviceRedundancyGroup to be an explicit nil

### UnsetControllerDeviceRedundancyGroup
`func (o *BulkWritableControllerRequest) UnsetControllerDeviceRedundancyGroup()`

UnsetControllerDeviceRedundancyGroup ensures that no value is present for ControllerDeviceRedundancyGroup, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableControllerRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableControllerRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableControllerRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableControllerRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableControllerRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableControllerRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableControllerRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableControllerRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableControllerRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableControllerRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableControllerRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableControllerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


