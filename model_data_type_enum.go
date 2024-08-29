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

// DataTypeEnum the model 'DataTypeEnum'
type DataTypeEnum string

// List of DataTypeEnum
const (
	DATATYPEENUM_TEXT DataTypeEnum = "text"
	DATATYPEENUM_INTEGER DataTypeEnum = "integer"
	DATATYPEENUM_BOOLEAN DataTypeEnum = "boolean"
	DATATYPEENUM_DATE DataTypeEnum = "date"
	DATATYPEENUM_URL DataTypeEnum = "url"
	DATATYPEENUM_SELECT DataTypeEnum = "select"
	DATATYPEENUM_MULTI_SELECT DataTypeEnum = "multi-select"
	DATATYPEENUM_JSON DataTypeEnum = "json"
	DATATYPEENUM_MARKDOWN DataTypeEnum = "markdown"
	DATATYPEENUM_CONTACT_OR_TEAM DataTypeEnum = "contact-or-team"
	DATATYPEENUM_DATETIME DataTypeEnum = "datetime"
	DATATYPEENUM_FLOAT DataTypeEnum = "float"
)

// All allowed values of DataTypeEnum enum
var AllowedDataTypeEnumEnumValues = []DataTypeEnum{
	"text",
	"integer",
	"boolean",
	"date",
	"url",
	"select",
	"multi-select",
	"json",
	"markdown",
	"contact-or-team",
	"datetime",
	"float",
}

func (v *DataTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataTypeEnum(value)
	for _, existing := range AllowedDataTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataTypeEnum", value)
}

// NewDataTypeEnumFromValue returns a pointer to a valid DataTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataTypeEnumFromValue(v string) (*DataTypeEnum, error) {
	ev := DataTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataTypeEnum: valid values are %v", v, AllowedDataTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataTypeEnum) IsValid() bool {
	for _, existing := range AllowedDataTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataTypeEnum value
func (v DataTypeEnum) Ptr() *DataTypeEnum {
	return &v
}

type NullableDataTypeEnum struct {
	value *DataTypeEnum
	isSet bool
}

func (v NullableDataTypeEnum) Get() *DataTypeEnum {
	return v.value
}

func (v *NullableDataTypeEnum) Set(val *DataTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTypeEnum(val *DataTypeEnum) *NullableDataTypeEnum {
	return &NullableDataTypeEnum{value: val, isSet: true}
}

func (v NullableDataTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
