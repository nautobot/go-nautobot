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

// DeviceRedundancyGroupFailoverStrategyValue the model 'DeviceRedundancyGroupFailoverStrategyValue'
type DeviceRedundancyGroupFailoverStrategyValue string

// List of DeviceRedundancyGroup_failover_strategy_value
const (
	DEVICEREDUNDANCYGROUPFAILOVERSTRATEGYVALUE_EMPTY DeviceRedundancyGroupFailoverStrategyValue = ""
	DEVICEREDUNDANCYGROUPFAILOVERSTRATEGYVALUE_ACTIVE_ACTIVE DeviceRedundancyGroupFailoverStrategyValue = "active-active"
	DEVICEREDUNDANCYGROUPFAILOVERSTRATEGYVALUE_ACTIVE_PASSIVE DeviceRedundancyGroupFailoverStrategyValue = "active-passive"
)

// All allowed values of DeviceRedundancyGroupFailoverStrategyValue enum
var AllowedDeviceRedundancyGroupFailoverStrategyValueEnumValues = []DeviceRedundancyGroupFailoverStrategyValue{
	"",
	"active-active",
	"active-passive",
}

func (v *DeviceRedundancyGroupFailoverStrategyValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceRedundancyGroupFailoverStrategyValue(value)
	for _, existing := range AllowedDeviceRedundancyGroupFailoverStrategyValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceRedundancyGroupFailoverStrategyValue", value)
}

// NewDeviceRedundancyGroupFailoverStrategyValueFromValue returns a pointer to a valid DeviceRedundancyGroupFailoverStrategyValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceRedundancyGroupFailoverStrategyValueFromValue(v string) (*DeviceRedundancyGroupFailoverStrategyValue, error) {
	ev := DeviceRedundancyGroupFailoverStrategyValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceRedundancyGroupFailoverStrategyValue: valid values are %v", v, AllowedDeviceRedundancyGroupFailoverStrategyValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceRedundancyGroupFailoverStrategyValue) IsValid() bool {
	for _, existing := range AllowedDeviceRedundancyGroupFailoverStrategyValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceRedundancyGroup_failover_strategy_value value
func (v DeviceRedundancyGroupFailoverStrategyValue) Ptr() *DeviceRedundancyGroupFailoverStrategyValue {
	return &v
}

type NullableDeviceRedundancyGroupFailoverStrategyValue struct {
	value *DeviceRedundancyGroupFailoverStrategyValue
	isSet bool
}

func (v NullableDeviceRedundancyGroupFailoverStrategyValue) Get() *DeviceRedundancyGroupFailoverStrategyValue {
	return v.value
}

func (v *NullableDeviceRedundancyGroupFailoverStrategyValue) Set(val *DeviceRedundancyGroupFailoverStrategyValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRedundancyGroupFailoverStrategyValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRedundancyGroupFailoverStrategyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRedundancyGroupFailoverStrategyValue(val *DeviceRedundancyGroupFailoverStrategyValue) *NullableDeviceRedundancyGroupFailoverStrategyValue {
	return &NullableDeviceRedundancyGroupFailoverStrategyValue{value: val, isSet: true}
}

func (v NullableDeviceRedundancyGroupFailoverStrategyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRedundancyGroupFailoverStrategyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

