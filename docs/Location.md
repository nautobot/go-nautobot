# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**TreeDepth** | **NullableInt32** |  | [readonly] 
**TimeZone** | Pointer to **NullableString** |  | [optional] 
**CircuitCount** | Pointer to **int32** |  | [optional] [readonly] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] 
**PrefixCount** | Pointer to **int32** |  | [optional] [readonly] 
**RackCount** | Pointer to **int32** |  | [optional] [readonly] 
**VirtualMachineCount** | Pointer to **int32** |  | [optional] [readonly] 
**VlanCount** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to **string** | Local facility ID or description | [optional] 
**Asn** | Pointer to **NullableInt64** | 32-bit autonomous system number | [optional] 
**PhysicalAddress** | Pointer to **string** |  | [optional] 
**ShippingAddress** | Pointer to **string** |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | GPS coordinate (latitude) | [optional] 
**Longitude** | Pointer to **NullableFloat64** | GPS coordinate (longitude) | [optional] 
**ContactName** | Pointer to **string** |  | [optional] 
**ContactPhone** | Pointer to **string** |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**LocationType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewLocation

`func NewLocation(id string, objectType string, display string, url string, naturalSlug string, treeDepth NullableInt32, name string, locationType BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Location) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Location) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Location) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Location) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Location) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Location) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Location) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Location) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Location) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Location) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Location) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Location) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetTreeDepth

`func (o *Location) GetTreeDepth() int32`

GetTreeDepth returns the TreeDepth field if non-nil, zero value otherwise.

### GetTreeDepthOk

`func (o *Location) GetTreeDepthOk() (*int32, bool)`

GetTreeDepthOk returns a tuple with the TreeDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeDepth

`func (o *Location) SetTreeDepth(v int32)`

SetTreeDepth sets TreeDepth field to given value.


### SetTreeDepthNil

`func (o *Location) SetTreeDepthNil(b bool)`

 SetTreeDepthNil sets the value for TreeDepth to be an explicit nil

### UnsetTreeDepth
`func (o *Location) UnsetTreeDepth()`

UnsetTreeDepth ensures that no value is present for TreeDepth, not even an explicit nil
### GetTimeZone

`func (o *Location) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Location) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Location) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Location) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *Location) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *Location) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetCircuitCount

`func (o *Location) GetCircuitCount() int32`

GetCircuitCount returns the CircuitCount field if non-nil, zero value otherwise.

### GetCircuitCountOk

`func (o *Location) GetCircuitCountOk() (*int32, bool)`

GetCircuitCountOk returns a tuple with the CircuitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitCount

`func (o *Location) SetCircuitCount(v int32)`

SetCircuitCount sets CircuitCount field to given value.

### HasCircuitCount

`func (o *Location) HasCircuitCount() bool`

HasCircuitCount returns a boolean if a field has been set.

### GetDeviceCount

`func (o *Location) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *Location) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *Location) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *Location) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetPrefixCount

`func (o *Location) GetPrefixCount() int32`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *Location) GetPrefixCountOk() (*int32, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *Location) SetPrefixCount(v int32)`

SetPrefixCount sets PrefixCount field to given value.

### HasPrefixCount

`func (o *Location) HasPrefixCount() bool`

HasPrefixCount returns a boolean if a field has been set.

### GetRackCount

`func (o *Location) GetRackCount() int32`

GetRackCount returns the RackCount field if non-nil, zero value otherwise.

### GetRackCountOk

`func (o *Location) GetRackCountOk() (*int32, bool)`

GetRackCountOk returns a tuple with the RackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackCount

`func (o *Location) SetRackCount(v int32)`

SetRackCount sets RackCount field to given value.

### HasRackCount

`func (o *Location) HasRackCount() bool`

HasRackCount returns a boolean if a field has been set.

### GetVirtualMachineCount

`func (o *Location) GetVirtualMachineCount() int32`

GetVirtualMachineCount returns the VirtualMachineCount field if non-nil, zero value otherwise.

### GetVirtualMachineCountOk

`func (o *Location) GetVirtualMachineCountOk() (*int32, bool)`

GetVirtualMachineCountOk returns a tuple with the VirtualMachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineCount

`func (o *Location) SetVirtualMachineCount(v int32)`

SetVirtualMachineCount sets VirtualMachineCount field to given value.

### HasVirtualMachineCount

`func (o *Location) HasVirtualMachineCount() bool`

HasVirtualMachineCount returns a boolean if a field has been set.

### GetVlanCount

`func (o *Location) GetVlanCount() int32`

GetVlanCount returns the VlanCount field if non-nil, zero value otherwise.

### GetVlanCountOk

`func (o *Location) GetVlanCountOk() (*int32, bool)`

GetVlanCountOk returns a tuple with the VlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCount

`func (o *Location) SetVlanCount(v int32)`

SetVlanCount sets VlanCount field to given value.

### HasVlanCount

`func (o *Location) HasVlanCount() bool`

HasVlanCount returns a boolean if a field has been set.

### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Location) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Location) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Location) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Location) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *Location) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *Location) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *Location) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *Location) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetAsn

`func (o *Location) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *Location) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *Location) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *Location) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *Location) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *Location) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetPhysicalAddress

`func (o *Location) GetPhysicalAddress() string`

GetPhysicalAddress returns the PhysicalAddress field if non-nil, zero value otherwise.

### GetPhysicalAddressOk

`func (o *Location) GetPhysicalAddressOk() (*string, bool)`

GetPhysicalAddressOk returns a tuple with the PhysicalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAddress

`func (o *Location) SetPhysicalAddress(v string)`

SetPhysicalAddress sets PhysicalAddress field to given value.

### HasPhysicalAddress

`func (o *Location) HasPhysicalAddress() bool`

HasPhysicalAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *Location) GetShippingAddress() string`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *Location) GetShippingAddressOk() (*string, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *Location) SetShippingAddress(v string)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *Location) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *Location) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Location) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Location) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Location) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *Location) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *Location) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *Location) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Location) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Location) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Location) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *Location) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *Location) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetContactName

`func (o *Location) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *Location) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *Location) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *Location) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactPhone

`func (o *Location) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *Location) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *Location) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.

### HasContactPhone

`func (o *Location) HasContactPhone() bool`

HasContactPhone returns a boolean if a field has been set.

### GetContactEmail

`func (o *Location) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Location) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Location) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Location) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetComments

`func (o *Location) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Location) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Location) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Location) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetParent

`func (o *Location) GetParent() BulkWritableCircuitRequestTenant`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Location) GetParentOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Location) SetParent(v BulkWritableCircuitRequestTenant)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Location) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *Location) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *Location) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetLocationType

`func (o *Location) GetLocationType() BulkWritableCableRequestStatus`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *Location) GetLocationTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *Location) SetLocationType(v BulkWritableCableRequestStatus)`

SetLocationType sets LocationType field to given value.


### GetStatus

`func (o *Location) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Location) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Location) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *Location) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Location) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Location) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Location) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Location) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Location) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCreated

`func (o *Location) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Location) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Location) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Location) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Location) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Location) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Location) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Location) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Location) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Location) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *Location) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Location) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Location) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *Location) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Location) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Location) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Location) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *Location) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Location) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Location) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Location) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


