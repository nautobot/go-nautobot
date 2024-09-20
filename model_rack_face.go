/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// RackFace - struct for RackFace
type RackFace struct {
	BlankEnum *BlankEnum
	FaceEnum *FaceEnum
}

// BlankEnumAsRackFace is a convenience function that returns BlankEnum wrapped in RackFace
func BlankEnumAsRackFace(v *BlankEnum) RackFace {
	return RackFace{
		BlankEnum: v,
	}
}

// FaceEnumAsRackFace is a convenience function that returns FaceEnum wrapped in RackFace
func FaceEnumAsRackFace(v *FaceEnum) RackFace {
	return RackFace{
		FaceEnum: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RackFace) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into FaceEnum
	err = newStrictDecoder(data).Decode(&dst.FaceEnum)
	if err == nil {
		jsonFaceEnum, _ := json.Marshal(dst.FaceEnum)
		if string(jsonFaceEnum) == "{}" { // empty struct
			dst.FaceEnum = nil
		} else {
			if err = validator.Validate(dst.FaceEnum); err != nil {
				dst.FaceEnum = nil
			} else {
				match++
			}
		}
	} else {
		dst.FaceEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.FaceEnum = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RackFace)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RackFace)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RackFace) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.FaceEnum != nil {
		return json.Marshal(&src.FaceEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RackFace) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.FaceEnum != nil {
		return obj.FaceEnum
	}

	// all schemas are nil
	return nil
}

type NullableRackFace struct {
	value *RackFace
	isSet bool
}

func (v NullableRackFace) Get() *RackFace {
	return v.value
}

func (v *NullableRackFace) Set(val *RackFace) {
	v.value = val
	v.isSet = true
}

func (v NullableRackFace) IsSet() bool {
	return v.isSet
}

func (v *NullableRackFace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackFace(val *RackFace) *NullableRackFace {
	return &NullableRackFace{value: val, isSet: true}
}

func (v NullableRackFace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackFace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


