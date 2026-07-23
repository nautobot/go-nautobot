# BulkWritableVPNPhase1PolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableVPNPhase1PolicyRequest

`func NewBulkWritableVPNPhase1PolicyRequest(id string, name string, ) *BulkWritableVPNPhase1PolicyRequest`

NewBulkWritableVPNPhase1PolicyRequest instantiates a new BulkWritableVPNPhase1PolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVPNPhase1PolicyRequestWithDefaults

`func NewBulkWritableVPNPhase1PolicyRequestWithDefaults() *BulkWritableVPNPhase1PolicyRequest`

NewBulkWritableVPNPhase1PolicyRequestWithDefaults instantiates a new BulkWritableVPNPhase1PolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVPNPhase1PolicyRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVPNPhase1PolicyRequest) SetId(v string)`

SetId sets Id field to given value.


### GetEncryptionAlgorithm

`func (o *BulkWritableVPNPhase1PolicyRequest) GetEncryptionAlgorithm() []EncryptionAlgorithmEnum`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetEncryptionAlgorithmOk() (*[]EncryptionAlgorithmEnum, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *BulkWritableVPNPhase1PolicyRequest) SetEncryptionAlgorithm(v []EncryptionAlgorithmEnum)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *BulkWritableVPNPhase1PolicyRequest) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetIntegrityAlgorithm

`func (o *BulkWritableVPNPhase1PolicyRequest) GetIntegrityAlgorithm() []IntegrityAlgorithmEnum`

GetIntegrityAlgorithm returns the IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetIntegrityAlgorithmOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetIntegrityAlgorithmOk() (*[]IntegrityAlgorithmEnum, bool)`

GetIntegrityAlgorithmOk returns a tuple with the IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityAlgorithm

`func (o *BulkWritableVPNPhase1PolicyRequest) SetIntegrityAlgorithm(v []IntegrityAlgorithmEnum)`

SetIntegrityAlgorithm sets IntegrityAlgorithm field to given value.

### HasIntegrityAlgorithm

`func (o *BulkWritableVPNPhase1PolicyRequest) HasIntegrityAlgorithm() bool`

HasIntegrityAlgorithm returns a boolean if a field has been set.

### GetDhGroup

`func (o *BulkWritableVPNPhase1PolicyRequest) GetDhGroup() []VPNPhase2PolicyChoices`

GetDhGroup returns the DhGroup field if non-nil, zero value otherwise.

### GetDhGroupOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetDhGroupOk() (*[]VPNPhase2PolicyChoices, bool)`

GetDhGroupOk returns a tuple with the DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhGroup

`func (o *BulkWritableVPNPhase1PolicyRequest) SetDhGroup(v []VPNPhase2PolicyChoices)`

SetDhGroup sets DhGroup field to given value.

### HasDhGroup

`func (o *BulkWritableVPNPhase1PolicyRequest) HasDhGroup() bool`

HasDhGroup returns a boolean if a field has been set.

### GetName

`func (o *BulkWritableVPNPhase1PolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableVPNPhase1PolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableVPNPhase1PolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableVPNPhase1PolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableVPNPhase1PolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIkeVersion

`func (o *BulkWritableVPNPhase1PolicyRequest) GetIkeVersion() BulkWritableVPNPhase1PolicyRequestIkeVersion`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetIkeVersionOk() (*BulkWritableVPNPhase1PolicyRequestIkeVersion, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *BulkWritableVPNPhase1PolicyRequest) SetIkeVersion(v BulkWritableVPNPhase1PolicyRequestIkeVersion)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *BulkWritableVPNPhase1PolicyRequest) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetAggressiveMode

`func (o *BulkWritableVPNPhase1PolicyRequest) GetAggressiveMode() bool`

GetAggressiveMode returns the AggressiveMode field if non-nil, zero value otherwise.

### GetAggressiveModeOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetAggressiveModeOk() (*bool, bool)`

