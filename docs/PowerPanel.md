# PowerPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**PanelType** | Pointer to [**PowerPanelPanelType**](PowerPanelPanelType.md) |  | [optional] 
**PowerPath** | Pointer to [**PowerFeedPowerPath**](PowerFeedPowerPath.md) |  | [optional] 
**PowerFeedCount** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**BreakerPositionCount** | Pointer to **NullableInt32** | Total number of breaker positions in the panel (e.g., 42) | [optional] 
**Location** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**RackGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPowerPanel

`func NewPowerPanel(objectType string, display string, url string, naturalSlug string, name string, location BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *PowerPanel`

NewPowerPanel instantiates a new PowerPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPanelWithDefaults

`func NewPowerPanelWithDefaults() *PowerPanel`

NewPowerPanelWithDefaults instantiates a new PowerPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerPanel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerPanel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerPanel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PowerPanel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *PowerPanel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PowerPanel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PowerPanel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *PowerPanel) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PowerPanel) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PowerPanel) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *PowerPanel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PowerPanel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PowerPanel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *PowerPanel) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *PowerPanel) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *PowerPanel) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetPanelType

`func (o *PowerPanel) GetPanelType() PowerPanelPanelType`

GetPanelType returns the PanelType field if non-nil, zero value otherwise.

### GetPanelTypeOk

`func (o *PowerPanel) GetPanelTypeOk() (*PowerPanelPanelType, bool)`

GetPanelTypeOk returns a tuple with the PanelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelType

`func (o *PowerPanel) SetPanelType(v PowerPanelPanelType)`

SetPanelType sets PanelType field to given value.

### HasPanelType

`func (o *PowerPanel) HasPanelType() bool`

HasPanelType returns a boolean if a field has been set.

### GetPowerPath

`func (o *PowerPanel) GetPowerPath() PowerFeedPowerPath`

GetPowerPath returns the PowerPath field if non-nil, zero value otherwise.

### GetPowerPathOk

`func (o *PowerPanel) GetPowerPathOk() (*PowerFeedPowerPath, bool)`

GetPowerPathOk returns a tuple with the PowerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPath

`func (o *PowerPanel) SetPowerPath(v PowerFeedPowerPath)`

SetPowerPath sets PowerPath field to given value.

### HasPowerPath

`func (o *PowerPanel) HasPowerPath() bool`

HasPowerPath returns a boolean if a field has been set.

### GetPowerFeedCount

`func (o *PowerPanel) GetPowerFeedCount() int32`

GetPowerFeedCount returns the PowerFeedCount field if non-nil, zero value otherwise.

### GetPowerFeedCountOk

`func (o *PowerPanel) GetPowerFeedCountOk() (*int32, bool)`

GetPowerFeedCountOk returns a tuple with the PowerFeedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFeedCount

`func (o *PowerPanel) SetPowerFeedCount(v int32)`

SetPowerFeedCount sets PowerFeedCount field to given value.

### HasPowerFeedCount

`func (o *PowerPanel) HasPowerFeedCount() bool`

HasPowerFeedCount returns a boolean if a field has been set.

### GetName

`func (o *PowerPanel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerPanel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerPanel) SetName(v string)`

SetName sets Name field to given value.


### GetBreakerPositionCount

`func (o *PowerPanel) GetBreakerPositionCount() int32`

GetBreakerPositionCount returns the BreakerPositionCount field if non-nil, zero value otherwise.

### GetBreakerPositionCountOk

`func (o *PowerPanel) GetBreakerPositionCountOk() (*int32, bool)`

GetBreakerPositionCountOk returns a tuple with the BreakerPositionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakerPositionCount

`func (o *PowerPanel) SetBreakerPositionCount(v int32)`

SetBreakerPositionCount sets BreakerPositionCount field to given value.

### HasBreakerPositionCount

`func (o *PowerPanel) HasBreakerPositionCount() bool`

HasBreakerPositionCount returns a boolean if a field has been set.

### SetBreakerPositionCountNil

`func (o *PowerPanel) SetBreakerPositionCountNil(b bool)`

 SetBreakerPositionCountNil sets the value for BreakerPositionCount to be an explicit nil

### UnsetBreakerPositionCount
`func (o *PowerPanel) UnsetBreakerPositionCount()`

UnsetBreakerPositionCount ensures that no value is present for BreakerPositionCount, not even an explicit nil
### GetLocation

`func (o *PowerPanel) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PowerPanel) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PowerPanel) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.


### GetRackGroup

`func (o *PowerPanel) GetRackGroup() BulkWritableCircuitRequestTenant`

GetRackGroup returns the RackGroup field if non-nil, zero value otherwise.

### GetRackGroupOk

`func (o *PowerPanel) GetRackGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRackGroupOk returns a tuple with the RackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackGroup

`func (o *PowerPanel) SetRackGroup(v BulkWritableCircuitRequestTenant)`

SetRackGroup sets RackGroup field to given value.

### HasRackGroup

`func (o *PowerPanel) HasRackGroup() bool`

HasRackGroup returns a boolean if a field has been set.

### SetRackGroupNil

`func (o *PowerPanel) SetRackGroupNil(b bool)`

 SetRackGroupNil sets the value for RackGroup to be an explicit nil

### UnsetRackGroup
`func (o *PowerPanel) UnsetRackGroup()`

UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
### GetCreated

`func (o *PowerPanel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PowerPanel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PowerPanel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *PowerPanel) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *PowerPanel) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *PowerPanel) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PowerPanel) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PowerPanel) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *PowerPanel) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *PowerPanel) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *PowerPanel) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *PowerPanel) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *PowerPanel) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *PowerPanel) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PowerPanel) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PowerPanel) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PowerPanel) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *PowerPanel) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PowerPanel) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PowerPanel) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PowerPanel) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


