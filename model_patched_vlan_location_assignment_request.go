/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedVLANLocationAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedVLANLocationAssignmentRequest{}

// PatchedVLANLocationAssignmentRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedVLANLocationAssignmentRequest struct {
	Vlan *BulkWritableCableRequestStatus `json:"vlan,omitempty"`
	Location *BulkWritableCableRequestStatus `json:"location,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedVLANLocationAssignmentRequest PatchedVLANLocationAssignmentRequest

// NewPatchedVLANLocationAssignmentRequest instantiates a new PatchedVLANLocationAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedVLANLocationAssignmentRequest() *PatchedVLANLocationAssignmentRequest {
	this := PatchedVLANLocationAssignmentRequest{}
	return &this
}

// NewPatchedVLANLocationAssignmentRequestWithDefaults instantiates a new PatchedVLANLocationAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedVLANLocationAssignmentRequestWithDefaults() *PatchedVLANLocationAssignmentRequest {
	this := PatchedVLANLocationAssignmentRequest{}
	return &this
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *PatchedVLANLocationAssignmentRequest) GetVlan() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Vlan) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVLANLocationAssignmentRequest) GetVlanOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *PatchedVLANLocationAssignmentRequest) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Vlan field.
func (o *PatchedVLANLocationAssignmentRequest) SetVlan(v BulkWritableCableRequestStatus) {
	o.Vlan = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PatchedVLANLocationAssignmentRequest) GetLocation() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Location) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedVLANLocationAssignmentRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PatchedVLANLocationAssignmentRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Location field.
func (o *PatchedVLANLocationAssignmentRequest) SetLocation(v BulkWritableCableRequestStatus) {
	o.Location = &v
}

func (o PatchedVLANLocationAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedVLANLocationAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedVLANLocationAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedVLANLocationAssignmentRequest := _PatchedVLANLocationAssignmentRequest{}

	err = json.Unmarshal(data, &varPatchedVLANLocationAssignmentRequest)

	if err != nil {
		return err
	}

	*o = PatchedVLANLocationAssignmentRequest(varPatchedVLANLocationAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vlan")
		delete(additionalProperties, "location")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedVLANLocationAssignmentRequest struct {
	value *PatchedVLANLocationAssignmentRequest
	isSet bool
}

func (v NullablePatchedVLANLocationAssignmentRequest) Get() *PatchedVLANLocationAssignmentRequest {
	return v.value
}

func (v *NullablePatchedVLANLocationAssignmentRequest) Set(val *PatchedVLANLocationAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedVLANLocationAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedVLANLocationAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedVLANLocationAssignmentRequest(val *PatchedVLANLocationAssignmentRequest) *NullablePatchedVLANLocationAssignmentRequest {
	return &NullablePatchedVLANLocationAssignmentRequest{value: val, isSet: true}
}

func (v NullablePatchedVLANLocationAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedVLANLocationAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


