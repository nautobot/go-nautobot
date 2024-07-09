# WritablePrefixRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | **string** |  | 
**Location** | Pointer to [**NullableBulkWritablePrefixRequestLocation**](BulkWritablePrefixRequestLocation.md) |  | [optional] 
**Type** | Pointer to [**PrefixTypeChoices**](PrefixTypeChoices.md) |  | [optional] 
**DateAllocated** | Pointer to **NullableTime** | Date this prefix was allocated to an RIR, reserved in IPAM, etc. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
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

### NewWritablePrefixRequest

`func NewWritablePrefixRequest(prefix string, status BulkWritableCableRequestStatus, ) *WritablePrefixRequest`

NewWritablePrefixRequest instantiates a new WritablePrefixRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePrefixRequestWithDefaults

`func NewWritablePrefixRequestWithDefaults() *WritablePrefixRequest`

NewWritablePrefixRequestWithDefaults instantiates a new WritablePrefixRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *WritablePrefixRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *WritablePrefixRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *WritablePrefixRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetLocation

`func (o *WritablePrefixRequest) GetLocation() BulkWritablePrefixRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WritablePrefixRequest) GetLocationOk() (*BulkWritablePrefixRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WritablePrefixRequest) SetLocation(v BulkWritablePrefixRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WritablePrefixRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *WritablePrefixRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *WritablePrefixRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetType

`func (o *WritablePrefixRequest) GetType() PrefixTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritablePrefixRequest) GetTypeOk() (*PrefixTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritablePrefixRequest) SetType(v PrefixTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *WritablePrefixRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDateAllocated

`func (o *WritablePrefixRequest) GetDateAllocated() time.Time`

GetDateAllocated returns the DateAllocated field if non-nil, zero value otherwise.

### GetDateAllocatedOk

`func (o *WritablePrefixRequest) GetDateAllocatedOk() (*time.Time, bool)`

GetDateAllocatedOk returns a tuple with the DateAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAllocated

`func (o *WritablePrefixRequest) SetDateAllocated(v time.Time)`

SetDateAllocated sets DateAllocated field to given value.

### HasDateAllocated

`func (o *WritablePrefixRequest) HasDateAllocated() bool`

HasDateAllocated returns a boolean if a field has been set.

### SetDateAllocatedNil

`func (o *WritablePrefixRequest) SetDateAllocatedNil(b bool)`

 SetDateAllocatedNil sets the value for DateAllocated to be an explicit nil

### UnsetDateAllocated
`func (o *WritablePrefixRequest) UnsetDateAllocated()`

UnsetDateAllocated ensures that no value is present for DateAllocated, not even an explicit nil
### GetDescription

`func (o *WritablePrefixRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritablePrefixRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritablePrefixRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritablePrefixRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *WritablePrefixRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritablePrefixRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritablePrefixRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *WritablePrefixRequest) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WritablePrefixRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WritablePrefixRequest) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *WritablePrefixRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *WritablePrefixRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *WritablePrefixRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetParent

`func (o *WritablePrefixRequest) GetParent() BulkWritablePrefixRequestParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WritablePrefixRequest) GetParentOk() (*BulkWritablePrefixRequestParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WritablePrefixRequest) SetParent(v BulkWritablePrefixRequestParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *WritablePrefixRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *WritablePrefixRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *WritablePrefixRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetNamespace

`func (o *WritablePrefixRequest) GetNamespace() BulkWritableCableRequestStatus`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *WritablePrefixRequest) GetNamespaceOk() (*BulkWritableCableRequestStatus, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *WritablePrefixRequest) SetNamespace(v BulkWritableCableRequestStatus)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *WritablePrefixRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTenant

`func (o *WritablePrefixRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritablePrefixRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritablePrefixRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritablePrefixRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritablePrefixRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritablePrefixRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetVlan

`func (o *WritablePrefixRequest) GetVlan() BulkWritableCircuitRequestTenant`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *WritablePrefixRequest) GetVlanOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *WritablePrefixRequest) SetVlan(v BulkWritableCircuitRequestTenant)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *WritablePrefixRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *WritablePrefixRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *WritablePrefixRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetRir

`func (o *WritablePrefixRequest) GetRir() BulkWritablePrefixRequestRir`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *WritablePrefixRequest) GetRirOk() (*BulkWritablePrefixRequestRir, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *WritablePrefixRequest) SetRir(v BulkWritablePrefixRequestRir)`

SetRir sets Rir field to given value.

### HasRir

`func (o *WritablePrefixRequest) HasRir() bool`

HasRir returns a boolean if a field has been set.

### SetRirNil

`func (o *WritablePrefixRequest) SetRirNil(b bool)`

 SetRirNil sets the value for Rir to be an explicit nil

### UnsetRir
`func (o *WritablePrefixRequest) UnsetRir()`

UnsetRir ensures that no value is present for Rir, not even an explicit nil
### GetTags

`func (o *WritablePrefixRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritablePrefixRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritablePrefixRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritablePrefixRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritablePrefixRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePrefixRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePrefixRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePrefixRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritablePrefixRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritablePrefixRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritablePrefixRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritablePrefixRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


