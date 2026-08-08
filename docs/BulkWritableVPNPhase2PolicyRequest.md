# BulkWritableVPNPhase2PolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EncryptionAlgorithm** | Pointer to [**[]EncryptionAlgorithmEnum**](EncryptionAlgorithmEnum.md) |  | [optional] 
**IntegrityAlgorithm** | Pointer to [**[]IntegrityAlgorithmEnum**](IntegrityAlgorithmEnum.md) |  | [optional] 
**PfsGroup** | Pointer to [**[]VPNPhase2PolicyChoices**](VPNPhase2PolicyChoices.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Lifetime** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewBulkWritableVPNPhase2PolicyRequest

`func NewBulkWritableVPNPhase2PolicyRequest(id string, name string, ) *BulkWritableVPNPhase2PolicyRequest`

NewBulkWritableVPNPhase2PolicyRequest instantiates a new BulkWritableVPNPhase2PolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVPNPhase2PolicyRequestWithDefaults

`func NewBulkWritableVPNPhase2PolicyRequestWithDefaults() *BulkWritableVPNPhase2PolicyRequest`

NewBulkWritableVPNPhase2PolicyRequestWithDefaults instantiates a new BulkWritableVPNPhase2PolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVPNPhase2PolicyRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVPNPhase2PolicyRequest) SetId(v string)`

SetId sets Id field to given value.


### GetEncryptionAlgorithm

`func (o *BulkWritableVPNPhase2PolicyRequest) GetEncryptionAlgorithm() []EncryptionAlgorithmEnum`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetEncryptionAlgorithmOk() (*[]EncryptionAlgorithmEnum, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *BulkWritableVPNPhase2PolicyRequest) SetEncryptionAlgorithm(v []EncryptionAlgorithmEnum)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *BulkWritableVPNPhase2PolicyRequest) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetIntegrityAlgorithm

`func (o *BulkWritableVPNPhase2PolicyRequest) GetIntegrityAlgorithm() []IntegrityAlgorithmEnum`

GetIntegrityAlgorithm returns the IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetIntegrityAlgorithmOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetIntegrityAlgorithmOk() (*[]IntegrityAlgorithmEnum, bool)`

GetIntegrityAlgorithmOk returns a tuple with the IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityAlgorithm

`func (o *BulkWritableVPNPhase2PolicyRequest) SetIntegrityAlgorithm(v []IntegrityAlgorithmEnum)`

SetIntegrityAlgorithm sets IntegrityAlgorithm field to given value.

### HasIntegrityAlgorithm

`func (o *BulkWritableVPNPhase2PolicyRequest) HasIntegrityAlgorithm() bool`

HasIntegrityAlgorithm returns a boolean if a field has been set.

### GetPfsGroup

`func (o *BulkWritableVPNPhase2PolicyRequest) GetPfsGroup() []VPNPhase2PolicyChoices`

GetPfsGroup returns the PfsGroup field if non-nil, zero value otherwise.

### GetPfsGroupOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetPfsGroupOk() (*[]VPNPhase2PolicyChoices, bool)`

GetPfsGroupOk returns a tuple with the PfsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsGroup

`func (o *BulkWritableVPNPhase2PolicyRequest) SetPfsGroup(v []VPNPhase2PolicyChoices)`

SetPfsGroup sets PfsGroup field to given value.

### HasPfsGroup

`func (o *BulkWritableVPNPhase2PolicyRequest) HasPfsGroup() bool`

HasPfsGroup returns a boolean if a field has been set.

### GetName

`func (o *BulkWritableVPNPhase2PolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableVPNPhase2PolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableVPNPhase2PolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableVPNPhase2PolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableVPNPhase2PolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLifetime

`func (o *BulkWritableVPNPhase2PolicyRequest) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *BulkWritableVPNPhase2PolicyRequest) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *BulkWritableVPNPhase2PolicyRequest) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### SetLifetimeNil

`func (o *BulkWritableVPNPhase2PolicyRequest) SetLifetimeNil(b bool)`

 SetLifetimeNil sets the value for Lifetime to be an explicit nil

### UnsetLifetime
`func (o *BulkWritableVPNPhase2PolicyRequest) UnsetLifetime()`

UnsetLifetime ensures that no value is present for Lifetime, not even an explicit nil
### GetTenant

`func (o *BulkWritableVPNPhase2PolicyRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableVPNPhase2PolicyRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableVPNPhase2PolicyRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableVPNPhase2PolicyRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableVPNPhase2PolicyRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableVPNPhase2PolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableVPNPhase2PolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableVPNPhase2PolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableVPNPhase2PolicyRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableVPNPhase2PolicyRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableVPNPhase2PolicyRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableVPNPhase2PolicyRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableVPNPhase2PolicyRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableVPNPhase2PolicyRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableVPNPhase2PolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


