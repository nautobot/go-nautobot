# WritableInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**Mode** | Pointer to [**PatchedWritableInterfaceRequestMode**](PatchedWritableInterfaceRequestMode.md) |  | [optional] 
**Type** | [**InterfaceTypeChoices**](InterfaceTypeChoices.md) |  | 
**MgmtOnly** | Pointer to **bool** | This interface is used only for out-of-band management | [optional] 
**Device** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
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

### NewWritableInterfaceRequest

`func NewWritableInterfaceRequest(name string, type_ InterfaceTypeChoices, device BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, ) *WritableInterfaceRequest`

NewWritableInterfaceRequest instantiates a new WritableInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableInterfaceRequestWithDefaults

`func NewWritableInterfaceRequestWithDefaults() *WritableInterfaceRequest`

NewWritableInterfaceRequestWithDefaults instantiates a new WritableInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *WritableInterfaceRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *WritableInterfaceRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *WritableInterfaceRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *WritableInterfaceRequest) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *WritableInterfaceRequest) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *WritableInterfaceRequest) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetName

`func (o *WritableInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritableInterfaceRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableInterfaceRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableInterfaceRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableInterfaceRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *WritableInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *WritableInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WritableInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WritableInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WritableInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMtu

`func (o *WritableInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *WritableInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *WritableInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *WritableInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *WritableInterfaceRequest) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *WritableInterfaceRequest) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetMode

`func (o *WritableInterfaceRequest) GetMode() PatchedWritableInterfaceRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WritableInterfaceRequest) GetModeOk() (*PatchedWritableInterfaceRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WritableInterfaceRequest) SetMode(v PatchedWritableInterfaceRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WritableInterfaceRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetType

`func (o *WritableInterfaceRequest) GetType() InterfaceTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableInterfaceRequest) GetTypeOk() (*InterfaceTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableInterfaceRequest) SetType(v InterfaceTypeChoices)`

SetType sets Type field to given value.


### GetMgmtOnly

`func (o *WritableInterfaceRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *WritableInterfaceRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *WritableInterfaceRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *WritableInterfaceRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDevice

`func (o *WritableInterfaceRequest) GetDevice() BulkWritableCableRequestStatus`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritableInterfaceRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritableInterfaceRequest) SetDevice(v BulkWritableCableRequestStatus)`

SetDevice sets Device field to given value.


### GetStatus

`func (o *WritableInterfaceRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableInterfaceRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableInterfaceRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetParentInterface

`func (o *WritableInterfaceRequest) GetParentInterface() BulkWritableInterfaceRequestParentInterface`

GetParentInterface returns the ParentInterface field if non-nil, zero value otherwise.

### GetParentInterfaceOk

`func (o *WritableInterfaceRequest) GetParentInterfaceOk() (*BulkWritableInterfaceRequestParentInterface, bool)`

GetParentInterfaceOk returns a tuple with the ParentInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInterface

`func (o *WritableInterfaceRequest) SetParentInterface(v BulkWritableInterfaceRequestParentInterface)`

SetParentInterface sets ParentInterface field to given value.

### HasParentInterface

`func (o *WritableInterfaceRequest) HasParentInterface() bool`

HasParentInterface returns a boolean if a field has been set.

### SetParentInterfaceNil

`func (o *WritableInterfaceRequest) SetParentInterfaceNil(b bool)`

 SetParentInterfaceNil sets the value for ParentInterface to be an explicit nil

### UnsetParentInterface
`func (o *WritableInterfaceRequest) UnsetParentInterface()`

UnsetParentInterface ensures that no value is present for ParentInterface, not even an explicit nil
### GetBridge

`func (o *WritableInterfaceRequest) GetBridge() BridgeInterface`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *WritableInterfaceRequest) GetBridgeOk() (*BridgeInterface, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *WritableInterfaceRequest) SetBridge(v BridgeInterface)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *WritableInterfaceRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *WritableInterfaceRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *WritableInterfaceRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetLag

`func (o *WritableInterfaceRequest) GetLag() ParentLAG`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *WritableInterfaceRequest) GetLagOk() (*ParentLAG, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *WritableInterfaceRequest) SetLag(v ParentLAG)`

SetLag sets Lag field to given value.

### HasLag

`func (o *WritableInterfaceRequest) HasLag() bool`

HasLag returns a boolean if a field has been set.

### SetLagNil

`func (o *WritableInterfaceRequest) SetLagNil(b bool)`

 SetLagNil sets the value for Lag to be an explicit nil

### UnsetLag
`func (o *WritableInterfaceRequest) UnsetLag()`

UnsetLag ensures that no value is present for Lag, not even an explicit nil
### GetUntaggedVlan

`func (o *WritableInterfaceRequest) GetUntaggedVlan() BulkWritableCircuitRequestTenant`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *WritableInterfaceRequest) GetUntaggedVlanOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *WritableInterfaceRequest) SetUntaggedVlan(v BulkWritableCircuitRequestTenant)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *WritableInterfaceRequest) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *WritableInterfaceRequest) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *WritableInterfaceRequest) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetVrf

`func (o *WritableInterfaceRequest) GetVrf() BulkWritableCircuitRequestTenant`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *WritableInterfaceRequest) GetVrfOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *WritableInterfaceRequest) SetVrf(v BulkWritableCircuitRequestTenant)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *WritableInterfaceRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *WritableInterfaceRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *WritableInterfaceRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTaggedVlans

`func (o *WritableInterfaceRequest) GetTaggedVlans() []TaggedVLANs`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *WritableInterfaceRequest) GetTaggedVlansOk() (*[]TaggedVLANs, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *WritableInterfaceRequest) SetTaggedVlans(v []TaggedVLANs)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *WritableInterfaceRequest) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetTags

`func (o *WritableInterfaceRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableInterfaceRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableInterfaceRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableInterfaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableInterfaceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableInterfaceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableInterfaceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableInterfaceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritableInterfaceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritableInterfaceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritableInterfaceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritableInterfaceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


