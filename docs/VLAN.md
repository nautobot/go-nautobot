# VLAN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**PrefixCount** | Pointer to **int32** |  | [optional] [readonly] 
**Vid** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**VlanGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Role** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Locations** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVLAN

`func NewVLAN(id string, objectType string, display string, url string, naturalSlug string, vid int32, name string, status BulkWritableCableRequestStatus, locations []BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *VLAN`

NewVLAN instantiates a new VLAN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVLANWithDefaults

`func NewVLANWithDefaults() *VLAN`

NewVLANWithDefaults instantiates a new VLAN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VLAN) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VLAN) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VLAN) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *VLAN) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VLAN) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VLAN) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VLAN) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VLAN) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VLAN) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VLAN) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VLAN) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VLAN) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VLAN) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VLAN) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VLAN) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetPrefixCount

`func (o *VLAN) GetPrefixCount() int32`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *VLAN) GetPrefixCountOk() (*int32, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *VLAN) SetPrefixCount(v int32)`

SetPrefixCount sets PrefixCount field to given value.

### HasPrefixCount

`func (o *VLAN) HasPrefixCount() bool`

HasPrefixCount returns a boolean if a field has been set.

### GetVid

`func (o *VLAN) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *VLAN) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *VLAN) SetVid(v int32)`

SetVid sets Vid field to given value.


### GetName

`func (o *VLAN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VLAN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VLAN) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VLAN) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VLAN) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VLAN) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VLAN) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVlanGroup

`func (o *VLAN) GetVlanGroup() BulkWritableCircuitRequestTenant`

GetVlanGroup returns the VlanGroup field if non-nil, zero value otherwise.

### GetVlanGroupOk

`func (o *VLAN) GetVlanGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVlanGroupOk returns a tuple with the VlanGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroup

`func (o *VLAN) SetVlanGroup(v BulkWritableCircuitRequestTenant)`

SetVlanGroup sets VlanGroup field to given value.

### HasVlanGroup

`func (o *VLAN) HasVlanGroup() bool`

HasVlanGroup returns a boolean if a field has been set.

### SetVlanGroupNil

`func (o *VLAN) SetVlanGroupNil(b bool)`

 SetVlanGroupNil sets the value for VlanGroup to be an explicit nil

### UnsetVlanGroup
`func (o *VLAN) UnsetVlanGroup()`

UnsetVlanGroup ensures that no value is present for VlanGroup, not even an explicit nil
### GetStatus

`func (o *VLAN) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VLAN) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VLAN) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *VLAN) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VLAN) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VLAN) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *VLAN) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VLAN) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VLAN) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *VLAN) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VLAN) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VLAN) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VLAN) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VLAN) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VLAN) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetLocations

`func (o *VLAN) GetLocations() []BulkWritableCableRequestStatus`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *VLAN) GetLocationsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *VLAN) SetLocations(v []BulkWritableCableRequestStatus)`

SetLocations sets Locations field to given value.


### GetCreated

`func (o *VLAN) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VLAN) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VLAN) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VLAN) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VLAN) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VLAN) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VLAN) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VLAN) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VLAN) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VLAN) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *VLAN) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VLAN) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VLAN) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VLAN) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *VLAN) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VLAN) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VLAN) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *VLAN) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VLAN) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VLAN) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VLAN) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


