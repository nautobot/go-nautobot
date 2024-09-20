/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the PatchedBulkWritableCircuitTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableCircuitTerminationRequest{}

// PatchedBulkWritableCircuitTerminationRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritableCircuitTerminationRequest struct {
	Id string `json:"id"`
	TermSide *TermSideEnum `json:"term_side,omitempty"`
	PortSpeed NullableInt32 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed NullableInt32 `json:"upstream_speed,omitempty"`
	XconnectId *string `json:"xconnect_id,omitempty"`
	PpInfo *string `json:"pp_info,omitempty"`
	Description *string `json:"description,omitempty"`
	Circuit *BulkWritableCableRequestStatus `json:"circuit,omitempty"`
	Location NullableBulkWritableCircuitRequestTenant `json:"location,omitempty"`
	ProviderNetwork NullableBulkWritableCircuitRequestTenant `json:"provider_network,omitempty"`
	CloudNetwork NullableBulkWritableCircuitRequestTenant `json:"cloud_network,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableCircuitTerminationRequest PatchedBulkWritableCircuitTerminationRequest

// NewPatchedBulkWritableCircuitTerminationRequest instantiates a new PatchedBulkWritableCircuitTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableCircuitTerminationRequest(id string) *PatchedBulkWritableCircuitTerminationRequest {
	this := PatchedBulkWritableCircuitTerminationRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableCircuitTerminationRequestWithDefaults instantiates a new PatchedBulkWritableCircuitTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableCircuitTerminationRequestWithDefaults() *PatchedBulkWritableCircuitTerminationRequest {
	this := PatchedBulkWritableCircuitTerminationRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableCircuitTerminationRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableCircuitTerminationRequest) SetId(v string) {
	o.Id = v
}

// GetTermSide returns the TermSide field value if set, zero value otherwise.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetTermSide() TermSideEnum {
	if o == nil || IsNil(o.TermSide) {
		var ret TermSideEnum
		return ret
	}
	return *o.TermSide
}

// GetTermSideOk returns a tuple with the TermSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetTermSideOk() (*TermSideEnum, bool) {
	if o == nil || IsNil(o.TermSide) {
		return nil, false
	}
	return o.TermSide, true
}

// HasTermSide returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasTermSide() bool {
	if o != nil && !IsNil(o.TermSide) {
		return true
	}

	return false
}

// SetTermSide gets a reference to the given TermSideEnum and assigns it to the TermSide field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetTermSide(v TermSideEnum) {
	o.TermSide = &v
}

// GetPortSpeed returns the PortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableCircuitTerminationRequest) GetPortSpeed() int32 {
	if o == nil || IsNil(o.PortSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.PortSpeed.Get()
}

// GetPortSpeedOk returns a tuple with the PortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableCircuitTerminationRequest) GetPortSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortSpeed.Get(), o.PortSpeed.IsSet()
}

// HasPortSpeed returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasPortSpeed() bool {
	if o != nil && o.PortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPortSpeed gets a reference to the given NullableInt32 and assigns it to the PortSpeed field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetPortSpeed(v int32) {
	o.PortSpeed.Set(&v)
}
// SetPortSpeedNil sets the value for PortSpeed to be an explicit nil
func (o *PatchedBulkWritableCircuitTerminationRequest) SetPortSpeedNil() {
	o.PortSpeed.Set(nil)
}

// UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
func (o *PatchedBulkWritableCircuitTerminationRequest) UnsetPortSpeed() {
	o.PortSpeed.Unset()
}

// GetUpstreamSpeed returns the UpstreamSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableCircuitTerminationRequest) GetUpstreamSpeed() int32 {
	if o == nil || IsNil(o.UpstreamSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.UpstreamSpeed.Get()
}

// GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableCircuitTerminationRequest) GetUpstreamSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamSpeed.Get(), o.UpstreamSpeed.IsSet()
}

// HasUpstreamSpeed returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasUpstreamSpeed() bool {
	if o != nil && o.UpstreamSpeed.IsSet() {
		return true
	}

	return false
}

// SetUpstreamSpeed gets a reference to the given NullableInt32 and assigns it to the UpstreamSpeed field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetUpstreamSpeed(v int32) {
	o.UpstreamSpeed.Set(&v)
}
// SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil
func (o *PatchedBulkWritableCircuitTerminationRequest) SetUpstreamSpeedNil() {
	o.UpstreamSpeed.Set(nil)
}

// UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
func (o *PatchedBulkWritableCircuitTerminationRequest) UnsetUpstreamSpeed() {
	o.UpstreamSpeed.Unset()
}

// GetXconnectId returns the XconnectId field value if set, zero value otherwise.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetXconnectId() string {
	if o == nil || IsNil(o.XconnectId) {
		var ret string
		return ret
	}
	return *o.XconnectId
}

// GetXconnectIdOk returns a tuple with the XconnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetXconnectIdOk() (*string, bool) {
	if o == nil || IsNil(o.XconnectId) {
		return nil, false
	}
	return o.XconnectId, true
}

// HasXconnectId returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasXconnectId() bool {
	if o != nil && !IsNil(o.XconnectId) {
		return true
	}

	return false
}

// SetXconnectId gets a reference to the given string and assigns it to the XconnectId field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetXconnectId(v string) {
	o.XconnectId = &v
}

// GetPpInfo returns the PpInfo field value if set, zero value otherwise.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetPpInfo() string {
	if o == nil || IsNil(o.PpInfo) {
		var ret string
		return ret
	}
	return *o.PpInfo
}

// GetPpInfoOk returns a tuple with the PpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetPpInfoOk() (*string, bool) {
	if o == nil || IsNil(o.PpInfo) {
		return nil, false
	}
	return o.PpInfo, true
}

// HasPpInfo returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasPpInfo() bool {
	if o != nil && !IsNil(o.PpInfo) {
		return true
	}

	return false
}

// SetPpInfo gets a reference to the given string and assigns it to the PpInfo field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetPpInfo(v string) {
	o.PpInfo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCircuit returns the Circuit field value if set, zero value otherwise.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetCircuit() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Circuit) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetCircuitOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Circuit) {
		return nil, false
	}
	return o.Circuit, true
}

// HasCircuit returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasCircuit() bool {
	if o != nil && !IsNil(o.Circuit) {
		return true
	}

	return false
}

// SetCircuit gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Circuit field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetCircuit(v BulkWritableCableRequestStatus) {
	o.Circuit = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableCircuitTerminationRequest) GetLocation() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Location.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableCircuitTerminationRequest) GetLocationOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Location field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetLocation(v BulkWritableCircuitRequestTenant) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *PatchedBulkWritableCircuitTerminationRequest) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *PatchedBulkWritableCircuitTerminationRequest) UnsetLocation() {
	o.Location.Unset()
}

// GetProviderNetwork returns the ProviderNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableCircuitTerminationRequest) GetProviderNetwork() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ProviderNetwork.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ProviderNetwork.Get()
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableCircuitTerminationRequest) GetProviderNetworkOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderNetwork.Get(), o.ProviderNetwork.IsSet()
}

// HasProviderNetwork returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasProviderNetwork() bool {
	if o != nil && o.ProviderNetwork.IsSet() {
		return true
	}

	return false
}

// SetProviderNetwork gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ProviderNetwork field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetProviderNetwork(v BulkWritableCircuitRequestTenant) {
	o.ProviderNetwork.Set(&v)
}
// SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil
func (o *PatchedBulkWritableCircuitTerminationRequest) SetProviderNetworkNil() {
	o.ProviderNetwork.Set(nil)
}

// UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
func (o *PatchedBulkWritableCircuitTerminationRequest) UnsetProviderNetwork() {
	o.ProviderNetwork.Unset()
}

// GetCloudNetwork returns the CloudNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableCircuitTerminationRequest) GetCloudNetwork() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.CloudNetwork.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.CloudNetwork.Get()
}

// GetCloudNetworkOk returns a tuple with the CloudNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableCircuitTerminationRequest) GetCloudNetworkOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudNetwork.Get(), o.CloudNetwork.IsSet()
}

// HasCloudNetwork returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasCloudNetwork() bool {
	if o != nil && o.CloudNetwork.IsSet() {
		return true
	}

	return false
}

// SetCloudNetwork gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the CloudNetwork field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetCloudNetwork(v BulkWritableCircuitRequestTenant) {
	o.CloudNetwork.Set(&v)
}
// SetCloudNetworkNil sets the value for CloudNetwork to be an explicit nil
func (o *PatchedBulkWritableCircuitTerminationRequest) SetCloudNetworkNil() {
	o.CloudNetwork.Set(nil)
}

// UnsetCloudNetwork ensures that no value is present for CloudNetwork, not even an explicit nil
func (o *PatchedBulkWritableCircuitTerminationRequest) UnsetCloudNetwork() {
	o.CloudNetwork.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableCircuitTerminationRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableCircuitTerminationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedBulkWritableCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableCircuitTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.TermSide) {
		toSerialize["term_side"] = o.TermSide
	}
	if o.PortSpeed.IsSet() {
		toSerialize["port_speed"] = o.PortSpeed.Get()
	}
	if o.UpstreamSpeed.IsSet() {
		toSerialize["upstream_speed"] = o.UpstreamSpeed.Get()
	}
	if !IsNil(o.XconnectId) {
		toSerialize["xconnect_id"] = o.XconnectId
	}
	if !IsNil(o.PpInfo) {
		toSerialize["pp_info"] = o.PpInfo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Circuit) {
		toSerialize["circuit"] = o.Circuit
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.ProviderNetwork.IsSet() {
		toSerialize["provider_network"] = o.ProviderNetwork.Get()
	}
	if o.CloudNetwork.IsSet() {
		toSerialize["cloud_network"] = o.CloudNetwork.Get()
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

func (o *PatchedBulkWritableCircuitTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPatchedBulkWritableCircuitTerminationRequest := _PatchedBulkWritableCircuitTerminationRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableCircuitTerminationRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableCircuitTerminationRequest(varPatchedBulkWritableCircuitTerminationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "term_side")
		delete(additionalProperties, "port_speed")
		delete(additionalProperties, "upstream_speed")
		delete(additionalProperties, "xconnect_id")
		delete(additionalProperties, "pp_info")
		delete(additionalProperties, "description")
		delete(additionalProperties, "circuit")
		delete(additionalProperties, "location")
		delete(additionalProperties, "provider_network")
		delete(additionalProperties, "cloud_network")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableCircuitTerminationRequest struct {
	value *PatchedBulkWritableCircuitTerminationRequest
	isSet bool
}

func (v NullablePatchedBulkWritableCircuitTerminationRequest) Get() *PatchedBulkWritableCircuitTerminationRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableCircuitTerminationRequest) Set(val *PatchedBulkWritableCircuitTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableCircuitTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableCircuitTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableCircuitTerminationRequest(val *PatchedBulkWritableCircuitTerminationRequest) *NullablePatchedBulkWritableCircuitTerminationRequest {
	return &NullablePatchedBulkWritableCircuitTerminationRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableCircuitTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


