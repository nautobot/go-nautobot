# ControllerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### NewControllerRequest

`func NewControllerRequest(name string, status ApprovalWorkflowStageResponseApprovalWorkflowStage, location ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *ControllerRequest`

NewControllerRequest instantiates a new ControllerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerRequestWithDefaults

`func NewControllerRequestWithDefaults() *ControllerRequest`

NewControllerRequestWithDefaults instantiates a new ControllerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControllerRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCapabilities

`func (o *ControllerRequest) GetCapabilities() []CapabilitiesEnum`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ControllerRequest) GetCapabilitiesOk() (*[]CapabilitiesEnum, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ControllerRequest) SetCapabilities(v []CapabilitiesEnum)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ControllerRequest) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *ControllerRequest) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *ControllerRequest) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetName

`func (o *ControllerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ControllerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ControllerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ControllerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ControllerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ControllerRequest) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ControllerRequest) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ControllerRequest) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.


### GetLocation

`func (o *ControllerRequest) GetLocation() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ControllerRequest) GetLocationOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ControllerRequest) SetLocation(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLocation sets Location field to given value.


### GetPlatform

`func (o *ControllerRequest) GetPlatform() ApprovalWorkflowUser`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ControllerRequest) GetPlatformOk() (*ApprovalWorkflowUser, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ControllerRequest) SetPlatform(v ApprovalWorkflowUser)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ControllerRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *ControllerRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *ControllerRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetRole

`func (o *ControllerRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ControllerRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ControllerRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *ControllerRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ControllerRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ControllerRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *ControllerRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ControllerRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ControllerRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ControllerRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *ControllerRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *ControllerRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetExternalIntegration

`func (o *ControllerRequest) GetExternalIntegration() ApprovalWorkflowUser`

GetExternalIntegration returns the ExternalIntegration field if non-nil, zero value otherwise.

### GetExternalIntegrationOk

`func (o *ControllerRequest) GetExternalIntegrationOk() (*ApprovalWorkflowUser, bool)`

GetExternalIntegrationOk returns a tuple with the ExternalIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegration

`func (o *ControllerRequest) SetExternalIntegration(v ApprovalWorkflowUser)`

SetExternalIntegration sets ExternalIntegration field to given value.

### HasExternalIntegration

`func (o *ControllerRequest) HasExternalIntegration() bool`

HasExternalIntegration returns a boolean if a field has been set.

### SetExternalIntegrationNil

`func (o *ControllerRequest) SetExternalIntegrationNil(b bool)`

 SetExternalIntegrationNil sets the value for ExternalIntegration to be an explicit nil

### UnsetExternalIntegration
`func (o *ControllerRequest) UnsetExternalIntegration()`

UnsetExternalIntegration ensures that no value is present for ExternalIntegration, not even an explicit nil
### GetControllerDevice

`func (o *ControllerRequest) GetControllerDevice() ApprovalWorkflowUser`

GetControllerDevice returns the ControllerDevice field if non-nil, zero value otherwise.

### GetControllerDeviceOk

`func (o *ControllerRequest) GetControllerDeviceOk() (*ApprovalWorkflowUser, bool)`

GetControllerDeviceOk returns a tuple with the ControllerDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDevice

`func (o *ControllerRequest) SetControllerDevice(v ApprovalWorkflowUser)`

SetControllerDevice sets ControllerDevice field to given value.

### HasControllerDevice

`func (o *ControllerRequest) HasControllerDevice() bool`

HasControllerDevice returns a boolean if a field has been set.

### SetControllerDeviceNil

`func (o *ControllerRequest) SetControllerDeviceNil(b bool)`

 SetControllerDeviceNil sets the value for ControllerDevice to be an explicit nil

### UnsetControllerDevice
`func (o *ControllerRequest) UnsetControllerDevice()`

UnsetControllerDevice ensures that no value is present for ControllerDevice, not even an explicit nil
### GetControllerDeviceRedundancyGroup

`func (o *ControllerRequest) GetControllerDeviceRedundancyGroup() ApprovalWorkflowUser`

GetControllerDeviceRedundancyGroup returns the ControllerDeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetControllerDeviceRedundancyGroupOk

`func (o *ControllerRequest) GetControllerDeviceRedundancyGroupOk() (*ApprovalWorkflowUser, bool)`

GetControllerDeviceRedundancyGroupOk returns a tuple with the ControllerDeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDeviceRedundancyGroup

`func (o *ControllerRequest) SetControllerDeviceRedundancyGroup(v ApprovalWorkflowUser)`

SetControllerDeviceRedundancyGroup sets ControllerDeviceRedundancyGroup field to given value.

### HasControllerDeviceRedundancyGroup

`func (o *ControllerRequest) HasControllerDeviceRedundancyGroup() bool`

HasControllerDeviceRedundancyGroup returns a boolean if a field has been set.

### SetControllerDeviceRedundancyGroupNil

`func (o *ControllerRequest) SetControllerDeviceRedundancyGroupNil(b bool)`

 SetControllerDeviceRedundancyGroupNil sets the value for ControllerDeviceRedundancyGroup to be an explicit nil

### UnsetControllerDeviceRedundancyGroup
`func (o *ControllerRequest) UnsetControllerDeviceRedundancyGroup()`

UnsetControllerDeviceRedundancyGroup ensures that no value is present for ControllerDeviceRedundancyGroup, not even an explicit nil
### GetCustomFields

`func (o *ControllerRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ControllerRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ControllerRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ControllerRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ControllerRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ControllerRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ControllerRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ControllerRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *ControllerRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ControllerRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ControllerRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ControllerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


