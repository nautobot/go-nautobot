# PatchedBulkWritableLocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableLocationRequest

`func NewPatchedBulkWritableLocationRequest(id string, ) *PatchedBulkWritableLocationRequest`

NewPatchedBulkWritableLocationRequest instantiates a new PatchedBulkWritableLocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableLocationRequestWithDefaults

`func NewPatchedBulkWritableLocationRequestWithDefaults() *PatchedBulkWritableLocationRequest`

NewPatchedBulkWritableLocationRequestWithDefaults instantiates a new PatchedBulkWritableLocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableLocationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableLocationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableLocationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetTimeZone

`func (o *PatchedBulkWritableLocationRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PatchedBulkWritableLocationRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PatchedBulkWritableLocationRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PatchedBulkWritableLocationRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *PatchedBulkWritableLocationRequest) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *PatchedBulkWritableLocationRequest) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetName

`func (o *PatchedBulkWritableLocationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableLocationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableLocationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableLocationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableLocationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableLocationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableLocationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableLocationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *PatchedBulkWritableLocationRequest) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *PatchedBulkWritableLocationRequest) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *PatchedBulkWritableLocationRequest) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *PatchedBulkWritableLocationRequest) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetAsn

`func (o *PatchedBulkWritableLocationRequest) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *PatchedBulkWritableLocationRequest) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *PatchedBulkWritableLocationRequest) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *PatchedBulkWritableLocationRequest) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *PatchedBulkWritableLocationRequest) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *PatchedBulkWritableLocationRequest) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetPhysicalAddress

`func (o *PatchedBulkWritableLocationRequest) GetPhysicalAddress() string`

GetPhysicalAddress returns the PhysicalAddress field if non-nil, zero value otherwise.

### GetPhysicalAddressOk

`func (o *PatchedBulkWritableLocationRequest) GetPhysicalAddressOk() (*string, bool)`

GetPhysicalAddressOk returns a tuple with the PhysicalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAddress

`func (o *PatchedBulkWritableLocationRequest) SetPhysicalAddress(v string)`

SetPhysicalAddress sets PhysicalAddress field to given value.

### HasPhysicalAddress

`func (o *PatchedBulkWritableLocationRequest) HasPhysicalAddress() bool`

HasPhysicalAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *PatchedBulkWritableLocationRequest) GetShippingAddress() string`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *PatchedBulkWritableLocationRequest) GetShippingAddressOk() (*string, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *PatchedBulkWritableLocationRequest) SetShippingAddress(v string)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *PatchedBulkWritableLocationRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *PatchedBulkWritableLocationRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PatchedBulkWritableLocationRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PatchedBulkWritableLocationRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PatchedBulkWritableLocationRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *PatchedBulkWritableLocationRequest) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *PatchedBulkWritableLocationRequest) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *PatchedBulkWritableLocationRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PatchedBulkWritableLocationRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PatchedBulkWritableLocationRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PatchedBulkWritableLocationRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *PatchedBulkWritableLocationRequest) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *PatchedBulkWritableLocationRequest) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetContactName

`func (o *PatchedBulkWritableLocationRequest) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *PatchedBulkWritableLocationRequest) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *PatchedBulkWritableLocationRequest) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *PatchedBulkWritableLocationRequest) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactPhone

`func (o *PatchedBulkWritableLocationRequest) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *PatchedBulkWritableLocationRequest) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *PatchedBulkWritableLocationRequest) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.

### HasContactPhone

`func (o *PatchedBulkWritableLocationRequest) HasContactPhone() bool`

HasContactPhone returns a boolean if a field has been set.

### GetContactEmail

`func (o *PatchedBulkWritableLocationRequest) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *PatchedBulkWritableLocationRequest) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *PatchedBulkWritableLocationRequest) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *PatchedBulkWritableLocationRequest) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetComments

`func (o *PatchedBulkWritableLocationRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkWritableLocationRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkWritableLocationRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkWritableLocationRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetParent

`func (o *PatchedBulkWritableLocationRequest) GetParent() BulkWritableCircuitRequestTenant`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedBulkWritableLocationRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedBulkWritableLocationRequest) SetParent(v BulkWritableCircuitRequestTenant)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedBulkWritableLocationRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedBulkWritableLocationRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedBulkWritableLocationRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetLocationType

`func (o *PatchedBulkWritableLocationRequest) GetLocationType() BulkWritableCableRequestStatus`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *PatchedBulkWritableLocationRequest) GetLocationTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *PatchedBulkWritableLocationRequest) SetLocationType(v BulkWritableCableRequestStatus)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *PatchedBulkWritableLocationRequest) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkWritableLocationRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableLocationRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableLocationRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableLocationRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkWritableLocationRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableLocationRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableLocationRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableLocationRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableLocationRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableLocationRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableLocationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableLocationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableLocationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableLocationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableLocationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableLocationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableLocationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableLocationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableLocationRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableLocationRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableLocationRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableLocationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

