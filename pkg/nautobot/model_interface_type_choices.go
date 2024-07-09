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

// InterfaceTypeChoices the model 'InterfaceTypeChoices'
type InterfaceTypeChoices string

// List of InterfaceTypeChoices
const (
	INTERFACETYPECHOICES_VIRTUAL InterfaceTypeChoices = "virtual"
	INTERFACETYPECHOICES_BRIDGE InterfaceTypeChoices = "bridge"
	INTERFACETYPECHOICES_LAG InterfaceTypeChoices = "lag"
	INTERFACETYPECHOICES__100BASE_FX InterfaceTypeChoices = "100base-fx"
	INTERFACETYPECHOICES__100BASE_LFX InterfaceTypeChoices = "100base-lfx"
	INTERFACETYPECHOICES__100BASE_TX InterfaceTypeChoices = "100base-tx"
	INTERFACETYPECHOICES__100BASE_T1 InterfaceTypeChoices = "100base-t1"
	INTERFACETYPECHOICES__1000BASE_T InterfaceTypeChoices = "1000base-t"
	INTERFACETYPECHOICES__2_5GBASE_T InterfaceTypeChoices = "2.5gbase-t"
	INTERFACETYPECHOICES__5GBASE_T InterfaceTypeChoices = "5gbase-t"
	INTERFACETYPECHOICES__10GBASE_T InterfaceTypeChoices = "10gbase-t"
	INTERFACETYPECHOICES__10GBASE_CX4 InterfaceTypeChoices = "10gbase-cx4"
	INTERFACETYPECHOICES__1000BASE_X_GBIC InterfaceTypeChoices = "1000base-x-gbic"
	INTERFACETYPECHOICES__1000BASE_X_SFP InterfaceTypeChoices = "1000base-x-sfp"
	INTERFACETYPECHOICES__10GBASE_X_SFPP InterfaceTypeChoices = "10gbase-x-sfpp"
	INTERFACETYPECHOICES__10GBASE_X_XFP InterfaceTypeChoices = "10gbase-x-xfp"
	INTERFACETYPECHOICES__10GBASE_X_XENPAK InterfaceTypeChoices = "10gbase-x-xenpak"
	INTERFACETYPECHOICES__10GBASE_X_X2 InterfaceTypeChoices = "10gbase-x-x2"
	INTERFACETYPECHOICES__25GBASE_X_SFP28 InterfaceTypeChoices = "25gbase-x-sfp28"
	INTERFACETYPECHOICES__50GBASE_X_SFP56 InterfaceTypeChoices = "50gbase-x-sfp56"
	INTERFACETYPECHOICES__40GBASE_X_QSFPP InterfaceTypeChoices = "40gbase-x-qsfpp"
	INTERFACETYPECHOICES__50GBASE_X_SFP28 InterfaceTypeChoices = "50gbase-x-sfp28"
	INTERFACETYPECHOICES__100GBASE_X_CFP InterfaceTypeChoices = "100gbase-x-cfp"
	INTERFACETYPECHOICES__100GBASE_X_CFP2 InterfaceTypeChoices = "100gbase-x-cfp2"
	INTERFACETYPECHOICES__200GBASE_X_CFP2 InterfaceTypeChoices = "200gbase-x-cfp2"
	INTERFACETYPECHOICES__400GBASE_X_CFP2 InterfaceTypeChoices = "400gbase-x-cfp2"
	INTERFACETYPECHOICES__100GBASE_X_CFP4 InterfaceTypeChoices = "100gbase-x-cfp4"
	INTERFACETYPECHOICES__100GBASE_X_CPAK InterfaceTypeChoices = "100gbase-x-cpak"
	INTERFACETYPECHOICES__100GBASE_X_QSFP28 InterfaceTypeChoices = "100gbase-x-qsfp28"
	INTERFACETYPECHOICES__100GBASE_X_CXP InterfaceTypeChoices = "100gbase-x-cxp"
	INTERFACETYPECHOICES__100GBASE_X_QSFPDD InterfaceTypeChoices = "100gbase-x-qsfpdd"
	INTERFACETYPECHOICES__100GBASE_X_DSFP InterfaceTypeChoices = "100gbase-x-dsfp"
	INTERFACETYPECHOICES__100GBASE_X_SFPDD InterfaceTypeChoices = "100gbase-x-sfpdd"
	INTERFACETYPECHOICES__200GBASE_X_QSFP56 InterfaceTypeChoices = "200gbase-x-qsfp56"
	INTERFACETYPECHOICES__200GBASE_X_QSFPDD InterfaceTypeChoices = "200gbase-x-qsfpdd"
	INTERFACETYPECHOICES__400GBASE_X_QSFP112 InterfaceTypeChoices = "400gbase-x-qsfp112"
	INTERFACETYPECHOICES__400GBASE_X_QSFPDD InterfaceTypeChoices = "400gbase-x-qsfpdd"
	INTERFACETYPECHOICES__400GBASE_X_OSFP InterfaceTypeChoices = "400gbase-x-osfp"
	INTERFACETYPECHOICES__400GBASE_X_OSFP_RHS InterfaceTypeChoices = "400gbase-x-osfp-rhs"
	INTERFACETYPECHOICES__400GBASE_X_CDFP InterfaceTypeChoices = "400gbase-x-cdfp"
	INTERFACETYPECHOICES__400GBASE_X_CFP8 InterfaceTypeChoices = "400gbase-x-cfp8"
	INTERFACETYPECHOICES__800GBASE_X_QSFPDD InterfaceTypeChoices = "800gbase-x-qsfpdd"
	INTERFACETYPECHOICES__800GBASE_X_OSFP InterfaceTypeChoices = "800gbase-x-osfp"
	INTERFACETYPECHOICES__1000BASE_KX InterfaceTypeChoices = "1000base-kx"
	INTERFACETYPECHOICES__10GBASE_KR InterfaceTypeChoices = "10gbase-kr"
	INTERFACETYPECHOICES__10GBASE_KX4 InterfaceTypeChoices = "10gbase-kx4"
	INTERFACETYPECHOICES__25GBASE_KR InterfaceTypeChoices = "25gbase-kr"
	INTERFACETYPECHOICES__40GBASE_KR4 InterfaceTypeChoices = "40gbase-kr4"
	INTERFACETYPECHOICES__50GBASE_KR InterfaceTypeChoices = "50gbase-kr"
	INTERFACETYPECHOICES__100GBASE_KP4 InterfaceTypeChoices = "100gbase-kp4"
	INTERFACETYPECHOICES__100GBASE_KR2 InterfaceTypeChoices = "100gbase-kr2"
	INTERFACETYPECHOICES__100GBASE_KR4 InterfaceTypeChoices = "100gbase-kr4"
	INTERFACETYPECHOICES_IEEE802_11A InterfaceTypeChoices = "ieee802.11a"
	INTERFACETYPECHOICES_IEEE802_11G InterfaceTypeChoices = "ieee802.11g"
	INTERFACETYPECHOICES_IEEE802_11N InterfaceTypeChoices = "ieee802.11n"
	INTERFACETYPECHOICES_IEEE802_11AC InterfaceTypeChoices = "ieee802.11ac"
	INTERFACETYPECHOICES_IEEE802_11AD InterfaceTypeChoices = "ieee802.11ad"
	INTERFACETYPECHOICES_IEEE802_11AX InterfaceTypeChoices = "ieee802.11ax"
	INTERFACETYPECHOICES_IEEE802_11AY InterfaceTypeChoices = "ieee802.11ay"
	INTERFACETYPECHOICES_IEEE802_15_1 InterfaceTypeChoices = "ieee802.15.1"
	INTERFACETYPECHOICES_OTHER_WIRELESS InterfaceTypeChoices = "other-wireless"
	INTERFACETYPECHOICES_GSM InterfaceTypeChoices = "gsm"
	INTERFACETYPECHOICES_CDMA InterfaceTypeChoices = "cdma"
	INTERFACETYPECHOICES_LTE InterfaceTypeChoices = "lte"
	INTERFACETYPECHOICES_SONET_OC3 InterfaceTypeChoices = "sonet-oc3"
	INTERFACETYPECHOICES_SONET_OC12 InterfaceTypeChoices = "sonet-oc12"
	INTERFACETYPECHOICES_SONET_OC48 InterfaceTypeChoices = "sonet-oc48"
	INTERFACETYPECHOICES_SONET_OC192 InterfaceTypeChoices = "sonet-oc192"
	INTERFACETYPECHOICES_SONET_OC768 InterfaceTypeChoices = "sonet-oc768"
	INTERFACETYPECHOICES_SONET_OC1920 InterfaceTypeChoices = "sonet-oc1920"
	INTERFACETYPECHOICES_SONET_OC3840 InterfaceTypeChoices = "sonet-oc3840"
	INTERFACETYPECHOICES__1GFC_SFP InterfaceTypeChoices = "1gfc-sfp"
	INTERFACETYPECHOICES__2GFC_SFP InterfaceTypeChoices = "2gfc-sfp"
	INTERFACETYPECHOICES__4GFC_SFP InterfaceTypeChoices = "4gfc-sfp"
	INTERFACETYPECHOICES__8GFC_SFPP InterfaceTypeChoices = "8gfc-sfpp"
	INTERFACETYPECHOICES__16GFC_SFPP InterfaceTypeChoices = "16gfc-sfpp"
	INTERFACETYPECHOICES__32GFC_SFP28 InterfaceTypeChoices = "32gfc-sfp28"
	INTERFACETYPECHOICES__32GFC_SFPP InterfaceTypeChoices = "32gfc-sfpp"
	INTERFACETYPECHOICES__64GFC_QSFPP InterfaceTypeChoices = "64gfc-qsfpp"
	INTERFACETYPECHOICES__64GFC_SFPDD InterfaceTypeChoices = "64gfc-sfpdd"
	INTERFACETYPECHOICES__64GFC_SFPP InterfaceTypeChoices = "64gfc-sfpp"
	INTERFACETYPECHOICES__128GFC_SFP28 InterfaceTypeChoices = "128gfc-sfp28"
	INTERFACETYPECHOICES_INFINIBAND_SDR InterfaceTypeChoices = "infiniband-sdr"
	INTERFACETYPECHOICES_INFINIBAND_DDR InterfaceTypeChoices = "infiniband-ddr"
	INTERFACETYPECHOICES_INFINIBAND_QDR InterfaceTypeChoices = "infiniband-qdr"
	INTERFACETYPECHOICES_INFINIBAND_FDR10 InterfaceTypeChoices = "infiniband-fdr10"
	INTERFACETYPECHOICES_INFINIBAND_FDR InterfaceTypeChoices = "infiniband-fdr"
	INTERFACETYPECHOICES_INFINIBAND_EDR InterfaceTypeChoices = "infiniband-edr"
	INTERFACETYPECHOICES_INFINIBAND_HDR InterfaceTypeChoices = "infiniband-hdr"
	INTERFACETYPECHOICES_INFINIBAND_NDR InterfaceTypeChoices = "infiniband-ndr"
	INTERFACETYPECHOICES_INFINIBAND_XDR InterfaceTypeChoices = "infiniband-xdr"
	INTERFACETYPECHOICES_T1 InterfaceTypeChoices = "t1"
	INTERFACETYPECHOICES_E1 InterfaceTypeChoices = "e1"
	INTERFACETYPECHOICES_T3 InterfaceTypeChoices = "t3"
	INTERFACETYPECHOICES_E3 InterfaceTypeChoices = "e3"
	INTERFACETYPECHOICES_DA15 InterfaceTypeChoices = "da15"
	INTERFACETYPECHOICES_DA26 InterfaceTypeChoices = "da26"
	INTERFACETYPECHOICES_DA31 InterfaceTypeChoices = "da31"
	INTERFACETYPECHOICES_DB25 InterfaceTypeChoices = "db25"
	INTERFACETYPECHOICES_DB44 InterfaceTypeChoices = "db44"
	INTERFACETYPECHOICES_DB60 InterfaceTypeChoices = "db60"
	INTERFACETYPECHOICES_DC37 InterfaceTypeChoices = "dc37"
	INTERFACETYPECHOICES_DC62 InterfaceTypeChoices = "dc62"
	INTERFACETYPECHOICES_DC79 InterfaceTypeChoices = "dc79"
	INTERFACETYPECHOICES_DD50 InterfaceTypeChoices = "dd50"
	INTERFACETYPECHOICES_DD78 InterfaceTypeChoices = "dd78"
	INTERFACETYPECHOICES_DD100 InterfaceTypeChoices = "dd100"
	INTERFACETYPECHOICES_DE9 InterfaceTypeChoices = "de9"
	INTERFACETYPECHOICES_DE15 InterfaceTypeChoices = "de15"
	INTERFACETYPECHOICES_DE19 InterfaceTypeChoices = "de19"
	INTERFACETYPECHOICES_DF104 InterfaceTypeChoices = "df104"
	INTERFACETYPECHOICES_XDSL InterfaceTypeChoices = "xdsl"
	INTERFACETYPECHOICES_DOCSIS InterfaceTypeChoices = "docsis"
	INTERFACETYPECHOICES_GPON InterfaceTypeChoices = "gpon"
	INTERFACETYPECHOICES_XG_PON InterfaceTypeChoices = "xg-pon"
	INTERFACETYPECHOICES_XGS_PON InterfaceTypeChoices = "xgs-pon"
	INTERFACETYPECHOICES_NG_PON2 InterfaceTypeChoices = "ng-pon2"
	INTERFACETYPECHOICES_EPON InterfaceTypeChoices = "epon"
	INTERFACETYPECHOICES__10G_EPON InterfaceTypeChoices = "10g-epon"
	INTERFACETYPECHOICES_CISCO_STACKWISE InterfaceTypeChoices = "cisco-stackwise"
	INTERFACETYPECHOICES_CISCO_STACKWISE_PLUS InterfaceTypeChoices = "cisco-stackwise-plus"
	INTERFACETYPECHOICES_CISCO_FLEXSTACK InterfaceTypeChoices = "cisco-flexstack"
	INTERFACETYPECHOICES_CISCO_FLEXSTACK_PLUS InterfaceTypeChoices = "cisco-flexstack-plus"
	INTERFACETYPECHOICES_CISCO_STACKWISE_80 InterfaceTypeChoices = "cisco-stackwise-80"
	INTERFACETYPECHOICES_CISCO_STACKWISE_160 InterfaceTypeChoices = "cisco-stackwise-160"
	INTERFACETYPECHOICES_CISCO_STACKWISE_320 InterfaceTypeChoices = "cisco-stackwise-320"
	INTERFACETYPECHOICES_CISCO_STACKWISE_480 InterfaceTypeChoices = "cisco-stackwise-480"
	INTERFACETYPECHOICES_CISCO_STACKWISE_1T InterfaceTypeChoices = "cisco-stackwise-1t"
	INTERFACETYPECHOICES_JUNIPER_VCP InterfaceTypeChoices = "juniper-vcp"
	INTERFACETYPECHOICES_EXTREME_SUMMITSTACK InterfaceTypeChoices = "extreme-summitstack"
	INTERFACETYPECHOICES_EXTREME_SUMMITSTACK_128 InterfaceTypeChoices = "extreme-summitstack-128"
	INTERFACETYPECHOICES_EXTREME_SUMMITSTACK_256 InterfaceTypeChoices = "extreme-summitstack-256"
	INTERFACETYPECHOICES_EXTREME_SUMMITSTACK_512 InterfaceTypeChoices = "extreme-summitstack-512"
	INTERFACETYPECHOICES_OTHER InterfaceTypeChoices = "other"
)

