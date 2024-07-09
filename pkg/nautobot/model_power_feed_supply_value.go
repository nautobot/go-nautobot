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

// PowerFeedSupplyValue the model 'PowerFeedSupplyValue'
type PowerFeedSupplyValue string

// List of PowerFeed_supply_value
const (
	POWERFEEDSUPPLYVALUE_AC PowerFeedSupplyValue = "ac"
	POWERFEEDSUPPLYVALUE_DC PowerFeedSupplyValue = "dc"
)

// All allowed values of PowerFeedSupplyValue enum
var AllowedPowerFeedSupplyValueEnumValues = []PowerFeedSupplyValue{
	"ac",
	"dc",
}

func (v *PowerFeedSupplyValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerFeedSupplyValue(value)
	for _, existing := range AllowedPowerFeedSupplyValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerFeedSupplyValue", value)
}

// NewPowerFeedSupplyValueFromValue returns a pointer to a valid PowerFeedSupplyValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerFeedSupplyValueFromValue(v string) (*PowerFeedSupplyValue, error) {
	ev := PowerFeedSupplyValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerFeedSupplyValue: valid values are %v", v, AllowedPowerFeedSupplyValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerFeedSupplyValue) IsValid() bool {
	for _, existing := range AllowedPowerFeedSupplyValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerFeed_supply_value value
func (v PowerFeedSupplyValue) Ptr() *PowerFeedSupplyValue {
	return &v
}

type NullablePowerFeedSupplyValue struct {
	value *PowerFeedSupplyValue
	isSet bool
}

func (v NullablePowerFeedSupplyValue) Get() *PowerFeedSupplyValue {
	return v.value
}

func (v *NullablePowerFeedSupplyValue) Set(val *PowerFeedSupplyValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedSupplyValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedSupplyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedSupplyValue(val *PowerFeedSupplyValue) *NullablePowerFeedSupplyValue {
	return &NullablePowerFeedSupplyValue{value: val, isSet: true}
}

func (v NullablePowerFeedSupplyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedSupplyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
