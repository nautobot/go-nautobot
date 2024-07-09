/*
API Documentation

Source of truth and network automation platform

API version: 2.2.5 (2.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the BulkWritableInterfaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableInterfaceRequest{}

// BulkWritableInterfaceRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableInterfaceRequest struct {
	Id string `json:"id"`
	Type InterfaceTypeChoices `json:"type"`
	Mode *ModeEnum `json:"mode,omitempty"`
	MacAddress NullableString `json:"mac_address,omitempty"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Mtu NullableInt32 `json:"mtu,omitempty"`
	// This interface is used only for out-of-band management
	MgmtOnly *bool `json:"mgmt_only,omitempty"`
	Device BulkWritableCableRequestStatus `json:"device"`
	Status BulkWritableCableRequestStatus `json:"status"`
	ParentInterface NullableBulkWritableInterfaceRequestParentInterface `json:"parent_interface,omitempty"`
	Bridge NullableBridgeInterface `json:"bridge,omitempty"`
	Lag NullableParentLAG `json:"lag,omitempty"`
	UntaggedVlan NullableBulkWritableCircuitRequestTenant `json:"untagged_vlan,omitempty"`
	Vrf NullableBulkWritableCircuitRequestTenant `json:"vrf,omitempty"`
	TaggedVlans []TaggedVLANs `json:"tagged_vlans,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableInterfaceRequest BulkWritableInterfaceRequest

// NewBulkWritableInterfaceRequest instantiates a new BulkWritableInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableInterfaceRequest(id string, type_ InterfaceTypeChoices, name string, device BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus) *BulkWritableInterfaceRequest {
	this := BulkWritableInterfaceRequest{}
	this.Id = id
	this.Type = type_
	this.Name = name
	this.Device = device
	this.Status = status
	return &this
}

// NewBulkWritableInterfaceRequestWithDefaults instantiates a new BulkWritableInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableInterfaceRequestWithDefaults() *BulkWritableInterfaceRequest {
	this := BulkWritableInterfaceRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableInterfaceRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableInterfaceRequest) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *BulkWritableInterfaceRequest) GetType() InterfaceTypeChoices {
	if o == nil {
		var ret InterfaceTypeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetTypeOk() (*InterfaceTypeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BulkWritableInterfaceRequest) SetType(v InterfaceTypeChoices) {
	o.Type = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRequest) GetMode() ModeEnum {
	if o == nil || IsNil(o.Mode) {
		var ret ModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetModeOk() (*ModeEnum, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given ModeEnum and assigns it to the Mode field.
func (o *BulkWritableInterfaceRequest) SetMode(v ModeEnum) {
	o.Mode = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableInterfaceRequest) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress.Get()) {
		var ret string
		return ret
	}
	return *o.MacAddress.Get()
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableInterfaceRequest) GetMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacAddress.Get(), o.MacAddress.IsSet()
}

// HasMacAddress returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasMacAddress() bool {
	if o != nil && o.MacAddress.IsSet() {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given NullableString and assigns it to the MacAddress field.
func (o *BulkWritableInterfaceRequest) SetMacAddress(v string) {
	o.MacAddress.Set(&v)
}
// SetMacAddressNil sets the value for MacAddress to be an explicit nil
func (o *BulkWritableInterfaceRequest) SetMacAddressNil() {
	o.MacAddress.Set(nil)
}

// UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
func (o *BulkWritableInterfaceRequest) UnsetMacAddress() {
	o.MacAddress.Unset()
}

// GetName returns the Name field value
func (o *BulkWritableInterfaceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableInterfaceRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BulkWritableInterfaceRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableInterfaceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BulkWritableInterfaceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableInterfaceRequest) GetMtu() int32 {
	if o == nil || IsNil(o.Mtu.Get()) {
		var ret int32
		return ret
	}
	return *o.Mtu.Get()
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableInterfaceRequest) GetMtuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mtu.Get(), o.Mtu.IsSet()
}

// HasMtu returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasMtu() bool {
	if o != nil && o.Mtu.IsSet() {
		return true
	}

	return false
}

// SetMtu gets a reference to the given NullableInt32 and assigns it to the Mtu field.
func (o *BulkWritableInterfaceRequest) SetMtu(v int32) {
	o.Mtu.Set(&v)
}
// SetMtuNil sets the value for Mtu to be an explicit nil
func (o *BulkWritableInterfaceRequest) SetMtuNil() {
	o.Mtu.Set(nil)
}

// UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
func (o *BulkWritableInterfaceRequest) UnsetMtu() {
	o.Mtu.Unset()
}

// GetMgmtOnly returns the MgmtOnly field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRequest) GetMgmtOnly() bool {
	if o == nil || IsNil(o.MgmtOnly) {
		var ret bool
		return ret
	}
	return *o.MgmtOnly
}

// GetMgmtOnlyOk returns a tuple with the MgmtOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetMgmtOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.MgmtOnly) {
		return nil, false
	}
	return o.MgmtOnly, true
}

// HasMgmtOnly returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasMgmtOnly() bool {
	if o != nil && !IsNil(o.MgmtOnly) {
		return true
	}

	return false
}

// SetMgmtOnly gets a reference to the given bool and assigns it to the MgmtOnly field.
func (o *BulkWritableInterfaceRequest) SetMgmtOnly(v bool) {
	o.MgmtOnly = &v
}

// GetDevice returns the Device field value
func (o *BulkWritableInterfaceRequest) GetDevice() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *BulkWritableInterfaceRequest) SetDevice(v BulkWritableCableRequestStatus) {
	o.Device = v
}

// GetStatus returns the Status field value
func (o *BulkWritableInterfaceRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkWritableInterfaceRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetParentInterface returns the ParentInterface field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableInterfaceRequest) GetParentInterface() BulkWritableInterfaceRequestParentInterface {
	if o == nil || IsNil(o.ParentInterface.Get()) {
		var ret BulkWritableInterfaceRequestParentInterface
		return ret
	}
	return *o.ParentInterface.Get()
}

// GetParentInterfaceOk returns a tuple with the ParentInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableInterfaceRequest) GetParentInterfaceOk() (*BulkWritableInterfaceRequestParentInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentInterface.Get(), o.ParentInterface.IsSet()
}

// HasParentInterface returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasParentInterface() bool {
	if o != nil && o.ParentInterface.IsSet() {
		return true
	}

	return false
}

// SetParentInterface gets a reference to the given NullableBulkWritableInterfaceRequestParentInterface and assigns it to the ParentInterface field.
func (o *BulkWritableInterfaceRequest) SetParentInterface(v BulkWritableInterfaceRequestParentInterface) {
	o.ParentInterface.Set(&v)
}
// SetParentInterfaceNil sets the value for ParentInterface to be an explicit nil
func (o *BulkWritableInterfaceRequest) SetParentInterfaceNil() {
	o.ParentInterface.Set(nil)
}

// UnsetParentInterface ensures that no value is present for ParentInterface, not even an explicit nil
func (o *BulkWritableInterfaceRequest) UnsetParentInterface() {
	o.ParentInterface.Unset()
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableInterfaceRequest) GetBridge() BridgeInterface {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret BridgeInterface
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableInterfaceRequest) GetBridgeOk() (*BridgeInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableBridgeInterface and assigns it to the Bridge field.
func (o *BulkWritableInterfaceRequest) SetBridge(v BridgeInterface) {
	o.Bridge.Set(&v)
}
// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *BulkWritableInterfaceRequest) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *BulkWritableInterfaceRequest) UnsetBridge() {
	o.Bridge.Unset()
}

// GetLag returns the Lag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableInterfaceRequest) GetLag() ParentLAG {
	if o == nil || IsNil(o.Lag.Get()) {
		var ret ParentLAG
		return ret
	}
	return *o.Lag.Get()
}

// GetLagOk returns a tuple with the Lag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableInterfaceRequest) GetLagOk() (*ParentLAG, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lag.Get(), o.Lag.IsSet()
}

// HasLag returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasLag() bool {
	if o != nil && o.Lag.IsSet() {
		return true
	}

	return false
}

// SetLag gets a reference to the given NullableParentLAG and assigns it to the Lag field.
func (o *BulkWritableInterfaceRequest) SetLag(v ParentLAG) {
	o.Lag.Set(&v)
}
// SetLagNil sets the value for Lag to be an explicit nil
func (o *BulkWritableInterfaceRequest) SetLagNil() {
	o.Lag.Set(nil)
}

// UnsetLag ensures that no value is present for Lag, not even an explicit nil
func (o *BulkWritableInterfaceRequest) UnsetLag() {
	o.Lag.Unset()
}

// GetUntaggedVlan returns the UntaggedVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableInterfaceRequest) GetUntaggedVlan() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.UntaggedVlan.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.UntaggedVlan.Get()
}

// GetUntaggedVlanOk returns a tuple with the UntaggedVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableInterfaceRequest) GetUntaggedVlanOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.UntaggedVlan.Get(), o.UntaggedVlan.IsSet()
}

// HasUntaggedVlan returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasUntaggedVlan() bool {
	if o != nil && o.UntaggedVlan.IsSet() {
		return true
	}

	return false
}

// SetUntaggedVlan gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the UntaggedVlan field.
func (o *BulkWritableInterfaceRequest) SetUntaggedVlan(v BulkWritableCircuitRequestTenant) {
	o.UntaggedVlan.Set(&v)
}
// SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil
func (o *BulkWritableInterfaceRequest) SetUntaggedVlanNil() {
	o.UntaggedVlan.Set(nil)
}

// UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
func (o *BulkWritableInterfaceRequest) UnsetUntaggedVlan() {
	o.UntaggedVlan.Unset()
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableInterfaceRequest) GetVrf() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableInterfaceRequest) GetVrfOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Vrf field.
func (o *BulkWritableInterfaceRequest) SetVrf(v BulkWritableCircuitRequestTenant) {
	o.Vrf.Set(&v)
}
// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *BulkWritableInterfaceRequest) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *BulkWritableInterfaceRequest) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTaggedVlans returns the TaggedVlans field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRequest) GetTaggedVlans() []TaggedVLANs {
	if o == nil || IsNil(o.TaggedVlans) {
		var ret []TaggedVLANs
		return ret
	}
	return o.TaggedVlans
}

// GetTaggedVlansOk returns a tuple with the TaggedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetTaggedVlansOk() ([]TaggedVLANs, bool) {
	if o == nil || IsNil(o.TaggedVlans) {
		return nil, false
	}
	return o.TaggedVlans, true
}

// HasTaggedVlans returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasTaggedVlans() bool {
	if o != nil && !IsNil(o.TaggedVlans) {
		return true
	}

	return false
}

// SetTaggedVlans gets a reference to the given []TaggedVLANs and assigns it to the TaggedVlans field.
func (o *BulkWritableInterfaceRequest) SetTaggedVlans(v []TaggedVLANs) {
	o.TaggedVlans = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableInterfaceRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableInterfaceRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableInterfaceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o BulkWritableInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableInterfaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if o.MacAddress.IsSet() {
		toSerialize["mac_address"] = o.MacAddress.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Mtu.IsSet() {
		toSerialize["mtu"] = o.Mtu.Get()
	}
	if !IsNil(o.MgmtOnly) {
		toSerialize["mgmt_only"] = o.MgmtOnly
	}
	toSerialize["device"] = o.Device
	toSerialize["status"] = o.Status
	if o.ParentInterface.IsSet() {
		toSerialize["parent_interface"] = o.ParentInterface.Get()
	}
	if o.Bridge.IsSet() {
		toSerialize["bridge"] = o.Bridge.Get()
	}
	if o.Lag.IsSet() {
		toSerialize["lag"] = o.Lag.Get()
	}
	if o.UntaggedVlan.IsSet() {
		toSerialize["untagged_vlan"] = o.UntaggedVlan.Get()
	}
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
	}
	if !IsNil(o.TaggedVlans) {
		toSerialize["tagged_vlans"] = o.TaggedVlans
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkWritableInterfaceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"name",
		"device",
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBulkWritableInterfaceRequest := _BulkWritableInterfaceRequest{}

	err = json.Unmarshal(data, &varBulkWritableInterfaceRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableInterfaceRequest(varBulkWritableInterfaceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "mac_address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "mtu")
		delete(additionalProperties, "mgmt_only")
		delete(additionalProperties, "device")
		delete(additionalProperties, "status")
		delete(additionalProperties, "parent_interface")
		delete(additionalProperties, "bridge")
		delete(additionalProperties, "lag")
		delete(additionalProperties, "untagged_vlan")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "tagged_vlans")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableInterfaceRequest struct {
	value *BulkWritableInterfaceRequest
	isSet bool
}

func (v NullableBulkWritableInterfaceRequest) Get() *BulkWritableInterfaceRequest {
	return v.value
}

func (v *NullableBulkWritableInterfaceRequest) Set(val *BulkWritableInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableInterfaceRequest(val *BulkWritableInterfaceRequest) *NullableBulkWritableInterfaceRequest {
	return &NullableBulkWritableInterfaceRequest{value: val, isSet: true}
}

func (v NullableBulkWritableInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

