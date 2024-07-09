# PatchedVLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**NullableBulkWritablePrefixRequestLocation**](BulkWritablePrefixRequestLocation.md) |  | [optional] 
**Vid** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**VlanGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedVLANRequest

`func NewPatchedVLANRequest() *PatchedVLANRequest`

NewPatchedVLANRequest instantiates a new PatchedVLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVLANRequestWithDefaults

`func NewPatchedVLANRequestWithDefaults() *PatchedVLANRequest`

NewPatchedVLANRequestWithDefaults instantiates a new PatchedVLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *PatchedVLANRequest) GetLocation() BulkWritablePrefixRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedVLANRequest) GetLocationOk() (*BulkWritablePrefixRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedVLANRequest) SetLocation(v BulkWritablePrefixRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedVLANRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedVLANRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedVLANRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetVid

`func (o *PatchedVLANRequest) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *PatchedVLANRequest) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *PatchedVLANRequest) SetVid(v int32)`

SetVid sets Vid field to given value.

### HasVid

`func (o *PatchedVLANRequest) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetName

`func (o *PatchedVLANRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedVLANRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedVLANRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedVLANRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedVLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedVLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedVLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedVLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVlanGroup

`func (o *PatchedVLANRequest) GetVlanGroup() BulkWritableCircuitRequestTenant`

GetVlanGroup returns the VlanGroup field if non-nil, zero value otherwise.

### GetVlanGroupOk

`func (o *PatchedVLANRequest) GetVlanGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVlanGroupOk returns a tuple with the VlanGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroup

`func (o *PatchedVLANRequest) SetVlanGroup(v BulkWritableCircuitRequestTenant)`

SetVlanGroup sets VlanGroup field to given value.

### HasVlanGroup

`func (o *PatchedVLANRequest) HasVlanGroup() bool`

HasVlanGroup returns a boolean if a field has been set.

### SetVlanGroupNil

`func (o *PatchedVLANRequest) SetVlanGroupNil(b bool)`

 SetVlanGroupNil sets the value for VlanGroup to be an explicit nil

### UnsetVlanGroup
`func (o *PatchedVLANRequest) UnsetVlanGroup()`

UnsetVlanGroup ensures that no value is present for VlanGroup, not even an explicit nil
### GetStatus

`func (o *PatchedVLANRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedVLANRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedVLANRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedVLANRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedVLANRequest) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedVLANRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedVLANRequest) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedVLANRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedVLANRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedVLANRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *PatchedVLANRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedVLANRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedVLANRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedVLANRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedVLANRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedVLANRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTags

`func (o *PatchedVLANRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedVLANRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedVLANRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedVLANRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedVLANRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedVLANRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedVLANRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedVLANRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedVLANRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedVLANRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedVLANRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedVLANRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


