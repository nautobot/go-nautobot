# VPNPhase1Policy

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
**DhGroup** | Pointer to [**[]VPNPhase2PolicyChoices**](VPNPhase2PolicyChoices.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IkeVersion** | Pointer to [**BulkWritableVPNPhase1PolicyRequestIkeVersion**](BulkWritableVPNPhase1PolicyRequestIkeVersion.md) |  | [optional] 
**AggressiveMode** | Pointer to **bool** | Only applicable to IKEv1 | [optional] 
**LifetimeSeconds** | Pointer to **NullableInt32** |  | [optional] 
**LifetimeKb** | Pointer to **NullableInt32** |  | [optional] 
**AuthenticationMethod** | Pointer to [**BulkWritableVPNPhase1PolicyRequestAuthenticationMethod**](BulkWritableVPNPhase1PolicyRequestAuthenticationMethod.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewVPNPhase1Policy

`func NewVPNPhase1Policy(objectType string, display string, url string, naturalSlug string, name string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *VPNPhase1Policy`

NewVPNPhase1Policy instantiates a new VPNPhase1Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNPhase1PolicyWithDefaults

`func NewVPNPhase1PolicyWithDefaults() *VPNPhase1Policy`

NewVPNPhase1PolicyWithDefaults instantiates a new VPNPhase1Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNPhase1Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNPhase1Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNPhase1Policy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNPhase1Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VPNPhase1Policy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VPNPhase1Policy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VPNPhase1Policy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VPNPhase1Policy) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VPNPhase1Policy) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VPNPhase1Policy) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VPNPhase1Policy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VPNPhase1Policy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VPNPhase1Policy) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VPNPhase1Policy) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VPNPhase1Policy) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VPNPhase1Policy) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetEncryptionAlgorithm

`func (o *VPNPhase1Policy) GetEncryptionAlgorithm() []EncryptionAlgorithmEnum`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *VPNPhase1Policy) GetEncryptionAlgorithmOk() (*[]EncryptionAlgorithmEnum, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *VPNPhase1Policy) SetEncryptionAlgorithm(v []EncryptionAlgorithmEnum)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *VPNPhase1Policy) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetIntegrityAlgorithm

`func (o *VPNPhase1Policy) GetIntegrityAlgorithm() []IntegrityAlgorithmEnum`

GetIntegrityAlgorithm returns the IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetIntegrityAlgorithmOk

`func (o *VPNPhase1Policy) GetIntegrityAlgorithmOk() (*[]IntegrityAlgorithmEnum, bool)`

GetIntegrityAlgorithmOk returns a tuple with the IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityAlgorithm

`func (o *VPNPhase1Policy) SetIntegrityAlgorithm(v []IntegrityAlgorithmEnum)`

SetIntegrityAlgorithm sets IntegrityAlgorithm field to given value.

### HasIntegrityAlgorithm

`func (o *VPNPhase1Policy) HasIntegrityAlgorithm() bool`

HasIntegrityAlgorithm returns a boolean if a field has been set.

### GetDhGroup

`func (o *VPNPhase1Policy) GetDhGroup() []VPNPhase2PolicyChoices`

GetDhGroup returns the DhGroup field if non-nil, zero value otherwise.

### GetDhGroupOk

`func (o *VPNPhase1Policy) GetDhGroupOk() (*[]VPNPhase2PolicyChoices, bool)`

GetDhGroupOk returns a tuple with the DhGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhGroup

`func (o *VPNPhase1Policy) SetDhGroup(v []VPNPhase2PolicyChoices)`

SetDhGroup sets DhGroup field to given value.

### HasDhGroup

`func (o *VPNPhase1Policy) HasDhGroup() bool`

HasDhGroup returns a boolean if a field has been set.

### GetName

