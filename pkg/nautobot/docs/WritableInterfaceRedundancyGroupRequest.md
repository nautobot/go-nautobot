# WritableInterfaceRedundancyGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to [**RedundancyProtocol**](RedundancyProtocol.md) |  | [optional] 
**ProtocolGroupId** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**SecretsGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**VirtualIp** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewWritableInterfaceRedundancyGroupRequest

`func NewWritableInterfaceRedundancyGroupRequest(name string, status BulkWritableCableRequestStatus, ) *WritableInterfaceRedundancyGroupRequest`

NewWritableInterfaceRedundancyGroupRequest instantiates a new WritableInterfaceRedundancyGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableInterfaceRedundancyGroupRequestWithDefaults

`func NewWritableInterfaceRedundancyGroupRequestWithDefaults() *WritableInterfaceRedundancyGroupRequest`

NewWritableInterfaceRedundancyGroupRequestWithDefaults instantiates a new WritableInterfaceRedundancyGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableInterfaceRedundancyGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableInterfaceRedundancyGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableInterfaceRedundancyGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WritableInterfaceRedundancyGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableInterfaceRedundancyGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableInterfaceRedundancyGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableInterfaceRedundancyGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProtocol

`func (o *WritableInterfaceRedundancyGroupRequest) GetProtocol() RedundancyProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *WritableInterfaceRedundancyGroupRequest) GetProtocolOk() (*RedundancyProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *WritableInterfaceRedundancyGroupRequest) SetProtocol(v RedundancyProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *WritableInterfaceRedundancyGroupRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProtocolGroupId

`func (o *WritableInterfaceRedundancyGroupRequest) GetProtocolGroupId() string`

GetProtocolGroupId returns the ProtocolGroupId field if non-nil, zero value otherwise.

### GetProtocolGroupIdOk

`func (o *WritableInterfaceRedundancyGroupRequest) GetProtocolGroupIdOk() (*string, bool)`

GetProtocolGroupIdOk returns a tuple with the ProtocolGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolGroupId

`func (o *WritableInterfaceRedundancyGroupRequest) SetProtocolGroupId(v string)`

SetProtocolGroupId sets ProtocolGroupId field to given value.

### HasProtocolGroupId

`func (o *WritableInterfaceRedundancyGroupRequest) HasProtocolGroupId() bool`

HasProtocolGroupId returns a boolean if a field has been set.

### GetStatus

`func (o *WritableInterfaceRedundancyGroupRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableInterfaceRedundancyGroupRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableInterfaceRedundancyGroupRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetSecretsGroup

`func (o *WritableInterfaceRedundancyGroupRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *WritableInterfaceRedundancyGroupRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *WritableInterfaceRedundancyGroupRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *WritableInterfaceRedundancyGroupRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *WritableInterfaceRedundancyGroupRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *WritableInterfaceRedundancyGroupRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetVirtualIp

`func (o *WritableInterfaceRedundancyGroupRequest) GetVirtualIp() BulkWritableCircuitRequestTenant`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *WritableInterfaceRedundancyGroupRequest) GetVirtualIpOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *WritableInterfaceRedundancyGroupRequest) SetVirtualIp(v BulkWritableCircuitRequestTenant)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *WritableInterfaceRedundancyGroupRequest) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### SetVirtualIpNil

`func (o *WritableInterfaceRedundancyGroupRequest) SetVirtualIpNil(b bool)`

 SetVirtualIpNil sets the value for VirtualIp to be an explicit nil

### UnsetVirtualIp
`func (o *WritableInterfaceRedundancyGroupRequest) UnsetVirtualIp()`

UnsetVirtualIp ensures that no value is present for VirtualIp, not even an explicit nil
### GetTags

`func (o *WritableInterfaceRedundancyGroupRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableInterfaceRedundancyGroupRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableInterfaceRedundancyGroupRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableInterfaceRedundancyGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableInterfaceRedundancyGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableInterfaceRedundancyGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableInterfaceRedundancyGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableInterfaceRedundancyGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritableInterfaceRedundancyGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritableInterfaceRedundancyGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritableInterfaceRedundancyGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritableInterfaceRedundancyGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


