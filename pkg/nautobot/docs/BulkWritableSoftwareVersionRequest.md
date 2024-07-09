# BulkWritableSoftwareVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewBulkWritableSoftwareVersionRequest

`func NewBulkWritableSoftwareVersionRequest(id string, version string, platform BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, ) *BulkWritableSoftwareVersionRequest`

NewBulkWritableSoftwareVersionRequest instantiates a new BulkWritableSoftwareVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableSoftwareVersionRequestWithDefaults

`func NewBulkWritableSoftwareVersionRequestWithDefaults() *BulkWritableSoftwareVersionRequest`

NewBulkWritableSoftwareVersionRequestWithDefaults instantiates a new BulkWritableSoftwareVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableSoftwareVersionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableSoftwareVersionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableSoftwareVersionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *BulkWritableSoftwareVersionRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BulkWritableSoftwareVersionRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BulkWritableSoftwareVersionRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetAlias

`func (o *BulkWritableSoftwareVersionRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *BulkWritableSoftwareVersionRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *BulkWritableSoftwareVersionRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *BulkWritableSoftwareVersionRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetReleaseDate

`func (o *BulkWritableSoftwareVersionRequest) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *BulkWritableSoftwareVersionRequest) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *BulkWritableSoftwareVersionRequest) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *BulkWritableSoftwareVersionRequest) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *BulkWritableSoftwareVersionRequest) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *BulkWritableSoftwareVersionRequest) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetEndOfSupportDate

`func (o *BulkWritableSoftwareVersionRequest) GetEndOfSupportDate() string`

GetEndOfSupportDate returns the EndOfSupportDate field if non-nil, zero value otherwise.

### GetEndOfSupportDateOk

`func (o *BulkWritableSoftwareVersionRequest) GetEndOfSupportDateOk() (*string, bool)`

GetEndOfSupportDateOk returns a tuple with the EndOfSupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfSupportDate

`func (o *BulkWritableSoftwareVersionRequest) SetEndOfSupportDate(v string)`

SetEndOfSupportDate sets EndOfSupportDate field to given value.

### HasEndOfSupportDate

`func (o *BulkWritableSoftwareVersionRequest) HasEndOfSupportDate() bool`

HasEndOfSupportDate returns a boolean if a field has been set.

### SetEndOfSupportDateNil

`func (o *BulkWritableSoftwareVersionRequest) SetEndOfSupportDateNil(b bool)`

 SetEndOfSupportDateNil sets the value for EndOfSupportDate to be an explicit nil

### UnsetEndOfSupportDate
`func (o *BulkWritableSoftwareVersionRequest) UnsetEndOfSupportDate()`

UnsetEndOfSupportDate ensures that no value is present for EndOfSupportDate, not even an explicit nil
### GetDocumentationUrl

`func (o *BulkWritableSoftwareVersionRequest) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *BulkWritableSoftwareVersionRequest) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *BulkWritableSoftwareVersionRequest) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *BulkWritableSoftwareVersionRequest) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetLongTermSupport

`func (o *BulkWritableSoftwareVersionRequest) GetLongTermSupport() bool`

GetLongTermSupport returns the LongTermSupport field if non-nil, zero value otherwise.

### GetLongTermSupportOk

`func (o *BulkWritableSoftwareVersionRequest) GetLongTermSupportOk() (*bool, bool)`

GetLongTermSupportOk returns a tuple with the LongTermSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermSupport

`func (o *BulkWritableSoftwareVersionRequest) SetLongTermSupport(v bool)`

SetLongTermSupport sets LongTermSupport field to given value.

### HasLongTermSupport

`func (o *BulkWritableSoftwareVersionRequest) HasLongTermSupport() bool`

HasLongTermSupport returns a boolean if a field has been set.

### GetPreRelease

`func (o *BulkWritableSoftwareVersionRequest) GetPreRelease() bool`

GetPreRelease returns the PreRelease field if non-nil, zero value otherwise.

### GetPreReleaseOk

`func (o *BulkWritableSoftwareVersionRequest) GetPreReleaseOk() (*bool, bool)`

GetPreReleaseOk returns a tuple with the PreRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRelease

`func (o *BulkWritableSoftwareVersionRequest) SetPreRelease(v bool)`

SetPreRelease sets PreRelease field to given value.

### HasPreRelease

`func (o *BulkWritableSoftwareVersionRequest) HasPreRelease() bool`

HasPreRelease returns a boolean if a field has been set.

### GetPlatform

`func (o *BulkWritableSoftwareVersionRequest) GetPlatform() BulkWritableCableRequestStatus`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *BulkWritableSoftwareVersionRequest) GetPlatformOk() (*BulkWritableCableRequestStatus, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *BulkWritableSoftwareVersionRequest) SetPlatform(v BulkWritableCableRequestStatus)`

SetPlatform sets Platform field to given value.


### GetStatus

`func (o *BulkWritableSoftwareVersionRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableSoftwareVersionRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableSoftwareVersionRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *BulkWritableSoftwareVersionRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableSoftwareVersionRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableSoftwareVersionRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableSoftwareVersionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableSoftwareVersionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableSoftwareVersionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableSoftwareVersionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableSoftwareVersionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableSoftwareVersionRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableSoftwareVersionRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableSoftwareVersionRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableSoftwareVersionRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


