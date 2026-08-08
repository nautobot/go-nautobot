# VPNPhase1PolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**EncryptionAlgorithm** | Pointer to [**[]EncryptionAlgorithmEnum**](EncryptionAlgorithmEnum.md) |  | [optional] 
**IntegrityAlgorithm** | Pointer to [**[]IntegrityAlgorithmEnum**](IntegrityAlgorithmEnum.md) |  | [optional] 
**DhGroup** | Pointer to [**[]VPNPhase2PolicyChoices**](VPNPhase2PolicyChoices.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IkeVersion** | Pointer to [**BulkWritableVPNPhase1PolicyRequestIkeVersion**](BulkWritableVPNPhase1PolicyRequestIkeVersion.md) |  | [optional] 
**AggressiveMode** | Pointer to **bool** | Only applicable to IKEv1 | [optional] 
**LifetimeSeconds** | Pointer to **NullableInt32** |  | [optional] 
**LifetimeKb** | Pointer to **NullableInt32** |  | [optional] 
**AuthenticationMethod** | Pointer to [**BulkWritableVPNPhase1PolicyRequestAuthenticationMethod**](BulkWritableVPNPhase1PolicyRequestAuthenticationMethod.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewVPNPhase1PolicyRequest

`func NewVPNPhase1PolicyRequest(name string, ) *VPNPhase1PolicyRequest`

NewVPNPhase1PolicyRequest instantiates a new VPNPhase1PolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNPhase1PolicyRequestWithDefaults

`func NewVPNPhase1PolicyRequestWithDefaults() *VPNPhase1PolicyRequest`

NewVPNPhase1PolicyRequestWithDefaults instantiates a new VPNPhase1PolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNPhase1PolicyRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNPhase1PolicyRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNPhase1PolicyRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNPhase1PolicyRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *VPNPhase1PolicyRequest) GetEncryptionAlgorithm() []EncryptionAlgorithmEnum`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *VPNPhase1PolicyRequest) GetEncryptionAlgorithmOk() (*[]EncryptionAlgorithmEnum, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *VPNPhase1PolicyRequest) SetEncryptionAlgorithm(v []EncryptionAlgorithmEnum)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *VPNPhase1PolicyRequest) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetIntegrityAlgorithm

`func (o *VPNPhase1PolicyRequest) GetIntegrityAlgorithm() []IntegrityAlgorithmEnum`

GetIntegrityAlgorithm returns the IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetIntegrityAlgorithmOk

`func (o *VPNPhase1PolicyRequest) GetIntegrityAlgorithmOk() (*[]IntegrityAlgorithmEnum, bool)`

GetIntegrityAlgorithmOk returns a tuple with the IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityAlgorithm

`func (o *VPNPhase1PolicyRequest) SetIntegrityAlgorithm(v []IntegrityAlgorithmEnum)`

SetIntegrityAlgorithm sets IntegrityAlgorithm field to given value.

### HasIntegrityAlgorithm

`func (o *VPNPhase1PolicyRequest) HasIntegrityAlgorithm() bool`

HasIntegrityAlgorithm returns a boolean if a field has been set.

### GetDhGroup

`func (o *VPNPhase1PolicyRequest) GetDhGroup() []VPNPhase2PolicyChoices`

GetDhGroup returns the DhGroup field if non-nil, zero value otherwise.

### GetDhGroupOk

`func (o *VPNPhase1PolicyRequest) GetDhGroupOk() (*[]VPNPhase2PolicyChoices, bool)`

GetDhGroupOk returns a tuple with the DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhGroup

`func (o *VPNPhase1PolicyRequest) SetDhGroup(v []VPNPhase2PolicyChoices)`

SetDhGroup sets DhGroup field to given value.

### HasDhGroup

`func (o *VPNPhase1PolicyRequest) HasDhGroup() bool`

HasDhGroup returns a boolean if a field has been set.

### GetName

`func (o *VPNPhase1PolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPNPhase1PolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPNPhase1PolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VPNPhase1PolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VPNPhase1PolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VPNPhase1PolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VPNPhase1PolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIkeVersion

`func (o *VPNPhase1PolicyRequest) GetIkeVersion() BulkWritableVPNPhase1PolicyRequestIkeVersion`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *VPNPhase1PolicyRequest) GetIkeVersionOk() (*BulkWritableVPNPhase1PolicyRequestIkeVersion, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *VPNPhase1PolicyRequest) SetIkeVersion(v BulkWritableVPNPhase1PolicyRequestIkeVersion)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *VPNPhase1PolicyRequest) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetAggressiveMode

