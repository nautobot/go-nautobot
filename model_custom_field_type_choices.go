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

// CustomFieldTypeChoices the model 'CustomFieldTypeChoices'
type CustomFieldTypeChoices string

// List of CustomFieldTypeChoices
const (
	CUSTOMFIELDTYPECHOICES_TEXT CustomFieldTypeChoices = "text"
	CUSTOMFIELDTYPECHOICES_INTEGER CustomFieldTypeChoices = "integer"
	CUSTOMFIELDTYPECHOICES_BOOLEAN CustomFieldTypeChoices = "boolean"
	CUSTOMFIELDTYPECHOICES_DATE CustomFieldTypeChoices = "date"
	CUSTOMFIELDTYPECHOICES_URL CustomFieldTypeChoices = "url"
	CUSTOMFIELDTYPECHOICES_SELECT CustomFieldTypeChoices = "select"
	CUSTOMFIELDTYPECHOICES_MULTI_SELECT CustomFieldTypeChoices = "multi-select"
	CUSTOMFIELDTYPECHOICES_JSON CustomFieldTypeChoices = "json"
	CUSTOMFIELDTYPECHOICES_MARKDOWN CustomFieldTypeChoices = "markdown"
)

// All allowed values of CustomFieldTypeChoices enum
var AllowedCustomFieldTypeChoicesEnumValues = []CustomFieldTypeChoices{
	"text",
	"integer",
	"boolean",
	"date",
	"url",
	"select",
	"multi-select",
	"json",
	"markdown",
}

func (v *CustomFieldTypeChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomFieldTypeChoices(value)
	for _, existing := range AllowedCustomFieldTypeChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomFieldTypeChoices", value)
}

// NewCustomFieldTypeChoicesFromValue returns a pointer to a valid CustomFieldTypeChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomFieldTypeChoicesFromValue(v string) (*CustomFieldTypeChoices, error) {
	ev := CustomFieldTypeChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomFieldTypeChoices: valid values are %v", v, AllowedCustomFieldTypeChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldTypeChoices) IsValid() bool {
	for _, existing := range AllowedCustomFieldTypeChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomFieldTypeChoices value
func (v CustomFieldTypeChoices) Ptr() *CustomFieldTypeChoices {
	return &v
}

type NullableCustomFieldTypeChoices struct {
	value *CustomFieldTypeChoices
	isSet bool
}

func (v NullableCustomFieldTypeChoices) Get() *CustomFieldTypeChoices {
	return v.value
}

func (v *NullableCustomFieldTypeChoices) Set(val *CustomFieldTypeChoices) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldTypeChoices) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldTypeChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldTypeChoices(val *CustomFieldTypeChoices) *NullableCustomFieldTypeChoices {
	return &NullableCustomFieldTypeChoices{value: val, isSet: true}
}

func (v NullableCustomFieldTypeChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldTypeChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

