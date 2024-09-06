/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// FrontPortTypeLabel the model 'FrontPortTypeLabel'
type FrontPortTypeLabel string

// List of FrontPort_type_label
const (
	FRONTPORTTYPELABEL__8_P8_C FrontPortTypeLabel = "8P8C"
	FRONTPORTTYPELABEL__8_P6_C FrontPortTypeLabel = "8P6C"
	FRONTPORTTYPELABEL__8_P4_C FrontPortTypeLabel = "8P4C"
	FRONTPORTTYPELABEL__8_P2_C FrontPortTypeLabel = "8P2C"
	FRONTPORTTYPELABEL__6_P6_C FrontPortTypeLabel = "6P6C"
	FRONTPORTTYPELABEL__6_P4_C FrontPortTypeLabel = "6P4C"
	FRONTPORTTYPELABEL__6_P2_C FrontPortTypeLabel = "6P2C"
	FRONTPORTTYPELABEL__4_P4_C FrontPortTypeLabel = "4P4C"
	FRONTPORTTYPELABEL__4_P2_C FrontPortTypeLabel = "4P2C"
	FRONTPORTTYPELABEL_GG45 FrontPortTypeLabel = "GG45"
	FRONTPORTTYPELABEL_TERA_4_P FrontPortTypeLabel = "TERA 4P"
	FRONTPORTTYPELABEL_TERA_2_P FrontPortTypeLabel = "TERA 2P"
	FRONTPORTTYPELABEL_TERA_1_P FrontPortTypeLabel = "TERA 1P"
	FRONTPORTTYPELABEL__110_PUNCH FrontPortTypeLabel = "110 Punch"
	FRONTPORTTYPELABEL_BNC FrontPortTypeLabel = "BNC"
	FRONTPORTTYPELABEL_F_CONNECTOR FrontPortTypeLabel = "F Connector"
	FRONTPORTTYPELABEL_N_CONNECTOR FrontPortTypeLabel = "N Connector"
	FRONTPORTTYPELABEL_MRJ21 FrontPortTypeLabel = "MRJ21"
	FRONTPORTTYPELABEL_FC FrontPortTypeLabel = "FC"
	FRONTPORTTYPELABEL_LC FrontPortTypeLabel = "LC"
	FRONTPORTTYPELABEL_LC_PC FrontPortTypeLabel = "LC/PC"
	FRONTPORTTYPELABEL_LC_UPC FrontPortTypeLabel = "LC/UPC"
	FRONTPORTTYPELABEL_LC_APC FrontPortTypeLabel = "LC/APC"
	FRONTPORTTYPELABEL_LSH FrontPortTypeLabel = "LSH"
	FRONTPORTTYPELABEL_LSH_PC FrontPortTypeLabel = "LSH/PC"
	FRONTPORTTYPELABEL_LSH_UPC FrontPortTypeLabel = "LSH/UPC"
	FRONTPORTTYPELABEL_LSH_APC FrontPortTypeLabel = "LSH/APC"
	FRONTPORTTYPELABEL_LX_5 FrontPortTypeLabel = "LX.5"
	FRONTPORTTYPELABEL_LX_5_PC FrontPortTypeLabel = "LX.5/PC"
	FRONTPORTTYPELABEL_LX_5_UPC FrontPortTypeLabel = "LX.5/UPC"
	FRONTPORTTYPELABEL_LX_5_APC FrontPortTypeLabel = "LX.5/APC"
	FRONTPORTTYPELABEL_MPO FrontPortTypeLabel = "MPO"
	FRONTPORTTYPELABEL_MTRJ FrontPortTypeLabel = "MTRJ"
	FRONTPORTTYPELABEL_SC FrontPortTypeLabel = "SC"
	FRONTPORTTYPELABEL_SC_PC FrontPortTypeLabel = "SC/PC"
	FRONTPORTTYPELABEL_SC_UPC FrontPortTypeLabel = "SC/UPC"
	FRONTPORTTYPELABEL_SC_APC FrontPortTypeLabel = "SC/APC"
	FRONTPORTTYPELABEL_ST FrontPortTypeLabel = "ST"
	FRONTPORTTYPELABEL_CS FrontPortTypeLabel = "CS"
	FRONTPORTTYPELABEL_SN FrontPortTypeLabel = "SN"
	FRONTPORTTYPELABEL_SMA_905 FrontPortTypeLabel = "SMA 905"
	FRONTPORTTYPELABEL_SMA_906 FrontPortTypeLabel = "SMA 906"
	FRONTPORTTYPELABEL_URM_P2 FrontPortTypeLabel = "URM-P2"
	FRONTPORTTYPELABEL_URM_P4 FrontPortTypeLabel = "URM-P4"
	FRONTPORTTYPELABEL_URM_P8 FrontPortTypeLabel = "URM-P8"
	FRONTPORTTYPELABEL_SPLICE FrontPortTypeLabel = "Splice"
	FRONTPORTTYPELABEL_OTHER FrontPortTypeLabel = "Other"
)

// All allowed values of FrontPortTypeLabel enum
var AllowedFrontPortTypeLabelEnumValues = []FrontPortTypeLabel{
	"8P8C",
	"8P6C",
	"8P4C",
	"8P2C",
	"6P6C",
	"6P4C",
	"6P2C",
	"4P4C",
	"4P2C",
	"GG45",
	"TERA 4P",
	"TERA 2P",
	"TERA 1P",
	"110 Punch",
	"BNC",
	"F Connector",
	"N Connector",
	"MRJ21",
	"FC",
	"LC",
	"LC/PC",
	"LC/UPC",
	"LC/APC",
	"LSH",
	"LSH/PC",
	"LSH/UPC",
	"LSH/APC",
	"LX.5",
	"LX.5/PC",
	"LX.5/UPC",
	"LX.5/APC",
	"MPO",
	"MTRJ",
	"SC",
	"SC/PC",
	"SC/UPC",
	"SC/APC",
	"ST",
	"CS",
	"SN",
	"SMA 905",
	"SMA 906",
	"URM-P2",
	"URM-P4",
	"URM-P8",
	"Splice",
	"Other",
}

func (v *FrontPortTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FrontPortTypeLabel(value)
	for _, existing := range AllowedFrontPortTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FrontPortTypeLabel", value)
}

// NewFrontPortTypeLabelFromValue returns a pointer to a valid FrontPortTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFrontPortTypeLabelFromValue(v string) (*FrontPortTypeLabel, error) {
	ev := FrontPortTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FrontPortTypeLabel: valid values are %v", v, AllowedFrontPortTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FrontPortTypeLabel) IsValid() bool {
	for _, existing := range AllowedFrontPortTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FrontPort_type_label value
func (v FrontPortTypeLabel) Ptr() *FrontPortTypeLabel {
	return &v
}

type NullableFrontPortTypeLabel struct {
	value *FrontPortTypeLabel
	isSet bool
}

func (v NullableFrontPortTypeLabel) Get() *FrontPortTypeLabel {
	return v.value
}

func (v *NullableFrontPortTypeLabel) Set(val *FrontPortTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontPortTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontPortTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontPortTypeLabel(val *FrontPortTypeLabel) *NullableFrontPortTypeLabel {
	return &NullableFrontPortTypeLabel{value: val, isSet: true}
}

func (v NullableFrontPortTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontPortTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

