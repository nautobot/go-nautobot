# SoftwareVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Version** | **string** |  | 
**Alias** | Pointer to **string** | Optional alternative label for this version | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**EndOfSupportDate** | Pointer to **NullableString** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**LongTermSupport** | Pointer to **bool** | Is a Long Term Support version | [optional] 
**PreRelease** | Pointer to **bool** | Is a Pre-Release version | [optional] 
**Platform** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSoftwareVersion

`func NewSoftwareVersion(id string, objectType string, display string, url string, naturalSlug string, version string, platform BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *SoftwareVersion`

NewSoftwareVersion instantiates a new SoftwareVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareVersionWithDefaults

`func NewSoftwareVersionWithDefaults() *SoftwareVersion`

NewSoftwareVersionWithDefaults instantiates a new SoftwareVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SoftwareVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SoftwareVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SoftwareVersion) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *SoftwareVersion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SoftwareVersion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SoftwareVersion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *SoftwareVersion) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SoftwareVersion) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SoftwareVersion) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *SoftwareVersion) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SoftwareVersion) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SoftwareVersion) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *SoftwareVersion) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *SoftwareVersion) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *SoftwareVersion) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetVersion

`func (o *SoftwareVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwareVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwareVersion) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetAlias

`func (o *SoftwareVersion) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *SoftwareVersion) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *SoftwareVersion) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *SoftwareVersion) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetReleaseDate

`func (o *SoftwareVersion) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SoftwareVersion) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SoftwareVersion) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *SoftwareVersion) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *SoftwareVersion) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *SoftwareVersion) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetEndOfSupportDate

`func (o *SoftwareVersion) GetEndOfSupportDate() string`

GetEndOfSupportDate returns the EndOfSupportDate field if non-nil, zero value otherwise.

### GetEndOfSupportDateOk

`func (o *SoftwareVersion) GetEndOfSupportDateOk() (*string, bool)`

GetEndOfSupportDateOk returns a tuple with the EndOfSupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfSupportDate

`func (o *SoftwareVersion) SetEndOfSupportDate(v string)`

SetEndOfSupportDate sets EndOfSupportDate field to given value.

### HasEndOfSupportDate

`func (o *SoftwareVersion) HasEndOfSupportDate() bool`

HasEndOfSupportDate returns a boolean if a field has been set.

### SetEndOfSupportDateNil

`func (o *SoftwareVersion) SetEndOfSupportDateNil(b bool)`

 SetEndOfSupportDateNil sets the value for EndOfSupportDate to be an explicit nil

### UnsetEndOfSupportDate
`func (o *SoftwareVersion) UnsetEndOfSupportDate()`

UnsetEndOfSupportDate ensures that no value is present for EndOfSupportDate, not even an explicit nil
### GetDocumentationUrl

`func (o *SoftwareVersion) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *SoftwareVersion) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *SoftwareVersion) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *SoftwareVersion) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetLongTermSupport

`func (o *SoftwareVersion) GetLongTermSupport() bool`

GetLongTermSupport returns the LongTermSupport field if non-nil, zero value otherwise.

### GetLongTermSupportOk

`func (o *SoftwareVersion) GetLongTermSupportOk() (*bool, bool)`

GetLongTermSupportOk returns a tuple with the LongTermSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermSupport

`func (o *SoftwareVersion) SetLongTermSupport(v bool)`

SetLongTermSupport sets LongTermSupport field to given value.

### HasLongTermSupport

`func (o *SoftwareVersion) HasLongTermSupport() bool`

HasLongTermSupport returns a boolean if a field has been set.

### GetPreRelease

`func (o *SoftwareVersion) GetPreRelease() bool`

GetPreRelease returns the PreRelease field if non-nil, zero value otherwise.

### GetPreReleaseOk

`func (o *SoftwareVersion) GetPreReleaseOk() (*bool, bool)`

GetPreReleaseOk returns a tuple with the PreRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRelease

`func (o *SoftwareVersion) SetPreRelease(v bool)`

SetPreRelease sets PreRelease field to given value.

### HasPreRelease

`func (o *SoftwareVersion) HasPreRelease() bool`

HasPreRelease returns a boolean if a field has been set.

### GetPlatform

`func (o *SoftwareVersion) GetPlatform() BulkWritableCableRequestStatus`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SoftwareVersion) GetPlatformOk() (*BulkWritableCableRequestStatus, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SoftwareVersion) SetPlatform(v BulkWritableCableRequestStatus)`

SetPlatform sets Platform field to given value.


### GetStatus

`func (o *SoftwareVersion) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SoftwareVersion) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SoftwareVersion) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *SoftwareVersion) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SoftwareVersion) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SoftwareVersion) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *SoftwareVersion) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SoftwareVersion) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *SoftwareVersion) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SoftwareVersion) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SoftwareVersion) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *SoftwareVersion) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *SoftwareVersion) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *SoftwareVersion) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SoftwareVersion) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SoftwareVersion) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SoftwareVersion) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *SoftwareVersion) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *SoftwareVersion) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *SoftwareVersion) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *SoftwareVersion) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SoftwareVersion) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SoftwareVersion) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SoftwareVersion) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


