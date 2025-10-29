# PatchedBulkWritableWirelessNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Ssid** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**WirelessNetworkModeChoices**](WirelessNetworkModeChoices.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Authentication** | Pointer to [**AuthenticationEnum**](AuthenticationEnum.md) |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**SecretsGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableWirelessNetworkRequest

`func NewPatchedBulkWritableWirelessNetworkRequest(id string, ) *PatchedBulkWritableWirelessNetworkRequest`

NewPatchedBulkWritableWirelessNetworkRequest instantiates a new PatchedBulkWritableWirelessNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableWirelessNetworkRequestWithDefaults

`func NewPatchedBulkWritableWirelessNetworkRequestWithDefaults() *PatchedBulkWritableWirelessNetworkRequest`

NewPatchedBulkWritableWirelessNetworkRequestWithDefaults instantiates a new PatchedBulkWritableWirelessNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSsid

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetMode

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetMode() WirelessNetworkModeChoices`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetModeOk() (*WirelessNetworkModeChoices, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetMode(v WirelessNetworkModeChoices)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthentication

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetAuthentication() AuthenticationEnum`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetAuthenticationOk() (*AuthenticationEnum, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetAuthentication(v AuthenticationEnum)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetHidden

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedBulkWritableWirelessNetworkRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetTenant

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableWirelessNetworkRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableWirelessNetworkRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableWirelessNetworkRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableWirelessNetworkRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


