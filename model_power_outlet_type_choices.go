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

// PowerOutletTypeChoices the model 'PowerOutletTypeChoices'
type PowerOutletTypeChoices string

// List of PowerOutletTypeChoices
const (
	POWEROUTLETTYPECHOICES_IEC_60320_C5 PowerOutletTypeChoices = "iec-60320-c5"
	POWEROUTLETTYPECHOICES_IEC_60320_C7 PowerOutletTypeChoices = "iec-60320-c7"
	POWEROUTLETTYPECHOICES_IEC_60320_C13 PowerOutletTypeChoices = "iec-60320-c13"
	POWEROUTLETTYPECHOICES_IEC_60320_C15 PowerOutletTypeChoices = "iec-60320-c15"
	POWEROUTLETTYPECHOICES_IEC_60320_C19 PowerOutletTypeChoices = "iec-60320-c19"
	POWEROUTLETTYPECHOICES_IEC_60320_C21 PowerOutletTypeChoices = "iec-60320-c21"
	POWEROUTLETTYPECHOICES_IEC_60309_P_N_E_4H PowerOutletTypeChoices = "iec-60309-p-n-e-4h"
	POWEROUTLETTYPECHOICES_IEC_60309_P_N_E_6H PowerOutletTypeChoices = "iec-60309-p-n-e-6h"
	POWEROUTLETTYPECHOICES_IEC_60309_P_N_E_9H PowerOutletTypeChoices = "iec-60309-p-n-e-9h"
	POWEROUTLETTYPECHOICES_IEC_60309_2P_E_4H PowerOutletTypeChoices = "iec-60309-2p-e-4h"
	POWEROUTLETTYPECHOICES_IEC_60309_2P_E_6H PowerOutletTypeChoices = "iec-60309-2p-e-6h"
	POWEROUTLETTYPECHOICES_IEC_60309_2P_E_9H PowerOutletTypeChoices = "iec-60309-2p-e-9h"
	POWEROUTLETTYPECHOICES_IEC_60309_3P_E_4H PowerOutletTypeChoices = "iec-60309-3p-e-4h"
	POWEROUTLETTYPECHOICES_IEC_60309_3P_E_6H PowerOutletTypeChoices = "iec-60309-3p-e-6h"
	POWEROUTLETTYPECHOICES_IEC_60309_3P_E_9H PowerOutletTypeChoices = "iec-60309-3p-e-9h"
	POWEROUTLETTYPECHOICES_IEC_60309_3P_N_E_4H PowerOutletTypeChoices = "iec-60309-3p-n-e-4h"
	POWEROUTLETTYPECHOICES_IEC_60309_3P_N_E_6H PowerOutletTypeChoices = "iec-60309-3p-n-e-6h"
	POWEROUTLETTYPECHOICES_IEC_60309_3P_N_E_9H PowerOutletTypeChoices = "iec-60309-3p-n-e-9h"
	POWEROUTLETTYPECHOICES_IEC_60906_1 PowerOutletTypeChoices = "iec-60906-1"
	POWEROUTLETTYPECHOICES_NBR_14136_10A PowerOutletTypeChoices = "nbr-14136-10a"
	POWEROUTLETTYPECHOICES_NBR_14136_20A PowerOutletTypeChoices = "nbr-14136-20a"
	POWEROUTLETTYPECHOICES_NEMA_1_15R PowerOutletTypeChoices = "nema-1-15r"
	POWEROUTLETTYPECHOICES_NEMA_5_15R PowerOutletTypeChoices = "nema-5-15r"
	POWEROUTLETTYPECHOICES_NEMA_5_20R PowerOutletTypeChoices = "nema-5-20r"
	POWEROUTLETTYPECHOICES_NEMA_5_30R PowerOutletTypeChoices = "nema-5-30r"
	POWEROUTLETTYPECHOICES_NEMA_5_50R PowerOutletTypeChoices = "nema-5-50r"
	POWEROUTLETTYPECHOICES_NEMA_6_15R PowerOutletTypeChoices = "nema-6-15r"
	POWEROUTLETTYPECHOICES_NEMA_6_20R PowerOutletTypeChoices = "nema-6-20r"
	POWEROUTLETTYPECHOICES_NEMA_6_30R PowerOutletTypeChoices = "nema-6-30r"
	POWEROUTLETTYPECHOICES_NEMA_6_50R PowerOutletTypeChoices = "nema-6-50r"
	POWEROUTLETTYPECHOICES_NEMA_10_30R PowerOutletTypeChoices = "nema-10-30r"
	POWEROUTLETTYPECHOICES_NEMA_10_50R PowerOutletTypeChoices = "nema-10-50r"
	POWEROUTLETTYPECHOICES_NEMA_14_20R PowerOutletTypeChoices = "nema-14-20r"
	POWEROUTLETTYPECHOICES_NEMA_14_30R PowerOutletTypeChoices = "nema-14-30r"
	POWEROUTLETTYPECHOICES_NEMA_14_50R PowerOutletTypeChoices = "nema-14-50r"
	POWEROUTLETTYPECHOICES_NEMA_14_60R PowerOutletTypeChoices = "nema-14-60r"
	POWEROUTLETTYPECHOICES_NEMA_15_15R PowerOutletTypeChoices = "nema-15-15r"
	POWEROUTLETTYPECHOICES_NEMA_15_20R PowerOutletTypeChoices = "nema-15-20r"
	POWEROUTLETTYPECHOICES_NEMA_15_30R PowerOutletTypeChoices = "nema-15-30r"
	POWEROUTLETTYPECHOICES_NEMA_15_50R PowerOutletTypeChoices = "nema-15-50r"
	POWEROUTLETTYPECHOICES_NEMA_15_60R PowerOutletTypeChoices = "nema-15-60r"
	POWEROUTLETTYPECHOICES_NEMA_L1_15R PowerOutletTypeChoices = "nema-l1-15r"
	POWEROUTLETTYPECHOICES_NEMA_L5_15R PowerOutletTypeChoices = "nema-l5-15r"
	POWEROUTLETTYPECHOICES_NEMA_L5_20R PowerOutletTypeChoices = "nema-l5-20r"
	POWEROUTLETTYPECHOICES_NEMA_L5_30R PowerOutletTypeChoices = "nema-l5-30r"
	POWEROUTLETTYPECHOICES_NEMA_L5_50R PowerOutletTypeChoices = "nema-l5-50r"
	POWEROUTLETTYPECHOICES_NEMA_L6_15R PowerOutletTypeChoices = "nema-l6-15r"
	POWEROUTLETTYPECHOICES_NEMA_L6_20R PowerOutletTypeChoices = "nema-l6-20r"
	POWEROUTLETTYPECHOICES_NEMA_L6_30R PowerOutletTypeChoices = "nema-l6-30r"
	POWEROUTLETTYPECHOICES_NEMA_L6_50R PowerOutletTypeChoices = "nema-l6-50r"
	POWEROUTLETTYPECHOICES_NEMA_L10_30R PowerOutletTypeChoices = "nema-l10-30r"
	POWEROUTLETTYPECHOICES_NEMA_L14_20R PowerOutletTypeChoices = "nema-l14-20r"
	POWEROUTLETTYPECHOICES_NEMA_L14_30R PowerOutletTypeChoices = "nema-l14-30r"
	POWEROUTLETTYPECHOICES_NEMA_L14_50R PowerOutletTypeChoices = "nema-l14-50r"
	POWEROUTLETTYPECHOICES_NEMA_L14_60R PowerOutletTypeChoices = "nema-l14-60r"
	POWEROUTLETTYPECHOICES_NEMA_L15_20R PowerOutletTypeChoices = "nema-l15-20r"
	POWEROUTLETTYPECHOICES_NEMA_L15_30R PowerOutletTypeChoices = "nema-l15-30r"
	POWEROUTLETTYPECHOICES_NEMA_L15_50R PowerOutletTypeChoices = "nema-l15-50r"
	POWEROUTLETTYPECHOICES_NEMA_L15_60R PowerOutletTypeChoices = "nema-l15-60r"
	POWEROUTLETTYPECHOICES_NEMA_L21_20R PowerOutletTypeChoices = "nema-l21-20r"
	POWEROUTLETTYPECHOICES_NEMA_L21_30R PowerOutletTypeChoices = "nema-l21-30r"
	POWEROUTLETTYPECHOICES_NEMA_L22_30R PowerOutletTypeChoices = "nema-l22-30r"
	POWEROUTLETTYPECHOICES_CS6360_C PowerOutletTypeChoices = "CS6360C"
	POWEROUTLETTYPECHOICES_CS6364_C PowerOutletTypeChoices = "CS6364C"
	POWEROUTLETTYPECHOICES_CS8164_C PowerOutletTypeChoices = "CS8164C"
	POWEROUTLETTYPECHOICES_CS8264_C PowerOutletTypeChoices = "CS8264C"
	POWEROUTLETTYPECHOICES_CS8364_C PowerOutletTypeChoices = "CS8364C"
	POWEROUTLETTYPECHOICES_CS8464_C PowerOutletTypeChoices = "CS8464C"
	POWEROUTLETTYPECHOICES_ITA_E PowerOutletTypeChoices = "ita-e"
	POWEROUTLETTYPECHOICES_ITA_F PowerOutletTypeChoices = "ita-f"
	POWEROUTLETTYPECHOICES_ITA_G PowerOutletTypeChoices = "ita-g"
	POWEROUTLETTYPECHOICES_ITA_H PowerOutletTypeChoices = "ita-h"
	POWEROUTLETTYPECHOICES_ITA_I PowerOutletTypeChoices = "ita-i"
	POWEROUTLETTYPECHOICES_ITA_J PowerOutletTypeChoices = "ita-j"
	POWEROUTLETTYPECHOICES_ITA_K PowerOutletTypeChoices = "ita-k"
	POWEROUTLETTYPECHOICES_ITA_L PowerOutletTypeChoices = "ita-l"
	POWEROUTLETTYPECHOICES_ITA_M PowerOutletTypeChoices = "ita-m"
	POWEROUTLETTYPECHOICES_ITA_N PowerOutletTypeChoices = "ita-n"
	POWEROUTLETTYPECHOICES_ITA_O PowerOutletTypeChoices = "ita-o"
	POWEROUTLETTYPECHOICES_ITA_MULTISTANDARD PowerOutletTypeChoices = "ita-multistandard"
	POWEROUTLETTYPECHOICES_USB_A PowerOutletTypeChoices = "usb-a"
	POWEROUTLETTYPECHOICES_USB_MICRO_B PowerOutletTypeChoices = "usb-micro-b"
	POWEROUTLETTYPECHOICES_USB_C PowerOutletTypeChoices = "usb-c"
	POWEROUTLETTYPECHOICES_DC_TERMINAL PowerOutletTypeChoices = "dc-terminal"
	POWEROUTLETTYPECHOICES_HDOT_CX PowerOutletTypeChoices = "hdot-cx"
	POWEROUTLETTYPECHOICES_SAF_D_GRID PowerOutletTypeChoices = "saf-d-grid"
	POWEROUTLETTYPECHOICES_NEUTRIK_POWERCON_20A PowerOutletTypeChoices = "neutrik-powercon-20a"
	POWEROUTLETTYPECHOICES_NEUTRIK_POWERCON_32A PowerOutletTypeChoices = "neutrik-powercon-32a"
	POWEROUTLETTYPECHOICES_NEUTRIK_POWERCON_TRUE1 PowerOutletTypeChoices = "neutrik-powercon-true1"
	POWEROUTLETTYPECHOICES_NEUTRIK_POWERCON_TRUE1_TOP PowerOutletTypeChoices = "neutrik-powercon-true1-top"
	POWEROUTLETTYPECHOICES_UBIQUITI_SMARTPOWER PowerOutletTypeChoices = "ubiquiti-smartpower"
	POWEROUTLETTYPECHOICES_HARDWIRED PowerOutletTypeChoices = "hardwired"
	POWEROUTLETTYPECHOICES_OTHER PowerOutletTypeChoices = "other"
)