// All allowed values of InterfaceTypeChoices enum
var AllowedInterfaceTypeChoicesEnumValues = []InterfaceTypeChoices{
	"virtual",
	"bridge",
	"lag",
	"100base-fx",
	"100base-lfx",
	"100base-tx",
	"100base-t1",
	"1000base-t",
	"2.5gbase-t",
	"5gbase-t",
	"10gbase-t",
	"10gbase-cx4",
	"1000base-x-gbic",
	"1000base-x-sfp",
	"10gbase-x-sfpp",
	"10gbase-x-xfp",
	"10gbase-x-xenpak",
	"10gbase-x-x2",
	"25gbase-x-sfp28",
	"50gbase-x-sfp56",
	"40gbase-x-qsfpp",
	"50gbase-x-sfp28",
	"100gbase-x-cfp",
	"100gbase-x-cfp2",
	"200gbase-x-cfp2",
	"400gbase-x-cfp2",
	"100gbase-x-cfp4",
	"100gbase-x-cpak",
	"100gbase-x-qsfp28",
	"100gbase-x-cxp",
	"100gbase-x-qsfpdd",
	"100gbase-x-dsfp",
	"100gbase-x-sfpdd",
	"200gbase-x-qsfp56",
	"200gbase-x-qsfpdd",
	"400gbase-x-qsfp112",
	"400gbase-x-qsfpdd",
	"400gbase-x-osfp",
	"400gbase-x-osfp-rhs",
	"400gbase-x-cdfp",
	"400gbase-x-cfp8",
	"800gbase-x-qsfpdd",
	"800gbase-x-osfp",
	"1000base-kx",
	"10gbase-kr",
	"10gbase-kx4",
	"25gbase-kr",
	"40gbase-kr4",
	"50gbase-kr",
	"100gbase-kp4",
	"100gbase-kr2",
	"100gbase-kr4",
	"ieee802.11a",
	"ieee802.11g",
	"ieee802.11n",
	"ieee802.11ac",
	"ieee802.11ad",
	"ieee802.11ax",
	"ieee802.11ay",
	"ieee802.15.1",
	"other-wireless",
	"gsm",
	"cdma",
	"lte",
	"sonet-oc3",
	"sonet-oc12",
	"sonet-oc48",
	"sonet-oc192",
	"sonet-oc768",
	"sonet-oc1920",
	"sonet-oc3840",
	"1gfc-sfp",
	"2gfc-sfp",
	"4gfc-sfp",
	"8gfc-sfpp",
	"16gfc-sfpp",
	"32gfc-sfp28",
	"32gfc-sfpp",
	"64gfc-qsfpp",
	"64gfc-sfpdd",
	"64gfc-sfpp",
	"128gfc-sfp28",
	"infiniband-sdr",
	"infiniband-ddr",
	"infiniband-qdr",
	"infiniband-fdr10",
	"infiniband-fdr",
	"infiniband-edr",
	"infiniband-hdr",
	"infiniband-ndr",
	"infiniband-xdr",
	"t1",
	"e1",
	"t3",
	"e3",
	"da15",
	"da26",
	"da31",
	"db25",
	"db44",
	"db60",
	"dc37",
	"dc62",
	"dc79",
	"dd50",
	"dd78",
	"dd100",
	"de9",
	"de15",
	"de19",
	"df104",
	"xdsl",
	"docsis",
	"gpon",
	"xg-pon",
	"xgs-pon",
	"ng-pon2",
	"epon",
	"10g-epon",
	"cisco-stackwise",
	"cisco-stackwise-plus",
	"cisco-flexstack",
	"cisco-flexstack-plus",
	"cisco-stackwise-80",
	"cisco-stackwise-160",
	"cisco-stackwise-320",
	"cisco-stackwise-480",
	"cisco-stackwise-1t",
	"juniper-vcp",
	"extreme-summitstack",
	"extreme-summitstack-128",
	"extreme-summitstack-256",
	"extreme-summitstack-512",
	"other",
}

