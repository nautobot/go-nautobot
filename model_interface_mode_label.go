/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// InterfaceModeLabel the model 'InterfaceModeLabel'
type InterfaceModeLabel string

// List of Interface_mode_label
const (
	INTERFACEMODELABEL_ACCESS InterfaceModeLabel = "Access"
	INTERFACEMODELABEL_TAGGED InterfaceModeLabel = "Tagged"
	INTERFACEMODELABEL_TAGGED__ALL InterfaceModeLabel = "Tagged (All)"
)

// All allowed values of InterfaceModeLabel enum
var AllowedInterfaceModeLabelEnumValues = []InterfaceModeLabel{
	"Access",
	"Tagged",
	"Tagged (All)",
}

func (v *InterfaceModeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceModeLabel(value)
	for _, existing := range AllowedInterfaceModeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceModeLabel", value)
}

// NewInterfaceModeLabelFromValue returns a pointer to a valid InterfaceModeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceModeLabelFromValue(v string) (*InterfaceModeLabel, error) {
	ev := InterfaceModeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceModeLabel: valid values are %v", v, AllowedInterfaceModeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceModeLabel) IsValid() bool {
	for _, existing := range AllowedInterfaceModeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_mode_label value
func (v InterfaceModeLabel) Ptr() *InterfaceModeLabel {
	return &v
}

type NullableInterfaceModeLabel struct {
	value *InterfaceModeLabel
	isSet bool
}

func (v NullableInterfaceModeLabel) Get() *InterfaceModeLabel {
	return v.value
}

func (v *NullableInterfaceModeLabel) Set(val *InterfaceModeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceModeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceModeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceModeLabel(val *InterfaceModeLabel) *NullableInterfaceModeLabel {
	return &NullableInterfaceModeLabel{value: val, isSet: true}
}

func (v NullableInterfaceModeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceModeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
