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

// DcimRacksElevationListRenderParameter the model 'DcimRacksElevationListRenderParameter'
type DcimRacksElevationListRenderParameter string

// List of dcim_racks_elevation_list_render_parameter
const (
	DCIMRACKSELEVATIONLISTRENDERPARAMETER_JSON DcimRacksElevationListRenderParameter = "json"
	DCIMRACKSELEVATIONLISTRENDERPARAMETER_SVG DcimRacksElevationListRenderParameter = "svg"
)

// All allowed values of DcimRacksElevationListRenderParameter enum
var AllowedDcimRacksElevationListRenderParameterEnumValues = []DcimRacksElevationListRenderParameter{
	"json",
	"svg",
}

func (v *DcimRacksElevationListRenderParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimRacksElevationListRenderParameter(value)
	for _, existing := range AllowedDcimRacksElevationListRenderParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimRacksElevationListRenderParameter", value)
}

// NewDcimRacksElevationListRenderParameterFromValue returns a pointer to a valid DcimRacksElevationListRenderParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimRacksElevationListRenderParameterFromValue(v string) (*DcimRacksElevationListRenderParameter, error) {
	ev := DcimRacksElevationListRenderParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimRacksElevationListRenderParameter: valid values are %v", v, AllowedDcimRacksElevationListRenderParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimRacksElevationListRenderParameter) IsValid() bool {
	for _, existing := range AllowedDcimRacksElevationListRenderParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_racks_elevation_list_render_parameter value
func (v DcimRacksElevationListRenderParameter) Ptr() *DcimRacksElevationListRenderParameter {
	return &v
}

type NullableDcimRacksElevationListRenderParameter struct {
	value *DcimRacksElevationListRenderParameter
	isSet bool
}

func (v NullableDcimRacksElevationListRenderParameter) Get() *DcimRacksElevationListRenderParameter {
	return v.value
}

func (v *NullableDcimRacksElevationListRenderParameter) Set(val *DcimRacksElevationListRenderParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimRacksElevationListRenderParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimRacksElevationListRenderParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimRacksElevationListRenderParameter(val *DcimRacksElevationListRenderParameter) *NullableDcimRacksElevationListRenderParameter {
	return &NullableDcimRacksElevationListRenderParameter{value: val, isSet: true}
}

func (v NullableDcimRacksElevationListRenderParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimRacksElevationListRenderParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

