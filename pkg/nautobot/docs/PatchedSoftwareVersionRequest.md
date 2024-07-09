# PatchedSoftwareVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** | Optional alternative label for this version | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**EndOfSupportDate** | Pointer to **NullableString** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**LongTermSupport** | Pointer to **bool** | Is a Long Term Support version | [optional] 
**PreRelease** | Pointer to **bool** | Is a Pre-Release version | [optional] 
**Platform** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedSoftwareVersionRequest

`func NewPatchedSoftwareVersionRequest() *PatchedSoftwareVersionRequest`

NewPatchedSoftwareVersionRequest instantiates a new PatchedSoftwareVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSoftwareVersionRequestWithDefaults

`func NewPatchedSoftwareVersionRequestWithDefaults() *PatchedSoftwareVersionRequest`

NewPatchedSoftwareVersionRequestWithDefaults instantiates a new PatchedSoftwareVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *PatchedSoftwareVersionRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchedSoftwareVersionRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchedSoftwareVersionRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchedSoftwareVersionRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAlias

`func (o *PatchedSoftwareVersionRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PatchedSoftwareVersionRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PatchedSoftwareVersionRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PatchedSoftwareVersionRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetReleaseDate

`func (o *PatchedSoftwareVersionRequest) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *PatchedSoftwareVersionRequest) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *PatchedSoftwareVersionRequest) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *PatchedSoftwareVersionRequest) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *PatchedSoftwareVersionRequest) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *PatchedSoftwareVersionRequest) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetEndOfSupportDate

`func (o *PatchedSoftwareVersionRequest) GetEndOfSupportDate() string`

GetEndOfSupportDate returns the EndOfSupportDate field if non-nil, zero value otherwise.

### GetEndOfSupportDateOk

`func (o *PatchedSoftwareVersionRequest) GetEndOfSupportDateOk() (*string, bool)`

GetEndOfSupportDateOk returns a tuple with the EndOfSupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfSupportDate

`func (o *PatchedSoftwareVersionRequest) SetEndOfSupportDate(v string)`

SetEndOfSupportDate sets EndOfSupportDate field to given value.

### HasEndOfSupportDate

`func (o *PatchedSoftwareVersionRequest) HasEndOfSupportDate() bool`

HasEndOfSupportDate returns a boolean if a field has been set.

### SetEndOfSupportDateNil

`func (o *PatchedSoftwareVersionRequest) SetEndOfSupportDateNil(b bool)`

 SetEndOfSupportDateNil sets the value for EndOfSupportDate to be an explicit nil

### UnsetEndOfSupportDate
`func (o *PatchedSoftwareVersionRequest) UnsetEndOfSupportDate()`

UnsetEndOfSupportDate ensures that no value is present for EndOfSupportDate, not even an explicit nil
### GetDocumentationUrl

`func (o *PatchedSoftwareVersionRequest) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *PatchedSoftwareVersionRequest) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *PatchedSoftwareVersionRequest) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *PatchedSoftwareVersionRequest) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetLongTermSupport

`func (o *PatchedSoftwareVersionRequest) GetLongTermSupport() bool`

GetLongTermSupport returns the LongTermSupport field if non-nil, zero value otherwise.

### GetLongTermSupportOk

`func (o *PatchedSoftwareVersionRequest) GetLongTermSupportOk() (*bool, bool)`

GetLongTermSupportOk returns a tuple with the LongTermSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermSupport

`func (o *PatchedSoftwareVersionRequest) SetLongTermSupport(v bool)`

SetLongTermSupport sets LongTermSupport field to given value.

### HasLongTermSupport

`func (o *PatchedSoftwareVersionRequest) HasLongTermSupport() bool`

HasLongTermSupport returns a boolean if a field has been set.

### GetPreRelease

`func (o *PatchedSoftwareVersionRequest) GetPreRelease() bool`

GetPreRelease returns the PreRelease field if non-nil, zero value otherwise.

### GetPreReleaseOk

`func (o *PatchedSoftwareVersionRequest) GetPreReleaseOk() (*bool, bool)`

GetPreReleaseOk returns a tuple with the PreRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRelease

`func (o *PatchedSoftwareVersionRequest) SetPreRelease(v bool)`

SetPreRelease sets PreRelease field to given value.

### HasPreRelease

`func (o *PatchedSoftwareVersionRequest) HasPreRelease() bool`

HasPreRelease returns a boolean if a field has been set.

### GetPlatform

`func (o *PatchedSoftwareVersionRequest) GetPlatform() BulkWritableCableRequestStatus`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedSoftwareVersionRequest) GetPlatformOk() (*BulkWritableCableRequestStatus, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedSoftwareVersionRequest) SetPlatform(v BulkWritableCableRequestStatus)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedSoftwareVersionRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedSoftwareVersionRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedSoftwareVersionRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedSoftwareVersionRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedSoftwareVersionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PatchedSoftwareVersionRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedSoftwareVersionRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedSoftwareVersionRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedSoftwareVersionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedSoftwareVersionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedSoftwareVersionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedSoftwareVersionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedSoftwareVersionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedSoftwareVersionRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedSoftwareVersionRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedSoftwareVersionRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedSoftwareVersionRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


