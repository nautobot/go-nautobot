# GitRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvidedContents** | Pointer to [**[]BulkWritableGitRepositoryRequestProvidedContentsInner**](BulkWritableGitRepositoryRequestProvidedContentsInner.md) |  | [optional] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] 
**RemoteUrl** | **string** | Only HTTP and HTTPS URLs are presently supported | 
**Branch** | Pointer to **string** |  | [optional] 
**CurrentHead** | Pointer to **string** | Commit hash of the most recent fetch from the selected branch. Used for syncing between workers. | [optional] 
**SecretsGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewGitRepositoryRequest

`func NewGitRepositoryRequest(name string, remoteUrl string, ) *GitRepositoryRequest`

NewGitRepositoryRequest instantiates a new GitRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRepositoryRequestWithDefaults

`func NewGitRepositoryRequestWithDefaults() *GitRepositoryRequest`

NewGitRepositoryRequestWithDefaults instantiates a new GitRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvidedContents

`func (o *GitRepositoryRequest) GetProvidedContents() []BulkWritableGitRepositoryRequestProvidedContentsInner`

GetProvidedContents returns the ProvidedContents field if non-nil, zero value otherwise.

### GetProvidedContentsOk

`func (o *GitRepositoryRequest) GetProvidedContentsOk() (*[]BulkWritableGitRepositoryRequestProvidedContentsInner, bool)`

GetProvidedContentsOk returns a tuple with the ProvidedContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidedContents

`func (o *GitRepositoryRequest) SetProvidedContents(v []BulkWritableGitRepositoryRequestProvidedContentsInner)`

SetProvidedContents sets ProvidedContents field to given value.

### HasProvidedContents

`func (o *GitRepositoryRequest) HasProvidedContents() bool`

HasProvidedContents returns a boolean if a field has been set.

### GetName

`func (o *GitRepositoryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitRepositoryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitRepositoryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *GitRepositoryRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GitRepositoryRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GitRepositoryRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *GitRepositoryRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetRemoteUrl

`func (o *GitRepositoryRequest) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *GitRepositoryRequest) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *GitRepositoryRequest) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.


### GetBranch

`func (o *GitRepositoryRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitRepositoryRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitRepositoryRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitRepositoryRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCurrentHead

`func (o *GitRepositoryRequest) GetCurrentHead() string`

GetCurrentHead returns the CurrentHead field if non-nil, zero value otherwise.

### GetCurrentHeadOk

`func (o *GitRepositoryRequest) GetCurrentHeadOk() (*string, bool)`

GetCurrentHeadOk returns a tuple with the CurrentHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentHead

`func (o *GitRepositoryRequest) SetCurrentHead(v string)`

SetCurrentHead sets CurrentHead field to given value.

### HasCurrentHead

`func (o *GitRepositoryRequest) HasCurrentHead() bool`

HasCurrentHead returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *GitRepositoryRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *GitRepositoryRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *GitRepositoryRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *GitRepositoryRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *GitRepositoryRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *GitRepositoryRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCustomFields

`func (o *GitRepositoryRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GitRepositoryRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GitRepositoryRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GitRepositoryRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *GitRepositoryRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GitRepositoryRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GitRepositoryRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GitRepositoryRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


