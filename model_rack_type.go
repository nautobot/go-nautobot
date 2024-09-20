/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the RackType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackType{}

// RackType struct for RackType
type RackType struct {
	Value *RackTypeValue `json:"value,omitempty"`
	Label *RackTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackType RackType

// NewRackType instantiates a new RackType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackType() *RackType {
	this := RackType{}
	return &this
}

// NewRackTypeWithDefaults instantiates a new RackType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackTypeWithDefaults() *RackType {
	this := RackType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RackType) GetValue() RackTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret RackTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetValueOk() (*RackTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RackType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given RackTypeValue and assigns it to the Value field.
func (o *RackType) SetValue(v RackTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RackType) GetLabel() RackTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret RackTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetLabelOk() (*RackTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RackType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given RackTypeLabel and assigns it to the Label field.
func (o *RackType) SetLabel(v RackTypeLabel) {
	o.Label = &v
}

func (o RackType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RackType) UnmarshalJSON(data []byte) (err error) {
	varRackType := _RackType{}

	err = json.Unmarshal(data, &varRackType)

	if err != nil {
		return err
	}

	*o = RackType(varRackType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackType struct {
	value *RackType
	isSet bool
}

func (v NullableRackType) Get() *RackType {
	return v.value
}

func (v *NullableRackType) Set(val *RackType) {
	v.value = val
	v.isSet = true
}

func (v NullableRackType) IsSet() bool {
	return v.isSet
}

func (v *NullableRackType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackType(val *RackType) *NullableRackType {
	return &NullableRackType{value: val, isSet: true}
}

func (v NullableRackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


