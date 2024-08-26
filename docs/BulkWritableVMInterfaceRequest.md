# BulkWritableVMInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Mode** | Pointer to [**ModeEnum**](ModeEnum.md) |  | [optional] 
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
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableVMInterfaceRequest

`func NewBulkWritableVMInterfaceRequest(id string, name string, status BulkWritableCableRequestStatus, virtualMachine BulkWritableCableRequestStatus, ) *BulkWritableVMInterfaceRequest`

NewBulkWritableVMInterfaceRequest instantiates a new BulkWritableVMInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableVMInterfaceRequestWithDefaults

`func NewBulkWritableVMInterfaceRequestWithDefaults() *BulkWritableVMInterfaceRequest`

NewBulkWritableVMInterfaceRequestWithDefaults instantiates a new BulkWritableVMInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableVMInterfaceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableVMInterfaceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableVMInterfaceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetMode

`func (o *BulkWritableVMInterfaceRequest) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BulkWritableVMInterfaceRequest) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BulkWritableVMInterfaceRequest) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *BulkWritableVMInterfaceRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMacAddress

`func (o *BulkWritableVMInterfaceRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BulkWritableVMInterfaceRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BulkWritableVMInterfaceRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BulkWritableVMInterfaceRequest) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *BulkWritableVMInterfaceRequest) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *BulkWritableVMInterfaceRequest) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetEnabled

`func (o *BulkWritableVMInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkWritableVMInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkWritableVMInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkWritableVMInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMtu

`func (o *BulkWritableVMInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *BulkWritableVMInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *BulkWritableVMInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *BulkWritableVMInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *BulkWritableVMInterfaceRequest) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *BulkWritableVMInterfaceRequest) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetName

`func (o *BulkWritableVMInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableVMInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableVMInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableVMInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableVMInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableVMInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableVMInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *BulkWritableVMInterfaceRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableVMInterfaceRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableVMInterfaceRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetParentInterface

`func (o *BulkWritableVMInterfaceRequest) GetParentInterface() BulkWritableInterfaceRequestParentInterface`

GetParentInterface returns the ParentInterface field if non-nil, zero value otherwise.

### GetParentInterfaceOk

`func (o *BulkWritableVMInterfaceRequest) GetParentInterfaceOk() (*BulkWritableInterfaceRequestParentInterface, bool)`

GetParentInterfaceOk returns a tuple with the ParentInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInterface

`func (o *BulkWritableVMInterfaceRequest) SetParentInterface(v BulkWritableInterfaceRequestParentInterface)`

SetParentInterface sets ParentInterface field to given value.

### HasParentInterface

`func (o *BulkWritableVMInterfaceRequest) HasParentInterface() bool`

HasParentInterface returns a boolean if a field has been set.

### SetParentInterfaceNil

`func (o *BulkWritableVMInterfaceRequest) SetParentInterfaceNil(b bool)`

 SetParentInterfaceNil sets the value for ParentInterface to be an explicit nil

### UnsetParentInterface
`func (o *BulkWritableVMInterfaceRequest) UnsetParentInterface()`

UnsetParentInterface ensures that no value is present for ParentInterface, not even an explicit nil
### GetBridge

`func (o *BulkWritableVMInterfaceRequest) GetBridge() BridgeInterface`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *BulkWritableVMInterfaceRequest) GetBridgeOk() (*BridgeInterface, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *BulkWritableVMInterfaceRequest) SetBridge(v BridgeInterface)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *BulkWritableVMInterfaceRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *BulkWritableVMInterfaceRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *BulkWritableVMInterfaceRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetVirtualMachine

`func (o *BulkWritableVMInterfaceRequest) GetVirtualMachine() BulkWritableCableRequestStatus`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *BulkWritableVMInterfaceRequest) GetVirtualMachineOk() (*BulkWritableCableRequestStatus, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *BulkWritableVMInterfaceRequest) SetVirtualMachine(v BulkWritableCableRequestStatus)`

SetVirtualMachine sets VirtualMachine field to given value.


### GetUntaggedVlan

`func (o *BulkWritableVMInterfaceRequest) GetUntaggedVlan() BulkWritableCircuitRequestTenant`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *BulkWritableVMInterfaceRequest) GetUntaggedVlanOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *BulkWritableVMInterfaceRequest) SetUntaggedVlan(v BulkWritableCircuitRequestTenant)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *BulkWritableVMInterfaceRequest) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *BulkWritableVMInterfaceRequest) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *BulkWritableVMInterfaceRequest) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetVrf

`func (o *BulkWritableVMInterfaceRequest) GetVrf() BulkWritableCircuitRequestTenant`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *BulkWritableVMInterfaceRequest) GetVrfOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *BulkWritableVMInterfaceRequest) SetVrf(v BulkWritableCircuitRequestTenant)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *BulkWritableVMInterfaceRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *BulkWritableVMInterfaceRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *BulkWritableVMInterfaceRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTaggedVlans

`func (o *BulkWritableVMInterfaceRequest) GetTaggedVlans() []TaggedVLANs`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *BulkWritableVMInterfaceRequest) GetTaggedVlansOk() (*[]TaggedVLANs, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *BulkWritableVMInterfaceRequest) SetTaggedVlans(v []TaggedVLANs)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *BulkWritableVMInterfaceRequest) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableVMInterfaceRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableVMInterfaceRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableVMInterfaceRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableVMInterfaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableVMInterfaceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableVMInterfaceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableVMInterfaceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableVMInterfaceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableVMInterfaceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableVMInterfaceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableVMInterfaceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableVMInterfaceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


