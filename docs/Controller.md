# Controller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Capabilities** | Pointer to [**[]ControllerCapabilitiesInner**](ControllerCapabilitiesInner.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Location** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Platform** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ExternalIntegration** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ControllerDevice** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ControllerDeviceRedundancyGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewController

`func NewController(objectType string, display string, url string, naturalSlug string, name string, status BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Controller`

NewController instantiates a new Controller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerWithDefaults

`func NewControllerWithDefaults() *Controller`

NewControllerWithDefaults instantiates a new Controller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Controller) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Controller) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Controller) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Controller) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *Controller) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Controller) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Controller) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Controller) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Controller) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Controller) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Controller) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Controller) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Controller) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Controller) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Controller) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Controller) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCapabilities

`func (o *Controller) GetCapabilities() []ControllerCapabilitiesInner`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Controller) GetCapabilitiesOk() (*[]ControllerCapabilitiesInner, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Controller) SetCapabilities(v []ControllerCapabilitiesInner)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Controller) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *Controller) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *Controller) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetName

`func (o *Controller) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Controller) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Controller) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Controller) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Controller) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Controller) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Controller) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *Controller) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Controller) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Controller) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetLocation

`func (o *Controller) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Controller) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Controller) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.


### GetPlatform

`func (o *Controller) GetPlatform() ApprovalWorkflowUser`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Controller) GetPlatformOk() (*ApprovalWorkflowUser, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Controller) SetPlatform(v ApprovalWorkflowUser)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Controller) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *Controller) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Controller) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetRole

`func (o *Controller) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Controller) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Controller) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *Controller) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *Controller) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *Controller) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *Controller) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Controller) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Controller) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Controller) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Controller) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Controller) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetExternalIntegration

`func (o *Controller) GetExternalIntegration() ApprovalWorkflowUser`

GetExternalIntegration returns the ExternalIntegration field if non-nil, zero value otherwise.

### GetExternalIntegrationOk

`func (o *Controller) GetExternalIntegrationOk() (*ApprovalWorkflowUser, bool)`

GetExternalIntegrationOk returns a tuple with the ExternalIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIntegration

`func (o *Controller) SetExternalIntegration(v ApprovalWorkflowUser)`

SetExternalIntegration sets ExternalIntegration field to given value.

### HasExternalIntegration

`func (o *Controller) HasExternalIntegration() bool`

HasExternalIntegration returns a boolean if a field has been set.

### SetExternalIntegrationNil

`func (o *Controller) SetExternalIntegrationNil(b bool)`

 SetExternalIntegrationNil sets the value for ExternalIntegration to be an explicit nil

### UnsetExternalIntegration
`func (o *Controller) UnsetExternalIntegration()`

UnsetExternalIntegration ensures that no value is present for ExternalIntegration, not even an explicit nil
### GetControllerDevice

`func (o *Controller) GetControllerDevice() ApprovalWorkflowUser`

GetControllerDevice returns the ControllerDevice field if non-nil, zero value otherwise.

### GetControllerDeviceOk

`func (o *Controller) GetControllerDeviceOk() (*ApprovalWorkflowUser, bool)`

GetControllerDeviceOk returns a tuple with the ControllerDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDevice

`func (o *Controller) SetControllerDevice(v ApprovalWorkflowUser)`

SetControllerDevice sets ControllerDevice field to given value.

### HasControllerDevice

`func (o *Controller) HasControllerDevice() bool`

HasControllerDevice returns a boolean if a field has been set.

### SetControllerDeviceNil

`func (o *Controller) SetControllerDeviceNil(b bool)`

 SetControllerDeviceNil sets the value for ControllerDevice to be an explicit nil

### UnsetControllerDevice
`func (o *Controller) UnsetControllerDevice()`

UnsetControllerDevice ensures that no value is present for ControllerDevice, not even an explicit nil
### GetControllerDeviceRedundancyGroup

`func (o *Controller) GetControllerDeviceRedundancyGroup() ApprovalWorkflowUser`

GetControllerDeviceRedundancyGroup returns the ControllerDeviceRedundancyGroup field if non-nil, zero value otherwise.

### GetControllerDeviceRedundancyGroupOk

`func (o *Controller) GetControllerDeviceRedundancyGroupOk() (*ApprovalWorkflowUser, bool)`

GetControllerDeviceRedundancyGroupOk returns a tuple with the ControllerDeviceRedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDeviceRedundancyGroup

`func (o *Controller) SetControllerDeviceRedundancyGroup(v ApprovalWorkflowUser)`

SetControllerDeviceRedundancyGroup sets ControllerDeviceRedundancyGroup field to given value.

### HasControllerDeviceRedundancyGroup

`func (o *Controller) HasControllerDeviceRedundancyGroup() bool`

HasControllerDeviceRedundancyGroup returns a boolean if a field has been set.

### SetControllerDeviceRedundancyGroupNil

`func (o *Controller) SetControllerDeviceRedundancyGroupNil(b bool)`

 SetControllerDeviceRedundancyGroupNil sets the value for ControllerDeviceRedundancyGroup to be an explicit nil

### UnsetControllerDeviceRedundancyGroup
`func (o *Controller) UnsetControllerDeviceRedundancyGroup()`

UnsetControllerDeviceRedundancyGroup ensures that no value is present for ControllerDeviceRedundancyGroup, not even an explicit nil
### GetCreated

`func (o *Controller) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Controller) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Controller) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Controller) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Controller) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Controller) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Controller) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Controller) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Controller) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Controller) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *Controller) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Controller) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Controller) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *Controller) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Controller) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Controller) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Controller) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *Controller) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Controller) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Controller) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Controller) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