// All allowed values of PowerOutletTypeChoices enum
var AllowedPowerOutletTypeChoicesEnumValues = []PowerOutletTypeChoices{
	"iec-60320-c5",
	"iec-60320-c7",
	"iec-60320-c13",
	"iec-60320-c15",
	"iec-60320-c19",
	"iec-60320-c21",
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
	"nema-1-15r",
	"nema-5-15r",
	"nema-5-20r",
	"nema-5-30r",
	"nema-5-50r",
	"nema-6-15r",
	"nema-6-20r",
	"nema-6-30r",
	"nema-6-50r",
	"nema-10-30r",
	"nema-10-50r",
	"nema-14-20r",
	"nema-14-30r",
	"nema-14-50r",
	"nema-14-60r",
	"nema-15-15r",
	"nema-15-20r",
	"nema-15-30r",
	"nema-15-50r",
	"nema-15-60r",
	"nema-l1-15r",
	"nema-l5-15r",
	"nema-l5-20r",
	"nema-l5-30r",
	"nema-l5-50r",
	"nema-l6-15r",
	"nema-l6-20r",
	"nema-l6-30r",
	"nema-l6-50r",
	"nema-l10-30r",
	"nema-l14-20r",
	"nema-l14-30r",
	"nema-l14-50r",
	"nema-l14-60r",
	"nema-l15-20r",
	"nema-l15-30r",
	"nema-l15-50r",
	"nema-l15-60r",
	"nema-l21-20r",
	"nema-l21-30r",
	"nema-l22-30r",
	"CS6360C",
	"CS6364C",
	"CS8164C",
	"CS8264C",
	"CS8364C",
	"CS8464C",
	"ita-e",
	"ita-f",
	"ita-g",
	"ita-h",
	"ita-i",
	"ita-j",
	"ita-k",
	"ita-l",
	"ita-m",
	"ita-n",
	"ita-o",
	"ita-multistandard",
	"usb-a",
	"usb-micro-b",
	"usb-c",
	"dc-terminal",
	"hdot-cx",
	"saf-d-grid",
	"neutrik-powercon-20a",
	"neutrik-powercon-32a",
	"neutrik-powercon-true1",
	"neutrik-powercon-true1-top",
	"ubiquiti-smartpower",
	"hardwired",
	"other",
}

