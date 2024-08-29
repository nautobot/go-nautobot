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

// LogLevelEnum the model 'LogLevelEnum'
type LogLevelEnum string

// List of LogLevelEnum
const (
	LOGLEVELENUM_DEBUG LogLevelEnum = "debug"
	LOGLEVELENUM_INFO LogLevelEnum = "info"
	LOGLEVELENUM_WARNING LogLevelEnum = "warning"
	LOGLEVELENUM_ERROR LogLevelEnum = "error"
	LOGLEVELENUM_CRITICAL LogLevelEnum = "critical"
)

// All allowed values of LogLevelEnum enum
var AllowedLogLevelEnumEnumValues = []LogLevelEnum{
	"debug",
	"info",
	"warning",
	"error",
	"critical",
}

func (v *LogLevelEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogLevelEnum(value)
	for _, existing := range AllowedLogLevelEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogLevelEnum", value)
}

// NewLogLevelEnumFromValue returns a pointer to a valid LogLevelEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogLevelEnumFromValue(v string) (*LogLevelEnum, error) {
	ev := LogLevelEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogLevelEnum: valid values are %v", v, AllowedLogLevelEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogLevelEnum) IsValid() bool {
	for _, existing := range AllowedLogLevelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogLevelEnum value
func (v LogLevelEnum) Ptr() *LogLevelEnum {
	return &v
}

type NullableLogLevelEnum struct {
	value *LogLevelEnum
	isSet bool
}

func (v NullableLogLevelEnum) Get() *LogLevelEnum {
	return v.value
}

func (v *NullableLogLevelEnum) Set(val *LogLevelEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableLogLevelEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableLogLevelEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogLevelEnum(val *LogLevelEnum) *NullableLogLevelEnum {
	return &NullableLogLevelEnum{value: val, isSet: true}
}

func (v NullableLogLevelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogLevelEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
