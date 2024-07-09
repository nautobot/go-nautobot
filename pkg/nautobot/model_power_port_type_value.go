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

// PowerPortTypeValue the model 'PowerPortTypeValue'
type PowerPortTypeValue string

// List of PowerPort_type_value
const (
	POWERPORTTYPEVALUE_IEC_60320_C6 PowerPortTypeValue = "iec-60320-c6"
	POWERPORTTYPEVALUE_IEC_60320_C8 PowerPortTypeValue = "iec-60320-c8"
	POWERPORTTYPEVALUE_IEC_60320_C14 PowerPortTypeValue = "iec-60320-c14"
	POWERPORTTYPEVALUE_IEC_60320_C16 PowerPortTypeValue = "iec-60320-c16"
	POWERPORTTYPEVALUE_IEC_60320_C20 PowerPortTypeValue = "iec-60320-c20"
	POWERPORTTYPEVALUE_IEC_60320_C22 PowerPortTypeValue = "iec-60320-c22"
	POWERPORTTYPEVALUE_IEC_60309_P_N_E_4H PowerPortTypeValue = "iec-60309-p-n-e-4h"
	POWERPORTTYPEVALUE_IEC_60309_P_N_E_6H PowerPortTypeValue = "iec-60309-p-n-e-6h"
	POWERPORTTYPEVALUE_IEC_60309_P_N_E_9H PowerPortTypeValue = "iec-60309-p-n-e-9h"
	POWERPORTTYPEVALUE_IEC_60309_2P_E_4H PowerPortTypeValue = "iec-60309-2p-e-4h"
	POWERPORTTYPEVALUE_IEC_60309_2P_E_6H PowerPortTypeValue = "iec-60309-2p-e-6h"
	POWERPORTTYPEVALUE_IEC_60309_2P_E_9H PowerPortTypeValue = "iec-60309-2p-e-9h"
	POWERPORTTYPEVALUE_IEC_60309_3P_E_4H PowerPortTypeValue = "iec-60309-3p-e-4h"
	POWERPORTTYPEVALUE_IEC_60309_3P_E_6H PowerPortTypeValue = "iec-60309-3p-e-6h"
	POWERPORTTYPEVALUE_IEC_60309_3P_E_9H PowerPortTypeValue = "iec-60309-3p-e-9h"
	POWERPORTTYPEVALUE_IEC_60309_3P_N_E_4H PowerPortTypeValue = "iec-60309-3p-n-e-4h"
	POWERPORTTYPEVALUE_IEC_60309_3P_N_E_6H PowerPortTypeValue = "iec-60309-3p-n-e-6h"
	POWERPORTTYPEVALUE_IEC_60309_3P_N_E_9H PowerPortTypeValue = "iec-60309-3p-n-e-9h"
	POWERPORTTYPEVALUE_IEC_60906_1 PowerPortTypeValue = "iec-60906-1"
	POWERPORTTYPEVALUE_NBR_14136_10A PowerPortTypeValue = "nbr-14136-10a"
	POWERPORTTYPEVALUE_NBR_14136_20A PowerPortTypeValue = "nbr-14136-20a"
	POWERPORTTYPEVALUE_NEMA_1_15P PowerPortTypeValue = "nema-1-15p"
	POWERPORTTYPEVALUE_NEMA_5_15P PowerPortTypeValue = "nema-5-15p"
	POWERPORTTYPEVALUE_NEMA_5_20P PowerPortTypeValue = "nema-5-20p"
	POWERPORTTYPEVALUE_NEMA_5_30P PowerPortTypeValue = "nema-5-30p"
	POWERPORTTYPEVALUE_NEMA_5_50P PowerPortTypeValue = "nema-5-50p"
	POWERPORTTYPEVALUE_NEMA_6_15P PowerPortTypeValue = "nema-6-15p"
	POWERPORTTYPEVALUE_NEMA_6_20P PowerPortTypeValue = "nema-6-20p"
	POWERPORTTYPEVALUE_NEMA_6_30P PowerPortTypeValue = "nema-6-30p"
	POWERPORTTYPEVALUE_NEMA_6_50P PowerPortTypeValue = "nema-6-50p"
	POWERPORTTYPEVALUE_NEMA_10_30P PowerPortTypeValue = "nema-10-30p"
	POWERPORTTYPEVALUE_NEMA_10_50P PowerPortTypeValue = "nema-10-50p"
	POWERPORTTYPEVALUE_NEMA_14_20P PowerPortTypeValue = "nema-14-20p"
	POWERPORTTYPEVALUE_NEMA_14_30P PowerPortTypeValue = "nema-14-30p"
	POWERPORTTYPEVALUE_NEMA_14_50P PowerPortTypeValue = "nema-14-50p"
	POWERPORTTYPEVALUE_NEMA_14_60P PowerPortTypeValue = "nema-14-60p"
	POWERPORTTYPEVALUE_NEMA_15_15P PowerPortTypeValue = "nema-15-15p"
	POWERPORTTYPEVALUE_NEMA_15_20P PowerPortTypeValue = "nema-15-20p"
	POWERPORTTYPEVALUE_NEMA_15_30P PowerPortTypeValue = "nema-15-30p"
	POWERPORTTYPEVALUE_NEMA_15_50P PowerPortTypeValue = "nema-15-50p"
	POWERPORTTYPEVALUE_NEMA_15_60P PowerPortTypeValue = "nema-15-60p"
	POWERPORTTYPEVALUE_NEMA_L1_15P PowerPortTypeValue = "nema-l1-15p"
	POWERPORTTYPEVALUE_NEMA_L5_15P PowerPortTypeValue = "nema-l5-15p"
	POWERPORTTYPEVALUE_NEMA_L5_20P PowerPortTypeValue = "nema-l5-20p"
	POWERPORTTYPEVALUE_NEMA_L5_30P PowerPortTypeValue = "nema-l5-30p"
	POWERPORTTYPEVALUE_NEMA_L5_50P PowerPortTypeValue = "nema-l5-50p"
	POWERPORTTYPEVALUE_NEMA_L6_15P PowerPortTypeValue = "nema-l6-15p"
	POWERPORTTYPEVALUE_NEMA_L6_20P PowerPortTypeValue = "nema-l6-20p"
	POWERPORTTYPEVALUE_NEMA_L6_30P PowerPortTypeValue = "nema-l6-30p"
	POWERPORTTYPEVALUE_NEMA_L6_50P PowerPortTypeValue = "nema-l6-50p"
	POWERPORTTYPEVALUE_NEMA_L10_30P PowerPortTypeValue = "nema-l10-30p"
	POWERPORTTYPEVALUE_NEMA_L14_20P PowerPortTypeValue = "nema-l14-20p"
	POWERPORTTYPEVALUE_NEMA_L14_30P PowerPortTypeValue = "nema-l14-30p"
	POWERPORTTYPEVALUE_NEMA_L14_50P PowerPortTypeValue = "nema-l14-50p"
	POWERPORTTYPEVALUE_NEMA_L14_60P PowerPortTypeValue = "nema-l14-60p"
	POWERPORTTYPEVALUE_NEMA_L15_20P PowerPortTypeValue = "nema-l15-20p"
	POWERPORTTYPEVALUE_NEMA_L15_30P PowerPortTypeValue = "nema-l15-30p"
	POWERPORTTYPEVALUE_NEMA_L15_50P PowerPortTypeValue = "nema-l15-50p"
	POWERPORTTYPEVALUE_NEMA_L15_60P PowerPortTypeValue = "nema-l15-60p"
	POWERPORTTYPEVALUE_NEMA_L21_20P PowerPortTypeValue = "nema-l21-20p"
	POWERPORTTYPEVALUE_NEMA_L21_30P PowerPortTypeValue = "nema-l21-30p"
	POWERPORTTYPEVALUE_NEMA_L22_30P PowerPortTypeValue = "nema-l22-30p"
	POWERPORTTYPEVALUE_CS6361C PowerPortTypeValue = "cs6361c"
	POWERPORTTYPEVALUE_CS6365C PowerPortTypeValue = "cs6365c"
	POWERPORTTYPEVALUE_CS8165C PowerPortTypeValue = "cs8165c"
	POWERPORTTYPEVALUE_CS8265C PowerPortTypeValue = "cs8265c"
	POWERPORTTYPEVALUE_CS8365C PowerPortTypeValue = "cs8365c"
	POWERPORTTYPEVALUE_CS8465C PowerPortTypeValue = "cs8465c"
	POWERPORTTYPEVALUE_ITA_C PowerPortTypeValue = "ita-c"
	POWERPORTTYPEVALUE_ITA_E PowerPortTypeValue = "ita-e"
	POWERPORTTYPEVALUE_ITA_F PowerPortTypeValue = "ita-f"
	POWERPORTTYPEVALUE_ITA_EF PowerPortTypeValue = "ita-ef"
	POWERPORTTYPEVALUE_ITA_G PowerPortTypeValue = "ita-g"
	POWERPORTTYPEVALUE_ITA_H PowerPortTypeValue = "ita-h"
	POWERPORTTYPEVALUE_ITA_I PowerPortTypeValue = "ita-i"
	POWERPORTTYPEVALUE_ITA_J PowerPortTypeValue = "ita-j"
	POWERPORTTYPEVALUE_ITA_K PowerPortTypeValue = "ita-k"
	POWERPORTTYPEVALUE_ITA_L PowerPortTypeValue = "ita-l"
	POWERPORTTYPEVALUE_ITA_M PowerPortTypeValue = "ita-m"
	POWERPORTTYPEVALUE_ITA_N PowerPortTypeValue = "ita-n"
	POWERPORTTYPEVALUE_ITA_O PowerPortTypeValue = "ita-o"
	POWERPORTTYPEVALUE_USB_A PowerPortTypeValue = "usb-a"
	POWERPORTTYPEVALUE_USB_B PowerPortTypeValue = "usb-b"
	POWERPORTTYPEVALUE_USB_C PowerPortTypeValue = "usb-c"
	POWERPORTTYPEVALUE_USB_MINI_A PowerPortTypeValue = "usb-mini-a"
	POWERPORTTYPEVALUE_USB_MINI_B PowerPortTypeValue = "usb-mini-b"
	POWERPORTTYPEVALUE_USB_MICRO_A PowerPortTypeValue = "usb-micro-a"
	POWERPORTTYPEVALUE_USB_MICRO_B PowerPortTypeValue = "usb-micro-b"
	POWERPORTTYPEVALUE_USB_MICRO_AB PowerPortTypeValue = "usb-micro-ab"
	POWERPORTTYPEVALUE_USB_3_B PowerPortTypeValue = "usb-3-b"
	POWERPORTTYPEVALUE_USB_3_MICRO_B PowerPortTypeValue = "usb-3-micro-b"
	POWERPORTTYPEVALUE_DC_TERMINAL PowerPortTypeValue = "dc-terminal"
	POWERPORTTYPEVALUE_SAF_D_GRID PowerPortTypeValue = "saf-d-grid"
	POWERPORTTYPEVALUE_NEUTRIK_POWERCON_20 PowerPortTypeValue = "neutrik-powercon-20"
	POWERPORTTYPEVALUE_NEUTRIK_POWERCON_32 PowerPortTypeValue = "neutrik-powercon-32"
	POWERPORTTYPEVALUE_NEUTRIK_POWERCON_TRUE1 PowerPortTypeValue = "neutrik-powercon-true1"
	POWERPORTTYPEVALUE_NEUTRIK_POWERCON_TRUE1_TOP PowerPortTypeValue = "neutrik-powercon-true1-top"
	POWERPORTTYPEVALUE_UBIQUITI_SMARTPOWER PowerPortTypeValue = "ubiquiti-smartpower"
	POWERPORTTYPEVALUE_HARDWIRED PowerPortTypeValue = "hardwired"
	POWERPORTTYPEVALUE_OTHER PowerPortTypeValue = "other"
)

