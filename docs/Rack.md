# Rack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Type** | Pointer to [**RackType**](RackType.md) |  | [optional] 
**Width** | Pointer to [**RackWidth**](RackWidth.md) |  | [optional] 
**OuterUnit** | Pointer to [**RackOuterUnit**](RackOuterUnit.md) |  | [optional] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] 
**PowerFeedCount** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**FacilityId** | Pointer to **NullableString** | Locally-assigned identifier | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this rack | [optional] 
**UHeight** | Pointer to **int32** | Height in rack units | [optional] 
**DescUnits** | Pointer to **bool** | Units are numbered top-to-bottom | [optional] 
**OuterWidth** | Pointer to **NullableInt32** | Outer dimension of rack (width) | [optional] 
**OuterDepth** | Pointer to **NullableInt32** | Outer dimension of rack (depth) | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Role** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Location** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**RackGroup** | Pointer to [**NullableBulkWritableRackRequestRackGroup**](BulkWritableRackRequestRackGroup.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewRack

`func NewRack(id string, objectType string, display string, url string, naturalSlug string, name string, status BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Rack`

NewRack instantiates a new Rack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackWithDefaults

`func NewRackWithDefaults() *Rack`

NewRackWithDefaults instantiates a new Rack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rack) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rack) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rack) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Rack) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Rack) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Rack) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Rack) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Rack) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Rack) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Rack) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Rack) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Rack) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Rack) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Rack) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Rack) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetType

`func (o *Rack) GetType() RackType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Rack) GetTypeOk() (*RackType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Rack) SetType(v RackType)`

SetType sets Type field to given value.

### HasType

`func (o *Rack) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWidth

`func (o *Rack) GetWidth() RackWidth`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Rack) GetWidthOk() (*RackWidth, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Rack) SetWidth(v RackWidth)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Rack) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetOuterUnit

`func (o *Rack) GetOuterUnit() RackOuterUnit`

GetOuterUnit returns the OuterUnit field if non-nil, zero value otherwise.

### GetOuterUnitOk

`func (o *Rack) GetOuterUnitOk() (*RackOuterUnit, bool)`

GetOuterUnitOk returns a tuple with the OuterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterUnit

`func (o *Rack) SetOuterUnit(v RackOuterUnit)`

SetOuterUnit sets OuterUnit field to given value.

### HasOuterUnit

`func (o *Rack) HasOuterUnit() bool`

HasOuterUnit returns a boolean if a field has been set.

### GetDeviceCount

`func (o *Rack) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *Rack) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *Rack) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *Rack) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetPowerFeedCount

`func (o *Rack) GetPowerFeedCount() int32`

GetPowerFeedCount returns the PowerFeedCount field if non-nil, zero value otherwise.

### GetPowerFeedCountOk

`func (o *Rack) GetPowerFeedCountOk() (*int32, bool)`

GetPowerFeedCountOk returns a tuple with the PowerFeedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFeedCount

`func (o *Rack) SetPowerFeedCount(v int32)`

SetPowerFeedCount sets PowerFeedCount field to given value.

### HasPowerFeedCount

`func (o *Rack) HasPowerFeedCount() bool`

HasPowerFeedCount returns a boolean if a field has been set.

### GetName

`func (o *Rack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Rack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Rack) SetName(v string)`

SetName sets Name field to given value.


### GetFacilityId

`func (o *Rack) GetFacilityId() string`

GetFacilityId returns the FacilityId field if non-nil, zero value otherwise.

### GetFacilityIdOk

`func (o *Rack) GetFacilityIdOk() (*string, bool)`

GetFacilityIdOk returns a tuple with the FacilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityId

`func (o *Rack) SetFacilityId(v string)`

SetFacilityId sets FacilityId field to given value.

### HasFacilityId

`func (o *Rack) HasFacilityId() bool`

HasFacilityId returns a boolean if a field has been set.

### SetFacilityIdNil

`func (o *Rack) SetFacilityIdNil(b bool)`

 SetFacilityIdNil sets the value for FacilityId to be an explicit nil

### UnsetFacilityId
`func (o *Rack) UnsetFacilityId()`

UnsetFacilityId ensures that no value is present for FacilityId, not even an explicit nil
### GetSerial

