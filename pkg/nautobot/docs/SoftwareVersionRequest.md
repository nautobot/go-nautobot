# SoftwareVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Alias** | Pointer to **string** | Optional alternative label for this version | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**EndOfSupportDate** | Pointer to **NullableString** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**LongTermSupport** | Pointer to **bool** | Is a Long Term Support version | [optional] 
**PreRelease** | Pointer to **bool** | Is a Pre-Release version | [optional] 
**Platform** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewSoftwareVersionRequest

`func NewSoftwareVersionRequest(version string, platform BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, ) *SoftwareVersionRequest`

NewSoftwareVersionRequest instantiates a new SoftwareVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareVersionRequestWithDefaults

`func NewSoftwareVersionRequestWithDefaults() *SoftwareVersionRequest`

NewSoftwareVersionRequestWithDefaults instantiates a new SoftwareVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SoftwareVersionRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SoftwareVersionRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SoftwareVersionRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetAlias

`func (o *SoftwareVersionRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *SoftwareVersionRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *SoftwareVersionRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *SoftwareVersionRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetReleaseDate

`func (o *SoftwareVersionRequest) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SoftwareVersionRequest) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SoftwareVersionRequest) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *SoftwareVersionRequest) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *SoftwareVersionRequest) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *SoftwareVersionRequest) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetEndOfSupportDate

`func (o *SoftwareVersionRequest) GetEndOfSupportDate() string`

GetEndOfSupportDate returns the EndOfSupportDate field if non-nil, zero value otherwise.

### GetEndOfSupportDateOk

`func (o *SoftwareVersionRequest) GetEndOfSupportDateOk() (*string, bool)`

GetEndOfSupportDateOk returns a tuple with the EndOfSupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfSupportDate

`func (o *SoftwareVersionRequest) SetEndOfSupportDate(v string)`

SetEndOfSupportDate sets EndOfSupportDate field to given value.

### HasEndOfSupportDate

`func (o *SoftwareVersionRequest) HasEndOfSupportDate() bool`

HasEndOfSupportDate returns a boolean if a field has been set.

### SetEndOfSupportDateNil

`func (o *SoftwareVersionRequest) SetEndOfSupportDateNil(b bool)`

 SetEndOfSupportDateNil sets the value for EndOfSupportDate to be an explicit nil

### UnsetEndOfSupportDate
`func (o *SoftwareVersionRequest) UnsetEndOfSupportDate()`

UnsetEndOfSupportDate ensures that no value is present for EndOfSupportDate, not even an explicit nil
### GetDocumentationUrl

`func (o *SoftwareVersionRequest) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *SoftwareVersionRequest) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *SoftwareVersionRequest) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *SoftwareVersionRequest) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetLongTermSupport

`func (o *SoftwareVersionRequest) GetLongTermSupport() bool`

GetLongTermSupport returns the LongTermSupport field if non-nil, zero value otherwise.

### GetLongTermSupportOk

`func (o *SoftwareVersionRequest) GetLongTermSupportOk() (*bool, bool)`

GetLongTermSupportOk returns a tuple with the LongTermSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermSupport

`func (o *SoftwareVersionRequest) SetLongTermSupport(v bool)`

SetLongTermSupport sets LongTermSupport field to given value.

### HasLongTermSupport

`func (o *SoftwareVersionRequest) HasLongTermSupport() bool`

HasLongTermSupport returns a boolean if a field has been set.

### GetPreRelease

`func (o *SoftwareVersionRequest) GetPreRelease() bool`

GetPreRelease returns the PreRelease field if non-nil, zero value otherwise.

### GetPreReleaseOk

`func (o *SoftwareVersionRequest) GetPreReleaseOk() (*bool, bool)`

GetPreReleaseOk returns a tuple with the PreRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRelease

`func (o *SoftwareVersionRequest) SetPreRelease(v bool)`

SetPreRelease sets PreRelease field to given value.

### HasPreRelease

`func (o *SoftwareVersionRequest) HasPreRelease() bool`

HasPreRelease returns a boolean if a field has been set.

### GetPlatform

`func (o *SoftwareVersionRequest) GetPlatform() BulkWritableCableRequestStatus`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SoftwareVersionRequest) GetPlatformOk() (*BulkWritableCableRequestStatus, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SoftwareVersionRequest) SetPlatform(v BulkWritableCableRequestStatus)`

SetPlatform sets Platform field to given value.


### GetStatus

`func (o *SoftwareVersionRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SoftwareVersionRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SoftwareVersionRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *SoftwareVersionRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SoftwareVersionRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SoftwareVersionRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SoftwareVersionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *SoftwareVersionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SoftwareVersionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SoftwareVersionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SoftwareVersionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *SoftwareVersionRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SoftwareVersionRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SoftwareVersionRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SoftwareVersionRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


