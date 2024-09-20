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

// BulkWritableExternalIntegrationRequestHttpMethod - struct for BulkWritableExternalIntegrationRequestHttpMethod
type BulkWritableExternalIntegrationRequestHttpMethod struct {
	BlankEnum *BlankEnum
	HttpMethodEnum *HttpMethodEnum
}

// BlankEnumAsBulkWritableExternalIntegrationRequestHttpMethod is a convenience function that returns BlankEnum wrapped in BulkWritableExternalIntegrationRequestHttpMethod
func BlankEnumAsBulkWritableExternalIntegrationRequestHttpMethod(v *BlankEnum) BulkWritableExternalIntegrationRequestHttpMethod {
	return BulkWritableExternalIntegrationRequestHttpMethod{
		BlankEnum: v,
	}
}

// HttpMethodEnumAsBulkWritableExternalIntegrationRequestHttpMethod is a convenience function that returns HttpMethodEnum wrapped in BulkWritableExternalIntegrationRequestHttpMethod
func HttpMethodEnumAsBulkWritableExternalIntegrationRequestHttpMethod(v *HttpMethodEnum) BulkWritableExternalIntegrationRequestHttpMethod {
	return BulkWritableExternalIntegrationRequestHttpMethod{
		HttpMethodEnum: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BulkWritableExternalIntegrationRequestHttpMethod) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into HttpMethodEnum
	err = newStrictDecoder(data).Decode(&dst.HttpMethodEnum)
	if err == nil {
		jsonHttpMethodEnum, _ := json.Marshal(dst.HttpMethodEnum)
		if string(jsonHttpMethodEnum) == "{}" { // empty struct
			dst.HttpMethodEnum = nil
		} else {
			if err = validator.Validate(dst.HttpMethodEnum); err != nil {
				dst.HttpMethodEnum = nil
			} else {
				match++
			}
		}
	} else {
		dst.HttpMethodEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.HttpMethodEnum = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BulkWritableExternalIntegrationRequestHttpMethod)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BulkWritableExternalIntegrationRequestHttpMethod)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BulkWritableExternalIntegrationRequestHttpMethod) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.HttpMethodEnum != nil {
		return json.Marshal(&src.HttpMethodEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BulkWritableExternalIntegrationRequestHttpMethod) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.HttpMethodEnum != nil {
		return obj.HttpMethodEnum
	}

	// all schemas are nil
	return nil
}

type NullableBulkWritableExternalIntegrationRequestHttpMethod struct {
	value *BulkWritableExternalIntegrationRequestHttpMethod
	isSet bool
}

func (v NullableBulkWritableExternalIntegrationRequestHttpMethod) Get() *BulkWritableExternalIntegrationRequestHttpMethod {
	return v.value
}

func (v *NullableBulkWritableExternalIntegrationRequestHttpMethod) Set(val *BulkWritableExternalIntegrationRequestHttpMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableExternalIntegrationRequestHttpMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableExternalIntegrationRequestHttpMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableExternalIntegrationRequestHttpMethod(val *BulkWritableExternalIntegrationRequestHttpMethod) *NullableBulkWritableExternalIntegrationRequestHttpMethod {
	return &NullableBulkWritableExternalIntegrationRequestHttpMethod{value: val, isSet: true}
}

func (v NullableBulkWritableExternalIntegrationRequestHttpMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableExternalIntegrationRequestHttpMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