func (v *InterfaceTypeChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceTypeChoices(value)
	for _, existing := range AllowedInterfaceTypeChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceTypeChoices", value)
}

// NewInterfaceTypeChoicesFromValue returns a pointer to a valid InterfaceTypeChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceTypeChoicesFromValue(v string) (*InterfaceTypeChoices, error) {
	ev := InterfaceTypeChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceTypeChoices: valid values are %v", v, AllowedInterfaceTypeChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceTypeChoices) IsValid() bool {
	for _, existing := range AllowedInterfaceTypeChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InterfaceTypeChoices value
func (v InterfaceTypeChoices) Ptr() *InterfaceTypeChoices {
	return &v
}

type NullableInterfaceTypeChoices struct {
	value *InterfaceTypeChoices
	isSet bool
}

func (v NullableInterfaceTypeChoices) Get() *InterfaceTypeChoices {
	return v.value
}

func (v *NullableInterfaceTypeChoices) Set(val *InterfaceTypeChoices) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTypeChoices) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTypeChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTypeChoices(val *InterfaceTypeChoices) *NullableInterfaceTypeChoices {
	return &NullableInterfaceTypeChoices{value: val, isSet: true}
}

func (v NullableInterfaceTypeChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTypeChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
