# VPNPhase2Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**EncryptionAlgorithm** | Pointer to [**[]EncryptionAlgorithmEnum**](EncryptionAlgorithmEnum.md) |  | [optional] 
**IntegrityAlgorithm** | Pointer to [**[]IntegrityAlgorithmEnum**](IntegrityAlgorithmEnum.md) |  | [optional] 
**PfsGroup** | Pointer to [**[]VPNPhase2PolicyChoices**](VPNPhase2PolicyChoices.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Lifetime** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewVPNPhase2Policy

`func NewVPNPhase2Policy(objectType string, display string, url string, naturalSlug string, name string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *VPNPhase2Policy`

NewVPNPhase2Policy instantiates a new VPNPhase2Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNPhase2PolicyWithDefaults

`func NewVPNPhase2PolicyWithDefaults() *VPNPhase2Policy`

NewVPNPhase2PolicyWithDefaults instantiates a new VPNPhase2Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNPhase2Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNPhase2Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNPhase2Policy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNPhase2Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VPNPhase2Policy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VPNPhase2Policy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VPNPhase2Policy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VPNPhase2Policy) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VPNPhase2Policy) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VPNPhase2Policy) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VPNPhase2Policy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VPNPhase2Policy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VPNPhase2Policy) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VPNPhase2Policy) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VPNPhase2Policy) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VPNPhase2Policy) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetEncryptionAlgorithm

`func (o *VPNPhase2Policy) GetEncryptionAlgorithm() []EncryptionAlgorithmEnum`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *VPNPhase2Policy) GetEncryptionAlgorithmOk() (*[]EncryptionAlgorithmEnum, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *VPNPhase2Policy) SetEncryptionAlgorithm(v []EncryptionAlgorithmEnum)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *VPNPhase2Policy) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetIntegrityAlgorithm

`func (o *VPNPhase2Policy) GetIntegrityAlgorithm() []IntegrityAlgorithmEnum`

GetIntegrityAlgorithm returns the IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetIntegrityAlgorithmOk

`func (o *VPNPhase2Policy) GetIntegrityAlgorithmOk() (*[]IntegrityAlgorithmEnum, bool)`

GetIntegrityAlgorithmOk returns a tuple with the IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityAlgorithm

`func (o *VPNPhase2Policy) SetIntegrityAlgorithm(v []IntegrityAlgorithmEnum)`

SetIntegrityAlgorithm sets IntegrityAlgorithm field to given value.

### HasIntegrityAlgorithm

`func (o *VPNPhase2Policy) HasIntegrityAlgorithm() bool`

HasIntegrityAlgorithm returns a boolean if a field has been set.

### GetPfsGroup

`func (o *VPNPhase2Policy) GetPfsGroup() []VPNPhase2PolicyChoices`

GetPfsGroup returns the PfsGroup field if non-nil, zero value otherwise.

### GetPfsGroupOk

`func (o *VPNPhase2Policy) GetPfsGroupOk() (*[]VPNPhase2PolicyChoices, bool)`

GetPfsGroupOk returns a tuple with the PfsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsGroup

`func (o *VPNPhase2Policy) SetPfsGroup(v []VPNPhase2PolicyChoices)`

SetPfsGroup sets PfsGroup field to given value.

### HasPfsGroup

`func (o *VPNPhase2Policy) HasPfsGroup() bool`

HasPfsGroup returns a boolean if a field has been set.

### GetName

`func (o *VPNPhase2Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPNPhase2Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPNPhase2Policy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VPNPhase2Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VPNPhase2Policy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VPNPhase2Policy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VPNPhase2Policy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLifetime

`func (o *VPNPhase2Policy) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *VPNPhase2Policy) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *VPNPhase2Policy) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *VPNPhase2Policy) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### SetLifetimeNil

`func (o *VPNPhase2Policy) SetLifetimeNil(b bool)`

 SetLifetimeNil sets the value for Lifetime to be an explicit nil

### UnsetLifetime
`func (o *VPNPhase2Policy) UnsetLifetime()`

UnsetLifetime ensures that no value is present for Lifetime, not even an explicit nil
### GetTenant

`func (o *VPNPhase2Policy) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VPNPhase2Policy) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VPNPhase2Policy) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VPNPhase2Policy) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VPNPhase2Policy) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VPNPhase2Policy) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCreated

`func (o *VPNPhase2Policy) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VPNPhase2Policy) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VPNPhase2Policy) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VPNPhase2Policy) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VPNPhase2Policy) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VPNPhase2Policy) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VPNPhase2Policy) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VPNPhase2Policy) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VPNPhase2Policy) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VPNPhase2Policy) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *VPNPhase2Policy) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VPNPhase2Policy) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VPNPhase2Policy) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *VPNPhase2Policy) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNPhase2Policy) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNPhase2Policy) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNPhase2Policy) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *VPNPhase2Policy) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VPNPhase2Policy) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VPNPhase2Policy) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VPNPhase2Policy) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


