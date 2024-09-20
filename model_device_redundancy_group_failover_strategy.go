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

// checks if the DeviceRedundancyGroupFailoverStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceRedundancyGroupFailoverStrategy{}

// DeviceRedundancyGroupFailoverStrategy struct for DeviceRedundancyGroupFailoverStrategy
type DeviceRedundancyGroupFailoverStrategy struct {
	Value *DeviceRedundancyGroupFailoverStrategyValue `json:"value,omitempty"`
	Label *DeviceRedundancyGroupFailoverStrategyLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceRedundancyGroupFailoverStrategy DeviceRedundancyGroupFailoverStrategy

// NewDeviceRedundancyGroupFailoverStrategy instantiates a new DeviceRedundancyGroupFailoverStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceRedundancyGroupFailoverStrategy() *DeviceRedundancyGroupFailoverStrategy {
	this := DeviceRedundancyGroupFailoverStrategy{}
	return &this
}

// NewDeviceRedundancyGroupFailoverStrategyWithDefaults instantiates a new DeviceRedundancyGroupFailoverStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceRedundancyGroupFailoverStrategyWithDefaults() *DeviceRedundancyGroupFailoverStrategy {
	this := DeviceRedundancyGroupFailoverStrategy{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeviceRedundancyGroupFailoverStrategy) GetValue() DeviceRedundancyGroupFailoverStrategyValue {
	if o == nil || IsNil(o.Value) {
		var ret DeviceRedundancyGroupFailoverStrategyValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRedundancyGroupFailoverStrategy) GetValueOk() (*DeviceRedundancyGroupFailoverStrategyValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeviceRedundancyGroupFailoverStrategy) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given DeviceRedundancyGroupFailoverStrategyValue and assigns it to the Value field.
func (o *DeviceRedundancyGroupFailoverStrategy) SetValue(v DeviceRedundancyGroupFailoverStrategyValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DeviceRedundancyGroupFailoverStrategy) GetLabel() DeviceRedundancyGroupFailoverStrategyLabel {
	if o == nil || IsNil(o.Label) {
		var ret DeviceRedundancyGroupFailoverStrategyLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRedundancyGroupFailoverStrategy) GetLabelOk() (*DeviceRedundancyGroupFailoverStrategyLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DeviceRedundancyGroupFailoverStrategy) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given DeviceRedundancyGroupFailoverStrategyLabel and assigns it to the Label field.
func (o *DeviceRedundancyGroupFailoverStrategy) SetLabel(v DeviceRedundancyGroupFailoverStrategyLabel) {
	o.Label = &v
}

func (o DeviceRedundancyGroupFailoverStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceRedundancyGroupFailoverStrategy) ToMap() (map[string]interface{}, error) {
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

func (o *DeviceRedundancyGroupFailoverStrategy) UnmarshalJSON(data []byte) (err error) {
	varDeviceRedundancyGroupFailoverStrategy := _DeviceRedundancyGroupFailoverStrategy{}

	err = json.Unmarshal(data, &varDeviceRedundancyGroupFailoverStrategy)

	if err != nil {
		return err
	}

	*o = DeviceRedundancyGroupFailoverStrategy(varDeviceRedundancyGroupFailoverStrategy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceRedundancyGroupFailoverStrategy struct {
	value *DeviceRedundancyGroupFailoverStrategy
	isSet bool
}

func (v NullableDeviceRedundancyGroupFailoverStrategy) Get() *DeviceRedundancyGroupFailoverStrategy {
	return v.value
}

func (v *NullableDeviceRedundancyGroupFailoverStrategy) Set(val *DeviceRedundancyGroupFailoverStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRedundancyGroupFailoverStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRedundancyGroupFailoverStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRedundancyGroupFailoverStrategy(val *DeviceRedundancyGroupFailoverStrategy) *NullableDeviceRedundancyGroupFailoverStrategy {
	return &NullableDeviceRedundancyGroupFailoverStrategy{value: val, isSet: true}
}

func (v NullableDeviceRedundancyGroupFailoverStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRedundancyGroupFailoverStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


