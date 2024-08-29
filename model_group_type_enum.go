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

// GroupTypeEnum the model 'GroupTypeEnum'
type GroupTypeEnum string

// List of GroupTypeEnum
const (
	GROUPTYPEENUM_DYNAMIC_FILTER GroupTypeEnum = "dynamic-filter"
	GROUPTYPEENUM_DYNAMIC_SET GroupTypeEnum = "dynamic-set"
	GROUPTYPEENUM_STATIC GroupTypeEnum = "static"
)

// All allowed values of GroupTypeEnum enum
var AllowedGroupTypeEnumEnumValues = []GroupTypeEnum{
	"dynamic-filter",
	"dynamic-set",
	"static",
}

func (v *GroupTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupTypeEnum(value)
	for _, existing := range AllowedGroupTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupTypeEnum", value)
}

// NewGroupTypeEnumFromValue returns a pointer to a valid GroupTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupTypeEnumFromValue(v string) (*GroupTypeEnum, error) {
	ev := GroupTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupTypeEnum: valid values are %v", v, AllowedGroupTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupTypeEnum) IsValid() bool {
	for _, existing := range AllowedGroupTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupTypeEnum value
func (v GroupTypeEnum) Ptr() *GroupTypeEnum {
	return &v
}

type NullableGroupTypeEnum struct {
	value *GroupTypeEnum
	isSet bool
}

func (v NullableGroupTypeEnum) Get() *GroupTypeEnum {
	return v.value
}

func (v *NullableGroupTypeEnum) Set(val *GroupTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupTypeEnum(val *GroupTypeEnum) *NullableGroupTypeEnum {
	return &NullableGroupTypeEnum{value: val, isSet: true}
}

func (v NullableGroupTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
