# PatchedControllerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Location** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Platform** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Role** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**ExternalIntegration** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**ControllerDevice** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**ControllerDeviceRedundancyGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedControllerRequest

`func NewPatchedControllerRequest() *PatchedControllerRequest`

NewPatchedControllerRequest instantiates a new PatchedControllerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedControllerRequestWithDefaults

`func NewPatchedControllerRequestWithDefaults() *PatchedControllerRequest`

NewPatchedControllerRequestWithDefaults instantiates a new PatchedControllerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedControllerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedControllerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedControllerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedControllerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedControllerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedControllerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedControllerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedControllerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedControllerRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedControllerRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedControllerRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedControllerRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedControllerRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedControllerRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedControllerRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedControllerRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPlatform

`func (o *PatchedControllerRequest) GetPlatform() BulkWritableCircuitRequestTenant`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedControllerRequest) GetPlatformOk() (*BulkWritableCircuitRequestTenant, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedControllerRequest) SetPlatform(v BulkWritableCircuitRequestTenant)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedControllerRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PatchedControllerRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PatchedControllerRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetRole

`func (o *PatchedControllerRequest) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedControllerRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedControllerRequest) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedControllerRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedControllerRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedControllerRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *PatchedControllerRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedControllerRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedControllerRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedControllerRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedControllerRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedControllerRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetExternalIntegration

`func (o *PatchedControllerRequest) GetExternalIntegration() BulkWritableCircuitRequestTenant`

GetExternalIntegration returns the ExternalIntegration field if non-nil, zero value otherwise.

### GetExternalIntegrationOk

`func (o *PatchedControllerRequest) GetExternalIntegrationOk() (*BulkWritableCircuitRequestTenant, bool)`

GetExternalIntegrationOk returns a tuple with the ExternalIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegration

`func (o *PatchedControllerRequest) SetExternalIntegration(v BulkWritableCircuitRequestTenant)`

SetExternalIntegration sets ExternalIntegration field to given value.

### HasExternalIntegration

`func (o *PatchedControllerRequest) HasExternalIntegration() bool`

HasExternalIntegration returns a boolean if a field has been set.

### SetExternalIntegrationNil

`func (o *PatchedControllerRequest) SetExternalIntegrationNil(b bool)`

 SetExternalIntegrationNil sets the value for ExternalIntegration to be an explicit nil

### UnsetExternalIntegration
`func (o *PatchedControllerRequest) UnsetExternalIntegration()`

UnsetExternalIntegration ensures that no value is present for ExternalIntegration, not even an explicit nil
### GetControllerDevice

`func (o *PatchedControllerRequest) GetControllerDevice() BulkWritableCircuitRequestTenant`

GetControllerDevice returns the ControllerDevice field if non-nil, zero value otherwise.

### GetControllerDeviceOk

`func (o *PatchedControllerRequest) GetControllerDeviceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetControllerDeviceOk returns a tuple with the ControllerDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDevice

`func (o *PatchedControllerRequest) SetControllerDevice(v BulkWritableCircuitRequestTenant)`

SetControllerDevice sets ControllerDevice field to given value.

### HasControllerDevice

`func (o *PatchedControllerRequest) HasControllerDevice() bool`

HasControllerDevice returns a boolean if a field has been set.

### SetControllerDeviceNil

`func (o *PatchedControllerRequest) SetControllerDeviceNil(b bool)`

 SetControllerDeviceNil sets the value for ControllerDevice to be an explicit nil

### UnsetControllerDevice
`func (o *PatchedControllerRequest) UnsetControllerDevice()`

UnsetControllerDevice ensures that no value is present for ControllerDevice, not even an explicit nil
### GetControllerDeviceRedundancyGroup

`func (o *PatchedControllerRequest) GetControllerDeviceRedundancyGroup() BulkWritableCircuitRequestTenant`

GetControllerDeviceRedundancyGroup returns the ControllerDeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetControllerDeviceRedundancyGroupOk

`func (o *PatchedControllerRequest) GetControllerDeviceRedundancyGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetControllerDeviceRedundancyGroupOk returns a tuple with the ControllerDeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDeviceRedundancyGroup

`func (o *PatchedControllerRequest) SetControllerDeviceRedundancyGroup(v BulkWritableCircuitRequestTenant)`

SetControllerDeviceRedundancyGroup sets ControllerDeviceRedundancyGroup field to given value.

### HasControllerDeviceRedundancyGroup

`func (o *PatchedControllerRequest) HasControllerDeviceRedundancyGroup() bool`

HasControllerDeviceRedundancyGroup returns a boolean if a field has been set.

### SetControllerDeviceRedundancyGroupNil

`func (o *PatchedControllerRequest) SetControllerDeviceRedundancyGroupNil(b bool)`

 SetControllerDeviceRedundancyGroupNil sets the value for ControllerDeviceRedundancyGroup to be an explicit nil

### UnsetControllerDeviceRedundancyGroup
`func (o *PatchedControllerRequest) UnsetControllerDeviceRedundancyGroup()`

UnsetControllerDeviceRedundancyGroup ensures that no value is present for ControllerDeviceRedundancyGroup, not even an explicit nil
### GetTags

`func (o *PatchedControllerRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedControllerRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedControllerRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedControllerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedControllerRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedControllerRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedControllerRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedControllerRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedControllerRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedControllerRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedControllerRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedControllerRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


