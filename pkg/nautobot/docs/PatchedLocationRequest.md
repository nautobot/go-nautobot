# PatchedLocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeZone** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
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
**LocationType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedLocationRequest

`func NewPatchedLocationRequest() *PatchedLocationRequest`

NewPatchedLocationRequest instantiates a new PatchedLocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLocationRequestWithDefaults

`func NewPatchedLocationRequestWithDefaults() *PatchedLocationRequest`

NewPatchedLocationRequestWithDefaults instantiates a new PatchedLocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeZone

`func (o *PatchedLocationRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PatchedLocationRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PatchedLocationRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PatchedLocationRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *PatchedLocationRequest) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *PatchedLocationRequest) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetName

`func (o *PatchedLocationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedLocationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedLocationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedLocationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedLocationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedLocationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedLocationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedLocationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *PatchedLocationRequest) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *PatchedLocationRequest) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *PatchedLocationRequest) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *PatchedLocationRequest) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetAsn

`func (o *PatchedLocationRequest) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *PatchedLocationRequest) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *PatchedLocationRequest) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *PatchedLocationRequest) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *PatchedLocationRequest) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *PatchedLocationRequest) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetPhysicalAddress

`func (o *PatchedLocationRequest) GetPhysicalAddress() string`

GetPhysicalAddress returns the PhysicalAddress field if non-nil, zero value otherwise.

### GetPhysicalAddressOk

`func (o *PatchedLocationRequest) GetPhysicalAddressOk() (*string, bool)`

GetPhysicalAddressOk returns a tuple with the PhysicalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAddress

`func (o *PatchedLocationRequest) SetPhysicalAddress(v string)`

SetPhysicalAddress sets PhysicalAddress field to given value.

### HasPhysicalAddress

`func (o *PatchedLocationRequest) HasPhysicalAddress() bool`

HasPhysicalAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *PatchedLocationRequest) GetShippingAddress() string`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *PatchedLocationRequest) GetShippingAddressOk() (*string, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *PatchedLocationRequest) SetShippingAddress(v string)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *PatchedLocationRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *PatchedLocationRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PatchedLocationRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PatchedLocationRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PatchedLocationRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *PatchedLocationRequest) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *PatchedLocationRequest) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *PatchedLocationRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PatchedLocationRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PatchedLocationRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PatchedLocationRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *PatchedLocationRequest) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *PatchedLocationRequest) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetContactName

`func (o *PatchedLocationRequest) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *PatchedLocationRequest) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *PatchedLocationRequest) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *PatchedLocationRequest) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactPhone

`func (o *PatchedLocationRequest) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *PatchedLocationRequest) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *PatchedLocationRequest) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.

### HasContactPhone

`func (o *PatchedLocationRequest) HasContactPhone() bool`

HasContactPhone returns a boolean if a field has been set.

### GetContactEmail

`func (o *PatchedLocationRequest) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *PatchedLocationRequest) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *PatchedLocationRequest) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *PatchedLocationRequest) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetComments

`func (o *PatchedLocationRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedLocationRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedLocationRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedLocationRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetParent

`func (o *PatchedLocationRequest) GetParent() BulkWritableCircuitRequestTenant`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedLocationRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedLocationRequest) SetParent(v BulkWritableCircuitRequestTenant)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedLocationRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedLocationRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedLocationRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetLocationType

`func (o *PatchedLocationRequest) GetLocationType() BulkWritableCableRequestStatus`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *PatchedLocationRequest) GetLocationTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *PatchedLocationRequest) SetLocationType(v BulkWritableCableRequestStatus)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *PatchedLocationRequest) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedLocationRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedLocationRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedLocationRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedLocationRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedLocationRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedLocationRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedLocationRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedLocationRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedLocationRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedLocationRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTags

`func (o *PatchedLocationRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedLocationRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedLocationRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedLocationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedLocationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedLocationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedLocationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedLocationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedLocationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedLocationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedLocationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedLocationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


