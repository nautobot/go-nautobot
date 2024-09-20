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

// checks if the CableType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CableType{}

// CableType struct for CableType
type CableType struct {
	Value *CableTypeValue `json:"value,omitempty"`
	Label *CableTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CableType CableType

// NewCableType instantiates a new CableType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCableType() *CableType {
	this := CableType{}
	return &this
}

// NewCableTypeWithDefaults instantiates a new CableType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCableTypeWithDefaults() *CableType {
	this := CableType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CableType) GetValue() CableTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret CableTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableType) GetValueOk() (*CableTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CableType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CableTypeValue and assigns it to the Value field.
func (o *CableType) SetValue(v CableTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CableType) GetLabel() CableTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret CableTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableType) GetLabelOk() (*CableTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CableType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given CableTypeLabel and assigns it to the Label field.
func (o *CableType) SetLabel(v CableTypeLabel) {
	o.Label = &v
}

func (o CableType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CableType) ToMap() (map[string]interface{}, error) {
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

func (o *CableType) UnmarshalJSON(data []byte) (err error) {
	varCableType := _CableType{}

	err = json.Unmarshal(data, &varCableType)

	if err != nil {
		return err
	}

	*o = CableType(varCableType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCableType struct {
	value *CableType
	isSet bool
}

func (v NullableCableType) Get() *CableType {
	return v.value
}

func (v *NullableCableType) Set(val *CableType) {
	v.value = val
	v.isSet = true
}

func (v NullableCableType) IsSet() bool {
	return v.isSet
}

func (v *NullableCableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableType(val *CableType) *NullableCableType {
	return &NullableCableType{value: val, isSet: true}
}

func (v NullableCableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


