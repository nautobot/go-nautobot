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

// PortTypeChoices the model 'PortTypeChoices'
type PortTypeChoices string

// List of PortTypeChoices
const (
	PORTTYPECHOICES__8P8C PortTypeChoices = "8p8c"
	PORTTYPECHOICES__8P6C PortTypeChoices = "8p6c"
	PORTTYPECHOICES__8P4C PortTypeChoices = "8p4c"
	PORTTYPECHOICES__8P2C PortTypeChoices = "8p2c"
	PORTTYPECHOICES__6P6C PortTypeChoices = "6p6c"
	PORTTYPECHOICES__6P4C PortTypeChoices = "6p4c"
	PORTTYPECHOICES__6P2C PortTypeChoices = "6p2c"
	PORTTYPECHOICES__4P4C PortTypeChoices = "4p4c"
	PORTTYPECHOICES__4P2C PortTypeChoices = "4p2c"
	PORTTYPECHOICES_GG45 PortTypeChoices = "gg45"
	PORTTYPECHOICES_TERA_4P PortTypeChoices = "tera-4p"
	PORTTYPECHOICES_TERA_2P PortTypeChoices = "tera-2p"
	PORTTYPECHOICES_TERA_1P PortTypeChoices = "tera-1p"
	PORTTYPECHOICES__110_PUNCH PortTypeChoices = "110-punch"
	PORTTYPECHOICES_BNC PortTypeChoices = "bnc"
	PORTTYPECHOICES_F PortTypeChoices = "f"
	PORTTYPECHOICES_N PortTypeChoices = "n"
	PORTTYPECHOICES_MRJ21 PortTypeChoices = "mrj21"
	PORTTYPECHOICES_FC PortTypeChoices = "fc"
	PORTTYPECHOICES_LC PortTypeChoices = "lc"
	PORTTYPECHOICES_LC_PC PortTypeChoices = "lc-pc"
	PORTTYPECHOICES_LC_UPC PortTypeChoices = "lc-upc"
	PORTTYPECHOICES_LC_APC PortTypeChoices = "lc-apc"
	PORTTYPECHOICES_LSH PortTypeChoices = "lsh"
	PORTTYPECHOICES_LSH_PC PortTypeChoices = "lsh-pc"
	PORTTYPECHOICES_LSH_UPC PortTypeChoices = "lsh-upc"
	PORTTYPECHOICES_LSH_APC PortTypeChoices = "lsh-apc"
	PORTTYPECHOICES_LX5 PortTypeChoices = "lx5"
	PORTTYPECHOICES_LX5_PC PortTypeChoices = "lx5-pc"
	PORTTYPECHOICES_LX5_UPC PortTypeChoices = "lx5-upc"
	PORTTYPECHOICES_LX5_APC PortTypeChoices = "lx5-apc"
	PORTTYPECHOICES_MPO PortTypeChoices = "mpo"
	PORTTYPECHOICES_MTRJ PortTypeChoices = "mtrj"
	PORTTYPECHOICES_SC PortTypeChoices = "sc"
	PORTTYPECHOICES_SC_PC PortTypeChoices = "sc-pc"
	PORTTYPECHOICES_SC_UPC PortTypeChoices = "sc-upc"
	PORTTYPECHOICES_SC_APC PortTypeChoices = "sc-apc"
	PORTTYPECHOICES_ST PortTypeChoices = "st"
	PORTTYPECHOICES_CS PortTypeChoices = "cs"
	PORTTYPECHOICES_SN PortTypeChoices = "sn"
	PORTTYPECHOICES_SMA_905 PortTypeChoices = "sma-905"
	PORTTYPECHOICES_SMA_906 PortTypeChoices = "sma-906"
	PORTTYPECHOICES_URM_P2 PortTypeChoices = "urm-p2"
	PORTTYPECHOICES_URM_P4 PortTypeChoices = "urm-p4"
	PORTTYPECHOICES_URM_P8 PortTypeChoices = "urm-p8"
	PORTTYPECHOICES_SPLICE PortTypeChoices = "splice"
	PORTTYPECHOICES_OTHER PortTypeChoices = "other"
)

// All allowed values of PortTypeChoices enum
var AllowedPortTypeChoicesEnumValues = []PortTypeChoices{
	"8p8c",
	"8p6c",
	"8p4c",
	"8p2c",
	"6p6c",
	"6p4c",
	"6p2c",
	"4p4c",
	"4p2c",
	"gg45",
	"tera-4p",
	"tera-2p",
	"tera-1p",
	"110-punch",
	"bnc",
	"f",
	"n",
	"mrj21",
	"fc",
	"lc",
	"lc-pc",
	"lc-upc",
	"lc-apc",
	"lsh",
	"lsh-pc",
	"lsh-upc",
	"lsh-apc",
	"lx5",
	"lx5-pc",
	"lx5-upc",
	"lx5-apc",
	"mpo",
	"mtrj",
	"sc",
	"sc-pc",
	"sc-upc",
	"sc-apc",
	"st",
	"cs",
	"sn",
	"sma-905",
	"sma-906",
	"urm-p2",
	"urm-p4",
	"urm-p8",
	"splice",
	"other",
}

func (v *PortTypeChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortTypeChoices(value)
	for _, existing := range AllowedPortTypeChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortTypeChoices", value)
}

// NewPortTypeChoicesFromValue returns a pointer to a valid PortTypeChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortTypeChoicesFromValue(v string) (*PortTypeChoices, error) {
	ev := PortTypeChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortTypeChoices: valid values are %v", v, AllowedPortTypeChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortTypeChoices) IsValid() bool {
	for _, existing := range AllowedPortTypeChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortTypeChoices value
func (v PortTypeChoices) Ptr() *PortTypeChoices {
	return &v
}

type NullablePortTypeChoices struct {
	value *PortTypeChoices
	isSet bool
}

func (v NullablePortTypeChoices) Get() *PortTypeChoices {
	return v.value
}

func (v *NullablePortTypeChoices) Set(val *PortTypeChoices) {
	v.value = val
	v.isSet = true
}

func (v NullablePortTypeChoices) IsSet() bool {
	return v.isSet
}

func (v *NullablePortTypeChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortTypeChoices(val *PortTypeChoices) *NullablePortTypeChoices {
	return &NullablePortTypeChoices{value: val, isSet: true}
}

func (v NullablePortTypeChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortTypeChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