GetAggressiveModeOk returns a tuple with the AggressiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggressiveMode

`func (o *BulkWritableVPNPhase1PolicyRequest) SetAggressiveMode(v bool)`

SetAggressiveMode sets AggressiveMode field to given value.

### HasAggressiveMode

`func (o *BulkWritableVPNPhase1PolicyRequest) HasAggressiveMode() bool`

HasAggressiveMode returns a boolean if a field has been set.

### GetLifetimeSeconds

`func (o *BulkWritableVPNPhase1PolicyRequest) GetLifetimeSeconds() int32`

GetLifetimeSeconds returns the LifetimeSeconds field if non-nil, zero value otherwise.

### GetLifetimeSecondsOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetLifetimeSecondsOk() (*int32, bool)`

GetLifetimeSecondsOk returns a tuple with the LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeSeconds

`func (o *BulkWritableVPNPhase1PolicyRequest) SetLifetimeSeconds(v int32)`

SetLifetimeSeconds sets LifetimeSeconds field to given value.

### HasLifetimeSeconds

`func (o *BulkWritableVPNPhase1PolicyRequest) HasLifetimeSeconds() bool`

HasLifetimeSeconds returns a boolean if a field has been set.

### SetLifetimeSecondsNil

`func (o *BulkWritableVPNPhase1PolicyRequest) SetLifetimeSecondsNil(b bool)`

 SetLifetimeSecondsNil sets the value for LifetimeSeconds to be an explicit nil

### UnsetLifetimeSeconds
`func (o *BulkWritableVPNPhase1PolicyRequest) UnsetLifetimeSeconds()`

UnsetLifetimeSeconds ensures that no value is present for LifetimeSeconds, not even an explicit nil
### GetLifetimeKb

`func (o *BulkWritableVPNPhase1PolicyRequest) GetLifetimeKb() int32`

GetLifetimeKb returns the LifetimeKb field if non-nil, zero value otherwise.

### GetLifetimeKbOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetLifetimeKbOk() (*int32, bool)`

GetLifetimeKbOk returns a tuple with the LifetimeKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeKb

`func (o *BulkWritableVPNPhase1PolicyRequest) SetLifetimeKb(v int32)`

SetLifetimeKb sets LifetimeKb field to given value.

### HasLifetimeKb

`func (o *BulkWritableVPNPhase1PolicyRequest) HasLifetimeKb() bool`

HasLifetimeKb returns a boolean if a field has been set.

### SetLifetimeKbNil

`func (o *BulkWritableVPNPhase1PolicyRequest) SetLifetimeKbNil(b bool)`

 SetLifetimeKbNil sets the value for LifetimeKb to be an explicit nil

### UnsetLifetimeKb
`func (o *BulkWritableVPNPhase1PolicyRequest) UnsetLifetimeKb()`

UnsetLifetimeKb ensures that no value is present for LifetimeKb, not even an explicit nil
### GetAuthenticationMethod

`func (o *BulkWritableVPNPhase1PolicyRequest) GetAuthenticationMethod() BulkWritableVPNPhase1PolicyRequestAuthenticationMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetAuthenticationMethodOk() (*BulkWritableVPNPhase1PolicyRequestAuthenticationMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *BulkWritableVPNPhase1PolicyRequest) SetAuthenticationMethod(v BulkWritableVPNPhase1PolicyRequestAuthenticationMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *BulkWritableVPNPhase1PolicyRequest) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetTenant

`func (o *BulkWritableVPNPhase1PolicyRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableVPNPhase1PolicyRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableVPNPhase1PolicyRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableVPNPhase1PolicyRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableVPNPhase1PolicyRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableVPNPhase1PolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableVPNPhase1PolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableVPNPhase1PolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableVPNPhase1PolicyRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableVPNPhase1PolicyRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableVPNPhase1PolicyRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableVPNPhase1PolicyRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableVPNPhase1PolicyRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableVPNPhase1PolicyRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableVPNPhase1PolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


