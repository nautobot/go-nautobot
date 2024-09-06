/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// PatchedWritableRackRequestType - struct for PatchedWritableRackRequestType
type PatchedWritableRackRequestType struct {
	BlankEnum *BlankEnum
	RackTypeChoices *RackTypeChoices
}

// BlankEnumAsPatchedWritableRackRequestType is a convenience function that returns BlankEnum wrapped in PatchedWritableRackRequestType
func BlankEnumAsPatchedWritableRackRequestType(v *BlankEnum) PatchedWritableRackRequestType {
	return PatchedWritableRackRequestType{
		BlankEnum: v,
	}
}

// RackTypeChoicesAsPatchedWritableRackRequestType is a convenience function that returns RackTypeChoices wrapped in PatchedWritableRackRequestType
func RackTypeChoicesAsPatchedWritableRackRequestType(v *RackTypeChoices) PatchedWritableRackRequestType {
	return PatchedWritableRackRequestType{
		RackTypeChoices: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritableRackRequestType) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into RackTypeChoices
	err = newStrictDecoder(data).Decode(&dst.RackTypeChoices)
	if err == nil {
		jsonRackTypeChoices, _ := json.Marshal(dst.RackTypeChoices)
		if string(jsonRackTypeChoices) == "{}" { // empty struct
			dst.RackTypeChoices = nil
		} else {
			if err = validator.Validate(dst.RackTypeChoices); err != nil {
				dst.RackTypeChoices = nil
			} else {
				match++
			}
		}
	} else {
		dst.RackTypeChoices = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.RackTypeChoices = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PatchedWritableRackRequestType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PatchedWritableRackRequestType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritableRackRequestType) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.RackTypeChoices != nil {
		return json.Marshal(&src.RackTypeChoices)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritableRackRequestType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.RackTypeChoices != nil {
		return obj.RackTypeChoices
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritableRackRequestType struct {
	value *PatchedWritableRackRequestType
	isSet bool
}

func (v NullablePatchedWritableRackRequestType) Get() *PatchedWritableRackRequestType {
	return v.value
}

func (v *NullablePatchedWritableRackRequestType) Set(val *PatchedWritableRackRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackRequestType(val *PatchedWritableRackRequestType) *NullablePatchedWritableRackRequestType {
	return &NullablePatchedWritableRackRequestType{value: val, isSet: true}
}

func (v NullablePatchedWritableRackRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


