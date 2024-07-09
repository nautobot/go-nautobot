# InterfaceRedundancyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Protocol** | [**InterfaceRedundancyGroupProtocol**](InterfaceRedundancyGroupProtocol.md) |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ProtocolGroupId** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**SecretsGroup** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**VirtualIp** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Interfaces** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInterfaceRedundancyGroup

`func NewInterfaceRedundancyGroup(id string, objectType string, display string, url string, naturalSlug string, protocol InterfaceRedundancyGroupProtocol, name string, status BulkWritableCableRequestStatus, interfaces []BulkWritableCableRequestStatus, created time.Time, lastUpdated NullableTime, notesUrl string, ) *InterfaceRedundancyGroup`

NewInterfaceRedundancyGroup instantiates a new InterfaceRedundancyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceRedundancyGroupWithDefaults

`func NewInterfaceRedundancyGroupWithDefaults() *InterfaceRedundancyGroup`

NewInterfaceRedundancyGroupWithDefaults instantiates a new InterfaceRedundancyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterfaceRedundancyGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterfaceRedundancyGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterfaceRedundancyGroup) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *InterfaceRedundancyGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InterfaceRedundancyGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InterfaceRedundancyGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *InterfaceRedundancyGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InterfaceRedundancyGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InterfaceRedundancyGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *InterfaceRedundancyGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InterfaceRedundancyGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InterfaceRedundancyGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *InterfaceRedundancyGroup) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *InterfaceRedundancyGroup) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *InterfaceRedundancyGroup) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetProtocol

`func (o *InterfaceRedundancyGroup) GetProtocol() InterfaceRedundancyGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InterfaceRedundancyGroup) GetProtocolOk() (*InterfaceRedundancyGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InterfaceRedundancyGroup) SetProtocol(v InterfaceRedundancyGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetName

`func (o *InterfaceRedundancyGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceRedundancyGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceRedundancyGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InterfaceRedundancyGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterfaceRedundancyGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterfaceRedundancyGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterfaceRedundancyGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProtocolGroupId

`func (o *InterfaceRedundancyGroup) GetProtocolGroupId() string`

GetProtocolGroupId returns the ProtocolGroupId field if non-nil, zero value otherwise.

### GetProtocolGroupIdOk

`func (o *InterfaceRedundancyGroup) GetProtocolGroupIdOk() (*string, bool)`

GetProtocolGroupIdOk returns a tuple with the ProtocolGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolGroupId

`func (o *InterfaceRedundancyGroup) SetProtocolGroupId(v string)`

SetProtocolGroupId sets ProtocolGroupId field to given value.

### HasProtocolGroupId

`func (o *InterfaceRedundancyGroup) HasProtocolGroupId() bool`

HasProtocolGroupId returns a boolean if a field has been set.

### GetStatus

`func (o *InterfaceRedundancyGroup) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InterfaceRedundancyGroup) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InterfaceRedundancyGroup) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetSecretsGroup

`func (o *InterfaceRedundancyGroup) GetSecretsGroup() BulkWritableCircuitRequestTenant`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *InterfaceRedundancyGroup) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *InterfaceRedundancyGroup) SetSecretsGroup(v BulkWritableCircuitRequestTenant)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *InterfaceRedundancyGroup) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *InterfaceRedundancyGroup) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *InterfaceRedundancyGroup) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetVirtualIp

`func (o *InterfaceRedundancyGroup) GetVirtualIp() BulkWritableCircuitRequestTenant`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *InterfaceRedundancyGroup) GetVirtualIpOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *InterfaceRedundancyGroup) SetVirtualIp(v BulkWritableCircuitRequestTenant)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *InterfaceRedundancyGroup) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### SetVirtualIpNil

`func (o *InterfaceRedundancyGroup) SetVirtualIpNil(b bool)`

 SetVirtualIpNil sets the value for VirtualIp to be an explicit nil

### UnsetVirtualIp
`func (o *InterfaceRedundancyGroup) UnsetVirtualIp()`

UnsetVirtualIp ensures that no value is present for VirtualIp, not even an explicit nil
### GetInterfaces

`func (o *InterfaceRedundancyGroup) GetInterfaces() []BulkWritableCableRequestStatus`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InterfaceRedundancyGroup) GetInterfacesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InterfaceRedundancyGroup) SetInterfaces(v []BulkWritableCableRequestStatus)`

SetInterfaces sets Interfaces field to given value.


### GetCreated

`func (o *InterfaceRedundancyGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InterfaceRedundancyGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InterfaceRedundancyGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *InterfaceRedundancyGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InterfaceRedundancyGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InterfaceRedundancyGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *InterfaceRedundancyGroup) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *InterfaceRedundancyGroup) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *InterfaceRedundancyGroup) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InterfaceRedundancyGroup) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InterfaceRedundancyGroup) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InterfaceRedundancyGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *InterfaceRedundancyGroup) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *InterfaceRedundancyGroup) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *InterfaceRedundancyGroup) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *InterfaceRedundancyGroup) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InterfaceRedundancyGroup) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InterfaceRedundancyGroup) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InterfaceRedundancyGroup) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


