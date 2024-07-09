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

// PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy - struct for PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy
type PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy struct {
	BlankEnum *BlankEnum
	FailoverStrategyEnum *FailoverStrategyEnum
}

// BlankEnumAsPatchedWritableDeviceRedundancyGroupRequestFailoverStrategy is a convenience function that returns BlankEnum wrapped in PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy
func BlankEnumAsPatchedWritableDeviceRedundancyGroupRequestFailoverStrategy(v *BlankEnum) PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy {
	return PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy{
		BlankEnum: v,
	}
}

// FailoverStrategyEnumAsPatchedWritableDeviceRedundancyGroupRequestFailoverStrategy is a convenience function that returns FailoverStrategyEnum wrapped in PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy
func FailoverStrategyEnumAsPatchedWritableDeviceRedundancyGroupRequestFailoverStrategy(v *FailoverStrategyEnum) PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy {
	return PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy{
		FailoverStrategyEnum: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into FailoverStrategyEnum
	err = newStrictDecoder(data).Decode(&dst.FailoverStrategyEnum)
	if err == nil {
		jsonFailoverStrategyEnum, _ := json.Marshal(dst.FailoverStrategyEnum)
		if string(jsonFailoverStrategyEnum) == "{}" { // empty struct
			dst.FailoverStrategyEnum = nil
		} else {
			if err = validator.Validate(dst.FailoverStrategyEnum); err != nil {
				dst.FailoverStrategyEnum = nil
			} else {
				match++
			}
		}
	} else {
		dst.FailoverStrategyEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.FailoverStrategyEnum = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.FailoverStrategyEnum != nil {
		return json.Marshal(&src.FailoverStrategyEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.FailoverStrategyEnum != nil {
		return obj.FailoverStrategyEnum
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritableDeviceRedundancyGroupRequestFailoverStrategy struct {
	value *PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy
	isSet bool
}

func (v NullablePatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) Get() *PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy {
	return v.value
}

func (v *NullablePatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) Set(val *PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableDeviceRedundancyGroupRequestFailoverStrategy(val *PatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) *NullablePatchedWritableDeviceRedundancyGroupRequestFailoverStrategy {
	return &NullablePatchedWritableDeviceRedundancyGroupRequestFailoverStrategy{value: val, isSet: true}
}

func (v NullablePatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableDeviceRedundancyGroupRequestFailoverStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

