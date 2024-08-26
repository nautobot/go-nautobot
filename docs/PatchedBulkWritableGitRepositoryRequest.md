# PatchedBulkWritableGitRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProvidedContents** | Pointer to [**[]BulkWritableGitRepositoryRequestProvidedContentsInner**](BulkWritableGitRepositoryRequestProvidedContentsInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] 
**RemoteUrl** | Pointer to **string** | Only HTTP and HTTPS URLs are presently supported | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**CurrentHead** | Pointer to **string** | Commit hash of the most recent fetch from the selected branch. Used for syncing between workers. | [optional] 
**SecretsGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableGitRepositoryRequest

`func NewPatchedBulkWritableGitRepositoryRequest(id string, ) *PatchedBulkWritableGitRepositoryRequest`

NewPatchedBulkWritableGitRepositoryRequest instantiates a new PatchedBulkWritableGitRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableGitRepositoryRequestWithDefaults

`func NewPatchedBulkWritableGitRepositoryRequestWithDefaults() *PatchedBulkWritableGitRepositoryRequest`

NewPatchedBulkWritableGitRepositoryRequestWithDefaults instantiates a new PatchedBulkWritableGitRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableGitRepositoryRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableGitRepositoryRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableGitRepositoryRequest) SetId(v string)`

SetId sets Id field to given value.


### GetProvidedContents

`func (o *PatchedBulkWritableGitRepositoryRequest) GetProvidedContents() []BulkWritableGitRepositoryRequestProvidedContentsInner`

GetProvidedContents returns the ProvidedContents field if non-nil, zero value otherwise.

### GetProvidedContentsOk

`func (o *PatchedBulkWritableGitRepositoryRequest) GetProvidedContentsOk() (*[]BulkWritableGitRepositoryRequestProvidedContentsInner, bool)`

GetProvidedContentsOk returns a tuple with the ProvidedContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidedContents

`func (o *PatchedBulkWritableGitRepositoryRequest) SetProvidedContents(v []BulkWritableGitRepositoryRequestProvidedContentsInner)`

SetProvidedContents sets ProvidedContents field to given value.

### HasProvidedContents

`func (o *PatchedBulkWritableGitRepositoryRequest) HasProvidedContents() bool`

HasProvidedContents returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableGitRepositoryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableGitRepositoryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableGitRepositoryRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableGitRepositoryRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedBulkWritableGitRepositoryRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedBulkWritableGitRepositoryRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedBulkWritableGitRepositoryRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedBulkWritableGitRepositoryRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetRemoteUrl

`func (o *PatchedBulkWritableGitRepositoryRequest) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *PatchedBulkWritableGitRepositoryRequest) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *PatchedBulkWritableGitRepositoryRequest) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.

### HasRemoteUrl

`func (o *PatchedBulkWritableGitRepositoryRequest) HasRemoteUrl() bool`

HasRemoteUrl returns a boolean if a field has been set.

### GetBranch

`func (o *PatchedBulkWritableGitRepositoryRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PatchedBulkWritableGitRepositoryRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PatchedBulkWritableGitRepositoryRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PatchedBulkWritableGitRepositoryRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCurrentHead

`func (o *PatchedBulkWritableGitRepositoryRequest) GetCurrentHead() string`

GetCurrentHead returns the CurrentHead field if non-nil, zero value otherwise.

### GetCurrentHeadOk

`func (o *PatchedBulkWritableGitRepositoryRequest) GetCurrentHeadOk() (*string, bool)`

GetCurrentHeadOk returns a tuple with the CurrentHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentHead

`func (o *PatchedBulkWritableGitRepositoryRequest) SetCurrentHead(v string)`

SetCurrentHead sets CurrentHead field to given value.

### HasCurrentHead

`func (o *PatchedBulkWritableGitRepositoryRequest) HasCurrentHead() bool`

HasCurrentHead returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *PatchedBulkWritableGitRepositoryRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedBulkWritableGitRepositoryRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedBulkWritableGitRepositoryRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedBulkWritableGitRepositoryRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedBulkWritableGitRepositoryRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedBulkWritableGitRepositoryRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableGitRepositoryRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableGitRepositoryRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableGitRepositoryRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableGitRepositoryRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableGitRepositoryRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableGitRepositoryRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableGitRepositoryRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableGitRepositoryRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