`func (o *VPNPhase1Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPNPhase1Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPNPhase1Policy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VPNPhase1Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VPNPhase1Policy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VPNPhase1Policy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VPNPhase1Policy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIkeVersion

`func (o *VPNPhase1Policy) GetIkeVersion() BulkWritableVPNPhase1PolicyRequestIkeVersion`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *VPNPhase1Policy) GetIkeVersionOk() (*BulkWritableVPNPhase1PolicyRequestIkeVersion, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *VPNPhase1Policy) SetIkeVersion(v BulkWritableVPNPhase1PolicyRequestIkeVersion)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *VPNPhase1Policy) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### GetAggressiveMode

`func (o *VPNPhase1Policy) GetAggressiveMode() bool`

GetAggressiveMode returns the AggressiveMode field if non-nil, zero value otherwise.

### GetAggressiveModeOk

`func (o *VPNPhase1Policy) GetAggressiveModeOk() (*bool, bool)`

GetAggressiveModeOk returns a tuple with the AggressiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggressiveMode

`func (o *VPNPhase1Policy) SetAggressiveMode(v bool)`

SetAggressiveMode sets AggressiveMode field to given value.

### HasAggressiveMode

`func (o *VPNPhase1Policy) HasAggressiveMode() bool`

HasAggressiveMode returns a boolean if a field has been set.

### GetLifetimeSeconds

`func (o *VPNPhase1Policy) GetLifetimeSeconds() int32`

GetLifetimeSeconds returns the LifetimeSeconds field if non-nil, zero value otherwise.

### GetLifetimeSecondsOk

`func (o *VPNPhase1Policy) GetLifetimeSecondsOk() (*int32, bool)`

GetLifetimeSecondsOk returns a tuple with the LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeSeconds

`func (o *VPNPhase1Policy) SetLifetimeSeconds(v int32)`

SetLifetimeSeconds sets LifetimeSeconds field to given value.

### HasLifetimeSeconds

`func (o *VPNPhase1Policy) HasLifetimeSeconds() bool`

HasLifetimeSeconds returns a boolean if a field has been set.

### SetLifetimeSecondsNil

`func (o *VPNPhase1Policy) SetLifetimeSecondsNil(b bool)`

 SetLifetimeSecondsNil sets the value for LifetimeSeconds to be an explicit nil

### UnsetLifetimeSeconds
`func (o *VPNPhase1Policy) UnsetLifetimeSeconds()`

UnsetLifetimeSeconds ensures that no value is present for LifetimeSeconds, not even an explicit nil
### GetLifetimeKb

`func (o *VPNPhase1Policy) GetLifetimeKb() int32`

GetLifetimeKb returns the LifetimeKb field if non-nil, zero value otherwise.

### GetLifetimeKbOk

`func (o *VPNPhase1Policy) GetLifetimeKbOk() (*int32, bool)`

GetLifetimeKbOk returns a tuple with the LifetimeKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeKb

`func (o *VPNPhase1Policy) SetLifetimeKb(v int32)`

SetLifetimeKb sets LifetimeKb field to given value.

### HasLifetimeKb

`func (o *VPNPhase1Policy) HasLifetimeKb() bool`

HasLifetimeKb returns a boolean if a field has been set.

### SetLifetimeKbNil

`func (o *VPNPhase1Policy) SetLifetimeKbNil(b bool)`

 SetLifetimeKbNil sets the value for LifetimeKb to be an explicit nil

### UnsetLifetimeKb
`func (o *VPNPhase1Policy) UnsetLifetimeKb()`

UnsetLifetimeKb ensures that no value is present for LifetimeKb, not even an explicit nil
### GetAuthenticationMethod

`func (o *VPNPhase1Policy) GetAuthenticationMethod() BulkWritableVPNPhase1PolicyRequestAuthenticationMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *VPNPhase1Policy) GetAuthenticationMethodOk() (*BulkWritableVPNPhase1PolicyRequestAuthenticationMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *VPNPhase1Policy) SetAuthenticationMethod(v BulkWritableVPNPhase1PolicyRequestAuthenticationMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *VPNPhase1Policy) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetTenant

`func (o *VPNPhase1Policy) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VPNPhase1Policy) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VPNPhase1Policy) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VPNPhase1Policy) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VPNPhase1Policy) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VPNPhase1Policy) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCreated

`func (o *VPNPhase1Policy) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VPNPhase1Policy) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VPNPhase1Policy) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VPNPhase1Policy) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VPNPhase1Policy) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VPNPhase1Policy) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VPNPhase1Policy) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VPNPhase1Policy) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VPNPhase1Policy) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VPNPhase1Policy) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *VPNPhase1Policy) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VPNPhase1Policy) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VPNPhase1Policy) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *VPNPhase1Policy) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNPhase1Policy) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNPhase1Policy) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNPhase1Policy) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *VPNPhase1Policy) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VPNPhase1Policy) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VPNPhase1Policy) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VPNPhase1Policy) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


