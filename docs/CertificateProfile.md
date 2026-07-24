# CertificateProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Name** | **string** |  | 
**CertificateType** | Pointer to [**BulkWritableCertificateProfileRequestCertificateType**](BulkWritableCertificateProfileRequestCertificateType.md) |  | [optional] 
**CertificateFilePath** | Pointer to **string** |  | [optional] 
**ChainFilePath** | Pointer to **string** |  | [optional] 
**KeyFilePath** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **NullableTime** |  | [optional] 
**Cipher** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewCertificateProfile

`func NewCertificateProfile(objectType string, display string, url string, naturalSlug string, name string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *CertificateProfile`

NewCertificateProfile instantiates a new CertificateProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileWithDefaults

`func NewCertificateProfileWithDefaults() *CertificateProfile`

NewCertificateProfileWithDefaults instantiates a new CertificateProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *CertificateProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CertificateProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CertificateProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *CertificateProfile) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CertificateProfile) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CertificateProfile) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *CertificateProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CertificateProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CertificateProfile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *CertificateProfile) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *CertificateProfile) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *CertificateProfile) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetName

`func (o *CertificateProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfile) SetName(v string)`

SetName sets Name field to given value.


### GetCertificateType

`func (o *CertificateProfile) GetCertificateType() BulkWritableCertificateProfileRequestCertificateType`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *CertificateProfile) GetCertificateTypeOk() (*BulkWritableCertificateProfileRequestCertificateType, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *CertificateProfile) SetCertificateType(v BulkWritableCertificateProfileRequestCertificateType)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *CertificateProfile) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetCertificateFilePath

`func (o *CertificateProfile) GetCertificateFilePath() string`

GetCertificateFilePath returns the CertificateFilePath field if non-nil, zero value otherwise.

### GetCertificateFilePathOk

`func (o *CertificateProfile) GetCertificateFilePathOk() (*string, bool)`

GetCertificateFilePathOk returns a tuple with the CertificateFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFilePath

`func (o *CertificateProfile) SetCertificateFilePath(v string)`

SetCertificateFilePath sets CertificateFilePath field to given value.

### HasCertificateFilePath

`func (o *CertificateProfile) HasCertificateFilePath() bool`

HasCertificateFilePath returns a boolean if a field has been set.

### GetChainFilePath

`func (o *CertificateProfile) GetChainFilePath() string`

GetChainFilePath returns the ChainFilePath field if non-nil, zero value otherwise.

### GetChainFilePathOk

`func (o *CertificateProfile) GetChainFilePathOk() (*string, bool)`

GetChainFilePathOk returns a tuple with the ChainFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainFilePath

`func (o *CertificateProfile) SetChainFilePath(v string)`

SetChainFilePath sets ChainFilePath field to given value.

### HasChainFilePath

`func (o *CertificateProfile) HasChainFilePath() bool`

HasChainFilePath returns a boolean if a field has been set.

### GetKeyFilePath

`func (o *CertificateProfile) GetKeyFilePath() string`

GetKeyFilePath returns the KeyFilePath field if non-nil, zero value otherwise.

### GetKeyFilePathOk

`func (o *CertificateProfile) GetKeyFilePathOk() (*string, bool)`

GetKeyFilePathOk returns a tuple with the KeyFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFilePath

`func (o *CertificateProfile) SetKeyFilePath(v string)`

SetKeyFilePath sets KeyFilePath field to given value.

### HasKeyFilePath

`func (o *CertificateProfile) HasKeyFilePath() bool`

HasKeyFilePath returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CertificateProfile) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CertificateProfile) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CertificateProfile) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CertificateProfile) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *CertificateProfile) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *CertificateProfile) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetCipher

`func (o *CertificateProfile) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *CertificateProfile) GetCipherOk() (*string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipher

`func (o *CertificateProfile) SetCipher(v string)`

SetCipher sets Cipher field to given value.

### HasCipher

`func (o *CertificateProfile) HasCipher() bool`

HasCipher returns a boolean if a field has been set.

### GetTenant

`func (o *CertificateProfile) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CertificateProfile) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CertificateProfile) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CertificateProfile) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *CertificateProfile) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *CertificateProfile) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCreated

`func (o *CertificateProfile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CertificateProfile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CertificateProfile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *CertificateProfile) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CertificateProfile) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *CertificateProfile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CertificateProfile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CertificateProfile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *CertificateProfile) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CertificateProfile) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *CertificateProfile) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *CertificateProfile) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *CertificateProfile) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *CertificateProfile) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CertificateProfile) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CertificateProfile) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CertificateProfile) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *CertificateProfile) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CertificateProfile) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CertificateProfile) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CertificateProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


