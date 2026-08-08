# PatchedBulkWritableCertificateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableCertificateProfileRequest

`func NewPatchedBulkWritableCertificateProfileRequest(id string, ) *PatchedBulkWritableCertificateProfileRequest`

NewPatchedBulkWritableCertificateProfileRequest instantiates a new PatchedBulkWritableCertificateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableCertificateProfileRequestWithDefaults

`func NewPatchedBulkWritableCertificateProfileRequestWithDefaults() *PatchedBulkWritableCertificateProfileRequest`

NewPatchedBulkWritableCertificateProfileRequestWithDefaults instantiates a new PatchedBulkWritableCertificateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableCertificateProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableCertificateProfileRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableCertificateProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableCertificateProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableCertificateProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertificateType

`func (o *PatchedBulkWritableCertificateProfileRequest) GetCertificateType() BulkWritableCertificateProfileRequestCertificateType`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetCertificateTypeOk() (*BulkWritableCertificateProfileRequestCertificateType, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *PatchedBulkWritableCertificateProfileRequest) SetCertificateType(v BulkWritableCertificateProfileRequestCertificateType)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *PatchedBulkWritableCertificateProfileRequest) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetCertificateFilePath

`func (o *PatchedBulkWritableCertificateProfileRequest) GetCertificateFilePath() string`

GetCertificateFilePath returns the CertificateFilePath field if non-nil, zero value otherwise.

### GetCertificateFilePathOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetCertificateFilePathOk() (*string, bool)`

GetCertificateFilePathOk returns a tuple with the CertificateFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFilePath

`func (o *PatchedBulkWritableCertificateProfileRequest) SetCertificateFilePath(v string)`

SetCertificateFilePath sets CertificateFilePath field to given value.

### HasCertificateFilePath

`func (o *PatchedBulkWritableCertificateProfileRequest) HasCertificateFilePath() bool`

HasCertificateFilePath returns a boolean if a field has been set.

### GetChainFilePath

`func (o *PatchedBulkWritableCertificateProfileRequest) GetChainFilePath() string`

GetChainFilePath returns the ChainFilePath field if non-nil, zero value otherwise.

### GetChainFilePathOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetChainFilePathOk() (*string, bool)`

GetChainFilePathOk returns a tuple with the ChainFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainFilePath

`func (o *PatchedBulkWritableCertificateProfileRequest) SetChainFilePath(v string)`

SetChainFilePath sets ChainFilePath field to given value.

### HasChainFilePath

`func (o *PatchedBulkWritableCertificateProfileRequest) HasChainFilePath() bool`

HasChainFilePath returns a boolean if a field has been set.

### GetKeyFilePath

`func (o *PatchedBulkWritableCertificateProfileRequest) GetKeyFilePath() string`

GetKeyFilePath returns the KeyFilePath field if non-nil, zero value otherwise.

### GetKeyFilePathOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetKeyFilePathOk() (*string, bool)`

GetKeyFilePathOk returns a tuple with the KeyFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFilePath

`func (o *PatchedBulkWritableCertificateProfileRequest) SetKeyFilePath(v string)`

SetKeyFilePath sets KeyFilePath field to given value.

### HasKeyFilePath

`func (o *PatchedBulkWritableCertificateProfileRequest) HasKeyFilePath() bool`

HasKeyFilePath returns a boolean if a field has been set.

### GetExpirationDate

`func (o *PatchedBulkWritableCertificateProfileRequest) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PatchedBulkWritableCertificateProfileRequest) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PatchedBulkWritableCertificateProfileRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PatchedBulkWritableCertificateProfileRequest) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PatchedBulkWritableCertificateProfileRequest) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetCipher

`func (o *PatchedBulkWritableCertificateProfileRequest) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetCipherOk() (*string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipher

`func (o *PatchedBulkWritableCertificateProfileRequest) SetCipher(v string)`

SetCipher sets Cipher field to given value.

### HasCipher

`func (o *PatchedBulkWritableCertificateProfileRequest) HasCipher() bool`

HasCipher returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkWritableCertificateProfileRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableCertificateProfileRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableCertificateProfileRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableCertificateProfileRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableCertificateProfileRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableCertificateProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableCertificateProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableCertificateProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableCertificateProfileRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableCertificateProfileRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableCertificateProfileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableCertificateProfileRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableCertificateProfileRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableCertificateProfileRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableCertificateProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


