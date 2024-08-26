# PatchedWritableInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**Mode** | Pointer to [**PatchedWritableInterfaceRequestMode**](PatchedWritableInterfaceRequestMode.md) |  | [optional] 
**Type** | Pointer to [**InterfaceTypeChoices**](InterfaceTypeChoices.md) |  | [optional] 
**MgmtOnly** | Pointer to **bool** | This interface is used only for out-of-band management | [optional] 
**Device** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**ParentInterface** | Pointer to [**NullableBulkWritableInterfaceRequestParentInterface**](BulkWritableInterfaceRequestParentInterface.md) |  | [optional] 
**Bridge** | Pointer to [**NullableBridgeInterface**](BridgeInterface.md) |  | [optional] 
**Lag** | Pointer to [**NullableParentLAG**](ParentLAG.md) |  | [optional] 
**UntaggedVlan** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Vrf** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**TaggedVlans** | Pointer to [**[]TaggedVLANs**](TaggedVLANs.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedWritableInterfaceRequest

`func NewPatchedWritableInterfaceRequest() *PatchedWritableInterfaceRequest`

NewPatchedWritableInterfaceRequest instantiates a new PatchedWritableInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableInterfaceRequestWithDefaults

`func NewPatchedWritableInterfaceRequestWithDefaults() *PatchedWritableInterfaceRequest`

NewPatchedWritableInterfaceRequestWithDefaults instantiates a new PatchedWritableInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *PatchedWritableInterfaceRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *PatchedWritableInterfaceRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *PatchedWritableInterfaceRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *PatchedWritableInterfaceRequest) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *PatchedWritableInterfaceRequest) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *PatchedWritableInterfaceRequest) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetName

`func (o *PatchedWritableInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableInterfaceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableInterfaceRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableInterfaceRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableInterfaceRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableInterfaceRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedWritableInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedWritableInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedWritableInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedWritableInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMtu

`func (o *PatchedWritableInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PatchedWritableInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PatchedWritableInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *PatchedWritableInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *PatchedWritableInterfaceRequest) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *PatchedWritableInterfaceRequest) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetMode

`func (o *PatchedWritableInterfaceRequest) GetMode() PatchedWritableInterfaceRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedWritableInterfaceRequest) GetModeOk() (*PatchedWritableInterfaceRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedWritableInterfaceRequest) SetMode(v PatchedWritableInterfaceRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedWritableInterfaceRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableInterfaceRequest) GetType() InterfaceTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableInterfaceRequest) GetTypeOk() (*InterfaceTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableInterfaceRequest) SetType(v InterfaceTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableInterfaceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMgmtOnly

`func (o *PatchedWritableInterfaceRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *PatchedWritableInterfaceRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *PatchedWritableInterfaceRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *PatchedWritableInterfaceRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDevice

`func (o *PatchedWritableInterfaceRequest) GetDevice() BulkWritableCableRequestStatus`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedWritableInterfaceRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedWritableInterfaceRequest) SetDevice(v BulkWritableCableRequestStatus)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedWritableInterfaceRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableInterfaceRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableInterfaceRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableInterfaceRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableInterfaceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParentInterface

`func (o *PatchedWritableInterfaceRequest) GetParentInterface() BulkWritableInterfaceRequestParentInterface`

GetParentInterface returns the ParentInterface field if non-nil, zero value otherwise.

### GetParentInterfaceOk

`func (o *PatchedWritableInterfaceRequest) GetParentInterfaceOk() (*BulkWritableInterfaceRequestParentInterface, bool)`

GetParentInterfaceOk returns a tuple with the ParentInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInterface

`func (o *PatchedWritableInterfaceRequest) SetParentInterface(v BulkWritableInterfaceRequestParentInterface)`

SetParentInterface sets ParentInterface field to given value.

### HasParentInterface

`func (o *PatchedWritableInterfaceRequest) HasParentInterface() bool`

HasParentInterface returns a boolean if a field has been set.

### SetParentInterfaceNil

`func (o *PatchedWritableInterfaceRequest) SetParentInterfaceNil(b bool)`

 SetParentInterfaceNil sets the value for ParentInterface to be an explicit nil

### UnsetParentInterface
`func (o *PatchedWritableInterfaceRequest) UnsetParentInterface()`

UnsetParentInterface ensures that no value is present for ParentInterface, not even an explicit nil
### GetBridge

`func (o *PatchedWritableInterfaceRequest) GetBridge() BridgeInterface`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *PatchedWritableInterfaceRequest) GetBridgeOk() (*BridgeInterface, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *PatchedWritableInterfaceRequest) SetBridge(v BridgeInterface)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *PatchedWritableInterfaceRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *PatchedWritableInterfaceRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *PatchedWritableInterfaceRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetLag

`func (o *PatchedWritableInterfaceRequest) GetLag() ParentLAG`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *PatchedWritableInterfaceRequest) GetLagOk() (*ParentLAG, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *PatchedWritableInterfaceRequest) SetLag(v ParentLAG)`

SetLag sets Lag field to given value.

### HasLag

`func (o *PatchedWritableInterfaceRequest) HasLag() bool`

HasLag returns a boolean if a field has been set.

### SetLagNil

`func (o *PatchedWritableInterfaceRequest) SetLagNil(b bool)`

 SetLagNil sets the value for Lag to be an explicit nil

### UnsetLag
`func (o *PatchedWritableInterfaceRequest) UnsetLag()`

UnsetLag ensures that no value is present for Lag, not even an explicit nil
### GetUntaggedVlan

`func (o *PatchedWritableInterfaceRequest) GetUntaggedVlan() BulkWritableCircuitRequestTenant`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *PatchedWritableInterfaceRequest) GetUntaggedVlanOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *PatchedWritableInterfaceRequest) SetUntaggedVlan(v BulkWritableCircuitRequestTenant)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *PatchedWritableInterfaceRequest) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *PatchedWritableInterfaceRequest) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *PatchedWritableInterfaceRequest) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetVrf

`func (o *PatchedWritableInterfaceRequest) GetVrf() BulkWritableCircuitRequestTenant`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedWritableInterfaceRequest) GetVrfOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedWritableInterfaceRequest) SetVrf(v BulkWritableCircuitRequestTenant)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedWritableInterfaceRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *PatchedWritableInterfaceRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *PatchedWritableInterfaceRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTaggedVlans

`func (o *PatchedWritableInterfaceRequest) GetTaggedVlans() []TaggedVLANs`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *PatchedWritableInterfaceRequest) GetTaggedVlansOk() (*[]TaggedVLANs, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *PatchedWritableInterfaceRequest) SetTaggedVlans(v []TaggedVLANs)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *PatchedWritableInterfaceRequest) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableInterfaceRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableInterfaceRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableInterfaceRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableInterfaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableInterfaceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableInterfaceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableInterfaceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableInterfaceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritableInterfaceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritableInterfaceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritableInterfaceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritableInterfaceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