// All allowed values of PowerPortTypeValue enum
var AllowedPowerPortTypeValueEnumValues = []PowerPortTypeValue{
	"iec-60320-c6",
	"iec-60320-c8",
	"iec-60320-c14",
	"iec-60320-c16",
	"iec-60320-c20",
	"iec-60320-c22",
	"iec-60309-p-n-e-4h",
	"iec-60309-p-n-e-6h",
	"iec-60309-p-n-e-9h",
	"iec-60309-2p-e-4h",
	"iec-60309-2p-e-6h",
	"iec-60309-2p-e-9h",
	"iec-60309-3p-e-4h",
	"iec-60309-3p-e-6h",
	"iec-60309-3p-e-9h",
	"iec-60309-3p-n-e-4h",
	"iec-60309-3p-n-e-6h",
	"iec-60309-3p-n-e-9h",
	"iec-60906-1",
	"nbr-14136-10a",
	"nbr-14136-20a",
	"nema-1-15p",
	"nema-5-15p",
	"nema-5-20p",
	"nema-5-30p",
	"nema-5-50p",
	"nema-6-15p",
	"nema-6-20p",
	"nema-6-30p",
	"nema-6-50p",
	"nema-10-30p",
	"nema-10-50p",
	"nema-14-20p",
	"nema-14-30p",
	"nema-14-50p",
	"nema-14-60p",
	"nema-15-15p",
	"nema-15-20p",
	"nema-15-30p",
	"nema-15-50p",
	"nema-15-60p",
	"nema-l1-15p",
	"nema-l5-15p",
	"nema-l5-20p",
	"nema-l5-30p",
	"nema-l5-50p",
	"nema-l6-15p",
	"nema-l6-20p",
	"nema-l6-30p",
	"nema-l6-50p",
	"nema-l10-30p",
	"nema-l14-20p",
	"nema-l14-30p",
	"nema-l14-50p",
	"nema-l14-60p",
	"nema-l15-20p",
	"nema-l15-30p",
	"nema-l15-50p",
	"nema-l15-60p",
	"nema-l21-20p",
	"nema-l21-30p",
	"nema-l22-30p",
	"cs6361c",
	"cs6365c",
	"cs8165c",
	"cs8265c",
	"cs8365c",
	"cs8465c",
	"ita-c",
	"ita-e",
	"ita-f",
	"ita-ef",
	"ita-g",
	"ita-h",
	"ita-i",
	"ita-j",
	"ita-k",
	"ita-l",
	"ita-m",
	"ita-n",
	"ita-o",
	"usb-a",
	"usb-b",
	"usb-c",
	"usb-mini-a",
	"usb-mini-b",
	"usb-micro-a",
	"usb-micro-b",
	"usb-micro-ab",
	"usb-3-b",
	"usb-3-micro-b",
	"dc-terminal",
	"saf-d-grid",
	"neutrik-powercon-20",
	"neutrik-powercon-32",
	"neutrik-powercon-true1",
	"neutrik-powercon-true1-top",
	"ubiquiti-smartpower",
	"hardwired",
	"other",
}

func (v *PowerPortTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerPortTypeValue(value)
	for _, existing := range AllowedPowerPortTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerPortTypeValue", value)
}

// NewPowerPortTypeValueFromValue returns a pointer to a valid PowerPortTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerPortTypeValueFromValue(v string) (*PowerPortTypeValue, error) {
	ev := PowerPortTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerPortTypeValue: valid values are %v", v, AllowedPowerPortTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerPortTypeValue) IsValid() bool {
	for _, existing := range AllowedPowerPortTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerPort_type_value value
func (v PowerPortTypeValue) Ptr() *PowerPortTypeValue {
	return &v
}

type NullablePowerPortTypeValue struct {
	value *PowerPortTypeValue
	isSet bool
}

func (v NullablePowerPortTypeValue) Get() *PowerPortTypeValue {
	return v.value
}

func (v *NullablePowerPortTypeValue) Set(val *PowerPortTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerPortTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerPortTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerPortTypeValue(val *PowerPortTypeValue) *NullablePowerPortTypeValue {
	return &NullablePowerPortTypeValue{value: val, isSet: true}
}

func (v NullablePowerPortTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerPortTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
