# PatchedWritableInterfaceRedundancyGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to [**RedundancyProtocol**](RedundancyProtocol.md) |  | [optional] 
**ProtocolGroupId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**SecretsGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**VirtualIp** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedWritableInterfaceRedundancyGroupRequest

`func NewPatchedWritableInterfaceRedundancyGroupRequest() *PatchedWritableInterfaceRedundancyGroupRequest`

NewPatchedWritableInterfaceRedundancyGroupRequest instantiates a new PatchedWritableInterfaceRedundancyGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableInterfaceRedundancyGroupRequestWithDefaults

`func NewPatchedWritableInterfaceRedundancyGroupRequestWithDefaults() *PatchedWritableInterfaceRedundancyGroupRequest`

NewPatchedWritableInterfaceRedundancyGroupRequestWithDefaults instantiates a new PatchedWritableInterfaceRedundancyGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetProtocol() RedundancyProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetProtocolOk() (*RedundancyProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetProtocol(v RedundancyProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProtocolGroupId

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetProtocolGroupId() string`

GetProtocolGroupId returns the ProtocolGroupId field if non-nil, zero value otherwise.

### GetProtocolGroupIdOk

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetProtocolGroupIdOk() (*string, bool)`

GetProtocolGroupIdOk returns a tuple with the ProtocolGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolGroupId

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetProtocolGroupId(v string)`

SetProtocolGroupId sets ProtocolGroupId field to given value.

### HasProtocolGroupId

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasProtocolGroupId() bool`

HasProtocolGroupId returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedWritableInterfaceRedundancyGroupRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetVirtualIp

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetVirtualIp() BulkWritableCircuitRequestTenant`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetVirtualIpOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetVirtualIp(v BulkWritableCircuitRequestTenant)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### SetVirtualIpNil

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetVirtualIpNil(b bool)`

 SetVirtualIpNil sets the value for VirtualIp to be an explicit nil

### UnsetVirtualIp
`func (o *PatchedWritableInterfaceRedundancyGroupRequest) UnsetVirtualIp()`

UnsetVirtualIp ensures that no value is present for VirtualIp, not even an explicit nil
### GetCustomFields

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


