# BulkWritableWirelessNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Ssid** | **string** |  | 
**Mode** | [**WirelessNetworkModeChoices**](WirelessNetworkModeChoices.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Authentication** | [**AuthenticationEnum**](AuthenticationEnum.md) |  | 
**Hidden** | Pointer to **bool** |  | [optional] 
**SecretsGroup** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewBulkWritableWirelessNetworkRequest

`func NewBulkWritableWirelessNetworkRequest(id string, name string, ssid string, mode WirelessNetworkModeChoices, authentication AuthenticationEnum, ) *BulkWritableWirelessNetworkRequest`

NewBulkWritableWirelessNetworkRequest instantiates a new BulkWritableWirelessNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableWirelessNetworkRequestWithDefaults

`func NewBulkWritableWirelessNetworkRequestWithDefaults() *BulkWritableWirelessNetworkRequest`

NewBulkWritableWirelessNetworkRequestWithDefaults instantiates a new BulkWritableWirelessNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableWirelessNetworkRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableWirelessNetworkRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableWirelessNetworkRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableWirelessNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableWirelessNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableWirelessNetworkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableWirelessNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableWirelessNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableWirelessNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableWirelessNetworkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSsid

`func (o *BulkWritableWirelessNetworkRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *BulkWritableWirelessNetworkRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *BulkWritableWirelessNetworkRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetMode

`func (o *BulkWritableWirelessNetworkRequest) GetMode() WirelessNetworkModeChoices`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BulkWritableWirelessNetworkRequest) GetModeOk() (*WirelessNetworkModeChoices, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BulkWritableWirelessNetworkRequest) SetMode(v WirelessNetworkModeChoices)`

SetMode sets Mode field to given value.


### GetEnabled

`func (o *BulkWritableWirelessNetworkRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkWritableWirelessNetworkRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkWritableWirelessNetworkRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkWritableWirelessNetworkRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthentication

`func (o *BulkWritableWirelessNetworkRequest) GetAuthentication() AuthenticationEnum`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *BulkWritableWirelessNetworkRequest) GetAuthenticationOk() (*AuthenticationEnum, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *BulkWritableWirelessNetworkRequest) SetAuthentication(v AuthenticationEnum)`

SetAuthentication sets Authentication field to given value.


### GetHidden

`func (o *BulkWritableWirelessNetworkRequest) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *BulkWritableWirelessNetworkRequest) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *BulkWritableWirelessNetworkRequest) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *BulkWritableWirelessNetworkRequest) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *BulkWritableWirelessNetworkRequest) GetSecretsGroup() ApprovalWorkflowUser`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *BulkWritableWirelessNetworkRequest) GetSecretsGroupOk() (*ApprovalWorkflowUser, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *BulkWritableWirelessNetworkRequest) SetSecretsGroup(v ApprovalWorkflowUser)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *BulkWritableWirelessNetworkRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *BulkWritableWirelessNetworkRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *BulkWritableWirelessNetworkRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetTenant

`func (o *BulkWritableWirelessNetworkRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableWirelessNetworkRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableWirelessNetworkRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableWirelessNetworkRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableWirelessNetworkRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableWirelessNetworkRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableWirelessNetworkRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableWirelessNetworkRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableWirelessNetworkRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableWirelessNetworkRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableWirelessNetworkRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableWirelessNetworkRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableWirelessNetworkRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableWirelessNetworkRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableWirelessNetworkRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableWirelessNetworkRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableWirelessNetworkRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableWirelessNetworkRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


