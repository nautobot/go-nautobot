# ControllerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Location** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
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

### NewControllerRequest

`func NewControllerRequest(name string, status BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus, ) *ControllerRequest`

NewControllerRequest instantiates a new ControllerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerRequestWithDefaults

`func NewControllerRequestWithDefaults() *ControllerRequest`

NewControllerRequestWithDefaults instantiates a new ControllerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *ControllerRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ControllerRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ControllerRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetLocation

`func (o *ControllerRequest) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ControllerRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ControllerRequest) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.


### GetPlatform

`func (o *ControllerRequest) GetPlatform() BulkWritableCircuitRequestTenant`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ControllerRequest) GetPlatformOk() (*BulkWritableCircuitRequestTenant, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ControllerRequest) SetPlatform(v BulkWritableCircuitRequestTenant)`

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

`func (o *ControllerRequest) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ControllerRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ControllerRequest) SetRole(v BulkWritableCircuitRequestTenant)`

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

`func (o *ControllerRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ControllerRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ControllerRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

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

`func (o *ControllerRequest) GetExternalIntegration() BulkWritableCircuitRequestTenant`

GetExternalIntegration returns the ExternalIntegration field if non-nil, zero value otherwise.

### GetExternalIntegrationOk

`func (o *ControllerRequest) GetExternalIntegrationOk() (*BulkWritableCircuitRequestTenant, bool)`

GetExternalIntegrationOk returns a tuple with the ExternalIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegration

`func (o *ControllerRequest) SetExternalIntegration(v BulkWritableCircuitRequestTenant)`

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

`func (o *ControllerRequest) GetControllerDevice() BulkWritableCircuitRequestTenant`

GetControllerDevice returns the ControllerDevice field if non-nil, zero value otherwise.

### GetControllerDeviceOk

`func (o *ControllerRequest) GetControllerDeviceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetControllerDeviceOk returns a tuple with the ControllerDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDevice

`func (o *ControllerRequest) SetControllerDevice(v BulkWritableCircuitRequestTenant)`

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

`func (o *ControllerRequest) GetControllerDeviceRedundancyGroup() BulkWritableCircuitRequestTenant`

GetControllerDeviceRedundancyGroup returns the ControllerDeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetControllerDeviceRedundancyGroupOk

`func (o *ControllerRequest) GetControllerDeviceRedundancyGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetControllerDeviceRedundancyGroupOk returns a tuple with the ControllerDeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDeviceRedundancyGroup

`func (o *ControllerRequest) SetControllerDeviceRedundancyGroup(v BulkWritableCircuitRequestTenant)`

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
### GetTags

`func (o *ControllerRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ControllerRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ControllerRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ControllerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

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

`func (o *ControllerRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ControllerRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ControllerRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ControllerRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


