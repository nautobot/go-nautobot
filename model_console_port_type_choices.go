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

// ConsolePortTypeChoices the model 'ConsolePortTypeChoices'
type ConsolePortTypeChoices string

// List of ConsolePortTypeChoices
const (
	CONSOLEPORTTYPECHOICES_DE_9 ConsolePortTypeChoices = "de-9"
	CONSOLEPORTTYPECHOICES_DB_25 ConsolePortTypeChoices = "db-25"
	CONSOLEPORTTYPECHOICES_RJ_11 ConsolePortTypeChoices = "rj-11"
	CONSOLEPORTTYPECHOICES_RJ_12 ConsolePortTypeChoices = "rj-12"
	CONSOLEPORTTYPECHOICES_RJ_45 ConsolePortTypeChoices = "rj-45"
	CONSOLEPORTTYPECHOICES_MINI_DIN_8 ConsolePortTypeChoices = "mini-din-8"
	CONSOLEPORTTYPECHOICES_USB_A ConsolePortTypeChoices = "usb-a"
	CONSOLEPORTTYPECHOICES_USB_B ConsolePortTypeChoices = "usb-b"
	CONSOLEPORTTYPECHOICES_USB_C ConsolePortTypeChoices = "usb-c"
	CONSOLEPORTTYPECHOICES_USB_MINI_A ConsolePortTypeChoices = "usb-mini-a"
	CONSOLEPORTTYPECHOICES_USB_MINI_B ConsolePortTypeChoices = "usb-mini-b"
	CONSOLEPORTTYPECHOICES_USB_MICRO_A ConsolePortTypeChoices = "usb-micro-a"
	CONSOLEPORTTYPECHOICES_USB_MICRO_B ConsolePortTypeChoices = "usb-micro-b"
	CONSOLEPORTTYPECHOICES_USB_MICRO_AB ConsolePortTypeChoices = "usb-micro-ab"
	CONSOLEPORTTYPECHOICES_OTHER ConsolePortTypeChoices = "other"
)

// All allowed values of ConsolePortTypeChoices enum
var AllowedConsolePortTypeChoicesEnumValues = []ConsolePortTypeChoices{
	"de-9",
	"db-25",
	"rj-11",
	"rj-12",
	"rj-45",
	"mini-din-8",
	"usb-a",
	"usb-b",
	"usb-c",
	"usb-mini-a",
	"usb-mini-b",
	"usb-micro-a",
	"usb-micro-b",
	"usb-micro-ab",
	"other",
}

func (v *ConsolePortTypeChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsolePortTypeChoices(value)
	for _, existing := range AllowedConsolePortTypeChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsolePortTypeChoices", value)
}

// NewConsolePortTypeChoicesFromValue returns a pointer to a valid ConsolePortTypeChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsolePortTypeChoicesFromValue(v string) (*ConsolePortTypeChoices, error) {
	ev := ConsolePortTypeChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsolePortTypeChoices: valid values are %v", v, AllowedConsolePortTypeChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsolePortTypeChoices) IsValid() bool {
	for _, existing := range AllowedConsolePortTypeChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsolePortTypeChoices value
func (v ConsolePortTypeChoices) Ptr() *ConsolePortTypeChoices {
	return &v
}

type NullableConsolePortTypeChoices struct {
	value *ConsolePortTypeChoices
	isSet bool
}

func (v NullableConsolePortTypeChoices) Get() *ConsolePortTypeChoices {
	return v.value
}

func (v *NullableConsolePortTypeChoices) Set(val *ConsolePortTypeChoices) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePortTypeChoices) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePortTypeChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePortTypeChoices(val *ConsolePortTypeChoices) *NullableConsolePortTypeChoices {
	return &NullableConsolePortTypeChoices{value: val, isSet: true}
}

func (v NullableConsolePortTypeChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePortTypeChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