`func (o *Rack) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Rack) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Rack) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Rack) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *Rack) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *Rack) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *Rack) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *Rack) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *Rack) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *Rack) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetUHeight

`func (o *Rack) GetUHeight() int32`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *Rack) GetUHeightOk() (*int32, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *Rack) SetUHeight(v int32)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *Rack) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetDescUnits

`func (o *Rack) GetDescUnits() bool`

GetDescUnits returns the DescUnits field if non-nil, zero value otherwise.

### GetDescUnitsOk

`func (o *Rack) GetDescUnitsOk() (*bool, bool)`

GetDescUnitsOk returns a tuple with the DescUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescUnits

`func (o *Rack) SetDescUnits(v bool)`

SetDescUnits sets DescUnits field to given value.

### HasDescUnits

`func (o *Rack) HasDescUnits() bool`

HasDescUnits returns a boolean if a field has been set.

### GetOuterWidth

`func (o *Rack) GetOuterWidth() int32`

GetOuterWidth returns the OuterWidth field if non-nil, zero value otherwise.

### GetOuterWidthOk

`func (o *Rack) GetOuterWidthOk() (*int32, bool)`

GetOuterWidthOk returns a tuple with the OuterWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterWidth

`func (o *Rack) SetOuterWidth(v int32)`

SetOuterWidth sets OuterWidth field to given value.

### HasOuterWidth

`func (o *Rack) HasOuterWidth() bool`

HasOuterWidth returns a boolean if a field has been set.

### SetOuterWidthNil

`func (o *Rack) SetOuterWidthNil(b bool)`

 SetOuterWidthNil sets the value for OuterWidth to be an explicit nil

### UnsetOuterWidth
`func (o *Rack) UnsetOuterWidth()`

UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
### GetOuterDepth

`func (o *Rack) GetOuterDepth() int32`

GetOuterDepth returns the OuterDepth field if non-nil, zero value otherwise.

### GetOuterDepthOk

`func (o *Rack) GetOuterDepthOk() (*int32, bool)`

GetOuterDepthOk returns a tuple with the OuterDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterDepth

`func (o *Rack) SetOuterDepth(v int32)`

SetOuterDepth sets OuterDepth field to given value.

### HasOuterDepth

`func (o *Rack) HasOuterDepth() bool`

HasOuterDepth returns a boolean if a field has been set.

### SetOuterDepthNil

`func (o *Rack) SetOuterDepthNil(b bool)`

 SetOuterDepthNil sets the value for OuterDepth to be an explicit nil

### UnsetOuterDepth
`func (o *Rack) UnsetOuterDepth()`

UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
### GetComments

`func (o *Rack) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Rack) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Rack) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Rack) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetStatus

`func (o *Rack) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Rack) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Rack) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *Rack) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Rack) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Rack) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *Rack) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *Rack) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *Rack) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetLocation

`func (o *Rack) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Rack) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Rack) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.


### GetRackGroup

`func (o *Rack) GetRackGroup() BulkWritableRackRequestRackGroup`

GetRackGroup returns the RackGroup field if non-nil, zero value otherwise.

### GetRackGroupOk

`func (o *Rack) GetRackGroupOk() (*BulkWritableRackRequestRackGroup, bool)`

GetRackGroupOk returns a tuple with the RackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackGroup

`func (o *Rack) SetRackGroup(v BulkWritableRackRequestRackGroup)`

SetRackGroup sets RackGroup field to given value.

### HasRackGroup

`func (o *Rack) HasRackGroup() bool`

HasRackGroup returns a boolean if a field has been set.

### SetRackGroupNil

`func (o *Rack) SetRackGroupNil(b bool)`

 SetRackGroupNil sets the value for RackGroup to be an explicit nil

### UnsetRackGroup
`func (o *Rack) UnsetRackGroup()`

UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
### GetTenant

`func (o *Rack) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Rack) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Rack) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Rack) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Rack) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Rack) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCreated

`func (o *Rack) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Rack) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Rack) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Rack) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Rack) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Rack) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Rack) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Rack) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Rack) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Rack) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *Rack) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Rack) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Rack) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *Rack) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Rack) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Rack) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Rack) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *Rack) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Rack) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Rack) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Rack) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


