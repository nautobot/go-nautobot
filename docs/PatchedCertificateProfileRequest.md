# PatchedCertificateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CertificateType** | Pointer to [**BulkWritableCertificateProfileRequestCertificateType**](BulkWritableCertificateProfileRequestCertificateType.md) |  | [optional] 
**CertificateFilePath** | Pointer to **string** |  | [optional] 
**ChainFilePath** | Pointer to **string** |  | [optional] 
**KeyFilePath** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **NullableTime** |  | [optional] 
**Cipher** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedCertificateProfileRequest

`func NewPatchedCertificateProfileRequest() *PatchedCertificateProfileRequest`

NewPatchedCertificateProfileRequest instantiates a new PatchedCertificateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCertificateProfileRequestWithDefaults

`func NewPatchedCertificateProfileRequestWithDefaults() *PatchedCertificateProfileRequest`

NewPatchedCertificateProfileRequestWithDefaults instantiates a new PatchedCertificateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedCertificateProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedCertificateProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedCertificateProfileRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedCertificateProfileRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedCertificateProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCertificateProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCertificateProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCertificateProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertificateType

`func (o *PatchedCertificateProfileRequest) GetCertificateType() BulkWritableCertificateProfileRequestCertificateType`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *PatchedCertificateProfileRequest) GetCertificateTypeOk() (*BulkWritableCertificateProfileRequestCertificateType, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *PatchedCertificateProfileRequest) SetCertificateType(v BulkWritableCertificateProfileRequestCertificateType)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *PatchedCertificateProfileRequest) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetCertificateFilePath

`func (o *PatchedCertificateProfileRequest) GetCertificateFilePath() string`

GetCertificateFilePath returns the CertificateFilePath field if non-nil, zero value otherwise.

### GetCertificateFilePathOk

`func (o *PatchedCertificateProfileRequest) GetCertificateFilePathOk() (*string, bool)`

GetCertificateFilePathOk returns a tuple with the CertificateFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFilePath

`func (o *PatchedCertificateProfileRequest) SetCertificateFilePath(v string)`

SetCertificateFilePath sets CertificateFilePath field to given value.

### HasCertificateFilePath

`func (o *PatchedCertificateProfileRequest) HasCertificateFilePath() bool`

HasCertificateFilePath returns a boolean if a field has been set.

### GetChainFilePath

`func (o *PatchedCertificateProfileRequest) GetChainFilePath() string`

GetChainFilePath returns the ChainFilePath field if non-nil, zero value otherwise.

### GetChainFilePathOk

`func (o *PatchedCertificateProfileRequest) GetChainFilePathOk() (*string, bool)`

GetChainFilePathOk returns a tuple with the ChainFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainFilePath

`func (o *PatchedCertificateProfileRequest) SetChainFilePath(v string)`

SetChainFilePath sets ChainFilePath field to given value.

### HasChainFilePath

`func (o *PatchedCertificateProfileRequest) HasChainFilePath() bool`

HasChainFilePath returns a boolean if a field has been set.

### GetKeyFilePath

`func (o *PatchedCertificateProfileRequest) GetKeyFilePath() string`

GetKeyFilePath returns the KeyFilePath field if non-nil, zero value otherwise.

### GetKeyFilePathOk

`func (o *PatchedCertificateProfileRequest) GetKeyFilePathOk() (*string, bool)`

GetKeyFilePathOk returns a tuple with the KeyFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFilePath

`func (o *PatchedCertificateProfileRequest) SetKeyFilePath(v string)`

SetKeyFilePath sets KeyFilePath field to given value.

### HasKeyFilePath

`func (o *PatchedCertificateProfileRequest) HasKeyFilePath() bool`

HasKeyFilePath returns a boolean if a field has been set.

### GetExpirationDate

`func (o *PatchedCertificateProfileRequest) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PatchedCertificateProfileRequest) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PatchedCertificateProfileRequest) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PatchedCertificateProfileRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PatchedCertificateProfileRequest) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PatchedCertificateProfileRequest) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetCipher

`func (o *PatchedCertificateProfileRequest) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *PatchedCertificateProfileRequest) GetCipherOk() (*string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipher

`func (o *PatchedCertificateProfileRequest) SetCipher(v string)`

SetCipher sets Cipher field to given value.

### HasCipher

`func (o *PatchedCertificateProfileRequest) HasCipher() bool`

HasCipher returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedCertificateProfileRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedCertificateProfileRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedCertificateProfileRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedCertificateProfileRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedCertificateProfileRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedCertificateProfileRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedCertificateProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedCertificateProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedCertificateProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedCertificateProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedCertificateProfileRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedCertificateProfileRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedCertificateProfileRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedCertificateProfileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedCertificateProfileRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedCertificateProfileRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedCertificateProfileRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedCertificateProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


