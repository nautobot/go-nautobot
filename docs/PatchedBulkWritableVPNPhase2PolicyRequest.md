# PatchedBulkWritableVPNPhase2PolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**EncryptionAlgorithm** | Pointer to [**[]EncryptionAlgorithmEnum**](EncryptionAlgorithmEnum.md) |  | [optional] 
**IntegrityAlgorithm** | Pointer to [**[]IntegrityAlgorithmEnum**](IntegrityAlgorithmEnum.md) |  | [optional] 
**PfsGroup** | Pointer to [**[]VPNPhase2PolicyChoices**](VPNPhase2PolicyChoices.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Lifetime** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVPNPhase2PolicyRequest

`func NewPatchedBulkWritableVPNPhase2PolicyRequest(id string, ) *PatchedBulkWritableVPNPhase2PolicyRequest`

NewPatchedBulkWritableVPNPhase2PolicyRequest instantiates a new PatchedBulkWritableVPNPhase2PolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVPNPhase2PolicyRequestWithDefaults

`func NewPatchedBulkWritableVPNPhase2PolicyRequestWithDefaults() *PatchedBulkWritableVPNPhase2PolicyRequest`

NewPatchedBulkWritableVPNPhase2PolicyRequestWithDefaults instantiates a new PatchedBulkWritableVPNPhase2PolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetId(v string)`

SetId sets Id field to given value.


### GetEncryptionAlgorithm

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetEncryptionAlgorithm() []EncryptionAlgorithmEnum`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetEncryptionAlgorithmOk() (*[]EncryptionAlgorithmEnum, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetEncryptionAlgorithm(v []EncryptionAlgorithmEnum)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetIntegrityAlgorithm

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetIntegrityAlgorithm() []IntegrityAlgorithmEnum`

GetIntegrityAlgorithm returns the IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetIntegrityAlgorithmOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetIntegrityAlgorithmOk() (*[]IntegrityAlgorithmEnum, bool)`

GetIntegrityAlgorithmOk returns a tuple with the IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityAlgorithm

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetIntegrityAlgorithm(v []IntegrityAlgorithmEnum)`

SetIntegrityAlgorithm sets IntegrityAlgorithm field to given value.

### HasIntegrityAlgorithm

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) HasIntegrityAlgorithm() bool`

HasIntegrityAlgorithm returns a boolean if a field has been set.

### GetPfsGroup

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetPfsGroup() []VPNPhase2PolicyChoices`

GetPfsGroup returns the PfsGroup field if non-nil, zero value otherwise.

### GetPfsGroupOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetPfsGroupOk() (*[]VPNPhase2PolicyChoices, bool)`

GetPfsGroupOk returns a tuple with the PfsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsGroup

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetPfsGroup(v []VPNPhase2PolicyChoices)`

SetPfsGroup sets PfsGroup field to given value.

### HasPfsGroup

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) HasPfsGroup() bool`

HasPfsGroup returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLifetime

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### SetLifetimeNil

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetLifetimeNil(b bool)`

 SetLifetimeNil sets the value for Lifetime to be an explicit nil

### UnsetLifetime
`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) UnsetLifetime()`

UnsetLifetime ensures that no value is present for Lifetime, not even an explicit nil
### GetTenant

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableVPNPhase2PolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