`func (o *VPNPhase1PolicyRequest) GetAggressiveMode() bool`

GetAggressiveMode returns the AggressiveMode field if non-nil, zero value otherwise.

### GetAggressiveModeOk

`func (o *VPNPhase1PolicyRequest) GetAggressiveModeOk() (*bool, bool)`

GetAggressiveModeOk returns a tuple with the AggressiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggressiveMode

`func (o *VPNPhase1PolicyRequest) SetAggressiveMode(v bool)`

SetAggressiveMode sets AggressiveMode field to given value.

### HasAggressiveMode

`func (o *VPNPhase1PolicyRequest) HasAggressiveMode() bool`

HasAggressiveMode returns a boolean if a field has been set.

### GetLifetimeSeconds

`func (o *VPNPhase1PolicyRequest) GetLifetimeSeconds() int32`

GetLifetimeSeconds returns the LifetimeSeconds field if non-nil, zero value otherwise.

### GetLifetimeSecondsOk

`func (o *VPNPhase1PolicyRequest) GetLifetimeSecondsOk() (*int32, bool)`

GetLifetimeSecondsOk returns a tuple with the LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeSeconds

`func (o *VPNPhase1PolicyRequest) SetLifetimeSeconds(v int32)`

SetLifetimeSeconds sets LifetimeSeconds field to given value.

### HasLifetimeSeconds

`func (o *VPNPhase1PolicyRequest) HasLifetimeSeconds() bool`

HasLifetimeSeconds returns a boolean if a field has been set.

### SetLifetimeSecondsNil

`func (o *VPNPhase1PolicyRequest) SetLifetimeSecondsNil(b bool)`

 SetLifetimeSecondsNil sets the value for LifetimeSeconds to be an explicit nil

### UnsetLifetimeSeconds
`func (o *VPNPhase1PolicyRequest) UnsetLifetimeSeconds()`

UnsetLifetimeSeconds ensures that no value is present for LifetimeSeconds, not even an explicit nil
### GetLifetimeKb

`func (o *VPNPhase1PolicyRequest) GetLifetimeKb() int32`

GetLifetimeKb returns the LifetimeKb field if non-nil, zero value otherwise.

### GetLifetimeKbOk

`func (o *VPNPhase1PolicyRequest) GetLifetimeKbOk() (*int32, bool)`

GetLifetimeKbOk returns a tuple with the LifetimeKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeKb

`func (o *VPNPhase1PolicyRequest) SetLifetimeKb(v int32)`

SetLifetimeKb sets LifetimeKb field to given value.

### HasLifetimeKb

`func (o *VPNPhase1PolicyRequest) HasLifetimeKb() bool`

HasLifetimeKb returns a boolean if a field has been set.

### SetLifetimeKbNil

`func (o *VPNPhase1PolicyRequest) SetLifetimeKbNil(b bool)`

 SetLifetimeKbNil sets the value for LifetimeKb to be an explicit nil

### UnsetLifetimeKb
`func (o *VPNPhase1PolicyRequest) UnsetLifetimeKb()`

UnsetLifetimeKb ensures that no value is present for LifetimeKb, not even an explicit nil
### GetAuthenticationMethod

`func (o *VPNPhase1PolicyRequest) GetAuthenticationMethod() BulkWritableVPNPhase1PolicyRequestAuthenticationMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *VPNPhase1PolicyRequest) GetAuthenticationMethodOk() (*BulkWritableVPNPhase1PolicyRequestAuthenticationMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *VPNPhase1PolicyRequest) SetAuthenticationMethod(v BulkWritableVPNPhase1PolicyRequestAuthenticationMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *VPNPhase1PolicyRequest) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetTenant

`func (o *VPNPhase1PolicyRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VPNPhase1PolicyRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VPNPhase1PolicyRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VPNPhase1PolicyRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VPNPhase1PolicyRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VPNPhase1PolicyRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *VPNPhase1PolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNPhase1PolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNPhase1PolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNPhase1PolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *VPNPhase1PolicyRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *VPNPhase1PolicyRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *VPNPhase1PolicyRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *VPNPhase1PolicyRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *VPNPhase1PolicyRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VPNPhase1PolicyRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VPNPhase1PolicyRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VPNPhase1PolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


