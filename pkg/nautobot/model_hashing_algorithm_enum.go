/*
API Documentation

Source of truth and network automation platform

API version: 2.2.5 (2.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// HashingAlgorithmEnum the model 'HashingAlgorithmEnum'
type HashingAlgorithmEnum string

// List of HashingAlgorithmEnum
const (
	HASHINGALGORITHMENUM_MD5 HashingAlgorithmEnum = "md5"
	HASHINGALGORITHMENUM_SHA1 HashingAlgorithmEnum = "sha1"
	HASHINGALGORITHMENUM_SHA224 HashingAlgorithmEnum = "sha224"
	HASHINGALGORITHMENUM_SHA384 HashingAlgorithmEnum = "sha384"
	HASHINGALGORITHMENUM_SHA256 HashingAlgorithmEnum = "sha256"
	HASHINGALGORITHMENUM_SHA512 HashingAlgorithmEnum = "sha512"
	HASHINGALGORITHMENUM_SHA3 HashingAlgorithmEnum = "sha3"
	HASHINGALGORITHMENUM_BLAKE2 HashingAlgorithmEnum = "blake2"
	HASHINGALGORITHMENUM_BLAKE3 HashingAlgorithmEnum = "blake3"
)

// All allowed values of HashingAlgorithmEnum enum
var AllowedHashingAlgorithmEnumEnumValues = []HashingAlgorithmEnum{
	"md5",
	"sha1",
	"sha224",
	"sha384",
	"sha256",
	"sha512",
	"sha3",
	"blake2",
	"blake3",
}

func (v *HashingAlgorithmEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HashingAlgorithmEnum(value)
	for _, existing := range AllowedHashingAlgorithmEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HashingAlgorithmEnum", value)
}

// NewHashingAlgorithmEnumFromValue returns a pointer to a valid HashingAlgorithmEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHashingAlgorithmEnumFromValue(v string) (*HashingAlgorithmEnum, error) {
	ev := HashingAlgorithmEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HashingAlgorithmEnum: valid values are %v", v, AllowedHashingAlgorithmEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HashingAlgorithmEnum) IsValid() bool {
	for _, existing := range AllowedHashingAlgorithmEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HashingAlgorithmEnum value
func (v HashingAlgorithmEnum) Ptr() *HashingAlgorithmEnum {
	return &v
}

type NullableHashingAlgorithmEnum struct {
	value *HashingAlgorithmEnum
	isSet bool
}

func (v NullableHashingAlgorithmEnum) Get() *HashingAlgorithmEnum {
	return v.value
}

func (v *NullableHashingAlgorithmEnum) Set(val *HashingAlgorithmEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableHashingAlgorithmEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableHashingAlgorithmEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashingAlgorithmEnum(val *HashingAlgorithmEnum) *NullableHashingAlgorithmEnum {
	return &NullableHashingAlgorithmEnum{value: val, isSet: true}
}

func (v NullableHashingAlgorithmEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashingAlgorithmEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
