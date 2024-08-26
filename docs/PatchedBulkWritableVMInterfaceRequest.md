# PatchedBulkWritableVMInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Mode** | Pointer to [**ModeEnum**](ModeEnum.md) |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**ParentInterface** | Pointer to [**NullableBulkWritableInterfaceRequestParentInterface**](BulkWritableInterfaceRequestParentInterface.md) |  | [optional] 
**Bridge** | Pointer to [**NullableBridgeInterface**](BridgeInterface.md) |  | [optional] 
**VirtualMachine** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**UntaggedVlan** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Vrf** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**TaggedVlans** | Pointer to [**[]TaggedVLANs**](TaggedVLANs.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVMInterfaceRequest

`func NewPatchedBulkWritableVMInterfaceRequest(id string, ) *PatchedBulkWritableVMInterfaceRequest`

NewPatchedBulkWritableVMInterfaceRequest instantiates a new PatchedBulkWritableVMInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVMInterfaceRequestWithDefaults

`func NewPatchedBulkWritableVMInterfaceRequestWithDefaults() *PatchedBulkWritableVMInterfaceRequest`

NewPatchedBulkWritableVMInterfaceRequestWithDefaults instantiates a new PatchedBulkWritableVMInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVMInterfaceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVMInterfaceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMode

`func (o *PatchedBulkWritableVMInterfaceRequest) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedBulkWritableVMInterfaceRequest) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedBulkWritableVMInterfaceRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMacAddress

`func (o *PatchedBulkWritableVMInterfaceRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *PatchedBulkWritableVMInterfaceRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *PatchedBulkWritableVMInterfaceRequest) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *PatchedBulkWritableVMInterfaceRequest) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *PatchedBulkWritableVMInterfaceRequest) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetEnabled

`func (o *PatchedBulkWritableVMInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkWritableVMInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkWritableVMInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMtu

`func (o *PatchedBulkWritableVMInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PatchedBulkWritableVMInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *PatchedBulkWritableVMInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *PatchedBulkWritableVMInterfaceRequest) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *PatchedBulkWritableVMInterfaceRequest) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetName

`func (o *PatchedBulkWritableVMInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableVMInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableVMInterfaceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableVMInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableVMInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableVMInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkWritableVMInterfaceRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableVMInterfaceRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableVMInterfaceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParentInterface

`func (o *PatchedBulkWritableVMInterfaceRequest) GetParentInterface() BulkWritableInterfaceRequestParentInterface`

GetParentInterface returns the ParentInterface field if non-nil, zero value otherwise.

### GetParentInterfaceOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetParentInterfaceOk() (*BulkWritableInterfaceRequestParentInterface, bool)`

GetParentInterfaceOk returns a tuple with the ParentInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInterface

`func (o *PatchedBulkWritableVMInterfaceRequest) SetParentInterface(v BulkWritableInterfaceRequestParentInterface)`

SetParentInterface sets ParentInterface field to given value.

### HasParentInterface

`func (o *PatchedBulkWritableVMInterfaceRequest) HasParentInterface() bool`

HasParentInterface returns a boolean if a field has been set.

### SetParentInterfaceNil

`func (o *PatchedBulkWritableVMInterfaceRequest) SetParentInterfaceNil(b bool)`

 SetParentInterfaceNil sets the value for ParentInterface to be an explicit nil

### UnsetParentInterface
`func (o *PatchedBulkWritableVMInterfaceRequest) UnsetParentInterface()`

UnsetParentInterface ensures that no value is present for ParentInterface, not even an explicit nil
### GetBridge

`func (o *PatchedBulkWritableVMInterfaceRequest) GetBridge() BridgeInterface`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetBridgeOk() (*BridgeInterface, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *PatchedBulkWritableVMInterfaceRequest) SetBridge(v BridgeInterface)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *PatchedBulkWritableVMInterfaceRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *PatchedBulkWritableVMInterfaceRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *PatchedBulkWritableVMInterfaceRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetVirtualMachine

`func (o *PatchedBulkWritableVMInterfaceRequest) GetVirtualMachine() BulkWritableCableRequestStatus`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetVirtualMachineOk() (*BulkWritableCableRequestStatus, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *PatchedBulkWritableVMInterfaceRequest) SetVirtualMachine(v BulkWritableCableRequestStatus)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *PatchedBulkWritableVMInterfaceRequest) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### GetUntaggedVlan

`func (o *PatchedBulkWritableVMInterfaceRequest) GetUntaggedVlan() BulkWritableCircuitRequestTenant`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetUntaggedVlanOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *PatchedBulkWritableVMInterfaceRequest) SetUntaggedVlan(v BulkWritableCircuitRequestTenant)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *PatchedBulkWritableVMInterfaceRequest) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *PatchedBulkWritableVMInterfaceRequest) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *PatchedBulkWritableVMInterfaceRequest) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetVrf

`func (o *PatchedBulkWritableVMInterfaceRequest) GetVrf() BulkWritableCircuitRequestTenant`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetVrfOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedBulkWritableVMInterfaceRequest) SetVrf(v BulkWritableCircuitRequestTenant)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedBulkWritableVMInterfaceRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *PatchedBulkWritableVMInterfaceRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *PatchedBulkWritableVMInterfaceRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTaggedVlans

`func (o *PatchedBulkWritableVMInterfaceRequest) GetTaggedVlans() []TaggedVLANs`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetTaggedVlansOk() (*[]TaggedVLANs, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *PatchedBulkWritableVMInterfaceRequest) SetTaggedVlans(v []TaggedVLANs)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *PatchedBulkWritableVMInterfaceRequest) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableVMInterfaceRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableVMInterfaceRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableVMInterfaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableVMInterfaceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableVMInterfaceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableVMInterfaceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableVMInterfaceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableVMInterfaceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableVMInterfaceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableVMInterfaceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


