# PatchedBulkWritableSoftwareVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** | Optional alternative label for this version | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**EndOfSupportDate** | Pointer to **NullableString** |  | [optional] 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**LongTermSupport** | Pointer to **bool** | Is a Long Term Support version | [optional] 
**PreRelease** | Pointer to **bool** | Is a Pre-Release version | [optional] 
**Platform** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableSoftwareVersionRequest

`func NewPatchedBulkWritableSoftwareVersionRequest(id string, ) *PatchedBulkWritableSoftwareVersionRequest`

NewPatchedBulkWritableSoftwareVersionRequest instantiates a new PatchedBulkWritableSoftwareVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableSoftwareVersionRequestWithDefaults

`func NewPatchedBulkWritableSoftwareVersionRequestWithDefaults() *PatchedBulkWritableSoftwareVersionRequest`

NewPatchedBulkWritableSoftwareVersionRequestWithDefaults instantiates a new PatchedBulkWritableSoftwareVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAlias

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetReleaseDate

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *PatchedBulkWritableSoftwareVersionRequest) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetEndOfSupportDate

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetEndOfSupportDate() string`

GetEndOfSupportDate returns the EndOfSupportDate field if non-nil, zero value otherwise.

### GetEndOfSupportDateOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetEndOfSupportDateOk() (*string, bool)`

GetEndOfSupportDateOk returns a tuple with the EndOfSupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfSupportDate

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetEndOfSupportDate(v string)`

SetEndOfSupportDate sets EndOfSupportDate field to given value.

### HasEndOfSupportDate

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasEndOfSupportDate() bool`

HasEndOfSupportDate returns a boolean if a field has been set.

### SetEndOfSupportDateNil

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetEndOfSupportDateNil(b bool)`

 SetEndOfSupportDateNil sets the value for EndOfSupportDate to be an explicit nil

### UnsetEndOfSupportDate
`func (o *PatchedBulkWritableSoftwareVersionRequest) UnsetEndOfSupportDate()`

UnsetEndOfSupportDate ensures that no value is present for EndOfSupportDate, not even an explicit nil
### GetDocumentationUrl

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetLongTermSupport

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetLongTermSupport() bool`

GetLongTermSupport returns the LongTermSupport field if non-nil, zero value otherwise.

### GetLongTermSupportOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetLongTermSupportOk() (*bool, bool)`

GetLongTermSupportOk returns a tuple with the LongTermSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermSupport

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetLongTermSupport(v bool)`

SetLongTermSupport sets LongTermSupport field to given value.

### HasLongTermSupport

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasLongTermSupport() bool`

HasLongTermSupport returns a boolean if a field has been set.

### GetPreRelease

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetPreRelease() bool`

GetPreRelease returns the PreRelease field if non-nil, zero value otherwise.

### GetPreReleaseOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetPreReleaseOk() (*bool, bool)`

GetPreReleaseOk returns a tuple with the PreRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRelease

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetPreRelease(v bool)`

SetPreRelease sets PreRelease field to given value.

### HasPreRelease

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasPreRelease() bool`

HasPreRelease returns a boolean if a field has been set.

### GetPlatform

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetPlatform() BulkWritableCableRequestStatus`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetPlatformOk() (*BulkWritableCableRequestStatus, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetPlatform(v BulkWritableCableRequestStatus)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableSoftwareVersionRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableSoftwareVersionRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableSoftwareVersionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


