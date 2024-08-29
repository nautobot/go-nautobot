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

// InterfaceRedundancyGroupProtocolChoices the model 'InterfaceRedundancyGroupProtocolChoices'
type InterfaceRedundancyGroupProtocolChoices string

// List of InterfaceRedundancyGroupProtocolChoices
const (
	INTERFACEREDUNDANCYGROUPPROTOCOLCHOICES_HSRP InterfaceRedundancyGroupProtocolChoices = "hsrp"
	INTERFACEREDUNDANCYGROUPPROTOCOLCHOICES_VRRP InterfaceRedundancyGroupProtocolChoices = "vrrp"
	INTERFACEREDUNDANCYGROUPPROTOCOLCHOICES_GLBP InterfaceRedundancyGroupProtocolChoices = "glbp"
	INTERFACEREDUNDANCYGROUPPROTOCOLCHOICES_CARP InterfaceRedundancyGroupProtocolChoices = "carp"
)

// All allowed values of InterfaceRedundancyGroupProtocolChoices enum
var AllowedInterfaceRedundancyGroupProtocolChoicesEnumValues = []InterfaceRedundancyGroupProtocolChoices{
	"hsrp",
	"vrrp",
	"glbp",
	"carp",
}

func (v *InterfaceRedundancyGroupProtocolChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceRedundancyGroupProtocolChoices(value)
	for _, existing := range AllowedInterfaceRedundancyGroupProtocolChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceRedundancyGroupProtocolChoices", value)
}

// NewInterfaceRedundancyGroupProtocolChoicesFromValue returns a pointer to a valid InterfaceRedundancyGroupProtocolChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceRedundancyGroupProtocolChoicesFromValue(v string) (*InterfaceRedundancyGroupProtocolChoices, error) {
	ev := InterfaceRedundancyGroupProtocolChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceRedundancyGroupProtocolChoices: valid values are %v", v, AllowedInterfaceRedundancyGroupProtocolChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceRedundancyGroupProtocolChoices) IsValid() bool {
	for _, existing := range AllowedInterfaceRedundancyGroupProtocolChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InterfaceRedundancyGroupProtocolChoices value
func (v InterfaceRedundancyGroupProtocolChoices) Ptr() *InterfaceRedundancyGroupProtocolChoices {
	return &v
}

type NullableInterfaceRedundancyGroupProtocolChoices struct {
	value *InterfaceRedundancyGroupProtocolChoices
	isSet bool
}

func (v NullableInterfaceRedundancyGroupProtocolChoices) Get() *InterfaceRedundancyGroupProtocolChoices {
	return v.value
}

func (v *NullableInterfaceRedundancyGroupProtocolChoices) Set(val *InterfaceRedundancyGroupProtocolChoices) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceRedundancyGroupProtocolChoices) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceRedundancyGroupProtocolChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceRedundancyGroupProtocolChoices(val *InterfaceRedundancyGroupProtocolChoices) *NullableInterfaceRedundancyGroupProtocolChoices {
	return &NullableInterfaceRedundancyGroupProtocolChoices{value: val, isSet: true}
}

func (v NullableInterfaceRedundancyGroupProtocolChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceRedundancyGroupProtocolChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
