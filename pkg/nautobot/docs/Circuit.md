# Circuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Cid** | **string** |  | 
**InstallDate** | Pointer to **NullableString** |  | [optional] 
**CommitRate** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Provider** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CircuitType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CircuitTerminationA** | [**NullableCircuitCircuitTerminationA**](CircuitCircuitTerminationA.md) |  | 
**CircuitTerminationZ** | [**NullableCircuitCircuitTerminationA**](CircuitCircuitTerminationA.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCircuit

`func NewCircuit(id string, objectType string, display string, url string, naturalSlug string, cid string, status BulkWritableCableRequestStatus, provider BulkWritableCableRequestStatus, circuitType BulkWritableCableRequestStatus, circuitTerminationA NullableCircuitCircuitTerminationA, circuitTerminationZ NullableCircuitCircuitTerminationA, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Circuit`

NewCircuit instantiates a new Circuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitWithDefaults

`func NewCircuitWithDefaults() *Circuit`

NewCircuitWithDefaults instantiates a new Circuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Circuit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Circuit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Circuit) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Circuit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Circuit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Circuit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Circuit) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Circuit) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Circuit) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Circuit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Circuit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Circuit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Circuit) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Circuit) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Circuit) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCid

`func (o *Circuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Circuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Circuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetInstallDate

`func (o *Circuit) GetInstallDate() string`

GetInstallDate returns the InstallDate field if non-nil, zero value otherwise.

### GetInstallDateOk

`func (o *Circuit) GetInstallDateOk() (*string, bool)`

GetInstallDateOk returns a tuple with the InstallDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDate

`func (o *Circuit) SetInstallDate(v string)`

SetInstallDate sets InstallDate field to given value.

### HasInstallDate

`func (o *Circuit) HasInstallDate() bool`

HasInstallDate returns a boolean if a field has been set.

### SetInstallDateNil

`func (o *Circuit) SetInstallDateNil(b bool)`

 SetInstallDateNil sets the value for InstallDate to be an explicit nil

### UnsetInstallDate
`func (o *Circuit) UnsetInstallDate()`

UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
### GetCommitRate

`func (o *Circuit) GetCommitRate() int32`

GetCommitRate returns the CommitRate field if non-nil, zero value otherwise.

### GetCommitRateOk

`func (o *Circuit) GetCommitRateOk() (*int32, bool)`

GetCommitRateOk returns a tuple with the CommitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitRate

`func (o *Circuit) SetCommitRate(v int32)`

SetCommitRate sets CommitRate field to given value.

### HasCommitRate

`func (o *Circuit) HasCommitRate() bool`

HasCommitRate returns a boolean if a field has been set.

### SetCommitRateNil

`func (o *Circuit) SetCommitRateNil(b bool)`

 SetCommitRateNil sets the value for CommitRate to be an explicit nil

### UnsetCommitRate
`func (o *Circuit) UnsetCommitRate()`

UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
### GetDescription

`func (o *Circuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Circuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Circuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Circuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *Circuit) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Circuit) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Circuit) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Circuit) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetStatus

`func (o *Circuit) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Circuit) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Circuit) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetProvider

`func (o *Circuit) GetProvider() BulkWritableCableRequestStatus`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Circuit) GetProviderOk() (*BulkWritableCableRequestStatus, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Circuit) SetProvider(v BulkWritableCableRequestStatus)`

SetProvider sets Provider field to given value.


### GetCircuitType

`func (o *Circuit) GetCircuitType() BulkWritableCableRequestStatus`

GetCircuitType returns the CircuitType field if non-nil, zero value otherwise.

### GetCircuitTypeOk

`func (o *Circuit) GetCircuitTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetCircuitTypeOk returns a tuple with the CircuitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitType

`func (o *Circuit) SetCircuitType(v BulkWritableCableRequestStatus)`

SetCircuitType sets CircuitType field to given value.


### GetTenant

`func (o *Circuit) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Circuit) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Circuit) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Circuit) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Circuit) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Circuit) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCircuitTerminationA

`func (o *Circuit) GetCircuitTerminationA() CircuitCircuitTerminationA`

GetCircuitTerminationA returns the CircuitTerminationA field if non-nil, zero value otherwise.

### GetCircuitTerminationAOk

`func (o *Circuit) GetCircuitTerminationAOk() (*CircuitCircuitTerminationA, bool)`

GetCircuitTerminationAOk returns a tuple with the CircuitTerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitTerminationA

`func (o *Circuit) SetCircuitTerminationA(v CircuitCircuitTerminationA)`

SetCircuitTerminationA sets CircuitTerminationA field to given value.


### SetCircuitTerminationANil

`func (o *Circuit) SetCircuitTerminationANil(b bool)`

 SetCircuitTerminationANil sets the value for CircuitTerminationA to be an explicit nil

### UnsetCircuitTerminationA
`func (o *Circuit) UnsetCircuitTerminationA()`

UnsetCircuitTerminationA ensures that no value is present for CircuitTerminationA, not even an explicit nil
### GetCircuitTerminationZ

`func (o *Circuit) GetCircuitTerminationZ() CircuitCircuitTerminationA`

GetCircuitTerminationZ returns the CircuitTerminationZ field if non-nil, zero value otherwise.

### GetCircuitTerminationZOk

`func (o *Circuit) GetCircuitTerminationZOk() (*CircuitCircuitTerminationA, bool)`

GetCircuitTerminationZOk returns a tuple with the CircuitTerminationZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitTerminationZ

`func (o *Circuit) SetCircuitTerminationZ(v CircuitCircuitTerminationA)`

SetCircuitTerminationZ sets CircuitTerminationZ field to given value.


### SetCircuitTerminationZNil

`func (o *Circuit) SetCircuitTerminationZNil(b bool)`

 SetCircuitTerminationZNil sets the value for CircuitTerminationZ to be an explicit nil

### UnsetCircuitTerminationZ
`func (o *Circuit) UnsetCircuitTerminationZ()`

UnsetCircuitTerminationZ ensures that no value is present for CircuitTerminationZ, not even an explicit nil
### GetCreated

`func (o *Circuit) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Circuit) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Circuit) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Circuit) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Circuit) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Circuit) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Circuit) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Circuit) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Circuit) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Circuit) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *Circuit) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Circuit) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Circuit) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Circuit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *Circuit) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Circuit) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Circuit) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *Circuit) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Circuit) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Circuit) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Circuit) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


