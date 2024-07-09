# VMInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Mode** | Pointer to [**InterfaceMode**](InterfaceMode.md) |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**ParentInterface** | Pointer to [**NullableBulkWritableInterfaceRequestParentInterface**](BulkWritableInterfaceRequestParentInterface.md) |  | [optional] 
**Bridge** | Pointer to [**NullableBridgeInterface**](BridgeInterface.md) |  | [optional] 
**VirtualMachine** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**UntaggedVlan** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Vrf** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**TaggedVlans** | Pointer to [**[]TaggedVLANs**](TaggedVLANs.md) |  | [optional] 
**IpAddresses** | [**[]IPAddresses**](IPAddresses.md) |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVMInterface

`func NewVMInterface(id string, objectType string, display string, url string, naturalSlug string, name string, status BulkWritableCableRequestStatus, virtualMachine BulkWritableCableRequestStatus, ipAddresses []IPAddresses, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *VMInterface`

NewVMInterface instantiates a new VMInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInterfaceWithDefaults

`func NewVMInterfaceWithDefaults() *VMInterface`

NewVMInterfaceWithDefaults instantiates a new VMInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMInterface) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMInterface) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMInterface) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *VMInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VMInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VMInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VMInterface) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VMInterface) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VMInterface) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VMInterface) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VMInterface) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VMInterface) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VMInterface) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VMInterface) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VMInterface) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetMode

`func (o *VMInterface) GetMode() InterfaceMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VMInterface) GetModeOk() (*InterfaceMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VMInterface) SetMode(v InterfaceMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VMInterface) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMacAddress

`func (o *VMInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VMInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VMInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VMInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *VMInterface) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *VMInterface) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetEnabled

`func (o *VMInterface) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VMInterface) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VMInterface) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VMInterface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMtu

`func (o *VMInterface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VMInterface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VMInterface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VMInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *VMInterface) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *VMInterface) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetName

`func (o *VMInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMInterface) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VMInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VMInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VMInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VMInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *VMInterface) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VMInterface) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VMInterface) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetParentInterface

`func (o *VMInterface) GetParentInterface() BulkWritableInterfaceRequestParentInterface`

GetParentInterface returns the ParentInterface field if non-nil, zero value otherwise.

### GetParentInterfaceOk

`func (o *VMInterface) GetParentInterfaceOk() (*BulkWritableInterfaceRequestParentInterface, bool)`

GetParentInterfaceOk returns a tuple with the ParentInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInterface

`func (o *VMInterface) SetParentInterface(v BulkWritableInterfaceRequestParentInterface)`

SetParentInterface sets ParentInterface field to given value.

### HasParentInterface

`func (o *VMInterface) HasParentInterface() bool`

HasParentInterface returns a boolean if a field has been set.

### SetParentInterfaceNil

`func (o *VMInterface) SetParentInterfaceNil(b bool)`

 SetParentInterfaceNil sets the value for ParentInterface to be an explicit nil

### UnsetParentInterface
`func (o *VMInterface) UnsetParentInterface()`

UnsetParentInterface ensures that no value is present for ParentInterface, not even an explicit nil
### GetBridge

`func (o *VMInterface) GetBridge() BridgeInterface`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *VMInterface) GetBridgeOk() (*BridgeInterface, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *VMInterface) SetBridge(v BridgeInterface)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *VMInterface) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *VMInterface) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *VMInterface) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetVirtualMachine

`func (o *VMInterface) GetVirtualMachine() BulkWritableCableRequestStatus`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VMInterface) GetVirtualMachineOk() (*BulkWritableCableRequestStatus, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VMInterface) SetVirtualMachine(v BulkWritableCableRequestStatus)`

SetVirtualMachine sets VirtualMachine field to given value.


### GetUntaggedVlan

`func (o *VMInterface) GetUntaggedVlan() BulkWritableCircuitRequestTenant`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *VMInterface) GetUntaggedVlanOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *VMInterface) SetUntaggedVlan(v BulkWritableCircuitRequestTenant)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *VMInterface) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *VMInterface) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *VMInterface) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetVrf

`func (o *VMInterface) GetVrf() BulkWritableCircuitRequestTenant`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VMInterface) GetVrfOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VMInterface) SetVrf(v BulkWritableCircuitRequestTenant)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *VMInterface) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *VMInterface) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *VMInterface) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTaggedVlans

`func (o *VMInterface) GetTaggedVlans() []TaggedVLANs`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *VMInterface) GetTaggedVlansOk() (*[]TaggedVLANs, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *VMInterface) SetTaggedVlans(v []TaggedVLANs)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *VMInterface) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetIpAddresses

`func (o *VMInterface) GetIpAddresses() []IPAddresses`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *VMInterface) GetIpAddressesOk() (*[]IPAddresses, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *VMInterface) SetIpAddresses(v []IPAddresses)`

SetIpAddresses sets IpAddresses field to given value.


### GetCreated

`func (o *VMInterface) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VMInterface) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VMInterface) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VMInterface) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VMInterface) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VMInterface) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VMInterface) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VMInterface) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VMInterface) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VMInterface) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *VMInterface) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VMInterface) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VMInterface) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VMInterface) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *VMInterface) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VMInterface) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VMInterface) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *VMInterface) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VMInterface) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VMInterface) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VMInterface) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


