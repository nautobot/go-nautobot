# BulkWritableGitRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProvidedContents** | Pointer to [**[]BulkWritableGitRepositoryRequestProvidedContentsInner**](BulkWritableGitRepositoryRequestProvidedContentsInner.md) |  | [optional] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] 
**RemoteUrl** | **string** | Only HTTP and HTTPS URLs are presently supported | 
**Branch** | Pointer to **string** | Branch, tag, or commit | [optional] 
**SecretsGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableGitRepositoryRequest

`func NewBulkWritableGitRepositoryRequest(id string, name string, remoteUrl string, ) *BulkWritableGitRepositoryRequest`

NewBulkWritableGitRepositoryRequest instantiates a new BulkWritableGitRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableGitRepositoryRequestWithDefaults

`func NewBulkWritableGitRepositoryRequestWithDefaults() *BulkWritableGitRepositoryRequest`

NewBulkWritableGitRepositoryRequestWithDefaults instantiates a new BulkWritableGitRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableGitRepositoryRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableGitRepositoryRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableGitRepositoryRequest) SetId(v string)`

SetId sets Id field to given value.


### GetProvidedContents

`func (o *BulkWritableGitRepositoryRequest) GetProvidedContents() []BulkWritableGitRepositoryRequestProvidedContentsInner`

GetProvidedContents returns the ProvidedContents field if non-nil, zero value otherwise.

### GetProvidedContentsOk

`func (o *BulkWritableGitRepositoryRequest) GetProvidedContentsOk() (*[]BulkWritableGitRepositoryRequestProvidedContentsInner, bool)`

GetProvidedContentsOk returns a tuple with the ProvidedContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidedContents

`func (o *BulkWritableGitRepositoryRequest) SetProvidedContents(v []BulkWritableGitRepositoryRequestProvidedContentsInner)`

SetProvidedContents sets ProvidedContents field to given value.

### HasProvidedContents

`func (o *BulkWritableGitRepositoryRequest) HasProvidedContents() bool`

HasProvidedContents returns a boolean if a field has been set.

### GetName

`func (o *BulkWritableGitRepositoryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableGitRepositoryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableGitRepositoryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BulkWritableGitRepositoryRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BulkWritableGitRepositoryRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BulkWritableGitRepositoryRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BulkWritableGitRepositoryRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetRemoteUrl

`func (o *BulkWritableGitRepositoryRequest) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *BulkWritableGitRepositoryRequest) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *BulkWritableGitRepositoryRequest) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.


### GetBranch

`func (o *BulkWritableGitRepositoryRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *BulkWritableGitRepositoryRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *BulkWritableGitRepositoryRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *BulkWritableGitRepositoryRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *BulkWritableGitRepositoryRequest) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *BulkWritableGitRepositoryRequest) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *BulkWritableGitRepositoryRequest) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *BulkWritableGitRepositoryRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *BulkWritableGitRepositoryRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *BulkWritableGitRepositoryRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableGitRepositoryRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableGitRepositoryRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableGitRepositoryRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableGitRepositoryRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableGitRepositoryRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableGitRepositoryRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableGitRepositoryRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableGitRepositoryRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableGitRepositoryRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableGitRepositoryRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableGitRepositoryRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableGitRepositoryRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