func (v *PowerOutletTypeChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerOutletTypeChoices(value)
	for _, existing := range AllowedPowerOutletTypeChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerOutletTypeChoices", value)
}

// NewPowerOutletTypeChoicesFromValue returns a pointer to a valid PowerOutletTypeChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerOutletTypeChoicesFromValue(v string) (*PowerOutletTypeChoices, error) {
	ev := PowerOutletTypeChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerOutletTypeChoices: valid values are %v", v, AllowedPowerOutletTypeChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerOutletTypeChoices) IsValid() bool {
	for _, existing := range AllowedPowerOutletTypeChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerOutletTypeChoices value
func (v PowerOutletTypeChoices) Ptr() *PowerOutletTypeChoices {
	return &v
}

type NullablePowerOutletTypeChoices struct {
	value *PowerOutletTypeChoices
	isSet bool
}

func (v NullablePowerOutletTypeChoices) Get() *PowerOutletTypeChoices {
	return v.value
}

func (v *NullablePowerOutletTypeChoices) Set(val *PowerOutletTypeChoices) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutletTypeChoices) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutletTypeChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutletTypeChoices(val *PowerOutletTypeChoices) *NullablePowerOutletTypeChoices {
	return &NullablePowerOutletTypeChoices{value: val, isSet: true}
}

func (v NullablePowerOutletTypeChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutletTypeChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

