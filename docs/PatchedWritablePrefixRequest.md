# PatchedWritablePrefixRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**NullableBulkWritablePrefixRequestLocation**](BulkWritablePrefixRequestLocation.md) |  | [optional] 
**Type** | Pointer to [**PrefixTypeChoices**](PrefixTypeChoices.md) |  | [optional] 
**DateAllocated** | Pointer to **NullableTime** | Date this prefix was allocated to an RIR, reserved in IPAM, etc. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Parent** | Pointer to [**NullableBulkWritablePrefixRequestParent**](BulkWritablePrefixRequestParent.md) |  | [optional] 
**Namespace** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Vlan** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Rir** | Pointer to [**NullableBulkWritablePrefixRequestRir**](BulkWritablePrefixRequestRir.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedWritablePrefixRequest

`func NewPatchedWritablePrefixRequest() *PatchedWritablePrefixRequest`

NewPatchedWritablePrefixRequest instantiates a new PatchedWritablePrefixRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritablePrefixRequestWithDefaults

`func NewPatchedWritablePrefixRequestWithDefaults() *PatchedWritablePrefixRequest`

NewPatchedWritablePrefixRequestWithDefaults instantiates a new PatchedWritablePrefixRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *PatchedWritablePrefixRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PatchedWritablePrefixRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PatchedWritablePrefixRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PatchedWritablePrefixRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedWritablePrefixRequest) GetLocation() BulkWritablePrefixRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedWritablePrefixRequest) GetLocationOk() (*BulkWritablePrefixRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedWritablePrefixRequest) SetLocation(v BulkWritablePrefixRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedWritablePrefixRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedWritablePrefixRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedWritablePrefixRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetType

`func (o *PatchedWritablePrefixRequest) GetType() PrefixTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritablePrefixRequest) GetTypeOk() (*PrefixTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritablePrefixRequest) SetType(v PrefixTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritablePrefixRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDateAllocated

`func (o *PatchedWritablePrefixRequest) GetDateAllocated() time.Time`

GetDateAllocated returns the DateAllocated field if non-nil, zero value otherwise.

### GetDateAllocatedOk

`func (o *PatchedWritablePrefixRequest) GetDateAllocatedOk() (*time.Time, bool)`

GetDateAllocatedOk returns a tuple with the DateAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAllocated

`func (o *PatchedWritablePrefixRequest) SetDateAllocated(v time.Time)`

SetDateAllocated sets DateAllocated field to given value.

### HasDateAllocated

`func (o *PatchedWritablePrefixRequest) HasDateAllocated() bool`

HasDateAllocated returns a boolean if a field has been set.

### SetDateAllocatedNil

`func (o *PatchedWritablePrefixRequest) SetDateAllocatedNil(b bool)`

 SetDateAllocatedNil sets the value for DateAllocated to be an explicit nil

### UnsetDateAllocated
`func (o *PatchedWritablePrefixRequest) UnsetDateAllocated()`

UnsetDateAllocated ensures that no value is present for DateAllocated, not even an explicit nil
### GetDescription

`func (o *PatchedWritablePrefixRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritablePrefixRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritablePrefixRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritablePrefixRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritablePrefixRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritablePrefixRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritablePrefixRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritablePrefixRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritablePrefixRequest) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritablePrefixRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritablePrefixRequest) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritablePrefixRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedWritablePrefixRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedWritablePrefixRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetParent

`func (o *PatchedWritablePrefixRequest) GetParent() BulkWritablePrefixRequestParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedWritablePrefixRequest) GetParentOk() (*BulkWritablePrefixRequestParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedWritablePrefixRequest) SetParent(v BulkWritablePrefixRequestParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedWritablePrefixRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedWritablePrefixRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedWritablePrefixRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetNamespace

`func (o *PatchedWritablePrefixRequest) GetNamespace() BulkWritableCableRequestStatus`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchedWritablePrefixRequest) GetNamespaceOk() (*BulkWritableCableRequestStatus, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchedWritablePrefixRequest) SetNamespace(v BulkWritableCableRequestStatus)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PatchedWritablePrefixRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritablePrefixRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritablePrefixRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritablePrefixRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritablePrefixRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritablePrefixRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritablePrefixRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetVlan

`func (o *PatchedWritablePrefixRequest) GetVlan() BulkWritableCircuitRequestTenant`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PatchedWritablePrefixRequest) GetVlanOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PatchedWritablePrefixRequest) SetVlan(v BulkWritableCircuitRequestTenant)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PatchedWritablePrefixRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *PatchedWritablePrefixRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *PatchedWritablePrefixRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetRir

`func (o *PatchedWritablePrefixRequest) GetRir() BulkWritablePrefixRequestRir`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *PatchedWritablePrefixRequest) GetRirOk() (*BulkWritablePrefixRequestRir, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *PatchedWritablePrefixRequest) SetRir(v BulkWritablePrefixRequestRir)`

SetRir sets Rir field to given value.

### HasRir

`func (o *PatchedWritablePrefixRequest) HasRir() bool`

HasRir returns a boolean if a field has been set.

### SetRirNil

`func (o *PatchedWritablePrefixRequest) SetRirNil(b bool)`

 SetRirNil sets the value for Rir to be an explicit nil

### UnsetRir
`func (o *PatchedWritablePrefixRequest) UnsetRir()`

UnsetRir ensures that no value is present for Rir, not even an explicit nil
### GetTags

`func (o *PatchedWritablePrefixRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritablePrefixRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritablePrefixRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritablePrefixRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritablePrefixRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritablePrefixRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritablePrefixRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritablePrefixRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritablePrefixRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritablePrefixRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritablePrefixRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritablePrefixRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


