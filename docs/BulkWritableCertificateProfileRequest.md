# BulkWritableCertificateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
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

### NewBulkWritableCertificateProfileRequest

`func NewBulkWritableCertificateProfileRequest(id string, name string, ) *BulkWritableCertificateProfileRequest`

NewBulkWritableCertificateProfileRequest instantiates a new BulkWritableCertificateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableCertificateProfileRequestWithDefaults

`func NewBulkWritableCertificateProfileRequestWithDefaults() *BulkWritableCertificateProfileRequest`

NewBulkWritableCertificateProfileRequestWithDefaults instantiates a new BulkWritableCertificateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableCertificateProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableCertificateProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableCertificateProfileRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableCertificateProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableCertificateProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableCertificateProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCertificateType

`func (o *BulkWritableCertificateProfileRequest) GetCertificateType() BulkWritableCertificateProfileRequestCertificateType`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *BulkWritableCertificateProfileRequest) GetCertificateTypeOk() (*BulkWritableCertificateProfileRequestCertificateType, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *BulkWritableCertificateProfileRequest) SetCertificateType(v BulkWritableCertificateProfileRequestCertificateType)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *BulkWritableCertificateProfileRequest) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetCertificateFilePath

`func (o *BulkWritableCertificateProfileRequest) GetCertificateFilePath() string`

GetCertificateFilePath returns the CertificateFilePath field if non-nil, zero value otherwise.

### GetCertificateFilePathOk

`func (o *BulkWritableCertificateProfileRequest) GetCertificateFilePathOk() (*string, bool)`

GetCertificateFilePathOk returns a tuple with the CertificateFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFilePath

`func (o *BulkWritableCertificateProfileRequest) SetCertificateFilePath(v string)`

SetCertificateFilePath sets CertificateFilePath field to given value.

### HasCertificateFilePath

`func (o *BulkWritableCertificateProfileRequest) HasCertificateFilePath() bool`

HasCertificateFilePath returns a boolean if a field has been set.

### GetChainFilePath

`func (o *BulkWritableCertificateProfileRequest) GetChainFilePath() string`

GetChainFilePath returns the ChainFilePath field if non-nil, zero value otherwise.

### GetChainFilePathOk

`func (o *BulkWritableCertificateProfileRequest) GetChainFilePathOk() (*string, bool)`

GetChainFilePathOk returns a tuple with the ChainFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainFilePath

`func (o *BulkWritableCertificateProfileRequest) SetChainFilePath(v string)`

SetChainFilePath sets ChainFilePath field to given value.

### HasChainFilePath

`func (o *BulkWritableCertificateProfileRequest) HasChainFilePath() bool`

HasChainFilePath returns a boolean if a field has been set.

### GetKeyFilePath

`func (o *BulkWritableCertificateProfileRequest) GetKeyFilePath() string`

GetKeyFilePath returns the KeyFilePath field if non-nil, zero value otherwise.

### GetKeyFilePathOk

`func (o *BulkWritableCertificateProfileRequest) GetKeyFilePathOk() (*string, bool)`

GetKeyFilePathOk returns a tuple with the KeyFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFilePath

`func (o *BulkWritableCertificateProfileRequest) SetKeyFilePath(v string)`

SetKeyFilePath sets KeyFilePath field to given value.

### HasKeyFilePath

`func (o *BulkWritableCertificateProfileRequest) HasKeyFilePath() bool`

HasKeyFilePath returns a boolean if a field has been set.

### GetExpirationDate

`func (o *BulkWritableCertificateProfileRequest) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *BulkWritableCertificateProfileRequest) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *BulkWritableCertificateProfileRequest) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *BulkWritableCertificateProfileRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *BulkWritableCertificateProfileRequest) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *BulkWritableCertificateProfileRequest) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetCipher

`func (o *BulkWritableCertificateProfileRequest) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *BulkWritableCertificateProfileRequest) GetCipherOk() (*string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipher

`func (o *BulkWritableCertificateProfileRequest) SetCipher(v string)`

SetCipher sets Cipher field to given value.

### HasCipher

`func (o *BulkWritableCertificateProfileRequest) HasCipher() bool`

HasCipher returns a boolean if a field has been set.

### GetTenant

`func (o *BulkWritableCertificateProfileRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableCertificateProfileRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableCertificateProfileRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableCertificateProfileRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableCertificateProfileRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableCertificateProfileRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableCertificateProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableCertificateProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableCertificateProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableCertificateProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableCertificateProfileRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableCertificateProfileRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableCertificateProfileRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableCertificateProfileRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableCertificateProfileRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableCertificateProfileRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableCertificateProfileRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableCertificateProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


