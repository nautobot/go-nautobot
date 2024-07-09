/*
API Documentation

Source of truth and network automation platform

API version: 2.2.5 (2.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// ParentChildStatus - Parent devices house child devices in device bays. Leave blank if this device type is neither a parent nor a child.
type ParentChildStatus struct {
	BlankEnum *BlankEnum
	SubdeviceRoleEnum *SubdeviceRoleEnum
}

// BlankEnumAsParentChildStatus is a convenience function that returns BlankEnum wrapped in ParentChildStatus
func BlankEnumAsParentChildStatus(v *BlankEnum) ParentChildStatus {
	return ParentChildStatus{
		BlankEnum: v,
	}
}

// SubdeviceRoleEnumAsParentChildStatus is a convenience function that returns SubdeviceRoleEnum wrapped in ParentChildStatus
func SubdeviceRoleEnumAsParentChildStatus(v *SubdeviceRoleEnum) ParentChildStatus {
	return ParentChildStatus{
		SubdeviceRoleEnum: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ParentChildStatus) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlankEnum
	err = newStrictDecoder(data).Decode(&dst.BlankEnum)
	if err == nil {
		jsonBlankEnum, _ := json.Marshal(dst.BlankEnum)
		if string(jsonBlankEnum) == "{}" { // empty struct
			dst.BlankEnum = nil
		} else {
			if err = validator.Validate(dst.BlankEnum); err != nil {
				dst.BlankEnum = nil
			} else {
				match++
			}
		}
	} else {
		dst.BlankEnum = nil
	}

	// try to unmarshal data into SubdeviceRoleEnum
	err = newStrictDecoder(data).Decode(&dst.SubdeviceRoleEnum)
	if err == nil {
		jsonSubdeviceRoleEnum, _ := json.Marshal(dst.SubdeviceRoleEnum)
		if string(jsonSubdeviceRoleEnum) == "{}" { // empty struct
			dst.SubdeviceRoleEnum = nil
		} else {
			if err = validator.Validate(dst.SubdeviceRoleEnum); err != nil {
				dst.SubdeviceRoleEnum = nil
			} else {
				match++
			}
		}
	} else {
		dst.SubdeviceRoleEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.SubdeviceRoleEnum = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ParentChildStatus)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ParentChildStatus)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ParentChildStatus) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.SubdeviceRoleEnum != nil {
		return json.Marshal(&src.SubdeviceRoleEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ParentChildStatus) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.SubdeviceRoleEnum != nil {
		return obj.SubdeviceRoleEnum
	}

	// all schemas are nil
	return nil
}

type NullableParentChildStatus struct {
	value *ParentChildStatus
	isSet bool
}

func (v NullableParentChildStatus) Get() *ParentChildStatus {
	return v.value
}

func (v *NullableParentChildStatus) Set(val *ParentChildStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableParentChildStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableParentChildStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentChildStatus(val *ParentChildStatus) *NullableParentChildStatus {
	return &NullableParentChildStatus{value: val, isSet: true}
}

func (v NullableParentChildStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentChildStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

